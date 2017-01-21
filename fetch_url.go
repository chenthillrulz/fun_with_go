package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"	
)

func main() {
	fmt.Println("Hello, playground")
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fectch %v \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", b)
}
