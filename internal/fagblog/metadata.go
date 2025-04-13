package fagblog

import (
	"github.com/BurntSushi/toml"
	"log"
)

type HeaderLink struct {
	Name string
	Href string
}

type Author struct {
	Name       string
	AvatarHref string
	Blurb      string
}

type SiteMetadata struct {
	Title       string
	Author      Author
	HeaderLinks []HeaderLink
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
