// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of DKIM tokens for a domain identity.
//
// When you execute the VerifyDomainDkim operation, the domain that you specify is
// added to the list of identities that are associated with your account. This is
// true even if you haven't already associated the domain with your account by
// using the VerifyDomainIdentity operation. However, you can't send email from
// the domain until you either successfully [verify it]or you successfully [set up DKIM for it].
//
// You use the tokens that are generated by this operation to create CNAME
// records. When Amazon SES detects that you've added these records to the DNS
// configuration for a domain, you can start sending email from that domain. You
// can start sending email even if you haven't added the TXT record provided by the
// VerifyDomainIdentity operation to the DNS configuration for your domain. All
// email that you send from the domain is authenticated using DKIM.
//
// To create the CNAME records for DKIM authentication, use the following values:
//
//   - Name: token._domainkey.example.com
//
//   - Type: CNAME
//
//   - Value: token.dkim.amazonses.com
//
// In the preceding example, replace token with one of the tokens that are
// generated when you execute this operation. Replace example.com with your domain.
// Repeat this process for each token that's generated by this operation.
//
// You can execute this operation no more than once per second.
//
// [verify it]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#verify-domain-procedure
// [set up DKIM for it]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
func (c *Client) VerifyDomainDkim(ctx context.Context, params *VerifyDomainDkimInput, optFns ...func(*Options)) (*VerifyDomainDkimOutput, error) {
	if params == nil {
		params = &VerifyDomainDkimInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyDomainDkim", params, optFns, c.addOperationVerifyDomainDkimMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyDomainDkimOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to generate the CNAME records needed to set up Easy DKIM
// with Amazon SES. For more information about setting up Easy DKIM, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
type VerifyDomainDkimInput struct {

	// The name of the domain to be verified for Easy DKIM signing.
	//
	// This member is required.
	Domain *string

	noSmithyDocumentSerde
}

// Returns CNAME records that you must publish to the DNS server of your domain to
// set up Easy DKIM with Amazon SES.
type VerifyDomainDkimOutput struct {

	// A set of character strings that represent the domain's identity. If the
	// identity is an email address, the tokens represent the domain of that address.
	//
	// Using these tokens, you need to create DNS CNAME records that point to DKIM
	// public keys that are hosted by Amazon SES. Amazon Web Services eventually
	// detects that you've updated your DNS records. This detection process might take
	// up to 72 hours. After successful detection, Amazon SES is able to DKIM-sign
	// email originating from that domain. (This only applies to domain identities, not
	// email address identities.)
	//
	// For more information about creating DNS records using DKIM tokens, see the [Amazon SES Developer Guide].
	//
	// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
	//
	// This member is required.
	DkimTokens []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyDomainDkimMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpVerifyDomainDkim{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpVerifyDomainDkim{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyDomainDkim"); err != nil {
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
	if err = addOpVerifyDomainDkimValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyDomainDkim(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyDomainDkim(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyDomainDkim",
	}
}