package path_size

import (
	"fmt"
	"os"
)

func getFileSize(path string) (int, error) {
	finfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	return int(finfo.Size()), nil
}

func getDirFilesSize(path string) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, file := range files {
		if !file.IsDir() {
			if info, err := file.Info(); err == nil {
				result += int(info.Size())
			}
		}
	}

	return result, nil
}

func GetSize(path string) (int, error) {
	fmt.Println(path)
	fmt.Println("_________")
	fmt.Println("")

	stat, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if stat.IsDir() {
		return getDirFilesSize(path)
	}
	// fmt.Println(os.Lstat(path))
	return getFileSize(path)
}
