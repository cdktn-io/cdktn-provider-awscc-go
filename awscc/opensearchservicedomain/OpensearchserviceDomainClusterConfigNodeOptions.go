// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainClusterConfigNodeOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchservice_domain#node_config OpensearchserviceDomain#node_config}.
	NodeConfig *OpensearchserviceDomainClusterConfigNodeOptionsNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchservice_domain#node_type OpensearchserviceDomain#node_type}.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

