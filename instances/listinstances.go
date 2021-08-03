package instances

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// func: List EC2 Instances
func ListEc2Instances(awsregion string) string {

	// table header
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Tag"},
			{Align: simpletable.AlignCenter, Text: "Instance Id"},
			{Align: simpletable.AlignCenter, Text: "PublicIpAddress"},
			{Align: simpletable.AlignCenter, Text: "State"},
		},
	}

	// aws session credentials
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// ec2 session
	ec2_session := ec2.New(aws_session, aws.NewConfig().WithRegion(awsregion))

	// ec2 describe instance
	result, err := ec2_session.DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		return err.Error()
	} else {
		if len(result.Reservations) == 0 {
			return "EC2 Instances not available"
		} else {
			for row, i := range result.Reservations {
				var tag string
				var publicip string

				// tag
				if len(i.Instances[0].Tags) == 0 {
					tag = "NA"
				} else {
					tag = *i.Instances[0].Tags[0].Value
				}

				// publicip
				if len(*i.Instances[0].PublicDnsName) == 0 {
					publicip = "NA"
				} else {
					publicip = *i.Instances[0].NetworkInterfaces[0].Association.PublicIp
				}

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row)},
					{Align: simpletable.AlignRight, Text: tag},
					{Align: simpletable.AlignRight, Text: *i.Instances[0].InstanceId},
					{Align: simpletable.AlignRight, Text: publicip},
					{Align: simpletable.AlignRight, Text: *i.Instances[0].State.Name},
				}
				table.Body.Cells = append(table.Body.Cells, r)
			}
			table.SetStyle(simpletable.StyleCompactLite)
			fmt.Printf("Listing EC2 instances from %s region\n\n", awsregion)
			return table.String()
		}
	}
}
