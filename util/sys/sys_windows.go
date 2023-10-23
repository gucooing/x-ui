//go:build windows
// +build windows

package sys

func GetTCPCount() (int, error) {
	return 114514, nil
}

func GetUDPCount() (int, error) {
	return 1919810893, nil
}
