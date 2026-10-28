// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/ecsgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsGatewayConfigurationHostSharedAlbOutputReference interface {
	cdktn.ComplexObject
	AccessLogs() EcsGatewayConfigurationHostSharedAlbAccessLogsOutputReference
	AccessLogsInput() interface{}
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
	CurrentAlbCount() *float64
	// Experimental.
	Fqn() *string
	HttpDesync() *string
	SetHttpDesync(val *string)
	HttpDesyncInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAlbCount() *float64
	SetMaxAlbCount(val *float64)
	MaxAlbCountInput() *float64
	NetworkConfiguration() EcsGatewayConfigurationHostSharedAlbNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TestTrafficRules() EcsGatewayConfigurationHostSharedAlbTestTrafficRulesOutputReference
	TestTrafficRulesInput() interface{}
	WebAclArn() *string
	SetWebAclArn(val *string)
	WebAclArnInput() *string
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
	PutAccessLogs(value *EcsGatewayConfigurationHostSharedAlbAccessLogs)
	PutNetworkConfiguration(value *EcsGatewayConfigurationHostSharedAlbNetworkConfiguration)
	PutTestTrafficRules(value *EcsGatewayConfigurationHostSharedAlbTestTrafficRules)
	ResetAccessLogs()
	ResetHttpDesync()
	ResetMaxAlbCount()
	ResetNetworkConfiguration()
	ResetPort()
	ResetProtocol()
	ResetScheme()
	ResetTestTrafficRules()
	ResetWebAclArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsGatewayConfigurationHostSharedAlbOutputReference
type jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) AccessLogs() EcsGatewayConfigurationHostSharedAlbAccessLogsOutputReference {
	var returns EcsGatewayConfigurationHostSharedAlbAccessLogsOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) AccessLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) CurrentAlbCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentAlbCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) HttpDesync() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpDesync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) HttpDesyncInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpDesyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) MaxAlbCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlbCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) MaxAlbCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlbCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) NetworkConfiguration() EcsGatewayConfigurationHostSharedAlbNetworkConfigurationOutputReference {
	var returns EcsGatewayConfigurationHostSharedAlbNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) TestTrafficRules() EcsGatewayConfigurationHostSharedAlbTestTrafficRulesOutputReference {
	var returns EcsGatewayConfigurationHostSharedAlbTestTrafficRulesOutputReference
	_jsii_.Get(
		j,
		"testTrafficRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) TestTrafficRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testTrafficRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) WebAclArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) WebAclArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclArnInput",
		&returns,
	)
	return returns
}


func NewEcsGatewayConfigurationHostSharedAlbOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsGatewayConfigurationHostSharedAlbOutputReference {
	_init_.Initialize()

	if err := validateNewEcsGatewayConfigurationHostSharedAlbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsGateway.EcsGatewayConfigurationHostSharedAlbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsGatewayConfigurationHostSharedAlbOutputReference_Override(e EcsGatewayConfigurationHostSharedAlbOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsGateway.EcsGatewayConfigurationHostSharedAlbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetHttpDesync(val *string) {
	if err := j.validateSetHttpDesyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpDesync",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetMaxAlbCount(val *float64) {
	if err := j.validateSetMaxAlbCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAlbCount",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference)SetWebAclArn(val *string) {
	if err := j.validateSetWebAclArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclArn",
		val,
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) PutAccessLogs(value *EcsGatewayConfigurationHostSharedAlbAccessLogs) {
	if err := e.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) PutNetworkConfiguration(value *EcsGatewayConfigurationHostSharedAlbNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) PutTestTrafficRules(value *EcsGatewayConfigurationHostSharedAlbTestTrafficRules) {
	if err := e.validatePutTestTrafficRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTestTrafficRules",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetHttpDesync() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpDesync",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetMaxAlbCount() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxAlbCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetScheme() {
	_jsii_.InvokeVoid(
		e,
		"resetScheme",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetTestTrafficRules() {
	_jsii_.InvokeVoid(
		e,
		"resetTestTrafficRules",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ResetWebAclArn() {
	_jsii_.InvokeVoid(
		e,
		"resetWebAclArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsGatewayConfigurationHostSharedAlbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

