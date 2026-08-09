// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsexternalanthropicworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/awsexternalanthropicworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsexternalanthropicWorkspaceDataResidencyOutputReference interface {
	cdktn.ComplexObject
	AllowedInferenceGeos() *[]*string
	SetAllowedInferenceGeos(val *[]*string)
	AllowedInferenceGeosInput() *[]*string
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
	DefaultInferenceGeo() *string
	SetDefaultInferenceGeo(val *string)
	DefaultInferenceGeoInput() *string
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
	WorkspaceGeo() *string
	SetWorkspaceGeo(val *string)
	WorkspaceGeoInput() *string
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
	ResetAllowedInferenceGeos()
	ResetDefaultInferenceGeo()
	ResetWorkspaceGeo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsexternalanthropicWorkspaceDataResidencyOutputReference
type jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) AllowedInferenceGeos() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInferenceGeos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) AllowedInferenceGeosInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInferenceGeosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) DefaultInferenceGeo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInferenceGeo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) DefaultInferenceGeoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInferenceGeoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) WorkspaceGeo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceGeo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) WorkspaceGeoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceGeoInput",
		&returns,
	)
	return returns
}


func NewAwsexternalanthropicWorkspaceDataResidencyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsexternalanthropicWorkspaceDataResidencyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsexternalanthropicWorkspaceDataResidencyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.awsexternalanthropicWorkspace.AwsexternalanthropicWorkspaceDataResidencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAwsexternalanthropicWorkspaceDataResidencyOutputReference_Override(a AwsexternalanthropicWorkspaceDataResidencyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.awsexternalanthropicWorkspace.AwsexternalanthropicWorkspaceDataResidencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetAllowedInferenceGeos(val *[]*string) {
	if err := j.validateSetAllowedInferenceGeosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInferenceGeos",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetDefaultInferenceGeo(val *string) {
	if err := j.validateSetDefaultInferenceGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInferenceGeo",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference)SetWorkspaceGeo(val *string) {
	if err := j.validateSetWorkspaceGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceGeo",
		val,
	)
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ResetAllowedInferenceGeos() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedInferenceGeos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ResetDefaultInferenceGeo() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultInferenceGeo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ResetWorkspaceGeo() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceGeo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsexternalanthropicWorkspaceDataResidencyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

