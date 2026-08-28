// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebbrowsersettings


type WorkspaceswebBrowserSettingsWebContentFilteringPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspacesweb_browser_settings#allowed_urls WorkspaceswebBrowserSettings#allowed_urls}.
	AllowedUrls *[]*string `field:"optional" json:"allowedUrls" yaml:"allowedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspacesweb_browser_settings#blocked_categories WorkspaceswebBrowserSettings#blocked_categories}.
	BlockedCategories *[]*string `field:"optional" json:"blockedCategories" yaml:"blockedCategories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspacesweb_browser_settings#blocked_urls WorkspaceswebBrowserSettings#blocked_urls}.
	BlockedUrls *[]*string `field:"optional" json:"blockedUrls" yaml:"blockedUrls"`
}

