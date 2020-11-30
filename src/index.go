package main

import (
    "net/http"
    "log"
    "github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init (){
    opts := renderer.Options{
        ParseGlobPattern:"./src/views/*.html",
    }
    rnd = renderer.New(opts)
}

func home (w http.ResponseWriter, r *http.Request){
    rnd.HTML(w,http.StatusOK,"example",nil)
}

func about (w http.ResponseWriter, r *http.Request){
    rnd.HTML(w,http.StatusOK, "about-me",nil)
}

func main(){    
    port:=":3000"
    mux := http.NewServeMux()
    mux.HandleFunc("/home", home)
    mux.HandleFunc("/about",about)

    log.Println("listening on port: ", port)
    log.Println("ensure that you have the url links correctly setted")

    http.ListenAndServe(port,mux)
}
