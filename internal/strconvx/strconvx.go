package strconvx

import (
	"fmt"
	"strconv"
	"strings"
)

func StrTimeToFloat64(t string) (float64, error) {
	strSplit := strings.Split(t, " ")
	if len(strSplit) == 1 {
		return strconv.ParseFloat(strSplit[0], 64)
		//return strconv.ParseInt(strSplit[0], 10, 64)
	}

	if len(strSplit) == 2 {
		f64, err := strconv.ParseFloat(strSplit[0], 64)
		if err != nil {
			return 0, err
		}

		if strSplit[1] == "ns/op" {
			return f64, nil
		}

		if strSplit[1] == "µs/op" {
			return f64 * 1_000, nil
		}

		if strSplit[1] == "ms/op" {
			return f64 * 1_000_000, nil
		}

		if strSplit[1] == "s/op" {
			return f64 * 1_000_000_000, nil
		}

		return f64, nil
	}

	return 0, fmt.Errorf("invalid format")
}

func StrSizeToFloat64(t string) (float64, error) {
	strSplit := strings.Split(t, " ")
	if len(strSplit) == 1 {
		return strconv.ParseFloat(strSplit[0], 64)
		//return strconv.ParseInt(strSplit[0], 10, 64)
	}

	if len(strSplit) == 2 {
		f64, err := strconv.ParseFloat(strSplit[0], 64)
		if err != nil {
			return 0, err
		}

		if strSplit[1] == "B/op" {
			return f64, nil
		}

		if strSplit[1] == "KB/op" {
			return f64 * 1_000, nil
		}

		if strSplit[1] == "MB/op" {
			return f64 * 1_000_000, nil
		}

		if strSplit[1] == "GB/op" {
			return f64 * 1_000_000_000, nil
		}

		return f64, nil
	}

	return 0, fmt.Errorf("invalid format")
}

func StrAllocCountToFloat64(t string) (float64, error) {
	strSplit := strings.Split(t, " ")
	f64, err := strconv.ParseFloat(strSplit[0], 64)
	if err != nil {
		return 0, err
	}

	return f64, nil
}
