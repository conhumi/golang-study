package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    epsiron := math.Pow(2, 1-53) // 計算機イプシロン（十分に小さな差）
    zCurrent := x
    dz := x / 2
    for dz > epsiron {
        zNext := zCurrent - (math.Pow(zCurrent, 2) - x) / (2 * zCurrent)
        dz = zCurrent - zNext
        zCurrent = zNext
    }
    return zCurrent
}

func main() {
    fmt.Println(Sqrt(2))
}

