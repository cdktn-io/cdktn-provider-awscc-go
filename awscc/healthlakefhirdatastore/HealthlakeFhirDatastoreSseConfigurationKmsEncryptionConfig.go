// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig struct {
	// The type of customer-managed-key (CMK) used for encryption.
	//
	// The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_fhir_datastore#cmk_type HealthlakeFhirDatastore#cmk_type}
	CmkType *string `field:"optional" json:"cmkType" yaml:"cmkType"`
	// The KMS encryption key id/alias used to encrypt the Data Store contents at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_fhir_datastore#kms_key_id HealthlakeFhirDatastore#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

