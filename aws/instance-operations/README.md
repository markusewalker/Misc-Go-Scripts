# AWS Instance Operations

### Description
Go script that performs the following AWS instance operations:

- Start instances
- Stop instances
- Reboot instances
- Terminate instances
- List instances

### Prerequisites
You will need to have the AWS CLI already installed on your client machine. Additionally, you will need to already have run the command `aws configure` and properly fill out the information to connect to your AWS account.

### Getting Started
To utilize this script, please follow the below workflow:

1. Clone the script into your environment.
2. Run the following commands to ensure you have the correct dependencies: `go mod init instance`; `go mod tidy`.
3. To start, stop, reboot or terminate an instance, run command: `go run instance.go <start|stop|reboot|terminate> <AMI ID>`.
4. To list an instance, run command: `go run instance.go list <name of instance>`.
     - Currently, you will modify line 133 to the name of the instance you wish to list, otherwise this operation will not work.