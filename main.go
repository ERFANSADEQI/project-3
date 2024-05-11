package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type Subscription struct {
	Conn *websocket.Conn
	Topic string
}

var subscriptions = make([] *Subscription, 0)