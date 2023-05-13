package config

import "html/template"

// AppConfig hoplds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
