package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Scope struct {
	Key         string `yaml:"key"`
	Description string `yaml:"description"`
}

type Signal struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Scope       []Scope `yaml:"scope"`
}

type SignalFile struct {
	Group       string   `yaml:"group"`
	Description string   `yaml:"description"`
	Signals     []Signal `yaml:"signals"`
}

type SignalGroup struct {
	Title       string
	Description string
	Signals     []Signal
}

type TemplateData struct {
	Groups []SignalGroup
}

func main() {
	signalsDir := "./signals"
	files, err := filepath.Glob(filepath.Join(signalsDir, "*.signal.yaml"))
	if err != nil {
		log.Fatalf("Failed to list signal yaml files: %v", err)
	}

	var groups []SignalGroup
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read %s: %v", file, err)
		}

		var signalFile SignalFile
		if err := yaml.Unmarshal(content, &signalFile); err != nil {
			log.Fatalf("Failed to unmarshal %s: %v", file, err)
		}

		title := strings.Title(strings.ReplaceAll(signalFile.Group, "_", " "))
		groups = append(groups, SignalGroup{
			Title:       title,
			Description: signalFile.Description,
			Signals:     signalFile.Signals,
		})
	}

	tmplContent, err := ioutil.ReadFile(filepath.Join(signalsDir, "signal.tmpl"))
	if err != nil {
		log.Fatalf("Failed to read template: %v", err)
	}

	tmpl, err := template.New("signals").Parse(string(tmplContent))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	outFile, err := os.Create(filepath.Join(signalsDir, "all.md"))
	if err != nil {
		log.Fatalf("Failed to create all.md: %v", err)
	}
	defer outFile.Close()

	data := TemplateData{Groups: groups}
	if err := tmpl.Execute(outFile, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
	fmt.Println("Generated signals/all.md")
}
