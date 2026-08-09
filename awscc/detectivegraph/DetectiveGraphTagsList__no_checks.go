// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package detectivegraph

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DetectiveGraphTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DetectiveGraphTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DetectiveGraphTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DetectiveGraphTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DetectiveGraphTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DetectiveGraphTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DetectiveGraphTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDetectiveGraphTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

