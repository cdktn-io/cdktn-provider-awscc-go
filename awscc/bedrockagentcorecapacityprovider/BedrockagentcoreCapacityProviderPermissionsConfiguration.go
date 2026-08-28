// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderPermissionsConfiguration struct {
	// The ARN of the IAM role that operators use to manage the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_capacity_provider#capacity_provider_operator_role_arn BedrockagentcoreCapacityProvider#capacity_provider_operator_role_arn}
	CapacityProviderOperatorRoleArn *string `field:"required" json:"capacityProviderOperatorRoleArn" yaml:"capacityProviderOperatorRoleArn"`
}

