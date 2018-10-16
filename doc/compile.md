## go源码依赖包编译注意
- golang.org依赖包下载问题
    当执行go get下载源码依赖包的时候，依赖golang.org的包由于被墙掉，无法直接下载,会出现如下错误
    https fetch: Get https://golang.org/x/sys/cpu?go-get=1: dial tcp 216.239.37.1:443: connect: connection refused

- 解决方法,从github的go源码包中去下载
    比这个资源下载不了：https://golang.org/x/sys/cpu?go-get=1:
    - 进入你的GOPATH目录
    - mkdir golang.org/x -p;cd golang.org/x
    - git clone https://github.com/golang/sys.git

    - 其他包类同，git clone https://github.com/golang/XXXXXX.git
        其中xxxxxx表示需要下载的包
