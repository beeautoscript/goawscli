package instances

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// func: Stop EC2 Instances
func StartEc2Instances(instanceid, awsregion string) string {
	if len(instanceid) == 0 {
		return "Please pass valid EC2 Instance Id"
	} else {
		// aws credential
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// ec2 session
		svc := ec2.New(sess, aws.NewConfig().WithRegion(awsregion))
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceid),
			},
		}
		result, err := svc.StartInstances(input)
		if err != nil {
			return err.Error()
		} else {
			return result.String()
		}

	}
}
