// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargetsS3Targets struct {
	// The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#dlq_event_queue_arn GlueCrawler#dlq_event_queue_arn}
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#event_queue_arn GlueCrawler#event_queue_arn}
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#exclusions GlueCrawler#exclusions}
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path to the Amazon S3 target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#path GlueCrawler#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset.
	//
	// If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#sample_size GlueCrawler#sample_size}
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}

