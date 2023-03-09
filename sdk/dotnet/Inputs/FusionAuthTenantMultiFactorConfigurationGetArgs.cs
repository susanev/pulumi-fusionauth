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

    public sealed class FusionAuthTenantMultiFactorConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticator")]
        public Input<Inputs.FusionAuthTenantMultiFactorConfigurationAuthenticatorGetArgs>? Authenticator { get; set; }

        [Input("email")]
        public Input<Inputs.FusionAuthTenantMultiFactorConfigurationEmailGetArgs>? Email { get; set; }

        [Input("sms")]
        public Input<Inputs.FusionAuthTenantMultiFactorConfigurationSmsGetArgs>? Sms { get; set; }

        public FusionAuthTenantMultiFactorConfigurationGetArgs()
        {
        }
        public static new FusionAuthTenantMultiFactorConfigurationGetArgs Empty => new FusionAuthTenantMultiFactorConfigurationGetArgs();
    }
}
