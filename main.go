package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome files in golang!")
	content := "This needs to go in file - LearnCodeOnline.in"

	file, err := os.Create("./mylco.txt")
	checkNilErr(err)

	// if err != nil {
	// 	panic(err)
	// }

	lengh, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Lengh is: ", lengh)
	defer file.Close()

	readFile("./mylco.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)
	fmt.Println("Test data inside the file is \n", string(databyte))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
