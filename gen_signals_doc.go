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

type Signal struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Scope       []struct {
		Key         string `yaml:"key"`
		Description string `yaml:"description"`
	} `yaml:"scope"`
}

type SignalGroup struct {
	Title   string
	Signals []Signal
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

		var signals []Signal
		// The top-level key is not known, so unmarshal into a map
		var raw map[string][]Signal
		if err := yaml.Unmarshal(content, &raw); err != nil {
			log.Fatalf("Failed to unmarshal %s: %v", file, err)
		}
		for _, v := range raw {
			signals = v
			break
		}

		// Title from filename
		title := groupTitleFromFilename(filepath.Base(file))
		groups = append(groups, SignalGroup{
			Title:   title,
			Signals: signals,
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

func groupTitleFromFilename(filename string) string {
	base := strings.TrimSuffix(filename, ".signal.yaml")
	switch base {
	case "aws_cloudwatch":
		return "AWS CloudWatch Metrics"
	case "kubernetes":
		return "Kubernetes Metrics"
	default:
		return strings.Title(strings.ReplaceAll(base, "_", " "))
	}
}
