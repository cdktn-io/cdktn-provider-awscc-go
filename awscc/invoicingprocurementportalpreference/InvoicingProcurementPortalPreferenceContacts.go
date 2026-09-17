// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference


type InvoicingProcurementPortalPreferenceContacts struct {
	// The email address of the contact person or role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#email InvoicingProcurementPortalPreference#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the contact person or role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#name InvoicingProcurementPortalPreference#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

