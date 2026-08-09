// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package medialivenode

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MedialiveNodeTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MedialiveNodeTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MedialiveNodeTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MedialiveNodeTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MedialiveNodeTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MedialiveNodeTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MedialiveNodeTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMedialiveNodeTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

