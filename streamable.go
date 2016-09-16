package main

import (
	"flag"
	"fmt"
	"github.com/maxkueng/go-streamable"
)

func main() {
	username := flag.String("u", "", "Your username")
	pass := flag.String("p", "", "Your password")
	filepath := flag.String("f", "", "File to upload")
	flag.Parse()

	if *username == "" || *pass == "" || *filepath == "" {
		fmt.Println("Can't leave out flags")
		fmt.Println("Example: streamable -u=username -p=password -f=video.mp4")
		fmt.Println("Use streamable --help for flag info")
		return
	}

	client := streamable.New()
	client.SetCredentials(*username, *pass)
	info, err := client.UploadVideo(*filepath)
	if err != nil {
		fmt.Printf("Upload err: %v", err)
		return
	}
	fmt.Printf("https://streamable.com/%s\n", info.Shortcode)
}
