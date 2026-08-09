// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectInstanceConfig struct {
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
	// The attributes for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_instance#attributes ConnectInstance#attributes}
	Attributes *ConnectInstanceAttributes `field:"required" json:"attributes" yaml:"attributes"`
	// Specifies the type of directory integration for new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_instance#identity_management_type ConnectInstance#identity_management_type}
	IdentityManagementType *string `field:"required" json:"identityManagementType" yaml:"identityManagementType"`
	// Existing directoryId user wants to map to the new Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_instance#directory_id ConnectInstance#directory_id}
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Alias of the new directory created as part of new instance creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_instance#instance_alias ConnectInstance#instance_alias}
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_instance#tags ConnectInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

