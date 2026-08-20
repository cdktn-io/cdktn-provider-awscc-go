// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchquotashare

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchQuotaShareConfig struct {
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
	// A list that specifies the quantity and type of compute capacity allocated to the quota share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#capacity_limits BatchQuotaShare#capacity_limits}
	CapacityLimits interface{} `field:"required" json:"capacityLimits" yaml:"capacityLimits"`
	// The AWS Batch job queue associated with the quota share.
	//
	// This can be the job queue name or ARN. A job queue must be in the `VALID` state before you can associate it with a quota share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#job_queue BatchQuotaShare#job_queue}
	JobQueue *string `field:"required" json:"jobQueue" yaml:"jobQueue"`
	// Specifies the preemption behavior for jobs in a quota share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#preemption_configuration BatchQuotaShare#preemption_configuration}
	PreemptionConfiguration *BatchQuotaSharePreemptionConfiguration `field:"required" json:"preemptionConfiguration" yaml:"preemptionConfiguration"`
	// The name of the quota share.
	//
	// It can be up to 128 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#quota_share_name BatchQuotaShare#quota_share_name}
	QuotaShareName *string `field:"required" json:"quotaShareName" yaml:"quotaShareName"`
	// Specifies whether a quota share reserves, lends, or both lends and borrows idle compute capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#resource_sharing_configuration BatchQuotaShare#resource_sharing_configuration}
	ResourceSharingConfiguration *BatchQuotaShareResourceSharingConfiguration `field:"required" json:"resourceSharingConfiguration" yaml:"resourceSharingConfiguration"`
	// The state of the quota share.
	//
	// If the quota share is `ENABLED`, it is able to accept jobs. If the quota share is `DISABLED`, new jobs won't be accepted but jobs already submitted can finish. The default state is `ENABLED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#state BatchQuotaShare#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags that you apply to the quota share to help you categorize and organize your resources.
	//
	// Each tag consists of a key and an optional value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#tags BatchQuotaShare#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

