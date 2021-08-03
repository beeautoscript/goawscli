package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func: Create S3 bucket
func DeleteS3Bucket(name, awsregion string) string {
	// aws session credentials
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// s3 session
	s3_session := s3.New(aws_session, aws.NewConfig().WithRegion(awsregion))

	// create bucket
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(name),
	}
	result, err := s3_session.DeleteBucket(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			return aerr.Error()
		}
	}
	return result.String()
}
