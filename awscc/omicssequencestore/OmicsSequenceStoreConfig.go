// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicssequencestore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsSequenceStoreConfig struct {
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
	// A name for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#name OmicsSequenceStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Location of the access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#access_log_location OmicsSequenceStore#access_log_location}
	AccessLogLocation *string `field:"optional" json:"accessLogLocation" yaml:"accessLogLocation"`
	// A description for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#description OmicsSequenceStore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#e_tag_algorithm_family OmicsSequenceStore#e_tag_algorithm_family}.
	ETagAlgorithmFamily *string `field:"optional" json:"eTagAlgorithmFamily" yaml:"eTagAlgorithmFamily"`
	// An S3 location that is used to store files that have failed a direct upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#fallback_location OmicsSequenceStore#fallback_location}
	FallbackLocation *string `field:"optional" json:"fallbackLocation" yaml:"fallbackLocation"`
	// The tags keys to propagate to the S3 objects associated with read sets in the sequence store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#propagated_set_level_tags OmicsSequenceStore#propagated_set_level_tags}
	PropagatedSetLevelTags *[]*string `field:"optional" json:"propagatedSetLevelTags" yaml:"propagatedSetLevelTags"`
	// The resource policy that controls S3 access on the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#s3_access_policy OmicsSequenceStore#s3_access_policy}
	S3AccessPolicy *string `field:"optional" json:"s3AccessPolicy" yaml:"s3AccessPolicy"`
	// Server-side encryption (SSE) settings for a store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#sse_config OmicsSequenceStore#sse_config}
	SseConfig *OmicsSequenceStoreSseConfig `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_sequence_store#tags OmicsSequenceStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

