package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joerx/gojournal/pkg/data"
)

// Generic message sent be the server
type serverStatus struct {
	Message    string
	APIVersion string
}

func prepare(w http.ResponseWriter) {

}

// Index handles the root url, returning a simple status message
func Index(w http.ResponseWriter, r *http.Request) {
	m := serverStatus{Message: "Hello World!", APIVersion: "1.0"}
	sendJSONResponse(w, m)

}

// GetEntries returns the list of journal entries
func GetEntries(w http.ResponseWriter, r *http.Request) {
	a := &data.Attachment{
		ID:   2,
		URL:  "http://foo.bar/image.png",
		Type: "image/png",
		Metadata: &data.ImageMeta{
			Filename: "image.png",
		},
	}
	e := data.Entry{
		Title:       "Hello World",
		ID:          1,
		Content:     "A dumb sample entry",
		Attachments: []*data.Attachment{a},
	}
	d := []data.Entry{e}
	sendJSONResponse(w, d)
}

// encode data as JSON and send it to the client with the correct content-type header set
func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
