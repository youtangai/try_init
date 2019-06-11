# init関数の順番調査
複数パッケージをインポートする際，そのパッケージ間で依存関係が存在する場合，どのように処理されるのか調査を行った．

# 方法
方法1
1. パッケージcで定義されたexported変数Cにcのinit関数で10秒待ってから値を代入する．
2. パッケージbで定義されたexported変数Bにbのinit関数でcパッケージのCの値を代入する
3. パッケージaで定義されたexported変数Aにaのinit関数でbパッケージのBの値を代入する
4. mainパッケージのinit関数でパッケージaのAを取得し，標準出力に表示する
  
また，これの逆順として方法2も行った．

# 環境(go envの出力内容)
```
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/nagaiyouta/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/nagaiyouta/go"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/Users/nagaiyouta/work/Go/try_init/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/41/qpfvv7zn2sgbxrp9k017w0xw0000gn/T/go-build714259205=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

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