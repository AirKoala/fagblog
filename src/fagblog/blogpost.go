package fagblog

import (
	"errors"
	"html/template"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"time"
)

type BlogPostMetadata struct {
	Title     string
	Timestamp time.Time
	Summary   string
}

type BlogPost struct {
	Metadata BlogPostMetadata
	Content  template.HTML
}

// Parses a TOML file and returns a BlogPost struct.
func GetPost(dirPath string, postName string) (BlogPost, error) {
	post := BlogPost{}
	postDirPath := dirPath + "/" + postName

	// Check if the post directory exists
	// If it doesn't exist, return os.ErrNotExist
	if _, err := os.Stat(postDirPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("Post directory does not exist: %s", postDirPath)
		return post, err
	}

	metadata, err := GetPostMetadata(dirPath, postName)
	if err != nil {
		return post, err
	}

	post.Metadata = metadata

	content, err := os.ReadFile(postDirPath + "/index.html")
	if err != nil {
		log.Printf("Error reading file %s: %v", postDirPath+"/index.html", err)
	}

	post.Content = template.HTML(content)

	return post, nil
}

func GetPostMetadata(dirPath string, postName string) (BlogPostMetadata, error) {
	metadata := BlogPostMetadata{}
	postDirPath := dirPath + "/" + postName

	// Check if the post directory exists
	// If it doesn't exist, return os.ErrNotExist
	if _, err := os.Stat(postDirPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("Post directory does not exist: %s", postDirPath)
		return metadata, err
	}

	_, err := toml.DecodeFile(postDirPath+"/meta.toml", &metadata)

	if err != nil {
		log.Printf("Error decoding TOML file: %v", err)
		return metadata, err
	}

	return metadata, nil
}

// GetPosts returns a list of all posts in the specified directory.
func GetPosts(dirPath string) ([]string, error) {
	posts := make([]string, 0)

	// Check if the post directory exists
	// If it doesn't exist, return os.ErrNotExist
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("Post directory does not exist: %s", dirPath)
		return posts, err
	}

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		log.Printf("Error reading directory %s: %v", dirPath, err)
		return posts, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			posts = append(posts, entry.Name())
		}
	}

	return posts, nil
}
