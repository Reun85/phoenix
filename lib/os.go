package lib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var r = "\x1b[0m"

func Mkdir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			Err(err)
		}
	}
}

func Rm(file string) {
	if err := os.RemoveAll(file); err != nil {
		Err(err)
	}
}

func Ln(target, linkName string) {
	if err := os.Symlink(target, linkName); err != nil {
		Err(err)
	}
}

func Cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		Err(err)
	}
	return cwd
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func WriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		Err(err)
	}
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		Err(err)
	}
	return data
}

func Exec(cmd string, args ...string) *exec.Cmd {
	_, err := exec.LookPath(cmd)
	if err != nil {
		Err(`executable "` + Magenta(cmd) + `" not found in $PATH`)
	}
	return exec.Command(cmd, args...)
}

func Invert(str string) string {
	return "\x1b[7m" + str + r
}

func Red(str string) string {
	return "\x1b[31m" + str + r
}

func Green(str string) string {
	return "\x1b[32m" + str + r
}

func Yellow(str string) string {
	return "\x1b[33m" + str + r
}

func Blue(str string) string {
	return "\x1b[34m" + str + r
}

func Magenta(str string) string {
	return "\x1b[35m" + str + r
}

func Cyan(str string) string {
	return "\x1b[36m" + str + r
}

func Err(err any) {
	switch v := err.(type) {
	case string:
		fmt.Fprintln(os.Stderr, Red("error: ")+v)
	case error:
		fmt.Fprintln(os.Stderr, Red("error: ")+v.Error())
	}
	os.Exit(1)
}

func ListDir(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		Err(err)
	}
	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names
}

func CopyFile(src, dst string) {
	input, err := os.ReadFile(src)
	if err != nil {
		Err(err)
	}
	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		Err(err)
	}
}

func MoveFile(src, dst string) {
	err := os.Rename(src, dst)
	if err != nil {
		Err(err)
	}
}

// GetEnv fetches an environment variable, returning a default value if it is not set.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func SetEnv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		Err(err)
	}
}

func TempFile(prefix string) string {
	file, err := os.CreateTemp("", prefix+"*")
	if err != nil {
		Err(err)
	}
	defer file.Close()
	return file.Name()
}

// HomeDir returns the current user's home directory.
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		Err(err)
	}
	return home
}

// FileInfo provides details about a file.
func FileInfo(path string) os.FileInfo {
	info, err := os.Stat(path)
	if err != nil {
		Err(err)
	}
	return info
}

func FileMode(path string) string {
	info := FileInfo(path)
	return info.Mode().String()
}

func FileSize(path string) int64 {
	info := FileInfo(path)
	return info.Size()
}

// ModTime returns the last modification time of a file.
func ModTime(path string) time.Time {
	info := FileInfo(path)
	return info.ModTime()
}

func AbsPath(relPath string) string {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		Err(err)
	}
	return absPath
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func Chmod(path string, mode os.FileMode) {
	err := os.Chmod(path, mode)
	if err != nil {
		Err(err)
	}
}

func CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		Err(err)
	}
	defer file.Close()
}

// TODO: change this to mostly use embedded variables
func defaultConfigDir() string {
	dotconf, _ := os.UserConfigDir()
	return filepath.Join(dotconf, dotconfname)
}

func logging(logFile string) (io.Writer, io.Writer, *os.File) {
	if logFile == "" {
		return os.Stdout, os.Stderr, nil
	}

	Mkdir(filepath.Dir(logFile))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		Err(err)
	}

	return io.MultiWriter(os.Stdout, file), io.MultiWriter(os.Stderr, file), file
}
