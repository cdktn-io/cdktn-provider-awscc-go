// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy


type ResiliencehubResiliencyPolicyPolicy struct {
	// Failure Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehub_resiliency_policy#az ResiliencehubResiliencyPolicy#az}
	Az *ResiliencehubResiliencyPolicyPolicyAz `field:"required" json:"az" yaml:"az"`
	// Failure Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehub_resiliency_policy#hardware ResiliencehubResiliencyPolicy#hardware}
	Hardware *ResiliencehubResiliencyPolicyPolicyHardware `field:"required" json:"hardware" yaml:"hardware"`
	// Failure Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehub_resiliency_policy#software ResiliencehubResiliencyPolicy#software}
	SoftwareAttribute *ResiliencehubResiliencyPolicyPolicySoftware `field:"required" json:"softwareAttribute" yaml:"softwareAttribute"`
	// Failure Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehub_resiliency_policy#region ResiliencehubResiliencyPolicy#region}
	Region *ResiliencehubResiliencyPolicyPolicyRegion `field:"optional" json:"region" yaml:"region"`
}

