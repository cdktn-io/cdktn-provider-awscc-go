// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastorePreloadDataConfig struct {
	// The type of preloaded data. Only Synthea preloaded data is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/healthlake_fhir_datastore#preload_data_type HealthlakeFhirDatastore#preload_data_type}
	PreloadDataType *string `field:"optional" json:"preloadDataType" yaml:"preloadDataType"`
}

