// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2ListenerConfig struct {
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
	// The actions for the default rule.
	//
	// You cannot define a condition for a default rule.
	//  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#default_actions Elasticloadbalancingv2Listener#default_actions}
	DefaultActions interface{} `field:"required" json:"defaultActions" yaml:"defaultActions"`
	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#load_balancer_arn Elasticloadbalancingv2Listener#load_balancer_arn}
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#alpn_policy Elasticloadbalancingv2Listener#alpn_policy}
	AlpnPolicy *[]*string `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// The default SSL server certificate for a secure listener.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//  For an HTTPS listener, update requires some interruptions. For a TLS listener, update requires no interruption.
	//  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#certificates Elasticloadbalancingv2Listener#certificates}
	Certificates interface{} `field:"optional" json:"certificates" yaml:"certificates"`
	// The listener attributes. Attributes that you do not modify retain their current values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#listener_attributes Elasticloadbalancingv2Listener#listener_attributes}
	ListenerAttributes interface{} `field:"optional" json:"listenerAttributes" yaml:"listenerAttributes"`
	// The mutual authentication configuration information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#mutual_authentication Elasticloadbalancingv2Listener#mutual_authentication}
	MutualAuthentication *Elasticloadbalancingv2ListenerMutualAuthentication `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// The port on which the load balancer is listening. You can't specify a port for a Gateway Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#port Elasticloadbalancingv2Listener#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, TCP_UDP, QUIC, and TCP_QUIC. You can’t specify the UDP, TCP_UDP, QUIC, or TCP_QUIC protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#protocol Elasticloadbalancingv2Listener#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/describe-ssl-policies.html) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html) in the *Network Load Balancers Guide*.
	//  [HTTPS listeners] Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic. To decrease the possibility of an interruption if your load balancer is handling a high volume of traffic, create an additional load balancer or request an LCU reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#ssl_policy Elasticloadbalancingv2Listener#ssl_policy}
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#tags Elasticloadbalancingv2Listener#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

