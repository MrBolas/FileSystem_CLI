package models

import "fmt"

type Folder struct {
	subFolders []*Folder
	files      []string
	name       string
}

func NewFs() Folder {
	return Folder{
		name: "/",
	}
}

func (f *Folder) MakeDir(path []string) {

	// check if dir exists
	for _, subfolder := range f.subFolders {
		if subfolder.name == path[0] {
			subfolder.MakeDir(path[1:])
			return
		}
	}

	// if not create and update path
	newFolder := Folder{
		name: path[0],
	}
	f.subFolders = append(f.subFolders, &newFolder)
	fmt.Println(path[0], " directory created.")

	// The desired folder is deeper
	if len(path) > 1 {
		newFolder.MakeDir(path[1:])
	}
}

func (f *Folder) MakeFile(path []string, filename string) {

	// On the file
	if len(path) == 0 {
		f.files = append(f.files, filename)
		fmt.Println(filename, "file created.")
		return
	}

	// check if dir exists
	for _, subfolder := range f.subFolders {
		if subfolder.name == path[0] {
			subfolder.MakeFile(path[1:], filename)
			return
		}
	}
}

func (f *Folder) List(path []string) {

	// if in target folder
	if len(path) == 0 {

		//list subfolders
		for _, subfolder := range f.subFolders {
			fmt.Println("folder -", subfolder.name)
		}

		// list files
		for _, file := range f.files {
			fmt.Println("file -", file)
		}

		return
	}

	// scan path
	for _, subfolder := range f.subFolders {
		if subfolder.name == path[0] {
			subfolder.List(path[1:])
			return
		}
	}

	fmt.Println("Directory", path, "does not exist.")
}

func (f *Folder) RmFile(path []string, filename string) {
	// On the file
	if len(path) == 0 {
		for fileIndex, file := range f.files {
			if file == filename {
				f.files = removeFileByIndex(f.files, fileIndex)
				fmt.Println(filename, "file deleted.")
			}
		}
		return
	}

	// check if dir exists
	for _, subfolder := range f.subFolders {
		if subfolder.name == path[0] {
			subfolder.RmFile(path[1:], filename)
			return
		}
	}
}

func (f *Folder) RmDir(path []string) {

	// On the file
	if len(path) == 1 {
		for folderIndex, subfolder := range f.subFolders {
			if subfolder.name == path[0] {
				f.subFolders = removeFolderByIndex(f.subFolders, folderIndex)
				fmt.Println(subfolder.name, "directory deleted.")
			}
		}
		return
	}

	// check if its the target dir
	for _, subfolder := range f.subFolders {
		if subfolder.name == path[0] {
			subfolder.RmDir(path[1:])
			return
		}
	}
}

func removeFolderByIndex(folders []*Folder, index int) []*Folder {
	ret := make([]*Folder, 0)
	ret = append(ret, folders[:index]...)
	return append(ret, folders[index+1:]...)
}

func removeFileByIndex(files []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, files[:index]...)
	return append(ret, files[index+1:]...)
}
