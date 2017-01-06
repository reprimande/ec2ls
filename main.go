package main

import (
	"os"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const Version string = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "ec2ls"
	app.Usage = "ls for AWS EC2"
	app.Version = Version
	app.Author = "Naoki Nomoto"
	app.Email = "happygrind@gmail.com"
	app.Action = func(c *cli.Context) {
		run()
	}

	app.Run(os.Args)
}

func run() int {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	for _, res := range resp.Reservations {
		for _, inst := range res.Instances {
			var name = ""
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			instanceID := *inst.InstanceId
			state := *inst.State.Name

			var publicDNSName = ""
			var publicIpAddress = ""
			var privateIpAddress = ""
			if (inst.PublicDnsName != nil) {
				publicDNSName = *inst.PublicDnsName
			}

			if (inst.PublicIpAddress != nil) {
				publicIpAddress = *inst.PublicIpAddress
			}

			if (inst.PrivateIpAddress != nil) {
				privateIpAddress = *inst.PrivateIpAddress
			}

			fmt.Printf(
				"\x1b[33m%-8s\x1b[0m\t%-20s\t%-30s\t%-15s\t%-50s\t%-15s\n",
				state, instanceID, name, publicIpAddress, publicDNSName, privateIpAddress)
		}
	}
	return 0
}
