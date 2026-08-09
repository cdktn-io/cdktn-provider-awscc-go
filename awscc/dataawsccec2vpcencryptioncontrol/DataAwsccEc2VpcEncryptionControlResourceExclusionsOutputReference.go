// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2vpcencryptioncontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2vpcencryptioncontrol/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference interface {
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
	EgressOnlyInternetGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	ElasticFileSystem() DataAwsccEc2VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccEc2VpcEncryptionControlResourceExclusions
	SetInternalValue(val *DataAwsccEc2VpcEncryptionControlResourceExclusions)
	InternetGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	Lambda() DataAwsccEc2VpcEncryptionControlResourceExclusionsLambdaOutputReference
	NatGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VirtualPrivateGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	VpcLattice() DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	VpcPeering() DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
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

// The jsii proxy struct for DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference
type jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) EgressOnlyInternetGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) ElasticFileSystem() DataAwsccEc2VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	_jsii_.Get(
		j,
		"elasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) InternalValue() *DataAwsccEc2VpcEncryptionControlResourceExclusions {
	var returns *DataAwsccEc2VpcEncryptionControlResourceExclusions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) InternetGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) Lambda() DataAwsccEc2VpcEncryptionControlResourceExclusionsLambdaOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) NatGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsNatGatewayOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) VirtualPrivateGateway() DataAwsccEc2VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	_jsii_.Get(
		j,
		"virtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) VpcLattice() DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	_jsii_.Get(
		j,
		"vpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) VpcPeering() DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference {
	var returns DataAwsccEc2VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
	_jsii_.Get(
		j,
		"vpcPeering",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2VpcEncryptionControl.DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference_Override(d DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2VpcEncryptionControl.DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference)SetInternalValue(val *DataAwsccEc2VpcEncryptionControlResourceExclusions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2VpcEncryptionControlResourceExclusionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

