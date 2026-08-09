// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmeendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/certificatemanageracmeendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference interface {
	cdktn.ComplexObject
	AllowedKeyAlgorithms() *[]*string
	SetAllowedKeyAlgorithms(val *[]*string)
	AllowedKeyAlgorithmsInput() *[]*string
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
	ResetAllowedKeyAlgorithms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference
type jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) AllowedKeyAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKeyAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) AllowedKeyAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKeyAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference {
	_init_.Initialize()

	if err := validateNewCertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.certificatemanagerAcmeEndpoint.CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference_Override(c CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.certificatemanagerAcmeEndpoint.CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetAllowedKeyAlgorithms(val *[]*string) {
	if err := j.validateSetAllowedKeyAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedKeyAlgorithms",
		val,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) ResetAllowedKeyAlgorithms() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedKeyAlgorithms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatemanagerAcmeEndpointCertificateAuthorityPublicCertificateAuthorityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

