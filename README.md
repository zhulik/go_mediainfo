# mediainfo
Golang binding for [libmediainfo](https://mediaarea.net/en/MediaInfo)

Duration, Bitrate, Codec, Streams and a lot of other meta-information about media files can be extracted through it.

For now supports only media files with one stream. Bindings for MediaInfoList is not provided. It can be easy fixed if anybody ask me.

Works through MediaInfoDLL/MediaInfoDLL.h(dynamic load and so on), so your mediainfo installation should has it.

Supports direct reading files by name and reading data from []byte buffers(without copying it for C calls).

Documentation for libmediainfo is poor and ascetic, can be found [here](https://mediaarea.net/en/MediaInfo/Support/SDK).

Your advices and suggestions are welcome!

## Example
```go
package main

import (
	"fmt"
	"github.com/zhulik/go_mediainfo"
	"io/ioutil"
	"os"
)

func main() {
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
	fmt.Println(mi.Get("BitRate")) // Print bitrate
}

```

Read the [documentation](https://godoc.org/github.com/zhulik/go_mediainfo) for other functions

## Contributing
You know=)