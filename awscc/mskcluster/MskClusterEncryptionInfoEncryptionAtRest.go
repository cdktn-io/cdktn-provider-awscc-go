// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterEncryptionInfoEncryptionAtRest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_cluster#data_volume_kms_key_id MskCluster#data_volume_kms_key_id}.
	DataVolumeKmsKeyId *string `field:"optional" json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

