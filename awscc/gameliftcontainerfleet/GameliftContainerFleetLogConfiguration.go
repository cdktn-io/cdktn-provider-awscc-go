// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetLogConfiguration struct {
	// Configures the service that provides logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet#log_destination GameliftContainerFleet#log_destination}
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// If log destination is CLOUDWATCH, logs are sent to the specified log group in Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet#log_group_arn GameliftContainerFleet#log_group_arn}
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// The name of the S3 bucket to pull logs from if S3 is the LogDestination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet#s3_bucket_name GameliftContainerFleet#s3_bucket_name}
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
}

