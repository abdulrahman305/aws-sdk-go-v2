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

// Lists the content associations.
//
// For more information about content associations--what they are and when they
// are used--see [Integrate Amazon Q in Connect with step-by-step guides]in the Amazon Connect Administrator Guide.
//
// [Integrate Amazon Q in Connect with step-by-step guides]: https://docs.aws.amazon.com/connect/latest/adminguide/integrate-q-with-guides.html
func (c *Client) ListContentAssociations(ctx context.Context, params *ListContentAssociationsInput, optFns ...func(*Options)) (*ListContentAssociationsOutput, error) {
	if params == nil {
		params = &ListContentAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContentAssociations", params, optFns, c.addOperationListContentAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContentAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContentAssociationsInput struct {

	// The identifier of the content.
	//
	// This member is required.
	ContentId *string

	// The identifier of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListContentAssociationsOutput struct {

	// Summary information about content associations.
	//
	// This member is required.
	ContentAssociationSummaries []types.ContentAssociationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContentAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContentAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContentAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListContentAssociations"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpListContentAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContentAssociations(options.Region), middleware.Before); err != nil {
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

// ListContentAssociationsPaginatorOptions is the paginator options for
// ListContentAssociations
type ListContentAssociationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContentAssociationsPaginator is a paginator for ListContentAssociations
type ListContentAssociationsPaginator struct {
	options   ListContentAssociationsPaginatorOptions
	client    ListContentAssociationsAPIClient
	params    *ListContentAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListContentAssociationsPaginator returns a new
// ListContentAssociationsPaginator
func NewListContentAssociationsPaginator(client ListContentAssociationsAPIClient, params *ListContentAssociationsInput, optFns ...func(*ListContentAssociationsPaginatorOptions)) *ListContentAssociationsPaginator {
	if params == nil {
		params = &ListContentAssociationsInput{}
	}

	options := ListContentAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContentAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContentAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContentAssociations page.
func (p *ListContentAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContentAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListContentAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListContentAssociationsAPIClient is a client that implements the
// ListContentAssociations operation.
type ListContentAssociationsAPIClient interface {
	ListContentAssociations(context.Context, *ListContentAssociationsInput, ...func(*Options)) (*ListContentAssociationsOutput, error)
}

var _ ListContentAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListContentAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListContentAssociations",
	}
}
