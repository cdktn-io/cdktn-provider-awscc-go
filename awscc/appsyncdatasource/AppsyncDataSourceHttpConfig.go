// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceHttpConfig struct {
	// The authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#authorization_config AppsyncDataSource#authorization_config}
	AuthorizationConfig *AppsyncDataSourceHttpConfigAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#endpoint AppsyncDataSource#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

