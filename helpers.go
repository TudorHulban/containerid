package linuxid

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func pickNumbersFrom(s string, howMany uint) string {
	isNumber := func(letter rune) bool {
		if int(letter) >= 48 && int(letter) <= 57 {
			return true
		}

		return false
	}

	var j int
	var res []string

	for _, char := range s {
		if isNumber(char) {
			res = append(res, string(char))

			j++
		}

		if j == int(howMany) {
			return strings.Join(res, "")
		}
	}

	return ""
}

// randInt generates a random uint32
func randInt() (uint32, error) {
	buf := make([]byte, 3)

	if _, err := rand.Reader.Read(buf); err != nil {
		return 0, fmt.Errorf("generate random number: %W;", err)
	}

	return uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2]), nil
}

func readLinuxMachineID() (string, error) {
	buf, err := ioutil.ReadFile("/etc/machine-id")
	if err != nil || len(buf) == 0 {
		buf, err = ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	}

	return string(buf), err
}

// readMachineID returns machine ID
func readMachineID() (string, error) {
	idHardware, errID := readLinuxMachineID()
	if errID != nil || len(idHardware) == 0 {
		hostname, errHo := os.Hostname()
		if errHo != nil {
			return "", fmt.Errorf("readMachineID os.Hostname: %w", errHo)
		}

		idHardware = hostname
	}

	return idHardware, nil
}
