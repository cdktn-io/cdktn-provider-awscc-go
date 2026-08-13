// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance


type LightsailInstanceNetworking struct {
	// Monthly Transfer of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#monthly_transfer LightsailInstance#monthly_transfer}
	MonthlyTransfer *LightsailInstanceNetworkingMonthlyTransfer `field:"optional" json:"monthlyTransfer" yaml:"monthlyTransfer"`
	// Ports to the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#ports LightsailInstance#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

