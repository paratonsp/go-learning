package helper

import (
	"fmt"
	"io"
	"learning-go/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func FileUploadLocal(r *http.Request, s string) (string, error) {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("assets")
	if err != nil {
		return "", err
	}
	defer file.Close()

	path := filepath.Join(".", "assets/"+s)
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + handler.Filename
	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()

	io.Copy(f, file)

	return fullPath, nil
}

func FileUploadS3(session *session.Session, r *http.Request, s string) (string, error) {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("assets")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	path := models.AWS_S3_BUCKET + "-" + s
	fullPath := models.AWS_S3_ENDPOINT + "/" + path + "/" + handler.Filename

	_, err = s3manager.NewUploader(session).Upload(&s3manager.UploadInput{
		Bucket: aws.String(path),
		ACL:    aws.String("public-read"),
		Key:    aws.String(handler.Filename),
		Body:   file,
	})

	return fullPath, err
}
