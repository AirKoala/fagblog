package fagblog

import "html/template"

type Context struct {
	BlogMetadata BlogMetadata
	Templates map[string]*template.Template
}
