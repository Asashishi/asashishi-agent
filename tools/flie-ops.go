package tools

import (
	"asashishi-agent/conf"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var rollBackMap = map[string]string{}

func listFiles(path string) []string {
	var (
		err     error
		entires []os.DirEntry
		paths   []string = []string{}
	)
	if entires, err = os.ReadDir(path); err != nil {
		panic(err)
	}
	for _, entry := range entires {
		if !entry.IsDir() {
			paths = append(paths, filepath.Join(path, entry.Name()))
		} else {
			var next bool = false
			for _, name := range conf.Env.FilesExcepts {
				if strings.Contains(entry.Name(), name) {
					next = true
					break
				}
			}
			if next {
				continue
			}
			paths = append(paths, filepath.Join(path, entry.Name()))
			paths = append(paths, listFiles(filepath.Join(path, entry.Name()))...)
		}
	}
	return paths
}

func removeFileCacheForDir(path string) {
	var (
		err     error
		entries []os.DirEntry
	)
	if entries, err = os.ReadDir(path); err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			removeFileCacheForDir(filepath.Join(path, entry.Name()))
		} else {
			delete(rollBackMap, filepath.Join(path, entry.Name()))
		}
	}
}

func GetFileList(path string) []string {
	var (
		err     error
		dirPath string
	)
	if path == "" {
		if dirPath, err = os.Getwd(); err != nil {
			return []string{}
		}
		return listFiles(dirPath)
	} else {
		return listFiles(path)
	}
}

func MoveContent(oPath string, nPath string) bool {
	var err error
	if err = os.Rename(oPath, nPath); err != nil {
		return false
	}
	return true
}

func CreateDir(path string) bool {
	var err error
	if err = os.Mkdir(path, 0755); err != nil {
		return false
	}
	return true
}

func RemoveDir(path string) bool {
	var (
		err error
	)
	removeFileCacheForDir(path)
	if err = os.RemoveAll(path); err != nil {
		return false
	}
	return true
}

func CreateFile(path string) bool {
	var (
		err  error
		file *os.File
	)
	if file, err = os.Create(path); err != nil {
		return false
	}
	defer file.Close()
	return true
}

func RemoveFile(path string) bool {
	var err error
	if err = os.Remove(path); err != nil {
		return false
	}
	delete(rollBackMap, path)
	return true
}

func ReadFileContent(path string) string {
	var (
		err      error
		data     []byte
		contents string = ""
	)
	if data, err = os.ReadFile(path); err != nil {
		return contents
	}

	contents = string(data)
	if _, ok := rollBackMap[path]; !ok {
		rollBackMap[path] = contents
	}
	return contents
}

func AppendContentAtTail(path string, content string) bool {
	var (
		err         error
		data        []byte
		newContents string
	)
	if data, err = os.ReadFile(path); err != nil {
		return false
	}
	newContents = string(data) + "\n" + content
	if err = os.WriteFile(path, []byte(newContents), 0644); err != nil {
		return false
	}
	return true
}

func SearchFileContent(path string, content string) [][]int {
	var (
		err       error
		data      []byte
		regex     *regexp.Regexp
		positions [][]int = [][]int{}
	)
	if data, err = os.ReadFile(path); err != nil {
		return positions
	}
	regex = regexp.MustCompile(regexp.QuoteMeta(content))
	positions = regex.FindAllStringIndex(string(data), -1)
	return positions
}

func ReplaceFileContentByPosition(path string, position []int, content string) bool {
	var (
		err        error
		data       []byte
		dataStr    string
		strBulider strings.Builder
	)
	if data, err = os.ReadFile(path); err != nil {
		return false
	}
	dataStr = string(data)
	strBulider.WriteString(dataStr[:position[0]])
	strBulider.WriteString(content)
	strBulider.WriteString(dataStr[position[1]:])
	if err = os.WriteFile(path, []byte(strBulider.String()), 0644); err != nil {
		return false
	}
	return true
}

func DeleteFileContentByPosition(path string, position []int) bool {
	var (
		err        error
		data       []byte
		dataStr    string
		strBulider strings.Builder
	)
	if data, err = os.ReadFile(path); err != nil {
		return false
	}
	dataStr = string(data)
	strBulider.WriteString(dataStr[:position[0]])
	strBulider.WriteString(dataStr[position[1]:])
	if err = os.WriteFile(path, []byte(strBulider.String()), 0644); err != nil {
		return false
	}
	return true
}

func DeleteFileContent(path string) bool {
	var (
		err  error
		file *os.File
	)
	if file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
		return false
	}
	defer file.Close()
	if _, err := file.WriteString(""); err != nil {
		return false
	}
	return true
}

func RenewFileCache(path string) bool {
	var (
		err  error
		data []byte
	)
	if data, err = os.ReadFile(path); err != nil {
		return false
	}
	rollBackMap[path] = string(data)
	return true
}

func FileContentRollBack(path string) bool {
	var (
		err  error
		file *os.File
	)
	if file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		return false
	}
	defer file.Close()
	if _, err := file.WriteString(rollBackMap[path]); err != nil {
		return false
	}
	return true
}
