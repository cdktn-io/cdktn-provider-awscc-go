// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataset


type ScnDatasetPartitionSpecFields struct {
	// The name of the partition field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/scn_dataset#name ScnDataset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The transformation of the partition field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/scn_dataset#transform ScnDataset#transform}
	Transform *ScnDatasetPartitionSpecFieldsTransform `field:"optional" json:"transform" yaml:"transform"`
}

