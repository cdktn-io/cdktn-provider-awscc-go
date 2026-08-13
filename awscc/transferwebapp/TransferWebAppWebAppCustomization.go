// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppWebAppCustomization struct {
	// Specifies a favicon to display in the browser tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_web_app#favicon_file TransferWebApp#favicon_file}
	FaviconFile *string `field:"optional" json:"faviconFile" yaml:"faviconFile"`
	// Specifies a logo to display on the web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_web_app#logo_file TransferWebApp#logo_file}
	LogoFile *string `field:"optional" json:"logoFile" yaml:"logoFile"`
	// Specifies a title to display on the web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_web_app#title TransferWebApp#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

