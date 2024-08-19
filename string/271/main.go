package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	return strings.Join(strs, "bekateak")
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return strings.Split(strs, "bekateak")
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
