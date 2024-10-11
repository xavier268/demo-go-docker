package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/xavier268/rgen"
	// To ensure importing external libs works, we will use my rgen lib as a demo ...")
)

func main() {
	fmt.Println("This message is sent from the demo program")

	pattern := "a(b|c)+a?c"
	maxlen := 4

	for s := range rgen.All(pattern, maxlen) {
		fmt.Print(s, " ")
		if len(s) > maxlen {
			fmt.Println("Too long !")
			os.Exit(2)
		}
		if b, err := regexp.Match(pattern, []byte(s)); err != nil || !b {
			fmt.Println("KO", err)
			os.Exit(3)
		}
		fmt.Println("OK")
	}
	fmt.Println("Demo successful")
}
