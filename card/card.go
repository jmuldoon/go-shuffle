package card

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Card is the structure holding the FaceEnum and ValueEnum of each card
type Card struct {
	Suite string
	Face  string
	Value int
}

// Faces repesents the yaml reading for the card name and value, to be extended
// potentially later for an arbitrary value/data
type Faces struct {
	Name  string `yaml:"Name"`
	Value int    `yaml:"Value"`
}

// Set repesents the yaml reading for the card set's suite and faces
type Set struct {
	Suites    []string `yaml:"Suites"`
	CardFaces []Faces  `yaml:"Faces"`
}

// ReadSet allows for the user to readin a cardset for the given .yaml path.
func (s *Set) ReadSet(p string) *Set {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(data, &s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s
}
