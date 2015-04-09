package main

import (
	"os"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ec2"
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
	svc := ec2.New(&aws.Config{Region: "ap-northeast-1"})
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
			instanceID := *inst.InstanceID
			state := *inst.State.Name
			publicDNSName := *inst.PublicDNSName

			fmt.Printf("%-8s\t%-8s\t%-40s\t%-50s\n", state, instanceID, name, publicDNSName)
		}
	}
	return 0
}
