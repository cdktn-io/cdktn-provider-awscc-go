// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptunegraphGraphConfig struct {
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
	// Memory for the Graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#provisioned_memory NeptunegraphGraph#provisioned_memory}
	ProvisionedMemory *float64 `field:"required" json:"provisionedMemory" yaml:"provisionedMemory"`
	// Value that indicates whether the Graph has deletion protection enabled.
	//
	// The graph can't be deleted when deletion protection is enabled.
	//
	// _Default_: If not specified, the default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#deletion_protection NeptunegraphGraph#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Contains a user-supplied name for the Graph.
	//
	// If you don't specify a name, we generate a unique Graph Name using a combination of Stack Name and a UUID comprising of 4 characters.
	//
	// _Important_: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#graph_name NeptunegraphGraph#graph_name}
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// The details of the import task to use to create the graph.
	//
	// When specified, the graph is created using CreateGraphUsingImportTask and data is imported from the supplied source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#import_task NeptunegraphGraph#import_task}
	ImportTask *NeptunegraphGraphImportTask `field:"optional" json:"importTask" yaml:"importTask"`
	// The ARN of the KMS key used to encrypt data in the Neptune Analytics graph.
	//
	// If not specified, the graph is encrypted with an AWS managed key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#kms_key_identifier NeptunegraphGraph#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.
	//
	// When the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.
	//
	// When the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.
	//
	// _Default_: If not specified, the default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#public_connectivity NeptunegraphGraph#public_connectivity}
	PublicConnectivity interface{} `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.
	//
	// Replica Count should always be less than or equal to 2.
	//
	// _Default_: If not specified, the default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#replica_count NeptunegraphGraph#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// The tags associated with this graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#tags NeptunegraphGraph#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Vector Search Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#vector_search_configuration NeptunegraphGraph#vector_search_configuration}
	VectorSearchConfiguration *NeptunegraphGraphVectorSearchConfiguration `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

