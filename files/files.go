package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	r "github.com/basileb/custom_text_editor/renderer"
)

func getSufix(s string, separator string) (string, bool) {
	separatorIdx := strings.LastIndex(s, separator)
	if separatorIdx == -1 {
		return "", false
	}

	return s[separatorIdx+1:], true

}

func removePathGetFilename(filepath string) string {
	lastSlash := strings.LastIndex(filepath, "/")
	if lastSlash == -1 {
		return filepath
	}
	return filepath[lastSlash+1:]
}

func GetFileExtension(filepath string) r.Language {
	filename := removePathGetFilename(filepath)
	suffix, found := getSufix(filename, ".")
	if !found {
		return r.NONE
	}

	switch suffix {
	case "c":
		return r.C
	default:
		return r.NONE
	}
}

func OpenFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func WriteFile(filename string, userText []string) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	// remember to close the file
	defer f.Close()

	for _, line := range userText {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func diffLine(l1, l2 *string) bool {
	if len(*l1) != len(*l2) {
		return false
	}
	return l1 == l2
}

func DiffText(t1, t2 []string) bool {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	identical := true

	if len(t1) != len(t2) {
		return false
	}

	for i := range t2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !diffLine(&t1[i], &t2[i]) {
				mutex.Lock()
				defer mutex.Unlock()
				identical = false
			}
		}(i)
	}

	wg.Wait()
	return identical
}
