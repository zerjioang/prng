package main

import (
	"fmt"
	"github.com/zerjioang/prng"
)

func main() {
	uuid := prng.New()
	fmt.Println(uuid)
}
