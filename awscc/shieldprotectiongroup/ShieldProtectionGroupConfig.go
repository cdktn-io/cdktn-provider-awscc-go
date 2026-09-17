// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldprotectiongroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShieldProtectionGroupConfig struct {
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
	// Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.
	//
	// * Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.
	// * Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.
	// * Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront and origin resources for CloudFront distributions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#aggregation ShieldProtectionGroup#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The criteria to use to choose the protected resources for inclusion in the group.
	//
	// You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#pattern ShieldProtectionGroup#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name of the protection group.
	//
	// You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#protection_group_id ShieldProtectionGroup#protection_group_id}
	ProtectionGroupId *string `field:"required" json:"protectionGroupId" yaml:"protectionGroupId"`
	// The Amazon Resource Names (ARNs) of the resources to include in the protection group.
	//
	// You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#members ShieldProtectionGroup#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// The resource type to include in the protection group.
	//
	// All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#resource_type ShieldProtectionGroup#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// One or more tag key-value pairs for the Protection object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/shield_protection_group#tags ShieldProtectionGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

