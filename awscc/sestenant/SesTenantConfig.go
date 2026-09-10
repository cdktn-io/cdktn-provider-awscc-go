// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sestenant

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesTenantConfig struct {
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
	// The name of the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_tenant#tenant_name SesTenant#tenant_name}
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// The list of resources to associate with the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_tenant#resource_associations SesTenant#resource_associations}
	ResourceAssociations interface{} `field:"optional" json:"resourceAssociations" yaml:"resourceAssociations"`
	// The tags (keys and values) associated with the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_tenant#tags SesTenant#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

