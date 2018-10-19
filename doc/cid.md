# CID
- 参考网址[github](https://github.com/ipld/cid)
## 概念
- 自描述的
- ipfs id 与 dag对象均为cid格式
- 用于内容寻址

## CIDV0
- 以 Qm开头
- 默认multibase：base58btc
- 默认cid-version:0
- 默认multicodec：sha2-256
- 默认hash长度：256bit

## CIDV1
>格式:`<multibase-prefix><cid-version><multicodec-content-type><multihash-content-address>`

## 解析示例
- [解析成人类可识别的cid示例](/../humancid)
