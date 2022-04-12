package main

import (
	"fmt"
	"ucc4/package_a"
	"ucc4/package_b"
)

func main() {
	fmt.Println(package_a.Hello())
	fmt.Println(package_b.Hello())
}
