// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Identifier        []Identifier               `json:"identifier,omitempty"`
	BasedOn           *Reference                 `json:"basedOn,omitempty"`
	Parent            *Reference                 `json:"parent,omitempty"`
	Status            ObservationStatus          `json:"status"`
	Method            *CodeableConcept           `json:"method,omitempty"`
	Code              *CodeableConcept           `json:"code,omitempty"`
	Subject           Reference                  `json:"subject"`
	Encounter         *Reference                 `json:"encounter,omitempty"`
	Condition         *Reference                 `json:"condition,omitempty"`
	Performer         *Reference                 `json:"performer,omitempty"`
	ReasonCode        []CodeableConcept          `json:"reasonCode,omitempty"`
	ReasonReference   []Reference                `json:"reasonReference,omitempty"`
	Basis             []Reference                `json:"basis,omitempty"`
	Prediction        []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation        *string                    `json:"mitigation,omitempty"`
	Note              []Annotation               `json:"note,omitempty"`
}
type RiskAssessmentPrediction struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Outcome           *CodeableConcept `json:"outcome,omitempty"`
	QualitativeRisk   *CodeableConcept `json:"qualitativeRisk,omitempty"`
	RelativeRisk      *string          `json:"relativeRisk,omitempty"`
	Rationale         *string          `json:"rationale,omitempty"`
}
type OtherRiskAssessment RiskAssessment

// MarshalJSON marshals the given RiskAssessment as JSON into a byte slice
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
}

// UnmarshalRiskAssessment unmarshals a RiskAssessment.
func UnmarshalRiskAssessment(b []byte) (RiskAssessment, error) {
	var riskAssessment RiskAssessment
	if err := json.Unmarshal(b, &riskAssessment); err != nil {
		return riskAssessment, err
	}
	return riskAssessment, nil
}