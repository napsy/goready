package main

import (
	"github.com/hoisie/web"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"time"
	"net/http"
	"html/template"
)

type Deployment struct {
	Name           string
	Description    string
	Hostname       string
	IPAddress      string
	deploymentKey  string // deploymeny key
	Configuration  string // custom configuration in json
}

func VerifyDeploymentKey() bool {
	return true
}
func GetDeployments(ctx *web.Context) {
	deployments := []*Deployment{
		&Deployment{"Testna", "Testna postavitev", "localhost", "127.0.0.1", "AAABBB", ""},
		&Deployment{"Druga", "Druga testna postavitev", "localhost", "127.0.0.1", "CCCDDD", ""},
		&Deployment{"Beta staging", "Beta staging server", "localhost", "127.0.0.1", "CCCDDD", ""},
	}
	b, err := json.Marshal(deployments)
	if err != nil {
		fmt.Println("error:", err)
	}
	ctx.ContentType("json")
	ctx.Write(b)
}

func GetIndex(ctx *web.Context) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(ctx.ResponseWriter, nil)
}

func wsHandler(ws *websocket.Conn) {
	op := struct {
		Type string `json:type`
		Data struct {
			Name string
			Operation string
		} `json:data`
	}{}

	websocket.JSON.Receive(ws, &op)

	switch op.Data.Operation {
	case "upgrade":
		fmt.Printf("Performing upograde of deploymeny '%s' ...\n", op.Data.Name)
	}
	websocket.JSON.Send(ws, map[string]string{"type":"message", "data":"OK"})
	time.Sleep(time.Second * 2)
	websocket.JSON.Send(ws, map[string]string{"type":"message", "data":"DONE"})
	ws.Close()
}
func main() {
	web.Get("/static", http.FileServer(http.Dir("static")))
	web.Get("/ws", websocket.Handler(wsHandler))
	web.Get("/", GetIndex)
	web.Get("/deployments", GetDeployments)
	web.Run("0.0.0.0:9999")
}

