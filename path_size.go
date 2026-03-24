package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getFileSize(path string) (int, error) {
	finfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	return int(finfo.Size()), nil
}

func getDirFilesSize(path string, hidden bool, recursive bool) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, file := range files {
		isHidden := strings.HasPrefix(file.Name(), ".")

		if !file.IsDir() {
			if info, err := file.Info(); err == nil {
				if isHidden && !hidden {
					continue
				}

				result += int(info.Size())
			}
		}

		if file.IsDir() && recursive {
			dirPath := filepath.Join(path, file.Name())
			dirSize, err := getDirFilesSize(dirPath, hidden, recursive)
			if err != nil {
				return 0, err
			}
			result += dirSize
		}
	}

	return result, nil
}

func GetPathSize(path string, recursive, human, hidden bool) (string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if stat.IsDir() {
		res, err := getDirFilesSize(path, hidden, recursive)
		if err != nil {
			return "", err
		}
		return FormatSize(res, human), nil
	}

	res, err := getFileSize(path)
	if err != nil {
		return "", err
	}
	return FormatSize(res, human), nil
}

func FormatSize(size int, human bool) string {
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
