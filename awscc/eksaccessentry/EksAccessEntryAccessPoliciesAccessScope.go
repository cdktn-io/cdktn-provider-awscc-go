// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksaccessentry


type EksAccessEntryAccessPoliciesAccessScope struct {
	// The namespaces to associate with the access scope. Only specify if Type is set to 'namespace'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_access_entry#namespaces EksAccessEntry#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// The type of the access scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_access_entry#type EksAccessEntry#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

