package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.Open("/app/shit.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	count := 0
	reader := bufio.NewReader(data)
	buffer := bytes.NewBuffer(make([]byte, 0))
	part := make([]byte, 4096)
	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	log.Println(buffer.Len())
	log.Println(buffer)
}