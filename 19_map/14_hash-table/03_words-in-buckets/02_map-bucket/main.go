package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	// Create map with a key of int
	// and a value of another map
	// with a key of string, which will be the word
	// and a value of int, which will be the number of times the word occurs
	buckets := make(map[int]map[string]int)
	// Create slices to hold words words
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n][word]++
	}
	// Print words in a bucket
	for k, v := range buckets[6] {
		fmt.Println(v, "\t-", k)
	}
	fmt.Println(len(buckets))
	fmt.Println(len(buckets[6]))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets

}
