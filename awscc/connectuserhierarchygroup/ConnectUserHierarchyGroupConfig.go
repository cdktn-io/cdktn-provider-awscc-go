// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuserhierarchygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUserHierarchyGroupConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_group#instance_arn ConnectUserHierarchyGroup#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the user hierarchy group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_group#name ConnectUserHierarchyGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) for the parent user hierarchy group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_group#parent_group_arn ConnectUserHierarchyGroup#parent_group_arn}
	ParentGroupArn *string `field:"optional" json:"parentGroupArn" yaml:"parentGroupArn"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_group#tags ConnectUserHierarchyGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

