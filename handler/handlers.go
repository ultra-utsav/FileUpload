package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

//IndexHandler to render html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")

	er := t.Execute(w, nil)

	if er != nil {
		log.Fatal("Error in rendering index webpage, ", er.Error())
	}
}

// ResumableUpload handles reumable uploads
func ResumableUpload(w http.ResponseWriter, r *http.Request) {
	tempFolder := "./files"

	//check if files folder exists or not, if not create it
	if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
		os.Mkdir(tempFolder, os.ModePerm)
	}

	query := r.URL.Query()
	resumableIdentifier, _ := query["resumableIdentifier"]

	filePath := fmt.Sprintf("%s/%s", tempFolder, resumableIdentifier[0])

	//parse multipart form :: 32 << 20 max 32 mb each chunk file size
	er := r.ParseMultipartForm(32 << 20)

	if er != nil {
		log.Println("Error in parsing multipart form: ", er.Error())

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			er := os.Remove(filePath)

			if er != nil {
				log.Println("Error in deleting file,", er.Error())
			}
		}

		return
	}

	// file is chunk file from request
	file, _, er := r.FormFile("file")

	if er != nil {
		log.Println("Error in fetching file from multipart form: ", er.Error())
		return
	}

	defer file.Close()

	//f is final file which combines all chunk file
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		print(err.Error())
	}

	defer f.Close()

	written, er := io.Copy(f, file)

	if er != nil {
		log.Fatal("Error in appending data in file, ", er.Error())
	}

	log.Printf("ChunkNumber : %v , Total written bytes: %v", query["resumableChunkNumber"][0], written)
}
