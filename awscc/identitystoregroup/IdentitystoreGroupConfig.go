// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoregroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentitystoreGroupConfig struct {
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
	// A string containing the name of the group. This value is commonly displayed when the group is referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_group#display_name IdentitystoreGroup#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The globally unique identifier for the identity store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_group#identity_store_id IdentitystoreGroup#identity_store_id}
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// A string containing the description of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_group#description IdentitystoreGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

