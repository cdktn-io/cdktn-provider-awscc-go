// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheuser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticacheUserConfig struct {
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
	// The target cache engine for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#engine ElasticacheUser#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#user_id ElasticacheUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#user_name ElasticacheUser#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Access permissions string used for this user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#access_string ElasticacheUser#access_string}
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#authentication_mode ElasticacheUser#authentication_mode}.
	AuthenticationMode *ElasticacheUserAuthenticationMode `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Indicates a password is not required for this user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#no_password_required ElasticacheUser#no_password_required}
	NoPasswordRequired interface{} `field:"optional" json:"noPasswordRequired" yaml:"noPasswordRequired"`
	// Passwords used for this user account. You can create up to two passwords for each user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#passwords ElasticacheUser#passwords}
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// An array of key-value pairs to apply to this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_user#tags ElasticacheUser#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

