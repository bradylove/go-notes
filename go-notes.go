package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Note struct {
	Type    string
	Message string
	LineNum int
}

type File struct {
	Name  string
	Notes []Note
}

var TYPE_TODO string = "todo"
var Files []File

func main() {
	handleFilesForDir(".")

	// Todo Move to a easy to customize formatter
	for _, x := range Files {
		fmt.Println(x.Name)

		for _, y := range x.Notes {
			fmt.Println("["+strconv.FormatInt(int64(y.LineNum), 10)+"] ["+strings.ToUpper(y.Type)+"]\t", y.Message)
		}
	}
}

// TODO: Write Tests
func handleFilesForDir(basepath string) {
	files, err := ioutil.ReadDir(basepath)

	for _, x := range files {
		fullPath := filepath.Join(basepath, x.Name())

		if x.IsDir() {
			CheckErrF(err, "Unable to read directory, please check your permissions")

			handleFilesForDir(fullPath)
		} else {
			r, err := regexp.Compile("go|rb|java|js|coffee")
			CheckErrF(err, "Invalid regexp")

			match := r.MatchString(x.Name())

			if match {
				f := searchForTodos(fullPath)

				if len(f.Notes) > 0 {
					Files = append(Files, f)
				}
			}
		}
	}
}

// Todo: Make usable with more then just "TODO"
func searchForTodos(filepath string) File {
	file := File{
		Name: filepath,
	}

	// Setup REGEXP
	r, err := regexp.Compile("(?i:(#|//).*[^\"]todo:?[^\"])(.*)")
	CheckErrF(err, "Invalid regexp")

	// Open file for reading
	f, err := os.Open(filepath)
	CheckErrF(err, "Failed to open file for reading")
	defer f.Close()

	br := bufio.NewReader(f)
	read := 0
	for {
		b, pre, err := br.ReadLine()
		if err == io.EOF {
			break
		} else {
			if !pre {
				read++
			}

			if m := r.FindSubmatch(b); len(m) > 0 {
				x := strings.TrimSpace(string(m[2]))

				n := Note{
					Type:    TYPE_TODO,
					Message: x,
					LineNum: read,
				}

				file.Notes = append(file.Notes, n)
			}
		}
	}

	return file
}

func pathToName(path string) string {
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}

// Todo Make similar function that doesn't exit
func CheckErrF(err error, msg string) {
	if err != nil {
		fmt.Println(msg, ":", err)
		os.Exit(1)
	}
}