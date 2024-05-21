# frozen_string_literal: true

require 'sinatra'
require 'json'
require 'optimizely/optimizely_factory'

set :bind, '0.0.0.0'

optimizely_client = Optimizely::OptimizelyFactory.default_instance('RiowyaMnbPxLa4dPWrDqu')

post '/decide' do
  content_type :json

  payload = JSON.parse(request.body.read)
  user_id = payload['userId']
  user_attributes = payload['userAttributes']

  if optimizely_client.is_valid
    user = optimizely_client.create_user_context(user_id, user_attributes)
    decision = user.decide('product_sort')

    if decision.variation_key
      sort_method = decision.variables['sort_method']
      flag_status = decision.enabled ? 'on' : 'off'

      {
        flag_status: flag_status,
        user_id: user.user_id,
        flag_variation: decision.variation_key,
        sort_method: sort_method,
        rule_key: decision.rule_key
      }.to_json
    else
      status 400
      { error: "decision error: #{decision.reasons.join("' '")}" }.to_json
    end
  else
    status 400
    { error: "Optimizely client invalid. Verify in Settings>Environments that you used the primary environment's SDK key" }.to_json
  end
end
