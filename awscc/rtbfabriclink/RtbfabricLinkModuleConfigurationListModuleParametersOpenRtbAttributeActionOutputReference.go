// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabriclink/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference interface {
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
	HeaderTag() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionHeaderTagOutputReference
	HeaderTagInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoBid() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionNoBidOutputReference
	NoBidInput() interface{}
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
	PutHeaderTag(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionHeaderTag)
	PutNoBid(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionNoBid)
	ResetHeaderTag()
	ResetNoBid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference
type jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) HeaderTag() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionHeaderTagOutputReference {
	var returns RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionHeaderTagOutputReference
	_jsii_.Get(
		j,
		"headerTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) HeaderTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) NoBid() RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionNoBidOutputReference {
	var returns RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionNoBidOutputReference
	_jsii_.Get(
		j,
		"noBid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) NoBidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noBidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference_Override(r RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) PutHeaderTag(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionHeaderTag) {
	if err := r.validatePutHeaderTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHeaderTag",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) PutNoBid(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionNoBid) {
	if err := r.validatePutNoBidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNoBid",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ResetHeaderTag() {
	_jsii_.InvokeVoid(
		r,
		"resetHeaderTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ResetNoBid() {
	_jsii_.InvokeVoid(
		r,
		"resetNoBid",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

