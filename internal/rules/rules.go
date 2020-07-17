package rules

import (
	"bufio"
	"encoding/json"
	"github.com/gabriel-vasile/mimetype"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

type MatchResult struct {
	Filename string `json:"filename"`
	Results  string `json:"results"`
	Rule     string `json:"rule"`
	Name     string `json:"name"`
}

var results []MatchResult

func SaveResults() {
	file, _ := json.MarshalIndent(results, "", "")
	_ = ioutil.WriteFile("columbo-report.json", file, 0644)
}

// Processes a single line match printing the result if found
func (c *Rule) ProcessLineMatch(destination string) error {
	err := filepath.Walk(destination, func(path string, info os.FileInfo, err error) error {
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
				found, _ := regexp.MatchString(c.LineMatch, scanner.Text())
				if found {
					results = append(results,
						MatchResult{
							Filename: path,
							Results:      scanner.Text(),
							Rule: c.Id,
							Name: c.Id,
						})
				}
			}

		}

		return nil
	})

	return err
}

// Processes a start and end marker match printing the result if found
func (c *Rule) ProcessStartEndMarker(destination string) error {

	var isMatching bool
	var matchedLines []string

	err := filepath.Walk(destination, func(path string, info os.FileInfo, err error) error {
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
				startFound, _ := regexp.MatchString(c.StartMarker, scanner.Text())
				endFound, _ := regexp.MatchString(c.EndMarker, scanner.Text())
				if startFound {
					isMatching = true
					matchedLines = append(matchedLines, scanner.Text())
				} else if endFound {
					isMatching = false
					matchedLines = append(matchedLines, scanner.Text())
				} else if isMatching {
					matchedLines = append(matchedLines, scanner.Text())
				}
			}
			if len(matchedLines) > 0 {
				results = append(results, MatchResult{
					Filename: path,
					Results:      strings.Join(matchedLines, "\n"),
					Rule: c.Id,
					Name: c.Id,
				})
				matchedLines = nil
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
