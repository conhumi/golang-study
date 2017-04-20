package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    var picture [][]uint8 
    for y := 0; y < dy; y++ {
        var line []uint8
        for x := 0; x < dx; x++ {
            pixcel := y^x
            line = append(line, uint8(pixcel))
        }
        picture = append(picture, line)
    }
    return picture
}

func main() {
    pic.Show(Pic)
}

