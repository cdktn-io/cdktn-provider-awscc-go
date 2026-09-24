// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfigurationEksLabelSelectorMatchExpressions struct {
	// Label key the requirement applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#key Resiliencehubv2Service#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Operator applied to the label key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#operator Resiliencehubv2Service#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Label values the requirement compares against.
	//
	// Up to 20 values. Required for IN and NOT_IN; omit for EXISTS and DOES_NOT_EXIST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#values Resiliencehubv2Service#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

