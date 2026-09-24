// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMetricsCollection struct {
	// The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch. The only valid value is ``1Minute``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#granularity AutoscalingAutoScalingGroup#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Identifies the metrics to enable.
	//
	// You can specify one or more of the following metrics:
	//   +   ``GroupMinSize``
	//   +   ``GroupMaxSize``
	//   +   ``GroupDesiredCapacity``
	//   +   ``GroupInServiceInstances``
	//   +   ``GroupPendingInstances``
	//   +   ``GroupStandbyInstances``
	//   +   ``GroupTerminatingInstances``
	//   +   ``GroupTotalInstances``
	//   +   ``GroupInServiceCapacity``
	//   +   ``GroupPendingCapacity``
	//   +   ``GroupStandbyCapacity``
	//   +   ``GroupTerminatingCapacity``
	//   +   ``GroupTotalCapacity``
	//   +   ``WarmPoolDesiredCapacity``
	//   +   ``WarmPoolWarmedCapacity``
	//   +   ``WarmPoolPendingCapacity``
	//   +   ``WarmPoolTerminatingCapacity``
	//   +   ``WarmPoolTotalCapacity``
	//   +   ``GroupAndWarmPoolDesiredCapacity``
	//   +   ``GroupAndWarmPoolTotalCapacity``
	//
	//  If you specify ``Granularity`` and don't specify any metrics, all metrics are enabled.
	//  For more information, see [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html) in the *Amazon EC2 Auto Scaling User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#metrics AutoscalingAutoScalingGroup#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

