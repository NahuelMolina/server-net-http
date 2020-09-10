package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
    "log"
)

type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename:= p.Title + ".txt"
    return ioutil.WriteFile(filename,p.Body,0600)
}

func loadPage(title string)(*Page, error){
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil,err
    }
    return &Page{Title:title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>",p.Title,p.Body)
}

func seconfHandler (w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"Merhaba %s",r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/init/", secondHanlder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


