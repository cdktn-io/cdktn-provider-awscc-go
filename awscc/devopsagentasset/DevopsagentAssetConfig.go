// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentAssetConfig struct {
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
	// The unique identifier of the parent Agent Space. The asset is created as a child of this agent space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_asset#agent_space_id DevopsagentAsset#agent_space_id}
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The type of asset.
	//
	// The Asset API treats this as an open string; call ListAssetTypes for the current authoritative set of supported types. As of launch, customer-creatable types include skill, agents_md, and attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_asset#asset_type DevopsagentAsset#asset_type}
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Inline file list. Mutually exclusive with Zip; enforced by the handler at Create/Update time. Write-only: not repopulated by Read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_asset#files DevopsagentAsset#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// Asset metadata document.
	//
	// Required and optional keys depend on AssetType. Values may be strings, numbers, booleans, or lists of any of those - validated server-side; see the public Asset API docs for the per-type metadata schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_asset#metadata DevopsagentAsset#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Base64-encoded zip bundle containing all files for the asset.
	//
	// Mutually exclusive with Files; enforced by the handler at Create/Update time. Write-only: not repopulated by Read. Server treats a zip as 'replace all files' (max 6 MiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_asset#zip DevopsagentAsset#zip}
	Zip *string `field:"optional" json:"zip" yaml:"zip"`
}

