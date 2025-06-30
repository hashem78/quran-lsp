package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("/home/mythi/development/quran-lsp/language-server/lsp.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	log.Println("Hello World")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		log.Println(string(bytes))
	}
}

func LspSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {

	return 0, nil, nil
}
