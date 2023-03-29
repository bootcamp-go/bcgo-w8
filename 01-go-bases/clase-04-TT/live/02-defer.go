package main

import (
	"io"
	"os"
)

func main() {
	// Task: read data from a file and send it to other file
	// - f1: file_1.txt
	// - f2: file_2.txt

	if err := SendData("live/file_1.txt", "live/file_2.txt"); err != nil {
		println(err.Error())
	}

	println("done")
}

func SendData(filename, other string) (err error) {
	var f1 *os.File
	f1, err = os.Open(filename)
	if err != nil {
		return
	}
	defer func() {
		println("closing f1")
		f1.Close()
	}()

	var f2 *os.File
	f2, err = os.OpenFile(other, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer func() {
		println("closing f2")
		f1.Close()
	}()

	// panic("stop here")

	// send data
	var b []byte
	b, err = io.ReadAll(f1)
	if err != nil {
		return
	}

	_, err = f2.Write(b)
	return
}