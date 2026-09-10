// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputEncryption struct {
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_flow_output#algorithm MediaconnectFlowOutput#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The type of key that is used for the encryption.
	//
	// If no keyType is provided, the service will use the default setting (static-key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_flow_output#key_type MediaconnectFlowOutput#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_flow_output#role_arn MediaconnectFlowOutput#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	//
	// This parameter is required for static key encryption and is not valid for SPEKE encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_flow_output#secret_arn MediaconnectFlowOutput#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

