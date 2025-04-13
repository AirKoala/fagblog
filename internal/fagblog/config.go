package fagblog

type Config struct {
	// The port on which the server will listen
	Port int

	// The directory where the templates are stored
	TemplateDir string

	// The directory where the blog content is stored
	ContentDir string

	StaticDir string
}
