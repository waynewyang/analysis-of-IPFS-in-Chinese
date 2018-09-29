## IPFS

### GO语言基础
- [x] [context用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/context)
- [x] [空struct用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/nullstruct)
- [x] [interface理解](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/interface)

### IPFS协议总览及计划
- [x] [协议总览](https://github.com/ipfs/specs/tree/master/architecture)
- [x] [merkledag瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/mdag.waist.png) 与 [IP瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/ip.waist.png)
- [ ] 总结文档

### IPFS API熟悉
- [x] IPFS API

### 基础模块分析
- [x] [multiformats](https://gitlab.com/waynewyang/ipfs/blob/master/doc/multiformat.md)
- [x] [cid](https://gitlab.com/waynewyang/ipfs/blob/master/doc/cid.md)
- [ ] [datastores](https://gitlab.com/waynewyang/ipfs/blob/master/doc/datastores.md)
	- [ ] leveldb
	- [ ] go-ds-flatfs

### IPFS主要流程分析
- [ ] ipfs init
- [ ] ipfs daemon
- [ ] ipfs daemon
- [ ] ipfs add
- [ ] cancel

### 协议层源码主逻辑分析
- [ ] network
	- [ ] Relay
	- [ ] 穿透
- [ ] routing
	- [x] DHT寻址原理
	- [x] 源码中修改K桶 a并发与k值提高效率
	- [ ] 细节分析
- [ ] bitswap
- [ ] IPLD
- [ ] IPNS

### libp2p使用
- [ ] 转发使用
	- [x] 命令
	- [ ] 编码
- [ ] 待规划

### 基于IPFS&libP2P协议完善度
- [ ] 指定单目录存储，FILECOIN提前考虑
- [ ] libP2P穿透差

### 基于IPFS应用规划
- [ ] 先河网盘
- [ ] 共享圈
- [ ] 分布式鉴权云存储
- [ ] 分布式聊天
- [ ] VIPFS