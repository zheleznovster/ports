package managers

import (
	"fmt"
	"ports/internal/database"
	"ports/internal/parsers"
)

// Manager a db and a parser (and perhaps an http server)
type Manager struct {
	Db     *database.Database
	Parser parsers.Parser
}

// NewManager creates a parser to initialize a database from the contents of a json file with ports data
func NewManager(datapath string) (*Manager, error) {
	parser, err := parsers.NewParser(datapath)
	if err != nil {
		return nil, fmt.Errorf("NewManager() error: %w", err)
	}
	return &Manager{
		// empty database
		Db: database.NewDatabase(),
		// empty parser
		Parser: parser,
	}, nil
}

//nolint: forbidigo
// LoadData parses input file and fills the database with ports data
func (manager *Manager) LoadData() error {
	if manager == nil {
		return fmt.Errorf("LoadData() failed: Manager must be initialized first")
	}
	return manager.Parser.ParseFile(manager.SaveRecordToDb)
}

func (manager *Manager) SaveRecordToDb(key string, rec interface{}) error {
	// save each port record in database
	if _, ok := manager.Db.Get(key); !ok {
		manager.Db.Insert(key, rec)
	} else {
		manager.Db.Update(key, rec)
	}
	return nil
}
