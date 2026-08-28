// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#path_prefix_hierarchy AppflowFlow#path_prefix_hierarchy}.
	PathPrefixHierarchy *[]*string `field:"optional" json:"pathPrefixHierarchy" yaml:"pathPrefixHierarchy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#prefix_format AppflowFlow#prefix_format}.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#prefix_type AppflowFlow#prefix_type}.
	PrefixType *string `field:"optional" json:"prefixType" yaml:"prefixType"`
}

