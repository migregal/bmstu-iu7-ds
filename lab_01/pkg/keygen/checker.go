package keygen

func CheckKey(key string) (bool, error) {
	machineKey, err := GetKey()
	if err != nil {
		return false, err
	}

	return machineKey == key, nil
}
