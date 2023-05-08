package sample_go_with_vulnerable_dep

import "github.com/hashicorp/vault/shamir"

func Hello() {
	_, err := shamir.Combine([][]byte{})
	if err != nil {
		return
	}
}
