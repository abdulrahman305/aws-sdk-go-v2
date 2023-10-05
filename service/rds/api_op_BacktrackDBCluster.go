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
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Backtracks a DB cluster to a specific time, without creating a new DB cluster.
// For more information on backtracking, see Backtracking an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html)
// in the Amazon Aurora User Guide. This action applies only to Aurora MySQL DB
// clusters.
func (c *Client) BacktrackDBCluster(ctx context.Context, params *BacktrackDBClusterInput, optFns ...func(*Options)) (*BacktrackDBClusterOutput, error) {
	if params == nil {
		params = &BacktrackDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BacktrackDBCluster", params, optFns, c.addOperationBacktrackDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BacktrackDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BacktrackDBClusterInput struct {

	// The timestamp of the time to backtrack the DB cluster to, specified in ISO 8601
	// format. For more information about ISO 8601, see the ISO8601 Wikipedia page. (http://en.wikipedia.org/wiki/ISO_8601)
	// If the specified time isn't a consistent time for the DB cluster, Aurora
	// automatically chooses the nearest possible consistent time for the DB cluster.
	// Constraints:
	//   - Must contain a valid ISO 8601 timestamp.
	//   - Can't contain a timestamp set in the future.
	// Example: 2017-07-08T18:00Z
	//
	// This member is required.
	BacktrackTo *time.Time

	// The DB cluster identifier of the DB cluster to be backtracked. This parameter
	// is stored as a lowercase string. Constraints:
	//   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	//   - First character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// Specifies whether to force the DB cluster to backtrack when binary logging is
	// enabled. Otherwise, an error occurs when binary logging is enabled.
	Force *bool

	// Specifies whether to backtrack the DB cluster to the earliest possible
	// backtrack time when BacktrackTo is set to a timestamp earlier than the earliest
	// backtrack time. When this parameter is disabled and BacktrackTo is set to a
	// timestamp earlier than the earliest backtrack time, an error occurs.
	UseEarliestTimeOnPointInTimeUnavailable *bool

	noSmithyDocumentSerde
}

// This data type is used as a response element in the DescribeDBClusterBacktracks
// action.
type BacktrackDBClusterOutput struct {

	// Contains the backtrack identifier.
	BacktrackIdentifier *string

	// The timestamp of the time at which the backtrack was requested.
	BacktrackRequestCreationTime *time.Time

	// The timestamp of the time to which the DB cluster was backtracked.
	BacktrackTo *time.Time

	// The timestamp of the time from which the DB cluster was backtracked.
	BacktrackedFrom *time.Time

	// Contains a user-supplied DB cluster identifier. This identifier is the unique
	// key that identifies a DB cluster.
	DBClusterIdentifier *string

	// The status of the backtrack. This property returns one of the following values:
	//   - applying - The backtrack is currently being applied to or rolled back from
	//   the DB cluster.
	//   - completed - The backtrack has successfully been applied to or rolled back
	//   from the DB cluster.
	//   - failed - An error occurred while the backtrack was applied to or rolled back
	//   from the DB cluster.
	//   - pending - The backtrack is currently pending application to or rollback from
	//   the DB cluster.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBacktrackDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBacktrackDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBacktrackDBCluster{}, middleware.After)
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
	if err = addBacktrackDBClusterResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBacktrackDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBacktrackDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBacktrackDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "BacktrackDBCluster",
	}
}

type opBacktrackDBClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opBacktrackDBClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opBacktrackDBClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addBacktrackDBClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opBacktrackDBClusterResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
