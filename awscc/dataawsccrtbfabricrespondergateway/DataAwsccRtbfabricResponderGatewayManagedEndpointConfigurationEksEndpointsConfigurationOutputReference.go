// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccrtbfabricrespondergateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccrtbfabricrespondergateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference interface {
	cdktn.ComplexObject
	ClusterApiServerCaCertificateChain() *string
	ClusterApiServerEndpointUri() *string
	ClusterName() *string
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
	EndpointsResourceNamespace() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration
	SetInternalValue(val *DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration)
	RoleArn() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference
type jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerCaCertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerCaCertificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterApiServerEndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApiServerEndpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) EndpointsResourceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointsResourceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InternalValue() *DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration {
	var returns *DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRtbfabricResponderGateway.DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference_Override(d DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRtbfabricResponderGateway.DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetInternalValue(val *DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

