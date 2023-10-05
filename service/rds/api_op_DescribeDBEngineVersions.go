// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the properties of specific versions of DB engines.
func (c *Client) DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBEngineVersions", params, optFns, c.addOperationDescribeDBEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBEngineVersionsInput struct {

	// The name of a specific DB parameter group family to return details for.
	// Constraints:
	//   - If supplied, must match an existing DB parameter group family.
	DBParameterGroupFamily *string

	// Specifies whether to return only the default version of the specified engine or
	// the engine and major version combination.
	DefaultOnly bool

	// The database engine to return version details for. Valid Values:
	//   - aurora-mysql
	//   - aurora-postgresql
	//   - custom-oracle-ee
	//   - mariadb
	//   - mysql
	//   - oracle-ee
	//   - oracle-ee-cdb
	//   - oracle-se2
	//   - oracle-se2-cdb
	//   - postgres
	//   - sqlserver-ee
	//   - sqlserver-se
	//   - sqlserver-ex
	//   - sqlserver-web
	Engine *string

	// A specific database engine version to return details for. Example: 5.1.49
	EngineVersion *string

	// A filter that specifies one or more DB engine versions to describe. Supported
	// filters:
	//   - db-parameter-group-family - Accepts parameter groups family names. The
	//   results list only includes information about the DB engine versions for these
	//   parameter group families.
	//   - engine - Accepts engine names. The results list only includes information
	//   about the DB engine versions for these engines.
	//   - engine-mode - Accepts DB engine modes. The results list only includes
	//   information about the DB engine versions for these engine modes. Valid DB engine
	//   modes are the following:
	//   - global
	//   - multimaster
	//   - parallelquery
	//   - provisioned
	//   - serverless
	//   - engine-version - Accepts engine versions. The results list only includes
	//   information about the DB engine versions for these engine versions.
	//   - status - Accepts engine version statuses. The results list only includes
	//   information about the DB engine versions for these statuses. Valid statuses are
	//   the following:
	//   - available
	//   - deprecated
	Filters []types.Filter

	// Specifies whether to also list the engine versions that aren't available. The
	// default is to list only available engine versions.
	IncludeAll *bool

	// Specifies whether to list the supported character sets for each engine version.
	// If this parameter is enabled and the requested engine supports the
	// CharacterSetName parameter for CreateDBInstance , the response includes a list
	// of supported character sets for each engine version. For RDS Custom, the default
	// is not to list supported character sets. If you enable this parameter, RDS
	// Custom returns no results.
	ListSupportedCharacterSets *bool

	// Specifies whether to list the supported time zones for each engine version. If
	// this parameter is enabled and the requested engine supports the TimeZone
	// parameter for CreateDBInstance , the response includes a list of supported time
	// zones for each engine version. For RDS Custom, the default is not to list
	// supported time zones. If you enable this parameter, RDS Custom returns no
	// results.
	ListSupportedTimezones *bool

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the DescribeDBEngineVersions
// action.
type DescribeDBEngineVersionsOutput struct {

	// A list of DBEngineVersion elements.
	DBEngineVersions []types.DBEngineVersion

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addDescribeDBEngineVersionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeDBEngineVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBEngineVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeDBEngineVersionsAPIClient is a client that implements the
// DescribeDBEngineVersions operation.
type DescribeDBEngineVersionsAPIClient interface {
	DescribeDBEngineVersions(context.Context, *DescribeDBEngineVersionsInput, ...func(*Options)) (*DescribeDBEngineVersionsOutput, error)
}

var _ DescribeDBEngineVersionsAPIClient = (*Client)(nil)

// DescribeDBEngineVersionsPaginatorOptions is the paginator options for
// DescribeDBEngineVersions
type DescribeDBEngineVersionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBEngineVersionsPaginator is a paginator for DescribeDBEngineVersions
type DescribeDBEngineVersionsPaginator struct {
	options   DescribeDBEngineVersionsPaginatorOptions
	client    DescribeDBEngineVersionsAPIClient
	params    *DescribeDBEngineVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBEngineVersionsPaginator returns a new
// DescribeDBEngineVersionsPaginator
func NewDescribeDBEngineVersionsPaginator(client DescribeDBEngineVersionsAPIClient, params *DescribeDBEngineVersionsInput, optFns ...func(*DescribeDBEngineVersionsPaginatorOptions)) *DescribeDBEngineVersionsPaginator {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	options := DescribeDBEngineVersionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBEngineVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBEngineVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBEngineVersions page.
func (p *DescribeDBEngineVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBEngineVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBEngineVersions",
	}
}

type opDescribeDBEngineVersionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeDBEngineVersionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeDBEngineVersionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "rds"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "rds"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("rds")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeDBEngineVersionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeDBEngineVersionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
