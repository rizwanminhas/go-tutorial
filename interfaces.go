package main

import (
	"bytes"
	"fmt"
)

func main() {

	// 1-
	var w Writer = ConsoleWriter{} // instead of ConsoleWrite we could use any writer that has a function with the same signature as Write
	w.Write([]byte("rizwan minhas"))

	// 2-
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// 3- This will only print 8 bytes per line.
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("my name is rizwan minhas and this is a test."))
	wc.Close()

	// 4- type conversion
	bwc, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println("everything is ok!", bwc)
	} else {
		fmt.Println("error:", ok)
	}

	// 5- empty interface, you will have to use type conversion because an empty interface doesn't have any methods
	var myObj interface{} = NewBufferedWriterCloser()
	if wc2, ok := myObj.(WriterCloser); ok {
		wc2.Write([]byte("this also works"))
		wc2.Close()
	}

}

// 1- contains the behavior, a function that takes an array of bytes and returns an int or an error
type Writer interface {
	Write([]byte) (int, error)
}

// 1- In order to use Writer interface we don't explisitly use extends or implements, in golang we implicitly use interfaces by having functions with same signatures.
type ConsoleWriter struct{}

// 1-
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, error := fmt.Println(string(data))
	return n, error
}

// 2-
type Incrementer interface {
	Increment() int
}

// 2-
type IntCounter int

// 2-
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// 3-
type Closer interface {
	Close() error
}

// 3 - combining multiple interfaces.
type WriterCloser interface {
	Writer
	Closer
}

// 3-
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// 3-
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// 3-
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// 3- a constructor function
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
