package main

import (
	"fmt"
	"os"

	"example.com/fileio/fileutils"
)

func main() {

	rootPath, _ := os.Getwd()

	fmt.Println(rootPath)

	c, err := fileutils.ReadTextFile(rootPath + "/FileIO" + "/text.txt")

	if err == nil {
		fmt.Printf("File : %s\n", c)
	} else {
		fmt.Println("Error: ", err)
	}
}
