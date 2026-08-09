// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabriclink/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference interface {
	cdktn.ComplexObject
	Action() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference
	ActionInput() interface{}
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
	FilterConfiguration() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfigurationList
	FilterConfigurationInput() interface{}
	FilterType() *string
	SetFilterType(val *string)
	FilterTypeInput() *string
	// Experimental.
	Fqn() *string
	HoldbackPercentage() *float64
	SetHoldbackPercentage(val *float64)
	HoldbackPercentageInput() *float64
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
	PutAction(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeAction)
	PutFilterConfiguration(value interface{})
	ResetAction()
	ResetFilterConfiguration()
	ResetFilterType()
	ResetHoldbackPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference
type jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) Action() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference {
	var returns RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) FilterConfiguration() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfigurationList {
	var returns RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfigurationList
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) FilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) FilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) HoldbackPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdbackPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) HoldbackPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdbackPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference_Override(r RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetFilterType(val *string) {
	if err := j.validateSetFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterType",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetHoldbackPercentage(val *float64) {
	if err := j.validateSetHoldbackPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"holdbackPercentage",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) PutAction(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeAction) {
	if err := r.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAction",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) PutFilterConfiguration(value interface{}) {
	if err := r.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		r,
		"resetAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ResetFilterType() {
	_jsii_.InvokeVoid(
		r,
		"resetFilterType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ResetHoldbackPercentage() {
	_jsii_.InvokeVoid(
		r,
		"resetHoldbackPercentage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

