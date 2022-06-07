package linuxid

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func pickNumbersFrom(s string, howMany uint) string {
	isNumber := func(letter string) bool {
		numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		for _, number := range numbers {
			if letter == number {
				return true
			}
		}

		return false
	}

	var j int
	var res []string

	for i := len(s); i >= 0; i-- {
		if isNumber(s[i-1 : i]) {
			res = append(res, s[i-1:i])

			j++
		}

		if j == int(howMany) {
			return strings.Join(res, "")
		}
	}

	return ""
}

// randInt generates a random uint32 as 4 digit string
// time cost is around 500 ns.
func randInt() (string, error) {
	buf := make([]byte, 3)

	if _, err := rand.Reader.Read(buf); err != nil {
		return "", fmt.Errorf("generate random number: %W;", err)
	}

	res := uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])

	return strconv.Itoa(int(res)), nil
}

// readMachineID returns machine ID
func readContainerID() (string, error) {
	idHardware, errID := ioutil.ReadFile("/etc/machine-id")
	if errID != nil || len(idHardware) == 0 {
		hostname, errHo := os.Hostname()
		if errHo != nil {
			return "", fmt.Errorf("readContainerID os.Hostname: %w", errHo)
		}

		return hostname, nil
	}

	return string(idHardware), nil
}
