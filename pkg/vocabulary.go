/*
Copyright 2018 The Vergo Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

// Vocabulary represents a vocabulary containing correct words
type Vocabulary struct {
	// FileName is the vocabulary file path
	FileName string
	// KnownWords indicates all the words count in this vocabulary
	KnownWords map[string]int64

	re *regexp.Regexp
	sync.Mutex
}

// NewVocabulary will create a new Vocabulary
func NewVocabulary(fileName string) (*Vocabulary, error) {
	v := &Vocabulary{FileName: fileName}
	err := v.init()
	if err != nil {
		return nil, err
	}
	return v, nil
}

// init will read the file line by line and compose the vocabulary
func (v *Vocabulary) init() error {
	v.KnownWords = make(map[string]int64)

	// matching all the words in upper and lower case
	// words like what's, re-enable are considered as valid as well
	re, err := regexp.Compile("([a-zA-Z'-]+)*")
	if err != nil {
		return err
	}
	v.re = re

	reader, err := os.Open(v.FileName)
	if err != nil {
		return err
	}
	defer reader.Close()
	bufReader := bufio.NewReader(reader)

	lineChan := make(chan string)
	defer close(lineChan)
	lineEOFChan := make(chan bool, 1)
	defer close(lineEOFChan)

	var wg sync.WaitGroup
	go v.compose(lineChan, lineEOFChan, &wg)

	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			lineEOFChan <- true
			if err == io.EOF {
				break
			}
			return err
		}
		wg.Add(1)
		lineChan <- line
	}

	wg.Wait()
	return nil
}

// compose will find and count valid words
func (v *Vocabulary) compose(lineChan chan string, lineEOFChan chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case line := <-lineChan:
			{
				words := v.re.FindAllString(line, -1)
				for _, word := range words {
					word = strings.ToLower(word)
					v.Lock()
					if _, found := v.KnownWords[word]; !found {
						v.KnownWords[word] = 1
					} else {
						v.KnownWords[word]++
					}
					v.Unlock()
				}
				wg.Done()
			}
		case <-lineEOFChan:
			{
				return
			}
		}
	}
}
