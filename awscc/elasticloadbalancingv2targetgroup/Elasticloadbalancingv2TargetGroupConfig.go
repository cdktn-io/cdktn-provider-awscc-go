// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2targetgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2TargetGroupConfig struct {
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
	// Indicates whether health checks are enabled.
	//
	// If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_enabled Elasticloadbalancingv2TargetGroup#health_check_enabled}
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_interval_seconds Elasticloadbalancingv2TargetGroup#health_check_interval_seconds}
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_path Elasticloadbalancingv2TargetGroup#health_check_path}
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_port Elasticloadbalancingv2TargetGroup#health_check_port}
	HealthCheckPort *string `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_protocol Elasticloadbalancingv2TargetGroup#health_check_protocol}
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#health_check_timeout_seconds Elasticloadbalancingv2TargetGroup#health_check_timeout_seconds}
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#healthy_threshold_count Elasticloadbalancingv2TargetGroup#healthy_threshold_count}
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The type of IP address used for this target group. The possible values are ipv4 and ipv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#ip_address_type Elasticloadbalancingv2TargetGroup#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#matcher Elasticloadbalancingv2TargetGroup#matcher}
	Matcher *Elasticloadbalancingv2TargetGroupMatcher `field:"optional" json:"matcher" yaml:"matcher"`
	// The name of the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#name Elasticloadbalancingv2TargetGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#port Elasticloadbalancingv2TargetGroup#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#protocol Elasticloadbalancingv2TargetGroup#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#protocol_version Elasticloadbalancingv2TargetGroup#protocol_version}
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#tags Elasticloadbalancingv2TargetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The port that the target control agent uses to communicate the available capacity of targets to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#target_control_port Elasticloadbalancingv2TargetGroup#target_control_port}
	TargetControlPort *float64 `field:"optional" json:"targetControlPort" yaml:"targetControlPort"`
	// The attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#target_group_attributes Elasticloadbalancingv2TargetGroup#target_group_attributes}
	TargetGroupAttributes interface{} `field:"optional" json:"targetGroupAttributes" yaml:"targetGroupAttributes"`
	// The targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#targets Elasticloadbalancingv2TargetGroup#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#target_type Elasticloadbalancingv2TargetGroup#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#unhealthy_threshold_count Elasticloadbalancingv2TargetGroup#unhealthy_threshold_count}
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#vpc_id Elasticloadbalancingv2TargetGroup#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

