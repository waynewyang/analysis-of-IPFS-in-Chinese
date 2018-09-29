package main

import (
	"flag"
	"fmt"
	ld "gx/ipfs/QmbBhyDKsY4mbY6xsKt3qu9Y7FPvMJ6qbD8AMjYYvPRw1g/goleveldb/leveldb"
	"io/ioutil"
	"os"
)

// flags
var input_dir string
var output_file string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s -i dir -o file \n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&input_dir, "i", "", "input leveldb dir ")
	flag.StringVar(&output_file, "o", "", "output file")
}

func main() {
	var item string

	flag.Parse()

	if input_dir == "" || output_file == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	db, err := ld.OpenFile(input_dir, nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	writeWithIoutil(output_file, "")

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		item = fmt.Sprintf("key:%s,value:%s\n", iter.Key(), iter.Value())
		appendToFile(output_file, item)
	}
	iter.Release()
}

func writeWithIoutil(name string, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) != nil {
		os.Exit(1)
	}
}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
