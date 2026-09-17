// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimevoiceconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeVoiceConnectorConfig struct {
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
	// The name of the Voice Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_voice_connector#name ChimeVoiceConnector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Enables or disables encryption for the Voice Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_voice_connector#require_encryption ChimeVoiceConnector#require_encryption}
	RequireEncryption interface{} `field:"required" json:"requireEncryption" yaml:"requireEncryption"`
	// The AWS Region in which the Voice Connector is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_voice_connector#aws_region ChimeVoiceConnector#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The type of network for the Voice Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_voice_connector#network_type ChimeVoiceConnector#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The tags assigned to the Voice Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_voice_connector#tags ChimeVoiceConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

