const express = require('express');
const optimizely = require('@optimizely/optimizely-sdk');
const bodyParser = require('body-parser');
const morgan = require('morgan');


const app = express();
app.use(bodyParser.json());
app.use(morgan('combined'));

const optimizelyClient = optimizely.createInstance({
  sdkKey: 'RiowyaMnbPxLa4dPWrDqu',
});

app.post('/decide', (req, res) => {
  const data = req.body;
  const userId = data.userId;
  const userAttributes = data.userAttributes;

  const user = optimizelyClient.createUserContext(userId, userAttributes);
  const decision = user.decide('product_sort');

  if (!decision.variationKey) {
    return res.status(400).json({ error: `decision error ${decision.reasons.join(', ')}` });
  }

  const sortMethod = decision.variables.sort_method;
  const flagStatus = decision.enabled ? 'on' : 'off';

  res.json({
    flag_status: flagStatus,
    user_id: user.userId,
    flag_variation: decision.variationKey,
    sort_method: sortMethod,
    rule_key: decision.ruleKey,
  });
});

app.listen(3000, () => {
  console.log('Server is running on port 3000');
});
