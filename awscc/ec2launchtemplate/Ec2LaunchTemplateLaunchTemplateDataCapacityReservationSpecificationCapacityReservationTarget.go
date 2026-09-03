// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTarget struct {
	// The ID of the Capacity Reservation in which to run the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#capacity_reservation_id Ec2LaunchTemplate#capacity_reservation_id}
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// The ARN of the Capacity Reservation resource group in which to run the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#capacity_reservation_resource_group_arn Ec2LaunchTemplate#capacity_reservation_resource_group_arn}
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

