package fagblog

import "html/template"

type Context struct {
	BlogMetadata BlogMetadata
	Templates *template.Template
}
