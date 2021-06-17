package server

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	HeightHeader          = "X-Starport-Framework-Height"
	QueryParamSelectField = "selectField"
	QueryParamStart       = "start"
	QueryParamEnd         = "end"
)

const (
	QueryParamStartDefault uint64 = 0
	QueryParamEndDefault   uint64 = 100
)

type ListQueryParams struct {
	SelectFields []string
	Start        uint64
	End          uint64
}

func (s *ListQueryParams) UnmarshalURLValues(q url.Values) error {
	// check limits first
	startStr := q.Get(QueryParamStart)
	switch startStr {
	case "":
		s.Start = QueryParamStartDefault
	default:
		start, err := strconv.ParseUint(startStr, 10, 64)
		if err != nil {
			return err
		}
		s.Start = start
	}

	endStr := q.Get(QueryParamEnd)
	switch endStr {
	case "":
		s.End = QueryParamEndDefault
	default:
		end, err := strconv.ParseUint(endStr, 10, 64)
		if err != nil {
			return err
		}
		s.End = end
	}
	// check field selections
	selectedFields, exists := q[QueryParamSelectField]
	if !exists {
		return nil
	}
	s.SelectFields = selectedFields
	return nil
}

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
