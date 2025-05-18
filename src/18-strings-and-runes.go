package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    const s = "สวัสดีt"
    const test = "ø"; // len of this is 2 because strings are represented as []byte

    fmt.Println("len:", len(s))
    fmt.Println("test:", len(test))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))
    fmt.Println("Rune count test:", utf8.RuneCountInString(test))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }

}

func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
