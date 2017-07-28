package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing url.")
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("HTML:\n\n", string(bytes))
	}
}
