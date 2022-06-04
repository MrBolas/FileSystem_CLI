package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/fs/models"
)

func main() {

	// define fs
	Fs := models.NewFs()

	// command loop
	// Support commands:
	// mkdir
	// mkfile
	// cd - active directory implementation ( save for later )
	// ls
	// rm dir
	// rm file

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------Balls Shell-----------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// mkdir folder1/folder2
		if strings.HasPrefix(text, "mkdir") {
			//sanitizing input
			path, err := sanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			Fs.MakeDir(path)
		}

		// rmdir folder1/folder2
		if strings.HasPrefix(text, "rmdir") {
			//sanitizing input
			path, err := sanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			Fs.RmDir(path)
		}

		// mkfile folder1/file2
		if strings.HasPrefix(text, "mkfile") {
			//sanitizing input
			path, err := sanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			filename := path[len(path)-1]
			path = path[0 : len(path)-1]

			Fs.MakeFile(path, filename)
		}

		// rmfile folder1/file2
		if strings.HasPrefix(text, "rmfile") {
			//sanitizing input
			path, err := sanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			filename := path[len(path)-1]
			path = path[0 : len(path)-1]

			Fs.RmFile(path, filename)
		}

		// ls folder1/folder2
		if strings.HasPrefix(text, "ls") {
			//sanitizing input
			path, err := sanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Fs.List(path)
		}

		fmt.Println("Unknown command")
	}
}

func sanitizeInputs(command string) ([]string, error) {

	arguments := strings.Split(command, " ")
	if len(arguments) < 2 {
		return []string{}, errors.New("not enough arguments")
	}

	path := strings.Split(arguments[1], "/")
	if len(path) < 1 {
		return []string{}, errors.New("invalid argument")
	}

	return path, nil
}
