package io

import (
	"os"
)

func GetInput(day int) ([]byte, error) {
	data, err := os.ReadFile(getInputAssetName(day))

	if err != nil {
		return nil, err
	}
	return data, nil
}
