// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthimagingdatastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HealthimagingDatastoreConfig struct {
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
	// User friendly name for Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthimaging_datastore#datastore_name HealthimagingDatastore#datastore_name}
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// ARN referencing a KMS key or KMS key alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthimaging_datastore#kms_key_arn HealthimagingDatastore#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A Map of key value pairs for Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthimaging_datastore#tags HealthimagingDatastore#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

