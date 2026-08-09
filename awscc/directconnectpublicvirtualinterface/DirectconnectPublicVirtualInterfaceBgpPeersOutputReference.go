// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectpublicvirtualinterface

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/directconnectpublicvirtualinterface/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectPublicVirtualInterfaceBgpPeersOutputReference interface {
	cdktn.ComplexObject
	AddressFamily() *string
	SetAddressFamily(val *string)
	AddressFamilyInput() *string
	AmazonAddress() *string
	SetAmazonAddress(val *string)
	AmazonAddressInput() *string
	Asn() *string
	SetAsn(val *string)
	AsnInput() *string
	AuthKey() *string
	SetAuthKey(val *string)
	AuthKeyInput() *string
	BgpPeerId() *string
	SetBgpPeerId(val *string)
	BgpPeerIdInput() *string
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
	CustomerAddress() *string
	SetCustomerAddress(val *string)
	CustomerAddressInput() *string
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
	ResetAmazonAddress()
	ResetAuthKey()
	ResetBgpPeerId()
	ResetCustomerAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DirectconnectPublicVirtualInterfaceBgpPeersOutputReference
type jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AmazonAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AmazonAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) Asn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AsnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AuthKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) AuthKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) BgpPeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpPeerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) BgpPeerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpPeerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) CustomerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) CustomerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDirectconnectPublicVirtualInterfaceBgpPeersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DirectconnectPublicVirtualInterfaceBgpPeersOutputReference {
	_init_.Initialize()

	if err := validateNewDirectconnectPublicVirtualInterfaceBgpPeersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.directconnectPublicVirtualInterface.DirectconnectPublicVirtualInterfaceBgpPeersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDirectconnectPublicVirtualInterfaceBgpPeersOutputReference_Override(d DirectconnectPublicVirtualInterfaceBgpPeersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.directconnectPublicVirtualInterface.DirectconnectPublicVirtualInterfaceBgpPeersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetAmazonAddress(val *string) {
	if err := j.validateSetAmazonAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonAddress",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetAsn(val *string) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetAuthKey(val *string) {
	if err := j.validateSetAuthKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authKey",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetBgpPeerId(val *string) {
	if err := j.validateSetBgpPeerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpPeerId",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetCustomerAddress(val *string) {
	if err := j.validateSetCustomerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAddress",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ResetAmazonAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetAmazonAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ResetAuthKey() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ResetBgpPeerId() {
	_jsii_.InvokeVoid(
		d,
		"resetBgpPeerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ResetCustomerAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DirectconnectPublicVirtualInterfaceBgpPeersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

