package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

type Fact struct {
    Fact string `json:"fact"`
}

func FetchFact() (string, error) {
    resp, err := http.Get("https://catfact.ninja/fact")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    var fact Fact
    json.Unmarshal(body, &fact)
    return fact.Fact, nil
}