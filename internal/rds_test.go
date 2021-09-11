package internal_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/karingisi/mockproject/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// MockRDS is a mock RDSAPI implementation
type mockRDSClient struct {
	rdsiface.RDSAPI
}

func (m *mockRDSClient) DescribeDBInstances(*rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	return &rds.DescribeDBInstancesOutput{DBInstances: []*rds.DBInstance{{DBInstanceArn: aws.String("DBInstanceArn")}}}, nil
}

var _ = Describe("Rds", func() {
	Describe("DescribeMyRDSInstances()", func() {
		Context("Random Resource Identifier", func() {
			It("should not crush", func() {
				mockRDSClient := &mockRDSClient{}

				mockedInput := &rds.DescribeDBInstancesInput{DBInstanceIdentifier: aws.String("siltest-rds")}
				actual, _ := internal.DescribeMyRDSInstances(mockRDSClient, "Random-ResourceIdentifier", mockedInput)
				expected := &rds.DescribeDBInstancesOutput{DBInstances: []*rds.DBInstance{{DBInstanceArn: aws.String("DBInstanceArn")}}}
				//err := errors.New("DBInstanceNotFound")
				//Expect(err).To(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
