// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupAvailabilityZoneImpairmentPolicy struct {
	// Specifies the health check behavior for the impaired Availability Zone in an active zonal shift.
	//
	// If you select ``Replace unhealthy``, instances that appear unhealthy will be replaced in all Availability Zones. If you select ``Ignore unhealthy``, instances will not be replaced in the Availability Zone with the active zonal shift. For more information, see [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html) in the *Amazon EC2 Auto Scaling User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#impaired_zone_health_check_behavior AutoscalingAutoScalingGroup#impaired_zone_health_check_behavior}
	ImpairedZoneHealthCheckBehavior *string `field:"optional" json:"impairedZoneHealthCheckBehavior" yaml:"impairedZoneHealthCheckBehavior"`
	// If ``true``, enable zonal shift for your Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#zonal_shift_enabled AutoscalingAutoScalingGroup#zonal_shift_enabled}
	ZonalShiftEnabled interface{} `field:"optional" json:"zonalShiftEnabled" yaml:"zonalShiftEnabled"`
}

