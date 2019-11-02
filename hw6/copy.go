package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"

	"github.com/cheggaaa/pb/v3"
)

func readSize(filepath string, offset int64, limit int64) (int64, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	filesize := info.Size()

	if filesize == 0 {
		return 0, fmt.Errorf("could not determine file size")
	}
	if filesize <= offset {
		return 0, fmt.Errorf("offset more or equal than file size")
	}

	readsize := filesize - offset
	if limit == READ_ALL_THE__REST_OF_THE_FILE {
		return readsize, nil
	}
	if readsize > limit {
		return limit, nil
	}
	return readsize, nil
}

func copySection(src io.ReaderAt, dest io.Writer, offset int64, readsize int64) error {
	reader := io.NewSectionReader(src, offset, readsize)

	progress := pb.Full.Start64(readsize)
	progressReader := progress.NewProxyReader(reader)
	_, err := io.Copy(dest, progressReader)
	if err != nil {
		return err
	}
	progress.Finish()

	return nil
}

func Copy(from string, to string, limit int64, offset int64) error {
	src, err := os.Open(from)
	defer src.Close()
	if err != nil {
		return err
	}

	readsize, err := readSize(from, offset, limit)
	if err != nil {
		return err
	}

	dest, err := os.Create(to)
	defer dest.Close()
	if err != nil {
		return err
	}

	if err := copySection(src, dest, offset, readsize); err != nil {
		if os.Remove(to) != nil {
			return errors.Wrap(err, "can not remove failed dest file")
		}
		return err
	}
	return nil
}
