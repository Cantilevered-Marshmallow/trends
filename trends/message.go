package trends

import (
    "net/http"
    "encoding/json"
)

type Message struct {
    Text             string
    YoutubeVideoId   string
    GoogleImageId    string
    ChatId           int
    UserFacebookId   string
    RedditAttachment struct {
                         URL       string
                         Title     string
                         Thumbnail string
                     }
}

var MessagesCache map[int][]Message

func cacheMessage (message *Message) {
    chatId := message.ChatId
    messages := MessagesCache[chatId]
    if messages[0].Text != "" {
        messages = append(messages, *message)
    } else {
        messages = []Message{*message}
    }
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {

    message := new(Message)
    err := json.NewDecoder(r.Body).Decode(message)
    if err != nil {
        return
    }

    cacheMessage(message)

    w.WriteHeader(201)
}
