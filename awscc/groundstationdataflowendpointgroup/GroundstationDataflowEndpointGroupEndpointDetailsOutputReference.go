// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/groundstationdataflowendpointgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroundstationDataflowEndpointGroupEndpointDetailsOutputReference interface {
	cdktn.ComplexObject
	AwsGroundStationAgentEndpoint() GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointOutputReference
	AwsGroundStationAgentEndpointInput() interface{}
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
	Endpoint() GroundstationDataflowEndpointGroupEndpointDetailsEndpointOutputReference
	EndpointInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecurityDetails() GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetailsOutputReference
	SecurityDetailsInput() interface{}
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
	PutAwsGroundStationAgentEndpoint(value *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpoint)
	PutEndpoint(value *GroundstationDataflowEndpointGroupEndpointDetailsEndpoint)
	PutSecurityDetails(value *GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetails)
	ResetAwsGroundStationAgentEndpoint()
	ResetEndpoint()
	ResetSecurityDetails()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroundstationDataflowEndpointGroupEndpointDetailsOutputReference
type jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) AwsGroundStationAgentEndpoint() GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointOutputReference {
	var returns GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointOutputReference
	_jsii_.Get(
		j,
		"awsGroundStationAgentEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) AwsGroundStationAgentEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsGroundStationAgentEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) Endpoint() GroundstationDataflowEndpointGroupEndpointDetailsEndpointOutputReference {
	var returns GroundstationDataflowEndpointGroupEndpointDetailsEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) SecurityDetails() GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetailsOutputReference {
	var returns GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetailsOutputReference
	_jsii_.Get(
		j,
		"securityDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) SecurityDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroundstationDataflowEndpointGroupEndpointDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GroundstationDataflowEndpointGroupEndpointDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewGroundstationDataflowEndpointGroupEndpointDetailsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationDataflowEndpointGroup.GroundstationDataflowEndpointGroupEndpointDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGroundstationDataflowEndpointGroupEndpointDetailsOutputReference_Override(g GroundstationDataflowEndpointGroupEndpointDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.groundstationDataflowEndpointGroup.GroundstationDataflowEndpointGroupEndpointDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) PutAwsGroundStationAgentEndpoint(value *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpoint) {
	if err := g.validatePutAwsGroundStationAgentEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsGroundStationAgentEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) PutEndpoint(value *GroundstationDataflowEndpointGroupEndpointDetailsEndpoint) {
	if err := g.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) PutSecurityDetails(value *GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetails) {
	if err := g.validatePutSecurityDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityDetails",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ResetAwsGroundStationAgentEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsGroundStationAgentEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ResetSecurityDetails() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityDetails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GroundstationDataflowEndpointGroupEndpointDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

