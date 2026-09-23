package main

import "fmt"

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".png":  "Images",
	".jpeg": "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".mp3":  "Music",
	".wav":  "Music",
	".mp4":  "Video",
	".avi":  "Video",
	".zip":  "Archives",
	".rar":  "Archives",
}

func main() {
	fmt.Print(DefaultRules)
}