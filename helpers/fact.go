package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Fact struct {
    Fact string `json:"fact"`
}

func FetchFact() string {
    resp, err := http.Get("https://catfact.ninja/fact")
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    var fact Fact
    json.Unmarshal(body, &fact)
    return fact.Fact
}