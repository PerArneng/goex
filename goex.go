package main

import (
	"fmt"
	"io"
	"os"
	"errors"
)

type ErrorHandlerReader interface {

	Read(buff []byte, 
			  errv ...func(err error)) (n int)

}

type ReaderWrapper struct {
	r io.Reader 
}

func (self *ReaderWrapper) Read(
				buff []byte, 
			  	errv ...func(err error)) (n int) {

	return 0
}


func NewReaderWrapper(r io.Reader) *ReaderWrapper {
	return &ReaderWrapper{r}
}

func OpenFile(fileName string, errv ...func(err error)) (file *os.File) {

	return nil
}

func main() {

	exitErrorHandler := func(err error) {
		fmt.Printf("An error occurred '%s'\nShutting Down!\n", 
				   err.Error())
		os.Exit(1)
	}

	exitErrorHandler(errors.New("Test Error"))

	file := OpenFile("input.txt", exitErrorHandler)

	fmt.Printf("dummy %s\n", file)
 //   if err != nil { panic(err) }


}