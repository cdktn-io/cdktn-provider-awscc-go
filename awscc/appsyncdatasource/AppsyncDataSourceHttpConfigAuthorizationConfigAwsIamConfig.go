// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceHttpConfigAuthorizationConfigAwsIamConfig struct {
	// The signing Region for AWS Identity and Access Management authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_data_source#signing_region AppsyncDataSource#signing_region}
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS Identity and Access Management authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_data_source#signing_service_name AppsyncDataSource#signing_service_name}
	SigningServiceName *string `field:"optional" json:"signingServiceName" yaml:"signingServiceName"`
}

