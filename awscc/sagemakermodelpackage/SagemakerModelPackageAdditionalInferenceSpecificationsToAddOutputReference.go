// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelpackage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference interface {
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
	Containers() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersList
	ContainersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	SupportedContentTypes() *[]*string
	SetSupportedContentTypes(val *[]*string)
	SupportedContentTypesInput() *[]*string
	SupportedRealtimeInferenceInstanceTypes() *[]*string
	SetSupportedRealtimeInferenceInstanceTypes(val *[]*string)
	SupportedRealtimeInferenceInstanceTypesInput() *[]*string
	SupportedResponseMimeTypes() *[]*string
	SetSupportedResponseMimeTypes(val *[]*string)
	SupportedResponseMimeTypesInput() *[]*string
	SupportedTransformInstanceTypes() *[]*string
	SetSupportedTransformInstanceTypes(val *[]*string)
	SupportedTransformInstanceTypesInput() *[]*string
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
	PutContainers(value interface{})
	ResetContainers()
	ResetDescription()
	ResetName()
	ResetSupportedContentTypes()
	ResetSupportedRealtimeInferenceInstanceTypes()
	ResetSupportedResponseMimeTypes()
	ResetSupportedTransformInstanceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference
type jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) Containers() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersList {
	var returns SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedContentTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedContentTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedRealtimeInferenceInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedRealtimeInferenceInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedResponseMimeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedResponseMimeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedTransformInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) SupportedTransformInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference_Override(s SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetSupportedContentTypes(val *[]*string) {
	if err := j.validateSetSupportedContentTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedContentTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetSupportedRealtimeInferenceInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedRealtimeInferenceInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetSupportedResponseMimeTypes(val *[]*string) {
	if err := j.validateSetSupportedResponseMimeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedResponseMimeTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetSupportedTransformInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedTransformInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedTransformInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) PutContainers(value interface{}) {
	if err := s.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetContainers() {
	_jsii_.InvokeVoid(
		s,
		"resetContainers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetSupportedContentTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedContentTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetSupportedRealtimeInferenceInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedRealtimeInferenceInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetSupportedResponseMimeTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedResponseMimeTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ResetSupportedTransformInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedTransformInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

