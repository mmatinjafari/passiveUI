package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

// Serve the main HTML page
func mainHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

// Gather URLs and send results to the client
func gatherUrlsHandler(w http.ResponseWriter, r *http.Request) {
    domain := r.URL.Query().Get("domain")
    if domain == "" {
        http.Error(w, "Domain parameter is required", http.StatusBadRequest)
        return
    }

    // Run commands to gather URLs
    commands := []string{
        fmt.Sprintf("echo https://%s", domain),
        fmt.Sprintf("echo %s | waybackurls | sort -u", domain),
        fmt.Sprintf("echo %s | gau --subs -u", domain),
    }

    results := ""
    for _, cmdStr := range commands {
        cmd := exec.Command("bash", "-c", cmdStr)
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Fprintf(w, "Command failed: %s\nError: %s\n", cmdStr, err)
            return
        }
        results += string(output)
    }

    // Send results back to client
    fmt.Fprintf(w, results)
}

func main() {
    // Route to serve index page and handle URL gathering
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/gather-urls", gatherUrlsHandler)
    
    // Route to serve static assets like CSS and JavaScript
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
