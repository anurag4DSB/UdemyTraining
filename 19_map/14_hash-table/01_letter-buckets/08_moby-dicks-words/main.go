package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	//scan the page
	//New Scanner takes a reader and res.Body implements the reader interface(so it is a reader)
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
