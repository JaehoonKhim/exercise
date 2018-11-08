package main

import "fmt"

func AddComment(memo string) {
	fmt.Printf("add %s", memo)
}

func RemoveComment(memo string) {
	fmt.Printf("remove %s", memo)
}

func Done(memo string) {
	fmt.Printf("done %s", memo)
}
