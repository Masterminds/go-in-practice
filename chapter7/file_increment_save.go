package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_plus.html")
		t.Execute(w, nil)
	} else {
		mr, err := r.MultipartReader()
		values := make(map[string][]string)

		if err != nil {
			panic("Failed to read multipart message")
		}

		maxValueBytes := int64(10 << 20)
		for {

			part, err := mr.NextPart()
			if err == io.EOF {

				break
			}

			name := part.FormName()
			if name == "" {
				continue
			}

			filename := part.FileName()
			var b bytes.Buffer
			if filename == "" {
				n, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				maxValueBytes -= n
				if maxValueBytes == 0 {
					fmt.Fprint(w, "multipart message too large")
					return
				}
				values[name] = append(values[name], b.String())
				continue
			}

			dst, err := os.Create("/tmp/dstfile." + filename)
			defer dst.Close()
			if err != nil {
				return
			}
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Println("Upload done")
		fmt.Println(values)

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
