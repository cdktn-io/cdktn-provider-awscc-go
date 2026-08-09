// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qldbstream

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QldbStreamTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QldbStreamTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QldbStreamTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QldbStreamTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QldbStreamTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QldbStreamTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QldbStreamTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQldbStreamTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

