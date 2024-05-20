package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	stringLiteral := "<html>\n\t<body>\"Hello\"</body>\n</html>"
	rawStringLiteral := `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(stringLiteral)
	fmt.Println(rawStringLiteral)

	i := 1 + 2
	//int to string
	fmt.Println(strconv.Itoa(i))

	fmt.Println(strconv.FormatBool((true)) + " " + strconv.FormatBool(false))

	//len counts bytes in a name
	name := "John"
	fmt.Println(len(name))
	//unicode can be multiple bytes
	name = "Señor"
	fmt.Println(len(name))
	//use rune to count characters
	fmt.Println(len([]rune(name)))
	//use utf8 to count characters
	fmt.Println(utf8.RuneCountInString(name))

	word := "hello"
	repeat := strings.Repeat("!!??!!", 3)
	fmt.Println(repeat + strings.ToUpper(word) + repeat)

	word = "      hello   "
	fmt.Println(strings.TrimSpace(word))
	fmt.Println(strings.TrimRight(word, " "))

	word = "$$hellos@@"
	fmt.Println(strings.Trim(word, "$@s"))
}