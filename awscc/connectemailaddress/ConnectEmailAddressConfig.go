// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectemailaddress

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectEmailAddressConfig struct {
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
	// Email address to be created for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#email_address ConnectEmailAddress#email_address}
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#instance_arn ConnectEmailAddress#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// List of alias configurations for the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#alias_configurations ConnectEmailAddress#alias_configurations}
	AliasConfigurations interface{} `field:"optional" json:"aliasConfigurations" yaml:"aliasConfigurations"`
	// A description for the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#description ConnectEmailAddress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name for the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#display_name ConnectEmailAddress#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_email_address#tags ConnectEmailAddress#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

