package main

import (
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"

	"github.com/UnderTreeTech/letteravatar"
)

var names = []string{
	"Alice",
	"Bob",
	"Carol",
	"Dave",
	"Eve",
	"Frank",
	"Gloria",
	"Henry",
	"Isabella",
	"James",
	"Жозефина",
	"Ярослав",
	"中国文化",
}

func main() {
	font, err := loadFont("./PingFang-SC-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	opt := &letteravatar.Options{
		FontSize: 20,
		Font:     font,
	}
	for _, name := range names {

		img, err := letteravatar.Draw(100, ([]rune(name))[:5], opt)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(name + ".png")
		if err != nil {
			log.Fatal(err)
		}

		err = png.Encode(file, img)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadFont(path string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	return f, nil
}
