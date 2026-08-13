// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurationsFilterExpressions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#expression DatazoneDataSource#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The search filter expression type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#type DatazoneDataSource#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

