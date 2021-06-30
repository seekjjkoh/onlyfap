import tensorflow as tf
import pandas as pd
import numpy as np

if __name__ == "__main__":
	test_df = pd.read_csv("./data/test.csv")
	submission_df = pd.read_csv("./data/sample_submission.csv")
	X = test_df.values/256
	model = tf.keras.models.load_model("./output/tf-dnn")
	preds = model.predict(X)
	submission_df["Label"] = np.argmax(preds, axis=1)
	submission_df.to_csv("./data/py_dnn_submission.csv", index=False)
