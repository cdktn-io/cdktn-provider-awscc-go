// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterEncryptionInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#encryption_at_rest MskCluster#encryption_at_rest}.
	EncryptionAtRest *MskClusterEncryptionInfoEncryptionAtRest `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#encryption_in_transit MskCluster#encryption_in_transit}.
	EncryptionInTransit *MskClusterEncryptionInfoEncryptionInTransit `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

