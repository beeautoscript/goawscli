package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func: Create S3 bucket
func CreateS3Bucket(name, awsregion string) string {
	// aws session credentials
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// s3 session
	s3_session := s3.New(aws_session, aws.NewConfig().WithRegion(awsregion))

	// create bucket
	input := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}
	result, err := s3_session.CreateBucket(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			return aerr.Error()
		}
	}
	return result.String()
}
