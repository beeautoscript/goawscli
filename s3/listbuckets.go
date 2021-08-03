package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func: Create S3 bucket
func ListS3Bucket(awsregion string) string {
	// aws session credentials
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// s3 session
	s3_session := s3.New(aws_session, aws.NewConfig().WithRegion(awsregion))

	// create bucket
	input := &s3.ListBucketsInput{}
	result, err := s3_session.ListBuckets(input)

	if err != nil {
		return err.Error()
	}
	if len(result.Buckets) == 0 {
		return "Bucket list is empty"
	}
	return result.String()
}
