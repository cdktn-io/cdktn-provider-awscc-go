// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorEgressConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_connector#vpc_lattice TransferConnector#vpc_lattice}.
	VpcLattice *TransferConnectorEgressConfigVpcLattice `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
}

