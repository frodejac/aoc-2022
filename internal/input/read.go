package input

import (
	"fmt"
	"os"
)

const ASSET_DIR = "assets"

func GetInput(day int) ([]byte, error) {
	data, err := os.ReadFile(getAssetName(day))

	if err != nil {
		return nil, err
	}
	return data, nil
}

func getAssetName(day int) string {
	return fmt.Sprintf("%s/day%02d/input.txt", ASSET_DIR, day)
}
