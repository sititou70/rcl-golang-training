package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outType = flag.String("t", "jpeg", "output image type: jpeg | png | gif")

func main() {
	flag.Parse()

	if err := convertImg(*outType, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "converting error: %v\n", err)
		os.Exit(1)
	}
}

func convertImg(outType string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch outType {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	}

	return nil
}
