package handlers

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

type Files struct{}

func NewFiles() *Files {
	return &Files{}
}

func (f *Files) SaveFile(rw http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	openedFile, err := os.OpenFile("files/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	written, err := io.Copy(openedFile, file)
	if err != nil {
		return
	}

	rw.Write([]byte(header.Filename + " is uploaded. Written " + strconv.Itoa(int(written)) + "bytes"))
	defer file.Close()
}

func (f *Files) DownloadFile(rw http.ResponseWriter, r *http.Request) {

}
