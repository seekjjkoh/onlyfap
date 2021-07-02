package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
)

var (
	modelPath = flag.String("modelPath", "./output/tf-dnn", "Pre-saved model filepath")
	output    = flag.String("output", "./data/go_dnn_submission.csv", "Output submission csv file")
)

const testFile = "./data/test.csv"
const submissionFile = "./data/sample_submission.csv"

func ReadInput(inputFile string) *tf.Tensor {
	f, _ := os.Open(inputFile)
	defer f.Close()
	r := csv.NewReader(f)
	r.Read() // skip header
	inputs := make([][784]float32, 0)

	i := 0
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		inputs = append(inputs, [784]float32{})
		for pixelIndex, v := range rec {
			floatV, _ := strconv.ParseFloat(v, 32)
			inputs[i][pixelIndex] = float32(floatV) / 256
		}
		i++
	}

	tensors, err := tf.NewTensor(inputs)
	if err != nil {
		panic(err)
	}
	return tensors
}

func ReadSubmission(inputFile string) [][]string {
	f, _ := os.Open(inputFile)
	defer f.Close()

	r := csv.NewReader(f)
	submission, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return submission
}

func WriteCSV(filename string, rows [][]string) {
	f, _ := os.Create(filename)
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(rows)
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

func DoNothing(something ...interface{}) {}

func main() {
	flag.Parse()
	testTensors := ReadInput(testFile)
	submissions := ReadSubmission(submissionFile)
	model := tg.LoadModel(*modelPath, []string{"serve"}, nil)
	DoNothing(testTensors, submissions, model) // make debugging easier

	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_inputs", 0): testTensors,
	})

	fmt.Println("output shape:", results[0].Shape())
	preds := results[0]
	predsInFloats := preds.Value().([][]float32)
	for i := 1; i < len(submissions); i++ {
		res := ArgMax(predsInFloats[i-1])
		submissions[i][1] = fmt.Sprint(res) // cast to string
	}
	WriteCSV(*output, submissions)
}
