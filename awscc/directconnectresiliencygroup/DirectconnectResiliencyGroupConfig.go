// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectresiliencygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectResiliencyGroupConfig struct {
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
	// The name of the resiliency group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/directconnect_resiliency_group#resiliency_group_name DirectconnectResiliencyGroup#resiliency_group_name}
	ResiliencyGroupName *string `field:"required" json:"resiliencyGroupName" yaml:"resiliencyGroupName"`
	// The ARNs of the connections that are members of the resiliency group.
	//
	// When specified, this is the complete, authoritative membership list: on update, connections not present are disassociated, and an explicit empty list removes all members. Omit the property to leave the existing membership unmanaged (unchanged on update). On create, an omitted or empty list means no connections are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/directconnect_resiliency_group#connection_arns DirectconnectResiliencyGroup#connection_arns}
	ConnectionArns *[]*string `field:"optional" json:"connectionArns" yaml:"connectionArns"`
	// The resiliency model that the resiliency group is intended to meet. Required when creating a resiliency group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/directconnect_resiliency_group#intended_resiliency_model DirectconnectResiliencyGroup#intended_resiliency_model}
	IntendedResiliencyModel *string `field:"optional" json:"intendedResiliencyModel" yaml:"intendedResiliencyModel"`
	// The tags associated with the resiliency group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/directconnect_resiliency_group#tags DirectconnectResiliencyGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

