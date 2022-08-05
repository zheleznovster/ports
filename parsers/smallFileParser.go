package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"ports/models"
)

type SmallFileParser struct {
	FileParser
}

func (parser *SmallFileParser) ParseFile(processRecord func(key string, rec interface{}) error) error {
	file, err := os.Open(parser.Path)
	if err != nil {
		return fmt.Errorf("ParseFile error: %w", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("ParseFile Error: %w", err))
		}
	}()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ParseFile error: %w", err)
	}

	inputMap := make(map[string]models.Port)
	err = json.Unmarshal(data, &inputMap)
	if err != nil {
		return fmt.Errorf("ParseFile error: %w", err)
	}
	for portCode, portRecord := range inputMap {
		err := processRecord(portCode, portRecord)
		if err != nil {
			return fmt.Errorf("ParseFile error: %w", err)
		}
		fmt.Printf("Port code: %v\n", portCode)
		fmt.Printf("Port Data: %+v\n", portRecord)
	}
	fmt.Println("SmallFileParser Loaded Data")
	return nil
}
