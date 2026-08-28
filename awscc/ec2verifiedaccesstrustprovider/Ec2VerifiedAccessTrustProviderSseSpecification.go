// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccesstrustprovider


type Ec2VerifiedAccessTrustProviderSseSpecification struct {
	// Whether to encrypt the policy with the provided key or disable encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_trust_provider#customer_managed_key_enabled Ec2VerifiedAccessTrustProvider#customer_managed_key_enabled}
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// KMS Key Arn used to encrypt the group policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_trust_provider#kms_key_arn Ec2VerifiedAccessTrustProvider#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

