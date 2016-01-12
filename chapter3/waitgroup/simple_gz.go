package main

import (
	"compress/gzip"
	"io"
	"os"
	//"sync"
)

func main() {
	for _, file := range os.Args {
		compress(file)
	}
	//var wg sync.WaitGroup
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
