package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/* Gets the most recently modified file */

func get_recent(path string) string {
	var old_time int64
	recent_file := ""

	walkfn := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.ModTime().Unix() > old_time {
			old_time = info.ModTime().Unix()
			recent_file = path
		}
		return nil
	}

	err := filepath.Walk(path, walkfn)
	if err != nil {
		fmt.Printf("Failed to read dir.: %v", err)
	}

	return recent_file
}

func main() {
	fmt.Println(get_recent(os.Args[1]))
}
