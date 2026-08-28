// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicinginvoiceunit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InvoicingInvoiceUnitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#invoice_receiver InvoicingInvoiceUnit#invoice_receiver}.
	InvoiceReceiver *string `field:"required" json:"invoiceReceiver" yaml:"invoiceReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#name InvoicingInvoiceUnit#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#rule InvoicingInvoiceUnit#rule}.
	Rule *InvoicingInvoiceUnitRule `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#description InvoicingInvoiceUnit#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#resource_tags InvoicingInvoiceUnit#resource_tags}.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/invoicing_invoice_unit#tax_inheritance_disabled InvoicingInvoiceUnit#tax_inheritance_disabled}.
	TaxInheritanceDisabled interface{} `field:"optional" json:"taxInheritanceDisabled" yaml:"taxInheritanceDisabled"`
}

