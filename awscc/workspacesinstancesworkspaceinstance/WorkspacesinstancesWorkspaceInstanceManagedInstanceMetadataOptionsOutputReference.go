// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesinstancesworkspaceinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HttpEndpoint() *string
	SetHttpEndpoint(val *string)
	HttpEndpointInput() *string
	HttpProtocolIpv6() *string
	SetHttpProtocolIpv6(val *string)
	HttpProtocolIpv6Input() *string
	HttpPutResponseHopLimit() *float64
	SetHttpPutResponseHopLimit(val *float64)
	HttpPutResponseHopLimitInput() *float64
	HttpTokens() *string
	SetHttpTokens(val *string)
	HttpTokensInput() *string
	InstanceMetadataTags() *string
	SetInstanceMetadataTags(val *string)
	InstanceMetadataTagsInput() *string
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
	ResetHttpEndpoint()
	ResetHttpProtocolIpv6()
	ResetHttpPutResponseHopLimit()
	ResetHttpTokens()
	ResetInstanceMetadataTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference
type jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpProtocolIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpProtocolIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpPutResponseHopLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpPutResponseHopLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) HttpTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) InstanceMetadataTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) InstanceMetadataTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference_Override(w WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetHttpEndpoint(val *string) {
	if err := j.validateSetHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpoint",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetHttpProtocolIpv6(val *string) {
	if err := j.validateSetHttpProtocolIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpProtocolIpv6",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetHttpPutResponseHopLimit(val *float64) {
	if err := j.validateSetHttpPutResponseHopLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPutResponseHopLimit",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetHttpTokens(val *string) {
	if err := j.validateSetHttpTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpTokens",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetInstanceMetadataTags(val *string) {
	if err := j.validateSetInstanceMetadataTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMetadataTags",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ResetHttpProtocolIpv6() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpProtocolIpv6",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ResetHttpPutResponseHopLimit() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpPutResponseHopLimit",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ResetHttpTokens() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpTokens",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ResetInstanceMetadataTags() {
	_jsii_.InvokeVoid(
		w,
		"resetInstanceMetadataTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

