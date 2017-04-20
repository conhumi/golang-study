package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(slice []byte) (int, error) {
    size := len(slice)
    for i := 0; i < size; i++ {
        slice[i] = byte('A')
    }
    return size, nil
}

func main() {
    reader.Validate(MyReader{})
}

