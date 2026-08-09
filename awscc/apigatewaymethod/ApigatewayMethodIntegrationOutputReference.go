// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaymethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewaymethod/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayMethodIntegrationOutputReference interface {
	cdktn.ComplexObject
	CacheKeyParameters() *[]*string
	SetCacheKeyParameters(val *[]*string)
	CacheKeyParametersInput() *[]*string
	CacheNamespace() *string
	SetCacheNamespace(val *string)
	CacheNamespaceInput() *string
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
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	ContentHandling() *string
	SetContentHandling(val *string)
	ContentHandlingInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	// Experimental.
	Fqn() *string
	IntegrationHttpMethod() *string
	SetIntegrationHttpMethod(val *string)
	IntegrationHttpMethodInput() *string
	IntegrationResponses() ApigatewayMethodIntegrationIntegrationResponsesList
	IntegrationResponsesInput() interface{}
	IntegrationTarget() *string
	SetIntegrationTarget(val *string)
	IntegrationTargetInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	PassthroughBehaviorInput() *string
	RequestParameters() *map[string]*string
	SetRequestParameters(val *map[string]*string)
	RequestParametersInput() *map[string]*string
	RequestTemplates() *map[string]*string
	SetRequestTemplates(val *map[string]*string)
	RequestTemplatesInput() *map[string]*string
	ResponseTransferMode() *string
	SetResponseTransferMode(val *string)
	ResponseTransferModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutInMillis() *float64
	SetTimeoutInMillis(val *float64)
	TimeoutInMillisInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutIntegrationResponses(value interface{})
	ResetCacheKeyParameters()
	ResetCacheNamespace()
	ResetConnectionId()
	ResetConnectionType()
	ResetContentHandling()
	ResetCredentials()
	ResetIntegrationHttpMethod()
	ResetIntegrationResponses()
	ResetIntegrationTarget()
	ResetPassthroughBehavior()
	ResetRequestParameters()
	ResetRequestTemplates()
	ResetResponseTransferMode()
	ResetTimeoutInMillis()
	ResetType()
	ResetUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigatewayMethodIntegrationOutputReference
type jsiiProxy_ApigatewayMethodIntegrationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CacheKeyParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CacheKeyParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CacheNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CacheNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ContentHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ContentHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationHttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationHttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationResponses() ApigatewayMethodIntegrationIntegrationResponsesList {
	var returns ApigatewayMethodIntegrationIntegrationResponsesList
	_jsii_.Get(
		j,
		"integrationResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integrationResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) IntegrationTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) PassthroughBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) RequestParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) RequestParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) RequestTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) RequestTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResponseTransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTransferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResponseTransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTransferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) TimeoutInMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) TimeoutInMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewApigatewayMethodIntegrationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApigatewayMethodIntegrationOutputReference {
	_init_.Initialize()

	if err := validateNewApigatewayMethodIntegrationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayMethodIntegrationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayMethod.ApigatewayMethodIntegrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayMethodIntegrationOutputReference_Override(a ApigatewayMethodIntegrationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayMethod.ApigatewayMethodIntegrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetCacheKeyParameters(val *[]*string) {
	if err := j.validateSetCacheKeyParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheKeyParameters",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetCacheNamespace(val *string) {
	if err := j.validateSetCacheNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheNamespace",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetContentHandling(val *string) {
	if err := j.validateSetContentHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentHandling",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetCredentials(val *string) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetIntegrationHttpMethod(val *string) {
	if err := j.validateSetIntegrationHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationHttpMethod",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetIntegrationTarget(val *string) {
	if err := j.validateSetIntegrationTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationTarget",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetPassthroughBehavior(val *string) {
	if err := j.validateSetPassthroughBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetRequestParameters(val *map[string]*string) {
	if err := j.validateSetRequestParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetRequestTemplates(val *map[string]*string) {
	if err := j.validateSetRequestTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetResponseTransferMode(val *string) {
	if err := j.validateSetResponseTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTransferMode",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetTimeoutInMillis(val *float64) {
	if err := j.validateSetTimeoutInMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMillis",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ApigatewayMethodIntegrationOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) PutIntegrationResponses(value interface{}) {
	if err := a.validatePutIntegrationResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegrationResponses",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetCacheKeyParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheKeyParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetCacheNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetConnectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetContentHandling() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetIntegrationHttpMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationHttpMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetIntegrationResponses() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationResponses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetIntegrationTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetPassthroughBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPassthroughBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetRequestTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetResponseTransferMode() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTransferMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetTimeoutInMillis() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInMillis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ResetUri() {
	_jsii_.InvokeVoid(
		a,
		"resetUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayMethodIntegrationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

