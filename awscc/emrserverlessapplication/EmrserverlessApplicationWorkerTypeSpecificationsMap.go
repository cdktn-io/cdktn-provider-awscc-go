// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/emrserverlessapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrserverlessApplicationWorkerTypeSpecificationsMap interface {
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
	Get(key *string) EmrserverlessApplicationWorkerTypeSpecificationsOutputReference
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

// The jsii proxy struct for EmrserverlessApplicationWorkerTypeSpecificationsMap
type jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrserverlessApplicationWorkerTypeSpecificationsMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EmrserverlessApplicationWorkerTypeSpecificationsMap {
	_init_.Initialize()

	if err := validateNewEmrserverlessApplicationWorkerTypeSpecificationsMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.emrserverlessApplication.EmrserverlessApplicationWorkerTypeSpecificationsMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrserverlessApplicationWorkerTypeSpecificationsMap_Override(e EmrserverlessApplicationWorkerTypeSpecificationsMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.emrserverlessApplication.EmrserverlessApplicationWorkerTypeSpecificationsMap",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) Get(key *string) EmrserverlessApplicationWorkerTypeSpecificationsOutputReference {
	if err := e.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns EmrserverlessApplicationWorkerTypeSpecificationsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationWorkerTypeSpecificationsMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

