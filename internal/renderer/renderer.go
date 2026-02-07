// Package renderer provides a simple template renderer for HTML templates
package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

type Renderer struct {
	templateDir string
}

func NewRenderer(dir string) *Renderer {
	return &Renderer{
		templateDir: dir,
	}
}

// Render executes a template file with the provided data
func (r *Renderer) Render(templateFile string, data any) (string, error) {
	path := filepath.Join(r.templateDir, templateFile)

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", fmt.Errorf("failed to parse template file %s: %w", templateFile, err)
	}
	
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
