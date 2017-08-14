package rap

import (
	"image/color"
	"math/rand"
	"time"
)

func Random() color.RGBA {
	rand.Seed(time.Now().UTC().UnixNano())

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))

	return color.RGBA{r, g, b, 255}
}
