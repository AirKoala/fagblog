package fagblog

import (
	"errors"
	"html/template"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"time"
)

type BlogPost struct {
	Title     string
	Timestamp time.Time
	Content   template.HTML
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

	_, err := toml.DecodeFile(postDirPath+"/meta.toml", &post)

	if err != nil {
		log.Printf("Error decoding TOML file: %v", err)
		return post, err
	}

	content, err := os.ReadFile(postDirPath + "/index.html")
	if err != nil {
		log.Printf("Error reading file %s: %v", postDirPath+"/index.html", err)
	}

	post.Content = template.HTML(content)

	return post, nil
}
