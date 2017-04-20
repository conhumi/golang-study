package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %.0f", e)
}

func Sqrt(x float64) (float64, error) {
    if !(x < 0) {
        epsiron := math.Pow(2, 1-53) // 計算機イプシロン（十分に小さな差）
        zCurrent := x
        dz := x / 2
        for dz > epsiron {
            zNext := zCurrent - (math.Pow(float64(zCurrent), 2)-x)/(2*zCurrent)
            dz = zCurrent - zNext
            zCurrent = zNext
        }
        return zCurrent, nil
    }

    return 0, ErrNegativeSqrt(-2)
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

