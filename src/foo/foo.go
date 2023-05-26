// filename: foo.go
package main

// generate dyn lib with: go build -buildmode=c-shared -o foo.so foo.go
import (
	"C"
	"fmt"
	"log"

	"github.com/hanwen/go-mtpfs/mtp"
	"github.com/hanwen/usb"
)

//export GetKey
func GetKey() *C.char {
	theKey := "123-456-789"
	return C.CString(theKey)
}

//export GetMTPInfo
func GetMTPInfo() *C.char {
	return C.CString(listMTPDevices())
}

func listMTPDevices() string {
	c := usb.NewContext()

	devs, err := mtp.FindDevices(c)
	if err != nil {
		log.Fatal(err)
	}
	if len(devs) == 0 {
		log.Fatal("No MTP devices found")
	}
	println("FOUND", fmt.Sprintf("%d", len(devs)), "DEVICES")

	for _, dev := range devs {
		var info mtp.DeviceInfo
		if err := dev.Open(); err != nil {
			log.Fatal(err)
		}
		if err := dev.GetDeviceInfo(&info); err != nil {
			log.Fatal(err)
		}
		// TODO: handle closing error?
		defer dev.Close()
		return info.String()
	}
	return ""
}

func main() {
	// println(listMTPDevices())
}
