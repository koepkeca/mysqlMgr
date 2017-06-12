package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"time"
)

type dataElement struct {
	example_data_1 string
	example_data_2 string
	example_data_3 int32
	example_data_4 int32
}

type randomDataPool struct {
	words []string
	sz    int64
}

func (r *randomDataPool) Load() (e error) {
	srcFile := "/usr/share/dict/words"
	fB, e := ioutil.ReadFile(srcFile)
	if e != nil {
		return
	}
	tmpByte := bytes.Split(fB, []byte("\n"))
	for _, next := range tmpByte {
		r.words = append(r.words, string(next))
		r.sz = r.sz + 1
	}
	return
}

func (r *randomDataPool) NewDataElement() (d dataElement) {
	rand.Seed(time.Now().UTC().UnixNano())
	d.example_data_1 = r.words[rand.Int63n(r.sz-1)]
	d.example_data_2 = r.words[rand.Int63n(r.sz-1)]
	d.example_data_3 = rand.Int31()
	d.example_data_4 = rand.Int31()
	return
}
