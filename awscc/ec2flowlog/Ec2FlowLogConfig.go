// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2flowlog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2FlowLogConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#resource_id Ec2FlowLog#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The type of resource for which to create the flow log.
	//
	// For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#resource_type Ec2FlowLog#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#deliver_cross_account_role Ec2FlowLog#deliver_cross_account_role}
	DeliverCrossAccountRole *string `field:"optional" json:"deliverCrossAccountRole" yaml:"deliverCrossAccountRole"`
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account.
	//
	// If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#deliver_logs_permission_arn Ec2FlowLog#deliver_logs_permission_arn}
	DeliverLogsPermissionArn *string `field:"optional" json:"deliverLogsPermissionArn" yaml:"deliverLogsPermissionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#destination_options Ec2FlowLog#destination_options}.
	DestinationOptions *Ec2FlowLogDestinationOptions `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// Specifies the destination to which the flow log data is to be published.
	//
	// Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#log_destination Ec2FlowLog#log_destination}
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#log_destination_type Ec2FlowLog#log_destination_type}
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#log_format Ec2FlowLog#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.
	//
	// If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#log_group_name Ec2FlowLog#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#max_aggregation_interval Ec2FlowLog#max_aggregation_interval}
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The resource types and associated tags for EC2 resources associated with the EC2 Tags feature for log enrichment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#tag_field_specifications Ec2FlowLog#tag_field_specifications}
	TagFieldSpecifications interface{} `field:"optional" json:"tagFieldSpecifications" yaml:"tagFieldSpecifications"`
	// The tags to apply to the flow logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#tags Ec2FlowLog#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#traffic_type Ec2FlowLog#traffic_type}
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

