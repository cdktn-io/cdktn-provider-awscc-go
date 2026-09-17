// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeueenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineQueueEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue_environment#farm_id DeadlineQueueEnvironment#farm_id}.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue_environment#priority DeadlineQueueEnvironment#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue_environment#queue_id DeadlineQueueEnvironment#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue_environment#template DeadlineQueueEnvironment#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue_environment#template_type DeadlineQueueEnvironment#template_type}.
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
}

