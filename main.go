package main

import (
	"fmt"

	"github.com/MachmannZheng/GolangPlayGround/fileOperation/CSV"
	"github.com/MachmannZheng/GolangPlayGround/testDelete"
)

func main() {
	fmt.Println("the main function is running")
	testDelete.ReadCSV1()
	CSV.ReadCSV()
}
