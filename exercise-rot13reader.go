package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13reader rot13Reader) Read(slice []byte) (int, error) {

    lowerLUT := make([]byte, 26)
    strings.NewReader("nopqrstuvwxyzabcdefghijklm").Read(lowerLUT)
    upperLUT := make([]byte, 26)
    strings.NewReader("NOPQRSTUVWXYZABCDEFGHIJKLM").Read(upperLUT)

    size, err := rot13reader.r.Read(slice)

    for i := 0; i < size; i++ {
        if byte('A') <= slice[i] && slice[i] <= byte('Z') {
            slice[i] = upperLUT[slice[i]-byte('A')]
        } else if byte('a') <= slice[i] && slice[i] <= byte('z') {
            slice[i] = lowerLUT[slice[i]-byte('a')]
        }
    }
    return size, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

