package main

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
            Data item
        }
    }
}

type item struct {
    Title     string
    URL       string
    Thumbnail string
}

func (i item) String() string {
    return fmt.Sprintf("%s\n%s\n%s", i.Title, i.URL, i.Thumbnail)
}

type jsonResponse struct {
    Links []item
}

var data []item

func grabData() {
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
        return;
    }
    data = make([]item, len(r.Data.Children))
    for i, child := range r.Data.Children {
        if !govalidator.IsURL(child.Data.URL) {
            continue;
        }
        if !govalidator.IsURL(child.Data.Thumbnail) {
            child.Data.Thumbnail = ""
        }
        data[i] = child.Data
    }
}

func getTrendsHandler(w http.ResponseWriter, r *http.Request) {


    trends := jsonResponse{data}

    js, err := json.Marshal(trends)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}


func cron() {
    for {
        grabData()
        time.Sleep(60*time.Second)
    }
}

func main() {

    go cron()

    http.HandleFunc("/trends", getTrendsHandler)
    log.Fatal(http.ListenAndServe(":5555", nil))

}
