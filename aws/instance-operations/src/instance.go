// Inspired By		: https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/example_code/ec2/start_stop_instances.go
// Date Modified	: 7/20/22

// Description		: Go script to handle multiple operations involving Amazon EC2 instances.

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var sess = session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
var svc = ec2.New(sess)

var ami = os.Args[2]

func startInstance() {
	// Does a check to verify if you have the needed privilges needed to start the instance.
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(ami),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is DryRunOperation, then set DryRun to false to start the instance.
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StartInstances(input)
		if err != nil {
			fmt.Println("ERROR. Unable to start the instance...", err)
		} else {
			fmt.Println("Starting the instance...")
			fmt.Println("\nSTATUS of AMI", ami, ":\n\n", result.StartingInstances)
		}
	} else {
		fmt.Println("ERROR. Unable to start the instance.", err)
	}
}

func stopInstance() {
	// Does a check to verify if you have the needed privilges needed to stop the instance.
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(ami),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is DryRunOperation, then set DryRun to false to stop the instance.
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StopInstances(input)
		if err != nil {
			fmt.Println("ERROR. Unable to stop the instance.", err)
		} else {
			fmt.Println("Stopping the instance...")
			fmt.Println("\nSTATUS of AMI", ami, ":\n\n", result.StoppingInstances)
		}
	} else {
		fmt.Println("ERROR. Unable to stop the instance.", err)
	}
}

func rebootInstance() {
	// Does a check to verify if you have the needed privilges needed to reboot the instance.
	input := &ec2.RebootInstancesInput{
		InstanceIds: []*string{
			aws.String(ami),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.RebootInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is DryRunOperation, then set DryRun to false to reboot the instance.
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.RebootInstances(input)
		if err != nil {
			fmt.Println("ERROR. Unable to reboot the instance.", err)
		} else {
			fmt.Println("Rebooting the instance...")
			fmt.Println("\nSTATUS of AMI", ami, ":\n\n", result)
		}
	} else {
		fmt.Println("ERROR. Unable to reboot the instance.", err)
	}
}

func terminateInstance() {
	// Does a check to verify if you have the needed privilges needed to terminate the instance.
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			aws.String(ami),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.TerminateInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is DryRunOperation, then set DryRun to false to reboot the instance.
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.TerminateInstances(input)
		if err != nil {
			fmt.Println("ERROR. Unable to terminate the instance.", err)
		} else {
			fmt.Println("Terminating the instance...")
			fmt.Println("\nSTATUS of AMI", ami, ":\n\n", result)
		}
	} else {
		fmt.Println("ERROR. Unable to reboot the instance.", err)
	}
}

func listInstance() {
	// Changes the input to the list of instances to what you would like to see.
	name := "mew-*"
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
		},
	}

	res, _ := svc.DescribeInstances(params)

	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("\nINSTANCE:", *instance.InstanceId)
			fmt.Println("\tSTATE:", *instance.State.Name)

			for _, tag := range instance.Tags {
				fmt.Println("\tTAG:", *tag.Key, *tag.Value)
			}
		}
	}
}

func main() {
	fmt.Println("\x1B[96m==========================================")
	fmt.Println("\t AWS Instance Operations")
	fmt.Println("==========================================\x1B[0m")

	choice := os.Args[1]

	switch {
	case choice == "start":
		startInstance()
	case choice == "stop":
		stopInstance()
	case choice == "reboot":
		rebootInstance()
	case choice == "terminate":
		terminateInstance()
	case choice == "list":
		listInstance()
	}
}
