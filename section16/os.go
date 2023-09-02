package main

import (
	"os"
)

func main1() {

	f, _ := os.Create("./section16/foo.txt")

	f.Write([]byte("Hello world!"))

	f.WriteAt([]byte("Goodbye world!"), 6)

	f.Seek(0, os.SEEK_END)

	f.WriteString("This is the end...")
}
