## ADVANCED COMMANDS
### 未分析
- [ ] ipfs name pubsub
- [ ] ipfs p2p
- [ ] filestore

### daemon        
> Start a long-running daemon process

### mount         
> Mount an IPFS read-only mountpoint
>挂载到本地操作系统，挂载之后可以做os下面的文件操作

### resolve       
> Resolve any type of name

### name          
>Publish and resolve IPNS names

- ipfs name publish
- ipfs name pubsub?
- ipfs name pubsub cancel?
- ipfs name pubsub state?
- ipfs name pubsub subs?
- ipfs name resolve
- 解析ipns名称

### key           
> Create and list IPNS name keypairs
> 用于生成ipns命名空间

- ipfs key gen
- ipfs key list
- ipfs key rename
- ipfs key rm

### dns           
> Resolve DNS links


### pin           
>Pin objects to local storage

- ipfs pin add
- ipfs pin ls
- ipfs pin rm
> 删除本地资源

- ipfs pin update 
> 更新新的，删除旧的

- ipfs pin verify

### repo          
> Manipulate the IPFS repository

- ipfs repo
- ipfs repo fsck
- ipfs repo gc
>执行手动垃圾回收

- ipfs repo stat  
>repo状态

- ipfs repo verify  
>校验资源合法性

- ipfs repo version 
>repo版本展示

### stats         
> Various operational stats

- ipfs stats
- ipfs stats bitswap
- ipfs stats bw：带宽状态
- ipfs stats repo

### p2p? 
> Libp2p stream mounting

- ipfs p2p listener
- ipfs p2p listener close
- ipfs p2p listener ls
- ipfs p2p listener open
- ipfs p2p stream
- ipfs p2p stream close
- ipfs p2p stream dial
- ipfs p2p stream ls

### filestore     
> Manage the filestore (experimental)

- ipfs filestore
- ipfs filestore dups
- ipfs filestore ls
- ipfs filestore verify
