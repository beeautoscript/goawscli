package instances

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/xuri/excelize/v2"
)

func GenerateCsvReportInstances(csvname, awsregion string) string {
	// excel report
	excel_f := excelize.NewFile()
	index := excel_f.NewSheet("Sheet1")
	// aws credential
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// ec2 session
	svc := ec2.New(sess, aws.NewConfig().WithRegion(awsregion))
	// ec2 describe instance
	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		return err.Error()
	} else {
		if len(result.Reservations) == 0 {
			return "EC2 Instances not available"
		} else {
			// create excel data
			excel_f.SetCellValue("Sheet1", "A1", "Tag")
			excel_f.SetCellValue("Sheet1", "B1", "Intance ID")
			excel_f.SetCellValue("Sheet1", "C1", "Public IPv4 Address")
			excel_f.SetCellValue("Sheet1", "D1", "State")
			excel_f.SetActiveSheet(index)

			for row, i := range result.Reservations {
				// tag
				var tag string
				if len(i.Instances[0].Tags) == 0 {
					tag = "NA"
				} else {
					tag = *i.Instances[0].Tags[0].Value
				}
				// public ip address
				var publicip string
				if len(*i.Instances[0].PublicDnsName) == 0 {
					publicip = "NA"
				} else {
					publicip = *i.Instances[0].NetworkInterfaces[0].Association.PublicIp
				}
				excel_f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", "A", row+2), tag)
				excel_f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", "B", row+2), *i.Instances[0].InstanceId)
				excel_f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", "C", row+2), publicip)
				excel_f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", "D", row+2), *i.Instances[0].State.Name)
				excel_f.SetActiveSheet(index)
			}

			if err := excel_f.SaveAs(csvname); err != nil {
				return err.Error()
			}

		}
	}
	return "CSV Report Generated Successfully"
}
