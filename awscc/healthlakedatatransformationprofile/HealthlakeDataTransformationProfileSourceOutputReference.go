// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/healthlakedatatransformationprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HealthlakeDataTransformationProfileSourceOutputReference interface {
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
	ExistingVersionedProfileId() HealthlakeDataTransformationProfileSourceExistingVersionedProfileIdOutputReference
	ExistingVersionedProfileIdInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProfileMapping() HealthlakeDataTransformationProfileSourceProfileMappingOutputReference
	ProfileMappingInput() interface{}
	StarterProfile() HealthlakeDataTransformationProfileSourceStarterProfileOutputReference
	StarterProfileInput() interface{}
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
	PutExistingVersionedProfileId(value *HealthlakeDataTransformationProfileSourceExistingVersionedProfileId)
	PutProfileMapping(value *HealthlakeDataTransformationProfileSourceProfileMapping)
	PutStarterProfile(value *HealthlakeDataTransformationProfileSourceStarterProfile)
	ResetExistingVersionedProfileId()
	ResetProfileMapping()
	ResetStarterProfile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthlakeDataTransformationProfileSourceOutputReference
type jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ExistingVersionedProfileId() HealthlakeDataTransformationProfileSourceExistingVersionedProfileIdOutputReference {
	var returns HealthlakeDataTransformationProfileSourceExistingVersionedProfileIdOutputReference
	_jsii_.Get(
		j,
		"existingVersionedProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ExistingVersionedProfileIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"existingVersionedProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ProfileMapping() HealthlakeDataTransformationProfileSourceProfileMappingOutputReference {
	var returns HealthlakeDataTransformationProfileSourceProfileMappingOutputReference
	_jsii_.Get(
		j,
		"profileMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ProfileMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) StarterProfile() HealthlakeDataTransformationProfileSourceStarterProfileOutputReference {
	var returns HealthlakeDataTransformationProfileSourceStarterProfileOutputReference
	_jsii_.Get(
		j,
		"starterProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) StarterProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starterProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthlakeDataTransformationProfileSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) HealthlakeDataTransformationProfileSourceOutputReference {
	_init_.Initialize()

	if err := validateNewHealthlakeDataTransformationProfileSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.healthlakeDataTransformationProfile.HealthlakeDataTransformationProfileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthlakeDataTransformationProfileSourceOutputReference_Override(h HealthlakeDataTransformationProfileSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.healthlakeDataTransformationProfile.HealthlakeDataTransformationProfileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) PutExistingVersionedProfileId(value *HealthlakeDataTransformationProfileSourceExistingVersionedProfileId) {
	if err := h.validatePutExistingVersionedProfileIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExistingVersionedProfileId",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) PutProfileMapping(value *HealthlakeDataTransformationProfileSourceProfileMapping) {
	if err := h.validatePutProfileMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putProfileMapping",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) PutStarterProfile(value *HealthlakeDataTransformationProfileSourceStarterProfile) {
	if err := h.validatePutStarterProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStarterProfile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ResetExistingVersionedProfileId() {
	_jsii_.InvokeVoid(
		h,
		"resetExistingVersionedProfileId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ResetProfileMapping() {
	_jsii_.InvokeVoid(
		h,
		"resetProfileMapping",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ResetStarterProfile() {
	_jsii_.InvokeVoid(
		h,
		"resetStarterProfile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeDataTransformationProfileSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

