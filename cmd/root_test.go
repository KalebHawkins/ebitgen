package cmd

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/*.tmpl
var tefs embed.FS

var testTemplateDir = "testdata"
var testTemplateName = "test.go.tmpl"

func TestLoadTemplates(t *testing.T) {
	err := loadTemplates(tefs, testTemplateDir)
	assert.NoError(t, err)
	assert.NotNil(t, templates[testTemplateName])
}
