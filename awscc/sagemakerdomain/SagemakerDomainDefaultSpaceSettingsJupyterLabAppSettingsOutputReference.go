// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerdomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference interface {
	cdktn.ComplexObject
	AppLifecycleManagement() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference
	AppLifecycleManagementInput() interface{}
	BuiltInLifecycleConfigArn() *string
	SetBuiltInLifecycleConfigArn(val *string)
	BuiltInLifecycleConfigArnInput() *string
	CodeRepositories() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList
	CodeRepositoriesInput() interface{}
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
	CustomImages() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList
	CustomImagesInput() interface{}
	DefaultResourceSpec() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference
	DefaultResourceSpecInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleConfigArns() *[]*string
	SetLifecycleConfigArns(val *[]*string)
	LifecycleConfigArnsInput() *[]*string
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
	PutAppLifecycleManagement(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement)
	PutCodeRepositories(value interface{})
	PutCustomImages(value interface{})
	PutDefaultResourceSpec(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec)
	ResetAppLifecycleManagement()
	ResetBuiltInLifecycleConfigArn()
	ResetCodeRepositories()
	ResetCustomImages()
	ResetDefaultResourceSpec()
	ResetLifecycleConfigArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) AppLifecycleManagement() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) AppLifecycleManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) CodeRepositories() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList
	_jsii_.Get(
		j,
		"codeRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) CodeRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) CustomImages() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList
	_jsii_.Get(
		j,
		"customImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) CustomImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) DefaultResourceSpec() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) DefaultResourceSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference_Override(s SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) PutAppLifecycleManagement(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement) {
	if err := s.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) PutCodeRepositories(value interface{}) {
	if err := s.validatePutCodeRepositoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCodeRepositories",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) PutCustomImages(value interface{}) {
	if err := s.validatePutCustomImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomImages",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) PutDefaultResourceSpec(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec) {
	if err := s.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		s,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		s,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetCodeRepositories() {
	_jsii_.InvokeVoid(
		s,
		"resetCodeRepositories",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetCustomImages() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomImages",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

