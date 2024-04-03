// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Defines the information necessary to create an audience model. An audience
// model is a machine learning model that Clean Rooms ML trains to measure
// similarity between users. Clean Rooms ML manages training and storing the
// audience model. The audience model can be used in multiple calls to the
// StartAudienceGenerationJob API.
func (c *Client) CreateAudienceModel(ctx context.Context, params *CreateAudienceModelInput, optFns ...func(*Options)) (*CreateAudienceModelOutput, error) {
	if params == nil {
		params = &CreateAudienceModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAudienceModel", params, optFns, c.addOperationCreateAudienceModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAudienceModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAudienceModelInput struct {

	// The name of the audience model resource.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the training dataset for this audience model.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The description of the audience model.
	Description *string

	// The Amazon Resource Name (ARN) of the KMS key. This key is used to encrypt and
	// decrypt customer-owned data in the trained ML model and the associated data.
	KmsKeyArn *string

	// The optional metadata that you apply to the resource to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50.
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8.
	//   - Maximum value length - 256 Unicode characters in UTF-8.
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case sensitive.
	//   - Do not use aws:, AWS:, or any upper or lowercase combination of such as a
	//   prefix for keys as it is reserved for AWS use. You cannot edit or delete tag
	//   keys with this prefix. Values can have this prefix. If a tag value has aws as
	//   its prefix but the key does not, then Clean Rooms ML considers it to be a user
	//   tag and will count against the limit of 50 tags. Tags with only the key prefix
	//   of aws do not count against your tags per resource limit.
	Tags map[string]string

	// The end date and time of the training window.
	TrainingDataEndTime *time.Time

	// The start date and time of the training window.
	TrainingDataStartTime *time.Time

	noSmithyDocumentSerde
}

type CreateAudienceModelOutput struct {

	// The Amazon Resource Name (ARN) of the audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAudienceModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAudienceModel"); err != nil {
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
	if err = addOpCreateAudienceModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAudienceModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAudienceModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAudienceModel",
	}
}
