// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsworkflowversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/omicsworkflowversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsWorkflowVersionParameterTemplateMap interface {
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
	Get(key *string) OmicsWorkflowVersionParameterTemplateOutputReference
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

// The jsii proxy struct for OmicsWorkflowVersionParameterTemplateMap
type jsiiProxy_OmicsWorkflowVersionParameterTemplateMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOmicsWorkflowVersionParameterTemplateMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OmicsWorkflowVersionParameterTemplateMap {
	_init_.Initialize()

	if err := validateNewOmicsWorkflowVersionParameterTemplateMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OmicsWorkflowVersionParameterTemplateMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.omicsWorkflowVersion.OmicsWorkflowVersionParameterTemplateMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOmicsWorkflowVersionParameterTemplateMap_Override(o OmicsWorkflowVersionParameterTemplateMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.omicsWorkflowVersion.OmicsWorkflowVersionParameterTemplateMap",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) Get(key *string) OmicsWorkflowVersionParameterTemplateOutputReference {
	if err := o.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns OmicsWorkflowVersionParameterTemplateOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OmicsWorkflowVersionParameterTemplateMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

