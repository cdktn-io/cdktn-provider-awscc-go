// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParametersCapacityReservationSpecificationCapacityReservationTarget struct {
	// The ID of the Capacity Reservation in which to run the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#capacity_reservation_id BedrockagentcoreCapacityProvider#capacity_reservation_id}
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// The ARN of the Capacity Reservation resource group in which to run the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#capacity_reservation_resource_group_arn BedrockagentcoreCapacityProvider#capacity_reservation_resource_group_arn}
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

