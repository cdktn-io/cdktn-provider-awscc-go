// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcegroupstagsynctask

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourcegroupsTagSyncTaskConfig struct {
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
	// The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resourcegroups_tag_sync_task#group ResourcegroupsTagSyncTask#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resourcegroups_tag_sync_task#role_arn ResourcegroupsTagSyncTask#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The tag key.
	//
	// Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resourcegroups_tag_sync_task#tag_key ResourcegroupsTagSyncTask#tag_key}
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// The tag value.
	//
	// Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resourcegroups_tag_sync_task#tag_value ResourcegroupsTagSyncTask#tag_value}
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
}

