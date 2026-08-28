// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider


type DmsDataProviderSettingsIbmDb2LuwSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_data_provider#certificate_arn DmsDataProvider#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_data_provider#database_name DmsDataProvider#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_data_provider#port DmsDataProvider#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_data_provider#server_name DmsDataProvider#server_name}.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_data_provider#ssl_mode DmsDataProvider#ssl_mode}.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

