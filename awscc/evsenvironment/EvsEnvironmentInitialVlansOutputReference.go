// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evsenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvsEnvironmentInitialVlansOutputReference interface {
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
	EdgeVTep() EvsEnvironmentInitialVlansEdgeVTepOutputReference
	EdgeVTepInput() interface{}
	ExpansionVlan1() EvsEnvironmentInitialVlansExpansionVlan1OutputReference
	ExpansionVlan1Input() interface{}
	ExpansionVlan2() EvsEnvironmentInitialVlansExpansionVlan2OutputReference
	ExpansionVlan2Input() interface{}
	// Experimental.
	Fqn() *string
	Hcx() EvsEnvironmentInitialVlansHcxOutputReference
	HcxInput() interface{}
	HcxNetworkAclId() *string
	SetHcxNetworkAclId(val *string)
	HcxNetworkAclIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsHcxPublic() interface{}
	SetIsHcxPublic(val interface{})
	IsHcxPublicInput() interface{}
	NsxUpLink() EvsEnvironmentInitialVlansNsxUpLinkOutputReference
	NsxUpLinkInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmkManagement() EvsEnvironmentInitialVlansVmkManagementOutputReference
	VmkManagementInput() interface{}
	VmManagement() EvsEnvironmentInitialVlansVmManagementOutputReference
	VmManagementInput() interface{}
	VMotion() EvsEnvironmentInitialVlansVMotionOutputReference
	VMotionInput() interface{}
	VSan() EvsEnvironmentInitialVlansVSanOutputReference
	VSanInput() interface{}
	VTep() EvsEnvironmentInitialVlansVTepOutputReference
	VTepInput() interface{}
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
	PutEdgeVTep(value *EvsEnvironmentInitialVlansEdgeVTep)
	PutExpansionVlan1(value *EvsEnvironmentInitialVlansExpansionVlan1)
	PutExpansionVlan2(value *EvsEnvironmentInitialVlansExpansionVlan2)
	PutHcx(value *EvsEnvironmentInitialVlansHcx)
	PutNsxUpLink(value *EvsEnvironmentInitialVlansNsxUpLink)
	PutVmkManagement(value *EvsEnvironmentInitialVlansVmkManagement)
	PutVmManagement(value *EvsEnvironmentInitialVlansVmManagement)
	PutVMotion(value *EvsEnvironmentInitialVlansVMotion)
	PutVSan(value *EvsEnvironmentInitialVlansVSan)
	PutVTep(value *EvsEnvironmentInitialVlansVTep)
	ResetEdgeVTep()
	ResetExpansionVlan1()
	ResetExpansionVlan2()
	ResetHcx()
	ResetHcxNetworkAclId()
	ResetIsHcxPublic()
	ResetNsxUpLink()
	ResetVmkManagement()
	ResetVmManagement()
	ResetVMotion()
	ResetVSan()
	ResetVTep()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EvsEnvironmentInitialVlansOutputReference
