// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package wisdomcontent

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WisdomContentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WisdomContentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WisdomContentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WisdomContentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WisdomContentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WisdomContentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WisdomContentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWisdomContentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

