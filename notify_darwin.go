package nfy

/*
// Compiler flags.
#cgo CFLAGS: -Wall -x objective-c -arch x86_64 -std=gnu99 -fobjc-arc
// Linker flags.
#cgo LDFLAGS: -framework Foundation -arch x86_64

#import "notify_darwin.h"
*/
import "C"
import (
	"unsafe"
)

// Notify displays a NSUserNotification.
func Notify(title, message string) error {
	sound := "Ping"
	t := C.CString(title)
	m := C.CString(message)
	s := C.CString(sound)
	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(m))
	defer C.free(unsafe.Pointer(s))

	C.BannerNotify(t, m, s)

	return nil
}
