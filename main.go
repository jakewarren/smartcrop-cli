package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/muesli/smartcrop"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version string

var (
	app   = kingpin.New("smartcrop-cli", "smartcrop finds good image crops for arbitrary crop sizes")
	input = app.Arg("input", "input file").Required().ExistingFile()

	width      = app.Flag("width", "preferred width of image (px)").Short('w').Default("250").Int()
	height     = app.Flag("height", "preferred height of image (px)").Short('h').Default("250").Int()
	quality    = app.Flag("quality", "image quality of output file").Short('q').Default("90").Int()
	outputFile = app.Flag("output-file", "name of the output file").Short('o').Default("./smartcrop.jpg").String()
)

func main() {
	app.Version(version).VersionFlag.Short('V')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetHandler(cli.New(os.Stderr))

	fi, _ := os.Open(*input)
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		log.WithError(err).Fatal("Error decoding image")
	}

	topCrop, err := smartcrop.SmartCrop(img, *width, *height)
	if err != nil {
		log.WithError(err).Fatal("Error cropping image")
	}
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	sub, ok := img.(SubImager)
	if ok {
		cropImage := sub.SubImage(topCrop)
		writeImage("jpeg", cropImage, *outputFile)
	} else {
		log.WithError(err).Fatal("No SubImage support")
	}
}

func writeImage(imgtype string, img image.Image, name string) error {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		panic(err)
	}

	switch imgtype {
	case "png":
		return writeImageToPng(img, name)
	case "jpeg":
		return writeImageToJpeg(img, name)
	}

	return errors.New("Unknown image type")
}

func writeImageToJpeg(img image.Image, name string) error {
	fso, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fso.Close()

	return jpeg.Encode(fso, img, &jpeg.Options{Quality: *quality})
}

func writeImageToPng(img image.Image, name string) error {
	fso, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fso.Close()

	return png.Encode(fso, img)
}
