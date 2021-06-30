from os import name
import pandas as pd
import numpy as np
import tensorflow as tf


def create_dnn(x, y):
	model = tf.keras.Sequential()
	model.add(tf.keras.layers.InputLayer(input_shape=(784,), name="inputs"))
	model.add(tf.keras.layers.Dense(512, activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.Dropout(0.5))
	model.add(tf.keras.layers.Dense(128, activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.Dropout(0.5))
	model.add(tf.keras.layers.Dense(10, activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.Softmax())
	print(model.summary())
	model.compile(optimizer="adam", loss="sparse_categorical_crossentropy", metrics=["acc"])
	model.fit(x, y, batch_size=32, epochs=8, shuffle=True)
	return model


def export_model(model, filename):
	tf.keras.models.save_model(model, filename)

if __name__ == "__main__":
	print("running with tf", tf.__version__)
	train_df = pd.read_csv("./data/train.csv")
	y = train_df.label.values
	X = train_df.drop("label", axis=1).values/256
	model = create_dnn(X, y)
	export_model(model, "output/tf-dnn")
	train_preds = np.argmax(model.predict(X), axis=1)
	train_acc = np.sum(train_preds == y)/len(y)
	print("train_accuracy", train_acc)
