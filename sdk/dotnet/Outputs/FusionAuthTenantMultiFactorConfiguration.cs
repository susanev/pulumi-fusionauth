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
    public sealed class FusionAuthTenantMultiFactorConfiguration
    {
        public readonly Outputs.FusionAuthTenantMultiFactorConfigurationAuthenticator? Authenticator;
        public readonly Outputs.FusionAuthTenantMultiFactorConfigurationEmail? Email;
        /// <summary>
        /// When set to `Enabled` and a user has one or more two-factor methods configured, the user will be required to complete a two-factor challenge during login. When set to `Disabled`, even when a user has configured one or more two-factor methods, the user will not be required to complete a two-factor challenge during login.
        /// </summary>
        public readonly string? LoginPolicy;
        public readonly Outputs.FusionAuthTenantMultiFactorConfigurationSms? Sms;

        [OutputConstructor]
        private FusionAuthTenantMultiFactorConfiguration(
            Outputs.FusionAuthTenantMultiFactorConfigurationAuthenticator? authenticator,

            Outputs.FusionAuthTenantMultiFactorConfigurationEmail? email,

            string? loginPolicy,

            Outputs.FusionAuthTenantMultiFactorConfigurationSms? sms)
        {
            Authenticator = authenticator;
            Email = email;
            LoginPolicy = loginPolicy;
            Sms = sms;
        }
    }
}
