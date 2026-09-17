// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksaddon


type EksAddonNamespaceConfig struct {
	// The custom namespace for creating the add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_addon#namespace EksAddon#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

