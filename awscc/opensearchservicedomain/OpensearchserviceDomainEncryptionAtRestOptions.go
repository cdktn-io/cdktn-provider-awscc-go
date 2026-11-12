// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainEncryptionAtRestOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The type of encryption at rest applied to the domain's data.
	//
	// Valid values are DISK and NATIVE. DISK is the default and uses volume-level encryption. NATIVE uses engine-native, index-level encryption and requires encryption at rest to be enabled and OpenSearch version 3.3 or later. After the mode is set to NATIVE, it can't be changed back to DISK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/opensearchservice_domain#encryption_mode OpensearchserviceDomain#encryption_mode}
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/opensearchservice_domain#kms_key_id OpensearchserviceDomain#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

