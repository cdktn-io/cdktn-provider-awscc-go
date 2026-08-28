// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataset


type ScnDatasetPartitionSpec struct {
	// The partition fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#fields ScnDataset#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

