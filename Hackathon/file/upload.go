package file

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"time"
)

const MaxFileSize = 8 << 20 // 8MB

type File struct {
	Name        string
	ContentType string
	Size        int64
}

// Upload is http function for process uploading image
func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		outputHTML(w, r, "./html/upload.html")
	}

	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, MaxFileSize)

	// receive file upload
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")

	// validate
	if !isImage(contentType) {
		http.Error(w, "file must be image", http.StatusForbidden)
		return
	}

	f := File{
		Name:        fileHeader.Filename,
		ContentType: contentType,
		Size:        fileHeader.Size,
	}

	// insert meta data to database
	err = insert(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeFile(file, f.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("upload file success"))
}

// insert save image metada to database
func insert(f File) error {
	db, err := sql.Open("mysql", "test:test@tcp(mysql:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := "INSERT INTO file(name, content_type, size) VALUES (?, ?, ?)"

	_, err = db.Query(q, f.Name, f.ContentType, f.Size)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// writeFile write received file to tmp folder
func writeFile(f multipart.File, originalFileName string) error {
	// named file before saving
	fileName := fmt.Sprintf("./tmp/%d%s", time.Now().Unix(), originalFileName)

	// open a file for writing
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// copy uploaded file's content
	_, err = io.Copy(file, f)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// isImage check if a file is image or not
func isImage(contentType string) bool {
	switch {
	case regexp.MustCompile("image/*").Match([]byte(contentType)):
		return true
	default:
		return false
	}
}

// outputHTML serve html content
func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
