// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecuretunnelingtunnel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsecuretunnelingTunnelConfig struct {
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
	// A short text description of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsecuretunneling_tunnel#description IotsecuretunnelingTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination configuration for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsecuretunneling_tunnel#destination_config IotsecuretunnelingTunnel#destination_config}
	DestinationConfig *IotsecuretunnelingTunnelDestinationConfig `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// A collection of tag metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsecuretunneling_tunnel#tags IotsecuretunnelingTunnel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Timeout configuration for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsecuretunneling_tunnel#timeout_config IotsecuretunnelingTunnel#timeout_config}
	TimeoutConfig *IotsecuretunnelingTunnelTimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

