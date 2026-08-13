// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterLoggingInfoBrokerLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#cloudwatch_logs MskCluster#cloudwatch_logs}.
	CloudwatchLogs *MskClusterLoggingInfoBrokerLogsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#firehose MskCluster#firehose}.
	Firehose *MskClusterLoggingInfoBrokerLogsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#s3 MskCluster#s3}.
	S3 *MskClusterLoggingInfoBrokerLogsS3 `field:"optional" json:"s3" yaml:"s3"`
}

