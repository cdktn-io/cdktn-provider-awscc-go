// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2codesecurityintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2CodeSecurityIntegrationConfig struct {
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
	// Create Integration Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_integration#create_integration_details Inspectorv2CodeSecurityIntegration#create_integration_details}
	CreateIntegrationDetails *Inspectorv2CodeSecurityIntegrationCreateIntegrationDetails `field:"optional" json:"createIntegrationDetails" yaml:"createIntegrationDetails"`
	// Code Security Integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_integration#name Inspectorv2CodeSecurityIntegration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_integration#tags Inspectorv2CodeSecurityIntegration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Integration Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_integration#type Inspectorv2CodeSecurityIntegration#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Update Integration Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_integration#update_integration_details Inspectorv2CodeSecurityIntegration#update_integration_details}
	UpdateIntegrationDetails *Inspectorv2CodeSecurityIntegrationUpdateIntegrationDetails `field:"optional" json:"updateIntegrationDetails" yaml:"updateIntegrationDetails"`
}

