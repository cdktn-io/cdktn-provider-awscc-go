// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportappaccountalias

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SupportappAccountAliasConfig struct {
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
	// An account alias associated with a customer's account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/supportapp_account_alias#account_alias SupportappAccountAlias#account_alias}
	AccountAlias *string `field:"required" json:"accountAlias" yaml:"accountAlias"`
}

