// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServicesTls struct {
	// The signer certificate authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#issuer_certificate_authority EcsService#issuer_certificate_authority}
	IssuerCertificateAuthority *EcsServiceServiceConnectConfigurationServicesTlsIssuerCertificateAuthority `field:"optional" json:"issuerCertificateAuthority" yaml:"issuerCertificateAuthority"`
	// The AWS Key Management Service key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#kms_key EcsService#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Amazon Resource Name (ARN) of the IAM role that's associated with the Service Connect TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#role_arn EcsService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

