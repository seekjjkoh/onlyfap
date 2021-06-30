package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

var modelPath = flag.String("modelPath", "./output/tf-dnn", "Pre-saved model filepath")

const PORT = ":1989"

var model = tg.LoadModel(*modelPath, []string{"serve"}, nil)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ArgMax(preds []float32) int {
	maxSoFar := preds[0]
	maxSoFarIndex := 0
	for i, p := range preds {
		if p > maxSoFar {
			maxSoFar = p
			maxSoFarIndex = i
		}
	}
	return maxSoFarIndex
}

func Predict(w http.ResponseWriter, r *http.Request) {
	inputs := make([]float32, 0)
	json.NewDecoder(r.Body).Decode(&inputs)
	// preprocess
	for i, v := range inputs {
		inputs[i] = v / 256
	}
	tensor, _ := tf.NewTensor([1][]float32{inputs})
	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_inputs", 0): tensor,
	})
	preds := results[0]
	pred := ArgMax(preds.Value().([][]float32)[0])
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"result\":%d}", pred)))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/predict", Predict)
	log.Println("Go server serving at port", PORT)
	http.ListenAndServe(PORT, nil)
}
