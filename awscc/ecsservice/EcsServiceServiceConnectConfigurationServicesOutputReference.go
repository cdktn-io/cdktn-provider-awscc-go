// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceServiceConnectConfigurationServicesOutputReference interface {
	cdktn.ComplexObject
	ClientAliases() EcsServiceServiceConnectConfigurationServicesClientAliasesList
	ClientAliasesInput() interface{}
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
	DiscoveryName() *string
	SetDiscoveryName(val *string)
	DiscoveryNameInput() *string
	// Experimental.
	Fqn() *string
	IngressPortOverride() *float64
	SetIngressPortOverride(val *float64)
	IngressPortOverrideInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timeout() EcsServiceServiceConnectConfigurationServicesTimeoutOutputReference
	TimeoutInput() interface{}
	Tls() EcsServiceServiceConnectConfigurationServicesTlsOutputReference
	TlsInput() interface{}
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
	PutClientAliases(value interface{})
	PutTimeout(value *EcsServiceServiceConnectConfigurationServicesTimeout)
	PutTls(value *EcsServiceServiceConnectConfigurationServicesTls)
	ResetClientAliases()
	ResetDiscoveryName()
	ResetIngressPortOverride()
	ResetPortName()
	ResetTimeout()
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceServiceConnectConfigurationServicesOutputReference
type jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ClientAliases() EcsServiceServiceConnectConfigurationServicesClientAliasesList {
	var returns EcsServiceServiceConnectConfigurationServicesClientAliasesList
	_jsii_.Get(
		j,
		"clientAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ClientAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) DiscoveryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) DiscoveryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) IngressPortOverride() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPortOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) IngressPortOverrideInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPortOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) Timeout() EcsServiceServiceConnectConfigurationServicesTimeoutOutputReference {
	var returns EcsServiceServiceConnectConfigurationServicesTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) TimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) Tls() EcsServiceServiceConnectConfigurationServicesTlsOutputReference {
	var returns EcsServiceServiceConnectConfigurationServicesTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


func NewEcsServiceServiceConnectConfigurationServicesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsServiceServiceConnectConfigurationServicesOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceServiceConnectConfigurationServicesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceServiceConnectConfigurationServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsServiceServiceConnectConfigurationServicesOutputReference_Override(e EcsServiceServiceConnectConfigurationServicesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsServiceServiceConnectConfigurationServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetDiscoveryName(val *string) {
	if err := j.validateSetDiscoveryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryName",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetIngressPortOverride(val *float64) {
	if err := j.validateSetIngressPortOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressPortOverride",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) PutClientAliases(value interface{}) {
	if err := e.validatePutClientAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putClientAliases",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) PutTimeout(value *EcsServiceServiceConnectConfigurationServicesTimeout) {
	if err := e.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeout",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) PutTls(value *EcsServiceServiceConnectConfigurationServicesTls) {
	if err := e.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTls",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetClientAliases() {
	_jsii_.InvokeVoid(
		e,
		"resetClientAliases",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetDiscoveryName() {
	_jsii_.InvokeVoid(
		e,
		"resetDiscoveryName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetIngressPortOverride() {
	_jsii_.InvokeVoid(
		e,
		"resetIngressPortOverride",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		e,
		"resetPortName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		e,
		"resetTls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceServiceConnectConfigurationServicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

