package parsers

import (
	"fmt"
	"os"
)

const (
	FileSizeThreshold = 1048576 // megabyte
)

type FileParser struct {
	Path string
}

type Parser interface {
	ParseFile(processRecord func(key string, rec interface{}) error) error
}

func NewParser(path string) (Parser, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("NewParser error: %w", err)
	}
	size := fileInfo.Size()
	switch {
	case size == 0:
		return nil, fmt.Errorf("NewParser error: json file must be non empty")
	case size <= FileSizeThreshold:
		return &SmallFileParser{FileParser{path}}, nil
	case size > FileSizeThreshold:
		return &LargeFileParser{FileParser: FileParser{path}}, nil
	}

	return nil, nil
}
