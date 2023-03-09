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

    public sealed class FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TThe length of the secure generator used for generating the the two factor code Id.
        /// </summary>
        [Input("length", required: true)]
        public Input<int> Length { get; set; } = null!;

        /// <summary>
        /// The type of the secure generator used for generating the two factor one time code Id.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs()
        {
        }
        public static new FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs Empty => new FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs();
    }
}
