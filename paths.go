package main

import (
	"os"
	"path/filepath"
)

const (
	AppName  = "middleman"
	CertName = "ca.crt"
	KeyName  = "ca.key"
)

func createDirsOrPanic(dirname string) {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		panic(err)
	}

}

var getAppConfigDir = func() string {
	dirName, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appConfigDir := filepath.Join(dirName, AppName)
	return appConfigDir
}

func getCertKeyPath() (string, string) {
	appConfigDir := getAppConfigDir()
	certPath, keyPath := filepath.Join(appConfigDir, CertName), filepath.Join(appConfigDir, KeyName)
	err := os.MkdirAll(filepath.Dir(certPath), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(keyPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	return certPath, keyPath
}

func getConfigPath() string {
	appConfigDir := getAppConfigDir()
	configPath := filepath.Join(appConfigDir, "config.json")
	os.MkdirAll(filepath.Dir(configPath), os.ModePerm)
	return configPath
}

func getLogFilePath() string {
	appConfigDir := getAppConfigDir()
	logFilePath := filepath.Join(appConfigDir, "log.txt")
	err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	return logFilePath
}

func getDatabasePath() string {
	appConfigDir := getAppConfigDir()
	databasePath := filepath.Join(appConfigDir, "database.json")
	err := os.MkdirAll(filepath.Dir(databasePath), os.ModePerm)
	if err != nil {
		panic(err)
	}

	return databasePath

}

func readLastNLines(filePath string, n int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	var lines []string
	buf := make([]byte, 1024)
	for offset := fi.Size(); offset > 0; {
		// Read backwards in chunks
		chunkSize := 1024
		if offset < int64(chunkSize) {
			chunkSize = int(offset)
		}
		offset -= int64(chunkSize)

		// Read the chunk
		file.Seek(offset, 0)
		n, err := file.Read(buf[:chunkSize])
		if err != nil {
			return nil, err
		}

		// Process the chunk
		for i := n - 1; i >= 0; i-- {
			if buf[i] == '\n' {
				lines = append(lines, "")
				if len(lines) >= n {
					break
				}
			} else {
				lines[len(lines)-1] = string(buf[i]) + lines[len(lines)-1]
			}
		}

		if len(lines) >= n {
			break
		}
	}

	// Reverse lines to maintain the correct order
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// Return the last n lines
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}

	return lines, nil
}
