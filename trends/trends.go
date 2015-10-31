package main

import (
    "fmt"
    "encoding/json"
    "net/http"
)

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


func main() {
    resp, err := http.Get("http://www.reddit.com/r/all/hot.json")
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
    items := make([]Item, len(r.Data.Children))
    for i, child := range r.Data.Children {
        items[i] = child.Data
    }
    // fmt.Println(items)
    fmt.Println(items[0].Title)
    fmt.Println()
    fmt.Println(items[0].URL)
    fmt.Println()
    fmt.Println(items[0].Thumbnail)

}
