// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package imagebuilderimage

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImagebuilderImageWorkflowsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageWorkflowsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageWorkflowsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImagebuilderImageWorkflowsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

