// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Outputs
{

    [OutputType]
    public sealed class FusionAuthTenantJwtConfiguration
    {
        /// <summary>
        /// The unique id of the signing key used to sign the access token.
        /// </summary>
        public readonly string AccessTokenKeyId;
        /// <summary>
        /// The unique id of the signing key used to sign the Id token.
        /// </summary>
        public readonly string IdTokenKeyId;
        /// <summary>
        /// The refresh token expiration policy.
        /// </summary>
        public readonly string? RefreshTokenExpirationPolicy;
        /// <summary>
        /// When enabled, the refresh token will be revoked when a user action, such as locking an account based on a number of failed login attempts, prevents user login.
        /// </summary>
        public readonly bool? RefreshTokenRevocationPolicyOnLoginPrevented;
        /// <summary>
        /// When enabled, the refresh token will be revoked when a user changes their password."
        /// </summary>
        public readonly bool? RefreshTokenRevocationPolicyOnPasswordChange;
        /// <summary>
        /// The length of time in minutes a Refresh Token is valid from the time it was issued. Value must be greater than 0.
        /// </summary>
        public readonly int RefreshTokenTimeToLiveInMinutes;
        /// <summary>
        /// The refresh token usage policy.
        /// </summary>
        public readonly string? RefreshTokenUsagePolicy;
        /// <summary>
        /// The length of time in seconds this JWT is valid from the time it was issued. Value must be greater than 0.
        /// * `login_configuration`
        /// </summary>
        public readonly int TimeToLiveInSeconds;

        [OutputConstructor]
        private FusionAuthTenantJwtConfiguration(
            string accessTokenKeyId,

            string idTokenKeyId,

            string? refreshTokenExpirationPolicy,

            bool? refreshTokenRevocationPolicyOnLoginPrevented,

            bool? refreshTokenRevocationPolicyOnPasswordChange,

            int refreshTokenTimeToLiveInMinutes,

            string? refreshTokenUsagePolicy,

            int timeToLiveInSeconds)
        {
            AccessTokenKeyId = accessTokenKeyId;
            IdTokenKeyId = idTokenKeyId;
            RefreshTokenExpirationPolicy = refreshTokenExpirationPolicy;
            RefreshTokenRevocationPolicyOnLoginPrevented = refreshTokenRevocationPolicyOnLoginPrevented;
            RefreshTokenRevocationPolicyOnPasswordChange = refreshTokenRevocationPolicyOnPasswordChange;
            RefreshTokenTimeToLiveInMinutes = refreshTokenTimeToLiveInMinutes;
            RefreshTokenUsagePolicy = refreshTokenUsagePolicy;
            TimeToLiveInSeconds = timeToLiveInSeconds;
        }
    }
}
