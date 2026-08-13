// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3SourceUploadSettings struct {
	// <p>Whether the file has a header row, or the files each have a header row.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_set#contains_header QuicksightDataSet#contains_header}
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// <p>The delimiter between values in the file.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_set#delimiter QuicksightDataSet#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_set#format QuicksightDataSet#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// <p>A row number to start reading data from.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_set#start_from_row QuicksightDataSet#start_from_row}
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_set#text_qualifier QuicksightDataSet#text_qualifier}.
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

