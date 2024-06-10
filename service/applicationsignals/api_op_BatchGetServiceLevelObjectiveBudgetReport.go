// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationsignals

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this operation to retrieve one or more service level objective (SLO) budget
// reports.
//
// An error budget is the amount of time in unhealthy periods that your service
// can accumulate during an interval before your overall SLO budget health is
// breached and the SLO is considered to be unmet. For example, an SLO with a
// threshold of 99.95% and a monthly interval translates to an error budget of 21.9
// minutes of downtime in a 30-day month.
//
// Budget reports include a health indicator, the attainment value, and remaining
// budget.
//
// For more information about SLO error budgets, see [SLO concepts].
//
// [SLO concepts]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html#CloudWatch-ServiceLevelObjectives-concepts
func (c *Client) BatchGetServiceLevelObjectiveBudgetReport(ctx context.Context, params *BatchGetServiceLevelObjectiveBudgetReportInput, optFns ...func(*Options)) (*BatchGetServiceLevelObjectiveBudgetReportOutput, error) {
	if params == nil {
		params = &BatchGetServiceLevelObjectiveBudgetReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetServiceLevelObjectiveBudgetReport", params, optFns, c.addOperationBatchGetServiceLevelObjectiveBudgetReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetServiceLevelObjectiveBudgetReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetServiceLevelObjectiveBudgetReportInput struct {

	// An array containing the IDs of the service level objectives that you want to
	// include in the report.
	//
	// This member is required.
	SloIds []string

	// The date and time that you want the report to be for. It is expressed as the
	// number of milliseconds since Jan 1, 1970 00:00:00 UTC.
	//
	// This member is required.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

type BatchGetServiceLevelObjectiveBudgetReportOutput struct {

	// An array of structures, where each structure includes an error indicating that
	// one of the requests in the array was not valid.
	//
	// This member is required.
	Errors []types.ServiceLevelObjectiveBudgetReportError

	// An array of structures, where each structure is one budget report.
	//
	// This member is required.
	Reports []types.ServiceLevelObjectiveBudgetReport

	// The date and time that the report is for. It is expressed as the number of
	// milliseconds since Jan 1, 1970 00:00:00 UTC.
	//
	// This member is required.
	Timestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetServiceLevelObjectiveBudgetReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetServiceLevelObjectiveBudgetReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetServiceLevelObjectiveBudgetReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetServiceLevelObjectiveBudgetReport"); err != nil {
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
	if err = addOpBatchGetServiceLevelObjectiveBudgetReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetServiceLevelObjectiveBudgetReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetServiceLevelObjectiveBudgetReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetServiceLevelObjectiveBudgetReport",
	}
}
