package helper

import (
	"errors"
	"strconv"
)

func ConvertID(strID string) (int, error) {
	id, _ := strconv.Atoi(strID)
	if id < 1 {
		return 0, errors.New("param must be a number greater than 0")
	}

	return id, nil
}
