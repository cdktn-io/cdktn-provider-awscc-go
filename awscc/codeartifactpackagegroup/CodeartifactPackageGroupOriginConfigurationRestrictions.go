// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeartifactpackagegroup


type CodeartifactPackageGroupOriginConfigurationRestrictions struct {
	// The external upstream restriction determines if new package versions can be ingested or retained from external connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codeartifact_package_group#external_upstream CodeartifactPackageGroup#external_upstream}
	ExternalUpstream *CodeartifactPackageGroupOriginConfigurationRestrictionsExternalUpstream `field:"optional" json:"externalUpstream" yaml:"externalUpstream"`
	// The internal upstream restriction determines if new package versions can be ingested or retained from upstream repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codeartifact_package_group#internal_upstream CodeartifactPackageGroup#internal_upstream}
	InternalUpstream *CodeartifactPackageGroupOriginConfigurationRestrictionsInternalUpstream `field:"optional" json:"internalUpstream" yaml:"internalUpstream"`
	// The publish restriction determines if new package versions can be published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codeartifact_package_group#publish CodeartifactPackageGroup#publish}
	Publish *CodeartifactPackageGroupOriginConfigurationRestrictionsPublish `field:"optional" json:"publish" yaml:"publish"`
}

