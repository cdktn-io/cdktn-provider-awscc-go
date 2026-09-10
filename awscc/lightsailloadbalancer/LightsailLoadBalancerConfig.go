// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailloadbalancer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailLoadBalancerConfig struct {
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
	// The instance port where you're creating your load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#instance_port LightsailLoadBalancer#instance_port}
	InstancePort *float64 `field:"required" json:"instancePort" yaml:"instancePort"`
	// The name of your load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#load_balancer_name LightsailLoadBalancer#load_balancer_name}
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The names of the instances attached to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#attached_instances LightsailLoadBalancer#attached_instances}
	AttachedInstances *[]*string `field:"optional" json:"attachedInstances" yaml:"attachedInstances"`
	// The path you provided to perform the load balancer health check.
	//
	// If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#health_check_path LightsailLoadBalancer#health_check_path}
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The IP address type for the load balancer.
	//
	// The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#ip_address_type LightsailLoadBalancer#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Configuration option to enable session stickiness.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#session_stickiness_enabled LightsailLoadBalancer#session_stickiness_enabled}
	SessionStickinessEnabled interface{} `field:"optional" json:"sessionStickinessEnabled" yaml:"sessionStickinessEnabled"`
	// Configuration option to adjust session stickiness cookie duration parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#session_stickiness_lb_cookie_duration_seconds LightsailLoadBalancer#session_stickiness_lb_cookie_duration_seconds}
	SessionStickinessLbCookieDurationSeconds *string `field:"optional" json:"sessionStickinessLbCookieDurationSeconds" yaml:"sessionStickinessLbCookieDurationSeconds"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#tags LightsailLoadBalancer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the TLS policy to apply to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_load_balancer#tls_policy_name LightsailLoadBalancer#tls_policy_name}
	TlsPolicyName *string `field:"optional" json:"tlsPolicyName" yaml:"tlsPolicyName"`
}

