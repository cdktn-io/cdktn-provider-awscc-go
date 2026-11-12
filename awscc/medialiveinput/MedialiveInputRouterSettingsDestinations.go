// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputRouterSettingsDestinations struct {
	// The Availability Zone (AZ) name this destination is created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#availability_zone_name MedialiveInput#availability_zone_name}
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// The ARN of the MediaConnect Router output currently connected to this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#router_output_arn MedialiveInput#router_output_arn}
	RouterOutputArn *string `field:"optional" json:"routerOutputArn" yaml:"routerOutputArn"`
}

