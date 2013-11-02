package main

import (
	"github.com/hoisie/web"
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
)

func VerifyDeploymentKey() bool {
	return true
}
func GetDeployments(ctx *web.Context) {
	fmt.Printf("************************************\n")
	deployments, err := LoadDeployments()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, deployment := range(deployments) {
		fmt.Printf(" Deployment %s ...\n", deployment.Name)
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

func main() {

	// Uncomment this to fill the database with dummy values
	/*err := InsertDeployments()
	if err != nil {
		fmt.Println(err)
		return
	}*/
	web.Get("/static", http.FileServer(http.Dir("static")))
	web.Get("/", GetIndex)
	web.Get("/deployments", GetDeployments)
	web.Run("0.0.0.0:9999")
}

