// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package panoramapackage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PanoramaPackageConfig struct {
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
	// A name for the package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/panorama_package#package_name PanoramaPackage#package_name}
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// A storage location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/panorama_package#storage_location PanoramaPackage#storage_location}
	StorageLocation *PanoramaPackageStorageLocation `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Tags for the package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/panorama_package#tags PanoramaPackage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

