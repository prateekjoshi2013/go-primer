package main

import (
	"fmt"
	"log"
	"strconv"

	v4 "github.com/gofrs/uuid"
	v3 "github.com/gofrs/uuid/v3"
)

func main() {
	id1, _ := v4.NewV4()
	fmt.Println(id1)
	id2, _ := v3.NewV4()
	fmt.Println(id2)
	i, err := strconv.ParseInt("2A", 16, 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("0x2A: ", i)

	i, err = strconv.ParseInt("0x2A", 0, 64)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("0x2A: ", i)
}
