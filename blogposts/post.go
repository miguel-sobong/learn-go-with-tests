package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile fs.File) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{Title: readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner)}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignores ---
	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
