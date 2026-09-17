// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2codesecurityscanconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2CodeSecurityScanConfigurationConfig struct {
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
	// Code Security Scan Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#configuration Inspectorv2CodeSecurityScanConfiguration#configuration}
	Configuration *Inspectorv2CodeSecurityScanConfigurationConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Configuration Level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#level Inspectorv2CodeSecurityScanConfiguration#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Code Security Scan Configuration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#name Inspectorv2CodeSecurityScanConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Scope Settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#scope_settings Inspectorv2CodeSecurityScanConfiguration#scope_settings}
	ScopeSettings *Inspectorv2CodeSecurityScanConfigurationScopeSettings `field:"optional" json:"scopeSettings" yaml:"scopeSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#tags Inspectorv2CodeSecurityScanConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

