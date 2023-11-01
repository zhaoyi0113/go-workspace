package main

import (
	"fmt"

	"github.com/zhaoyi0113/go-workspace/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
