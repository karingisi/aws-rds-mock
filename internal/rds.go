package internal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	log "github.com/sirupsen/logrus"
)

const (
	Region       = "ap-southeast-2"
	DBIdentifier = "iview-testing-9-primary"
)

func NewRDSClient() rdsiface.RDSAPI {
	s := session.Must(session.NewSession())
	client := rds.New(s, aws.NewConfig().WithRegion(Region))

	return client
}

func DescribeMyRDSInstances(r rdsiface.RDSAPI, resourceIdentifier string, input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	fmt.Println("Describe My RDS Instance")
	result, err := r.DescribeDBInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBInstanceNotFoundFault:
				log.WithFields(log.Fields{"err": aerr.Error()}).Error(rds.ErrCodeDBInstanceNotFoundFault)
			default:
				log.Error(err.Error())
			}
		} else {
			log.Error(err.Error())
		}
		return nil, err
	}

	return result, nil
}
