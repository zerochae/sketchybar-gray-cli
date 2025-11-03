package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Config struct {
	path string
	data map[string]string
}

func New(path string) *Config {
	return &Config{
		path: path,
		data: make(map[string]string),
	}
}

func NewUser() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	path := filepath.Join(home, ".config", "sketchybar", "user.sketchybarrc")
	return New(path)
}

func NewDefault() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	path := filepath.Join(home, ".config", "sketchybar", "sketchybarrc")
	return New(path)
}

func (c *Config) Load() error {
	data, err := os.ReadFile(c.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return c.loadBashFormat(data)
}

func (c *Config) loadBashFormat(data []byte) error {
	bashConfig := make(map[string]string)
	lines := string(data)

	for _, line := range splitLines(lines) {
		if key, value, ok := parseBashExport(line); ok {
			bashConfig[key] = value
		}
	}

	c.data = bashConfig
	return nil
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, ch := range s {
		if ch == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func parseBashExport(line string) (string, string, bool) {
	trimmed := ""
	for _, ch := range line {
		if ch != ' ' && ch != '\t' {
			trimmed += string(ch)
		} else if trimmed != "" {
			trimmed += string(ch)
		}
	}

	if len(trimmed) < 7 || trimmed[:6] != "export" {
		return "", "", false
	}

	rest := trimmed[6:]
	for len(rest) > 0 && (rest[0] == ' ' || rest[0] == '\t') {
		rest = rest[1:]
	}

	eqIdx := -1
	for i, ch := range rest {
		if ch == '=' {
			eqIdx = i
			break
		}
	}

	if eqIdx == -1 {
		return "", "", false
	}

	key := rest[:eqIdx]
	value := rest[eqIdx+1:]

	if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}

	return key, value, true
}

func (c *Config) Get(key string) (string, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Config) Set(key, value string) {
	normalizedKey := NormalizeKey(key)
	c.data[normalizedKey] = value
}

func (c *Config) List() map[string]string {
	result := make(map[string]string)
	for k, v := range c.data {
		result[k] = v
	}
	return result
}

func (c *Config) Save() error {
	dir := filepath.Dir(c.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	var content string
	content += "#!/usr/bin/env bash\n"
	content += "\n"
	content += "# User Custom Configuration\n"
	content += "# Managed by gsbar - https://github.com/zerochae/gsbar\n"
	content += "\n"

	for key, value := range c.data {
		exportKey := NormalizeKey(key)
		content += fmt.Sprintf("export %s=\"%s\"\n", exportKey, value)
	}

	return os.WriteFile(c.path, []byte(content), 0o644)
}

func GetValueCascade(key string) (string, error) {
	normalizedKey := NormalizeKey(key)

	userCfg := NewUser()
	if userCfg == nil {
		return "", fmt.Errorf("failed to get user config path")
	}

	if err := userCfg.Load(); err != nil {
		return "", err
	}

	if val, ok := userCfg.Get(normalizedKey); ok {
		return val, nil
	}

	defaultCfg := NewDefault()
	if defaultCfg == nil {
		return "", fmt.Errorf("failed to get default config path")
	}

	if err := defaultCfg.Load(); err != nil {
		return "", err
	}

	if val, ok := defaultCfg.Get(normalizedKey); ok {
		return val, nil
	}

	return "", fmt.Errorf("key not found: %s", key)
}

func toScreamingSnakeCase(camelCase string) string {
	var result strings.Builder

	for i, ch := range camelCase {
		if i > 0 && ch >= 'A' && ch <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(ch)
	}

	return "SBAR_" + strings.ToUpper(result.String())
}

func NormalizeKey(key string) string {
	if strings.HasPrefix(key, "SBAR_") || regexp.MustCompile(`^[A-Z_]+$`).MatchString(key) {
		return key
	}
	return toScreamingSnakeCase(key)
}
