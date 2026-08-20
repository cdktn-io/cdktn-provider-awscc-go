// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinestorageprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineStorageProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_storage_profile#display_name DeadlineStorageProfile#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_storage_profile#farm_id DeadlineStorageProfile#farm_id}.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_storage_profile#os_family DeadlineStorageProfile#os_family}.
	OsFamily *string `field:"required" json:"osFamily" yaml:"osFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_storage_profile#file_system_locations DeadlineStorageProfile#file_system_locations}.
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
}

