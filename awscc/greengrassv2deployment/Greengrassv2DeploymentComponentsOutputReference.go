// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/greengrassv2deployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2DeploymentComponentsOutputReference interface {
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
	ComponentVersion() *string
	SetComponentVersion(val *string)
	ComponentVersionInput() *string
	ConfigurationUpdate() Greengrassv2DeploymentComponentsConfigurationUpdateOutputReference
	ConfigurationUpdateInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RunWith() Greengrassv2DeploymentComponentsRunWithOutputReference
	RunWithInput() interface{}
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
	PutConfigurationUpdate(value *Greengrassv2DeploymentComponentsConfigurationUpdate)
	PutRunWith(value *Greengrassv2DeploymentComponentsRunWith)
	ResetComponentVersion()
	ResetConfigurationUpdate()
	ResetRunWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2DeploymentComponentsOutputReference
type jsiiProxy_Greengrassv2DeploymentComponentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ComponentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ComponentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ConfigurationUpdate() Greengrassv2DeploymentComponentsConfigurationUpdateOutputReference {
	var returns Greengrassv2DeploymentComponentsConfigurationUpdateOutputReference
	_jsii_.Get(
		j,
		"configurationUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ConfigurationUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) RunWith() Greengrassv2DeploymentComponentsRunWithOutputReference {
	var returns Greengrassv2DeploymentComponentsRunWithOutputReference
	_jsii_.Get(
		j,
		"runWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) RunWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGreengrassv2DeploymentComponentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) Greengrassv2DeploymentComponentsOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentComponentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2DeploymentComponentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewGreengrassv2DeploymentComponentsOutputReference_Override(g Greengrassv2DeploymentComponentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetComponentVersion(val *string) {
	if err := j.validateSetComponentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentVersion",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) PutConfigurationUpdate(value *Greengrassv2DeploymentComponentsConfigurationUpdate) {
	if err := g.validatePutConfigurationUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigurationUpdate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) PutRunWith(value *Greengrassv2DeploymentComponentsRunWith) {
	if err := g.validatePutRunWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRunWith",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ResetComponentVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetComponentVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ResetConfigurationUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigurationUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ResetRunWith() {
	_jsii_.InvokeVoid(
		g,
		"resetRunWith",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_Greengrassv2DeploymentComponentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

