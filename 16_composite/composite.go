package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

type Folder struct {
	components []File
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("searching recusively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c File) {
	f.components = append(f.components, c)
}
