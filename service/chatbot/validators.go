// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateChimeWebhookConfiguration struct {
}

func (*validateOpCreateChimeWebhookConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateChimeWebhookConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateChimeWebhookConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateChimeWebhookConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMicrosoftTeamsChannelConfiguration struct {
}

func (*validateOpCreateMicrosoftTeamsChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMicrosoftTeamsChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMicrosoftTeamsChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMicrosoftTeamsChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSlackChannelConfiguration struct {
}

func (*validateOpCreateSlackChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSlackChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSlackChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSlackChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteChimeWebhookConfiguration struct {
}

func (*validateOpDeleteChimeWebhookConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteChimeWebhookConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteChimeWebhookConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteChimeWebhookConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMicrosoftTeamsChannelConfiguration struct {
}

func (*validateOpDeleteMicrosoftTeamsChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMicrosoftTeamsChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMicrosoftTeamsChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMicrosoftTeamsChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMicrosoftTeamsConfiguredTeam struct {
}

func (*validateOpDeleteMicrosoftTeamsConfiguredTeam) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMicrosoftTeamsConfiguredTeam) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMicrosoftTeamsConfiguredTeamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMicrosoftTeamsConfiguredTeamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMicrosoftTeamsUserIdentity struct {
}

func (*validateOpDeleteMicrosoftTeamsUserIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMicrosoftTeamsUserIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMicrosoftTeamsUserIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMicrosoftTeamsUserIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSlackChannelConfiguration struct {
}

func (*validateOpDeleteSlackChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSlackChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSlackChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSlackChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSlackUserIdentity struct {
}

func (*validateOpDeleteSlackUserIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSlackUserIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSlackUserIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSlackUserIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSlackWorkspaceAuthorization struct {
}

func (*validateOpDeleteSlackWorkspaceAuthorization) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSlackWorkspaceAuthorization) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSlackWorkspaceAuthorizationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSlackWorkspaceAuthorizationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMicrosoftTeamsChannelConfiguration struct {
}

func (*validateOpGetMicrosoftTeamsChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMicrosoftTeamsChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMicrosoftTeamsChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMicrosoftTeamsChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateChimeWebhookConfiguration struct {
}

func (*validateOpUpdateChimeWebhookConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateChimeWebhookConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateChimeWebhookConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateChimeWebhookConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMicrosoftTeamsChannelConfiguration struct {
}

func (*validateOpUpdateMicrosoftTeamsChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMicrosoftTeamsChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMicrosoftTeamsChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMicrosoftTeamsChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSlackChannelConfiguration struct {
}

func (*validateOpUpdateSlackChannelConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSlackChannelConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSlackChannelConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSlackChannelConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateChimeWebhookConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateChimeWebhookConfiguration{}, middleware.After)
}

func addOpCreateMicrosoftTeamsChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMicrosoftTeamsChannelConfiguration{}, middleware.After)
}

func addOpCreateSlackChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSlackChannelConfiguration{}, middleware.After)
}

func addOpDeleteChimeWebhookConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteChimeWebhookConfiguration{}, middleware.After)
}

func addOpDeleteMicrosoftTeamsChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMicrosoftTeamsChannelConfiguration{}, middleware.After)
}

func addOpDeleteMicrosoftTeamsConfiguredTeamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMicrosoftTeamsConfiguredTeam{}, middleware.After)
}

func addOpDeleteMicrosoftTeamsUserIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMicrosoftTeamsUserIdentity{}, middleware.After)
}

func addOpDeleteSlackChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSlackChannelConfiguration{}, middleware.After)
}

func addOpDeleteSlackUserIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSlackUserIdentity{}, middleware.After)
}

func addOpDeleteSlackWorkspaceAuthorizationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSlackWorkspaceAuthorization{}, middleware.After)
}

func addOpGetMicrosoftTeamsChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMicrosoftTeamsChannelConfiguration{}, middleware.After)
}

func addOpUpdateChimeWebhookConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateChimeWebhookConfiguration{}, middleware.After)
}

func addOpUpdateMicrosoftTeamsChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMicrosoftTeamsChannelConfiguration{}, middleware.After)
}

func addOpUpdateSlackChannelConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSlackChannelConfiguration{}, middleware.After)
}

func validateOpCreateChimeWebhookConfigurationInput(v *CreateChimeWebhookConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateChimeWebhookConfigurationInput"}
	if v.WebhookDescription == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WebhookDescription"))
	}
	if v.WebhookUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WebhookUrl"))
	}
	if v.SnsTopicArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsTopicArns"))
	}
	if v.IamRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IamRoleArn"))
	}
	if v.ConfigurationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMicrosoftTeamsChannelConfigurationInput(v *CreateMicrosoftTeamsChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMicrosoftTeamsChannelConfigurationInput"}
	if v.ChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelId"))
	}
	if v.TeamId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TeamId"))
	}
	if v.TenantId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TenantId"))
	}
	if v.IamRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IamRoleArn"))
	}
	if v.ConfigurationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSlackChannelConfigurationInput(v *CreateSlackChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSlackChannelConfigurationInput"}
	if v.SlackTeamId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackTeamId"))
	}
	if v.SlackChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackChannelId"))
	}
	if v.IamRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IamRoleArn"))
	}
	if v.ConfigurationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteChimeWebhookConfigurationInput(v *DeleteChimeWebhookConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteChimeWebhookConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMicrosoftTeamsChannelConfigurationInput(v *DeleteMicrosoftTeamsChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMicrosoftTeamsChannelConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMicrosoftTeamsConfiguredTeamInput(v *DeleteMicrosoftTeamsConfiguredTeamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMicrosoftTeamsConfiguredTeamInput"}
	if v.TeamId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TeamId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMicrosoftTeamsUserIdentityInput(v *DeleteMicrosoftTeamsUserIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMicrosoftTeamsUserIdentityInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSlackChannelConfigurationInput(v *DeleteSlackChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSlackChannelConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSlackUserIdentityInput(v *DeleteSlackUserIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSlackUserIdentityInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if v.SlackTeamId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackTeamId"))
	}
	if v.SlackUserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackUserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSlackWorkspaceAuthorizationInput(v *DeleteSlackWorkspaceAuthorizationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSlackWorkspaceAuthorizationInput"}
	if v.SlackTeamId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackTeamId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMicrosoftTeamsChannelConfigurationInput(v *GetMicrosoftTeamsChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMicrosoftTeamsChannelConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateChimeWebhookConfigurationInput(v *UpdateChimeWebhookConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateChimeWebhookConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMicrosoftTeamsChannelConfigurationInput(v *UpdateMicrosoftTeamsChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMicrosoftTeamsChannelConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if v.ChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSlackChannelConfigurationInput(v *UpdateSlackChannelConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSlackChannelConfigurationInput"}
	if v.ChatConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChatConfigurationArn"))
	}
	if v.SlackChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SlackChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
