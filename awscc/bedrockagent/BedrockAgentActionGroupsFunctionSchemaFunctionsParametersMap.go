// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap interface {
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
	Get(key *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference
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

// The jsii proxy struct for BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap
type jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap {
	_init_.Initialize()

	if err := validateNewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap_Override(b BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Get(key *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference {
	if err := b.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

