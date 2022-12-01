package io

import (
	"os"
)

func GetSolution(day int, part int) ([]byte, error) {
	data, err := os.ReadFile(getSolutionAssetName(day, part))

	if err != nil {
		return nil, err
	}
	return data, nil
}
