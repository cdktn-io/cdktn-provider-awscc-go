// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteriaMlDetectionConfig struct {
	// The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_security_profile#confidence_level IotSecurityProfile#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

