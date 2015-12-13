package mediainfo

// #cgo CFLAGS: -DUNICODE
// #cgo LDFLAGS: -lz -lzen -lpthread -lstdc++ -lmediainfo -ldl
// #include "go_mediainfo.h"
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

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

func NewMediaInfo() *MediaInfo {
	result := &MediaInfo{handle: C.GoMediaInfo_New()}
	runtime.SetFinalizer(result, func(f *MediaInfo) {
		C.GoMediaInfo_Delete(f.handle)
	})
	return result
}

func (mi *MediaInfo) Open(path string) error {
	s := C.GoMediaInfo_Open(mi.handle, C.CString(path))
	if s == 0 {
		return errors.New(fmt.Sprintf("MediaInfo can't open file: %s", path))
	}
	return nil
}

func (mi *MediaInfo) Close() {
	C.GoMediaInfo_Close(mi.handle)
}

func (mi *MediaInfo) Get(param string) string {
	return C.GoString(C.GoMediaInfoGet(mi.handle, C.CString(param)))
}
