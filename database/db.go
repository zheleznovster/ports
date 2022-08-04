package database

import (
	"sync"
)

type Database struct {
	Data map[string]interface{}
	Mtx  sync.RWMutex
}

// NewDatabase initializes a database and the underlying map
func NewDatabase() *Database {
	return &Database{
		Data: make(map[string]interface{}, 1000),
		Mtx:  sync.RWMutex{},
	}
}

// Get returns the value referenced by the key and true, if the key is found
// does nothing and returns nil and false, if the key is not found
// does nothing and returns nil and false if key is empty
func (c *Database) Get(key string) (interface{}, bool) {
	if key == "" {
		return nil, false
	}

	if c.Data == nil {
		return nil, false
	}

	c.Mtx.Lock()
	defer c.Mtx.Unlock()

	value, ok := c.Data[key]

	return value, ok
}

// Insert inserts a record only if it doesn't exist
// does nothing and returns false if the record already exists
// does nothing and returns false if key is empty
// allows empty values
func (c *Database) Insert(key string, value interface{}) bool {
	if key == "" {
		return false
	}

	if c.Data == nil {
		return false
	}

	c.Mtx.Lock()
	defer c.Mtx.Unlock()

	if _, ok := c.Data[key]; ok {
		return false
	}

	c.Data[key] = value

	return true
}

// Update updates a record and returns true only if it exists
// does not update a record and returns false if the record doesn't already exist
// does nothing and returns false if key is empty
func (c *Database) Update(key string, value interface{}) bool {
	if key == "" {
		return false
	}

	if c.Data == nil {
		return false
	}

	c.Mtx.Lock()
	defer c.Mtx.Unlock()

	if _, ok := c.Data[key]; !ok {
		return false
	}

	c.Data[key] = value

	return true
}

// Delete a record and returns true only if it exists
// does nothing and returns false if the record doesn't already exist
// does nothing and returns false if key is empty
func (c *Database) Delete(key string) bool {
	if key == "" {
		return false
	}

	if c.Data == nil {
		return false
	}

	c.Mtx.Lock()
	defer c.Mtx.Unlock()

	// before deleting, make sure the element exists
	// returning false allows the user to maybe stop trying to delete something that doesn't exist
	if _, ok := c.Data[key]; !ok {
		return false
	}

	delete(c.Data, key)

	return true
}
