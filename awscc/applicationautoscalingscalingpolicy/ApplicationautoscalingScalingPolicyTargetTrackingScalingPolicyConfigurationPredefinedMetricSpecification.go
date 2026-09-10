// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	// The metric type. The ``ALBRequestCountPerTarget`` metric type applies only to Spot fleet requests and ECS services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scaling_policy#predefined_metric_type ApplicationautoscalingScalingPolicy#predefined_metric_type}
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is ``ALBRequestCountPerTarget`` and there is a target group attached to the Spot Fleet or ECS service.
	//  You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//  ``app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff``.
	//  Where:
	//   +  app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	//   +  targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	//  To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scaling_policy#resource_label ApplicationautoscalingScalingPolicy#resource_label}
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

