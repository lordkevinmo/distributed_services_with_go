package server

import (
	"fmt"
	"sync"
)

// Log is the data structure of the Log
type Log struct {
	mu      sync.Mutex
	records []Record
}

// NewLog is reference of the Log with DI
func NewLog() *Log {
	return &Log{}
}

// Append take record as a parameter and append it to the log.
// The function return the record offset and the error if it exists
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read take an offset and return an Record and error if exists
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}
	return c.records[offset], nil
}

// Record is the Data structure responsible of write the Log
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// ErrorOffsetNotFound is the error type
var ErrorOffsetNotFound = fmt.Errorf("offset not found")
