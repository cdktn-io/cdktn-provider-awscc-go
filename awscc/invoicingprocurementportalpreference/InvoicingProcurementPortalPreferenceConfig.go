// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InvoicingProcurementPortalPreferenceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The domain identifier for the buyer in the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#buyer_domain InvoicingProcurementPortalPreference#buyer_domain}
	BuyerDomain *string `field:"required" json:"buyerDomain" yaml:"buyerDomain"`
	// The unique identifier for the buyer in the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#buyer_identifier InvoicingProcurementPortalPreference#buyer_identifier}
	BuyerIdentifier *string `field:"required" json:"buyerIdentifier" yaml:"buyerIdentifier"`
	// List of contact information for portal administrators and technical contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#contacts InvoicingProcurementPortalPreference#contacts}
	Contacts interface{} `field:"required" json:"contacts" yaml:"contacts"`
	// Indicates whether e-invoice delivery is enabled for this procurement portal preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_enabled InvoicingProcurementPortalPreference#einvoice_delivery_enabled}
	EinvoiceDeliveryEnabled interface{} `field:"required" json:"einvoiceDeliveryEnabled" yaml:"einvoiceDeliveryEnabled"`
	// The name of the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#procurement_portal_name InvoicingProcurementPortalPreference#procurement_portal_name}
	ProcurementPortalName *string `field:"required" json:"procurementPortalName" yaml:"procurementPortalName"`
	// Indicates whether purchase order retrieval is enabled for this procurement portal preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#purchase_order_retrieval_enabled InvoicingProcurementPortalPreference#purchase_order_retrieval_enabled}
	PurchaseOrderRetrievalEnabled interface{} `field:"required" json:"purchaseOrderRetrievalEnabled" yaml:"purchaseOrderRetrievalEnabled"`
	// The domain identifier for the supplier in the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#supplier_domain InvoicingProcurementPortalPreference#supplier_domain}
	SupplierDomain *string `field:"required" json:"supplierDomain" yaml:"supplierDomain"`
	// The unique identifier for the supplier in the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#supplier_identifier InvoicingProcurementPortalPreference#supplier_identifier}
	SupplierIdentifier *string `field:"required" json:"supplierIdentifier" yaml:"supplierIdentifier"`
	// Specifies the preferences for e-invoice delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#einvoice_delivery_preference InvoicingProcurementPortalPreference#einvoice_delivery_preference}
	EinvoiceDeliveryPreference *InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreference `field:"optional" json:"einvoiceDeliveryPreference" yaml:"einvoiceDeliveryPreference"`
	// The endpoint URL where e-invoices are delivered to the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#procurement_portal_instance_endpoint InvoicingProcurementPortalPreference#procurement_portal_instance_endpoint}
	ProcurementPortalInstanceEndpoint *string `field:"optional" json:"procurementPortalInstanceEndpoint" yaml:"procurementPortalInstanceEndpoint"`
	// The shared secret or authentication credential used for secure communication with the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#procurement_portal_shared_secret InvoicingProcurementPortalPreference#procurement_portal_shared_secret}
	ProcurementPortalSharedSecret *string `field:"optional" json:"procurementPortalSharedSecret" yaml:"procurementPortalSharedSecret"`
	// Specifies criteria for selecting which invoices should be processed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#selector InvoicingProcurementPortalPreference#selector}
	Selector *InvoicingProcurementPortalPreferenceSelector `field:"optional" json:"selector" yaml:"selector"`
	// The tags associated with this procurement portal preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#tags InvoicingProcurementPortalPreference#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Configuration settings for the test environment of the procurement portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference#test_env_preference InvoicingProcurementPortalPreference#test_env_preference}
	TestEnvPreference *InvoicingProcurementPortalPreferenceTestEnvPreference `field:"optional" json:"testEnvPreference" yaml:"testEnvPreference"`
}

