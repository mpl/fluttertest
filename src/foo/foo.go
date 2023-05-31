// filename: foo.go
package main

import (
	"C"
	/*
	   "github.com/hanwen/go-mtpfs/mtp"
	   "github.com/hanwen/usb"
	*/)

//export GetKey
func GetKey() *C.char {
	theKey := "123-456-789"
	return C.CString(theKey)
}

/*
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
			log.Printf("OPEN: %v\n", err)
			continue
		}
		if err := dev.GetDeviceInfo(&info); err != nil {
			log.Printf("GETINFO: %v\n", err)
			continue
		}
		// TODO: handle closing error?
		defer dev.Close()
		println("INFO: ", info.String())
		return info.String()
	}
	return ""
}
*/

func main() {
}
