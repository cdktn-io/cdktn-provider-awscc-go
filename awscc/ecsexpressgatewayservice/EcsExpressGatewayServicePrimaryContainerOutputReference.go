// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsexpressgatewayservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsExpressGatewayServicePrimaryContainerOutputReference interface {
	cdktn.ComplexObject
	AwsLogsConfiguration() EcsExpressGatewayServicePrimaryContainerAwsLogsConfigurationOutputReference
	AwsLogsConfigurationInput() interface{}
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	ContainerPort() *float64
	SetContainerPort(val *float64)
	ContainerPortInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() EcsExpressGatewayServicePrimaryContainerEnvironmentList
	EnvironmentInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RepositoryCredentials() EcsExpressGatewayServicePrimaryContainerRepositoryCredentialsOutputReference
	RepositoryCredentialsInput() interface{}
	Secrets() EcsExpressGatewayServicePrimaryContainerSecretsList
	SecretsInput() interface{}
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
	PutAwsLogsConfiguration(value *EcsExpressGatewayServicePrimaryContainerAwsLogsConfiguration)
	PutEnvironment(value interface{})
	PutRepositoryCredentials(value *EcsExpressGatewayServicePrimaryContainerRepositoryCredentials)
	PutSecrets(value interface{})
	ResetAwsLogsConfiguration()
	ResetCommand()
	ResetContainerPort()
	ResetEnvironment()
	ResetImage()
	ResetRepositoryCredentials()
	ResetSecrets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsExpressGatewayServicePrimaryContainerOutputReference
type jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) AwsLogsConfiguration() EcsExpressGatewayServicePrimaryContainerAwsLogsConfigurationOutputReference {
	var returns EcsExpressGatewayServicePrimaryContainerAwsLogsConfigurationOutputReference
	_jsii_.Get(
		j,
		"awsLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) AwsLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ContainerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Environment() EcsExpressGatewayServicePrimaryContainerEnvironmentList {
	var returns EcsExpressGatewayServicePrimaryContainerEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) RepositoryCredentials() EcsExpressGatewayServicePrimaryContainerRepositoryCredentialsOutputReference {
	var returns EcsExpressGatewayServicePrimaryContainerRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Secrets() EcsExpressGatewayServicePrimaryContainerSecretsList {
	var returns EcsExpressGatewayServicePrimaryContainerSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsExpressGatewayServicePrimaryContainerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsExpressGatewayServicePrimaryContainerOutputReference {
	_init_.Initialize()

	if err := validateNewEcsExpressGatewayServicePrimaryContainerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsExpressGatewayService.EcsExpressGatewayServicePrimaryContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsExpressGatewayServicePrimaryContainerOutputReference_Override(e EcsExpressGatewayServicePrimaryContainerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsExpressGatewayService.EcsExpressGatewayServicePrimaryContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetContainerPort(val *float64) {
	if err := j.validateSetContainerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerPort",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) PutAwsLogsConfiguration(value *EcsExpressGatewayServicePrimaryContainerAwsLogsConfiguration) {
	if err := e.validatePutAwsLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAwsLogsConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) PutEnvironment(value interface{}) {
	if err := e.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) PutRepositoryCredentials(value *EcsExpressGatewayServicePrimaryContainerRepositoryCredentials) {
	if err := e.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) PutSecrets(value interface{}) {
	if err := e.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecrets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetAwsLogsConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsLogsConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetContainerPort() {
	_jsii_.InvokeVoid(
		e,
		"resetContainerPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		e,
		"resetImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		e,
		"resetSecrets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServicePrimaryContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

