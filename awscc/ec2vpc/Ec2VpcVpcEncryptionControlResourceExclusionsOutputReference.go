// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference interface {
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
	EgressOnlyInternetGateway() Ec2VpcVpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	ElasticFileSystem() Ec2VpcVpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *Ec2VpcVpcEncryptionControlResourceExclusions
	SetInternalValue(val *Ec2VpcVpcEncryptionControlResourceExclusions)
	InternetGateway() Ec2VpcVpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	Lambda() Ec2VpcVpcEncryptionControlResourceExclusionsLambdaOutputReference
	NatGateway() Ec2VpcVpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VirtualPrivateGateway() Ec2VpcVpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	VpcLattice() Ec2VpcVpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	VpcPeering() Ec2VpcVpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
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

// The jsii proxy struct for Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference
type jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) EgressOnlyInternetGateway() Ec2VpcVpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) ElasticFileSystem() Ec2VpcVpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	_jsii_.Get(
		j,
		"elasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) InternalValue() *Ec2VpcVpcEncryptionControlResourceExclusions {
	var returns *Ec2VpcVpcEncryptionControlResourceExclusions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) InternetGateway() Ec2VpcVpcEncryptionControlResourceExclusionsInternetGatewayOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) Lambda() Ec2VpcVpcEncryptionControlResourceExclusionsLambdaOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) NatGateway() Ec2VpcVpcEncryptionControlResourceExclusionsNatGatewayOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) VirtualPrivateGateway() Ec2VpcVpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	_jsii_.Get(
		j,
		"virtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) VpcLattice() Ec2VpcVpcEncryptionControlResourceExclusionsVpcLatticeOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	_jsii_.Get(
		j,
		"vpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) VpcPeering() Ec2VpcVpcEncryptionControlResourceExclusionsVpcPeeringOutputReference {
	var returns Ec2VpcVpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
	_jsii_.Get(
		j,
		"vpcPeering",
		&returns,
	)
	return returns
}


func NewEc2VpcVpcEncryptionControlResourceExclusionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VpcVpcEncryptionControlResourceExclusionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Vpc.Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VpcVpcEncryptionControlResourceExclusionsOutputReference_Override(e Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Vpc.Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference)SetInternalValue(val *Ec2VpcVpcEncryptionControlResourceExclusions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VpcVpcEncryptionControlResourceExclusionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

