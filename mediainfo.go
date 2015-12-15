package mediainfo

// #cgo CFLAGS: -DUNICODE
// #cgo LDFLAGS: -lz -lzen -lpthread -lstdc++ -lmediainfo -ldl
// #include "go_mediainfo.h"
import "C"

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

// MediaInfo - represents MediaInfo class, all interaction with libmediainfo through it
type MediaInfo struct {
	handle unsafe.Pointer
}

func init() {
	C.setlocale(C.LC_CTYPE, C.CString(""))
	C.MediaInfoDLL_Load()

	if C.MediaInfoDLL_IsLoaded() == 0 {
		panic("Cannot load mediainfo")
	}
}

// NewMediaInfo - constructs new MediaInfo
func NewMediaInfo() *MediaInfo {
	result := &MediaInfo{handle: C.GoMediaInfo_New()}
	runtime.SetFinalizer(result, func(f *MediaInfo) {
		f.Close()
		C.GoMediaInfo_Delete(f.handle)
	})
	return result
}

// OpenFile - opens file
func (mi *MediaInfo) OpenFile(path string) error {
	s := C.GoMediaInfo_OpenFile(mi.handle, C.CString(path))
	if s == 0 {
		return fmt.Errorf("MediaInfo can't open file: %s", path)
	}
	return nil
}

// OpenMemory - opens memory buffer
func (mi *MediaInfo) OpenMemory(bytes []byte) error {
	if len(bytes) == 0 {
		return fmt.Errorf("Buffer is empty")
	}
	s := C.GoMediaInfo_OpenMemory(mi.handle, (*C.char)(unsafe.Pointer(&bytes[0])), C.size_t(len(bytes)))
	if s == 0 {
		return fmt.Errorf("MediaInfo can't open memory buffer")
	}
	return nil
}

// Close - closes file
func (mi *MediaInfo) Close() {
	C.GoMediaInfo_Close(mi.handle)
}

// Get - allow to read info from file
func (mi *MediaInfo) Get(param string) string {
	return C.GoString(C.GoMediaInfoGet(mi.handle, C.CString(param)))
}

// Inform returns string with suffary file information, like mediainfo util
func (mi *MediaInfo) Inform() string {
	return C.GoString(C.GoMediaInfoInform(mi.handle))
}

// AvailableParameters returns string with all available Get params and it's descriptions
func (mi *MediaInfo) AvailableParameters() string {
	return C.GoString(C.GoMediaInfoAvailableParameters(mi.handle))
}

// Duration returns file duration
func (mi *MediaInfo) Duration() int {
	duration, _ := strconv.Atoi(mi.Get("Duration"))
	return duration
}
