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
	"strings"
)

type Word struct {
	SingleWord string
	correct    string
	Vocabularies []*Vocabulary
}

// All edits that are one edit away from `SingleWord`
// TODO: what if the length is 1?
func (word *Word) Hamming1() []string {
	var variants []string
	var wordSplits []Split

	// splits
	for i := 0; i < len(word.SingleWord)+1; i++ {
		wordSplits = append(wordSplits, Split{
			Left:  word.SingleWord[:i],
			Right: word.SingleWord[i:],
		})
	}

	for _, split := range wordSplits {
		// deletes
		if len(split.Right) >= 1 {
			variants = append(variants, split.Left+split.Right[1:])
		}
		// transposes
		if len(split.Right) > 1 {
			variants = append(variants, split.Left+string(split.Right[1])+string(split.Right[0])+split.Right[2:])
		}

		for _, letter := range strings.Split(DefaultAlphabet, "") {
			// replaces
			if len(split.Right) >= 1 {
				variants = append(variants, split.Left+letter+split.Right[1:])
			}
			// inserts
			variants = append(variants, split.Left+letter+split.Right)
		}
	}

	return variants
}

// All edits that are two edits away from `SingleWord`.
func (word *Word) Hamming2() []string {
	var variants []string
	for _, word1 := range word.Hamming1() {
		newWord := Word{SingleWord: word1}
		for _, word2 := range newWord.Hamming1() {
			variants = append(variants, word2)
		}
	}
	return variants
}

type Split struct {
	Left  string
	Right string
}
