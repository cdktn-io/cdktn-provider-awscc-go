// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteriaStatisticalThreshold struct {
	// The percentile which resolves to a threshold value by which compliance with a behavior is determined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#statistic IotSecurityProfile#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

