// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsaccessscope


type Ec2NetworkInsightsAccessScopeExcludePaths struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_insights_access_scope#destination Ec2NetworkInsightsAccessScope#destination}.
	Destination *Ec2NetworkInsightsAccessScopeExcludePathsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_insights_access_scope#source Ec2NetworkInsightsAccessScope#source}.
	Source *Ec2NetworkInsightsAccessScopeExcludePathsSource `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_insights_access_scope#through_resources Ec2NetworkInsightsAccessScope#through_resources}.
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

