package utils

import (
	"errors"
	"strings"
)

var ErrResourceFormat = errors.New("invalid resource request")

func CpuResource(resource string) (uint64, error) {
	var token, n = resource, len(resource)
	if strings.HasSuffix(resource, "m") {
		token = resource[0 : n-1]
	}

	var request int
	for _, ch := range token {
		if ch >= '0' && ch <= '9' {
			request = (request * 10) + int(ch-'0')
		} else {
			return 0, ErrResourceFormat
		}
	}
	return uint64(request), nil
}

func MemoryResource(resource string) (uint64, error) {
	n := len(resource)
	var token, gauge, request = "", "", 0
	if resource[n-1] == 'i' {
		token = resource[0 : n-2]
		gauge = resource[n-2:]
	} else {
		token = resource[0 : n-1]
		gauge = resource[n-1:]
	}
	for _, ch := range token {
		if ch >= '0' && ch <= '9' {
			request = (request * 10) + int(ch-'0')
		} else {
			return 0, ErrResourceFormat
		}
	}
	return uint64(request) * Legitimate[gauge], nil
}

func LegalMemoryGauge(gauge string) bool {
	if _, ok := Legitimate[gauge]; ok {
		return true
	}
	return false
}

const (
	kiloBytes = 1024
	megaBytes = 1024 * kiloBytes
	gigaBytes = 1024 * megaBytes
	teraBytes = 1024 * gigaBytes
	petaBytes = 1024 * teraBytes
	exaBytes  = 1024 * petaBytes
)

var Legitimate = map[string]uint64{
	"Ei": exaBytes, "Pi": petaBytes,
	"Ti": teraBytes, "Gi": gigaBytes,
	"Mi": megaBytes, "Ki": kiloBytes,
	"E": exaBytes, "P": petaBytes,
	"T": teraBytes, "G": gigaBytes,
	"M": megaBytes, "k": kiloBytes,
}

func Convert(gauge string) uint64 {
	if _, ok := Legitimate[gauge]; ok {
		return Legitimate[gauge]
	}
	return 0
}
