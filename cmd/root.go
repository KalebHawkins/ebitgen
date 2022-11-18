/*
Copyright © 2022 Kaleb Hawkins <KalebHawkins@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const (
	templateDir  = "templates"
	defaultTempl = "default.go.tmpl"
)

var (
	//go:embed templates/*.tmpl
	efs embed.FS

	templates map[string]*template.Template

	width   int
	height  int
	title   string
	outFile string
)

type Project struct {
	Width  int
	Height int
	Title  string
}

var rootCmd = &cobra.Command{
	Use:   "ebitgen",
	Short: "A tool to generate a new Ebiten project.",
	Long:  `Ebitgen is a tool for generating new Ebiten projects from templates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadTemplates(efs, templateDir); err != nil {
			return err
		}

		f, err := os.Create(outFile)
		if err != nil {
			return err
		}

		p := &Project{
			Width:  width,
			Height: height,
			Title:  title,
		}

		if err := parseTemplate(f, p, defaultTempl); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&width, "width", "w", 640, "The width of the game window")
	rootCmd.Flags().IntVarP(&height, "height", "H", 480, "The height of the game window")
	rootCmd.Flags().StringVarP(&title, "title", "t", "New Game", "The title of the game window")
	rootCmd.Flags().StringVarP(&outFile, "outfile", "o", "game.go", "The file to write to")

}

func loadTemplates(efs embed.FS, templateDir string) error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	fs.WalkDir(efs, templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		pt, err := template.ParseFS(efs, path)
		if err != nil {
			return nil
		}

		templates[d.Name()] = pt

		return nil
	})

	return nil
}

func parseTemplate(w io.Writer, p *Project, templateName string) error {
	return templates[templateName].Execute(w, p)
}
