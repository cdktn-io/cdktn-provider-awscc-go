// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionOrigin struct {
	// The name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_distribution#name LightsailDistribution#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_distribution#protocol_policy LightsailDistribution#protocol_policy}
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// The AWS Region name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_distribution#region_name LightsailDistribution#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

