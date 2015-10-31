package trends

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "bytes"
    "time"
)

func checkJSONFormat (body *bytes.Buffer) bool {
    r := new(jsonResponse)
    err := json.NewDecoder(body).Decode(r)
    if err != nil {
        return false
    }
    if (len(r.Links) > 0){
        return true
    }
    return false
}

func TestGetJson (t *testing.T) {
    go cron()
    time.Sleep(1*time.Second)
    req, _ := http.NewRequest("GET", "", nil)
    w := httptest.NewRecorder()
    getTrendsHandler(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("Home page didn't return %v", http.StatusOK)
    }
    if !checkJSONFormat(w.Body) {
        t.Errorf("Something wrong with JSON response")
    }
}
