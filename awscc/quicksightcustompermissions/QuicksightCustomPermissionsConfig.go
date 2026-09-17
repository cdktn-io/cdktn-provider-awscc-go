// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightcustompermissions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightCustomPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_custom_permissions#aws_account_id QuicksightCustomPermissions#aws_account_id}.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_custom_permissions#custom_permissions_name QuicksightCustomPermissions#custom_permissions_name}.
	CustomPermissionsName *string `field:"required" json:"customPermissionsName" yaml:"customPermissionsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_custom_permissions#capabilities QuicksightCustomPermissions#capabilities}.
	Capabilities *QuicksightCustomPermissionsCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_custom_permissions#tags QuicksightCustomPermissions#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

