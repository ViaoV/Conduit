package api

import (
	"time"
)

type ApiError struct {
	Error string `json:"error"`
}

type PutMessageRequest struct {
	Mailboxes []string `json:"mailboxes"`
	Body      string   `json:"body"`
}

type PutMessageResponse struct {
	MessageSize int64    `json:"messageSize"`
	Mailboxes   []string `json:"mailboxes"`
}

type GetMessageRequest struct {
	Token   string `json:"token"`
	Mailbox string `json:"mailbox"`
}

type GetMessageResponse struct {
	Message      string    `json:"message"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"createdAt"`
	ReceiveCount int64     `json:"receiveCount"`
}

func (r *GetMessageResponse) IsEmpty() bool {
	if r.Body == "" {
		return true
	} else {
		return false
	}
}

type DeleteMessageRequest struct {
	Message string `json:"message"`
}

type DeleteMessageResponse struct {
	Message string `json:"message"`
}