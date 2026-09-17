// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageraccount


type CertificatemanagerAccountExpiryEventsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/certificatemanager_account#days_before_expiry CertificatemanagerAccount#days_before_expiry}.
	DaysBeforeExpiry *float64 `field:"optional" json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
}

