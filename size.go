package imago

import (
	"image"
)

func Height(i image.Image) int {
	return i.Bounds().Size().Y
}

func Width(i image.Image) int {
	return i.Bounds().Size().X
}
