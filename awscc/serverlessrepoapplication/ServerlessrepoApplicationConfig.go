// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serverlessrepoapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServerlessrepoApplicationConfig struct {
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
	// The name of the author publishing the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#author ServerlessrepoApplication#author}
	Author *string `field:"required" json:"author" yaml:"author"`
	// The description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#description ServerlessrepoApplication#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#name ServerlessrepoApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A URL with more information about the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#home_page_url ServerlessrepoApplication#home_page_url}
	HomePageUrl *string `field:"optional" json:"homePageUrl" yaml:"homePageUrl"`
	// Labels to improve discovery of apps in search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#labels ServerlessrepoApplication#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// A local text file that contains the license of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#license_body ServerlessrepoApplication#license_body}
	LicenseBody *string `field:"optional" json:"licenseBody" yaml:"licenseBody"`
	// A text readme file in Markdown language that contains a more detailed description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#readme_body ServerlessrepoApplication#readme_body}
	ReadmeBody *string `field:"optional" json:"readmeBody" yaml:"readmeBody"`
	// The semantic version of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#semantic_version ServerlessrepoApplication#semantic_version}
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// A link to a public repository for the source code of your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#source_code_url ServerlessrepoApplication#source_code_url}
	SourceCodeUrl *string `field:"optional" json:"sourceCodeUrl" yaml:"sourceCodeUrl"`
	// A valid identifier from https://spdx.org/licenses/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#spdx_license_id ServerlessrepoApplication#spdx_license_id}
	SpdxLicenseId *string `field:"optional" json:"spdxLicenseId" yaml:"spdxLicenseId"`
	// The local raw packaged AWS SAM template file of your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/serverlessrepo_application#template_body ServerlessrepoApplication#template_body}
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
}

