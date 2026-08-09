// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/grafanaworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GrafanaWorkspaceSamlConfigurationOutputReference interface {
	cdktn.ComplexObject
	AllowedOrganizations() *[]*string
	SetAllowedOrganizations(val *[]*string)
	AllowedOrganizationsInput() *[]*string
	AssertionAttributes() GrafanaWorkspaceSamlConfigurationAssertionAttributesOutputReference
	AssertionAttributesInput() interface{}
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
	IdpMetadata() GrafanaWorkspaceSamlConfigurationIdpMetadataOutputReference
	IdpMetadataInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoginValidityDuration() *float64
	SetLoginValidityDuration(val *float64)
	LoginValidityDurationInput() *float64
	RoleValues() GrafanaWorkspaceSamlConfigurationRoleValuesOutputReference
	RoleValuesInput() interface{}
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
	PutAssertionAttributes(value *GrafanaWorkspaceSamlConfigurationAssertionAttributes)
	PutIdpMetadata(value *GrafanaWorkspaceSamlConfigurationIdpMetadata)
	PutRoleValues(value *GrafanaWorkspaceSamlConfigurationRoleValues)
	ResetAllowedOrganizations()
	ResetAssertionAttributes()
	ResetIdpMetadata()
	ResetLoginValidityDuration()
	ResetRoleValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GrafanaWorkspaceSamlConfigurationOutputReference
type jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) AllowedOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) AllowedOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) AssertionAttributes() GrafanaWorkspaceSamlConfigurationAssertionAttributesOutputReference {
	var returns GrafanaWorkspaceSamlConfigurationAssertionAttributesOutputReference
	_jsii_.Get(
		j,
		"assertionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) AssertionAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) IdpMetadata() GrafanaWorkspaceSamlConfigurationIdpMetadataOutputReference {
	var returns GrafanaWorkspaceSamlConfigurationIdpMetadataOutputReference
	_jsii_.Get(
		j,
		"idpMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) IdpMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idpMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) LoginValidityDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) LoginValidityDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) RoleValues() GrafanaWorkspaceSamlConfigurationRoleValuesOutputReference {
	var returns GrafanaWorkspaceSamlConfigurationRoleValuesOutputReference
	_jsii_.Get(
		j,
		"roleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) RoleValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGrafanaWorkspaceSamlConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GrafanaWorkspaceSamlConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGrafanaWorkspaceSamlConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.grafanaWorkspace.GrafanaWorkspaceSamlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGrafanaWorkspaceSamlConfigurationOutputReference_Override(g GrafanaWorkspaceSamlConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.grafanaWorkspace.GrafanaWorkspaceSamlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetAllowedOrganizations(val *[]*string) {
	if err := j.validateSetAllowedOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrganizations",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetLoginValidityDuration(val *float64) {
	if err := j.validateSetLoginValidityDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginValidityDuration",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) PutAssertionAttributes(value *GrafanaWorkspaceSamlConfigurationAssertionAttributes) {
	if err := g.validatePutAssertionAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAssertionAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) PutIdpMetadata(value *GrafanaWorkspaceSamlConfigurationIdpMetadata) {
	if err := g.validatePutIdpMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdpMetadata",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) PutRoleValues(value *GrafanaWorkspaceSamlConfigurationRoleValues) {
	if err := g.validatePutRoleValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRoleValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ResetAllowedOrganizations() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedOrganizations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ResetAssertionAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAssertionAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ResetIdpMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetIdpMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ResetLoginValidityDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetLoginValidityDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ResetRoleValues() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

