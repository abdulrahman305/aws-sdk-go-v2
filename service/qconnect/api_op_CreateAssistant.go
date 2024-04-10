// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q in Connect assistant.
func (c *Client) CreateAssistant(ctx context.Context, params *CreateAssistantInput, optFns ...func(*Options)) (*CreateAssistantOutput, error) {
	if params == nil {
		params = &CreateAssistantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssistant", params, optFns, c.addOperationCreateAssistantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssistantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssistantInput struct {

	// The name of the assistant.
	//
	// This member is required.
	Name *string

	// The type of assistant.
	//
	// This member is required.
	Type types.AssistantType

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The description of the assistant.
	Description *string

	// The configuration information for the customer managed key used for encryption.
	// The customer managed key must have a policy that allows kms:CreateGrant ,
	// kms:DescribeKey , kms:Decrypt , and kms:GenerateDataKey* permissions to the IAM
	// identity using the key to invoke Amazon Q in Connect. To use Amazon Q in Connect
	// with chat, the key policy must also allow kms:Decrypt , kms:GenerateDataKey* ,
	// and kms:DescribeKey permissions to the connect.amazonaws.com service principal.
	// For more information about setting up a customer managed key for Amazon Q in
	// Connect, see Enable Amazon Q in Connect for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-q.html)
	// .
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssistantOutput struct {

	// Information about the assistant.
	Assistant *types.AssistantData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssistantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAssistant"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateAssistantMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAssistantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssistant(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateAssistant struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAssistant) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAssistant) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssistantInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssistantInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAssistantMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAssistant{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAssistant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAssistant",
	}
}
