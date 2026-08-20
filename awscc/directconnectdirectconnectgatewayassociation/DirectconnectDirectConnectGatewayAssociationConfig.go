// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectdirectconnectgatewayassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectDirectConnectGatewayAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID or ARN of the virtual private gateway or transit gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association#associated_gateway_id DirectconnectDirectConnectGatewayAssociation#associated_gateway_id}
	AssociatedGatewayId *string `field:"required" json:"associatedGatewayId" yaml:"associatedGatewayId"`
	// The ID or ARN of the Direct Connect gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association#direct_connect_gateway_id DirectconnectDirectConnectGatewayAssociation#direct_connect_gateway_id}
	DirectConnectGatewayId *string `field:"required" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// The Amazon Resource Name (ARN) of the role to accept the Direct Connect Gateway association proposal. Needs directconnect:AcceptDirectConnectGatewayAssociationProposal permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association#accept_direct_connect_gateway_association_proposal_role_arn DirectconnectDirectConnectGatewayAssociation#accept_direct_connect_gateway_association_proposal_role_arn}
	AcceptDirectConnectGatewayAssociationProposalRoleArn *string `field:"optional" json:"acceptDirectConnectGatewayAssociationProposalRoleArn" yaml:"acceptDirectConnectGatewayAssociationProposalRoleArn"`
	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	//
	// This parameter is required when you create an association to a transit gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association#allowed_prefixes_to_direct_connect_gateway DirectconnectDirectConnectGatewayAssociation#allowed_prefixes_to_direct_connect_gateway}
	AllowedPrefixesToDirectConnectGateway *[]*string `field:"optional" json:"allowedPrefixesToDirectConnectGateway" yaml:"allowedPrefixesToDirectConnectGateway"`
}

