// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file contains the schemas for HL7 messages, segments and values for HL7v2 version 2.3.0.
// It has been auto-generated from the HL7v2 specification.

package hl7

import "reflect"


// AD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type AD struct {
	StreetAddress *ST `hl7:"false,Street Address"`
	OtherDesignation *ST `hl7:"false,Other Designation"`
	City *ST `hl7:"false,City"`
	StateOrProvince *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode *ST `hl7:"false,Zip Or Postal Code"`
	Country *ID `hl7:"false,Country"`
	AddressType *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// CD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CD struct {
	ChannelIdentifier *CM `hl7:"false,Channel Identifier"`
	ElectrodeNames *CM `hl7:"false,Electrode Names"`
	ChannelSensitivityUnits *CM `hl7:"false,Channel Sensitivity/Units"`
	CalibrationParameters *CM `hl7:"false,Calibration Parameters"`
	SamplingFrequency *NM `hl7:"false,Sampling Frequency"`
	MinimumMaximumDataValues *CM `hl7:"false,Minimum/Maximum Data Values"`
}

// CE represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CE struct {
	Identifier *ID `hl7:"false,Identifier"`
	Text *ST `hl7:"false,Text"`
	NameOfCodingSystem *ST `hl7:"false,Name Of Coding System"`
	AlternateIdentifier *ID `hl7:"false,Alternate Identifier"`
	AlternateText *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem *ST `hl7:"false,Name Of Alternate Coding System"`
}

// CF represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CF struct {
	Identifier *ID `hl7:"false,Identifier"`
	FormattedText *FT `hl7:"false,Formatted Text"`
	NameOfCodingSystem *ST `hl7:"false,Name Of Coding System"`
	AlternateIdentifier *ID `hl7:"false,Alternate Identifier"`
	AlternateFormattedText *FT `hl7:"false,Alternate Formatted Text"`
	NameOfAlternateCodingSystem *ST `hl7:"false,Name Of Alternate Coding System"`
}

// CK represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CK struct {
	IDNumber *NM `hl7:"false,ID Number"`
	CheckDigit *ST `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
}

// CK_ACCOUNT_NO represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CK_ACCOUNT_NO struct {
	AccountNumber *NM `hl7:"false,Account Number"`
	CheckDigit *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID *ID `hl7:"false,Facility ID"`
}

// CK_PAT_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CK_PAT_ID struct {
	PatientID *ST `hl7:"false,Patient ID"`
	CheckDigit *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID *ID `hl7:"false,Facility ID"`
}

// CM_ABS_RANGE represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_ABS_RANGE struct {
	Range *CM `hl7:"false,Range"`
	NumericChange *NM `hl7:"false,Numeric Change"`
	PercentPerChange *NM `hl7:"false,Percent Per Change"`
	Days *NM `hl7:"false,Days"`
}

// CM_AUI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_AUI struct {
	AuthorizationNumber *ST `hl7:"false,Authorization Number"`
	Date *TS `hl7:"false,Date"`
	Source *ST `hl7:"false,Source"`
}

// CM_BATCH_TOTAL represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_BATCH_TOTAL struct {
	BatchTotal1 *NM `hl7:"false,Batch Total 1"`
	BatchTotal2 *NM `hl7:"false,Batch Total 2"`
}

// CM_CCD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_CCD struct {
	WhenToChargeCode *ID `hl7:"false,When To Charge Code"`
	DateTime *TS `hl7:"false,Date/Time"`
}

// CM_DDI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DDI struct {
	DelayDays *NM `hl7:"false,Delay Days"`
	Amount *NM `hl7:"false,Amount"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_DIN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DIN struct {
	Date *TS `hl7:"false,Date"`
	InstitutionName *CE `hl7:"false,Institution Name"`
}

// CM_DLD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DLD struct {
	DischargeLocation *ID `hl7:"false,Discharge Location"`
	EffectiveDate *TS `hl7:"false,Effective Date"`
}

// CM_DLT represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DLT struct {
	Range *CM `hl7:"false,Range"`
	NumericThreshold *NM `hl7:"false,Numeric Threshold"`
	Change *ST `hl7:"false,Change"`
	LengthOfTimeDays *NM `hl7:"false,Length Of Time-Days"`
}

// CM_DTN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DTN struct {
	DayType *IS `hl7:"false,Day Type"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_EIP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_EIP struct {
	ParentSPlacerOrderNumber *EI `hl7:"false,Parent´s Placer Order Number"`
	ParentSFillerOrderNumber *EI `hl7:"false,Parent´s Filler Order Number"`
}

// CM_ELD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_ELD struct {
	SegmentID *ST `hl7:"false,Segment ID"`
	Sequence *NM `hl7:"false,Sequence"`
	FieldPosition *NM `hl7:"false,Field Position"`
	CodeIdentifyingError *CE `hl7:"false,Code Identifying Error"`
}

// CM_FILLER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_FILLER struct {
	UniqueFillerId *ID `hl7:"false,Unique Filler Id"`
	FillerApplicationID *ID `hl7:"false,Filler Application ID"`
}

// CM_FINANCE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_FINANCE struct {
	FinancialClassID *ID `hl7:"false,Financial Class ID"`
	EffectiveDate *TS `hl7:"false,Effective Date"`
}

// CM_GROUP_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_GROUP_ID struct {
	UniqueGroupId *ID `hl7:"false,Unique Group Id"`
	PlacerApplicationId *ID `hl7:"false,Placer Application Id"`
}

// CM_INTERNAL_LOCATION represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_INTERNAL_LOCATION struct {
	NurseUnitStation *ID `hl7:"false,Nurse Unit (Station)"`
	Room *ID `hl7:"false,Room"`
	Bed *ID `hl7:"false,Bed"`
	FacilityID *ID `hl7:"false,Facility ID"`
	BedStatus *ID `hl7:"false,Bed Status"`
	Etage *ID `hl7:"false,Etage"`
	Klinik *ID `hl7:"false,Klinik"`
	Zentrum *ID `hl7:"false,Zentrum"`
}

// CM_JOB_CODE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_JOB_CODE struct {
	JobCode *ID `hl7:"false,Job Code"`
	EmployeeClassification *ID `hl7:"false,Employee Classification"`
}

// CM_LA1 represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_LA1 struct {
	PointOfCare *ST `hl7:"false,Point Of Care"`
	Room *IS `hl7:"false,Room"`
	Bed *IS `hl7:"false,Bed"`
	Facility *HD `hl7:"false,Facility"`
	LocationStatus *IS `hl7:"false,Location Status"`
	PersonLocationType *IS `hl7:"false,Person Location Type"`
	Building *IS `hl7:"false,Building"`
	Floor *ST `hl7:"false,Floor"`
	StreetAddress *ST `hl7:"false,Street Address"`
	OtherDesignation *ST `hl7:"false,Other Designation"`
	City *ST `hl7:"false,City"`
	StateOrProvince *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode *ST `hl7:"false,Zip Or Postal Code"`
	Country *ID `hl7:"false,Country"`
	AddressType *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// CM_LICENSE_NO represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_LICENSE_NO struct {
	LicenseNumber *ST `hl7:"false,License Number"`
	IssuingStateProvinceCountry *ST `hl7:"false,Issuing State,Province,Country"`
}

// CM_MOC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_MOC struct {
	DollarAmount *MO `hl7:"false,Dollar Amount"`
	ChargeCode *CE `hl7:"false,Charge Code"`
}

// CM_MSG represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_MSG struct {
	MessageType *ID `hl7:"false,Message Type"`
	TriggerEvent *ID `hl7:"false,Trigger Event"`
}

// CM_NDL represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_NDL struct {
	Name *CN `hl7:"false,Name"`
	StartDateTime *TS `hl7:"false,Start Date/Time"`
	EndDateTime *TS `hl7:"false,End Date/Time"`
	PointOfCare *IS `hl7:"false,Point Of Care"`
	Room *IS `hl7:"false,Room"`
	Bed *IS `hl7:"false,Bed"`
	Facility *HD `hl7:"false,Facility"`
	LocationStatus *IS `hl7:"false,Location Status"`
	PersonLocationType *IS `hl7:"false,Person Location Type"`
	Building *IS `hl7:"false,Building"`
	Floor *ST `hl7:"false,Floor"`
}

// CM_OCD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_OCD struct {
	OccurrenceCode *CE `hl7:"false,Occurrence Code"`
	OccurrenceDate *DT `hl7:"false,Occurrence Date"`
}

// CM_OSP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_OSP struct {
	OccurrenceSpanCode *CE `hl7:"false,Occurrence Span Code"`
	OccurrenceSpanStartDate *DT `hl7:"false,Occurrence Span Start Date"`
	OccurrenceSpanStopDate *DT `hl7:"false,Occurrence Span Stop Date"`
}

// CM_PAT_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PAT_ID struct {
	PatientID *ST `hl7:"false,Patient ID"`
	CheckDigit *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID *ID `hl7:"false,Facility ID"`
	Type *ID `hl7:"false,Type"`
}

// CM_PAT_ID_0192 represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PAT_ID_0192 struct {
	PatientID *ST `hl7:"false,Patient ID"`
	CheckDigit *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID *ID `hl7:"false,Facility ID"`
	Type *ID `hl7:"false,Type"`
}

// CM_PCF represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PCF struct {
	PreCertificationPatientType *IS `hl7:"false,Pre-Certification Patient Type"`
	PreCertificationRequired *ID `hl7:"false,Pre-Certification Required"`
	PreCertificationWindwow *TS `hl7:"false,Pre-Certification Windwow"`
}

// CM_PEN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PEN struct {
	PenaltyType *IS `hl7:"false,Penalty Type"`
	PenaltyAmount *NM `hl7:"false,Penalty Amount"`
}

// CM_PI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PI struct {
	IDNumber *ST `hl7:"false,ID Number"`
	TypeOfIDNumber *IS `hl7:"false,Type Of ID Number"`
	OtherQualifyingInfo *ST `hl7:"false,Other Qualifying Info"`
}

// CM_PIP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PIP struct {
	Privilege *CE `hl7:"false,Privilege"`
	PrivilegeClass *CE `hl7:"false,Privilege Class"`
	ExpirationDate *DT `hl7:"false,Expiration Date"`
	ActivationDate *DT `hl7:"false,Activation Date"`
}

// CM_PLACER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PLACER struct {
	UniquePlacerId *ID `hl7:"false,Unique Placer Id"`
	PlacerApplication *ID `hl7:"false,Placer Application"`
}

// CM_PLN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PLN struct {
	IDNumber *ST `hl7:"false,ID Number"`
	TypeOfIDNumber *IS `hl7:"false,Type Of ID Number"`
	StateOtherQualifyingInfo *ST `hl7:"false,State/Other Qualifying Info"`
	ExpirationDate *DT `hl7:"false,Expiration Date"`
}

// CM_POSITION represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_POSITION struct {
	Saal *ST `hl7:"false,Saal"`
	Tisch *ST `hl7:"false,Tisch"`
	Stuhl *ST `hl7:"false,Stuhl"`
}

// CM_PRACTITIONER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PRACTITIONER struct {
	ProcedurePractitionerID *CN `hl7:"false,Procedure Practitioner  ID"`
	ProcedurePractitionerType *ID `hl7:"false,Procedure Practitioner Type"`
}

// CM_PRL represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PRL struct {
	OBX3ObservationIdentifierOfParentResult *CE `hl7:"false,OBX-3 Observation Identifier Of Parent Result"`
	OBX4SubIDOfParentResult *ST `hl7:"false,OBX-4 Sub-ID Of Parent Result"`
	PartOfOBX5ObservationResultFromParent *TX `hl7:"false,Part Of OBX-5 Observation Result From Parent"`
}

// CM_PTA represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PTA struct {
	PolicyType *IS `hl7:"false,Policy Type"`
	AmountClass *IS `hl7:"false,Amount Class"`
	Amount *NM `hl7:"false,Amount"`
}

// CM_RANGE represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RANGE struct {
	LowValue *CE `hl7:"false,Low Value"`
	HighValue *CE `hl7:"false,High Value"`
}

// CM_RFR represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RFR struct {
	ReferenceRange *CM `hl7:"false,Reference Range"`
	Sex *IS `hl7:"false,Sex"`
	AgeRange *CM `hl7:"false,Age Range"`
	AgeGestation *CM `hl7:"false,Age Gestation"`
	Species *TX `hl7:"false,Species"`
	RaceSubspecies *ST `hl7:"false,Race/Subspecies"`
	Conditions *TX `hl7:"false,Conditions"`
}

// CM_RI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RI struct {
	RepeatPattern *IS `hl7:"false,Repeat Pattern"`
	ExplicitTimeInterval *ST `hl7:"false,Explicit Time Interval"`
}

// CM_RMC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RMC struct {
	RoomType *IS `hl7:"false,Room Type"`
	AmountType *IS `hl7:"false,Amount Type"`
	CoverageAmount *NM `hl7:"false,Coverage Amount"`
}

// CM_SPD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_SPD struct {
	SpecialtyName *ST `hl7:"false,Specialty Name"`
	GoverningBoard *ST `hl7:"false,Governing Board"`
	EligibleOrCertified *ID `hl7:"false,Eligible Or Certified"`
	DateOfCertification *DT `hl7:"false,Date Of Certification"`
}

// CM_SPS represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_SPS struct {
	SpecimenSourceNameOrCode *CE `hl7:"false,Specimen Source Name Or Code"`
	Additives *TX `hl7:"false,Additives"`
	Freetext *TX `hl7:"false,Freetext"`
	BodySite *CE `hl7:"false,Body Site"`
	SiteModifier *CE `hl7:"false,Site Modifier"`
	CollectionModifierMethodCode *CE `hl7:"false,Collection Modifier Method Code"`
}

// CM_UVC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_UVC struct {
	ValueCode *IS `hl7:"false,Value Code"`
	ValueAmount *NM `hl7:"false,Value Amount"`
}

// CM_VR represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_VR struct {
	FirstDataCodeValue *ST `hl7:"false,First Data Code Value"`
	LastDataCodeCalue *ST `hl7:"false,Last Data Code Calue"`
}

// CM_WVI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_WVI struct {
	ChannelNumber *NM `hl7:"false,Channel Number"`
	ChannelName *ST `hl7:"false,Channel Name"`
}

// CN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CN struct {
	IDNumber *ST `hl7:"false,ID Number"`
	FamilyName *ST `hl7:"false,Family Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G., MD)"`
	SourceTable *ID `hl7:"false,Source Table"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
}

// CN_PERSON represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CN_PERSON struct {
	IDNumber *ID `hl7:"false,ID Number"`
	FamiliyName *ST `hl7:"false,Familiy Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G. MD)"`
	SourceTableId *ID `hl7:"false,Source Table Id"`
}

// CN_PHYSICIAN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CN_PHYSICIAN struct {
	PhysicianID *ID `hl7:"false,Physician ID"`
	FamiliyName *ST `hl7:"false,Familiy Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G. MD)"`
	SourceTableId *ID `hl7:"false,Source Table Id"`
	Adresse *AD `hl7:"false,Adresse"`
	Telefon *TN `hl7:"false,Telefon"`
	Faxnummer *TN `hl7:"false,Faxnummer"`
	OnlineNummer *TN `hl7:"false,Online-Nummer"`
	EMail *ST `hl7:"false,E-Mail"`
}

// CP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CP struct {
	Price *MO `hl7:"false,Price"`
	PriceType *ID `hl7:"false,Price Type"`
	FromValue *NM `hl7:"false,From Value"`
	ToValue *NM `hl7:"false,To Value"`
	RangeUnits *CE `hl7:"false,Range Units"`
	RangeType *ID `hl7:"false,Range Type"`
}

// CQ represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CQ struct {
	Quantity *NM `hl7:"false,Quantity"`
	Units *CE `hl7:"false,Units"`
}

// CQ_QUANTITY represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CQ_QUANTITY struct {
	Quantity *ST `hl7:"false,Quantity"`
	Units *ST `hl7:"false,Units"`
}

// CX represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CX struct {
	ID *ST `hl7:"false,ID"`
	CheckDigit *ST `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
	IdentifierTypeCode *IS `hl7:"false,Identifier Type Code"`
	AssigningFacility *HD `hl7:"false,Assigning Facility"`
}

// DLN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type DLN struct {
	DriverSLicenseNumber *ST `hl7:"false,Driver´s License Number"`
	IssuingStateProvinceCountry *IS `hl7:"false,Issuing State, Province, Country"`
	ExpirationDate *DT `hl7:"false,Expiration Date"`
}

// DR represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type DR struct {
	RangeStartDateTime *TS `hl7:"false,Range Start Date/Time"`
	RangeEndDateTime *TS `hl7:"false,Range End Date/Time"`
}

// ED represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type ED struct {
	SourceApplication *HD `hl7:"false,Source Application"`
	TypeOfData *ID `hl7:"false,Type Of Data"`
	DataSubtype *ID `hl7:"false,Data Subtype"`
	Encoding *ID `hl7:"false,Encoding"`
	Data *ST `hl7:"false,Data"`
}

// EI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type EI struct {
	EntityIdentifier *ST `hl7:"false,Entity Identifier"`
	NamespaceID *IS `hl7:"false,Namespace ID"`
	UniversalID *ST `hl7:"false,Universal ID"`
	UniversalIDType *ID `hl7:"false,Universal ID Type"`
}

// FC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type FC struct {
	FinancialClass *IS `hl7:"false,Financial Class"`
	EffectiveDate *TS `hl7:"false,Effective Date"`
}

// HD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type HD struct {
	NamespaceID *IS `hl7:"false,Namespace ID"`
	UniversalID *ST `hl7:"false,Universal ID"`
	UniversalIDType *ID `hl7:"false,Universal ID Type"`
}

// JCC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type JCC struct {
	JobCode *IS `hl7:"false,Job Code"`
	JobClass *IS `hl7:"false,Job Class"`
}

// MA represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type MA struct {
	Sample1FromChannel1 *NM `hl7:"false,Sample 1 From Channel 1"`
	Sample1FromChannel2 *NM `hl7:"false,Sample 1 From Channel 2"`
	Sample1FromChannel3 *NM `hl7:"false,Sample 1 From Channel 3"`
	Sample2FromChannel1 *NM `hl7:"false,Sample 2 From Channel 1"`
	Sample2FromChannel2 *NM `hl7:"false,Sample 2 From Channel 2"`
	Sample2FromChannel3 *NM `hl7:"false,Sample 2 From Channel 3"`
}

// MO represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type MO struct {
	Quantity *NM `hl7:"false,Quantity"`
	Denomination *ID `hl7:"false,Denomination"`
}

// NA represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type NA struct {
	Value1 *NM `hl7:"false,Value1"`
	Value2 *NM `hl7:"false,Value2"`
	Value3 *NM `hl7:"false,Value3"`
	Value4 *NM `hl7:"false,Value4"`
}

// PL represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type PL struct {
	PointOfCare *ID `hl7:"false,Point Of Care"`
	Room *IS `hl7:"false,Room"`
	Bed *IS `hl7:"false,Bed"`
	Facility *HD `hl7:"false,Facility"`
	LocationStatus *IS `hl7:"false,Location Status"`
	PersonLocationType *IS `hl7:"false,Person Location Type"`
	Building *IS `hl7:"false,Building"`
	Floor *ST `hl7:"false,Floor"`
	LocationType *ST `hl7:"false,Location Type"`
}

// PN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type PN struct {
	FamiliyName *ST `hl7:"false,Familiy Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G. MD)"`
}

// PPN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type PPN struct {
	IDNumber *ST `hl7:"false,ID Number"`
	FamilyName *ST `hl7:"false,Family Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G., MD)"`
	SourceTable *ID `hl7:"false,Source Table"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
	NameTypeCode *ID `hl7:"false,Name Type Code"`
	IdentifierCheckDigit *ST `hl7:"false,Identifier Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	IdentifierTypeCode *IS `hl7:"false,Identifier Type Code"`
	AssigningFacility *HD `hl7:"false,Assigning Facility"`
	DateTimeActionPerformed *TS `hl7:"false,Date/Time Action Performed"`
}

// PT represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type PT struct {
	ProcessingID *ST `hl7:"false,Processing ID"`
	ProcessingMode *ST `hl7:"false,Processing Mode"`
}

// QIP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type QIP struct {
	FieldName *ST `hl7:"false,Field Name"`
	Value1Value2Value3 *ST `hl7:"false,Value1&Value2&Value3"`
}

// QSC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type QSC struct {
	NameOfField *ST `hl7:"false,Name Of Field"`
	RelationalOperator *ID `hl7:"false,Relational Operator"`
	Value *ST `hl7:"false,Value"`
	RelationalConjunction *ID `hl7:"false,Relational Conjunction"`
}

// RCD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type RCD struct {
	ItemNumber *ST `hl7:"false,Item Number"`
	DateType *ST `hl7:"false,Date Type"`
	MaximumColumnWidth *NM `hl7:"false,Maximum Column Width"`
}

// RI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type RI struct {
	RepeatPattern *IS `hl7:"false,Repeat Pattern"`
	ExplicitTimeInterval *ST `hl7:"false,Explicit Time Interval"`
}

// RP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type RP struct {
	Pointer *ST `hl7:"false,Pointer"`
	ApplicationID *HD `hl7:"false,Application ID"`
	TypeOfData *ID `hl7:"false,Type Of Data"`
	Subtype *ID `hl7:"false,Subtype"`
}

// SCV represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type SCV struct {
	ParameterClass *IS `hl7:"false,Parameter Class"`
	ParameterValue *ST `hl7:"false,Parameter Value"`
}

// SN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type SN struct {
	Comparator *ST `hl7:"false,Comparator"`
	Num1 *NM `hl7:"false,Num1"`
	SeparatorOrSuffix *ST `hl7:"false,Separator Or Suffix"`
	Num2 *NM `hl7:"false,Num2"`
}

// TQ represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type TQ struct {
	Quantity *CQ `hl7:"false,Quantity"`
	Interval *CM `hl7:"false,Interval"`
	Duration *ST `hl7:"false,Duration"`
	StartDateTime *TS `hl7:"false,Start Date/Time"`
	EndDateTime *TS `hl7:"false,End Date/Time"`
	Priority *ST `hl7:"false,Priority"`
	Condition *ST `hl7:"false,Condition"`
	Text *TX `hl7:"false,Text"`
	Conjunction *ST `hl7:"false,Conjunction"`
	OrderSequencing *CM `hl7:"false,Order Sequencing"`
}

// VH represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type VH struct {
	StartDayRange *ID `hl7:"false,Start Day Range"`
	EndDayRange *ID `hl7:"false,End Day Range"`
	StartHourRange *TM `hl7:"false,Start Hour Range"`
	EndHourRange *TM `hl7:"false,End Hour Range"`
}

// XAD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type XAD struct {
	StreetAddress *ST `hl7:"false,Street Address"`
	OtherDesignation *ST `hl7:"false,Other Designation"`
	City *ST `hl7:"false,City"`
	StateOrProvince *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode *ST `hl7:"false,Zip Or Postal Code"`
	Country *ID `hl7:"false,Country"`
	AddressType *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
	CountyParishCode *IS `hl7:"false,County/Parish Code"`
	CensusTract *IS `hl7:"false,Census Tract"`
}

// XCN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type XCN struct {
	IDNumber *ST `hl7:"false,ID Number"`
	FamilyName *ST `hl7:"false,Family Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G., MD)"`
	SourceTable *ID `hl7:"false,Source Table"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
	NameType *ID `hl7:"false,Name Type"`
	IdentifierCheckDigit *ST `hl7:"false,Identifier Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	IdentifierTypeCode *IS `hl7:"false,Identifier Type Code"`
	AssigningFacilityID *HD `hl7:"false,Assigning Facility ID"`
}

// XON represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type XON struct {
	OrganizationName *ST `hl7:"false,Organization Name"`
	OrganizationNameTypeCode *IS `hl7:"false,Organization Name Type Code"`
	IDNumber *NM `hl7:"false,ID Number"`
	CheckDigit *ST `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority *HD `hl7:"false,Assigning Authority"`
	IdentifierTypeCode *IS `hl7:"false,Identifier Type Code"`
	AssigningFacilityID *HD `hl7:"false,Assigning Facility ID"`
}

// XPN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type XPN struct {
	FamilyName *ST `hl7:"false,Family Name"`
	GivenName *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD *ST `hl7:"false,Degree (E.G., MD)"`
	NameTypeCode *ID `hl7:"false,Name Type Code"`
	NameRepresentationCode *ID `hl7:"false,Name Representation Code"`
}

// XTN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type XTN struct {
	9999999999X99999CAnyText *TN `hl7:"false,[(999)] 999-9999 [X99999][C Any Text]"`
	TelecommunicationUseCode *ID `hl7:"false,Telecommunication Use Code"`
	TelecommunicationEquipmentType *ID `hl7:"false,Telecommunication Equipment Type"`
	EmailAddress *ST `hl7:"false,Email Address"`
	CountryCode *NM `hl7:"false,Country Code"`
	AreaCityCode *NM `hl7:"false,Area/City Code"`
	PhoneNumber *NM `hl7:"false,Phone Number"`
	Extension *NM `hl7:"false,Extension"`
	AnyText *ST `hl7:"false,Any Text"`
}

// ACC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ACC struct {
	AccidentDateTime *TS `hl7:"false,Accident Date/Time"` // ACC-1
	AccidentCode *CE `hl7:"false,Accident Code"` // ACC-2
	AccidentLocation *ST `hl7:"false,Accident Location"` // ACC-3
	AutoAccidentState *CE `hl7:"false,Auto Accident State"` // ACC-4
	AccidentJobRelatedIndicator *ID `hl7:"false,Accident Job Related Indicator"` // ACC-5
	AccidentDeathIndicator *ID `hl7:"false,Accident Death Indicator"` // ACC-6
}

func (s *ACC) SegmentName() string {
	return "ACC"
}

// ADD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ADD struct {
	AddendumContinuationPointer *ST `hl7:"false,Addendum Continuation Pointer"` // ADD-1
}

func (s *ADD) SegmentName() string {
	return "ADD"
}

// AIG represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AIG struct {
	SetIDAIG *SI `hl7:"true,Set ID - AIG"` // AIG-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // AIG-2
	ResourceID *CE `hl7:"false,Resource ID"` // AIG-3
	ResourceType *CE `hl7:"true,Resource Type"` // AIG-4
	ResourceGroup []CE `hl7:"false,Resource Group"` // AIG-5
	ResourceQuantity *NM `hl7:"false,Resource Quantity"` // AIG-6
	ResourceQuantityUnits *CE `hl7:"false,Resource Quantity Units"` // AIG-7
	StartDateTime *TS `hl7:"false,Start Date/Time"` // AIG-8
	StartDateTimeOffset *NM `hl7:"false,Start Date/Time Offset"` // AIG-9
	StartDateTimeOffsetUnits *CE `hl7:"false,Start Date/Time Offset Units"` // AIG-10
	Duration *NM `hl7:"false,Duration"` // AIG-11
	DurationUnits *CE `hl7:"false,Duration Units"` // AIG-12
	AllowSubstitutionCode *IS `hl7:"false,Allow Substitution Code"` // AIG-13
	FillerStatusCode *CE `hl7:"false,Filler Status Code"` // AIG-14
}

func (s *AIG) SegmentName() string {
	return "AIG"
}

// AIL represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AIL struct {
	SetIDAIL *SI `hl7:"true,Set ID - AIL"` // AIL-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // AIL-2
	LocationResourceID *PL `hl7:"true,Location Resource ID"` // AIL-3
	LocationType *CE `hl7:"false,Location Type"` // AIL-4
	LocationGroup *CE `hl7:"false,Location Group"` // AIL-5
	StartDateTime *TS `hl7:"false,Start Date/Time"` // AIL-6
	StartDateTimeOffset *NM `hl7:"false,Start Date/Time Offset"` // AIL-7
	StartDateTimeOffsetUnits *CE `hl7:"false,Start Date/Time Offset Units"` // AIL-8
	Duration *NM `hl7:"false,Duration"` // AIL-9
	DurationUnits *CE `hl7:"false,Duration Units"` // AIL-10
	AllowSubstitutionCode *IS `hl7:"false,Allow Substitution Code"` // AIL-11
	FillerStatusCode *CE `hl7:"false,Filler Status Code"` // AIL-12
}

func (s *AIL) SegmentName() string {
	return "AIL"
}

// AIP represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AIP struct {
	SetIDAIP *SI `hl7:"true,Set ID - AIP"` // AIP-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // AIP-2
	PersonnelResourceID *XCN `hl7:"false,Personnel Resource ID"` // AIP-3
	ResourceRole *CE `hl7:"true,Resource Role"` // AIP-4
	ResourceGroup []CE `hl7:"false,Resource Group"` // AIP-5
	StartDateTime *TS `hl7:"false,Start Date/Time"` // AIP-6
	StartDateTimeOffset *NM `hl7:"false,Start Date/Time Offset"` // AIP-7
	StartDateTimeOffsetUnits *CE `hl7:"false,Start Date/Time Offset Units"` // AIP-8
	Duration *NM `hl7:"false,Duration"` // AIP-9
	DurationUnits *CE `hl7:"false,Duration Units"` // AIP-10
	AllowSubstitutionCode *IS `hl7:"false,Allow Substitution Code"` // AIP-11
	FillerStatusCode *CE `hl7:"false,Filler Status Code"` // AIP-12
}

func (s *AIP) SegmentName() string {
	return "AIP"
}

// AIS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AIS struct {
	SetIDAIS *SI `hl7:"true,Set ID - AIS"` // AIS-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // AIS-2
	UniversalServiceIdentifier *CE `hl7:"true,Universal Service Identifier"` // AIS-3
	StartDateTime *TS `hl7:"false,Start Date/Time"` // AIS-4
	StartDateTimeOffset *NM `hl7:"false,Start Date/Time Offset"` // AIS-5
	StartDateTimeOffsetUnits *CE `hl7:"false,Start Date/Time Offset Units"` // AIS-6
	Duration *NM `hl7:"false,Duration"` // AIS-7
	DurationUnits *CE `hl7:"false,Duration Units"` // AIS-8
	AllowSubstitutionCode *IS `hl7:"false,Allow Substitution Code"` // AIS-9
	FillerStatusCode *CE `hl7:"false,Filler Status Code"` // AIS-10
}

func (s *AIS) SegmentName() string {
	return "AIS"
}

// AL1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AL1 struct {
	SetIDAL1 *SI `hl7:"true,Set ID - AL1"` // AL1-1
	AllergyType *ID `hl7:"false,Allergy Type"` // AL1-2
	AllergyCodeMnemonicDescription *CE `hl7:"true,Allergy Code/Mnemonic/ Description"` // AL1-3
	AllergySeverity *ID `hl7:"false,Allergy Severity"` // AL1-4
	AllergyReaction *ST `hl7:"false,Allergy Reaction"` // AL1-5
	IdentificationDate *DT `hl7:"false,Identification Date"` // AL1-6
}

func (s *AL1) SegmentName() string {
	return "AL1"
}

// APR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type APR struct {
	TimeSelectionCriteria []SCV `hl7:"false,Time Selection Criteria"` // APR-1
	ResourceSelectionCriteria []SCV `hl7:"false,Resource Selection Criteria"` // APR-2
	LocationSelectionCriteria []SCV `hl7:"false,Location Selection Criteria"` // APR-3
	SlotSpacingCriteria *NM `hl7:"false,Slot Spacing Criteria"` // APR-4
	FillerOverrideCriteria []SCV `hl7:"false,Filler Override Criteria"` // APR-5
}

func (s *APR) SegmentName() string {
	return "APR"
}

// ARQ represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ARQ struct {
	PlacerAppointmentID *EI `hl7:"true,Placer Appointment ID"` // ARQ-1
	FillerAppointmentID *EI `hl7:"false,Filler Appointment ID"` // ARQ-2
	OccurrenceNumber *NM `hl7:"false,Occurrence Number"` // ARQ-3
	PlacerGroupNumber *EI `hl7:"false,Placer Group Number"` // ARQ-4
	ScheduleID *CE `hl7:"false,Schedule ID"` // ARQ-5
	RequestEventReason *CE `hl7:"false,Request Event Reason"` // ARQ-6
	AppointmentReason *CE `hl7:"false,Appointment Reason"` // ARQ-7
	AppointmentType *CE `hl7:"false,Appointment Type"` // ARQ-8
	AppointmentDuration *NM `hl7:"false,Appointment Duration"` // ARQ-9
	AppointmentDurationUnits *CE `hl7:"false,Appointment Duration Units"` // ARQ-10
	RequestedStartDateTimeRange []DR `hl7:"false,Requested Start Date/Time Range"` // ARQ-11
	Priority *ST `hl7:"false,Priority"` // ARQ-12
	RepeatingInterval *RI `hl7:"false,Repeating Interval"` // ARQ-13
	RepeatingIntervalDuration *ST `hl7:"false,Repeating Interval Duration"` // ARQ-14
	PlacerContactPerson *XCN `hl7:"false,Placer Contact Person"` // ARQ-15
	PlacerContactPhoneNumber *XTN `hl7:"false,Placer Contact Phone Number"` // ARQ-16
	PlacerContactAddress *XAD `hl7:"false,Placer Contact Address"` // ARQ-17
	PlacerContactLocation *PL `hl7:"false,Placer Contact Location"` // ARQ-18
	EnteredByPerson *XCN `hl7:"true,Entered By Person"` // ARQ-19
	EnteredByPhoneNumber []XTN `hl7:"false,Entered By Phone Number"` // ARQ-20
	EnteredByLocation *PL `hl7:"false,Entered By Location"` // ARQ-21
	ParentPlacerAppointmentID *EI `hl7:"false,Parent Placer Appointment ID"` // ARQ-22
	ParentFillerAppointmentID *EI `hl7:"false,Parent Filler Appointment ID"` // ARQ-23
}

func (s *ARQ) SegmentName() string {
	return "ARQ"
}

// AUT represents the corresponding HL7 segment.
// Definition from HL7 2.3
type AUT struct {
	AuthorizingPayorPlanCode *CE `hl7:"false,Authorizing Payor, Plan Code"` // AUT-1
	AuthorizingPayorCompanyID *CE `hl7:"true,Authorizing Payor, Company ID"` // AUT-2
	AuthorizingPayorCompanyName *ST `hl7:"false,Authorizing Payor, Company Name"` // AUT-3
	AuthorizationEffectiveDate *TS `hl7:"false,Authorization Effective Date"` // AUT-4
	AuthorizationExpirationDate *TS `hl7:"false,Authorization Expiration Date"` // AUT-5
	AuthorizationIdentifier *EI `hl7:"false,Authorization Identifier"` // AUT-6
	ReimbursementLimit *CP `hl7:"false,Reimbursement Limit"` // AUT-7
	RequestedNumberOfTreatments *NM `hl7:"false,Requested Number Of Treatments"` // AUT-8
	AuthorizedNumberOfTreatments *NM `hl7:"false,Authorized Number Of Treatments"` // AUT-9
	ProcessDate *TS `hl7:"false,Process Date"` // AUT-10
}

func (s *AUT) SegmentName() string {
	return "AUT"
}

// BHS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type BHS struct {
	BatchFieldSeparator *ST `hl7:"true,Batch Field Separator"` // BHS-1
	BatchEncodingCharacters *ST `hl7:"true,Batch Encoding Characters"` // BHS-2
	BatchSendingApplication *ST `hl7:"false,Batch Sending Application"` // BHS-3
	BatchSendingFacility *ST `hl7:"false,Batch Sending Facility"` // BHS-4
	BatchReceivingApplication *ST `hl7:"false,Batch Receiving Application"` // BHS-5
	BatchReceivingFacility *ST `hl7:"false,Batch Receiving Facility"` // BHS-6
	BatchCreationDateTime *TS `hl7:"false,Batch Creation Date/Time"` // BHS-7
	BatchSecurity *ST `hl7:"false,Batch Security"` // BHS-8
	BatchNameIDType *ST `hl7:"false,Batch Name/ID/Type"` // BHS-9
	BatchComment *ST `hl7:"false,Batch Comment"` // BHS-10
	BatchControlID *ST `hl7:"false,Batch Control ID"` // BHS-11
	ReferenceBatchControlID *ST `hl7:"false,Reference Batch Control ID"` // BHS-12
}

func (s *BHS) SegmentName() string {
	return "BHS"
}

// BLG represents the corresponding HL7 segment.
// Definition from HL7 2.3
type BLG struct {
	WhenToCharge *CM `hl7:"false,When To Charge"` // BLG-1
	ChargeType *ID `hl7:"false,Charge Type"` // BLG-2
	AccountID *CK `hl7:"false,Account ID"` // BLG-3
}

func (s *BLG) SegmentName() string {
	return "BLG"
}

// BTS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type BTS struct {
	BatchMessageCount *ST `hl7:"false,Batch Message Count"` // BTS-1
	BatchComment *ST `hl7:"false,Batch Comment"` // BTS-2
	BatchTotals []NM `hl7:"false,Batch Totals"` // BTS-3
}

func (s *BTS) SegmentName() string {
	return "BTS"
}

// CDM represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CDM struct {
	PrimaryKeyValue *CE `hl7:"true,Primary Key Value"` // CDM-1
	ChargeCodeAlias []CE `hl7:"false,Charge Code Alias"` // CDM-2
	ChargeDescriptionShort *ST `hl7:"true,Charge Description Short"` // CDM-3
	ChargeDescriptionLong *ST `hl7:"false,Charge Description Long"` // CDM-4
	DescriptionOverrideIndicator *ID `hl7:"false,Description Override Indicator"` // CDM-5
	ExplodingCharges []CE `hl7:"false,Exploding Charges"` // CDM-6
	ProcedureCode []CE `hl7:"false,Procedure Code"` // CDM-7
	ActiveInactiveFlag *ID `hl7:"false,Active/Inactive Flag"` // CDM-8
	InventoryNumber []CE `hl7:"false,Inventory Number"` // CDM-9
	ResourceLoad *NM `hl7:"false,Resource Load"` // CDM-10
	ContractNumber []CK `hl7:"false,Contract Number"` // CDM-11
	ContractOrganization *XON `hl7:"false,Contract Organization"` // CDM-12
	RoomFeeIndicator *ID `hl7:"false,Room Fee Indicator"` // CDM-13
}

func (s *CDM) SegmentName() string {
	return "CDM"
}

// CM0 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CM0 struct {
	CM0SetID *SI `hl7:"false,CM0 - Set ID"` // CM0-1
	SponsorStudyID *EI `hl7:"true,Sponsor Study ID"` // CM0-2
	AlternateStudyID []CE `hl7:"false,Alternate Study ID"` // CM0-3
	TitleOfStudy *ST `hl7:"true,Title Of Study"` // CM0-4
	ChairmanOfStudy *XCN `hl7:"false,Chairman Of Study"` // CM0-5
	LastIRBApprovalDate *DT `hl7:"false,Last IRB Approval Date"` // CM0-6
	TotalAccrualToDate *NM `hl7:"false,Total Accrual To Date"` // CM0-7
	LastAccrualDate *DT `hl7:"false,Last Accrual Date"` // CM0-8
	ContactForStudy *XCN `hl7:"false,Contact For Study"` // CM0-9
	ContactSTelNumber *XTN `hl7:"false,Contact'S Tel. Number"` // CM0-10
	ContactSAddress *XAD `hl7:"false,Contact'S Address"` // CM0-11
}

func (s *CM0) SegmentName() string {
	return "CM0"
}

// CM1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CM1 struct {
	CM1SetID *SI `hl7:"true,CM1 - Set ID"` // CM1-1
	StudyPhaseIdentifier *CE `hl7:"false,Study Phase Identifier"` // CM1-2
	DescriptionOfStudyPhase *ST `hl7:"true,Description Of Study Phase"` // CM1-3
}

func (s *CM1) SegmentName() string {
	return "CM1"
}

// CM2 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CM2 struct {
	CM2SetID *SI `hl7:"false,CM2 - Set ID"` // CM2-1
	ScheduledTimePoint *CE `hl7:"true,Scheduled Time Point"` // CM2-2
	DescriptionOfTimePoint *ST `hl7:"false,Description Of Time Point"` // CM2-3
	EventsScheduledThisTimePoint []CE `hl7:"true,Events Scheduled This Time Point"` // CM2-4
}

func (s *CM2) SegmentName() string {
	return "CM2"
}

// CSP represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CSP struct {
	StudyPhaseIdentifier *CE `hl7:"false,Study Phase Identifier"` // CSP-1
	DateTimeStudyPhaseBegan *TS `hl7:"true,Date/Time Study Phase Began"` // CSP-2
	DateTimeStudyPhaseEnded *TS `hl7:"false,Date/Time Study Phase Ended"` // CSP-3
	StudyPhaseEvaluability *CE `hl7:"false,Study Phase Evaluability"` // CSP-4
}

func (s *CSP) SegmentName() string {
	return "CSP"
}

// CSR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CSR struct {
	SponsorStudyID *EI `hl7:"true,Sponsor Study ID"` // CSR-1
	AlternateStudyID *EI `hl7:"false,Alternate Study ID"` // CSR-2
	InstitutionRegisteringThePatient *CE `hl7:"false,Institution Registering The Patient"` // CSR-3
	SponsorPatientID *CX `hl7:"true,Sponsor Patient ID"` // CSR-4
	AlternatePatientID *CX `hl7:"false,Alternate Patient ID"` // CSR-5
	DateTimeOfPatientStudyRegistration *TS `hl7:"false,Date/Time Of Patient Study Registration"` // CSR-6
	PersonPerformingStudyRegistration *XCN `hl7:"false,Person Performing Study Registration"` // CSR-7
	StudyAuthorizingProvider *XCN `hl7:"true,Study Authorizing Provider"` // CSR-8
	DateTimePatientStudyConsentSigned *TS `hl7:"false,Date/Time Patient Study Consent Signed"` // CSR-9
	PatientStudyEligibilityStatus *CE `hl7:"false,Patient Study Eligibility Status"` // CSR-10
	StudyRandomizationDateTime []TS `hl7:"false,Study Randomization Date/Time"` // CSR-11
	StudyRandomizedArm []CE `hl7:"false,Study Randomized Arm"` // CSR-12
	StratumForStudyRandomization []CE `hl7:"false,Stratum For Study Randomization"` // CSR-13
	PatientEvaluabilityStatus *CE `hl7:"false,Patient Evaluability Status"` // CSR-14
	DateTimeEndedStudy *TS `hl7:"false,Date/Time Ended Study"` // CSR-15
	ReasonEndedStudy *CE `hl7:"false,Reason Ended Study"` // CSR-16
}

func (s *CSR) SegmentName() string {
	return "CSR"
}

// CSS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CSS struct {
	StudyScheduledTimePoint *CE `hl7:"false,Study Scheduled Time Point"` // CSS-1
	StudyScheduledPatientTimePoint *TS `hl7:"false,Study Scheduled Patient Time Point"` // CSS-2
	StudyQualityControlCodes []CE `hl7:"false,Study Quality Control Codes"` // CSS-3
}

func (s *CSS) SegmentName() string {
	return "CSS"
}

// CTD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CTD struct {
	ContactRole *CE `hl7:"true,Contact Role"` // CTD-1
	ContactName []XPN `hl7:"false,Contact Name"` // CTD-2
	ContactAddress []XAD `hl7:"false,Contact Address"` // CTD-3
	ContactLocation *PL `hl7:"false,Contact Location"` // CTD-4
	ContactCommunicationInformation []XTN `hl7:"false,Contact Communication Information"` // CTD-5
	PreferredMethodOfContact *CE `hl7:"false,Preferred Method Of Contact"` // CTD-6
	ContactIdentifiers []CM `hl7:"false,Contact Identifiers"` // CTD-7
}

func (s *CTD) SegmentName() string {
	return "CTD"
}

// CTI represents the corresponding HL7 segment.
// Definition from HL7 2.3
type CTI struct {
	SponsorStudyID *EI `hl7:"true,Sponsor Study ID"` // CTI-1
	StudyPhaseIdentifier *CE `hl7:"false,Study Phase Identifier"` // CTI-2
	StudyScheduledTimePoint *CE `hl7:"false,Study Scheduled Time Point"` // CTI-3
}

func (s *CTI) SegmentName() string {
	return "CTI"
}

// DB1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type DB1 struct {
	SetIDDB1 *SI `hl7:"true,Set ID - DB1"` // DB1-1
	DisabledPersonCode *IS `hl7:"false,Disabled Person Code"` // DB1-2
	DisabledPersonIdentifier []CX `hl7:"false,Disabled Person Identifier"` // DB1-3
	DisabledIndicator *ID `hl7:"false,Disabled Indicator"` // DB1-4
	DisabilityStartDate *DT `hl7:"false,Disability Start Date"` // DB1-5
	DisabilityEndDate *DT `hl7:"false,Disability End Date"` // DB1-6
	DisabilityReturnToWorkDate *DT `hl7:"false,Disability Return To Work Date"` // DB1-7
	DisabilityUnableToWorkDate *DT `hl7:"false,Disability Unable To Work Date"` // DB1-8
}

func (s *DB1) SegmentName() string {
	return "DB1"
}

// DG1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type DG1 struct {
	SetIDDiagnosis *SI `hl7:"true,Set ID - Diagnosis"` // DG1-1
	DiagnosisCodingMethod *ID `hl7:"false,Diagnosis Coding Method"` // DG1-2
	DiagnosisCode *CE `hl7:"false,Diagnosis Code"` // DG1-3
	DiagnosisDescription *ST `hl7:"false,Diagnosis Description"` // DG1-4
	DiagnosisDateTime *TS `hl7:"false,Diagnosis Date/Time"` // DG1-5
	DiagnosisType *IS `hl7:"true,Diagnosis Type"` // DG1-6
	MajorDiagnosticCategory *CE `hl7:"false,Major Diagnostic Category"` // DG1-7
	DiagnosticRelatedGroup *CE `hl7:"false,Diagnostic Related Group"` // DG1-8
	DRGApprovalIndicator *ID `hl7:"false,DRG Approval Indicator"` // DG1-9
	DRGGrouperReviewCode *IS `hl7:"false,DRG Grouper Review Code"` // DG1-10
	OutlierType *CE `hl7:"false,Outlier Type"` // DG1-11
	OutlierDays *NM `hl7:"false,Outlier Days"` // DG1-12
	OutlierCost *CP `hl7:"false,Outlier Cost"` // DG1-13
	GrouperVersionAndType *ST `hl7:"false,Grouper Version And Type"` // DG1-14
	DiagnosisPriority *NM `hl7:"false,Diagnosis Priority"` // DG1-15
	DiagnosingClinician []XCN `hl7:"false,Diagnosing Clinician"` // DG1-16
	DiagnosisClassification *IS `hl7:"false,Diagnosis Classification"` // DG1-17
	ConfidentialIndicator *ID `hl7:"false,Confidential Indicator"` // DG1-18
	AttestationDateTime *TS `hl7:"false,Attestation Date/Time"` // DG1-19
}

func (s *DG1) SegmentName() string {
	return "DG1"
}

// DRG represents the corresponding HL7 segment.
// Definition from HL7 2.3
type DRG struct {
	DiagnosticRelatedGroup *CE `hl7:"false,Diagnostic Related Group"` // DRG-1
	DRGAssignedDateTime *TS `hl7:"false,DRG Assigned Date/Time"` // DRG-2
	DRGApprovalIndicator *ID `hl7:"false,DRG Approval Indicator"` // DRG-3
	DRGGrouperReviewCode *IS `hl7:"false,DRG Grouper Review Code"` // DRG-4
	OutlierType *CE `hl7:"false,Outlier Type"` // DRG-5
	OutlierDays *NM `hl7:"false,Outlier Days"` // DRG-6
	OutlierCost *CP `hl7:"false,Outlier Cost"` // DRG-7
	DRGPayor *IS `hl7:"false,DRG Payor"` // DRG-8
	OutlierReimbursement *CP `hl7:"false,Outlier Reimbursement"` // DRG-9
	ConfidentialIndicator *ID `hl7:"false,Confidential Indicator"` // DRG-10
}

func (s *DRG) SegmentName() string {
	return "DRG"
}

// DSC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type DSC struct {
	ContinuationPointer *ST `hl7:"false,Continuation Pointer"` // DSC-1
}

func (s *DSC) SegmentName() string {
	return "DSC"
}

// DSP represents the corresponding HL7 segment.
// Definition from HL7 2.3
type DSP struct {
	SetIDDisplayData *SI `hl7:"false,Set ID - Display Data"` // DSP-1
	DisplayLevel *SI `hl7:"false,Display Level"` // DSP-2
	DataLine *TX `hl7:"true,Data Line"` // DSP-3
	LogicalBreakPoint *ST `hl7:"false,Logical Break Point"` // DSP-4
	ResultID *TX `hl7:"false,Result ID"` // DSP-5
}

func (s *DSP) SegmentName() string {
	return "DSP"
}

// EQL represents the corresponding HL7 segment.
// Definition from HL7 2.3
type EQL struct {
	QueryTag *ST `hl7:"false,Query Tag"` // EQL-1
	QueryResponseFormatCode *ID `hl7:"true,Query/ Response Format Code"` // EQL-2
	EQLQueryName *CE `hl7:"true,EQL Query Name"` // EQL-3
	EQLQueryStatement *ST `hl7:"true,EQL Query Statement"` // EQL-4
}

func (s *EQL) SegmentName() string {
	return "EQL"
}

// ERQ represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ERQ struct {
	QueryTag *ST `hl7:"false,Query Tag"` // ERQ-1
	EventIdentifier *CE `hl7:"true,Event Identifier"` // ERQ-2
	InputParameterList []QIP `hl7:"false,Input Parameter List"` // ERQ-3
}

func (s *ERQ) SegmentName() string {
	return "ERQ"
}

// ERR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ERR struct {
	ErrorCodeAndLocation []CM `hl7:"true,Error Code And Location"` // ERR-1
}

func (s *ERR) SegmentName() string {
	return "ERR"
}

// EVN represents the corresponding HL7 segment.
// Definition from HL7 2.3
type EVN struct {
	EventTypeCode *ID `hl7:"true,Event Type Code"` // EVN-1
	RecordedDateTime *TS `hl7:"false,Recorded Date/Time"` // EVN-2
	DateTimePlannedEvent *TS `hl7:"false,Date/Time Planned Event"` // EVN-3
	EventReasonCode *ID `hl7:"false,Event Reason Code"` // EVN-4
	OperatorID *CN `hl7:"false,Operator ID"` // EVN-5
	EventOccured *TS `hl7:"false,Event Occured"` // EVN-6
}

func (s *EVN) SegmentName() string {
	return "EVN"
}

// FAC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type FAC struct {
	FacilityID []EI `hl7:"true,Facility ID"` // FAC-1
	FacilityType *ID `hl7:"false,Facility Type"` // FAC-2
	FacilityAddress *XAD `hl7:"true,Facility Address"` // FAC-3
	FacilityTelecommunication *XTN `hl7:"true,Facility Telecommunication"` // FAC-4
	ContactPerson []XCN `hl7:"false,Contact Person"` // FAC-5
	ContactTitle []ST `hl7:"false,Contact Title"` // FAC-6
	ContactAddress []XAD `hl7:"false,Contact Address"` // FAC-7
	ContactTelecommunication []XTN `hl7:"false,Contact Telecommunication"` // FAC-8
	SignatureAuthority *XCN `hl7:"true,Signature Authority"` // FAC-9
	SignatureAuthorityTitle *ST `hl7:"false,Signature Authority Title"` // FAC-10
	SignatureAuthorityAddress *XAD `hl7:"false,Signature Authority Address"` // FAC-11
	SignatureAuthorityTelecommunication *XTN `hl7:"false,Signature Authority Telecommunication"` // FAC-12
}

func (s *FAC) SegmentName() string {
	return "FAC"
}

// FHS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type FHS struct {
	FileFieldSeparator *ST `hl7:"true,File Field Separator"` // FHS-1
	FileEncodingCharacters *ST `hl7:"true,File Encoding Characters"` // FHS-2
	FileSendingApplication *ST `hl7:"false,File Sending Application"` // FHS-3
	FileSendingFacility *ST `hl7:"false,File Sending Facility"` // FHS-4
	FileReceivingApplication *ST `hl7:"false,File Receiving Application"` // FHS-5
	FileReceivingFacility *ST `hl7:"false,File Receiving Facility"` // FHS-6
	FileCreationDateTime *TS `hl7:"false,File Creation Date/Time"` // FHS-7
	FileSecurity *ST `hl7:"false,File Security"` // FHS-8
	FileNameID *ST `hl7:"false,File Name/ID"` // FHS-9
	FileHeaderComment *ST `hl7:"false,File Header Comment"` // FHS-10
	FileControlID *ST `hl7:"false,File Control ID"` // FHS-11
	ReferenceFileControlID *ST `hl7:"false,Reference File Control ID"` // FHS-12
}

func (s *FHS) SegmentName() string {
	return "FHS"
}

// FT1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type FT1 struct {
	SetIDFinancialTransaction *SI `hl7:"false,Set ID - Financial Transaction"` // FT1-1
	TransactionID *ST `hl7:"false,Transaction ID"` // FT1-2
	TransactionBatchID *ST `hl7:"false,Transaction Batch ID"` // FT1-3
	TransactionDate *TS `hl7:"true,Transaction Date"` // FT1-4
	TransactionPostingDate *TS `hl7:"false,Transaction Posting Date"` // FT1-5
	TransactionType *ID `hl7:"true,Transaction Type"` // FT1-6
	TransactionCode *CE `hl7:"true,Transaction Code"` // FT1-7
	TransactionDescription *ST `hl7:"false,Transaction Description"` // FT1-8
	TransactionDescriptionAlternate *ST `hl7:"false,Transaction Description - Alternate"` // FT1-9
	TransactionQuantity *NM `hl7:"false,Transaction Quantity"` // FT1-10
	TransactionAmountExtended *CP `hl7:"false,Transaction Amount - Extended"` // FT1-11
	TransactionAmountUnit *CP `hl7:"false,Transaction Amount - Unit"` // FT1-12
	DepartmentCode *CE `hl7:"false,Department Code"` // FT1-13
	InsurancePlanID *CE `hl7:"false,Insurance Plan ID"` // FT1-14
	InsuranceAmount *CP `hl7:"false,Insurance Amount"` // FT1-15
	AssignedPatientLocation *PL `hl7:"false,Assigned Patient Location"` // FT1-16
	FeeSchedule *ID `hl7:"false,Fee Schedule"` // FT1-17
	PatientType *ID `hl7:"false,Patient Type"` // FT1-18
	DiagnosisCode []CE `hl7:"false,Diagnosis Code"` // FT1-19
	PerformedByCode *XCN `hl7:"false,Performed By Code"` // FT1-20
	OrderedByCode *XCN `hl7:"false,Ordered By Code"` // FT1-21
	UnitCost *NM `hl7:"false,Unit Cost"` // FT1-22
	FillerOrderNumber *EI `hl7:"false,Filler Order Number"` // FT1-23
	EnteredByCode *XCN `hl7:"false,Entered By Code"` // FT1-24
	ProcedureCode *CE `hl7:"false,Procedure Code"` // FT1-25
}

func (s *FT1) SegmentName() string {
	return "FT1"
}

// FTS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type FTS struct {
	FileBatchCount *NM `hl7:"false,File Batch Count"` // FTS-1
	FileTrailerComment *ST `hl7:"false,File Trailer Comment"` // FTS-2
}

func (s *FTS) SegmentName() string {
	return "FTS"
}

// GOL represents the corresponding HL7 segment.
// Definition from HL7 2.3
type GOL struct {
	ActionCode *ID `hl7:"true,Action Code"` // GOL-1
	ActionDateTime *TS `hl7:"true,Action Date/Time"` // GOL-2
	GoalID *CE `hl7:"true,Goal ID"` // GOL-3
	GoalInstanceID *EI `hl7:"true,Goal Instance ID"` // GOL-4
	EpisodeOfCareID *EI `hl7:"false,Episode Of Care ID"` // GOL-5
	GoalListPriority *NM `hl7:"false,Goal List Priority"` // GOL-6
	GoalEstablishedDateTime *TS `hl7:"false,Goal Established Date/Time"` // GOL-7
	ExpectedGoalAchievementDateTime *TS `hl7:"false,Expected Goal Achievement Date/Time"` // GOL-8
	GoalClassification *CE `hl7:"false,Goal Classification"` // GOL-9
	GoalManagementDiscipline *CE `hl7:"false,Goal Management Discipline"` // GOL-10
	CurrentGoalReviewStatus *CE `hl7:"false,Current Goal Review Status"` // GOL-11
	CurrentGoalReviewDateTime *TS `hl7:"false,Current Goal Review Date/Time"` // GOL-12
	NextGoalReviewDateTime *TS `hl7:"false,Next Goal Review Date/Time"` // GOL-13
	PreviousGoalReviewDateTime *TS `hl7:"false,Previous Goal Review Date/Time"` // GOL-14
	GoalReviewInterval *TQ `hl7:"false,Goal Review Interval"` // GOL-15
	GoalEvaluation *CE `hl7:"false,Goal Evaluation"` // GOL-16
	GoalEvaluationComment []ST `hl7:"false,Goal Evaluation Comment"` // GOL-17
	GoalLifeCycleStatus *CE `hl7:"false,Goal Life Cycle Status"` // GOL-18
	GoalLifeCycleStatusDateTime *TS `hl7:"false,Goal Life Cycle Status Date/Time"` // GOL-19
	GoalTargetType []CE `hl7:"false,Goal Target Type"` // GOL-20
	GoalTargetName []XPN `hl7:"false,Goal Target Name"` // GOL-21
}

func (s *GOL) SegmentName() string {
	return "GOL"
}

// GT1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type GT1 struct {
	SetIDGuarantor *SI `hl7:"true,Set ID - Guarantor"` // GT1-1
	GuarantorNumber []CX `hl7:"false,Guarantor Number"` // GT1-2
	GuarantorName []XPN `hl7:"true,Guarantor Name"` // GT1-3
	GuarantorSpouseName []XPN `hl7:"false,Guarantor Spouse Name"` // GT1-4
	GuarantorAddress []XAD `hl7:"false,Guarantor Address"` // GT1-5
	GuarantorPhNumHome []XTN `hl7:"false,Guarantor Ph Num- Home"` // GT1-6
	GuarantorPhNumBusiness []XTN `hl7:"false,Guarantor Ph Num-Business"` // GT1-7
	GuarantorDateTimeOfBirth *TS `hl7:"false,Guarantor Date/Time Of Birth"` // GT1-8
	GuarantorSex *IS `hl7:"false,Guarantor Sex"` // GT1-9
	GuarantorType *IS `hl7:"false,Guarantor Type"` // GT1-10
	GuarantorRelationship *IS `hl7:"false,Guarantor Relationship"` // GT1-11
	GuarantorSSN *ST `hl7:"false,Guarantor SSN"` // GT1-12
	GuarantorDateBegin *DT `hl7:"false,Guarantor Date - Begin"` // GT1-13
	GuarantorDateEnd *DT `hl7:"false,Guarantor Date - End"` // GT1-14
	GuarantorPriority *NM `hl7:"false,Guarantor Priority"` // GT1-15
	GuarantorEmployerName []XPN `hl7:"false,Guarantor Employer Name"` // GT1-16
	GuarantorEmployerAddress []XAD `hl7:"false,Guarantor Employer Address"` // GT1-17
	GuarantorEmployPhoneNumber []XTN `hl7:"false,Guarantor Employ Phone Number"` // GT1-18
	GuarantorEmployeeIDNumber []CX `hl7:"false,Guarantor Employee ID Number"` // GT1-19
	GuarantorEmploymentStatus *IS `hl7:"false,Guarantor Employment Status"` // GT1-20
	GuarantorOrganization []XON `hl7:"false,Guarantor Organization"` // GT1-21
	GuarantorBillingHoldFlag *ID `hl7:"false,Guarantor Billing Hold Flag"` // GT1-22
	GuarantorCreditRatingCode *CE `hl7:"false,Guarantor Credit Rating Code"` // GT1-23
	GuarantorDeathDateAndTime *TS `hl7:"false,Guarantor Death Date And Time"` // GT1-24
	GuarantorDeathFlag *ID `hl7:"false,Guarantor Death Flag"` // GT1-25
	GuarantorChargeAdjustmentCode *CE `hl7:"false,Guarantor Charge Adjustment Code"` // GT1-26
	GuarantorHouseholdAnnualIncome *CP `hl7:"false,Guarantor Household Annual Income"` // GT1-27
	GuarantorHouseholdSize *NM `hl7:"false,Guarantor Household Size"` // GT1-28
	GuarantorEmployerIDNumber []CX `hl7:"false,Guarantor Employer ID Number"` // GT1-29
	GuarantorMaritalStatusCode *IS `hl7:"false,Guarantor Marital Status Code"` // GT1-30
	GuarantorHireEffectiveDate *DT `hl7:"false,Guarantor Hire Effective Date"` // GT1-31
	EmploymentStopDate *DT `hl7:"false,Employment Stop Date"` // GT1-32
	LivingDependency *IS `hl7:"false,Living Dependency"` // GT1-33
	AmbulatoryStatus *IS `hl7:"false,Ambulatory Status"` // GT1-34
	Citizenship *IS `hl7:"false,Citizenship"` // GT1-35
	PrimaryLanguage *CE `hl7:"false,Primary Language"` // GT1-36
	LivingArrangement *IS `hl7:"false,Living Arrangement"` // GT1-37
	PublicityIndicator *CE `hl7:"false,Publicity Indicator"` // GT1-38
	ProtectionIndicator *ID `hl7:"false,Protection Indicator"` // GT1-39
	StudentIndicator *IS `hl7:"false,Student Indicator"` // GT1-40
	Religion *IS `hl7:"false,Religion"` // GT1-41
	MotherSMaidenName *XPN `hl7:"false,Mother’s Maiden Name"` // GT1-42
	NationalityCode *CE `hl7:"false,Nationality Code"` // GT1-43
	EthnicGroup *IS `hl7:"false,Ethnic Group"` // GT1-44
	ContactPersonSName []XPN `hl7:"false,Contact Person'S Name"` // GT1-45
	ContactPersonSTelephoneNumber []XTN `hl7:"false,Contact Person’s Telephone Number"` // GT1-46
	ContactReason *CE `hl7:"false,Contact Reason"` // GT1-47
	ContactRelationshipCode *IS `hl7:"false,Contact Relationship Code"` // GT1-48
	JobTitle *ST `hl7:"false,Job Title"` // GT1-49
	JobCodeClass *JCC `hl7:"false,Job Code/Class"` // GT1-50
	GuarantorEmployerSOrganizationName []XON `hl7:"false,Guarantor Employer'S Organization Name"` // GT1-51
	Handicap *IS `hl7:"false,Handicap"` // GT1-52
	JobStatus *IS `hl7:"false,Job Status"` // GT1-53
	GuarantorFinancialClass *FC `hl7:"false,Guarantor Financial Class"` // GT1-54
	GuarantorRace *IS `hl7:"false,Guarantor Race"` // GT1-55
}

func (s *GT1) SegmentName() string {
	return "GT1"
}

// IN1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type IN1 struct {
	SetIDInsurance *SI `hl7:"true,Set ID - Insurance"` // IN1-1
	InsurancePlanID *CE `hl7:"false,Insurance Plan ID"` // IN1-2
	InsuranceCompanyID *CX `hl7:"true,Insurance Company ID"` // IN1-3
	InsuranceCompanyName *XON `hl7:"false,Insurance Company Name"` // IN1-4
	InsuranceCompanyAddress *XAD `hl7:"false,Insurance Company Address"` // IN1-5
	InsuranceCoContactPpers *XPN `hl7:"false,Insurance Co. Contact Ppers"` // IN1-6
	InsuranceCoPhoneNumber []XTN `hl7:"false,Insurance Co Phone Number"` // IN1-7
	GroupNumber *ST `hl7:"false,Group Number"` // IN1-8
	GroupName *XON `hl7:"false,Group Name"` // IN1-9
	InsuredSGroupEmployerID *CX `hl7:"false,Insured'S Group Employer ID"` // IN1-10
	InsuredSGroupEmpName *XON `hl7:"false,Insured'S Group Emp Name"` // IN1-11
	PlanEffectiveDate *DT `hl7:"false,Plan Effective Date"` // IN1-12
	PlanExpirationDate *DT `hl7:"false,Plan Expiration Date"` // IN1-13
	AuthorizationInformation *CM `hl7:"false,Authorization Information"` // IN1-14
	PlanType *IS `hl7:"false,Plan Type"` // IN1-15
	NameOfInsured *XPN `hl7:"false,Name Of Insured"` // IN1-16
	InsuredSRelationshipToPatient *IS `hl7:"false,Insured'S Relationship To Patient"` // IN1-17
	InsuredSDateOfBirth *TS `hl7:"false,Insured'S Date Of Birth"` // IN1-18
	InsuredSAddress *XAD `hl7:"false,Insured'S Address"` // IN1-19
	AssignmentOfBenefits *IS `hl7:"false,Assignment Of Benefits"` // IN1-20
	CoordinationOfBenefits *IS `hl7:"false,Coordination Of Benefits"` // IN1-21
	CoordOfBenPriority *ST `hl7:"false,Coord Of Ben. Priority"` // IN1-22
	NoticeOfAdmissionCode *ID `hl7:"false,Notice Of Admission Code"` // IN1-23
	NoticeOfAdmissionDate *DT `hl7:"false,Notice Of Admission Date"` // IN1-24
	RptOfEigibilityCode *ID `hl7:"false,Rpt Of Eigibility Code"` // IN1-25
	RptOfEligibilityDate *DT `hl7:"false,Rpt Of Eligibility Date"` // IN1-26
	ReleaseInformationCode *IS `hl7:"false,Release Information Code"` // IN1-27
	PreAdmitCert *ST `hl7:"false,Pre-Admit Cert"` // IN1-28
	VerificationDateTime *TS `hl7:"false,Verification Date/Time"` // IN1-29
	VerificationBy *XPN `hl7:"false,Verification By"` // IN1-30
	TypeOfAgreementCode *IS `hl7:"false,Type Of Agreement Code"` // IN1-31
	BillingStatus *IS `hl7:"false,Billing Status"` // IN1-32
	LifetimeReserveDays *NM `hl7:"false,Lifetime Reserve Days"` // IN1-33
	DelayBeforeLifetimeReserveDays *NM `hl7:"false,Delay Before Lifetime Reserve Days"` // IN1-34
	CompanyPlanCode *IS `hl7:"false,Company Plan Code"` // IN1-35
	PolicyNumber *ST `hl7:"false,Policy Number"` // IN1-36
	PolicyDeductible *CP `hl7:"false,Policy Deductible"` // IN1-37
	PolicyLimitAmount *CP `hl7:"false,Policy Limit - Amount"` // IN1-38
	PolicyLimitDays *NM `hl7:"false,Policy Limit - Days"` // IN1-39
	RoomRateSemiPrivate *CP `hl7:"false,Room Rate - Semi-Private"` // IN1-40
	RoomRatePrivate *CP `hl7:"false,Room Rate - Private"` // IN1-41
	InsuredSEmploymentStatus *CE `hl7:"false,Insured'S Employment Status"` // IN1-42
	InsuredSSex *IS `hl7:"false,Insured'S Sex"` // IN1-43
	InsuredSEmployerAddress *XAD `hl7:"false,Insured'S Employer Address"` // IN1-44
	VerificationStatus *ST `hl7:"false,Verification Status"` // IN1-45
	PriorInsurancePlanID *IS `hl7:"false,Prior Insurance Plan ID"` // IN1-46
	CoverageType *IS `hl7:"false,Coverage Type"` // IN1-47
	Handicap *IS `hl7:"false,Handicap"` // IN1-48
	InsuredSIDNumber *CX `hl7:"false,Insured'S ID Number"` // IN1-49
}

func (s *IN1) SegmentName() string {
	return "IN1"
}

// IN2 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type IN2 struct {
	InsuredSEmployeeID *CX `hl7:"false,Insured'S Employee ID"` // IN2-1
	InsuredSSocialSecurityNumber *ST `hl7:"false,Insured'S Social Security Number"` // IN2-2
	InsuredSEmployerName *XCN `hl7:"false,Insured'S Employer Name"` // IN2-3
	EmployerInformationData *IS `hl7:"false,Employer Information Data"` // IN2-4
	MailClaimParty *IS `hl7:"false,Mail Claim Party"` // IN2-5
	MedicareHealthInsCardNumber *ST `hl7:"false,Medicare Health Ins Card Number"` // IN2-6
	MedicaidCaseName *XPN `hl7:"false,Medicaid Case Name"` // IN2-7
	MedicaidCaseNumber *ST `hl7:"false,Medicaid Case Number"` // IN2-8
	ChampusSponsorName *XPN `hl7:"false,Champus Sponsor Name"` // IN2-9
	ChampusIDNumber *ST `hl7:"false,Champus ID Number"` // IN2-10
	DependentOfChampusRecipient *CE `hl7:"false,Dependent Of Champus Recipient"` // IN2-11
	ChampusOrganization *ST `hl7:"false,Champus Organization"` // IN2-12
	ChampusStation *ST `hl7:"false,Champus Station"` // IN2-13
	ChampusService *IS `hl7:"false,Champus Service"` // IN2-14
	ChampusRankGrade *IS `hl7:"false,Champus Rank/Grade"` // IN2-15
	ChampusStatus *IS `hl7:"false,Champus Status"` // IN2-16
	ChampusRetireDate *DT `hl7:"false,Champus Retire Date"` // IN2-17
	ChampusNonAvailCertOnFile *ID `hl7:"false,Champus Non-Avail Cert On File"` // IN2-18
	BabyCoverage *ID `hl7:"false,Baby Coverage"` // IN2-19
	CombineBabyBill *ID `hl7:"false,Combine Baby Bill"` // IN2-20
	BloodDeductible *ST `hl7:"false,Blood Deductible"` // IN2-21
	SpecialCoverageApprovalName *XPN `hl7:"false,Special Coverage Approval Name"` // IN2-22
	SpecialCoverageApprovalTitle *ST `hl7:"false,Special Coverage Approval Title"` // IN2-23
	NonCoveredInsuranceCode []ST `hl7:"false,Non-Covered Insurance Code"` // IN2-24
	PayorID *CX `hl7:"false,Payor ID"` // IN2-25
	PayorSubscriberID *CX `hl7:"false,Payor Subscriber ID"` // IN2-26
	EligibilitySource *IS `hl7:"false,Eligibility Source"` // IN2-27
	RoomCoverageTypeAmount []CM `hl7:"false,Room Coverage Type/Amount"` // IN2-28
	PolicyTypeAmount []CM `hl7:"false,Policy Type/Amount"` // IN2-29
	DailyDeductible *CM `hl7:"false,Daily Deductible"` // IN2-30
	LivingDependency *IS `hl7:"false,Living Dependency"` // IN2-31
	AmbulatoryStatus *IS `hl7:"false,Ambulatory Status"` // IN2-32
	Citizenship *IS `hl7:"false,Citizenship"` // IN2-33
	PrimaryLanguage *CE `hl7:"false,Primary Language"` // IN2-34
	LivingArrangement *IS `hl7:"false,Living Arrangement"` // IN2-35
	PublicityIndicator *CE `hl7:"false,Publicity Indicator"` // IN2-36
	ProtectionIndicator *ID `hl7:"false,Protection Indicator"` // IN2-37
	StudentIndicator *IS `hl7:"false,Student Indicator"` // IN2-38
	Religion *IS `hl7:"false,Religion"` // IN2-39
	MotherSMaidenName *XPN `hl7:"false,Mother’s Maiden Name"` // IN2-40
	NationalityCode *CE `hl7:"false,Nationality Code"` // IN2-41
	EthnicGroup *IS `hl7:"false,Ethnic Group"` // IN2-42
	MaritalStatus []IS `hl7:"false,Marital Status"` // IN2-43
	EmploymentStartDate *DT `hl7:"false,Employment Start Date"` // IN2-44
	EmploymentStopDate *DT `hl7:"false,Employment Stop Date"` // IN2-45
	JobTitle *ST `hl7:"false,Job Title"` // IN2-46
	JobCodeClass *JCC `hl7:"false,Job Code/Class"` // IN2-47
	JobStatus *IS `hl7:"false,Job Status"` // IN2-48
	EmployerContactPersonName []XPN `hl7:"false,Employer Contact Person Name"` // IN2-49
	EmployerContactPersonPhoneNumber []XTN `hl7:"false,Employer Contact Person Phone Number"` // IN2-50
	EmployerContactReason *IS `hl7:"false,Employer Contact Reason"` // IN2-51
	InsuredSContactPersonSName []XPN `hl7:"false,Insured’s Contact Person’s Name"` // IN2-52
	InsuredSContactPersonTelephoneNumber []XTN `hl7:"false,Insured’s Contact Person Telephone Number"` // IN2-53
	InsuredSContactPersonReason []IS `hl7:"false,Insured’s Contact Person Reason"` // IN2-54
	RelationshipToThePatientStartDate *DT `hl7:"false,Relationship To The Patient Start Date"` // IN2-55
	RelationshipToThePatientStopDate []DT `hl7:"false,Relationship To The Patient Stop Date"` // IN2-56
	InsuranceCoContactReason *IS `hl7:"false,Insurance Co. Contact Reason"` // IN2-57
	InsuranceCoContactPhoneNumber *XTN `hl7:"false,Insurance Co. Contact Phone Number"` // IN2-58
	PolicyScope *IS `hl7:"false,Policy Scope"` // IN2-59
	PolicySource *IS `hl7:"false,Policy Source"` // IN2-60
	PatientMemberNumber *CX `hl7:"false,Patient Member Number"` // IN2-61
	GuarantorSRelationshipToInsured *IS `hl7:"false,Guarantor’s Relationship To Insured"` // IN2-62
	InsuredSTelephoneNumberHome []XTN `hl7:"false,Insured’s Telephone Number - Home"` // IN2-63
	InsuredSEmployerTelephoneNumber []XTN `hl7:"false,Insured’s Employer Telephone Number"` // IN2-64
	MilitaryHandicappedProgram *CE `hl7:"false,Military Handicapped Program"` // IN2-65
	SuspendFlag *ID `hl7:"false,Suspend Flag"` // IN2-66
	CoPayLimitFlag *ID `hl7:"false,Co-Pay Limit Flag"` // IN2-67
	StoplossLimitFlag *ID `hl7:"false,Stoploss Limit Flag"` // IN2-68
	InsuredOrganizationNameAndID []XON `hl7:"false,Insured Organization Name And ID"` // IN2-69
	InsuredEmployerOrganizationNameAndID []XON `hl7:"false,Insured Employer Organization Name And ID"` // IN2-70
	Race *IS `hl7:"false,Race"` // IN2-71
	PatientRelationshipToInsured *ID `hl7:"false,Patient Relationship To Insured"` // IN2-72
}

func (s *IN2) SegmentName() string {
	return "IN2"
}

// IN3 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type IN3 struct {
	SetIDInsuranceCertification *SI `hl7:"true,Set ID - Insurance Certification"` // IN3-1
	CertificationNumber *CX `hl7:"false,Certification Number"` // IN3-2
	CertifiedBy []XCN `hl7:"false,Certified By"` // IN3-3
	CertificationRequired *ID `hl7:"false,Certification Required"` // IN3-4
	Penalty *CM `hl7:"false,Penalty"` // IN3-5
	CertificationDateTime *TS `hl7:"false,Certification Date/Time"` // IN3-6
	CertificationModifyDateTime *TS `hl7:"false,Certification Modify Date/Time"` // IN3-7
	Operator []XCN `hl7:"false,Operator"` // IN3-8
	CertificationBeginDate *DT `hl7:"false,Certification Begin Date"` // IN3-9
	CertificationEndDate *DT `hl7:"false,Certification End Date"` // IN3-10
	Days *CM `hl7:"false,Days"` // IN3-11
	NonConcurCodeDescription *CE `hl7:"false,Non-Concur Code/Description"` // IN3-12
	NonConcurEffectiveDateTime *TS `hl7:"false,Non-Concur Effective Date/Time"` // IN3-13
	PhysicianReviewer []XCN `hl7:"false,Physician Reviewer"` // IN3-14
	CertificationContact *ST `hl7:"false,Certification Contact"` // IN3-15
	CertificationContactPhoneNumber []XTN `hl7:"false,Certification Contact Phone Number"` // IN3-16
	AppealReason *CE `hl7:"false,Appeal Reason"` // IN3-17
	CertificationAgency *CE `hl7:"false,Certification Agency"` // IN3-18
	CertificationAgencyPhoneNumber []XTN `hl7:"false,Certification Agency Phone Number"` // IN3-19
	PreCertificationRequiredWindow []CM `hl7:"false,Pre-Certification Required/Window"` // IN3-20
	CaseManager *ST `hl7:"false,Case Manager"` // IN3-21
	SecondOpinionDate *DT `hl7:"false,Second Opinion Date"` // IN3-22
	SecondOpinionStatus *IS `hl7:"false,Second Opinion Status"` // IN3-23
	SecondOpinionDocumentationReceived []IS `hl7:"false,Second Opinion Documentation Received"` // IN3-24
	SecondOpinionPhysician []XCN `hl7:"false,Second Opinion Physician"` // IN3-25
}

func (s *IN3) SegmentName() string {
	return "IN3"
}

// LCC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type LCC struct {
	PrimaryKeyValue *PL `hl7:"true,Primary Key Value"` // LCC-1
	LocationDepartment *IS `hl7:"true,Location Department"` // LCC-2
	AccommodationType []CE `hl7:"false,Accommodation Type"` // LCC-3
	ChargeCode []CE `hl7:"true,Charge Code"` // LCC-4
}

func (s *LCC) SegmentName() string {
	return "LCC"
}

// LCH represents the corresponding HL7 segment.
// Definition from HL7 2.3
type LCH struct {
	PrimaryKeyValue *PL `hl7:"true,Primary Key Value"` // LCH-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // LCH-2
	SegmentUniqueKey *EI `hl7:"false,Segment Unique Key"` // LCH-3
	LocationCharacteristicID *CE `hl7:"true,Location Characteristic ID"` // LCH-4
	LocationCharacteristicValue *CE `hl7:"true,Location Characteristic Value"` // LCH-5
}

func (s *LCH) SegmentName() string {
	return "LCH"
}

// LDP represents the corresponding HL7 segment.
// Definition from HL7 2.3
type LDP struct {
	LDPPrimaryKeyValue *PL `hl7:"true,LDP Primary Key Value"` // LDP-1
	LocationDepartment *IS `hl7:"true,Location Department"` // LDP-2
	LocationService []IS `hl7:"false,Location Service"` // LDP-3
	SpecialityType []CE `hl7:"false,Speciality Type"` // LDP-4
	ValidPatientClasses []ID `hl7:"false,Valid Patient Classes"` // LDP-5
	ActiveInactiveFlag *ID `hl7:"false,Active/Inactive Flag"` // LDP-6
	ActivationDate *TS `hl7:"false,Activation Date"` // LDP-7
	InactivationDateLDP *TS `hl7:"false,Inactivation Date - LDP"` // LDP-8
	InactivatedReason *ST `hl7:"false,Inactivated Reason"` // LDP-9
	VisitingHours []VH `hl7:"false,Visiting Hours"` // LDP-10
	ContactPhone *XTN `hl7:"false,Contact Phone"` // LDP-11
}

func (s *LDP) SegmentName() string {
	return "LDP"
}

// LOC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type LOC struct {
	PrimaryKeyValue *PL `hl7:"true,Primary Key Value"` // LOC-1
	LocationDescription *ST `hl7:"false,Location Description"` // LOC-2
	LocationType []IS `hl7:"true,Location Type"` // LOC-3
	OrganizationName *XON `hl7:"false,Organization Name"` // LOC-4
	LocationAddress *XAD `hl7:"false,Location Address"` // LOC-5
	LocationPhone []XTN `hl7:"false,Location Phone"` // LOC-6
	LicenseNumber []CE `hl7:"false,License Number"` // LOC-7
	LocationEquipment []ID `hl7:"false,Location Equipment"` // LOC-8
}

func (s *LOC) SegmentName() string {
	return "LOC"
}

// LRL represents the corresponding HL7 segment.
// Definition from HL7 2.3
type LRL struct {
	PrimaryKeyValue *PL `hl7:"true,Primary Key Value"` // LRL-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // LRL-2
	SegmentUniqueKey *EI `hl7:"false,Segment Unique Key"` // LRL-3
	LocationRelationshipID *CE `hl7:"false,Location Relationship ID"` // LRL-4
	OrganizationalLocationRelationshipValue *XON `hl7:"false,Organizational Location Relationship Value"` // LRL-5
	PatientLocationRelationshipValue *PL `hl7:"false,Patient Location Relationship Value"` // LRL-6
}

func (s *LRL) SegmentName() string {
	return "LRL"
}

// MFA represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MFA struct {
	RecordLevelEventCode *ID `hl7:"true,Record-Level Event Code"` // MFA-1
	MFNControlID *ST `hl7:"false,MFN Control ID"` // MFA-2
	EventCompletionDateTime *TS `hl7:"false,Event Completion Date/Time"` // MFA-3
	ErrorReturnCodeAndOrText *CE `hl7:"true,Error Return Code And/Or Text"` // MFA-4
	PrimaryKeyValue []CE `hl7:"true,Primary Key Value"` // MFA-5
}

func (s *MFA) SegmentName() string {
	return "MFA"
}

// MFE represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MFE struct {
	RecordLevelEventCode *ID `hl7:"true,Record-Level Event Code"` // MFE-1
	MFNControlID *ST `hl7:"false,MFN Control ID"` // MFE-2
	EffectiveDateTime *TS `hl7:"false,Effective Date/Time"` // MFE-3
	PrimaryKeyValue []CE `hl7:"true,Primary Key Value"` // MFE-4
}

func (s *MFE) SegmentName() string {
	return "MFE"
}

// MFI represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MFI struct {
	MasterFileIdentifier *CE `hl7:"true,Master File Identifier"` // MFI-1
	MasterFileApplicationIdentifier *HD `hl7:"false,Master File Application Identifier"` // MFI-2
	FileLevelEventCode *ID `hl7:"true,File-Level Event Code"` // MFI-3
	EnteredDateTime *TS `hl7:"false,Entered Date/Time"` // MFI-4
	EffectiveDateTime *TS `hl7:"false,Effective Date/Time"` // MFI-5
	ResponseLevelCode *ID `hl7:"true,Response Level Code"` // MFI-6
}

func (s *MFI) SegmentName() string {
	return "MFI"
}

// MRG represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MRG struct {
	PriorPatientIDInternal []CX `hl7:"true,Prior Patient ID - Internal"` // MRG-1
	PriorAlternatePatientID []CX `hl7:"false,Prior Alternate Patient ID"` // MRG-2
	PriorPatientAccountNumber *CX `hl7:"false,Prior Patient Account Number"` // MRG-3
	PriorPatientIDExternal *CX `hl7:"false,Prior Patient ID - External"` // MRG-4
	PriorVisitNumber *CX `hl7:"false,Prior Visit Number"` // MRG-5
	PriorAlternateVisitID *CX `hl7:"false,Prior Alternate Visit ID"` // MRG-6
	PriorPatientName *CX `hl7:"false,Prior Patient Name"` // MRG-7
}

func (s *MRG) SegmentName() string {
	return "MRG"
}

// MSA represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MSA struct {
	AcknowledgementCode *ID `hl7:"true,Acknowledgement Code"` // MSA-1
	MessageControlID *ST `hl7:"true,Message Control ID"` // MSA-2
	TextMessage *ST `hl7:"false,Text Message"` // MSA-3
	ExpectedSequenceNumber *NM `hl7:"false,Expected Sequence Number"` // MSA-4
	DelayedAcknowledgementType *ID `hl7:"false,Delayed Acknowledgement Type"` // MSA-5
	ErrorCondition *CE `hl7:"false,Error Condition"` // MSA-6
}

func (s *MSA) SegmentName() string {
	return "MSA"
}

// MSH represents the corresponding HL7 segment.
// Definition from HL7 2.3
type MSH struct {
	// Missing: MSH.1
	EncodingCharacters *Delimiters `hl7:"true,Encoding Characters"` // MSH-2
	SendingApplication *HD `hl7:"false,Sending Application"` // MSH-3
	SendingFacility *HD `hl7:"false,Sending Facility"` // MSH-4
	ReceivingApplication *HD `hl7:"false,Receiving Application"` // MSH-5
	ReceivingFacility *HD `hl7:"false,Receiving Facility"` // MSH-6
	DateTimeOfMessage *TS `hl7:"false,Date / Time Of Message"` // MSH-7
	Security *ST `hl7:"false,Security"` // MSH-8
	MessageType *CM `hl7:"true,Message Type"` // MSH-9
	MessageControlID *ST `hl7:"true,Message Control ID"` // MSH-10
	ProcessingID *PT `hl7:"true,Processing ID"` // MSH-11
	VersionID *ID `hl7:"true,Version ID"` // MSH-12
	SequenceNumber *NM `hl7:"false,Sequence Number"` // MSH-13
	ContinuationPointer *ST `hl7:"false,Continuation Pointer"` // MSH-14
	AcceptAcknowledgementType *ID `hl7:"false,Accept Acknowledgement Type"` // MSH-15
	ApplicationAcknowledgementType *ID `hl7:"false,Application Acknowledgement Type"` // MSH-16
	CountryCode *ID `hl7:"false,Country Code"` // MSH-17
	CharacterSet *ID `hl7:"false,Character Set"` // MSH-18
	PrincipalLanguageOfMessage *CE `hl7:"false,Principal Language Of Message"` // MSH-19
}

func (s *MSH) SegmentName() string {
	return "MSH"
}

// NCK represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NCK struct {
	SystemDateTime *TS `hl7:"false,System Date/Time"` // NCK-1
}

func (s *NCK) SegmentName() string {
	return "NCK"
}

// NK1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NK1 struct {
	SetIDNextOfKin *SI `hl7:"true,Set ID - Next Of Kin"` // NK1-1
	Name []XPN `hl7:"false,Name"` // NK1-2
	Relationship *CE `hl7:"false,Relationship"` // NK1-3
	Address []XAD `hl7:"false,Address"` // NK1-4
	PhoneNumber []XTN `hl7:"false,Phone Number"` // NK1-5
	BusinessPhoneNumber []XTN `hl7:"false,Business Phone Number"` // NK1-6
	ContactRole *CE `hl7:"false,Contact Role"` // NK1-7
	StartDate *DT `hl7:"false,Start Date"` // NK1-8
	EndDate *DT `hl7:"false,End Date"` // NK1-9
	NextOfKinAssociatedPartiesJobTitle *ST `hl7:"false,Next Of Kin/Associated Parties Job Title"` // NK1-10
	NextOfKinJobAssociatedPartiesCodeClass *JCC `hl7:"false,Next Of Kin Job/Associated Parties Code/Class"` // NK1-11
	NextOfKinAssociatedPartiesEmployeeNumber *CX `hl7:"false,Next Of Kin/Associated Parties Employee Number"` // NK1-12
	OrganizationName []XON `hl7:"false,Organization Name"` // NK1-13
	MaritalStatus []IS `hl7:"false,Marital Status"` // NK1-14
	Sex *IS `hl7:"false,Sex"` // NK1-15
	DateOfBirth *TS `hl7:"false,Date Of Birth"` // NK1-16
	LivingDependency *IS `hl7:"false,Living Dependency"` // NK1-17
	AmbulatoryStatus *IS `hl7:"false,Ambulatory Status"` // NK1-18
	Citizenship *IS `hl7:"false,Citizenship"` // NK1-19
	PrimaryLanguage *CE `hl7:"false,Primary Language"` // NK1-20
	LivingArrangement *IS `hl7:"false,Living Arrangement"` // NK1-21
	PublicityIndicator *CE `hl7:"false,Publicity Indicator"` // NK1-22
	ProtectionIndicator *ID `hl7:"false,Protection Indicator"` // NK1-23
	StudentIndicator *IS `hl7:"false,Student Indicator"` // NK1-24
	Religion *IS `hl7:"false,Religion"` // NK1-25
	MotherSMaidenName *XPN `hl7:"false,Mother’s Maiden Name"` // NK1-26
	NationalityCode *CE `hl7:"false,Nationality Code"` // NK1-27
	EthnicGroup *IS `hl7:"false,Ethnic Group"` // NK1-28
	ContactReason *CE `hl7:"false,Contact Reason"` // NK1-29
	ContactPersonSName []XPN `hl7:"false,Contact Person'S Name"` // NK1-30
	ContactPersonSTelephoneNumber []XTN `hl7:"false,Contact Person’s Telephone Number"` // NK1-31
	ContactPersonSAddress []XAD `hl7:"false,Contact Person’s Address"` // NK1-32
	AssociatedPartySIdentifiers []CX `hl7:"false,Associated Party’s Identifiers"` // NK1-33
	JobStatus *IS `hl7:"false,Job Status"` // NK1-34
	Race *IS `hl7:"false,Race"` // NK1-35
	Handicap *IS `hl7:"false,Handicap"` // NK1-36
	ContactPersonSocialSecurityNumber *ST `hl7:"false,Contact Person Social Security Number"` // NK1-37
}

func (s *NK1) SegmentName() string {
	return "NK1"
}

// NPU represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NPU struct {
	BedLocation *PL `hl7:"true,Bed Location"` // NPU-1
	BedStatus *IS `hl7:"false,Bed Status"` // NPU-2
}

func (s *NPU) SegmentName() string {
	return "NPU"
}

// NSC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NSC struct {
	NetworkChangeType *ID `hl7:"false,Network Change Type"` // NSC-1
	CurrentCPU *ST `hl7:"false,Current CPU"` // NSC-2
	CurrentFileserver *ST `hl7:"false,Current Fileserver"` // NSC-3
	CurrentApplication *ST `hl7:"false,Current Application"` // NSC-4
	CurrentFacility *ST `hl7:"false,Current Facility"` // NSC-5
	NewCPU *ST `hl7:"false,New CPU"` // NSC-6
	NewFileserver *ST `hl7:"false,New Fileserver"` // NSC-7
	NewApplication *ST `hl7:"false,New Application"` // NSC-8
}

func (s *NSC) SegmentName() string {
	return "NSC"
}

// NST represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NST struct {
	StatisticsAvailable *ID `hl7:"false,Statistics Available"` // NST-1
	SourceIdentifier *ST `hl7:"false,Source Identifier"` // NST-2
	SourceType *ID `hl7:"false,Source Type"` // NST-3
	StatisticsStart *TS `hl7:"false,Statistics Start"` // NST-4
	StatisticsEnd *TS `hl7:"false,Statistics End"` // NST-5
	ReceiveCharacterCount *NM `hl7:"false,Receive Character Count"` // NST-6
	SendCharacterCount *NM `hl7:"false,Send Character Count"` // NST-7
	MessagesReceived *NM `hl7:"false,Messages Received"` // NST-8
	MessagesSent *NM `hl7:"false,Messages Sent"` // NST-9
	ChecksumErrorsReceived *NM `hl7:"false,Checksum Errors Received"` // NST-10
	LengthErrorsReceived *NM `hl7:"false,Length Errors Received"` // NST-11
	OtherErrorsReceived *NM `hl7:"false,Other Errors Received"` // NST-12
	ConnectTimeouts *NM `hl7:"false,Connect Timeouts"` // NST-13
	ReceiveTimeouts *NM `hl7:"false,Receive Timeouts"` // NST-14
	NetworkErrors *NM `hl7:"false,Network Errors"` // NST-15
}

func (s *NST) SegmentName() string {
	return "NST"
}

// NTE represents the corresponding HL7 segment.
// Definition from HL7 2.3
type NTE struct {
	SetIDNotesAndComments *SI `hl7:"false,Set ID - Notes And Comments"` // NTE-1
	SourceOfComment *ID `hl7:"false,Source Of Comment"` // NTE-2
	Comment []FT `hl7:"false,Comment"` // NTE-3
}

func (s *NTE) SegmentName() string {
	return "NTE"
}

// OBR represents the corresponding HL7 segment.
// Definition from HL7 SYNTHETIC
type OBR struct {
	SetIDOBR *SI `hl7:"false,Set ID - OBR"` // OBR-1
	PlacerOrderNumber *EI `hl7:"false,Placer Order Number"` // OBR-2
	FillerOrderNumber *EI `hl7:"false,Filler Order Number"` // OBR-3
	UniversalServiceIdentifier *CWE `hl7:"true,Universal Service Identifier"` // OBR-4
	Priority *ID `hl7:"false,Priority"` // OBR-5
	RequestedDateTime *TS `hl7:"false,Requested Date/Time"` // OBR-6
	ObservationDateTime *TS `hl7:"false,Observation Date/Time"` // OBR-7
	ObservationEndDateTime *TS `hl7:"false,Observation End Date/Time"` // OBR-8
	CollectionVolume *CQ `hl7:"false,Collection Volume"` // OBR-9
	CollectorIdentifier []XCN `hl7:"false,Collector Identifier"` // OBR-10
	SpecimenActionCode *ID `hl7:"false,Specimen Action Code"` // OBR-11
	DangerCode *CE `hl7:"false,Danger Code"` // OBR-12
	RelevantClinicalInformation *ST `hl7:"false,Relevant Clinical Information"` // OBR-13
	SpecimenReceivedDateTime *TS `hl7:"false,Specimen Received Date/Time"` // OBR-14
	SpecimenSource *SPS `hl7:"false,Specimen Source"` // OBR-15
	OrderingProvider []XCN `hl7:"false,Ordering Provider"` // OBR-16
	OrderCallbackPhoneNumber []XTN `hl7:"false,Order Callback Phone Number"` // OBR-17
	PlacerField1 *ST `hl7:"false,Placer Field 1"` // OBR-18
	PlacerField2 *ST `hl7:"false,Placer Field 2"` // OBR-19
	FillerField1 *ST `hl7:"false,Filler Field 1"` // OBR-20
	FillerField2 *ST `hl7:"false,Filler Field 2"` // OBR-21
	ResultsRptStatusChngDateTime *TS `hl7:"false,Results Rpt/Status Chng - Date/Time"` // OBR-22
	ChargeToPractice *MOC `hl7:"false,Charge To Practice"` // OBR-23
	DiagnosticServSectID *ID `hl7:"false,Diagnostic Serv Sect ID"` // OBR-24
	ResultStatus *ID `hl7:"false,Result Status"` // OBR-25
	ParentResult *PRL `hl7:"false,Parent Result"` // OBR-26
	QuantityTiming []TQ `hl7:"false,Quantity/Timing"` // OBR-27
	ResultCopiesTo []XCN `hl7:"false,Result Copies To"` // OBR-28
	Parent *EIP `hl7:"false,Parent"` // OBR-29
	TransportationMode *ID `hl7:"false,Transportation Mode"` // OBR-30
	ReasonForStudy []CE `hl7:"false,Reason For Study"` // OBR-31
	PrincipalResultInterpreter *NDL `hl7:"false,Principal Result Interpreter"` // OBR-32
	AssistantResultInterpreter []NDL `hl7:"false,Assistant Result Interpreter"` // OBR-33
	Technician []NDL `hl7:"false,Technician"` // OBR-34
	Transcriptionist []NDL `hl7:"false,Transcriptionist"` // OBR-35
	ScheduledDateTime *TS `hl7:"false,Scheduled Date/Time"` // OBR-36
	NumberOfSampleContainers *NM `hl7:"false,Number Of Sample Containers *"` // OBR-37
	TransportLogisticsOfCollectedSample []CE `hl7:"false,Transport Logistics Of Collected Sample"` // OBR-38
	CollectorSComment []CE `hl7:"false,Collector'S Comment *"` // OBR-39
	TransportArrangementResponsibility *CE `hl7:"false,Transport Arrangement Responsibility"` // OBR-40
	TransportArranged *ID `hl7:"false,Transport Arranged"` // OBR-41
	EscortRequired *ID `hl7:"false,Escort Required"` // OBR-42
	PlannedPatientTransportComment []CE `hl7:"false,Planned Patient Transport Comment"` // OBR-43
	ProcedureCode *CE `hl7:"false,Procedure Code"` // OBR-44
	ProcedureCodeModifier []CE `hl7:"false,Procedure Code Modifier"` // OBR-45
	PlacerSupplementalServiceInformation []CE `hl7:"false,Placer Supplemental Service Information"` // OBR-46
	FillerSupplementalServiceInformation []CE `hl7:"false,Filler Supplemental Service Information"` // OBR-47
	MedicallyNecessaryDuplicateProcedureReason *CWE `hl7:"false,Medically Necessary Duplicate Procedure Reason."` // OBR-48
	ResultHandling *IS `hl7:"false,Result Handling"` // OBR-49
	ParentUniversalServiceIdentifier *CWE `hl7:"false,Parent Universal Service Identifier"` // OBR-50
}

func (s *OBR) SegmentName() string {
	return "OBR"
}

// OBX represents the corresponding HL7 segment.
// Definition from HL7 SYNTHETIC
type OBX struct {
	SetIDOBX *SI `hl7:"false,Set ID - OBX"` // OBX-1
	ValueType *ID `hl7:"false,Value Type"` // OBX-2
	ObservationIdentifier *CWE `hl7:"true,Observation Identifier"` // OBX-3
	ObservationSubID *ST `hl7:"false,Observation Sub-ID"` // OBX-4
	ObservationValue []Any `hl7:"false,Observation Value"` // OBX-5
	Units *CE `hl7:"false,Units"` // OBX-6
	ReferencesRange *ST `hl7:"false,References Range"` // OBX-7
	AbnormalFlags []IS `hl7:"false,Abnormal Flags"` // OBX-8
	Probability *NM `hl7:"false,Probability"` // OBX-9
	NatureOfAbnormalTest []ID `hl7:"false,Nature Of Abnormal Test"` // OBX-10
	ObservationResultStatus *ID `hl7:"true,Observation Result Status"` // OBX-11
	EffectiveDateOfReferenceRangeValues *TS `hl7:"false,Effective Date Of Reference Range Values"` // OBX-12
	UserDefinedAccessChecks *ST `hl7:"false,User Defined Access Checks"` // OBX-13
	DateTimeOfTheObservation *TS `hl7:"false,Date/Time Of The Observation"` // OBX-14
	ProducerSReference *CE `hl7:"false,Producer'S Reference"` // OBX-15
	ResponsibleObserver []XCN `hl7:"false,Responsible Observer"` // OBX-16
	ObservationMethod []CE `hl7:"false,Observation Method"` // OBX-17
	EquipmentInstanceIdentifier []EI `hl7:"false,Equipment Instance Identifier"` // OBX-18
	DateTimeOfTheAnalysis *TS `hl7:"false,Date/Time Of The Analysis"` // OBX-19
	ReservedForHarmonizationWithVersion26A []XON `hl7:"false,Reserved For Harmonization With Version 2.6 A"` // OBX-20
	ReservedForHarmonizationWithVersion26B []XON `hl7:"false,Reserved For Harmonization With Version 2.6 B"` // OBX-21
	ReservedForHarmonizationWithVersion26C []XON `hl7:"false,Reserved For Harmonization With Version 2.6 C"` // OBX-22
	PerformingOrganizationName *XON `hl7:"false,Performing Organization Name"` // OBX-23
	PerformingOrganizationAddress *XAD `hl7:"false,Performing Organization Address"` // OBX-24
	PerformingOrganizationMedicalDirector *XCN `hl7:"false,Performing Organization Medical Director"` // OBX-25
}

func (s *OBX) SegmentName() string {
	return "OBX"
}

// ODS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ODS struct {
	Type *ID `hl7:"true,Type"` // ODS-1
	ServicePeriod []CE `hl7:"false,Service Period"` // ODS-2
	DietSupplementOrPreferenceCode []CE `hl7:"true,Diet, Supplement, Or Preference Code"` // ODS-3
	TextInstruction *ST `hl7:"false,Text Instruction"` // ODS-4
}

func (s *ODS) SegmentName() string {
	return "ODS"
}

// ODT represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ODT struct {
	TrayType *CE `hl7:"true,Tray Type"` // ODT-1
	ServicePeriod []CE `hl7:"false,Service Period"` // ODT-2
	TextInstruction *ST `hl7:"false,Text Instruction"` // ODT-3
}

func (s *ODT) SegmentName() string {
	return "ODT"
}

// OM1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM1 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM1-1
	ProducerSTestObservationID *CE `hl7:"true,Producer'S Test/Observation ID"` // OM1-2
	PermittedDataTypes []ID `hl7:"false,Permitted Data Types"` // OM1-3
	SpecimenRequired *ID `hl7:"true,Specimen Required"` // OM1-4
	ProducerID *CE `hl7:"true,Producer ID"` // OM1-5
	ObservationDescription *CE `hl7:"false,Observation Description"` // OM1-6
	OtherTestObservationIDsForTheObservation *CE `hl7:"false,Other Test/Observation IDs For The Observation"` // OM1-7
	OtherNames []ST `hl7:"true,Other Names"` // OM1-8
	PreferredReportNameForTheObservation *ST `hl7:"false,Preferred Report Name For The Observation"` // OM1-9
	PreferredShortNameOrMnemonicForObservation *ST `hl7:"false,Preferred Short Name Or Mnemonic For Observation"` // OM1-10
	PreferredLongNameForTheObservation *ST `hl7:"false,Preferred Long Name For The Observation"` // OM1-11
	Orderability *ID `hl7:"false,Orderability"` // OM1-12
	IdentityOfInstrumentUsedToPerfromThisStudy []CE `hl7:"false,Identity Of Instrument Used To Perfrom This Study"` // OM1-13
	CodedRepresentationOfMethod *CE `hl7:"false,Coded Representation Of Method"` // OM1-14
	Portable *ID `hl7:"false,Portable"` // OM1-15
	ObservationProducingDepartmentSection []CE `hl7:"false,Observation Producing Department/Section"` // OM1-16
	TelephoneNumberOfSection *TN `hl7:"false,Telephone Number Of Section"` // OM1-17
	NatureOfTestObservation *ID `hl7:"false,Nature Of Test/Observation"` // OM1-18
	ReportSubheader *CE `hl7:"false,Report Subheader"` // OM1-19
	ReportDisplayOrder *ST `hl7:"false,Report Display Order"` // OM1-20
	DateTimeStampForAnyChangeInDefAttriForObs *TS `hl7:"false,Date/Time Stamp For Any Change In Def Attri For Obs"` // OM1-21
	EffectiveDateTimeOfChangeInTestProcThatMakeResultsNonComparable *TS `hl7:"false,Effective Date/Time Of Change In Test Proc. That Make Results Non-Comparable"` // OM1-22
	TypicalTurnAroundTime *NM `hl7:"false,Typical Turn-Around Time"` // OM1-23
	ProcessingTime *NM `hl7:"false,Processing Time"` // OM1-24
	ProcessingPriority []ID `hl7:"false,Processing Priority"` // OM1-25
	ReportingPriority *ID `hl7:"false,Reporting Priority"` // OM1-26
	OutsideSiteSWhereObservationMayBePerformed []CE `hl7:"false,Outside Site(S) Where Observation May Be Performed"` // OM1-27
	AddressOfOutsideSiteS *AD `hl7:"false,Address Of Outside Site(S)"` // OM1-28
	PhoneNumberOfOutsideSite *TN `hl7:"false,Phone Number Of Outside Site"` // OM1-29
	ConfidentialityCode *ID `hl7:"false,Confidentiality Code"` // OM1-30
	ObservationsRequiredToInterpretTheObservation *CE `hl7:"false,Observations Required To Interpret The Observation"` // OM1-31
	InterpretationOfObservations *TX `hl7:"false,Interpretation Of Observations"` // OM1-32
	ContraindicationsToObservations *CE `hl7:"false,Contraindications To Observations"` // OM1-33
	ReflexTestsObservations []CE `hl7:"false,Reflex Tests/Observations"` // OM1-34
	RulesThatTriggerReflexTesting *ST `hl7:"false,Rules That Trigger Reflex Testing"` // OM1-35
	FixedCannedMessage *CE `hl7:"false,Fixed Canned Message"` // OM1-36
	PatientPreparation *TX `hl7:"false,Patient Preparation"` // OM1-37
	ProcedureMedication *CE `hl7:"false,Procedure Medication"` // OM1-38
	FactorsThatMayEffectTheObservation *TX `hl7:"false,Factors That May Effect The Observation"` // OM1-39
	TestObservationPerformanceSchedule []ST `hl7:"false,Test/Observation Performance Schedule"` // OM1-40
	DescriptionOfTestMethods *TX `hl7:"false,Description Of Test Methods"` // OM1-41
	KindOfQuantityObserved *CE `hl7:"false,Kind Of Quantity Observed"` // OM1-42
	PointVersusInterval *CE `hl7:"false,Point Versus Interval"` // OM1-43
	ChallengeInformation *TX `hl7:"false,Challenge Information"` // OM1-44
	RelationshipModifier *CE `hl7:"false,Relationship Modifier"` // OM1-45
	TargetAnatomicSiteOfTest *CE `hl7:"false,Target Anatomic Site Of Test"` // OM1-46
	ModalityOfImagingMeasurement *CE `hl7:"false,Modality Of Imaging Measurement"` // OM1-47
}

func (s *OM1) SegmentName() string {
	return "OM1"
}

// OM2 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM2 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM2-1
	UnitsOfMeasure *CE `hl7:"false,Units Of Measure"` // OM2-2
	RangeOfDecimalPrecision []NM `hl7:"false,Range Of Decimal Precision"` // OM2-3
	CorrespondingSIUnitsOfMeasure *CE `hl7:"false,Corresponding SI Units Of Measure"` // OM2-4
	SIConversionFactor *TX `hl7:"false,SI Conversion Factor"` // OM2-5
	ReferenceNormalRangeOrdinalContinuousObs *CM `hl7:"false,Reference (Normal) Range - Ordinal & Continuous Obs"` // OM2-6
	CriticalRangeForOrdinalContinuousObs *CM `hl7:"false,Critical Range For Ordinal & Continuous Obs"` // OM2-7
	AbsoluteRangeForOrdinalContinuousObs *CM `hl7:"false,Absolute Range For Ordinal & Continuous Obs"` // OM2-8
	DeltaCheckCriteria []CM `hl7:"false,Delta Check Criteria"` // OM2-9
	MinimumMeaningfulIncrements *NM `hl7:"false,Minimum Meaningful Increments"` // OM2-10
}

func (s *OM2) SegmentName() string {
	return "OM2"
}

// OM3 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM3 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM3-1
	PreferredCodingSystem *CE `hl7:"false,Preferred Coding System"` // OM3-2
	ValidCodedAnswers *CE `hl7:"false,Valid Coded Answers"` // OM3-3
	NormalTextCodesForCategoricalObservations []CE `hl7:"false,Normal Text/Codes For Categorical Observations"` // OM3-4
	AbnormalTextCodesForCategoricalObservations *CE `hl7:"false,Abnormal Text/Codes For Categorical Observations"` // OM3-5
	CriticalTextCodesForCategoricalObservations *CE `hl7:"false,Critical Text Codes For Categorical Observations"` // OM3-6
	ValueType *ID `hl7:"true,Value Type"` // OM3-7
}

func (s *OM3) SegmentName() string {
	return "OM3"
}

// OM4 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM4 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM4-1
	DerivedSpecimen *ID `hl7:"false,Derived Specimen"` // OM4-2
	ContainerDescription *TX `hl7:"false,Container Description"` // OM4-3
	ContainerVolume *NM `hl7:"false,Container Volume"` // OM4-4
	ContainerUnits *CE `hl7:"false,Container Units"` // OM4-5
	Specimen *CE `hl7:"false,Specimen"` // OM4-6
	Additive *CE `hl7:"false,Additive"` // OM4-7
	Preparation *TX `hl7:"false,Preparation"` // OM4-8
	SpecialHandlingRequirements *TX `hl7:"false,Special Handling Requirements"` // OM4-9
	NormalCollectionVolume *CQ `hl7:"false,Normal Collection Volume"` // OM4-10
	MinimumCollectionVolume *CQ `hl7:"false,Minimum Collection Volume"` // OM4-11
	SpecimenRequirements *TX `hl7:"false,Specimen Requirements"` // OM4-12
	SpecimenPriorities *ID `hl7:"false,Specimen Priorities"` // OM4-13
	SpecimenRetentionTime *CQ `hl7:"false,Specimen Retention Time"` // OM4-14
}

func (s *OM4) SegmentName() string {
	return "OM4"
}

// OM5 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM5 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM5-1
	TestObservationsIncludedWAnOrderedTestBattery []CE `hl7:"false,Test/Observations Included W/An Ordered Test Battery"` // OM5-2
	ObservationIDSuffixes *ST `hl7:"false,Observation ID Suffixes"` // OM5-3
}

func (s *OM5) SegmentName() string {
	return "OM5"
}

// OM6 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type OM6 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM6-1
	DerivationRule *TX `hl7:"false,Derivation Rule"` // OM6-2
}

func (s *OM6) SegmentName() string {
	return "OM6"
}

// ORC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ORC struct {
	OrderControl *ID `hl7:"true,Order Control"` // ORC-1
	PlacerOrderNumber []EI `hl7:"false,Placer Order Number"` // ORC-2
	FillerOrderNumber *EI `hl7:"false,Filler Order Number"` // ORC-3
	PlacerGroupNumber *EI `hl7:"false,Placer Group Number"` // ORC-4
	OrderStatus *ID `hl7:"false,Order Status"` // ORC-5
	ResponseFlag *ID `hl7:"false,Response Flag"` // ORC-6
	QuantityTiming *TQ `hl7:"true,Quantity/Timing"` // ORC-7
	Parent *CM `hl7:"false,Parent"` // ORC-8
	DateTimeOfTransaction *TS `hl7:"false,Date/Time Of Transaction"` // ORC-9
	EnteredBy *XCN `hl7:"false,Entered By"` // ORC-10
	VerifiedBy *XCN `hl7:"false,Verified By"` // ORC-11
	OrderingProvider []XCN `hl7:"false,Ordering Provider"` // ORC-12
	EntererSLocation *PL `hl7:"false,Enterer'S Location"` // ORC-13
	CallBackPhoneNumber []TN `hl7:"false,Call Back Phone Number"` // ORC-14
	OrderEffectiveDateTime *TS `hl7:"false,Order Effective Date/Time"` // ORC-15
	OrderControlCodeReason *CE `hl7:"false,Order Control Code Reason"` // ORC-16
	EnteringOrganization *CE `hl7:"false,Entering Organization"` // ORC-17
	EnteringDevice *CE `hl7:"false,Entering Device"` // ORC-18
	ActionBy *XCN `hl7:"false,Action By"` // ORC-19
}

func (s *ORC) SegmentName() string {
	return "ORC"
}

// ORO represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ORO struct {
	ORDERITEMID *CE `hl7:"false,ORDER ITEM ID"` // ORO-1
	SUBSTITUTEALLOWED *ID `hl7:"false,SUBSTITUTE ALLOWED"` // ORO-2
	RESULTSCOPIESTO []CN `hl7:"false,RESULTS COPIES TO"` // ORO-3
	STOCKLOCATION *ID `hl7:"false,STOCK LOCATION"` // ORO-4
}

func (s *ORO) SegmentName() string {
	return "ORO"
}

// PCR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PCR struct {
	ImplicatedProduct *CE `hl7:"true,Implicated Product"` // PCR-1
	GenericProduct *IS `hl7:"false,Generic Product"` // PCR-2
	ProductClass *CE `hl7:"false,Product Class"` // PCR-3
	TotalDurationOfTherapy *CQ `hl7:"false,Total Duration Of Therapy"` // PCR-4
	ProductManufactureDate *TS `hl7:"false,Product Manufacture Date"` // PCR-5
	ProductExpirationDate *TS `hl7:"false,Product Expiration Date"` // PCR-6
	ProductImplantationDate *TS `hl7:"false,Product Implantation Date"` // PCR-7
	ProductExplantationDate *TS `hl7:"false,Product Explantation Date"` // PCR-8
	SingleUseDevice *IS `hl7:"false,Single Use Device"` // PCR-9
	IndicationForProductUse *CE `hl7:"false,Indication For Product Use"` // PCR-10
	ProductProblem *IS `hl7:"false,Product Problem"` // PCR-11
	ProductSerialLotNumber []ST `hl7:"false,Product Serial/Lot Number"` // PCR-12
	ProductAvailableForInspection *IS `hl7:"false,Product Available For Inspection"` // PCR-13
	ProductEvaluationPerformed *CE `hl7:"false,Product Evaluation Performed"` // PCR-14
	ProductEvaluationStatus *CE `hl7:"false,Product Evaluation Status"` // PCR-15
	ProductEvaluationResults *CE `hl7:"false,Product Evaluation Results"` // PCR-16
	EvaluatedProductSource *ID `hl7:"false,Evaluated Product Source"` // PCR-17
	DateProductReturnedToManufacturer *TS `hl7:"false,Date Product Returned To Manufacturer"` // PCR-18
	DeviceOperatorQualifications *ID `hl7:"false,Device Operator Qualifications"` // PCR-19
	RelatednessAssessment *ID `hl7:"false,Relatedness Assessment"` // PCR-20
	ActionTakenInResponseToTheEvent []ID `hl7:"false,Action Taken In Response To The Event"` // PCR-21
	EventCausalityObservations []ID `hl7:"false,Event Causality Observations"` // PCR-22
	IndirectExposureMechanism []ID `hl7:"false,Indirect Exposure Mechanism"` // PCR-23
}

func (s *PCR) SegmentName() string {
	return "PCR"
}

// PD1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PD1 struct {
	LivingDependency *IS `hl7:"false,Living Dependency"` // PD1-1
	LivingArrangement *IS `hl7:"false,Living Arrangement"` // PD1-2
	PatientPrimaryFacility []XON `hl7:"false,Patient Primary Facility"` // PD1-3
	PatientPrimaryCareProviderNameIDNo []XCN `hl7:"false,Patient Primary Care Provider Name & ID No."` // PD1-4
	StudentIndicator *IS `hl7:"false,Student Indicator"` // PD1-5
	Handicap *IS `hl7:"false,Handicap"` // PD1-6
	LivingWill *IS `hl7:"false,Living Will"` // PD1-7
	OrganDonor *IS `hl7:"false,Organ Donor"` // PD1-8
	SeparateBill *ID `hl7:"false,Separate Bill"` // PD1-9
	DuplicatePatient []CX `hl7:"false,Duplicate Patient"` // PD1-10
	PublicityIndicator *CE `hl7:"false,Publicity Indicator"` // PD1-11
	ProtectionIndicator *ID `hl7:"false,Protection Indicator"` // PD1-12
}

func (s *PD1) SegmentName() string {
	return "PD1"
}

// PDC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PDC struct {
	ManufacturerDistributor *XON `hl7:"true,Manufacturer/Distributor"` // PDC-1
	Country *CE `hl7:"true,Country"` // PDC-2
	BrandName *ST `hl7:"true,Brand Name"` // PDC-3
	DeviceFamilyName *ST `hl7:"false,Device Family Name"` // PDC-4
	GenericName *CE `hl7:"false,Generic Name"` // PDC-5
	ModelIdentifier []ST `hl7:"false,Model Identifier"` // PDC-6
	CatalogueIdentifier *ST `hl7:"false,Catalogue Identifier"` // PDC-7
	OtherIdentifier []ST `hl7:"false,Other Identifier"` // PDC-8
	ProductCode *CE `hl7:"false,Product Code"` // PDC-9
	MarketingBasis *ID `hl7:"false,Marketing Basis"` // PDC-10
	MarketingApprovalIdentifier *ST `hl7:"false,Marketing Approval Identifier"` // PDC-11
	LabeledShelfLife *CQ `hl7:"false,Labeled Shelf Life"` // PDC-12
	ExpectedShelfLife *CQ `hl7:"false,Expected Shelf Life"` // PDC-13
	DateFirstMarked *TS `hl7:"false,Date First Marked"` // PDC-14
	DateLastMarked *TS `hl7:"false,Date Last Marked"` // PDC-15
}

func (s *PDC) SegmentName() string {
	return "PDC"
}

// PEO represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PEO struct {
	EventIdentifiersUsed []CE `hl7:"false,Event Identifiers Used"` // PEO-1
	EventSymptomDiagnosisCode []CE `hl7:"false,Event Symptom/Diagnosis Code"` // PEO-2
	EventOnsetDateTime *TS `hl7:"true,Event Onset Date/Time"` // PEO-3
	EventExacerbationDateTime *TS `hl7:"false,Event Exacerbation Date/Time"` // PEO-4
	EventImprovedDateTime *TS `hl7:"false,Event Improved Date/Time"` // PEO-5
	EventEndedDataTime *TS `hl7:"false,Event Ended Data/Time"` // PEO-6
	EventLocationOccurredAddress *XAD `hl7:"false,Event Location Occurred Address"` // PEO-7
	EventQualification []ID `hl7:"false,Event Qualification"` // PEO-8
	EventSerious *ID `hl7:"false,Event Serious"` // PEO-9
	EventExpected *ID `hl7:"false,Event Expected"` // PEO-10
	EventOutcome []ID `hl7:"false,Event Outcome"` // PEO-11
	PatientOutcome *ID `hl7:"false,Patient Outcome"` // PEO-12
	EventDescriptionFromOthers []FT `hl7:"false,Event Description From Others"` // PEO-13
	EventFromOriginalReporter []FT `hl7:"false,Event From Original Reporter"` // PEO-14
	EventDescriptionFromPatient []FT `hl7:"false,Event Description From Patient"` // PEO-15
	EventDescriptionFromPractitioner []FT `hl7:"false,Event Description From Practitioner"` // PEO-16
	EventDescriptionFromAutopsy []FT `hl7:"false,Event Description From Autopsy"` // PEO-17
	CauseOfDeath []CE `hl7:"false,Cause Of Death"` // PEO-18
	PrimaryObserverName *XPN `hl7:"false,Primary Observer Name"` // PEO-19
	PrimaryObserverAddress []XAD `hl7:"false,Primary Observer Address"` // PEO-20
	PrimaryObserverTelephone []XTN `hl7:"false,Primary Observer Telephone"` // PEO-21
	PrimaryObserverSQualification *ID `hl7:"false,Primary Observer’s Qualification"` // PEO-22
	ConfirmationProvidedBy *ID `hl7:"false,Confirmation Provided By"` // PEO-23
	PrimaryObserverAwareDateTime *TS `hl7:"false,Primary Observer Aware Date/Time"` // PEO-24
	PrimaryObserverSIdentityMayBeDivulged *ID `hl7:"false,Primary Observer’s Identity May Be Divulged"` // PEO-25
}

func (s *PEO) SegmentName() string {
	return "PEO"
}

// PES represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PES struct {
	SenderOrganizationName *XON `hl7:"false,Sender Organization Name"` // PES-1
	SenderIndividualName []XCN `hl7:"false,Sender Individual Name"` // PES-2
	SenderAddress []XAD `hl7:"false,Sender Address"` // PES-3
	SenderTelephone []XTN `hl7:"false,Sender Telephone"` // PES-4
	SenderEventIdentifier *EI `hl7:"false,Sender Event Identifier"` // PES-5
	SenderSequenceNumber *NM `hl7:"false,Sender Sequence Number"` // PES-6
	SenderEventDescription []FT `hl7:"false,Sender Event Description"` // PES-7
	SenderComment *FT `hl7:"false,Sender Comment"` // PES-8
	SenderAwareDateTime *TS `hl7:"false,Sender Aware Date/Time"` // PES-9
	EventReportDate *TS `hl7:"true,Event Report Date"` // PES-10
	EventReportTimingType []ID `hl7:"false,Event Report Timing/Type"` // PES-11
	EventReportSource *ID `hl7:"false,Event Report Source"` // PES-12
	EventReportedTo []ID `hl7:"false,Event Reported To"` // PES-13
}

func (s *PES) SegmentName() string {
	return "PES"
}

// PID represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PID struct {
	SetIDPatientID *SI `hl7:"false,Set ID - Patient ID"` // PID-1
	PatientIDExternalID *CX `hl7:"false,Patient ID (External ID)"` // PID-2
	PatientIDInternalID []CX `hl7:"true,Patient ID (Internal ID)"` // PID-3
	AlternatePatientID *CX `hl7:"false,Alternate Patient ID"` // PID-4
	PatientName *XPN `hl7:"true,Patient Name"` // PID-5
	MotherSMaidenName *XPN `hl7:"false,Mother'S Maiden Name"` // PID-6
	DateOfBirth *TS `hl7:"false,Date Of Birth"` // PID-7
	Sex *IS `hl7:"false,Sex"` // PID-8
	PatientAlias []XPN `hl7:"false,Patient Alias"` // PID-9
	Race *CWE `hl7:"false,Race"` // PID-10
	PatientAddress []XAD `hl7:"false,Patient Address"` // PID-11
	CountyCode *IS `hl7:"false,County Code"` // PID-12
	PhoneNumberHome []XTN `hl7:"false,Phone Number - Home"` // PID-13
	PhoneNumberBusiness []XTN `hl7:"false,Phone Number - Business"` // PID-14
	PrimaryLanguage *CE `hl7:"false,Primary Language"` // PID-15
	MaritalStatus []IS `hl7:"false,Marital Status"` // PID-16
	Religion *IS `hl7:"false,Religion"` // PID-17
	PatientAccountNumber *CX `hl7:"false,Patient Account Number"` // PID-18
	SSNNumberPatient *ST `hl7:"false,SSN Number - Patient"` // PID-19
	DriverSLicenseNumber *DLN `hl7:"false,Driver'S License Number"` // PID-20
	MotherSIdentifier *CX `hl7:"false,Mother'S Identifier"` // PID-21
	EthnicGroup *CWE `hl7:"false,Ethnic Group"` // PID-22
	BirthPlace *ST `hl7:"false,Birth Place"` // PID-23
	MultipleBirthIndicator *ID `hl7:"false,Multiple Birth Indicator"` // PID-24
	BirthOrder *NM `hl7:"false,Birth Order"` // PID-25
	Citizenship *IS `hl7:"false,Citizenship"` // PID-26
	VeteransMilitaryStatus *CE `hl7:"false,Veterans Military Status"` // PID-27
	NationalityCode *CE `hl7:"false,Nationality Code"` // PID-28
	PatientDeathDateAndTime *TS `hl7:"false,Patient Death Date And Time"` // PID-29
	PatientDeathIndicator *ID `hl7:"false,Patient Death Indicator"` // PID-30
}

func (s *PID) SegmentName() string {
	return "PID"
}

// PR1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PR1 struct {
	SetIDProcedure *SI `hl7:"true,Set ID - Procedure"` // PR1-1
	ProcedureCodingMethod *IS `hl7:"true,Procedure Coding Method"` // PR1-2
	ProcedureCode *CE `hl7:"false,Procedure Code"` // PR1-3
	ProcedureDescription *ST `hl7:"false,Procedure Description"` // PR1-4
	ProcedureDateTime *TS `hl7:"false,Procedure Date/Time"` // PR1-5
	ProcedureType *ID `hl7:"true,Procedure Type"` // PR1-6
	ProcedureMinutes *NM `hl7:"false,Procedure Minutes"` // PR1-7
	Anesthesiologist []XCN `hl7:"false,Anesthesiologist"` // PR1-8
	AnesthesiaCode *IS `hl7:"false,Anesthesia Code"` // PR1-9
	AnesthesiaMinutes *NM `hl7:"false,Anesthesia Minutes"` // PR1-10
	Surgeon []XCN `hl7:"false,Surgeon"` // PR1-11
	ProcedurePractitioner []XCN `hl7:"false,Procedure Practitioner"` // PR1-12
	ConsentCode *CE `hl7:"false,Consent Code"` // PR1-13
	ProcedurePriority *NM `hl7:"false,Procedure Priority"` // PR1-14
	AssociatedDiagnosisCode *CE `hl7:"false,Associated Diagnosis Code"` // PR1-15
}

func (s *PR1) SegmentName() string {
	return "PR1"
}

// PRA represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PRA struct {
	PRAPrimaryKeyValue *ST `hl7:"true,PRA - Primary Key Value"` // PRA-1
	PractionerGroup []CE `hl7:"false,Practioner Group"` // PRA-2
	PractionerCategory []ID `hl7:"false,Practioner Category"` // PRA-3
	ProviderBilling *ID `hl7:"false,Provider Billing"` // PRA-4
	Specialty []CM `hl7:"false,Specialty"` // PRA-5
	PractitionerIDNumbers []CM `hl7:"false,Practitioner ID Numbers"` // PRA-6
	Privileges []CM `hl7:"false,Privileges"` // PRA-7
}

func (s *PRA) SegmentName() string {
	return "PRA"
}

// PRB represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PRB struct {
	ActionCode *ID `hl7:"true,Action Code"` // PRB-1
	ActionDateTime *TS `hl7:"true,Action Date/Time"` // PRB-2
	ProblemID *CE `hl7:"true,Problem ID"` // PRB-3
	ProblemInstanceID *EI `hl7:"true,Problem Instance ID"` // PRB-4
	EpisodeOfCareID *EI `hl7:"false,Episode Of Care ID"` // PRB-5
	ProblemListPriority *NM `hl7:"false,Problem List Priority"` // PRB-6
	ProblemEstablishedDateTime *TS `hl7:"false,Problem Established Date/Time"` // PRB-7
	AnticipatedProblemResolutionDateTime *TS `hl7:"false,Anticipated Problem Resolution Date/Time"` // PRB-8
	ActualProblemResolutionDateTime *TS `hl7:"false,Actual Problem Resolution Date/Time"` // PRB-9
	ProblemClassification *CE `hl7:"false,Problem Classification"` // PRB-10
	ProblemManagementDiscipline []CE `hl7:"false,Problem Management Discipline"` // PRB-11
	ProblemPersistence *CE `hl7:"false,Problem Persistence"` // PRB-12
	ProblemConfirmationStatus *CE `hl7:"false,Problem Confirmation Status"` // PRB-13
	ProblemLifeCycleStatus *CE `hl7:"false,Problem Life Cycle Status"` // PRB-14
	ProblemLifeCycleStatusDateTime *TS `hl7:"false,Problem Life Cycle Status Date/Time"` // PRB-15
	ProblemDateOfOnset *TS `hl7:"false,Problem Date Of Onset"` // PRB-16
	ProblemOnsetText *ST `hl7:"false,Problem Onset Text"` // PRB-17
	ProblemRanking *CE `hl7:"false,Problem Ranking"` // PRB-18
	CertaintyOfProblem *CE `hl7:"false,Certainty Of Problem"` // PRB-19
	ProbabilityOfProblem01 *NM `hl7:"false,Probability Of Problem (0-1)"` // PRB-20
	IndividualAwarenessOfProblem *CE `hl7:"false,Individual Awareness Of Problem"` // PRB-21
	ProblemPrognosis *CE `hl7:"false,Problem Prognosis"` // PRB-22
	IndividualAwarenessOfPrognosis *CE `hl7:"false,Individual Awareness Of Prognosis"` // PRB-23
	FamilySignificantOtherAwarenessOfProblemPrognosis *ST `hl7:"false,Family/Significant Other Awareness Of Problem/Prognosis"` // PRB-24
	SecuritySensitivity *CE `hl7:"false,Security/Sensitivity"` // PRB-25
}

func (s *PRB) SegmentName() string {
	return "PRB"
}

// PRC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PRC struct {
	PrimaryKeyValue *CE `hl7:"true,Primary Key Value"` // PRC-1
	FacilityID []EI `hl7:"true,Facility ID"` // PRC-2
	Department []CE `hl7:"false,Department"` // PRC-3
	ValidPatientClasses []ID `hl7:"false,Valid Patient Classes"` // PRC-4
	Price []CP `hl7:"false,Price"` // PRC-5
	Formula []ST `hl7:"false,Formula"` // PRC-6
	MinimumQuantity *NM `hl7:"false,Minimum Quantity"` // PRC-7
	MaximumQuantity *NM `hl7:"false,Maximum Quantity"` // PRC-8
	MinimumPrice *MO `hl7:"false,Minimum Price"` // PRC-9
	MaximumPrice *MO `hl7:"false,Maximum Price"` // PRC-10
	EffectiveStartDate *TS `hl7:"false,Effective Start Date"` // PRC-11
	EffectiveEndDate *TS `hl7:"false,Effective End Date"` // PRC-12
	PriceOverrideFlag *ID `hl7:"false,Price Override Flag"` // PRC-13
	BillingCategory []CE `hl7:"false,Billing Category"` // PRC-14
	ChargeableFlag *ID `hl7:"false,Chargeable Flag"` // PRC-15
	ActiveInactiveFlag *ID `hl7:"false,Active/Inactive Flag"` // PRC-16
	Cost *MO `hl7:"false,Cost"` // PRC-17
	ChargeOnIndicator *ID `hl7:"false,Charge On Indicator"` // PRC-18
}

func (s *PRC) SegmentName() string {
	return "PRC"
}

// PRD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PRD struct {
	Role []CE `hl7:"true,Role"` // PRD-1
	ProviderName []XPN `hl7:"false,Provider Name"` // PRD-2
	ProviderAddress *XAD `hl7:"false,Provider Address"` // PRD-3
	ProviderLocation *PL `hl7:"false,Provider Location"` // PRD-4
	ProviderCommunicationInformation []XTN `hl7:"false,Provider Communication Information"` // PRD-5
	PreferredMethodOfContact *CE `hl7:"false,Preferred Method Of Contact"` // PRD-6
	ProviderIdentifiers []CM `hl7:"false,Provider Identifiers"` // PRD-7
	EffectiveStartDateOfRole *TS `hl7:"false,Effective Start Date Of Role"` // PRD-8
	EffectiveEndDateOfRole *TS `hl7:"false,Effective End Date Of Role"` // PRD-9
}

func (s *PRD) SegmentName() string {
	return "PRD"
}

// PSH represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PSH struct {
	ReportType *ST `hl7:"true,Report Type"` // PSH-1
	ReportFormIdentifier *ST `hl7:"false,Report Form Identifier"` // PSH-2
	ReportDate *TS `hl7:"true,Report Date"` // PSH-3
	ReportIntervalStartDate *TS `hl7:"false,Report Interval Start Date"` // PSH-4
	ReportIntervalEndDate *TS `hl7:"false,Report Interval End Date"` // PSH-5
	QuantityManufactured *CQ `hl7:"false,Quantity Manufactured"` // PSH-6
	QuantityDistributed *CQ `hl7:"false,Quantity Distributed"` // PSH-7
	QuantityDistributedMethod *ID `hl7:"false,Quantity Distributed Method"` // PSH-8
	QuantityDistributedComment *FT `hl7:"false,Quantity Distributed Comment"` // PSH-9
	QuantityInUse *CQ `hl7:"false,Quantity In Use"` // PSH-10
	QuantityInUseMethod *ID `hl7:"false,Quantity In Use Method"` // PSH-11
	QuantityInUseComment *FT `hl7:"false,Quantity In Use Comment"` // PSH-12
	NumberOfProductExperienceReportsFiledByFacility []NM `hl7:"false,Number Of Product Experience Reports Filed By Facility"` // PSH-13
	NumberOfProductExperienceReportsFiledByDistributor []NM `hl7:"false,Number Of Product Experience Reports Filed By Distributor"` // PSH-14
}

func (s *PSH) SegmentName() string {
	return "PSH"
}

// PTH represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PTH struct {
	ActionCode *ID `hl7:"true,Action Code"` // PTH-1
	PathwayID *CE `hl7:"true,Pathway ID"` // PTH-2
	PathwayInstanceID *EI `hl7:"true,Pathway Instance ID"` // PTH-3
	PathwayEstablishedDateTime *TS `hl7:"true,Pathway Established Date/Time"` // PTH-4
	PathwayLifecycleStatus *CE `hl7:"false,Pathway Lifecycle Status"` // PTH-5
	ChangePathwayLifecycleStatusDateTime *TS `hl7:"false,Change Pathway Lifecycle Status Date/Time"` // PTH-6
}

func (s *PTH) SegmentName() string {
	return "PTH"
}

// PV1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PV1 struct {
	SetIDPatientVisit *SI `hl7:"false,Set ID - Patient Visit"` // PV1-1
	PatientClass *ID `hl7:"true,Patient Class"` // PV1-2
	AssignedPatientLocation *PL `hl7:"false,Assigned Patient Location"` // PV1-3
	AdmissionType *ID `hl7:"false,Admission Type"` // PV1-4
	PreadmitNumber *CX `hl7:"false,Preadmit Number"` // PV1-5
	PriorPatientLocation *PL `hl7:"false,Prior Patient Location"` // PV1-6
	AttendingDoctor *XCN `hl7:"false,Attending Doctor"` // PV1-7
	ReferringDoctor *XCN `hl7:"false,Referring Doctor"` // PV1-8
	ConsultingDoctor []XCN `hl7:"false,Consulting Doctor"` // PV1-9
	HospitalService *ID `hl7:"false,Hospital Service"` // PV1-10
	TemporaryLocation *PL `hl7:"false,Temporary Location"` // PV1-11
	PreadmitTestIndicator *ID `hl7:"false,Preadmit Test Indicator"` // PV1-12
	ReadmissionIndicator *ID `hl7:"false,Readmission Indicator"` // PV1-13
	AdmitSource *ID `hl7:"false,Admit Source"` // PV1-14
	AmbulatoryStatus *IS `hl7:"false,Ambulatory Status"` // PV1-15
	VIPIndicator *ID `hl7:"false,VIP Indicator"` // PV1-16
	AdmittingDoctor *XCN `hl7:"false,Admitting Doctor"` // PV1-17
	PatientType *ID `hl7:"false,Patient Type"` // PV1-18
	VisitNumber *CX `hl7:"false,Visit Number"` // PV1-19
	FinancialClass []FC `hl7:"false,Financial Class"` // PV1-20
	ChargePriceIndicator *ID `hl7:"false,Charge Price Indicator"` // PV1-21
	CourtesyCode *ID `hl7:"false,Courtesy Code"` // PV1-22
	CreditRating *ID `hl7:"false,Credit Rating"` // PV1-23
	ContractCode []ID `hl7:"false,Contract Code"` // PV1-24
	ContractEffectiveDate []DT `hl7:"false,Contract Effective Date"` // PV1-25
	ContractAmount []NM `hl7:"false,Contract Amount"` // PV1-26
	ContractPeriod []NM `hl7:"false,Contract Period"` // PV1-27
	InterestCode *ID `hl7:"false,Interest Code"` // PV1-28
	TransferToBadDebtCode *ID `hl7:"false,Transfer To Bad Debt Code"` // PV1-29
	TransferToBadDebtDate *DT `hl7:"false,Transfer To Bad Debt Date"` // PV1-30
	BadDebtAgencyCode *ID `hl7:"false,Bad Debt Agency Code"` // PV1-31
	BadDebtTransferAmount *NM `hl7:"false,Bad Debt Transfer Amount"` // PV1-32
	BadDebtRecoveryAmount *NM `hl7:"false,Bad Debt Recovery Amount"` // PV1-33
	DeleteAccountIndicator *ID `hl7:"false,Delete Account Indicator"` // PV1-34
	DeleteAccountDate *DT `hl7:"false,Delete Account Date"` // PV1-35
	DischargeDisposition *ID `hl7:"false,Discharge Disposition"` // PV1-36
	DischargedToLocation *CM `hl7:"false,Discharged To Location"` // PV1-37
	DietType *ID `hl7:"false,Diet Type"` // PV1-38
	ServicingFacility *ID `hl7:"false,Servicing Facility"` // PV1-39
	BedStatus *IS `hl7:"false,Bed Status"` // PV1-40
	AccountStatus *ID `hl7:"false,Account Status"` // PV1-41
	PendingLocation *PL `hl7:"false,Pending Location"` // PV1-42
	PriorTemporaryLocation *PL `hl7:"false,Prior Temporary Location"` // PV1-43
	AdmitDateTime *TS `hl7:"false,Admit Date/Time"` // PV1-44
	DischargeDateTime *TS `hl7:"false,Discharge Date/Time"` // PV1-45
	CurrentPatientBalance *NM `hl7:"false,Current Patient Balance"` // PV1-46
	TotalCharges *NM `hl7:"false,Total Charges"` // PV1-47
	TotalAdjustments *NM `hl7:"false,Total Adjustments"` // PV1-48
	TotalPayments *NM `hl7:"false,Total Payments"` // PV1-49
	AlternateVisitID *CX `hl7:"false,Alternate Visit ID"` // PV1-50
	VisitIndicator *IS `hl7:"false,Visit Indicator"` // PV1-51
	OtherHealthcareProvider []XCN `hl7:"false,Other Healthcare Provider"` // PV1-52
}

func (s *PV1) SegmentName() string {
	return "PV1"
}

// PV2 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type PV2 struct {
	PriorPendingLocation *PL `hl7:"false,Prior Pending Location"` // PV2-1
	AccommodationCode *CE `hl7:"false,Accommodation Code"` // PV2-2
	AdmitReason *CE `hl7:"false,Admit Reason"` // PV2-3
	TransferReason *CE `hl7:"false,Transfer Reason"` // PV2-4
	PatientValuables []ST `hl7:"false,Patient Valuables"` // PV2-5
	PatientValuablesLocation *ST `hl7:"false,Patient Valuables Location"` // PV2-6
	VisitUserCode *IS `hl7:"false,Visit User Code"` // PV2-7
	ExpectedAdmitDate *TS `hl7:"false,Expected Admit Date"` // PV2-8
	ExpectedDischargeDate *TS `hl7:"false,Expected Discharge Date"` // PV2-9
	EstimatedLengthOfInpatientStay *NM `hl7:"false,Estimated Length Of Inpatient Stay"` // PV2-10
	ActualLengthOfInpatientStay *NM `hl7:"false,Actual Length Of Inpatient Stay"` // PV2-11
	VisitDescription *ST `hl7:"false,Visit Description"` // PV2-12
	ReferralSourceCode *XCN `hl7:"false,Referral Source Code"` // PV2-13
	PreviousServiceDate *DT `hl7:"false,Previous Service Date"` // PV2-14
	EmploymentIllnessRelatedIndicator *ID `hl7:"false,Employment Illness Related Indicator"` // PV2-15
	PurgeStatusCode *IS `hl7:"false,Purge Status Code"` // PV2-16
	PurgeStatusDate *DT `hl7:"false,Purge Status Date"` // PV2-17
	SpecialProgramCode *IS `hl7:"false,Special Program Code"` // PV2-18
	RetentionIndicator *ID `hl7:"false,Retention Indicator"` // PV2-19
	ExpectedNumberOfInsurancePlans *NM `hl7:"false,Expected Number Of Insurance Plans"` // PV2-20
	VisitPublicityCode *IS `hl7:"false,Visit Publicity Code"` // PV2-21
	VisitProtectionIndicator *ID `hl7:"false,Visit Protection Indicator"` // PV2-22
	ClinicOrganizationName []XON `hl7:"false,Clinic Organization Name"` // PV2-23
	PatientStatusCode *IS `hl7:"false,Patient Status Code"` // PV2-24
	VisitPriorityCode *IS `hl7:"false,Visit Priority Code"` // PV2-25
	PreviousTreatmentDate *DT `hl7:"false,Previous Treatment Date"` // PV2-26
	ExpectedDischargeDisposition *IS `hl7:"false,Expected Discharge Disposition"` // PV2-27
	SignatureOnFileDate *DT `hl7:"false,Signature On File Date"` // PV2-28
	FirstSimilarIllnessDate *DT `hl7:"false,First Similar Illness Date"` // PV2-29
	PatientChargeAdjustmentCode *IS `hl7:"false,Patient Charge Adjustment Code"` // PV2-30
	RecurringServiceCode *IS `hl7:"false,Recurring Service Code"` // PV2-31
	BillingMediaCode *ID `hl7:"false,Billing Media Code"` // PV2-32
	ExpectedSurgeryDateTime *TS `hl7:"false,Expected Surgery Date & Time"` // PV2-33
	MilitaryPartnershipCode *ID `hl7:"false,Military Partnership Code"` // PV2-34
	MilitaryNonAvailabiltiyCode *ID `hl7:"false,Military Non-Availabiltiy Code"` // PV2-35
	NewbornBabyIndicator *ID `hl7:"false,Newborn Baby Indicator"` // PV2-36
	BabyDetainedIndicator *ID `hl7:"false,Baby Detained Indicator"` // PV2-37
}

func (s *PV2) SegmentName() string {
	return "PV2"
}

// QAK represents the corresponding HL7 segment.
// Definition from HL7 2.3
type QAK struct {
	QueryTag *ST `hl7:"false,Query Tag"` // QAK-1
	QueryResponseStatus *ID `hl7:"false,Query Response Status"` // QAK-2
}

func (s *QAK) SegmentName() string {
	return "QAK"
}

// QRD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type QRD struct {
	QueryDateTime *TS `hl7:"false,Query Date/Time"` // QRD-1
	QueryFormatCode *ID `hl7:"true,Query Format Code"` // QRD-2
	QueryPriority *ID `hl7:"true,Query Priority"` // QRD-3
	QueryID *ST `hl7:"true,Query ID"` // QRD-4
	DeferredResponseType *ID `hl7:"false,Deferred Response Type"` // QRD-5
	DeferredResponseDateTime *TS `hl7:"false,Deferred Response Date/Time"` // QRD-6
	QuantityLimitedRequest *CQ `hl7:"true,Quantity Limited Request"` // QRD-7
	WhoSubjectFilter []XCN `hl7:"true,Who Subject Filter"` // QRD-8
	WhatSubjectFilter []CE `hl7:"true,What Subject Filter"` // QRD-9
	WhatDepartmentDataCode []CE `hl7:"true,What Department Data Code"` // QRD-10
	WhatDataCodeValueQualifier []CM `hl7:"false,What Data Code Value Qualifier"` // QRD-11
	QueryResultsLevel *ID `hl7:"false,Query Results Level"` // QRD-12
}

func (s *QRD) SegmentName() string {
	return "QRD"
}

// QRF represents the corresponding HL7 segment.
// Definition from HL7 2.3
type QRF struct {
	WhereSubjectFilter []ST `hl7:"true,Where Subject Filter"` // QRF-1
	WhenDataStartDateTime *TS `hl7:"false,When Data Start Date/Time"` // QRF-2
	WhenDataEndDateTime *TS `hl7:"false,When Data End Date/Time"` // QRF-3
	WhatUserQualifier []ST `hl7:"false,What User Qualifier"` // QRF-4
	OtherQRYSubjectFilter []ST `hl7:"false,Other QRY Subject Filter"` // QRF-5
	WhichDateTimeQualifier []ID `hl7:"false,Which Date/Time Qualifier"` // QRF-6
	WhichDateTimeStatusQualifier []ID `hl7:"false,Which Date/Time Status Qualifier"` // QRF-7
	DateTimeSelectionQualifier []ID `hl7:"false,Date/Time Selection Qualifier"` // QRF-8
	WhenQuantityTimingQualifier *TQ `hl7:"false,When Quantity/Timing Qualifier"` // QRF-9
}

func (s *QRF) SegmentName() string {
	return "QRF"
}

// RDF represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RDF struct {
	NumberOfColumnsPerRow *NM `hl7:"true,Number Of Columns Per Row"` // RDF-1
	ColumnDescription []RCD `hl7:"true,Column Description"` // RDF-2
}

func (s *RDF) SegmentName() string {
	return "RDF"
}

// RDT represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RDT struct {
	ColumnValue *Any `hl7:"true,Column Value"` // RDT-1
}

func (s *RDT) SegmentName() string {
	return "RDT"
}

// RF1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RF1 struct {
	ReferralStatus *CE `hl7:"false,Referral Status"` // RF1-1
	ReferralPriority *CE `hl7:"false,Referral Priority"` // RF1-2
	ReferralType *CE `hl7:"false,Referral Type"` // RF1-3
	ReferralDisposition []CE `hl7:"false,Referral Disposition"` // RF1-4
	ReferralCategory *CE `hl7:"false,Referral Category"` // RF1-5
	OriginatingReferralIdentifier *EI `hl7:"true,Originating Referral Identifier"` // RF1-6
	EffectiveDate *TS `hl7:"false,Effective Date"` // RF1-7
	ExpirationDate *TS `hl7:"false,Expiration Date"` // RF1-8
	ProcessDate *TS `hl7:"false,Process Date"` // RF1-9
	ReferralReason []CE `hl7:"false,Referral Reason"` // RF1-10
	ExternalReferralIdentifier []EI `hl7:"false,External Referral Identifier"` // RF1-11
}

func (s *RF1) SegmentName() string {
	return "RF1"
}

// RGS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RGS struct {
	SetIDRGS *SI `hl7:"true,Set ID - RGS"` // RGS-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // RGS-2
	ResourceGroupID *CE `hl7:"false,Resource Group ID"` // RGS-3
}

func (s *RGS) SegmentName() string {
	return "RGS"
}

// ROL represents the corresponding HL7 segment.
// Definition from HL7 2.3
type ROL struct {
	RoleInstanceID *EI `hl7:"true,Role Instance ID"` // ROL-1
	ActionCode *ID `hl7:"true,Action Code"` // ROL-2
	Role *CE `hl7:"false,Role"` // ROL-3
	RolePerson *XCN `hl7:"true,Role Person"` // ROL-4
	RoleBeginDateTime *TS `hl7:"false,Role Begin Date/Time"` // ROL-5
	RoleEndDateTime *TS `hl7:"false,Role End Date/Time"` // ROL-6
	RoleDuration *CE `hl7:"false,Role Duration"` // ROL-7
	RoleActionAssumptionReason *CE `hl7:"false,Role Action (Assumption) Reason"` // ROL-8
}

func (s *ROL) SegmentName() string {
	return "ROL"
}

// RQ1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RQ1 struct {
	AnticipatedPrice *ST `hl7:"false,Anticipated Price"` // RQ1-1
	ManufacturedID *CE `hl7:"false,Manufactured ID"` // RQ1-2
	ManufacturerSCatalog *ST `hl7:"false,Manufacturer'S Catalog"` // RQ1-3
	VendorID *CE `hl7:"false,Vendor ID"` // RQ1-4
	VendorCatalog *ST `hl7:"false,Vendor Catalog"` // RQ1-5
	Taxable *ID `hl7:"false,Taxable"` // RQ1-6
	SubstituteAllowed *ID `hl7:"false,Substitute Allowed"` // RQ1-7
}

func (s *RQ1) SegmentName() string {
	return "RQ1"
}

// RQD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RQD struct {
	RequisitionLineNumber *SI `hl7:"false,Requisition Line Number"` // RQD-1
	ItemCodeInternal *CE `hl7:"false,Item Code - Internal"` // RQD-2
	ItemCodeExternal *CE `hl7:"false,Item Code - External"` // RQD-3
	HospitalItemCode *CE `hl7:"false,Hospital Item Code"` // RQD-4
	RequisitionQuantity *NM `hl7:"false,Requisition Quantity"` // RQD-5
	RequisitionUnitOfMeasure *CE `hl7:"false,Requisition Unit Of Measure"` // RQD-6
	DepartmentCostCenter *ID `hl7:"false,Department Cost Center"` // RQD-7
	ItemNaturalAccountCode *ID `hl7:"false,Item Natural Account Code"` // RQD-8
	DeliverToID *CE `hl7:"false,Deliver To ID"` // RQD-9
	DateNeeded *DT `hl7:"false,Date Needed"` // RQD-10
}

func (s *RQD) SegmentName() string {
	return "RQD"
}

// RX1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type RX1 struct {
	UNUSED1 *ST `hl7:"false,UNUSED"` // RX1-1
	UNUSED2 *ST `hl7:"false,UNUSED"` // RX1-2
	ROUTE *ST `hl7:"false,ROUTE"` // RX1-3
	SITEADMINISTERED *ST `hl7:"false,SITE ADMINISTERED"` // RX1-4
	IVSOLUTIONRATE *CQ `hl7:"false,IV SOLUTION RATE"` // RX1-5
	DRUGSTRENGTH *CQ `hl7:"false,DRUG STRENGTH"` // RX1-6
	FINALCONCENTRATION *NM `hl7:"false,FINAL CONCENTRATION"` // RX1-7
	FINALVOLUMEINML *NM `hl7:"false,FINAL VOLUME IN ML."` // RX1-8
	DRUGDOSE *CM `hl7:"false,DRUG DOSE"` // RX1-9
	DRUGROLE *ID `hl7:"false,DRUG ROLE"` // RX1-10
	PRESCRIPTIONSEQUENCE *NM `hl7:"false,PRESCRIPTION SEQUENCE #"` // RX1-11
	QUANTITYDISPENSED *CQ `hl7:"false,QUANTITY DISPENSED"` // RX1-12
	UNUSED3 *ST `hl7:"false,UNUSED"` // RX1-13
	DRUGID *CE `hl7:"false,DRUG ID"` // RX1-14
	COMPONENTDRUGIDS []ID `hl7:"false,COMPONENT DRUG IDS"` // RX1-15
	PRESCRIPTIONTYPE *ID `hl7:"false,PRESCRIPTION TYPE"` // RX1-16
	SUBSTITUTIONSTATUS *ID `hl7:"false,SUBSTITUTION STATUS"` // RX1-17
	RXORDERSTATUS *ID `hl7:"false,RX ORDER STATUS"` // RX1-18
	NUMBEROFREFILLS *NM `hl7:"false,NUMBER OF REFILLS"` // RX1-19
	UNUSED4 *ST `hl7:"false,UNUSED"` // RX1-20
	REFILLSREMAINING *NM `hl7:"false,REFILLS REMAINING"` // RX1-21
	DEACLASS *ID `hl7:"false,DEA CLASS"` // RX1-22
	ORDERINGMDSDEANUMBER *NM `hl7:"false,ORDERING MD'S DEA NUMBER"` // RX1-23
	UNUSED5 *ST `hl7:"false,UNUSED"` // RX1-24
	LASTREFILLDATETIME *TS `hl7:"false,LAST REFILL DATE/TIME"` // RX1-25
	RXNUMBER *ST `hl7:"false,RX NUMBER"` // RX1-26
	PRNSTATUS *ID `hl7:"false,PRN STATUS"` // RX1-27
	PHARMACYINSTRUCTIONS []TX `hl7:"false,PHARMACY INSTRUCTIONS"` // RX1-28
	PATIENTINSTRUCTIONS []TX `hl7:"false,PATIENT INSTRUCTIONS"` // RX1-29
	INSTRUCTIONS []TX `hl7:"false,INSTRUCTIONS"` // RX1-30
}

func (s *RX1) SegmentName() string {
	return "RX1"
}

// RXA represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXA struct {
	GiveSubIDCounter *ID `hl7:"true,Give Sub-ID Counter"` // RXA-1
	AdministrationSubIDCounter *NM `hl7:"true,Administration Sub-ID Counter"` // RXA-2
	DateTimeStartOfAdministration *TS `hl7:"true,Date/Time Start Of Administration"` // RXA-3
	DateTimeEndOfAdministration *TS `hl7:"true,Date/Time End Of Administration"` // RXA-4
	AdministeredCode *CE `hl7:"true,Administered Code"` // RXA-5
	AdministeredAmount *NM `hl7:"true,Administered Amount"` // RXA-6
	AdministeredUnits *CE `hl7:"false,Administered Units"` // RXA-7
	AdministeredDosageForm *CE `hl7:"false,Administered Dosage Form"` // RXA-8
	AdministrationNotes []CE `hl7:"false,Administration Notes"` // RXA-9
	AdministeringProvider *XCN `hl7:"false,Administering Provider"` // RXA-10
	AdministeredAtLocation *CM `hl7:"false,Administered-At Location"` // RXA-11
	AdministeredPerTimeUnit *ST `hl7:"false,Administered Per (Time Unit)"` // RXA-12
	AdministeredStrength *NM `hl7:"false,Administered Strength"` // RXA-13
	AdministeredStrengthUnits *CE `hl7:"false,Administered Strength Units"` // RXA-14
	SubstanceLotNumber []ST `hl7:"false,Substance Lot Number"` // RXA-15
	SubstanceExpirationDate []TS `hl7:"false,Substance Expiration Date"` // RXA-16
	SubstanceManufacturerName []CE `hl7:"false,Substance Manufacturer Name"` // RXA-17
	SubstanceRefusalReason []CE `hl7:"false,Substance Refusal Reason"` // RXA-18
	Indication []CE `hl7:"false,Indication"` // RXA-19
	CompletionStatus *ID `hl7:"false,Completion Status"` // RXA-20
	ActionCodeRXA *ID `hl7:"false,Action Code-RXA"` // RXA-21
	SystemEntryDateTime *TS `hl7:"false,System Entry Date/Time"` // RXA-22
}

func (s *RXA) SegmentName() string {
	return "RXA"
}

// RXC represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXC struct {
	RXComponentType *ID `hl7:"true,RX Component Type"` // RXC-1
	ComponentCode *CE `hl7:"true,Component Code"` // RXC-2
	ComponentAmount *NM `hl7:"true,Component Amount"` // RXC-3
	ComponentUnits *CE `hl7:"true,Component Units"` // RXC-4
	ComponentStrength *NM `hl7:"false,Component Strength"` // RXC-5
	ComponentStrengthUnits *CE `hl7:"false,Component Strength Units"` // RXC-6
}

func (s *RXC) SegmentName() string {
	return "RXC"
}

// RXD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXD struct {
	DispenseSubIDCounter *NM `hl7:"true,Dispense Sub-ID Counter"` // RXD-1
	DispenseGiveCode *CE `hl7:"true,Dispense/Give Code"` // RXD-2
	DateTimeDispensed *TS `hl7:"true,Date/Time Dispensed"` // RXD-3
	ActualDispenseAmount *NM `hl7:"true,Actual Dispense Amount"` // RXD-4
	ActualDispenseUnits *CE `hl7:"false,Actual Dispense Units"` // RXD-5
	ActualDosageForm *CE `hl7:"false,Actual Dosage Form"` // RXD-6
	PrescriptionNumber *ST `hl7:"true,Prescription Number"` // RXD-7
	NumberOfRefillsRemaining *NM `hl7:"false,Number Of Refills Remaining"` // RXD-8
	DispenseNotes []CE `hl7:"false,Dispense Notes"` // RXD-9
	DispensingProvider *XCN `hl7:"false,Dispensing Provider"` // RXD-10
	SubstitutionStatus *ID `hl7:"false,Substitution Status"` // RXD-11
	TotalDailyDose *CQ `hl7:"false,Total Daily Dose"` // RXD-12
	DispenseToLocation *CM `hl7:"false,Dispense-To Location"` // RXD-13
	NeedsHumanReview *ID `hl7:"false,Needs Human Review"` // RXD-14
	PharmacyTreatmentSupplierSSpecialDispensingInstructions []CE `hl7:"false,Pharmacy/Treatment Supplier'S Special Dispensing Instructions"` // RXD-15
	ActualStrength *NM `hl7:"false,Actual Strength"` // RXD-16
	ActualStrengthUnit *CE `hl7:"false,Actual Strength Unit"` // RXD-17
	SubstanceLotNumber []ST `hl7:"false,Substance Lot Number"` // RXD-18
	SubstanceExpirationDate []TS `hl7:"false,Substance Expiration Date"` // RXD-19
	SubstanceManufacturerName []CE `hl7:"false,Substance Manufacturer Name"` // RXD-20
	Indication *CE `hl7:"false,Indication"` // RXD-21
	DispensePackageSize *NM `hl7:"false,Dispense Package Size"` // RXD-22
	DispensePackageSizeUnit *CE `hl7:"false,Dispense Package Size Unit"` // RXD-23
	DispensePackageMethod *ID `hl7:"false,Dispense Package Method"` // RXD-24
}

func (s *RXD) SegmentName() string {
	return "RXD"
}

// RXE represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXE struct {
	QuantityTiming *TQ `hl7:"true,Quantity/Timing"` // RXE-1
	GiveCode *CE `hl7:"true,Give Code"` // RXE-2
	GiveAmountMinimum *NM `hl7:"true,Give Amount - Minimum"` // RXE-3
	GiveAmountMaximum *NM `hl7:"false,Give Amount - Maximum"` // RXE-4
	GiveUnits *CE `hl7:"true,Give Units"` // RXE-5
	GiveDosageForm *CE `hl7:"false,Give Dosage Form"` // RXE-6
	ProviderSAdministrationInstructions []CE `hl7:"false,Provider'S Administration Instructions"` // RXE-7
	DeliverToLocation *CM `hl7:"false,Deliver To Location"` // RXE-8
	SubstitutionStatus *ID `hl7:"false,Substitution Status"` // RXE-9
	DispenseAmount *NM `hl7:"false,Dispense Amount"` // RXE-10
	DispenseUnits *CE `hl7:"false,Dispense Units"` // RXE-11
	NumberOfRefills *NM `hl7:"false,Number Of Refills"` // RXE-12
	OrderingProviderSDEANumber *CN `hl7:"false,Ordering Provider'S DEA Number"` // RXE-13
	PharmacistTreatmentSupplierSVerifierID *CN `hl7:"false,Pharmacist/Treatment Supplier'S Verifier ID"` // RXE-14
	PrescriptionNumber *ST `hl7:"false,Prescription Number"` // RXE-15
	NumberOfRefillsRemaining *NM `hl7:"false,Number Of Refills Remaining"` // RXE-16
	NumberOfRefillsDosesDispensed *NM `hl7:"false,Number Of Refills/Doses Dispensed"` // RXE-17
	DateTimeOfMostRecentRefillOrDoseDispensed *TS `hl7:"false,Date / Time Of Most Recent Refill Or Dose Dispensed"` // RXE-18
	TotalDailyDose *CQ `hl7:"false,Total Daily Dose"` // RXE-19
	NeedsHumanReview *ID `hl7:"false,Needs Human Review"` // RXE-20
	PharmacyTreatmentSupplierSSpecialDispensingInstructions []CE `hl7:"false,Pharmacy/Treatment Supplier'S Special Dispensing Instructions"` // RXE-21
	GivePerTimeUnit *ST `hl7:"false,Give Per (Time Unit)"` // RXE-22
	GiveRateAmount *ST `hl7:"false,Give Rate Amount"` // RXE-23
	GiveRateUnits *CE `hl7:"false,Give Rate Units"` // RXE-24
	GiveStrength *NM `hl7:"false,Give Strength"` // RXE-25
	GiveStrengthUnits *CE `hl7:"false,Give Strength Units"` // RXE-26
	GiveIndication *CE `hl7:"false,Give Indication"` // RXE-27
	DispensePackageSize *NM `hl7:"false,Dispense Package Size"` // RXE-28
	DispensePackageSizeUnit *CE `hl7:"false,Dispense Package Size Unit"` // RXE-29
	DispensePackageMethod *ID `hl7:"false,Dispense Package Method"` // RXE-30
}

func (s *RXE) SegmentName() string {
	return "RXE"
}

// RXG represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXG struct {
	GiveSubIDCounter *ID `hl7:"false,Give Sub-ID Counter"` // RXG-1
	DispenseSubIDCounter *NM `hl7:"true,Dispense Sub-ID Counter"` // RXG-2
	QuantityTiming *TQ `hl7:"true,Quantity/Timing"` // RXG-3
	GiveCode *CE `hl7:"true,Give Code"` // RXG-4
	GiveAmountMinimum *NM `hl7:"true,Give Amount - Minimum"` // RXG-5
	GiveAmountMaximum *NM `hl7:"false,Give Amount - Maximum"` // RXG-6
	GiveUnits *CE `hl7:"true,Give Units"` // RXG-7
	GiveDosageForm *CE `hl7:"false,Give Dosage Form"` // RXG-8
	AdministrationNotes []CE `hl7:"false,Administration Notes"` // RXG-9
	SubstitutionStatus *ID `hl7:"false,Substitution Status"` // RXG-10
	DispenseToLocation *CM `hl7:"false,Dispense-To Location"` // RXG-11
	NeedsHumanReview *ID `hl7:"false,Needs Human Review"` // RXG-12
	PharmacySpecialAdministrationInstructions *CE `hl7:"false,Pharmacy Special Administration Instructions"` // RXG-13
	GivePerTimeUnit *ST `hl7:"false,Give Per (Time Unit)"` // RXG-14
	GiveRateAmount *ST `hl7:"false,Give Rate Amount"` // RXG-15
	GiveRateUnits *CE `hl7:"false,Give Rate Units"` // RXG-16
	GiveStrength *NM `hl7:"false,Give Strength"` // RXG-17
	GiveStrengthUnits *CE `hl7:"false,Give Strength Units"` // RXG-18
	SubstanceLotNumber []ST `hl7:"false,Substance Lot Number"` // RXG-19
	SubstanceExpirationDate []TS `hl7:"false,Substance Expiration Date"` // RXG-20
	SubstanceManufacturerName []CE `hl7:"false,Substance Manufacturer Name"` // RXG-21
	Indication *CE `hl7:"false,Indication"` // RXG-22
}

func (s *RXG) SegmentName() string {
	return "RXG"
}

// RXO represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXO struct {
	RequestedGiveCode *CE `hl7:"true,Requested Give Code"` // RXO-1
	RequestedGiveAmountMinimum *NM `hl7:"true,Requested Give Amount - Minimum"` // RXO-2
	RequestedGiveAmountMaximum *NM `hl7:"false,Requested Give Amount - Maximum"` // RXO-3
	RequestedGiveUnits *CE `hl7:"true,Requested Give Units"` // RXO-4
	RequestedDosageForm *CE `hl7:"false,Requested Dosage Form"` // RXO-5
	ProviderSPharmacyInstructions []CE `hl7:"false,Provider'S Pharmacy Instructions"` // RXO-6
	ProviderSAdministrationInstructions []CE `hl7:"false,Provider'S Administration Instructions"` // RXO-7
	DeliverToLocation *CM `hl7:"false,Deliver To Location"` // RXO-8
	AllowSubstitutions *ID `hl7:"false,Allow Substitutions"` // RXO-9
	RequestedDispenseCode *CE `hl7:"false,Requested Dispense Code"` // RXO-10
	RequestedDispenseAmount *NM `hl7:"false,Requested Dispense Amount"` // RXO-11
	RequestedDispenseUnits *CE `hl7:"false,Requested Dispense Units"` // RXO-12
	NumberOfRefills *NM `hl7:"false,Number Of Refills"` // RXO-13
	OrderingProviderSDEANumber *CN `hl7:"false,Ordering Provider'S DEA Number"` // RXO-14
	PharmacistTreatmentSupplierSVerifierID *CN `hl7:"false,Pharmacist/Treatment Supplier'S Verifier ID"` // RXO-15
	NeedsHumanReview *ID `hl7:"false,Needs Human Review"` // RXO-16
	RequestedGivePerTimeUnit *ST `hl7:"false,Requested Give Per (Time Unit)"` // RXO-17
	RequestedGiveStrength *NM `hl7:"false,Requested Give Strength"` // RXO-18
	RequestedGiveStrengthUnits *CE `hl7:"false,Requested Give Strength Units"` // RXO-19
	Indication *CE `hl7:"false,Indication"` // RXO-20
	RequestedGiveRateAmount *ST `hl7:"false,Requested Give Rate Amount"` // RXO-21
	RequestedGiveRateUnits *CE `hl7:"false,Requested Give Rate Units"` // RXO-22
}

func (s *RXO) SegmentName() string {
	return "RXO"
}

// RXR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type RXR struct {
	Route *CE `hl7:"true,Route"` // RXR-1
	Site *CE `hl7:"false,Site"` // RXR-2
	AdministrationDevice *CE `hl7:"false,Administration Device"` // RXR-3
	AdministrationMethod *CE `hl7:"false,Administration Method"` // RXR-4
}

func (s *RXR) SegmentName() string {
	return "RXR"
}

// SCH represents the corresponding HL7 segment.
// Definition from HL7 2.3
type SCH struct {
	PlacerAppointmentID *EI `hl7:"true,Placer Appointment ID"` // SCH-1
	FillerAppointmentID *EI `hl7:"false,Filler Appointment ID"` // SCH-2
	OccurrenceNumber *NM `hl7:"false,Occurrence Number"` // SCH-3
	PlacerGroupNumber *EI `hl7:"false,Placer Group Number"` // SCH-4
	ScheduleID *CE `hl7:"false,Schedule ID"` // SCH-5
	EventReason *CE `hl7:"true,Event Reason"` // SCH-6
	AppointmentReason *CE `hl7:"false,Appointment Reason"` // SCH-7
	AppointmentType *CE `hl7:"false,Appointment Type"` // SCH-8
	AppointmentDuration *NM `hl7:"false,Appointment Duration"` // SCH-9
	AppointmentDurationUnits *CE `hl7:"false,Appointment Duration Units"` // SCH-10
	AppointmentTimingQuantity []TQ `hl7:"true,Appointment Timing Quantity"` // SCH-11
	PlacerContactPerson *XCN `hl7:"false,Placer Contact Person"` // SCH-12
	PlacerContactPhoneNumber *XTN `hl7:"false,Placer Contact Phone Number"` // SCH-13
	PlacerContactAddress *XAD `hl7:"false,Placer Contact Address"` // SCH-14
	PlacerContactLocation *PL `hl7:"false,Placer Contact Location"` // SCH-15
	FillerContactPerson *XCN `hl7:"true,Filler Contact Person"` // SCH-16
	FillerContactPhoneNumber *XTN `hl7:"false,Filler Contact Phone Number"` // SCH-17
	FillerContactAddress *XAD `hl7:"false,Filler Contact Address"` // SCH-18
	FillerContactLocation *PL `hl7:"false,Filler Contact Location"` // SCH-19
	EnteredByPerson *XCN `hl7:"true,Entered By Person"` // SCH-20
	EnteredByPhoneNumber []XTN `hl7:"false,Entered By Phone Number"` // SCH-21
	EnteredByLocation *PL `hl7:"false,Entered By Location"` // SCH-22
	ParentPlacerAppointmentID *EI `hl7:"false,Parent Placer Appointment ID"` // SCH-23
	ParentFillerAppointmentID *EI `hl7:"false,Parent Filler Appointment ID"` // SCH-24
	FillerStatusCode *CE `hl7:"false,Filler Status Code"` // SCH-25
}

func (s *SCH) SegmentName() string {
	return "SCH"
}

// SPR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type SPR struct {
	QueryTag *ST `hl7:"false,Query Tag"` // SPR-1
	QueryResponseFormatCode *ID `hl7:"true,Query/ Response Format Code"` // SPR-2
	StoredProcedureName *CE `hl7:"true,Stored Procedure Name"` // SPR-3
	InputParameterList []QIP `hl7:"false,Input Parameter List"` // SPR-4
}

func (s *SPR) SegmentName() string {
	return "SPR"
}

// STF represents the corresponding HL7 segment.
// Definition from HL7 2.3
type STF struct {
	STFPrimaryKeyValue *CE `hl7:"true,STF - Primary Key Value"` // STF-1
	StaffIDCode []CE `hl7:"false,Staff ID Code"` // STF-2
	StaffName *XPN `hl7:"false,Staff Name"` // STF-3
	StaffType []ID `hl7:"false,Staff Type"` // STF-4
	Sex *IS `hl7:"false,Sex"` // STF-5
	DateOfBirth *TS `hl7:"false,Date Of Birth"` // STF-6
	ActiveInactiveFlag *ID `hl7:"false,Active/Inactive Flag"` // STF-7
	Department []CE `hl7:"false,Department"` // STF-8
	Service []CE `hl7:"false,Service"` // STF-9
	Phone []TN `hl7:"false,Phone"` // STF-10
	OfficeHomeAddress []AD `hl7:"false,Office/Home Address"` // STF-11
	ActivationDate []CM `hl7:"false,Activation Date"` // STF-12
	InactivationDate []CM `hl7:"false,Inactivation Date"` // STF-13
	BackupPersonID []CE `hl7:"false,Backup Person ID"` // STF-14
	EMailAddress []ST `hl7:"false,E-Mail Address"` // STF-15
	PreferredMethodOfContact *CE `hl7:"false,Preferred Method Of Contact"` // STF-16
	MaritalStatus []IS `hl7:"false,Marital Status"` // STF-17
	JobTitle *ST `hl7:"false,Job Title"` // STF-18
	JobCodeClass *JCC `hl7:"false,Job Code/Class"` // STF-19
	EmploymentStatus *IS `hl7:"false,Employment Status"` // STF-20
	AdditionalInsuredOnAuto *ID `hl7:"false,Additional Insured On Auto"` // STF-21
	DriverSLicenseNumber *DLN `hl7:"false,Driver'S License Number"` // STF-22
	CopyAutoIns *ID `hl7:"false,Copy Auto Ins"` // STF-23
	AutoInsExpires *DT `hl7:"false,Auto Ins. Expires"` // STF-24
	DateLastDMVReview *DT `hl7:"false,Date Last DMV Review"` // STF-25
	DateNextDMVReview *DT `hl7:"false,Date Next DMV Review"` // STF-26
}

func (s *STF) SegmentName() string {
	return "STF"
}

// TXA represents the corresponding HL7 segment.
// Definition from HL7 2.3
type TXA struct {
	SetIDTXA *SI `hl7:"true,Set ID- TXA"` // TXA-1
	DocumentType *IS `hl7:"true,Document Type"` // TXA-2
	DocumentContentPresentation *ID `hl7:"false,Document Content Presentation"` // TXA-3
	ActivityDateTime *TS `hl7:"false,Activity Date/Time"` // TXA-4
	PrimaryActivityProviderCodeName *XCN `hl7:"false,Primary Activity Provider Code/Name"` // TXA-5
	OriginationDateTime *TS `hl7:"false,Origination Date/Time"` // TXA-6
	TranscriptionDateTime *TS `hl7:"false,Transcription Date/Time"` // TXA-7
	EditDateTime []TS `hl7:"false,Edit Date/Time"` // TXA-8
	OriginatorCodeName *XCN `hl7:"false,Originator Code/Name"` // TXA-9
	AssignedDocumentAuthenticator []XCN `hl7:"false,Assigned Document Authenticator"` // TXA-10
	TranscriptionistCodeName *XCN `hl7:"false,Transcriptionist Code/Name"` // TXA-11
	UniqueDocumentNumber *EI `hl7:"true,Unique Document Number"` // TXA-12
	ParentDocumentNumber *EI `hl7:"false,Parent Document Number"` // TXA-13
	PlacerOrderNumber []EI `hl7:"false,Placer Order Number"` // TXA-14
	FillerOrderNumber *EI `hl7:"false,Filler Order Number"` // TXA-15
	UniqueDocumentFileName *ST `hl7:"false,Unique Document File Name"` // TXA-16
	DocumentCompletionStatus []ID `hl7:"true,Document Completion Status"` // TXA-17
	DocumentConfidentialityStatus *ID `hl7:"false,Document Confidentiality Status"` // TXA-18
	DocumentAvailabilityStatus *ID `hl7:"false,Document Availability Status"` // TXA-19
	DocumentStorageStatus *ID `hl7:"false,Document Storage Status"` // TXA-20
	DocumentChangeReason *ST `hl7:"false,Document Change Reason"` // TXA-21
	AuthenticationPersonTimeStamp []PPN `hl7:"false,Authentication Person, Time Stamp"` // TXA-22
	DistributedCopiesCodeAndNameOfRecipients []XCN `hl7:"false,Distributed Copies (Code And Name Of Recipients)"` // TXA-23
}

func (s *TXA) SegmentName() string {
	return "TXA"
}

// UB1 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type UB1 struct {
	SetIDUB1 *SI `hl7:"false,Set ID - UB1"` // UB1-1
	BloodDeductible43 *NM `hl7:"false,Blood Deductible  (43)"` // UB1-2
	BloodFurnishedPintsOf40 *NM `hl7:"false,Blood Furnished Pints Of (40)"` // UB1-3
	BloodReplacedPints41 *NM `hl7:"false,Blood Replaced Pints (41)"` // UB1-4
	BloodNotReplacedPints42 *NM `hl7:"false,Blood Not Replaced Pints(42)"` // UB1-5
	CoInsuranceDays25 *NM `hl7:"false,Co Insurance Days (25)"` // UB1-6
	ConditionCode3539 []IS `hl7:"false,Condition Code (35-39)"` // UB1-7
	CoveredDays23 *NM `hl7:"false,Covered Days   (23)"` // UB1-8
	NonCoveredDays24 *NM `hl7:"false,Non Covered Days   (24)"` // UB1-9
	ValueAmountCode4649 []CM `hl7:"false,Value Amount & Code (46-49)"` // UB1-10
	NumberOfGraceDays90 *NM `hl7:"false,Number Of Grace Days (90)"` // UB1-11
	SpecProgramIndicator44 *CE `hl7:"false,Spec Program Indicator (44)"` // UB1-12
	PSROURApprovalIndicator87 *ID `hl7:"false,PSRO/UR Approval Indicator (87)"` // UB1-13
	PSROURApprovedStayFm88 *DT `hl7:"false,PSRO/UR Approved Stay Fm (88)"` // UB1-14
	PSROURApprovedStayTo89 *DT `hl7:"false,PSRO/UR Approved Stay To (89)"` // UB1-15
	Occurrence2832 []CM `hl7:"false,Occurrence (28 32)"` // UB1-16
	OccurrenceSpan33 *ID `hl7:"false,Occurrence Span (33)"` // UB1-17
	OccurSpanStartDate33 *DT `hl7:"false,Occur Span Start Date(33)"` // UB1-18
	OccurSpanEndDate33 *DT `hl7:"false,Occur Span End Date (33)"` // UB1-19
	UB82Locator2 *ST `hl7:"false,UB 82 Locator 2"` // UB1-20
	UB82Locator9 *ST `hl7:"false,UB 82 Locator 9"` // UB1-21
	UB82Locator27 *ST `hl7:"false,UB 82 Locator 27"` // UB1-22
	UB82Locator45 *ST `hl7:"false,UB 82 Locator 45"` // UB1-23
}

func (s *UB1) SegmentName() string {
	return "UB1"
}

// UB2 represents the corresponding HL7 segment.
// Definition from HL7 2.3
type UB2 struct {
	SetIDUB2 *SI `hl7:"false,Set ID - UB2"` // UB2-1
	CoInsuranceDays9 *ST `hl7:"false,Co-Insurance Days (9)"` // UB2-2
	ConditionCode2430 []ID `hl7:"false,Condition Code (24-30)"` // UB2-3
	CoveredDays7 *ST `hl7:"false,Covered Days (7)"` // UB2-4
	NonCoveredDays8 *ST `hl7:"false,Non-Covered Days (8)"` // UB2-5
	ValueAmountCode []CM `hl7:"false,Value Amount & Code"` // UB2-6
	OccurrenceCodeDate3235 []CM `hl7:"false,Occurrence Code & Date (32-35)"` // UB2-7
	OccurrenceSpanCodeDates36 []CM `hl7:"false,Occurrence Span Code/Dates (36)"` // UB2-8
	UB92Locator2State []ST `hl7:"false,UB92 Locator 2 (State)"` // UB2-9
	UB92Locator11State []ST `hl7:"false,UB92 Locator 11 (State)"` // UB2-10
	UB92Locator31National *ST `hl7:"false,UB92 Locator 31 (National)"` // UB2-11
	DocumentControlNumber []ST `hl7:"false,Document Control Number"` // UB2-12
	UB92Locator49National []ST `hl7:"false,UB92 Locator 49 (National)"` // UB2-13
	UB92Locator56State []ST `hl7:"false,UB92 Locator 56 (State)"` // UB2-14
	UB92Locator57National *ST `hl7:"false,UB92 Locator 57 (National)"` // UB2-15
	UB92Locator78State []ST `hl7:"false,UB92 Locator 78 (State)"` // UB2-16
	SpecialVisitCount *NM `hl7:"false,Special Visit Count"` // UB2-17
}

func (s *UB2) SegmentName() string {
	return "UB2"
}

// URD represents the corresponding HL7 segment.
// Definition from HL7 2.3
type URD struct {
	RUDateTime *TS `hl7:"false,R/U Date/Time"` // URD-1
	ReportPriority *ID `hl7:"false,Report Priority"` // URD-2
	RUWhoSubjectDefinition []XCN `hl7:"true,R/U Who Subject Definition"` // URD-3
	RUWhatSubjectDefinition []CE `hl7:"false,R/U What Subject Definition"` // URD-4
	RUWhatDepartmentCode []CE `hl7:"false,R/U What Department Code"` // URD-5
	RUDisplayPrintLocations []ST `hl7:"false,R/U Display/Print Locations"` // URD-6
	RUResultsLevel *ID `hl7:"false,R/U Results Level"` // URD-7
}

func (s *URD) SegmentName() string {
	return "URD"
}

// URS represents the corresponding HL7 segment.
// Definition from HL7 2.3
type URS struct {
	RUWhereSubjectDefinition []ST `hl7:"true,R/U Where Subject Definition"` // URS-1
	RUWhenDataStartDateTime *TS `hl7:"false,R/U When Data Start Date/Time"` // URS-2
	RUWhenDataEndDateTime *TS `hl7:"false,R/U When Data End Date/Time"` // URS-3
	RUWhatUserQualifier []ST `hl7:"false,R/U What User Qualifier"` // URS-4
	RUOtherResultsSubjectDefinition []ST `hl7:"false,R/U Other Results Subject Definition"` // URS-5
	RUWhichDateTimeQualifier []ID `hl7:"false,R/U Which Date/Time Qualifier"` // URS-6
	RUWhichDateTimeStatusQualifier []ID `hl7:"false,R/U Which Date/Time Status Qualifier"` // URS-7
	RUDateTimeSelectionQualifier []ID `hl7:"false,R/U Date/Time Selection Qualifier"` // URS-8
	RUQuantityTimingQualifier *TQ `hl7:"false,R/U Quantity/Timing Qualifier"` // URS-9
}

func (s *URS) SegmentName() string {
	return "URS"
}

// VAR represents the corresponding HL7 segment.
// Definition from HL7 2.3
type VAR struct {
	VarianceInstanceID *EI `hl7:"true,Variance Instance ID"` // VAR-1
	DocumentedDateTime *TS `hl7:"true,Documented Date/Time"` // VAR-2
	StatedVarianceDateTime *TS `hl7:"false,Stated Variance Date/Time"` // VAR-3
	VarianceOriginator *XCN `hl7:"false,Variance Originator"` // VAR-4
	VarianceClassification *CE `hl7:"false,Variance Classification"` // VAR-5
	VarianceDescription []ST `hl7:"false,Variance Description"` // VAR-6
}

func (s *VAR) SegmentName() string {
	return "VAR"
}

// VTQ represents the corresponding HL7 segment.
// Definition from HL7 2.3
type VTQ struct {
	QueryTag *ST `hl7:"false,Query Tag"` // VTQ-1
	QueryResponseFormatCode *ID `hl7:"true,Query/ Response Format Code"` // VTQ-2
	VTQueryName *CE `hl7:"true,VT Query Name"` // VTQ-3
	VirtualTableName *CE `hl7:"true,Virtual Table Name"` // VTQ-4
	SelectionCriteria []QSC `hl7:"false,Selection Criteria"` // VTQ-5
}

func (s *VTQ) SegmentName() string {
	return "VTQ"
}


// ACK represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ACK struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	Other []interface{}
}

func (s *ACK) MessageTypeName() string {
	return "ACK"
}
// ADR_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QUERY_RESPONSE []ADR_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ADR_A19) MessageTypeName() string {
	return "ADR_A19"
}
// ADR_A19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADR_A19_INSURANCE) MessageTypeSubStructName() string {
	return "ADR_A19_INSURANCE"
}
// ADR_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19_QUERY_RESPONSE struct {
	EVN *EVN `hl7:"false,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADR_A19_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADR_A19_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ADR_A19_QUERY_RESPONSE"
}
// ADT_A01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A01 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	PROCEDURE []ADT_A01_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A01_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A01) MessageTypeName() string {
	return "ADT_A01"
}
// ADT_A01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A01_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A01_INSURANCE"
}
// ADT_A01_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A01_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A01_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A01_PROCEDURE"
}
// ADT_A02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A02 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A02) MessageTypeName() string {
	return "ADT_A02"
}
// ADT_A03 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A03 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	PROCEDURE []ADT_A03_PROCEDURE `hl7:"false,PROCEDURE"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A03) MessageTypeName() string {
	return "ADT_A03"
}
// ADT_A03_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A03_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A03_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A03_PROCEDURE"
}
// ADT_A04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A04 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A04_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A04) MessageTypeName() string {
	return "ADT_A04"
}
// ADT_A04_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A04_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A04_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A04_INSURANCE"
}
// ADT_A05 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A05 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A05_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A05) MessageTypeName() string {
	return "ADT_A05"
}
// ADT_A05_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A05_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A05_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A05_INSURANCE"
}
// ADT_A06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A06 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"false,MRG"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	DRG1 *DRG `hl7:"false,DRG1"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG2 *DRG `hl7:"false,DRG2"`
	PROCEDURE []ADT_A06_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A06_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A06) MessageTypeName() string {
	return "ADT_A06"
}
// ADT_A06_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A06_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A06_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A06_INSURANCE"
}
// ADT_A06_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A06_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A06_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A06_PROCEDURE"
}
// ADT_A07 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A07 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	MRG *MRG `hl7:"false,MRG"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A07_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A07) MessageTypeName() string {
	return "ADT_A07"
}
// ADT_A07_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A07_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A07_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A07_INSURANCE"
}
// ADT_A08 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A08 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A08_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A08) MessageTypeName() string {
	return "ADT_A08"
}
// ADT_A08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A08_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A08_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A08_INSURANCE"
}
// ADT_A09 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A09 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A09) MessageTypeName() string {
	return "ADT_A09"
}
// ADT_A10 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A10 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A10) MessageTypeName() string {
	return "ADT_A10"
}
// ADT_A11 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A11 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A11) MessageTypeName() string {
	return "ADT_A11"
}
// ADT_A12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A12 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A12) MessageTypeName() string {
	return "ADT_A12"
}
// ADT_A13 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A13 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A13_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A13) MessageTypeName() string {
	return "ADT_A13"
}
// ADT_A13_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A13_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A13_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A13_INSURANCE"
}
// ADT_A14 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A14 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A14_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A14) MessageTypeName() string {
	return "ADT_A14"
}
// ADT_A14_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A14_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A14_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A14_INSURANCE"
}
// ADT_A15 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A15 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A15) MessageTypeName() string {
	return "ADT_A15"
}
// ADT_A16 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A16 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 *DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	Other []interface{}
}

func (s *ADT_A16) MessageTypeName() string {
	return "ADT_A16"
}
// ADT_A17 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A17 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID1 *PID `hl7:"true,PID1"`
	PD11 *PD1 `hl7:"false,PD11"`
	PV11 *PV1 `hl7:"true,PV11"`
	PV21 *PV2 `hl7:"false,PV21"`
	DB11 []DB1 `hl7:"false,DB11"`
	OBX1 []OBX `hl7:"false,OBX1"`
	PID2 *PID `hl7:"true,PID2"`
	PD12 *PD1 `hl7:"false,PD12"`
	PV12 *PV1 `hl7:"true,PV12"`
	PV22 *PV2 `hl7:"false,PV22"`
	DB12 []DB1 `hl7:"false,DB12"`
	OBX2 []OBX `hl7:"false,OBX2"`
	Other []interface{}
}

func (s *ADT_A17) MessageTypeName() string {
	return "ADT_A17"
}
// ADT_A17_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A17_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A17_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A17_PATIENT"
}
// ADT_A18 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A18 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"false,MRG"`
	PV1 *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A18) MessageTypeName() string {
	return "ADT_A18"
}
// ADT_A20 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A20 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	NPU *NPU `hl7:"true,NPU"`
	Other []interface{}
}

func (s *ADT_A20) MessageTypeName() string {
	return "ADT_A20"
}
// ADT_A21 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A21 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A21) MessageTypeName() string {
	return "ADT_A21"
}
// ADT_A22 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A22 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A22) MessageTypeName() string {
	return "ADT_A22"
}
// ADT_A23 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A23 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A23) MessageTypeName() string {
	return "ADT_A23"
}
// ADT_A24 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A24 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID1 *PID `hl7:"true,PID1"`
	PD11 *PD1 `hl7:"false,PD11"`
	PV11 *PV1 `hl7:"false,PV11"`
	DB11 []DB1 `hl7:"false,DB11"`
	PID2 *PID `hl7:"true,PID2"`
	PD12 *PD1 `hl7:"false,PD12"`
	PV12 *PV1 `hl7:"false,PV12"`
	DB12 []DB1 `hl7:"false,DB12"`
	Other []interface{}
}

func (s *ADT_A24) MessageTypeName() string {
	return "ADT_A24"
}
// ADT_A25 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A25 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A25) MessageTypeName() string {
	return "ADT_A25"
}
// ADT_A26 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A26 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A26) MessageTypeName() string {
	return "ADT_A26"
}
// ADT_A27 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A27 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A27) MessageTypeName() string {
	return "ADT_A27"
}
// ADT_A28 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A28 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A28_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A28) MessageTypeName() string {
	return "ADT_A28"
}
// ADT_A28_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A28_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A28_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A28_INSURANCE"
}
// ADT_A29 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A29 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A29) MessageTypeName() string {
	return "ADT_A29"
}
// ADT_A30 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A30 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A30) MessageTypeName() string {
	return "ADT_A30"
}
// ADT_A31 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A31 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	PR1 []PR1 `hl7:"false,PR1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ADT_A31_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ADT_A31) MessageTypeName() string {
	return "ADT_A31"
}
// ADT_A31_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A31_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A31_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A31_INSURANCE"
}
// ADT_A32 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A32 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A32) MessageTypeName() string {
	return "ADT_A32"
}
// ADT_A33 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A33 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A33) MessageTypeName() string {
	return "ADT_A33"
}
// ADT_A34 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A34 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A34) MessageTypeName() string {
	return "ADT_A34"
}
// ADT_A35 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A35 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A35) MessageTypeName() string {
	return "ADT_A35"
}
// ADT_A36 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A36 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A36) MessageTypeName() string {
	return "ADT_A36"
}
// ADT_A37 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A37 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID1 *PID `hl7:"true,PID1"`
	PV11 *PV1 `hl7:"false,PV11"`
	PID2 *PID `hl7:"true,PID2"`
	PV12 *PV1 `hl7:"false,PV12"`
	Other []interface{}
}

func (s *ADT_A37) MessageTypeName() string {
	return "ADT_A37"
}
// ADT_A38 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A38 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	Other []interface{}
}

func (s *ADT_A38) MessageTypeName() string {
	return "ADT_A38"
}
// ADT_A39 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A39 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []ADT_A39_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *ADT_A39) MessageTypeName() string {
	return "ADT_A39"
}
// ADT_A39_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A39_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	PV1 *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ADT_A39_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A39_PATIENT"
}
// ADT_A40 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A40 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []ADT_A40_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *ADT_A40) MessageTypeName() string {
	return "ADT_A40"
}
// ADT_A40_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A40_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	PV1 *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ADT_A40_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A40_PATIENT"
}
// ADT_A43 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A43 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []ADT_A43_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *ADT_A43) MessageTypeName() string {
	return "ADT_A43"
}
// ADT_A43_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A43_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A43_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A43_PATIENT"
}
// ADT_A44 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A44 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []ADT_A44_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *ADT_A44) MessageTypeName() string {
	return "ADT_A44"
}
// ADT_A44_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A44_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A44_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A44_PATIENT"
}
// ADT_A45 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A45 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MERGE_INFO []ADT_A45_MERGE_INFO `hl7:"true,MERGE_INFO"`
	Other []interface{}
}

func (s *ADT_A45) MessageTypeName() string {
	return "ADT_A45"
}
// ADT_A45_MERGE_INFO represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A45_MERGE_INFO struct {
	MRG *MRG `hl7:"true,MRG"`
	PV1 *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A45_MERGE_INFO) MessageTypeSubStructName() string {
	return "ADT_A45_MERGE_INFO"
}
// ADT_A50 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ADT_A50 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	MRG *MRG `hl7:"true,MRG"`
	PV1 *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A50) MessageTypeName() string {
	return "ADT_A50"
}
// ARD_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	QUERY_RESPONSE []ARD_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ARD_A19) MessageTypeName() string {
	return "ARD_A19"
}
// ARD_A19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ARD_A19_INSURANCE) MessageTypeSubStructName() string {
	return "ARD_A19_INSURANCE"
}
// ARD_A19_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ARD_A19_PROCEDURE) MessageTypeSubStructName() string {
	return "ARD_A19_PROCEDURE"
}
// ARD_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_QUERY_RESPONSE struct {
	EVN *EVN `hl7:"false,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NK1 []NK1 `hl7:"false,NK1"`
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	PROCEDURE []ARD_A19_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []ARD_A19_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *ARD_A19_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ARD_A19_QUERY_RESPONSE"
}
// BAR_P01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P01 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	VISIT []BAR_P01_VISIT `hl7:"true,VISIT"`
	Other []interface{}
}

func (s *BAR_P01) MessageTypeName() string {
	return "BAR_P01"
}
// BAR_P01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *BAR_P01_INSURANCE) MessageTypeSubStructName() string {
	return "BAR_P01_INSURANCE"
}
// BAR_P01_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P01_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *BAR_P01_PROCEDURE) MessageTypeSubStructName() string {
	return "BAR_P01_PROCEDURE"
}
// BAR_P01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P01_VISIT struct {
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	AL1 []AL1 `hl7:"false,AL1"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	PROCEDURE []BAR_P01_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1 []GT1 `hl7:"false,GT1"`
	NK1 []NK1 `hl7:"false,NK1"`
	INSURANCE []BAR_P01_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	UB1 *UB1 `hl7:"false,UB1"`
	UB2 *UB2 `hl7:"false,UB2"`
	Other []interface{}
}

func (s *BAR_P01_VISIT) MessageTypeSubStructName() string {
	return "BAR_P01_VISIT"
}
// BAR_P02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P02 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []BAR_P02_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *BAR_P02) MessageTypeName() string {
	return "BAR_P02"
}
// BAR_P02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"false,PV1"`
	DB1 []DB1 `hl7:"false,DB1"`
	Other []interface{}
}

func (s *BAR_P02_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P02_PATIENT"
}
// BAR_P06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P06 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PATIENT []BAR_P06_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *BAR_P06) MessageTypeName() string {
	return "BAR_P06"
}
// BAR_P06_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type BAR_P06_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *BAR_P06_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P06_PATIENT"
}
// CRM_C01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CRM_C01 struct {
	MSH *MSH `hl7:"true,MSH"`
	PATIENT []CRM_C01_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *CRM_C01) MessageTypeName() string {
	return "CRM_C01"
}
// CRM_C01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CRM_C01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	CSR *CSR `hl7:"true,CSR"`
	CSP []CSP `hl7:"false,CSP"`
	Other []interface{}
}

func (s *CRM_C01_PATIENT) MessageTypeSubStructName() string {
	return "CRM_C01_PATIENT"
}
// CSU_C09 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09 struct {
	MSH *MSH `hl7:"true,MSH"`
	PATIENT []CSU_C09_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *CSU_C09) MessageTypeName() string {
	return "CSU_C09"
}
// CSU_C09_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	VISIT *CSU_C09_VISIT `hl7:"false,VISIT"`
	CSR *CSR `hl7:"true,CSR"`
	STUDY_PHASE []CSU_C09_STUDY_PHASE `hl7:"true,STUDY_PHASE"`
	Other []interface{}
}

func (s *CSU_C09_PATIENT) MessageTypeSubStructName() string {
	return "CSU_C09_PATIENT"
}
// CSU_C09_RX_ADMIN represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_RX_ADMIN struct {
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"true,RXR"`
	Other []interface{}
}

func (s *CSU_C09_RX_ADMIN) MessageTypeSubStructName() string {
	return "CSU_C09_RX_ADMIN"
}
// CSU_C09_STUDY_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_STUDY_OBSERVATION struct {
	ORC *ORC `hl7:"false,ORC"`
	OBR *OBR `hl7:"true,OBR"`
	OBX []OBX `hl7:"true,OBX"`
	Other []interface{}
}

func (s *CSU_C09_STUDY_OBSERVATION) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_OBSERVATION"
}
// CSU_C09_STUDY_PHARM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_STUDY_PHARM struct {
	ORC *ORC `hl7:"false,ORC"`
	RX_ADMIN []CSU_C09_RX_ADMIN `hl7:"true,RX_ADMIN"`
	Other []interface{}
}

func (s *CSU_C09_STUDY_PHARM) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_PHARM"
}
// CSU_C09_STUDY_PHASE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_STUDY_PHASE struct {
	CSP *CSP `hl7:"false,CSP"`
	STUDY_SCHEDULE []CSU_C09_STUDY_SCHEDULE `hl7:"true,STUDY_SCHEDULE"`
	Other []interface{}
}

func (s *CSU_C09_STUDY_PHASE) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_PHASE"
}
// CSU_C09_STUDY_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_STUDY_SCHEDULE struct {
	CSS *CSS `hl7:"false,CSS"`
	STUDY_OBSERVATION []CSU_C09_STUDY_OBSERVATION `hl7:"false,STUDY_OBSERVATION"`
	STUDY_PHARM []CSU_C09_STUDY_PHARM `hl7:"true,STUDY_PHARM"`
	Other []interface{}
}

func (s *CSU_C09_STUDY_SCHEDULE) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_SCHEDULE"
}
// CSU_C09_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type CSU_C09_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *CSU_C09_VISIT) MessageTypeSubStructName() string {
	return "CSU_C09_VISIT"
}
// DFT_P03 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DFT_P03 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DB1 []DB1 `hl7:"false,DB1"`
	OBX []OBX `hl7:"false,OBX"`
	FINANCIAL []DFT_P03_FINANCIAL `hl7:"true,FINANCIAL"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG *DRG `hl7:"false,DRG"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []DFT_P03_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	Other []interface{}
}

func (s *DFT_P03) MessageTypeName() string {
	return "DFT_P03"
}
// DFT_P03_FINANCIAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DFT_P03_FINANCIAL struct {
	FT1 *FT1 `hl7:"true,FT1"`
	FINANCIAL_PROCEDURE []DFT_P03_FINANCIAL_PROCEDURE `hl7:"false,FINANCIAL_PROCEDURE"`
	Other []interface{}
}

func (s *DFT_P03_FINANCIAL) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL"
}
// DFT_P03_FINANCIAL_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DFT_P03_FINANCIAL_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	ROL []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P03_FINANCIAL_PROCEDURE) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL_PROCEDURE"
}
// DFT_P03_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DFT_P03_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *DFT_P03_INSURANCE) MessageTypeSubStructName() string {
	return "DFT_P03_INSURANCE"
}
// DOC_T12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DOC_T12 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	RESULT []DOC_T12_RESULT `hl7:"true,RESULT"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DOC_T12) MessageTypeName() string {
	return "DOC_T12"
}
// DOC_T12_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DOC_T12_RESULT struct {
	EVN *EVN `hl7:"false,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	TXA *TXA `hl7:"true,TXA"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *DOC_T12_RESULT) MessageTypeSubStructName() string {
	return "DOC_T12_RESULT"
}
// DSR_P04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_P04 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_P04) MessageTypeName() string {
	return "DSR_P04"
}
// DSR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DSR_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"false,QAK"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_Q01) MessageTypeName() string {
	return "DSR_Q01"
}
// DSR_Q03 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type DSR_Q03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"false,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"false,QAK"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_Q03) MessageTypeName() string {
	return "DSR_Q03"
}
// DSR_R03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_R03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"false,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_R03) MessageTypeName() string {
	return "DSR_R03"
}
// EDR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type EDR_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"true,QAK"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EDR_Q01) MessageTypeName() string {
	return "EDR_Q01"
}
// EQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type EQQ_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	EQL *EQL `hl7:"true,EQL"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EQQ_Q01) MessageTypeName() string {
	return "EQQ_Q01"
}
// ERP_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ERP_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"true,QAK"`
	ERQ *ERQ `hl7:"true,ERQ"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ERP_Q01) MessageTypeName() string {
	return "ERP_Q01"
}
// MCF_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type MCF_Q02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	Other []interface{}
}

func (s *MCF_Q02) MessageTypeName() string {
	return "MCF_Q02"
}
// MDM_T01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MDM_T01 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	TXA *TXA `hl7:"true,TXA"`
	Other []interface{}
}

func (s *MDM_T01) MessageTypeName() string {
	return "MDM_T01"
}
// MDM_T02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MDM_T02 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"true,PV1"`
	TXA *TXA `hl7:"true,TXA"`
	OBX []OBX `hl7:"true,OBX"`
	Other []interface{}
}

func (s *MDM_T02) MessageTypeName() string {
	return "MDM_T02"
}
// MFD_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M01) MessageTypeName() string {
	return "MFD_M01"
}
// MFD_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M02) MessageTypeName() string {
	return "MFD_M02"
}
// MFD_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M03) MessageTypeName() string {
	return "MFD_M03"
}
// MFK_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFK_M01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M01) MessageTypeName() string {
	return "MFK_M01"
}
// MFK_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFK_M02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M02) MessageTypeName() string {
	return "MFK_M02"
}
// MFK_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFK_M03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	MFI *MFI `hl7:"true,MFI"`
	MFA []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M03) MessageTypeName() string {
	return "MFK_M03"
}
// MFN_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF []MFN_M01_MF `hl7:"true,MF"`
	Other []interface{}
}

func (s *MFN_M01) MessageTypeName() string {
	return "MFN_M01"
}
// MFN_M01_MF represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M01_MF struct {
	MFE *MFE `hl7:"true,MFE"`
	Other []interface{}
}

func (s *MFN_M01_MF) MessageTypeSubStructName() string {
	return "MFN_M01_MF"
}
// MFN_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_STAFF []MFN_M02_MF_STAFF `hl7:"true,MF_STAFF"`
	Other []interface{}
}

func (s *MFN_M02) MessageTypeName() string {
	return "MFN_M02"
}
// MFN_M02_MF_STAFF represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M02_MF_STAFF struct {
	MFE *MFE `hl7:"true,MFE"`
	STF *STF `hl7:"true,STF"`
	PRA *PRA `hl7:"false,PRA"`
	Other []interface{}
}

func (s *MFN_M02_MF_STAFF) MessageTypeSubStructName() string {
	return "MFN_M02_MF_STAFF"
}
// MFN_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST []MFN_M03_MF_TEST `hl7:"true,MF_TEST"`
	Other []interface{}
}

func (s *MFN_M03) MessageTypeName() string {
	return "MFN_M03"
}
// MFN_M03_MF_TEST represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M03_MF_TEST struct {
	MFE *MFE `hl7:"true,MFE"`
	OM1 *OM1 `hl7:"true,OM1"`
	// Missing: anyHL7Segment
	Other []interface{}
}

func (s *MFN_M03_MF_TEST) MessageTypeSubStructName() string {
	return "MFN_M03_MF_TEST"
}
// MFN_M05 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M05 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_LOCATION []MFN_M05_MF_LOCATION `hl7:"true,MF_LOCATION"`
	Other []interface{}
}

func (s *MFN_M05) MessageTypeName() string {
	return "MFN_M05"
}
// MFN_M05_MF_LOCATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M05_MF_LOCATION struct {
	MFE *MFE `hl7:"true,MFE"`
	LOC *LOC `hl7:"true,LOC"`
	LCH []LCH `hl7:"false,LCH"`
	LRL []LRL `hl7:"false,LRL"`
	MF_LOC_DEPT []MFN_M05_MF_LOC_DEPT `hl7:"true,MF_LOC_DEPT"`
	Other []interface{}
}

func (s *MFN_M05_MF_LOCATION) MessageTypeSubStructName() string {
	return "MFN_M05_MF_LOCATION"
}
// MFN_M05_MF_LOC_DEPT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M05_MF_LOC_DEPT struct {
	LDP *LDP `hl7:"true,LDP"`
	LCH []LCH `hl7:"false,LCH"`
	LCC []LCC `hl7:"false,LCC"`
	Other []interface{}
}

func (s *MFN_M05_MF_LOC_DEPT) MessageTypeSubStructName() string {
	return "MFN_M05_MF_LOC_DEPT"
}
// MFN_M06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M06 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_CDM []MFN_M06_MF_CDM `hl7:"true,MF_CDM"`
	Other []interface{}
}

func (s *MFN_M06) MessageTypeName() string {
	return "MFN_M06"
}
// MFN_M06_MF_CDM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M06_MF_CDM struct {
	MFE *MFE `hl7:"true,MFE"`
	CDM *CDM `hl7:"true,CDM"`
	PRC []PRC `hl7:"false,PRC"`
	Other []interface{}
}

func (s *MFN_M06_MF_CDM) MessageTypeSubStructName() string {
	return "MFN_M06_MF_CDM"
}
// MFN_M07 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M07 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_CLIN_STUDY []MFN_M07_MF_CLIN_STUDY `hl7:"true,MF_CLIN_STUDY"`
	Other []interface{}
}

func (s *MFN_M07) MessageTypeName() string {
	return "MFN_M07"
}
// MFN_M07_MF_CLIN_STUDY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M07_MF_CLIN_STUDY struct {
	MFE *MFE `hl7:"true,MFE"`
	CM0 *CM0 `hl7:"true,CM0"`
	MF_PHASE_SCHED_DETAIL []MFN_M07_MF_PHASE_SCHED_DETAIL `hl7:"false,MF_PHASE_SCHED_DETAIL"`
	Other []interface{}
}

func (s *MFN_M07_MF_CLIN_STUDY) MessageTypeSubStructName() string {
	return "MFN_M07_MF_CLIN_STUDY"
}
// MFN_M07_MF_PHASE_SCHED_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M07_MF_PHASE_SCHED_DETAIL struct {
	CM1 *CM1 `hl7:"true,CM1"`
	CM2 []CM2 `hl7:"false,CM2"`
	Other []interface{}
}

func (s *MFN_M07_MF_PHASE_SCHED_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M07_MF_PHASE_SCHED_DETAIL"
}
// MFN_M08 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M08 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST_NUMERIC []MFN_M08_MF_TEST_NUMERIC `hl7:"true,MF_TEST_NUMERIC"`
	Other []interface{}
}

func (s *MFN_M08) MessageTypeName() string {
	return "MFN_M08"
}
// MFN_M08_MF_NUMERIC_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M08_MF_NUMERIC_OBSERVATION struct {
	OM2 *OM2 `hl7:"false,OM2"`
	OM3 *OM3 `hl7:"false,OM3"`
	OM4 *OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M08_MF_NUMERIC_OBSERVATION) MessageTypeSubStructName() string {
	return "MFN_M08_MF_NUMERIC_OBSERVATION"
}
// MFN_M08_MF_TEST_NUMERIC represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M08_MF_TEST_NUMERIC struct {
	MFE *MFE `hl7:"true,MFE"`
	OM1 *OM1 `hl7:"true,OM1"`
	MF_NUMERIC_OBSERVATION *MFN_M08_MF_NUMERIC_OBSERVATION `hl7:"false,MF_NUMERIC_OBSERVATION"`
	Other []interface{}
}

func (s *MFN_M08_MF_TEST_NUMERIC) MessageTypeSubStructName() string {
	return "MFN_M08_MF_TEST_NUMERIC"
}
// MFN_M09 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M09 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST_CATEGORICAL []MFN_M09_MF_TEST_CATEGORICAL `hl7:"true,MF_TEST_CATEGORICAL"`
	Other []interface{}
}

func (s *MFN_M09) MessageTypeName() string {
	return "MFN_M09"
}
// MFN_M09_MF_TEST_CATEGORICAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M09_MF_TEST_CATEGORICAL struct {
	MFE *MFE `hl7:"true,MFE"`
	MF_TEST_CAT_DETAIL *MFN_M09_MF_TEST_CAT_DETAIL `hl7:"false,MF_TEST_CAT_DETAIL"`
	Other []interface{}
}

func (s *MFN_M09_MF_TEST_CATEGORICAL) MessageTypeSubStructName() string {
	return "MFN_M09_MF_TEST_CATEGORICAL"
}
// MFN_M09_MF_TEST_CAT_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M09_MF_TEST_CAT_DETAIL struct {
	OM3 *OM3 `hl7:"true,OM3"`
	OM4 []OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M09_MF_TEST_CAT_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M09_MF_TEST_CAT_DETAIL"
}
// MFN_M10 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M10 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST_BATTERIES []MFN_M10_MF_TEST_BATTERIES `hl7:"true,MF_TEST_BATTERIES"`
	Other []interface{}
}

func (s *MFN_M10) MessageTypeName() string {
	return "MFN_M10"
}
// MFN_M10_MF_TEST_BATTERIES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M10_MF_TEST_BATTERIES struct {
	MF_TEST_BATT_DETAIL *MFN_M10_MF_TEST_BATT_DETAIL `hl7:"false,MF_TEST_BATT_DETAIL"`
	Other []interface{}
}

func (s *MFN_M10_MF_TEST_BATTERIES) MessageTypeSubStructName() string {
	return "MFN_M10_MF_TEST_BATTERIES"
}
// MFN_M10_MF_TEST_BATT_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M10_MF_TEST_BATT_DETAIL struct {
	OM5 *OM5 `hl7:"true,OM5"`
	OM4 []OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M10_MF_TEST_BATT_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M10_MF_TEST_BATT_DETAIL"
}
// MFN_M11 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M11 struct {
	MSH *MSH `hl7:"true,MSH"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST_CALCULATED []MFN_M11_MF_TEST_CALCULATED `hl7:"true,MF_TEST_CALCULATED"`
	Other []interface{}
}

func (s *MFN_M11) MessageTypeName() string {
	return "MFN_M11"
}
// MFN_M11_MF_TEST_CALCULATED represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M11_MF_TEST_CALCULATED struct {
	MFE *MFE `hl7:"true,MFE"`
	OM1 *OM1 `hl7:"true,OM1"`
	MF_TEST_CALC_DETAIL *MFN_M11_MF_TEST_CALC_DETAIL `hl7:"false,MF_TEST_CALC_DETAIL"`
	Other []interface{}
}

func (s *MFN_M11_MF_TEST_CALCULATED) MessageTypeSubStructName() string {
	return "MFN_M11_MF_TEST_CALCULATED"
}
// MFN_M11_MF_TEST_CALC_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M11_MF_TEST_CALC_DETAIL struct {
	OM6 *OM6 `hl7:"true,OM6"`
	OM2 *OM2 `hl7:"true,OM2"`
	Other []interface{}
}

func (s *MFN_M11_MF_TEST_CALC_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M11_MF_TEST_CALC_DETAIL"
}
// MFQ_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M01 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M01) MessageTypeName() string {
	return "MFQ_M01"
}
// MFQ_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M02 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M02) MessageTypeName() string {
	return "MFQ_M02"
}
// MFQ_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M03 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M03) MessageTypeName() string {
	return "MFQ_M03"
}
// MFR_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	MFI *MFI `hl7:"true,MFI"`
	MF []MFR_M01_MF `hl7:"true,MF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFR_M01) MessageTypeName() string {
	return "MFR_M01"
}
// MFR_M01_MF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M01_MF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M01_MF) MessageTypeSubStructName() string {
	return "MFR_M01_MF"
}
// MFR_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	MFI *MFI `hl7:"true,MFI"`
	MF_STAFF []MFR_M02_MF_STAFF `hl7:"true,MF_STAFF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFR_M02) MessageTypeName() string {
	return "MFR_M02"
}
// MFR_M02_MF_STAFF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M02_MF_STAFF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M02_MF_STAFF) MessageTypeSubStructName() string {
	return "MFR_M02_MF_STAFF"
}
// MFR_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	MFI *MFI `hl7:"true,MFI"`
	MF_TEST []MFR_M03_MF_TEST `hl7:"true,MF_TEST"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFR_M03) MessageTypeName() string {
	return "MFR_M03"
}
// MFR_M03_MF_TEST represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M03_MF_TEST struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M03_MF_TEST) MessageTypeSubStructName() string {
	return "MFR_M03_MF_TEST"
}
// NMD_N01_APP_STATS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_APP_STATS struct {
	NST *NST `hl7:"true,NST"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_APP_STATS) MessageTypeSubStructName() string {
	return "NMD_N01_APP_STATS"
}
// NMD_N01_APP_STATUS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_APP_STATUS struct {
	NSC *NSC `hl7:"true,NSC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_APP_STATUS) MessageTypeSubStructName() string {
	return "NMD_N01_APP_STATUS"
}
// NMD_N01_CLOCK represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_CLOCK struct {
	NCK *NCK `hl7:"true,NCK"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_CLOCK) MessageTypeSubStructName() string {
	return "NMD_N01_CLOCK"
}
// NMD_N01_CLOCK_AND_STATS_WITH_NOTES represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES struct {
	CLOCK *NMD_N01_CLOCK `hl7:"false,CLOCK"`
	APP_STATS *NMD_N01_APP_STATS `hl7:"false,APP_STATS"`
	APP_STATUS *NMD_N01_APP_STATUS `hl7:"false,APP_STATUS"`
	Other []interface{}
}

func (s *NMD_N01_CLOCK_AND_STATS_WITH_NOTES) MessageTypeSubStructName() string {
	return "NMD_N01_CLOCK_AND_STATS_WITH_NOTES"
}
// NMD_N01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01 struct {
	MSH *MSH `hl7:"true,MSH"`
	CLOCK_AND_STATS_WITH_NOTES []NMD_N01_CLOCK_AND_STATS_WITH_NOTES `hl7:"true,CLOCK_AND_STATS_WITH_NOTES"`
	Other []interface{}
}

func (s *NMD_N01) MessageTypeName() string {
	return "NMD_N01"
}
// NMQ_N02_CLOCK_AND_STATISTICS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02_CLOCK_AND_STATISTICS struct {
	NCK *NCK `hl7:"false,NCK"`
	NST *NST `hl7:"false,NST"`
	NSC *NSC `hl7:"false,NSC"`
	Other []interface{}
}

func (s *NMQ_N02_CLOCK_AND_STATISTICS) MessageTypeSubStructName() string {
	return "NMQ_N02_CLOCK_AND_STATISTICS"
}
// NMQ_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRY_WITH_DETAIL *NMQ_N02_QRY_WITH_DETAIL `hl7:"false,QRY_WITH_DETAIL"`
	CLOCK_AND_STATISTICS []NMQ_N02_CLOCK_AND_STATISTICS `hl7:"true,CLOCK_AND_STATISTICS"`
	Other []interface{}
}

func (s *NMQ_N02) MessageTypeName() string {
	return "NMQ_N02"
}
// NMQ_N02_QRY_WITH_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02_QRY_WITH_DETAIL struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *NMQ_N02_QRY_WITH_DETAIL) MessageTypeSubStructName() string {
	return "NMQ_N02_QRY_WITH_DETAIL"
}
// NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT struct {
	NCK *NCK `hl7:"false,NCK"`
	NTE1 []NTE `hl7:"false,NTE1"`
	NST *NST `hl7:"false,NST"`
	NTE2 []NTE `hl7:"false,NTE2"`
	NSC *NSC `hl7:"false,NSC"`
	NTE3 []NTE `hl7:"false,NTE3"`
	Other []interface{}
}

func (s *NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT) MessageTypeSubStructName() string {
	return "NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT"
}
// NMR_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMR_N02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"false,QRD"`
	CLOCK_AND_STATS_WITH_NOTES_ALT []NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT `hl7:"true,CLOCK_AND_STATS_WITH_NOTES_ALT"`
	Other []interface{}
}

func (s *NMR_N02) MessageTypeName() string {
	return "NMR_N02"
}
// OMD_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *OMD_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER_DIET []OMD_O01_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []OMD_O01_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other []interface{}
}

func (s *OMD_O01) MessageTypeName() string {
	return "OMD_O01"
}
// OMD_O01_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_DIET struct {
	ODS []ODS `hl7:"true,ODS"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []OMD_O01_OBSERVATION `hl7:"true,OBSERVATION"`
	Other []interface{}
}

func (s *OMD_O01_DIET) MessageTypeSubStructName() string {
	return "OMD_O01_DIET"
}
// OMD_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMD_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMD_O01_INSURANCE"
}
// OMD_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMD_O01_OBSERVATION"
}
// OMD_O01_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_ORDER_DIET struct {
	ORC *ORC `hl7:"true,ORC"`
	DIET *OMD_O01_DIET `hl7:"false,DIET"`
	Other []interface{}
}

func (s *OMD_O01_ORDER_DIET) MessageTypeSubStructName() string {
	return "OMD_O01_ORDER_DIET"
}
// OMD_O01_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_ORDER_TRAY struct {
	ORC *ORC `hl7:"true,ORC"`
	ODT []ODT `hl7:"true,ODT"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O01_ORDER_TRAY) MessageTypeSubStructName() string {
	return "OMD_O01_ORDER_TRAY"
}
// OMD_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *OMD_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []OMD_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *OMD_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMD_O01_PATIENT"
}
// OMD_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMD_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMD_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMD_O01_PATIENT_VISIT"
}
// OMN_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *OMN_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []OMN_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *OMN_O01) MessageTypeName() string {
	return "OMN_O01"
}
// OMN_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMN_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMN_O01_INSURANCE"
}
// OMN_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMN_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMN_O01_OBSERVATION"
}
// OMN_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *OMN_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG *BLG `hl7:"false,BLG"`
	Other []interface{}
}

func (s *OMN_O01_ORDER) MessageTypeSubStructName() string {
	return "OMN_O01_ORDER"
}
// OMN_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_ORDER_DETAIL struct {
	RQD *RQD `hl7:"true,RQD"`
	RQ1 *RQ1 `hl7:"false,RQ1"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []OMN_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *OMN_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "OMN_O01_ORDER_DETAIL"
}
// OMN_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *OMN_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []OMN_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *OMN_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMN_O01_PATIENT"
}
// OMN_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMN_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMN_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMN_O01_PATIENT_VISIT"
}
// OMS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *OMS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []OMS_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *OMS_O01) MessageTypeName() string {
	return "OMS_O01"
}
// OMS_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMS_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMS_O01_INSURANCE"
}
// OMS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMS_O01_OBSERVATION"
}
// OMS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *OMS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG *BLG `hl7:"false,BLG"`
	Other []interface{}
}

func (s *OMS_O01_ORDER) MessageTypeSubStructName() string {
	return "OMS_O01_ORDER"
}
// OMS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_ORDER_DETAIL struct {
	RQD *RQD `hl7:"true,RQD"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []OMS_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *OMS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "OMS_O01_ORDER_DETAIL"
}
// OMS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *OMS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []OMS_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *OMS_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMS_O01_PATIENT"
}
// OMS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OMS_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMS_O01_PATIENT_VISIT"
}
// ORD_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORD_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *ORD_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *ORD_O02) MessageTypeName() string {
	return "ORD_O02"
}
// ORD_O02_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORD_O02_ORDER_DIET struct {
	ORC *ORC `hl7:"true,ORC"`
	ODS []ODS `hl7:"false,ODS"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_ORDER_DIET) MessageTypeSubStructName() string {
	return "ORD_O02_ORDER_DIET"
}
// ORD_O02_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORD_O02_ORDER_TRAY struct {
	ORC *ORC `hl7:"true,ORC"`
	ODT []ODT `hl7:"false,ODT"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_ORDER_TRAY) MessageTypeSubStructName() string {
	return "ORD_O02_ORDER_TRAY"
}
// ORD_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORD_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORD_O02_PATIENT"
}
// ORD_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORD_O02_RESPONSE struct {
	PATIENT *ORD_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER_DIET []ORD_O02_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []ORD_O02_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other []interface{}
}

func (s *ORD_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORD_O02_RESPONSE"
}
// ORF_R04 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORF_R04 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	QUERY_RESPONSE []ORF_R04_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ORF_R04) MessageTypeName() string {
	return "ORF_R04"
}
// ORF_R04_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORF_R04_OBSERVATION struct {
	OBX *OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_OBSERVATION) MessageTypeSubStructName() string {
	return "ORF_R04_OBSERVATION"
}
// ORF_R04_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORF_R04_ORDER struct {
	ORC *ORC `hl7:"false,ORC"`
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []ORF_R04_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *ORF_R04_ORDER) MessageTypeSubStructName() string {
	return "ORF_R04_ORDER"
}
// ORF_R04_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORF_R04_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_PATIENT) MessageTypeSubStructName() string {
	return "ORF_R04_PATIENT"
}
// ORF_R04_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORF_R04_QUERY_RESPONSE struct {
	PATIENT *ORF_R04_PATIENT `hl7:"false,PATIENT"`
	ORDER []ORF_R04_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ORF_R04_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ORF_R04_QUERY_RESPONSE"
}
// ORM_O01_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_CHOICE struct {
	// Only one of the following fields will be set.
	OBR *OBR `hl7:"true,OBR"`
	RQD *RQD `hl7:"true,RQD"`
	RQ1 *RQ1 `hl7:"true,RQ1"`
	RXO *RXO `hl7:"true,RXO"`
	ODS *ODS `hl7:"true,ODS"`
	ODT *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *ORM_O01_CHOICE) MessageTypeSubStructName() string {
	return "ORM_O01_CHOICE"
}
// ORM_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *ORM_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []ORM_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ORM_O01) MessageTypeName() string {
	return "ORM_O01"
}
// ORM_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ORM_O01_INSURANCE) MessageTypeSubStructName() string {
	return "ORM_O01_INSURANCE"
}
// ORM_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORM_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORM_O01_OBSERVATION"
}
// ORM_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *ORM_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	CTI *CTI `hl7:"false,CTI"`
	BLG *BLG `hl7:"false,BLG"`
	Other []interface{}
}

func (s *ORM_O01_ORDER) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER"
}
// ORM_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_ORDER_DETAIL struct {
	CHOICE *ORM_O01_CHOICE `hl7:"true,CHOICE"`
	NTE []NTE `hl7:"false,NTE"`
	DG1 []DG1 `hl7:"false,DG1"`
	OBSERVATION []ORM_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *ORM_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER_DETAIL"
}
// ORM_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *ORM_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []ORM_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *ORM_O01_PATIENT) MessageTypeSubStructName() string {
	return "ORM_O01_PATIENT"
}
// ORM_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORM_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORM_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "ORM_O01_PATIENT_VISIT"
}
// ORN_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORN_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *ORN_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *ORN_O02) MessageTypeName() string {
	return "ORN_O02"
}
// ORN_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORN_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	RQD *RQD `hl7:"true,RQD"`
	RQ1 *RQ1 `hl7:"false,RQ1"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O02_ORDER) MessageTypeSubStructName() string {
	return "ORN_O02_ORDER"
}
// ORN_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORN_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORN_O02_PATIENT"
}
// ORN_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORN_O02_RESPONSE struct {
	PATIENT *ORN_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER []ORN_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ORN_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORN_O02_RESPONSE"
}
// ORR_O02_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORR_O02_CHOICE struct {
	// Only one of the following fields will be set.
	OBR *OBR `hl7:"true,OBR"`
	RQD *RQD `hl7:"true,RQD"`
	RQ1 *RQ1 `hl7:"true,RQ1"`
	RXO *RXO `hl7:"true,RXO"`
	ODS *ODS `hl7:"true,ODS"`
	ODT *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *ORR_O02_CHOICE) MessageTypeSubStructName() string {
	return "ORR_O02_CHOICE"
}
// ORR_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORR_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *ORR_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *ORR_O02) MessageTypeName() string {
	return "ORR_O02"
}
// ORR_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORR_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	CHOICE *ORR_O02_CHOICE `hl7:"true,CHOICE"`
	NTE []NTE `hl7:"false,NTE"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *ORR_O02_ORDER) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER"
}
// ORR_O02_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02_ORDER_DETAIL struct {
	CHOICE *ORR_O02_CHOICE `hl7:"true,CHOICE"`
	Other []interface{}
}

func (s *ORR_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER_DETAIL"
}
// ORR_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORR_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORR_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORR_O02_PATIENT"
}
// ORR_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORR_O02_RESPONSE struct {
	PATIENT *ORR_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER []ORR_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ORR_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORR_O02_RESPONSE"
}
// ORU_R01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01 struct {
	MSH *MSH `hl7:"true,MSH"`
	RESPONSE []ORU_R01_RESPONSE `hl7:"true,RESPONSE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ORU_R01) MessageTypeName() string {
	return "ORU_R01"
}
// ORU_R01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_OBSERVATION struct {
	OBX *OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_OBSERVATION"
}
// ORU_R01_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_ORDER_OBSERVATION struct {
	ORC *ORC `hl7:"false,ORC"`
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []ORU_R01_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *ORU_R01_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_ORDER_OBSERVATION"
}
// ORU_R01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	VISIT *ORU_R01_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *ORU_R01_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT"
}
// ORU_R01_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01_PATIENT_RESULT struct {
	PATIENT *ORU_R01_PATIENT `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R01_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *ORU_R01_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT_RESULT"
}
// ORU_R01_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_RESPONSE struct {
	PATIENT *ORU_R01_PATIENT `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R01_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *ORU_R01_RESPONSE) MessageTypeSubStructName() string {
	return "ORU_R01_RESPONSE"
}
// ORU_R01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORU_R01_VISIT) MessageTypeSubStructName() string {
	return "ORU_R01_VISIT"
}
// ORU_R03 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03 struct {
	MSH *MSH `hl7:"true,MSH"`
	PATIENT_RESULT []ORU_R03_PATIENT_RESULT `hl7:"true,PATIENT_RESULT"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ORU_R03) MessageTypeName() string {
	return "ORU_R03"
}
// ORU_R03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_OBSERVATION struct {
	OBX *OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R03_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R03_OBSERVATION"
}
// ORU_R03_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_ORDER_OBSERVATION struct {
	ORC *ORC `hl7:"false,ORC"`
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []ORU_R03_OBSERVATION `hl7:"true,OBSERVATION"`
	Other []interface{}
}

func (s *ORU_R03_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R03_ORDER_OBSERVATION"
}
// ORU_R03_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	PV1 *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ORU_R03_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R03_PATIENT"
}
// ORU_R03_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_PATIENT_RESULT struct {
	PATIENT *ORU_R03_PATIENT `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R03_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *ORU_R03_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R03_PATIENT_RESULT"
}
// ORU_R32 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32 struct {
	MSH *MSH `hl7:"true,MSH"`
	PATIENT_RESULT []ORU_R32_PATIENT_RESULT `hl7:"true,PATIENT_RESULT"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ORU_R32) MessageTypeName() string {
	return "ORU_R32"
}
// ORU_R32_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_OBSERVATION struct {
	OBX *OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R32_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_OBSERVATION"
}
// ORU_R32_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_ORDER_OBSERVATION struct {
	ORC *ORC `hl7:"false,ORC"`
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []ORU_R32_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *ORU_R32_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_ORDER_OBSERVATION"
}
// ORU_R32_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NK1 []NK1 `hl7:"false,NK1"`
	NTE []NTE `hl7:"false,NTE"`
	VISIT *ORU_R32_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *ORU_R32_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT"
}
// ORU_R32_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_PATIENT_RESULT struct {
	PATIENT *ORU_R32_PATIENT `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R32_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *ORU_R32_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT_RESULT"
}
// ORU_R32_VISIT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORU_R32_VISIT) MessageTypeSubStructName() string {
	return "ORU_R32_VISIT"
}
// OSQ_Q06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OSQ_Q06 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *OSQ_Q06) MessageTypeName() string {
	return "OSQ_Q06"
}
// OSR_Q06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OSR_Q06 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	RESPONSE *OSR_Q06_RESPONSE `hl7:"false,RESPONSE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *OSR_Q06) MessageTypeName() string {
	return "OSR_Q06"
}
// OSR_Q06_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OSR_Q06_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	OBR *OBR `hl7:"false,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *OSR_Q06_ORDER) MessageTypeSubStructName() string {
	return "OSR_Q06_ORDER"
}
// OSR_Q06_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OSR_Q06_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OSR_Q06_PATIENT) MessageTypeSubStructName() string {
	return "OSR_Q06_PATIENT"
}
// OSR_Q06_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type OSR_Q06_RESPONSE struct {
	PATIENT *OSR_Q06_PATIENT `hl7:"false,PATIENT"`
	ORDER []OSR_Q06_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *OSR_Q06_RESPONSE) MessageTypeSubStructName() string {
	return "OSR_Q06_RESPONSE"
}
// PEX_P07_ASSOCIATED_PERSON represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_ASSOCIATED_PERSON struct {
	NK1 *NK1 `hl7:"true,NK1"`
	ASSOCIATED_RX_ORDER *PEX_P07_ASSOCIATED_RX_ORDER `hl7:"false,ASSOCIATED_RX_ORDER"`
	ASSOCIATED_RX_ADMIN []PEX_P07_ASSOCIATED_RX_ADMIN `hl7:"false,ASSOCIATED_RX_ADMIN"`
	PRB []PRB `hl7:"false,PRB"`
	OBX []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *PEX_P07_ASSOCIATED_PERSON) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_PERSON"
}
// PEX_P07_ASSOCIATED_RX_ADMIN represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_ASSOCIATED_RX_ADMIN struct {
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_ASSOCIATED_RX_ADMIN) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_RX_ADMIN"
}
// PEX_P07_ASSOCIATED_RX_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_ASSOCIATED_RX_ORDER struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_ASSOCIATED_RX_ORDER) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_RX_ORDER"
}
// PEX_P07 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	VISIT *PEX_P07_VISIT `hl7:"false,VISIT"`
	EXPERIENCE []PEX_P07_EXPERIENCE `hl7:"true,EXPERIENCE"`
	Other []interface{}
}

func (s *PEX_P07) MessageTypeName() string {
	return "PEX_P07"
}
// PEX_P07_EXPERIENCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_EXPERIENCE struct {
	PES *PES `hl7:"true,PES"`
	PEX_OBSERVATION []PEX_P07_PEX_OBSERVATION `hl7:"true,PEX_OBSERVATION"`
	Other []interface{}
}

func (s *PEX_P07_EXPERIENCE) MessageTypeSubStructName() string {
	return "PEX_P07_EXPERIENCE"
}
// PEX_P07_PEX_CAUSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_PEX_CAUSE struct {
	PCR *PCR `hl7:"true,PCR"`
	RX_ORDER *PEX_P07_RX_ORDER `hl7:"false,RX_ORDER"`
	RX_ADMINISTRATION []PEX_P07_RX_ADMINISTRATION `hl7:"false,RX_ADMINISTRATION"`
	PRB []PRB `hl7:"false,PRB"`
	OBX []OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	ASSOCIATED_PERSON *PEX_P07_ASSOCIATED_PERSON `hl7:"false,ASSOCIATED_PERSON"`
	STUDY []PEX_P07_STUDY `hl7:"false,STUDY"`
	Other []interface{}
}

func (s *PEX_P07_PEX_CAUSE) MessageTypeSubStructName() string {
	return "PEX_P07_PEX_CAUSE"
}
// PEX_P07_PEX_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_PEX_OBSERVATION struct {
	PEO *PEO `hl7:"true,PEO"`
	PEX_CAUSE []PEX_P07_PEX_CAUSE `hl7:"true,PEX_CAUSE"`
	Other []interface{}
}

func (s *PEX_P07_PEX_OBSERVATION) MessageTypeSubStructName() string {
	return "PEX_P07_PEX_OBSERVATION"
}
// PEX_P07_RX_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_RX_ADMINISTRATION struct {
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_RX_ADMINISTRATION) MessageTypeSubStructName() string {
	return "PEX_P07_RX_ADMINISTRATION"
}
// PEX_P07_RX_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_RX_ORDER struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_RX_ORDER) MessageTypeSubStructName() string {
	return "PEX_P07_RX_ORDER"
}
// PEX_P07_STUDY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_STUDY struct {
	CSR *CSR `hl7:"true,CSR"`
	CSP []CSP `hl7:"false,CSP"`
	Other []interface{}
}

func (s *PEX_P07_STUDY) MessageTypeSubStructName() string {
	return "PEX_P07_STUDY"
}
// PEX_P07_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PEX_P07_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PEX_P07_VISIT) MessageTypeSubStructName() string {
	return "PEX_P07_VISIT"
}
// PGL_PC6 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6 struct {
	MSH *MSH `hl7:"true,MSH"`
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PGL_PC6_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	GOAL []PGL_PC6_GOAL `hl7:"true,GOAL"`
	Other []interface{}
}

func (s *PGL_PC6) MessageTypeName() string {
	return "PGL_PC6"
}
// PGL_PC6_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PGL_PC6_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	PATHWAY []PGL_PC6_PATHWAY `hl7:"false,PATHWAY"`
	OBSERVATION []PGL_PC6_OBSERVATION `hl7:"false,OBSERVATION"`
	PROBLEM []PGL_PC6_PROBLEM `hl7:"false,PROBLEM"`
	ORDER []PGL_PC6_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PGL_PC6_GOAL) MessageTypeSubStructName() string {
	return "PGL_PC6_GOAL"
}
// PGL_PC6_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PGL_PC6_GOAL_ROLE"
}
// PGL_PC6_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PGL_PC6_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_OBSERVATION"
}
// PGL_PC6_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PGL_PC6_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PGL_PC6_ORDER) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER"
}
// PGL_PC6_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PGL_PC6_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PGL_PC6_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER_DETAIL"
}
// PGL_PC6_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER_OBSERVATION"
}
// PGL_PC6_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_PATHWAY) MessageTypeSubStructName() string {
	return "PGL_PC6_PATHWAY"
}
// PGL_PC6_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PGL_PC6_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PGL_PC6_PATIENT_VISIT"
}
// PGL_PC6_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PGL_PC6_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PGL_PC6_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other []interface{}
}

func (s *PGL_PC6_PROBLEM) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM"
}
// PGL_PC6_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PGL_PC6_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM_OBSERVATION"
}
// PGL_PC6_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PGL_PC6_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM_ROLE"
}
// PIN_I07 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PIN_I07 struct {
	MSH *MSH `hl7:"true,MSH"`
	PROVIDER []PIN_I07_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *PIN_I07_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PIN_I07) MessageTypeName() string {
	return "PIN_I07"
}
// PIN_I07_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PIN_I07_GUARANTOR_INSURANCE struct {
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []PIN_I07_INSURANCE `hl7:"true,INSURANCE"`
	Other []interface{}
}

func (s *PIN_I07_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "PIN_I07_GUARANTOR_INSURANCE"
}
// PIN_I07_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PIN_I07_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *PIN_I07_INSURANCE) MessageTypeSubStructName() string {
	return "PIN_I07_INSURANCE"
}
// PIN_I07_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PIN_I07_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *PIN_I07_PROVIDER) MessageTypeSubStructName() string {
	return "PIN_I07_PROVIDER"
}
// PPG_PCG represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG struct {
	MSH *MSH `hl7:"true,MSH"`
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PPG_PCG_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY []PPG_PCG_PATHWAY `hl7:"true,PATHWAY"`
	Other []interface{}
}

func (s *PPG_PCG) MessageTypeName() string {
	return "PPG_PCG"
}
// PPG_PCG_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PPG_PCG_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPG_PCG_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM []PPG_PCG_PROBLEM `hl7:"false,PROBLEM"`
	ORDER []PPG_PCG_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PPG_PCG_GOAL) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL"
}
// PPG_PCG_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPG_PCG_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL_OBSERVATION"
}
// PPG_PCG_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL_ROLE"
}
// PPG_PCG_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PPG_PCG_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PPG_PCG_ORDER) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER"
}
// PPG_PCG_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPG_PCG_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PPG_PCG_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER_DETAIL"
}
// PPG_PCG_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER_OBSERVATION"
}
// PPG_PCG_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PATHWAY_ROLE []PPG_PCG_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	GOAL []PPG_PCG_GOAL `hl7:"false,GOAL"`
	Other []interface{}
}

func (s *PPG_PCG_PATHWAY) MessageTypeSubStructName() string {
	return "PPG_PCG_PATHWAY"
}
// PPG_PCG_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PATHWAY_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_PATHWAY_ROLE"
}
// PPG_PCG_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPG_PCG_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPG_PCG_PATIENT_VISIT"
}
// PPG_PCG_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PPG_PCG_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPG_PCG_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other []interface{}
}

func (s *PPG_PCG_PROBLEM) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM"
}
// PPG_PCG_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPG_PCG_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM_OBSERVATION"
}
// PPG_PCG_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPG_PCG_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM_ROLE"
}
// PPP_PCB represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB struct {
	MSH *MSH `hl7:"true,MSH"`
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PPP_PCB_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY []PPP_PCB_PATHWAY `hl7:"true,PATHWAY"`
	Other []interface{}
}

func (s *PPP_PCB) MessageTypeName() string {
	return "PPP_PCB"
}
// PPP_PCB_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PPP_PCB_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPP_PCB_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other []interface{}
}

func (s *PPP_PCB_GOAL) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL"
}
// PPP_PCB_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPP_PCB_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL_OBSERVATION"
}
// PPP_PCB_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL_ROLE"
}
// PPP_PCB_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PPP_PCB_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PPP_PCB_ORDER) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER"
}
// PPP_PCB_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPP_PCB_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PPP_PCB_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER_DETAIL"
}
// PPP_PCB_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER_OBSERVATION"
}
// PPP_PCB_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PATHWAY_ROLE []PPP_PCB_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	PROBLEM []PPP_PCB_PROBLEM `hl7:"false,PROBLEM"`
	Other []interface{}
}

func (s *PPP_PCB_PATHWAY) MessageTypeSubStructName() string {
	return "PPP_PCB_PATHWAY"
}
// PPP_PCB_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PATHWAY_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_PATHWAY_ROLE"
}
// PPP_PCB_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPP_PCB_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPP_PCB_PATIENT_VISIT"
}
// PPP_PCB_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PPP_PCB_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPP_PCB_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL []PPP_PCB_GOAL `hl7:"false,GOAL"`
	ORDER []PPP_PCB_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PPP_PCB_PROBLEM) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM"
}
// PPP_PCB_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPP_PCB_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM_OBSERVATION"
}
// PPP_PCB_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPP_PCB_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM_ROLE"
}
// PPR_PC1 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1 struct {
	MSH *MSH `hl7:"true,MSH"`
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PPR_PC1_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PROBLEM []PPR_PC1_PROBLEM `hl7:"true,PROBLEM"`
	Other []interface{}
}

func (s *PPR_PC1) MessageTypeName() string {
	return "PPR_PC1"
}
// PPR_PC1_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PPR_PC1_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPR_PC1_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other []interface{}
}

func (s *PPR_PC1_GOAL) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL"
}
// PPR_PC1_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPR_PC1_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL_OBSERVATION"
}
// PPR_PC1_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL_ROLE"
}
// PPR_PC1_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PPR_PC1_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PPR_PC1_ORDER) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER"
}
// PPR_PC1_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPR_PC1_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PPR_PC1_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER_DETAIL"
}
// PPR_PC1_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER_OBSERVATION"
}
// PPR_PC1_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_PATHWAY) MessageTypeSubStructName() string {
	return "PPR_PC1_PATHWAY"
}
// PPR_PC1_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPR_PC1_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPR_PC1_PATIENT_VISIT"
}
// PPR_PC1_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PPR_PC1_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PATHWAY []PPR_PC1_PATHWAY `hl7:"false,PATHWAY"`
	PROBLEM_OBSERVATION []PPR_PC1_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL []PPR_PC1_GOAL `hl7:"false,GOAL"`
	ORDER []PPR_PC1_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PPR_PC1_PROBLEM) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM"
}
// PPR_PC1_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPR_PC1_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM_OBSERVATION"
}
// PPR_PC1_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPR_PC1_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM_ROLE"
}
// PPT_PCL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	PATIENT []PPT_PCL_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *PPT_PCL) MessageTypeName() string {
	return "PPT_PCL"
}
// PPT_PCL_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PPT_PCL_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPT_PCL_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM []PPT_PCL_PROBLEM `hl7:"false,PROBLEM"`
	ORDER []PPT_PCL_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PPT_PCL_GOAL) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL"
}
// PPT_PCL_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPT_PCL_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL_OBSERVATION"
}
// PPT_PCL_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL_ROLE"
}
// PPT_PCL_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PPT_PCL_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PPT_PCL_ORDER) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER"
}
// PPT_PCL_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPT_PCL_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PPT_PCL_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER_DETAIL"
}
// PPT_PCL_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER_OBSERVATION"
}
// PPT_PCL_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PATHWAY_ROLE []PPT_PCL_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	GOAL []PPT_PCL_GOAL `hl7:"false,GOAL"`
	Other []interface{}
}

func (s *PPT_PCL_PATHWAY) MessageTypeSubStructName() string {
	return "PPT_PCL_PATHWAY"
}
// PPT_PCL_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PATHWAY_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_PATHWAY_ROLE"
}
// PPT_PCL_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PPT_PCL_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY []PPT_PCL_PATHWAY `hl7:"true,PATHWAY"`
	Other []interface{}
}

func (s *PPT_PCL_PATIENT) MessageTypeSubStructName() string {
	return "PPT_PCL_PATIENT"
}
// PPT_PCL_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPT_PCL_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPT_PCL_PATIENT_VISIT"
}
// PPT_PCL_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PPT_PCL_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPT_PCL_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other []interface{}
}

func (s *PPT_PCL_PROBLEM) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM"
}
// PPT_PCL_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPT_PCL_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM_OBSERVATION"
}
// PPT_PCL_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPT_PCL_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM_ROLE"
}
// PPV_PCA represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	PATIENT []PPV_PCA_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *PPV_PCA) MessageTypeName() string {
	return "PPV_PCA"
}
// PPV_PCA_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PPV_PCA_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_PATHWAY []PPV_PCA_GOAL_PATHWAY `hl7:"false,GOAL_PATHWAY"`
	GOAL_OBSERVATION []PPV_PCA_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM []PPV_PCA_PROBLEM `hl7:"false,PROBLEM"`
	ORDER []PPV_PCA_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL"
}
// PPV_PCA_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_OBSERVATION"
}
// PPV_PCA_GOAL_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_GOAL_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_PATHWAY) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_PATHWAY"
}
// PPV_PCA_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_ROLE"
}
// PPV_PCA_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PPV_PCA_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PPV_PCA_ORDER) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER"
}
// PPV_PCA_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPV_PCA_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PPV_PCA_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER_DETAIL"
}
// PPV_PCA_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER_OBSERVATION"
}
// PPV_PCA_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PPV_PCA_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	GOAL []PPV_PCA_GOAL `hl7:"true,GOAL"`
	Other []interface{}
}

func (s *PPV_PCA_PATIENT) MessageTypeSubStructName() string {
	return "PPV_PCA_PATIENT"
}
// PPV_PCA_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPV_PCA_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPV_PCA_PATIENT_VISIT"
}
// PPV_PCA_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PPV_PCA_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPV_PCA_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other []interface{}
}

func (s *PPV_PCA_PROBLEM) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM"
}
// PPV_PCA_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPV_PCA_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM_OBSERVATION"
}
// PPV_PCA_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PPV_PCA_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM_ROLE"
}
// PRR_PC5 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	PATIENT []PRR_PC5_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *PRR_PC5) MessageTypeName() string {
	return "PRR_PC5"
}
// PRR_PC5_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PRR_PC5_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PRR_PC5_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other []interface{}
}

func (s *PRR_PC5_GOAL) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL"
}
// PRR_PC5_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PRR_PC5_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL_OBSERVATION"
}
// PRR_PC5_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL_ROLE"
}
// PRR_PC5_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PRR_PC5_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PRR_PC5_ORDER) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER"
}
// PRR_PC5_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PRR_PC5_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PRR_PC5_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER_DETAIL"
}
// PRR_PC5_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER_OBSERVATION"
}
// PRR_PC5_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PRR_PC5_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PROBLEM []PRR_PC5_PROBLEM `hl7:"true,PROBLEM"`
	Other []interface{}
}

func (s *PRR_PC5_PATIENT) MessageTypeSubStructName() string {
	return "PRR_PC5_PATIENT"
}
// PRR_PC5_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PRR_PC5_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PRR_PC5_PATIENT_VISIT"
}
// PRR_PC5_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PRR_PC5_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_PATHWAY []PRR_PC5_PROBLEM_PATHWAY `hl7:"false,PROBLEM_PATHWAY"`
	PROBLEM_OBSERVATION []PRR_PC5_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL []PRR_PC5_GOAL `hl7:"false,GOAL"`
	ORDER []PRR_PC5_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM"
}
// PRR_PC5_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_OBSERVATION"
}
// PRR_PC5_PROBLEM_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PROBLEM_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_PATHWAY) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_PATHWAY"
}
// PRR_PC5_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PRR_PC5_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_ROLE"
}
// PTR_PCF represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QRD *QRD `hl7:"true,QRD"`
	PATIENT []PTR_PCF_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *PTR_PCF) MessageTypeName() string {
	return "PTR_PCF"
}
// PTR_PCF_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_GOAL struct {
	GOL *GOL `hl7:"true,GOL"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	GOAL_ROLE []PTR_PCF_GOAL_ROLE `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PTR_PCF_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other []interface{}
}

func (s *PTR_PCF_GOAL) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL"
}
// PTR_PCF_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_GOAL_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PTR_PCF_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL_OBSERVATION"
}
// PTR_PCF_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_GOAL_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL_ROLE"
}
// PTR_PCF_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *PTR_PCF_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *PTR_PCF_ORDER) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER"
}
// PTR_PCF_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_ORDER_DETAIL struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	ORDER_OBSERVATION []PTR_PCF_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other []interface{}
}

func (s *PTR_PCF_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER_DETAIL"
}
// PTR_PCF_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_ORDER_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER_OBSERVATION"
}
// PTR_PCF_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PATHWAY struct {
	PTH *PTH `hl7:"true,PTH"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PATHWAY_ROLE []PTR_PCF_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	PROBLEM []PTR_PCF_PROBLEM `hl7:"false,PROBLEM"`
	Other []interface{}
}

func (s *PTR_PCF_PATHWAY) MessageTypeSubStructName() string {
	return "PTR_PCF_PATHWAY"
}
// PTR_PCF_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PATHWAY_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_PATHWAY_ROLE"
}
// PTR_PCF_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PATIENT_VISIT *PTR_PCF_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY []PTR_PCF_PATHWAY `hl7:"true,PATHWAY"`
	Other []interface{}
}

func (s *PTR_PCF_PATIENT) MessageTypeSubStructName() string {
	return "PTR_PCF_PATIENT"
}
// PTR_PCF_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PTR_PCF_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PTR_PCF_PATIENT_VISIT"
}
// PTR_PCF_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PROBLEM struct {
	PRB *PRB `hl7:"true,PRB"`
	NTE []NTE `hl7:"false,NTE"`
	VAR []VAR `hl7:"false,VAR"`
	PROBLEM_ROLE []PTR_PCF_PROBLEM_ROLE `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PTR_PCF_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL []PTR_PCF_GOAL `hl7:"false,GOAL"`
	ORDER []PTR_PCF_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *PTR_PCF_PROBLEM) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM"
}
// PTR_PCF_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PROBLEM_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PTR_PCF_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM_OBSERVATION"
}
// PTR_PCF_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type PTR_PCF_PROBLEM_ROLE struct {
	ROL *ROL `hl7:"true,ROL"`
	VAR []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM_ROLE"
}
// QCK_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QCK_Q02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"false,QAK"`
	Other []interface{}
}

func (s *QCK_Q02) MessageTypeName() string {
	return "QCK_Q02"
}
// QRY_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_A19 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_A19) MessageTypeName() string {
	return "QRY_A19"
}
// QRY_P04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_P04 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_P04) MessageTypeName() string {
	return "QRY_P04"
}
// QRY_PC4 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_PC4 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_PC4) MessageTypeName() string {
	return "QRY_PC4"
}
// QRY_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_Q01) MessageTypeName() string {
	return "QRY_Q01"
}
// QRY_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_Q02 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_Q02) MessageTypeName() string {
	return "QRY_Q02"
}
// QRY_R02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_R02 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"true,QRF"`
	Other []interface{}
}

func (s *QRY_R02) MessageTypeName() string {
	return "QRY_R02"
}
// QRY_T12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type QRY_T12 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_T12) MessageTypeName() string {
	return "QRY_T12"
}
// RAR_RAR represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAR_RAR struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	DEFINITION []RAR_RAR_DEFINITION `hl7:"true,DEFINITION"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RAR_RAR) MessageTypeName() string {
	return "RAR_RAR"
}
// RAR_RAR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAR_RAR_DEFINITION struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT *RAR_RAR_PATIENT `hl7:"false,PATIENT"`
	ORDER []RAR_RAR_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RAR_RAR_DEFINITION) MessageTypeSubStructName() string {
	return "RAR_RAR_DEFINITION"
}
// RAR_RAR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAR_RAR_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RAR_RAR_ENCODING) MessageTypeSubStructName() string {
	return "RAR_RAR_ENCODING"
}
// RAR_RAR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAR_RAR_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ENCODING *RAR_RAR_ENCODING `hl7:"false,ENCODING"`
	RXA []RXA `hl7:"true,RXA"`
	Other []interface{}
}

func (s *RAR_RAR_ORDER) MessageTypeSubStructName() string {
	return "RAR_RAR_ORDER"
}
// RAR_RAR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAR_RAR_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAR_RAR_PATIENT) MessageTypeSubStructName() string {
	return "RAR_RAR_PATIENT"
}
// RAS_O01_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_COMPONENTS struct {
	RXC []RXC `hl7:"true,RXC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O01_COMPONENTS) MessageTypeSubStructName() string {
	return "RAS_O01_COMPONENTS"
}
// RAS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RAS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []RAS_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RAS_O01) MessageTypeName() string {
	return "RAS_O01"
}
// RAS_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RAS_O01_ENCODING) MessageTypeSubStructName() string {
	return "RAS_O01_ENCODING"
}
// RAS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RAS_O01_OBSERVATION"
}
// RAS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RAS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING *RAS_O01_ENCODING `hl7:"false,ENCODING"`
	RXA []RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"true,RXR"`
	OBSERVATION []RAS_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	CTI []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *RAS_O01_ORDER) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER"
}
// RAS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RAS_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other []interface{}
}

func (s *RAS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER_DETAIL"
}
// RAS_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE []NTE `hl7:"true,NTE"`
	RXR []RXR `hl7:"true,RXR"`
	COMPONENTS *RAS_O01_COMPONENTS `hl7:"false,COMPONENTS"`
	Other []interface{}
}

func (s *RAS_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER_DETAIL_SUPPLEMENT"
}
// RAS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	AL1 []AL1 `hl7:"false,AL1"`
	PATIENT_VISIT *RAS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other []interface{}
}

func (s *RAS_O01_PATIENT) MessageTypeSubStructName() string {
	return "RAS_O01_PATIENT"
}
// RAS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RAS_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RAS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RAS_O01_PATIENT_VISIT"
}
// RCI_I05 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCI_I05 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PROVIDER []RCI_I05_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	OBSERVATION []RCI_I05_OBSERVATION `hl7:"false,OBSERVATION"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RCI_I05) MessageTypeName() string {
	return "RCI_I05"
}
// RCI_I05_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCI_I05_OBSERVATION struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	RESULTS []RCI_I05_RESULTS `hl7:"false,RESULTS"`
	Other []interface{}
}

func (s *RCI_I05_OBSERVATION) MessageTypeSubStructName() string {
	return "RCI_I05_OBSERVATION"
}
// RCI_I05_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCI_I05_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RCI_I05_PROVIDER) MessageTypeSubStructName() string {
	return "RCI_I05_PROVIDER"
}
// RCI_I05_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCI_I05_RESULTS struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RCI_I05_RESULTS) MessageTypeSubStructName() string {
	return "RCI_I05_RESULTS"
}
// RCL_I06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCL_I06 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PROVIDER []RCL_I06_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	NTE []NTE `hl7:"false,NTE"`
	DSP []DSP `hl7:"false,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RCL_I06) MessageTypeName() string {
	return "RCL_I06"
}
// RCL_I06_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RCL_I06_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RCL_I06_PROVIDER) MessageTypeSubStructName() string {
	return "RCL_I06_PROVIDER"
}
// RDE_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_COMPONENT struct {
	RXC []RXC `hl7:"true,RXC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDE_O01_COMPONENT"
}
// RDE_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RDE_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []RDE_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RDE_O01) MessageTypeName() string {
	return "RDE_O01"
}
// RDE_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RDE_O01_INSURANCE) MessageTypeSubStructName() string {
	return "RDE_O01_INSURANCE"
}
// RDE_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_OBSERVATION struct {
	OBX *OBX `hl7:"false,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDE_O01_OBSERVATION"
}
// RDE_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RDE_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	OBSERVATION []RDE_O01_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI *CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *RDE_O01_ORDER) MessageTypeSubStructName() string {
	return "RDE_O01_ORDER"
}
// RDE_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	NTE []NTE `hl7:"false,NTE"`
	RXR []RXR `hl7:"true,RXR"`
	COMPONENT *RDE_O01_COMPONENT `hl7:"false,COMPONENT"`
	Other []interface{}
}

func (s *RDE_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDE_O01_ORDER_DETAIL"
}
// RDE_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *RDE_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []RDE_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *RDE_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDE_O01_PATIENT"
}
// RDE_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDE_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDE_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDE_O01_PATIENT_VISIT"
}
// RDO_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_COMPONENT struct {
	RXC []RXC `hl7:"true,RXC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDO_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDO_O01_COMPONENT"
}
// RDO_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RDO_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []RDO_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RDO_O01) MessageTypeName() string {
	return "RDO_O01"
}
// RDO_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RDO_O01_INSURANCE) MessageTypeSubStructName() string {
	return "RDO_O01_INSURANCE"
}
// RDO_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDO_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDO_O01_OBSERVATION"
}
// RDO_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RDO_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG *BLG `hl7:"false,BLG"`
	Other []interface{}
}

func (s *RDO_O01_ORDER) MessageTypeSubStructName() string {
	return "RDO_O01_ORDER"
}
// RDO_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	NTE []NTE `hl7:"false,NTE"`
	RXR []RXR `hl7:"true,RXR"`
	COMPONENT *RDO_O01_COMPONENT `hl7:"false,COMPONENT"`
	OBSERVATION []RDO_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *RDO_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDO_O01_ORDER_DETAIL"
}
// RDO_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT_VISIT *RDO_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []RDO_O01_INSURANCE `hl7:"false,INSURANCE"`
	GT1 *GT1 `hl7:"false,GT1"`
	AL1 []AL1 `hl7:"false,AL1"`
	Other []interface{}
}

func (s *RDO_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDO_O01_PATIENT"
}
// RDO_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDO_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDO_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDO_O01_PATIENT_VISIT"
}
// RDR_RDR represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	DEFINITION []RDR_RDR_DEFINITION `hl7:"true,DEFINITION"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RDR_RDR) MessageTypeName() string {
	return "RDR_RDR"
}
// RDR_RDR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR_DEFINITION struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT *RDR_RDR_PATIENT `hl7:"false,PATIENT"`
	ORDER []RDR_RDR_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RDR_RDR_DEFINITION) MessageTypeSubStructName() string {
	return "RDR_RDR_DEFINITION"
}
// RDR_RDR_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR_DISPENSE struct {
	RXD *RXD `hl7:"true,RXD"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDR_RDR_DISPENSE) MessageTypeSubStructName() string {
	return "RDR_RDR_DISPENSE"
}
// RDR_RDR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR *RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDR_RDR_ENCODING) MessageTypeSubStructName() string {
	return "RDR_RDR_ENCODING"
}
// RDR_RDR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ENCODING *RDR_RDR_ENCODING `hl7:"false,ENCODING"`
	DISPENSE []RDR_RDR_DISPENSE `hl7:"true,DISPENSE"`
	Other []interface{}
}

func (s *RDR_RDR_ORDER) MessageTypeSubStructName() string {
	return "RDR_RDR_ORDER"
}
// RDR_RDR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDR_RDR_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDR_RDR_PATIENT) MessageTypeSubStructName() string {
	return "RDR_RDR_PATIENT"
}
// RDS_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_COMPONENT struct {
	RXC []RXC `hl7:"true,RXC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDS_O01_COMPONENT"
}
// RDS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RDS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []RDS_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RDS_O01) MessageTypeName() string {
	return "RDS_O01"
}
// RDS_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDS_O01_ENCODING) MessageTypeSubStructName() string {
	return "RDS_O01_ENCODING"
}
// RDS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDS_O01_OBSERVATION"
}
// RDS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RDS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING *RDS_O01_ENCODING `hl7:"false,ENCODING"`
	RXD *RXD `hl7:"true,RXD"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	OBSERVATION []RDS_O01_OBSERVATION `hl7:"true,OBSERVATION"`
	Other []interface{}
}

func (s *RDS_O01_ORDER) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER"
}
// RDS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RDS_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other []interface{}
}

func (s *RDS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER_DETAIL"
}
// RDS_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE []NTE `hl7:"true,NTE"`
	RXR []RXR `hl7:"true,RXR"`
	COMPONENT *RDS_O01_COMPONENT `hl7:"false,COMPONENT"`
	Other []interface{}
}

func (s *RDS_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER_DETAIL_SUPPLEMENT"
}
// RDS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NTE []NTE `hl7:"false,NTE"`
	AL1 []AL1 `hl7:"false,AL1"`
	PATIENT_VISIT *RDS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other []interface{}
}

func (s *RDS_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDS_O01_PATIENT"
}
// RDS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RDS_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDS_O01_PATIENT_VISIT"
}
// REF_I12_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_AUTHORIZATION struct {
	AUT *AUT `hl7:"true,AUT"`
	CTD *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_AUTHORIZATION) MessageTypeSubStructName() string {
	return "REF_I12_AUTHORIZATION"
}
// REF_I12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12 struct {
	MSH *MSH `hl7:"true,MSH"`
	RF1 *RF1 `hl7:"false,RF1"`
	AUTHORIZATION *REF_I12_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	PROVIDER []REF_I12_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []REF_I12_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	PROCEDURE []REF_I12_PROCEDURE `hl7:"false,PROCEDURE"`
	RESULTS []REF_I12_RESULTS `hl7:"false,RESULTS"`
	VISIT *REF_I12_VISIT `hl7:"false,VISIT"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *REF_I12) MessageTypeName() string {
	return "REF_I12"
}
// REF_I12_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *REF_I12_INSURANCE) MessageTypeSubStructName() string {
	return "REF_I12_INSURANCE"
}
// REF_I12_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *REF_I12_OBSERVATION) MessageTypeSubStructName() string {
	return "REF_I12_OBSERVATION"
}
// REF_I12_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	AUTHORIZATION *REF_I12_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	Other []interface{}
}

func (s *REF_I12_PROCEDURE) MessageTypeSubStructName() string {
	return "REF_I12_PROCEDURE"
}
// REF_I12_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_PROVIDER) MessageTypeSubStructName() string {
	return "REF_I12_PROVIDER"
}
// REF_I12_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_RESULTS struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []REF_I12_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *REF_I12_RESULTS) MessageTypeSubStructName() string {
	return "REF_I12_RESULTS"
}
// REF_I12_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *REF_I12_VISIT) MessageTypeSubStructName() string {
	return "REF_I12_VISIT"
}
// RER_RER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RER_RER struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	DEFINITION []RER_RER_DEFINITION `hl7:"true,DEFINITION"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RER_RER) MessageTypeName() string {
	return "RER_RER"
}
// RER_RER_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RER_RER_DEFINITION struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT *RER_RER_PATIENT `hl7:"false,PATIENT"`
	ORDER []RER_RER_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RER_RER_DEFINITION) MessageTypeSubStructName() string {
	return "RER_RER_DEFINITION"
}
// RER_RER_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RER_RER_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RER_RER_ORDER) MessageTypeSubStructName() string {
	return "RER_RER_ORDER"
}
// RER_RER_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RER_RER_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RER_RER_PATIENT) MessageTypeSubStructName() string {
	return "RER_RER_PATIENT"
}
// RGR_RGR represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGR_RGR struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	DEFINITION []RGR_RGR_DEFINITION `hl7:"true,DEFINITION"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RGR_RGR) MessageTypeName() string {
	return "RGR_RGR"
}
// RGR_RGR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGR_RGR_DEFINITION struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT *RGR_RGR_PATIENT `hl7:"false,PATIENT"`
	ORDER []RGR_RGR_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RGR_RGR_DEFINITION) MessageTypeSubStructName() string {
	return "RGR_RGR_DEFINITION"
}
// RGR_RGR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGR_RGR_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGR_RGR_ENCODING) MessageTypeSubStructName() string {
	return "RGR_RGR_ENCODING"
}
// RGR_RGR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGR_RGR_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ENCODING *RGR_RGR_ENCODING `hl7:"false,ENCODING"`
	RXG []RXG `hl7:"true,RXG"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGR_RGR_ORDER) MessageTypeSubStructName() string {
	return "RGR_RGR_ORDER"
}
// RGR_RGR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGR_RGR_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGR_RGR_PATIENT) MessageTypeSubStructName() string {
	return "RGR_RGR_PATIENT"
}
// RGV_O01_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_COMPONENTS struct {
	RXC []RXC `hl7:"true,RXC"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O01_COMPONENTS) MessageTypeSubStructName() string {
	return "RGV_O01_COMPONENTS"
}
// RGV_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01 struct {
	MSH *MSH `hl7:"true,MSH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RGV_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER []RGV_O01_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RGV_O01) MessageTypeName() string {
	return "RGV_O01"
}
// RGV_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_ENCODING struct {
	RXE *RXE `hl7:"true,RXE"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGV_O01_ENCODING) MessageTypeSubStructName() string {
	return "RGV_O01_ENCODING"
}
// RGV_O01_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_GIVE struct {
	RXG *RXG `hl7:"true,RXG"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	OBSERVATION []RGV_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *RGV_O01_GIVE) MessageTypeSubStructName() string {
	return "RGV_O01_GIVE"
}
// RGV_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RGV_O01_OBSERVATION"
}
// RGV_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RGV_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING *RGV_O01_ENCODING `hl7:"false,ENCODING"`
	GIVE []RGV_O01_GIVE `hl7:"true,GIVE"`
	Other []interface{}
}

func (s *RGV_O01_ORDER) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER"
}
// RGV_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RGV_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other []interface{}
}

func (s *RGV_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER_DETAIL"
}
// RGV_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE []NTE `hl7:"true,NTE"`
	RXR []RXR `hl7:"true,RXR"`
	COMPONENTS *RGV_O01_COMPONENTS `hl7:"false,COMPONENTS"`
	Other []interface{}
}

func (s *RGV_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER_DETAIL_SUPPLEMENT"
}
// RGV_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	AL1 []AL1 `hl7:"false,AL1"`
	PATIENT_VISIT *RGV_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other []interface{}
}

func (s *RGV_O01_PATIENT) MessageTypeSubStructName() string {
	return "RGV_O01_PATIENT"
}
// RGV_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RGV_O01_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RGV_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RGV_O01_PATIENT_VISIT"
}
// ROR_ROR represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ROR_ROR struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	DEFINITION []ROR_ROR_DEFINITION `hl7:"true,DEFINITION"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ROR_ROR) MessageTypeName() string {
	return "ROR_ROR"
}
// ROR_ROR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ROR_ROR_DEFINITION struct {
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT *ROR_ROR_PATIENT `hl7:"false,PATIENT"`
	ORDER []ROR_ROR_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ROR_ROR_DEFINITION) MessageTypeSubStructName() string {
	return "ROR_ROR_DEFINITION"
}
// ROR_ROR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ROR_ROR_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	RXO *RXO `hl7:"true,RXO"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *ROR_ROR_ORDER) MessageTypeSubStructName() string {
	return "ROR_ROR_ORDER"
}
// ROR_ROR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ROR_ROR_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ROR_ROR_PATIENT) MessageTypeSubStructName() string {
	return "ROR_ROR_PATIENT"
}
// RPA_I08_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_AUTHORIZATION struct {
	AUT *AUT `hl7:"true,AUT"`
	CTD *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPA_I08_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RPA_I08_AUTHORIZATION"
}
// RPA_I08 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	RF1 *RF1 `hl7:"false,RF1"`
	AUTHORIZATION *RPA_I08_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	PROVIDER []RPA_I08_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []RPA_I08_INSURANCE `hl7:"false,INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	PROCEDURE []RPA_I08_PROCEDURE `hl7:"true,PROCEDURE"`
	OBSERVATION []RPA_I08_OBSERVATION `hl7:"false,OBSERVATION"`
	VISIT *RPA_I08_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *RPA_I08) MessageTypeName() string {
	return "RPA_I08"
}
// RPA_I08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RPA_I08_INSURANCE) MessageTypeSubStructName() string {
	return "RPA_I08_INSURANCE"
}
// RPA_I08_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_OBSERVATION struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	RESULTS []RPA_I08_RESULTS `hl7:"false,RESULTS"`
	Other []interface{}
}

func (s *RPA_I08_OBSERVATION) MessageTypeSubStructName() string {
	return "RPA_I08_OBSERVATION"
}
// RPA_I08_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	AUTHORIZATION *RPA_I08_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	Other []interface{}
}

func (s *RPA_I08_PROCEDURE) MessageTypeSubStructName() string {
	return "RPA_I08_PROCEDURE"
}
// RPA_I08_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPA_I08_PROVIDER) MessageTypeSubStructName() string {
	return "RPA_I08_PROVIDER"
}
// RPA_I08_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_RESULTS struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RPA_I08_RESULTS) MessageTypeSubStructName() string {
	return "RPA_I08_RESULTS"
}
// RPA_I08_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPA_I08_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RPA_I08_VISIT) MessageTypeSubStructName() string {
	return "RPA_I08_VISIT"
}
// RPI_I01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPI_I01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	PROVIDER []RPI_I01_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RPI_I01_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RPI_I01) MessageTypeName() string {
	return "RPI_I01"
}
// RPI_I01_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPI_I01_GUARANTOR_INSURANCE struct {
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []RPI_I01_INSURANCE `hl7:"true,INSURANCE"`
	Other []interface{}
}

func (s *RPI_I01_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I01_GUARANTOR_INSURANCE"
}
// RPI_I01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPI_I01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RPI_I01_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I01_INSURANCE"
}
// RPI_I01_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPI_I01_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPI_I01_PROVIDER) MessageTypeSubStructName() string {
	return "RPI_I01_PROVIDER"
}
// RPL_I02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPL_I02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	PROVIDER []RPL_I02_PROVIDER `hl7:"true,PROVIDER"`
	NTE []NTE `hl7:"false,NTE"`
	DSP []DSP `hl7:"false,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RPL_I02) MessageTypeName() string {
	return "RPL_I02"
}
// RPL_I02_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RPL_I02_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPL_I02_PROVIDER) MessageTypeSubStructName() string {
	return "RPL_I02_PROVIDER"
}
// RQA_I08_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_AUTHORIZATION struct {
	AUT *AUT `hl7:"true,AUT"`
	CTD *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQA_I08_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RQA_I08_AUTHORIZATION"
}
// RQA_I08 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08 struct {
	MSH *MSH `hl7:"true,MSH"`
	RF1 *RF1 `hl7:"false,RF1"`
	AUTHORIZATION *RQA_I08_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	PROVIDER []RQA_I08_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RQA_I08_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	ACC *ACC `hl7:"false,ACC"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	PROCEDURE []RQA_I08_PROCEDURE `hl7:"false,PROCEDURE"`
	OBSERVATION []RQA_I08_OBSERVATION `hl7:"false,OBSERVATION"`
	VISIT *RQA_I08_VISIT `hl7:"false,VISIT"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQA_I08) MessageTypeName() string {
	return "RQA_I08"
}
// RQA_I08_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_GUARANTOR_INSURANCE struct {
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []RQA_I08_INSURANCE `hl7:"true,INSURANCE"`
	Other []interface{}
}

func (s *RQA_I08_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RQA_I08_GUARANTOR_INSURANCE"
}
// RQA_I08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RQA_I08_INSURANCE) MessageTypeSubStructName() string {
	return "RQA_I08_INSURANCE"
}
// RQA_I08_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_OBSERVATION struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	RESULTS []RQA_I08_RESULTS `hl7:"false,RESULTS"`
	Other []interface{}
}

func (s *RQA_I08_OBSERVATION) MessageTypeSubStructName() string {
	return "RQA_I08_OBSERVATION"
}
// RQA_I08_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	AUTHORIZATION *RQA_I08_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	Other []interface{}
}

func (s *RQA_I08_PROCEDURE) MessageTypeSubStructName() string {
	return "RQA_I08_PROCEDURE"
}
// RQA_I08_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQA_I08_PROVIDER) MessageTypeSubStructName() string {
	return "RQA_I08_PROVIDER"
}
// RQA_I08_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_RESULTS struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQA_I08_RESULTS) MessageTypeSubStructName() string {
	return "RQA_I08_RESULTS"
}
// RQA_I08_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQA_I08_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RQA_I08_VISIT) MessageTypeSubStructName() string {
	return "RQA_I08_VISIT"
}
// RQC_I05 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQC_I05 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PROVIDER []RQC_I05_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GT1 []GT1 `hl7:"false,GT1"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQC_I05) MessageTypeName() string {
	return "RQC_I05"
}
// RQC_I05_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQC_I05_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQC_I05_PROVIDER) MessageTypeSubStructName() string {
	return "RQC_I05_PROVIDER"
}
// RQC_I06 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQC_I06 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PROVIDER []RQC_I06_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GT1 *GT1 `hl7:"false,GT1"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQC_I06) MessageTypeName() string {
	return "RQC_I06"
}
// RQC_I06_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQC_I06_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQC_I06_PROVIDER) MessageTypeSubStructName() string {
	return "RQC_I06_PROVIDER"
}
// RQI_I01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQI_I01 struct {
	MSH *MSH `hl7:"true,MSH"`
	PROVIDER []RQI_I01_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RQI_I01_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQI_I01) MessageTypeName() string {
	return "RQI_I01"
}
// RQI_I01_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQI_I01_GUARANTOR_INSURANCE struct {
	GT1 []GT1 `hl7:"false,GT1"`
	INSURANCE []RQI_I01_INSURANCE `hl7:"true,INSURANCE"`
	Other []interface{}
}

func (s *RQI_I01_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RQI_I01_GUARANTOR_INSURANCE"
}
// RQI_I01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQI_I01_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RQI_I01_INSURANCE) MessageTypeSubStructName() string {
	return "RQI_I01_INSURANCE"
}
// RQI_I01_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQI_I01_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQI_I01_PROVIDER) MessageTypeSubStructName() string {
	return "RQI_I01_PROVIDER"
}
// RQP_I04 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQP_I04 struct {
	MSH *MSH `hl7:"true,MSH"`
	PROVIDER []RQP_I04_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	GT1 []GT1 `hl7:"false,GT1"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQP_I04) MessageTypeName() string {
	return "RQP_I04"
}
// RQP_I04_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQP_I04_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQP_I04_PROVIDER) MessageTypeSubStructName() string {
	return "RQP_I04_PROVIDER"
}
// RQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQQ_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	ERQ *ERQ `hl7:"true,ERQ"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RQQ_Q01) MessageTypeName() string {
	return "RQQ_Q01"
}
// RRA_O02_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRA_O02_ADMINISTRATION struct {
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"true,RXR"`
	Other []interface{}
}

func (s *RRA_O02_ADMINISTRATION) MessageTypeSubStructName() string {
	return "RRA_O02_ADMINISTRATION"
}
// RRA_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRA_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *RRA_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *RRA_O02) MessageTypeName() string {
	return "RRA_O02"
}
// RRA_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRA_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ADMINISTRATION []RRA_O02_ADMINISTRATION `hl7:"false,ADMINISTRATION"`
	Other []interface{}
}

func (s *RRA_O02_ORDER) MessageTypeSubStructName() string {
	return "RRA_O02_ORDER"
}
// RRA_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRA_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRA_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRA_O02_PATIENT"
}
// RRA_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRA_O02_RESPONSE struct {
	PATIENT *RRA_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER []RRA_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RRA_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRA_O02_RESPONSE"
}
// RRD_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRD_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *RRD_O02_PATIENT `hl7:"false,PATIENT"`
	Other []interface{}
}

func (s *RRD_O02) MessageTypeName() string {
	return "RRD_O02"
}
// RRD_O02_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRD_O02_DISPENSE struct {
	RXD *RXD `hl7:"true,RXD"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRD_O02_DISPENSE) MessageTypeSubStructName() string {
	return "RRD_O02_DISPENSE"
}
// RRD_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRD_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	DISPENSE *RRD_O02_DISPENSE `hl7:"false,DISPENSE"`
	Other []interface{}
}

func (s *RRD_O02_ORDER) MessageTypeSubStructName() string {
	return "RRD_O02_ORDER"
}
// RRD_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRD_O02_PATIENT struct {
	RESPONSE *RRD_O02_RESPONSE `hl7:"false,RESPONSE"`
	ORDER []RRD_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RRD_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRD_O02_PATIENT"
}
// RRD_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRD_O02_RESPONSE struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRD_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRD_O02_RESPONSE"
}
// RRG_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRG_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *RRG_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *RRG_O02) MessageTypeName() string {
	return "RRG_O02"
}
// RRG_O02_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRG_O02_GIVE struct {
	RXG *RXG `hl7:"true,RXG"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRG_O02_GIVE) MessageTypeSubStructName() string {
	return "RRG_O02_GIVE"
}
// RRG_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRG_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	GIVE *RRG_O02_GIVE `hl7:"false,GIVE"`
	Other []interface{}
}

func (s *RRG_O02_ORDER) MessageTypeSubStructName() string {
	return "RRG_O02_ORDER"
}
// RRG_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRG_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRG_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRG_O02_PATIENT"
}
// RRG_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRG_O02_RESPONSE struct {
	PATIENT *RRG_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER []RRG_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RRG_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRG_O02_RESPONSE"
}
// RRI_I12_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_AUTHORIZATION struct {
	AUT *AUT `hl7:"true,AUT"`
	CTD *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RRI_I12_AUTHORIZATION"
}
// RRI_I12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"false,MSA"`
	RF1 *RF1 `hl7:"false,RF1"`
	AUTHORIZATION *RRI_I12_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	PROVIDER []RRI_I12_PROVIDER `hl7:"true,PROVIDER"`
	PID *PID `hl7:"true,PID"`
	ACC *ACC `hl7:"false,ACC"`
	DG1 []DG1 `hl7:"false,DG1"`
	DRG []DRG `hl7:"false,DRG"`
	AL1 []AL1 `hl7:"false,AL1"`
	PROCEDURE []RRI_I12_PROCEDURE `hl7:"false,PROCEDURE"`
	RESULTS []RRI_I12_RESULTS `hl7:"false,RESULTS"`
	VISIT *RRI_I12_VISIT `hl7:"false,VISIT"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRI_I12) MessageTypeName() string {
	return "RRI_I12"
}
// RRI_I12_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRI_I12_OBSERVATION) MessageTypeSubStructName() string {
	return "RRI_I12_OBSERVATION"
}
// RRI_I12_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_PROCEDURE struct {
	PR1 *PR1 `hl7:"true,PR1"`
	AUTHORIZATION *RRI_I12_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	Other []interface{}
}

func (s *RRI_I12_PROCEDURE) MessageTypeSubStructName() string {
	return "RRI_I12_PROCEDURE"
}
// RRI_I12_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_PROVIDER struct {
	PRD *PRD `hl7:"true,PRD"`
	CTD []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_PROVIDER) MessageTypeSubStructName() string {
	return "RRI_I12_PROVIDER"
}
// RRI_I12_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_RESULTS struct {
	OBR *OBR `hl7:"true,OBR"`
	NTE []NTE `hl7:"false,NTE"`
	OBSERVATION []RRI_I12_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *RRI_I12_RESULTS) MessageTypeSubStructName() string {
	return "RRI_I12_RESULTS"
}
// RRI_I12_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RRI_I12_VISIT) MessageTypeSubStructName() string {
	return "RRI_I12_VISIT"
}
// RRO_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRO_O02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	NTE []NTE `hl7:"false,NTE"`
	RESPONSE *RRO_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other []interface{}
}

func (s *RRO_O02) MessageTypeName() string {
	return "RRO_O02"
}
// RRO_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRO_O02_ORDER struct {
	ORC *ORC `hl7:"true,ORC"`
	ORDER_DETAIL *RRO_O02_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other []interface{}
}

func (s *RRO_O02_ORDER) MessageTypeSubStructName() string {
	return "RRO_O02_ORDER"
}
// RRO_O02_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRO_O02_ORDER_DETAIL struct {
	RXO *RXO `hl7:"true,RXO"`
	NTE1 []NTE `hl7:"false,NTE1"`
	RXR []RXR `hl7:"true,RXR"`
	RXC []RXC `hl7:"false,RXC"`
	NTE2 []NTE `hl7:"false,NTE2"`
	Other []interface{}
}

func (s *RRO_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RRO_O02_ORDER_DETAIL"
}
// RRO_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRO_O02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRO_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRO_O02_PATIENT"
}
// RRO_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRO_O02_RESPONSE struct {
	PATIENT *RRO_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER []RRO_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *RRO_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRO_O02_RESPONSE"
}
// SIU_S12 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12 struct {
	MSH *MSH `hl7:"true,MSH"`
	SCH *SCH `hl7:"true,SCH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT []SIU_S12_PATIENT `hl7:"false,PATIENT"`
	RESOURCES []SIU_S12_RESOURCES `hl7:"true,RESOURCES"`
	Other []interface{}
}

func (s *SIU_S12) MessageTypeName() string {
	return "SIU_S12"
}
// SIU_S12_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_GENERAL_RESOURCE struct {
	AIG *AIG `hl7:"true,AIG"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_GENERAL_RESOURCE"
}
// SIU_S12_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_LOCATION_RESOURCE struct {
	AIL *AIL `hl7:"true,AIL"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_LOCATION_RESOURCE"
}
// SIU_S12_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SIU_S12_PATIENT) MessageTypeSubStructName() string {
	return "SIU_S12_PATIENT"
}
// SIU_S12_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_PERSONNEL_RESOURCE struct {
	AIP *AIP `hl7:"true,AIP"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_PERSONNEL_RESOURCE"
}
// SIU_S12_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_RESOURCES struct {
	RGS *RGS `hl7:"true,RGS"`
	SERVICE []SIU_S12_SERVICE `hl7:"false,SERVICE"`
	GENERAL_RESOURCE []SIU_S12_GENERAL_RESOURCE `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE []SIU_S12_LOCATION_RESOURCE `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SIU_S12_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other []interface{}
}

func (s *SIU_S12_RESOURCES) MessageTypeSubStructName() string {
	return "SIU_S12_RESOURCES"
}
// SIU_S12_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SIU_S12_SERVICE struct {
	AIS *AIS `hl7:"true,AIS"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_SERVICE) MessageTypeSubStructName() string {
	return "SIU_S12_SERVICE"
}
// SPQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SPQ_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	SPR *SPR `hl7:"true,SPR"`
	RDF *RDF `hl7:"false,RDF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *SPQ_Q01) MessageTypeName() string {
	return "SPQ_Q01"
}
// SQM_S25 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	REQUEST *SQM_S25_REQUEST `hl7:"false,REQUEST"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *SQM_S25) MessageTypeName() string {
	return "SQM_S25"
}
// SQM_S25_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_GENERAL_RESOURCE struct {
	AIG *AIG `hl7:"true,AIG"`
	APR *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_GENERAL_RESOURCE"
}
// SQM_S25_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_LOCATION_RESOURCE struct {
	AIL *AIL `hl7:"true,AIL"`
	APR *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_LOCATION_RESOURCE"
}
// SQM_S25_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_PERSONNEL_RESOURCE struct {
	AIP *AIP `hl7:"true,AIP"`
	APR *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_PERSONNEL_RESOURCE"
}
// SQM_S25_REQUEST represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_REQUEST struct {
	ARQ *ARQ `hl7:"true,ARQ"`
	APR *APR `hl7:"false,APR"`
	PID *PID `hl7:"false,PID"`
	RESOURCES []SQM_S25_RESOURCES `hl7:"true,RESOURCES"`
	Other []interface{}
}

func (s *SQM_S25_REQUEST) MessageTypeSubStructName() string {
	return "SQM_S25_REQUEST"
}
// SQM_S25_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_RESOURCES struct {
	RGS *RGS `hl7:"true,RGS"`
	SERVICE []SQM_S25_SERVICE `hl7:"false,SERVICE"`
	GENERAL_RESOURCE []SQM_S25_GENERAL_RESOURCE `hl7:"false,GENERAL_RESOURCE"`
	PERSONNEL_RESOURCE []SQM_S25_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	LOCATION_RESOURCE []SQM_S25_LOCATION_RESOURCE `hl7:"false,LOCATION_RESOURCE"`
	Other []interface{}
}

func (s *SQM_S25_RESOURCES) MessageTypeSubStructName() string {
	return "SQM_S25_RESOURCES"
}
// SQM_S25_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQM_S25_SERVICE struct {
	AIS *AIS `hl7:"true,AIS"`
	APR *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_SERVICE) MessageTypeSubStructName() string {
	return "SQM_S25_SERVICE"
}
// SQR_S25 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"true,QAK"`
	SCHEDULE []SQR_S25_SCHEDULE `hl7:"false,SCHEDULE"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *SQR_S25) MessageTypeName() string {
	return "SQR_S25"
}
// SQR_S25_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_GENERAL_RESOURCE struct {
	AIG *AIG `hl7:"true,AIG"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_GENERAL_RESOURCE"
}
// SQR_S25_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_LOCATION_RESOURCE struct {
	AIL *AIL `hl7:"true,AIL"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_LOCATION_RESOURCE"
}
// SQR_S25_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DG1 *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SQR_S25_PATIENT) MessageTypeSubStructName() string {
	return "SQR_S25_PATIENT"
}
// SQR_S25_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_PERSONNEL_RESOURCE struct {
	AIP *AIP `hl7:"true,AIP"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_PERSONNEL_RESOURCE"
}
// SQR_S25_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_RESOURCES struct {
	RGS *RGS `hl7:"true,RGS"`
	SERVICE []SQR_S25_SERVICE `hl7:"false,SERVICE"`
	GENERAL_RESOURCE []SQR_S25_GENERAL_RESOURCE `hl7:"false,GENERAL_RESOURCE"`
	PERSONNEL_RESOURCE []SQR_S25_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	LOCATION_RESOURCE []SQR_S25_LOCATION_RESOURCE `hl7:"false,LOCATION_RESOURCE"`
	Other []interface{}
}

func (s *SQR_S25_RESOURCES) MessageTypeSubStructName() string {
	return "SQR_S25_RESOURCES"
}
// SQR_S25_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_SCHEDULE struct {
	SCH *SCH `hl7:"true,SCH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT *SQR_S25_PATIENT `hl7:"false,PATIENT"`
	RESOURCES []SQR_S25_RESOURCES `hl7:"true,RESOURCES"`
	Other []interface{}
}

func (s *SQR_S25_SCHEDULE) MessageTypeSubStructName() string {
	return "SQR_S25_SCHEDULE"
}
// SQR_S25_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SQR_S25_SERVICE struct {
	AIS *AIS `hl7:"true,AIS"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_SERVICE) MessageTypeSubStructName() string {
	return "SQR_S25_SERVICE"
}
// SRM_S01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01 struct {
	MSH *MSH `hl7:"true,MSH"`
	ARQ *ARQ `hl7:"true,ARQ"`
	APR *APR `hl7:"false,APR"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT []SRM_S01_PATIENT `hl7:"false,PATIENT"`
	RESOURCES []SRM_S01_RESOURCES `hl7:"true,RESOURCES"`
	Other []interface{}
}

func (s *SRM_S01) MessageTypeName() string {
	return "SRM_S01"
}
// SRM_S01_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_GENERAL_RESOURCE struct {
	AIG *AIG `hl7:"true,AIG"`
	APR *APR `hl7:"false,APR"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_GENERAL_RESOURCE"
}
// SRM_S01_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_LOCATION_RESOURCE struct {
	AIL *AIL `hl7:"true,AIL"`
	APR *APR `hl7:"false,APR"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_LOCATION_RESOURCE"
}
// SRM_S01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	OBX []OBX `hl7:"false,OBX"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SRM_S01_PATIENT) MessageTypeSubStructName() string {
	return "SRM_S01_PATIENT"
}
// SRM_S01_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_PERSONNEL_RESOURCE struct {
	AIP *AIP `hl7:"true,AIP"`
	APR *APR `hl7:"false,APR"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_PERSONNEL_RESOURCE"
}
// SRM_S01_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_RESOURCES struct {
	RGS *RGS `hl7:"true,RGS"`
	SERVICE []SRM_S01_SERVICE `hl7:"false,SERVICE"`
	GENERAL_RESOURCE []SRM_S01_GENERAL_RESOURCE `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE []SRM_S01_LOCATION_RESOURCE `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SRM_S01_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other []interface{}
}

func (s *SRM_S01_RESOURCES) MessageTypeSubStructName() string {
	return "SRM_S01_RESOURCES"
}
// SRM_S01_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRM_S01_SERVICE struct {
	AIS *AIS `hl7:"true,AIS"`
	APR *APR `hl7:"false,APR"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_SERVICE) MessageTypeSubStructName() string {
	return "SRM_S01_SERVICE"
}
// SRR_S01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	SCHEDULE *SRR_S01_SCHEDULE `hl7:"false,SCHEDULE"`
	Other []interface{}
}

func (s *SRR_S01) MessageTypeName() string {
	return "SRR_S01"
}
// SRR_S01_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_GENERAL_RESOURCE struct {
	AIG *AIG `hl7:"true,AIG"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_GENERAL_RESOURCE"
}
// SRR_S01_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_LOCATION_RESOURCE struct {
	AIL *AIL `hl7:"true,AIL"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_LOCATION_RESOURCE"
}
// SRR_S01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	PV1 *PV1 `hl7:"false,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	DG1 []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SRR_S01_PATIENT) MessageTypeSubStructName() string {
	return "SRR_S01_PATIENT"
}
// SRR_S01_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_PERSONNEL_RESOURCE struct {
	AIP *AIP `hl7:"true,AIP"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_PERSONNEL_RESOURCE"
}
// SRR_S01_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_RESOURCES struct {
	RGS *RGS `hl7:"true,RGS"`
	SERVICE []SRR_S01_SERVICE `hl7:"false,SERVICE"`
	GENERAL_RESOURCE []SRR_S01_GENERAL_RESOURCE `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE []SRR_S01_LOCATION_RESOURCE `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SRR_S01_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other []interface{}
}

func (s *SRR_S01_RESOURCES) MessageTypeSubStructName() string {
	return "SRR_S01_RESOURCES"
}
// SRR_S01_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_SCHEDULE struct {
	SCH *SCH `hl7:"true,SCH"`
	NTE []NTE `hl7:"false,NTE"`
	PATIENT []SRR_S01_PATIENT `hl7:"false,PATIENT"`
	RESOURCES []SRR_S01_RESOURCES `hl7:"true,RESOURCES"`
	Other []interface{}
}

func (s *SRR_S01_SCHEDULE) MessageTypeSubStructName() string {
	return "SRR_S01_SCHEDULE"
}
// SRR_S01_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SRR_S01_SERVICE struct {
	AIS *AIS `hl7:"true,AIS"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_SERVICE) MessageTypeSubStructName() string {
	return "SRR_S01_SERVICE"
}
// SUR_P09 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SUR_P09 struct {
	MSH *MSH `hl7:"true,MSH"`
	FACILITY []SUR_P09_FACILITY `hl7:"true,FACILITY"`
	Other []interface{}
}

func (s *SUR_P09) MessageTypeName() string {
	return "SUR_P09"
}
// SUR_P09_FACILITY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SUR_P09_FACILITY struct {
	FAC *FAC `hl7:"true,FAC"`
	PRODUCT []SUR_P09_PRODUCT `hl7:"true,PRODUCT"`
	PSH *PSH `hl7:"true,PSH"`
	FACILITY_DETAIL []SUR_P09_FACILITY_DETAIL `hl7:"true,FACILITY_DETAIL"`
	Other []interface{}
}

func (s *SUR_P09_FACILITY) MessageTypeSubStructName() string {
	return "SUR_P09_FACILITY"
}
// SUR_P09_FACILITY_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SUR_P09_FACILITY_DETAIL struct {
	FAC *FAC `hl7:"true,FAC"`
	PDC *PDC `hl7:"true,PDC"`
	NTE *NTE `hl7:"true,NTE"`
	Other []interface{}
}

func (s *SUR_P09_FACILITY_DETAIL) MessageTypeSubStructName() string {
	return "SUR_P09_FACILITY_DETAIL"
}
// SUR_P09_PRODUCT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SUR_P09_PRODUCT struct {
	PSH *PSH `hl7:"true,PSH"`
	PDC *PDC `hl7:"true,PDC"`
	Other []interface{}
}

func (s *SUR_P09_PRODUCT) MessageTypeSubStructName() string {
	return "SUR_P09_PRODUCT"
}
// TBR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type TBR_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"true,QAK"`
	RDF *RDF `hl7:"true,RDF"`
	RDT []RDT `hl7:"true,RDT"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *TBR_Q01) MessageTypeName() string {
	return "TBR_Q01"
}
// UDM_Q05 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type UDM_Q05 struct {
	MSH *MSH `hl7:"true,MSH"`
	URD *URD `hl7:"true,URD"`
	URS *URS `hl7:"false,URS"`
	DSP []DSP `hl7:"true,DSP"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *UDM_Q05) MessageTypeName() string {
	return "UDM_Q05"
}
// VQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VQQ_Q01 struct {
	MSH *MSH `hl7:"true,MSH"`
	VTQ *VTQ `hl7:"true,VTQ"`
	RDF *RDF `hl7:"false,RDF"`
	DSC *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *VQQ_Q01) MessageTypeName() string {
	return "VQQ_Q01"
}
// VXQ_V01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXQ_V01 struct {
	MSH *MSH `hl7:"true,MSH"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *VXQ_V01) MessageTypeName() string {
	return "VXQ_V01"
}
// VXR_V03 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXR_V03 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NK1 []NK1 `hl7:"false,NK1"`
	PATIENT_VISIT *VXR_V03_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE []VXR_V03_INSURANCE `hl7:"false,INSURANCE"`
	ORDER []VXR_V03_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *VXR_V03) MessageTypeName() string {
	return "VXR_V03"
}
// VXR_V03_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXR_V03_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *VXR_V03_INSURANCE) MessageTypeSubStructName() string {
	return "VXR_V03_INSURANCE"
}
// VXR_V03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXR_V03_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *VXR_V03_OBSERVATION) MessageTypeSubStructName() string {
	return "VXR_V03_OBSERVATION"
}
// VXR_V03_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXR_V03_ORDER struct {
	ORC *ORC `hl7:"false,ORC"`
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"false,RXR"`
	OBSERVATION []VXR_V03_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *VXR_V03_ORDER) MessageTypeSubStructName() string {
	return "VXR_V03_ORDER"
}
// VXR_V03_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXR_V03_PATIENT_VISIT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *VXR_V03_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "VXR_V03_PATIENT_VISIT"
}
// VXU_V04 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXU_V04 struct {
	MSH *MSH `hl7:"true,MSH"`
	PID *PID `hl7:"true,PID"`
	PD1 *PD1 `hl7:"false,PD1"`
	NK1 []NK1 `hl7:"false,NK1"`
	PATIENT *VXU_V04_PATIENT `hl7:"false,PATIENT"`
	INSURANCE []VXU_V04_INSURANCE `hl7:"false,INSURANCE"`
	ORDER []VXU_V04_ORDER `hl7:"false,ORDER"`
	Other []interface{}
}

func (s *VXU_V04) MessageTypeName() string {
	return "VXU_V04"
}
// VXU_V04_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXU_V04_INSURANCE struct {
	IN1 *IN1 `hl7:"true,IN1"`
	IN2 *IN2 `hl7:"false,IN2"`
	IN3 *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *VXU_V04_INSURANCE) MessageTypeSubStructName() string {
	return "VXU_V04_INSURANCE"
}
// VXU_V04_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXU_V04_OBSERVATION struct {
	OBX *OBX `hl7:"true,OBX"`
	NTE []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *VXU_V04_OBSERVATION) MessageTypeSubStructName() string {
	return "VXU_V04_OBSERVATION"
}
// VXU_V04_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXU_V04_ORDER struct {
	ORC *ORC `hl7:"false,ORC"`
	RXA *RXA `hl7:"true,RXA"`
	RXR *RXR `hl7:"false,RXR"`
	OBSERVATION []VXU_V04_OBSERVATION `hl7:"false,OBSERVATION"`
	Other []interface{}
}

func (s *VXU_V04_ORDER) MessageTypeSubStructName() string {
	return "VXU_V04_ORDER"
}
// VXU_V04_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXU_V04_PATIENT struct {
	PV1 *PV1 `hl7:"true,PV1"`
	PV2 *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *VXU_V04_PATIENT) MessageTypeSubStructName() string {
	return "VXU_V04_PATIENT"
}
// VXX_V02 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXX_V02 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	QRD *QRD `hl7:"true,QRD"`
	QRF *QRF `hl7:"false,QRF"`
	PATIENT []VXX_V02_PATIENT `hl7:"true,PATIENT"`
	Other []interface{}
}

func (s *VXX_V02) MessageTypeName() string {
	return "VXX_V02"
}
// VXX_V02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VXX_V02_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	NK1 []NK1 `hl7:"false,NK1"`
	Other []interface{}
}

func (s *VXX_V02_PATIENT) MessageTypeSubStructName() string {
	return "VXX_V02_PATIENT"
}

// GenericHL7Segment represents the corresponding HL7 segment type.
type GenericHL7Segment struct {
	segment []byte
}

func (s *GenericHL7Segment) SegmentName() string {
	return "GenericHL7Segment"
}

// ACC returns the first ACC segment within the message, or nil if there isn't one.
func (m *Message) ACC() (*ACC, error) {
	ps, err := m.Parse("ACC")
	pst, ok := ps.(*ACC)
	if ok {
		return pst, err
	}
	return nil, err
}
// ADD returns the first ADD segment within the message, or nil if there isn't one.
func (m *Message) ADD() (*ADD, error) {
	ps, err := m.Parse("ADD")
	pst, ok := ps.(*ADD)
	if ok {
		return pst, err
	}
	return nil, err
}
// AIG returns the first AIG segment within the message, or nil if there isn't one.
func (m *Message) AIG() (*AIG, error) {
	ps, err := m.Parse("AIG")
	pst, ok := ps.(*AIG)
	if ok {
		return pst, err
	}
	return nil, err
}
// AIL returns the first AIL segment within the message, or nil if there isn't one.
func (m *Message) AIL() (*AIL, error) {
	ps, err := m.Parse("AIL")
	pst, ok := ps.(*AIL)
	if ok {
		return pst, err
	}
	return nil, err
}
// AIP returns the first AIP segment within the message, or nil if there isn't one.
func (m *Message) AIP() (*AIP, error) {
	ps, err := m.Parse("AIP")
	pst, ok := ps.(*AIP)
	if ok {
		return pst, err
	}
	return nil, err
}
// AIS returns the first AIS segment within the message, or nil if there isn't one.
func (m *Message) AIS() (*AIS, error) {
	ps, err := m.Parse("AIS")
	pst, ok := ps.(*AIS)
	if ok {
		return pst, err
	}
	return nil, err
}
// AL1 returns the first AL1 segment within the message, or nil if there isn't one.
func (m *Message) AL1() (*AL1, error) {
	ps, err := m.Parse("AL1")
	pst, ok := ps.(*AL1)
	if ok {
		return pst, err
	}
	return nil, err
}
// APR returns the first APR segment within the message, or nil if there isn't one.
func (m *Message) APR() (*APR, error) {
	ps, err := m.Parse("APR")
	pst, ok := ps.(*APR)
	if ok {
		return pst, err
	}
	return nil, err
}
// ARQ returns the first ARQ segment within the message, or nil if there isn't one.
func (m *Message) ARQ() (*ARQ, error) {
	ps, err := m.Parse("ARQ")
	pst, ok := ps.(*ARQ)
	if ok {
		return pst, err
	}
	return nil, err
}
// AUT returns the first AUT segment within the message, or nil if there isn't one.
func (m *Message) AUT() (*AUT, error) {
	ps, err := m.Parse("AUT")
	pst, ok := ps.(*AUT)
	if ok {
		return pst, err
	}
	return nil, err
}
// BHS returns the first BHS segment within the message, or nil if there isn't one.
func (m *Message) BHS() (*BHS, error) {
	ps, err := m.Parse("BHS")
	pst, ok := ps.(*BHS)
	if ok {
		return pst, err
	}
	return nil, err
}
// BLG returns the first BLG segment within the message, or nil if there isn't one.
func (m *Message) BLG() (*BLG, error) {
	ps, err := m.Parse("BLG")
	pst, ok := ps.(*BLG)
	if ok {
		return pst, err
	}
	return nil, err
}
// BTS returns the first BTS segment within the message, or nil if there isn't one.
func (m *Message) BTS() (*BTS, error) {
	ps, err := m.Parse("BTS")
	pst, ok := ps.(*BTS)
	if ok {
		return pst, err
	}
	return nil, err
}
// CDM returns the first CDM segment within the message, or nil if there isn't one.
func (m *Message) CDM() (*CDM, error) {
	ps, err := m.Parse("CDM")
	pst, ok := ps.(*CDM)
	if ok {
		return pst, err
	}
	return nil, err
}
// CM0 returns the first CM0 segment within the message, or nil if there isn't one.
func (m *Message) CM0() (*CM0, error) {
	ps, err := m.Parse("CM0")
	pst, ok := ps.(*CM0)
	if ok {
		return pst, err
	}
	return nil, err
}
// CM1 returns the first CM1 segment within the message, or nil if there isn't one.
func (m *Message) CM1() (*CM1, error) {
	ps, err := m.Parse("CM1")
	pst, ok := ps.(*CM1)
	if ok {
		return pst, err
	}
	return nil, err
}
// CM2 returns the first CM2 segment within the message, or nil if there isn't one.
func (m *Message) CM2() (*CM2, error) {
	ps, err := m.Parse("CM2")
	pst, ok := ps.(*CM2)
	if ok {
		return pst, err
	}
	return nil, err
}
// CSP returns the first CSP segment within the message, or nil if there isn't one.
func (m *Message) CSP() (*CSP, error) {
	ps, err := m.Parse("CSP")
	pst, ok := ps.(*CSP)
	if ok {
		return pst, err
	}
	return nil, err
}
// CSR returns the first CSR segment within the message, or nil if there isn't one.
func (m *Message) CSR() (*CSR, error) {
	ps, err := m.Parse("CSR")
	pst, ok := ps.(*CSR)
	if ok {
		return pst, err
	}
	return nil, err
}
// CSS returns the first CSS segment within the message, or nil if there isn't one.
func (m *Message) CSS() (*CSS, error) {
	ps, err := m.Parse("CSS")
	pst, ok := ps.(*CSS)
	if ok {
		return pst, err
	}
	return nil, err
}
// CTD returns the first CTD segment within the message, or nil if there isn't one.
func (m *Message) CTD() (*CTD, error) {
	ps, err := m.Parse("CTD")
	pst, ok := ps.(*CTD)
	if ok {
		return pst, err
	}
	return nil, err
}
// CTI returns the first CTI segment within the message, or nil if there isn't one.
func (m *Message) CTI() (*CTI, error) {
	ps, err := m.Parse("CTI")
	pst, ok := ps.(*CTI)
	if ok {
		return pst, err
	}
	return nil, err
}
// DB1 returns the first DB1 segment within the message, or nil if there isn't one.
func (m *Message) DB1() (*DB1, error) {
	ps, err := m.Parse("DB1")
	pst, ok := ps.(*DB1)
	if ok {
		return pst, err
	}
	return nil, err
}
// DG1 returns the first DG1 segment within the message, or nil if there isn't one.
func (m *Message) DG1() (*DG1, error) {
	ps, err := m.Parse("DG1")
	pst, ok := ps.(*DG1)
	if ok {
		return pst, err
	}
	return nil, err
}
// DRG returns the first DRG segment within the message, or nil if there isn't one.
func (m *Message) DRG() (*DRG, error) {
	ps, err := m.Parse("DRG")
	pst, ok := ps.(*DRG)
	if ok {
		return pst, err
	}
	return nil, err
}
// DSC returns the first DSC segment within the message, or nil if there isn't one.
func (m *Message) DSC() (*DSC, error) {
	ps, err := m.Parse("DSC")
	pst, ok := ps.(*DSC)
	if ok {
		return pst, err
	}
	return nil, err
}
// DSP returns the first DSP segment within the message, or nil if there isn't one.
func (m *Message) DSP() (*DSP, error) {
	ps, err := m.Parse("DSP")
	pst, ok := ps.(*DSP)
	if ok {
		return pst, err
	}
	return nil, err
}
// EQL returns the first EQL segment within the message, or nil if there isn't one.
func (m *Message) EQL() (*EQL, error) {
	ps, err := m.Parse("EQL")
	pst, ok := ps.(*EQL)
	if ok {
		return pst, err
	}
	return nil, err
}
// ERQ returns the first ERQ segment within the message, or nil if there isn't one.
func (m *Message) ERQ() (*ERQ, error) {
	ps, err := m.Parse("ERQ")
	pst, ok := ps.(*ERQ)
	if ok {
		return pst, err
	}
	return nil, err
}
// ERR returns the first ERR segment within the message, or nil if there isn't one.
func (m *Message) ERR() (*ERR, error) {
	ps, err := m.Parse("ERR")
	pst, ok := ps.(*ERR)
	if ok {
		return pst, err
	}
	return nil, err
}
// EVN returns the first EVN segment within the message, or nil if there isn't one.
func (m *Message) EVN() (*EVN, error) {
	ps, err := m.Parse("EVN")
	pst, ok := ps.(*EVN)
	if ok {
		return pst, err
	}
	return nil, err
}
// FAC returns the first FAC segment within the message, or nil if there isn't one.
func (m *Message) FAC() (*FAC, error) {
	ps, err := m.Parse("FAC")
	pst, ok := ps.(*FAC)
	if ok {
		return pst, err
	}
	return nil, err
}
// FHS returns the first FHS segment within the message, or nil if there isn't one.
func (m *Message) FHS() (*FHS, error) {
	ps, err := m.Parse("FHS")
	pst, ok := ps.(*FHS)
	if ok {
		return pst, err
	}
	return nil, err
}
// FT1 returns the first FT1 segment within the message, or nil if there isn't one.
func (m *Message) FT1() (*FT1, error) {
	ps, err := m.Parse("FT1")
	pst, ok := ps.(*FT1)
	if ok {
		return pst, err
	}
	return nil, err
}
// FTS returns the first FTS segment within the message, or nil if there isn't one.
func (m *Message) FTS() (*FTS, error) {
	ps, err := m.Parse("FTS")
	pst, ok := ps.(*FTS)
	if ok {
		return pst, err
	}
	return nil, err
}
// GOL returns the first GOL segment within the message, or nil if there isn't one.
func (m *Message) GOL() (*GOL, error) {
	ps, err := m.Parse("GOL")
	pst, ok := ps.(*GOL)
	if ok {
		return pst, err
	}
	return nil, err
}
// GT1 returns the first GT1 segment within the message, or nil if there isn't one.
func (m *Message) GT1() (*GT1, error) {
	ps, err := m.Parse("GT1")
	pst, ok := ps.(*GT1)
	if ok {
		return pst, err
	}
	return nil, err
}
// IN1 returns the first IN1 segment within the message, or nil if there isn't one.
func (m *Message) IN1() (*IN1, error) {
	ps, err := m.Parse("IN1")
	pst, ok := ps.(*IN1)
	if ok {
		return pst, err
	}
	return nil, err
}
// IN2 returns the first IN2 segment within the message, or nil if there isn't one.
func (m *Message) IN2() (*IN2, error) {
	ps, err := m.Parse("IN2")
	pst, ok := ps.(*IN2)
	if ok {
		return pst, err
	}
	return nil, err
}
// IN3 returns the first IN3 segment within the message, or nil if there isn't one.
func (m *Message) IN3() (*IN3, error) {
	ps, err := m.Parse("IN3")
	pst, ok := ps.(*IN3)
	if ok {
		return pst, err
	}
	return nil, err
}
// LCC returns the first LCC segment within the message, or nil if there isn't one.
func (m *Message) LCC() (*LCC, error) {
	ps, err := m.Parse("LCC")
	pst, ok := ps.(*LCC)
	if ok {
		return pst, err
	}
	return nil, err
}
// LCH returns the first LCH segment within the message, or nil if there isn't one.
func (m *Message) LCH() (*LCH, error) {
	ps, err := m.Parse("LCH")
	pst, ok := ps.(*LCH)
	if ok {
		return pst, err
	}
	return nil, err
}
// LDP returns the first LDP segment within the message, or nil if there isn't one.
func (m *Message) LDP() (*LDP, error) {
	ps, err := m.Parse("LDP")
	pst, ok := ps.(*LDP)
	if ok {
		return pst, err
	}
	return nil, err
}
// LOC returns the first LOC segment within the message, or nil if there isn't one.
func (m *Message) LOC() (*LOC, error) {
	ps, err := m.Parse("LOC")
	pst, ok := ps.(*LOC)
	if ok {
		return pst, err
	}
	return nil, err
}
// LRL returns the first LRL segment within the message, or nil if there isn't one.
func (m *Message) LRL() (*LRL, error) {
	ps, err := m.Parse("LRL")
	pst, ok := ps.(*LRL)
	if ok {
		return pst, err
	}
	return nil, err
}
// MFA returns the first MFA segment within the message, or nil if there isn't one.
func (m *Message) MFA() (*MFA, error) {
	ps, err := m.Parse("MFA")
	pst, ok := ps.(*MFA)
	if ok {
		return pst, err
	}
	return nil, err
}
// MFE returns the first MFE segment within the message, or nil if there isn't one.
func (m *Message) MFE() (*MFE, error) {
	ps, err := m.Parse("MFE")
	pst, ok := ps.(*MFE)
	if ok {
		return pst, err
	}
	return nil, err
}
// MFI returns the first MFI segment within the message, or nil if there isn't one.
func (m *Message) MFI() (*MFI, error) {
	ps, err := m.Parse("MFI")
	pst, ok := ps.(*MFI)
	if ok {
		return pst, err
	}
	return nil, err
}
// MRG returns the first MRG segment within the message, or nil if there isn't one.
func (m *Message) MRG() (*MRG, error) {
	ps, err := m.Parse("MRG")
	pst, ok := ps.(*MRG)
	if ok {
		return pst, err
	}
	return nil, err
}
// MSA returns the first MSA segment within the message, or nil if there isn't one.
func (m *Message) MSA() (*MSA, error) {
	ps, err := m.Parse("MSA")
	pst, ok := ps.(*MSA)
	if ok {
		return pst, err
	}
	return nil, err
}
// MSH returns the first MSH segment within the message, or nil if there isn't one.
func (m *Message) MSH() (*MSH, error) {
	ps, err := m.Parse("MSH")
	pst, ok := ps.(*MSH)
	if ok {
		return pst, err
	}
	return nil, err
}
// NCK returns the first NCK segment within the message, or nil if there isn't one.
func (m *Message) NCK() (*NCK, error) {
	ps, err := m.Parse("NCK")
	pst, ok := ps.(*NCK)
	if ok {
		return pst, err
	}
	return nil, err
}
// NK1 returns the first NK1 segment within the message, or nil if there isn't one.
func (m *Message) NK1() (*NK1, error) {
	ps, err := m.Parse("NK1")
	pst, ok := ps.(*NK1)
	if ok {
		return pst, err
	}
	return nil, err
}
// NPU returns the first NPU segment within the message, or nil if there isn't one.
func (m *Message) NPU() (*NPU, error) {
	ps, err := m.Parse("NPU")
	pst, ok := ps.(*NPU)
	if ok {
		return pst, err
	}
	return nil, err
}
// NSC returns the first NSC segment within the message, or nil if there isn't one.
func (m *Message) NSC() (*NSC, error) {
	ps, err := m.Parse("NSC")
	pst, ok := ps.(*NSC)
	if ok {
		return pst, err
	}
	return nil, err
}
// NST returns the first NST segment within the message, or nil if there isn't one.
func (m *Message) NST() (*NST, error) {
	ps, err := m.Parse("NST")
	pst, ok := ps.(*NST)
	if ok {
		return pst, err
	}
	return nil, err
}
// NTE returns the first NTE segment within the message, or nil if there isn't one.
func (m *Message) NTE() (*NTE, error) {
	ps, err := m.Parse("NTE")
	pst, ok := ps.(*NTE)
	if ok {
		return pst, err
	}
	return nil, err
}
// OBR returns the first OBR segment within the message, or nil if there isn't one.
func (m *Message) OBR() (*OBR, error) {
	ps, err := m.Parse("OBR")
	pst, ok := ps.(*OBR)
	if ok {
		return pst, err
	}
	return nil, err
}
// OBX returns the first OBX segment within the message, or nil if there isn't one.
func (m *Message) OBX() (*OBX, error) {
	ps, err := m.Parse("OBX")
	pst, ok := ps.(*OBX)
	if ok {
		return pst, err
	}
	return nil, err
}
// ODS returns the first ODS segment within the message, or nil if there isn't one.
func (m *Message) ODS() (*ODS, error) {
	ps, err := m.Parse("ODS")
	pst, ok := ps.(*ODS)
	if ok {
		return pst, err
	}
	return nil, err
}
// ODT returns the first ODT segment within the message, or nil if there isn't one.
func (m *Message) ODT() (*ODT, error) {
	ps, err := m.Parse("ODT")
	pst, ok := ps.(*ODT)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM1 returns the first OM1 segment within the message, or nil if there isn't one.
func (m *Message) OM1() (*OM1, error) {
	ps, err := m.Parse("OM1")
	pst, ok := ps.(*OM1)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM2 returns the first OM2 segment within the message, or nil if there isn't one.
func (m *Message) OM2() (*OM2, error) {
	ps, err := m.Parse("OM2")
	pst, ok := ps.(*OM2)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM3 returns the first OM3 segment within the message, or nil if there isn't one.
func (m *Message) OM3() (*OM3, error) {
	ps, err := m.Parse("OM3")
	pst, ok := ps.(*OM3)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM4 returns the first OM4 segment within the message, or nil if there isn't one.
func (m *Message) OM4() (*OM4, error) {
	ps, err := m.Parse("OM4")
	pst, ok := ps.(*OM4)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM5 returns the first OM5 segment within the message, or nil if there isn't one.
func (m *Message) OM5() (*OM5, error) {
	ps, err := m.Parse("OM5")
	pst, ok := ps.(*OM5)
	if ok {
		return pst, err
	}
	return nil, err
}
// OM6 returns the first OM6 segment within the message, or nil if there isn't one.
func (m *Message) OM6() (*OM6, error) {
	ps, err := m.Parse("OM6")
	pst, ok := ps.(*OM6)
	if ok {
		return pst, err
	}
	return nil, err
}
// ORC returns the first ORC segment within the message, or nil if there isn't one.
func (m *Message) ORC() (*ORC, error) {
	ps, err := m.Parse("ORC")
	pst, ok := ps.(*ORC)
	if ok {
		return pst, err
	}
	return nil, err
}
// ORO returns the first ORO segment within the message, or nil if there isn't one.
func (m *Message) ORO() (*ORO, error) {
	ps, err := m.Parse("ORO")
	pst, ok := ps.(*ORO)
	if ok {
		return pst, err
	}
	return nil, err
}
// PCR returns the first PCR segment within the message, or nil if there isn't one.
func (m *Message) PCR() (*PCR, error) {
	ps, err := m.Parse("PCR")
	pst, ok := ps.(*PCR)
	if ok {
		return pst, err
	}
	return nil, err
}
// PD1 returns the first PD1 segment within the message, or nil if there isn't one.
func (m *Message) PD1() (*PD1, error) {
	ps, err := m.Parse("PD1")
	pst, ok := ps.(*PD1)
	if ok {
		return pst, err
	}
	return nil, err
}
// PDC returns the first PDC segment within the message, or nil if there isn't one.
func (m *Message) PDC() (*PDC, error) {
	ps, err := m.Parse("PDC")
	pst, ok := ps.(*PDC)
	if ok {
		return pst, err
	}
	return nil, err
}
// PEO returns the first PEO segment within the message, or nil if there isn't one.
func (m *Message) PEO() (*PEO, error) {
	ps, err := m.Parse("PEO")
	pst, ok := ps.(*PEO)
	if ok {
		return pst, err
	}
	return nil, err
}
// PES returns the first PES segment within the message, or nil if there isn't one.
func (m *Message) PES() (*PES, error) {
	ps, err := m.Parse("PES")
	pst, ok := ps.(*PES)
	if ok {
		return pst, err
	}
	return nil, err
}
// PID returns the first PID segment within the message, or nil if there isn't one.
func (m *Message) PID() (*PID, error) {
	ps, err := m.Parse("PID")
	pst, ok := ps.(*PID)
	if ok {
		return pst, err
	}
	return nil, err
}
// PR1 returns the first PR1 segment within the message, or nil if there isn't one.
func (m *Message) PR1() (*PR1, error) {
	ps, err := m.Parse("PR1")
	pst, ok := ps.(*PR1)
	if ok {
		return pst, err
	}
	return nil, err
}
// PRA returns the first PRA segment within the message, or nil if there isn't one.
func (m *Message) PRA() (*PRA, error) {
	ps, err := m.Parse("PRA")
	pst, ok := ps.(*PRA)
	if ok {
		return pst, err
	}
	return nil, err
}
// PRB returns the first PRB segment within the message, or nil if there isn't one.
func (m *Message) PRB() (*PRB, error) {
	ps, err := m.Parse("PRB")
	pst, ok := ps.(*PRB)
	if ok {
		return pst, err
	}
	return nil, err
}
// PRC returns the first PRC segment within the message, or nil if there isn't one.
func (m *Message) PRC() (*PRC, error) {
	ps, err := m.Parse("PRC")
	pst, ok := ps.(*PRC)
	if ok {
		return pst, err
	}
	return nil, err
}
// PRD returns the first PRD segment within the message, or nil if there isn't one.
func (m *Message) PRD() (*PRD, error) {
	ps, err := m.Parse("PRD")
	pst, ok := ps.(*PRD)
	if ok {
		return pst, err
	}
	return nil, err
}
// PSH returns the first PSH segment within the message, or nil if there isn't one.
func (m *Message) PSH() (*PSH, error) {
	ps, err := m.Parse("PSH")
	pst, ok := ps.(*PSH)
	if ok {
		return pst, err
	}
	return nil, err
}
// PTH returns the first PTH segment within the message, or nil if there isn't one.
func (m *Message) PTH() (*PTH, error) {
	ps, err := m.Parse("PTH")
	pst, ok := ps.(*PTH)
	if ok {
		return pst, err
	}
	return nil, err
}
// PV1 returns the first PV1 segment within the message, or nil if there isn't one.
func (m *Message) PV1() (*PV1, error) {
	ps, err := m.Parse("PV1")
	pst, ok := ps.(*PV1)
	if ok {
		return pst, err
	}
	return nil, err
}
// PV2 returns the first PV2 segment within the message, or nil if there isn't one.
func (m *Message) PV2() (*PV2, error) {
	ps, err := m.Parse("PV2")
	pst, ok := ps.(*PV2)
	if ok {
		return pst, err
	}
	return nil, err
}
// QAK returns the first QAK segment within the message, or nil if there isn't one.
func (m *Message) QAK() (*QAK, error) {
	ps, err := m.Parse("QAK")
	pst, ok := ps.(*QAK)
	if ok {
		return pst, err
	}
	return nil, err
}
// QRD returns the first QRD segment within the message, or nil if there isn't one.
func (m *Message) QRD() (*QRD, error) {
	ps, err := m.Parse("QRD")
	pst, ok := ps.(*QRD)
	if ok {
		return pst, err
	}
	return nil, err
}
// QRF returns the first QRF segment within the message, or nil if there isn't one.
func (m *Message) QRF() (*QRF, error) {
	ps, err := m.Parse("QRF")
	pst, ok := ps.(*QRF)
	if ok {
		return pst, err
	}
	return nil, err
}
// RDF returns the first RDF segment within the message, or nil if there isn't one.
func (m *Message) RDF() (*RDF, error) {
	ps, err := m.Parse("RDF")
	pst, ok := ps.(*RDF)
	if ok {
		return pst, err
	}
	return nil, err
}
// RDT returns the first RDT segment within the message, or nil if there isn't one.
func (m *Message) RDT() (*RDT, error) {
	ps, err := m.Parse("RDT")
	pst, ok := ps.(*RDT)
	if ok {
		return pst, err
	}
	return nil, err
}
// RF1 returns the first RF1 segment within the message, or nil if there isn't one.
func (m *Message) RF1() (*RF1, error) {
	ps, err := m.Parse("RF1")
	pst, ok := ps.(*RF1)
	if ok {
		return pst, err
	}
	return nil, err
}
// RGS returns the first RGS segment within the message, or nil if there isn't one.
func (m *Message) RGS() (*RGS, error) {
	ps, err := m.Parse("RGS")
	pst, ok := ps.(*RGS)
	if ok {
		return pst, err
	}
	return nil, err
}
// ROL returns the first ROL segment within the message, or nil if there isn't one.
func (m *Message) ROL() (*ROL, error) {
	ps, err := m.Parse("ROL")
	pst, ok := ps.(*ROL)
	if ok {
		return pst, err
	}
	return nil, err
}
// RQ1 returns the first RQ1 segment within the message, or nil if there isn't one.
func (m *Message) RQ1() (*RQ1, error) {
	ps, err := m.Parse("RQ1")
	pst, ok := ps.(*RQ1)
	if ok {
		return pst, err
	}
	return nil, err
}
// RQD returns the first RQD segment within the message, or nil if there isn't one.
func (m *Message) RQD() (*RQD, error) {
	ps, err := m.Parse("RQD")
	pst, ok := ps.(*RQD)
	if ok {
		return pst, err
	}
	return nil, err
}
// RX1 returns the first RX1 segment within the message, or nil if there isn't one.
func (m *Message) RX1() (*RX1, error) {
	ps, err := m.Parse("RX1")
	pst, ok := ps.(*RX1)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXA returns the first RXA segment within the message, or nil if there isn't one.
func (m *Message) RXA() (*RXA, error) {
	ps, err := m.Parse("RXA")
	pst, ok := ps.(*RXA)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXC returns the first RXC segment within the message, or nil if there isn't one.
func (m *Message) RXC() (*RXC, error) {
	ps, err := m.Parse("RXC")
	pst, ok := ps.(*RXC)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXD returns the first RXD segment within the message, or nil if there isn't one.
func (m *Message) RXD() (*RXD, error) {
	ps, err := m.Parse("RXD")
	pst, ok := ps.(*RXD)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXE returns the first RXE segment within the message, or nil if there isn't one.
func (m *Message) RXE() (*RXE, error) {
	ps, err := m.Parse("RXE")
	pst, ok := ps.(*RXE)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXG returns the first RXG segment within the message, or nil if there isn't one.
func (m *Message) RXG() (*RXG, error) {
	ps, err := m.Parse("RXG")
	pst, ok := ps.(*RXG)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXO returns the first RXO segment within the message, or nil if there isn't one.
func (m *Message) RXO() (*RXO, error) {
	ps, err := m.Parse("RXO")
	pst, ok := ps.(*RXO)
	if ok {
		return pst, err
	}
	return nil, err
}
// RXR returns the first RXR segment within the message, or nil if there isn't one.
func (m *Message) RXR() (*RXR, error) {
	ps, err := m.Parse("RXR")
	pst, ok := ps.(*RXR)
	if ok {
		return pst, err
	}
	return nil, err
}
// SCH returns the first SCH segment within the message, or nil if there isn't one.
func (m *Message) SCH() (*SCH, error) {
	ps, err := m.Parse("SCH")
	pst, ok := ps.(*SCH)
	if ok {
		return pst, err
	}
	return nil, err
}
// SPR returns the first SPR segment within the message, or nil if there isn't one.
func (m *Message) SPR() (*SPR, error) {
	ps, err := m.Parse("SPR")
	pst, ok := ps.(*SPR)
	if ok {
		return pst, err
	}
	return nil, err
}
// STF returns the first STF segment within the message, or nil if there isn't one.
func (m *Message) STF() (*STF, error) {
	ps, err := m.Parse("STF")
	pst, ok := ps.(*STF)
	if ok {
		return pst, err
	}
	return nil, err
}
// TXA returns the first TXA segment within the message, or nil if there isn't one.
func (m *Message) TXA() (*TXA, error) {
	ps, err := m.Parse("TXA")
	pst, ok := ps.(*TXA)
	if ok {
		return pst, err
	}
	return nil, err
}
// UB1 returns the first UB1 segment within the message, or nil if there isn't one.
func (m *Message) UB1() (*UB1, error) {
	ps, err := m.Parse("UB1")
	pst, ok := ps.(*UB1)
	if ok {
		return pst, err
	}
	return nil, err
}
// UB2 returns the first UB2 segment within the message, or nil if there isn't one.
func (m *Message) UB2() (*UB2, error) {
	ps, err := m.Parse("UB2")
	pst, ok := ps.(*UB2)
	if ok {
		return pst, err
	}
	return nil, err
}
// URD returns the first URD segment within the message, or nil if there isn't one.
func (m *Message) URD() (*URD, error) {
	ps, err := m.Parse("URD")
	pst, ok := ps.(*URD)
	if ok {
		return pst, err
	}
	return nil, err
}
// URS returns the first URS segment within the message, or nil if there isn't one.
func (m *Message) URS() (*URS, error) {
	ps, err := m.Parse("URS")
	pst, ok := ps.(*URS)
	if ok {
		return pst, err
	}
	return nil, err
}
// VAR returns the first VAR segment within the message, or nil if there isn't one.
func (m *Message) VAR() (*VAR, error) {
	ps, err := m.Parse("VAR")
	pst, ok := ps.(*VAR)
	if ok {
		return pst, err
	}
	return nil, err
}
// VTQ returns the first VTQ segment within the message, or nil if there isn't one.
func (m *Message) VTQ() (*VTQ, error) {
	ps, err := m.Parse("VTQ")
	pst, ok := ps.(*VTQ)
	if ok {
		return pst, err
	}
	return nil, err
}

// AllACC returns a slice containing all ACC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllACC() ([]*ACC, error) {
	pss, err := m.ParseAll("ACC")
	return pss.([]*ACC), err
}
// AllADD returns a slice containing all ADD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllADD() ([]*ADD, error) {
	pss, err := m.ParseAll("ADD")
	return pss.([]*ADD), err
}
// AllAIG returns a slice containing all AIG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIG() ([]*AIG, error) {
	pss, err := m.ParseAll("AIG")
	return pss.([]*AIG), err
}
// AllAIL returns a slice containing all AIL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIL() ([]*AIL, error) {
	pss, err := m.ParseAll("AIL")
	return pss.([]*AIL), err
}
// AllAIP returns a slice containing all AIP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIP() ([]*AIP, error) {
	pss, err := m.ParseAll("AIP")
	return pss.([]*AIP), err
}
// AllAIS returns a slice containing all AIS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIS() ([]*AIS, error) {
	pss, err := m.ParseAll("AIS")
	return pss.([]*AIS), err
}
// AllAL1 returns a slice containing all AL1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAL1() ([]*AL1, error) {
	pss, err := m.ParseAll("AL1")
	return pss.([]*AL1), err
}
// AllAPR returns a slice containing all APR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAPR() ([]*APR, error) {
	pss, err := m.ParseAll("APR")
	return pss.([]*APR), err
}
// AllARQ returns a slice containing all ARQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllARQ() ([]*ARQ, error) {
	pss, err := m.ParseAll("ARQ")
	return pss.([]*ARQ), err
}
// AllAUT returns a slice containing all AUT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAUT() ([]*AUT, error) {
	pss, err := m.ParseAll("AUT")
	return pss.([]*AUT), err
}
// AllBHS returns a slice containing all BHS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBHS() ([]*BHS, error) {
	pss, err := m.ParseAll("BHS")
	return pss.([]*BHS), err
}
// AllBLG returns a slice containing all BLG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBLG() ([]*BLG, error) {
	pss, err := m.ParseAll("BLG")
	return pss.([]*BLG), err
}
// AllBTS returns a slice containing all BTS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBTS() ([]*BTS, error) {
	pss, err := m.ParseAll("BTS")
	return pss.([]*BTS), err
}
// AllCDM returns a slice containing all CDM segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCDM() ([]*CDM, error) {
	pss, err := m.ParseAll("CDM")
	return pss.([]*CDM), err
}
// AllCM0 returns a slice containing all CM0 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM0() ([]*CM0, error) {
	pss, err := m.ParseAll("CM0")
	return pss.([]*CM0), err
}
// AllCM1 returns a slice containing all CM1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM1() ([]*CM1, error) {
	pss, err := m.ParseAll("CM1")
	return pss.([]*CM1), err
}
// AllCM2 returns a slice containing all CM2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM2() ([]*CM2, error) {
	pss, err := m.ParseAll("CM2")
	return pss.([]*CM2), err
}
// AllCSP returns a slice containing all CSP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSP() ([]*CSP, error) {
	pss, err := m.ParseAll("CSP")
	return pss.([]*CSP), err
}
// AllCSR returns a slice containing all CSR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSR() ([]*CSR, error) {
	pss, err := m.ParseAll("CSR")
	return pss.([]*CSR), err
}
// AllCSS returns a slice containing all CSS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSS() ([]*CSS, error) {
	pss, err := m.ParseAll("CSS")
	return pss.([]*CSS), err
}
// AllCTD returns a slice containing all CTD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCTD() ([]*CTD, error) {
	pss, err := m.ParseAll("CTD")
	return pss.([]*CTD), err
}
// AllCTI returns a slice containing all CTI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCTI() ([]*CTI, error) {
	pss, err := m.ParseAll("CTI")
	return pss.([]*CTI), err
}
// AllDB1 returns a slice containing all DB1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDB1() ([]*DB1, error) {
	pss, err := m.ParseAll("DB1")
	return pss.([]*DB1), err
}
// AllDG1 returns a slice containing all DG1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDG1() ([]*DG1, error) {
	pss, err := m.ParseAll("DG1")
	return pss.([]*DG1), err
}
// AllDRG returns a slice containing all DRG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDRG() ([]*DRG, error) {
	pss, err := m.ParseAll("DRG")
	return pss.([]*DRG), err
}
// AllDSC returns a slice containing all DSC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDSC() ([]*DSC, error) {
	pss, err := m.ParseAll("DSC")
	return pss.([]*DSC), err
}
// AllDSP returns a slice containing all DSP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDSP() ([]*DSP, error) {
	pss, err := m.ParseAll("DSP")
	return pss.([]*DSP), err
}
// AllEQL returns a slice containing all EQL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEQL() ([]*EQL, error) {
	pss, err := m.ParseAll("EQL")
	return pss.([]*EQL), err
}
// AllERQ returns a slice containing all ERQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllERQ() ([]*ERQ, error) {
	pss, err := m.ParseAll("ERQ")
	return pss.([]*ERQ), err
}
// AllERR returns a slice containing all ERR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllERR() ([]*ERR, error) {
	pss, err := m.ParseAll("ERR")
	return pss.([]*ERR), err
}
// AllEVN returns a slice containing all EVN segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEVN() ([]*EVN, error) {
	pss, err := m.ParseAll("EVN")
	return pss.([]*EVN), err
}
// AllFAC returns a slice containing all FAC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFAC() ([]*FAC, error) {
	pss, err := m.ParseAll("FAC")
	return pss.([]*FAC), err
}
// AllFHS returns a slice containing all FHS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFHS() ([]*FHS, error) {
	pss, err := m.ParseAll("FHS")
	return pss.([]*FHS), err
}
// AllFT1 returns a slice containing all FT1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFT1() ([]*FT1, error) {
	pss, err := m.ParseAll("FT1")
	return pss.([]*FT1), err
}
// AllFTS returns a slice containing all FTS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFTS() ([]*FTS, error) {
	pss, err := m.ParseAll("FTS")
	return pss.([]*FTS), err
}
// AllGOL returns a slice containing all GOL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGOL() ([]*GOL, error) {
	pss, err := m.ParseAll("GOL")
	return pss.([]*GOL), err
}
// AllGT1 returns a slice containing all GT1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGT1() ([]*GT1, error) {
	pss, err := m.ParseAll("GT1")
	return pss.([]*GT1), err
}
// AllIN1 returns a slice containing all IN1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN1() ([]*IN1, error) {
	pss, err := m.ParseAll("IN1")
	return pss.([]*IN1), err
}
// AllIN2 returns a slice containing all IN2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN2() ([]*IN2, error) {
	pss, err := m.ParseAll("IN2")
	return pss.([]*IN2), err
}
// AllIN3 returns a slice containing all IN3 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN3() ([]*IN3, error) {
	pss, err := m.ParseAll("IN3")
	return pss.([]*IN3), err
}
// AllLCC returns a slice containing all LCC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLCC() ([]*LCC, error) {
	pss, err := m.ParseAll("LCC")
	return pss.([]*LCC), err
}
// AllLCH returns a slice containing all LCH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLCH() ([]*LCH, error) {
	pss, err := m.ParseAll("LCH")
	return pss.([]*LCH), err
}
// AllLDP returns a slice containing all LDP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLDP() ([]*LDP, error) {
	pss, err := m.ParseAll("LDP")
	return pss.([]*LDP), err
}
// AllLOC returns a slice containing all LOC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLOC() ([]*LOC, error) {
	pss, err := m.ParseAll("LOC")
	return pss.([]*LOC), err
}
// AllLRL returns a slice containing all LRL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLRL() ([]*LRL, error) {
	pss, err := m.ParseAll("LRL")
	return pss.([]*LRL), err
}
// AllMFA returns a slice containing all MFA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFA() ([]*MFA, error) {
	pss, err := m.ParseAll("MFA")
	return pss.([]*MFA), err
}
// AllMFE returns a slice containing all MFE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFE() ([]*MFE, error) {
	pss, err := m.ParseAll("MFE")
	return pss.([]*MFE), err
}
// AllMFI returns a slice containing all MFI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFI() ([]*MFI, error) {
	pss, err := m.ParseAll("MFI")
	return pss.([]*MFI), err
}
// AllMRG returns a slice containing all MRG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMRG() ([]*MRG, error) {
	pss, err := m.ParseAll("MRG")
	return pss.([]*MRG), err
}
// AllMSA returns a slice containing all MSA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMSA() ([]*MSA, error) {
	pss, err := m.ParseAll("MSA")
	return pss.([]*MSA), err
}
// AllMSH returns a slice containing all MSH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMSH() ([]*MSH, error) {
	pss, err := m.ParseAll("MSH")
	return pss.([]*MSH), err
}
// AllNCK returns a slice containing all NCK segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNCK() ([]*NCK, error) {
	pss, err := m.ParseAll("NCK")
	return pss.([]*NCK), err
}
// AllNK1 returns a slice containing all NK1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNK1() ([]*NK1, error) {
	pss, err := m.ParseAll("NK1")
	return pss.([]*NK1), err
}
// AllNPU returns a slice containing all NPU segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNPU() ([]*NPU, error) {
	pss, err := m.ParseAll("NPU")
	return pss.([]*NPU), err
}
// AllNSC returns a slice containing all NSC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNSC() ([]*NSC, error) {
	pss, err := m.ParseAll("NSC")
	return pss.([]*NSC), err
}
// AllNST returns a slice containing all NST segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNST() ([]*NST, error) {
	pss, err := m.ParseAll("NST")
	return pss.([]*NST), err
}
// AllNTE returns a slice containing all NTE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNTE() ([]*NTE, error) {
	pss, err := m.ParseAll("NTE")
	return pss.([]*NTE), err
}
// AllOBR returns a slice containing all OBR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOBR() ([]*OBR, error) {
	pss, err := m.ParseAll("OBR")
	return pss.([]*OBR), err
}
// AllOBX returns a slice containing all OBX segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOBX() ([]*OBX, error) {
	pss, err := m.ParseAll("OBX")
	return pss.([]*OBX), err
}
// AllODS returns a slice containing all ODS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllODS() ([]*ODS, error) {
	pss, err := m.ParseAll("ODS")
	return pss.([]*ODS), err
}
// AllODT returns a slice containing all ODT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllODT() ([]*ODT, error) {
	pss, err := m.ParseAll("ODT")
	return pss.([]*ODT), err
}
// AllOM1 returns a slice containing all OM1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM1() ([]*OM1, error) {
	pss, err := m.ParseAll("OM1")
	return pss.([]*OM1), err
}
// AllOM2 returns a slice containing all OM2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM2() ([]*OM2, error) {
	pss, err := m.ParseAll("OM2")
	return pss.([]*OM2), err
}
// AllOM3 returns a slice containing all OM3 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM3() ([]*OM3, error) {
	pss, err := m.ParseAll("OM3")
	return pss.([]*OM3), err
}
// AllOM4 returns a slice containing all OM4 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM4() ([]*OM4, error) {
	pss, err := m.ParseAll("OM4")
	return pss.([]*OM4), err
}
// AllOM5 returns a slice containing all OM5 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM5() ([]*OM5, error) {
	pss, err := m.ParseAll("OM5")
	return pss.([]*OM5), err
}
// AllOM6 returns a slice containing all OM6 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM6() ([]*OM6, error) {
	pss, err := m.ParseAll("OM6")
	return pss.([]*OM6), err
}
// AllORC returns a slice containing all ORC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORC() ([]*ORC, error) {
	pss, err := m.ParseAll("ORC")
	return pss.([]*ORC), err
}
// AllORO returns a slice containing all ORO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORO() ([]*ORO, error) {
	pss, err := m.ParseAll("ORO")
	return pss.([]*ORO), err
}
// AllPCR returns a slice containing all PCR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPCR() ([]*PCR, error) {
	pss, err := m.ParseAll("PCR")
	return pss.([]*PCR), err
}
// AllPD1 returns a slice containing all PD1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPD1() ([]*PD1, error) {
	pss, err := m.ParseAll("PD1")
	return pss.([]*PD1), err
}
// AllPDC returns a slice containing all PDC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPDC() ([]*PDC, error) {
	pss, err := m.ParseAll("PDC")
	return pss.([]*PDC), err
}
// AllPEO returns a slice containing all PEO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPEO() ([]*PEO, error) {
	pss, err := m.ParseAll("PEO")
	return pss.([]*PEO), err
}
// AllPES returns a slice containing all PES segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPES() ([]*PES, error) {
	pss, err := m.ParseAll("PES")
	return pss.([]*PES), err
}
// AllPID returns a slice containing all PID segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPID() ([]*PID, error) {
	pss, err := m.ParseAll("PID")
	return pss.([]*PID), err
}
// AllPR1 returns a slice containing all PR1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPR1() ([]*PR1, error) {
	pss, err := m.ParseAll("PR1")
	return pss.([]*PR1), err
}
// AllPRA returns a slice containing all PRA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRA() ([]*PRA, error) {
	pss, err := m.ParseAll("PRA")
	return pss.([]*PRA), err
}
// AllPRB returns a slice containing all PRB segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRB() ([]*PRB, error) {
	pss, err := m.ParseAll("PRB")
	return pss.([]*PRB), err
}
// AllPRC returns a slice containing all PRC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRC() ([]*PRC, error) {
	pss, err := m.ParseAll("PRC")
	return pss.([]*PRC), err
}
// AllPRD returns a slice containing all PRD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRD() ([]*PRD, error) {
	pss, err := m.ParseAll("PRD")
	return pss.([]*PRD), err
}
// AllPSH returns a slice containing all PSH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPSH() ([]*PSH, error) {
	pss, err := m.ParseAll("PSH")
	return pss.([]*PSH), err
}
// AllPTH returns a slice containing all PTH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPTH() ([]*PTH, error) {
	pss, err := m.ParseAll("PTH")
	return pss.([]*PTH), err
}
// AllPV1 returns a slice containing all PV1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPV1() ([]*PV1, error) {
	pss, err := m.ParseAll("PV1")
	return pss.([]*PV1), err
}
// AllPV2 returns a slice containing all PV2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPV2() ([]*PV2, error) {
	pss, err := m.ParseAll("PV2")
	return pss.([]*PV2), err
}
// AllQAK returns a slice containing all QAK segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQAK() ([]*QAK, error) {
	pss, err := m.ParseAll("QAK")
	return pss.([]*QAK), err
}
// AllQRD returns a slice containing all QRD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQRD() ([]*QRD, error) {
	pss, err := m.ParseAll("QRD")
	return pss.([]*QRD), err
}
// AllQRF returns a slice containing all QRF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQRF() ([]*QRF, error) {
	pss, err := m.ParseAll("QRF")
	return pss.([]*QRF), err
}
// AllRDF returns a slice containing all RDF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRDF() ([]*RDF, error) {
	pss, err := m.ParseAll("RDF")
	return pss.([]*RDF), err
}
// AllRDT returns a slice containing all RDT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRDT() ([]*RDT, error) {
	pss, err := m.ParseAll("RDT")
	return pss.([]*RDT), err
}
// AllRF1 returns a slice containing all RF1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRF1() ([]*RF1, error) {
	pss, err := m.ParseAll("RF1")
	return pss.([]*RF1), err
}
// AllRGS returns a slice containing all RGS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRGS() ([]*RGS, error) {
	pss, err := m.ParseAll("RGS")
	return pss.([]*RGS), err
}
// AllROL returns a slice containing all ROL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllROL() ([]*ROL, error) {
	pss, err := m.ParseAll("ROL")
	return pss.([]*ROL), err
}
// AllRQ1 returns a slice containing all RQ1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRQ1() ([]*RQ1, error) {
	pss, err := m.ParseAll("RQ1")
	return pss.([]*RQ1), err
}
// AllRQD returns a slice containing all RQD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRQD() ([]*RQD, error) {
	pss, err := m.ParseAll("RQD")
	return pss.([]*RQD), err
}
// AllRX1 returns a slice containing all RX1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRX1() ([]*RX1, error) {
	pss, err := m.ParseAll("RX1")
	return pss.([]*RX1), err
}
// AllRXA returns a slice containing all RXA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXA() ([]*RXA, error) {
	pss, err := m.ParseAll("RXA")
	return pss.([]*RXA), err
}
// AllRXC returns a slice containing all RXC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXC() ([]*RXC, error) {
	pss, err := m.ParseAll("RXC")
	return pss.([]*RXC), err
}
// AllRXD returns a slice containing all RXD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXD() ([]*RXD, error) {
	pss, err := m.ParseAll("RXD")
	return pss.([]*RXD), err
}
// AllRXE returns a slice containing all RXE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXE() ([]*RXE, error) {
	pss, err := m.ParseAll("RXE")
	return pss.([]*RXE), err
}
// AllRXG returns a slice containing all RXG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXG() ([]*RXG, error) {
	pss, err := m.ParseAll("RXG")
	return pss.([]*RXG), err
}
// AllRXO returns a slice containing all RXO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXO() ([]*RXO, error) {
	pss, err := m.ParseAll("RXO")
	return pss.([]*RXO), err
}
// AllRXR returns a slice containing all RXR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXR() ([]*RXR, error) {
	pss, err := m.ParseAll("RXR")
	return pss.([]*RXR), err
}
// AllSCH returns a slice containing all SCH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSCH() ([]*SCH, error) {
	pss, err := m.ParseAll("SCH")
	return pss.([]*SCH), err
}
// AllSPR returns a slice containing all SPR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSPR() ([]*SPR, error) {
	pss, err := m.ParseAll("SPR")
	return pss.([]*SPR), err
}
// AllSTF returns a slice containing all STF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSTF() ([]*STF, error) {
	pss, err := m.ParseAll("STF")
	return pss.([]*STF), err
}
// AllTXA returns a slice containing all TXA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllTXA() ([]*TXA, error) {
	pss, err := m.ParseAll("TXA")
	return pss.([]*TXA), err
}
// AllUB1 returns a slice containing all UB1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllUB1() ([]*UB1, error) {
	pss, err := m.ParseAll("UB1")
	return pss.([]*UB1), err
}
// AllUB2 returns a slice containing all UB2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllUB2() ([]*UB2, error) {
	pss, err := m.ParseAll("UB2")
	return pss.([]*UB2), err
}
// AllURD returns a slice containing all URD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllURD() ([]*URD, error) {
	pss, err := m.ParseAll("URD")
	return pss.([]*URD), err
}
// AllURS returns a slice containing all URS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllURS() ([]*URS, error) {
	pss, err := m.ParseAll("URS")
	return pss.([]*URS), err
}
// AllVAR returns a slice containing all VAR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllVAR() ([]*VAR, error) {
	pss, err := m.ParseAll("VAR")
	return pss.([]*VAR), err
}
// AllVTQ returns a slice containing all VTQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllVTQ() ([]*VTQ, error) {
	pss, err := m.ParseAll("VTQ")
	return pss.([]*VTQ), err
}

// v2 API
type ACKv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
}

func (m *ACKv2) MSH() *MSH {
	return m.msh
}

func (m *ACKv2) MSA() *MSA {
	return m.msa
}

func (m *ACKv2) ERR() *ERR {
	return m.err
}


func (m ACKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
	}, nil
}
type ADR_A19v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*ADR_A19_PIDv2 // Required
	dsc *DSC
}

func (m *ADR_A19v2) MSH() *MSH {
	return m.msh
}

func (m *ADR_A19v2) MSA() *MSA {
	return m.msa
}

func (m *ADR_A19v2) ERR() *ERR {
	return m.err
}

func (m *ADR_A19v2) QRD() *QRD {
	return m.qrd
}

func (m *ADR_A19v2) GroupByPID() []*ADR_A19_PIDv2 {
	return m.pid
}

func (m *ADR_A19v2) DSC() *DSC {
	return m.dsc
}


func (m ADR_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ADR_A19_PIDv2 struct {
	evn *EVN
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADR_A19_PID_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADR_A19_PIDv2) EVN() *EVN {
	return m.evn
}

func (m *ADR_A19_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADR_A19_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADR_A19_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ADR_A19_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ADR_A19_PIDv2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADR_A19_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADR_A19_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADR_A19_PIDv2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADR_A19_PIDv2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADR_A19_PIDv2) GroupByIN1() []*ADR_A19_PID_IN1v2 {
	return m.in1
}

func (m *ADR_A19_PIDv2) ACC() *ACC {
	return m.acc
}

func (m *ADR_A19_PIDv2) UB1() *UB1 {
	return m.ub1
}

func (m *ADR_A19_PIDv2) UB2() *UB2 {
	return m.ub2
}


func (m ADR_A19_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADR_A19_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADR_A19_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADR_A19_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADR_A19_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADR_A19_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A01v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	drg *DRG
	pr1 []*ADT_A01_PR1v2
	gt1 []*GT1
	in1 []*ADT_A01_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A01v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A01v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A01v2) PID() *PID {
	return m.pid
}

func (m *ADT_A01v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A01v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A01v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A01v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A01v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A01v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A01v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A01v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A01v2) DRG() *DRG {
	return m.drg
}

func (m *ADT_A01v2) GroupByPR1() []*ADT_A01_PR1v2 {
	return m.pr1
}

func (m *ADT_A01v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A01v2) GroupByIN1() []*ADT_A01_IN1v2 {
	return m.in1
}

func (m *ADT_A01v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A01v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A01v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"drg": m.drg,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A01_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *ADT_A01_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *ADT_A01_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m ADT_A01_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type ADT_A01_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A01_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A01_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A01_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A01_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A02v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
}

func (m *ADT_A02v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A02v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A02v2) PID() *PID {
	return m.pid
}

func (m *ADT_A02v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A02v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A02v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A02v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A02v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
	}, nil
}
type ADT_A03v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	dg1 []*DG1
	drg *DRG
	pr1 []*ADT_A03_PR1v2
	obx []*OBX
}

func (m *ADT_A03v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A03v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A03v2) PID() *PID {
	return m.pid
}

func (m *ADT_A03v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A03v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A03v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A03v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A03v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A03v2) DRG() *DRG {
	return m.drg
}

func (m *ADT_A03v2) GroupByPR1() []*ADT_A03_PR1v2 {
	return m.pr1
}

func (m *ADT_A03v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"dg1": m.dg1,
		"drg": m.drg,
		"pr1": m.pr1,
		"obx": m.obx,
	}, nil
}
type ADT_A03_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *ADT_A03_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *ADT_A03_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m ADT_A03_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type ADT_A04v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A04_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A04v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A04v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A04v2) PID() *PID {
	return m.pid
}

func (m *ADT_A04v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A04v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A04v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A04v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A04v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A04v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A04v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A04v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A04v2) GroupByIN1() []*ADT_A04_IN1v2 {
	return m.in1
}

func (m *ADT_A04v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A04v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A04v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A04_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A04_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A04_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A04_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A04_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A05v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A05_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A05v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A05v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A05v2) PID() *PID {
	return m.pid
}

func (m *ADT_A05v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A05v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A05v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A05v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A05v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A05v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A05v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A05v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A05v2) GroupByIN1() []*ADT_A05_IN1v2 {
	return m.in1
}

func (m *ADT_A05v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A05v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A05v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A05_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A05_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A05_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A05_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A05_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A06v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	mrg *MRG
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	drg1 *DRG
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	drg2 *DRG
	pr1 []*ADT_A06_PR1v2
	gt1 []*GT1
	in1 []*ADT_A06_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A06v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A06v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A06v2) PID() *PID {
	return m.pid
}

func (m *ADT_A06v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A06v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A06v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A06v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A06v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A06v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A06v2) DRG1() *DRG {
	return m.drg1
}

func (m *ADT_A06v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A06v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A06v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A06v2) DRG2() *DRG {
	return m.drg2
}

func (m *ADT_A06v2) GroupByPR1() []*ADT_A06_PR1v2 {
	return m.pr1
}

func (m *ADT_A06v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A06v2) GroupByIN1() []*ADT_A06_IN1v2 {
	return m.in1
}

func (m *ADT_A06v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A06v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A06v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"drg1": m.drg1,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"drg2": m.drg2,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A06_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *ADT_A06_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *ADT_A06_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m ADT_A06_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type ADT_A06_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A06_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A06_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A06_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A06_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A07v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A07_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A07v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A07v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A07v2) PID() *PID {
	return m.pid
}

func (m *ADT_A07v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A07v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A07v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A07v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A07v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A07v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A07v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A07v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A07v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A07v2) GroupByIN1() []*ADT_A07_IN1v2 {
	return m.in1
}

func (m *ADT_A07v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A07v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A07v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A07_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A07_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A07_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A07_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A07_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A08v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A08_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A08v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A08v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A08v2) PID() *PID {
	return m.pid
}

func (m *ADT_A08v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A08v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A08v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A08v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A08v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A08v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A08v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A08v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A08v2) GroupByIN1() []*ADT_A08_IN1v2 {
	return m.in1
}

func (m *ADT_A08v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A08v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A08v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A08_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A08_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A08_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A08_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A08_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A09v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A09v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A09v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A09v2) PID() *PID {
	return m.pid
}

func (m *ADT_A09v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A09v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A09v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A09v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A09v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A09v2) AllDG1() []*DG1 {
	return m.dg1
}


func (m ADT_A09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type ADT_A10v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A10v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A10v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A10v2) PID() *PID {
	return m.pid
}

func (m *ADT_A10v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A10v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A10v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A10v2) AllDG1() []*DG1 {
	return m.dg1
}


func (m ADT_A10v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type ADT_A11v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A11v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A11v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A11v2) PID() *PID {
	return m.pid
}

func (m *ADT_A11v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A11v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A11v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A11v2) AllDG1() []*DG1 {
	return m.dg1
}


func (m ADT_A11v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type ADT_A12v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	dg1 *DG1
}

func (m *ADT_A12v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A12v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A12v2) PID() *PID {
	return m.pid
}

func (m *ADT_A12v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A12v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A12v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A12v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A12v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A12v2) DG1() *DG1 {
	return m.dg1
}


func (m ADT_A12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type ADT_A13v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A13_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A13v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A13v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A13v2) PID() *PID {
	return m.pid
}

func (m *ADT_A13v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A13v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A13v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A13v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A13v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A13v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A13v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A13v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A13v2) GroupByIN1() []*ADT_A13_IN1v2 {
	return m.in1
}

func (m *ADT_A13v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A13v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A13v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A13v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A13_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A13_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A13_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A13_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A13_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A14v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A14_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A14v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A14v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A14v2) PID() *PID {
	return m.pid
}

func (m *ADT_A14v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A14v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A14v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A14v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A14v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A14v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A14v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A14v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A14v2) GroupByIN1() []*ADT_A14_IN1v2 {
	return m.in1
}

func (m *ADT_A14v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A14v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A14v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A14v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A14_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A14_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A14_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A14_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A14_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A15v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A15v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A15v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A15v2) PID() *PID {
	return m.pid
}

func (m *ADT_A15v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A15v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A15v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A15v2) AllDG1() []*DG1 {
	return m.dg1
}


func (m ADT_A15v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type ADT_A16v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	dg1 *DG1
	drg *DRG
}

func (m *ADT_A16v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A16v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A16v2) PID() *PID {
	return m.pid
}

func (m *ADT_A16v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A16v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A16v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A16v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A16v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A16v2) DG1() *DG1 {
	return m.dg1
}

func (m *ADT_A16v2) DRG() *DRG {
	return m.drg
}


func (m ADT_A16v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"dg1": m.dg1,
		"drg": m.drg,
	}, nil
}
type ADT_A17v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid1 *PID // Required
	pd11 *PD1
	pv11 *PV1 // Required
	pv21 *PV2
	db11 []*DB1
	obx1 []*OBX
	pid2 *PID // Required
	pd12 *PD1
	pv12 *PV1 // Required
	pv22 *PV2
	db12 []*DB1
	obx2 []*OBX
}

func (m *ADT_A17v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A17v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A17v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A17v2) PD11() *PD1 {
	return m.pd11
}

func (m *ADT_A17v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A17v2) PV21() *PV2 {
	return m.pv21
}

func (m *ADT_A17v2) AllDB11() []*DB1 {
	return m.db11
}

func (m *ADT_A17v2) AllOBX1() []*OBX {
	return m.obx1
}

func (m *ADT_A17v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A17v2) PD12() *PD1 {
	return m.pd12
}

func (m *ADT_A17v2) PV12() *PV1 {
	return m.pv12
}

func (m *ADT_A17v2) PV22() *PV2 {
	return m.pv22
}

func (m *ADT_A17v2) AllDB12() []*DB1 {
	return m.db12
}

func (m *ADT_A17v2) AllOBX2() []*OBX {
	return m.obx2
}


func (m ADT_A17v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid1": m.pid1,
		"pd11": m.pd11,
		"pv11": m.pv11,
		"pv21": m.pv21,
		"db11": m.db11,
		"obx1": m.obx1,
		"pid2": m.pid2,
		"pd12": m.pd12,
		"pv12": m.pv12,
		"pv22": m.pv22,
		"db12": m.db12,
		"obx2": m.obx2,
	}, nil
}
type ADT_A18v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	mrg *MRG
	pv1 *PV1 // Required
}

func (m *ADT_A18v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A18v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A18v2) PID() *PID {
	return m.pid
}

func (m *ADT_A18v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A18v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A18v2) PV1() *PV1 {
	return m.pv1
}


func (m ADT_A18v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}
type ADT_A20v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	npu *NPU // Required
}

func (m *ADT_A20v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A20v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A20v2) NPU() *NPU {
	return m.npu
}


func (m ADT_A20v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"npu": m.npu,
	}, nil
}
type ADT_A21v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A21v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A21v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A21v2) PID() *PID {
	return m.pid
}

func (m *ADT_A21v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A21v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A21v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A21v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A22v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A22v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A22v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A22v2) PID() *PID {
	return m.pid
}

func (m *ADT_A22v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A22v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A22v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A22v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A23v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A23v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A23v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A23v2) PID() *PID {
	return m.pid
}

func (m *ADT_A23v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A23v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A23v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A23v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A24v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid1 *PID // Required
	pd11 *PD1
	pv11 *PV1
	db11 []*DB1
	pid2 *PID // Required
	pd12 *PD1
	pv12 *PV1
	db12 []*DB1
}

func (m *ADT_A24v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A24v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A24v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A24v2) PD11() *PD1 {
	return m.pd11
}

func (m *ADT_A24v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A24v2) AllDB11() []*DB1 {
	return m.db11
}

func (m *ADT_A24v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A24v2) PD12() *PD1 {
	return m.pd12
}

func (m *ADT_A24v2) PV12() *PV1 {
	return m.pv12
}

func (m *ADT_A24v2) AllDB12() []*DB1 {
	return m.db12
}


func (m ADT_A24v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid1": m.pid1,
		"pd11": m.pd11,
		"pv11": m.pv11,
		"db11": m.db11,
		"pid2": m.pid2,
		"pd12": m.pd12,
		"pv12": m.pv12,
		"db12": m.db12,
	}, nil
}
type ADT_A25v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A25v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A25v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A25v2) PID() *PID {
	return m.pid
}

func (m *ADT_A25v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A25v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A25v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A25v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A26v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A26v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A26v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A26v2) PID() *PID {
	return m.pid
}

func (m *ADT_A26v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A26v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A26v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A26v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A27v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A27v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A27v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A27v2) PID() *PID {
	return m.pid
}

func (m *ADT_A27v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A27v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A27v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A27v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A28v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A28_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A28v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A28v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A28v2) PID() *PID {
	return m.pid
}

func (m *ADT_A28v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A28v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A28v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A28v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A28v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A28v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A28v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A28v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A28v2) GroupByIN1() []*ADT_A28_IN1v2 {
	return m.in1
}

func (m *ADT_A28v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A28v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A28v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A28v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A28_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A28_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A28_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A28_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A28_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A29v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A29v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A29v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A29v2) PID() *PID {
	return m.pid
}

func (m *ADT_A29v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A29v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A29v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A29v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A30v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
}

func (m *ADT_A30v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A30v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A30v2) PID() *PID {
	return m.pid
}

func (m *ADT_A30v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A30v2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A30v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
	}, nil
}
type ADT_A31v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A31_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A31v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A31v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A31v2) PID() *PID {
	return m.pid
}

func (m *ADT_A31v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A31v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A31v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A31v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A31v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A31v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A31v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A31v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A31v2) GroupByIN1() []*ADT_A31_IN1v2 {
	return m.in1
}

func (m *ADT_A31v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A31v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A31v2) UB2() *UB2 {
	return m.ub2
}


func (m ADT_A31v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ADT_A31_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A31_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A31_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A31_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ADT_A31_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ADT_A32v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A32v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A32v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A32v2) PID() *PID {
	return m.pid
}

func (m *ADT_A32v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A32v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A32v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A32v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A33v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A33v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A33v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A33v2) PID() *PID {
	return m.pid
}

func (m *ADT_A33v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A33v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A33v2) AllOBX() []*OBX {
	return m.obx
}


func (m ADT_A33v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}
type ADT_A34v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A34v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A34v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A34v2) PID() *PID {
	return m.pid
}

func (m *ADT_A34v2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A34v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}
type ADT_A35v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A35v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A35v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A35v2) PID() *PID {
	return m.pid
}

func (m *ADT_A35v2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A35v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}
type ADT_A36v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A36v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A36v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A36v2) PID() *PID {
	return m.pid
}

func (m *ADT_A36v2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A36v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}
type ADT_A37v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid1 *PID // Required
	pv11 *PV1
	pid2 *PID // Required
	pv12 *PV1
}

func (m *ADT_A37v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A37v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A37v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A37v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A37v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A37v2) PV12() *PV1 {
	return m.pv12
}


func (m ADT_A37v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid1": m.pid1,
		"pv11": m.pv11,
		"pid2": m.pid2,
		"pv12": m.pv12,
	}, nil
}
type ADT_A38v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	dg1 []*DG1
	drg *DRG
}

func (m *ADT_A38v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A38v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A38v2) PID() *PID {
	return m.pid
}

func (m *ADT_A38v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A38v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A38v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A38v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ADT_A38v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A38v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A38v2) DRG() *DRG {
	return m.drg
}


func (m ADT_A38v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"dg1": m.dg1,
		"drg": m.drg,
	}, nil
}
type ADT_A39v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*ADT_A39_PIDv2 // Required
}

func (m *ADT_A39v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A39v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A39v2) GroupByPID() []*ADT_A39_PIDv2 {
	return m.pid
}


func (m ADT_A39v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type ADT_A39_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
	pv1 *PV1
}

func (m *ADT_A39_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A39_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A39_PIDv2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A39_PIDv2) PV1() *PV1 {
	return m.pv1
}


func (m ADT_A39_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}
type ADT_A40v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*ADT_A40_PIDv2 // Required
}

func (m *ADT_A40v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A40v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A40v2) GroupByPID() []*ADT_A40_PIDv2 {
	return m.pid
}


func (m ADT_A40v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type ADT_A40_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
	pv1 *PV1
}

func (m *ADT_A40_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A40_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A40_PIDv2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A40_PIDv2) PV1() *PV1 {
	return m.pv1
}


func (m ADT_A40_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}
type ADT_A43v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*ADT_A43_PIDv2 // Required
}

func (m *ADT_A43v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A43v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A43v2) GroupByPID() []*ADT_A43_PIDv2 {
	return m.pid
}


func (m ADT_A43v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type ADT_A43_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
}

func (m *ADT_A43_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A43_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A43_PIDv2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A43_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
	}, nil
}
type ADT_A44v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*ADT_A44_PIDv2 // Required
}

func (m *ADT_A44v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A44v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A44v2) GroupByPID() []*ADT_A44_PIDv2 {
	return m.pid
}


func (m ADT_A44v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type ADT_A44_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
}

func (m *ADT_A44_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A44_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A44_PIDv2) MRG() *MRG {
	return m.mrg
}


func (m ADT_A44_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
	}, nil
}
type ADT_A45v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	mrg []*ADT_A45_MRGv2 // Required
}

func (m *ADT_A45v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A45v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A45v2) PID() *PID {
	return m.pid
}

func (m *ADT_A45v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A45v2) GroupByMRG() []*ADT_A45_MRGv2 {
	return m.mrg
}


func (m ADT_A45v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
	}, nil
}
type ADT_A45_MRGv2 struct {
	mrg *MRG // Required
	pv1 *PV1 // Required
}

func (m *ADT_A45_MRGv2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A45_MRGv2) PV1() *PV1 {
	return m.pv1
}


func (m ADT_A45_MRGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}
type ADT_A50v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
	pv1 *PV1 // Required
}

func (m *ADT_A50v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A50v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A50v2) PID() *PID {
	return m.pid
}

func (m *ADT_A50v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A50v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A50v2) PV1() *PV1 {
	return m.pv1
}


func (m ADT_A50v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}
type ARD_A19v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	pid []*ARD_A19_PIDv2 // Required
	dsc *DSC
}

func (m *ARD_A19v2) MSH() *MSH {
	return m.msh
}

func (m *ARD_A19v2) MSA() *MSA {
	return m.msa
}

func (m *ARD_A19v2) ERR() *ERR {
	return m.err
}

func (m *ARD_A19v2) QRD() *QRD {
	return m.qrd
}

func (m *ARD_A19v2) QRF() *QRF {
	return m.qrf
}

func (m *ARD_A19v2) GroupByPID() []*ARD_A19_PIDv2 {
	return m.pid
}

func (m *ARD_A19v2) DSC() *DSC {
	return m.dsc
}


func (m ARD_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ARD_A19_PIDv2 struct {
	evn *EVN
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	drg *DRG
	pr1 []*ARD_A19_PID_PR1v2
	gt1 []*GT1
	in1 []*ARD_A19_PID_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ARD_A19_PIDv2) EVN() *EVN {
	return m.evn
}

func (m *ARD_A19_PIDv2) PID() *PID {
	return m.pid
}

func (m *ARD_A19_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ARD_A19_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ARD_A19_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ARD_A19_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ARD_A19_PIDv2) AllDB1() []*DB1 {
	return m.db1
}

func (m *ARD_A19_PIDv2) AllOBX() []*OBX {
	return m.obx
}

func (m *ARD_A19_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ARD_A19_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ARD_A19_PIDv2) DRG() *DRG {
	return m.drg
}

func (m *ARD_A19_PIDv2) GroupByPR1() []*ARD_A19_PID_PR1v2 {
	return m.pr1
}

func (m *ARD_A19_PIDv2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ARD_A19_PIDv2) GroupByIN1() []*ARD_A19_PID_IN1v2 {
	return m.in1
}

func (m *ARD_A19_PIDv2) ACC() *ACC {
	return m.acc
}

func (m *ARD_A19_PIDv2) UB1() *UB1 {
	return m.ub1
}

func (m *ARD_A19_PIDv2) UB2() *UB2 {
	return m.ub2
}


func (m ARD_A19_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"drg": m.drg,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type ARD_A19_PID_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *ARD_A19_PID_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *ARD_A19_PID_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m ARD_A19_PID_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type ARD_A19_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ARD_A19_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ARD_A19_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ARD_A19_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ARD_A19_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type BAR_P01v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 []*BAR_P01_PV1v2 // Required
}

func (m *BAR_P01v2) MSH() *MSH {
	return m.msh
}

func (m *BAR_P01v2) EVN() *EVN {
	return m.evn
}

func (m *BAR_P01v2) PID() *PID {
	return m.pid
}

func (m *BAR_P01v2) PD1() *PD1 {
	return m.pd1
}

func (m *BAR_P01v2) GroupByPV1() []*BAR_P01_PV1v2 {
	return m.pv1
}


func (m BAR_P01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
	}, nil
}
type BAR_P01_PV1v2 struct {
	pv1 *PV1
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	drg *DRG
	pr1 []*BAR_P01_PV1_PR1v2
	gt1 []*GT1
	nk1 []*NK1
	in1 []*BAR_P01_PV1_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *BAR_P01_PV1v2) PV1() *PV1 {
	return m.pv1
}

func (m *BAR_P01_PV1v2) PV2() *PV2 {
	return m.pv2
}

func (m *BAR_P01_PV1v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *BAR_P01_PV1v2) AllOBX() []*OBX {
	return m.obx
}

func (m *BAR_P01_PV1v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *BAR_P01_PV1v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *BAR_P01_PV1v2) DRG() *DRG {
	return m.drg
}

func (m *BAR_P01_PV1v2) GroupByPR1() []*BAR_P01_PV1_PR1v2 {
	return m.pr1
}

func (m *BAR_P01_PV1v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *BAR_P01_PV1v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *BAR_P01_PV1v2) GroupByIN1() []*BAR_P01_PV1_IN1v2 {
	return m.in1
}

func (m *BAR_P01_PV1v2) ACC() *ACC {
	return m.acc
}

func (m *BAR_P01_PV1v2) UB1() *UB1 {
	return m.ub1
}

func (m *BAR_P01_PV1v2) UB2() *UB2 {
	return m.ub2
}


func (m BAR_P01_PV1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"drg": m.drg,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"nk1": m.nk1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}
type BAR_P01_PV1_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *BAR_P01_PV1_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *BAR_P01_PV1_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m BAR_P01_PV1_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type BAR_P01_PV1_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *BAR_P01_PV1_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *BAR_P01_PV1_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *BAR_P01_PV1_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m BAR_P01_PV1_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type BAR_P02v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*BAR_P02_PIDv2 // Required
}

func (m *BAR_P02v2) MSH() *MSH {
	return m.msh
}

func (m *BAR_P02v2) EVN() *EVN {
	return m.evn
}

func (m *BAR_P02v2) GroupByPID() []*BAR_P02_PIDv2 {
	return m.pid
}


func (m BAR_P02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type BAR_P02_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1
	db1 []*DB1
}

func (m *BAR_P02_PIDv2) PID() *PID {
	return m.pid
}

func (m *BAR_P02_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *BAR_P02_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *BAR_P02_PIDv2) AllDB1() []*DB1 {
	return m.db1
}


func (m BAR_P02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"db1": m.db1,
	}, nil
}
type BAR_P06v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid []*BAR_P06_PIDv2 // Required
}

func (m *BAR_P06v2) MSH() *MSH {
	return m.msh
}

func (m *BAR_P06v2) EVN() *EVN {
	return m.evn
}

func (m *BAR_P06v2) GroupByPID() []*BAR_P06_PIDv2 {
	return m.pid
}


func (m BAR_P06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}
type BAR_P06_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
}

func (m *BAR_P06_PIDv2) PID() *PID {
	return m.pid
}

func (m *BAR_P06_PIDv2) PV1() *PV1 {
	return m.pv1
}


func (m BAR_P06_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}
type CRM_C01v2 struct {
	msh *MSH // Required
	pid []*CRM_C01_PIDv2 // Required
}

func (m *CRM_C01v2) MSH() *MSH {
	return m.msh
}

func (m *CRM_C01v2) GroupByPID() []*CRM_C01_PIDv2 {
	return m.pid
}


func (m CRM_C01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
	}, nil
}
type CRM_C01_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
	csr *CSR // Required
	csp []*CSP
}

func (m *CRM_C01_PIDv2) PID() *PID {
	return m.pid
}

func (m *CRM_C01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *CRM_C01_PIDv2) CSR() *CSR {
	return m.csr
}

func (m *CRM_C01_PIDv2) AllCSP() []*CSP {
	return m.csp
}


func (m CRM_C01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"csr": m.csr,
		"csp": m.csp,
	}, nil
}
type CSU_C09v2 struct {
	msh *MSH // Required
	pid []*CSU_C09_PIDv2 // Required
}

func (m *CSU_C09v2) MSH() *MSH {
	return m.msh
}

func (m *CSU_C09v2) GroupByPID() []*CSU_C09_PIDv2 {
	return m.pid
}


func (m CSU_C09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
	}, nil
}
type CSU_C09_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	csr *CSR // Required
	csp []*CSU_C09_PID_CSPv2 // Required
}

func (m *CSU_C09_PIDv2) PID() *PID {
	return m.pid
}

func (m *CSU_C09_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *CSU_C09_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *CSU_C09_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *CSU_C09_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *CSU_C09_PIDv2) CSR() *CSR {
	return m.csr
}

func (m *CSU_C09_PIDv2) GroupByCSP() []*CSU_C09_PID_CSPv2 {
	return m.csp
}


func (m CSU_C09_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"csr": m.csr,
		"csp": m.csp,
	}, nil
}
type CSU_C09_PID_CSPv2 struct {
	csp *CSP
	css []*CSU_C09_PID_CSP_CSSv2 // Required
}

func (m *CSU_C09_PID_CSPv2) CSP() *CSP {
	return m.csp
}

func (m *CSU_C09_PID_CSPv2) GroupByCSS() []*CSU_C09_PID_CSP_CSSv2 {
	return m.css
}


func (m CSU_C09_PID_CSPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"csp": m.csp,
		"css": m.css,
	}, nil
}
type CSU_C09_PID_CSP_CSSv2 struct {
	css *CSS
	obr []*CSU_C09_PID_CSP_CSS_OBRv2
	orc []*CSU_C09_PID_CSP_CSS_ORCv2 // Required
}

func (m *CSU_C09_PID_CSP_CSSv2) CSS() *CSS {
	return m.css
}

func (m *CSU_C09_PID_CSP_CSSv2) GroupByOBR() []*CSU_C09_PID_CSP_CSS_OBRv2 {
	return m.obr
}

func (m *CSU_C09_PID_CSP_CSSv2) GroupByORC() []*CSU_C09_PID_CSP_CSS_ORCv2 {
	return m.orc
}


func (m CSU_C09_PID_CSP_CSSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"css": m.css,
		"obr": m.obr,
		"orc": m.orc,
	}, nil
}
type CSU_C09_PID_CSP_CSS_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	obx []*OBX // Required
}

func (m *CSU_C09_PID_CSP_CSS_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *CSU_C09_PID_CSP_CSS_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *CSU_C09_PID_CSP_CSS_OBRv2) AllOBX() []*OBX {
	return m.obx
}


func (m CSU_C09_PID_CSP_CSS_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"obx": m.obx,
	}, nil
}
type CSU_C09_PID_CSP_CSS_ORCv2 struct {
	orc *ORC
	rxa []*CSU_C09_PID_CSP_CSS_ORC_RXAv2 // Required
}

func (m *CSU_C09_PID_CSP_CSS_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *CSU_C09_PID_CSP_CSS_ORCv2) GroupByRXA() []*CSU_C09_PID_CSP_CSS_ORC_RXAv2 {
	return m.rxa
}


func (m CSU_C09_PID_CSP_CSS_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxa": m.rxa,
	}, nil
}
type CSU_C09_PID_CSP_CSS_ORC_RXAv2 struct {
	rxa *RXA // Required
	rxr *RXR // Required
}

func (m *CSU_C09_PID_CSP_CSS_ORC_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *CSU_C09_PID_CSP_CSS_ORC_RXAv2) RXR() *RXR {
	return m.rxr
}


func (m CSU_C09_PID_CSP_CSS_ORC_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxa": m.rxa,
		"rxr": m.rxr,
	}, nil
}
type DFT_P03v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	pv1 *PV1
	pv2 *PV2
	db1 []*DB1
	obx []*OBX
	ft1 []*DFT_P03_FT1v2 // Required
	dg1 []*DG1
	drg *DRG
	gt1 []*GT1
	in1 []*DFT_P03_IN1v2
	acc *ACC
}

func (m *DFT_P03v2) MSH() *MSH {
	return m.msh
}

func (m *DFT_P03v2) EVN() *EVN {
	return m.evn
}

func (m *DFT_P03v2) PID() *PID {
	return m.pid
}

func (m *DFT_P03v2) PD1() *PD1 {
	return m.pd1
}

func (m *DFT_P03v2) PV1() *PV1 {
	return m.pv1
}

func (m *DFT_P03v2) PV2() *PV2 {
	return m.pv2
}

func (m *DFT_P03v2) AllDB1() []*DB1 {
	return m.db1
}

func (m *DFT_P03v2) AllOBX() []*OBX {
	return m.obx
}

func (m *DFT_P03v2) GroupByFT1() []*DFT_P03_FT1v2 {
	return m.ft1
}

func (m *DFT_P03v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *DFT_P03v2) DRG() *DRG {
	return m.drg
}

func (m *DFT_P03v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *DFT_P03v2) GroupByIN1() []*DFT_P03_IN1v2 {
	return m.in1
}

func (m *DFT_P03v2) ACC() *ACC {
	return m.acc
}


func (m DFT_P03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"db1": m.db1,
		"obx": m.obx,
		"ft1": m.ft1,
		"dg1": m.dg1,
		"drg": m.drg,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
	}, nil
}
type DFT_P03_FT1v2 struct {
	ft1 *FT1 // Required
	pr1 []*DFT_P03_FT1_PR1v2
}

func (m *DFT_P03_FT1v2) FT1() *FT1 {
	return m.ft1
}

func (m *DFT_P03_FT1v2) GroupByPR1() []*DFT_P03_FT1_PR1v2 {
	return m.pr1
}


func (m DFT_P03_FT1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ft1": m.ft1,
		"pr1": m.pr1,
	}, nil
}
type DFT_P03_FT1_PR1v2 struct {
	pr1 *PR1 // Required
	rol []*ROL
}

func (m *DFT_P03_FT1_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *DFT_P03_FT1_PR1v2) AllROL() []*ROL {
	return m.rol
}


func (m DFT_P03_FT1_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"rol": m.rol,
	}, nil
}
type DFT_P03_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *DFT_P03_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *DFT_P03_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *DFT_P03_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m DFT_P03_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type DOC_T12v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*DOC_T12_PIDv2 // Required
	dsc *DSC
}

func (m *DOC_T12v2) MSH() *MSH {
	return m.msh
}

func (m *DOC_T12v2) MSA() *MSA {
	return m.msa
}

func (m *DOC_T12v2) ERR() *ERR {
	return m.err
}

func (m *DOC_T12v2) QRD() *QRD {
	return m.qrd
}

func (m *DOC_T12v2) GroupByPID() []*DOC_T12_PIDv2 {
	return m.pid
}

func (m *DOC_T12v2) DSC() *DSC {
	return m.dsc
}


func (m DOC_T12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type DOC_T12_PIDv2 struct {
	evn *EVN
	pid *PID // Required
	pv1 *PV1 // Required
	txa *TXA // Required
	obx []*OBX
}

func (m *DOC_T12_PIDv2) EVN() *EVN {
	return m.evn
}

func (m *DOC_T12_PIDv2) PID() *PID {
	return m.pid
}

func (m *DOC_T12_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *DOC_T12_PIDv2) TXA() *TXA {
	return m.txa
}

func (m *DOC_T12_PIDv2) AllOBX() []*OBX {
	return m.obx
}


func (m DOC_T12_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"txa": m.txa,
		"obx": m.obx,
	}, nil
}
type DSR_P04v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_P04v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_P04v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_P04v2) ERR() *ERR {
	return m.err
}

func (m *DSR_P04v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_P04v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_P04v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_P04v2) DSC() *DSC {
	return m.dsc
}


func (m DSR_P04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type DSR_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_Q01v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_Q01v2) ERR() *ERR {
	return m.err
}

func (m *DSR_Q01v2) QAK() *QAK {
	return m.qak
}

func (m *DSR_Q01v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_Q01v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_Q01v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m DSR_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type DSR_Q03v2 struct {
	msh *MSH // Required
	msa *MSA
	err *ERR
	qak *QAK
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_Q03v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_Q03v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_Q03v2) ERR() *ERR {
	return m.err
}

func (m *DSR_Q03v2) QAK() *QAK {
	return m.qak
}

func (m *DSR_Q03v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_Q03v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_Q03v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_Q03v2) DSC() *DSC {
	return m.dsc
}


func (m DSR_Q03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type DSR_R03v2 struct {
	msh *MSH // Required
	msa *MSA
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_R03v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_R03v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_R03v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_R03v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_R03v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_R03v2) DSC() *DSC {
	return m.dsc
}


func (m DSR_R03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type EDR_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK // Required
	dsp []*DSP // Required
	dsc *DSC
}

func (m *EDR_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *EDR_Q01v2) MSA() *MSA {
	return m.msa
}

func (m *EDR_Q01v2) ERR() *ERR {
	return m.err
}

func (m *EDR_Q01v2) QAK() *QAK {
	return m.qak
}

func (m *EDR_Q01v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *EDR_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m EDR_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type EQQ_Q01v2 struct {
	msh *MSH // Required
	eql *EQL // Required
	dsc *DSC
}

func (m *EQQ_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *EQQ_Q01v2) EQL() *EQL {
	return m.eql
}

func (m *EQQ_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m EQQ_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"eql": m.eql,
		"dsc": m.dsc,
	}, nil
}
type ERP_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK // Required
	erq *ERQ // Required
	dsc *DSC
}

func (m *ERP_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *ERP_Q01v2) MSA() *MSA {
	return m.msa
}

func (m *ERP_Q01v2) ERR() *ERR {
	return m.err
}

func (m *ERP_Q01v2) QAK() *QAK {
	return m.qak
}

func (m *ERP_Q01v2) ERQ() *ERQ {
	return m.erq
}

func (m *ERP_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m ERP_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"erq": m.erq,
		"dsc": m.dsc,
	}, nil
}
type MCF_Q02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
}

func (m *MCF_Q02v2) MSH() *MSH {
	return m.msh
}

func (m *MCF_Q02v2) MSA() *MSA {
	return m.msa
}


func (m MCF_Q02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
	}, nil
}
type MDM_T01v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	txa *TXA // Required
}

func (m *MDM_T01v2) MSH() *MSH {
	return m.msh
}

func (m *MDM_T01v2) EVN() *EVN {
	return m.evn
}

func (m *MDM_T01v2) PID() *PID {
	return m.pid
}

func (m *MDM_T01v2) PV1() *PV1 {
	return m.pv1
}

func (m *MDM_T01v2) TXA() *TXA {
	return m.txa
}


func (m MDM_T01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"txa": m.txa,
	}, nil
}
type MDM_T02v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	txa *TXA // Required
	obx []*OBX // Required
}

func (m *MDM_T02v2) MSH() *MSH {
	return m.msh
}

func (m *MDM_T02v2) EVN() *EVN {
	return m.evn
}

func (m *MDM_T02v2) PID() *PID {
	return m.pid
}

func (m *MDM_T02v2) PV1() *PV1 {
	return m.pv1
}

func (m *MDM_T02v2) TXA() *TXA {
	return m.txa
}

func (m *MDM_T02v2) AllOBX() []*OBX {
	return m.obx
}


func (m MDM_T02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"txa": m.txa,
		"obx": m.obx,
	}, nil
}
type MFD_M01v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M01v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFD_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFD_M02v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M02v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFD_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFD_M03v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M03v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFD_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFK_M01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M01v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M01v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFK_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFK_M02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M02v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M02v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFK_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFK_M03v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M03v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M03v2) ERR() *ERR {
	return m.err
}

func (m *MFK_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M03v2) AllMFA() []*MFA {
	return m.mfa
}


func (m MFK_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}
type MFN_M01v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M01_MFEv2 // Required
}

func (m *MFN_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M01v2) GroupByMFE() []*MFN_M01_MFEv2 {
	return m.mfe
}


func (m MFN_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M01_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFN_M01_MFEv2) MFE() *MFE {
	return m.mfe
}


func (m MFN_M01_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}
type MFN_M02v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M02_MFEv2 // Required
}

func (m *MFN_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M02v2) GroupByMFE() []*MFN_M02_MFEv2 {
	return m.mfe
}


func (m MFN_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M02_MFEv2 struct {
	mfe *MFE // Required
	stf *STF // Required
	pra *PRA
}

func (m *MFN_M02_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M02_MFEv2) STF() *STF {
	return m.stf
}

func (m *MFN_M02_MFEv2) PRA() *PRA {
	return m.pra
}


func (m MFN_M02_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"stf": m.stf,
		"pra": m.pra,
	}, nil
}
type MFN_M03v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M03_MFEv2 // Required
}

func (m *MFN_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M03v2) GroupByMFE() []*MFN_M03_MFEv2 {
	return m.mfe
}


func (m MFN_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M03_MFEv2 struct {
	mfe *MFE // Required
	om1 *OM1 // Required
}

func (m *MFN_M03_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M03_MFEv2) OM1() *OM1 {
	return m.om1
}


func (m MFN_M03_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"om1": m.om1,
	}, nil
}
type MFN_M05v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M05_MFEv2 // Required
}

func (m *MFN_M05v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M05v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M05v2) GroupByMFE() []*MFN_M05_MFEv2 {
	return m.mfe
}


func (m MFN_M05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M05_MFEv2 struct {
	mfe *MFE // Required
	loc *LOC // Required
	lch []*LCH
	lrl []*LRL
	ldp []*MFN_M05_MFE_LDPv2 // Required
}

func (m *MFN_M05_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M05_MFEv2) LOC() *LOC {
	return m.loc
}

func (m *MFN_M05_MFEv2) AllLCH() []*LCH {
	return m.lch
}

func (m *MFN_M05_MFEv2) AllLRL() []*LRL {
	return m.lrl
}

func (m *MFN_M05_MFEv2) GroupByLDP() []*MFN_M05_MFE_LDPv2 {
	return m.ldp
}


func (m MFN_M05_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"loc": m.loc,
		"lch": m.lch,
		"lrl": m.lrl,
		"ldp": m.ldp,
	}, nil
}
type MFN_M05_MFE_LDPv2 struct {
	ldp *LDP // Required
	lch []*LCH
	lcc []*LCC
}

func (m *MFN_M05_MFE_LDPv2) LDP() *LDP {
	return m.ldp
}

func (m *MFN_M05_MFE_LDPv2) AllLCH() []*LCH {
	return m.lch
}

func (m *MFN_M05_MFE_LDPv2) AllLCC() []*LCC {
	return m.lcc
}


func (m MFN_M05_MFE_LDPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ldp": m.ldp,
		"lch": m.lch,
		"lcc": m.lcc,
	}, nil
}
type MFN_M06v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M06_MFEv2 // Required
}

func (m *MFN_M06v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M06v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M06v2) GroupByMFE() []*MFN_M06_MFEv2 {
	return m.mfe
}


func (m MFN_M06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M06_MFEv2 struct {
	mfe *MFE // Required
	cdm *CDM // Required
	prc []*PRC
}

func (m *MFN_M06_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M06_MFEv2) CDM() *CDM {
	return m.cdm
}

func (m *MFN_M06_MFEv2) AllPRC() []*PRC {
	return m.prc
}


func (m MFN_M06_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"cdm": m.cdm,
		"prc": m.prc,
	}, nil
}
type MFN_M07v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M07_MFEv2 // Required
}

func (m *MFN_M07v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M07v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M07v2) GroupByMFE() []*MFN_M07_MFEv2 {
	return m.mfe
}


func (m MFN_M07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M07_MFEv2 struct {
	mfe *MFE // Required
	cm0 *CM0 // Required
	cm1 []*MFN_M07_MFE_CM1v2
}

func (m *MFN_M07_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M07_MFEv2) CM0() *CM0 {
	return m.cm0
}

func (m *MFN_M07_MFEv2) GroupByCM1() []*MFN_M07_MFE_CM1v2 {
	return m.cm1
}


func (m MFN_M07_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"cm0": m.cm0,
		"cm1": m.cm1,
	}, nil
}
type MFN_M07_MFE_CM1v2 struct {
	cm1 *CM1 // Required
	cm2 []*CM2
}

func (m *MFN_M07_MFE_CM1v2) CM1() *CM1 {
	return m.cm1
}

func (m *MFN_M07_MFE_CM1v2) AllCM2() []*CM2 {
	return m.cm2
}


func (m MFN_M07_MFE_CM1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"cm1": m.cm1,
		"cm2": m.cm2,
	}, nil
}
type MFN_M08v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M08_MFEv2 // Required
}

func (m *MFN_M08v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M08v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M08v2) GroupByMFE() []*MFN_M08_MFEv2 {
	return m.mfe
}


func (m MFN_M08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M08_MFEv2 struct {
	mfe *MFE // Required
	om1 *OM1 // Required
	om2 *OM2
	om3 *OM3
	om4 *OM4
}

func (m *MFN_M08_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M08_MFEv2) OM1() *OM1 {
	return m.om1
}

func (m *MFN_M08_MFEv2) OM2() *OM2 {
	return m.om2
}

func (m *MFN_M08_MFEv2) OM3() *OM3 {
	return m.om3
}

func (m *MFN_M08_MFEv2) OM4() *OM4 {
	return m.om4
}


func (m MFN_M08_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"om1": m.om1,
		"om2": m.om2,
		"om3": m.om3,
		"om4": m.om4,
	}, nil
}
type MFN_M09v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M09_MFEv2 // Required
}

func (m *MFN_M09v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M09v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M09v2) GroupByMFE() []*MFN_M09_MFEv2 {
	return m.mfe
}


func (m MFN_M09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M09_MFEv2 struct {
	mfe *MFE // Required
	om3 *OM3 // Required
	om4 []*OM4
}

func (m *MFN_M09_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M09_MFEv2) OM3() *OM3 {
	return m.om3
}

func (m *MFN_M09_MFEv2) AllOM4() []*OM4 {
	return m.om4
}


func (m MFN_M09_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"om3": m.om3,
		"om4": m.om4,
	}, nil
}
type MFN_M10v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	om5 []*MFN_M10_OM5v2 // Required
}

func (m *MFN_M10v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M10v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M10v2) GroupByOM5() []*MFN_M10_OM5v2 {
	return m.om5
}


func (m MFN_M10v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"om5": m.om5,
	}, nil
}
type MFN_M10_OM5v2 struct {
	om5 *OM5 // Required
	om4 []*OM4
}

func (m *MFN_M10_OM5v2) OM5() *OM5 {
	return m.om5
}

func (m *MFN_M10_OM5v2) AllOM4() []*OM4 {
	return m.om4
}


func (m MFN_M10_OM5v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"om5": m.om5,
		"om4": m.om4,
	}, nil
}
type MFN_M11v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfe []*MFN_M11_MFEv2 // Required
}

func (m *MFN_M11v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M11v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M11v2) GroupByMFE() []*MFN_M11_MFEv2 {
	return m.mfe
}


func (m MFN_M11v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}
type MFN_M11_MFEv2 struct {
	mfe *MFE // Required
	om1 *OM1 // Required
	om6 *OM6 // Required
	om2 *OM2 // Required
}

func (m *MFN_M11_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m *MFN_M11_MFEv2) OM1() *OM1 {
	return m.om1
}

func (m *MFN_M11_MFEv2) OM6() *OM6 {
	return m.om6
}

func (m *MFN_M11_MFEv2) OM2() *OM2 {
	return m.om2
}


func (m MFN_M11_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
		"om1": m.om1,
		"om6": m.om6,
		"om2": m.om2,
	}, nil
}
type MFQ_M01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M01v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M01v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M01v2) DSC() *DSC {
	return m.dsc
}


func (m MFQ_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type MFQ_M02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M02v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M02v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M02v2) DSC() *DSC {
	return m.dsc
}


func (m MFQ_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type MFQ_M03v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M03v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M03v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M03v2) DSC() *DSC {
	return m.dsc
}


func (m MFQ_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type MFR_M01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI // Required
	mfe []*MFR_M01_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M01v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M01v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M01v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M01v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M01v2) GroupByMFE() []*MFR_M01_MFEv2 {
	return m.mfe
}

func (m *MFR_M01v2) DSC() *DSC {
	return m.dsc
}


func (m MFR_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}
type MFR_M01_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M01_MFEv2) MFE() *MFE {
	return m.mfe
}


func (m MFR_M01_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}
type MFR_M02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI // Required
	mfe []*MFR_M02_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M02v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M02v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M02v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M02v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M02v2) GroupByMFE() []*MFR_M02_MFEv2 {
	return m.mfe
}

func (m *MFR_M02v2) DSC() *DSC {
	return m.dsc
}


func (m MFR_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}
type MFR_M02_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M02_MFEv2) MFE() *MFE {
	return m.mfe
}


func (m MFR_M02_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}
type MFR_M03v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI // Required
	mfe []*MFR_M03_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M03v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M03v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M03v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M03v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M03v2) GroupByMFE() []*MFR_M03_MFEv2 {
	return m.mfe
}

func (m *MFR_M03v2) DSC() *DSC {
	return m.dsc
}


func (m MFR_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}
type MFR_M03_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M03_MFEv2) MFE() *MFE {
	return m.mfe
}


func (m MFR_M03_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}
type NMD_N01v2 struct {
	msh *MSH // Required
	nck []*NMD_N01_NCKv2 // Required
}

func (m *NMD_N01v2) MSH() *MSH {
	return m.msh
}

func (m *NMD_N01v2) GroupByNCK() []*NMD_N01_NCKv2 {
	return m.nck
}


func (m NMD_N01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nck": m.nck,
	}, nil
}
type NMD_N01_NCKv2 struct {
	nck *NCK // Required
	nte []*NTE
	nst *NMD_N01_NCK_NSTv2
	nsc *NMD_N01_NCK_NSCv2
}

func (m *NMD_N01_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMD_N01_NCKv2) AllNTE() []*NTE {
	return m.nte
}

func (m *NMD_N01_NCKv2) GroupByNST() *NMD_N01_NCK_NSTv2 {
	return m.nst
}

func (m *NMD_N01_NCKv2) GroupByNSC() *NMD_N01_NCK_NSCv2 {
	return m.nsc
}


func (m NMD_N01_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck": m.nck,
		"nte": m.nte,
		"nst": m.nst,
		"nsc": m.nsc,
	}, nil
}
type NMD_N01_NCK_NSTv2 struct {
	nst *NST // Required
	nte []*NTE
}

func (m *NMD_N01_NCK_NSTv2) NST() *NST {
	return m.nst
}

func (m *NMD_N01_NCK_NSTv2) AllNTE() []*NTE {
	return m.nte
}


func (m NMD_N01_NCK_NSTv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nst": m.nst,
		"nte": m.nte,
	}, nil
}
type NMD_N01_NCK_NSCv2 struct {
	nsc *NSC // Required
	nte []*NTE
}

func (m *NMD_N01_NCK_NSCv2) NSC() *NSC {
	return m.nsc
}

func (m *NMD_N01_NCK_NSCv2) AllNTE() []*NTE {
	return m.nte
}


func (m NMD_N01_NCK_NSCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nsc": m.nsc,
		"nte": m.nte,
	}, nil
}
type NMQ_N02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	nck []*NMQ_N02_NCKv2 // Required
}

func (m *NMQ_N02v2) MSH() *MSH {
	return m.msh
}

func (m *NMQ_N02v2) QRD() *QRD {
	return m.qrd
}

func (m *NMQ_N02v2) QRF() *QRF {
	return m.qrf
}

func (m *NMQ_N02v2) GroupByNCK() []*NMQ_N02_NCKv2 {
	return m.nck
}


func (m NMQ_N02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"nck": m.nck,
	}, nil
}
type NMQ_N02_NCKv2 struct {
	nck *NCK
	nst *NST
	nsc *NSC
}

func (m *NMQ_N02_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMQ_N02_NCKv2) NST() *NST {
	return m.nst
}

func (m *NMQ_N02_NCKv2) NSC() *NSC {
	return m.nsc
}


func (m NMQ_N02_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck": m.nck,
		"nst": m.nst,
		"nsc": m.nsc,
	}, nil
}
type NMR_N02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD
	nck []*NMR_N02_NCKv2 // Required
}

func (m *NMR_N02v2) MSH() *MSH {
	return m.msh
}

func (m *NMR_N02v2) MSA() *MSA {
	return m.msa
}

func (m *NMR_N02v2) ERR() *ERR {
	return m.err
}

func (m *NMR_N02v2) QRD() *QRD {
	return m.qrd
}

func (m *NMR_N02v2) GroupByNCK() []*NMR_N02_NCKv2 {
	return m.nck
}


func (m NMR_N02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"nck": m.nck,
	}, nil
}
type NMR_N02_NCKv2 struct {
	nck *NCK
	nte1 []*NTE
	nst *NST
	nte2 []*NTE
	nsc *NSC
	nte3 []*NTE
}

func (m *NMR_N02_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMR_N02_NCKv2) AllNTE1() []*NTE {
	return m.nte1
}

func (m *NMR_N02_NCKv2) NST() *NST {
	return m.nst
}

func (m *NMR_N02_NCKv2) AllNTE2() []*NTE {
	return m.nte2
}

func (m *NMR_N02_NCKv2) NSC() *NSC {
	return m.nsc
}

func (m *NMR_N02_NCKv2) AllNTE3() []*NTE {
	return m.nte3
}


func (m NMR_N02_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck": m.nck,
		"nte1": m.nte1,
		"nst": m.nst,
		"nte2": m.nte2,
		"nsc": m.nsc,
		"nte3": m.nte3,
	}, nil
}
type OMD_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *OMD_O01_PIDv2
	orc1 []*OMD_O01_ORC1v2 // Required
	orc2 []*OMD_O01_ORC2v2
}

func (m *OMD_O01v2) MSH() *MSH {
	return m.msh
}

func (m *OMD_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMD_O01v2) GroupByPID() *OMD_O01_PIDv2 {
	return m.pid
}

func (m *OMD_O01v2) GroupByORC1() []*OMD_O01_ORC1v2 {
	return m.orc1
}

func (m *OMD_O01v2) GroupByORC2() []*OMD_O01_ORC2v2 {
	return m.orc2
}


func (m OMD_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc1": m.orc1,
		"orc2": m.orc2,
	}, nil
}
type OMD_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*OMD_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *OMD_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *OMD_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *OMD_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMD_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *OMD_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *OMD_O01_PIDv2) GroupByIN1() []*OMD_O01_PID_IN1v2 {
	return m.in1
}

func (m *OMD_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *OMD_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m OMD_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type OMD_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *OMD_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *OMD_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *OMD_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m OMD_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type OMD_O01_ORC1v2 struct {
	orc *ORC // Required
	ods []*ODS // Required
	nte []*NTE
	obx []*OMD_O01_ORC1_OBXv2 // Required
}

func (m *OMD_O01_ORC1v2) ORC() *ORC {
	return m.orc
}

func (m *OMD_O01_ORC1v2) AllODS() []*ODS {
	return m.ods
}

func (m *OMD_O01_ORC1v2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMD_O01_ORC1v2) GroupByOBX() []*OMD_O01_ORC1_OBXv2 {
	return m.obx
}


func (m OMD_O01_ORC1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"ods": m.ods,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type OMD_O01_ORC1_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *OMD_O01_ORC1_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *OMD_O01_ORC1_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m OMD_O01_ORC1_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type OMD_O01_ORC2v2 struct {
	orc *ORC // Required
	odt []*ODT // Required
	nte []*NTE
}

func (m *OMD_O01_ORC2v2) ORC() *ORC {
	return m.orc
}

func (m *OMD_O01_ORC2v2) AllODT() []*ODT {
	return m.odt
}

func (m *OMD_O01_ORC2v2) AllNTE() []*NTE {
	return m.nte
}


func (m OMD_O01_ORC2v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"odt": m.odt,
		"nte": m.nte,
	}, nil
}
type OMN_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *OMN_O01_PIDv2
	orc []*OMN_O01_ORCv2 // Required
}

func (m *OMN_O01v2) MSH() *MSH {
	return m.msh
}

func (m *OMN_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMN_O01v2) GroupByPID() *OMN_O01_PIDv2 {
	return m.pid
}

func (m *OMN_O01v2) GroupByORC() []*OMN_O01_ORCv2 {
	return m.orc
}


func (m OMN_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type OMN_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*OMN_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *OMN_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *OMN_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *OMN_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMN_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *OMN_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *OMN_O01_PIDv2) GroupByIN1() []*OMN_O01_PID_IN1v2 {
	return m.in1
}

func (m *OMN_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *OMN_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m OMN_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type OMN_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *OMN_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *OMN_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *OMN_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m OMN_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type OMN_O01_ORCv2 struct {
	orc *ORC // Required
	rqd *RQD // Required
	rq1 *RQ1
	nte []*NTE
	obx []*OMN_O01_ORC_OBXv2
	blg *BLG
}

func (m *OMN_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *OMN_O01_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *OMN_O01_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *OMN_O01_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMN_O01_ORCv2) GroupByOBX() []*OMN_O01_ORC_OBXv2 {
	return m.obx
}

func (m *OMN_O01_ORCv2) BLG() *BLG {
	return m.blg
}


func (m OMN_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rqd": m.rqd,
		"rq1": m.rq1,
		"nte": m.nte,
		"obx": m.obx,
		"blg": m.blg,
	}, nil
}
type OMN_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *OMN_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *OMN_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m OMN_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type OMS_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *OMS_O01_PIDv2
	orc []*OMS_O01_ORCv2 // Required
}

func (m *OMS_O01v2) MSH() *MSH {
	return m.msh
}

func (m *OMS_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMS_O01v2) GroupByPID() *OMS_O01_PIDv2 {
	return m.pid
}

func (m *OMS_O01v2) GroupByORC() []*OMS_O01_ORCv2 {
	return m.orc
}


func (m OMS_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type OMS_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*OMS_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *OMS_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *OMS_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *OMS_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMS_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *OMS_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *OMS_O01_PIDv2) GroupByIN1() []*OMS_O01_PID_IN1v2 {
	return m.in1
}

func (m *OMS_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *OMS_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m OMS_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type OMS_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *OMS_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *OMS_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *OMS_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m OMS_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type OMS_O01_ORCv2 struct {
	orc *ORC // Required
	rqd *RQD // Required
	nte []*NTE
	obx []*OMS_O01_ORC_OBXv2
	blg *BLG
}

func (m *OMS_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *OMS_O01_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *OMS_O01_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OMS_O01_ORCv2) GroupByOBX() []*OMS_O01_ORC_OBXv2 {
	return m.obx
}

func (m *OMS_O01_ORCv2) BLG() *BLG {
	return m.blg
}


func (m OMS_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rqd": m.rqd,
		"nte": m.nte,
		"obx": m.obx,
		"blg": m.blg,
	}, nil
}
type OMS_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *OMS_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *OMS_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m OMS_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ORD_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *ORD_O02_PIDv2
}

func (m *ORD_O02v2) MSH() *MSH {
	return m.msh
}

func (m *ORD_O02v2) MSA() *MSA {
	return m.msa
}

func (m *ORD_O02v2) ERR() *ERR {
	return m.err
}

func (m *ORD_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORD_O02v2) GroupByPID() *ORD_O02_PIDv2 {
	return m.pid
}


func (m ORD_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type ORD_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc1 []*ORD_O02_PID_ORC1v2 // Required
	orc2 []*ORD_O02_PID_ORC2v2
}

func (m *ORD_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORD_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORD_O02_PIDv2) GroupByORC1() []*ORD_O02_PID_ORC1v2 {
	return m.orc1
}

func (m *ORD_O02_PIDv2) GroupByORC2() []*ORD_O02_PID_ORC2v2 {
	return m.orc2
}


func (m ORD_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc1": m.orc1,
		"orc2": m.orc2,
	}, nil
}
type ORD_O02_PID_ORC1v2 struct {
	orc *ORC // Required
	ods []*ODS
	nte []*NTE
}

func (m *ORD_O02_PID_ORC1v2) ORC() *ORC {
	return m.orc
}

func (m *ORD_O02_PID_ORC1v2) AllODS() []*ODS {
	return m.ods
}

func (m *ORD_O02_PID_ORC1v2) AllNTE() []*NTE {
	return m.nte
}


func (m ORD_O02_PID_ORC1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"ods": m.ods,
		"nte": m.nte,
	}, nil
}
type ORD_O02_PID_ORC2v2 struct {
	orc *ORC // Required
	odt []*ODT
	nte []*NTE
}

func (m *ORD_O02_PID_ORC2v2) ORC() *ORC {
	return m.orc
}

func (m *ORD_O02_PID_ORC2v2) AllODT() []*ODT {
	return m.odt
}

func (m *ORD_O02_PID_ORC2v2) AllNTE() []*NTE {
	return m.nte
}


func (m ORD_O02_PID_ORC2v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"odt": m.odt,
		"nte": m.nte,
	}, nil
}
type ORF_R04v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	pid []*ORF_R04_PIDv2 // Required
	dsc *DSC
}

func (m *ORF_R04v2) MSH() *MSH {
	return m.msh
}

func (m *ORF_R04v2) MSA() *MSA {
	return m.msa
}

func (m *ORF_R04v2) QRD() *QRD {
	return m.qrd
}

func (m *ORF_R04v2) QRF() *QRF {
	return m.qrf
}

func (m *ORF_R04v2) GroupByPID() []*ORF_R04_PIDv2 {
	return m.pid
}

func (m *ORF_R04v2) DSC() *DSC {
	return m.dsc
}


func (m ORF_R04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ORF_R04_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	obr []*ORF_R04_PID_OBRv2 // Required
}

func (m *ORF_R04_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORF_R04_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORF_R04_PIDv2) GroupByOBR() []*ORF_R04_PID_OBRv2 {
	return m.obr
}


func (m ORF_R04_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"obr": m.obr,
	}, nil
}
type ORF_R04_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORF_R04_PID_OBR_OBXv2 // Required
	cti []*CTI
}

func (m *ORF_R04_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORF_R04_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORF_R04_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORF_R04_PID_OBRv2) GroupByOBX() []*ORF_R04_PID_OBR_OBXv2 {
	return m.obx
}

func (m *ORF_R04_PID_OBRv2) AllCTI() []*CTI {
	return m.cti
}


func (m ORF_R04_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}
type ORF_R04_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORF_R04_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORF_R04_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORF_R04_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ORM_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *ORM_O01_PIDv2
	orc []*ORM_O01_ORCv2 // Required
}

func (m *ORM_O01v2) MSH() *MSH {
	return m.msh
}

func (m *ORM_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORM_O01v2) GroupByPID() *ORM_O01_PIDv2 {
	return m.pid
}

func (m *ORM_O01v2) GroupByORC() []*ORM_O01_ORCv2 {
	return m.orc
}


func (m ORM_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type ORM_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*ORM_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *ORM_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORM_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ORM_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORM_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORM_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ORM_O01_PIDv2) GroupByIN1() []*ORM_O01_PID_IN1v2 {
	return m.in1
}

func (m *ORM_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *ORM_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m ORM_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type ORM_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ORM_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ORM_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ORM_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m ORM_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type ORM_O01_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	rqd *RQD // Required
	rq1 *RQ1 // Required
	rxo *RXO // Required
	ods *ODS // Required
	odt *ODT // Required
	nte []*NTE
	dg1 []*DG1
	obx []*ORM_O01_ORC_OBXv2
	cti *CTI
	blg *BLG
}

func (m *ORM_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORM_O01_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORM_O01_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *ORM_O01_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *ORM_O01_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *ORM_O01_ORCv2) ODS() *ODS {
	return m.ods
}

func (m *ORM_O01_ORCv2) ODT() *ODT {
	return m.odt
}

func (m *ORM_O01_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORM_O01_ORCv2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ORM_O01_ORCv2) GroupByOBX() []*ORM_O01_ORC_OBXv2 {
	return m.obx
}

func (m *ORM_O01_ORCv2) CTI() *CTI {
	return m.cti
}

func (m *ORM_O01_ORCv2) BLG() *BLG {
	return m.blg
}


func (m ORM_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"rqd": m.rqd,
		"rq1": m.rq1,
		"rxo": m.rxo,
		"ods": m.ods,
		"odt": m.odt,
		"nte": m.nte,
		"dg1": m.dg1,
		"obx": m.obx,
		"cti": m.cti,
		"blg": m.blg,
	}, nil
}
type ORM_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *ORM_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORM_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORM_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ORN_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *ORN_O02_PIDv2
}

func (m *ORN_O02v2) MSH() *MSH {
	return m.msh
}

func (m *ORN_O02v2) MSA() *MSA {
	return m.msa
}

func (m *ORN_O02v2) ERR() *ERR {
	return m.err
}

func (m *ORN_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORN_O02v2) GroupByPID() *ORN_O02_PIDv2 {
	return m.pid
}


func (m ORN_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type ORN_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*ORN_O02_PID_ORCv2 // Required
}

func (m *ORN_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORN_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORN_O02_PIDv2) GroupByORC() []*ORN_O02_PID_ORCv2 {
	return m.orc
}


func (m ORN_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type ORN_O02_PID_ORCv2 struct {
	orc *ORC // Required
	rqd *RQD // Required
	rq1 *RQ1
	nte []*NTE
}

func (m *ORN_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORN_O02_PID_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *ORN_O02_PID_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *ORN_O02_PID_ORCv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORN_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rqd": m.rqd,
		"rq1": m.rq1,
		"nte": m.nte,
	}, nil
}
type ORR_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *ORR_O02_PIDv2
}

func (m *ORR_O02v2) MSH() *MSH {
	return m.msh
}

func (m *ORR_O02v2) MSA() *MSA {
	return m.msa
}

func (m *ORR_O02v2) ERR() *ERR {
	return m.err
}

func (m *ORR_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORR_O02v2) GroupByPID() *ORR_O02_PIDv2 {
	return m.pid
}


func (m ORR_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type ORR_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*ORR_O02_PID_ORCv2 // Required
}

func (m *ORR_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORR_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORR_O02_PIDv2) GroupByORC() []*ORR_O02_PID_ORCv2 {
	return m.orc
}


func (m ORR_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type ORR_O02_PID_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	rqd *RQD // Required
	rq1 *RQ1 // Required
	rxo *RXO // Required
	ods *ODS // Required
	odt *ODT // Required
	nte []*NTE
	cti []*CTI
}

func (m *ORR_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORR_O02_PID_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORR_O02_PID_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *ORR_O02_PID_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *ORR_O02_PID_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *ORR_O02_PID_ORCv2) ODS() *ODS {
	return m.ods
}

func (m *ORR_O02_PID_ORCv2) ODT() *ODT {
	return m.odt
}

func (m *ORR_O02_PID_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORR_O02_PID_ORCv2) AllCTI() []*CTI {
	return m.cti
}


func (m ORR_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"rqd": m.rqd,
		"rq1": m.rq1,
		"rxo": m.rxo,
		"ods": m.ods,
		"odt": m.odt,
		"nte": m.nte,
		"cti": m.cti,
	}, nil
}
type ORU_R01v2 struct {
	msh *MSH // Required
	pid []*ORU_R01_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R01v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R01v2) GroupByPID() []*ORU_R01_PIDv2 {
	return m.pid
}

func (m *ORU_R01v2) DSC() *DSC {
	return m.dsc
}


func (m ORU_R01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ORU_R01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	obr []*ORU_R01_PID_OBRv2 // Required
}

func (m *ORU_R01_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ORU_R01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ORU_R01_PIDv2) GroupByOBR() []*ORU_R01_PID_OBRv2 {
	return m.obr
}


func (m ORU_R01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obr": m.obr,
	}, nil
}
type ORU_R01_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R01_PID_OBR_OBXv2 // Required
	cti []*CTI
}

func (m *ORU_R01_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R01_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R01_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R01_PID_OBRv2) GroupByOBX() []*ORU_R01_PID_OBR_OBXv2 {
	return m.obx
}

func (m *ORU_R01_PID_OBRv2) AllCTI() []*CTI {
	return m.cti
}


func (m ORU_R01_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}
type ORU_R01_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R01_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R01_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORU_R01_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ORU_R03v2 struct {
	msh *MSH // Required
	pid []*ORU_R03_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R03v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R03v2) GroupByPID() []*ORU_R03_PIDv2 {
	return m.pid
}

func (m *ORU_R03v2) DSC() *DSC {
	return m.dsc
}


func (m ORU_R03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ORU_R03_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	pv1 *PV1
	obr []*ORU_R03_PID_OBRv2 // Required
}

func (m *ORU_R03_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R03_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R03_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R03_PIDv2) GroupByOBR() []*ORU_R03_PID_OBRv2 {
	return m.obr
}


func (m ORU_R03_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"pv1": m.pv1,
		"obr": m.obr,
	}, nil
}
type ORU_R03_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R03_PID_OBR_OBXv2 // Required
}

func (m *ORU_R03_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R03_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R03_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R03_PID_OBRv2) GroupByOBX() []*ORU_R03_PID_OBR_OBXv2 {
	return m.obx
}


func (m ORU_R03_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type ORU_R03_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R03_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R03_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORU_R03_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ORU_R32v2 struct {
	msh *MSH // Required
	pid []*ORU_R32_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R32v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R32v2) GroupByPID() []*ORU_R32_PIDv2 {
	return m.pid
}

func (m *ORU_R32v2) DSC() *DSC {
	return m.dsc
}


func (m ORU_R32v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type ORU_R32_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	obr []*ORU_R32_PID_OBRv2 // Required
}

func (m *ORU_R32_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R32_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ORU_R32_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ORU_R32_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R32_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R32_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ORU_R32_PIDv2) GroupByOBR() []*ORU_R32_PID_OBRv2 {
	return m.obr
}


func (m ORU_R32_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obr": m.obr,
	}, nil
}
type ORU_R32_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R32_PID_OBR_OBXv2 // Required
	cti []*CTI
}

func (m *ORU_R32_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R32_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R32_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R32_PID_OBRv2) GroupByOBX() []*ORU_R32_PID_OBR_OBXv2 {
	return m.obx
}

func (m *ORU_R32_PID_OBRv2) AllCTI() []*CTI {
	return m.cti
}


func (m ORU_R32_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}
type ORU_R32_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R32_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R32_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m ORU_R32_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type OSQ_Q06v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *OSQ_Q06v2) MSH() *MSH {
	return m.msh
}

func (m *OSQ_Q06v2) QRD() *QRD {
	return m.qrd
}

func (m *OSQ_Q06v2) QRF() *QRF {
	return m.qrf
}

func (m *OSQ_Q06v2) DSC() *DSC {
	return m.dsc
}


func (m OSQ_Q06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type OSR_Q06v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	qrd *QRD // Required
	qrf *QRF
	pid *OSR_Q06_PIDv2
	dsc *DSC
}

func (m *OSR_Q06v2) MSH() *MSH {
	return m.msh
}

func (m *OSR_Q06v2) MSA() *MSA {
	return m.msa
}

func (m *OSR_Q06v2) ERR() *ERR {
	return m.err
}

func (m *OSR_Q06v2) AllNTE() []*NTE {
	return m.nte
}

func (m *OSR_Q06v2) QRD() *QRD {
	return m.qrd
}

func (m *OSR_Q06v2) QRF() *QRF {
	return m.qrf
}

func (m *OSR_Q06v2) GroupByPID() *OSR_Q06_PIDv2 {
	return m.pid
}

func (m *OSR_Q06v2) DSC() *DSC {
	return m.dsc
}


func (m OSR_Q06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}
type OSR_Q06_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*OSR_Q06_PID_ORCv2 // Required
}

func (m *OSR_Q06_PIDv2) PID() *PID {
	return m.pid
}

func (m *OSR_Q06_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OSR_Q06_PIDv2) GroupByORC() []*OSR_Q06_PID_ORCv2 {
	return m.orc
}


func (m OSR_Q06_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type OSR_Q06_PID_ORCv2 struct {
	orc *ORC // Required
	obr *OBR
	nte []*NTE
	cti []*CTI
}

func (m *OSR_Q06_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *OSR_Q06_PID_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *OSR_Q06_PID_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *OSR_Q06_PID_ORCv2) AllCTI() []*CTI {
	return m.cti
}


func (m OSR_Q06_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"cti": m.cti,
	}, nil
}
type PEX_P07v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	pes []*PEX_P07_PESv2 // Required
}

func (m *PEX_P07v2) MSH() *MSH {
	return m.msh
}

func (m *PEX_P07v2) EVN() *EVN {
	return m.evn
}

func (m *PEX_P07v2) PID() *PID {
	return m.pid
}

func (m *PEX_P07v2) PD1() *PD1 {
	return m.pd1
}

func (m *PEX_P07v2) AllNTE() []*NTE {
	return m.nte
}

func (m *PEX_P07v2) PV1() *PV1 {
	return m.pv1
}

func (m *PEX_P07v2) PV2() *PV2 {
	return m.pv2
}

func (m *PEX_P07v2) GroupByPES() []*PEX_P07_PESv2 {
	return m.pes
}


func (m PEX_P07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"pes": m.pes,
	}, nil
}
type PEX_P07_PESv2 struct {
	pes *PES // Required
	peo []*PEX_P07_PES_PEOv2 // Required
}

func (m *PEX_P07_PESv2) PES() *PES {
	return m.pes
}

func (m *PEX_P07_PESv2) GroupByPEO() []*PEX_P07_PES_PEOv2 {
	return m.peo
}


func (m PEX_P07_PESv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pes": m.pes,
		"peo": m.peo,
	}, nil
}
type PEX_P07_PES_PEOv2 struct {
	peo *PEO // Required
	pcr []*PEX_P07_PES_PEO_PCRv2 // Required
}

func (m *PEX_P07_PES_PEOv2) PEO() *PEO {
	return m.peo
}

func (m *PEX_P07_PES_PEOv2) GroupByPCR() []*PEX_P07_PES_PEO_PCRv2 {
	return m.pcr
}


func (m PEX_P07_PES_PEOv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"peo": m.peo,
		"pcr": m.pcr,
	}, nil
}
type PEX_P07_PES_PEO_PCRv2 struct {
	pcr *PCR // Required
	rxe *RXE // Required
	rxr []*RXR
	rxa []*PEX_P07_PES_PEO_PCR_RXAv2
	prb []*PRB
	obx []*OBX
	nte []*NTE
	nk1 *PEX_P07_PES_PEO_PCR_NK1v2
	csr []*PEX_P07_PES_PEO_PCR_CSRv2
}

func (m *PEX_P07_PES_PEO_PCRv2) PCR() *PCR {
	return m.pcr
}

func (m *PEX_P07_PES_PEO_PCRv2) RXE() *RXE {
	return m.rxe
}

func (m *PEX_P07_PES_PEO_PCRv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *PEX_P07_PES_PEO_PCRv2) GroupByRXA() []*PEX_P07_PES_PEO_PCR_RXAv2 {
	return m.rxa
}

func (m *PEX_P07_PES_PEO_PCRv2) AllPRB() []*PRB {
	return m.prb
}

func (m *PEX_P07_PES_PEO_PCRv2) AllOBX() []*OBX {
	return m.obx
}

func (m *PEX_P07_PES_PEO_PCRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PEX_P07_PES_PEO_PCRv2) GroupByNK1() *PEX_P07_PES_PEO_PCR_NK1v2 {
	return m.nk1
}

func (m *PEX_P07_PES_PEO_PCRv2) GroupByCSR() []*PEX_P07_PES_PEO_PCR_CSRv2 {
	return m.csr
}


func (m PEX_P07_PES_PEO_PCRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pcr": m.pcr,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxa": m.rxa,
		"prb": m.prb,
		"obx": m.obx,
		"nte": m.nte,
		"nk1": m.nk1,
		"csr": m.csr,
	}, nil
}
type PEX_P07_PES_PEO_PCR_RXAv2 struct {
	rxa *RXA // Required
	rxr *RXR
}

func (m *PEX_P07_PES_PEO_PCR_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *PEX_P07_PES_PEO_PCR_RXAv2) RXR() *RXR {
	return m.rxr
}


func (m PEX_P07_PES_PEO_PCR_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxa": m.rxa,
		"rxr": m.rxr,
	}, nil
}
type PEX_P07_PES_PEO_PCR_NK1v2 struct {
	nk1 *NK1 // Required
	rxe *RXE // Required
	rxr []*RXR
	rxa []*PEX_P07_PES_PEO_PCR_NK1_RXAv2
	prb []*PRB
	obx []*OBX
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) NK1() *NK1 {
	return m.nk1
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) RXE() *RXE {
	return m.rxe
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) AllRXR() []*RXR {
	return m.rxr
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) GroupByRXA() []*PEX_P07_PES_PEO_PCR_NK1_RXAv2 {
	return m.rxa
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) AllPRB() []*PRB {
	return m.prb
}

func (m *PEX_P07_PES_PEO_PCR_NK1v2) AllOBX() []*OBX {
	return m.obx
}


func (m PEX_P07_PES_PEO_PCR_NK1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nk1": m.nk1,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxa": m.rxa,
		"prb": m.prb,
		"obx": m.obx,
	}, nil
}
type PEX_P07_PES_PEO_PCR_NK1_RXAv2 struct {
	rxa *RXA // Required
	rxr *RXR
}

func (m *PEX_P07_PES_PEO_PCR_NK1_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *PEX_P07_PES_PEO_PCR_NK1_RXAv2) RXR() *RXR {
	return m.rxr
}


func (m PEX_P07_PES_PEO_PCR_NK1_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxa": m.rxa,
		"rxr": m.rxr,
	}, nil
}
type PEX_P07_PES_PEO_PCR_CSRv2 struct {
	csr *CSR // Required
	csp []*CSP
}

func (m *PEX_P07_PES_PEO_PCR_CSRv2) CSR() *CSR {
	return m.csr
}

func (m *PEX_P07_PES_PEO_PCR_CSRv2) AllCSP() []*CSP {
	return m.csp
}


func (m PEX_P07_PES_PEO_PCR_CSRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"csr": m.csr,
		"csp": m.csp,
	}, nil
}
type PGL_PC6v2 struct {
	msh *MSH // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	gol []*PGL_PC6_GOLv2 // Required
}

func (m *PGL_PC6v2) MSH() *MSH {
	return m.msh
}

func (m *PGL_PC6v2) PID() *PID {
	return m.pid
}

func (m *PGL_PC6v2) PV1() *PV1 {
	return m.pv1
}

func (m *PGL_PC6v2) PV2() *PV2 {
	return m.pv2
}

func (m *PGL_PC6v2) GroupByGOL() []*PGL_PC6_GOLv2 {
	return m.gol
}


func (m PGL_PC6v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"gol": m.gol,
	}, nil
}
type PGL_PC6_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PGL_PC6_GOL_ROLv2
	pth []*PGL_PC6_GOL_PTHv2
	obx []*PGL_PC6_GOL_OBXv2
	prb []*PGL_PC6_GOL_PRBv2
	orc []*PGL_PC6_GOL_ORCv2
}

func (m *PGL_PC6_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PGL_PC6_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PGL_PC6_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PGL_PC6_GOLv2) GroupByROL() []*PGL_PC6_GOL_ROLv2 {
	return m.rol
}

func (m *PGL_PC6_GOLv2) GroupByPTH() []*PGL_PC6_GOL_PTHv2 {
	return m.pth
}

func (m *PGL_PC6_GOLv2) GroupByOBX() []*PGL_PC6_GOL_OBXv2 {
	return m.obx
}

func (m *PGL_PC6_GOLv2) GroupByPRB() []*PGL_PC6_GOL_PRBv2 {
	return m.prb
}

func (m *PGL_PC6_GOLv2) GroupByORC() []*PGL_PC6_GOL_ORCv2 {
	return m.orc
}


func (m PGL_PC6_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"pth": m.pth,
		"obx": m.obx,
		"prb": m.prb,
		"orc": m.orc,
	}, nil
}
type PGL_PC6_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PGL_PC6_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PGL_PC6_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PGL_PC6_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PGL_PC6_GOL_PTHv2 struct {
	pth *PTH // Required
	var_ []*VAR
}

func (m *PGL_PC6_GOL_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PGL_PC6_GOL_PTHv2) AllVAR() []*VAR {
	return m.var_
}


func (m PGL_PC6_GOL_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"var_": m.var_,
	}, nil
}
type PGL_PC6_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PGL_PC6_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PGL_PC6_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PGL_PC6_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PGL_PC6_GOL_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PGL_PC6_GOL_PRB_ROLv2
	obx []*PGL_PC6_GOL_PRB_OBXv2
}

func (m *PGL_PC6_GOL_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PGL_PC6_GOL_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PGL_PC6_GOL_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PGL_PC6_GOL_PRBv2) GroupByROL() []*PGL_PC6_GOL_PRB_ROLv2 {
	return m.rol
}

func (m *PGL_PC6_GOL_PRBv2) GroupByOBX() []*PGL_PC6_GOL_PRB_OBXv2 {
	return m.obx
}


func (m PGL_PC6_GOL_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PGL_PC6_GOL_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PGL_PC6_GOL_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PGL_PC6_GOL_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PGL_PC6_GOL_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PGL_PC6_GOL_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PGL_PC6_GOL_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PGL_PC6_GOL_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PGL_PC6_GOL_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PGL_PC6_GOL_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PGL_PC6_GOL_ORC_OBXv2
}

func (m *PGL_PC6_GOL_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PGL_PC6_GOL_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PGL_PC6_GOL_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PGL_PC6_GOL_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PGL_PC6_GOL_ORCv2) GroupByOBX() []*PGL_PC6_GOL_ORC_OBXv2 {
	return m.obx
}


func (m PGL_PC6_GOL_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PGL_PC6_GOL_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PGL_PC6_GOL_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PGL_PC6_GOL_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PGL_PC6_GOL_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PGL_PC6_GOL_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PIN_I07v2 struct {
	msh *MSH // Required
	prd []*PIN_I07_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*PIN_I07_IN1v2 // Required
	nte []*NTE
}

func (m *PIN_I07v2) MSH() *MSH {
	return m.msh
}

func (m *PIN_I07v2) GroupByPRD() []*PIN_I07_PRDv2 {
	return m.prd
}

func (m *PIN_I07v2) PID() *PID {
	return m.pid
}

func (m *PIN_I07v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *PIN_I07v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *PIN_I07v2) GroupByIN1() []*PIN_I07_IN1v2 {
	return m.in1
}

func (m *PIN_I07v2) AllNTE() []*NTE {
	return m.nte
}


func (m PIN_I07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"nte": m.nte,
	}, nil
}
type PIN_I07_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *PIN_I07_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *PIN_I07_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m PIN_I07_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type PIN_I07_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *PIN_I07_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *PIN_I07_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *PIN_I07_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m PIN_I07_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type PPG_PCGv2 struct {
	msh *MSH // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	pth []*PPG_PCG_PTHv2 // Required
}

func (m *PPG_PCGv2) MSH() *MSH {
	return m.msh
}

func (m *PPG_PCGv2) PID() *PID {
	return m.pid
}

func (m *PPG_PCGv2) PV1() *PV1 {
	return m.pv1
}

func (m *PPG_PCGv2) PV2() *PV2 {
	return m.pv2
}

func (m *PPG_PCGv2) GroupByPTH() []*PPG_PCG_PTHv2 {
	return m.pth
}


func (m PPG_PCGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"pth": m.pth,
	}, nil
}
type PPG_PCG_PTHv2 struct {
	pth *PTH // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPG_PCG_PTH_ROLv2
	gol []*PPG_PCG_PTH_GOLv2
}

func (m *PPG_PCG_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PPG_PCG_PTHv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPG_PCG_PTHv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPG_PCG_PTHv2) GroupByROL() []*PPG_PCG_PTH_ROLv2 {
	return m.rol
}

func (m *PPG_PCG_PTHv2) GroupByGOL() []*PPG_PCG_PTH_GOLv2 {
	return m.gol
}


func (m PPG_PCG_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"gol": m.gol,
	}, nil
}
type PPG_PCG_PTH_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPG_PCG_PTH_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPG_PCG_PTH_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPG_PCG_PTH_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPG_PCG_PTH_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPG_PCG_PTH_GOL_ROLv2
	obx []*PPG_PCG_PTH_GOL_OBXv2
	prb []*PPG_PCG_PTH_GOL_PRBv2
	orc []*PPG_PCG_PTH_GOL_ORCv2
}

func (m *PPG_PCG_PTH_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PPG_PCG_PTH_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPG_PCG_PTH_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPG_PCG_PTH_GOLv2) GroupByROL() []*PPG_PCG_PTH_GOL_ROLv2 {
	return m.rol
}

func (m *PPG_PCG_PTH_GOLv2) GroupByOBX() []*PPG_PCG_PTH_GOL_OBXv2 {
	return m.obx
}

func (m *PPG_PCG_PTH_GOLv2) GroupByPRB() []*PPG_PCG_PTH_GOL_PRBv2 {
	return m.prb
}

func (m *PPG_PCG_PTH_GOLv2) GroupByORC() []*PPG_PCG_PTH_GOL_ORCv2 {
	return m.orc
}


func (m PPG_PCG_PTH_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
		"prb": m.prb,
		"orc": m.orc,
	}, nil
}
type PPG_PCG_PTH_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPG_PCG_PTH_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPG_PCG_PTH_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPG_PCG_PTH_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPG_PCG_PTH_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPG_PCG_PTH_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPG_PCG_PTH_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPG_PCG_PTH_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPG_PCG_PTH_GOL_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPG_PCG_PTH_GOL_PRB_ROLv2
	obx []*PPG_PCG_PTH_GOL_PRB_OBXv2
}

func (m *PPG_PCG_PTH_GOL_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PPG_PCG_PTH_GOL_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPG_PCG_PTH_GOL_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPG_PCG_PTH_GOL_PRBv2) GroupByROL() []*PPG_PCG_PTH_GOL_PRB_ROLv2 {
	return m.rol
}

func (m *PPG_PCG_PTH_GOL_PRBv2) GroupByOBX() []*PPG_PCG_PTH_GOL_PRB_OBXv2 {
	return m.obx
}


func (m PPG_PCG_PTH_GOL_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PPG_PCG_PTH_GOL_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPG_PCG_PTH_GOL_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPG_PCG_PTH_GOL_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPG_PCG_PTH_GOL_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPG_PCG_PTH_GOL_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPG_PCG_PTH_GOL_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPG_PCG_PTH_GOL_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPG_PCG_PTH_GOL_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPG_PCG_PTH_GOL_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PPG_PCG_PTH_GOL_ORC_OBXv2
}

func (m *PPG_PCG_PTH_GOL_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PPG_PCG_PTH_GOL_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PPG_PCG_PTH_GOL_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPG_PCG_PTH_GOL_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPG_PCG_PTH_GOL_ORCv2) GroupByOBX() []*PPG_PCG_PTH_GOL_ORC_OBXv2 {
	return m.obx
}


func (m PPG_PCG_PTH_GOL_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PPG_PCG_PTH_GOL_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PPG_PCG_PTH_GOL_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPG_PCG_PTH_GOL_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPG_PCG_PTH_GOL_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPG_PCG_PTH_GOL_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PPP_PCBv2 struct {
	msh *MSH // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	pth []*PPP_PCB_PTHv2 // Required
}

func (m *PPP_PCBv2) MSH() *MSH {
	return m.msh
}

func (m *PPP_PCBv2) PID() *PID {
	return m.pid
}

func (m *PPP_PCBv2) PV1() *PV1 {
	return m.pv1
}

func (m *PPP_PCBv2) PV2() *PV2 {
	return m.pv2
}

func (m *PPP_PCBv2) GroupByPTH() []*PPP_PCB_PTHv2 {
	return m.pth
}


func (m PPP_PCBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"pth": m.pth,
	}, nil
}
type PPP_PCB_PTHv2 struct {
	pth *PTH // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPP_PCB_PTH_ROLv2
	prb []*PPP_PCB_PTH_PRBv2
}

func (m *PPP_PCB_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PPP_PCB_PTHv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPP_PCB_PTHv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPP_PCB_PTHv2) GroupByROL() []*PPP_PCB_PTH_ROLv2 {
	return m.rol
}

func (m *PPP_PCB_PTHv2) GroupByPRB() []*PPP_PCB_PTH_PRBv2 {
	return m.prb
}


func (m PPP_PCB_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"prb": m.prb,
	}, nil
}
type PPP_PCB_PTH_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPP_PCB_PTH_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPP_PCB_PTH_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPP_PCB_PTH_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPP_PCB_PTH_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPP_PCB_PTH_PRB_ROLv2
	obx []*PPP_PCB_PTH_PRB_OBXv2
	gol []*PPP_PCB_PTH_PRB_GOLv2
	orc []*PPP_PCB_PTH_PRB_ORCv2
}

func (m *PPP_PCB_PTH_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PPP_PCB_PTH_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPP_PCB_PTH_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPP_PCB_PTH_PRBv2) GroupByROL() []*PPP_PCB_PTH_PRB_ROLv2 {
	return m.rol
}

func (m *PPP_PCB_PTH_PRBv2) GroupByOBX() []*PPP_PCB_PTH_PRB_OBXv2 {
	return m.obx
}

func (m *PPP_PCB_PTH_PRBv2) GroupByGOL() []*PPP_PCB_PTH_PRB_GOLv2 {
	return m.gol
}

func (m *PPP_PCB_PTH_PRBv2) GroupByORC() []*PPP_PCB_PTH_PRB_ORCv2 {
	return m.orc
}


func (m PPP_PCB_PTH_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
		"gol": m.gol,
		"orc": m.orc,
	}, nil
}
type PPP_PCB_PTH_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPP_PCB_PTH_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPP_PCB_PTH_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPP_PCB_PTH_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPP_PCB_PTH_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPP_PCB_PTH_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPP_PCB_PTH_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPP_PCB_PTH_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPP_PCB_PTH_PRB_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPP_PCB_PTH_PRB_GOL_ROLv2
	obx []*PPP_PCB_PTH_PRB_GOL_OBXv2
}

func (m *PPP_PCB_PTH_PRB_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PPP_PCB_PTH_PRB_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPP_PCB_PTH_PRB_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPP_PCB_PTH_PRB_GOLv2) GroupByROL() []*PPP_PCB_PTH_PRB_GOL_ROLv2 {
	return m.rol
}

func (m *PPP_PCB_PTH_PRB_GOLv2) GroupByOBX() []*PPP_PCB_PTH_PRB_GOL_OBXv2 {
	return m.obx
}


func (m PPP_PCB_PTH_PRB_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PPP_PCB_PTH_PRB_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPP_PCB_PTH_PRB_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPP_PCB_PTH_PRB_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPP_PCB_PTH_PRB_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPP_PCB_PTH_PRB_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPP_PCB_PTH_PRB_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPP_PCB_PTH_PRB_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPP_PCB_PTH_PRB_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPP_PCB_PTH_PRB_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PPP_PCB_PTH_PRB_ORC_OBXv2
}

func (m *PPP_PCB_PTH_PRB_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PPP_PCB_PTH_PRB_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PPP_PCB_PTH_PRB_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPP_PCB_PTH_PRB_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPP_PCB_PTH_PRB_ORCv2) GroupByOBX() []*PPP_PCB_PTH_PRB_ORC_OBXv2 {
	return m.obx
}


func (m PPP_PCB_PTH_PRB_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PPP_PCB_PTH_PRB_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PPP_PCB_PTH_PRB_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPP_PCB_PTH_PRB_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPP_PCB_PTH_PRB_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPP_PCB_PTH_PRB_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PPR_PC1v2 struct {
	msh *MSH // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	prb []*PPR_PC1_PRBv2 // Required
}

func (m *PPR_PC1v2) MSH() *MSH {
	return m.msh
}

func (m *PPR_PC1v2) PID() *PID {
	return m.pid
}

func (m *PPR_PC1v2) PV1() *PV1 {
	return m.pv1
}

func (m *PPR_PC1v2) PV2() *PV2 {
	return m.pv2
}

func (m *PPR_PC1v2) GroupByPRB() []*PPR_PC1_PRBv2 {
	return m.prb
}


func (m PPR_PC1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"prb": m.prb,
	}, nil
}
type PPR_PC1_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPR_PC1_PRB_ROLv2
	pth []*PPR_PC1_PRB_PTHv2
	obx []*PPR_PC1_PRB_OBXv2
	gol []*PPR_PC1_PRB_GOLv2
	orc []*PPR_PC1_PRB_ORCv2
}

func (m *PPR_PC1_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PPR_PC1_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPR_PC1_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPR_PC1_PRBv2) GroupByROL() []*PPR_PC1_PRB_ROLv2 {
	return m.rol
}

func (m *PPR_PC1_PRBv2) GroupByPTH() []*PPR_PC1_PRB_PTHv2 {
	return m.pth
}

func (m *PPR_PC1_PRBv2) GroupByOBX() []*PPR_PC1_PRB_OBXv2 {
	return m.obx
}

func (m *PPR_PC1_PRBv2) GroupByGOL() []*PPR_PC1_PRB_GOLv2 {
	return m.gol
}

func (m *PPR_PC1_PRBv2) GroupByORC() []*PPR_PC1_PRB_ORCv2 {
	return m.orc
}


func (m PPR_PC1_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"pth": m.pth,
		"obx": m.obx,
		"gol": m.gol,
		"orc": m.orc,
	}, nil
}
type PPR_PC1_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPR_PC1_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPR_PC1_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPR_PC1_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPR_PC1_PRB_PTHv2 struct {
	pth *PTH // Required
	var_ []*VAR
}

func (m *PPR_PC1_PRB_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PPR_PC1_PRB_PTHv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPR_PC1_PRB_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"var_": m.var_,
	}, nil
}
type PPR_PC1_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPR_PC1_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPR_PC1_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPR_PC1_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPR_PC1_PRB_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPR_PC1_PRB_GOL_ROLv2
	obx []*PPR_PC1_PRB_GOL_OBXv2
}

func (m *PPR_PC1_PRB_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PPR_PC1_PRB_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPR_PC1_PRB_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPR_PC1_PRB_GOLv2) GroupByROL() []*PPR_PC1_PRB_GOL_ROLv2 {
	return m.rol
}

func (m *PPR_PC1_PRB_GOLv2) GroupByOBX() []*PPR_PC1_PRB_GOL_OBXv2 {
	return m.obx
}


func (m PPR_PC1_PRB_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PPR_PC1_PRB_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPR_PC1_PRB_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPR_PC1_PRB_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPR_PC1_PRB_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPR_PC1_PRB_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPR_PC1_PRB_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPR_PC1_PRB_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPR_PC1_PRB_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPR_PC1_PRB_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PPR_PC1_PRB_ORC_OBXv2
}

func (m *PPR_PC1_PRB_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PPR_PC1_PRB_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PPR_PC1_PRB_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPR_PC1_PRB_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPR_PC1_PRB_ORCv2) GroupByOBX() []*PPR_PC1_PRB_ORC_OBXv2 {
	return m.obx
}


func (m PPR_PC1_PRB_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PPR_PC1_PRB_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PPR_PC1_PRB_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPR_PC1_PRB_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPR_PC1_PRB_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPR_PC1_PRB_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PPT_PCLv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*PPT_PCL_PIDv2 // Required
}

func (m *PPT_PCLv2) MSH() *MSH {
	return m.msh
}

func (m *PPT_PCLv2) MSA() *MSA {
	return m.msa
}

func (m *PPT_PCLv2) ERR() *ERR {
	return m.err
}

func (m *PPT_PCLv2) QRD() *QRD {
	return m.qrd
}

func (m *PPT_PCLv2) GroupByPID() []*PPT_PCL_PIDv2 {
	return m.pid
}


func (m PPT_PCLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
	}, nil
}
type PPT_PCL_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	pth []*PPT_PCL_PID_PTHv2 // Required
}

func (m *PPT_PCL_PIDv2) PID() *PID {
	return m.pid
}

func (m *PPT_PCL_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *PPT_PCL_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *PPT_PCL_PIDv2) GroupByPTH() []*PPT_PCL_PID_PTHv2 {
	return m.pth
}


func (m PPT_PCL_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"pth": m.pth,
	}, nil
}
type PPT_PCL_PID_PTHv2 struct {
	pth *PTH // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPT_PCL_PID_PTH_ROLv2
	gol []*PPT_PCL_PID_PTH_GOLv2
}

func (m *PPT_PCL_PID_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PPT_PCL_PID_PTHv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPT_PCL_PID_PTHv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPT_PCL_PID_PTHv2) GroupByROL() []*PPT_PCL_PID_PTH_ROLv2 {
	return m.rol
}

func (m *PPT_PCL_PID_PTHv2) GroupByGOL() []*PPT_PCL_PID_PTH_GOLv2 {
	return m.gol
}


func (m PPT_PCL_PID_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"gol": m.gol,
	}, nil
}
type PPT_PCL_PID_PTH_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPT_PCL_PID_PTH_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPT_PCL_PID_PTH_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPT_PCL_PID_PTH_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPT_PCL_PID_PTH_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPT_PCL_PID_PTH_GOL_ROLv2
	obx []*PPT_PCL_PID_PTH_GOL_OBXv2
	prb []*PPT_PCL_PID_PTH_GOL_PRBv2
	orc []*PPT_PCL_PID_PTH_GOL_ORCv2
}

func (m *PPT_PCL_PID_PTH_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PPT_PCL_PID_PTH_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPT_PCL_PID_PTH_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPT_PCL_PID_PTH_GOLv2) GroupByROL() []*PPT_PCL_PID_PTH_GOL_ROLv2 {
	return m.rol
}

func (m *PPT_PCL_PID_PTH_GOLv2) GroupByOBX() []*PPT_PCL_PID_PTH_GOL_OBXv2 {
	return m.obx
}

func (m *PPT_PCL_PID_PTH_GOLv2) GroupByPRB() []*PPT_PCL_PID_PTH_GOL_PRBv2 {
	return m.prb
}

func (m *PPT_PCL_PID_PTH_GOLv2) GroupByORC() []*PPT_PCL_PID_PTH_GOL_ORCv2 {
	return m.orc
}


func (m PPT_PCL_PID_PTH_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
		"prb": m.prb,
		"orc": m.orc,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPT_PCL_PID_PTH_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPT_PCL_PID_PTH_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPT_PCL_PID_PTH_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPT_PCL_PID_PTH_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPT_PCL_PID_PTH_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPT_PCL_PID_PTH_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPT_PCL_PID_PTH_GOL_PRB_ROLv2
	obx []*PPT_PCL_PID_PTH_GOL_PRB_OBXv2
}

func (m *PPT_PCL_PID_PTH_GOL_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PPT_PCL_PID_PTH_GOL_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPT_PCL_PID_PTH_GOL_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPT_PCL_PID_PTH_GOL_PRBv2) GroupByROL() []*PPT_PCL_PID_PTH_GOL_PRB_ROLv2 {
	return m.rol
}

func (m *PPT_PCL_PID_PTH_GOL_PRBv2) GroupByOBX() []*PPT_PCL_PID_PTH_GOL_PRB_OBXv2 {
	return m.obx
}


func (m PPT_PCL_PID_PTH_GOL_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPT_PCL_PID_PTH_GOL_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPT_PCL_PID_PTH_GOL_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPT_PCL_PID_PTH_GOL_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPT_PCL_PID_PTH_GOL_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPT_PCL_PID_PTH_GOL_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPT_PCL_PID_PTH_GOL_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PPT_PCL_PID_PTH_GOL_ORC_OBXv2
}

func (m *PPT_PCL_PID_PTH_GOL_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PPT_PCL_PID_PTH_GOL_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PPT_PCL_PID_PTH_GOL_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPT_PCL_PID_PTH_GOL_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPT_PCL_PID_PTH_GOL_ORCv2) GroupByOBX() []*PPT_PCL_PID_PTH_GOL_ORC_OBXv2 {
	return m.obx
}


func (m PPT_PCL_PID_PTH_GOL_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PPT_PCL_PID_PTH_GOL_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PPT_PCL_PID_PTH_GOL_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPT_PCL_PID_PTH_GOL_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPT_PCL_PID_PTH_GOL_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPT_PCL_PID_PTH_GOL_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PPV_PCAv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*PPV_PCA_PIDv2 // Required
}

func (m *PPV_PCAv2) MSH() *MSH {
	return m.msh
}

func (m *PPV_PCAv2) MSA() *MSA {
	return m.msa
}

func (m *PPV_PCAv2) ERR() *ERR {
	return m.err
}

func (m *PPV_PCAv2) QRD() *QRD {
	return m.qrd
}

func (m *PPV_PCAv2) GroupByPID() []*PPV_PCA_PIDv2 {
	return m.pid
}


func (m PPV_PCAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
	}, nil
}
type PPV_PCA_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	gol []*PPV_PCA_PID_GOLv2 // Required
}

func (m *PPV_PCA_PIDv2) PID() *PID {
	return m.pid
}

func (m *PPV_PCA_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *PPV_PCA_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *PPV_PCA_PIDv2) GroupByGOL() []*PPV_PCA_PID_GOLv2 {
	return m.gol
}


func (m PPV_PCA_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"gol": m.gol,
	}, nil
}
type PPV_PCA_PID_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPV_PCA_PID_GOL_ROLv2
	pth []*PPV_PCA_PID_GOL_PTHv2
	obx []*PPV_PCA_PID_GOL_OBXv2
	prb []*PPV_PCA_PID_GOL_PRBv2
	orc []*PPV_PCA_PID_GOL_ORCv2
}

func (m *PPV_PCA_PID_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PPV_PCA_PID_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPV_PCA_PID_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPV_PCA_PID_GOLv2) GroupByROL() []*PPV_PCA_PID_GOL_ROLv2 {
	return m.rol
}

func (m *PPV_PCA_PID_GOLv2) GroupByPTH() []*PPV_PCA_PID_GOL_PTHv2 {
	return m.pth
}

func (m *PPV_PCA_PID_GOLv2) GroupByOBX() []*PPV_PCA_PID_GOL_OBXv2 {
	return m.obx
}

func (m *PPV_PCA_PID_GOLv2) GroupByPRB() []*PPV_PCA_PID_GOL_PRBv2 {
	return m.prb
}

func (m *PPV_PCA_PID_GOLv2) GroupByORC() []*PPV_PCA_PID_GOL_ORCv2 {
	return m.orc
}


func (m PPV_PCA_PID_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"pth": m.pth,
		"obx": m.obx,
		"prb": m.prb,
		"orc": m.orc,
	}, nil
}
type PPV_PCA_PID_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPV_PCA_PID_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPV_PCA_PID_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPV_PCA_PID_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPV_PCA_PID_GOL_PTHv2 struct {
	pth *PTH // Required
	var_ []*VAR
}

func (m *PPV_PCA_PID_GOL_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PPV_PCA_PID_GOL_PTHv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPV_PCA_PID_GOL_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"var_": m.var_,
	}, nil
}
type PPV_PCA_PID_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPV_PCA_PID_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPV_PCA_PID_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPV_PCA_PID_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPV_PCA_PID_GOL_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PPV_PCA_PID_GOL_PRB_ROLv2
	obx []*PPV_PCA_PID_GOL_PRB_OBXv2
}

func (m *PPV_PCA_PID_GOL_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PPV_PCA_PID_GOL_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPV_PCA_PID_GOL_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPV_PCA_PID_GOL_PRBv2) GroupByROL() []*PPV_PCA_PID_GOL_PRB_ROLv2 {
	return m.rol
}

func (m *PPV_PCA_PID_GOL_PRBv2) GroupByOBX() []*PPV_PCA_PID_GOL_PRB_OBXv2 {
	return m.obx
}


func (m PPV_PCA_PID_GOL_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PPV_PCA_PID_GOL_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PPV_PCA_PID_GOL_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PPV_PCA_PID_GOL_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPV_PCA_PID_GOL_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PPV_PCA_PID_GOL_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PPV_PCA_PID_GOL_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPV_PCA_PID_GOL_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PPV_PCA_PID_GOL_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PPV_PCA_PID_GOL_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PPV_PCA_PID_GOL_ORC_OBXv2
}

func (m *PPV_PCA_PID_GOL_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PPV_PCA_PID_GOL_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PPV_PCA_PID_GOL_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPV_PCA_PID_GOL_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PPV_PCA_PID_GOL_ORCv2) GroupByOBX() []*PPV_PCA_PID_GOL_ORC_OBXv2 {
	return m.obx
}


func (m PPV_PCA_PID_GOL_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PPV_PCA_PID_GOL_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PPV_PCA_PID_GOL_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PPV_PCA_PID_GOL_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PPV_PCA_PID_GOL_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PPV_PCA_PID_GOL_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PRR_PC5v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*PRR_PC5_PIDv2 // Required
}

func (m *PRR_PC5v2) MSH() *MSH {
	return m.msh
}

func (m *PRR_PC5v2) MSA() *MSA {
	return m.msa
}

func (m *PRR_PC5v2) ERR() *ERR {
	return m.err
}

func (m *PRR_PC5v2) QRD() *QRD {
	return m.qrd
}

func (m *PRR_PC5v2) GroupByPID() []*PRR_PC5_PIDv2 {
	return m.pid
}


func (m PRR_PC5v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
	}, nil
}
type PRR_PC5_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	prb []*PRR_PC5_PID_PRBv2 // Required
}

func (m *PRR_PC5_PIDv2) PID() *PID {
	return m.pid
}

func (m *PRR_PC5_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *PRR_PC5_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *PRR_PC5_PIDv2) GroupByPRB() []*PRR_PC5_PID_PRBv2 {
	return m.prb
}


func (m PRR_PC5_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"prb": m.prb,
	}, nil
}
type PRR_PC5_PID_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PRR_PC5_PID_PRB_ROLv2
	pth []*PRR_PC5_PID_PRB_PTHv2
	obx []*PRR_PC5_PID_PRB_OBXv2
	gol []*PRR_PC5_PID_PRB_GOLv2
	orc []*PRR_PC5_PID_PRB_ORCv2
}

func (m *PRR_PC5_PID_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PRR_PC5_PID_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PRR_PC5_PID_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PRR_PC5_PID_PRBv2) GroupByROL() []*PRR_PC5_PID_PRB_ROLv2 {
	return m.rol
}

func (m *PRR_PC5_PID_PRBv2) GroupByPTH() []*PRR_PC5_PID_PRB_PTHv2 {
	return m.pth
}

func (m *PRR_PC5_PID_PRBv2) GroupByOBX() []*PRR_PC5_PID_PRB_OBXv2 {
	return m.obx
}

func (m *PRR_PC5_PID_PRBv2) GroupByGOL() []*PRR_PC5_PID_PRB_GOLv2 {
	return m.gol
}

func (m *PRR_PC5_PID_PRBv2) GroupByORC() []*PRR_PC5_PID_PRB_ORCv2 {
	return m.orc
}


func (m PRR_PC5_PID_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"pth": m.pth,
		"obx": m.obx,
		"gol": m.gol,
		"orc": m.orc,
	}, nil
}
type PRR_PC5_PID_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PRR_PC5_PID_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PRR_PC5_PID_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PRR_PC5_PID_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PRR_PC5_PID_PRB_PTHv2 struct {
	pth *PTH // Required
	var_ []*VAR
}

func (m *PRR_PC5_PID_PRB_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PRR_PC5_PID_PRB_PTHv2) AllVAR() []*VAR {
	return m.var_
}


func (m PRR_PC5_PID_PRB_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"var_": m.var_,
	}, nil
}
type PRR_PC5_PID_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PRR_PC5_PID_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PRR_PC5_PID_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PRR_PC5_PID_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PRR_PC5_PID_PRB_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PRR_PC5_PID_PRB_GOL_ROLv2
	obx []*PRR_PC5_PID_PRB_GOL_OBXv2
}

func (m *PRR_PC5_PID_PRB_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PRR_PC5_PID_PRB_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PRR_PC5_PID_PRB_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PRR_PC5_PID_PRB_GOLv2) GroupByROL() []*PRR_PC5_PID_PRB_GOL_ROLv2 {
	return m.rol
}

func (m *PRR_PC5_PID_PRB_GOLv2) GroupByOBX() []*PRR_PC5_PID_PRB_GOL_OBXv2 {
	return m.obx
}


func (m PRR_PC5_PID_PRB_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PRR_PC5_PID_PRB_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PRR_PC5_PID_PRB_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PRR_PC5_PID_PRB_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PRR_PC5_PID_PRB_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PRR_PC5_PID_PRB_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PRR_PC5_PID_PRB_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PRR_PC5_PID_PRB_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PRR_PC5_PID_PRB_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PRR_PC5_PID_PRB_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PRR_PC5_PID_PRB_ORC_OBXv2
}

func (m *PRR_PC5_PID_PRB_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PRR_PC5_PID_PRB_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PRR_PC5_PID_PRB_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PRR_PC5_PID_PRB_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PRR_PC5_PID_PRB_ORCv2) GroupByOBX() []*PRR_PC5_PID_PRB_ORC_OBXv2 {
	return m.obx
}


func (m PRR_PC5_PID_PRB_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PRR_PC5_PID_PRB_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PRR_PC5_PID_PRB_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PRR_PC5_PID_PRB_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PRR_PC5_PID_PRB_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PRR_PC5_PID_PRB_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type PTR_PCFv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	pid []*PTR_PCF_PIDv2 // Required
}

func (m *PTR_PCFv2) MSH() *MSH {
	return m.msh
}

func (m *PTR_PCFv2) MSA() *MSA {
	return m.msa
}

func (m *PTR_PCFv2) ERR() *ERR {
	return m.err
}

func (m *PTR_PCFv2) QRD() *QRD {
	return m.qrd
}

func (m *PTR_PCFv2) GroupByPID() []*PTR_PCF_PIDv2 {
	return m.pid
}


func (m PTR_PCFv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
	}, nil
}
type PTR_PCF_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	pth []*PTR_PCF_PID_PTHv2 // Required
}

func (m *PTR_PCF_PIDv2) PID() *PID {
	return m.pid
}

func (m *PTR_PCF_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *PTR_PCF_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *PTR_PCF_PIDv2) GroupByPTH() []*PTR_PCF_PID_PTHv2 {
	return m.pth
}


func (m PTR_PCF_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"pth": m.pth,
	}, nil
}
type PTR_PCF_PID_PTHv2 struct {
	pth *PTH // Required
	nte []*NTE
	var_ []*VAR
	rol []*PTR_PCF_PID_PTH_ROLv2
	prb []*PTR_PCF_PID_PTH_PRBv2
}

func (m *PTR_PCF_PID_PTHv2) PTH() *PTH {
	return m.pth
}

func (m *PTR_PCF_PID_PTHv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PTR_PCF_PID_PTHv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PTR_PCF_PID_PTHv2) GroupByROL() []*PTR_PCF_PID_PTH_ROLv2 {
	return m.rol
}

func (m *PTR_PCF_PID_PTHv2) GroupByPRB() []*PTR_PCF_PID_PTH_PRBv2 {
	return m.prb
}


func (m PTR_PCF_PID_PTHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pth": m.pth,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"prb": m.prb,
	}, nil
}
type PTR_PCF_PID_PTH_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PTR_PCF_PID_PTH_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PTR_PCF_PID_PTH_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PTR_PCF_PID_PTH_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PTR_PCF_PID_PTH_PRBv2 struct {
	prb *PRB // Required
	nte []*NTE
	var_ []*VAR
	rol []*PTR_PCF_PID_PTH_PRB_ROLv2
	obx []*PTR_PCF_PID_PTH_PRB_OBXv2
	gol []*PTR_PCF_PID_PTH_PRB_GOLv2
	orc []*PTR_PCF_PID_PTH_PRB_ORCv2
}

func (m *PTR_PCF_PID_PTH_PRBv2) PRB() *PRB {
	return m.prb
}

func (m *PTR_PCF_PID_PTH_PRBv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PTR_PCF_PID_PTH_PRBv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PTR_PCF_PID_PTH_PRBv2) GroupByROL() []*PTR_PCF_PID_PTH_PRB_ROLv2 {
	return m.rol
}

func (m *PTR_PCF_PID_PTH_PRBv2) GroupByOBX() []*PTR_PCF_PID_PTH_PRB_OBXv2 {
	return m.obx
}

func (m *PTR_PCF_PID_PTH_PRBv2) GroupByGOL() []*PTR_PCF_PID_PTH_PRB_GOLv2 {
	return m.gol
}

func (m *PTR_PCF_PID_PTH_PRBv2) GroupByORC() []*PTR_PCF_PID_PTH_PRB_ORCv2 {
	return m.orc
}


func (m PTR_PCF_PID_PTH_PRBv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prb": m.prb,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
		"gol": m.gol,
		"orc": m.orc,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PTR_PCF_PID_PTH_PRB_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PTR_PCF_PID_PTH_PRB_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PTR_PCF_PID_PTH_PRB_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PTR_PCF_PID_PTH_PRB_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PTR_PCF_PID_PTH_PRB_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PTR_PCF_PID_PTH_PRB_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_GOLv2 struct {
	gol *GOL // Required
	nte []*NTE
	var_ []*VAR
	rol []*PTR_PCF_PID_PTH_PRB_GOL_ROLv2
	obx []*PTR_PCF_PID_PTH_PRB_GOL_OBXv2
}

func (m *PTR_PCF_PID_PTH_PRB_GOLv2) GOL() *GOL {
	return m.gol
}

func (m *PTR_PCF_PID_PTH_PRB_GOLv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PTR_PCF_PID_PTH_PRB_GOLv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PTR_PCF_PID_PTH_PRB_GOLv2) GroupByROL() []*PTR_PCF_PID_PTH_PRB_GOL_ROLv2 {
	return m.rol
}

func (m *PTR_PCF_PID_PTH_PRB_GOLv2) GroupByOBX() []*PTR_PCF_PID_PTH_PRB_GOL_OBXv2 {
	return m.obx
}


func (m PTR_PCF_PID_PTH_PRB_GOLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"gol": m.gol,
		"nte": m.nte,
		"var_": m.var_,
		"rol": m.rol,
		"obx": m.obx,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_GOL_ROLv2 struct {
	rol *ROL // Required
	var_ []*VAR
}

func (m *PTR_PCF_PID_PTH_PRB_GOL_ROLv2) ROL() *ROL {
	return m.rol
}

func (m *PTR_PCF_PID_PTH_PRB_GOL_ROLv2) AllVAR() []*VAR {
	return m.var_
}


func (m PTR_PCF_PID_PTH_PRB_GOL_ROLv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rol": m.rol,
		"var_": m.var_,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_GOL_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *PTR_PCF_PID_PTH_PRB_GOL_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PTR_PCF_PID_PTH_PRB_GOL_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m PTR_PCF_PID_PTH_PRB_GOL_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	nte []*NTE
	var_ []*VAR
	obx []*PTR_PCF_PID_PTH_PRB_ORC_OBXv2
}

func (m *PTR_PCF_PID_PTH_PRB_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *PTR_PCF_PID_PTH_PRB_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *PTR_PCF_PID_PTH_PRB_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PTR_PCF_PID_PTH_PRB_ORCv2) AllVAR() []*VAR {
	return m.var_
}

func (m *PTR_PCF_PID_PTH_PRB_ORCv2) GroupByOBX() []*PTR_PCF_PID_PTH_PRB_ORC_OBXv2 {
	return m.obx
}


func (m PTR_PCF_PID_PTH_PRB_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"var_": m.var_,
		"obx": m.obx,
	}, nil
}
type PTR_PCF_PID_PTH_PRB_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
	var_ []*VAR
}

func (m *PTR_PCF_PID_PTH_PRB_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *PTR_PCF_PID_PTH_PRB_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m *PTR_PCF_PID_PTH_PRB_ORC_OBXv2) AllVAR() []*VAR {
	return m.var_
}


func (m PTR_PCF_PID_PTH_PRB_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
		"var_": m.var_,
	}, nil
}
type QCK_Q02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK
}

func (m *QCK_Q02v2) MSH() *MSH {
	return m.msh
}

func (m *QCK_Q02v2) MSA() *MSA {
	return m.msa
}

func (m *QCK_Q02v2) ERR() *ERR {
	return m.err
}

func (m *QCK_Q02v2) QAK() *QAK {
	return m.qak
}


func (m QCK_Q02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
	}, nil
}
type QRY_A19v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
}

func (m *QRY_A19v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_A19v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_A19v2) QRF() *QRF {
	return m.qrf
}


func (m QRY_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}
type QRY_P04v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_P04v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_P04v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_P04v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_P04v2) DSC() *DSC {
	return m.dsc
}


func (m QRY_P04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type QRY_PC4v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
}

func (m *QRY_PC4v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_PC4v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_PC4v2) QRF() *QRF {
	return m.qrf
}


func (m QRY_PC4v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}
type QRY_Q01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_Q01v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_Q01v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m QRY_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type QRY_Q02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_Q02v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_Q02v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_Q02v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_Q02v2) DSC() *DSC {
	return m.dsc
}


func (m QRY_Q02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}
type QRY_R02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF // Required
}

func (m *QRY_R02v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_R02v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_R02v2) QRF() *QRF {
	return m.qrf
}


func (m QRY_R02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}
type QRY_T12v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
}

func (m *QRY_T12v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_T12v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_T12v2) QRF() *QRF {
	return m.qrf
}


func (m QRY_T12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}
type RAR_RARv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd []*RAR_RAR_QRDv2 // Required
	dsc *DSC
}

func (m *RAR_RARv2) MSH() *MSH {
	return m.msh
}

func (m *RAR_RARv2) MSA() *MSA {
	return m.msa
}

func (m *RAR_RARv2) ERR() *ERR {
	return m.err
}

func (m *RAR_RARv2) GroupByQRD() []*RAR_RAR_QRDv2 {
	return m.qrd
}

func (m *RAR_RARv2) DSC() *DSC {
	return m.dsc
}


func (m RAR_RARv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"dsc": m.dsc,
	}, nil
}
type RAR_RAR_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	nte []*NTE
	orc []*RAR_RAR_QRD_ORCv2 // Required
}

func (m *RAR_RAR_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *RAR_RAR_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *RAR_RAR_QRDv2) PID() *PID {
	return m.pid
}

func (m *RAR_RAR_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RAR_RAR_QRDv2) GroupByORC() []*RAR_RAR_QRD_ORCv2 {
	return m.orc
}


func (m RAR_RAR_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RAR_RAR_QRD_ORCv2 struct {
	orc *ORC // Required
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
	rxa []*RXA // Required
}

func (m *RAR_RAR_QRD_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RAR_RAR_QRD_ORCv2) RXE() *RXE {
	return m.rxe
}

func (m *RAR_RAR_QRD_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RAR_RAR_QRD_ORCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RAR_RAR_QRD_ORCv2) AllRXA() []*RXA {
	return m.rxa
}


func (m RAR_RAR_QRD_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"rxa": m.rxa,
	}, nil
}
type RAS_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *RAS_O01_PIDv2
	orc []*RAS_O01_ORCv2 // Required
}

func (m *RAS_O01v2) MSH() *MSH {
	return m.msh
}

func (m *RAS_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RAS_O01v2) GroupByPID() *RAS_O01_PIDv2 {
	return m.pid
}

func (m *RAS_O01v2) GroupByORC() []*RAS_O01_ORCv2 {
	return m.orc
}


func (m RAS_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type RAS_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	al1 []*AL1
	pv1 *PV1 // Required
	pv2 *PV2
}

func (m *RAS_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *RAS_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *RAS_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RAS_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RAS_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *RAS_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}


func (m RAS_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"al1": m.al1,
		"pv1": m.pv1,
		"pv2": m.pv2,
	}, nil
}
type RAS_O01_ORCv2 struct {
	orc *ORC // Required
	rxo *RAS_O01_ORC_RXOv2
	rxe *RAS_O01_ORC_RXEv2
	rxa []*RXA // Required
	rxr *RXR // Required
	obx []*RAS_O01_ORC_OBXv2
	cti []*CTI
}

func (m *RAS_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RAS_O01_ORCv2) GroupByRXO() *RAS_O01_ORC_RXOv2 {
	return m.rxo
}

func (m *RAS_O01_ORCv2) GroupByRXE() *RAS_O01_ORC_RXEv2 {
	return m.rxe
}

func (m *RAS_O01_ORCv2) AllRXA() []*RXA {
	return m.rxa
}

func (m *RAS_O01_ORCv2) RXR() *RXR {
	return m.rxr
}

func (m *RAS_O01_ORCv2) GroupByOBX() []*RAS_O01_ORC_OBXv2 {
	return m.obx
}

func (m *RAS_O01_ORCv2) AllCTI() []*CTI {
	return m.cti
}


func (m RAS_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"rxe": m.rxe,
		"rxa": m.rxa,
		"rxr": m.rxr,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}
type RAS_O01_ORC_RXOv2 struct {
	rxo *RXO // Required
	nte []*NTE // Required
	rxr []*RXR // Required
	rxc *RAS_O01_ORC_RXO_RXCv2
}

func (m *RAS_O01_ORC_RXOv2) RXO() *RXO {
	return m.rxo
}

func (m *RAS_O01_ORC_RXOv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RAS_O01_ORC_RXOv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RAS_O01_ORC_RXOv2) GroupByRXC() *RAS_O01_ORC_RXO_RXCv2 {
	return m.rxc
}


func (m RAS_O01_ORC_RXOv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxo": m.rxo,
		"nte": m.nte,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RAS_O01_ORC_RXO_RXCv2 struct {
	rxc []*RXC // Required
	nte []*NTE
}

func (m *RAS_O01_ORC_RXO_RXCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RAS_O01_ORC_RXO_RXCv2) AllNTE() []*NTE {
	return m.nte
}


func (m RAS_O01_ORC_RXO_RXCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxc": m.rxc,
		"nte": m.nte,
	}, nil
}
type RAS_O01_ORC_RXEv2 struct {
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RAS_O01_ORC_RXEv2) RXE() *RXE {
	return m.rxe
}

func (m *RAS_O01_ORC_RXEv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RAS_O01_ORC_RXEv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RAS_O01_ORC_RXEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RAS_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RAS_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RAS_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RAS_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RCI_I05v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	prd []*RCI_I05_PRDv2 // Required
	pid *PID // Required
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	obr []*RCI_I05_OBRv2
	nte []*NTE
}

func (m *RCI_I05v2) MSH() *MSH {
	return m.msh
}

func (m *RCI_I05v2) MSA() *MSA {
	return m.msa
}

func (m *RCI_I05v2) QRD() *QRD {
	return m.qrd
}

func (m *RCI_I05v2) QRF() *QRF {
	return m.qrf
}

func (m *RCI_I05v2) GroupByPRD() []*RCI_I05_PRDv2 {
	return m.prd
}

func (m *RCI_I05v2) PID() *PID {
	return m.pid
}

func (m *RCI_I05v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *RCI_I05v2) AllDRG() []*DRG {
	return m.drg
}

func (m *RCI_I05v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RCI_I05v2) GroupByOBR() []*RCI_I05_OBRv2 {
	return m.obr
}

func (m *RCI_I05v2) AllNTE() []*NTE {
	return m.nte
}


func (m RCI_I05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"prd": m.prd,
		"pid": m.pid,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"obr": m.obr,
		"nte": m.nte,
	}, nil
}
type RCI_I05_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RCI_I05_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RCI_I05_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RCI_I05_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RCI_I05_OBRv2 struct {
	obr *OBR // Required
	nte []*NTE
	obx []*RCI_I05_OBR_OBXv2
}

func (m *RCI_I05_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *RCI_I05_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RCI_I05_OBRv2) GroupByOBX() []*RCI_I05_OBR_OBXv2 {
	return m.obx
}


func (m RCI_I05_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type RCI_I05_OBR_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RCI_I05_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RCI_I05_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RCI_I05_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RCL_I06v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	prd []*RCL_I06_PRDv2 // Required
	pid *PID // Required
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	nte []*NTE
	dsp []*DSP
	dsc *DSC
}

func (m *RCL_I06v2) MSH() *MSH {
	return m.msh
}

func (m *RCL_I06v2) MSA() *MSA {
	return m.msa
}

func (m *RCL_I06v2) QRD() *QRD {
	return m.qrd
}

func (m *RCL_I06v2) QRF() *QRF {
	return m.qrf
}

func (m *RCL_I06v2) GroupByPRD() []*RCL_I06_PRDv2 {
	return m.prd
}

func (m *RCL_I06v2) PID() *PID {
	return m.pid
}

func (m *RCL_I06v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *RCL_I06v2) AllDRG() []*DRG {
	return m.drg
}

func (m *RCL_I06v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RCL_I06v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RCL_I06v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *RCL_I06v2) DSC() *DSC {
	return m.dsc
}


func (m RCL_I06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"prd": m.prd,
		"pid": m.pid,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"nte": m.nte,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type RCL_I06_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RCL_I06_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RCL_I06_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RCL_I06_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RDE_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *RDE_O01_PIDv2
	orc []*RDE_O01_ORCv2 // Required
}

func (m *RDE_O01v2) MSH() *MSH {
	return m.msh
}

func (m *RDE_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDE_O01v2) GroupByPID() *RDE_O01_PIDv2 {
	return m.pid
}

func (m *RDE_O01v2) GroupByORC() []*RDE_O01_ORCv2 {
	return m.orc
}


func (m RDE_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type RDE_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*RDE_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *RDE_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *RDE_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *RDE_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDE_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *RDE_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *RDE_O01_PIDv2) GroupByIN1() []*RDE_O01_PID_IN1v2 {
	return m.in1
}

func (m *RDE_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *RDE_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m RDE_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type RDE_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RDE_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RDE_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RDE_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RDE_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RDE_O01_ORCv2 struct {
	orc *ORC // Required
	rxo *RDE_O01_ORC_RXOv2
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
	obx []*RDE_O01_ORC_OBXv2 // Required
	cti *CTI
}

func (m *RDE_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RDE_O01_ORCv2) GroupByRXO() *RDE_O01_ORC_RXOv2 {
	return m.rxo
}

func (m *RDE_O01_ORCv2) RXE() *RXE {
	return m.rxe
}

func (m *RDE_O01_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDE_O01_ORCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDE_O01_ORCv2) GroupByOBX() []*RDE_O01_ORC_OBXv2 {
	return m.obx
}

func (m *RDE_O01_ORCv2) CTI() *CTI {
	return m.cti
}


func (m RDE_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}
type RDE_O01_ORC_RXOv2 struct {
	rxo *RXO // Required
	nte []*NTE
	rxr []*RXR // Required
	rxc *RDE_O01_ORC_RXO_RXCv2
}

func (m *RDE_O01_ORC_RXOv2) RXO() *RXO {
	return m.rxo
}

func (m *RDE_O01_ORC_RXOv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDE_O01_ORC_RXOv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDE_O01_ORC_RXOv2) GroupByRXC() *RDE_O01_ORC_RXO_RXCv2 {
	return m.rxc
}


func (m RDE_O01_ORC_RXOv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxo": m.rxo,
		"nte": m.nte,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RDE_O01_ORC_RXO_RXCv2 struct {
	rxc []*RXC // Required
	nte []*NTE
}

func (m *RDE_O01_ORC_RXO_RXCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDE_O01_ORC_RXO_RXCv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDE_O01_ORC_RXO_RXCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxc": m.rxc,
		"nte": m.nte,
	}, nil
}
type RDE_O01_ORC_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *RDE_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RDE_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDE_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RDO_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *RDO_O01_PIDv2
	orc []*RDO_O01_ORCv2 // Required
}

func (m *RDO_O01v2) MSH() *MSH {
	return m.msh
}

func (m *RDO_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDO_O01v2) GroupByPID() *RDO_O01_PIDv2 {
	return m.pid
}

func (m *RDO_O01v2) GroupByORC() []*RDO_O01_ORCv2 {
	return m.orc
}


func (m RDO_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type RDO_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*RDO_O01_PID_IN1v2
	gt1 *GT1
	al1 []*AL1
}

func (m *RDO_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *RDO_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *RDO_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDO_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *RDO_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *RDO_O01_PIDv2) GroupByIN1() []*RDO_O01_PID_IN1v2 {
	return m.in1
}

func (m *RDO_O01_PIDv2) GT1() *GT1 {
	return m.gt1
}

func (m *RDO_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}


func (m RDO_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"gt1": m.gt1,
		"al1": m.al1,
	}, nil
}
type RDO_O01_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RDO_O01_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RDO_O01_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RDO_O01_PID_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RDO_O01_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RDO_O01_ORCv2 struct {
	orc *ORC // Required
	rxo *RXO // Required
	nte []*NTE
	rxr []*RXR // Required
	rxc *RDO_O01_ORC_RXCv2
	obx []*RDO_O01_ORC_OBXv2
	blg *BLG
}

func (m *RDO_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RDO_O01_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *RDO_O01_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDO_O01_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDO_O01_ORCv2) GroupByRXC() *RDO_O01_ORC_RXCv2 {
	return m.rxc
}

func (m *RDO_O01_ORCv2) GroupByOBX() []*RDO_O01_ORC_OBXv2 {
	return m.obx
}

func (m *RDO_O01_ORCv2) BLG() *BLG {
	return m.blg
}


func (m RDO_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"nte": m.nte,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"obx": m.obx,
		"blg": m.blg,
	}, nil
}
type RDO_O01_ORC_RXCv2 struct {
	rxc []*RXC // Required
	nte []*NTE
}

func (m *RDO_O01_ORC_RXCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDO_O01_ORC_RXCv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDO_O01_ORC_RXCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxc": m.rxc,
		"nte": m.nte,
	}, nil
}
type RDO_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RDO_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RDO_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDO_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RDR_RDRv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd []*RDR_RDR_QRDv2 // Required
	dsc *DSC
}

func (m *RDR_RDRv2) MSH() *MSH {
	return m.msh
}

func (m *RDR_RDRv2) MSA() *MSA {
	return m.msa
}

func (m *RDR_RDRv2) ERR() *ERR {
	return m.err
}

func (m *RDR_RDRv2) GroupByQRD() []*RDR_RDR_QRDv2 {
	return m.qrd
}

func (m *RDR_RDRv2) DSC() *DSC {
	return m.dsc
}


func (m RDR_RDRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"dsc": m.dsc,
	}, nil
}
type RDR_RDR_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	nte []*NTE
	orc []*RDR_RDR_QRD_ORCv2 // Required
}

func (m *RDR_RDR_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *RDR_RDR_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *RDR_RDR_QRDv2) PID() *PID {
	return m.pid
}

func (m *RDR_RDR_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDR_RDR_QRDv2) GroupByORC() []*RDR_RDR_QRD_ORCv2 {
	return m.orc
}


func (m RDR_RDR_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RDR_RDR_QRD_ORCv2 struct {
	orc *ORC // Required
	rxe *RXE // Required
	rxr *RXR // Required
	rxc []*RXC
	rxd []*RDR_RDR_QRD_ORC_RXDv2 // Required
}

func (m *RDR_RDR_QRD_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RDR_RDR_QRD_ORCv2) RXE() *RXE {
	return m.rxe
}

func (m *RDR_RDR_QRD_ORCv2) RXR() *RXR {
	return m.rxr
}

func (m *RDR_RDR_QRD_ORCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDR_RDR_QRD_ORCv2) GroupByRXD() []*RDR_RDR_QRD_ORC_RXDv2 {
	return m.rxd
}


func (m RDR_RDR_QRD_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"rxd": m.rxd,
	}, nil
}
type RDR_RDR_QRD_ORC_RXDv2 struct {
	rxd *RXD // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RDR_RDR_QRD_ORC_RXDv2) RXD() *RXD {
	return m.rxd
}

func (m *RDR_RDR_QRD_ORC_RXDv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDR_RDR_QRD_ORC_RXDv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RDR_RDR_QRD_ORC_RXDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxd": m.rxd,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RDS_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *RDS_O01_PIDv2
	orc []*RDS_O01_ORCv2 // Required
}

func (m *RDS_O01v2) MSH() *MSH {
	return m.msh
}

func (m *RDS_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDS_O01v2) GroupByPID() *RDS_O01_PIDv2 {
	return m.pid
}

func (m *RDS_O01v2) GroupByORC() []*RDS_O01_ORCv2 {
	return m.orc
}


func (m RDS_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type RDS_O01_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nte []*NTE
	al1 []*AL1
	pv1 *PV1 // Required
	pv2 *PV2
}

func (m *RDS_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *RDS_O01_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *RDS_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDS_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RDS_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *RDS_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}


func (m RDS_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nte": m.nte,
		"al1": m.al1,
		"pv1": m.pv1,
		"pv2": m.pv2,
	}, nil
}
type RDS_O01_ORCv2 struct {
	orc *ORC // Required
	rxo *RDS_O01_ORC_RXOv2
	rxe *RDS_O01_ORC_RXEv2
	rxd *RXD // Required
	rxr []*RXR // Required
	rxc []*RXC
	obx []*RDS_O01_ORC_OBXv2 // Required
}

func (m *RDS_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RDS_O01_ORCv2) GroupByRXO() *RDS_O01_ORC_RXOv2 {
	return m.rxo
}

func (m *RDS_O01_ORCv2) GroupByRXE() *RDS_O01_ORC_RXEv2 {
	return m.rxe
}

func (m *RDS_O01_ORCv2) RXD() *RXD {
	return m.rxd
}

func (m *RDS_O01_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDS_O01_ORCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDS_O01_ORCv2) GroupByOBX() []*RDS_O01_ORC_OBXv2 {
	return m.obx
}


func (m RDS_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"rxe": m.rxe,
		"rxd": m.rxd,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"obx": m.obx,
	}, nil
}
type RDS_O01_ORC_RXOv2 struct {
	rxo *RXO // Required
	nte []*NTE // Required
	rxr []*RXR // Required
	rxc *RDS_O01_ORC_RXO_RXCv2
}

func (m *RDS_O01_ORC_RXOv2) RXO() *RXO {
	return m.rxo
}

func (m *RDS_O01_ORC_RXOv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RDS_O01_ORC_RXOv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDS_O01_ORC_RXOv2) GroupByRXC() *RDS_O01_ORC_RXO_RXCv2 {
	return m.rxc
}


func (m RDS_O01_ORC_RXOv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxo": m.rxo,
		"nte": m.nte,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RDS_O01_ORC_RXO_RXCv2 struct {
	rxc []*RXC // Required
	nte []*NTE
}

func (m *RDS_O01_ORC_RXO_RXCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RDS_O01_ORC_RXO_RXCv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDS_O01_ORC_RXO_RXCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxc": m.rxc,
		"nte": m.nte,
	}, nil
}
type RDS_O01_ORC_RXEv2 struct {
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RDS_O01_ORC_RXEv2) RXE() *RXE {
	return m.rxe
}

func (m *RDS_O01_ORC_RXEv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RDS_O01_ORC_RXEv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RDS_O01_ORC_RXEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RDS_O01_ORC_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RDS_O01_ORC_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RDS_O01_ORC_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RDS_O01_ORC_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type REF_I12v2 struct {
	msh *MSH // Required
	rf1 *RF1
	aut *AUT // Required
	ctd *CTD
	prd []*REF_I12_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*REF_I12_IN1v2
	acc *ACC
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	pr1 []*REF_I12_PR1v2
	obr []*REF_I12_OBRv2
	pv1 *PV1 // Required
	pv2 *PV2
	nte []*NTE
}

func (m *REF_I12v2) MSH() *MSH {
	return m.msh
}

func (m *REF_I12v2) RF1() *RF1 {
	return m.rf1
}

func (m *REF_I12v2) AUT() *AUT {
	return m.aut
}

func (m *REF_I12v2) CTD() *CTD {
	return m.ctd
}

func (m *REF_I12v2) GroupByPRD() []*REF_I12_PRDv2 {
	return m.prd
}

func (m *REF_I12v2) PID() *PID {
	return m.pid
}

func (m *REF_I12v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *REF_I12v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *REF_I12v2) GroupByIN1() []*REF_I12_IN1v2 {
	return m.in1
}

func (m *REF_I12v2) ACC() *ACC {
	return m.acc
}

func (m *REF_I12v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *REF_I12v2) AllDRG() []*DRG {
	return m.drg
}

func (m *REF_I12v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *REF_I12v2) GroupByPR1() []*REF_I12_PR1v2 {
	return m.pr1
}

func (m *REF_I12v2) GroupByOBR() []*REF_I12_OBRv2 {
	return m.obr
}

func (m *REF_I12v2) PV1() *PV1 {
	return m.pv1
}

func (m *REF_I12v2) PV2() *PV2 {
	return m.pv2
}

func (m *REF_I12v2) AllNTE() []*NTE {
	return m.nte
}


func (m REF_I12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"rf1": m.rf1,
		"aut": m.aut,
		"ctd": m.ctd,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"pr1": m.pr1,
		"obr": m.obr,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"nte": m.nte,
	}, nil
}
type REF_I12_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *REF_I12_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *REF_I12_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m REF_I12_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type REF_I12_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *REF_I12_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *REF_I12_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *REF_I12_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m REF_I12_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type REF_I12_PR1v2 struct {
	pr1 *PR1 // Required
	aut *AUT // Required
	ctd *CTD
}

func (m *REF_I12_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *REF_I12_PR1v2) AUT() *AUT {
	return m.aut
}

func (m *REF_I12_PR1v2) CTD() *CTD {
	return m.ctd
}


func (m REF_I12_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"aut": m.aut,
		"ctd": m.ctd,
	}, nil
}
type REF_I12_OBRv2 struct {
	obr *OBR // Required
	nte []*NTE
	obx []*REF_I12_OBR_OBXv2
}

func (m *REF_I12_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *REF_I12_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *REF_I12_OBRv2) GroupByOBX() []*REF_I12_OBR_OBXv2 {
	return m.obx
}


func (m REF_I12_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type REF_I12_OBR_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *REF_I12_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *REF_I12_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m REF_I12_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RER_RERv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd []*RER_RER_QRDv2 // Required
	dsc *DSC
}

func (m *RER_RERv2) MSH() *MSH {
	return m.msh
}

func (m *RER_RERv2) MSA() *MSA {
	return m.msa
}

func (m *RER_RERv2) ERR() *ERR {
	return m.err
}

func (m *RER_RERv2) GroupByQRD() []*RER_RER_QRDv2 {
	return m.qrd
}

func (m *RER_RERv2) DSC() *DSC {
	return m.dsc
}


func (m RER_RERv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"dsc": m.dsc,
	}, nil
}
type RER_RER_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	nte []*NTE
	orc []*RER_RER_QRD_ORCv2 // Required
}

func (m *RER_RER_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *RER_RER_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *RER_RER_QRDv2) PID() *PID {
	return m.pid
}

func (m *RER_RER_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RER_RER_QRDv2) GroupByORC() []*RER_RER_QRD_ORCv2 {
	return m.orc
}


func (m RER_RER_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RER_RER_QRD_ORCv2 struct {
	orc *ORC // Required
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RER_RER_QRD_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RER_RER_QRD_ORCv2) RXE() *RXE {
	return m.rxe
}

func (m *RER_RER_QRD_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RER_RER_QRD_ORCv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RER_RER_QRD_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RGR_RGRv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd []*RGR_RGR_QRDv2 // Required
	dsc *DSC
}

func (m *RGR_RGRv2) MSH() *MSH {
	return m.msh
}

func (m *RGR_RGRv2) MSA() *MSA {
	return m.msa
}

func (m *RGR_RGRv2) ERR() *ERR {
	return m.err
}

func (m *RGR_RGRv2) GroupByQRD() []*RGR_RGR_QRDv2 {
	return m.qrd
}

func (m *RGR_RGRv2) DSC() *DSC {
	return m.dsc
}


func (m RGR_RGRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"dsc": m.dsc,
	}, nil
}
type RGR_RGR_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	nte []*NTE
	orc []*RGR_RGR_QRD_ORCv2 // Required
}

func (m *RGR_RGR_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *RGR_RGR_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *RGR_RGR_QRDv2) PID() *PID {
	return m.pid
}

func (m *RGR_RGR_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RGR_RGR_QRDv2) GroupByORC() []*RGR_RGR_QRD_ORCv2 {
	return m.orc
}


func (m RGR_RGR_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RGR_RGR_QRD_ORCv2 struct {
	orc *ORC // Required
	rxe *RGR_RGR_QRD_ORC_RXEv2
	rxg []*RXG // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RGR_RGR_QRD_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RGR_RGR_QRD_ORCv2) GroupByRXE() *RGR_RGR_QRD_ORC_RXEv2 {
	return m.rxe
}

func (m *RGR_RGR_QRD_ORCv2) AllRXG() []*RXG {
	return m.rxg
}

func (m *RGR_RGR_QRD_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RGR_RGR_QRD_ORCv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RGR_RGR_QRD_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxe": m.rxe,
		"rxg": m.rxg,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RGR_RGR_QRD_ORC_RXEv2 struct {
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RGR_RGR_QRD_ORC_RXEv2) RXE() *RXE {
	return m.rxe
}

func (m *RGR_RGR_QRD_ORC_RXEv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RGR_RGR_QRD_ORC_RXEv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RGR_RGR_QRD_ORC_RXEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RGV_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *RGV_O01_PIDv2
	orc []*RGV_O01_ORCv2 // Required
}

func (m *RGV_O01v2) MSH() *MSH {
	return m.msh
}

func (m *RGV_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RGV_O01v2) GroupByPID() *RGV_O01_PIDv2 {
	return m.pid
}

func (m *RGV_O01v2) GroupByORC() []*RGV_O01_ORCv2 {
	return m.orc
}


func (m RGV_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}
type RGV_O01_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	al1 []*AL1
	pv1 *PV1 // Required
	pv2 *PV2
}

func (m *RGV_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *RGV_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RGV_O01_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RGV_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *RGV_O01_PIDv2) PV2() *PV2 {
	return m.pv2
}


func (m RGV_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"al1": m.al1,
		"pv1": m.pv1,
		"pv2": m.pv2,
	}, nil
}
type RGV_O01_ORCv2 struct {
	orc *ORC // Required
	rxo *RXO // Required
	nte []*NTE // Required
	rxr []*RXR // Required
	rxc *RGV_O01_ORC_RXCv2
	rxe *RGV_O01_ORC_RXEv2
	rxg []*RGV_O01_ORC_RXGv2 // Required
}

func (m *RGV_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RGV_O01_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *RGV_O01_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RGV_O01_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RGV_O01_ORCv2) GroupByRXC() *RGV_O01_ORC_RXCv2 {
	return m.rxc
}

func (m *RGV_O01_ORCv2) GroupByRXE() *RGV_O01_ORC_RXEv2 {
	return m.rxe
}

func (m *RGV_O01_ORCv2) GroupByRXG() []*RGV_O01_ORC_RXGv2 {
	return m.rxg
}


func (m RGV_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"nte": m.nte,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"rxe": m.rxe,
		"rxg": m.rxg,
	}, nil
}
type RGV_O01_ORC_RXCv2 struct {
	rxc []*RXC // Required
	nte []*NTE
}

func (m *RGV_O01_ORC_RXCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RGV_O01_ORC_RXCv2) AllNTE() []*NTE {
	return m.nte
}


func (m RGV_O01_ORC_RXCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxc": m.rxc,
		"nte": m.nte,
	}, nil
}
type RGV_O01_ORC_RXEv2 struct {
	rxe *RXE // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RGV_O01_ORC_RXEv2) RXE() *RXE {
	return m.rxe
}

func (m *RGV_O01_ORC_RXEv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RGV_O01_ORC_RXEv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RGV_O01_ORC_RXEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxe": m.rxe,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RGV_O01_ORC_RXGv2 struct {
	rxg *RXG // Required
	rxr []*RXR // Required
	rxc []*RXC
	obx []*RGV_O01_ORC_RXG_OBXv2
}

func (m *RGV_O01_ORC_RXGv2) RXG() *RXG {
	return m.rxg
}

func (m *RGV_O01_ORC_RXGv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RGV_O01_ORC_RXGv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RGV_O01_ORC_RXGv2) GroupByOBX() []*RGV_O01_ORC_RXG_OBXv2 {
	return m.obx
}


func (m RGV_O01_ORC_RXGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxg": m.rxg,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"obx": m.obx,
	}, nil
}
type RGV_O01_ORC_RXG_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RGV_O01_ORC_RXG_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RGV_O01_ORC_RXG_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RGV_O01_ORC_RXG_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type ROR_RORv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd []*ROR_ROR_QRDv2 // Required
	dsc *DSC
}

func (m *ROR_RORv2) MSH() *MSH {
	return m.msh
}

func (m *ROR_RORv2) MSA() *MSA {
	return m.msa
}

func (m *ROR_RORv2) ERR() *ERR {
	return m.err
}

func (m *ROR_RORv2) GroupByQRD() []*ROR_ROR_QRDv2 {
	return m.qrd
}

func (m *ROR_RORv2) DSC() *DSC {
	return m.dsc
}


func (m ROR_RORv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"dsc": m.dsc,
	}, nil
}
type ROR_ROR_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	nte []*NTE
	orc []*ROR_ROR_QRD_ORCv2 // Required
}

func (m *ROR_ROR_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *ROR_ROR_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *ROR_ROR_QRDv2) PID() *PID {
	return m.pid
}

func (m *ROR_ROR_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ROR_ROR_QRDv2) GroupByORC() []*ROR_ROR_QRD_ORCv2 {
	return m.orc
}


func (m ROR_ROR_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type ROR_ROR_QRD_ORCv2 struct {
	orc *ORC // Required
	rxo *RXO // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *ROR_ROR_QRD_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ROR_ROR_QRD_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *ROR_ROR_QRD_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *ROR_ROR_QRD_ORCv2) AllRXC() []*RXC {
	return m.rxc
}


func (m ROR_ROR_QRD_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RPA_I08v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	rf1 *RF1
	aut *AUT // Required
	ctd *CTD
	prd []*RPA_I08_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*RPA_I08_IN1v2
	acc *ACC
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	pr1 []*RPA_I08_PR1v2 // Required
	obr []*RPA_I08_OBRv2
	pv1 *PV1 // Required
	pv2 *PV2
	nte []*NTE
}

func (m *RPA_I08v2) MSH() *MSH {
	return m.msh
}

func (m *RPA_I08v2) MSA() *MSA {
	return m.msa
}

func (m *RPA_I08v2) RF1() *RF1 {
	return m.rf1
}

func (m *RPA_I08v2) AUT() *AUT {
	return m.aut
}

func (m *RPA_I08v2) CTD() *CTD {
	return m.ctd
}

func (m *RPA_I08v2) GroupByPRD() []*RPA_I08_PRDv2 {
	return m.prd
}

func (m *RPA_I08v2) PID() *PID {
	return m.pid
}

func (m *RPA_I08v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RPA_I08v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RPA_I08v2) GroupByIN1() []*RPA_I08_IN1v2 {
	return m.in1
}

func (m *RPA_I08v2) ACC() *ACC {
	return m.acc
}

func (m *RPA_I08v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *RPA_I08v2) AllDRG() []*DRG {
	return m.drg
}

func (m *RPA_I08v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RPA_I08v2) GroupByPR1() []*RPA_I08_PR1v2 {
	return m.pr1
}

func (m *RPA_I08v2) GroupByOBR() []*RPA_I08_OBRv2 {
	return m.obr
}

func (m *RPA_I08v2) PV1() *PV1 {
	return m.pv1
}

func (m *RPA_I08v2) PV2() *PV2 {
	return m.pv2
}

func (m *RPA_I08v2) AllNTE() []*NTE {
	return m.nte
}


func (m RPA_I08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"rf1": m.rf1,
		"aut": m.aut,
		"ctd": m.ctd,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"pr1": m.pr1,
		"obr": m.obr,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"nte": m.nte,
	}, nil
}
type RPA_I08_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RPA_I08_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RPA_I08_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RPA_I08_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RPA_I08_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RPA_I08_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RPA_I08_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RPA_I08_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RPA_I08_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RPA_I08_PR1v2 struct {
	pr1 *PR1 // Required
	aut *AUT // Required
	ctd *CTD
}

func (m *RPA_I08_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *RPA_I08_PR1v2) AUT() *AUT {
	return m.aut
}

func (m *RPA_I08_PR1v2) CTD() *CTD {
	return m.ctd
}


func (m RPA_I08_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"aut": m.aut,
		"ctd": m.ctd,
	}, nil
}
type RPA_I08_OBRv2 struct {
	obr *OBR // Required
	nte []*NTE
	obx []*RPA_I08_OBR_OBXv2
}

func (m *RPA_I08_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *RPA_I08_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RPA_I08_OBRv2) GroupByOBX() []*RPA_I08_OBR_OBXv2 {
	return m.obx
}


func (m RPA_I08_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type RPA_I08_OBR_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RPA_I08_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RPA_I08_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RPA_I08_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RPI_I01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	prd []*RPI_I01_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*RPI_I01_IN1v2 // Required
	nte []*NTE
}

func (m *RPI_I01v2) MSH() *MSH {
	return m.msh
}

func (m *RPI_I01v2) MSA() *MSA {
	return m.msa
}

func (m *RPI_I01v2) GroupByPRD() []*RPI_I01_PRDv2 {
	return m.prd
}

func (m *RPI_I01v2) PID() *PID {
	return m.pid
}

func (m *RPI_I01v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RPI_I01v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RPI_I01v2) GroupByIN1() []*RPI_I01_IN1v2 {
	return m.in1
}

func (m *RPI_I01v2) AllNTE() []*NTE {
	return m.nte
}


func (m RPI_I01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"nte": m.nte,
	}, nil
}
type RPI_I01_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RPI_I01_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RPI_I01_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RPI_I01_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RPI_I01_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RPI_I01_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RPI_I01_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RPI_I01_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RPI_I01_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RPL_I02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	prd []*RPL_I02_PRDv2 // Required
	nte []*NTE
	dsp []*DSP
	dsc *DSC
}

func (m *RPL_I02v2) MSH() *MSH {
	return m.msh
}

func (m *RPL_I02v2) MSA() *MSA {
	return m.msa
}

func (m *RPL_I02v2) GroupByPRD() []*RPL_I02_PRDv2 {
	return m.prd
}

func (m *RPL_I02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RPL_I02v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *RPL_I02v2) DSC() *DSC {
	return m.dsc
}


func (m RPL_I02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"prd": m.prd,
		"nte": m.nte,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type RPL_I02_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RPL_I02_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RPL_I02_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RPL_I02_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQA_I08v2 struct {
	msh *MSH // Required
	rf1 *RF1
	aut *AUT // Required
	ctd *CTD
	prd []*RQA_I08_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*RQA_I08_IN1v2 // Required
	acc *ACC
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	pr1 []*RQA_I08_PR1v2
	obr []*RQA_I08_OBRv2
	pv1 *PV1 // Required
	pv2 *PV2
	nte []*NTE
}

func (m *RQA_I08v2) MSH() *MSH {
	return m.msh
}

func (m *RQA_I08v2) RF1() *RF1 {
	return m.rf1
}

func (m *RQA_I08v2) AUT() *AUT {
	return m.aut
}

func (m *RQA_I08v2) CTD() *CTD {
	return m.ctd
}

func (m *RQA_I08v2) GroupByPRD() []*RQA_I08_PRDv2 {
	return m.prd
}

func (m *RQA_I08v2) PID() *PID {
	return m.pid
}

func (m *RQA_I08v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RQA_I08v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RQA_I08v2) GroupByIN1() []*RQA_I08_IN1v2 {
	return m.in1
}

func (m *RQA_I08v2) ACC() *ACC {
	return m.acc
}

func (m *RQA_I08v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *RQA_I08v2) AllDRG() []*DRG {
	return m.drg
}

func (m *RQA_I08v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RQA_I08v2) GroupByPR1() []*RQA_I08_PR1v2 {
	return m.pr1
}

func (m *RQA_I08v2) GroupByOBR() []*RQA_I08_OBRv2 {
	return m.obr
}

func (m *RQA_I08v2) PV1() *PV1 {
	return m.pv1
}

func (m *RQA_I08v2) PV2() *PV2 {
	return m.pv2
}

func (m *RQA_I08v2) AllNTE() []*NTE {
	return m.nte
}


func (m RQA_I08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"rf1": m.rf1,
		"aut": m.aut,
		"ctd": m.ctd,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"pr1": m.pr1,
		"obr": m.obr,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"nte": m.nte,
	}, nil
}
type RQA_I08_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RQA_I08_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RQA_I08_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RQA_I08_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQA_I08_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RQA_I08_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RQA_I08_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RQA_I08_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RQA_I08_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RQA_I08_PR1v2 struct {
	pr1 *PR1 // Required
	aut *AUT // Required
	ctd *CTD
}

func (m *RQA_I08_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *RQA_I08_PR1v2) AUT() *AUT {
	return m.aut
}

func (m *RQA_I08_PR1v2) CTD() *CTD {
	return m.ctd
}


func (m RQA_I08_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"aut": m.aut,
		"ctd": m.ctd,
	}, nil
}
type RQA_I08_OBRv2 struct {
	obr *OBR // Required
	nte []*NTE
	obx []*RQA_I08_OBR_OBXv2
}

func (m *RQA_I08_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *RQA_I08_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RQA_I08_OBRv2) GroupByOBX() []*RQA_I08_OBR_OBXv2 {
	return m.obx
}


func (m RQA_I08_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type RQA_I08_OBR_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RQA_I08_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RQA_I08_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RQA_I08_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RQC_I05v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	prd []*RQC_I05_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	nte []*NTE
}

func (m *RQC_I05v2) MSH() *MSH {
	return m.msh
}

func (m *RQC_I05v2) QRD() *QRD {
	return m.qrd
}

func (m *RQC_I05v2) QRF() *QRF {
	return m.qrf
}

func (m *RQC_I05v2) GroupByPRD() []*RQC_I05_PRDv2 {
	return m.prd
}

func (m *RQC_I05v2) PID() *PID {
	return m.pid
}

func (m *RQC_I05v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RQC_I05v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RQC_I05v2) AllNTE() []*NTE {
	return m.nte
}


func (m RQC_I05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"nte": m.nte,
	}, nil
}
type RQC_I05_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RQC_I05_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RQC_I05_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RQC_I05_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQC_I06v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	prd []*RQC_I06_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 *GT1
	nte []*NTE
}

func (m *RQC_I06v2) MSH() *MSH {
	return m.msh
}

func (m *RQC_I06v2) QRD() *QRD {
	return m.qrd
}

func (m *RQC_I06v2) QRF() *QRF {
	return m.qrf
}

func (m *RQC_I06v2) GroupByPRD() []*RQC_I06_PRDv2 {
	return m.prd
}

func (m *RQC_I06v2) PID() *PID {
	return m.pid
}

func (m *RQC_I06v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RQC_I06v2) GT1() *GT1 {
	return m.gt1
}

func (m *RQC_I06v2) AllNTE() []*NTE {
	return m.nte
}


func (m RQC_I06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"nte": m.nte,
	}, nil
}
type RQC_I06_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RQC_I06_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RQC_I06_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RQC_I06_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQI_I01v2 struct {
	msh *MSH // Required
	prd []*RQI_I01_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	in1 []*RQI_I01_IN1v2 // Required
	nte []*NTE
}

func (m *RQI_I01v2) MSH() *MSH {
	return m.msh
}

func (m *RQI_I01v2) GroupByPRD() []*RQI_I01_PRDv2 {
	return m.prd
}

func (m *RQI_I01v2) PID() *PID {
	return m.pid
}

func (m *RQI_I01v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RQI_I01v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RQI_I01v2) GroupByIN1() []*RQI_I01_IN1v2 {
	return m.in1
}

func (m *RQI_I01v2) AllNTE() []*NTE {
	return m.nte
}


func (m RQI_I01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"in1": m.in1,
		"nte": m.nte,
	}, nil
}
type RQI_I01_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RQI_I01_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RQI_I01_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RQI_I01_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQI_I01_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *RQI_I01_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *RQI_I01_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *RQI_I01_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m RQI_I01_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type RQP_I04v2 struct {
	msh *MSH // Required
	prd []*RQP_I04_PRDv2 // Required
	pid *PID // Required
	nk1 []*NK1
	gt1 []*GT1
	nte []*NTE
}

func (m *RQP_I04v2) MSH() *MSH {
	return m.msh
}

func (m *RQP_I04v2) GroupByPRD() []*RQP_I04_PRDv2 {
	return m.prd
}

func (m *RQP_I04v2) PID() *PID {
	return m.pid
}

func (m *RQP_I04v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *RQP_I04v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *RQP_I04v2) AllNTE() []*NTE {
	return m.nte
}


func (m RQP_I04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"prd": m.prd,
		"pid": m.pid,
		"nk1": m.nk1,
		"gt1": m.gt1,
		"nte": m.nte,
	}, nil
}
type RQP_I04_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RQP_I04_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RQP_I04_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RQP_I04_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RQQ_Q01v2 struct {
	msh *MSH // Required
	erq *ERQ // Required
	dsc *DSC
}

func (m *RQQ_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *RQQ_Q01v2) ERQ() *ERQ {
	return m.erq
}

func (m *RQQ_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m RQQ_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"erq": m.erq,
		"dsc": m.dsc,
	}, nil
}
type RRA_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *RRA_O02_PIDv2
}

func (m *RRA_O02v2) MSH() *MSH {
	return m.msh
}

func (m *RRA_O02v2) MSA() *MSA {
	return m.msa
}

func (m *RRA_O02v2) ERR() *ERR {
	return m.err
}

func (m *RRA_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRA_O02v2) GroupByPID() *RRA_O02_PIDv2 {
	return m.pid
}


func (m RRA_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type RRA_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*RRA_O02_PID_ORCv2 // Required
}

func (m *RRA_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *RRA_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRA_O02_PIDv2) GroupByORC() []*RRA_O02_PID_ORCv2 {
	return m.orc
}


func (m RRA_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RRA_O02_PID_ORCv2 struct {
	orc *ORC // Required
	rxa []*RRA_O02_PID_ORC_RXAv2
}

func (m *RRA_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RRA_O02_PID_ORCv2) GroupByRXA() []*RRA_O02_PID_ORC_RXAv2 {
	return m.rxa
}


func (m RRA_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxa": m.rxa,
	}, nil
}
type RRA_O02_PID_ORC_RXAv2 struct {
	rxa *RXA // Required
	rxr *RXR // Required
}

func (m *RRA_O02_PID_ORC_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *RRA_O02_PID_ORC_RXAv2) RXR() *RXR {
	return m.rxr
}


func (m RRA_O02_PID_ORC_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rxa": m.rxa,
		"rxr": m.rxr,
	}, nil
}
type RRD_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *RRD_O02_PIDv2
}

func (m *RRD_O02v2) MSH() *MSH {
	return m.msh
}

func (m *RRD_O02v2) MSA() *MSA {
	return m.msa
}

func (m *RRD_O02v2) ERR() *ERR {
	return m.err
}

func (m *RRD_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRD_O02v2) GroupByPID() *RRD_O02_PIDv2 {
	return m.pid
}


func (m RRD_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type RRD_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*RRD_O02_PID_ORCv2 // Required
}

func (m *RRD_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *RRD_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRD_O02_PIDv2) GroupByORC() []*RRD_O02_PID_ORCv2 {
	return m.orc
}


func (m RRD_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RRD_O02_PID_ORCv2 struct {
	orc *ORC // Required
	rxd *RXD // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RRD_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RRD_O02_PID_ORCv2) RXD() *RXD {
	return m.rxd
}

func (m *RRD_O02_PID_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RRD_O02_PID_ORCv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RRD_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxd": m.rxd,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RRG_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *RRG_O02_PIDv2
}

func (m *RRG_O02v2) MSH() *MSH {
	return m.msh
}

func (m *RRG_O02v2) MSA() *MSA {
	return m.msa
}

func (m *RRG_O02v2) ERR() *ERR {
	return m.err
}

func (m *RRG_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRG_O02v2) GroupByPID() *RRG_O02_PIDv2 {
	return m.pid
}


func (m RRG_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type RRG_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*RRG_O02_PID_ORCv2 // Required
}

func (m *RRG_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *RRG_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRG_O02_PIDv2) GroupByORC() []*RRG_O02_PID_ORCv2 {
	return m.orc
}


func (m RRG_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RRG_O02_PID_ORCv2 struct {
	orc *ORC // Required
	rxg *RXG // Required
	rxr []*RXR // Required
	rxc []*RXC
}

func (m *RRG_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RRG_O02_PID_ORCv2) RXG() *RXG {
	return m.rxg
}

func (m *RRG_O02_PID_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RRG_O02_PID_ORCv2) AllRXC() []*RXC {
	return m.rxc
}


func (m RRG_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxg": m.rxg,
		"rxr": m.rxr,
		"rxc": m.rxc,
	}, nil
}
type RRI_I12v2 struct {
	msh *MSH // Required
	msa *MSA
	rf1 *RF1
	aut *AUT // Required
	ctd *CTD
	prd []*RRI_I12_PRDv2 // Required
	pid *PID // Required
	acc *ACC
	dg1 []*DG1
	drg []*DRG
	al1 []*AL1
	pr1 []*RRI_I12_PR1v2
	obr []*RRI_I12_OBRv2
	pv1 *PV1 // Required
	pv2 *PV2
	nte []*NTE
}

func (m *RRI_I12v2) MSH() *MSH {
	return m.msh
}

func (m *RRI_I12v2) MSA() *MSA {
	return m.msa
}

func (m *RRI_I12v2) RF1() *RF1 {
	return m.rf1
}

func (m *RRI_I12v2) AUT() *AUT {
	return m.aut
}

func (m *RRI_I12v2) CTD() *CTD {
	return m.ctd
}

func (m *RRI_I12v2) GroupByPRD() []*RRI_I12_PRDv2 {
	return m.prd
}

func (m *RRI_I12v2) PID() *PID {
	return m.pid
}

func (m *RRI_I12v2) ACC() *ACC {
	return m.acc
}

func (m *RRI_I12v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *RRI_I12v2) AllDRG() []*DRG {
	return m.drg
}

func (m *RRI_I12v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *RRI_I12v2) GroupByPR1() []*RRI_I12_PR1v2 {
	return m.pr1
}

func (m *RRI_I12v2) GroupByOBR() []*RRI_I12_OBRv2 {
	return m.obr
}

func (m *RRI_I12v2) PV1() *PV1 {
	return m.pv1
}

func (m *RRI_I12v2) PV2() *PV2 {
	return m.pv2
}

func (m *RRI_I12v2) AllNTE() []*NTE {
	return m.nte
}


func (m RRI_I12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"rf1": m.rf1,
		"aut": m.aut,
		"ctd": m.ctd,
		"prd": m.prd,
		"pid": m.pid,
		"acc": m.acc,
		"dg1": m.dg1,
		"drg": m.drg,
		"al1": m.al1,
		"pr1": m.pr1,
		"obr": m.obr,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"nte": m.nte,
	}, nil
}
type RRI_I12_PRDv2 struct {
	prd *PRD // Required
	ctd []*CTD
}

func (m *RRI_I12_PRDv2) PRD() *PRD {
	return m.prd
}

func (m *RRI_I12_PRDv2) AllCTD() []*CTD {
	return m.ctd
}


func (m RRI_I12_PRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"prd": m.prd,
		"ctd": m.ctd,
	}, nil
}
type RRI_I12_PR1v2 struct {
	pr1 *PR1 // Required
	aut *AUT // Required
	ctd *CTD
}

func (m *RRI_I12_PR1v2) PR1() *PR1 {
	return m.pr1
}

func (m *RRI_I12_PR1v2) AUT() *AUT {
	return m.aut
}

func (m *RRI_I12_PR1v2) CTD() *CTD {
	return m.ctd
}


func (m RRI_I12_PR1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pr1": m.pr1,
		"aut": m.aut,
		"ctd": m.ctd,
	}, nil
}
type RRI_I12_OBRv2 struct {
	obr *OBR // Required
	nte []*NTE
	obx []*RRI_I12_OBR_OBXv2
}

func (m *RRI_I12_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *RRI_I12_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRI_I12_OBRv2) GroupByOBX() []*RRI_I12_OBR_OBXv2 {
	return m.obx
}


func (m RRI_I12_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}
type RRI_I12_OBR_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *RRI_I12_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *RRI_I12_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m RRI_I12_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type RRO_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	nte []*NTE
	pid *RRO_O02_PIDv2
}

func (m *RRO_O02v2) MSH() *MSH {
	return m.msh
}

func (m *RRO_O02v2) MSA() *MSA {
	return m.msa
}

func (m *RRO_O02v2) ERR() *ERR {
	return m.err
}

func (m *RRO_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRO_O02v2) GroupByPID() *RRO_O02_PIDv2 {
	return m.pid
}


func (m RRO_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}
type RRO_O02_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	orc []*RRO_O02_PID_ORCv2 // Required
}

func (m *RRO_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *RRO_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *RRO_O02_PIDv2) GroupByORC() []*RRO_O02_PID_ORCv2 {
	return m.orc
}


func (m RRO_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}
type RRO_O02_PID_ORCv2 struct {
	orc *ORC // Required
	rxo *RXO // Required
	nte1 []*NTE
	rxr []*RXR // Required
	rxc []*RXC
	nte2 []*NTE
}

func (m *RRO_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *RRO_O02_PID_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *RRO_O02_PID_ORCv2) AllNTE1() []*NTE {
	return m.nte1
}

func (m *RRO_O02_PID_ORCv2) AllRXR() []*RXR {
	return m.rxr
}

func (m *RRO_O02_PID_ORCv2) AllRXC() []*RXC {
	return m.rxc
}

func (m *RRO_O02_PID_ORCv2) AllNTE2() []*NTE {
	return m.nte2
}


func (m RRO_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxo": m.rxo,
		"nte1": m.nte1,
		"rxr": m.rxr,
		"rxc": m.rxc,
		"nte2": m.nte2,
	}, nil
}
type SIU_S12v2 struct {
	msh *MSH // Required
	sch *SCH // Required
	nte []*NTE
	pid []*SIU_S12_PIDv2
	rgs []*SIU_S12_RGSv2 // Required
}

func (m *SIU_S12v2) MSH() *MSH {
	return m.msh
}

func (m *SIU_S12v2) SCH() *SCH {
	return m.sch
}

func (m *SIU_S12v2) AllNTE() []*NTE {
	return m.nte
}

func (m *SIU_S12v2) GroupByPID() []*SIU_S12_PIDv2 {
	return m.pid
}

func (m *SIU_S12v2) GroupByRGS() []*SIU_S12_RGSv2 {
	return m.rgs
}


func (m SIU_S12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"sch": m.sch,
		"nte": m.nte,
		"pid": m.pid,
		"rgs": m.rgs,
	}, nil
}
type SIU_S12_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *SIU_S12_PIDv2) PID() *PID {
	return m.pid
}

func (m *SIU_S12_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *SIU_S12_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *SIU_S12_PIDv2) AllOBX() []*OBX {
	return m.obx
}

func (m *SIU_S12_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}


func (m SIU_S12_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type SIU_S12_RGSv2 struct {
	rgs *RGS // Required
	ais []*SIU_S12_RGS_AISv2
	aig []*SIU_S12_RGS_AIGv2
	ail []*SIU_S12_RGS_AILv2
	aip []*SIU_S12_RGS_AIPv2
}

func (m *SIU_S12_RGSv2) RGS() *RGS {
	return m.rgs
}

func (m *SIU_S12_RGSv2) GroupByAIS() []*SIU_S12_RGS_AISv2 {
	return m.ais
}

func (m *SIU_S12_RGSv2) GroupByAIG() []*SIU_S12_RGS_AIGv2 {
	return m.aig
}

func (m *SIU_S12_RGSv2) GroupByAIL() []*SIU_S12_RGS_AILv2 {
	return m.ail
}

func (m *SIU_S12_RGSv2) GroupByAIP() []*SIU_S12_RGS_AIPv2 {
	return m.aip
}


func (m SIU_S12_RGSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rgs": m.rgs,
		"ais": m.ais,
		"aig": m.aig,
		"ail": m.ail,
		"aip": m.aip,
	}, nil
}
type SIU_S12_RGS_AISv2 struct {
	ais *AIS // Required
	nte []*NTE
}

func (m *SIU_S12_RGS_AISv2) AIS() *AIS {
	return m.ais
}

func (m *SIU_S12_RGS_AISv2) AllNTE() []*NTE {
	return m.nte
}


func (m SIU_S12_RGS_AISv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ais": m.ais,
		"nte": m.nte,
	}, nil
}
type SIU_S12_RGS_AIGv2 struct {
	aig *AIG // Required
	nte []*NTE
}

func (m *SIU_S12_RGS_AIGv2) AIG() *AIG {
	return m.aig
}

func (m *SIU_S12_RGS_AIGv2) AllNTE() []*NTE {
	return m.nte
}


func (m SIU_S12_RGS_AIGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aig": m.aig,
		"nte": m.nte,
	}, nil
}
type SIU_S12_RGS_AILv2 struct {
	ail *AIL // Required
	nte []*NTE
}

func (m *SIU_S12_RGS_AILv2) AIL() *AIL {
	return m.ail
}

func (m *SIU_S12_RGS_AILv2) AllNTE() []*NTE {
	return m.nte
}


func (m SIU_S12_RGS_AILv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ail": m.ail,
		"nte": m.nte,
	}, nil
}
type SIU_S12_RGS_AIPv2 struct {
	aip *AIP // Required
	nte []*NTE
}

func (m *SIU_S12_RGS_AIPv2) AIP() *AIP {
	return m.aip
}

func (m *SIU_S12_RGS_AIPv2) AllNTE() []*NTE {
	return m.nte
}


func (m SIU_S12_RGS_AIPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aip": m.aip,
		"nte": m.nte,
	}, nil
}
type SPQ_Q01v2 struct {
	msh *MSH // Required
	spr *SPR // Required
	rdf *RDF
	dsc *DSC
}

func (m *SPQ_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *SPQ_Q01v2) SPR() *SPR {
	return m.spr
}

func (m *SPQ_Q01v2) RDF() *RDF {
	return m.rdf
}

func (m *SPQ_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m SPQ_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"spr": m.spr,
		"rdf": m.rdf,
		"dsc": m.dsc,
	}, nil
}
type SQM_S25v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	arq *ARQ // Required
	apr *APR
	pid *PID
	rgs []*SQM_S25_RGSv2 // Required
	dsc *DSC
}

func (m *SQM_S25v2) MSH() *MSH {
	return m.msh
}

func (m *SQM_S25v2) QRD() *QRD {
	return m.qrd
}

func (m *SQM_S25v2) QRF() *QRF {
	return m.qrf
}

func (m *SQM_S25v2) ARQ() *ARQ {
	return m.arq
}

func (m *SQM_S25v2) APR() *APR {
	return m.apr
}

func (m *SQM_S25v2) PID() *PID {
	return m.pid
}

func (m *SQM_S25v2) GroupByRGS() []*SQM_S25_RGSv2 {
	return m.rgs
}

func (m *SQM_S25v2) DSC() *DSC {
	return m.dsc
}


func (m SQM_S25v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"arq": m.arq,
		"apr": m.apr,
		"pid": m.pid,
		"rgs": m.rgs,
		"dsc": m.dsc,
	}, nil
}
type SQM_S25_RGSv2 struct {
	rgs *RGS // Required
	ais []*SQM_S25_RGS_AISv2
	aig []*SQM_S25_RGS_AIGv2
	aip []*SQM_S25_RGS_AIPv2
	ail []*SQM_S25_RGS_AILv2
}

func (m *SQM_S25_RGSv2) RGS() *RGS {
	return m.rgs
}

func (m *SQM_S25_RGSv2) GroupByAIS() []*SQM_S25_RGS_AISv2 {
	return m.ais
}

func (m *SQM_S25_RGSv2) GroupByAIG() []*SQM_S25_RGS_AIGv2 {
	return m.aig
}

func (m *SQM_S25_RGSv2) GroupByAIP() []*SQM_S25_RGS_AIPv2 {
	return m.aip
}

func (m *SQM_S25_RGSv2) GroupByAIL() []*SQM_S25_RGS_AILv2 {
	return m.ail
}


func (m SQM_S25_RGSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rgs": m.rgs,
		"ais": m.ais,
		"aig": m.aig,
		"aip": m.aip,
		"ail": m.ail,
	}, nil
}
type SQM_S25_RGS_AISv2 struct {
	ais *AIS // Required
	apr *APR
}

func (m *SQM_S25_RGS_AISv2) AIS() *AIS {
	return m.ais
}

func (m *SQM_S25_RGS_AISv2) APR() *APR {
	return m.apr
}


func (m SQM_S25_RGS_AISv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ais": m.ais,
		"apr": m.apr,
	}, nil
}
type SQM_S25_RGS_AIGv2 struct {
	aig *AIG // Required
	apr *APR
}

func (m *SQM_S25_RGS_AIGv2) AIG() *AIG {
	return m.aig
}

func (m *SQM_S25_RGS_AIGv2) APR() *APR {
	return m.apr
}


func (m SQM_S25_RGS_AIGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aig": m.aig,
		"apr": m.apr,
	}, nil
}
type SQM_S25_RGS_AIPv2 struct {
	aip *AIP // Required
	apr *APR
}

func (m *SQM_S25_RGS_AIPv2) AIP() *AIP {
	return m.aip
}

func (m *SQM_S25_RGS_AIPv2) APR() *APR {
	return m.apr
}


func (m SQM_S25_RGS_AIPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aip": m.aip,
		"apr": m.apr,
	}, nil
}
type SQM_S25_RGS_AILv2 struct {
	ail *AIL // Required
	apr *APR
}

func (m *SQM_S25_RGS_AILv2) AIL() *AIL {
	return m.ail
}

func (m *SQM_S25_RGS_AILv2) APR() *APR {
	return m.apr
}


func (m SQM_S25_RGS_AILv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ail": m.ail,
		"apr": m.apr,
	}, nil
}
type SQR_S25v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK // Required
	sch []*SQR_S25_SCHv2
	dsc *DSC
}

func (m *SQR_S25v2) MSH() *MSH {
	return m.msh
}

func (m *SQR_S25v2) MSA() *MSA {
	return m.msa
}

func (m *SQR_S25v2) ERR() *ERR {
	return m.err
}

func (m *SQR_S25v2) QAK() *QAK {
	return m.qak
}

func (m *SQR_S25v2) GroupBySCH() []*SQR_S25_SCHv2 {
	return m.sch
}

func (m *SQR_S25v2) DSC() *DSC {
	return m.dsc
}


func (m SQR_S25v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"sch": m.sch,
		"dsc": m.dsc,
	}, nil
}
type SQR_S25_SCHv2 struct {
	sch *SCH // Required
	nte []*NTE
	pid *PID // Required
	pv1 *PV1
	pv2 *PV2
	dg1 *DG1
	rgs []*SQR_S25_SCH_RGSv2 // Required
}

func (m *SQR_S25_SCHv2) SCH() *SCH {
	return m.sch
}

func (m *SQR_S25_SCHv2) AllNTE() []*NTE {
	return m.nte
}

func (m *SQR_S25_SCHv2) PID() *PID {
	return m.pid
}

func (m *SQR_S25_SCHv2) PV1() *PV1 {
	return m.pv1
}

func (m *SQR_S25_SCHv2) PV2() *PV2 {
	return m.pv2
}

func (m *SQR_S25_SCHv2) DG1() *DG1 {
	return m.dg1
}

func (m *SQR_S25_SCHv2) GroupByRGS() []*SQR_S25_SCH_RGSv2 {
	return m.rgs
}


func (m SQR_S25_SCHv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"sch": m.sch,
		"nte": m.nte,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"dg1": m.dg1,
		"rgs": m.rgs,
	}, nil
}
type SQR_S25_SCH_RGSv2 struct {
	rgs *RGS // Required
	ais []*SQR_S25_SCH_RGS_AISv2
	aig []*SQR_S25_SCH_RGS_AIGv2
	aip []*SQR_S25_SCH_RGS_AIPv2
	ail []*SQR_S25_SCH_RGS_AILv2
}

func (m *SQR_S25_SCH_RGSv2) RGS() *RGS {
	return m.rgs
}

func (m *SQR_S25_SCH_RGSv2) GroupByAIS() []*SQR_S25_SCH_RGS_AISv2 {
	return m.ais
}

func (m *SQR_S25_SCH_RGSv2) GroupByAIG() []*SQR_S25_SCH_RGS_AIGv2 {
	return m.aig
}

func (m *SQR_S25_SCH_RGSv2) GroupByAIP() []*SQR_S25_SCH_RGS_AIPv2 {
	return m.aip
}

func (m *SQR_S25_SCH_RGSv2) GroupByAIL() []*SQR_S25_SCH_RGS_AILv2 {
	return m.ail
}


func (m SQR_S25_SCH_RGSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rgs": m.rgs,
		"ais": m.ais,
		"aig": m.aig,
		"aip": m.aip,
		"ail": m.ail,
	}, nil
}
type SQR_S25_SCH_RGS_AISv2 struct {
	ais *AIS // Required
	nte []*NTE
}

func (m *SQR_S25_SCH_RGS_AISv2) AIS() *AIS {
	return m.ais
}

func (m *SQR_S25_SCH_RGS_AISv2) AllNTE() []*NTE {
	return m.nte
}


func (m SQR_S25_SCH_RGS_AISv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ais": m.ais,
		"nte": m.nte,
	}, nil
}
type SQR_S25_SCH_RGS_AIGv2 struct {
	aig *AIG // Required
	nte []*NTE
}

func (m *SQR_S25_SCH_RGS_AIGv2) AIG() *AIG {
	return m.aig
}

func (m *SQR_S25_SCH_RGS_AIGv2) AllNTE() []*NTE {
	return m.nte
}


func (m SQR_S25_SCH_RGS_AIGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aig": m.aig,
		"nte": m.nte,
	}, nil
}
type SQR_S25_SCH_RGS_AIPv2 struct {
	aip *AIP // Required
	nte []*NTE
}

func (m *SQR_S25_SCH_RGS_AIPv2) AIP() *AIP {
	return m.aip
}

func (m *SQR_S25_SCH_RGS_AIPv2) AllNTE() []*NTE {
	return m.nte
}


func (m SQR_S25_SCH_RGS_AIPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aip": m.aip,
		"nte": m.nte,
	}, nil
}
type SQR_S25_SCH_RGS_AILv2 struct {
	ail *AIL // Required
	nte []*NTE
}

func (m *SQR_S25_SCH_RGS_AILv2) AIL() *AIL {
	return m.ail
}

func (m *SQR_S25_SCH_RGS_AILv2) AllNTE() []*NTE {
	return m.nte
}


func (m SQR_S25_SCH_RGS_AILv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ail": m.ail,
		"nte": m.nte,
	}, nil
}
type SRM_S01v2 struct {
	msh *MSH // Required
	arq *ARQ // Required
	apr *APR
	nte []*NTE
	pid []*SRM_S01_PIDv2
	rgs []*SRM_S01_RGSv2 // Required
}

func (m *SRM_S01v2) MSH() *MSH {
	return m.msh
}

func (m *SRM_S01v2) ARQ() *ARQ {
	return m.arq
}

func (m *SRM_S01v2) APR() *APR {
	return m.apr
}

func (m *SRM_S01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *SRM_S01v2) GroupByPID() []*SRM_S01_PIDv2 {
	return m.pid
}

func (m *SRM_S01v2) GroupByRGS() []*SRM_S01_RGSv2 {
	return m.rgs
}


func (m SRM_S01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"arq": m.arq,
		"apr": m.apr,
		"nte": m.nte,
		"pid": m.pid,
		"rgs": m.rgs,
	}, nil
}
type SRM_S01_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *SRM_S01_PIDv2) PID() *PID {
	return m.pid
}

func (m *SRM_S01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *SRM_S01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *SRM_S01_PIDv2) AllOBX() []*OBX {
	return m.obx
}

func (m *SRM_S01_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}


func (m SRM_S01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}
type SRM_S01_RGSv2 struct {
	rgs *RGS // Required
	ais []*SRM_S01_RGS_AISv2
	aig []*SRM_S01_RGS_AIGv2
	ail []*SRM_S01_RGS_AILv2
	aip []*SRM_S01_RGS_AIPv2
}

func (m *SRM_S01_RGSv2) RGS() *RGS {
	return m.rgs
}

func (m *SRM_S01_RGSv2) GroupByAIS() []*SRM_S01_RGS_AISv2 {
	return m.ais
}

func (m *SRM_S01_RGSv2) GroupByAIG() []*SRM_S01_RGS_AIGv2 {
	return m.aig
}

func (m *SRM_S01_RGSv2) GroupByAIL() []*SRM_S01_RGS_AILv2 {
	return m.ail
}

func (m *SRM_S01_RGSv2) GroupByAIP() []*SRM_S01_RGS_AIPv2 {
	return m.aip
}


func (m SRM_S01_RGSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rgs": m.rgs,
		"ais": m.ais,
		"aig": m.aig,
		"ail": m.ail,
		"aip": m.aip,
	}, nil
}
type SRM_S01_RGS_AISv2 struct {
	ais *AIS // Required
	apr *APR
	nte []*NTE
}

func (m *SRM_S01_RGS_AISv2) AIS() *AIS {
	return m.ais
}

func (m *SRM_S01_RGS_AISv2) APR() *APR {
	return m.apr
}

func (m *SRM_S01_RGS_AISv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRM_S01_RGS_AISv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ais": m.ais,
		"apr": m.apr,
		"nte": m.nte,
	}, nil
}
type SRM_S01_RGS_AIGv2 struct {
	aig *AIG // Required
	apr *APR
	nte []*NTE
}

func (m *SRM_S01_RGS_AIGv2) AIG() *AIG {
	return m.aig
}

func (m *SRM_S01_RGS_AIGv2) APR() *APR {
	return m.apr
}

func (m *SRM_S01_RGS_AIGv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRM_S01_RGS_AIGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aig": m.aig,
		"apr": m.apr,
		"nte": m.nte,
	}, nil
}
type SRM_S01_RGS_AILv2 struct {
	ail *AIL // Required
	apr *APR
	nte []*NTE
}

func (m *SRM_S01_RGS_AILv2) AIL() *AIL {
	return m.ail
}

func (m *SRM_S01_RGS_AILv2) APR() *APR {
	return m.apr
}

func (m *SRM_S01_RGS_AILv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRM_S01_RGS_AILv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ail": m.ail,
		"apr": m.apr,
		"nte": m.nte,
	}, nil
}
type SRM_S01_RGS_AIPv2 struct {
	aip *AIP // Required
	apr *APR
	nte []*NTE
}

func (m *SRM_S01_RGS_AIPv2) AIP() *AIP {
	return m.aip
}

func (m *SRM_S01_RGS_AIPv2) APR() *APR {
	return m.apr
}

func (m *SRM_S01_RGS_AIPv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRM_S01_RGS_AIPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aip": m.aip,
		"apr": m.apr,
		"nte": m.nte,
	}, nil
}
type SRR_S01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	sch *SCH // Required
	nte []*NTE
	pid []*SRR_S01_PIDv2
	rgs []*SRR_S01_RGSv2 // Required
}

func (m *SRR_S01v2) MSH() *MSH {
	return m.msh
}

func (m *SRR_S01v2) MSA() *MSA {
	return m.msa
}

func (m *SRR_S01v2) ERR() *ERR {
	return m.err
}

func (m *SRR_S01v2) SCH() *SCH {
	return m.sch
}

func (m *SRR_S01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *SRR_S01v2) GroupByPID() []*SRR_S01_PIDv2 {
	return m.pid
}

func (m *SRR_S01v2) GroupByRGS() []*SRR_S01_RGSv2 {
	return m.rgs
}


func (m SRR_S01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"sch": m.sch,
		"nte": m.nte,
		"pid": m.pid,
		"rgs": m.rgs,
	}, nil
}
type SRR_S01_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
	pv2 *PV2
	dg1 []*DG1
}

func (m *SRR_S01_PIDv2) PID() *PID {
	return m.pid
}

func (m *SRR_S01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *SRR_S01_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *SRR_S01_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}


func (m SRR_S01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"dg1": m.dg1,
	}, nil
}
type SRR_S01_RGSv2 struct {
	rgs *RGS // Required
	ais []*SRR_S01_RGS_AISv2
	aig []*SRR_S01_RGS_AIGv2
	ail []*SRR_S01_RGS_AILv2
	aip []*SRR_S01_RGS_AIPv2
}

func (m *SRR_S01_RGSv2) RGS() *RGS {
	return m.rgs
}

func (m *SRR_S01_RGSv2) GroupByAIS() []*SRR_S01_RGS_AISv2 {
	return m.ais
}

func (m *SRR_S01_RGSv2) GroupByAIG() []*SRR_S01_RGS_AIGv2 {
	return m.aig
}

func (m *SRR_S01_RGSv2) GroupByAIL() []*SRR_S01_RGS_AILv2 {
	return m.ail
}

func (m *SRR_S01_RGSv2) GroupByAIP() []*SRR_S01_RGS_AIPv2 {
	return m.aip
}


func (m SRR_S01_RGSv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"rgs": m.rgs,
		"ais": m.ais,
		"aig": m.aig,
		"ail": m.ail,
		"aip": m.aip,
	}, nil
}
type SRR_S01_RGS_AISv2 struct {
	ais *AIS // Required
	nte []*NTE
}

func (m *SRR_S01_RGS_AISv2) AIS() *AIS {
	return m.ais
}

func (m *SRR_S01_RGS_AISv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRR_S01_RGS_AISv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ais": m.ais,
		"nte": m.nte,
	}, nil
}
type SRR_S01_RGS_AIGv2 struct {
	aig *AIG // Required
	nte []*NTE
}

func (m *SRR_S01_RGS_AIGv2) AIG() *AIG {
	return m.aig
}

func (m *SRR_S01_RGS_AIGv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRR_S01_RGS_AIGv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aig": m.aig,
		"nte": m.nte,
	}, nil
}
type SRR_S01_RGS_AILv2 struct {
	ail *AIL // Required
	nte []*NTE
}

func (m *SRR_S01_RGS_AILv2) AIL() *AIL {
	return m.ail
}

func (m *SRR_S01_RGS_AILv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRR_S01_RGS_AILv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"ail": m.ail,
		"nte": m.nte,
	}, nil
}
type SRR_S01_RGS_AIPv2 struct {
	aip *AIP // Required
	nte []*NTE
}

func (m *SRR_S01_RGS_AIPv2) AIP() *AIP {
	return m.aip
}

func (m *SRR_S01_RGS_AIPv2) AllNTE() []*NTE {
	return m.nte
}


func (m SRR_S01_RGS_AIPv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"aip": m.aip,
		"nte": m.nte,
	}, nil
}
type SUR_P09v2 struct {
	msh *MSH // Required
	fac []*SUR_P09_FACv2 // Required
}

func (m *SUR_P09v2) MSH() *MSH {
	return m.msh
}

func (m *SUR_P09v2) GroupByFAC() []*SUR_P09_FACv2 {
	return m.fac
}


func (m SUR_P09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"fac": m.fac,
	}, nil
}
type SUR_P09_FACv2 struct {
	fac1 *FAC // Required
	psh1 []*SUR_P09_FAC_PSH1v2 // Required
	psh2 *PSH // Required
	fac2 []*SUR_P09_FAC_FAC2v2 // Required
}

func (m *SUR_P09_FACv2) FAC1() *FAC {
	return m.fac1
}

func (m *SUR_P09_FACv2) GroupByPSH1() []*SUR_P09_FAC_PSH1v2 {
	return m.psh1
}

func (m *SUR_P09_FACv2) PSH2() *PSH {
	return m.psh2
}

func (m *SUR_P09_FACv2) GroupByFAC2() []*SUR_P09_FAC_FAC2v2 {
	return m.fac2
}


func (m SUR_P09_FACv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"fac1": m.fac1,
		"psh1": m.psh1,
		"psh2": m.psh2,
		"fac2": m.fac2,
	}, nil
}
type SUR_P09_FAC_PSH1v2 struct {
	psh *PSH // Required
	pdc *PDC // Required
}

func (m *SUR_P09_FAC_PSH1v2) PSH() *PSH {
	return m.psh
}

func (m *SUR_P09_FAC_PSH1v2) PDC() *PDC {
	return m.pdc
}


func (m SUR_P09_FAC_PSH1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"psh": m.psh,
		"pdc": m.pdc,
	}, nil
}
type SUR_P09_FAC_FAC2v2 struct {
	fac *FAC // Required
	pdc *PDC // Required
	nte *NTE // Required
}

func (m *SUR_P09_FAC_FAC2v2) FAC() *FAC {
	return m.fac
}

func (m *SUR_P09_FAC_FAC2v2) PDC() *PDC {
	return m.pdc
}

func (m *SUR_P09_FAC_FAC2v2) NTE() *NTE {
	return m.nte
}


func (m SUR_P09_FAC_FAC2v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"fac": m.fac,
		"pdc": m.pdc,
		"nte": m.nte,
	}, nil
}
type TBR_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qak *QAK // Required
	rdf *RDF // Required
	rdt []*RDT // Required
	dsc *DSC
}

func (m *TBR_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *TBR_Q01v2) MSA() *MSA {
	return m.msa
}

func (m *TBR_Q01v2) ERR() *ERR {
	return m.err
}

func (m *TBR_Q01v2) QAK() *QAK {
	return m.qak
}

func (m *TBR_Q01v2) RDF() *RDF {
	return m.rdf
}

func (m *TBR_Q01v2) AllRDT() []*RDT {
	return m.rdt
}

func (m *TBR_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m TBR_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qak": m.qak,
		"rdf": m.rdf,
		"rdt": m.rdt,
		"dsc": m.dsc,
	}, nil
}
type UDM_Q05v2 struct {
	msh *MSH // Required
	urd *URD // Required
	urs *URS
	dsp []*DSP // Required
	dsc *DSC
}

func (m *UDM_Q05v2) MSH() *MSH {
	return m.msh
}

func (m *UDM_Q05v2) URD() *URD {
	return m.urd
}

func (m *UDM_Q05v2) URS() *URS {
	return m.urs
}

func (m *UDM_Q05v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *UDM_Q05v2) DSC() *DSC {
	return m.dsc
}


func (m UDM_Q05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"urd": m.urd,
		"urs": m.urs,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}
type VQQ_Q01v2 struct {
	msh *MSH // Required
	vtq *VTQ // Required
	rdf *RDF
	dsc *DSC
}

func (m *VQQ_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *VQQ_Q01v2) VTQ() *VTQ {
	return m.vtq
}

func (m *VQQ_Q01v2) RDF() *RDF {
	return m.rdf
}

func (m *VQQ_Q01v2) DSC() *DSC {
	return m.dsc
}


func (m VQQ_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"vtq": m.vtq,
		"rdf": m.rdf,
		"dsc": m.dsc,
	}, nil
}
type VXQ_V01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
}

func (m *VXQ_V01v2) MSH() *MSH {
	return m.msh
}

func (m *VXQ_V01v2) QRD() *QRD {
	return m.qrd
}

func (m *VXQ_V01v2) QRF() *QRF {
	return m.qrf
}


func (m VXQ_V01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}
type VXR_V03v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*VXR_V03_IN1v2
	rxa []*VXR_V03_RXAv2
}

func (m *VXR_V03v2) MSH() *MSH {
	return m.msh
}

func (m *VXR_V03v2) MSA() *MSA {
	return m.msa
}

func (m *VXR_V03v2) QRD() *QRD {
	return m.qrd
}

func (m *VXR_V03v2) QRF() *QRF {
	return m.qrf
}

func (m *VXR_V03v2) PID() *PID {
	return m.pid
}

func (m *VXR_V03v2) PD1() *PD1 {
	return m.pd1
}

func (m *VXR_V03v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *VXR_V03v2) PV1() *PV1 {
	return m.pv1
}

func (m *VXR_V03v2) PV2() *PV2 {
	return m.pv2
}

func (m *VXR_V03v2) GroupByIN1() []*VXR_V03_IN1v2 {
	return m.in1
}

func (m *VXR_V03v2) GroupByRXA() []*VXR_V03_RXAv2 {
	return m.rxa
}


func (m VXR_V03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"rxa": m.rxa,
	}, nil
}
type VXR_V03_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *VXR_V03_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *VXR_V03_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *VXR_V03_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m VXR_V03_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type VXR_V03_RXAv2 struct {
	orc *ORC
	rxa *RXA // Required
	rxr *RXR
	obx []*VXR_V03_RXA_OBXv2
}

func (m *VXR_V03_RXAv2) ORC() *ORC {
	return m.orc
}

func (m *VXR_V03_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *VXR_V03_RXAv2) RXR() *RXR {
	return m.rxr
}

func (m *VXR_V03_RXAv2) GroupByOBX() []*VXR_V03_RXA_OBXv2 {
	return m.obx
}


func (m VXR_V03_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxa": m.rxa,
		"rxr": m.rxr,
		"obx": m.obx,
	}, nil
}
type VXR_V03_RXA_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *VXR_V03_RXA_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *VXR_V03_RXA_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m VXR_V03_RXA_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type VXU_V04v2 struct {
	msh *MSH // Required
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	in1 []*VXU_V04_IN1v2
	rxa []*VXU_V04_RXAv2
}

func (m *VXU_V04v2) MSH() *MSH {
	return m.msh
}

func (m *VXU_V04v2) PID() *PID {
	return m.pid
}

func (m *VXU_V04v2) PD1() *PD1 {
	return m.pd1
}

func (m *VXU_V04v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *VXU_V04v2) PV1() *PV1 {
	return m.pv1
}

func (m *VXU_V04v2) PV2() *PV2 {
	return m.pv2
}

func (m *VXU_V04v2) GroupByIN1() []*VXU_V04_IN1v2 {
	return m.in1
}

func (m *VXU_V04v2) GroupByRXA() []*VXU_V04_RXAv2 {
	return m.rxa
}


func (m VXU_V04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"in1": m.in1,
		"rxa": m.rxa,
	}, nil
}
type VXU_V04_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *VXU_V04_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *VXU_V04_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *VXU_V04_IN1v2) IN3() *IN3 {
	return m.in3
}


func (m VXU_V04_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}
type VXU_V04_RXAv2 struct {
	orc *ORC
	rxa *RXA // Required
	rxr *RXR
	obx []*VXU_V04_RXA_OBXv2
}

func (m *VXU_V04_RXAv2) ORC() *ORC {
	return m.orc
}

func (m *VXU_V04_RXAv2) RXA() *RXA {
	return m.rxa
}

func (m *VXU_V04_RXAv2) RXR() *RXR {
	return m.rxr
}

func (m *VXU_V04_RXAv2) GroupByOBX() []*VXU_V04_RXA_OBXv2 {
	return m.obx
}


func (m VXU_V04_RXAv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"rxa": m.rxa,
		"rxr": m.rxr,
		"obx": m.obx,
	}, nil
}
type VXU_V04_RXA_OBXv2 struct {
	obx *OBX // Required
	nte []*NTE
}

func (m *VXU_V04_RXA_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *VXU_V04_RXA_OBXv2) AllNTE() []*NTE {
	return m.nte
}


func (m VXU_V04_RXA_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}
type VXX_V02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	pid []*VXX_V02_PIDv2 // Required
}

func (m *VXX_V02v2) MSH() *MSH {
	return m.msh
}

func (m *VXX_V02v2) MSA() *MSA {
	return m.msa
}

func (m *VXX_V02v2) QRD() *QRD {
	return m.qrd
}

func (m *VXX_V02v2) QRF() *QRF {
	return m.qrf
}

func (m *VXX_V02v2) GroupByPID() []*VXX_V02_PIDv2 {
	return m.pid
}


func (m VXX_V02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
	}, nil
}
type VXX_V02_PIDv2 struct {
	pid *PID // Required
	nk1 []*NK1
}

func (m *VXX_V02_PIDv2) PID() *PID {
	return m.pid
}

func (m *VXX_V02_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}


func (m VXX_V02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nk1": m.nk1,
	}, nil
}

// Types maps the name of an HL7 segment or message type to the type of the struct that
// represents that segment or message type.
var Types = map[string]reflect.Type{}
func init() {
	Types["ACC"] = reflect.TypeOf(ACC{})
	Types["ADD"] = reflect.TypeOf(ADD{})
	Types["AIG"] = reflect.TypeOf(AIG{})
	Types["AIL"] = reflect.TypeOf(AIL{})
	Types["AIP"] = reflect.TypeOf(AIP{})
	Types["AIS"] = reflect.TypeOf(AIS{})
	Types["AL1"] = reflect.TypeOf(AL1{})
	Types["APR"] = reflect.TypeOf(APR{})
	Types["ARQ"] = reflect.TypeOf(ARQ{})
	Types["AUT"] = reflect.TypeOf(AUT{})
	Types["BHS"] = reflect.TypeOf(BHS{})
	Types["BLG"] = reflect.TypeOf(BLG{})
	Types["BTS"] = reflect.TypeOf(BTS{})
	Types["CDM"] = reflect.TypeOf(CDM{})
	Types["CM0"] = reflect.TypeOf(CM0{})
	Types["CM1"] = reflect.TypeOf(CM1{})
	Types["CM2"] = reflect.TypeOf(CM2{})
	Types["CSP"] = reflect.TypeOf(CSP{})
	Types["CSR"] = reflect.TypeOf(CSR{})
	Types["CSS"] = reflect.TypeOf(CSS{})
	Types["CTD"] = reflect.TypeOf(CTD{})
	Types["CTI"] = reflect.TypeOf(CTI{})
	Types["DB1"] = reflect.TypeOf(DB1{})
	Types["DG1"] = reflect.TypeOf(DG1{})
	Types["DRG"] = reflect.TypeOf(DRG{})
	Types["DSC"] = reflect.TypeOf(DSC{})
	Types["DSP"] = reflect.TypeOf(DSP{})
	Types["EQL"] = reflect.TypeOf(EQL{})
	Types["ERQ"] = reflect.TypeOf(ERQ{})
	Types["ERR"] = reflect.TypeOf(ERR{})
	Types["EVN"] = reflect.TypeOf(EVN{})
	Types["FAC"] = reflect.TypeOf(FAC{})
	Types["FHS"] = reflect.TypeOf(FHS{})
	Types["FT1"] = reflect.TypeOf(FT1{})
	Types["FTS"] = reflect.TypeOf(FTS{})
	Types["GOL"] = reflect.TypeOf(GOL{})
	Types["GT1"] = reflect.TypeOf(GT1{})
	Types["IN1"] = reflect.TypeOf(IN1{})
	Types["IN2"] = reflect.TypeOf(IN2{})
	Types["IN3"] = reflect.TypeOf(IN3{})
	Types["LCC"] = reflect.TypeOf(LCC{})
	Types["LCH"] = reflect.TypeOf(LCH{})
	Types["LDP"] = reflect.TypeOf(LDP{})
	Types["LOC"] = reflect.TypeOf(LOC{})
	Types["LRL"] = reflect.TypeOf(LRL{})
	Types["MFA"] = reflect.TypeOf(MFA{})
	Types["MFE"] = reflect.TypeOf(MFE{})
	Types["MFI"] = reflect.TypeOf(MFI{})
	Types["MRG"] = reflect.TypeOf(MRG{})
	Types["MSA"] = reflect.TypeOf(MSA{})
	Types["MSH"] = reflect.TypeOf(MSH{})
	Types["NCK"] = reflect.TypeOf(NCK{})
	Types["NK1"] = reflect.TypeOf(NK1{})
	Types["NPU"] = reflect.TypeOf(NPU{})
	Types["NSC"] = reflect.TypeOf(NSC{})
	Types["NST"] = reflect.TypeOf(NST{})
	Types["NTE"] = reflect.TypeOf(NTE{})
	Types["OBR"] = reflect.TypeOf(OBR{})
	Types["OBX"] = reflect.TypeOf(OBX{})
	Types["ODS"] = reflect.TypeOf(ODS{})
	Types["ODT"] = reflect.TypeOf(ODT{})
	Types["OM1"] = reflect.TypeOf(OM1{})
	Types["OM2"] = reflect.TypeOf(OM2{})
	Types["OM3"] = reflect.TypeOf(OM3{})
	Types["OM4"] = reflect.TypeOf(OM4{})
	Types["OM5"] = reflect.TypeOf(OM5{})
	Types["OM6"] = reflect.TypeOf(OM6{})
	Types["ORC"] = reflect.TypeOf(ORC{})
	Types["ORO"] = reflect.TypeOf(ORO{})
	Types["PCR"] = reflect.TypeOf(PCR{})
	Types["PD1"] = reflect.TypeOf(PD1{})
	Types["PDC"] = reflect.TypeOf(PDC{})
	Types["PEO"] = reflect.TypeOf(PEO{})
	Types["PES"] = reflect.TypeOf(PES{})
	Types["PID"] = reflect.TypeOf(PID{})
	Types["PR1"] = reflect.TypeOf(PR1{})
	Types["PRA"] = reflect.TypeOf(PRA{})
	Types["PRB"] = reflect.TypeOf(PRB{})
	Types["PRC"] = reflect.TypeOf(PRC{})
	Types["PRD"] = reflect.TypeOf(PRD{})
	Types["PSH"] = reflect.TypeOf(PSH{})
	Types["PTH"] = reflect.TypeOf(PTH{})
	Types["PV1"] = reflect.TypeOf(PV1{})
	Types["PV2"] = reflect.TypeOf(PV2{})
	Types["QAK"] = reflect.TypeOf(QAK{})
	Types["QRD"] = reflect.TypeOf(QRD{})
	Types["QRF"] = reflect.TypeOf(QRF{})
	Types["RDF"] = reflect.TypeOf(RDF{})
	Types["RDT"] = reflect.TypeOf(RDT{})
	Types["RF1"] = reflect.TypeOf(RF1{})
	Types["RGS"] = reflect.TypeOf(RGS{})
	Types["ROL"] = reflect.TypeOf(ROL{})
	Types["RQ1"] = reflect.TypeOf(RQ1{})
	Types["RQD"] = reflect.TypeOf(RQD{})
	Types["RX1"] = reflect.TypeOf(RX1{})
	Types["RXA"] = reflect.TypeOf(RXA{})
	Types["RXC"] = reflect.TypeOf(RXC{})
	Types["RXD"] = reflect.TypeOf(RXD{})
	Types["RXE"] = reflect.TypeOf(RXE{})
	Types["RXG"] = reflect.TypeOf(RXG{})
	Types["RXO"] = reflect.TypeOf(RXO{})
	Types["RXR"] = reflect.TypeOf(RXR{})
	Types["SCH"] = reflect.TypeOf(SCH{})
	Types["SPR"] = reflect.TypeOf(SPR{})
	Types["STF"] = reflect.TypeOf(STF{})
	Types["TXA"] = reflect.TypeOf(TXA{})
	Types["UB1"] = reflect.TypeOf(UB1{})
	Types["UB2"] = reflect.TypeOf(UB2{})
	Types["URD"] = reflect.TypeOf(URD{})
	Types["URS"] = reflect.TypeOf(URS{})
	Types["VAR"] = reflect.TypeOf(VAR{})
	Types["VTQ"] = reflect.TypeOf(VTQ{})
	Types["ACK"] = reflect.TypeOf(ACK{})
	Types["ADR_A19"] = reflect.TypeOf(ADR_A19{})
	Types["ADR_A19_INSURANCE"] = reflect.TypeOf(ADR_A19_INSURANCE{})
	Types["ADR_A19_QUERY_RESPONSE"] = reflect.TypeOf(ADR_A19_QUERY_RESPONSE{})
	Types["ADT_A01"] = reflect.TypeOf(ADT_A01{})
	Types["ADT_A01_INSURANCE"] = reflect.TypeOf(ADT_A01_INSURANCE{})
	Types["ADT_A01_PROCEDURE"] = reflect.TypeOf(ADT_A01_PROCEDURE{})
	Types["ADT_A02"] = reflect.TypeOf(ADT_A02{})
	Types["ADT_A03"] = reflect.TypeOf(ADT_A03{})
	Types["ADT_A03_PROCEDURE"] = reflect.TypeOf(ADT_A03_PROCEDURE{})
	Types["ADT_A04"] = reflect.TypeOf(ADT_A04{})
	Types["ADT_A04_INSURANCE"] = reflect.TypeOf(ADT_A04_INSURANCE{})
	Types["ADT_A05"] = reflect.TypeOf(ADT_A05{})
	Types["ADT_A05_INSURANCE"] = reflect.TypeOf(ADT_A05_INSURANCE{})
	Types["ADT_A06"] = reflect.TypeOf(ADT_A06{})
	Types["ADT_A06_INSURANCE"] = reflect.TypeOf(ADT_A06_INSURANCE{})
	Types["ADT_A06_PROCEDURE"] = reflect.TypeOf(ADT_A06_PROCEDURE{})
	Types["ADT_A07"] = reflect.TypeOf(ADT_A07{})
	Types["ADT_A07_INSURANCE"] = reflect.TypeOf(ADT_A07_INSURANCE{})
	Types["ADT_A08"] = reflect.TypeOf(ADT_A08{})
	Types["ADT_A08_INSURANCE"] = reflect.TypeOf(ADT_A08_INSURANCE{})
	Types["ADT_A09"] = reflect.TypeOf(ADT_A09{})
	Types["ADT_A10"] = reflect.TypeOf(ADT_A10{})
	Types["ADT_A11"] = reflect.TypeOf(ADT_A11{})
	Types["ADT_A12"] = reflect.TypeOf(ADT_A12{})
	Types["ADT_A13"] = reflect.TypeOf(ADT_A13{})
	Types["ADT_A13_INSURANCE"] = reflect.TypeOf(ADT_A13_INSURANCE{})
	Types["ADT_A14"] = reflect.TypeOf(ADT_A14{})
	Types["ADT_A14_INSURANCE"] = reflect.TypeOf(ADT_A14_INSURANCE{})
	Types["ADT_A15"] = reflect.TypeOf(ADT_A15{})
	Types["ADT_A16"] = reflect.TypeOf(ADT_A16{})
	Types["ADT_A17"] = reflect.TypeOf(ADT_A17{})
	Types["ADT_A17_PATIENT"] = reflect.TypeOf(ADT_A17_PATIENT{})
	Types["ADT_A18"] = reflect.TypeOf(ADT_A18{})
	Types["ADT_A20"] = reflect.TypeOf(ADT_A20{})
	Types["ADT_A21"] = reflect.TypeOf(ADT_A21{})
	Types["ADT_A22"] = reflect.TypeOf(ADT_A22{})
	Types["ADT_A23"] = reflect.TypeOf(ADT_A23{})
	Types["ADT_A24"] = reflect.TypeOf(ADT_A24{})
	Types["ADT_A25"] = reflect.TypeOf(ADT_A25{})
	Types["ADT_A26"] = reflect.TypeOf(ADT_A26{})
	Types["ADT_A27"] = reflect.TypeOf(ADT_A27{})
	Types["ADT_A28"] = reflect.TypeOf(ADT_A28{})
	Types["ADT_A28_INSURANCE"] = reflect.TypeOf(ADT_A28_INSURANCE{})
	Types["ADT_A29"] = reflect.TypeOf(ADT_A29{})
	Types["ADT_A30"] = reflect.TypeOf(ADT_A30{})
	Types["ADT_A31"] = reflect.TypeOf(ADT_A31{})
	Types["ADT_A31_INSURANCE"] = reflect.TypeOf(ADT_A31_INSURANCE{})
	Types["ADT_A32"] = reflect.TypeOf(ADT_A32{})
	Types["ADT_A33"] = reflect.TypeOf(ADT_A33{})
	Types["ADT_A34"] = reflect.TypeOf(ADT_A34{})
	Types["ADT_A35"] = reflect.TypeOf(ADT_A35{})
	Types["ADT_A36"] = reflect.TypeOf(ADT_A36{})
	Types["ADT_A37"] = reflect.TypeOf(ADT_A37{})
	Types["ADT_A38"] = reflect.TypeOf(ADT_A38{})
	Types["ADT_A39"] = reflect.TypeOf(ADT_A39{})
	Types["ADT_A39_PATIENT"] = reflect.TypeOf(ADT_A39_PATIENT{})
	Types["ADT_A40"] = reflect.TypeOf(ADT_A40{})
	Types["ADT_A40_PATIENT"] = reflect.TypeOf(ADT_A40_PATIENT{})
	Types["ADT_A43"] = reflect.TypeOf(ADT_A43{})
	Types["ADT_A43_PATIENT"] = reflect.TypeOf(ADT_A43_PATIENT{})
	Types["ADT_A44"] = reflect.TypeOf(ADT_A44{})
	Types["ADT_A44_PATIENT"] = reflect.TypeOf(ADT_A44_PATIENT{})
	Types["ADT_A45"] = reflect.TypeOf(ADT_A45{})
	Types["ADT_A45_MERGE_INFO"] = reflect.TypeOf(ADT_A45_MERGE_INFO{})
	Types["ADT_A50"] = reflect.TypeOf(ADT_A50{})
	Types["ARD_A19"] = reflect.TypeOf(ARD_A19{})
	Types["ARD_A19_INSURANCE"] = reflect.TypeOf(ARD_A19_INSURANCE{})
	Types["ARD_A19_PROCEDURE"] = reflect.TypeOf(ARD_A19_PROCEDURE{})
	Types["ARD_A19_QUERY_RESPONSE"] = reflect.TypeOf(ARD_A19_QUERY_RESPONSE{})
	Types["BAR_P01"] = reflect.TypeOf(BAR_P01{})
	Types["BAR_P01_INSURANCE"] = reflect.TypeOf(BAR_P01_INSURANCE{})
	Types["BAR_P01_PROCEDURE"] = reflect.TypeOf(BAR_P01_PROCEDURE{})
	Types["BAR_P01_VISIT"] = reflect.TypeOf(BAR_P01_VISIT{})
	Types["BAR_P02"] = reflect.TypeOf(BAR_P02{})
	Types["BAR_P02_PATIENT"] = reflect.TypeOf(BAR_P02_PATIENT{})
	Types["BAR_P06"] = reflect.TypeOf(BAR_P06{})
	Types["BAR_P06_PATIENT"] = reflect.TypeOf(BAR_P06_PATIENT{})
	Types["CRM_C01"] = reflect.TypeOf(CRM_C01{})
	Types["CRM_C01_PATIENT"] = reflect.TypeOf(CRM_C01_PATIENT{})
	Types["CSU_C09"] = reflect.TypeOf(CSU_C09{})
	Types["CSU_C09_PATIENT"] = reflect.TypeOf(CSU_C09_PATIENT{})
	Types["CSU_C09_RX_ADMIN"] = reflect.TypeOf(CSU_C09_RX_ADMIN{})
	Types["CSU_C09_STUDY_OBSERVATION"] = reflect.TypeOf(CSU_C09_STUDY_OBSERVATION{})
	Types["CSU_C09_STUDY_PHARM"] = reflect.TypeOf(CSU_C09_STUDY_PHARM{})
	Types["CSU_C09_STUDY_PHASE"] = reflect.TypeOf(CSU_C09_STUDY_PHASE{})
	Types["CSU_C09_STUDY_SCHEDULE"] = reflect.TypeOf(CSU_C09_STUDY_SCHEDULE{})
	Types["CSU_C09_VISIT"] = reflect.TypeOf(CSU_C09_VISIT{})
	Types["DFT_P03"] = reflect.TypeOf(DFT_P03{})
	Types["DFT_P03_FINANCIAL"] = reflect.TypeOf(DFT_P03_FINANCIAL{})
	Types["DFT_P03_FINANCIAL_PROCEDURE"] = reflect.TypeOf(DFT_P03_FINANCIAL_PROCEDURE{})
	Types["DFT_P03_INSURANCE"] = reflect.TypeOf(DFT_P03_INSURANCE{})
	Types["DOC_T12"] = reflect.TypeOf(DOC_T12{})
	Types["DOC_T12_RESULT"] = reflect.TypeOf(DOC_T12_RESULT{})
	Types["DSR_P04"] = reflect.TypeOf(DSR_P04{})
	Types["DSR_Q01"] = reflect.TypeOf(DSR_Q01{})
	Types["DSR_Q03"] = reflect.TypeOf(DSR_Q03{})
	Types["DSR_R03"] = reflect.TypeOf(DSR_R03{})
	Types["EDR_Q01"] = reflect.TypeOf(EDR_Q01{})
	Types["EQQ_Q01"] = reflect.TypeOf(EQQ_Q01{})
	Types["ERP_Q01"] = reflect.TypeOf(ERP_Q01{})
	Types["MCF_Q02"] = reflect.TypeOf(MCF_Q02{})
	Types["MDM_T01"] = reflect.TypeOf(MDM_T01{})
	Types["MDM_T02"] = reflect.TypeOf(MDM_T02{})
	Types["MFD_M01"] = reflect.TypeOf(MFD_M01{})
	Types["MFD_M02"] = reflect.TypeOf(MFD_M02{})
	Types["MFD_M03"] = reflect.TypeOf(MFD_M03{})
	Types["MFK_M01"] = reflect.TypeOf(MFK_M01{})
	Types["MFK_M02"] = reflect.TypeOf(MFK_M02{})
	Types["MFK_M03"] = reflect.TypeOf(MFK_M03{})
	Types["MFN_M01"] = reflect.TypeOf(MFN_M01{})
	Types["MFN_M01_MF"] = reflect.TypeOf(MFN_M01_MF{})
	Types["MFN_M02"] = reflect.TypeOf(MFN_M02{})
	Types["MFN_M02_MF_STAFF"] = reflect.TypeOf(MFN_M02_MF_STAFF{})
	Types["MFN_M03"] = reflect.TypeOf(MFN_M03{})
	Types["MFN_M03_MF_TEST"] = reflect.TypeOf(MFN_M03_MF_TEST{})
	Types["MFN_M05"] = reflect.TypeOf(MFN_M05{})
	Types["MFN_M05_MF_LOCATION"] = reflect.TypeOf(MFN_M05_MF_LOCATION{})
	Types["MFN_M05_MF_LOC_DEPT"] = reflect.TypeOf(MFN_M05_MF_LOC_DEPT{})
	Types["MFN_M06"] = reflect.TypeOf(MFN_M06{})
	Types["MFN_M06_MF_CDM"] = reflect.TypeOf(MFN_M06_MF_CDM{})
	Types["MFN_M07"] = reflect.TypeOf(MFN_M07{})
	Types["MFN_M07_MF_CLIN_STUDY"] = reflect.TypeOf(MFN_M07_MF_CLIN_STUDY{})
	Types["MFN_M07_MF_PHASE_SCHED_DETAIL"] = reflect.TypeOf(MFN_M07_MF_PHASE_SCHED_DETAIL{})
	Types["MFN_M08"] = reflect.TypeOf(MFN_M08{})
	Types["MFN_M08_MF_NUMERIC_OBSERVATION"] = reflect.TypeOf(MFN_M08_MF_NUMERIC_OBSERVATION{})
	Types["MFN_M08_MF_TEST_NUMERIC"] = reflect.TypeOf(MFN_M08_MF_TEST_NUMERIC{})
	Types["MFN_M09"] = reflect.TypeOf(MFN_M09{})
	Types["MFN_M09_MF_TEST_CATEGORICAL"] = reflect.TypeOf(MFN_M09_MF_TEST_CATEGORICAL{})
	Types["MFN_M09_MF_TEST_CAT_DETAIL"] = reflect.TypeOf(MFN_M09_MF_TEST_CAT_DETAIL{})
	Types["MFN_M10"] = reflect.TypeOf(MFN_M10{})
	Types["MFN_M10_MF_TEST_BATTERIES"] = reflect.TypeOf(MFN_M10_MF_TEST_BATTERIES{})
	Types["MFN_M10_MF_TEST_BATT_DETAIL"] = reflect.TypeOf(MFN_M10_MF_TEST_BATT_DETAIL{})
	Types["MFN_M11"] = reflect.TypeOf(MFN_M11{})
	Types["MFN_M11_MF_TEST_CALCULATED"] = reflect.TypeOf(MFN_M11_MF_TEST_CALCULATED{})
	Types["MFN_M11_MF_TEST_CALC_DETAIL"] = reflect.TypeOf(MFN_M11_MF_TEST_CALC_DETAIL{})
	Types["MFQ_M01"] = reflect.TypeOf(MFQ_M01{})
	Types["MFQ_M02"] = reflect.TypeOf(MFQ_M02{})
	Types["MFQ_M03"] = reflect.TypeOf(MFQ_M03{})
	Types["MFR_M01"] = reflect.TypeOf(MFR_M01{})
	Types["MFR_M01_MF"] = reflect.TypeOf(MFR_M01_MF{})
	Types["MFR_M02"] = reflect.TypeOf(MFR_M02{})
	Types["MFR_M02_MF_STAFF"] = reflect.TypeOf(MFR_M02_MF_STAFF{})
	Types["MFR_M03"] = reflect.TypeOf(MFR_M03{})
	Types["MFR_M03_MF_TEST"] = reflect.TypeOf(MFR_M03_MF_TEST{})
	Types["NMD_N01_APP_STATS"] = reflect.TypeOf(NMD_N01_APP_STATS{})
	Types["NMD_N01_APP_STATUS"] = reflect.TypeOf(NMD_N01_APP_STATUS{})
	Types["NMD_N01_CLOCK"] = reflect.TypeOf(NMD_N01_CLOCK{})
	Types["NMD_N01_CLOCK_AND_STATS_WITH_NOTES"] = reflect.TypeOf(NMD_N01_CLOCK_AND_STATS_WITH_NOTES{})
	Types["NMD_N01"] = reflect.TypeOf(NMD_N01{})
	Types["NMQ_N02_CLOCK_AND_STATISTICS"] = reflect.TypeOf(NMQ_N02_CLOCK_AND_STATISTICS{})
	Types["NMQ_N02"] = reflect.TypeOf(NMQ_N02{})
	Types["NMQ_N02_QRY_WITH_DETAIL"] = reflect.TypeOf(NMQ_N02_QRY_WITH_DETAIL{})
	Types["NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT"] = reflect.TypeOf(NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT{})
	Types["NMR_N02"] = reflect.TypeOf(NMR_N02{})
	Types["OMD_O01"] = reflect.TypeOf(OMD_O01{})
	Types["OMD_O01_DIET"] = reflect.TypeOf(OMD_O01_DIET{})
	Types["OMD_O01_INSURANCE"] = reflect.TypeOf(OMD_O01_INSURANCE{})
	Types["OMD_O01_OBSERVATION"] = reflect.TypeOf(OMD_O01_OBSERVATION{})
	Types["OMD_O01_ORDER_DIET"] = reflect.TypeOf(OMD_O01_ORDER_DIET{})
	Types["OMD_O01_ORDER_TRAY"] = reflect.TypeOf(OMD_O01_ORDER_TRAY{})
	Types["OMD_O01_PATIENT"] = reflect.TypeOf(OMD_O01_PATIENT{})
	Types["OMD_O01_PATIENT_VISIT"] = reflect.TypeOf(OMD_O01_PATIENT_VISIT{})
	Types["OMN_O01"] = reflect.TypeOf(OMN_O01{})
	Types["OMN_O01_INSURANCE"] = reflect.TypeOf(OMN_O01_INSURANCE{})
	Types["OMN_O01_OBSERVATION"] = reflect.TypeOf(OMN_O01_OBSERVATION{})
	Types["OMN_O01_ORDER"] = reflect.TypeOf(OMN_O01_ORDER{})
	Types["OMN_O01_ORDER_DETAIL"] = reflect.TypeOf(OMN_O01_ORDER_DETAIL{})
	Types["OMN_O01_PATIENT"] = reflect.TypeOf(OMN_O01_PATIENT{})
	Types["OMN_O01_PATIENT_VISIT"] = reflect.TypeOf(OMN_O01_PATIENT_VISIT{})
	Types["OMS_O01"] = reflect.TypeOf(OMS_O01{})
	Types["OMS_O01_INSURANCE"] = reflect.TypeOf(OMS_O01_INSURANCE{})
	Types["OMS_O01_OBSERVATION"] = reflect.TypeOf(OMS_O01_OBSERVATION{})
	Types["OMS_O01_ORDER"] = reflect.TypeOf(OMS_O01_ORDER{})
	Types["OMS_O01_ORDER_DETAIL"] = reflect.TypeOf(OMS_O01_ORDER_DETAIL{})
	Types["OMS_O01_PATIENT"] = reflect.TypeOf(OMS_O01_PATIENT{})
	Types["OMS_O01_PATIENT_VISIT"] = reflect.TypeOf(OMS_O01_PATIENT_VISIT{})
	Types["ORD_O02"] = reflect.TypeOf(ORD_O02{})
	Types["ORD_O02_ORDER_DIET"] = reflect.TypeOf(ORD_O02_ORDER_DIET{})
	Types["ORD_O02_ORDER_TRAY"] = reflect.TypeOf(ORD_O02_ORDER_TRAY{})
	Types["ORD_O02_PATIENT"] = reflect.TypeOf(ORD_O02_PATIENT{})
	Types["ORD_O02_RESPONSE"] = reflect.TypeOf(ORD_O02_RESPONSE{})
	Types["ORF_R04"] = reflect.TypeOf(ORF_R04{})
	Types["ORF_R04_OBSERVATION"] = reflect.TypeOf(ORF_R04_OBSERVATION{})
	Types["ORF_R04_ORDER"] = reflect.TypeOf(ORF_R04_ORDER{})
	Types["ORF_R04_PATIENT"] = reflect.TypeOf(ORF_R04_PATIENT{})
	Types["ORF_R04_QUERY_RESPONSE"] = reflect.TypeOf(ORF_R04_QUERY_RESPONSE{})
	Types["ORM_O01_CHOICE"] = reflect.TypeOf(ORM_O01_CHOICE{})
	Types["ORM_O01"] = reflect.TypeOf(ORM_O01{})
	Types["ORM_O01_INSURANCE"] = reflect.TypeOf(ORM_O01_INSURANCE{})
	Types["ORM_O01_OBSERVATION"] = reflect.TypeOf(ORM_O01_OBSERVATION{})
	Types["ORM_O01_ORDER"] = reflect.TypeOf(ORM_O01_ORDER{})
	Types["ORM_O01_ORDER_DETAIL"] = reflect.TypeOf(ORM_O01_ORDER_DETAIL{})
	Types["ORM_O01_PATIENT"] = reflect.TypeOf(ORM_O01_PATIENT{})
	Types["ORM_O01_PATIENT_VISIT"] = reflect.TypeOf(ORM_O01_PATIENT_VISIT{})
	Types["ORN_O02"] = reflect.TypeOf(ORN_O02{})
	Types["ORN_O02_ORDER"] = reflect.TypeOf(ORN_O02_ORDER{})
	Types["ORN_O02_PATIENT"] = reflect.TypeOf(ORN_O02_PATIENT{})
	Types["ORN_O02_RESPONSE"] = reflect.TypeOf(ORN_O02_RESPONSE{})
	Types["ORR_O02_CHOICE"] = reflect.TypeOf(ORR_O02_CHOICE{})
	Types["ORR_O02"] = reflect.TypeOf(ORR_O02{})
	Types["ORR_O02_ORDER"] = reflect.TypeOf(ORR_O02_ORDER{})
	Types["ORR_O02_ORDER_DETAIL"] = reflect.TypeOf(ORR_O02_ORDER_DETAIL{})
	Types["ORR_O02_PATIENT"] = reflect.TypeOf(ORR_O02_PATIENT{})
	Types["ORR_O02_RESPONSE"] = reflect.TypeOf(ORR_O02_RESPONSE{})
	Types["ORU_R01"] = reflect.TypeOf(ORU_R01{})
	Types["ORU_R01_OBSERVATION"] = reflect.TypeOf(ORU_R01_OBSERVATION{})
	Types["ORU_R01_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R01_ORDER_OBSERVATION{})
	Types["ORU_R01_PATIENT"] = reflect.TypeOf(ORU_R01_PATIENT{})
	Types["ORU_R01_PATIENT_RESULT"] = reflect.TypeOf(ORU_R01_PATIENT_RESULT{})
	Types["ORU_R01_RESPONSE"] = reflect.TypeOf(ORU_R01_RESPONSE{})
	Types["ORU_R01_VISIT"] = reflect.TypeOf(ORU_R01_VISIT{})
	Types["ORU_R03"] = reflect.TypeOf(ORU_R03{})
	Types["ORU_R03_OBSERVATION"] = reflect.TypeOf(ORU_R03_OBSERVATION{})
	Types["ORU_R03_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R03_ORDER_OBSERVATION{})
	Types["ORU_R03_PATIENT"] = reflect.TypeOf(ORU_R03_PATIENT{})
	Types["ORU_R03_PATIENT_RESULT"] = reflect.TypeOf(ORU_R03_PATIENT_RESULT{})
	Types["ORU_R32"] = reflect.TypeOf(ORU_R32{})
	Types["ORU_R32_OBSERVATION"] = reflect.TypeOf(ORU_R32_OBSERVATION{})
	Types["ORU_R32_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R32_ORDER_OBSERVATION{})
	Types["ORU_R32_PATIENT"] = reflect.TypeOf(ORU_R32_PATIENT{})
	Types["ORU_R32_PATIENT_RESULT"] = reflect.TypeOf(ORU_R32_PATIENT_RESULT{})
	Types["ORU_R32_VISIT"] = reflect.TypeOf(ORU_R32_VISIT{})
	Types["OSQ_Q06"] = reflect.TypeOf(OSQ_Q06{})
	Types["OSR_Q06"] = reflect.TypeOf(OSR_Q06{})
	Types["OSR_Q06_ORDER"] = reflect.TypeOf(OSR_Q06_ORDER{})
	Types["OSR_Q06_PATIENT"] = reflect.TypeOf(OSR_Q06_PATIENT{})
	Types["OSR_Q06_RESPONSE"] = reflect.TypeOf(OSR_Q06_RESPONSE{})
	Types["PEX_P07_ASSOCIATED_PERSON"] = reflect.TypeOf(PEX_P07_ASSOCIATED_PERSON{})
	Types["PEX_P07_ASSOCIATED_RX_ADMIN"] = reflect.TypeOf(PEX_P07_ASSOCIATED_RX_ADMIN{})
	Types["PEX_P07_ASSOCIATED_RX_ORDER"] = reflect.TypeOf(PEX_P07_ASSOCIATED_RX_ORDER{})
	Types["PEX_P07"] = reflect.TypeOf(PEX_P07{})
	Types["PEX_P07_EXPERIENCE"] = reflect.TypeOf(PEX_P07_EXPERIENCE{})
	Types["PEX_P07_PEX_CAUSE"] = reflect.TypeOf(PEX_P07_PEX_CAUSE{})
	Types["PEX_P07_PEX_OBSERVATION"] = reflect.TypeOf(PEX_P07_PEX_OBSERVATION{})
	Types["PEX_P07_RX_ADMINISTRATION"] = reflect.TypeOf(PEX_P07_RX_ADMINISTRATION{})
	Types["PEX_P07_RX_ORDER"] = reflect.TypeOf(PEX_P07_RX_ORDER{})
	Types["PEX_P07_STUDY"] = reflect.TypeOf(PEX_P07_STUDY{})
	Types["PEX_P07_VISIT"] = reflect.TypeOf(PEX_P07_VISIT{})
	Types["PGL_PC6"] = reflect.TypeOf(PGL_PC6{})
	Types["PGL_PC6_GOAL"] = reflect.TypeOf(PGL_PC6_GOAL{})
	Types["PGL_PC6_GOAL_ROLE"] = reflect.TypeOf(PGL_PC6_GOAL_ROLE{})
	Types["PGL_PC6_OBSERVATION"] = reflect.TypeOf(PGL_PC6_OBSERVATION{})
	Types["PGL_PC6_ORDER"] = reflect.TypeOf(PGL_PC6_ORDER{})
	Types["PGL_PC6_ORDER_DETAIL"] = reflect.TypeOf(PGL_PC6_ORDER_DETAIL{})
	Types["PGL_PC6_ORDER_OBSERVATION"] = reflect.TypeOf(PGL_PC6_ORDER_OBSERVATION{})
	Types["PGL_PC6_PATHWAY"] = reflect.TypeOf(PGL_PC6_PATHWAY{})
	Types["PGL_PC6_PATIENT_VISIT"] = reflect.TypeOf(PGL_PC6_PATIENT_VISIT{})
	Types["PGL_PC6_PROBLEM"] = reflect.TypeOf(PGL_PC6_PROBLEM{})
	Types["PGL_PC6_PROBLEM_OBSERVATION"] = reflect.TypeOf(PGL_PC6_PROBLEM_OBSERVATION{})
	Types["PGL_PC6_PROBLEM_ROLE"] = reflect.TypeOf(PGL_PC6_PROBLEM_ROLE{})
	Types["PIN_I07"] = reflect.TypeOf(PIN_I07{})
	Types["PIN_I07_GUARANTOR_INSURANCE"] = reflect.TypeOf(PIN_I07_GUARANTOR_INSURANCE{})
	Types["PIN_I07_INSURANCE"] = reflect.TypeOf(PIN_I07_INSURANCE{})
	Types["PIN_I07_PROVIDER"] = reflect.TypeOf(PIN_I07_PROVIDER{})
	Types["PPG_PCG"] = reflect.TypeOf(PPG_PCG{})
	Types["PPG_PCG_GOAL"] = reflect.TypeOf(PPG_PCG_GOAL{})
	Types["PPG_PCG_GOAL_OBSERVATION"] = reflect.TypeOf(PPG_PCG_GOAL_OBSERVATION{})
	Types["PPG_PCG_GOAL_ROLE"] = reflect.TypeOf(PPG_PCG_GOAL_ROLE{})
	Types["PPG_PCG_ORDER"] = reflect.TypeOf(PPG_PCG_ORDER{})
	Types["PPG_PCG_ORDER_DETAIL"] = reflect.TypeOf(PPG_PCG_ORDER_DETAIL{})
	Types["PPG_PCG_ORDER_OBSERVATION"] = reflect.TypeOf(PPG_PCG_ORDER_OBSERVATION{})
	Types["PPG_PCG_PATHWAY"] = reflect.TypeOf(PPG_PCG_PATHWAY{})
	Types["PPG_PCG_PATHWAY_ROLE"] = reflect.TypeOf(PPG_PCG_PATHWAY_ROLE{})
	Types["PPG_PCG_PATIENT_VISIT"] = reflect.TypeOf(PPG_PCG_PATIENT_VISIT{})
	Types["PPG_PCG_PROBLEM"] = reflect.TypeOf(PPG_PCG_PROBLEM{})
	Types["PPG_PCG_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPG_PCG_PROBLEM_OBSERVATION{})
	Types["PPG_PCG_PROBLEM_ROLE"] = reflect.TypeOf(PPG_PCG_PROBLEM_ROLE{})
	Types["PPP_PCB"] = reflect.TypeOf(PPP_PCB{})
	Types["PPP_PCB_GOAL"] = reflect.TypeOf(PPP_PCB_GOAL{})
	Types["PPP_PCB_GOAL_OBSERVATION"] = reflect.TypeOf(PPP_PCB_GOAL_OBSERVATION{})
	Types["PPP_PCB_GOAL_ROLE"] = reflect.TypeOf(PPP_PCB_GOAL_ROLE{})
	Types["PPP_PCB_ORDER"] = reflect.TypeOf(PPP_PCB_ORDER{})
	Types["PPP_PCB_ORDER_DETAIL"] = reflect.TypeOf(PPP_PCB_ORDER_DETAIL{})
	Types["PPP_PCB_ORDER_OBSERVATION"] = reflect.TypeOf(PPP_PCB_ORDER_OBSERVATION{})
	Types["PPP_PCB_PATHWAY"] = reflect.TypeOf(PPP_PCB_PATHWAY{})
	Types["PPP_PCB_PATHWAY_ROLE"] = reflect.TypeOf(PPP_PCB_PATHWAY_ROLE{})
	Types["PPP_PCB_PATIENT_VISIT"] = reflect.TypeOf(PPP_PCB_PATIENT_VISIT{})
	Types["PPP_PCB_PROBLEM"] = reflect.TypeOf(PPP_PCB_PROBLEM{})
	Types["PPP_PCB_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPP_PCB_PROBLEM_OBSERVATION{})
	Types["PPP_PCB_PROBLEM_ROLE"] = reflect.TypeOf(PPP_PCB_PROBLEM_ROLE{})
	Types["PPR_PC1"] = reflect.TypeOf(PPR_PC1{})
	Types["PPR_PC1_GOAL"] = reflect.TypeOf(PPR_PC1_GOAL{})
	Types["PPR_PC1_GOAL_OBSERVATION"] = reflect.TypeOf(PPR_PC1_GOAL_OBSERVATION{})
	Types["PPR_PC1_GOAL_ROLE"] = reflect.TypeOf(PPR_PC1_GOAL_ROLE{})
	Types["PPR_PC1_ORDER"] = reflect.TypeOf(PPR_PC1_ORDER{})
	Types["PPR_PC1_ORDER_DETAIL"] = reflect.TypeOf(PPR_PC1_ORDER_DETAIL{})
	Types["PPR_PC1_ORDER_OBSERVATION"] = reflect.TypeOf(PPR_PC1_ORDER_OBSERVATION{})
	Types["PPR_PC1_PATHWAY"] = reflect.TypeOf(PPR_PC1_PATHWAY{})
	Types["PPR_PC1_PATIENT_VISIT"] = reflect.TypeOf(PPR_PC1_PATIENT_VISIT{})
	Types["PPR_PC1_PROBLEM"] = reflect.TypeOf(PPR_PC1_PROBLEM{})
	Types["PPR_PC1_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPR_PC1_PROBLEM_OBSERVATION{})
	Types["PPR_PC1_PROBLEM_ROLE"] = reflect.TypeOf(PPR_PC1_PROBLEM_ROLE{})
	Types["PPT_PCL"] = reflect.TypeOf(PPT_PCL{})
	Types["PPT_PCL_GOAL"] = reflect.TypeOf(PPT_PCL_GOAL{})
	Types["PPT_PCL_GOAL_OBSERVATION"] = reflect.TypeOf(PPT_PCL_GOAL_OBSERVATION{})
	Types["PPT_PCL_GOAL_ROLE"] = reflect.TypeOf(PPT_PCL_GOAL_ROLE{})
	Types["PPT_PCL_ORDER"] = reflect.TypeOf(PPT_PCL_ORDER{})
	Types["PPT_PCL_ORDER_DETAIL"] = reflect.TypeOf(PPT_PCL_ORDER_DETAIL{})
	Types["PPT_PCL_ORDER_OBSERVATION"] = reflect.TypeOf(PPT_PCL_ORDER_OBSERVATION{})
	Types["PPT_PCL_PATHWAY"] = reflect.TypeOf(PPT_PCL_PATHWAY{})
	Types["PPT_PCL_PATHWAY_ROLE"] = reflect.TypeOf(PPT_PCL_PATHWAY_ROLE{})
	Types["PPT_PCL_PATIENT"] = reflect.TypeOf(PPT_PCL_PATIENT{})
	Types["PPT_PCL_PATIENT_VISIT"] = reflect.TypeOf(PPT_PCL_PATIENT_VISIT{})
	Types["PPT_PCL_PROBLEM"] = reflect.TypeOf(PPT_PCL_PROBLEM{})
	Types["PPT_PCL_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPT_PCL_PROBLEM_OBSERVATION{})
	Types["PPT_PCL_PROBLEM_ROLE"] = reflect.TypeOf(PPT_PCL_PROBLEM_ROLE{})
	Types["PPV_PCA"] = reflect.TypeOf(PPV_PCA{})
	Types["PPV_PCA_GOAL"] = reflect.TypeOf(PPV_PCA_GOAL{})
	Types["PPV_PCA_GOAL_OBSERVATION"] = reflect.TypeOf(PPV_PCA_GOAL_OBSERVATION{})
	Types["PPV_PCA_GOAL_PATHWAY"] = reflect.TypeOf(PPV_PCA_GOAL_PATHWAY{})
	Types["PPV_PCA_GOAL_ROLE"] = reflect.TypeOf(PPV_PCA_GOAL_ROLE{})
	Types["PPV_PCA_ORDER"] = reflect.TypeOf(PPV_PCA_ORDER{})
	Types["PPV_PCA_ORDER_DETAIL"] = reflect.TypeOf(PPV_PCA_ORDER_DETAIL{})
	Types["PPV_PCA_ORDER_OBSERVATION"] = reflect.TypeOf(PPV_PCA_ORDER_OBSERVATION{})
	Types["PPV_PCA_PATIENT"] = reflect.TypeOf(PPV_PCA_PATIENT{})
	Types["PPV_PCA_PATIENT_VISIT"] = reflect.TypeOf(PPV_PCA_PATIENT_VISIT{})
	Types["PPV_PCA_PROBLEM"] = reflect.TypeOf(PPV_PCA_PROBLEM{})
	Types["PPV_PCA_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPV_PCA_PROBLEM_OBSERVATION{})
	Types["PPV_PCA_PROBLEM_ROLE"] = reflect.TypeOf(PPV_PCA_PROBLEM_ROLE{})
	Types["PRR_PC5"] = reflect.TypeOf(PRR_PC5{})
	Types["PRR_PC5_GOAL"] = reflect.TypeOf(PRR_PC5_GOAL{})
	Types["PRR_PC5_GOAL_OBSERVATION"] = reflect.TypeOf(PRR_PC5_GOAL_OBSERVATION{})
	Types["PRR_PC5_GOAL_ROLE"] = reflect.TypeOf(PRR_PC5_GOAL_ROLE{})
	Types["PRR_PC5_ORDER"] = reflect.TypeOf(PRR_PC5_ORDER{})
	Types["PRR_PC5_ORDER_DETAIL"] = reflect.TypeOf(PRR_PC5_ORDER_DETAIL{})
	Types["PRR_PC5_ORDER_OBSERVATION"] = reflect.TypeOf(PRR_PC5_ORDER_OBSERVATION{})
	Types["PRR_PC5_PATIENT"] = reflect.TypeOf(PRR_PC5_PATIENT{})
	Types["PRR_PC5_PATIENT_VISIT"] = reflect.TypeOf(PRR_PC5_PATIENT_VISIT{})
	Types["PRR_PC5_PROBLEM"] = reflect.TypeOf(PRR_PC5_PROBLEM{})
	Types["PRR_PC5_PROBLEM_OBSERVATION"] = reflect.TypeOf(PRR_PC5_PROBLEM_OBSERVATION{})
	Types["PRR_PC5_PROBLEM_PATHWAY"] = reflect.TypeOf(PRR_PC5_PROBLEM_PATHWAY{})
	Types["PRR_PC5_PROBLEM_ROLE"] = reflect.TypeOf(PRR_PC5_PROBLEM_ROLE{})
	Types["PTR_PCF"] = reflect.TypeOf(PTR_PCF{})
	Types["PTR_PCF_GOAL"] = reflect.TypeOf(PTR_PCF_GOAL{})
	Types["PTR_PCF_GOAL_OBSERVATION"] = reflect.TypeOf(PTR_PCF_GOAL_OBSERVATION{})
	Types["PTR_PCF_GOAL_ROLE"] = reflect.TypeOf(PTR_PCF_GOAL_ROLE{})
	Types["PTR_PCF_ORDER"] = reflect.TypeOf(PTR_PCF_ORDER{})
	Types["PTR_PCF_ORDER_DETAIL"] = reflect.TypeOf(PTR_PCF_ORDER_DETAIL{})
	Types["PTR_PCF_ORDER_OBSERVATION"] = reflect.TypeOf(PTR_PCF_ORDER_OBSERVATION{})
	Types["PTR_PCF_PATHWAY"] = reflect.TypeOf(PTR_PCF_PATHWAY{})
	Types["PTR_PCF_PATHWAY_ROLE"] = reflect.TypeOf(PTR_PCF_PATHWAY_ROLE{})
	Types["PTR_PCF_PATIENT"] = reflect.TypeOf(PTR_PCF_PATIENT{})
	Types["PTR_PCF_PATIENT_VISIT"] = reflect.TypeOf(PTR_PCF_PATIENT_VISIT{})
	Types["PTR_PCF_PROBLEM"] = reflect.TypeOf(PTR_PCF_PROBLEM{})
	Types["PTR_PCF_PROBLEM_OBSERVATION"] = reflect.TypeOf(PTR_PCF_PROBLEM_OBSERVATION{})
	Types["PTR_PCF_PROBLEM_ROLE"] = reflect.TypeOf(PTR_PCF_PROBLEM_ROLE{})
	Types["QCK_Q02"] = reflect.TypeOf(QCK_Q02{})
	Types["QRY_A19"] = reflect.TypeOf(QRY_A19{})
	Types["QRY_P04"] = reflect.TypeOf(QRY_P04{})
	Types["QRY_PC4"] = reflect.TypeOf(QRY_PC4{})
	Types["QRY_Q01"] = reflect.TypeOf(QRY_Q01{})
	Types["QRY_Q02"] = reflect.TypeOf(QRY_Q02{})
	Types["QRY_R02"] = reflect.TypeOf(QRY_R02{})
	Types["QRY_T12"] = reflect.TypeOf(QRY_T12{})
	Types["RAR_RAR"] = reflect.TypeOf(RAR_RAR{})
	Types["RAR_RAR_DEFINITION"] = reflect.TypeOf(RAR_RAR_DEFINITION{})
	Types["RAR_RAR_ENCODING"] = reflect.TypeOf(RAR_RAR_ENCODING{})
	Types["RAR_RAR_ORDER"] = reflect.TypeOf(RAR_RAR_ORDER{})
	Types["RAR_RAR_PATIENT"] = reflect.TypeOf(RAR_RAR_PATIENT{})
	Types["RAS_O01_COMPONENTS"] = reflect.TypeOf(RAS_O01_COMPONENTS{})
	Types["RAS_O01"] = reflect.TypeOf(RAS_O01{})
	Types["RAS_O01_ENCODING"] = reflect.TypeOf(RAS_O01_ENCODING{})
	Types["RAS_O01_OBSERVATION"] = reflect.TypeOf(RAS_O01_OBSERVATION{})
	Types["RAS_O01_ORDER"] = reflect.TypeOf(RAS_O01_ORDER{})
	Types["RAS_O01_ORDER_DETAIL"] = reflect.TypeOf(RAS_O01_ORDER_DETAIL{})
	Types["RAS_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RAS_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RAS_O01_PATIENT"] = reflect.TypeOf(RAS_O01_PATIENT{})
	Types["RAS_O01_PATIENT_VISIT"] = reflect.TypeOf(RAS_O01_PATIENT_VISIT{})
	Types["RCI_I05"] = reflect.TypeOf(RCI_I05{})
	Types["RCI_I05_OBSERVATION"] = reflect.TypeOf(RCI_I05_OBSERVATION{})
	Types["RCI_I05_PROVIDER"] = reflect.TypeOf(RCI_I05_PROVIDER{})
	Types["RCI_I05_RESULTS"] = reflect.TypeOf(RCI_I05_RESULTS{})
	Types["RCL_I06"] = reflect.TypeOf(RCL_I06{})
	Types["RCL_I06_PROVIDER"] = reflect.TypeOf(RCL_I06_PROVIDER{})
	Types["RDE_O01_COMPONENT"] = reflect.TypeOf(RDE_O01_COMPONENT{})
	Types["RDE_O01"] = reflect.TypeOf(RDE_O01{})
	Types["RDE_O01_INSURANCE"] = reflect.TypeOf(RDE_O01_INSURANCE{})
	Types["RDE_O01_OBSERVATION"] = reflect.TypeOf(RDE_O01_OBSERVATION{})
	Types["RDE_O01_ORDER"] = reflect.TypeOf(RDE_O01_ORDER{})
	Types["RDE_O01_ORDER_DETAIL"] = reflect.TypeOf(RDE_O01_ORDER_DETAIL{})
	Types["RDE_O01_PATIENT"] = reflect.TypeOf(RDE_O01_PATIENT{})
	Types["RDE_O01_PATIENT_VISIT"] = reflect.TypeOf(RDE_O01_PATIENT_VISIT{})
	Types["RDO_O01_COMPONENT"] = reflect.TypeOf(RDO_O01_COMPONENT{})
	Types["RDO_O01"] = reflect.TypeOf(RDO_O01{})
	Types["RDO_O01_INSURANCE"] = reflect.TypeOf(RDO_O01_INSURANCE{})
	Types["RDO_O01_OBSERVATION"] = reflect.TypeOf(RDO_O01_OBSERVATION{})
	Types["RDO_O01_ORDER"] = reflect.TypeOf(RDO_O01_ORDER{})
	Types["RDO_O01_ORDER_DETAIL"] = reflect.TypeOf(RDO_O01_ORDER_DETAIL{})
	Types["RDO_O01_PATIENT"] = reflect.TypeOf(RDO_O01_PATIENT{})
	Types["RDO_O01_PATIENT_VISIT"] = reflect.TypeOf(RDO_O01_PATIENT_VISIT{})
	Types["RDR_RDR"] = reflect.TypeOf(RDR_RDR{})
	Types["RDR_RDR_DEFINITION"] = reflect.TypeOf(RDR_RDR_DEFINITION{})
	Types["RDR_RDR_DISPENSE"] = reflect.TypeOf(RDR_RDR_DISPENSE{})
	Types["RDR_RDR_ENCODING"] = reflect.TypeOf(RDR_RDR_ENCODING{})
	Types["RDR_RDR_ORDER"] = reflect.TypeOf(RDR_RDR_ORDER{})
	Types["RDR_RDR_PATIENT"] = reflect.TypeOf(RDR_RDR_PATIENT{})
	Types["RDS_O01_COMPONENT"] = reflect.TypeOf(RDS_O01_COMPONENT{})
	Types["RDS_O01"] = reflect.TypeOf(RDS_O01{})
	Types["RDS_O01_ENCODING"] = reflect.TypeOf(RDS_O01_ENCODING{})
	Types["RDS_O01_OBSERVATION"] = reflect.TypeOf(RDS_O01_OBSERVATION{})
	Types["RDS_O01_ORDER"] = reflect.TypeOf(RDS_O01_ORDER{})
	Types["RDS_O01_ORDER_DETAIL"] = reflect.TypeOf(RDS_O01_ORDER_DETAIL{})
	Types["RDS_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RDS_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RDS_O01_PATIENT"] = reflect.TypeOf(RDS_O01_PATIENT{})
	Types["RDS_O01_PATIENT_VISIT"] = reflect.TypeOf(RDS_O01_PATIENT_VISIT{})
	Types["REF_I12_AUTHORIZATION"] = reflect.TypeOf(REF_I12_AUTHORIZATION{})
	Types["REF_I12"] = reflect.TypeOf(REF_I12{})
	Types["REF_I12_INSURANCE"] = reflect.TypeOf(REF_I12_INSURANCE{})
	Types["REF_I12_OBSERVATION"] = reflect.TypeOf(REF_I12_OBSERVATION{})
	Types["REF_I12_PROCEDURE"] = reflect.TypeOf(REF_I12_PROCEDURE{})
	Types["REF_I12_PROVIDER"] = reflect.TypeOf(REF_I12_PROVIDER{})
	Types["REF_I12_RESULTS"] = reflect.TypeOf(REF_I12_RESULTS{})
	Types["REF_I12_VISIT"] = reflect.TypeOf(REF_I12_VISIT{})
	Types["RER_RER"] = reflect.TypeOf(RER_RER{})
	Types["RER_RER_DEFINITION"] = reflect.TypeOf(RER_RER_DEFINITION{})
	Types["RER_RER_ORDER"] = reflect.TypeOf(RER_RER_ORDER{})
	Types["RER_RER_PATIENT"] = reflect.TypeOf(RER_RER_PATIENT{})
	Types["RGR_RGR"] = reflect.TypeOf(RGR_RGR{})
	Types["RGR_RGR_DEFINITION"] = reflect.TypeOf(RGR_RGR_DEFINITION{})
	Types["RGR_RGR_ENCODING"] = reflect.TypeOf(RGR_RGR_ENCODING{})
	Types["RGR_RGR_ORDER"] = reflect.TypeOf(RGR_RGR_ORDER{})
	Types["RGR_RGR_PATIENT"] = reflect.TypeOf(RGR_RGR_PATIENT{})
	Types["RGV_O01_COMPONENTS"] = reflect.TypeOf(RGV_O01_COMPONENTS{})
	Types["RGV_O01"] = reflect.TypeOf(RGV_O01{})
	Types["RGV_O01_ENCODING"] = reflect.TypeOf(RGV_O01_ENCODING{})
	Types["RGV_O01_GIVE"] = reflect.TypeOf(RGV_O01_GIVE{})
	Types["RGV_O01_OBSERVATION"] = reflect.TypeOf(RGV_O01_OBSERVATION{})
	Types["RGV_O01_ORDER"] = reflect.TypeOf(RGV_O01_ORDER{})
	Types["RGV_O01_ORDER_DETAIL"] = reflect.TypeOf(RGV_O01_ORDER_DETAIL{})
	Types["RGV_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RGV_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RGV_O01_PATIENT"] = reflect.TypeOf(RGV_O01_PATIENT{})
	Types["RGV_O01_PATIENT_VISIT"] = reflect.TypeOf(RGV_O01_PATIENT_VISIT{})
	Types["ROR_ROR"] = reflect.TypeOf(ROR_ROR{})
	Types["ROR_ROR_DEFINITION"] = reflect.TypeOf(ROR_ROR_DEFINITION{})
	Types["ROR_ROR_ORDER"] = reflect.TypeOf(ROR_ROR_ORDER{})
	Types["ROR_ROR_PATIENT"] = reflect.TypeOf(ROR_ROR_PATIENT{})
	Types["RPA_I08_AUTHORIZATION"] = reflect.TypeOf(RPA_I08_AUTHORIZATION{})
	Types["RPA_I08"] = reflect.TypeOf(RPA_I08{})
	Types["RPA_I08_INSURANCE"] = reflect.TypeOf(RPA_I08_INSURANCE{})
	Types["RPA_I08_OBSERVATION"] = reflect.TypeOf(RPA_I08_OBSERVATION{})
	Types["RPA_I08_PROCEDURE"] = reflect.TypeOf(RPA_I08_PROCEDURE{})
	Types["RPA_I08_PROVIDER"] = reflect.TypeOf(RPA_I08_PROVIDER{})
	Types["RPA_I08_RESULTS"] = reflect.TypeOf(RPA_I08_RESULTS{})
	Types["RPA_I08_VISIT"] = reflect.TypeOf(RPA_I08_VISIT{})
	Types["RPI_I01"] = reflect.TypeOf(RPI_I01{})
	Types["RPI_I01_GUARANTOR_INSURANCE"] = reflect.TypeOf(RPI_I01_GUARANTOR_INSURANCE{})
	Types["RPI_I01_INSURANCE"] = reflect.TypeOf(RPI_I01_INSURANCE{})
	Types["RPI_I01_PROVIDER"] = reflect.TypeOf(RPI_I01_PROVIDER{})
	Types["RPL_I02"] = reflect.TypeOf(RPL_I02{})
	Types["RPL_I02_PROVIDER"] = reflect.TypeOf(RPL_I02_PROVIDER{})
	Types["RQA_I08_AUTHORIZATION"] = reflect.TypeOf(RQA_I08_AUTHORIZATION{})
	Types["RQA_I08"] = reflect.TypeOf(RQA_I08{})
	Types["RQA_I08_GUARANTOR_INSURANCE"] = reflect.TypeOf(RQA_I08_GUARANTOR_INSURANCE{})
	Types["RQA_I08_INSURANCE"] = reflect.TypeOf(RQA_I08_INSURANCE{})
	Types["RQA_I08_OBSERVATION"] = reflect.TypeOf(RQA_I08_OBSERVATION{})
	Types["RQA_I08_PROCEDURE"] = reflect.TypeOf(RQA_I08_PROCEDURE{})
	Types["RQA_I08_PROVIDER"] = reflect.TypeOf(RQA_I08_PROVIDER{})
	Types["RQA_I08_RESULTS"] = reflect.TypeOf(RQA_I08_RESULTS{})
	Types["RQA_I08_VISIT"] = reflect.TypeOf(RQA_I08_VISIT{})
	Types["RQC_I05"] = reflect.TypeOf(RQC_I05{})
	Types["RQC_I05_PROVIDER"] = reflect.TypeOf(RQC_I05_PROVIDER{})
	Types["RQC_I06"] = reflect.TypeOf(RQC_I06{})
	Types["RQC_I06_PROVIDER"] = reflect.TypeOf(RQC_I06_PROVIDER{})
	Types["RQI_I01"] = reflect.TypeOf(RQI_I01{})
	Types["RQI_I01_GUARANTOR_INSURANCE"] = reflect.TypeOf(RQI_I01_GUARANTOR_INSURANCE{})
	Types["RQI_I01_INSURANCE"] = reflect.TypeOf(RQI_I01_INSURANCE{})
	Types["RQI_I01_PROVIDER"] = reflect.TypeOf(RQI_I01_PROVIDER{})
	Types["RQP_I04"] = reflect.TypeOf(RQP_I04{})
	Types["RQP_I04_PROVIDER"] = reflect.TypeOf(RQP_I04_PROVIDER{})
	Types["RQQ_Q01"] = reflect.TypeOf(RQQ_Q01{})
	Types["RRA_O02_ADMINISTRATION"] = reflect.TypeOf(RRA_O02_ADMINISTRATION{})
	Types["RRA_O02"] = reflect.TypeOf(RRA_O02{})
	Types["RRA_O02_ORDER"] = reflect.TypeOf(RRA_O02_ORDER{})
	Types["RRA_O02_PATIENT"] = reflect.TypeOf(RRA_O02_PATIENT{})
	Types["RRA_O02_RESPONSE"] = reflect.TypeOf(RRA_O02_RESPONSE{})
	Types["RRD_O02"] = reflect.TypeOf(RRD_O02{})
	Types["RRD_O02_DISPENSE"] = reflect.TypeOf(RRD_O02_DISPENSE{})
	Types["RRD_O02_ORDER"] = reflect.TypeOf(RRD_O02_ORDER{})
	Types["RRD_O02_PATIENT"] = reflect.TypeOf(RRD_O02_PATIENT{})
	Types["RRD_O02_RESPONSE"] = reflect.TypeOf(RRD_O02_RESPONSE{})
	Types["RRG_O02"] = reflect.TypeOf(RRG_O02{})
	Types["RRG_O02_GIVE"] = reflect.TypeOf(RRG_O02_GIVE{})
	Types["RRG_O02_ORDER"] = reflect.TypeOf(RRG_O02_ORDER{})
	Types["RRG_O02_PATIENT"] = reflect.TypeOf(RRG_O02_PATIENT{})
	Types["RRG_O02_RESPONSE"] = reflect.TypeOf(RRG_O02_RESPONSE{})
	Types["RRI_I12_AUTHORIZATION"] = reflect.TypeOf(RRI_I12_AUTHORIZATION{})
	Types["RRI_I12"] = reflect.TypeOf(RRI_I12{})
	Types["RRI_I12_OBSERVATION"] = reflect.TypeOf(RRI_I12_OBSERVATION{})
	Types["RRI_I12_PROCEDURE"] = reflect.TypeOf(RRI_I12_PROCEDURE{})
	Types["RRI_I12_PROVIDER"] = reflect.TypeOf(RRI_I12_PROVIDER{})
	Types["RRI_I12_RESULTS"] = reflect.TypeOf(RRI_I12_RESULTS{})
	Types["RRI_I12_VISIT"] = reflect.TypeOf(RRI_I12_VISIT{})
	Types["RRO_O02"] = reflect.TypeOf(RRO_O02{})
	Types["RRO_O02_ORDER"] = reflect.TypeOf(RRO_O02_ORDER{})
	Types["RRO_O02_ORDER_DETAIL"] = reflect.TypeOf(RRO_O02_ORDER_DETAIL{})
	Types["RRO_O02_PATIENT"] = reflect.TypeOf(RRO_O02_PATIENT{})
	Types["RRO_O02_RESPONSE"] = reflect.TypeOf(RRO_O02_RESPONSE{})
	Types["SIU_S12"] = reflect.TypeOf(SIU_S12{})
	Types["SIU_S12_GENERAL_RESOURCE"] = reflect.TypeOf(SIU_S12_GENERAL_RESOURCE{})
	Types["SIU_S12_LOCATION_RESOURCE"] = reflect.TypeOf(SIU_S12_LOCATION_RESOURCE{})
	Types["SIU_S12_PATIENT"] = reflect.TypeOf(SIU_S12_PATIENT{})
	Types["SIU_S12_PERSONNEL_RESOURCE"] = reflect.TypeOf(SIU_S12_PERSONNEL_RESOURCE{})
	Types["SIU_S12_RESOURCES"] = reflect.TypeOf(SIU_S12_RESOURCES{})
	Types["SIU_S12_SERVICE"] = reflect.TypeOf(SIU_S12_SERVICE{})
	Types["SPQ_Q01"] = reflect.TypeOf(SPQ_Q01{})
	Types["SQM_S25"] = reflect.TypeOf(SQM_S25{})
	Types["SQM_S25_GENERAL_RESOURCE"] = reflect.TypeOf(SQM_S25_GENERAL_RESOURCE{})
	Types["SQM_S25_LOCATION_RESOURCE"] = reflect.TypeOf(SQM_S25_LOCATION_RESOURCE{})
	Types["SQM_S25_PERSONNEL_RESOURCE"] = reflect.TypeOf(SQM_S25_PERSONNEL_RESOURCE{})
	Types["SQM_S25_REQUEST"] = reflect.TypeOf(SQM_S25_REQUEST{})
	Types["SQM_S25_RESOURCES"] = reflect.TypeOf(SQM_S25_RESOURCES{})
	Types["SQM_S25_SERVICE"] = reflect.TypeOf(SQM_S25_SERVICE{})
	Types["SQR_S25"] = reflect.TypeOf(SQR_S25{})
	Types["SQR_S25_GENERAL_RESOURCE"] = reflect.TypeOf(SQR_S25_GENERAL_RESOURCE{})
	Types["SQR_S25_LOCATION_RESOURCE"] = reflect.TypeOf(SQR_S25_LOCATION_RESOURCE{})
	Types["SQR_S25_PATIENT"] = reflect.TypeOf(SQR_S25_PATIENT{})
	Types["SQR_S25_PERSONNEL_RESOURCE"] = reflect.TypeOf(SQR_S25_PERSONNEL_RESOURCE{})
	Types["SQR_S25_RESOURCES"] = reflect.TypeOf(SQR_S25_RESOURCES{})
	Types["SQR_S25_SCHEDULE"] = reflect.TypeOf(SQR_S25_SCHEDULE{})
	Types["SQR_S25_SERVICE"] = reflect.TypeOf(SQR_S25_SERVICE{})
	Types["SRM_S01"] = reflect.TypeOf(SRM_S01{})
	Types["SRM_S01_GENERAL_RESOURCE"] = reflect.TypeOf(SRM_S01_GENERAL_RESOURCE{})
	Types["SRM_S01_LOCATION_RESOURCE"] = reflect.TypeOf(SRM_S01_LOCATION_RESOURCE{})
	Types["SRM_S01_PATIENT"] = reflect.TypeOf(SRM_S01_PATIENT{})
	Types["SRM_S01_PERSONNEL_RESOURCE"] = reflect.TypeOf(SRM_S01_PERSONNEL_RESOURCE{})
	Types["SRM_S01_RESOURCES"] = reflect.TypeOf(SRM_S01_RESOURCES{})
	Types["SRM_S01_SERVICE"] = reflect.TypeOf(SRM_S01_SERVICE{})
	Types["SRR_S01"] = reflect.TypeOf(SRR_S01{})
	Types["SRR_S01_GENERAL_RESOURCE"] = reflect.TypeOf(SRR_S01_GENERAL_RESOURCE{})
	Types["SRR_S01_LOCATION_RESOURCE"] = reflect.TypeOf(SRR_S01_LOCATION_RESOURCE{})
	Types["SRR_S01_PATIENT"] = reflect.TypeOf(SRR_S01_PATIENT{})
	Types["SRR_S01_PERSONNEL_RESOURCE"] = reflect.TypeOf(SRR_S01_PERSONNEL_RESOURCE{})
	Types["SRR_S01_RESOURCES"] = reflect.TypeOf(SRR_S01_RESOURCES{})
	Types["SRR_S01_SCHEDULE"] = reflect.TypeOf(SRR_S01_SCHEDULE{})
	Types["SRR_S01_SERVICE"] = reflect.TypeOf(SRR_S01_SERVICE{})
	Types["SUR_P09"] = reflect.TypeOf(SUR_P09{})
	Types["SUR_P09_FACILITY"] = reflect.TypeOf(SUR_P09_FACILITY{})
	Types["SUR_P09_FACILITY_DETAIL"] = reflect.TypeOf(SUR_P09_FACILITY_DETAIL{})
	Types["SUR_P09_PRODUCT"] = reflect.TypeOf(SUR_P09_PRODUCT{})
	Types["TBR_Q01"] = reflect.TypeOf(TBR_Q01{})
	Types["UDM_Q05"] = reflect.TypeOf(UDM_Q05{})
	Types["VQQ_Q01"] = reflect.TypeOf(VQQ_Q01{})
	Types["VXQ_V01"] = reflect.TypeOf(VXQ_V01{})
	Types["VXR_V03"] = reflect.TypeOf(VXR_V03{})
	Types["VXR_V03_INSURANCE"] = reflect.TypeOf(VXR_V03_INSURANCE{})
	Types["VXR_V03_OBSERVATION"] = reflect.TypeOf(VXR_V03_OBSERVATION{})
	Types["VXR_V03_ORDER"] = reflect.TypeOf(VXR_V03_ORDER{})
	Types["VXR_V03_PATIENT_VISIT"] = reflect.TypeOf(VXR_V03_PATIENT_VISIT{})
	Types["VXU_V04"] = reflect.TypeOf(VXU_V04{})
	Types["VXU_V04_INSURANCE"] = reflect.TypeOf(VXU_V04_INSURANCE{})
	Types["VXU_V04_OBSERVATION"] = reflect.TypeOf(VXU_V04_OBSERVATION{})
	Types["VXU_V04_ORDER"] = reflect.TypeOf(VXU_V04_ORDER{})
	Types["VXU_V04_PATIENT"] = reflect.TypeOf(VXU_V04_PATIENT{})
	Types["VXX_V02"] = reflect.TypeOf(VXX_V02{})
	Types["VXX_V02_PATIENT"] = reflect.TypeOf(VXX_V02_PATIENT{})
	Types["GenericHL7Segment"] = reflect.TypeOf(GenericHL7Segment{})
	Types["ACKv2"] = reflect.TypeOf(ACKv2{})
	Types["ADR_A19v2"] = reflect.TypeOf(ADR_A19v2{})
	Types["ADR_A19_PIDv2"] = reflect.TypeOf(ADR_A19_PIDv2{})
	Types["ADR_A19_PID_IN1v2"] = reflect.TypeOf(ADR_A19_PID_IN1v2{})
	Types["ADT_A01v2"] = reflect.TypeOf(ADT_A01v2{})
	Types["ADT_A01_PR1v2"] = reflect.TypeOf(ADT_A01_PR1v2{})
	Types["ADT_A01_IN1v2"] = reflect.TypeOf(ADT_A01_IN1v2{})
	Types["ADT_A02v2"] = reflect.TypeOf(ADT_A02v2{})
	Types["ADT_A03v2"] = reflect.TypeOf(ADT_A03v2{})
	Types["ADT_A03_PR1v2"] = reflect.TypeOf(ADT_A03_PR1v2{})
	Types["ADT_A04v2"] = reflect.TypeOf(ADT_A04v2{})
	Types["ADT_A04_IN1v2"] = reflect.TypeOf(ADT_A04_IN1v2{})
	Types["ADT_A05v2"] = reflect.TypeOf(ADT_A05v2{})
	Types["ADT_A05_IN1v2"] = reflect.TypeOf(ADT_A05_IN1v2{})
	Types["ADT_A06v2"] = reflect.TypeOf(ADT_A06v2{})
	Types["ADT_A06_PR1v2"] = reflect.TypeOf(ADT_A06_PR1v2{})
	Types["ADT_A06_IN1v2"] = reflect.TypeOf(ADT_A06_IN1v2{})
	Types["ADT_A07v2"] = reflect.TypeOf(ADT_A07v2{})
	Types["ADT_A07_IN1v2"] = reflect.TypeOf(ADT_A07_IN1v2{})
	Types["ADT_A08v2"] = reflect.TypeOf(ADT_A08v2{})
	Types["ADT_A08_IN1v2"] = reflect.TypeOf(ADT_A08_IN1v2{})
	Types["ADT_A09v2"] = reflect.TypeOf(ADT_A09v2{})
	Types["ADT_A10v2"] = reflect.TypeOf(ADT_A10v2{})
	Types["ADT_A11v2"] = reflect.TypeOf(ADT_A11v2{})
	Types["ADT_A12v2"] = reflect.TypeOf(ADT_A12v2{})
	Types["ADT_A13v2"] = reflect.TypeOf(ADT_A13v2{})
	Types["ADT_A13_IN1v2"] = reflect.TypeOf(ADT_A13_IN1v2{})
	Types["ADT_A14v2"] = reflect.TypeOf(ADT_A14v2{})
	Types["ADT_A14_IN1v2"] = reflect.TypeOf(ADT_A14_IN1v2{})
	Types["ADT_A15v2"] = reflect.TypeOf(ADT_A15v2{})
	Types["ADT_A16v2"] = reflect.TypeOf(ADT_A16v2{})
	Types["ADT_A17v2"] = reflect.TypeOf(ADT_A17v2{})
	Types["ADT_A18v2"] = reflect.TypeOf(ADT_A18v2{})
	Types["ADT_A20v2"] = reflect.TypeOf(ADT_A20v2{})
	Types["ADT_A21v2"] = reflect.TypeOf(ADT_A21v2{})
	Types["ADT_A22v2"] = reflect.TypeOf(ADT_A22v2{})
	Types["ADT_A23v2"] = reflect.TypeOf(ADT_A23v2{})
	Types["ADT_A24v2"] = reflect.TypeOf(ADT_A24v2{})
	Types["ADT_A25v2"] = reflect.TypeOf(ADT_A25v2{})
	Types["ADT_A26v2"] = reflect.TypeOf(ADT_A26v2{})
	Types["ADT_A27v2"] = reflect.TypeOf(ADT_A27v2{})
	Types["ADT_A28v2"] = reflect.TypeOf(ADT_A28v2{})
	Types["ADT_A28_IN1v2"] = reflect.TypeOf(ADT_A28_IN1v2{})
	Types["ADT_A29v2"] = reflect.TypeOf(ADT_A29v2{})
	Types["ADT_A30v2"] = reflect.TypeOf(ADT_A30v2{})
	Types["ADT_A31v2"] = reflect.TypeOf(ADT_A31v2{})
	Types["ADT_A31_IN1v2"] = reflect.TypeOf(ADT_A31_IN1v2{})
	Types["ADT_A32v2"] = reflect.TypeOf(ADT_A32v2{})
	Types["ADT_A33v2"] = reflect.TypeOf(ADT_A33v2{})
	Types["ADT_A34v2"] = reflect.TypeOf(ADT_A34v2{})
	Types["ADT_A35v2"] = reflect.TypeOf(ADT_A35v2{})
	Types["ADT_A36v2"] = reflect.TypeOf(ADT_A36v2{})
	Types["ADT_A37v2"] = reflect.TypeOf(ADT_A37v2{})
	Types["ADT_A38v2"] = reflect.TypeOf(ADT_A38v2{})
	Types["ADT_A39v2"] = reflect.TypeOf(ADT_A39v2{})
	Types["ADT_A39_PIDv2"] = reflect.TypeOf(ADT_A39_PIDv2{})
	Types["ADT_A40v2"] = reflect.TypeOf(ADT_A40v2{})
	Types["ADT_A40_PIDv2"] = reflect.TypeOf(ADT_A40_PIDv2{})
	Types["ADT_A43v2"] = reflect.TypeOf(ADT_A43v2{})
	Types["ADT_A43_PIDv2"] = reflect.TypeOf(ADT_A43_PIDv2{})
	Types["ADT_A44v2"] = reflect.TypeOf(ADT_A44v2{})
	Types["ADT_A44_PIDv2"] = reflect.TypeOf(ADT_A44_PIDv2{})
	Types["ADT_A45v2"] = reflect.TypeOf(ADT_A45v2{})
	Types["ADT_A45_MRGv2"] = reflect.TypeOf(ADT_A45_MRGv2{})
	Types["ADT_A50v2"] = reflect.TypeOf(ADT_A50v2{})
	Types["ARD_A19v2"] = reflect.TypeOf(ARD_A19v2{})
	Types["ARD_A19_PIDv2"] = reflect.TypeOf(ARD_A19_PIDv2{})
	Types["ARD_A19_PID_PR1v2"] = reflect.TypeOf(ARD_A19_PID_PR1v2{})
	Types["ARD_A19_PID_IN1v2"] = reflect.TypeOf(ARD_A19_PID_IN1v2{})
	Types["BAR_P01v2"] = reflect.TypeOf(BAR_P01v2{})
	Types["BAR_P01_PV1v2"] = reflect.TypeOf(BAR_P01_PV1v2{})
	Types["BAR_P01_PV1_PR1v2"] = reflect.TypeOf(BAR_P01_PV1_PR1v2{})
	Types["BAR_P01_PV1_IN1v2"] = reflect.TypeOf(BAR_P01_PV1_IN1v2{})
	Types["BAR_P02v2"] = reflect.TypeOf(BAR_P02v2{})
	Types["BAR_P02_PIDv2"] = reflect.TypeOf(BAR_P02_PIDv2{})
	Types["BAR_P06v2"] = reflect.TypeOf(BAR_P06v2{})
	Types["BAR_P06_PIDv2"] = reflect.TypeOf(BAR_P06_PIDv2{})
	Types["CRM_C01v2"] = reflect.TypeOf(CRM_C01v2{})
	Types["CRM_C01_PIDv2"] = reflect.TypeOf(CRM_C01_PIDv2{})
	Types["CSU_C09v2"] = reflect.TypeOf(CSU_C09v2{})
	Types["CSU_C09_PIDv2"] = reflect.TypeOf(CSU_C09_PIDv2{})
	Types["CSU_C09_PID_CSPv2"] = reflect.TypeOf(CSU_C09_PID_CSPv2{})
	Types["CSU_C09_PID_CSP_CSSv2"] = reflect.TypeOf(CSU_C09_PID_CSP_CSSv2{})
	Types["CSU_C09_PID_CSP_CSS_OBRv2"] = reflect.TypeOf(CSU_C09_PID_CSP_CSS_OBRv2{})
	Types["CSU_C09_PID_CSP_CSS_ORCv2"] = reflect.TypeOf(CSU_C09_PID_CSP_CSS_ORCv2{})
	Types["CSU_C09_PID_CSP_CSS_ORC_RXAv2"] = reflect.TypeOf(CSU_C09_PID_CSP_CSS_ORC_RXAv2{})
	Types["DFT_P03v2"] = reflect.TypeOf(DFT_P03v2{})
	Types["DFT_P03_FT1v2"] = reflect.TypeOf(DFT_P03_FT1v2{})
	Types["DFT_P03_FT1_PR1v2"] = reflect.TypeOf(DFT_P03_FT1_PR1v2{})
	Types["DFT_P03_IN1v2"] = reflect.TypeOf(DFT_P03_IN1v2{})
	Types["DOC_T12v2"] = reflect.TypeOf(DOC_T12v2{})
	Types["DOC_T12_PIDv2"] = reflect.TypeOf(DOC_T12_PIDv2{})
	Types["DSR_P04v2"] = reflect.TypeOf(DSR_P04v2{})
	Types["DSR_Q01v2"] = reflect.TypeOf(DSR_Q01v2{})
	Types["DSR_Q03v2"] = reflect.TypeOf(DSR_Q03v2{})
	Types["DSR_R03v2"] = reflect.TypeOf(DSR_R03v2{})
	Types["EDR_Q01v2"] = reflect.TypeOf(EDR_Q01v2{})
	Types["EQQ_Q01v2"] = reflect.TypeOf(EQQ_Q01v2{})
	Types["ERP_Q01v2"] = reflect.TypeOf(ERP_Q01v2{})
	Types["MCF_Q02v2"] = reflect.TypeOf(MCF_Q02v2{})
	Types["MDM_T01v2"] = reflect.TypeOf(MDM_T01v2{})
	Types["MDM_T02v2"] = reflect.TypeOf(MDM_T02v2{})
	Types["MFD_M01v2"] = reflect.TypeOf(MFD_M01v2{})
	Types["MFD_M02v2"] = reflect.TypeOf(MFD_M02v2{})
	Types["MFD_M03v2"] = reflect.TypeOf(MFD_M03v2{})
	Types["MFK_M01v2"] = reflect.TypeOf(MFK_M01v2{})
	Types["MFK_M02v2"] = reflect.TypeOf(MFK_M02v2{})
	Types["MFK_M03v2"] = reflect.TypeOf(MFK_M03v2{})
	Types["MFN_M01v2"] = reflect.TypeOf(MFN_M01v2{})
	Types["MFN_M01_MFEv2"] = reflect.TypeOf(MFN_M01_MFEv2{})
	Types["MFN_M02v2"] = reflect.TypeOf(MFN_M02v2{})
	Types["MFN_M02_MFEv2"] = reflect.TypeOf(MFN_M02_MFEv2{})
	Types["MFN_M03v2"] = reflect.TypeOf(MFN_M03v2{})
	Types["MFN_M03_MFEv2"] = reflect.TypeOf(MFN_M03_MFEv2{})
	Types["MFN_M05v2"] = reflect.TypeOf(MFN_M05v2{})
	Types["MFN_M05_MFEv2"] = reflect.TypeOf(MFN_M05_MFEv2{})
	Types["MFN_M05_MFE_LDPv2"] = reflect.TypeOf(MFN_M05_MFE_LDPv2{})
	Types["MFN_M06v2"] = reflect.TypeOf(MFN_M06v2{})
	Types["MFN_M06_MFEv2"] = reflect.TypeOf(MFN_M06_MFEv2{})
	Types["MFN_M07v2"] = reflect.TypeOf(MFN_M07v2{})
	Types["MFN_M07_MFEv2"] = reflect.TypeOf(MFN_M07_MFEv2{})
	Types["MFN_M07_MFE_CM1v2"] = reflect.TypeOf(MFN_M07_MFE_CM1v2{})
	Types["MFN_M08v2"] = reflect.TypeOf(MFN_M08v2{})
	Types["MFN_M08_MFEv2"] = reflect.TypeOf(MFN_M08_MFEv2{})
	Types["MFN_M09v2"] = reflect.TypeOf(MFN_M09v2{})
	Types["MFN_M09_MFEv2"] = reflect.TypeOf(MFN_M09_MFEv2{})
	Types["MFN_M10v2"] = reflect.TypeOf(MFN_M10v2{})
	Types["MFN_M10_OM5v2"] = reflect.TypeOf(MFN_M10_OM5v2{})
	Types["MFN_M11v2"] = reflect.TypeOf(MFN_M11v2{})
	Types["MFN_M11_MFEv2"] = reflect.TypeOf(MFN_M11_MFEv2{})
	Types["MFQ_M01v2"] = reflect.TypeOf(MFQ_M01v2{})
	Types["MFQ_M02v2"] = reflect.TypeOf(MFQ_M02v2{})
	Types["MFQ_M03v2"] = reflect.TypeOf(MFQ_M03v2{})
	Types["MFR_M01v2"] = reflect.TypeOf(MFR_M01v2{})
	Types["MFR_M01_MFEv2"] = reflect.TypeOf(MFR_M01_MFEv2{})
	Types["MFR_M02v2"] = reflect.TypeOf(MFR_M02v2{})
	Types["MFR_M02_MFEv2"] = reflect.TypeOf(MFR_M02_MFEv2{})
	Types["MFR_M03v2"] = reflect.TypeOf(MFR_M03v2{})
	Types["MFR_M03_MFEv2"] = reflect.TypeOf(MFR_M03_MFEv2{})
	Types["NMD_N01v2"] = reflect.TypeOf(NMD_N01v2{})
	Types["NMD_N01_NCKv2"] = reflect.TypeOf(NMD_N01_NCKv2{})
	Types["NMD_N01_NCK_NSTv2"] = reflect.TypeOf(NMD_N01_NCK_NSTv2{})
	Types["NMD_N01_NCK_NSCv2"] = reflect.TypeOf(NMD_N01_NCK_NSCv2{})
	Types["NMQ_N02v2"] = reflect.TypeOf(NMQ_N02v2{})
	Types["NMQ_N02_NCKv2"] = reflect.TypeOf(NMQ_N02_NCKv2{})
	Types["NMR_N02v2"] = reflect.TypeOf(NMR_N02v2{})
	Types["NMR_N02_NCKv2"] = reflect.TypeOf(NMR_N02_NCKv2{})
	Types["OMD_O01v2"] = reflect.TypeOf(OMD_O01v2{})
	Types["OMD_O01_PIDv2"] = reflect.TypeOf(OMD_O01_PIDv2{})
	Types["OMD_O01_PID_IN1v2"] = reflect.TypeOf(OMD_O01_PID_IN1v2{})
	Types["OMD_O01_ORC1v2"] = reflect.TypeOf(OMD_O01_ORC1v2{})
	Types["OMD_O01_ORC1_OBXv2"] = reflect.TypeOf(OMD_O01_ORC1_OBXv2{})
	Types["OMD_O01_ORC2v2"] = reflect.TypeOf(OMD_O01_ORC2v2{})
	Types["OMN_O01v2"] = reflect.TypeOf(OMN_O01v2{})
	Types["OMN_O01_PIDv2"] = reflect.TypeOf(OMN_O01_PIDv2{})
	Types["OMN_O01_PID_IN1v2"] = reflect.TypeOf(OMN_O01_PID_IN1v2{})
	Types["OMN_O01_ORCv2"] = reflect.TypeOf(OMN_O01_ORCv2{})
	Types["OMN_O01_ORC_OBXv2"] = reflect.TypeOf(OMN_O01_ORC_OBXv2{})
	Types["OMS_O01v2"] = reflect.TypeOf(OMS_O01v2{})
	Types["OMS_O01_PIDv2"] = reflect.TypeOf(OMS_O01_PIDv2{})
	Types["OMS_O01_PID_IN1v2"] = reflect.TypeOf(OMS_O01_PID_IN1v2{})
	Types["OMS_O01_ORCv2"] = reflect.TypeOf(OMS_O01_ORCv2{})
	Types["OMS_O01_ORC_OBXv2"] = reflect.TypeOf(OMS_O01_ORC_OBXv2{})
	Types["ORD_O02v2"] = reflect.TypeOf(ORD_O02v2{})
	Types["ORD_O02_PIDv2"] = reflect.TypeOf(ORD_O02_PIDv2{})
	Types["ORD_O02_PID_ORC1v2"] = reflect.TypeOf(ORD_O02_PID_ORC1v2{})
	Types["ORD_O02_PID_ORC2v2"] = reflect.TypeOf(ORD_O02_PID_ORC2v2{})
	Types["ORF_R04v2"] = reflect.TypeOf(ORF_R04v2{})
	Types["ORF_R04_PIDv2"] = reflect.TypeOf(ORF_R04_PIDv2{})
	Types["ORF_R04_PID_OBRv2"] = reflect.TypeOf(ORF_R04_PID_OBRv2{})
	Types["ORF_R04_PID_OBR_OBXv2"] = reflect.TypeOf(ORF_R04_PID_OBR_OBXv2{})
	Types["ORM_O01v2"] = reflect.TypeOf(ORM_O01v2{})
	Types["ORM_O01_PIDv2"] = reflect.TypeOf(ORM_O01_PIDv2{})
	Types["ORM_O01_PID_IN1v2"] = reflect.TypeOf(ORM_O01_PID_IN1v2{})
	Types["ORM_O01_ORCv2"] = reflect.TypeOf(ORM_O01_ORCv2{})
	Types["ORM_O01_ORC_OBXv2"] = reflect.TypeOf(ORM_O01_ORC_OBXv2{})
	Types["ORN_O02v2"] = reflect.TypeOf(ORN_O02v2{})
	Types["ORN_O02_PIDv2"] = reflect.TypeOf(ORN_O02_PIDv2{})
	Types["ORN_O02_PID_ORCv2"] = reflect.TypeOf(ORN_O02_PID_ORCv2{})
	Types["ORR_O02v2"] = reflect.TypeOf(ORR_O02v2{})
	Types["ORR_O02_PIDv2"] = reflect.TypeOf(ORR_O02_PIDv2{})
	Types["ORR_O02_PID_ORCv2"] = reflect.TypeOf(ORR_O02_PID_ORCv2{})
	Types["ORU_R01v2"] = reflect.TypeOf(ORU_R01v2{})
	Types["ORU_R01_PIDv2"] = reflect.TypeOf(ORU_R01_PIDv2{})
	Types["ORU_R01_PID_OBRv2"] = reflect.TypeOf(ORU_R01_PID_OBRv2{})
	Types["ORU_R01_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R01_PID_OBR_OBXv2{})
	Types["ORU_R03v2"] = reflect.TypeOf(ORU_R03v2{})
	Types["ORU_R03_PIDv2"] = reflect.TypeOf(ORU_R03_PIDv2{})
	Types["ORU_R03_PID_OBRv2"] = reflect.TypeOf(ORU_R03_PID_OBRv2{})
	Types["ORU_R03_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R03_PID_OBR_OBXv2{})
	Types["ORU_R32v2"] = reflect.TypeOf(ORU_R32v2{})
	Types["ORU_R32_PIDv2"] = reflect.TypeOf(ORU_R32_PIDv2{})
	Types["ORU_R32_PID_OBRv2"] = reflect.TypeOf(ORU_R32_PID_OBRv2{})
	Types["ORU_R32_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R32_PID_OBR_OBXv2{})
	Types["OSQ_Q06v2"] = reflect.TypeOf(OSQ_Q06v2{})
	Types["OSR_Q06v2"] = reflect.TypeOf(OSR_Q06v2{})
	Types["OSR_Q06_PIDv2"] = reflect.TypeOf(OSR_Q06_PIDv2{})
	Types["OSR_Q06_PID_ORCv2"] = reflect.TypeOf(OSR_Q06_PID_ORCv2{})
	Types["PEX_P07v2"] = reflect.TypeOf(PEX_P07v2{})
	Types["PEX_P07_PESv2"] = reflect.TypeOf(PEX_P07_PESv2{})
	Types["PEX_P07_PES_PEOv2"] = reflect.TypeOf(PEX_P07_PES_PEOv2{})
	Types["PEX_P07_PES_PEO_PCRv2"] = reflect.TypeOf(PEX_P07_PES_PEO_PCRv2{})
	Types["PEX_P07_PES_PEO_PCR_RXAv2"] = reflect.TypeOf(PEX_P07_PES_PEO_PCR_RXAv2{})
	Types["PEX_P07_PES_PEO_PCR_NK1v2"] = reflect.TypeOf(PEX_P07_PES_PEO_PCR_NK1v2{})
	Types["PEX_P07_PES_PEO_PCR_NK1_RXAv2"] = reflect.TypeOf(PEX_P07_PES_PEO_PCR_NK1_RXAv2{})
	Types["PEX_P07_PES_PEO_PCR_CSRv2"] = reflect.TypeOf(PEX_P07_PES_PEO_PCR_CSRv2{})
	Types["PGL_PC6v2"] = reflect.TypeOf(PGL_PC6v2{})
	Types["PGL_PC6_GOLv2"] = reflect.TypeOf(PGL_PC6_GOLv2{})
	Types["PGL_PC6_GOL_ROLv2"] = reflect.TypeOf(PGL_PC6_GOL_ROLv2{})
	Types["PGL_PC6_GOL_PTHv2"] = reflect.TypeOf(PGL_PC6_GOL_PTHv2{})
	Types["PGL_PC6_GOL_OBXv2"] = reflect.TypeOf(PGL_PC6_GOL_OBXv2{})
	Types["PGL_PC6_GOL_PRBv2"] = reflect.TypeOf(PGL_PC6_GOL_PRBv2{})
	Types["PGL_PC6_GOL_PRB_ROLv2"] = reflect.TypeOf(PGL_PC6_GOL_PRB_ROLv2{})
	Types["PGL_PC6_GOL_PRB_OBXv2"] = reflect.TypeOf(PGL_PC6_GOL_PRB_OBXv2{})
	Types["PGL_PC6_GOL_ORCv2"] = reflect.TypeOf(PGL_PC6_GOL_ORCv2{})
	Types["PGL_PC6_GOL_ORC_OBXv2"] = reflect.TypeOf(PGL_PC6_GOL_ORC_OBXv2{})
	Types["PIN_I07v2"] = reflect.TypeOf(PIN_I07v2{})
	Types["PIN_I07_PRDv2"] = reflect.TypeOf(PIN_I07_PRDv2{})
	Types["PIN_I07_IN1v2"] = reflect.TypeOf(PIN_I07_IN1v2{})
	Types["PPG_PCGv2"] = reflect.TypeOf(PPG_PCGv2{})
	Types["PPG_PCG_PTHv2"] = reflect.TypeOf(PPG_PCG_PTHv2{})
	Types["PPG_PCG_PTH_ROLv2"] = reflect.TypeOf(PPG_PCG_PTH_ROLv2{})
	Types["PPG_PCG_PTH_GOLv2"] = reflect.TypeOf(PPG_PCG_PTH_GOLv2{})
	Types["PPG_PCG_PTH_GOL_ROLv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_ROLv2{})
	Types["PPG_PCG_PTH_GOL_OBXv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_OBXv2{})
	Types["PPG_PCG_PTH_GOL_PRBv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_PRBv2{})
	Types["PPG_PCG_PTH_GOL_PRB_ROLv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_PRB_ROLv2{})
	Types["PPG_PCG_PTH_GOL_PRB_OBXv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_PRB_OBXv2{})
	Types["PPG_PCG_PTH_GOL_ORCv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_ORCv2{})
	Types["PPG_PCG_PTH_GOL_ORC_OBXv2"] = reflect.TypeOf(PPG_PCG_PTH_GOL_ORC_OBXv2{})
	Types["PPP_PCBv2"] = reflect.TypeOf(PPP_PCBv2{})
	Types["PPP_PCB_PTHv2"] = reflect.TypeOf(PPP_PCB_PTHv2{})
	Types["PPP_PCB_PTH_ROLv2"] = reflect.TypeOf(PPP_PCB_PTH_ROLv2{})
	Types["PPP_PCB_PTH_PRBv2"] = reflect.TypeOf(PPP_PCB_PTH_PRBv2{})
	Types["PPP_PCB_PTH_PRB_ROLv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_ROLv2{})
	Types["PPP_PCB_PTH_PRB_OBXv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_OBXv2{})
	Types["PPP_PCB_PTH_PRB_GOLv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_GOLv2{})
	Types["PPP_PCB_PTH_PRB_GOL_ROLv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_GOL_ROLv2{})
	Types["PPP_PCB_PTH_PRB_GOL_OBXv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_GOL_OBXv2{})
	Types["PPP_PCB_PTH_PRB_ORCv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_ORCv2{})
	Types["PPP_PCB_PTH_PRB_ORC_OBXv2"] = reflect.TypeOf(PPP_PCB_PTH_PRB_ORC_OBXv2{})
	Types["PPR_PC1v2"] = reflect.TypeOf(PPR_PC1v2{})
	Types["PPR_PC1_PRBv2"] = reflect.TypeOf(PPR_PC1_PRBv2{})
	Types["PPR_PC1_PRB_ROLv2"] = reflect.TypeOf(PPR_PC1_PRB_ROLv2{})
	Types["PPR_PC1_PRB_PTHv2"] = reflect.TypeOf(PPR_PC1_PRB_PTHv2{})
	Types["PPR_PC1_PRB_OBXv2"] = reflect.TypeOf(PPR_PC1_PRB_OBXv2{})
	Types["PPR_PC1_PRB_GOLv2"] = reflect.TypeOf(PPR_PC1_PRB_GOLv2{})
	Types["PPR_PC1_PRB_GOL_ROLv2"] = reflect.TypeOf(PPR_PC1_PRB_GOL_ROLv2{})
	Types["PPR_PC1_PRB_GOL_OBXv2"] = reflect.TypeOf(PPR_PC1_PRB_GOL_OBXv2{})
	Types["PPR_PC1_PRB_ORCv2"] = reflect.TypeOf(PPR_PC1_PRB_ORCv2{})
	Types["PPR_PC1_PRB_ORC_OBXv2"] = reflect.TypeOf(PPR_PC1_PRB_ORC_OBXv2{})
	Types["PPT_PCLv2"] = reflect.TypeOf(PPT_PCLv2{})
	Types["PPT_PCL_PIDv2"] = reflect.TypeOf(PPT_PCL_PIDv2{})
	Types["PPT_PCL_PID_PTHv2"] = reflect.TypeOf(PPT_PCL_PID_PTHv2{})
	Types["PPT_PCL_PID_PTH_ROLv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_ROLv2{})
	Types["PPT_PCL_PID_PTH_GOLv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOLv2{})
	Types["PPT_PCL_PID_PTH_GOL_ROLv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_ROLv2{})
	Types["PPT_PCL_PID_PTH_GOL_OBXv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_OBXv2{})
	Types["PPT_PCL_PID_PTH_GOL_PRBv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_PRBv2{})
	Types["PPT_PCL_PID_PTH_GOL_PRB_ROLv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_PRB_ROLv2{})
	Types["PPT_PCL_PID_PTH_GOL_PRB_OBXv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_PRB_OBXv2{})
	Types["PPT_PCL_PID_PTH_GOL_ORCv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_ORCv2{})
	Types["PPT_PCL_PID_PTH_GOL_ORC_OBXv2"] = reflect.TypeOf(PPT_PCL_PID_PTH_GOL_ORC_OBXv2{})
	Types["PPV_PCAv2"] = reflect.TypeOf(PPV_PCAv2{})
	Types["PPV_PCA_PIDv2"] = reflect.TypeOf(PPV_PCA_PIDv2{})
	Types["PPV_PCA_PID_GOLv2"] = reflect.TypeOf(PPV_PCA_PID_GOLv2{})
	Types["PPV_PCA_PID_GOL_ROLv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_ROLv2{})
	Types["PPV_PCA_PID_GOL_PTHv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_PTHv2{})
	Types["PPV_PCA_PID_GOL_OBXv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_OBXv2{})
	Types["PPV_PCA_PID_GOL_PRBv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_PRBv2{})
	Types["PPV_PCA_PID_GOL_PRB_ROLv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_PRB_ROLv2{})
	Types["PPV_PCA_PID_GOL_PRB_OBXv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_PRB_OBXv2{})
	Types["PPV_PCA_PID_GOL_ORCv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_ORCv2{})
	Types["PPV_PCA_PID_GOL_ORC_OBXv2"] = reflect.TypeOf(PPV_PCA_PID_GOL_ORC_OBXv2{})
	Types["PRR_PC5v2"] = reflect.TypeOf(PRR_PC5v2{})
	Types["PRR_PC5_PIDv2"] = reflect.TypeOf(PRR_PC5_PIDv2{})
	Types["PRR_PC5_PID_PRBv2"] = reflect.TypeOf(PRR_PC5_PID_PRBv2{})
	Types["PRR_PC5_PID_PRB_ROLv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_ROLv2{})
	Types["PRR_PC5_PID_PRB_PTHv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_PTHv2{})
	Types["PRR_PC5_PID_PRB_OBXv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_OBXv2{})
	Types["PRR_PC5_PID_PRB_GOLv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_GOLv2{})
	Types["PRR_PC5_PID_PRB_GOL_ROLv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_GOL_ROLv2{})
	Types["PRR_PC5_PID_PRB_GOL_OBXv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_GOL_OBXv2{})
	Types["PRR_PC5_PID_PRB_ORCv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_ORCv2{})
	Types["PRR_PC5_PID_PRB_ORC_OBXv2"] = reflect.TypeOf(PRR_PC5_PID_PRB_ORC_OBXv2{})
	Types["PTR_PCFv2"] = reflect.TypeOf(PTR_PCFv2{})
	Types["PTR_PCF_PIDv2"] = reflect.TypeOf(PTR_PCF_PIDv2{})
	Types["PTR_PCF_PID_PTHv2"] = reflect.TypeOf(PTR_PCF_PID_PTHv2{})
	Types["PTR_PCF_PID_PTH_ROLv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_ROLv2{})
	Types["PTR_PCF_PID_PTH_PRBv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRBv2{})
	Types["PTR_PCF_PID_PTH_PRB_ROLv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_ROLv2{})
	Types["PTR_PCF_PID_PTH_PRB_OBXv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_OBXv2{})
	Types["PTR_PCF_PID_PTH_PRB_GOLv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_GOLv2{})
	Types["PTR_PCF_PID_PTH_PRB_GOL_ROLv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_GOL_ROLv2{})
	Types["PTR_PCF_PID_PTH_PRB_GOL_OBXv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_GOL_OBXv2{})
	Types["PTR_PCF_PID_PTH_PRB_ORCv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_ORCv2{})
	Types["PTR_PCF_PID_PTH_PRB_ORC_OBXv2"] = reflect.TypeOf(PTR_PCF_PID_PTH_PRB_ORC_OBXv2{})
	Types["QCK_Q02v2"] = reflect.TypeOf(QCK_Q02v2{})
	Types["QRY_A19v2"] = reflect.TypeOf(QRY_A19v2{})
	Types["QRY_P04v2"] = reflect.TypeOf(QRY_P04v2{})
	Types["QRY_PC4v2"] = reflect.TypeOf(QRY_PC4v2{})
	Types["QRY_Q01v2"] = reflect.TypeOf(QRY_Q01v2{})
	Types["QRY_Q02v2"] = reflect.TypeOf(QRY_Q02v2{})
	Types["QRY_R02v2"] = reflect.TypeOf(QRY_R02v2{})
	Types["QRY_T12v2"] = reflect.TypeOf(QRY_T12v2{})
	Types["RAR_RARv2"] = reflect.TypeOf(RAR_RARv2{})
	Types["RAR_RAR_QRDv2"] = reflect.TypeOf(RAR_RAR_QRDv2{})
	Types["RAR_RAR_QRD_ORCv2"] = reflect.TypeOf(RAR_RAR_QRD_ORCv2{})
	Types["RAS_O01v2"] = reflect.TypeOf(RAS_O01v2{})
	Types["RAS_O01_PIDv2"] = reflect.TypeOf(RAS_O01_PIDv2{})
	Types["RAS_O01_ORCv2"] = reflect.TypeOf(RAS_O01_ORCv2{})
	Types["RAS_O01_ORC_RXOv2"] = reflect.TypeOf(RAS_O01_ORC_RXOv2{})
	Types["RAS_O01_ORC_RXO_RXCv2"] = reflect.TypeOf(RAS_O01_ORC_RXO_RXCv2{})
	Types["RAS_O01_ORC_RXEv2"] = reflect.TypeOf(RAS_O01_ORC_RXEv2{})
	Types["RAS_O01_ORC_OBXv2"] = reflect.TypeOf(RAS_O01_ORC_OBXv2{})
	Types["RCI_I05v2"] = reflect.TypeOf(RCI_I05v2{})
	Types["RCI_I05_PRDv2"] = reflect.TypeOf(RCI_I05_PRDv2{})
	Types["RCI_I05_OBRv2"] = reflect.TypeOf(RCI_I05_OBRv2{})
	Types["RCI_I05_OBR_OBXv2"] = reflect.TypeOf(RCI_I05_OBR_OBXv2{})
	Types["RCL_I06v2"] = reflect.TypeOf(RCL_I06v2{})
	Types["RCL_I06_PRDv2"] = reflect.TypeOf(RCL_I06_PRDv2{})
	Types["RDE_O01v2"] = reflect.TypeOf(RDE_O01v2{})
	Types["RDE_O01_PIDv2"] = reflect.TypeOf(RDE_O01_PIDv2{})
	Types["RDE_O01_PID_IN1v2"] = reflect.TypeOf(RDE_O01_PID_IN1v2{})
	Types["RDE_O01_ORCv2"] = reflect.TypeOf(RDE_O01_ORCv2{})
	Types["RDE_O01_ORC_RXOv2"] = reflect.TypeOf(RDE_O01_ORC_RXOv2{})
	Types["RDE_O01_ORC_RXO_RXCv2"] = reflect.TypeOf(RDE_O01_ORC_RXO_RXCv2{})
	Types["RDE_O01_ORC_OBXv2"] = reflect.TypeOf(RDE_O01_ORC_OBXv2{})
	Types["RDO_O01v2"] = reflect.TypeOf(RDO_O01v2{})
	Types["RDO_O01_PIDv2"] = reflect.TypeOf(RDO_O01_PIDv2{})
	Types["RDO_O01_PID_IN1v2"] = reflect.TypeOf(RDO_O01_PID_IN1v2{})
	Types["RDO_O01_ORCv2"] = reflect.TypeOf(RDO_O01_ORCv2{})
	Types["RDO_O01_ORC_RXCv2"] = reflect.TypeOf(RDO_O01_ORC_RXCv2{})
	Types["RDO_O01_ORC_OBXv2"] = reflect.TypeOf(RDO_O01_ORC_OBXv2{})
	Types["RDR_RDRv2"] = reflect.TypeOf(RDR_RDRv2{})
	Types["RDR_RDR_QRDv2"] = reflect.TypeOf(RDR_RDR_QRDv2{})
	Types["RDR_RDR_QRD_ORCv2"] = reflect.TypeOf(RDR_RDR_QRD_ORCv2{})
	Types["RDR_RDR_QRD_ORC_RXDv2"] = reflect.TypeOf(RDR_RDR_QRD_ORC_RXDv2{})
	Types["RDS_O01v2"] = reflect.TypeOf(RDS_O01v2{})
	Types["RDS_O01_PIDv2"] = reflect.TypeOf(RDS_O01_PIDv2{})
	Types["RDS_O01_ORCv2"] = reflect.TypeOf(RDS_O01_ORCv2{})
	Types["RDS_O01_ORC_RXOv2"] = reflect.TypeOf(RDS_O01_ORC_RXOv2{})
	Types["RDS_O01_ORC_RXO_RXCv2"] = reflect.TypeOf(RDS_O01_ORC_RXO_RXCv2{})
	Types["RDS_O01_ORC_RXEv2"] = reflect.TypeOf(RDS_O01_ORC_RXEv2{})
	Types["RDS_O01_ORC_OBXv2"] = reflect.TypeOf(RDS_O01_ORC_OBXv2{})
	Types["REF_I12v2"] = reflect.TypeOf(REF_I12v2{})
	Types["REF_I12_PRDv2"] = reflect.TypeOf(REF_I12_PRDv2{})
	Types["REF_I12_IN1v2"] = reflect.TypeOf(REF_I12_IN1v2{})
	Types["REF_I12_PR1v2"] = reflect.TypeOf(REF_I12_PR1v2{})
	Types["REF_I12_OBRv2"] = reflect.TypeOf(REF_I12_OBRv2{})
	Types["REF_I12_OBR_OBXv2"] = reflect.TypeOf(REF_I12_OBR_OBXv2{})
	Types["RER_RERv2"] = reflect.TypeOf(RER_RERv2{})
	Types["RER_RER_QRDv2"] = reflect.TypeOf(RER_RER_QRDv2{})
	Types["RER_RER_QRD_ORCv2"] = reflect.TypeOf(RER_RER_QRD_ORCv2{})
	Types["RGR_RGRv2"] = reflect.TypeOf(RGR_RGRv2{})
	Types["RGR_RGR_QRDv2"] = reflect.TypeOf(RGR_RGR_QRDv2{})
	Types["RGR_RGR_QRD_ORCv2"] = reflect.TypeOf(RGR_RGR_QRD_ORCv2{})
	Types["RGR_RGR_QRD_ORC_RXEv2"] = reflect.TypeOf(RGR_RGR_QRD_ORC_RXEv2{})
	Types["RGV_O01v2"] = reflect.TypeOf(RGV_O01v2{})
	Types["RGV_O01_PIDv2"] = reflect.TypeOf(RGV_O01_PIDv2{})
	Types["RGV_O01_ORCv2"] = reflect.TypeOf(RGV_O01_ORCv2{})
	Types["RGV_O01_ORC_RXCv2"] = reflect.TypeOf(RGV_O01_ORC_RXCv2{})
	Types["RGV_O01_ORC_RXEv2"] = reflect.TypeOf(RGV_O01_ORC_RXEv2{})
	Types["RGV_O01_ORC_RXGv2"] = reflect.TypeOf(RGV_O01_ORC_RXGv2{})
	Types["RGV_O01_ORC_RXG_OBXv2"] = reflect.TypeOf(RGV_O01_ORC_RXG_OBXv2{})
	Types["ROR_RORv2"] = reflect.TypeOf(ROR_RORv2{})
	Types["ROR_ROR_QRDv2"] = reflect.TypeOf(ROR_ROR_QRDv2{})
	Types["ROR_ROR_QRD_ORCv2"] = reflect.TypeOf(ROR_ROR_QRD_ORCv2{})
	Types["RPA_I08v2"] = reflect.TypeOf(RPA_I08v2{})
	Types["RPA_I08_PRDv2"] = reflect.TypeOf(RPA_I08_PRDv2{})
	Types["RPA_I08_IN1v2"] = reflect.TypeOf(RPA_I08_IN1v2{})
	Types["RPA_I08_PR1v2"] = reflect.TypeOf(RPA_I08_PR1v2{})
	Types["RPA_I08_OBRv2"] = reflect.TypeOf(RPA_I08_OBRv2{})
	Types["RPA_I08_OBR_OBXv2"] = reflect.TypeOf(RPA_I08_OBR_OBXv2{})
	Types["RPI_I01v2"] = reflect.TypeOf(RPI_I01v2{})
	Types["RPI_I01_PRDv2"] = reflect.TypeOf(RPI_I01_PRDv2{})
	Types["RPI_I01_IN1v2"] = reflect.TypeOf(RPI_I01_IN1v2{})
	Types["RPL_I02v2"] = reflect.TypeOf(RPL_I02v2{})
	Types["RPL_I02_PRDv2"] = reflect.TypeOf(RPL_I02_PRDv2{})
	Types["RQA_I08v2"] = reflect.TypeOf(RQA_I08v2{})
	Types["RQA_I08_PRDv2"] = reflect.TypeOf(RQA_I08_PRDv2{})
	Types["RQA_I08_IN1v2"] = reflect.TypeOf(RQA_I08_IN1v2{})
	Types["RQA_I08_PR1v2"] = reflect.TypeOf(RQA_I08_PR1v2{})
	Types["RQA_I08_OBRv2"] = reflect.TypeOf(RQA_I08_OBRv2{})
	Types["RQA_I08_OBR_OBXv2"] = reflect.TypeOf(RQA_I08_OBR_OBXv2{})
	Types["RQC_I05v2"] = reflect.TypeOf(RQC_I05v2{})
	Types["RQC_I05_PRDv2"] = reflect.TypeOf(RQC_I05_PRDv2{})
	Types["RQC_I06v2"] = reflect.TypeOf(RQC_I06v2{})
	Types["RQC_I06_PRDv2"] = reflect.TypeOf(RQC_I06_PRDv2{})
	Types["RQI_I01v2"] = reflect.TypeOf(RQI_I01v2{})
	Types["RQI_I01_PRDv2"] = reflect.TypeOf(RQI_I01_PRDv2{})
	Types["RQI_I01_IN1v2"] = reflect.TypeOf(RQI_I01_IN1v2{})
	Types["RQP_I04v2"] = reflect.TypeOf(RQP_I04v2{})
	Types["RQP_I04_PRDv2"] = reflect.TypeOf(RQP_I04_PRDv2{})
	Types["RQQ_Q01v2"] = reflect.TypeOf(RQQ_Q01v2{})
	Types["RRA_O02v2"] = reflect.TypeOf(RRA_O02v2{})
	Types["RRA_O02_PIDv2"] = reflect.TypeOf(RRA_O02_PIDv2{})
	Types["RRA_O02_PID_ORCv2"] = reflect.TypeOf(RRA_O02_PID_ORCv2{})
	Types["RRA_O02_PID_ORC_RXAv2"] = reflect.TypeOf(RRA_O02_PID_ORC_RXAv2{})
	Types["RRD_O02v2"] = reflect.TypeOf(RRD_O02v2{})
	Types["RRD_O02_PIDv2"] = reflect.TypeOf(RRD_O02_PIDv2{})
	Types["RRD_O02_PID_ORCv2"] = reflect.TypeOf(RRD_O02_PID_ORCv2{})
	Types["RRG_O02v2"] = reflect.TypeOf(RRG_O02v2{})
	Types["RRG_O02_PIDv2"] = reflect.TypeOf(RRG_O02_PIDv2{})
	Types["RRG_O02_PID_ORCv2"] = reflect.TypeOf(RRG_O02_PID_ORCv2{})
	Types["RRI_I12v2"] = reflect.TypeOf(RRI_I12v2{})
	Types["RRI_I12_PRDv2"] = reflect.TypeOf(RRI_I12_PRDv2{})
	Types["RRI_I12_PR1v2"] = reflect.TypeOf(RRI_I12_PR1v2{})
	Types["RRI_I12_OBRv2"] = reflect.TypeOf(RRI_I12_OBRv2{})
	Types["RRI_I12_OBR_OBXv2"] = reflect.TypeOf(RRI_I12_OBR_OBXv2{})
	Types["RRO_O02v2"] = reflect.TypeOf(RRO_O02v2{})
	Types["RRO_O02_PIDv2"] = reflect.TypeOf(RRO_O02_PIDv2{})
	Types["RRO_O02_PID_ORCv2"] = reflect.TypeOf(RRO_O02_PID_ORCv2{})
	Types["SIU_S12v2"] = reflect.TypeOf(SIU_S12v2{})
	Types["SIU_S12_PIDv2"] = reflect.TypeOf(SIU_S12_PIDv2{})
	Types["SIU_S12_RGSv2"] = reflect.TypeOf(SIU_S12_RGSv2{})
	Types["SIU_S12_RGS_AISv2"] = reflect.TypeOf(SIU_S12_RGS_AISv2{})
	Types["SIU_S12_RGS_AIGv2"] = reflect.TypeOf(SIU_S12_RGS_AIGv2{})
	Types["SIU_S12_RGS_AILv2"] = reflect.TypeOf(SIU_S12_RGS_AILv2{})
	Types["SIU_S12_RGS_AIPv2"] = reflect.TypeOf(SIU_S12_RGS_AIPv2{})
	Types["SPQ_Q01v2"] = reflect.TypeOf(SPQ_Q01v2{})
	Types["SQM_S25v2"] = reflect.TypeOf(SQM_S25v2{})
	Types["SQM_S25_RGSv2"] = reflect.TypeOf(SQM_S25_RGSv2{})
	Types["SQM_S25_RGS_AISv2"] = reflect.TypeOf(SQM_S25_RGS_AISv2{})
	Types["SQM_S25_RGS_AIGv2"] = reflect.TypeOf(SQM_S25_RGS_AIGv2{})
	Types["SQM_S25_RGS_AIPv2"] = reflect.TypeOf(SQM_S25_RGS_AIPv2{})
	Types["SQM_S25_RGS_AILv2"] = reflect.TypeOf(SQM_S25_RGS_AILv2{})
	Types["SQR_S25v2"] = reflect.TypeOf(SQR_S25v2{})
	Types["SQR_S25_SCHv2"] = reflect.TypeOf(SQR_S25_SCHv2{})
	Types["SQR_S25_SCH_RGSv2"] = reflect.TypeOf(SQR_S25_SCH_RGSv2{})
	Types["SQR_S25_SCH_RGS_AISv2"] = reflect.TypeOf(SQR_S25_SCH_RGS_AISv2{})
	Types["SQR_S25_SCH_RGS_AIGv2"] = reflect.TypeOf(SQR_S25_SCH_RGS_AIGv2{})
	Types["SQR_S25_SCH_RGS_AIPv2"] = reflect.TypeOf(SQR_S25_SCH_RGS_AIPv2{})
	Types["SQR_S25_SCH_RGS_AILv2"] = reflect.TypeOf(SQR_S25_SCH_RGS_AILv2{})
	Types["SRM_S01v2"] = reflect.TypeOf(SRM_S01v2{})
	Types["SRM_S01_PIDv2"] = reflect.TypeOf(SRM_S01_PIDv2{})
	Types["SRM_S01_RGSv2"] = reflect.TypeOf(SRM_S01_RGSv2{})
	Types["SRM_S01_RGS_AISv2"] = reflect.TypeOf(SRM_S01_RGS_AISv2{})
	Types["SRM_S01_RGS_AIGv2"] = reflect.TypeOf(SRM_S01_RGS_AIGv2{})
	Types["SRM_S01_RGS_AILv2"] = reflect.TypeOf(SRM_S01_RGS_AILv2{})
	Types["SRM_S01_RGS_AIPv2"] = reflect.TypeOf(SRM_S01_RGS_AIPv2{})
	Types["SRR_S01v2"] = reflect.TypeOf(SRR_S01v2{})
	Types["SRR_S01_PIDv2"] = reflect.TypeOf(SRR_S01_PIDv2{})
	Types["SRR_S01_RGSv2"] = reflect.TypeOf(SRR_S01_RGSv2{})
	Types["SRR_S01_RGS_AISv2"] = reflect.TypeOf(SRR_S01_RGS_AISv2{})
	Types["SRR_S01_RGS_AIGv2"] = reflect.TypeOf(SRR_S01_RGS_AIGv2{})
	Types["SRR_S01_RGS_AILv2"] = reflect.TypeOf(SRR_S01_RGS_AILv2{})
	Types["SRR_S01_RGS_AIPv2"] = reflect.TypeOf(SRR_S01_RGS_AIPv2{})
	Types["SUR_P09v2"] = reflect.TypeOf(SUR_P09v2{})
	Types["SUR_P09_FACv2"] = reflect.TypeOf(SUR_P09_FACv2{})
	Types["SUR_P09_FAC_PSH1v2"] = reflect.TypeOf(SUR_P09_FAC_PSH1v2{})
	Types["SUR_P09_FAC_FAC2v2"] = reflect.TypeOf(SUR_P09_FAC_FAC2v2{})
	Types["TBR_Q01v2"] = reflect.TypeOf(TBR_Q01v2{})
	Types["UDM_Q05v2"] = reflect.TypeOf(UDM_Q05v2{})
	Types["VQQ_Q01v2"] = reflect.TypeOf(VQQ_Q01v2{})
	Types["VXQ_V01v2"] = reflect.TypeOf(VXQ_V01v2{})
	Types["VXR_V03v2"] = reflect.TypeOf(VXR_V03v2{})
	Types["VXR_V03_IN1v2"] = reflect.TypeOf(VXR_V03_IN1v2{})
	Types["VXR_V03_RXAv2"] = reflect.TypeOf(VXR_V03_RXAv2{})
	Types["VXR_V03_RXA_OBXv2"] = reflect.TypeOf(VXR_V03_RXA_OBXv2{})
	Types["VXU_V04v2"] = reflect.TypeOf(VXU_V04v2{})
	Types["VXU_V04_IN1v2"] = reflect.TypeOf(VXU_V04_IN1v2{})
	Types["VXU_V04_RXAv2"] = reflect.TypeOf(VXU_V04_RXAv2{})
	Types["VXU_V04_RXA_OBXv2"] = reflect.TypeOf(VXU_V04_RXA_OBXv2{})
	Types["VXX_V02v2"] = reflect.TypeOf(VXX_V02v2{})
	Types["VXX_V02_PIDv2"] = reflect.TypeOf(VXX_V02_PIDv2{})
}
var FollowSets = map[string]StringSet{
	"ACKv2.err": StringSet{
	},
	"ACKv2.msa": StringSet{
		"ERR": true,
	},
	"ACKv2.msh": StringSet{
		"ERR": true,
		"MSA": true,
	},
	"ADR_A19_PID_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PID_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PID_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.acc": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.dg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pr1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.ub1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.ub2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.dsc": StringSet{
	},
	"ADR_A19v2.err": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.msa": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.qrd": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_PR1v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_PR1v2.rol": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.db1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.drg": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pd1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A01v2.ub2": StringSet{
	},
	"ADT_A02v2.db1": StringSet{
		"DB1": true,
		"OBX": true,
	},
	"ADT_A02v2.evn": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.msh": StringSet{
		"DB1": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A02v2.pd1": StringSet{
		"DB1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.pid": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.pv1": StringSet{
		"DB1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A02v2.pv2": StringSet{
		"DB1": true,
		"OBX": true,
	},
	"ADT_A03_PR1v2.pr1": StringSet{
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A03_PR1v2.rol": StringSet{
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A03v2.db1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A03v2.dg1": StringSet{
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A03v2.drg": StringSet{
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A03v2.evn": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"ADT_A03v2.msh": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"ADT_A03v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A03v2.pd1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"ADT_A03v2.pid": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"ADT_A03v2.pv1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
	},
	"ADT_A03v2.pv2": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"ADT_A04_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A04v2.ub2": StringSet{
	},
	"ADT_A05_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A05v2.ub2": StringSet{
	},
	"ADT_A06_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_PR1v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_PR1v2.rol": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.db1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.drg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.drg2": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.mrg": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pd1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A06v2.ub2": StringSet{
	},
	"ADT_A07_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.mrg": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A07v2.ub2": StringSet{
	},
	"ADT_A08_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A08v2.ub2": StringSet{
	},
	"ADT_A09v2.db1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
	},
	"ADT_A09v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A09v2.evn": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.msh": StringSet{
		"DB1": true,
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A09v2.pd1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.pid": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.pv1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A09v2.pv2": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
	},
	"ADT_A10v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A10v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A10v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A10v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A11v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A11v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A11v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A11v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A12v2.db1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
	},
	"ADT_A12v2.dg1": StringSet{
	},
	"ADT_A12v2.evn": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.msh": StringSet{
		"DB1": true,
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A12v2.pd1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.pid": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.pv1": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A12v2.pv2": StringSet{
		"DB1": true,
		"DG1": true,
		"OBX": true,
	},
	"ADT_A13_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A13v2.ub2": StringSet{
	},
	"ADT_A14_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A14v2.ub2": StringSet{
	},
	"ADT_A15v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A15v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A15v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A15v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A16v2.db1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A16v2.dg1": StringSet{
		"DRG": true,
	},
	"ADT_A16v2.drg": StringSet{
	},
	"ADT_A16v2.evn": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.msh": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.obx": StringSet{
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A16v2.pd1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.pid": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.pv1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A16v2.pv2": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A17v2.db11": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.db12": StringSet{
		"DB1": true,
		"OBX": true,
	},
	"ADT_A17v2.evn": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.msh": StringSet{
		"DB1": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.obx1": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.obx2": StringSet{
		"OBX": true,
	},
	"ADT_A17v2.pd11": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pd12": StringSet{
		"DB1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pid1": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pid2": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv11": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv12": StringSet{
		"DB1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A17v2.pv21": StringSet{
		"DB1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv22": StringSet{
		"DB1": true,
		"OBX": true,
	},
	"ADT_A18v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A18v2.mrg": StringSet{
		"PV1": true,
	},
	"ADT_A18v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A18v2.pd1": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A18v2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PV1": true,
	},
	"ADT_A18v2.pv1": StringSet{
	},
	"ADT_A20v2.evn": StringSet{
		"NPU": true,
	},
	"ADT_A20v2.msh": StringSet{
		"EVN": true,
		"NPU": true,
	},
	"ADT_A20v2.npu": StringSet{
	},
	"ADT_A21v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A21v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A21v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A22v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A22v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A22v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A23v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A23v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A23v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A24v2.db11": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.db12": StringSet{
		"DB1": true,
	},
	"ADT_A24v2.evn": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.msh": StringSet{
		"DB1": true,
		"EVN": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pd11": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pd12": StringSet{
		"DB1": true,
		"PV1": true,
	},
	"ADT_A24v2.pid1": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pid2": StringSet{
		"DB1": true,
		"PD1": true,
		"PV1": true,
	},
	"ADT_A24v2.pv11": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pv12": StringSet{
		"DB1": true,
	},
	"ADT_A25v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A25v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A25v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A26v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A26v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A26v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A27v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A27v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A27v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A28_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A28v2.ub2": StringSet{
	},
	"ADT_A29v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A29v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A29v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A30v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A30v2.mrg": StringSet{
	},
	"ADT_A30v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A30v2.pd1": StringSet{
		"MRG": true,
	},
	"ADT_A30v2.pid": StringSet{
		"MRG": true,
		"PD1": true,
	},
	"ADT_A31_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A31v2.ub2": StringSet{
	},
	"ADT_A32v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A32v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A32v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A33v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A33v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A33v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A34v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A34v2.mrg": StringSet{
	},
	"ADT_A34v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A34v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A35v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A35v2.mrg": StringSet{
	},
	"ADT_A35v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A35v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A36v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A36v2.mrg": StringSet{
	},
	"ADT_A36v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A36v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A37v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pid1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pid2": StringSet{
		"PV1": true,
	},
	"ADT_A37v2.pv11": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pv12": StringSet{
	},
	"ADT_A38v2.db1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A38v2.dg1": StringSet{
		"DG1": true,
		"DRG": true,
	},
	"ADT_A38v2.drg": StringSet{
	},
	"ADT_A38v2.evn": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A38v2.msh": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A38v2.obx": StringSet{
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A38v2.pd1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A38v2.pid": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A38v2.pv1": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A38v2.pv2": StringSet{
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"OBX": true,
	},
	"ADT_A39_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A39_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A39_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A39_PIDv2.pv1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A39v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A39v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pv1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A43_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A43_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A43_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A43v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A43v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A45_MRGv2.mrg": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A45_MRGv2.pv1": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A45v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A45v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A45v2.pd1": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A45v2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PV1": true,
	},
	"ADT_A50v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A50v2.mrg": StringSet{
		"PV1": true,
	},
	"ADT_A50v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A50v2.pd1": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A50v2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PV1": true,
	},
	"ADT_A50v2.pv1": StringSet{
	},
	"ARD_A19_PID_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PID_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PID_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PID_PR1v2.pr1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PID_PR1v2.rol": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.acc": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.db1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.dg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.drg": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.pd1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.ub1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19_PIDv2.ub2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19v2.dsc": StringSet{
	},
	"ARD_A19v2.err": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"QRF": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19v2.msa": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"QRF": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"QRF": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19v2.qrd": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRF": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"ARD_A19v2.qrf": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_PR1v2.pr1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_PR1v2.rol": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.acc": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.db1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.dg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.drg": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.ub1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.ub2": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.pd1": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PD1": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P02_PIDv2.db1": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P02_PIDv2.pd1": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P02_PIDv2.pid": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P02_PIDv2.pv1": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P02v2.evn": StringSet{
		"DB1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P02v2.msh": StringSet{
		"DB1": true,
		"EVN": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"BAR_P06_PIDv2.pid": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P06_PIDv2.pv1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P06v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P06v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"CRM_C01_PIDv2.csp": StringSet{
		"CSP": true,
		"CSR": true,
		"PID": true,
		"PV1": true,
	},
	"CRM_C01_PIDv2.csr": StringSet{
		"CSP": true,
		"CSR": true,
		"PID": true,
		"PV1": true,
	},
	"CRM_C01_PIDv2.pid": StringSet{
		"CSP": true,
		"CSR": true,
		"PID": true,
		"PV1": true,
	},
	"CRM_C01_PIDv2.pv1": StringSet{
		"CSP": true,
		"CSR": true,
		"PID": true,
		"PV1": true,
	},
	"CRM_C01v2.msh": StringSet{
		"CSP": true,
		"CSR": true,
		"PID": true,
		"PV1": true,
	},
	"CSU_C09_PID_CSP_CSS_OBRv2.obr": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSS_OBRv2.obx": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSS_OBRv2.orc": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSS_ORC_RXAv2.rxa": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSS_ORC_RXAv2.rxr": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSS_ORCv2.orc": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSP_CSSv2.css": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PID_CSPv2.csp": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.csr": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.nte": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.pd1": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.pid": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.pv1": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09_PIDv2.pv2": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"CSU_C09v2.msh": StringSet{
		"CSP": true,
		"CSR": true,
		"CSS": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"DFT_P03_FT1_PR1v2.pr1": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
	},
	"DFT_P03_FT1_PR1v2.rol": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
	},
	"DFT_P03_FT1v2.ft1": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"ROL": true,
	},
	"DFT_P03_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03v2.acc": StringSet{
	},
	"DFT_P03v2.db1": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"DFT_P03v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03v2.drg": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03v2.evn": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"DFT_P03v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
	},
	"DFT_P03v2.msh": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"EVN": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PD1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"DFT_P03v2.obx": StringSet{
		"ACC": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"DFT_P03v2.pd1": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"DFT_P03v2.pid": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PD1": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
	},
	"DFT_P03v2.pv1": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"ROL": true,
	},
	"DFT_P03v2.pv2": StringSet{
		"ACC": true,
		"DB1": true,
		"DG1": true,
		"DRG": true,
		"FT1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"ROL": true,
	},
	"DOC_T12_PIDv2.evn": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DOC_T12_PIDv2.obx": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DOC_T12_PIDv2.pid": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DOC_T12_PIDv2.pv1": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DOC_T12_PIDv2.txa": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DOC_T12v2.dsc": StringSet{
	},
	"DOC_T12v2.err": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"QRD": true,
		"TXA": true,
	},
	"DOC_T12v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"QRD": true,
		"TXA": true,
	},
	"DOC_T12v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"MSA": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"QRD": true,
		"TXA": true,
	},
	"DOC_T12v2.qrd": StringSet{
		"DSC": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"DSR_P04v2.dsc": StringSet{
	},
	"DSR_P04v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_P04v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_P04v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q01v2.dsc": StringSet{
	},
	"DSR_Q01v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q01v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.qak": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_Q01v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q03v2.dsc": StringSet{
	},
	"DSR_Q03v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q03v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QAK": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.qak": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_Q03v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_R03v2.dsc": StringSet{
	},
	"DSR_R03v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_R03v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_R03v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_R03v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_R03v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"EDR_Q01v2.dsc": StringSet{
	},
	"EDR_Q01v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"EDR_Q01v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QAK": true,
	},
	"EDR_Q01v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QAK": true,
	},
	"EDR_Q01v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QAK": true,
	},
	"EDR_Q01v2.qak": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"EQQ_Q01v2.dsc": StringSet{
	},
	"EQQ_Q01v2.eql": StringSet{
		"DSC": true,
	},
	"EQQ_Q01v2.msh": StringSet{
		"DSC": true,
		"EQL": true,
	},
	"ERP_Q01v2.dsc": StringSet{
	},
	"ERP_Q01v2.erq": StringSet{
		"DSC": true,
	},
	"ERP_Q01v2.err": StringSet{
		"DSC": true,
		"ERQ": true,
		"QAK": true,
	},
	"ERP_Q01v2.msa": StringSet{
		"DSC": true,
		"ERQ": true,
		"ERR": true,
		"QAK": true,
	},
	"ERP_Q01v2.msh": StringSet{
		"DSC": true,
		"ERQ": true,
		"ERR": true,
		"MSA": true,
		"QAK": true,
	},
	"ERP_Q01v2.qak": StringSet{
		"DSC": true,
		"ERQ": true,
	},
	"MCF_Q02v2.msa": StringSet{
	},
	"MCF_Q02v2.msh": StringSet{
		"MSA": true,
	},
	"MDM_T01v2.evn": StringSet{
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"MDM_T01v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"MDM_T01v2.pid": StringSet{
		"PV1": true,
		"TXA": true,
	},
	"MDM_T01v2.pv1": StringSet{
		"TXA": true,
	},
	"MDM_T01v2.txa": StringSet{
	},
	"MDM_T02v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"MDM_T02v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"TXA": true,
	},
	"MDM_T02v2.obx": StringSet{
		"OBX": true,
	},
	"MDM_T02v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"TXA": true,
	},
	"MDM_T02v2.pv1": StringSet{
		"OBX": true,
		"TXA": true,
	},
	"MDM_T02v2.txa": StringSet{
		"OBX": true,
	},
	"MFD_M01v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M01v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M01v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFD_M02v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M02v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M02v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFD_M03v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M03v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M03v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M01v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M01v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M01v2.msa": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M01v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFK_M02v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M02v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M02v2.msa": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M02v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFK_M03v2.err": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M03v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M03v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M03v2.msa": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
	},
	"MFK_M03v2.msh": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFN_M01_MFEv2.mfe": StringSet{
		"MFE": true,
	},
	"MFN_M01v2.mfi": StringSet{
		"MFE": true,
	},
	"MFN_M01v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
	},
	"MFN_M02_MFEv2.mfe": StringSet{
		"MFE": true,
		"PRA": true,
		"STF": true,
	},
	"MFN_M02_MFEv2.pra": StringSet{
		"MFE": true,
		"PRA": true,
		"STF": true,
	},
	"MFN_M02_MFEv2.stf": StringSet{
		"MFE": true,
		"PRA": true,
		"STF": true,
	},
	"MFN_M02v2.mfi": StringSet{
		"MFE": true,
		"PRA": true,
		"STF": true,
	},
	"MFN_M02v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
		"PRA": true,
		"STF": true,
	},
	"MFN_M03_MFEv2.mfe": StringSet{
		"MFE": true,
		"OM1": true,
	},
	"MFN_M03_MFEv2.om1": StringSet{
		"MFE": true,
		"OM1": true,
	},
	"MFN_M03v2.mfi": StringSet{
		"MFE": true,
		"OM1": true,
	},
	"MFN_M03v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
		"OM1": true,
	},
	"MFN_M05_MFE_LDPv2.lcc": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
	},
	"MFN_M05_MFE_LDPv2.lch": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
	},
	"MFN_M05_MFE_LDPv2.ldp": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
	},
	"MFN_M05_MFEv2.lch": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
	},
	"MFN_M05_MFEv2.loc": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
	},
	"MFN_M05_MFEv2.lrl": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
	},
	"MFN_M05_MFEv2.mfe": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
	},
	"MFN_M05v2.mfi": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
	},
	"MFN_M05v2.msh": StringSet{
		"LCC": true,
		"LCH": true,
		"LDP": true,
		"LOC": true,
		"LRL": true,
		"MFE": true,
		"MFI": true,
	},
	"MFN_M06_MFEv2.cdm": StringSet{
		"CDM": true,
		"MFE": true,
		"PRC": true,
	},
	"MFN_M06_MFEv2.mfe": StringSet{
		"CDM": true,
		"MFE": true,
		"PRC": true,
	},
	"MFN_M06_MFEv2.prc": StringSet{
		"CDM": true,
		"MFE": true,
		"PRC": true,
	},
	"MFN_M06v2.mfi": StringSet{
		"CDM": true,
		"MFE": true,
		"PRC": true,
	},
	"MFN_M06v2.msh": StringSet{
		"CDM": true,
		"MFE": true,
		"MFI": true,
		"PRC": true,
	},
	"MFN_M07_MFE_CM1v2.cm1": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
	},
	"MFN_M07_MFE_CM1v2.cm2": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
	},
	"MFN_M07_MFEv2.cm0": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
		"MFE": true,
	},
	"MFN_M07_MFEv2.mfe": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
		"MFE": true,
	},
	"MFN_M07v2.mfi": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
		"MFE": true,
	},
	"MFN_M07v2.msh": StringSet{
		"CM0": true,
		"CM1": true,
		"CM2": true,
		"MFE": true,
		"MFI": true,
	},
	"MFN_M08_MFEv2.mfe": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08_MFEv2.om1": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08_MFEv2.om2": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08_MFEv2.om3": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08_MFEv2.om4": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08v2.mfi": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M08v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
		"OM1": true,
		"OM2": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M09_MFEv2.mfe": StringSet{
		"MFE": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M09_MFEv2.om3": StringSet{
		"MFE": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M09_MFEv2.om4": StringSet{
		"MFE": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M09v2.mfi": StringSet{
		"MFE": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M09v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
		"OM3": true,
		"OM4": true,
	},
	"MFN_M10_OM5v2.om4": StringSet{
		"OM4": true,
		"OM5": true,
	},
	"MFN_M10_OM5v2.om5": StringSet{
		"OM4": true,
		"OM5": true,
	},
	"MFN_M10v2.mfi": StringSet{
		"OM4": true,
		"OM5": true,
	},
	"MFN_M10v2.msh": StringSet{
		"MFI": true,
		"OM4": true,
		"OM5": true,
	},
	"MFN_M11_MFEv2.mfe": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFN_M11_MFEv2.om1": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFN_M11_MFEv2.om2": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFN_M11_MFEv2.om6": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFN_M11v2.mfi": StringSet{
		"MFE": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFN_M11v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
		"OM1": true,
		"OM2": true,
		"OM6": true,
	},
	"MFQ_M01v2.dsc": StringSet{
	},
	"MFQ_M01v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M01v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M01v2.qrf": StringSet{
		"DSC": true,
	},
	"MFQ_M02v2.dsc": StringSet{
	},
	"MFQ_M02v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M02v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M02v2.qrf": StringSet{
		"DSC": true,
	},
	"MFQ_M03v2.dsc": StringSet{
	},
	"MFQ_M03v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M03v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M03v2.qrf": StringSet{
		"DSC": true,
	},
	"MFR_M01_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M01v2.dsc": StringSet{
	},
	"MFR_M01v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M01v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M01v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"MFR_M02_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M02v2.dsc": StringSet{
	},
	"MFR_M02v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M02v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M02v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"MFR_M03_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M03v2.dsc": StringSet{
	},
	"MFR_M03v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M03v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M03v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"NMD_N01_NCK_NSCv2.nsc": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSCv2.nte": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSTv2.nst": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSTv2.nte": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCKv2.nte": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01v2.msh": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMQ_N02_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02_NCKv2.nsc": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02_NCKv2.nst": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02v2.msh": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"QRD": true,
		"QRF": true,
	},
	"NMQ_N02v2.qrd": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"QRF": true,
	},
	"NMQ_N02v2.qrf": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMR_N02_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nsc": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nst": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte1": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte2": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte3": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02v2.err": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.msa": StringSet{
		"ERR": true,
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.qrd": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"OMD_O01_ORC1_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC1_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC1v2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC1v2.ods": StringSet{
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC1v2.orc": StringSet{
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC2v2.nte": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC2v2.odt": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_ORC2v2.orc": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
	},
	"OMD_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
	},
	"OMD_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"OMD_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV2": true,
	},
	"OMD_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"OMD_O01v2.msh": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"OMD_O01v2.nte": StringSet{
		"AL1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"OMN_O01_ORC_OBXv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORC_OBXv2.obx": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORCv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORCv2.rq1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_ORCv2.rqd": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01v2.msh": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMN_O01v2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
	},
	"OMS_O01_ORC_OBXv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RQD": true,
	},
	"OMS_O01_ORC_OBXv2.obx": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RQD": true,
	},
	"OMS_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_ORCv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_ORCv2.rqd": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RQD": true,
	},
	"OMS_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RQD": true,
	},
	"OMS_O01v2.msh": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQD": true,
	},
	"OMS_O01v2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQD": true,
	},
	"ORD_O02_PID_ORC1v2.nte": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PID_ORC1v2.ods": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PID_ORC1v2.orc": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PID_ORC2v2.nte": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PID_ORC2v2.odt": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PID_ORC2v2.orc": StringSet{
		"NTE": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
	},
	"ORD_O02v2.err": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
	},
	"ORD_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
	},
	"ORD_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
	},
	"ORD_O02v2.nte": StringSet{
		"NTE": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
	},
	"ORF_R04_PID_OBR_OBXv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PID_OBR_OBXv2.obx": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PID_OBRv2.cti": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PID_OBRv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PID_OBRv2.obr": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PID_OBRv2.orc": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_PIDv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
	},
	"ORF_R04_PIDv2.pid": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
	},
	"ORF_R04v2.dsc": StringSet{
	},
	"ORF_R04v2.msa": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04v2.msh": StringSet{
		"CTI": true,
		"DSC": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04v2.qrd": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
	},
	"ORF_R04v2.qrf": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
	},
	"ORM_O01_ORC_OBXv2.nte": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORC_OBXv2.obx": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.cti": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.dg1": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.nte": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.obr": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.ods": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.odt": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rq1": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rqd": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rxo": StringSet{
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01v2.msh": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01v2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"CTI": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORN_O02_PID_ORCv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02_PID_ORCv2.orc": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02_PID_ORCv2.rq1": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02_PID_ORCv2.rqd": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02v2.err": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORN_O02v2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
	},
	"ORR_O02_PID_ORCv2.cti": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.obr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.ods": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.odt": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.orc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rq1": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rqd": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rxo": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PIDv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PIDv2.pid": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.err": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.msa": StringSet{
		"CTI": true,
		"ERR": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.msh": StringSet{
		"CTI": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORU_R01_PID_OBR_OBXv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PID_OBR_OBXv2.obx": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PID_OBRv2.cti": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PID_OBRv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PID_OBRv2.obr": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PID_OBRv2.orc": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PIDv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PIDv2.pd1": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PIDv2.pid": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PIDv2.pv1": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01_PIDv2.pv2": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R01v2.dsc": StringSet{
	},
	"ORU_R01v2.msh": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R03_PID_OBR_OBXv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBR_OBXv2.obx": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.obr": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.pv1": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03v2.dsc": StringSet{
	},
	"ORU_R03v2.msh": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R32_PID_OBR_OBXv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBR_OBXv2.obx": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.cti": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.obr": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.orc": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.nk1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pd1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pid": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pv1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pv2": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32v2.dsc": StringSet{
	},
	"ORU_R32v2.msh": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"OSQ_Q06v2.dsc": StringSet{
	},
	"OSQ_Q06v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"OSQ_Q06v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"OSQ_Q06v2.qrf": StringSet{
		"DSC": true,
	},
	"OSR_Q06_PID_ORCv2.cti": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06_PID_ORCv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06_PID_ORCv2.obr": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06_PID_ORCv2.orc": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06_PIDv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06_PIDv2.pid": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
	},
	"OSR_Q06v2.dsc": StringSet{
	},
	"OSR_Q06v2.err": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"OSR_Q06v2.msa": StringSet{
		"CTI": true,
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"OSR_Q06v2.msh": StringSet{
		"CTI": true,
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"OSR_Q06v2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"OSR_Q06v2.qrd": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
	},
	"OSR_Q06v2.qrf": StringSet{
		"CTI": true,
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"PID": true,
	},
	"PEX_P07_PES_PEO_PCR_CSRv2.csp": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_CSRv2.csr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1_RXAv2.rxa": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1_RXAv2.rxr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1v2.nk1": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1v2.obx": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1v2.prb": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1v2.rxe": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_NK1v2.rxr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_RXAv2.rxa": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCR_RXAv2.rxr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.nte": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.obx": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.pcr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.prb": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.rxe": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEO_PCRv2.rxr": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PES_PEOv2.peo": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07_PESv2.pes": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.evn": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PD1": true,
		"PEO": true,
		"PES": true,
		"PID": true,
		"PRB": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.msh": StringSet{
		"CSP": true,
		"CSR": true,
		"EVN": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PD1": true,
		"PEO": true,
		"PES": true,
		"PID": true,
		"PRB": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.nte": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.pd1": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.pid": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PD1": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.pv1": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"PV2": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PEX_P07v2.pv2": StringSet{
		"CSP": true,
		"CSR": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"PCR": true,
		"PEO": true,
		"PES": true,
		"PRB": true,
		"RXA": true,
		"RXE": true,
		"RXR": true,
	},
	"PGL_PC6_GOL_OBXv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_OBXv2.obx": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORC_OBXv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORC_OBXv2.obx": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORC_OBXv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORCv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORCv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORCv2.orc": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ORCv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRB_OBXv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRB_OBXv2.obx": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRB_ROLv2.rol": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRB_ROLv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRBv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRBv2.prb": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PRBv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PTHv2.pth": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_PTHv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ROLv2.rol": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOL_ROLv2.var_": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6v2.msh": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6v2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6v2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PGL_PC6v2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PIN_I07_IN1v2.in1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"PIN_I07_IN1v2.in2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"PIN_I07_IN1v2.in3": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"PIN_I07_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"PIN_I07_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"PIN_I07v2.gt1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"PIN_I07v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"PIN_I07v2.nk1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"PIN_I07v2.nte": StringSet{
		"NTE": true,
	},
	"PIN_I07v2.pid": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"PPG_PCG_PTH_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTH_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTHv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCG_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCGv2.msh": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCGv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCGv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPG_PCGv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTH_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTHv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCB_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCBv2.msh": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCBv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCBv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPP_PCBv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1v2.msh": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1v2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1v2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPR_PC1v2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTH_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTHv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PID_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PIDv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PIDv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCL_PIDv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCLv2.err": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCLv2.msa": StringSet{
		"ERR": true,
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCLv2.msh": StringSet{
		"ERR": true,
		"GOL": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPT_PCLv2.qrd": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PID_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PIDv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PIDv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCA_PIDv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCAv2.err": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCAv2.msa": StringSet{
		"ERR": true,
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCAv2.msh": StringSet{
		"ERR": true,
		"GOL": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PPV_PCAv2.qrd": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PID_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PIDv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PIDv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5_PIDv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5v2.err": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5v2.msa": StringSet{
		"ERR": true,
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5v2.msh": StringSet{
		"ERR": true,
		"GOL": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PRR_PC5v2.qrd": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOL_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOL_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOL_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOL_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOLv2.gol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOLv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_GOLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORC_OBXv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORC_OBXv2.obx": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORC_OBXv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORCv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORCv2.obr": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORCv2.orc": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ORCv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRB_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRBv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRBv2.prb": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_PRBv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_ROLv2.rol": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTH_ROLv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTHv2.nte": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTHv2.pth": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PID_PTHv2.var_": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PIDv2.pid": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PIDv2.pv1": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCF_PIDv2.pv2": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCFv2.err": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCFv2.msa": StringSet{
		"ERR": true,
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCFv2.msh": StringSet{
		"ERR": true,
		"GOL": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"ROL": true,
		"VAR": true,
	},
	"PTR_PCFv2.qrd": StringSet{
		"GOL": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PRB": true,
		"PTH": true,
		"PV1": true,
		"PV2": true,
		"ROL": true,
		"VAR": true,
	},
	"QCK_Q02v2.err": StringSet{
		"QAK": true,
	},
	"QCK_Q02v2.msa": StringSet{
		"ERR": true,
		"QAK": true,
	},
	"QCK_Q02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"QAK": true,
	},
	"QCK_Q02v2.qak": StringSet{
	},
	"QRY_A19v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"QRY_A19v2.qrd": StringSet{
		"QRF": true,
	},
	"QRY_A19v2.qrf": StringSet{
	},
	"QRY_P04v2.dsc": StringSet{
	},
	"QRY_P04v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_P04v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_P04v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_PC4v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"QRY_PC4v2.qrd": StringSet{
		"QRF": true,
	},
	"QRY_PC4v2.qrf": StringSet{
	},
	"QRY_Q01v2.dsc": StringSet{
	},
	"QRY_Q01v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_Q01v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_Q01v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_Q02v2.dsc": StringSet{
	},
	"QRY_Q02v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_Q02v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_Q02v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_R02v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"QRY_R02v2.qrd": StringSet{
		"QRF": true,
	},
	"QRY_R02v2.qrf": StringSet{
	},
	"QRY_T12v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"QRY_T12v2.qrd": StringSet{
		"QRF": true,
	},
	"QRY_T12v2.qrf": StringSet{
	},
	"RAR_RAR_QRD_ORCv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRD_ORCv2.rxa": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRD_ORCv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRD_ORCv2.rxe": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRD_ORCv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RAR_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RARv2.dsc": StringSet{
	},
	"RAR_RARv2.err": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RARv2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAR_RARv2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RAS_O01_ORC_OBXv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_OBXv2.obx": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXEv2.rxc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXEv2.rxe": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXEv2.rxr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXO_RXCv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXO_RXCv2.rxc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXOv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXOv2.rxo": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORC_RXOv2.rxr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORCv2.cti": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORCv2.orc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORCv2.rxa": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_ORCv2.rxr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.pv1": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01_PIDv2.pv2": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01v2.msh": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RAS_O01v2.nte": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RCI_I05_OBR_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
	},
	"RCI_I05_OBR_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
	},
	"RCI_I05_OBRv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05_OBRv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05_PRDv2.ctd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
	},
	"RCI_I05_PRDv2.prd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
	},
	"RCI_I05v2.al1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05v2.dg1": StringSet{
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05v2.drg": StringSet{
		"AL1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05v2.msa": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RCI_I05v2.msh": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RCI_I05v2.nte": StringSet{
		"NTE": true,
	},
	"RCI_I05v2.pid": StringSet{
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"RCI_I05v2.qrd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
		"QRF": true,
	},
	"RCI_I05v2.qrf": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PRD": true,
	},
	"RCL_I06_PRDv2.ctd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RCL_I06_PRDv2.prd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RCL_I06v2.al1": StringSet{
		"AL1": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RCL_I06v2.dg1": StringSet{
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RCL_I06v2.drg": StringSet{
		"AL1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RCL_I06v2.dsc": StringSet{
	},
	"RCL_I06v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"RCL_I06v2.msa": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RCL_I06v2.msh": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"MSA": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RCL_I06v2.nte": StringSet{
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RCL_I06v2.pid": StringSet{
		"AL1": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RCL_I06v2.qrd": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRF": true,
	},
	"RCL_I06v2.qrf": StringSet{
		"AL1": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RDE_O01_ORC_OBXv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_OBXv2.obx": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_RXO_RXCv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_RXO_RXCv2.rxc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_RXOv2.nte": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_RXOv2.rxo": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORC_RXOv2.rxr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORCv2.cti": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORCv2.orc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORCv2.rxc": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORCv2.rxe": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_ORCv2.rxr": StringSet{
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"CTI": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01v2.msh": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDE_O01v2.nte": StringSet{
		"AL1": true,
		"CTI": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORC_OBXv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORC_OBXv2.obx": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORC_RXCv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORC_RXCv2.rxc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORCv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORCv2.rxo": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_ORCv2.rxr": StringSet{
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PID_IN1v2.in1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PID_IN1v2.in2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PID_IN1v2.in3": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.gt1": StringSet{
		"AL1": true,
		"BLG": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.pv1": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01_PIDv2.pv2": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01v2.msh": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDO_O01v2.nte": StringSet{
		"AL1": true,
		"BLG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORC_RXDv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORC_RXDv2.rxd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORC_RXDv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORCv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORCv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORCv2.rxe": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRD_ORCv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDR_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDRv2.dsc": StringSet{
	},
	"RDR_RDRv2.err": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDRv2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDR_RDRv2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXR": true,
	},
	"RDS_O01_ORC_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXEv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXEv2.rxe": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXEv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXO_RXCv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXO_RXCv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXOv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXOv2.rxo": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORC_RXOv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORCv2.orc": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORCv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORCv2.rxd": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_ORCv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.pd1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.pv1": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01_PIDv2.pv2": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01v2.msh": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"RDS_O01v2.nte": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXD": true,
		"RXE": true,
		"RXO": true,
		"RXR": true,
	},
	"REF_I12_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_OBR_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_OBR_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_OBRv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_OBRv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_PR1v2.aut": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_PR1v2.ctd": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_PR1v2.pr1": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_PRDv2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12_PRDv2.prd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.acc": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.al1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.aut": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.dg1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.drg": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"REF_I12v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.nte": StringSet{
		"NTE": true,
	},
	"REF_I12v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"REF_I12v2.pv1": StringSet{
		"NTE": true,
		"PV2": true,
	},
	"REF_I12v2.pv2": StringSet{
		"NTE": true,
	},
	"REF_I12v2.rf1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RER_RER_QRD_ORCv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRD_ORCv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRD_ORCv2.rxe": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRD_ORCv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RER_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RERv2.dsc": StringSet{
	},
	"RER_RERv2.err": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RERv2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RER_RERv2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORC_RXEv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORC_RXEv2.rxe": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORC_RXEv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORCv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORCv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORCv2.rxg": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRD_ORCv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGR_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGRv2.dsc": StringSet{
	},
	"RGR_RGRv2.err": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGRv2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGR_RGRv2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXCv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXCv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXEv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXEv2.rxe": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXEv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXG_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXG_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXGv2.rxc": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXGv2.rxg": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORC_RXGv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORCv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORCv2.orc": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORCv2.rxo": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_ORCv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_PIDv2.al1": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_PIDv2.nte": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_PIDv2.pid": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_PIDv2.pv1": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01_PIDv2.pv2": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01v2.msh": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"RGV_O01v2.nte": StringSet{
		"AL1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXC": true,
		"RXE": true,
		"RXG": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRD_ORCv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRD_ORCv2.rxc": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRD_ORCv2.rxo": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRD_ORCv2.rxr": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_ROR_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_RORv2.dsc": StringSet{
	},
	"ROR_RORv2.err": StringSet{
		"DSC": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_RORv2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"ROR_RORv2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RPA_I08_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_OBR_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_OBR_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_OBRv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_OBRv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_PR1v2.aut": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_PR1v2.ctd": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_PR1v2.pr1": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_PRDv2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08_PRDv2.prd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.acc": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.al1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.aut": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.dg1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.drg": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.msa": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"RPA_I08v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"RPA_I08v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.nte": StringSet{
		"NTE": true,
	},
	"RPA_I08v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RPA_I08v2.pv1": StringSet{
		"NTE": true,
		"PV2": true,
	},
	"RPA_I08v2.pv2": StringSet{
		"NTE": true,
	},
	"RPA_I08v2.rf1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RPI_I01_IN1v2.in1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RPI_I01_IN1v2.in2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RPI_I01_IN1v2.in3": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RPI_I01_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RPI_I01_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RPI_I01v2.gt1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RPI_I01v2.msa": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RPI_I01v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RPI_I01v2.nk1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"RPI_I01v2.nte": StringSet{
		"NTE": true,
	},
	"RPI_I01v2.pid": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"RPL_I02_PRDv2.ctd": StringSet{
		"CTD": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PRD": true,
	},
	"RPL_I02_PRDv2.prd": StringSet{
		"CTD": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PRD": true,
	},
	"RPL_I02v2.dsc": StringSet{
	},
	"RPL_I02v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"RPL_I02v2.msa": StringSet{
		"CTD": true,
		"DSC": true,
		"DSP": true,
		"NTE": true,
		"PRD": true,
	},
	"RPL_I02v2.msh": StringSet{
		"CTD": true,
		"DSC": true,
		"DSP": true,
		"MSA": true,
		"NTE": true,
		"PRD": true,
	},
	"RPL_I02v2.nte": StringSet{
		"DSC": true,
		"DSP": true,
		"NTE": true,
	},
	"RQA_I08_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_OBR_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_OBR_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_OBRv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_OBRv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_PR1v2.aut": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_PR1v2.ctd": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_PR1v2.pr1": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_PRDv2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08_PRDv2.prd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.acc": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.al1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.aut": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.dg1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.drg": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"RQA_I08v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.nte": StringSet{
		"NTE": true,
	},
	"RQA_I08v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RQA_I08v2.pv1": StringSet{
		"NTE": true,
		"PV2": true,
	},
	"RQA_I08v2.pv2": StringSet{
		"NTE": true,
	},
	"RQA_I08v2.rf1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RQC_I05_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQC_I05_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQC_I05v2.gt1": StringSet{
		"GT1": true,
		"NTE": true,
	},
	"RQC_I05v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RQC_I05v2.nk1": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQC_I05v2.nte": StringSet{
		"NTE": true,
	},
	"RQC_I05v2.pid": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQC_I05v2.qrd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRF": true,
	},
	"RQC_I05v2.qrf": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQC_I06_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQC_I06_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQC_I06v2.gt1": StringSet{
		"NTE": true,
	},
	"RQC_I06v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRD": true,
		"QRF": true,
	},
	"RQC_I06v2.nk1": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQC_I06v2.nte": StringSet{
		"NTE": true,
	},
	"RQC_I06v2.pid": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQC_I06v2.qrd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
		"QRF": true,
	},
	"RQC_I06v2.qrf": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQI_I01_IN1v2.in1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RQI_I01_IN1v2.in2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RQI_I01_IN1v2.in3": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RQI_I01_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQI_I01_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQI_I01v2.gt1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
	},
	"RQI_I01v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQI_I01v2.nk1": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"RQI_I01v2.nte": StringSet{
		"NTE": true,
	},
	"RQI_I01v2.pid": StringSet{
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
	},
	"RQP_I04_PRDv2.ctd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQP_I04_PRDv2.prd": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQP_I04v2.gt1": StringSet{
		"GT1": true,
		"NTE": true,
	},
	"RQP_I04v2.msh": StringSet{
		"CTD": true,
		"GT1": true,
		"NK1": true,
		"NTE": true,
		"PID": true,
		"PRD": true,
	},
	"RQP_I04v2.nk1": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQP_I04v2.nte": StringSet{
		"NTE": true,
	},
	"RQP_I04v2.pid": StringSet{
		"GT1": true,
		"NK1": true,
		"NTE": true,
	},
	"RQQ_Q01v2.dsc": StringSet{
	},
	"RQQ_Q01v2.erq": StringSet{
		"DSC": true,
	},
	"RQQ_Q01v2.msh": StringSet{
		"DSC": true,
		"ERQ": true,
	},
	"RRA_O02_PID_ORC_RXAv2.rxa": StringSet{
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02_PID_ORC_RXAv2.rxr": StringSet{
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02_PID_ORCv2.orc": StringSet{
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02v2.err": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXA": true,
		"RXR": true,
	},
	"RRA_O02v2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXA": true,
		"RXR": true,
	},
	"RRD_O02_PID_ORCv2.orc": StringSet{
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02_PID_ORCv2.rxc": StringSet{
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02_PID_ORCv2.rxd": StringSet{
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02_PID_ORCv2.rxr": StringSet{
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02v2.err": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRD_O02v2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXD": true,
		"RXR": true,
	},
	"RRG_O02_PID_ORCv2.orc": StringSet{
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02_PID_ORCv2.rxc": StringSet{
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02_PID_ORCv2.rxg": StringSet{
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02_PID_ORCv2.rxr": StringSet{
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02v2.err": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRG_O02v2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXG": true,
		"RXR": true,
	},
	"RRI_I12_OBR_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_OBR_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_OBRv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_OBRv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_PR1v2.aut": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_PR1v2.ctd": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_PR1v2.pr1": StringSet{
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_PRDv2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12_PRDv2.prd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.acc": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.al1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.aut": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.ctd": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.dg1": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.drg": StringSet{
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.msa": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"RRI_I12v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
		"RF1": true,
	},
	"RRI_I12v2.nte": StringSet{
		"NTE": true,
	},
	"RRI_I12v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
	},
	"RRI_I12v2.pv1": StringSet{
		"NTE": true,
		"PV2": true,
	},
	"RRI_I12v2.pv2": StringSet{
		"NTE": true,
	},
	"RRI_I12v2.rf1": StringSet{
		"ACC": true,
		"AL1": true,
		"AUT": true,
		"CTD": true,
		"DG1": true,
		"DRG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PRD": true,
		"PV1": true,
		"PV2": true,
	},
	"RRO_O02_PID_ORCv2.nte1": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PID_ORCv2.nte2": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PID_ORCv2.orc": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PID_ORCv2.rxc": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PID_ORCv2.rxo": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PID_ORCv2.rxr": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"ORC": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02v2.err": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02v2.msa": StringSet{
		"ERR": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"RRO_O02v2.nte": StringSet{
		"NTE": true,
		"ORC": true,
		"PID": true,
		"RXC": true,
		"RXO": true,
		"RXR": true,
	},
	"SIU_S12_PIDv2.dg1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12_PIDv2.obx": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12_PIDv2.pid": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12_PIDv2.pv1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12_PIDv2.pv2": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12_RGS_AIGv2.aig": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AIGv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AILv2.ail": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AILv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AIPv2.aip": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AIPv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AISv2.ais": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGS_AISv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SIU_S12_RGSv2.rgs": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
		"RGS": true,
	},
	"SIU_S12v2.msh": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SIU_S12v2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SIU_S12v2.sch": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SPQ_Q01v2.dsc": StringSet{
	},
	"SPQ_Q01v2.msh": StringSet{
		"DSC": true,
		"RDF": true,
		"SPR": true,
	},
	"SPQ_Q01v2.rdf": StringSet{
		"DSC": true,
	},
	"SPQ_Q01v2.spr": StringSet{
		"DSC": true,
		"RDF": true,
	},
	"SQM_S25_RGS_AIGv2.aig": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AIGv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AILv2.ail": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AILv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AIPv2.aip": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AIPv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AISv2.ais": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGS_AISv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
	},
	"SQM_S25_RGSv2.rgs": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
		"RGS": true,
	},
	"SQM_S25v2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
		"PID": true,
		"RGS": true,
	},
	"SQM_S25v2.arq": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
		"PID": true,
		"RGS": true,
	},
	"SQM_S25v2.dsc": StringSet{
	},
	"SQM_S25v2.msh": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"ARQ": true,
		"DSC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
		"RGS": true,
	},
	"SQM_S25v2.pid": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DSC": true,
		"RGS": true,
	},
	"SQM_S25v2.qrd": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"ARQ": true,
		"DSC": true,
		"PID": true,
		"QRF": true,
		"RGS": true,
	},
	"SQM_S25v2.qrf": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"ARQ": true,
		"DSC": true,
		"PID": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AIGv2.aig": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AIGv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AILv2.ail": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AILv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AIPv2.aip": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AIPv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AISv2.ais": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGS_AISv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCH_RGSv2.rgs": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SQR_S25_SCHv2.dg1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25_SCHv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25_SCHv2.pid": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25_SCHv2.pv1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25_SCHv2.pv2": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25_SCHv2.sch": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25v2.dsc": StringSet{
	},
	"SQR_S25v2.err": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QAK": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25v2.msa": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QAK": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25v2.msh": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QAK": true,
		"RGS": true,
		"SCH": true,
	},
	"SQR_S25v2.qak": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"DSC": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SRM_S01_PIDv2.dg1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01_PIDv2.obx": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01_PIDv2.pid": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01_PIDv2.pv1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01_PIDv2.pv2": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01_RGS_AIGv2.aig": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AIGv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AIGv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AILv2.ail": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AILv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AILv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AIPv2.aip": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AIPv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AIPv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AISv2.ais": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AISv2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGS_AISv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
	},
	"SRM_S01_RGSv2.rgs": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"NTE": true,
		"RGS": true,
	},
	"SRM_S01v2.apr": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01v2.arq": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01v2.msh": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"ARQ": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRM_S01v2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"APR": true,
		"DG1": true,
		"NTE": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01_PIDv2.dg1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01_PIDv2.pid": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01_PIDv2.pv1": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01_PIDv2.pv2": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01_RGS_AIGv2.aig": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AIGv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AILv2.ail": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AILv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AIPv2.aip": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AIPv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AISv2.ais": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGS_AISv2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
	},
	"SRR_S01_RGSv2.rgs": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"NTE": true,
		"RGS": true,
	},
	"SRR_S01v2.err": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SRR_S01v2.msa": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"ERR": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SRR_S01v2.msh": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"ERR": true,
		"MSA": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
		"SCH": true,
	},
	"SRR_S01v2.nte": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SRR_S01v2.sch": StringSet{
		"AIG": true,
		"AIL": true,
		"AIP": true,
		"AIS": true,
		"DG1": true,
		"NTE": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RGS": true,
	},
	"SUR_P09_FAC_FAC2v2.fac": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FAC_FAC2v2.nte": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FAC_FAC2v2.pdc": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FAC_PSH1v2.pdc": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FAC_PSH1v2.psh": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FACv2.fac1": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09_FACv2.psh2": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"SUR_P09v2.msh": StringSet{
		"FAC": true,
		"NTE": true,
		"PDC": true,
		"PSH": true,
	},
	"TBR_Q01v2.dsc": StringSet{
	},
	"TBR_Q01v2.err": StringSet{
		"DSC": true,
		"QAK": true,
		"RDF": true,
		"RDT": true,
	},
	"TBR_Q01v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"QAK": true,
		"RDF": true,
		"RDT": true,
	},
	"TBR_Q01v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MSA": true,
		"QAK": true,
		"RDF": true,
		"RDT": true,
	},
	"TBR_Q01v2.qak": StringSet{
		"DSC": true,
		"RDF": true,
		"RDT": true,
	},
	"TBR_Q01v2.rdf": StringSet{
		"DSC": true,
		"RDT": true,
	},
	"TBR_Q01v2.rdt": StringSet{
		"DSC": true,
		"RDT": true,
	},
	"UDM_Q05v2.dsc": StringSet{
	},
	"UDM_Q05v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"UDM_Q05v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"URD": true,
		"URS": true,
	},
	"UDM_Q05v2.urd": StringSet{
		"DSC": true,
		"DSP": true,
		"URS": true,
	},
	"UDM_Q05v2.urs": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"VQQ_Q01v2.dsc": StringSet{
	},
	"VQQ_Q01v2.msh": StringSet{
		"DSC": true,
		"RDF": true,
		"VTQ": true,
	},
	"VQQ_Q01v2.rdf": StringSet{
		"DSC": true,
	},
	"VQQ_Q01v2.vtq": StringSet{
		"DSC": true,
		"RDF": true,
	},
	"VXQ_V01v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"VXQ_V01v2.qrd": StringSet{
		"QRF": true,
	},
	"VXQ_V01v2.qrf": StringSet{
	},
	"VXR_V03_IN1v2.in1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_IN1v2.in2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_IN1v2.in3": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_RXA_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_RXA_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_RXAv2.orc": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_RXAv2.rxa": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03_RXAv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.msa": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.msh": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"QRF": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.nk1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.pd1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.pid": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.pv1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.pv2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.qrd": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"QRF": true,
		"RXA": true,
		"RXR": true,
	},
	"VXR_V03v2.qrf": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_IN1v2.in1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_IN1v2.in2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_IN1v2.in3": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_RXA_OBXv2.nte": StringSet{
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_RXA_OBXv2.obx": StringSet{
		"NTE": true,
		"OBX": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_RXAv2.orc": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_RXAv2.rxa": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04_RXAv2.rxr": StringSet{
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.msh": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.nk1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.pd1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.pid": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.pv1": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"PV2": true,
		"RXA": true,
		"RXR": true,
	},
	"VXU_V04v2.pv2": StringSet{
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NTE": true,
		"OBX": true,
		"ORC": true,
		"RXA": true,
		"RXR": true,
	},
	"VXX_V02_PIDv2.nk1": StringSet{
		"NK1": true,
		"PID": true,
	},
	"VXX_V02_PIDv2.pid": StringSet{
		"NK1": true,
		"PID": true,
	},
	"VXX_V02v2.msa": StringSet{
		"NK1": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"VXX_V02v2.msh": StringSet{
		"MSA": true,
		"NK1": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"VXX_V02v2.qrd": StringSet{
		"NK1": true,
		"PID": true,
		"QRF": true,
	},
	"VXX_V02v2.qrf": StringSet{
		"NK1": true,
		"PID": true,
	},
}
