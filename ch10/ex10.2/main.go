package main

import (
	"os"

	"ex10.2/archive"
	_ "ex10.2/archive/tar"
	_ "ex10.2/archive/zip"
)

func main() {
	files, err := archive.FileNames(os.Stdin)
	if err != nil {
		panic(err)
	}

	for _, name := range files {
		println(name)
	}
}
