// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastoreProfileConfiguration struct {
	// The FHIR profiles the Data Store validates resources against by default, for example us-core-3.1.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#default_profiles HealthlakeFhirDatastore#default_profiles}
	DefaultProfiles *[]*string `field:"optional" json:"defaultProfiles" yaml:"defaultProfiles"`
}

