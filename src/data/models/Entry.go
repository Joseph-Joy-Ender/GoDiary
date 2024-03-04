package models

import "time"

type Entry struct {
	id        string
	title     string
	body      string
	createdAt time.Time
}

func entry(id string, title string, body string) *Entry {
	return &Entry{
		id:        id,
		title:     title,
		body:      body,
		createdAt: time.Now(),
	}
}

func (receiver *Entry) getId() string {
	return receiver.id
}

func (receiver *Entry) setId(id string) {
	receiver.id = id
}

func (receiver *Entry) setTitle(title string) {
	receiver.title = title
}

func (receiver *Entry) getTitle() string {
	return receiver.title
}

func (receiver *Entry) setBody(body string) {
	receiver.body = body
}

func (receiver *Entry) getBody() string {
	return receiver.body

}