type jsiiProxy_EvsEnvironmentInitialVlansOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) EdgeVTep() EvsEnvironmentInitialVlansEdgeVTepOutputReference {
	var returns EvsEnvironmentInitialVlansEdgeVTepOutputReference
	_jsii_.Get(
		j,
		"edgeVTep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) EdgeVTepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeVTepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ExpansionVlan1() EvsEnvironmentInitialVlansExpansionVlan1OutputReference {
	var returns EvsEnvironmentInitialVlansExpansionVlan1OutputReference
	_jsii_.Get(
		j,
		"expansionVlan1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ExpansionVlan1Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expansionVlan1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ExpansionVlan2() EvsEnvironmentInitialVlansExpansionVlan2OutputReference {
	var returns EvsEnvironmentInitialVlansExpansionVlan2OutputReference
	_jsii_.Get(
		j,
		"expansionVlan2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ExpansionVlan2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expansionVlan2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) Hcx() EvsEnvironmentInitialVlansHcxOutputReference {
	var returns EvsEnvironmentInitialVlansHcxOutputReference
	_jsii_.Get(
		j,
		"hcx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) HcxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hcxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) HcxNetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcxNetworkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) HcxNetworkAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcxNetworkAclIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) IsHcxPublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHcxPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) IsHcxPublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHcxPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) NsxUpLink() EvsEnvironmentInitialVlansNsxUpLinkOutputReference {
	var returns EvsEnvironmentInitialVlansNsxUpLinkOutputReference
	_jsii_.Get(
		j,
		"nsxUpLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) NsxUpLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nsxUpLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VmkManagement() EvsEnvironmentInitialVlansVmkManagementOutputReference {
	var returns EvsEnvironmentInitialVlansVmkManagementOutputReference
	_jsii_.Get(
		j,
		"vmkManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VmkManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmkManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VmManagement() EvsEnvironmentInitialVlansVmManagementOutputReference {
	var returns EvsEnvironmentInitialVlansVmManagementOutputReference
	_jsii_.Get(
		j,
		"vmManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VmManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VMotion() EvsEnvironmentInitialVlansVMotionOutputReference {
	var returns EvsEnvironmentInitialVlansVMotionOutputReference
	_jsii_.Get(
		j,
		"vMotion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VMotionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vMotionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VSan() EvsEnvironmentInitialVlansVSanOutputReference {
	var returns EvsEnvironmentInitialVlansVSanOutputReference
	_jsii_.Get(
		j,
		"vSan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VSanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vSanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VTep() EvsEnvironmentInitialVlansVTepOutputReference {
	var returns EvsEnvironmentInitialVlansVTepOutputReference
	_jsii_.Get(
		j,
		"vTep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) VTepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vTepInput",
		&returns,
	)
	return returns
}


func NewEvsEnvironmentInitialVlansOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EvsEnvironmentInitialVlansOutputReference {
	_init_.Initialize()

	if err := validateNewEvsEnvironmentInitialVlansOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvsEnvironmentInitialVlansOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironmentInitialVlansOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEvsEnvironmentInitialVlansOutputReference_Override(e EvsEnvironmentInitialVlansOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironmentInitialVlansOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetHcxNetworkAclId(val *string) {
	if err := j.validateSetHcxNetworkAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hcxNetworkAclId",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetIsHcxPublic(val interface{}) {
	if err := j.validateSetIsHcxPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHcxPublic",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentInitialVlansOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutEdgeVTep(value *EvsEnvironmentInitialVlansEdgeVTep) {
	if err := e.validatePutEdgeVTepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEdgeVTep",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutExpansionVlan1(value *EvsEnvironmentInitialVlansExpansionVlan1) {
	if err := e.validatePutExpansionVlan1Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExpansionVlan1",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutExpansionVlan2(value *EvsEnvironmentInitialVlansExpansionVlan2) {
	if err := e.validatePutExpansionVlan2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExpansionVlan2",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutHcx(value *EvsEnvironmentInitialVlansHcx) {
	if err := e.validatePutHcxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHcx",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutNsxUpLink(value *EvsEnvironmentInitialVlansNsxUpLink) {
	if err := e.validatePutNsxUpLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNsxUpLink",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutVmkManagement(value *EvsEnvironmentInitialVlansVmkManagement) {
	if err := e.validatePutVmkManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVmkManagement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutVmManagement(value *EvsEnvironmentInitialVlansVmManagement) {
	if err := e.validatePutVmManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVmManagement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutVMotion(value *EvsEnvironmentInitialVlansVMotion) {
	if err := e.validatePutVMotionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVMotion",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutVSan(value *EvsEnvironmentInitialVlansVSan) {
	if err := e.validatePutVSanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVSan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) PutVTep(value *EvsEnvironmentInitialVlansVTep) {
	if err := e.validatePutVTepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVTep",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetEdgeVTep() {
	_jsii_.InvokeVoid(
		e,
		"resetEdgeVTep",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetExpansionVlan1() {
	_jsii_.InvokeVoid(
		e,
		"resetExpansionVlan1",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetExpansionVlan2() {
	_jsii_.InvokeVoid(
		e,
		"resetExpansionVlan2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetHcx() {
	_jsii_.InvokeVoid(
		e,
		"resetHcx",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetHcxNetworkAclId() {
	_jsii_.InvokeVoid(
		e,
		"resetHcxNetworkAclId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetIsHcxPublic() {
	_jsii_.InvokeVoid(
		e,
		"resetIsHcxPublic",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetNsxUpLink() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxUpLink",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetVmkManagement() {
	_jsii_.InvokeVoid(
		e,
		"resetVmkManagement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetVmManagement() {
	_jsii_.InvokeVoid(
		e,
		"resetVmManagement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetVMotion() {
	_jsii_.InvokeVoid(
		e,
		"resetVMotion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetVSan() {
	_jsii_.InvokeVoid(
		e,
		"resetVSan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ResetVTep() {
	_jsii_.InvokeVoid(
		e,
		"resetVTep",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EvsEnvironmentInitialVlansOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

