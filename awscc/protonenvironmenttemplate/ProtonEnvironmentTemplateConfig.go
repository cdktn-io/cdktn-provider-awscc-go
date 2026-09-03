// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package protonenvironmenttemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ProtonEnvironmentTemplateConfig struct {
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
	// <p>A description of the environment template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#description ProtonEnvironmentTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// <p>The environment template name as displayed in the developer interface.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#display_name ProtonEnvironmentTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// <p>A customer provided encryption key that Proton uses to encrypt data.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#encryption_key ProtonEnvironmentTemplate#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#name ProtonEnvironmentTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#provisioning ProtonEnvironmentTemplate#provisioning}.
	Provisioning *string `field:"optional" json:"provisioning" yaml:"provisioning"`
	// <p>An optional list of metadata items that you can associate with the Proton environment template.
	//
	// A tag is a key-value pair.</p>
	//          <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
	//         <i>Proton User Guide</i>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_environment_template#tags ProtonEnvironmentTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

