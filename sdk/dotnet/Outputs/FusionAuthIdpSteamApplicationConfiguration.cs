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
    public sealed class FusionAuthIdpSteamApplicationConfiguration
    {
        /// <summary>
        /// ID of the Application to apply this configuration to.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        public readonly string? ButtonText;
        /// <summary>
        /// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Determines if a UserRegistration is created for the User automatically or not. If a user doesn’t exist in FusionAuth and logs in through an identity provider, this boolean controls whether or not FusionAuth creates a registration for the User in the Application they are logging into.
        /// </summary>
        public readonly bool? CreateRegistration;
        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The top-level scope that you are requesting from Steam.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
        /// </summary>
        public readonly string? WebApiKey;

        [OutputConstructor]
        private FusionAuthIdpSteamApplicationConfiguration(
            string? applicationId,

            string? buttonText,

            string? clientId,

            bool? createRegistration,

            bool? enabled,

            string? scope,

            string? webApiKey)
        {
            ApplicationId = applicationId;
            ButtonText = buttonText;
            ClientId = clientId;
            CreateRegistration = createRegistration;
            Enabled = enabled;
            Scope = scope;
            WebApiKey = webApiKey;
        }
    }
}
