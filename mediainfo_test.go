package mediainfo_test

import (
	"fmt"
	"github.com/zhulik/go_mediainfo"
	"io/ioutil"
	"os"
	"testing"
)

const (
	ogg        = "testdata/test.ogg"
	mp3        = "testdata/test.mp3"
	non_exists = "testdata/non_exists.ogg"
)

func TestOpenWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(ogg)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(mp3)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithUnexistsFile(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(non_exists)
	if error == nil {
		t.Fail()
	}
}

func TestOpenMemoryWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(ogg)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(mp3)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithEmptyArray(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenMemory([]byte{})
	if error == nil {
		t.Fail()
	}
}

func TestInformWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestInformWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestDurationWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if mi.Duration() != 3494 {
		t.Fail()
	}
}

func TestDurationWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if mi.Duration() != 87771 {
		t.Fail()
	}
}

func TestCodecWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if mi.Codec() != "OGG" {
		t.Fail()
	}
}

func TestCodecWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if mi.Codec() != "MPEG Audio" {
		t.Fail()
	}
}

func TestFormatWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if mi.Format() != "OGG" {
		t.Fail()
	}
}

func TestFormatWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if mi.Format() != "MPEG Audio" {
		t.Fail()
	}
}

//----------------------------------------------------------------------------------------------------------------------
func BenchmarkOpenAndDurationWithOgg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		mi.OpenFile(ogg)

		mi.Duration()
	}
}

func BenchmarkOpenAndDurationWithMp3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		mi.OpenFile(mp3)

		mi.Duration()
	}
}

func BenchmarkOpenMemoryAndDurationWithOgg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		f, _ := os.Open(ogg)
		bytes, _ := ioutil.ReadAll(f)

		mi.OpenMemory(bytes)
		mi.Duration()
	}
}

func BenchmarkOpenMemoryAndDurationWithMp3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		f, _ := os.Open(mp3)
		bytes, _ := ioutil.ReadAll(f)

		mi.OpenMemory(bytes)
		mi.Duration()
	}
}

//----------------------------------------------------------------------------------------------------------------------

func ExampleUsage() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	mi := mediainfo.NewMediaInfo()
	err = mi.OpenMemory(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(mi.AvailableParameters()) // Print all supported params for Get
	fmt.Println(mi.Get("BitRate"))        // Print bitrate
}
