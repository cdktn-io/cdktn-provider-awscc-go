// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap
type jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreHarness.BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap_Override(b BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreHarness.BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) Get(key *string) BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference {
	if err := b.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

