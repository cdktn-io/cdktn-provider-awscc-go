// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceHttpConfigAuthorizationConfig struct {
	// The authorization type that the HTTP endpoint requires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_data_source#authorization_type AppsyncDataSource#authorization_type}
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The AWS Identity and Access Management settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_data_source#aws_iam_config AppsyncDataSource#aws_iam_config}
	AwsIamConfig *AppsyncDataSourceHttpConfigAuthorizationConfigAwsIamConfig `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

