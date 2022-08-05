package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type LargeFileParser struct {
	FileParser
	FilePointer *os.File
	Decoder     *json.Decoder
}

//nolint: forbidigo
func (parser *LargeFileParser) ParseFile(processRecord func(key string, rec interface{}) error) error {
	if err := parser.openFile(parser.Path); err != nil {
		return fmt.Errorf("(parser *LargeFileParser) ParseFile error: %w", err)
	}
	defer func() {
		err := parser.closeFile()
		if err != nil {
			fmt.Println(fmt.Errorf("(parser *LargeFileParser) ParseFile Error: %w", err))
		}
	}()

	// Read the array open bracket
	token, err := parser.Decoder.Token()
	if err != nil {
		return fmt.Errorf("(parser *LargeFileParser) ParseFile error: %w", err)
	}
	fmt.Println(token)
	if token != json.Delim('{') {
		return fmt.Errorf("(parser *LargeFileParser) ParseFile error: json must start with open curly brace {")
	}

	for {
		// parse next port record from json file
		portCode, portRecord, err := parser.parseNextRecord()
		if err != nil {
			return fmt.Errorf("ParseFile error: %w", err)
		}
		if len(portRecord) == 0 {
			break // found end of the file
		}

		err = processRecord(portCode, portRecord)
		if err != nil {
			return fmt.Errorf("(parser *LargeFileParser) ParseFile error: %w", err)
		}
		fmt.Printf("Port code: %v\n", portCode)
		fmt.Printf("Port Data: %+v\n", portRecord)

	}
	fmt.Println("(parser *LargeFileParser) ParseFile succeeded")
	return nil
}

// openFile opens a given json file and initializes a json decoder
func (parser *LargeFileParser) openFile(path string) error {
	var err error
	parser.FilePointer, err = os.Open(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("(parser *LargeFileParser) openFile error: %w", err)
	}

	parser.Decoder = json.NewDecoder(parser.FilePointer)

	return nil
}

// parseNextRecord reads the port code and the json record for a single port
func (parser *LargeFileParser) parseNextRecord() (string, map[string]interface{}, error) {

	if !parser.Decoder.More() {
		return "", nil, nil
	}

	// Read the port code
	portCode, err := parser.readPortCode()
	if err != nil {
		return "", nil, fmt.Errorf("parseNextRecord error: %w", err)
	}

	// Read the port record
	data := map[string]interface{}{}
	err = parser.Decoder.Decode(&data)
	if err != nil {
		return "", nil, fmt.Errorf("parseNextRecord error: %w", err)
	}
	return portCode, data, nil
}

func (parser *LargeFileParser) closeFile() error {
	return parser.FilePointer.Close()
}

var (
	openingBrace   = json.Delim('{')
	closingBrace   = json.Delim('}')
	openingBracket = json.Delim('[')
	closingBracket = json.Delim(']')
)

func (parser *LargeFileParser) readPortCode() (string, error) {
	var token, portCode interface{}
	var err error
	token, err = parser.Decoder.Token()
	if err != nil {
		return "", fmt.Errorf("readPortCode error: %w", err)
	}

	switch token {
	case openingBracket:
		return "", fmt.Errorf("readPortCode error: encountered an opening bracket")
	case openingBrace:
		// Read the port code
		portCode, err = parser.Decoder.Token()
		if err != nil {
			return "", fmt.Errorf("readPortCode error: %w", err)
		}
		switch portCode {
		case closingBrace, closingBracket:
			return "", fmt.Errorf("readPortCode error: %v", "invalid port code")
		}

	default:
		portCode = token
	}

	return portCode.(string), nil
}
