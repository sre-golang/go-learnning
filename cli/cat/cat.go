package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type CatArgs struct {
	FilePath string
}

func getAllArgs(args *CatArgs) {
	flag.StringVar(&args.FilePath, "f", "", "attachment file path")

	flag.Parse()
}

//读取附件内容
func ReadFile(filePath string) ([]byte, error) {
	if filePath == EMPTY_STRING {
		return nil, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return nil, err
	}

	att, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read file content failed:%v\n", err)
		return nil, err
	}

	return att, nil
}

func main() {
	var args *CatArgs = &CatArgs{}
	getAllArgs(args)

	content, err := ReadFile(args.FilePath)
	if err != nil {
		fmt.Printf("get file content faield:%v\n", err)
		return
	}

	fmt.Println(string(content))
}
