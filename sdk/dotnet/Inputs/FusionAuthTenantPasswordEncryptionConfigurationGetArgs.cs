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

    public sealed class FusionAuthTenantPasswordEncryptionConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default method for encrypting the User’s password.
        /// </summary>
        [Input("encryptionScheme")]
        public Input<string>? EncryptionScheme { get; set; }

        /// <summary>
        /// The factor used by the password encryption scheme. If not provided, the PasswordEncryptor provides a default value. Generally this will be used as an iteration count to generate the hash. The actual use of this value is up to the PasswordEncryptor implementation.
        /// </summary>
        [Input("encryptionSchemeFactor")]
        public Input<int>? EncryptionSchemeFactor { get; set; }

        /// <summary>
        /// When enabled a user’s hash configuration will be modified to match these configured settings. This can be useful to increase a password hash strength over time or upgrade imported users to a more secure encryption scheme after an initial import.
        /// </summary>
        [Input("modifyEncryptionSchemeOnLogin")]
        public Input<bool>? ModifyEncryptionSchemeOnLogin { get; set; }

        public FusionAuthTenantPasswordEncryptionConfigurationGetArgs()
        {
        }
    }
}
