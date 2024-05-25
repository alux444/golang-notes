package main

import (
	"fmt"
	"unicode"
)

func wrapText(text string) {
	widthLimit := 40
	curWidth := 0

	for _, r := range text {
		fmt.Printf("%c", r)
		curWidth++

		if curWidth > widthLimit && r != '\n' && unicode.IsSpace(r) {
			fmt.Println()
			curWidth = 0
		}
		if r == '\n' {
			curWidth = 0
		}
	}
	fmt.Println()
}

func main() {
	text := "Hello world, how is it going? It is ok.. The weather is beautiful."
	text2 := "Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır."
	text3 := "Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner"
	text4 := "Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler."
	wrapText(text)
	wrapText(text2)
	wrapText(text3)
	wrapText(text4)
}
