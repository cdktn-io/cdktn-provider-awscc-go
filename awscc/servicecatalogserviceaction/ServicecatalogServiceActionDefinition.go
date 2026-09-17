// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogserviceaction


type ServicecatalogServiceActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_service_action#key ServicecatalogServiceAction#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_service_action#value ServicecatalogServiceAction#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

