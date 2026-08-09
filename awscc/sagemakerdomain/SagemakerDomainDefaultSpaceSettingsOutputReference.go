// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerdomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerDomainDefaultSpaceSettingsOutputReference interface {
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
	CustomFileSystemConfigs() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList
	CustomFileSystemConfigsInput() interface{}
	CustomPosixUserConfig() SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference
	CustomPosixUserConfigInput() interface{}
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JupyterLabAppSettings() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference
	JupyterLabAppSettingsInput() interface{}
	JupyterServerAppSettings() SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference
	JupyterServerAppSettingsInput() interface{}
	KernelGatewayAppSettings() SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference
	KernelGatewayAppSettingsInput() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SpaceStorageSettings() SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference
	SpaceStorageSettingsInput() interface{}
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
	PutCustomFileSystemConfigs(value interface{})
	PutCustomPosixUserConfig(value *SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfig)
	PutJupyterLabAppSettings(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings)
	PutJupyterServerAppSettings(value *SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettings)
	PutKernelGatewayAppSettings(value *SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings)
	PutSpaceStorageSettings(value *SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings)
	ResetCustomFileSystemConfigs()
	ResetCustomPosixUserConfig()
	ResetExecutionRole()
	ResetJupyterLabAppSettings()
	ResetJupyterServerAppSettings()
	ResetKernelGatewayAppSettings()
	ResetSecurityGroups()
	ResetSpaceStorageSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultSpaceSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) CustomFileSystemConfigs() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList {
	var returns SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList
	_jsii_.Get(
		j,
		"customFileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) CustomFileSystemConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) CustomPosixUserConfig() SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) CustomPosixUserConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) JupyterLabAppSettings() SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) JupyterLabAppSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) JupyterServerAppSettings() SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) JupyterServerAppSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) KernelGatewayAppSettings() SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) KernelGatewayAppSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) SpaceStorageSettings() SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) SpaceStorageSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultSpaceSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultSpaceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultSpaceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultSpaceSettingsOutputReference_Override(s SagemakerDomainDefaultSpaceSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutCustomFileSystemConfigs(value interface{}) {
	if err := s.validatePutCustomFileSystemConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomFileSystemConfigs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutCustomPosixUserConfig(value *SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfig) {
	if err := s.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutJupyterLabAppSettings(value *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings) {
	if err := s.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutJupyterServerAppSettings(value *SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettings) {
	if err := s.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutKernelGatewayAppSettings(value *SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings) {
	if err := s.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) PutSpaceStorageSettings(value *SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings) {
	if err := s.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetCustomFileSystemConfigs() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomFileSystemConfigs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

