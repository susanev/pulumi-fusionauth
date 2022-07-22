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
    public sealed class FusionAuthSystemConfigurationCorsConfiguration
    {
        /// <summary>
        /// The Access-Control-Allow-Credentials response header values as described by MDN Access-Control-Allow-Credentials.
        /// </summary>
        public readonly bool? AllowCredentials;
        /// <summary>
        /// The Access-Control-Allow-Headers response header values as described by MDN Access-Control-Allow-Headers.
        /// </summary>
        public readonly ImmutableArray<string> AllowedHeaders;
        /// <summary>
        /// The Access-Control-Allow-Methods response header values as described by MDN Access-Control-Allow-Methods.
        /// </summary>
        public readonly ImmutableArray<string> AllowedMethods;
        /// <summary>
        /// The Access-Control-Allow-Origin response header values as described by MDN Access-Control-Allow-Origin. If the wildcard * is specified, no additional domains may be specified.
        /// </summary>
        public readonly ImmutableArray<string> AllowedOrigins;
        /// <summary>
        /// Whether or not FusionAuth should delete the login records based upon this configuration. When true the loginRecordConfiguration.delete.numberOfDaysToRetain will be used to identify login records that are eligible for deletion. When this value is set to false login records will be preserved forever.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The Access-Control-Expose-Headers response header values as described by MDN Access-Control-Expose-Headers.
        /// </summary>
        public readonly ImmutableArray<string> ExposedHeaders;
        /// <summary>
        /// The Access-Control-Max-Age response header values as described by MDN Access-Control-Max-Age.
        /// </summary>
        public readonly int? PreflightMaxAgeInSeconds;

        [OutputConstructor]
        private FusionAuthSystemConfigurationCorsConfiguration(
            bool? allowCredentials,

            ImmutableArray<string> allowedHeaders,

            ImmutableArray<string> allowedMethods,

            ImmutableArray<string> allowedOrigins,

            bool? enabled,

            ImmutableArray<string> exposedHeaders,

            int? preflightMaxAgeInSeconds)
        {
            AllowCredentials = allowCredentials;
            AllowedHeaders = allowedHeaders;
            AllowedMethods = allowedMethods;
            AllowedOrigins = allowedOrigins;
            Enabled = enabled;
            ExposedHeaders = exposedHeaders;
            PreflightMaxAgeInSeconds = preflightMaxAgeInSeconds;
        }
    }
}
