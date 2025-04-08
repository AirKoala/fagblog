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

type SiteMetadata struct {
	Title  string
	Author Author
}

func SiteMetadataFromToml(path string) (SiteMetadata, error) {
	parsedMetadata := SiteMetadata{}
	_, err := toml.DecodeFile(path, &parsedMetadata)

	if err != nil {
		log.Printf("Error decoding TOML file: %v", err)
		return parsedMetadata, err
	}

	return parsedMetadata, nil
}
