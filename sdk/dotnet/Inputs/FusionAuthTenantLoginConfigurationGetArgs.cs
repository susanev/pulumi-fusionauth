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

    public sealed class FusionAuthTenantLoginConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether to require an API key for the Login API when an `applicationId` is not provided. When an `applicationId` is provided to the Login API call, the application configuration will take precedence. In almost all cases, you will want to this to be `true`.
        /// </summary>
        [Input("requireAuthentication")]
        public Input<bool>? RequireAuthentication { get; set; }

        public FusionAuthTenantLoginConfigurationGetArgs()
        {
        }
    }
}
