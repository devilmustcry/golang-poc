package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"time"
)

func main() {
	region := "ap-southeast-1"

	// AWS configuration
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	// If not in prod environment, set credentials from configuration

	accessKey := ""
	secretKey := ""
	awsConfig.Credentials = credentials.NewStaticCredentials(accessKey, secretKey, "")

	// In prod environment, SDK will use IAM role attached to the EC2, ECS, etc.

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test")
	start := time.Now()
	end := start.Add(time.Second * 30)
	milli := end.UnixMilli() - start.UnixMilli()

	cloudClient := cloudwatch.New(sess)
	testDatum := cloudwatch.MetricDatum{
		Counts:            nil,
		Dimensions:        nil,
		MetricName:        aws.String("completionTime"),
		StatisticValues:   nil,
		StorageResolution: nil,
		Timestamp:         aws.Time(time.Now()),
		Unit:              aws.String("Milliseconds"),
		Value:             aws.Float64(float64(milli)),
		Values:            nil,
	}
	var metricDatas []*cloudwatch.MetricDatum
	metricDatas = append(metricDatas, &testDatum)
	cloudClient.PutMetricData(&cloudwatch.PutMetricDataInput{
		MetricData: metricDatas,
		Namespace:  aws.String("DMIND_SERVICE"),
	})
}
