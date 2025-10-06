package fs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrWrittenZeroBytes = errors.New("written 0 bytes.")
)

func Write(path, name string, r io.Reader) error {
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s", path, name),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}

	written, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	if written == 0 {
		return ErrWrittenZeroBytes
	}

	return nil
}

func ReadString(path string) (string, error) {
	b := new(bytes.Buffer)
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	written, err := io.Copy(b, f)
	if err != nil {
		return "", err
	}
	if written == 0 {
		return "", ErrWrittenZeroBytes
	}

	return b.String(), nil
}

func Descriptor(path, name string) (*os.File, error) {
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s", path, name),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return nil, err
	}
	return f, nil
}
