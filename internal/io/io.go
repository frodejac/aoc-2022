package io

import (
	"fmt"
)

const ASSET_DIR = "assets"

func getInputAssetName(day int) string {
	return fmt.Sprintf("%s/day%02d/input.txt", ASSET_DIR, day)
}

func getSolutionAssetName(day int, part int) string {
	return fmt.Sprintf("%s/day%02d/part%d.txt", ASSET_DIR, day, part)
}
