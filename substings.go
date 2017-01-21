package main

import (
	"os"
	"fmt"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./substrings <string>")
		return
	}

	str := os.Args[1]
	length := len(str)
	dict := make(map[string] bool)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			new_str := str[i:j+1]
			dict[new_str] = true
		}
	}

	final := ""
	for k, _ := range dict {
		fmt.Println(k)
		if k > final {
			final = k
		}
	}

	if final != "" {
		fmt.Println("The largest string is ", final)
	}
}
