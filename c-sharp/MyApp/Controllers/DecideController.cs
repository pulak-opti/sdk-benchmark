using Microsoft.AspNetCore.Mvc;
using OptimizelySDK;
using OptimizelySDK.Entity;
using System;
using System.Collections.Generic;


public class UserPayload
{
    public string? UserId { get; set; }
    public Dictionary<string, string>? UserAttributes { get; set; }
}

[ApiController]
[Route("[controller]")]
public class DecideController : ControllerBase
{
    [HttpPost]
    public IEnumerable<string> Post([FromBody] UserPayload payload)
    {
        var results = new List<string>();

        var optimizelyClient = OptimizelyFactory.NewDefaultInstance("RiowyaMnbPxLa4dPWrDqu");
        if (!optimizelyClient.IsValid)
        {
            results.Add(
                "Optimizely client invalid. " +
                "Verify in Settings>Environments that you used the primary environment's SDK key");
            optimizelyClient.Dispose();
            return results;
        }

        var userAttributes = new UserAttributes();
        if (payload.UserAttributes != null)
        {
            foreach (var attribute in payload.UserAttributes)
            {
                userAttributes.Add(attribute.Key, attribute.Value);
            }
        }

        var user = optimizelyClient.CreateUserContext(payload.UserId, userAttributes);
        var decision = user.Decide("product_sort");

        if (string.IsNullOrEmpty(decision.VariationKey))
        {
            results.Add(Environment.NewLine + Environment.NewLine +
                        "Decision error: " + string.Join(" ", decision.Reasons));
            return results;
        }

        var sortMethod = decision.Variables.ToDictionary()["sort_method"];
        var hasOnFlags = decision.Enabled;

        results.Add(Environment.NewLine +
                    $"Flag {(decision.Enabled ? "on" : "off")}. " +
                    $"User number {user.GetUserId()} saw " +
                    $"flag variation {decision.VariationKey} and got " +
                    $"products sorted by {sortMethod} config variable as part of " +
                    $"flag rule {decision.RuleKey}");

        return results;
    }
}

