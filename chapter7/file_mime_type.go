package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filename := "/tmp/" + h.Filename
		out, _ := os.Create(filename)
		defer out.Close()
		fmt.Println(h.Header["Content-Type"][0])

		buff := make([]byte, 512)
		_, err = f.Read(buff)
		filetype := http.DetectContentType(buff)
		fmt.Println(filetype)

		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
