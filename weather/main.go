package main

import (
    "fmt"
    "net/http"
    "io"
    "strings"
    "os"
)

func main() {
    resp, err := http.Get("https://wttr.in/Grenoble?format=Weather:+%l+%C+%t+%h")
    if err != nil {
        fmt.Println("erreur:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    weather := strings.TrimSpace(string(body))

    content, _ := os.ReadFile("README.md")
    markerStart := "<!-- weather-start -->"
    markerEnd := "<!-- weather-end -->"
    start := strings.Index(string(content), markerStart)
    end := strings.Index(string(content), markerEnd)
    if start == -1 || end == -1 || end < start {
        fmt.Println("marqueurs météo absents")
        os.Exit(1)
    }

    before := content[:start+len(markerStart)]
    after := content[end:]
    updated := fmt.Sprintf("%s\n\n<code>%s</code>\n%s", before, weather, after)
    os.WriteFile("README.md", []byte(updated), 0644)
}
