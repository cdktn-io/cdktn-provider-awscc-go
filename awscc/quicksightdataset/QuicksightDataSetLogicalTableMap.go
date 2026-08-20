// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMap struct {
	// <p>A display name for the logical table.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#alias QuicksightDataSet#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// <p>Transform operations that act on this logical table.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#data_transforms QuicksightDataSet#data_transforms}
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
	// <p>Information about the source of a logical table.
	//
	// This is a variant type structure. For
	//             this structure to be valid, only one of the attributes can be non-null.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#source QuicksightDataSet#source}
	Source *QuicksightDataSetLogicalTableMapSource `field:"optional" json:"source" yaml:"source"`
}

