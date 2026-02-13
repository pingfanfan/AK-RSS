package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Timezone   string           `yaml:"timezone"`
	Store      StoreConfig      `yaml:"store"`
	Feeds      FeedsConfig      `yaml:"feeds"`
	Rules      []RuleConfig     `yaml:"rules"`
	Summarizer SummarizerConfig `yaml:"summarizer"`
	Sinks      []SinkConfig     `yaml:"sinks"`
}

type StoreConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type FeedsConfig struct {
	OPMLPath string `yaml:"opml_path"`
}

type RuleConfig struct {
	Name             string   `yaml:"name"`
	IncludeKeywords  []string `yaml:"include_keywords"`
	ExcludeKeywords  []string `yaml:"exclude_keywords"`
	IncludeDomains   []string `yaml:"include_domains"`
	MinContentLength int      `yaml:"min_content_length"`
}

type SinkConfig struct {
	Type    string            `yaml:"type"`
	Name    string            `yaml:"name"`
	Enabled bool              `yaml:"enabled"`
	Config  map[string]string `yaml:"config"`
}

type SummarizerConfig struct {
	Type    string            `yaml:"type"`
	Enabled bool              `yaml:"enabled"`
	Config  map[string]string `yaml:"config"`
}

func Load(path string) (Config, error) {
	var cfg Config
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func Validate(cfg Config) error {
	if strings.TrimSpace(cfg.Feeds.OPMLPath) == "" {
		return errors.New("feeds.opml_path is required")
	}
	if len(cfg.Sinks) == 0 {
		return errors.New("at least one sink is required")
	}
	for i, s := range cfg.Sinks {
		if strings.TrimSpace(s.Type) == "" {
			return fmt.Errorf("sinks[%d].type is required", i)
		}
		if s.Name == "" {
			return fmt.Errorf("sinks[%d].name is required", i)
		}
	}
	return nil
}
