package keygen

import (
	"crypto/sha256"
	"os/exec"
	"runtime"
	"strings"
)

const KeyRegexp = "[[:alnum:]]{8}-(?:[[:alnum:]]{4}-){3}[[:alnum:]]{12}"

func GetKey() (string, error) {
	var (
		key []byte
		err error
	)

	switch runtime.GOOS {
	case "darwin":
		key, err = getMacOSKey()
	case "linux":
		key, err = getLinuxKey()
	}
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(key)
	return string(append(key[:4], hash[:]...)), err
}

func getMacOSKey() ([]byte, error) {
	cmd := "ioreg -d2 -c IOPlatformExpertDevice | awk -F\\\" '/IOPlatformUUID/{print $(NF-1)}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(string(out))), nil
}

func getLinuxKey() ([]byte, error) {
	// needed super-user mode
	cmd := "cat /sys/class/dmi/id/product_uuid"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err == nil {
		return []byte(strings.TrimSpace(string(out))), nil
	}

	// every other user
	cmd = "cat /var/lib/dbus/machine-id"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(string(out))), nil
}
