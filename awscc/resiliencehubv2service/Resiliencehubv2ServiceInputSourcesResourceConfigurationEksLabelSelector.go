// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfigurationEksLabelSelector struct {
	// Label selector requirements an object must satisfy to be discovered. Up to 20 requirements, all of which must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#match_expressions Resiliencehubv2Service#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// Label key/value pairs an object must carry to be discovered. Up to 20 pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#match_labels Resiliencehubv2Service#match_labels}
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

