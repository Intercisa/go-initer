package file

import (
	"fmt"
	"os"

	"barnasipiczki.com/go-initer/util"
)

func WriteMainIntoFile() {
	file, err := os.Create("main.go")
	if err != nil {
		fmt.Println("Error durring writing the main file! ", err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(util.GetMain())
}
