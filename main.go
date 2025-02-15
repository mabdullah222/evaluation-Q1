package main

import (
	"fmt"
	"question1/utils"
)

func main() {
	currDir := "/home///user//documents"
	relPath := "../projects/code"

	absPath := utils.AbsPath(currDir, relPath)
	fmt.Println("Absolute Path:", absPath)
}
