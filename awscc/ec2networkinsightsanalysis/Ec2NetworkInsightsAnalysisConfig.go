// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsanalysis

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkInsightsAnalysisConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_network_insights_analysis#network_insights_path_id Ec2NetworkInsightsAnalysis#network_insights_path_id}.
	NetworkInsightsPathId *string `field:"required" json:"networkInsightsPathId" yaml:"networkInsightsPathId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_network_insights_analysis#filter_in_arns Ec2NetworkInsightsAnalysis#filter_in_arns}.
	FilterInArns *[]*string `field:"optional" json:"filterInArns" yaml:"filterInArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_network_insights_analysis#tags Ec2NetworkInsightsAnalysis#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

