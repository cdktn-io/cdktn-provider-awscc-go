// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2stage


type Apigatewayv2StageAccessLogSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_stage#destination_arn Apigatewayv2Stage#destination_arn}.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_stage#format Apigatewayv2Stage#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

