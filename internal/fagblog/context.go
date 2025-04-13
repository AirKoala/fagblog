package fagblog

import "html/template"

type Context struct {
	SiteMetadata SiteMetadata
	Templates map[string]*template.Template
}
