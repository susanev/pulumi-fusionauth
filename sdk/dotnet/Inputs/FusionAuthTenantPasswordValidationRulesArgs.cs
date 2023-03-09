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

    public sealed class FusionAuthTenantPasswordValidationRulesArgs : global::Pulumi.ResourceArgs
    {
        [Input("breachDetection")]
        public Input<Inputs.FusionAuthTenantPasswordValidationRulesBreachDetectionArgs>? BreachDetection { get; set; }

        /// <summary>
        /// The maximum length of a password when a new user is created or a user requests a password change.
        /// </summary>
        [Input("maxLength")]
        public Input<int>? MaxLength { get; set; }

        /// <summary>
        /// The minimum length of a password when a new user is created or a user requests a password change.
        /// </summary>
        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        [Input("rememberPreviousPasswords")]
        public Input<Inputs.FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsArgs>? RememberPreviousPasswords { get; set; }

        /// <summary>
        /// Whether to force the user to use at least one non-alphanumeric character.
        /// </summary>
        [Input("requireNonAlpha")]
        public Input<bool>? RequireNonAlpha { get; set; }

        /// <summary>
        /// Whether to force the user to use at least one number.
        /// </summary>
        [Input("requireNumber")]
        public Input<bool>? RequireNumber { get; set; }

        /// <summary>
        /// Whether to force the user to use at least one uppercase and one lowercase character.
        /// </summary>
        [Input("requiredMixedCase")]
        public Input<bool>? RequiredMixedCase { get; set; }

        /// <summary>
        /// When enabled the user’s password will be validated during login. If the password does not meet the currently configured validation rules the user will be required to change their password.
        /// </summary>
        [Input("validateOnLogin")]
        public Input<bool>? ValidateOnLogin { get; set; }

        public FusionAuthTenantPasswordValidationRulesArgs()
        {
        }
        public static new FusionAuthTenantPasswordValidationRulesArgs Empty => new FusionAuthTenantPasswordValidationRulesArgs();
    }
}
