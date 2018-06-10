package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	files := args[1:]
	if len(files) == 0 {
		return errors.New("no files specified")
	}

	for _, fn := range files {
		file, err := os.Open(fn)
		if err != nil {
			return err
		}

		img, _, err := image.Decode(file)
		if err != nil {
			return err
		}

		if img.Bounds().Empty() {
			return errors.New("empty image found")
		}

		p := img.Bounds().Size()

		ext := filepath.Ext(file.Name())
		n := file.Name()[0 : len(file.Name())-len(ext)]

		name := fmt.Sprintf("%s-%dx%d%s", n, p.X, p.Y, ext)
		if err := os.Rename(file.Name(), name); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
