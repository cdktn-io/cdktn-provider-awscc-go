// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontrealtimelogconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontRealtimeLogConfigConfig struct {
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
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_realtime_log_config#end_points CloudfrontRealtimeLogConfig#end_points}
	EndPoints interface{} `field:"required" json:"endPoints" yaml:"endPoints"`
	// A list of fields that are included in each real-time log record.
	//
	// In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.
	//  For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_realtime_log_config#fields CloudfrontRealtimeLogConfig#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The unique name of this real-time log configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_realtime_log_config#name CloudfrontRealtimeLogConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The sampling rate for this real-time log configuration.
	//
	// The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_realtime_log_config#sampling_rate CloudfrontRealtimeLogConfig#sampling_rate}
	SamplingRate *float64 `field:"required" json:"samplingRate" yaml:"samplingRate"`
}

