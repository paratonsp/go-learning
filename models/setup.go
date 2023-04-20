package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var S3S *session.Session

const (
	AWS_S3_BUCKET     = "learning-go"
	AWS_S3_REGION     = "id-jkt-1"
	AWS_S3_ENDPOINT   = "https://is3.cloudhost.id/paratonsp-storage"
	AWS_S3_ACCESS_KEY = "00MX39J2CQKZI5O9TCJ9"
	AWS_S3_SECRET_KEY = "WHSrk1R0EQVRCFIYbvYxgZliQZUK4z4oHVQcUe0d"
)

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("paratonsp:Katon#11@tcp(103.54.170.102:3306)/learning_go"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
		fmt.Println(err.Error())
		return
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	DB = db
}

func ConnectStorage() {
	creds := credentials.NewStaticCredentials(AWS_S3_ACCESS_KEY, AWS_S3_SECRET_KEY, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials:      creds,
		Region:           aws.String(AWS_S3_REGION),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(AWS_S3_ENDPOINT),
	})
	if err != nil {
		fmt.Printf("NewSession error: %s\n", err)
		return
	}
	S3S = sess
}
