// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchschedulingpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchSchedulingPolicyConfig struct {
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
	// Fair Share Policy for the Job Queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_scheduling_policy#fairshare_policy BatchSchedulingPolicy#fairshare_policy}
	FairsharePolicy *BatchSchedulingPolicyFairsharePolicy `field:"optional" json:"fairsharePolicy" yaml:"fairsharePolicy"`
	// Name of Scheduling Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_scheduling_policy#name BatchSchedulingPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Quota Share Policy for the Job Queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_scheduling_policy#quota_share_policy BatchSchedulingPolicy#quota_share_policy}
	QuotaSharePolicy *BatchSchedulingPolicyQuotaSharePolicy `field:"optional" json:"quotaSharePolicy" yaml:"quotaSharePolicy"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_scheduling_policy#tags BatchSchedulingPolicy#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

