// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstancebot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeAppInstanceBotConfig struct {
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
	// The ARN of the AppInstance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/chime_app_instance_bot#app_instance_arn ChimeAppInstanceBot#app_instance_arn}
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// A structure that contains configuration data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/chime_app_instance_bot#configuration ChimeAppInstanceBot#configuration}
	Configuration *ChimeAppInstanceBotConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The metadata of the AppInstanceBot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/chime_app_instance_bot#metadata ChimeAppInstanceBot#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the AppInstanceBot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/chime_app_instance_bot#name ChimeAppInstanceBot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the AppInstanceBot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/chime_app_instance_bot#tags ChimeAppInstanceBot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

