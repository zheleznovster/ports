package managers

import (
	"encoding/json"
	"fmt"
	"os"
	"ports/database"
	"ports/parsers"
)

// Manager a db and a parser (and perhaps an http server)
type Manager struct {
	Db     *database.Database
	Parser *parsers.Parser
}

// NewManager creates a parser to initialize a database from the contents of a json file with ports data
func NewManager() *Manager {
	return &Manager{
		// empty database
		Db: database.NewDatabase(),
		// empty parser
		Parser: &parsers.Parser{},
	}
}

//nolint: forbidigo
// LoadData parses input file and fills the database with ports data
func (manager Manager) LoadData(path string) error {

	if err := manager.Parser.OpenFile(path); err != nil {
		return fmt.Errorf("LoadData error: %w", err)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("(manager *Manager) LoadData error: %w", err)
	}

	if fileInfo.Size() == 0 {
		return fmt.Errorf("(parser *Parser) LoadData error: json file must be non empty")
	}

	// Read the array open bracket
	var token any
	token, err = manager.Parser.Decoder.Token()
	if err != nil {
		return fmt.Errorf("(parser *Parser) LoadData error: %w", err)
	}
	fmt.Println(token)
	if token != json.Delim('{') {
		return fmt.Errorf("(parser *Parser) LoadData error: json must start with open curly brace {")
	}

	for {
		// parse next port record from json file
		portCode, portRecord, err := manager.Parser.ParseNextRecord()
		if err != nil {
			return fmt.Errorf("LoadData error: %w", err)
		}
		if len(portRecord) == 0 {
			break // found end of the file
		}

		// save port record in database
		if _, ok := manager.Db.Get(portCode); !ok {
			manager.Db.Insert(portCode, portRecord)
		} else {
			manager.Db.Update(portCode, portRecord)
		}

		fmt.Printf("Port code: %v\n", portCode)
		fmt.Printf("Port Data: %+v\n", portRecord)

	}
	fmt.Println("Data Loaded")
	return manager.Parser.CloseFile()
}
