package sidik

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	return dx*dy
}

func type18() {
	pic.Show(Pic)
}