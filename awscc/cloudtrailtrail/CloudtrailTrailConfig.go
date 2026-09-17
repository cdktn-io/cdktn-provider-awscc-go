// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrailtrail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudtrailTrailConfig struct {
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
	// Whether the CloudTrail is currently logging AWS API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#is_logging CloudtrailTrail#is_logging}
	IsLogging interface{} `field:"required" json:"isLogging" yaml:"isLogging"`
	// Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#s3_bucket_name CloudtrailTrail#s3_bucket_name}
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The advanced event selectors that were used to select events for the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#advanced_event_selectors CloudtrailTrail#advanced_event_selectors}
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// Specifies the aggregation configuration to aggregate CloudTrail Events. A maximum of 1 aggregation configuration is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#aggregation_configurations CloudtrailTrail#aggregation_configurations}
	AggregationConfigurations interface{} `field:"optional" json:"aggregationConfigurations" yaml:"aggregationConfigurations"`
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.
	//
	// Not required unless you specify CloudWatchLogsRoleArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#cloudwatch_logs_log_group_arn CloudtrailTrail#cloudwatch_logs_log_group_arn}
	CloudwatchLogsLogGroupArn *string `field:"optional" json:"cloudwatchLogsLogGroupArn" yaml:"cloudwatchLogsLogGroupArn"`
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#cloudwatch_logs_role_arn CloudtrailTrail#cloudwatch_logs_role_arn}
	CloudwatchLogsRoleArn *string `field:"optional" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Specifies whether log file validation is enabled. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#enable_log_file_validation CloudtrailTrail#enable_log_file_validation}
	EnableLogFileValidation interface{} `field:"optional" json:"enableLogFileValidation" yaml:"enableLogFileValidation"`
	// Use event selectors to further specify the management and data event settings for your trail.
	//
	// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#event_selectors CloudtrailTrail#event_selectors}
	EventSelectors interface{} `field:"optional" json:"eventSelectors" yaml:"eventSelectors"`
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#include_global_service_events CloudtrailTrail#include_global_service_events}
	IncludeGlobalServiceEvents interface{} `field:"optional" json:"includeGlobalServiceEvents" yaml:"includeGlobalServiceEvents"`
	// Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#insight_selectors CloudtrailTrail#insight_selectors}
	InsightSelectors interface{} `field:"optional" json:"insightSelectors" yaml:"insightSelectors"`
	// Specifies whether the trail applies only to the current region or to all regions.
	//
	// The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#is_multi_region_trail CloudtrailTrail#is_multi_region_trail}
	IsMultiRegionTrail interface{} `field:"optional" json:"isMultiRegionTrail" yaml:"isMultiRegionTrail"`
	// Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account.
	//
	// The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#is_organization_trail CloudtrailTrail#is_organization_trail}
	IsOrganizationTrail interface{} `field:"optional" json:"isOrganizationTrail" yaml:"isOrganizationTrail"`
	// Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#kms_key_id CloudtrailTrail#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	//
	// For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#s3_key_prefix CloudtrailTrail#s3_key_prefix}
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic defined for notification of log file delivery.
	//
	// The maximum length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#sns_topic_name CloudtrailTrail#sns_topic_name}
	SnsTopicName *string `field:"optional" json:"snsTopicName" yaml:"snsTopicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#tags CloudtrailTrail#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#trail_name CloudtrailTrail#trail_name}.
	TrailName *string `field:"optional" json:"trailName" yaml:"trailName"`
}

