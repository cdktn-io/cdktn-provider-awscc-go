// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package voiceiddomain


type VoiceidDomainServerSideEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/voiceid_domain#kms_key_id VoiceidDomain#kms_key_id}.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

