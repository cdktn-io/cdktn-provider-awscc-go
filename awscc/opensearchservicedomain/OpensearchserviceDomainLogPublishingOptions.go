// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainLogPublishingOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#cloudwatch_logs_log_group_arn OpensearchserviceDomain#cloudwatch_logs_log_group_arn}.
	CloudwatchLogsLogGroupArn *string `field:"optional" json:"cloudwatchLogsLogGroupArn" yaml:"cloudwatchLogsLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

