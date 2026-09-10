// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceSslProperties struct {
	// <p>A Boolean option to control whether SSL should be disabled.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_data_source#disable_ssl QuicksightDataSource#disable_ssl}
	DisableSsl interface{} `field:"optional" json:"disableSsl" yaml:"disableSsl"`
}

