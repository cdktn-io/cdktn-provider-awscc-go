// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mskreplicator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskReplicatorKafkaClustersOutputReference interface {
	cdktn.ComplexObject
	AmazonMskCluster() MskReplicatorKafkaClustersAmazonMskClusterOutputReference
	AmazonMskClusterInput() interface{}
	ApacheKafkaCluster() MskReplicatorKafkaClustersApacheKafkaClusterOutputReference
	ApacheKafkaClusterInput() interface{}
	ClientAuthentication() MskReplicatorKafkaClustersClientAuthenticationOutputReference
	ClientAuthenticationInput() interface{}
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
	EncryptionInTransit() MskReplicatorKafkaClustersEncryptionInTransitOutputReference
	EncryptionInTransitInput() interface{}
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
	VpcConfig() MskReplicatorKafkaClustersVpcConfigOutputReference
	VpcConfigInput() interface{}
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
	PutAmazonMskCluster(value *MskReplicatorKafkaClustersAmazonMskCluster)
	PutApacheKafkaCluster(value *MskReplicatorKafkaClustersApacheKafkaCluster)
	PutClientAuthentication(value *MskReplicatorKafkaClustersClientAuthentication)
	PutEncryptionInTransit(value *MskReplicatorKafkaClustersEncryptionInTransit)
	PutVpcConfig(value *MskReplicatorKafkaClustersVpcConfig)
	ResetAmazonMskCluster()
	ResetApacheKafkaCluster()
	ResetClientAuthentication()
	ResetEncryptionInTransit()
	ResetVpcConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskReplicatorKafkaClustersOutputReference
type jsiiProxy_MskReplicatorKafkaClustersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) AmazonMskCluster() MskReplicatorKafkaClustersAmazonMskClusterOutputReference {
	var returns MskReplicatorKafkaClustersAmazonMskClusterOutputReference
	_jsii_.Get(
		j,
		"amazonMskCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) AmazonMskClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonMskClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ApacheKafkaCluster() MskReplicatorKafkaClustersApacheKafkaClusterOutputReference {
	var returns MskReplicatorKafkaClustersApacheKafkaClusterOutputReference
	_jsii_.Get(
		j,
		"apacheKafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ApacheKafkaClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apacheKafkaClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ClientAuthentication() MskReplicatorKafkaClustersClientAuthenticationOutputReference {
	var returns MskReplicatorKafkaClustersClientAuthenticationOutputReference
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ClientAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) EncryptionInTransit() MskReplicatorKafkaClustersEncryptionInTransitOutputReference {
	var returns MskReplicatorKafkaClustersEncryptionInTransitOutputReference
	_jsii_.Get(
		j,
		"encryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) EncryptionInTransitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) VpcConfig() MskReplicatorKafkaClustersVpcConfigOutputReference {
	var returns MskReplicatorKafkaClustersVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


func NewMskReplicatorKafkaClustersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MskReplicatorKafkaClustersOutputReference {
	_init_.Initialize()

	if err := validateNewMskReplicatorKafkaClustersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskReplicatorKafkaClustersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMskReplicatorKafkaClustersOutputReference_Override(m MskReplicatorKafkaClustersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) PutAmazonMskCluster(value *MskReplicatorKafkaClustersAmazonMskCluster) {
	if err := m.validatePutAmazonMskClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAmazonMskCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) PutApacheKafkaCluster(value *MskReplicatorKafkaClustersApacheKafkaCluster) {
	if err := m.validatePutApacheKafkaClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putApacheKafkaCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) PutClientAuthentication(value *MskReplicatorKafkaClustersClientAuthentication) {
	if err := m.validatePutClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientAuthentication",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) PutEncryptionInTransit(value *MskReplicatorKafkaClustersEncryptionInTransit) {
	if err := m.validatePutEncryptionInTransitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionInTransit",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) PutVpcConfig(value *MskReplicatorKafkaClustersVpcConfig) {
	if err := m.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ResetAmazonMskCluster() {
	_jsii_.InvokeVoid(
		m,
		"resetAmazonMskCluster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ResetApacheKafkaCluster() {
	_jsii_.InvokeVoid(
		m,
		"resetApacheKafkaCluster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ResetClientAuthentication() {
	_jsii_.InvokeVoid(
		m,
		"resetClientAuthentication",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ResetEncryptionInTransit() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionInTransit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

