// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference


type InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreference struct {
	// The method to use for testing the connection to the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#connection_testing_method InvoicingProcurementPortalPreference#connection_testing_method}
	ConnectionTestingMethod *string `field:"optional" json:"connectionTestingMethod" yaml:"connectionTestingMethod"`
	// The ISO 8601 date-time when e-invoice delivery should be activated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_activation_date InvoicingProcurementPortalPreference#einvoice_delivery_activation_date}
	EinvoiceDeliveryActivationDate *string `field:"optional" json:"einvoiceDeliveryActivationDate" yaml:"einvoiceDeliveryActivationDate"`
	// The types of attachments to include with the e-invoice delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_attachment_types InvoicingProcurementPortalPreference#einvoice_delivery_attachment_types}
	EinvoiceDeliveryAttachmentTypes *[]*string `field:"optional" json:"einvoiceDeliveryAttachmentTypes" yaml:"einvoiceDeliveryAttachmentTypes"`
	// The types of e-invoice documents to be delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_document_types InvoicingProcurementPortalPreference#einvoice_delivery_document_types}
	EinvoiceDeliveryDocumentTypes *[]*string `field:"optional" json:"einvoiceDeliveryDocumentTypes" yaml:"einvoiceDeliveryDocumentTypes"`
	// The communication protocol to use for e-invoice delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#protocol InvoicingProcurementPortalPreference#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The sources of purchase order data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#purchase_order_data_sources InvoicingProcurementPortalPreference#purchase_order_data_sources}
	PurchaseOrderDataSources interface{} `field:"optional" json:"purchaseOrderDataSources" yaml:"purchaseOrderDataSources"`
}

