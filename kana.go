// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Kana accepts, as argument text or as standard input, Japanese
UTF-8 text. It copies the text to standard output, converting
katakana and hiragana to romaji, leaving the rest (including
kanji) unmodified.
*/
package main // import "robpike.io/cmd/kana"

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"robpike.io/nihongo"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, nihongo.RomajiReader(os.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	text := nihongo.RomajiString(strings.Join(os.Args[1:], " "))
	fmt.Printf("%s\n", text)
}
