// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph


type NeptunegraphGraphImportTaskImportOptions struct {
	// Options for importing data from a Neptune database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/neptunegraph_graph#neptune NeptunegraphGraph#neptune}
	Neptune *NeptunegraphGraphImportTaskImportOptionsNeptune `field:"optional" json:"neptune" yaml:"neptune"`
}

