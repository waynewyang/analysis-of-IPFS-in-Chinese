# namesys
## 目录
- [必要性](#必要性)
- [目的](#目的)
- [使用简介](#使用简介)
- 原理分析
	- [DNS原理分析](/namesys/dns)
	- [IPNS原理分析](/namesys/ipns)

##  必要性
- MerkleDAG的特性
> MerkleDAG作为IPFS的核心数据结构，使得使用内容hash进行寻址访问成为可能，任一叶子节点hash的变更，均会导致MerkleDAG的root hash的变化，确保了内容的可靠性。

> 优点：防篡改、去重、内容寻址

> 问题：
> 1. 哈希难以被人类记忆，想象一下我们的网站服务是一串串的哈希，非常头疼
>	访问方式：/ipns/QmabfJUjD6Gf3KBFBvjiR9XRPp8hBG6zDKDSfjVmRYfU7P or /ipfs/QmabfJUjD6Gf3KBFBvjiR9XRPp8hBG6zDKDSfjVmRYfU7P
> 2. 每次修改，均会导致哈希值的改变，这对动态数据存储非常不方便
> [修改导致哈希值变化示例](#修改导致哈希值变化示例)

- [回到目录](#目录)

## 目的
- 提供所指向对象可变的指针
>IPNS基于[SFS](http://en.wikipedia.org/wiki/Self-certifying_File_System)。它是PKI命名空间 - 名称只是公钥的哈希。控制私钥的人控制名称。记录由私钥签名并分布在任何地方（在IPFS中，通过路由系统）。这是一种在互联网上分配可变名称的平等主义方式，没有任何集中化或证书颁发机构。
	- 发布
	- 解析

- 提供人类可读的方式访问IPFS资源
> 通用域名的DNS解析

- [回到目录](#目录)

## 使用简介

### IPNS发布及访问
- ipfs name
> 将内容发布到ipns，并通过ipns的方式查看

```
wayne@wayne:~/ipnstest$ ipfs key gen --type=rsa --size=2048 arsyuntestkey
QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF
	
wayne@wayne:~/ipnstest$ ipfs key list -l
QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9 self          
QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF arsyuntestkey 
QmcSctrtBydCtyQRxcG43jSD6VLvBdANUzFbzjss3TfuF2 mykey  
	
wayne@wayne:~/ipnstest$ ll ~/.ipfs/keystore/
total 16
drwx------ 2 wayne wayne 4096 Oct 19 11:17 ./
drwxrwxr-x 5 wayne wayne 4096 Oct 18 16:46 ../
-rw-rw-r-- 1 wayne wayne 1197 Oct 19 11:17 arsyuntestkey
-rw-rw-r-- 1 wayne wayne 1196 Oct 19 11:12 mykey
wayne@wayne:~/ipnstest$ 
wayne@wayne:~/ipnstest$ echo "hello,arsyun,wayne" > test
wayne@wayne:~/ipnstest$ ipfs add test
added QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX test
19 B / 19 B [========================================================================================] 100.00%
	      
wayne@wayne:~/ipnstest$ ipfs name publish --key=arsyuntestkey /ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
Published to QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF: /ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
	
wayne@wayne:~/ipnstest$ ipfs cat  /ipns/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF
hello,arsyun,wayne
wayne@wayne:~/ipnstest$ ipfs cat /ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
hello,arsyun,wayne
```

### IPNS解析
- ipfs resolve
> ipns解析为ipfs地址

### DNS解析
- ipfs dns
> 通用域名，解析为ipfs地址

- [回到目录](#目录)

## 修改导致哈希值变化示例
```
wayne@wayne:~/test$ mkdir ipnstest
wayne@wayne:~/test$ echo hello > ipnstest/a
wayne@wayne:~/test$ echo arsyun > ipnstest/b
wayne@wayne:~/test$ ipfs add ipnstest -r
added QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN ipnstest/a
added QmUQ2Uug69QWK3z7SKnF4fBv7iGoe5FCgBqaUnz5SdaFtX ipnstest/b
added Qmf7KcuWGeqFqGU2twMq1VAvshUeB1PL8bSGE3kxjYugwx ipnstest
 13 B / 13 B [=========================================================================================] 100.00%wayne@wayne:~/test$ ipfs object get Qmf7KcuWGeqFqGU2twMq1VAvshUeB1PL8bSGE3kxjYugwx
{"Links":[{"Name":"a","Hash":"QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN","Size":14},{"Name":"b","Hash":"QmUQ2Uug69QWK3z7SKnF4fBv7iGoe5FCgBqaUnz5SdaFtX","Size":15}],"Data":"\u0008\u0001"}
wayne@wayne:~/test$ echo ars > ipnstest/b
wayne@wayne:~/test$ ipfs add ipnstest -r
added QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN ipnstest/a
added QmaeuJm6yoXFUMUMQMWs1hodPFsyMkwPGhtWangihTSruk ipnstest/b
added QmabfJUjD6Gf3KBFBvjiR9XRPp8hBG6zDKDSfjVmRYfU7P ipnstest
 10 B / 10 B [=========================================================================================] 100.00%wayne@wayne:~/test$ ipfs object get QmabfJUjD6Gf3KBFBvjiR9XRPp8hBG6zDKDSfjVmRYfU7P
{"Links":[{"Name":"a","Hash":"QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN","Size":14},{"Name":"b","Hash":"QmaeuJm6yoXFUMUMQMWs1hodPFsyMkwPGhtWangihTSruk","Size":12}],"Data":"\u0008\u0001"}
```

- [回到必要性](#必要性)

- [回到目录](#目录)

