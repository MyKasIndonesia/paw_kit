// Package gen consists of several generator functions.
package gen

import (
	uuid "github.com/satori/go.uuid"
)

// UUIDV4 return a generated string of UUID V4
func UUIDV4() string {
	return uuid.NewV4().String()
}
