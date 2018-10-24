# 主要API概要流程分析

## 目录
- [ipfs init](#init)
- [ipfs daemon](#daemon)
- [ipfs add](#add)
- [ipfs get](#get)
- [cancel](#cancel)

## init
go-ipfs/cmd/ipfs/init.go

- 解析json配置文件到内存中，解析到conf中
![](/主要api流程概要分析/1.png)

- 执行doInit函数进行初始化
![](/主要api流程概要分析/2.png)
![](/主要api流程概要分析/3.png)

- 初始化配置到内存
![](/主要api流程概要分析/4.png)

- 初始化资源目录
![](/主要api流程概要分析/5.png)

- 初始化pns空间(ipns名称是本节点id，且绑定到一个空目录)
![](/主要api流程概要分析/6.png)

[回到目录](#目录)

## daemon
go-ipfs/cmd/ipfs/daemon.go :func daemonFunc
- 解析传参、初始化配置
- 启动网络侦听服务
	- swarm服务 默认4001
	- api服务 默认5001
	- 网关服务  默认8080
- 完成启动

[回到目录](#目录)

## add
go-ipfs/core/commands/add.go->go-ipfs/core/coreunix/add.go

- 分片-> 调用dag层cache 到内存-> 如果有pin选项，持久化到本地（**分片方式是可选的**）

- DAG 调用交换层，广播宣告provide信息，存储到内容dht表中。

[回到目录](#目录)

## get
go-ipfs/core/commands/get.go
> 调用DAG层

> 本地有 -> 结束

> 本地无 -> 由dag服务调用交换、路由、网络获取对应的内容

[回到目录](#目录)

## cancel
> 通过context进行线程控制，资源回收

[回到目录](#目录)
