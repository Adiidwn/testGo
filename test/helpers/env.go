package helpers

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// EnvString returns the env string value if the key exists.
// Otherwise, returns initial value.
func EnvString(key, initial string) string {
	v, exists := os.LookupEnv(key)
	if !exists {
		return initial
	}
	return v
}

// EnvInt returns the env integer value if the key exists.
// Otherwise, returns initial value.
func EnvInt(key string, initial int) int {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return n
}

// EnvDuration returns the env duration value if the key exists.
// Otherwise, returns initial value.
func EnvDuration(key string, initial time.Duration) time.Duration {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		panic(err)
	}

	return d
}

// EnvBool returns the env boolean value if the key exists.
// Otherwise, returns initial value.
func EnvBool(key string, initial bool) bool {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}

	return b
}

// LoadEnv loads env values from file.
func LoadEnv(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	values := ParseEnv(f)
	return setEnv(values)
}

type env struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// ParseEnv knows how to parse env values from reader.
func ParseEnv(reader io.Reader) []env {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var values []env

	for scanner.Scan() {
		fullLine := scanner.Text()
		fullLine = strings.TrimSpace(fullLine)

		if len(fullLine) == 0 {
			continue
		}
		if strings.HasPrefix(fullLine, "#") || strings.HasPrefix(fullLine, "=") {
			continue
		}

		parts := strings.SplitN(fullLine, "=", 2)
		kv := env{
			Key: strings.TrimSpace(parts[0]),
			Val: strings.TrimSpace(parts[1]),
		}

		values = append(values, kv)
	}

	return values
}

func setEnv(values []env) error {
	for _, v := range values {
		if err := os.Setenv(v.Key, v.Val); err != nil {
			return err
		}
	}

	return nil
}
