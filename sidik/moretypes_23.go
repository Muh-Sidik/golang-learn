package sidik

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{s: 1}
}

func type23() {
	wc.Test(WordCount)
}