package imago

import (
	"image"
)

//Height of image
func Height(i image.Image) int {
	return i.Bounds().Size().Y
}

//Width of image
func Width(i image.Image) int {
	return i.Bounds().Size().X
}
