// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsRedshiftProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#credentials DatazoneConnection#credentials}.
	Credentials *DatazoneConnectionPropsRedshiftPropertiesCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#database_name DatazoneConnection#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#host DatazoneConnection#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Redshift Lineage Sync Configuration Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#lineage_sync DatazoneConnection#lineage_sync}
	LineageSync *DatazoneConnectionPropsRedshiftPropertiesLineageSync `field:"optional" json:"lineageSync" yaml:"lineageSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#port DatazoneConnection#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#storage DatazoneConnection#storage}.
	Storage *DatazoneConnectionPropsRedshiftPropertiesStorage `field:"optional" json:"storage" yaml:"storage"`
}

