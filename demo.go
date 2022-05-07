package main

import (
	"fmt"
	"os"

	// To ensure importing external libs works, we will use my revregex lib as a demo ...")
	"github.com/xavier268/revregex"
)

func main() {
	fmt.Println("This message is sent from the demo program")

	pattern := "a(b|c)+a?"
	generator, err := revregex.NewGen(pattern)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	chooser := revregex.NewRandChooser()
	result := generator.Next(chooser)
	fmt.Println("Generated : ", result)
	err = generator.Verify(result)
	if err != nil {
		fmt.Println("Verification failed ", err)
		os.Exit(2)
	} else {
		fmt.Println("Demo is sucessful !")
	}
}
