// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightspath


type Ec2NetworkInsightsPathFilterAtDestinationDestinationPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_insights_path#from_port Ec2NetworkInsightsPath#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_insights_path#to_port Ec2NetworkInsightsPath#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

