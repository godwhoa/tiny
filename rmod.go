package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type File struct {
	info os.FileInfo
	path string
}

/* Implements sort interface to sort by timestamp */
type ByTime []File

func (b ByTime) Len() int {
	return len(b)
}

func (b ByTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByTime) Less(i, j int) bool {
	return b[i].info.ModTime().Unix() > b[j].info.ModTime().Unix()
}

/*
 * Gets most recent file from dir.
 * Also allows you to search with specific match.
 */
func get_recent(path string, ext string) string {
	// Build up files array
	files := ByTime{}
	walkfn := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, File{info, path})
		}
		return nil
	}
	err := filepath.Walk(path, walkfn)
	if err != nil {
		fmt.Printf("Failed to read dir.: %v", err)
	}
	// Sort it
	sort.Sort(files)
	// If users want to go by extension
	if ext != "" {
		for _, f := range files {
			if strings.Contains(f.path, ext) {
				return f.path
			}
		}
		return ""
	}
	return files[0].path
}

func main() {
	out := ""
	switch len(os.Args) {
	case 1:
		out = "No arguments."
	case 2:
		out = get_recent(os.Args[1], "")
	case 3:
		out = get_recent(os.Args[1], os.Args[2])
	}
	fmt.Println(out)
}
