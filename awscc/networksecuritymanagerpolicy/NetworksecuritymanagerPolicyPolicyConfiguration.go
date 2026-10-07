// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritymanagerpolicy


type NetworksecuritymanagerPolicyPolicyConfiguration struct {
	// Controls automatic remediation of non-compliant resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#remediation_enabled NetworksecuritymanagerPolicy#remediation_enabled}
	RemediationEnabled interface{} `field:"optional" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Controls automatic cleanup of unused resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#resources_clean_up NetworksecuritymanagerPolicy#resources_clean_up}
	ResourcesCleanUp interface{} `field:"optional" json:"resourcesCleanUp" yaml:"resourcesCleanUp"`
	// WAF-specific policy settings. Populated only for WAF firewall type policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#waf_config NetworksecuritymanagerPolicy#waf_config}
	WafConfig *NetworksecuritymanagerPolicyPolicyConfigurationWafConfig `field:"optional" json:"wafConfig" yaml:"wafConfig"`
}

