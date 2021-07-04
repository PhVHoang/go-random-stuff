package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	_, err := rot.r.Read(b)
	if err == io.EOF {
		return 0, err
	}
	for i := range b {
		var base byte
		if b[i] >= 'a' && b[i] <= 'z' {
			base = 'a'
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			base = 'A'
		} else {
			continue
		}
		b[i] = (((b[i] - base) + 13) % 26) + base
	}	
	return len(b), nil	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
