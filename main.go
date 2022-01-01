package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	feed, err := gofeed.NewParser().ParseURL("https://zenn.dev/komisan19/feed")
	if err != nil {
		log.Fatalln(err)
		return
	}
	file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for _, item := range feed.Items {
		fmt.Fprintf(file, "[%v](%v)\n", item.Title, item.Link)
	}
}
