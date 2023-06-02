// filename: foo.go
package main

import (
	"C"

	"fmt"
	"log"

	"github.com/hanwen/go-mtpfs/mtp"
	"github.com/hanwen/usb"
)
import "strings"

//export GetKey
func GetKey() *C.char {
	theKey := "123-456-789"
	return C.CString(theKey)
}

//export GetMTPInfo
func GetMTPInfo() *C.char {
	return C.CString(listMTPDevices())
}

//export GetFile
func GetFile() *C.char {
	// TODO: other signature than a string lol
	if err := getFile("Pictures/Screenshots/Screenshot_20171209-162420.png"); err != nil {
		return C.CString(fmt.Sprintf("GETFILE: %v", err))
	}
	return C.CString("")
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
		dev.USBDebug = true

		if err := dev.Open(); err != nil {
			log.Printf("OPEN: %v\n", err)
			continue
		}
		// TODO: handle closing error?
		defer dev.Close()

		if err := dev.GetDeviceInfo(&info); err != nil {
			log.Printf("GETINFO: %v\n", err)
			continue
		}
		// println("INFO: ", info.String())
		return info.String()
	}
	return ""
}

func getFile(filePath string) error {
	dev, err := mtp.SelectDevice("")
	if err != nil {
		return err
	}
	defer dev.Close()

	info := mtp.DeviceInfo{}
	err = dev.GetDeviceInfo(&info)
	if err != nil {
		return fmt.Errorf("GetDeviceInfo failed:", err)
	}

	if !strings.Contains(info.MTPExtension, "android.com:") {
		return fmt.Errorf("no android extensions", info.MTPExtension)
	}

	if err = dev.Configure(); err != nil {
		return fmt.Errorf("Configure failed:", err)
	}

	sids := mtp.Uint32Array{}
	err = dev.GetStorageIDs(&sids)
	if err != nil {
		return fmt.Errorf("GetStorageIDs failed: %v", err)
	}

	if len(sids.Values) == 0 {
		return fmt.Errorf("No storages")
	}

	id := sids.Values[0]

	// dev.USBDebug = true
	fd, err := dev.GetFileHandle(id, mtp.NOPARENT_ID, filePath)
	if err != nil {
		return err
	}
	println("FOUND HANDLE: ", fd)
	return nil
}

func main() {
	println("GO MAIN IS NOT CALLED BY DART")
	println(listMTPDevices())
	if err := getFile("Pictures/Screenshots/Screenshot_20171209-162420.png"); err != nil {
		println("ERROR: ", err.Error())
	}
}

/*
	// 500 + 512 triggers the null read case on both sides.
	const testSize = 500 + 512
	name := fmt.Sprintf("mtp-doodle-test%x", rand.Int31())

	send := ObjectInfo{
		StorageID:        id,
		ObjectFormat:     OFC_Undefined,
		ParentObject:     0xFFFFFFFF,
		Filename:         name,
		CompressedSize:   uint32(testSize),
		ModificationDate: time.Now(),
		Keywords:         "bla",
	}
	data := make([]byte, testSize)
	for i := range data {
		data[i] = byte('0' + i%10)
	}

	_, _, handle, err := dev.SendObjectInfo(id, 0xFFFFFFFF, &send)
	if err != nil {
		t.Fatal("SendObjectInfo failed:", err)
	} else {
		buf := bytes.NewBuffer(data)
		t.Logf("Sent objectinfo handle: 0x%x\n", handle)
		err = dev.SendObject(buf, int64(len(data)))
		if err != nil {
			t.Log("SendObject failed:", err)
		}
	}

	magicStr := "life universe"
	magicOff := 21
	magicSize := 42

	err = dev.AndroidBeginEditObject(handle)
	if err != nil {
		t.Errorf("AndroidBeginEditObject: %v", err)
		return
	}
	err = dev.AndroidTruncate(handle, int64(magicSize))
	if err != nil {
		t.Errorf("AndroidTruncate: %v", err)
	}
	buf := bytes.NewBufferString(magicStr)
	err = dev.AndroidSendPartialObject(handle, int64(magicOff), uint32(buf.Len()), buf)
	if err != nil {
		t.Errorf("AndroidSendPartialObject: %v", err)
	}
	if buf.Len() > 0 {
		t.Errorf("buffer not consumed")
	}
	err = dev.AndroidEndEditObject(handle)
	if err != nil {
		t.Errorf("AndroidEndEditObject: %v", err)
	}

	buf = &bytes.Buffer{}
	err = dev.GetObject(handle, buf)
	if err != nil {
		t.Errorf("GetObject: %v", err)
	}

	if buf.Len() != magicSize {
		t.Errorf("truncate failed:: %v", err)
	}
	for i := 0; i < len(magicStr); i++ {
		data[i+magicOff] = magicStr[i]
	}
	want := string(data[:magicSize])
	if buf.String() != want {
		t.Errorf("read result was %q, want %q", buf.String(), want)
	}

	buf = &bytes.Buffer{}
	err = dev.AndroidGetPartialObject64(handle, buf, int64(magicOff), 5)
	if err != nil {
		t.Errorf("AndroidGetPartialObject64: %v", err)
	}
	want = magicStr[:5]
	got := buf.String()
	if got != want {
		t.Errorf("AndroidGetPartialObject64: got %q want %q", got, want)
	}

	// Try write at end of file.
	err = dev.AndroidBeginEditObject(handle)
	if err != nil {
		t.Errorf("AndroidBeginEditObject: %v", err)
		return
	}
	buf = bytes.NewBufferString(magicStr)
	err = dev.AndroidSendPartialObject(handle, int64(magicSize), uint32(buf.Len()), buf)
	if err != nil {
		t.Errorf("AndroidSendPartialObject: %v", err)
	}
	if buf.Len() > 0 {
		t.Errorf("buffer not consumed")
	}
	err = dev.AndroidEndEditObject(handle)
	if err != nil {
		t.Errorf("AndroidEndEditObject: %v", err)
	}
	buf = &bytes.Buffer{}
	err = dev.GetObject(handle, buf)
	if err != nil {
		t.Errorf("GetObject: %v", err)
	}
	want = string(data[:magicSize]) + magicStr
	got = buf.String()
	if got != want {
		t.Errorf("GetObject: got %q want %q", got, want)
	}

	err = dev.DeleteObject(handle)
	if err != nil {
		t.Fatalf("DeleteObject failed: %v", err)
	}

}
*/
