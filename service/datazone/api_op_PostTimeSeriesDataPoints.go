// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Posts time series data points to Amazon DataZone for the specified asset.
func (c *Client) PostTimeSeriesDataPoints(ctx context.Context, params *PostTimeSeriesDataPointsInput, optFns ...func(*Options)) (*PostTimeSeriesDataPointsOutput, error) {
	if params == nil {
		params = &PostTimeSeriesDataPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PostTimeSeriesDataPoints", params, optFns, c.addOperationPostTimeSeriesDataPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PostTimeSeriesDataPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostTimeSeriesDataPointsInput struct {

	// The ID of the Amazon DataZone domain in which you want to post time series data
	// points.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the asset for which you want to post time series data points.
	//
	// This member is required.
	EntityIdentifier *string

	// The type of the asset for which you want to post data points.
	//
	// This member is required.
	EntityType types.TimeSeriesEntityType

	// The forms that contain the data points that you want to post.
	//
	// This member is required.
	Forms []types.TimeSeriesDataPointFormInput

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

type PostTimeSeriesDataPointsOutput struct {

	// The ID of the Amazon DataZone domain in which you want to post time series data
	// points.
	DomainId *string

	// The ID of the asset for which you want to post time series data points.
	EntityId *string

	// The type of the asset for which you want to post data points.
	EntityType types.TimeSeriesEntityType

	// The forms that contain the data points that you have posted.
	Forms []types.TimeSeriesDataPointFormOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPostTimeSeriesDataPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPostTimeSeriesDataPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPostTimeSeriesDataPoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PostTimeSeriesDataPoints"); err != nil {
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
	if err = addIdempotencyToken_opPostTimeSeriesDataPointsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPostTimeSeriesDataPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPostTimeSeriesDataPoints(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPostTimeSeriesDataPoints struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPostTimeSeriesDataPoints) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPostTimeSeriesDataPoints) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PostTimeSeriesDataPointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PostTimeSeriesDataPointsInput ")
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
func addIdempotencyToken_opPostTimeSeriesDataPointsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPostTimeSeriesDataPoints{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPostTimeSeriesDataPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PostTimeSeriesDataPoints",
	}
}
