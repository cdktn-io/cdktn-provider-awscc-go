// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2codesecurityscanconfiguration


type Inspectorv2CodeSecurityScanConfigurationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_scan_configuration#continuous_integration_scan_configuration Inspectorv2CodeSecurityScanConfiguration#continuous_integration_scan_configuration}.
	ContinuousIntegrationScanConfiguration *Inspectorv2CodeSecurityScanConfigurationConfigurationContinuousIntegrationScanConfiguration `field:"optional" json:"continuousIntegrationScanConfiguration" yaml:"continuousIntegrationScanConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_scan_configuration#periodic_scan_configuration Inspectorv2CodeSecurityScanConfiguration#periodic_scan_configuration}.
	PeriodicScanConfiguration *Inspectorv2CodeSecurityScanConfigurationConfigurationPeriodicScanConfiguration `field:"optional" json:"periodicScanConfiguration" yaml:"periodicScanConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_code_security_scan_configuration#rule_set_categories Inspectorv2CodeSecurityScanConfiguration#rule_set_categories}.
	RuleSetCategories *[]*string `field:"optional" json:"ruleSetCategories" yaml:"ruleSetCategories"`
}

