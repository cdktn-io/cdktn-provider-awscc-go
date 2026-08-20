// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorEgressConfigVpcLattice struct {
	// Port to connect to on the target VPC Lattice resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_connector#port_number TransferConnector#port_number}
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
	// ARN of the VPC Lattice resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_connector#resource_configuration_arn TransferConnector#resource_configuration_arn}
	ResourceConfigurationArn *string `field:"optional" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
}

