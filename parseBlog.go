package main

import (
	"fmt"

	"github.com/Sri2103/blog/internal/blog"
)

func ParseBlog(b *blog.BlogExcerpt) {
	// Parse the blog post and print out relevant information

	title := b.Title
	date := b.Date

	fmt.Println("Title: ", title)
	fmt.Println("Date: ", date)
}
