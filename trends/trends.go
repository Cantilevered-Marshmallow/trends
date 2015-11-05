package trends

import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    "time"
    "github.com/asaskevich/govalidator"
)

var redditURL = "http://www.reddit.com/r/all/hot.json"

type redditResponse struct {
    Data struct {
        Children []struct {
            Data redditItem
        }
    }
}

type redditItem struct {
    Title     string
    URL       string
    Thumbnail string
}

func (i redditItem) String() string {
    return fmt.Sprintf("%s\n%s\n%s", i.Title, i.URL, i.Thumbnail)
}

type trendsResponse struct {
    Links []redditItem
}

var redditData []redditItem

func grabRedditData() {
    resp, err := http.Get(redditURL)
    if err != nil {
        return;
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return;
    }
    r := new(redditResponse)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return
    }

    redditData = make([]redditItem, len(r.Data.Children))
    for i, child := range r.Data.Children {
        if !govalidator.IsURL(child.Data.URL) {
            continue;
        }
        if !govalidator.IsURL(child.Data.Thumbnail) {
            child.Data.Thumbnail = ""
        }
        redditData[i] = child.Data
    }
}

func getTrendsHandler(w http.ResponseWriter, r *http.Request) {
    trends := trendsResponse{redditData}

    js, err := json.Marshal(trends)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func cron() {
    for {
        grabRedditData()
        time.Sleep(60*time.Second)
    }
}

func main() {

    go cron()
    http.HandleFunc("/trends", getTrendsHandler)
    http.HandleFunc("/message", postMessageHandler)
    log.Fatal(http.ListenAndServe(":5555", nil))
}
