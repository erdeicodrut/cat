package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	files := os.Args[1:]

	output := os.Stdout

	for _, fileName := range files {
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			log.Printf("File %v coulndn't open because %v", fileName, err)
		}

		text, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Could not read file %v, because %v", file.Name(), err)
		}

		_, err = output.Write(text)
		if err != nil {
			log.Printf("Could not print %v, because %v", file.Name(), err)
		}
	}
}
