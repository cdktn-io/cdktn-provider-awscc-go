// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridgesource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectBridgeSourceConfig struct {
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
	// The Amazon Resource Number (ARN) of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#bridge_arn MediaconnectBridgeSource#bridge_arn}
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#name MediaconnectBridgeSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#flow_source MediaconnectBridgeSource#flow_source}
	FlowSource *MediaconnectBridgeSourceFlowSource `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge. A network source originates at your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#network_source MediaconnectBridgeSource#network_source}
	NetworkSource *MediaconnectBridgeSourceNetworkSource `field:"optional" json:"networkSource" yaml:"networkSource"`
}

