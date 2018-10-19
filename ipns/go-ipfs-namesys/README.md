# namesys

## 对应命令，功能介绍

### 发布

- ipfs name
    发布到ipns  


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
	      
	wayne@wayne:~/ipnstest$ ipfs name publish --key=arsyuntestkey 			/ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
	Published to QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF: /ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
	
	wayne@wayne:~/ipnstest$ ipfs cat  /ipns/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF
	hello,arsyun,wayne
	wayne@wayne:~/ipnstest$ ipfs cat /ipfs/QmfQ3GSZz9WVVUhtu7bHy7gv2ixAodFc3KSVs4mNkHfCBX
	hello,arsyun,wayne

## 解析
- ipfs dns
    域名解析为ipfs地址
- ipfs resolve
    ipns解析为ipfs地址
