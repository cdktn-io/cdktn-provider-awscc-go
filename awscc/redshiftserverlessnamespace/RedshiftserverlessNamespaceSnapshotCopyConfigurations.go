// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessnamespace


type RedshiftserverlessNamespaceSnapshotCopyConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_namespace#destination_kms_key_id RedshiftserverlessNamespace#destination_kms_key_id}.
	DestinationKmsKeyId *string `field:"optional" json:"destinationKmsKeyId" yaml:"destinationKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_namespace#destination_region RedshiftserverlessNamespace#destination_region}.
	DestinationRegion *string `field:"optional" json:"destinationRegion" yaml:"destinationRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_namespace#snapshot_retention_period RedshiftserverlessNamespace#snapshot_retention_period}.
	SnapshotRetentionPeriod *float64 `field:"optional" json:"snapshotRetentionPeriod" yaml:"snapshotRetentionPeriod"`
}

