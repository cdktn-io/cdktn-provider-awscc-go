// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbuser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MemorydbUserConfig struct {
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
	// The name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_user#user_name MemorydbUser#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Access permissions string used for this user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_user#access_string MemorydbUser#access_string}
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_user#authentication_mode MemorydbUser#authentication_mode}.
	AuthenticationMode *MemorydbUserAuthenticationMode `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// An array of key-value pairs to apply to this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_user#tags MemorydbUser#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

