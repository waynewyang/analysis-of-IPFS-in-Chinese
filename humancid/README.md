## humancid
- CID to human readable cid
- ref [cid-utils.ipfs.team](http://cid-utils.ipfs.team/) 
## Usage
	humancid -i cid-string

## cid-v0 example
	wayne@wayne:~/humancid$ ./humancid -i QmYxFD7AAEgAHaYMzU52wjkL15HrNfiuNfVVaXzctf4iJd
	[cid version]: 0  
	[multibase  ]  
		   [code]: z  
		   [name]: base58btc  
	[multicodec ]  
		   [code]: 0x70  
		   [name]: protobuf 
	[mulitihash ]  
		   [code]: 0x12  
		   [name]: sha2-256  
		   [bits]: 256  
		   [value]: 9db334f4ed28fbd34066159c609ccfe153034f2425d776bc7bcc3b6c916c85ba  

## cid-v1 example
	wayne@wayne:~/humancid$ ./humancid -i zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA
	[cid version]: 1
	[multibase  ]
		   [code]: z
		   [name]: base58btc
	[multicodec ]
		   [code]: 0x55
		   [name]: raw
	[mulitihash ]
		   [code]: 0x12
		   [name]: sha2-256
		   [bits]: 256
		   [value]: 6e6ff7950a36187a801613426e858dce686cd7d7e3c0fc42ee0330072d245c95
