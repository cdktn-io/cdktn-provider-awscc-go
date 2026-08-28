// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamdirectoryconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamDirectoryConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_directory_config#directory_name AppstreamDirectoryConfig#directory_name}.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_directory_config#organizational_unit_distinguished_names AppstreamDirectoryConfig#organizational_unit_distinguished_names}.
	OrganizationalUnitDistinguishedNames *[]*string `field:"required" json:"organizationalUnitDistinguishedNames" yaml:"organizationalUnitDistinguishedNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_directory_config#service_account_credentials AppstreamDirectoryConfig#service_account_credentials}.
	ServiceAccountCredentials *AppstreamDirectoryConfigServiceAccountCredentials `field:"required" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_directory_config#certificate_based_auth_properties AppstreamDirectoryConfig#certificate_based_auth_properties}.
	CertificateBasedAuthProperties *AppstreamDirectoryConfigCertificateBasedAuthProperties `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
}

