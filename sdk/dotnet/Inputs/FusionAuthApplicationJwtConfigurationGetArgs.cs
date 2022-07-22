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

    public sealed class FusionAuthApplicationJwtConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the signing key used to sign the access token.
        /// </summary>
        [Input("accessTokenId")]
        public Input<string>? AccessTokenId { get; set; }

        /// <summary>
        /// Whether or not SAML Single Logout for this SAML IdP is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Id of the signing key used to sign the Id token.
        /// </summary>
        [Input("idTokenKeyId")]
        public Input<string>? IdTokenKeyId { get; set; }

        /// <summary>
        /// The length of time in minutes the JWT refresh token will live before it is expired and is not able to be exchanged for a JWT.
        /// </summary>
        [Input("refreshTokenTtlMinutes")]
        public Input<int>? RefreshTokenTtlMinutes { get; set; }

        /// <summary>
        /// The length of time in seconds the JWT will live before it is expired and no longer valid.
        /// </summary>
        [Input("ttlSeconds")]
        public Input<int>? TtlSeconds { get; set; }

        public FusionAuthApplicationJwtConfigurationGetArgs()
        {
        }
    }
}
