package wox

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Aws ..
type Aws struct {
	S3 *s3.S3
}

//NewAws instanciate a new aws struct
func NewAws() Aws {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := s3.New(sess)

	aws := Aws{S3: svc}

	return aws

}
