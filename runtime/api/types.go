package api

import (
	"net/http"
	"strconv"
)

const HeightHeader = "X-Starport-Framework-Height"

func getHeight(r *http.Request) (uint64, error) {
	heightStr := r.Header.Get(HeightHeader)
	if heightStr == "" {
		return 0, nil
	}
	height, err := strconv.ParseUint(heightStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return height, nil
}

type Error struct {
}
