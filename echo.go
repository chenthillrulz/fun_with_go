// Echo prints command line paramters

package main

import (
	"fmt"
	"os"
	//"strconv"
)

func main () {
	fmt.Println(os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		//fmt.Println (strconv.Itoa(i) + " " + os.Args[i])
		s := fmt.Sprintf ("%d %s", i, os.Args[i])
		fmt.Println (s)
	}
}
