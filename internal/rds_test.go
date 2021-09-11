package internal_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/karingisi/mockproject/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockRDSClient struct {
	rdsiface.RDSAPI
	DescribeDBInstancesOutput *rds.DescribeDBInstancesOutput
	Error                     error
}

func (m mockRDSClient) DescribeDBInstances(*rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	return m.DescribeDBInstancesOutput, m.Error
}

var _ = Describe("Rds", func() {
	Describe("DescribeMyRDSInstances()", func() {
		Context("Non existing Resource Identifier", func() {
			It("should return empty DescribeDBInstancesOutput{}", func() {
				mockedOutput := &rds.DescribeDBInstancesOutput{}
				client := internal.RDSClient{
					Client: mockRDSClient{DescribeDBInstancesOutput: mockedOutput},
				}
				mockedInput := &rds.DescribeDBInstancesInput{DBInstanceIdentifier: aws.String("non-exist-rds")}
				actual, _ := client.DescribeMyRDSInstances(mockedInput)
				expected := mockedOutput
				Expect(actual).To(Equal(expected))
			})
		})

		Context("Existing Resource Identifier", func() {
			It("Should return DescribeDBInstancesOutput", func() {
				mockedOutput := &rds.DescribeDBInstancesOutput{DBInstances: []*rds.DBInstance{{DBInstanceArn: aws.String("RandomDBInstanceArn")}}}
				rdsClient := internal.RDSClient{
					Client: mockRDSClient{DescribeDBInstancesOutput: mockedOutput},
				}
				mockedInput := &rds.DescribeDBInstancesInput{DBInstanceIdentifier: aws.String("random-rds")}
				actual, _ := rdsClient.DescribeMyRDSInstances(mockedInput)
				expected := mockedOutput
				Expect(actual).To(Equal(expected))
			})
		})

		Context("Nil Resource Identifier", func() {
			It("Should return error", func() {
				rdsClient := internal.RDSClient{
					Client: mockRDSClient{Error: errors.New("some error")},
				}
				mockedInput := &rds.DescribeDBInstancesInput{}
				_, err := rdsClient.DescribeMyRDSInstances(mockedInput)
				Expect(err).To(Equal(errors.New("some error")))
			})
		})

	})
})
