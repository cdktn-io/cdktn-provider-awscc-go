// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetFormatOptionsExcel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#header_row DatabrewDataset#header_row}.
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#sheet_indexes DatabrewDataset#sheet_indexes}.
	SheetIndexes *[]*float64 `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#sheet_names DatabrewDataset#sheet_names}.
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
}

