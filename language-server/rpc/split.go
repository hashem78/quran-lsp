package rpc

import (
	"errors"
	"strconv"
	"strings"
)

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	dataAsString := string(data)
	header, content, found := strings.Cut(dataAsString, "\r\n\r\n")
	if !found {
		return 0, nil, nil
	}

	_, contentLengthStr, found := strings.Cut(header, "Content-Length: ")

	if !found {
		return 0, nil, errors.New("Unable to find content length")
	}

	contentLength, err := strconv.Atoi(contentLengthStr)

	if err != nil {
		return 0, nil, err
	}

	totalLength := contentLength + 4 + len(header)

	return totalLength, []byte(content), nil
}
