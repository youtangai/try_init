# init関数の順番調査
複数パッケージをインポートする際，そのパッケージ間で依存関係が存在する場合，どのように処理されるのか調査を行った．

# 方法
方法1
1. パッケージcで定義されたexported変数Cにcのinit関数で値を代入する．
2. パッケージbで定義されたexported変数Bにbのinit関数でcパッケージのCの値を代入する
3. パッケージaで定義されたexported変数Aにaのinit関数でbパッケージのBの値を代入する
4. mainパッケージのinit関数でパッケージaのAを取得し，標準出力に表示する
  
また，これの逆順として方法2も行った．

# 結果
依存順にパッケージのinit関数が呼ばれることを確認した  
方法1
```
c init hello
b init hello
a init hello
main init hello
main main hello
```
  
方法2  
```
a init hello
b init hello
c init hello
main init hello
main main hello
```