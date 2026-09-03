// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#priority BatchJobQueue#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#compute_environment_order BatchJobQueue#compute_environment_order}.
	ComputeEnvironmentOrder interface{} `field:"optional" json:"computeEnvironmentOrder" yaml:"computeEnvironmentOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#job_queue_name BatchJobQueue#job_queue_name}.
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#job_queue_type BatchJobQueue#job_queue_type}.
	JobQueueType *string `field:"optional" json:"jobQueueType" yaml:"jobQueueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#job_state_time_limit_actions BatchJobQueue#job_state_time_limit_actions}.
	JobStateTimeLimitActions interface{} `field:"optional" json:"jobStateTimeLimitActions" yaml:"jobStateTimeLimitActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#scheduling_policy_arn BatchJobQueue#scheduling_policy_arn}.
	SchedulingPolicyArn *string `field:"optional" json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#service_environment_order BatchJobQueue#service_environment_order}.
	ServiceEnvironmentOrder interface{} `field:"optional" json:"serviceEnvironmentOrder" yaml:"serviceEnvironmentOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#state BatchJobQueue#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_job_queue#tags BatchJobQueue#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

