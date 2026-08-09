// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apsworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelSet() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLabelSetList
	LabelSetInput() interface{}
	Limits() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimitsOutputReference
	LimitsInput() interface{}
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
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutLabelSet(value interface{})
	PutLimits(value *ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimits)
	ResetLabelSet()
	ResetLimits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference
type jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) LabelSet() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLabelSetList {
	var returns ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLabelSetList
	_jsii_.Get(
		j,
		"labelSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) LabelSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) Limits() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimitsOutputReference {
	var returns ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) LimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference {
	_init_.Initialize()

	if err := validateNewApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apsWorkspace.ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference_Override(a ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apsWorkspace.ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) PutLabelSet(value interface{}) {
	if err := a.validatePutLabelSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) PutLimits(value *ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimits) {
	if err := a.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLimits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ResetLabelSet() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		a,
		"resetLimits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

