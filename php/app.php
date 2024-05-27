<?php
require 'vendor/autoload.php';

use Optimizely\OptimizelyFactory;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

$app = AppFactory::create();

$app->addBodyParsingMiddleware();


$app->post('/decide', function (Request $request, Response $response) {
    $optimizelyToken = 'RiowyaMnbPxLa4dPWrDqu';

    $optimizelyClient = OptimizelyFactory::createDefaultInstance($optimizelyToken);

    $data = $request->getParsedBody();
    
    if (!isset($data['userId']) || !isset($data['userAttributes'])) {
        $response->getBody()->write("Missing 'userId' or 'userAttributes' in request body");
        return $response->withStatus(400);
    }

    $userId = $data['userId'];
    $userAttributes = $data['userAttributes'];

    if ($optimizelyClient->isValid()) {
        $user = $optimizelyClient->createUserContext($userId, $userAttributes);
        $decision = $user->decide('product_sort');

        if (empty($decision->getVariationKey())) {
            $response->getBody()->write(sprintf("decision error:  %s", implode(', ', $decision->getReasons())));
        }

        $sortMethod = $decision->getVariables() ['sort_method'];
        $enabled = $decision->getEnabled();

        $response->getBody()->write(sprintf("\n\nFlag %s. User %s saw flag variation %s and got products sorted by %s config variable as part of flag rule: %s", $enabled ? "on" : "off", $user->getUserId(), $decision->getVariationKey(), $sortMethod, $decision->getRuleKey()));

        if (!$enabled) {
            $response->getBody()->write(sprintf("\n\n Flag was off. Some reasons could include: 
                \n Check your SDK key. Verify in Settings>Environments that you used the right key for the environment where your flag is toggled to ON. 
                \ncheck your key at  https://app.optimizely.com/v2/projects/%s settings/implementation", $optimizelyClient
                ->configManager
                ->getConfig()
                ->getProjectId()));
        }
    } else {
        $response->getBody()->write("Optimizely client invalid. Verify in Settings>Environments that you used the primary environment's SDK key");
    }

    return $response;
});

$app->get('/health', function (Request $request, Response $response) {
    $response->getBody()->write("Server is running");
    return $response->withStatus(200);
});


$app->run();
