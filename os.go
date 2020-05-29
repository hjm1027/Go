package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	str := ""
	if len(os.Args) > 1{
		str += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning",str)
}
