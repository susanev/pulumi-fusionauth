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

    public sealed class FusionAuthSystemConfigurationAuditLogConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("delete")]
        public Input<Inputs.FusionAuthSystemConfigurationAuditLogConfigurationDeleteArgs>? Delete { get; set; }

        public FusionAuthSystemConfigurationAuditLogConfigurationArgs()
        {
        }
        public static new FusionAuthSystemConfigurationAuditLogConfigurationArgs Empty => new FusionAuthSystemConfigurationAuditLogConfigurationArgs();
    }
}
