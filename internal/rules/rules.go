package rules

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"bufio"
	"path/filepath"
	"github.com/gabriel-vasile/mimetype"
	"regexp"
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

func (c *Rule) ProcessFiles(destination string) error {
	lineMatch := regexp.MustCompile(c.LineMatch)

	err := filepath.Walk(destination, func(path string, info os.FileInfo, err error) error {
		log.Println("-> processing ", path)
		mime, err := mimetype.DetectFile(path)
		if mime.Is("text/plain") {

			file, err := os.Open(path)
			if err != nil {
				log.Fatal("x unable to open ", path, " :: ", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				if lineMatch.FindStringIndex(scanner.Text()) == nil {
					log.Println("LINE MATCH :: ", scanner.Text())
				}
			}

		}

		return nil
	})

	return err
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
