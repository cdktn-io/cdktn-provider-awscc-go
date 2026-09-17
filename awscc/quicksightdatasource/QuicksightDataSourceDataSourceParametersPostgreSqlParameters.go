// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceDataSourceParametersPostgreSqlParameters struct {
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// <p>Host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// <p>Port.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

