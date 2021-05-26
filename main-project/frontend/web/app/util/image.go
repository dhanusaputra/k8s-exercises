package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadImageToVolume ...
func DownloadImageToVolume(url string, path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	f.Sync()

	return nil
}

// ReadImageFromVolume ...
func ReadImageFromVolume(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
