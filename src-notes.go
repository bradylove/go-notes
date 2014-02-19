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

// Note: These are the different types of notes
var TYPE_TODO string = "todo"
var TYPE_NOTE string = "note"

var Files []File

// Note: This is the main func
func main() {
	handleFilesForDir(".")

	// Todo Move to a easy to customize formatter
	for _, x := range Files {
		x.PrintNotes()
	}
}

func (f *File) PrintNotes() {
	fmt.Println("\033[1m" + f.Name + "\033[0m")

	for _, y := range f.Notes {
		fmt.Println("["+strconv.FormatInt(int64(y.LineNum), 10)+"] ["+strings.ToUpper(y.Type)+"]\t", y.Message)
	}
}

// TODO: Write Tests for todos
func handleFilesForDir(basepath string) {
	files, err := ioutil.ReadDir(basepath)

	for _, x := range files {
		fullPath := filepath.Join(basepath, x.Name())

		if x.IsDir() {
			CheckErrF(err, "Unable to read directory, please check your permissions")

			handleFilesForDir(fullPath)
		} else {
			r, err := regexp.Compile("go|rb|java|js|coffee|cc|cpp|h|cl|el|lisp|hs")
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

// Todo: Make usable with more then just "TODO" hmm
func searchForTodos(filepath string) File {
	file := File{
		Name: filepath,
	}

	// Setup REGEXP
	r, err := regexp.Compile("(?i:(#|//|;;|--).*[^\"](note|todo):?[^\"])(.*[^[todo]|[note])")
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
				x := strings.TrimSpace(string(m[3]))

				n := Note{
					Type:    strings.ToLower(string(m[2])),
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

// Todo Make similar function that doesn't exit okay?
func CheckErrF(err error, msg string) {
	if err != nil {
		fmt.Println(msg, ":", err)
		os.Exit(1)
	}
}