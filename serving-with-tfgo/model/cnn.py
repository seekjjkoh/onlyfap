import pandas as pd
import numpy as np
import tensorflow as tf


def create_cnn(x, y):
	model = tf.keras.Sequential()
	model.add(tf.keras.layers.InputLayer(input_shape=(784,), name="inputs"))
	model.add(tf.keras.layers.Reshape((28,28,1)))
	model.add(tf.keras.layers.Conv2D(32, 3, padding="valid", activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.MaxPooling2D())
	model.add(tf.keras.layers.BatchNormalization())
	model.add(tf.keras.layers.Dropout(0.5))
	model.add(tf.keras.layers.Conv2D(32, 5, padding="same", activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.BatchNormalization())
	model.add(tf.keras.layers.Dropout(0.5))
	model.add(tf.keras.layers.Flatten())
	model.add(tf.keras.layers.Dense(10, activation=tf.keras.activations.relu))
	model.add(tf.keras.layers.Softmax())
	model.compile(optimizer="adam", loss="sparse_categorical_crossentropy", metrics=["acc"])
	print(model.summary())
	model.fit(x, y, batch_size=32, epochs=8, shuffle=True)
	return model


def export_model(model, filename):
	tf.keras.models.save_model(model, filename)

if __name__ == "__main__":
	print("running with tf", tf.__version__)
	train_df = pd.read_csv("./data/train.csv")
	y = train_df.label.values
	X = train_df.drop("label", axis=1).values/256
	model = create_cnn(X, y)
	export_model(model, "output/tf-cnn")
	train_preds = np.argmax(model.predict(X), axis=1)
	train_acc = np.sum(train_preds == y)/len(y)
	print("train_accuracy", train_acc)
