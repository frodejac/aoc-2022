package io

import (
	"log"
	"os"
)

func GetInput(day int) ([]byte, error) {
	log.Printf("Getting input for day %d from %s\n", day, getInputAssetName(day))
	data, err := os.ReadFile(getInputAssetName(day))

	if err != nil {
		return nil, err
	}
	return data, nil
}
