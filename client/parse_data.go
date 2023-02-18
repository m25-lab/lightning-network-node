package client

import (
	"errors"
	"strings"
)

func ParseCallbackData(data string) (string, string, error) {
	splitedData := strings.Split(data, ":")
	if len(splitedData) != 2 {
		return "", "", errors.New("invalid data")
	}

	return splitedData[0], splitedData[1], nil
}
