package main

import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    "time"
)

var redditUrl string = "http://www.reddit.com/r/all/hot.json"

type Response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

type Item struct {
    Title    string
    URL      string
    Thumbnail string
}

func (i Item) String() string {
    return fmt.Sprintf("%s\n%s\n%s", i.Title, i.URL, i.Thumbnail)
}

type Trends struct {
    Topics []Item
}

var data []Item

func grabData() {
    resp, err := http.Get(redditUrl)
    if err != nil {
        return;
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return;
    }
    r := new(Response)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return;
    }
    data = make([]Item, len(r.Data.Children))
    for i, child := range r.Data.Children {
        data[i] = child.Data
    }
}

func getTrendshandler(w http.ResponseWriter, r *http.Request) {


    trends := Trends{data}
    
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
        time.Sleep(5*time.Second)
    }
}

func main() {

    go cron()

    http.HandleFunc("/trends", getTrendshandler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}
