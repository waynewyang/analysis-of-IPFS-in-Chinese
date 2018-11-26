go交叉编译arm上的程序

1.ubuntu系统中安装go

sudo apt-get install golang

2.交叉编译

$GOOS	    $GOARCH

android     arm

darwin      386

darwin      amd64

darwin      arm

darwin      arm64

dragonfly   amd64

freebsd     386

freebsd     amd64

freebsd     arm

linux       386

linux       amd64

linux       arm

linux       arm64

linux       ppc64

linux       ppc64le

linux       mips

linux       mipsle

linux       mips64

linux       mips64le

netbsd      386

netbsd      amd64

netbsd      arm

openbsd     386

openbsd     amd64

openbsd     arm

plan9       386

plan9       amd64

solaris     amd64

windows     386

windows     amd64

#编译一般go程序

GOARM=7 GOARCH=arm GOOS=linux go build -o outfile

#编译内嵌c代码的go程序

CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc GOARM=7 GOARCH=arm GOOS=linux go build -o outfile

ipfs编译：

GOARM=7 GOARCH=arm

然后make install  屏蔽profile中的GOBIN