// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#display_name DeadlineQueue#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#farm_id DeadlineQueue#farm_id}.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#allowed_storage_profile_ids DeadlineQueue#allowed_storage_profile_ids}.
	AllowedStorageProfileIds *[]*string `field:"optional" json:"allowedStorageProfileIds" yaml:"allowedStorageProfileIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#default_budget_action DeadlineQueue#default_budget_action}.
	DefaultBudgetAction *string `field:"optional" json:"defaultBudgetAction" yaml:"defaultBudgetAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#description DeadlineQueue#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#job_attachment_settings DeadlineQueue#job_attachment_settings}.
	JobAttachmentSettings *DeadlineQueueJobAttachmentSettings `field:"optional" json:"jobAttachmentSettings" yaml:"jobAttachmentSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#job_run_as_user DeadlineQueue#job_run_as_user}.
	JobRunAsUser *DeadlineQueueJobRunAsUser `field:"optional" json:"jobRunAsUser" yaml:"jobRunAsUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#required_file_system_location_names DeadlineQueue#required_file_system_location_names}.
	RequiredFileSystemLocationNames *[]*string `field:"optional" json:"requiredFileSystemLocationNames" yaml:"requiredFileSystemLocationNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#role_arn DeadlineQueue#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#scheduling_configuration DeadlineQueue#scheduling_configuration}.
	SchedulingConfiguration *DeadlineQueueSchedulingConfiguration `field:"optional" json:"schedulingConfiguration" yaml:"schedulingConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#tags DeadlineQueue#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

