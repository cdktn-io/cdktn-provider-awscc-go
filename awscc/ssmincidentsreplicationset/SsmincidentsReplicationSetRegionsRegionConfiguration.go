// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsreplicationset


type SsmincidentsReplicationSetRegionsRegionConfiguration struct {
	// The AWS Key Management Service key ID or Key Alias to use to encrypt your replication set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssmincidents_replication_set#sse_kms_key_id SsmincidentsReplicationSet#sse_kms_key_id}
	SseKmsKeyId *string `field:"optional" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
}

