package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {

	password := os.Args[1]
	long := os.Args[2]
	bang := os.Args[3]
	longb := []byte{}
	sep := func(a rune) bool {
		if a == rune('\\') {
			return true
		}
		return false
	}

	for _, tm := range strings.FieldsFunc(long, sep) {
		t, _ := hex.DecodeString(tm[1:3])
		longb = append(longb, t...)
	}

	h1 := md5.New()

	h2 := md5.New()
	h3 := md5.New()
	fmt.Fprint(h1, password)
	fmt.Fprintf(h2, "%s", h1.Sum(nil))
	fmt.Fprintf(h2, "%s", longb)
	fmt.Fprintf(h3, "%s", fmt.Sprintf("%X", (h2.Sum(nil))))
	fmt.Fprintf(h3, "%s", strings.ToUpper(bang))
	vcl := fmt.Sprintf("%X", h3.Sum(nil))
	fmt.Println(vcl)
	return
}
