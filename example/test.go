package main

import (
	"fmt"

	"github.com/disiqueira/tindergo"
)

func main() {

	t := tindergo.New()

	err := t.Authenticate(token)
	if err != nil {
		panic(err)
	}

	recs, err := t.RecsSocial()
	if err != nil {
		panic(err)
	}

	fmt.Println(recs)
	/*
		for _, r := range recs.Results {
			fmt.Println(r.)
		}*/
}
