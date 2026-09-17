// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset


type DevopsagentAssetFiles struct {
	// Base64-encoded binary contents of the file. Mutually exclusive with ContentText (max 6 MiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_asset#content_bytes DevopsagentAsset#content_bytes}
	ContentBytes *string `field:"optional" json:"contentBytes" yaml:"contentBytes"`
	// UTF-8 text contents of the file. Mutually exclusive with ContentBytes (max 1.5 MiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_asset#content_text DevopsagentAsset#content_text}
	ContentText *string `field:"optional" json:"contentText" yaml:"contentText"`
	// Per-file metadata document. Values may be strings, numbers, booleans, or lists of any of those (validated server-side).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_asset#metadata DevopsagentAsset#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Path of this file within the asset bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_asset#path DevopsagentAsset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

