python-model-train-dnn:
	python model/dnn.py

python-model-train-cnn:
	python model/cnn.py

python-offline-predict-dnn:
	python offline-prediction/python/dnn.py

python-offline-predict-cnn:
	python offline-prediction/python/cnn.py

go-offline-predict-dnn:
	go run offline-prediction/tfgo/main.go

go-offline-predict-cnn:
	go run offline-prediction/tfgo/main.go --modelPath=./output/tf-cnn --output=./data/go_cnn_submission.csv

python-run-server:
	python serving/flask/app.py

go-run-server:
	go run serving/go/main.go

test-curl:
	sh scripts/test_curl.sh
