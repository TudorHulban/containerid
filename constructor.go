package linuxid

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var _containerID string

func getContainerID() string {
	if len(_containerID) == 0 {
		id, err := readContainerID()
		if err != nil {
			return ""
		}

		_containerID = pickNumbersFrom(id, 5)
	}

	return _containerID
}

func NewID() (uint, error) {
	containerID := getContainerID()
	if len(containerID) == 0 {
		return 0, errors.New("there was an error")
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	random, errRa := randInt()
	if errRa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errRa)
	}

	parsed, errPa := strconv.ParseUint((now[:11] + containerID + random + now[16:19])[:20], 10, 64)
	if errPa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errPa)
	}

	return uint(parsed), nil
}
