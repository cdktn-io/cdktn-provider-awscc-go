// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package osispipeline

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OsisPipelineTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOsisPipelineTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

