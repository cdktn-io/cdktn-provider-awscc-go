// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomainobjecttype


type CustomerprofilesDomainObjectTypeFields struct {
	// The content type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_domain_object_type#content_type CustomerprofilesDomainObjectType#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The feature type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_domain_object_type#feature_type CustomerprofilesDomainObjectType#feature_type}
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
	// The source field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_domain_object_type#source CustomerprofilesDomainObjectType#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The target field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_domain_object_type#target CustomerprofilesDomainObjectType#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

