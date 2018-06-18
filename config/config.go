package config

import (
	"os"
	"bufio"
	"strings"
	"errors"
	"io"
)

type Config map[string]string

func LoadPropertiesFile(fileName string) (Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LoadProperties(f)
}

func LoadProperties(r io.Reader) (Config, error) {
	m := Config{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 0 || len(parts) > 2 {
			return nil, errors.New("invalid properties line: " + line)
		}
		m[parts[0]] = parts[1]
	}
	return m, nil
}

func (c Config) GetString(k string) (string, error) {
	if msg, ok := c[k]; ok {
		return msg, nil
	} else {
		return "", errors.New("missing property: " + k)
	}
}
