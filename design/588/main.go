package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

type (
	FileSystem struct {
		fs map[string]*Dir
	}

	File struct {
		name string
		body string
	}

	Dir struct {
		fsd   map[string]*Dir
		files []*File
	}
)

func Constructor() FileSystem {
	mp := make(map[string]*Dir)
	mp["/"] = &Dir{map[string]*Dir{}, []*File{}}
	return FileSystem{
		fs: mp,
	}

}

func (this *FileSystem) Ls(path string) []string {
	//fmt.Println("LS", path)
	dir := this.fs["/"]
	var result []string
	pa := strings.Split(path, "/")
	for i := 0; i < len(pa); i++ {
		if pa[i] == "" {
			continue
		}
		newDir, ok := dir.fsd[pa[i]]
		if !ok {
			break
		}

		dir = newDir
	}
	for _, v := range dir.files {
		if v.name == pa[len(pa)-1] {
			return []string{v.name}
		}
		result = append(result, v.name)
	}

	for v := range dir.fsd {
		result = append(result, v)
	}

	sort.Strings(result)

	return result
}

func (this *FileSystem) Mkdir(path string) {
	//fmt.Println("Mkdir", path)

	dir := this.fs["/"]

	pa := strings.Split(path, "/")

	for i := 0; i < len(pa); i++ {
		if pa[i] == "" {
			continue
		}
		newDir, ok := dir.fsd[pa[i]]
		if !ok {
			dir.fsd[pa[i]] = &Dir{map[string]*Dir{}, []*File{}}
			newDir = dir.fsd[pa[i]]
		}
		dir = newDir
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	//fmt.Println("AddContentToFile", filePath, content)
	dir := this.fs["/"]

	path := strings.Split(filePath, "/")
	for i := 0; i < len(path)-1; i++ {
		if path[i] == "" {
			continue
		}

		newDir, ok := dir.fsd[path[i]]
		if !ok {
			//fmt.Println(path, "not exists")
			return
		}

		dir = newDir
	}

	for i := 0; i < len(dir.files); i++ {
		if dir.files[i].name == path[len(path)-1] {
			dir.files[i].body += content
			return
		}
	}
	dir.files = append(dir.files, &File{path[len(path)-1], content})
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	//fmt.Println("ReadContentFromFile", filePath)
	dir := this.fs["/"]

	path := strings.Split(filePath, "/")
	for i := 0; i < len(path)-1; i++ {
		if path[i] == "" {
			continue
		}
		newDir, ok := dir.fsd[path[i]]
		if !ok {
			//fmt.Println(path, path[i], "not exists")
		}

		dir = newDir
	}

	for i := 0; i < len(dir.files); i++ {
		if dir.files[i].name == path[len(path)-1] {
			return dir.files[i].body
		}
	}

	return ""
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
