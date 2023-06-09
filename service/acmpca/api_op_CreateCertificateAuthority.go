// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a root or subordinate private certificate authority (CA). You must
// specify the CA configuration, an optional configuration for Online Certificate
// Status Protocol (OCSP) and/or a certificate revocation list (CRL), the CA type,
// and an optional idempotency token to avoid accidental creation of multiple CAs.
// The CA configuration specifies the name of the algorithm and key size to be used
// to create the CA private key, the type of signing algorithm that the CA uses,
// and X.500 subject information. The OCSP configuration can optionally specify a
// custom URL for the OCSP responder. The CRL configuration specifies the CRL
// expiration period in days (the validity period of the CRL), the Amazon S3 bucket
// that will contain the CRL, and a CNAME alias for the S3 bucket that is included
// in certificates issued by the CA. If successful, this action returns the Amazon
// Resource Name (ARN) of the CA. Both Amazon Web Services Private CA and the IAM
// principal must have permission to write to the S3 bucket that you specify. If
// the IAM principal making the call does not have permission to write to the
// bucket, then an exception is thrown. For more information, see Access policies
// for CRLs in Amazon S3 (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies)
// . Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see Encrypting Your CRLs (https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#crl-encryption)
// .
func (c *Client) CreateCertificateAuthority(ctx context.Context, params *CreateCertificateAuthorityInput, optFns ...func(*Options)) (*CreateCertificateAuthorityOutput, error) {
	if params == nil {
		params = &CreateCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCertificateAuthority", params, optFns, c.addOperationCreateCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateAuthorityInput struct {

	// Name and bit size of the private key algorithm, the name of the signing
	// algorithm, and X.500 certificate subject information.
	//
	// This member is required.
	CertificateAuthorityConfiguration *types.CertificateAuthorityConfiguration

	// The type of the certificate authority.
	//
	// This member is required.
	CertificateAuthorityType types.CertificateAuthorityType

	// Custom string that can be used to distinguish between calls to the
	// CreateCertificateAuthority action. Idempotency tokens for
	// CreateCertificateAuthority time out after five minutes. Therefore, if you call
	// CreateCertificateAuthority multiple times with the same idempotency token within
	// five minutes, Amazon Web Services Private CA recognizes that you are requesting
	// only certificate authority and will issue only one. If you change the
	// idempotency token for each call, Amazon Web Services Private CA recognizes that
	// you are requesting multiple certificate authorities.
	IdempotencyToken *string

	// Specifies a cryptographic key management compliance standard used for handling
	// CA keys. Default: FIPS_140_2_LEVEL_3_OR_HIGHER Some Amazon Web Services Regions
	// do not support the default. When creating a CA in these Regions, you must
	// provide FIPS_140_2_LEVEL_2_OR_HIGHER as the argument for
	// KeyStorageSecurityStandard . Failure to do this results in an
	// InvalidArgsException with the message, "A certificate authority cannot be
	// created in this region with the specified security standard." For information
	// about security standard support in various Regions, see Storage and security
	// compliance of Amazon Web Services Private CA private keys (https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys)
	// .
	KeyStorageSecurityStandard types.KeyStorageSecurityStandard

	// Contains information to enable Online Certificate Status Protocol (OCSP)
	// support, to enable a certificate revocation list (CRL), to enable both, or to
	// enable neither. The default is for both certificate validation mechanisms to be
	// disabled. The following requirements apply to revocation configurations.
	//   - A configuration disabling CRLs or OCSP must contain only the Enabled=False
	//   parameter, and will fail if other parameters such as CustomCname or
	//   ExpirationInDays are included.
	//   - In a CRL configuration, the S3BucketName parameter must conform to Amazon
	//   S3 bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)
	//   .
	//   - A configuration containing a custom Canonical Name (CNAME) parameter for
	//   CRLs or OCSP must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)
	//   restrictions on the use of special characters in a CNAME.
	//   - In a CRL or OCSP configuration, the value of a CNAME parameter must not
	//   include a protocol prefix such as "http://" or "https://".
	// For more information, see the OcspConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_OcspConfiguration.html)
	// and CrlConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CrlConfiguration.html)
	// types.
	RevocationConfiguration *types.RevocationConfiguration

	// Key-value pairs that will be attached to the new private CA. You can associate
	// up to 50 tags with a private CA. For information using tags with IAM to manage
	// permissions, see Controlling Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html)
	// .
	Tags []types.Tag

	// Specifies whether the CA issues general-purpose certificates that typically
	// require a revocation mechanism, or short-lived certificates that may optionally
	// omit revocation because they expire quickly. Short-lived certificate validity is
	// limited to seven days. The default value is GENERAL_PURPOSE.
	UsageMode types.CertificateAuthorityUsageMode

	noSmithyDocumentSerde
}

type CreateCertificateAuthorityOutput struct {

	// If successful, the Amazon Resource Name (ARN) of the certificate authority
	// (CA). This is of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	CertificateAuthorityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificateAuthority{}, middleware.After)
	if err != nil {
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificateAuthority(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "CreateCertificateAuthority",
	}
}
