// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_policy#static VerifiedpermissionsPolicy#static}.
	Static *VerifiedpermissionsPolicyDefinitionStatic `field:"optional" json:"static" yaml:"static"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_policy#template_linked VerifiedpermissionsPolicy#template_linked}.
	TemplateLinked *VerifiedpermissionsPolicyDefinitionTemplateLinked `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

