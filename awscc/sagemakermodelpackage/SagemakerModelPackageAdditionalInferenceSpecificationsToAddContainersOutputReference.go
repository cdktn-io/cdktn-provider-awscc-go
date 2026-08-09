// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelpackage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference interface {
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
	ContainerHostname() *string
	SetContainerHostname(val *string)
	ContainerHostnameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	Framework() *string
	SetFramework(val *string)
	FrameworkInput() *string
	FrameworkVersion() *string
	SetFrameworkVersion(val *string)
	FrameworkVersionInput() *string
	Image() *string
	SetImage(val *string)
	ImageDigest() *string
	SetImageDigest(val *string)
	ImageDigestInput() *string
	ImageInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelDataSource() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSourceOutputReference
	ModelDataSourceInput() interface{}
	ModelDataUrl() *string
	SetModelDataUrl(val *string)
	ModelDataUrlInput() *string
	ModelInput() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInputOutputReference
	ModelInputInput() interface{}
	NearestModelName() *string
	SetNearestModelName(val *string)
	NearestModelNameInput() *string
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
	PutModelDataSource(value *SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSource)
	PutModelInput(value *SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInput)
	ResetContainerHostname()
	ResetEnvironment()
	ResetFramework()
	ResetFrameworkVersion()
	ResetImage()
	ResetImageDigest()
	ResetModelDataSource()
	ResetModelDataUrl()
	ResetModelInput()
	ResetNearestModelName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference
type jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ContainerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ContainerHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) FrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) FrameworkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ImageDigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDigestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelDataSource() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSourceOutputReference {
	var returns SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSourceOutputReference
	_jsii_.Get(
		j,
		"modelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelDataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelInput() SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInputOutputReference {
	var returns SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInputOutputReference
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ModelInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) NearestModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearestModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) NearestModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearestModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference_Override(s SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetContainerHostname(val *string) {
	if err := j.validateSetContainerHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerHostname",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetFramework(val *string) {
	if err := j.validateSetFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetFrameworkVersion(val *string) {
	if err := j.validateSetFrameworkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkVersion",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetImageDigest(val *string) {
	if err := j.validateSetImageDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageDigest",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetModelDataUrl(val *string) {
	if err := j.validateSetModelDataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataUrl",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetNearestModelName(val *string) {
	if err := j.validateSetNearestModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nearestModelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) PutModelDataSource(value *SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSource) {
	if err := s.validatePutModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) PutModelInput(value *SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelInput) {
	if err := s.validatePutModelInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetContainerHostname() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerHostname",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetFramework() {
	_jsii_.InvokeVoid(
		s,
		"resetFramework",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetFrameworkVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetFrameworkVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		s,
		"resetImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetImageDigest() {
	_jsii_.InvokeVoid(
		s,
		"resetImageDigest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetModelDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetModelDataUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetModelInput() {
	_jsii_.InvokeVoid(
		s,
		"resetModelInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ResetNearestModelName() {
	_jsii_.InvokeVoid(
		s,
		"resetNearestModelName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

