package main

import (
	"fmt"
	"os"

	"github.com/disiqueira/tindergo"
)

func main() {

	t := tindergo.New()

	err := t.Authenticate(token)
	if err != nil {
		panic(err)
	}

	recs, err := t.RecsCore()
	if err != nil {
		panic(err)
	}

	for _, r := range recs {
		u, _ := t.User(r.ID)
		fmt.Println(u.Name)
		fmt.Println(u.Bio)
		os.Exit(2)
	}
}
