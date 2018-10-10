package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	norExported int
}

func main() {
	p1 := person{"James", "Bond", 23, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
