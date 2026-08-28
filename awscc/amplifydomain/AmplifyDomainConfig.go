// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifydomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AmplifyDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#app_id AmplifyDomain#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#domain_name AmplifyDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#sub_domain_settings AmplifyDomain#sub_domain_settings}.
	SubDomainSettings interface{} `field:"required" json:"subDomainSettings" yaml:"subDomainSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#auto_sub_domain_creation_patterns AmplifyDomain#auto_sub_domain_creation_patterns}.
	AutoSubDomainCreationPatterns *[]*string `field:"optional" json:"autoSubDomainCreationPatterns" yaml:"autoSubDomainCreationPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#auto_sub_domain_iam_role AmplifyDomain#auto_sub_domain_iam_role}.
	AutoSubDomainIamRole *string `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#certificate_settings AmplifyDomain#certificate_settings}.
	CertificateSettings *AmplifyDomainCertificateSettings `field:"optional" json:"certificateSettings" yaml:"certificateSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_domain#enable_auto_sub_domain AmplifyDomain#enable_auto_sub_domain}.
	EnableAutoSubDomain interface{} `field:"optional" json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
}

