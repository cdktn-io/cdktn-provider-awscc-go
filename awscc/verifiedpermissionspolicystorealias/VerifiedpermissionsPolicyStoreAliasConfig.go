// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicystorealias

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VerifiedpermissionsPolicyStoreAliasConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_store_alias#alias_name VerifiedpermissionsPolicyStoreAlias#alias_name}.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_store_alias#policy_store_id VerifiedpermissionsPolicyStoreAlias#policy_store_id}.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

