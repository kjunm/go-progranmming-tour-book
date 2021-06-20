package main

import (
	"flag"
	"fmt"
	"log"
)

// 标准库的基本使用和长短选项
func main() {
	var name string
	flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go语言", "帮助信息")
	flag.Parse()
	fmt.Printf("%v", flag.CommandLine)

	log.Printf("name: %s", name)
}
