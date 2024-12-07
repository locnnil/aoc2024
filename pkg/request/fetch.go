package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/locnnil/aoc2024.git/pkg/env"
)

const BaseURL = "https://adventofcode.com/2024/day/%v/input"

func ReadInput(day int) (string, error) {
	env.LoadEnv()
	tk := env.GetOrDie("SESSION_TOKEN")

	req, err := CreateRequest(day, tk)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(&req)
	if err != nil {
		return "", fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	return string(body), nil
}

func CreateRequest(day int, tk string) (http.Request, error) {
	if tk == "" {
		return http.Request{}, fmt.Errorf("no session token provided")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, day), nil)
	if err != nil {
		return http.Request{}, fmt.Errorf("error creat. HTTP request: %v", err)
	}

	// Add the session cookie
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", tk))
	return *req, nil
}
