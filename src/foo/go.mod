module github.com/mpl/fluttertest/src/foo

go 1.20

require (
	github.com/hanwen/go-mtpfs v1.0.0
	github.com/hanwen/usb v0.0.0-20141217151552-69aee4530ac7
)

replace (
	github.com/hanwen/go-mtpfs => ./notvendor/github.com/hanwen/go-mtpfs
	github.com/hanwen/usb => ./notvendor/github.com/hanwen/usb
)
