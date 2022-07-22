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
    public sealed class FusionAuthIdpSamlV2IdpInitiatedTenantConfiguration
    {
        public readonly bool? LimitUserLinkCountEnabled;
        public readonly int? LimitUserLinkCountMaximumLinks;
        public readonly string? TenantId;

        [OutputConstructor]
        private FusionAuthIdpSamlV2IdpInitiatedTenantConfiguration(
            bool? limitUserLinkCountEnabled,

            int? limitUserLinkCountMaximumLinks,

            string? tenantId)
        {
            LimitUserLinkCountEnabled = limitUserLinkCountEnabled;
            LimitUserLinkCountMaximumLinks = limitUserLinkCountMaximumLinks;
            TenantId = tenantId;
        }
    }
}
