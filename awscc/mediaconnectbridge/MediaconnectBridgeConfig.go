// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectBridgeConfig struct {
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
	// The name of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The placement Amazon Resource Number (ARN) of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#placement_arn MediaconnectBridge#placement_arn}
	PlacementArn *string `field:"required" json:"placementArn" yaml:"placementArn"`
	// The sources on this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#sources MediaconnectBridge#sources}
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#egress_gateway_bridge MediaconnectBridge#egress_gateway_bridge}.
	EgressGatewayBridge *MediaconnectBridgeEgressGatewayBridge `field:"optional" json:"egressGatewayBridge" yaml:"egressGatewayBridge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#ingress_gateway_bridge MediaconnectBridge#ingress_gateway_bridge}.
	IngressGatewayBridge *MediaconnectBridgeIngressGatewayBridge `field:"optional" json:"ingressGatewayBridge" yaml:"ingressGatewayBridge"`
	// The outputs on this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#outputs MediaconnectBridge#outputs}
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The settings for source failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#source_failover_config MediaconnectBridge#source_failover_config}
	SourceFailoverConfig *MediaconnectBridgeSourceFailoverConfig `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

