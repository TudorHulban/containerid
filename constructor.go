package linuxid

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type ID uint

var _pid = os.Getpid()
var _machineID string

func getMachineID() string {
	if len(_machineID) == 0 {
		id, err := readMachineID()
		if err != nil {
			return ""
		}

		_machineID = id
	}

	return _machineID
}

func NewID() (ID, error) {
	machineID := getMachineID()
	if len(machineID) == 0 {
		return 0, errors.New("there was an error")
	}

	time := strconv.FormatInt(time.Now().Unix(), 10)
	pid := strconv.Itoa(_pid)

	random, errRa := randInt()
	if errRa != nil {
		return 0, fmt.Errorf("NewID randInt: %w", errRa)
	}

	rand := strconv.Itoa(int(random))[:4]

	parsed, errPa := strconv.ParseUint(time+pickNumbersFrom(machineID, 3)+pid[len(pid)-3:]+rand, 10, 64)
	if errPa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errPa)
	}

	res := ID(parsed)

	return res, nil
}
