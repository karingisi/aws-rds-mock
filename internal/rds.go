package internal

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	log "github.com/sirupsen/logrus"
)

const (
	Region       = "ap-southeast-2"
	DBIdentifier = "some-db-identifier"
)

type RDSClient struct {
	Client rdsiface.RDSAPI
}

func (r *RDSClient) DescribeMyRDSInstances(input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	result, err := r.Client.DescribeDBInstances(input)
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
