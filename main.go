package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func ConvertToJpegFromPng(b []byte) []byte {
	nr := bytes.NewReader(b)
	buff := new(bytes.Buffer)
	img, err := png.Decode(nr)
	if err != nil {
		panic(err)
	}
	var rgba *image.RGBA
	if nrgba, ok := img.(*image.NRGBA); ok {
		if nrgba.Opaque() {
			rgba = &image.RGBA{
				Pix:    nrgba.Pix,
				Stride: nrgba.Stride,
				Rect:   nrgba.Rect,
			}
		}
	}
	if rgba != nil {
		log.Println("HEY using rbga")
		err = jpeg.Encode(buff, rgba, &jpeg.Options{Quality: 95})
	} else {
		log.Println("Hey not using rbga")
		err = jpeg.Encode(buff, img, &jpeg.Options{Quality: 95})
	}
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func readFile(f string) []byte {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return ConvertToJpegFromPng(b)
}

func main() {
	t1 := readFile("t1.png")
	t2 := readFile("t2.png")

	ioutil.WriteFile("t1.jpg", t1, os.ModePerm)
	ioutil.WriteFile("t2.jpg", t2, os.ModePerm)
}
