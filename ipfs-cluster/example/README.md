# 示例

## 建立集群(多个节点做类似操作)
- 启动IPFS daemon
    ipfs daemon
- 初始化cluster配置
    ipfs-cluster-service init
    其中cluster配置中secret值，集群中每个节点需要一样

```
  "cluster": {
    "id": "QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB",
    "peername": "wayne",
    "secret": "52117401c06e416b5a69911e7596d8d0537b9d62cabc2bbe22b39a8733596444",
    ...
    }
    - 启动cluster进程
        直接启动，或者指定链接具体的节点
        ipfs-cluster-service daemon 
        ipfs-cluster-service daemon --bootstrap /ip4/192.168.57.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
```

## 查看集群节点

```
wayne@wayne:~$ ipfs-cluster-ctl peers ls
QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB | wayneyang | Sees 2 other peers
  > Addresses:
    - /ip4/127.0.0.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.158.124/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.161.166/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.177.175/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.2.183/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.57.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.84.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
  > IPFS: QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/127.0.0.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.158.124/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.161.166/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.177.175/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.2.183/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.57.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.84.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001:0:9d38:953c:10da:2d22:8be1:2302/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::587f:a18c:8502:35c4/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::91e/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::f1c1:a2ca:4dde:b9be/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/::1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB | wayne | Sees 2 other peers
  > Addresses:
    - /ip4/127.0.0.1/tcp/9096/ipfs/QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB
    - /ip4/192.168.57.129/tcp/9096/ipfs/QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB
  > IPFS: QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip4/127.0.0.1/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip4/192.168.57.129/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip6/::1/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS | waynepc | Sees 2 other peers
  > Addresses:
    - /ip4/127.0.0.1/tcp/9096/ipfs/QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS
    - /ip4/169.254.98.114/tcp/9096/ipfs/QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS
    - /ip4/192.168.10.200/tcp/9096/ipfs/QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS
    - /ip4/192.168.2.210/tcp/9096/ipfs/QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS
  > IPFS: QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip4/127.0.0.1/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip4/169.254.98.114/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip4/192.168.10.200/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip4/192.168.2.210/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip6/2001:0:9d38:90d7:103b:2cf0:8be1:2302/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip6/2001::49a/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip6/2001::9c22:ed62:5982:89a2/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip6/2001::e199:b497:7c55:a1df/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
    - /ip6/::1/tcp/4001/ipfs/QmWDesnDLPTcGVL1svU4MCeKSV1vUwd7WjCGxJdqJsQUEL
```
    
## 删除节点

```
wayne@wayne:~$ ipfs-cluster-ctl peers rm QmbcEZ4pevmZoNd3cefmfPDEV6JNRkrHHnnFKwu4FQC8SS
wayne@wayne:~$ ipfs-cluster-ctl peers ls
QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB | wayneyang | Sees 1 other peers
  > Addresses:
    - /ip4/127.0.0.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.158.124/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.161.166/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/169.254.177.175/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.2.183/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.57.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
    - /ip4/192.168.84.1/tcp/9096/ipfs/QmVFgbaA1Y3wf7KMDKENRwVNrhRYP8h3Q2e58VDUCP5TvB
  > IPFS: QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/127.0.0.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.158.124/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.161.166/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/169.254.177.175/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.2.183/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.57.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip4/192.168.84.1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001:0:9d38:953c:10da:2d22:8be1:2302/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::587f:a18c:8502:35c4/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::91e/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/2001::f1c1:a2ca:4dde:b9be/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
    - /ip6/::1/tcp/4001/ipfs/QmY4NwwDbhfaET7oWvDWi4YExo6HRdowRRYayeR7fJTDJW
QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB | wayne | Sees 1 other peers
  > Addresses:
    - /ip4/127.0.0.1/tcp/9096/ipfs/QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB
    - /ip4/192.168.57.129/tcp/9096/ipfs/QmVGWa3hZaTjpwpUeATxUs68s7QKDyWPxFRwaoU39zZQkB
  > IPFS: QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip4/127.0.0.1/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip4/192.168.57.129/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
    - /ip6/::1/tcp/4001/ipfs/QmeaqaXGKUn9X9XvjrJwHsVKVks32oxdU2w1VSULih4hy9
```

## 集群添加文件&删除（彻底删除需要清除缓存 ipfs repo gc）

```
wayne@wayne:~$ ipfs-cluster-ctl  add testarsyun
added QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP testarsyun
wayne@wayne:~$ ipfs-cluster-ctl  status QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : PINNED | 2018-10-19T09:05:33Z
    > wayne           : PINNED | 2018-10-19T09:05:52Z
    > waynepc         : PINNED | 2018-10-19T09:05:53Z
wayne@wayne:~$ ipfs-cluster-ctl  pin rm QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : UNPINNED | 2018-10-19T09:06:09Z
    > wayne           : UNPINNED | 2018-10-19T09:06:28Z
    > waynepc         : UNPINNED | 2018-10-19T09:06:29Z
wayne@wayne:~$ ipfs-cluster-ctl  status QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : UNPINNED | 2018-10-19T09:06:13Z
    > wayne           : UNPINNED | 2018-10-19T09:06:32Z
    > waynepc         : UNPINNED | 2018-10-19T09:06:33Z
```

## 重新添加文件（有一台机器不在线的时候，上线后删除，因为pin操作，有queue队列记录）

```
wayne@wayne:~$ ipfs-cluster-ctl pin  add QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP    
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : PINNED | 2018-10-19T09:12:34Z
    > wayne           : PINNED | 2018-10-19T09:12:53Z
    > waynepc         : PINNED | 2018-10-19T09:12:54Z
wayne@wayne:~$ ipfs-cluster-ctl  pin rm QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP     
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : UNPINNED | 2018-10-19T09:12:48Z
    > wayne           : UNPINNED | 2018-10-19T09:13:06Z
    > <peer.ID bcEZ4p> : CLUSTER_ERROR: dial attempt failed: <peer.ID VGWa3h> --> <peer.ID bcEZ4p> dial attempt failed: context deadline exceeded | 2018-10-19T09:13:08Z
wayne@wayne:~$ ipfs-cluster-ctl  status QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP 
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : UNPINNED | 2018-10-19T09:12:55Z
    > wayne           : UNPINNED | 2018-10-19T09:13:14Z
    > <peer.ID bcEZ4p> : CLUSTER_ERROR: dial attempt failed: <peer.ID VGWa3h> --> <peer.ID bcEZ4p> dial attempt failed: context deadline exceeded | 2018-10-19T09:13:19Z
wayne@wayne:~$ ipfs-cluster-ctl  status QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP
QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP :
    > wayneyang       : UNPINNED | 2018-10-19T09:13:21Z
    > wayne           : UNPINNED | 2018-10-19T09:13:40Z
    > waynepc         : UNPINNED | 2018-10-19T09:13:41Z
```