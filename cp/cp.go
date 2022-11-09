package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var src string
var dst string
var offset int64
var limit int64

const (
	pageSize = 1000
)

func init() {
	flag.StringVar(&src, "from", "", "path to source file")
	flag.StringVar(&dst, "to", "", "path to destination file")
	flag.Int64Var(&offset, "offset", 0, "offset in source file")
	flag.Int64Var(&limit, "limit", -1, "max number of bytes to copy")
}

type progress struct {
	current int64
	total   int64
}

func printProgress(in <-chan progress, errs <-chan error, errsOut chan<- error) {
	for {
		select {
		case p := <-in:
			fmt.Printf("\033[2K\r%d/%d", p.current, p.total)
		case err := <-errs:
			errsOut <- err
			return
		}
	}
}

func copy(src string, dst string, offset int64, limit int64, in chan<- progress, errs chan<- error) {
	srcFile, err := os.Open(src)
	if err != nil {
		errs <- err
		return
	}
	fi, err := srcFile.Stat()
	if err != nil {
		errs <- err
		return
	}
	if fi.IsDir() {
		err = fmt.Errorf("dir copy not supported")
		errs <- err
		return
	}

	total := fi.Size()

	currentOffset, err := srcFile.Seek(offset, io.SeekStart)
	if err != nil {
		errs <- err
		return
	}

	total -= currentOffset
	if total > limit && limit > -1 {
		total = limit
	}

	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		errs <- err
		return
	}

	var written int64
	var read int64

	for read < total {
		written, err = io.CopyN(dstFile, srcFile, pageSize)
		if err != nil && err != io.EOF {
			errs <- err
			return
		}

		read += written
		in <- progress{
			current: read,
			total:   total,
		}
	}

	errs <- nil
}

func Copy(src string, dst string, offset int64, limit int64) error {
	progs := make(chan progress)
	errs := make(chan error)
	errsMain := make(chan error)

	go printProgress(progs, errs, errsMain)
	go copy(src, dst, offset, limit, progs, errs)

	return <-errsMain
}

func main() {
	flag.Parse()
	err := Copy(src, dst, offset, limit)
	if err != nil {
		fmt.Printf("\033[2K\r%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("\033[2K\rDone\n")
}
