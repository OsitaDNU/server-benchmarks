package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// BlogPost represents a mock blog post
type BlogPost struct {
    Title      string    `json:"title"`
    Content    string    `json:"content"`
	URL        string    `json:"url"`
    Author     string    `json:"author"`
    CreatedAt  time.Time `json:"created_at"`
    Tags       []string  `json:"tags"`
    Views      int       `json:"views"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    post := BlogPost{
        Title:     "Exploring the Tranquil Beauty of Nature: A Journey to Serendipity",
        Content:   `There is something incredibly captivating about the serene embrace of nature. From the gentle rustling of leaves to the melodic chirping of birds, nature offers a symphony of tranquility that beckons us to step away from the hustle and bustle of daily life. This blog post is a celebration of nature's beauty and the profound peace it brings to our souls.`,
		URL:       "http://localhost:8080/blog/1",
        Author:    "Author Name",
        CreatedAt: time.Now(),
        Tags:      []string{"Go", "Benchmark", "Web"},
        Views:     123,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
