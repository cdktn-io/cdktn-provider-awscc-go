// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastoreNlpConfiguration struct {
	// The NLP status for the Data Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#status HealthlakeFhirDatastore#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

