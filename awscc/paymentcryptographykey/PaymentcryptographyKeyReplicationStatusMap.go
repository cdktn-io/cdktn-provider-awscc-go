// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/paymentcryptographykey/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PaymentcryptographyKeyReplicationStatusMap interface {
	cdktn.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(key *string) PaymentcryptographyKeyReplicationStatusOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaymentcryptographyKeyReplicationStatusMap
type jsiiProxy_PaymentcryptographyKeyReplicationStatusMap struct {
	internal.Type__cdktnComplexMap
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPaymentcryptographyKeyReplicationStatusMap(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PaymentcryptographyKeyReplicationStatusMap {
	_init_.Initialize()

	if err := validateNewPaymentcryptographyKeyReplicationStatusMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaymentcryptographyKeyReplicationStatusMap{}

	_jsii_.Create(
		"@cdktn/provider-awscc.paymentcryptographyKey.PaymentcryptographyKeyReplicationStatusMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaymentcryptographyKeyReplicationStatusMap_Override(p PaymentcryptographyKeyReplicationStatusMap, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.paymentcryptographyKey.PaymentcryptographyKeyReplicationStatusMap",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) Get(key *string) PaymentcryptographyKeyReplicationStatusOutputReference {
	if err := p.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns PaymentcryptographyKeyReplicationStatusOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) InterpolationForAttribute(property *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

