// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesdedicatedippool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesDedicatedIpPoolConfig struct {
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
	// The name of the dedicated IP pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_dedicated_ip_pool#pool_name SesDedicatedIpPool#pool_name}
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_dedicated_ip_pool#scaling_mode SesDedicatedIpPool#scaling_mode}
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
	// The tags (keys and values) associated with the dedicated IP pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_dedicated_ip_pool#tags SesDedicatedIpPool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

