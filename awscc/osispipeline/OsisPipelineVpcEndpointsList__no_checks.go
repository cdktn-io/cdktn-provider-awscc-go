// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package osispipeline

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OsisPipelineVpcEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineVpcEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineVpcEndpointsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcEndpointsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOsisPipelineVpcEndpointsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

