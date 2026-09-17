// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydeployment


type ApigatewayDeploymentStageDescriptionAccessLogSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_deployment#destination_arn ApigatewayDeployment#destination_arn}.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_deployment#format ApigatewayDeployment#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

