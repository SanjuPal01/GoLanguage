package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("foo.*", "foobar")
	fmt.Println("matched:", matched, "err:", err)

	matched, err = regexp.Match(`\w+fo\w+`, []byte("seafood"))
	fmt.Println("matched:", matched, "err:", err)

	// any character followed by at (_at)
	regex, err := regexp.Compile(`\wat`)
	fmt.Println("regex:", regex, "err:", err)
	str := `bat mat pot sat cat rat pat vat hat`

	fmt.Println("Matched String: ", regex.MatchString(str))
	fmt.Println("Matched String: ", regex.Match([]byte(str)))
	fmt.Println("Find String: ", string(regex.Find([]byte(str))))
	fmt.Printf("FindAll String: %q\n", regex.FindAll([]byte(str), -1))
	fmt.Println("Find Index: ", regex.FindIndex([]byte(str)))
	fmt.Println("Find All Index: ", regex.FindAllIndex([]byte(str), -1))
	fmt.Printf("Replace all Literal: %q\n", regex.ReplaceAllLiteral([]byte(str), []byte("dog")))
	fmt.Printf("Replace all Literal String: %q\n", regex.ReplaceAllLiteralString(str, "dog"))
	fmt.Println("String: ", regex.String())

	var re *regexp.Regexp
	re, err = regexp.Compile(`^(img-\d+)\.png$`)
	f := "img-284621831.png"
	fmt.Println("Replace All string: ", re.ReplaceAllString(f, `$1`))
}
