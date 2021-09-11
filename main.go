package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/karingisi/mockproject/internal"
)

func main() {

	s := session.Must(session.NewSession())
	client := rds.New(s, aws.NewConfig().WithRegion(internal.Region))

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(internal.DBIdentifier),
	}
	c := internal.RDSClient{
		Client: client,
	}
	r, err := c.DescribeMyRDSInstances(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}
