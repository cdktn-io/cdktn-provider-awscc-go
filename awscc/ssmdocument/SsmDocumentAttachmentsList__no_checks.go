// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssmdocument

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmDocumentAttachmentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentAttachmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentAttachmentsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentAttachmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmDocumentAttachmentsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

