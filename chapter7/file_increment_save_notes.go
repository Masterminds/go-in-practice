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
		fmt.Println(os.TempDir())
	} else {

		mr, err := r.MultipartReader()
		if err != nil {
			panic("Failed to read multipart message: ")
		}

		length := r.ContentLength
		for {

			part, err := mr.NextPart()
			if err == io.EOF {

				break
			}
			var read int64
			var p float32

			filename := part.FileName()
			dst, err := os.Create("/tmp/dstfile." + filename)
			if err != nil {
				return
			}
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				read = read + int64(cBytes)
				//fmt.Printf("read: %v \n",read )
				p = float32(read) / float32(length) * 100
				fmt.Printf("progress: %v \n", p)
				dst.Write(buffer[0:cBytes])
			}
		}

		// f, h, err := r.FormFile("file")
		// if err != nil {
		// 	panic(err)
		// }
		// defer f.Close()
		// filename := "/tmp/" + h.Filename
		// out, _ := os.Create(filename)
		// defer out.Close()
		// fmt.Println(h.Header["Content-Type"][0])

		// buff := make([]byte, 512)
		// _, err = f.Read(buff)
		// filetype := http.DetectContentType(buff)
		// fmt.Println(filetype)

		// io.Copy(out, f)
		fmt.Println("Upload done")

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
