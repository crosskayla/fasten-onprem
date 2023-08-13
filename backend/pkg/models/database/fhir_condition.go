// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirCondition struct {
	models.ResourceBase
	// Abatement as age or age range
	// https://hl7.org/fhir/r4/search.html#quantity
	AbatementAge datatypes.JSON `gorm:"column:abatementAge;type:text;serializer:json" json:"abatementAge,omitempty"`
	// Date-related abatements (dateTime and period)
	// https://hl7.org/fhir/r4/search.html#date
	AbatementDate *time.Time `gorm:"column:abatementDate;type:datetime" json:"abatementDate,omitempty"`
	// Abatement as a string
	// https://hl7.org/fhir/r4/search.html#string
	AbatementString datatypes.JSON `gorm:"column:abatementString;type:text;serializer:json" json:"abatementString,omitempty"`
	// Person who asserts this condition
	// https://hl7.org/fhir/r4/search.html#reference
	Asserter datatypes.JSON `gorm:"column:asserter;type:text;serializer:json" json:"asserter,omitempty"`
	// Anatomical location, if relevant
	// https://hl7.org/fhir/r4/search.html#token
	BodySite datatypes.JSON `gorm:"column:bodySite;type:text;serializer:json" json:"bodySite,omitempty"`
	// The category of the condition
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// The clinical status of the condition
	// https://hl7.org/fhir/r4/search.html#token
	ClinicalStatus datatypes.JSON `gorm:"column:clinicalStatus;type:text;serializer:json" json:"clinicalStatus,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Code that identifies the allergy or intolerance
	   * [Condition](condition.html): Code for the condition
	   * [DeviceRequest](devicerequest.html): Code for what is being requested/ordered
	   * [DiagnosticReport](diagnosticreport.html): The code for the report, as opposed to codes for the atomic results, which are the names on the observation resource referred to from the result
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a condition code
	   * [List](list.html): What the purpose of this list is
	   * [Medication](medication.html): Returns medications for a specific code
	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication code
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine code
	   * [MedicationRequest](medicationrequest.html): Return prescriptions of this medication code
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication code
	   * [Observation](observation.html): The code of the observation type
	   * [Procedure](procedure.html): A code to identify a  procedure
	   * [ServiceRequest](servicerequest.html): What is being requested/ordered
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Code datatypes.JSON `gorm:"column:code;type:text;serializer:json" json:"code,omitempty"`
	// Encounter created as part of
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// Manifestation/symptom
	// https://hl7.org/fhir/r4/search.html#token
	Evidence datatypes.JSON `gorm:"column:evidence;type:text;serializer:json" json:"evidence,omitempty"`
	// Supporting information found elsewhere
	// https://hl7.org/fhir/r4/search.html#reference
	EvidenceDetail datatypes.JSON `gorm:"column:evidenceDetail;type:text;serializer:json" json:"evidenceDetail,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): External ids for this item
	   * [CarePlan](careplan.html): External Ids for this plan
	   * [CareTeam](careteam.html): External Ids for this team
	   * [Composition](composition.html): Version-independent identifier for the Composition
	   * [Condition](condition.html): A unique identifier of the condition record
	   * [Consent](consent.html): Identifier for this record (external references)
	   * [DetectedIssue](detectedissue.html): Unique id for the detected issue
	   * [DeviceRequest](devicerequest.html): Business identifier for request/order
	   * [DiagnosticReport](diagnosticreport.html): An identifier for the report
	   * [DocumentManifest](documentmanifest.html): Unique Identifier for the set of documents
	   * [DocumentReference](documentreference.html): Master Version Specific Identifier
	   * [Encounter](encounter.html): Identifier(s) by which this encounter is known
	   * [EpisodeOfCare](episodeofcare.html): Business Identifier(s) relevant for this EpisodeOfCare
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a record identifier
	   * [Goal](goal.html): External Ids for this goal
	   * [ImagingStudy](imagingstudy.html): Identifiers for the Study, such as DICOM Study Instance UID and Accession number
	   * [Immunization](immunization.html): Business identifier
	   * [List](list.html): Business identifier
	   * [MedicationAdministration](medicationadministration.html): Return administrations with this external identifier
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with this external identifier
	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this external identifier
	   * [MedicationStatement](medicationstatement.html): Return statements with this external identifier
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this external identifier
	   * [Observation](observation.html): The unique id for a particular observation
	   * [Procedure](procedure.html): A unique identifier for a procedure
	   * [RiskAssessment](riskassessment.html): Unique identifier for the assessment
	   * [ServiceRequest](servicerequest.html): Identifiers assigned to this order
	   * [SupplyDelivery](supplydelivery.html): External identifier
	   * [SupplyRequest](supplyrequest.html): Business Identifier for SupplyRequest
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this external identifier
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Onsets as age or age range
	// https://hl7.org/fhir/r4/search.html#quantity
	OnsetAge datatypes.JSON `gorm:"column:onsetAge;type:text;serializer:json" json:"onsetAge,omitempty"`
	// Date related onsets (dateTime and Period)
	// https://hl7.org/fhir/r4/search.html#date
	OnsetDate *time.Time `gorm:"column:onsetDate;type:datetime" json:"onsetDate,omitempty"`
	// Onsets as a string
	// https://hl7.org/fhir/r4/search.html#string
	OnsetInfo datatypes.JSON `gorm:"column:onsetInfo;type:text;serializer:json" json:"onsetInfo,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Date record was first recorded
	// https://hl7.org/fhir/r4/search.html#date
	RecordedDate *time.Time `gorm:"column:recordedDate;type:datetime" json:"recordedDate,omitempty"`
	// The severity of the condition
	// https://hl7.org/fhir/r4/search.html#token
	Severity datatypes.JSON `gorm:"column:severity;type:text;serializer:json" json:"severity,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// Simple summary (disease specific)
	// https://hl7.org/fhir/r4/search.html#token
	Stage datatypes.JSON `gorm:"column:stage;type:text;serializer:json" json:"stage,omitempty"`
	// Who has the condition?
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// unconfirmed | provisional | differential | confirmed | refuted | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	VerificationStatus datatypes.JSON `gorm:"column:verificationStatus;type:text;serializer:json" json:"verificationStatus,omitempty"`
}

