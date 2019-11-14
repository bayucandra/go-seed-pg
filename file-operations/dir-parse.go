package file_operations

import (
	"fmt"
	"os"
	"sort"
)

func DirParse( dirPath string ) (filePath []string, err error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	_ = f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	for _, val := range files {
		filePath = append( filePath, fmt.Sprintf("%s/%s", dirPath, val.Name()) )
	}

	return
}
