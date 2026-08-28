// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastoreTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_fhir_datastore#key HealthlakeFhirDatastore#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_fhir_datastore#value HealthlakeFhirDatastore#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