func (s *FhirCondition) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"abatementAge":       "quantity",
		"abatementDate":      "date",
		"abatementString":    "string",
		"asserter":           "reference",
		"bodySite":           "token",
		"category":           "token",
		"clinicalStatus":     "token",
		"code":               "token",
		"encounter":          "reference",
		"evidence":           "token",
		"evidenceDetail":     "reference",
		"identifier":         "token",
		"language":           "token",
		"lastUpdated":        "date",
		"onsetAge":           "quantity",
		"onsetDate":          "date",
		"onsetInfo":          "string",
		"profile":            "reference",
		"recordedDate":       "date",
		"severity":           "token",
		"sourceUri":          "uri",
		"stage":              "token",
		"subject":            "reference",
		"tag":                "token",
		"text":               "string",
		"type":               "special",
		"verificationStatus": "token",
	}
	return searchParameters
}
func (s *FhirCondition) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting AbatementAge
	abatementAgeResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.abatementAge | Condition.abatementRange'))")
	if err == nil && abatementAgeResult.String() != "undefined" {
		s.AbatementAge = []byte(abatementAgeResult.String())
	}
	// extracting AbatementDate
	abatementDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Condition.abatementDateTime | Condition.abatementPeriod')[0]")
	if err == nil && abatementDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, abatementDateResult.String())
		if err == nil {
			s.AbatementDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", abatementDateResult.String())
			if err == nil {
				s.AbatementDate = &d
			}
		}
	}
	// extracting AbatementString
	abatementStringResult, err := vm.RunString(` 
							AbatementStringResult = window.fhirpath.evaluate(fhirResource, 'Condition.abatementString')
							AbatementStringProcessed = AbatementStringResult.reduce((accumulator, currentValue) => {
								if (typeof currentValue === 'string') {
									//basic string
									accumulator.push(currentValue)
								} else if (currentValue.family  || currentValue.given) {
									//HumanName http://hl7.org/fhir/R4/datatypes.html#HumanName
									var humanNameParts = []
									if (currentValue.prefix) {
										humanNameParts = humanNameParts.concat(currentValue.prefix)
									}
									if (currentValue.given) {	
										humanNameParts = humanNameParts.concat(currentValue.given)
									}	
									if (currentValue.family) {	
										humanNameParts.push(currentValue.family)	
									}	
									if (currentValue.suffix) {	
										humanNameParts = humanNameParts.concat(currentValue.suffix)	
									}
									accumulator.push(humanNameParts.join(" "))
								} else if (currentValue.city || currentValue.state || currentValue.country || currentValue.postalCode) {
									//Address http://hl7.org/fhir/R4/datatypes.html#Address
									var addressParts = []		
									if (currentValue.line) {
										addressParts = addressParts.concat(currentValue.line)
									}
									if (currentValue.city) {
										addressParts.push(currentValue.city)
									}	
									if (currentValue.state) {	
										addressParts.push(currentValue.state)
									}	
									if (currentValue.postalCode) {
										addressParts.push(currentValue.postalCode)
									}	
									if (currentValue.country) {
										addressParts.push(currentValue.country)	
									}	
									accumulator.push(addressParts.join(" "))
								} else {
									//string, boolean
									accumulator.push(currentValue)
								}
								return accumulator
							}, [])
						
							if(AbatementStringProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AbatementStringProcessed)
							}
						 `)
	if err == nil && abatementStringResult.String() != "undefined" {
		s.AbatementString = []byte(abatementStringResult.String())
	}
	// extracting Asserter
	asserterResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.asserter'))")
	if err == nil && asserterResult.String() != "undefined" {
	}
	// extracting BodySite
	bodySiteResult, err := vm.RunString(` 
							BodySiteResult = window.fhirpath.evaluate(fhirResource, 'Condition.bodySite')
							BodySiteProcessed = BodySiteResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(BodySiteProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(BodySiteProcessed)
							}
						 `)
	if err == nil && bodySiteResult.String() != "undefined" {
		s.BodySite = []byte(bodySiteResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString(` 
							CategoryResult = window.fhirpath.evaluate(fhirResource, 'Condition.category')
							CategoryProcessed = CategoryResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(CategoryProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(CategoryProcessed)
							}
						 `)
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting ClinicalStatus
	clinicalStatusResult, err := vm.RunString(` 
							ClinicalStatusResult = window.fhirpath.evaluate(fhirResource, 'Condition.clinicalStatus')
							ClinicalStatusProcessed = ClinicalStatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ClinicalStatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ClinicalStatusProcessed)
							}
						 `)
	if err == nil && clinicalStatusResult.String() != "undefined" {
		s.ClinicalStatus = []byte(clinicalStatusResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString(` 
							CodeResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')
							CodeProcessed = CodeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(CodeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(CodeProcessed)
							}
						 `)
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.encounter'))")
	if err == nil && encounterResult.String() != "undefined" {
	}
	// extracting Evidence
	evidenceResult, err := vm.RunString(` 
							EvidenceResult = window.fhirpath.evaluate(fhirResource, 'Condition.evidence.code')
							EvidenceProcessed = EvidenceResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(EvidenceProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EvidenceProcessed)
							}
						 `)
	if err == nil && evidenceResult.String() != "undefined" {
		s.Evidence = []byte(evidenceResult.String())
	}
	// extracting EvidenceDetail
	evidenceDetailResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.evidence.detail'))")
	if err == nil && evidenceDetailResult.String() != "undefined" {
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(IdentifierProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(IdentifierProcessed)
							}
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'Resource.language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting OnsetAge
	onsetAgeResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.onsetAge | Condition.onsetRange'))")
	if err == nil && onsetAgeResult.String() != "undefined" {
		s.OnsetAge = []byte(onsetAgeResult.String())
	}
	// extracting OnsetDate
	onsetDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Condition.onsetDateTime | Condition.onsetPeriod')[0]")
	if err == nil && onsetDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, onsetDateResult.String())
		if err == nil {
			s.OnsetDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", onsetDateResult.String())
			if err == nil {
				s.OnsetDate = &d
			}
		}
	}
	// extracting OnsetInfo
	onsetInfoResult, err := vm.RunString(` 
							OnsetInfoResult = window.fhirpath.evaluate(fhirResource, 'Condition.onsetString')
							OnsetInfoProcessed = OnsetInfoResult.reduce((accumulator, currentValue) => {
								if (typeof currentValue === 'string') {
									//basic string
									accumulator.push(currentValue)
								} else if (currentValue.family  || currentValue.given) {
									//HumanName http://hl7.org/fhir/R4/datatypes.html#HumanName
									var humanNameParts = []
									if (currentValue.prefix) {
										humanNameParts = humanNameParts.concat(currentValue.prefix)
									}
									if (currentValue.given) {	
										humanNameParts = humanNameParts.concat(currentValue.given)
									}	
									if (currentValue.family) {	
										humanNameParts.push(currentValue.family)	
									}	
									if (currentValue.suffix) {	
										humanNameParts = humanNameParts.concat(currentValue.suffix)	
									}
									accumulator.push(humanNameParts.join(" "))
								} else if (currentValue.city || currentValue.state || currentValue.country || currentValue.postalCode) {
									//Address http://hl7.org/fhir/R4/datatypes.html#Address
									var addressParts = []		
									if (currentValue.line) {
										addressParts = addressParts.concat(currentValue.line)
									}
									if (currentValue.city) {
										addressParts.push(currentValue.city)
									}	
									if (currentValue.state) {	
										addressParts.push(currentValue.state)
									}	
									if (currentValue.postalCode) {
										addressParts.push(currentValue.postalCode)
									}	
									if (currentValue.country) {
										addressParts.push(currentValue.country)	
									}	
									accumulator.push(addressParts.join(" "))
								} else {
									//string, boolean
									accumulator.push(currentValue)
								}
								return accumulator
							}, [])
						
							if(OnsetInfoProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(OnsetInfoProcessed)
							}
						 `)
	if err == nil && onsetInfoResult.String() != "undefined" {
		s.OnsetInfo = []byte(onsetInfoResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
	}
	// extracting RecordedDate
	recordedDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Condition.recordedDate')[0]")
	if err == nil && recordedDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, recordedDateResult.String())
		if err == nil {
			s.RecordedDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", recordedDateResult.String())
			if err == nil {
				s.RecordedDate = &d
			}
		}
	}
	// extracting Severity
	severityResult, err := vm.RunString(` 
							SeverityResult = window.fhirpath.evaluate(fhirResource, 'Condition.severity')
							SeverityProcessed = SeverityResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(SeverityProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SeverityProcessed)
							}
						 `)
	if err == nil && severityResult.String() != "undefined" {
		s.Severity = []byte(severityResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Stage
	stageResult, err := vm.RunString(` 
							StageResult = window.fhirpath.evaluate(fhirResource, 'Condition.stage.summary')
							StageProcessed = StageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(StageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StageProcessed)
							}
						 `)
	if err == nil && stageResult.String() != "undefined" {
		s.Stage = []byte(stageResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Condition.subject'))")
	if err == nil && subjectResult.String() != "undefined" {
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting VerificationStatus
	verificationStatusResult, err := vm.RunString(` 
							VerificationStatusResult = window.fhirpath.evaluate(fhirResource, 'Condition.verificationStatus')
							VerificationStatusProcessed = VerificationStatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(VerificationStatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(VerificationStatusProcessed)
							}
						 `)
	if err == nil && verificationStatusResult.String() != "undefined" {
		s.VerificationStatus = []byte(verificationStatusResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirCondition) TableName() string {
	return "fhir_condition"
}
