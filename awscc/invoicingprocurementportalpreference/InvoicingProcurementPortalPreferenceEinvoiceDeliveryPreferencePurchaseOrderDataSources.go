// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference


type InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferencePurchaseOrderDataSources struct {
	// The type of e-invoice document that requires purchase order data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_document_type InvoicingProcurementPortalPreference#einvoice_delivery_document_type}
	EinvoiceDeliveryDocumentType *string `field:"optional" json:"einvoiceDeliveryDocumentType" yaml:"einvoiceDeliveryDocumentType"`
	// The type of source for purchase order data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#purchase_order_data_source_type InvoicingProcurementPortalPreference#purchase_order_data_source_type}
	PurchaseOrderDataSourceType *string `field:"optional" json:"purchaseOrderDataSourceType" yaml:"purchaseOrderDataSourceType"`
}

