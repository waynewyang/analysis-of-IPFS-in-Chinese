## IPFS

### GO语言基础
- [x] [context用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/context)
- [x] [空struct用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/nullstruct)
- [x] [interface理解](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/interface)

### IPFS协议总览
- [x] [协议总览](https://github.com/ipfs/specs/tree/master/architecture)
- [x] [merkledag瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/mdag.waist.png) 与 [IP瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/ip.waist.png)
- [x] [国内研究参考]
	- [x] [mornmist team](https://github.com/mornmist/Newcomer-Guide)
	- [x] [mornmist ipfs](https://github.com/mornmist/IPFS-For-Chinese)

### IPFS API熟悉
- [x] [参考](https://ipfs.docs.apiary.io)
- [x] API笔记
	- [x] [basic](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/basic.md)
	- [x] [data structure](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/datastructure.md)
	- [x] [advance](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/adv.md)
	- [x] [network](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/net.md)
	- [x] [tool](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/tool.md)
	- [x] [others](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/others.md) 

### libp2p接口使用
- [ ] 转发使用
	- [x] [命令](https://gitlab.com/waynewyang/ipfs/blob/master/doc/network.md)
	- [ ] 编码

### 基础模块分析
- [x] [multiformats](https://gitlab.com/waynewyang/ipfs/blob/master/doc/multiformat.md)
	- [x] [mafmt](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/mafmt)
	- [x] [addr-util](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/addr-util)
	- [x] [maddr-filter](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/maddr-filter)
	- [x] [multiaddr-filter](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/multiaddr-filter)
- [x] [cid](https://gitlab.com/waynewyang/ipfs/blob/master/doc/cid.md)
- [ ] [datastores](https://gitlab.com/waynewyang/ipfs/blob/master/doc/datastores.md)
	- [ ] leveldb
	- [ ] go-ds-flatfs
- [ ] [protobuf](https://gitlab.com/waynewyang/ipfs/blob/master/protobuf)
	- [ ] io(gogo-protobuf/io)
- [ ] [semver](https://gitlab.com/waynewyang/ipfs/blob/master/semver)

### IPFS主流程源码概要分析
- [ ] ipfs init
- [ ] ipfs daemon
- [ ] ipfs add
- [ ] ipfs get
- [ ] cancel

### 协议层源码主逻辑分析
- [ ] network
	- [x] [pnet](https://gitlab.com/waynewyang/ipfs/blob/master/go-libp2p-interface-pnet)
	- [ ] relay
	- [ ] 穿透
- [ ] routing
	- [x] DHT寻址原理
	- [x] 源码中修改K桶 a并发与k值提高效率
- [ ] bitswap
- [ ] IPLD
- [ ] IPNS

### IPFS&libP2P协议对于我们来说存在的主要问题
- 只能指定单目录存储，如果FILECOIN不做处理，硬盘的换盘会比较麻烦
- libP2P穿透差

### 基于IPFS应用规划
- [ ] 先河网盘
- [ ] 共享圈
- [ ] 分布式鉴权云存储
- [ ] 分布式聊天
- [ ] VIPFS
