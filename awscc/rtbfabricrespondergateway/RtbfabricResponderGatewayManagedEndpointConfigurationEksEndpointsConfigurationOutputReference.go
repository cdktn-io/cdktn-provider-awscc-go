// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricrespondergateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabricrespondergateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference interface {
	cdktn.ComplexObject
	ClusterApiServerCaCertificateChain() *string
	SetClusterApiServerCaCertificateChain(val *string)
	ClusterApiServerCaCertificateChainInput() *string
	ClusterApiServerEndpointUri() *string
	SetClusterApiServerEndpointUri(val *string)
	ClusterApiServerEndpointUriInput() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	EndpointsResourceName() *string
	SetEndpointsResourceName(val *string)
	EndpointsResourceNameInput() *string
	EndpointsResourceNamespace() *string
	SetEndpointsResourceNamespace(val *string)
	EndpointsResourceNamespaceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetClusterApiServerCaCertificateChain()
	ResetClusterApiServerEndpointUri()
	ResetClusterName()
	ResetEndpointsResourceName()
	ResetEndpointsResourceNamespace()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference
type jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerCaCertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerCaCertificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerCaCertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerCaCertificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerEndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerEndpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerEndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerEndpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricResponderGateway.RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference_Override(r RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricResponderGateway.RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetClusterApiServerCaCertificateChain(val *string) {
	if err := j.validateSetClusterApiServerCaCertificateChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterApiServerCaCertificateChain",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetClusterApiServerEndpointUri(val *string) {
	if err := j.validateSetClusterApiServerEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterApiServerEndpointUri",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetEndpointsResourceName(val *string) {
	if err := j.validateSetEndpointsResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointsResourceName",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetEndpointsResourceNamespace(val *string) {
	if err := j.validateSetEndpointsResourceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointsResourceNamespace",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetClusterApiServerCaCertificateChain() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterApiServerCaCertificateChain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetClusterApiServerEndpointUri() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterApiServerEndpointUri",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetEndpointsResourceName() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpointsResourceName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetEndpointsResourceNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpointsResourceNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

