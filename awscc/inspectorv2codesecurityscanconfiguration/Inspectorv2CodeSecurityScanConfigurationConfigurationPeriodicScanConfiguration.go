// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2codesecurityscanconfiguration


type Inspectorv2CodeSecurityScanConfigurationConfigurationPeriodicScanConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#frequency Inspectorv2CodeSecurityScanConfiguration#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_code_security_scan_configuration#frequency_expression Inspectorv2CodeSecurityScanConfiguration#frequency_expression}.
	FrequencyExpression *string `field:"optional" json:"frequencyExpression" yaml:"frequencyExpression"`
}

