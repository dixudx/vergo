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
	"io/ioutil"
	"os"
	"path/filepath"
)

func DefaultVocabularyFiles() ([]string, error) {
	var vocabularyFiles []string
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	vocabularyRoot := filepath.Join(rootDir, DefaultVocabularyFolder)
	files, err := ioutil.ReadDir(vocabularyRoot)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		vocabularyFiles = append(vocabularyFiles, filepath.Join(vocabularyRoot, f.Name()))
	}
	return vocabularyFiles, nil
}
