// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrocksession

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockSessionConfig struct {
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
	// The Amazon Resource Name (ARN) of the KMS key to use to encrypt the session data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_session#encryption_key_arn BedrockSession#encryption_key_arn}
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// A map of key-value pairs containing attributes to be persisted across the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_session#session_metadata BedrockSession#session_metadata}
	SessionMetadata *map[string]*string `field:"optional" json:"sessionMetadata" yaml:"sessionMetadata"`
	// A list of tags associated with the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_session#tags BedrockSession#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

