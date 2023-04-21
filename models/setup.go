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
	DB_USERNAME       = "paratonsp"
	DB_PASSWORD       = "Katon#11"
	DB_HOST           = "103.54.170.102:3306"
	DB_NAME           = "learning_go"
)

func ConnectDatabase() {
	url := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")/" + DB_NAME
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println("DB Connection Error: " + err.Error())
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
		fmt.Printf("S3 Session Error: " + err.Error())
		return
	}
	S3S = sess
}
