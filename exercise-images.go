package main

import "golang.org/x/tour/pic"
import (
    "image"
    "image/color"
)

type Image struct {
    width  int
    height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    value := uint8(x ^ y)
    return color.RGBA{value, value, 255, 255}
}

func main() {
    m := Image{256, 256}
    pic.ShowImage(m)
}
