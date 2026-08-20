// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfiguration struct {
	// KMS KeyId. Can be either key uuid or full key arn or key alias arn or short key alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#key_id Kinesisanalyticsv2Application#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Specifies whether application data is encrypted using service key: AWS_OWNED_KEY or customer key: CUSTOMER_MANAGED_KEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#key_type Kinesisanalyticsv2Application#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

