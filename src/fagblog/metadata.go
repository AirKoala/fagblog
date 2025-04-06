package fagblog

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Author struct {
	Name       string
	AvatarHref string
	Blurb      string
}

type BlogMetadata struct {
	Title  string
	Author Author
}

func BlogMetadataFromToml(path string) (BlogMetadata, error) {
	parsedMetadata := BlogMetadata{}
	_, err := toml.DecodeFile(path, &parsedMetadata)

	if err != nil {
		log.Printf("Error decoding TOML file: %v", err)
		return parsedMetadata, err
	}

	return parsedMetadata, nil
}
