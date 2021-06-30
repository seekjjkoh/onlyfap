from flask import Flask, request, jsonify
import numpy as np
import tensorflow as tf

# change this for different architecture
model_path="./output/tf-dnn"
PORT=1989
model=tf.keras.models.load_model(model_path)

app=Flask(__name__)

@app.route("/")
def index():
	return "Hello world"

@app.route("/predict", methods=["POST"])
def predict():
	image = np.array([np.array(request.get_json())/256])
	preds = model.predict(image)
	pred = np.argmax(preds[0])
	return jsonify({"result": int(pred)})


if __name__ == "__main__":
    print("Python Flask server serving at port:", PORT)
    app.run(port=PORT)
