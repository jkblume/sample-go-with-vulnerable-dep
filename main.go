package sample_go_with_vulnerable_dep

import (
	"fmt"
	"github.com/hashicorp/vault/shamir"
)

func init() {
	_, err := shamir.Combine([][]byte{})
	if err != nil {
		return
	}
}

func Hello() {
	_, err := shamir.Combine([][]byte{})
	if err != nil {
		return
	}
}

func Todo() {
	fmt.Println("Hello World")
}
