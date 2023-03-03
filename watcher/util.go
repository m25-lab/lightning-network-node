package watcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

func ParseEventDataNewBlock(block coretypes.ResultEvent) (*types.EventDataNewBlock, error) {
	byteData, err := json.Marshal(block.Data)
	if err != nil {
		return nil, err
	}

	ret := &types.EventDataNewBlock{}
	if err := json.Unmarshal(byteData, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func ReadLastLine(file *os.File) (string, error) {
	var cursor int64 = 0
	var line string = ""

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := stat.Size()
	if fileSize == 0 {
		return "", nil
	}

	for {
		cursor -= 1
		file.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		file.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) {
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line)

		if cursor == -fileSize {
			break
		}
	}

	return line, nil
}

func isFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if _, err := os.Create(pathLogFile); err != nil {
				return false, nil
			}
		} else {
			return false, err
		}
	}

	return true, nil
}

func loadHistory(pathLogFile string) (int64, error) {
	fileExists, err := isFileExists(pathLogFile)
	if err != nil {
		return 0, err
	}

	if !fileExists {
		return -1, nil
	}

	file, err := os.Open(pathLogFile)
	if err != nil {
		return 0, err
	}

	line, err := ReadLastLine(file)
	if err != nil {
		return 0, err
	}

	if line == "" {
		return -1, nil
	}

	var temp map[string]string
	err = json.Unmarshal([]byte(line), &temp)
	if err != nil {
		return 0, err
	}

	height, err := strconv.ParseInt(strings.Split(temp["msg"], " ")[2], 10, 64)
	if err != nil {
		return 0, err
	}

	return height, nil
}

func createFileLogger(pathLogFile string) (*log.Logger, error) {
	fileExists, err := isFileExists(pathLogFile)
	if err != nil {
		return nil, err
	}

	if !fileExists {
		if _, err := os.Create(pathLogFile); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(pathLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(file)

	return logger, nil

}
