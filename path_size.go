package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getSize(path string, hidden bool, recursive bool) (int64, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !stat.IsDir() {
		return stat.Size(), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	var result int64
	for _, file := range files {
		isHidden := strings.HasPrefix(file.Name(), ".")

		if !file.IsDir() {
			if info, err := file.Info(); err == nil {
				if isHidden && !hidden {
					continue
				}

				result += info.Size()
			}
		}

		if file.IsDir() && recursive {
			dirSize, _ := getSize(filepath.Join(path, file.Name()), hidden, recursive)
			result += dirSize
		}
	}

	return result, nil
}

func GetPathSize(path string, recursive, human, hidden bool) (string, error) {
	res, err := getSize(path, hidden, recursive)
	if err != nil {
		return "", err
	}
	return FormatSize(res, human), nil
}

func FormatSize(size int64, human bool) string {
	dimns := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	res := float64(size)
	dimension := "B"
	for i := 0; i < len(dimns); i++ {
		if res < 1000 {
			dimension = dimns[i]
			break
		}
		res /= 1024
	}

	if human && dimension != "B" {
		return fmt.Sprintf("%.1f%s", res, dimension)
	}

	return fmt.Sprintf("%dB", size)
}
