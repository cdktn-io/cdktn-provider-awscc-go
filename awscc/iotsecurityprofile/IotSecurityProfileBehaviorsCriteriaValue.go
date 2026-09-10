// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteriaValue struct {
	// If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#cidrs IotSecurityProfile#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#count IotSecurityProfile#count}
	Count *string `field:"optional" json:"count" yaml:"count"`
	// The numeral value of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#number IotSecurityProfile#number}
	Number *float64 `field:"optional" json:"number" yaml:"number"`
	// The numeral values of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#numbers IotSecurityProfile#numbers}
	Numbers *[]*float64 `field:"optional" json:"numbers" yaml:"numbers"`
	// If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#ports IotSecurityProfile#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// The string values of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_security_profile#strings IotSecurityProfile#strings}
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
}

