package main

import (
	"flag"
	"fmt"
	ds "gx/ipfs/QmVG5gxteQNEMhrS8prJSmU2C9rebtFuTd3SYZ5kE3YZ5k/go-datastore"
	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
	dshelp "gx/ipfs/Qmd39D2vUhmPKQA2fgykjo2JXwekHKeJUggmGRpYuVMA2Z/go-ipfs-ds-help"
	"os"
)

func main() {
	cidstring := flag.String("c", "", "input cid")
	dskeystring := flag.String("d", "", "input dskey")
	flag.Parse()
	if *cidstring == "" && *dskeystring == "" {
		fmt.Println("Please Input cid or dskey!")
		os.Exit(0)
	}
	if *cidstring != "" && *dskeystring != "" {
		fmt.Println("you can input cid or dskey,only input one type!")
		os.Exit(0)
	}

	if *cidstring != "" {
		c, err := cid.Decode(*cidstring)
		dsKey := dshelp.CidToDsKey(c)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(dsKey)
		}
	} else if *dskeystring != "" {
		c, err := dshelp.DsKeyToCid(ds.NewKey(*dskeystring))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(c)
		}
	}
}
