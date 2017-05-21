// Written by https://xojoc.pw. Apache license 2.0. No warranty.

// Package must provides "must" variants of common functions.
//
// These functions are handy in main programs and tests.
package must // import "xojoc.pw/must"

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"runtime"
)

// OK prints a stack trace and calls log.Fatal if err != nil.
// Otherwise does nothing.
func OK(err error) {
	if err == nil {
		return
	}
	buf := make([]byte, 8000)
	runtime.Stack(buf, false)
	fmt.Println(string(buf))
	log.Fatal(err)
}

func Open(path string) *os.File {
	file, err := os.Open(path)
	OK(err)
	return file
}

func Close(c io.Closer) {
	OK(c.Close())
}

func Create(path string) *os.File {
	file, err := os.Create(path)
	OK(err)
	return file
}

func Remove(path string) {
	OK(os.Remove(path))
}

func ReadImage(path string) image.Image {
	img, _, err := image.Decode(bytes.NewBuffer(ReadFile(path)))
	OK(err)
	return img
}

func WriteImage(file string, img image.Image) {
	ext := path.Ext(file)
	switch ext {
	case ".png":
		f := Create(file)
		OK(png.Encode(f, img))
	default:
		panic("unknown file extension：" + ext)
	}
}

func ReadFile(path string) []byte {
	d, err := ioutil.ReadFile(path)
	OK(err)
	return d
}

func ReadAll(r io.Reader) []byte {
	b, err := ioutil.ReadAll(r)
	OK(err)
	return b
}

func URL(s string) *url.URL {
	u, err := url.Parse(s)
	OK(err)
	return u
}
