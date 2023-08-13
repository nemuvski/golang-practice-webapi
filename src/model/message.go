package model

import (
	"strconv"

	"github.com/cockroachdb/errors"
)

type Message struct {
	ID			uint8		`json:"id"`
	Content	string	`json:"content"`
}

func ParseMessageID(id string) (uint8, error) {
	parsedID, err := strconv.ParseUint(id, 10, 8)
	if err != nil {
		return 0, errors.Wrap(err, "strconv.ParseUint")
	}
	return uint8(parsedID), nil
}
