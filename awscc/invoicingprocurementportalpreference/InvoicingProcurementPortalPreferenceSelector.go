// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference


type InvoicingProcurementPortalPreferenceSelector struct {
	// The Amazon Resource Name (ARN) of invoice unit identifiers to which this preference applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#invoice_unit_arns InvoicingProcurementPortalPreference#invoice_unit_arns}
	InvoiceUnitArns *[]*string `field:"optional" json:"invoiceUnitArns" yaml:"invoiceUnitArns"`
}

