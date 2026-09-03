// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceNetworkConfigurationEgressConfiguration struct {
	// Network egress type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#egress_type ApprunnerService#egress_type}
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VpcConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#vpc_connector_arn ApprunnerService#vpc_connector_arn}
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

