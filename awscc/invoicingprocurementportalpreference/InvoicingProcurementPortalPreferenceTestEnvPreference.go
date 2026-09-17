// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference


type InvoicingProcurementPortalPreferenceTestEnvPreference struct {
	// The domain identifier for the buyer in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#buyer_domain InvoicingProcurementPortalPreference#buyer_domain}
	BuyerDomain *string `field:"optional" json:"buyerDomain" yaml:"buyerDomain"`
	// The unique identifier for the buyer in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#buyer_identifier InvoicingProcurementPortalPreference#buyer_identifier}
	BuyerIdentifier *string `field:"optional" json:"buyerIdentifier" yaml:"buyerIdentifier"`
	// The endpoint URL for e-invoice delivery in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#procurement_portal_instance_endpoint InvoicingProcurementPortalPreference#procurement_portal_instance_endpoint}
	ProcurementPortalInstanceEndpoint *string `field:"optional" json:"procurementPortalInstanceEndpoint" yaml:"procurementPortalInstanceEndpoint"`
	// The shared secret for secure communication in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#procurement_portal_shared_secret InvoicingProcurementPortalPreference#procurement_portal_shared_secret}
	ProcurementPortalSharedSecret *string `field:"optional" json:"procurementPortalSharedSecret" yaml:"procurementPortalSharedSecret"`
	// The domain identifier for the supplier in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#supplier_domain InvoicingProcurementPortalPreference#supplier_domain}
	SupplierDomain *string `field:"optional" json:"supplierDomain" yaml:"supplierDomain"`
	// The unique identifier for the supplier in the test environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/invoicing_procurement_portal_preference#supplier_identifier InvoicingProcurementPortalPreference#supplier_identifier}
	SupplierIdentifier *string `field:"optional" json:"supplierIdentifier" yaml:"supplierIdentifier"`
}

