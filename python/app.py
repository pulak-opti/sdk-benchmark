from flask import Flask, request, jsonify
from optimizely import optimizely

app = Flask(__name__)

optimizely_client = optimizely.Optimizely(sdk_key="RiowyaMnbPxLa4dPWrDqu")

@app.route('/decide', methods=['POST'])
def decide():
    if not optimizely_client.config_manager.get_config():
        raise Exception("Optimizely client invalid. Verify in Settings>Environments that "
                        "you used the primary environment's SDK key")

    data = request.get_json()
    user_id = data.get('userId')
    user_attributes = data.get('userAttributes')

    user = optimizely_client.create_user_context(user_id, user_attributes)
    decision = user.decide("product_sort")

    if not decision.variation_key:
        return jsonify(error=f"decision error {', '.join(decision.reasons)}"), 400

    sort_method = decision.variables["sort_method"]

    if decision.enabled:
        flag_status = 'on'
    else:
        flag_status = 'off'

    return jsonify(
        flag_status=flag_status,
        user_id=user.user_id,
        flag_variation=decision.variation_key,
        sort_method=sort_method,
        rule_key=decision.rule_key
    )

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
