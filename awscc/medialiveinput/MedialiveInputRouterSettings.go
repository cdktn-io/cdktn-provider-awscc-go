// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputRouterSettings struct {
	// MediaConnect Router destinations associated with this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#destinations MedialiveInput#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Encryption configuration for the MediaConnect Router association. AUTOMATIC uses a service-managed secret; SECRETS_MANAGER requires SecretArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#encryption_type MedialiveInput#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The ARN of the Secrets Manager secret used to encrypt this input. Required when EncryptionType is SECRETS_MANAGER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#secret_arn MedialiveInput#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

