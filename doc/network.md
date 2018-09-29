## libp2p使用

### 参考的stun go源码
- [源码](https://github.com/ccding/go-stun)
- 可用的免费stun服务器
    1. stun.stunprotocol.org:3478

### 转发
- 透过转发进行连接
> 通过QmPyc6noLVKCp8Fx53rGLRLDXKNU7p9qN1njMW3tinEkTm节点relay到QmVf2XngNc4pnKCeJmTy5xyKjQdZDtAWGjbyGSR4eNhAr6

```
ipfs swarm connect /p2p-circuit/ipfs/QmPyc6noLVKCp8Fx53rGLRLDXKNU7p9qN1njMW3tinEkTm/ipfs/QmVf2XngNc4pnKCeJmTy5xyKjQdZDtAWGjbyGSR4eNhAr6
```
