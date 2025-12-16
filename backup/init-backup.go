package backup

import (
	"archive/zip"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const BarWidth float64 = 67

func BackupFiles() {
	var (
		err            error
		rootPath       string
		targetPath     string
		targetFile     string
		zipFile        *os.File
		zipWriter      *zip.Writer
		info           os.FileInfo
		filled         int             = 0
		totalFiles     int             = 0
		completedFiles int             = 0
		percent        float64         = 0
		cache          map[string]bool = map[string]bool{}
		timestamp      string          = tools.GetFormatedTime()
	)
	if rootPath, err = os.Getwd(); err != nil {
		panic(err)
	}
	targetPath = filepath.Join(rootPath, "backup", "files")
	targetFile = filepath.Join(rootPath, "backup", "files", fmt.Sprintf("%s.zip", timestamp))
	if info, err = os.Stat(targetPath); info == nil || !info.IsDir() || err != nil {
		if err = os.MkdirAll(targetPath, 0777); err != nil {
			panic(err)
		}
	}
	if zipFile, err = os.Create(targetFile); err != nil {
		panic(err)
	}
	zipWriter = zip.NewWriter(zipFile)

	defer zipFile.Close()
	defer zipWriter.Close()

	var countTotalFileNumber = func(path string, info os.FileInfo, err error) error {
		var filePathStr []string
		for _, except := range conf.Env.BackUpExcepts {
			filePathStr = strings.Split(path, global.Backslash)
			if strings.Contains(strings.Join(filePathStr[0:len(filePathStr)-1], global.Backslash), except) {
				return nil
			}
		}
		if err == nil && !info.IsDir() {
			totalFiles++
		}
		return nil
	}
	var fileHanldeFunc = func(path string, info os.FileInfo, err error) error {
		var (
			innerErr    error
			relPath     string
			srcFile     *os.File
			filePathStr []string
			writer      io.Writer
			fileHeader  *zip.FileHeader
		)
		if cache[path] {
			return nil
		}
		cache[path] = true
		if err != nil {
			return err
		} else if relPath, err = filepath.Rel(rootPath, path); err != nil {
			innerErr = err
			return innerErr
		}
		for _, except := range conf.Env.BackUpExcepts {
			filePathStr = strings.Split(path, global.Backslash)
			if strings.Contains(strings.Join(filePathStr[0:len(filePathStr)-1], global.Backslash), except) || relPath == conf.Env.AppName {
				innerErr = nil
				return innerErr
			}
		}
		if info.IsDir() {
			innerErr = nil
			return innerErr
		} else if fileHeader, err = zip.FileInfoHeader(info); err != nil {
			innerErr = err
			return innerErr
		} else {
			fileHeader.Name = relPath
			fileHeader.Method = zip.Deflate
			if writer, err = zipWriter.CreateHeader(fileHeader); err != nil {
				innerErr = err
				return innerErr
			} else if srcFile, err = os.Open(path); err != nil {
				innerErr = err
				return innerErr
			} else {
				defer srcFile.Close()
				if _, err = io.Copy(writer, srcFile); err != nil {
					innerErr = err
					return innerErr
				}
				completedFiles++
				percent = float64(completedFiles) / float64(totalFiles) * 100
				filled = int((percent / 100) * BarWidth)
				fmt.Printf(
					ProcessBar,
					strings.Repeat(Processed, filled)+strings.Repeat(global.SpaceString, int(BarWidth)-filled),
					percent,
					completedFiles,
					totalFiles,
				)
				if percent == 100 {
					fmt.Printf("\n")
				}
				return nil
			}
		}
	}
	if err = filepath.Walk(rootPath, countTotalFileNumber); err != nil {
		panic(err)
	} else if totalFiles > 0 {
		fmt.Println(StartBackupComment)
	} else {
		fmt.Println(NoFileToBackupComment)
		return
	}
	if err = filepath.Walk(rootPath, fileHanldeFunc); err != nil {
		panic(err)
	}
	fmt.Println(BackupCompletedComment)
}
