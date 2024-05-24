package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func bytesVsString() {
	str := "hey"
	bytes := []byte{104, 101, 121}
	fmt.Println(str)
	fmt.Printf("%s\n", bytes)
}

func runes() {
	//runes are typeless - they can be any integer type
	var aRune rune = 'a'
	fmt.Printf("%T\n", aRune)
}

func charDecEncoded(start byte, end byte) {
	// char, decimal, hex, encoded
	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	for n := start; n <= end; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}

func strVsBytesVsRunes() {
	str := "Yūgen ☯ 💀"
	// can't change a string
	// a string is a read-only byte-slice

	bytes := []byte(str)
	// can change a byte slice
	// bytes[0] = 'N'
	// bytes[1] = 'o'

	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))
	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0]) // ok
	fmt.Printf("2nd byte   : %c\n", str[1]) // not ok
	//string is a sequence of bytes, so can't access the 2nd byte properly, it's a multi-byte character

	// disadvantage: each one is 4 bytes
	// but indexes each individual character to one rune
	runes := []rune(str)
	fmt.Printf("%s\n", string(runes))
	for i, r := range runes {
		fmt.Printf("runes[%2d] = % -12x = %q\n", i, string(r), r)
	}
}

func decoding() {
	r, size := utf8.DecodeRuneInString("öykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	text := "öykü"
	//for range just gives the position of current rune rather than size
	//i.e index of each character

	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)
		fmt.Printf(" size: %d bytes.\n", size)

		//have to increase index by size of bytes
		i += size
	}
}

func changingRune() {
	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	//need to find start and end index of first character
	var size int

	//get index of byte
	_, size = utf8.DecodeRune(word)
	fmt.Println(size)
	// overwrite the current bytes with the new uppercased bytes
	copy(word[:size], bytes.ToUpper(word[:size]))

	// to get printed bytes/runes need to be encoded in a string
	fmt.Printf("%s = % [1]x\n", word)
}

func dumping() {
	// no information = nil header
	// empty := ""
	// dump(empty)

	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	for i := range hello {
		dump(hello[i : i+1])
	}

	//new string header, as new structure is created
	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

func dump(s string) {
	ptr := unsafe.StringData(s)
	fmt.Printf("%q: %+v\n", s, ptr)
}

func exercise() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range words {
		fmt.Printf("%q\n", s)

		fmt.Printf("Bytes: %d -> Runes: %d\n", len(s), utf8.RuneCountInString(s))
		fmt.Printf("Bytes: % x\n", s)
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}

		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("First rune: %c %d\n", r, size)
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("Last rune: %c %d\n", r, size)

		runes := []rune(s)
		fmt.Printf("%x\n", []rune(runes[:2]))
		fmt.Printf("%q\n", []rune(runes[len(runes)-2:]))
	}
}

func main() {
	bytesVsString()
	runes()
	charDecEncoded('\n', 'Z')
	strVsBytesVsRunes()
	decoding()
	changingRune()
	dumping()
	exercise()
}
