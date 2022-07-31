package parsers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Parser struct {
	FilePointer *os.File
	Decoder     *json.Decoder
}

// OpenFile opens a given json file and initializes a json decoder
func (parser *Parser) OpenFile(path string) error {
	var err error
	parser.FilePointer, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("(parser *Parser) OpenFile error: %w", err)
	}

	parser.Decoder = json.NewDecoder(parser.FilePointer)

	return nil
}

// ParseNextRecord reads the port code and the json record for a single port
func (parser *Parser) ParseNextRecord() (string, map[string]interface{}, error) {

	if !parser.Decoder.More() {
		return "", nil, nil
	}

	// Read the port code
	portCode, err := parser.readPortCode()
	if err != nil {
		return "", nil, fmt.Errorf("ParseNextRecord error: %w", err)
	}

	// Read the port record
	data := map[string]interface{}{}
	err = parser.Decoder.Decode(&data)
	if err != nil {
		return "", nil, fmt.Errorf("ParseNextRecord error: %w", err)
	}
	return portCode, data, nil
}

func (parser *Parser) CloseFile() error {
	return parser.FilePointer.Close()
}

var (
	openingBrace   = json.Delim('{')
	closingBrace   = json.Delim('}')
	openingBracket = json.Delim('[')
	closingBracket = json.Delim(']')
)

func (parser *Parser) readPortCode() (string, error) {
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
