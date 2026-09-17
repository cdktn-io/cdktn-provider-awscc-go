// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions struct {
	// Deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#block_duration_minutes Ec2LaunchTemplate#block_duration_minutes}
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// The behavior when a Spot Instance is interrupted. The default is ``terminate``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#instance_interruption_behavior Ec2LaunchTemplate#instance_interruption_behavior}
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The maximum hourly price you're willing to pay for the Spot Instances.
	//
	// We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
	//   If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#max_price Ec2LaunchTemplate#max_price}
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The Spot Instance request type.
	//
	// If you are using Spot Instances with an Auto Scaling group, use ``one-time`` requests, as the ASlong service handles requesting new Spot Instances whenever the group is below its desired capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#spot_instance_type Ec2LaunchTemplate#spot_instance_type}
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// The end date of the request, in UTC format (*YYYY-MM-DD*T*HH:MM:SS*Z).
	//
	// Supported only for persistent requests.
	//   +  For a persistent request, the request remains active until the ``ValidUntil`` date and time is reached. Otherwise, the request remains active until you cancel it.
	//   +  For a one-time request, ``ValidUntil`` is not supported. The request remains active until all instances launch or you cancel the request.
	//
	//  Default: 7 days from the current date
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#valid_until Ec2LaunchTemplate#valid_until}
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

