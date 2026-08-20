// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeuelimitassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineQueueLimitAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_queue_limit_association#farm_id DeadlineQueueLimitAssociation#farm_id}.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_queue_limit_association#limit_id DeadlineQueueLimitAssociation#limit_id}.
	LimitId *string `field:"required" json:"limitId" yaml:"limitId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_queue_limit_association#queue_id DeadlineQueueLimitAssociation#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

