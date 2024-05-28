package com.example;

import com.google.gson.Gson;
import com.optimizely.ab.OptimizelyFactory;
import com.optimizely.ab.config.parser.JsonParseException;
import com.optimizely.ab.Optimizely;
import com.optimizely.ab.OptimizelyUserContext;
import com.optimizely.ab.optimizelydecision.OptimizelyDecision;
import spark.Request;
import spark.Response;

import java.util.Map;
import java.util.Random;

import static spark.Spark.*;

public class Main {
    public static void main(String[] args) {
        port(8000);

        post("/decide", (req, res) -> {
            Gson gson = new Gson();
            Payload payload = gson.fromJson(req.body(), Payload.class);
            String userId = payload.userId;
            Map<String, String> userAttributes = payload.userAttributes;

            Optimizely optimizely = OptimizelyFactory.newDefaultInstance("RiowyaMnbPxLa4dPWrDqu");
            OptimizelyUserContext user = optimizely.createUserContext(userId, userAttributes);
            OptimizelyDecision decision = user.decide("product_sort");

            if (decision.getVariationKey() == null) {
                res.status(400);
                return gson.toJson(new Error("decision error: " + decision.getReasons()));
            }

            String sortMethod = decision.getVariables().getValue("sort_method", String.class);
            String flagStatus = decision.getEnabled() ? "on" : "off";


            return gson.toJson(new Decision(flagStatus, userId, decision.getVariationKey(), sortMethod, decision.getRuleKey()));
        });
    }

    static class Payload {
        String userId;
        Map<String, String> userAttributes;
    }

    static class Error {
        String error;

        Error(String error) {
            this.error = error;
        }
    }

    static class Decision {
        String flag_status;
        String user_id;
        String flag_variation;
        String sort_method;
        String rule_key;

        Decision(String flag_status, String user_id, String flag_variation, String sort_method, String rule_key) {
            this.flag_status = flag_status;
            this.user_id = user_id;
            this.flag_variation = flag_variation;
            this.sort_method = sort_method;
            this.rule_key = rule_key;
        }
    }
}
