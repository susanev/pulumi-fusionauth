// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Inputs
{

    public sealed class FusionAuthTenantExternalIdentifierConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time in seconds until a OAuth authorization code in no longer valid to be exchanged for an access token. This is essentially the time allowed between the start of an Authorization request during the Authorization code grant and when you request an access token using this authorization code on the Token endpoint.
        /// </summary>
        [Input("authorizationGrantIdTimeToLiveInSeconds", required: true)]
        public Input<int> AuthorizationGrantIdTimeToLiveInSeconds { get; set; } = null!;

        [Input("changePasswordIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationChangePasswordIdGeneratorArgs> ChangePasswordIdGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a change password Id is no longer valid and cannot be used by the Change Password API. Value must be greater than 0.
        /// </summary>
        [Input("changePasswordIdTimeToLiveInSeconds", required: true)]
        public Input<int> ChangePasswordIdTimeToLiveInSeconds { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a device code Id is no longer valid and cannot be used by the Token API. Value must be greater than 0.
        /// </summary>
        [Input("deviceCodeTimeToLiveInSeconds", required: true)]
        public Input<int> DeviceCodeTimeToLiveInSeconds { get; set; } = null!;

        [Input("deviceUserCodeIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs> DeviceUserCodeIdGenerator { get; set; } = null!;

        [Input("emailVerificationIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationIdGeneratorArgs> EmailVerificationIdGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a email verification Id is no longer valid and cannot be used by the Verify Email API. Value must be greater than 0.
        /// </summary>
        [Input("emailVerificationIdTimeToLiveInSeconds", required: true)]
        public Input<int> EmailVerificationIdTimeToLiveInSeconds { get; set; } = null!;

        [Input("emailVerificationOneTimeCodeGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationOneTimeCodeGeneratorArgs> EmailVerificationOneTimeCodeGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until an external authentication Id is no longer valid and cannot be used by the Token API. Value must be greater than 0.
        /// </summary>
        [Input("externalAuthenticationIdTimeToLiveInSeconds", required: true)]
        public Input<int> ExternalAuthenticationIdTimeToLiveInSeconds { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a One Time Password is no longer valid and cannot be used by the Login API. Value must be greater than 0.
        /// </summary>
        [Input("oneTimePasswordTimeToLiveInSeconds", required: true)]
        public Input<int> OneTimePasswordTimeToLiveInSeconds { get; set; } = null!;

        [Input("passwordlessLoginGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationPasswordlessLoginGeneratorArgs> PasswordlessLoginGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a passwordless code is no longer valid and cannot be used by the Passwordless API. Value must be greater than 0.
        /// </summary>
        [Input("passwordlessLoginTimeToLiveInSeconds", required: true)]
        public Input<int> PasswordlessLoginTimeToLiveInSeconds { get; set; } = null!;

        /// <summary>
        /// The number of seconds before the pending account link identifier is no longer valid to complete an account link request. Value must be greater than 0.
        /// </summary>
        [Input("pendingAccountLinkTimeToLiveInSeconds")]
        public Input<int>? PendingAccountLinkTimeToLiveInSeconds { get; set; }

        [Input("registrationVerificationIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationIdGeneratorArgs> RegistrationVerificationIdGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a registration verification Id is no longer valid and cannot be used by the Verify Registration API. Value must be greater than 0.
        /// </summary>
        [Input("registrationVerificationIdTimeToLiveInSeconds", required: true)]
        public Input<int> RegistrationVerificationIdTimeToLiveInSeconds { get; set; } = null!;

        [Input("registrationVerificationOneTimeCodeGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationOneTimeCodeGeneratorArgs> RegistrationVerificationOneTimeCodeGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds that a SAML AuthN request will be eligible for use to authenticate with FusionAuth.
        /// </summary>
        [Input("samlV2AuthnRequestIdTtlSeconds")]
        public Input<int>? SamlV2AuthnRequestIdTtlSeconds { get; set; }

        [Input("setupPasswordIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs> SetupPasswordIdGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until a setup password Id is no longer valid and cannot be used by the Change Password API. Value must be greater than 0.
        /// </summary>
        [Input("setupPasswordIdTimeToLiveInSeconds", required: true)]
        public Input<int> SetupPasswordIdTimeToLiveInSeconds { get; set; } = null!;

        /// <summary>
        /// The number of seconds before the Trust Token is no longer valid to complete a request that requires trust. Value must be greater than 0.
        /// </summary>
        [Input("trustTokenTimeToLiveInSeconds")]
        public Input<int>? TrustTokenTimeToLiveInSeconds { get; set; }

        /// <summary>
        /// The time in seconds until a two factor Id is no longer valid and cannot be used by the Two Factor Login API. Value must be greater than 0.
        /// </summary>
        [Input("twoFactorIdTimeToLiveInSeconds", required: true)]
        public Input<int> TwoFactorIdTimeToLiveInSeconds { get; set; } = null!;

        [Input("twoFactorOneTimeCodeIdGenerator", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationTwoFactorOneTimeCodeIdGeneratorArgs> TwoFactorOneTimeCodeIdGenerator { get; set; } = null!;

        /// <summary>
        /// The time in seconds until an issued Two Factor trust Id is no longer valid and the User will be required to complete Two Factor authentication during the next authentication attempt. Value must be greater than 0.
        /// </summary>
        [Input("twoFactorTrustIdTimeToLiveInSeconds", required: true)]
        public Input<int> TwoFactorTrustIdTimeToLiveInSeconds { get; set; } = null!;

        public FusionAuthTenantExternalIdentifierConfigurationArgs()
        {
        }
    }
}
