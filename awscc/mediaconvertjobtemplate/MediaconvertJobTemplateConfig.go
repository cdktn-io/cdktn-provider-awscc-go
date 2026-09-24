// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconvertjobtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconvertJobTemplateConfig struct {
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
	// Specify, in JSON format, the transcoding job settings for this job template.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#settings_json MediaconvertJobTemplate#settings_json}
	SettingsJson *string `field:"required" json:"settingsJson" yaml:"settingsJson"`
	// Accelerated transcoding can significantly speed up jobs with long, visually complex content.
	//
	// Outputs that use this feature incur pro-tier pricing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#acceleration_settings MediaconvertJobTemplate#acceleration_settings}
	AccelerationSettings *MediaconvertJobTemplateAccelerationSettings `field:"optional" json:"accelerationSettings" yaml:"accelerationSettings"`
	// Optional. A category for the job template you are creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#category MediaconvertJobTemplate#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional. A description of the job template you are creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#description MediaconvertJobTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#hop_destinations MediaconvertJobTemplate#hop_destinations}
	HopDestinations interface{} `field:"optional" json:"hopDestinations" yaml:"hopDestinations"`
	// The name of the job template you are creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#name MediaconvertJobTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify the relative priority for this job.
	//
	// In any given queue, the service begins processing the job with the highest value first. When more than one job has the same priority, the service begins processing the job that you submitted first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#priority MediaconvertJobTemplate#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Optional.
	//
	// The queue that jobs created from this template are assigned to. Specify the Amazon Resource Name (ARN) of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#queue MediaconvertJobTemplate#queue}
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch Events. Set the interval, in seconds, between status updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#status_update_interval MediaconvertJobTemplate#status_update_interval}
	StatusUpdateInterval *string `field:"optional" json:"statusUpdateInterval" yaml:"statusUpdateInterval"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#tags MediaconvertJobTemplate#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

