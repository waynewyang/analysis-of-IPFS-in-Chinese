# IPFS分析笔记——Wayne

> 作者：杨尉(waynewyang),转载请注明出处

> 此资源库编写于2018年9月-10月，开源于2019年5月。


![](/study.png) 
## 目录
- IPFS协议总览
	- [IPFS定义及参考技术](#ipfs定义及参考技术)
	- [协议总览](#协议总览)
- 基础
	- [语言基础](#语言基础)
	- [基础模块分析](#基础模块分析)
		- multiformat
		- cid
		- repo—看看IPFS持久化的数据是什么样的
- IPFS协议分析
	- [协议层分析](#协议层分析)
		- network, routing, exchange, merkledag, naming
	- [IPFS应用层数据结构](#应用层数据结构)
		- unixfs
	- [debug方法](#协议debug方法)
- IPFS API
	- [对外接口](#对外接口)
	- [IPFS主要API概要分析](#ipfs主流程概要分析)
		- init、daemon、add、get、cancel
- 应用
	- 官方应用
		- [IPFS-Cluster](#ipfs-cluster)
		- [ipfs-companion](#ipfs-companion)
		- [ipfs-desktop](#ipfs-desktop)
	- [第三方应用](#第三方应用)

### IPFS定义及参考技术
- 定义
	> 白皮书：IPFS是一种内容可寻址、版本化、点对点超媒体的分布式文件系统

	> 官网：IPFS是分布式Web,一种点对点超媒体协议，使网络更快，更安全，更开放。
	
- 参考技术
	- ice协议——网络连通性
	> 解释：解决peers之间的连通性

	- DHT技术，KAD、coral算法——分布式路由
	> 解释：解决peers之间的路由、内容路由（内容寻址）、以及分布式哈希key-value存储
	
	> 重要特征：内容寻址、或者说哈希寻址
	
	- bt——交换技术
	> 解释：解决peers之间的内容交换
	
	> 重要特征：P2P分布式、更快、更开放
	
	- git——数据结构
	> 解释：MerkleDAG借鉴git的数据存储方式，解决数据的防篡改、去重
	
	> 重要特征：防篡改、去重
	
	- sfs——命名系统
	> 解决哈希难以记忆、以及由于DAG特性导致的动态数据变更导致的更新成本问题
	
	> 重要特征：版本化（通过ipns的发布记录）、ipns解析（人类更容易识别操作）、dns解析

	- unixfs系统
	> 特征：unixfs挂载特性，挂载的ipfs类似一个文件系统，可以使用unix方式访问全球哈希。

```
第一步：在PEER A上，在IPFS添加another.txt文件
D:\go-ipfs>ipfs add another.txt
 29 B / 29 B [===================================================
added QmamQ2prV7FTfFm1eJc5o6QRA2BAkUJAbc5JCrLpu9dY2z another.txt

第二步：在peer B上，挂载ipfs文件系统（全球性的），直接通过unix的方式，cat第一步的哈希，可以查询到文件内容
wayne@wayne:~$ ipfs mount
IPFS mounted at: /ipfs
IPNS mounted at: /ipns
wayne@wayne:~/go/src$ cat /ipfs/QmamQ2prV7FTfFm1eJc5o6QRA2BAkUJAbc5JCrLpu9dY2z
"This is new another file" 
```

- [回到目录](#目录)

### 协议总览
- [x] [**IPFS项目进度**]
	- [pm记录](https://github.com/ipfs/pm)
	- ![](https://img.shields.io/badge/status-wip-orange.svg?style=flat-square) - 规范制定中，不成熟。
	- ![](https://img.shields.io/badge/status-draft-yellow.svg?style=flat-square) - 草稿完成，很有可能做较大变更。
	- ![](https://img.shields.io/badge/status-reliable-green.svg?style=flat-square) -接近稳定，只会做较小变更。
	- ![](https://img.shields.io/badge/status-stable-brightgreen.svg?style=flat-square) - 规范稳定，本质不变，可能做小的改善。
	- ![](https://img.shields.io/badge/status-permanent-blue.svg?style=flat-square) - 完全稳定，永久不变。
- [x] [**协议总览**](https://github.com/ipfs/specs/tree/master/architecture)
	- MerkleDAG vs IP 瘦腰图
		- ipfs通过MerkleDAG访问资源、http通过IP访问资源
		- 向下提供支撑，向上提供应用

![](/doc/image/mdag.waist.png) 

![](/doc/image/ip.waist.png)

- [回到目录](#目录)

### 语言基础
- [x] [go依赖包编译注意](/doc/compile.md)
- [x] [context用法](/go-basic/context)
- [x] [interface理解](/go-basic/interface)

- [回到目录](#目录)

### 基础模块分析
- [x] [**multiformats**](/doc/multiformat.md)
- [x] **基于multiaddr的一些支撑包（网络层）**
  - [x] [mafmt](/multiaddrs/mafmt)
  - [x] [addr-util](/multiaddrs/addr-util)
  - [x] [maddr-filter](/multiaddrs/maddr-filter)
  - [x] [multiaddr-filter](/multiaddrs/multiaddr-filter)
- [x] [**cid**](/doc/cid.md)
- [x] **repo**
	- [x] [参考 ](https://github.com/ipfs/specs/tree/master/repo) 
		- 文档未更新 
	- [x] [持久化数据分析](/datastores/README.md)
		- [x] [笔记](datastores/note.md)
		- [x] [示例](/datastores/example.md)
	- [ ] [细节源码分析](/datastores/源码分析.md)
- [ ] **其他**
	- [ ] [protobuf](/protobuf)
	- [ ] [semver](/semver)
	- [ ] [IPFS数据导入导出/DEX](https://github.com/ipfs/specs/tree/master/dex)

- [回到目录](#目录)

### 协议层分析
- [x] [**network**](/network)
	- [ ] [pnet](/go-libp2p-interface-pnet)
	- [x] [relay](/doc/network.md)
	- [ ] 穿透，待分析，目前端口转发OK，穿透比较差
- [x] [**routing**](/routing)
	- [x] [DHT寻址原理](/routing/dht.md)
	- [x] [源码中修改K桶 a并发与k值提高效率](/routing/修改并发属.md)
- [x] **block exchange**
	- [x] [block](/ipld/block)
	- [x] [笔记](/bitswap/bitswap)
- [x] **merkledag**
	- [x] DAG的相关的数据结构
	> cid -> block -> node(dag)

	> link -> node(dag)
	- [x] [对上层接口 ipld-format](/ipld/ipld-format)
	- [x] 具体实现分析笔记
		- [x] [ipld-cbor](/ipld/ipld-cbor)
		- [x] [go-merkledag](/ipld/go-merkledag)
	- [x] [dag-link数目大小限制，为什么是最多174个](/ipld/dag-link数目大小限制)
- [x] [**namesys**](/namesys)
	- [x] [dns](/namesys/dns) 
	- [x] [ipns](/namesys/ipns)

- [回到目录](#目录)

### 应用层数据结构
- [x] [unixfs](/unixfs)
	- [x] mount
		- [x] 挂载在unix文件系统之上进行操作
		- [x] [mount 示例](/unixfs/mount.md)
	- [x] [files(mfs)参考](https://github.com/ipfs/interface-ipfs-core/blob/master/SPEC/FILES.md#mutable-file-system)
    	- [x] [操作示例](/unixfs/files_op.md) 
	- [ ] 其他
	- [ ] 源码分析

- [回到目录](#目录)
	
### 协议debug方法
> ipfs log tail
- [回到目录](#目录)

### 对外接口
- [x] **参考**
	- [x] [参考](https://ipfs.docs.apiary.io)
	- [x] [public-api](https://github.com/ipfs/specs/tree/master/public-api)
	- [ ] [js  interface-ipfs-core](https://github.com/ipfs/interface-ipfs-core)
- [x] **API笔记**
	- [x] [basic](/doc/api/basic.md)
	- [x] [data structure](/doc/api/datastructure.md)
	- [ ] [advance](/doc/api/adv.md)
	- [x] [network](/doc/api/net.md)
	- [x] [tool](/doc/api/tool.md)
	- [ ] [others](/doc/api/others.md) 

- [回到目录](#目录)

### [IPFS主流程概要分析](/主要api流程概要分析)

- [回到目录](#目录)

### IPFS-Cluster
- [x] 参考
  - [x] [源码](https://github.com/ipfs/ipfs-cluster)
  - [x] [文档](https://cluster.ipfs.io/documentation/)
- [x] [笔记](/ipfs-cluster)
  - [x] [使用示例](/ipfs-cluster/example)
  - [ ] 源码分析
- [回到目录](#目录)

### ipfs-companion
- [网址](https://github.com/ipfs-shipyard/ipfs-companion)

- [回到目录](#目录)

### ipfs-desktop
- [网址](https://github.com/ipfs-shipyard/ipfs-desktop)

- [回到目录](#目录)

## 第三方应用
- [DTube - 分散式视频平台](https://d.tube/)
- [Decentraland的JanusVR查看器（来自ipfs的内容）](https://www.youtube.com/watch?v=841vXBopH68)
- [Textile Photos - 用于照片的数字钱包，端到端加密](https://www.textile.photos/)
- [Paratii - 分布式策展协议和流媒体引擎](Paratii - 分布式策展协议和流媒体引擎)
- [Peergos - 端到端加密，p2p文件存储和共享](https://peergos.org/)
- [更多收录的应用](https://awesome.ipfs.io/categories/apps/)

- 文章
	- [文章：纺织：将下一百万同行加入IPFS](https://medium.com/textileio/adding-the-next-million-peers-to-ipfs-76d356352d14)
	- [文章：未来开源的分散代码分发](https://medium.com/textileio/decentralized-code-distribution-for-the-future-of-open-source-2dc58f1153b2)
	- [文章：阿卡莎：修补差距](https://blog.akasha.world/2018/04/18/akasha-web-mend-the-gap/)
	- [Talk：真正无服务器，CRDT和IPFS（scalarconf）](https://www.youtube.com/watch?v=EscG2aytq10)
	- [会谈：DFS峰会上的IPFS闪电会谈](https://decentralizedweb.net/videos/talks-ipfs-lightning-talks/)
	- [教程：从头开始构建行星际ĐApp](https://medium.com/textileio/building-an-interplanetary-%C4%91app-from-scratch-51f9b8be5a74)
	- [教程：EC2的IPFS集群对等安装程序](https://medium.com/textileio/ipfs-cluster-peer-installer-for-ec2-ef2e4bfb1a74)
	- [教程：在Go中编写一个简单的P2P区块链](https://medium.com/@mycoralhealth/code-a-simple-p2p-blockchain-in-go-46662601f417)
	- [教程：从零到行星际英雄（基于浏览器的ĐApps与IPFS）](https://medium.freecodecamp.org/from-zero-to-interplanetary-hero-7e62f7d4427)

- [回到目录](#目录)
