package rules

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type RulesSpec struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Id          string `yaml:"id"`
	Description string `yaml:"description"`
	LineMatch   string `yaml:"line_match,omitempty"`
	StartMarker string `yaml:"start_marker,omitempty"`
	EndMarker   string `yaml:"end_marker,omitempty"`
}

func (c *RulesSpec) Parse(specFile string) *RulesSpec {
	filename, _ := filepath.Abs(specFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read %s %v", filename, err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal("Failed to unmarshal: %v", err)
		os.Exit(1)
	}
	return c
}
