// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectsecurityprofile


type ConnectSecurityProfileGranularAccessControlConfigurationDataTableAccessControlConfigurationPrimaryAttributeAccessControlConfigurationPrimaryAttributeValues struct {
	// Specifies the type of access granted. Currently, only "ALLOW" is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_security_profile#access_type ConnectSecurityProfile#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The name of the primary attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_security_profile#attribute_name ConnectSecurityProfile#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// An array of allowed primary values for the specified primary attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_security_profile#values ConnectSecurityProfile#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

