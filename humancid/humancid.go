package main

import (
	"flag"
	"fmt"
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	mb "gx/ipfs/QmSbvata2WqNkqGtZNg8MR3SKwnB8iQ7vTPJgWqB8bC5kR/go-multibase"
	c "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
	"os"
)

func main() {
	cidstring := flag.String("i", "", "input cid")
	flag.Parse()
	if *cidstring == "" {
		fmt.Println("Please Input CID!")
		os.Exit(0)
	}
	base, err := c.ExtractEncoding(*cidstring)
	if err != nil {
		fmt.Println("Unknow cid format!")
		os.Exit(0)
	}

	cid, err := c.Decode(*cidstring)
	if err != nil {
		fmt.Println("Unknow cid format!")
		os.Exit(0)
	}

	prefix := cid.Prefix()

	multihash := cid.Hash()
	multihashString := multihash.HexString()
	hashValue := string([]byte(multihashString)[4:])

	fmt.Println("[cid version]:", prefix.Version)
	fmt.Println("[multibase  ]")
	fmt.Printf("       [code]: %c\n", base)
	fmt.Println("       [name]:", mb.EncodingToStr[base])
	fmt.Println("[multicodec ]")
	fmt.Printf("       [code]: 0x%x\n", prefix.Codec)
	fmt.Println("       [name]:", c.CodecToStr[prefix.Codec])
	fmt.Println("[mulitihash ]")
	fmt.Printf("       [code]: 0x%x\n", prefix.MhType)
	fmt.Println("       [name]:", mh.Codes[prefix.MhType])
	fmt.Println("       [bits]:", 8*prefix.MhLength)
	fmt.Println("      [value]:", hashValue)
}
