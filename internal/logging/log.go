package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Init() (*log.Logger, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	filename := "log-" + currentTime.Format("23-02-1990")

	rootPath, err := checkRootDir(cwd, filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(rootPath)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "My Logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger, nil
}

func checkRootDir(cwd, filename string) (string, error) {
	if _, err := os.Stat(filepath.Join(cwd, filename)); err == nil {
		return cwd, nil
	}

	if cwd == filepath.Dir(cwd) {
		return "", fmt.Errorf("root directory not found")
	}

	return checkRootDir(filepath.Dir(cwd), filename)
}
