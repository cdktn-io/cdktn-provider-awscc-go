// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectlag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectLagConfig struct {
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
	// The bandwidth of the individual physical dedicated connections bundled by the LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#connections_bandwidth DirectconnectLag#connections_bandwidth}
	ConnectionsBandwidth *string `field:"required" json:"connectionsBandwidth" yaml:"connectionsBandwidth"`
	// The name of the LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#lag_name DirectconnectLag#lag_name}
	LagName *string `field:"required" json:"lagName" yaml:"lagName"`
	// The location for the LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#location DirectconnectLag#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#minimum_links DirectconnectLag#minimum_links}
	MinimumLinks *float64 `field:"optional" json:"minimumLinks" yaml:"minimumLinks"`
	// The name of the service provider associated with the requested LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#provider_name DirectconnectLag#provider_name}
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Indicates whether you want the LAG to support MAC Security (MACsec).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#request_mac_sec DirectconnectLag#request_mac_sec}
	RequestMacSec interface{} `field:"optional" json:"requestMacSec" yaml:"requestMacSec"`
	// The tags associated with the LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_lag#tags DirectconnectLag#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

