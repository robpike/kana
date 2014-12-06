// Copyright 2014 Rob Pike. All rights reserved.
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
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		do(string(data))
		return
	}
	do(strings.Join(os.Args[1:], " "))
}

func do(s string) {
	prevKana := true
	skip := false
	for i, r := range s {
		k, ok := kana[r]
		if skip {
			skip = false
			continue
		}
		if !ok {
			if prevKana {
				fmt.Printf(" ")
			}
			fmt.Printf("%c", r)
			prevKana = false
			continue
		}
		if !prevKana {
			fmt.Printf(" ")
		}
		prevKana = true
		// Is there a modifier?
		r2, _ := utf8.DecodeRuneInString(s[i+utf8.RuneLen(r):])
		if small[r2] {
			skip = true
			k2, ok := mod[r2]
			if ok {
				fmt.Printf("%s%s", k[:len(k)-1], k2[1:])
				continue
			}
			k2, ok = vowel[r2]
			if ok {
				fmt.Printf("%s-", k)
				continue
			}
			// Otherwise it's just odd.
			fmt.Printf("<%s.%s>", k, odd[r2])
			continue
		}
		fmt.Printf("%s", k)
	}
	fmt.Print("\n")
}

var kana = map[rune]string{
	'あ': "a",
	'い': "i",
	'う': "u",
	'え': "e",
	'お': "o",
	'か': "ka",
	'が': "ga",
	'き': "ki",
	'ぎ': "gi",
	'く': "ku",
	'ぐ': "gu",
	'け': "ke",
	'げ': "ge",
	'こ': "ko",
	'ご': "go",
	'さ': "sa",
	'ざ': "za",
	'し': "shi",
	'じ': "zi",
	'す': "su",
	'ず': "zu",
	'せ': "se",
	'ぜ': "ze",
	'そ': "so",
	'ぞ': "zo",
	'た': "ta",
	'だ': "da",
	'ち': "chi",
	'ぢ': "di",
	'つ': "tsu",
	'づ': "du",
	'て': "te",
	'で': "de",
	'と': "to",
	'ど': "do",
	'な': "na",
	'に': "ni",
	'ぬ': "nu",
	'ね': "ne",
	'の': "no",
	'は': "ha",
	'ば': "va",
	'ぱ': "pa",
	'ひ': "hi",
	'び': "vi",
	'ぴ': "pi",
	'ふ': "fu",
	'ぶ': "bu",
	'ぷ': "pu",
	'へ': "he",
	'べ': "ve",
	'ぺ': "pe",
	'ほ': "ho",
	'ぼ': "vo",
	'ぽ': "po",
	'ま': "ma",
	'み': "mi",
	'む': "mu",
	'め': "me",
	'も': "mo",
	'や': "ya",
	'ゆ': "yu",
	'よ': "yo",
	'ら': "ra",
	'り': "ri",
	'る': "ru",
	'れ': "re",
	'ろ': "ro",
	'わ': "wa",
	'ゐ': "wi",
	'ゑ': "we",
	'を': "wo",
	'ん': "n",
	'ゔ': "vu",
	'ア': "a",
	'イ': "i",
	'ウ': "u",
	'エ': "e",
	'オ': "o",
	'カ': "ka",
	'ガ': "ga",
	'キ': "ki",
	'ギ': "gi",
	'ク': "ku",
	'グ': "gu",
	'ケ': "ke",
	'ゲ': "ge",
	'コ': "ko",
	'ゴ': "go",
	'サ': "sa",
	'ザ': "za",
	'シ': "shi",
	'ジ': "zi",
	'ス': "su",
	'ズ': "zu",
	'セ': "se",
	'ゼ': "ze",
	'ソ': "so",
	'ゾ': "zo",
	'タ': "ta",
	'ダ': "da",
	'チ': "chi",
	'ヂ': "di",
	'ツ': "tsu",
	'ヅ': "du",
	'テ': "te",
	'デ': "de",
	'ト': "to",
	'ド': "do",
	'ナ': "na",
	'ニ': "ni",
	'ヌ': "nu",
	'ネ': "ne",
	'ノ': "no",
	'ハ': "ha",
	'バ': "va",
	'パ': "pa",
	'ヒ': "hi",
	'ビ': "vi",
	'ピ': "pi",
	'フ': "fu",
	'ブ': "bu",
	'プ': "pu",
	'ヘ': "he",
	'ベ': "ve",
	'ペ': "pe",
	'ホ': "ho",
	'ボ': "vo",
	'ポ': "po",
	'マ': "ma",
	'ミ': "mi",
	'ム': "mu",
	'メ': "me",
	'モ': "mo",
	'ヤ': "ya",
	'ユ': "yu",
	'ヨ': "yo",
	'ラ': "ra",
	'リ': "ri",
	'ル': "ru",
	'レ': "re",
	'ロ': "ro",
	'ワ': "wa",
	'ヰ': "wi",
	'ヱ': "we",
	'ヲ': "wo",
	'ン': "n",
	'ヴ': "vu",
}

var small = map[rune]bool{
	'ぁ': true,
	'ぃ': true,
	'ぅ': true,
	'ぇ': true,
	'ぉ': true,
	'っ': true,
	'ゃ': true,
	'ゅ': true,
	'ょ': true,
	'ゎ': true,
	'ゕ': true,
	'ゖ': true,

	'ァ': true,
	'ィ': true,
	'ゥ': true,
	'ェ': true,
	'ォ': true,
	'ッ': true,
	'ャ': true,
	'ュ': true,
	'ョ': true,
	'ヮ': true,
	'ヵ': true,
	'ヶ': true,
}

var vowel = map[rune]string{
	'ぁ': "a",
	'ぃ': "i",
	'ぅ': "u",
	'ぇ': "e",
	'ぉ': "o",

	'ァ': "a",
	'ィ': "i",
	'ゥ': "u",
	'ェ': "e",
	'ォ': "o",
}

var mod = map[rune]string{
	'ゃ': "ya",
	'ゅ': "yu",
	'ょ': "yo",

	'ャ': "ya",
	'ュ': "yu",
	'ョ': "yo",
}

var odd = map[rune]string{
	'っ': "hold",  // tsu == hold consonant
	'ゕ': "count", // ka == counting mark
	'ゖ': "count", // ke == counting mark
	'ッ': "hold",  // tsu == hold consonant
	'ヵ': "count", // ka == counting mark
	'ヶ': "count", // ke == counting mark
}
