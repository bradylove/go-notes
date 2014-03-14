package main

import (
	"bufio"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Options struct {
	ShowTodos  bool `short:"t" verbose:"todos" description:"Show todos"`
	ShowNotes  bool `short:"n" verbose:"notes" description:"Show notes"`
	ShowFixMes bool `short"f" verbose:"fixmes" description:"Show fixmes"`
	ShowDups   bool `short"d" verbose:"dups" description:"Show dups (duplicates)"`
}

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
	flag.Parse()

	FilterType = strings.ToLower(FilterType)

	fmt.Println(FilterType)

	var rootPath string

	if len(os.Args) > 1 {
		rootPath = os.Args[1]
	}

	if rootPath == "" {
		rootPath = "."
	}

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Println("Usage: src-notes: [directory]")
		return
	}

	handleFilesForDir(rootPath)

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
			f := searchForTodos(fullPath)

			if len(f.Notes) > 0 {
				Files = append(Files, f)
			}
		}
	}
}

func extFromName(fName string) string {
	chunks := strings.Split(fName, ".")

	if len(chunks) > 1 {
		return chunks[len(chunks)-1]
	} else {
		return ""
	}
}

// Todo: Make usable with more then just "TODO" hmm
func searchForTodos(filepath string) File {
	file := File{
		Name: filepath,
	}

	ext := extFromName(file.Name)
	cmt := Cmts[ext]

	if cmt == "" {
		return file
	}

	// Todo: Make regexp work with HTML
	// Fixme: Make regexp more dynamic
	r, err := regexp.Compile("(?i)(.*" + cmt + "\\s)(note|todo|fixme|dup):?[^\"](.*)")
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
