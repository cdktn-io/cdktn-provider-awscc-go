// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetColumnGroupsGeoSpatialColumnGroup struct {
	// <p>Columns in this hierarchy.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_set#country_code QuicksightDataSet#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// <p>A display name for the hierarchy.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

