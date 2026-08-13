// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetFormatOptionsCsv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_dataset#delimiter DatabrewDataset#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_dataset#header_row DatabrewDataset#header_row}.
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
}

