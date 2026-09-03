// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fsxvolume/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference interface {
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
	DefaultRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference
	DefaultRetentionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference
	MaximumRetentionInput() interface{}
	MinimumRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference
	MinimumRetentionInput() interface{}
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
	PutDefaultRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetention)
	PutMaximumRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetention)
	PutMinimumRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetention)
	ResetDefaultRetention()
	ResetMaximumRetention()
	ResetMinimumRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference
type jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) DefaultRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference
	_jsii_.Get(
		j,
		"defaultRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) DefaultRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) MaximumRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference
	_jsii_.Get(
		j,
		"maximumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) MaximumRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) MinimumRetention() FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference
	_jsii_.Get(
		j,
		"minimumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) MinimumRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minimumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference {
	_init_.Initialize()

	if err := validateNewFsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference_Override(f FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) PutDefaultRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetention) {
	if err := f.validatePutDefaultRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDefaultRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) PutMaximumRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetention) {
	if err := f.validatePutMaximumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMaximumRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) PutMinimumRetention(value *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetention) {
	if err := f.validatePutMinimumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMinimumRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ResetDefaultRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ResetMaximumRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetMaximumRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ResetMinimumRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetMinimumRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

