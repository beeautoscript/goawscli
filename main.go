package main

import (
	"fmt"
	"goawscli/instances"
	"goawscli/s3"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// func: main
func main() {
	// application cli setup
	app := cli.NewApp()

	app.Name = "Go AWS Cli"
	app.Usage = "Manage and Monitor AWS resources from cli"

	// EC2 Flags command
	// region flag
	regionFlag := []cli.Flag{
		&cli.StringFlag{
			Name:     "region",
			Value:    "us-east-1",
			Usage:    "aws region name",
			Required: true,
		},
	}
	// instance id flag
	instanceid_flag := []cli.Flag{
		&cli.StringFlag{
			Name:     "instanceid",
			Usage:    "ec2 instance id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "region",
			Value:    "us-east-1",
			Usage:    "aws region name",
			Required: true,
		},
	}
	// csv report file name
	csv_flag := []cli.Flag{
		&cli.StringFlag{
			Name:     "csv",
			Value:    "example.xlsx",
			Usage:    "name of csv file to be saved.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "region",
			Value:    "us-east-1",
			Usage:    "aws region name",
			Required: true,
		},
	}
	// S3 Flags
	bucket_name_flag := []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "new bucket name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "region",
			Value:    "us-east-1",
			Usage:    "aws region name",
			Required: true,
		},
	}
	// application commands
	app.Commands = []*cli.Command{
		{
			Name:  "ec2",
			Usage: "Manage ec2 instances",
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "List all ec2 instances",
					Flags: regionFlag,
					Action: func(c *cli.Context) error {
						ec2_list_result := instances.ListEc2Instances(c.String("region"))
						fmt.Println(ec2_list_result)
						return nil
					},
				},
				{
					Name:  "stop",
					Usage: "Stop running ec2 instance",
					Flags: instanceid_flag,
					Action: func(c *cli.Context) error {
						ec2_stop_instance := instances.StopEc2Instances(c.String("instanceid"), c.String("region"))
						fmt.Println(ec2_stop_instance)
						return nil
					},
				},
				{
					Name:  "start",
					Usage: "Start stopped ec2 instance",
					Flags: instanceid_flag,
					Action: func(c *cli.Context) error {
						ec2_start_instance := instances.StartEc2Instances(c.String("instanceid"), c.String("region"))
						fmt.Println(ec2_start_instance)
						return nil
					},
				},
				{
					Name:  "reboot",
					Usage: "Reboot running ec2 instance",
					Flags: instanceid_flag,
					Action: func(c *cli.Context) error {
						ec2_reboot_instance := instances.RebootEc2Instances(c.String("instanceid"), c.String("region"))
						fmt.Println(ec2_reboot_instance)
						return nil
					},
				},
				{
					Name:  "terminate",
					Usage: "Terminate ec2 instance",
					Flags: instanceid_flag,
					Action: func(c *cli.Context) error {
						ec2_terminate_instance := instances.TerminateEc2Instances(c.String("instanceid"), c.String("region"))
						fmt.Println(ec2_terminate_instance)
						return nil
					},
				},
				{
					Name:  "report",
					Usage: "Generate csv report of ec2 instances",
					Flags: csv_flag,
					Action: func(c *cli.Context) error {
						ec2_csv_report := instances.GenerateCsvReportInstances(c.String("csv"), c.String("region"))
						fmt.Println(ec2_csv_report)
						return nil
					},
				},
			},
		},
		{
			Name:  "s3",
			Usage: "Manage S3",
			Subcommands: []*cli.Command{
				{
					Name:  "new",
					Usage: "Create new bucket",
					Flags: bucket_name_flag,
					Action: func(c *cli.Context) error {
						s3_bucket_create := s3.CreateS3Bucket(c.String("name"), c.String("region"))
						fmt.Println(s3_bucket_create)
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "Delete bucket",
					Flags: bucket_name_flag,
					Action: func(c *cli.Context) error {
						s3_bucket_delete := s3.DeleteS3Bucket(c.String("name"), c.String("region"))
						fmt.Println(s3_bucket_delete)
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "List buckets",
					Flags: regionFlag,
					Action: func(c *cli.Context) error {
						s3_bucket_list := s3.ListS3Bucket(c.String("region"))
						fmt.Println(s3_bucket_list)
						return nil
					},
				},
			},
		},
	}

	// run application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
