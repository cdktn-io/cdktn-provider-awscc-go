// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreruntime/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference interface {
	cdktn.ComplexObject
	AllowedAudience() *[]*string
	SetAllowedAudience(val *[]*string)
	AllowedAudienceInput() *[]*string
	AllowedClients() *[]*string
	SetAllowedClients(val *[]*string)
	AllowedClientsInput() *[]*string
	AllowedScopes() *[]*string
	SetAllowedScopes(val *[]*string)
	AllowedScopesInput() *[]*string
	AllowedWorkloadConfiguration() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationOutputReference
	AllowedWorkloadConfigurationInput() interface{}
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
	CustomClaims() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsList
	CustomClaimsInput() interface{}
	DiscoveryUrl() *string
	SetDiscoveryUrl(val *string)
	DiscoveryUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateEndpoint() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOutputReference
	PrivateEndpointInput() interface{}
	PrivateEndpointOverrides() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList
	PrivateEndpointOverridesInput() interface{}
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
	PutAllowedWorkloadConfiguration(value *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration)
	PutCustomClaims(value interface{})
	PutPrivateEndpoint(value *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpoint)
	PutPrivateEndpointOverrides(value interface{})
	ResetAllowedAudience()
	ResetAllowedClients()
	ResetAllowedScopes()
	ResetAllowedWorkloadConfiguration()
	ResetCustomClaims()
	ResetDiscoveryUrl()
	ResetPrivateEndpoint()
	ResetPrivateEndpointOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference
type jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedAudience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedAudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedWorkloadConfiguration() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationOutputReference {
	var returns BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationOutputReference
	_jsii_.Get(
		j,
		"allowedWorkloadConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedWorkloadConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedWorkloadConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CustomClaims() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsList {
	var returns BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsList
	_jsii_.Get(
		j,
		"customClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CustomClaimsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) DiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpoint() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOutputReference {
	var returns BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOutputReference
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointOverrides() BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList {
	var returns BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList
	_jsii_.Get(
		j,
		"privateEndpointOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreRuntime.BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference_Override(b BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreRuntime.BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedAudience(val *[]*string) {
	if err := j.validateSetAllowedAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAudience",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedClients(val *[]*string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedScopes(val *[]*string) {
	if err := j.validateSetAllowedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedScopes",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetDiscoveryUrl(val *string) {
	if err := j.validateSetDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryUrl",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutAllowedWorkloadConfiguration(value *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration) {
	if err := b.validatePutAllowedWorkloadConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllowedWorkloadConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutCustomClaims(value interface{}) {
	if err := b.validatePutCustomClaimsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomClaims",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutPrivateEndpoint(value *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpoint) {
	if err := b.validatePutPrivateEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPrivateEndpoint",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutPrivateEndpointOverrides(value interface{}) {
	if err := b.validatePutPrivateEndpointOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPrivateEndpointOverrides",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedAudience() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedAudience",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedScopes() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedScopes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedWorkloadConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedWorkloadConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetCustomClaims() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomClaims",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetDiscoveryUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetDiscoveryUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetPrivateEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivateEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetPrivateEndpointOverrides() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivateEndpointOverrides",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

