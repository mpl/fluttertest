// filename: lib.go
package main

// generate dyn lib with: go build -buildmode=c-shared -o foo.so foo.go
import "C"

//export GetKey
func GetKey() *C.char {
	theKey := "123-456-789"
	return C.CString(theKey)
}

func main() {}
