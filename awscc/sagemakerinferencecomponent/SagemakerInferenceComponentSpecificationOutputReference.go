// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerinferencecomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerInferenceComponentSpecificationOutputReference interface {
	cdktn.ComplexObject
	BaseInferenceComponentName() *string
	SetBaseInferenceComponentName(val *string)
	BaseInferenceComponentNameInput() *string
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
	ComputeResourceRequirements() SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference
	ComputeResourceRequirementsInput() interface{}
	Container() SagemakerInferenceComponentSpecificationContainerOutputReference
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentDataCacheConfig() SagemakerInferenceComponentSpecificationCurrentDataCacheConfigOutputReference
	DataCacheConfig() SagemakerInferenceComponentSpecificationDataCacheConfigOutputReference
	DataCacheConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
	SchedulingConfig() SagemakerInferenceComponentSpecificationSchedulingConfigOutputReference
	SchedulingConfigInput() interface{}
	StartupParameters() SagemakerInferenceComponentSpecificationStartupParametersOutputReference
	StartupParametersInput() interface{}
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
	PutComputeResourceRequirements(value *SagemakerInferenceComponentSpecificationComputeResourceRequirements)
	PutContainer(value *SagemakerInferenceComponentSpecificationContainer)
	PutDataCacheConfig(value *SagemakerInferenceComponentSpecificationDataCacheConfig)
	PutSchedulingConfig(value *SagemakerInferenceComponentSpecificationSchedulingConfig)
	PutStartupParameters(value *SagemakerInferenceComponentSpecificationStartupParameters)
	ResetBaseInferenceComponentName()
	ResetComputeResourceRequirements()
	ResetContainer()
	ResetDataCacheConfig()
	ResetModelName()
	ResetSchedulingConfig()
	ResetStartupParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerInferenceComponentSpecificationOutputReference
type jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) BaseInferenceComponentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInferenceComponentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) BaseInferenceComponentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInferenceComponentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ComputeResourceRequirements() SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference {
	var returns SagemakerInferenceComponentSpecificationComputeResourceRequirementsOutputReference
	_jsii_.Get(
		j,
		"computeResourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ComputeResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeResourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) Container() SagemakerInferenceComponentSpecificationContainerOutputReference {
	var returns SagemakerInferenceComponentSpecificationContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) CurrentDataCacheConfig() SagemakerInferenceComponentSpecificationCurrentDataCacheConfigOutputReference {
	var returns SagemakerInferenceComponentSpecificationCurrentDataCacheConfigOutputReference
	_jsii_.Get(
		j,
		"currentDataCacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) DataCacheConfig() SagemakerInferenceComponentSpecificationDataCacheConfigOutputReference {
	var returns SagemakerInferenceComponentSpecificationDataCacheConfigOutputReference
	_jsii_.Get(
		j,
		"dataCacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) DataCacheConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) SchedulingConfig() SagemakerInferenceComponentSpecificationSchedulingConfigOutputReference {
	var returns SagemakerInferenceComponentSpecificationSchedulingConfigOutputReference
	_jsii_.Get(
		j,
		"schedulingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) SchedulingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) StartupParameters() SagemakerInferenceComponentSpecificationStartupParametersOutputReference {
	var returns SagemakerInferenceComponentSpecificationStartupParametersOutputReference
	_jsii_.Get(
		j,
		"startupParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) StartupParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startupParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerInferenceComponentSpecificationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerInferenceComponentSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerInferenceComponentSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerInferenceComponentSpecificationOutputReference_Override(s SagemakerInferenceComponentSpecificationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetBaseInferenceComponentName(val *string) {
	if err := j.validateSetBaseInferenceComponentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseInferenceComponentName",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) PutComputeResourceRequirements(value *SagemakerInferenceComponentSpecificationComputeResourceRequirements) {
	if err := s.validatePutComputeResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComputeResourceRequirements",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) PutContainer(value *SagemakerInferenceComponentSpecificationContainer) {
	if err := s.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainer",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) PutDataCacheConfig(value *SagemakerInferenceComponentSpecificationDataCacheConfig) {
	if err := s.validatePutDataCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataCacheConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) PutSchedulingConfig(value *SagemakerInferenceComponentSpecificationSchedulingConfig) {
	if err := s.validatePutSchedulingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchedulingConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) PutStartupParameters(value *SagemakerInferenceComponentSpecificationStartupParameters) {
	if err := s.validatePutStartupParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStartupParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetBaseInferenceComponentName() {
	_jsii_.InvokeVoid(
		s,
		"resetBaseInferenceComponentName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetComputeResourceRequirements() {
	_jsii_.InvokeVoid(
		s,
		"resetComputeResourceRequirements",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		s,
		"resetContainer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetDataCacheConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDataCacheConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		s,
		"resetModelName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetSchedulingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSchedulingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ResetStartupParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetStartupParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

