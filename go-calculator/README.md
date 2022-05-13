# Go Calculator

### Description
Go script that functions as a simple calculator performing addition, subtraction, multiplication and division. The user can specify which operation is performed however many times they wish to and can end the program at any time that they are satisfied.

### Getting Started
To utilize this script, please follow the below workflow:

(1) Clone the script into your environment.\
(2) Run the following commands to ensure you have the correct dependencies: `go mod init calculator; go mod tidy`\
(3) Run `go run calculator.go` to run the script. Alternatively, run `go build calculator.go` to run as a binary file.

See below an image of the script in action:
![Image of Calculator](https://github.com/markusewalker/Misc-Go-Scripts/blob/master/go-calculator/calculator.jpg)

### Unit Testing
Along with this script, you can perform unit testing with the `calculator_test.go` script.

Simply run `go test -v` in the `src` directory as shown below:

![Testing Result](https://github.com/markusewalker/Misc-Go-Scripts/blob/master/go-calculator/tests.jpg)
