// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstanceuser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeAppInstanceUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#app_instance_arn ChimeAppInstanceUser#app_instance_arn}.
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#app_instance_user_id ChimeAppInstanceUser#app_instance_user_id}.
	AppInstanceUserId *string `field:"required" json:"appInstanceUserId" yaml:"appInstanceUserId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#expiration_settings ChimeAppInstanceUser#expiration_settings}.
	ExpirationSettings *ChimeAppInstanceUserExpirationSettings `field:"optional" json:"expirationSettings" yaml:"expirationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#metadata ChimeAppInstanceUser#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#name ChimeAppInstanceUser#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#tags ChimeAppInstanceUser#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

