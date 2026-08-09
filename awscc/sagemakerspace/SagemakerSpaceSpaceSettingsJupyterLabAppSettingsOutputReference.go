// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference interface {
	cdktn.ComplexObject
	AppLifecycleManagement() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference
	AppLifecycleManagementInput() interface{}
	CodeRepositories() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList
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
	DefaultResourceSpec() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference
	DefaultResourceSpecInput() interface{}
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
	PutAppLifecycleManagement(value *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement)
	PutCodeRepositories(value interface{})
	PutDefaultResourceSpec(value *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec)
	ResetAppLifecycleManagement()
	ResetCodeRepositories()
	ResetDefaultResourceSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference
type jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) AppLifecycleManagement() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference {
	var returns SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) AppLifecycleManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) CodeRepositories() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList {
	var returns SagemakerSpaceSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList
	_jsii_.Get(
		j,
		"codeRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) CodeRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) DefaultResourceSpec() SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference {
	var returns SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) DefaultResourceSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference_Override(s SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) PutAppLifecycleManagement(value *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement) {
	if err := s.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) PutCodeRepositories(value interface{}) {
	if err := s.validatePutCodeRepositoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCodeRepositories",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) PutDefaultResourceSpec(value *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec) {
	if err := s.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		s,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ResetCodeRepositories() {
	_jsii_.InvokeVoid(
		s,
		"resetCodeRepositories",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsJupyterLabAppSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

