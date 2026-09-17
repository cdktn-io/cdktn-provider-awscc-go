// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicytemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VerifiedpermissionsPolicyTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_template#policy_store_id VerifiedpermissionsPolicyTemplate#policy_store_id}.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_template#statement VerifiedpermissionsPolicyTemplate#statement}.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_template#description VerifiedpermissionsPolicyTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy_template#name VerifiedpermissionsPolicyTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

