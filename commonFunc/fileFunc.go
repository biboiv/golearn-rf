package commonFunc

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type FileInfo struct {
	Name    string    // base name of the file
	Size    int64     // length in bytes for regular files; system-dependent for others
	ModTime time.Time // modification time
	IsDir   bool      // abbreviation for Mode().IsDir()

}

func FSortFileListByName(dir []FileInfo) []FileInfo {
	sort.Slice(dir, func(i, j int) bool { return dir[i].Name < dir[j].Name })
	return dir
}
func FSortFileListByDate(dir []FileInfo) []FileInfo {
	sort.Slice(dir, func(i, j int) bool { return dir[i].ModTime.Before(dir[j].ModTime) })
	return dir
}
func FDirExt(dir, ext string) ([]FileInfo, error) {
	//through the ass, but...
	//
	//alternative :
	//  "path/filepath"
	//  files, _ := filepath.Glob("*.txt") - return string[]

	all, err := FDir(dir)
	if err != nil {
		return nil, err
	}
	var rez []FileInfo
	for _, f := range all {
		if strings.Contains(strings.ToLower(f.Name), strings.ToLower("."+ext)) {
			rez = append(rez, f)
		}
	}
	return rez, nil
}

func FDir(dir string) ([]FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var rezfiles []FileInfo
	for _, file := range files {
		var f FileInfo
		f.Name = file.Name()
		f.ModTime = file.ModTime()
		f.Size = file.Size()
		f.IsDir = file.IsDir()
		rezfiles = append(rezfiles, f)
	}
	return rezfiles, err
}

func FfileExists(fn string) (bool, error) {
	stat, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false, nil
	}
	if stat.IsDir() {
		return false, errors.New(fn + " is a folder.")
	}
	return true, nil
}

func FwriteTXTFile(fn string, txt []string, mode int) error {
	// fn : file name
	// txt : text
	// mode : 0 - ceate new file if file not exists
	//        1 - append text, if file exists
	//        3 - always create new file
	var err error
	var file *os.File
	var ifExists bool
	ifExists, err = FfileExists(fn)
	if err != nil {
		return err
	}
	if ifExists && (mode == 0) {
		return errors.New("Error. File exists.")
	} else {
		file, err = os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}

	defer file.Close()
	for _, s := range txt {
		fmt.Fprintln(file, s)
	}

	return nil
}

func FreadTXTFile(fn string) ([]string, error) {
	var rez []string
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rez = append(rez, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return rez, nil
}
