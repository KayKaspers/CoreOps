package observe

// This file realizes the per-field provenance structure of Observation
// Contract section 13 and its minimal validity semantics.
//
// Provenance must allow later reconstruction of source, collection mechanism
// class, observation time, normalization/transformation, freshness and
// evidence origin. It is a specialization of the Field Provenance standard
// section 6 and introduces no new schema.
//
// Provenance is NOT trust, and this package derives none of the following from
// it:
//
//	provenance present         != trusted
//	provenance present         != validated
//	provenance complete        != value correct
//	parsed                     != validated provenance
//	collection mechanism class != selected technology
//
// No trust level, no authorization, no security level, no compliance
// statement, no support status, no health and no validation beyond the
// contract semantics is inferred here.

// SourceIdentity is the identity of the concrete source of one field
// (section 9). It is itself part of the provenance. It is not trust, not
// authority and not validation.
type SourceIdentity string

// SourceClass is the source class per Field Provenance standard section 7
// (section 9). It is not a trust assignment.
type SourceClass string

// SourceReference references the concrete source location or artifact
// (section 13).
type SourceReference string

// CollectionMechanismClass is the class of the collection path — deliberately
// a CLASS, not a tool (sections 9 and 13).
//
//	collection mechanism class != selected technology
//	collection mechanism class != P-3 decision
//
// The tool choice is not made. This package selects none, names none and
// implements none.
type CollectionMechanismClass string

// RawValueReference references the retained raw value (section 13).
type RawValueReference string

// TransformationReference references the transformation applied to produce a
// normalized value (section 13). It is conditional: it exists only where a
// normalization actually ran. This package runs none.
type TransformationReference string

// AuditReference is the audit reference of the field record (section 13). A
// reference is not a governed audit record: that classification belongs to the
// Audit/Evidence model (MOD-EVD-001) and does not follow automatically.
type AuditReference string

// ValidationStatus is the field validation status carried in the provenance
// set (section 13). Its value vocabulary is owned by the Field Provenance and
// Data Lineage standard, NOT by the Observation Contract; the contract neither
// adopts nor bounds it for this slice. This package therefore carries the slot
// and deliberately enumerates NO values: defining them here would create a
// vocabulary this work package has no authority to create.
type ValidationStatus string

// ConflictState is the conflict state carried in the provenance set
// (section 13). Its vocabulary is owned by the source-of-truth conflict model,
// not by the Observation Contract, and is deliberately not enumerated here for
// the same reason as ValidationStatus. A conflict stays visible and is never
// silently resolved (sections 8.3 and 15).
type ConflictState string

// RawReference couples the raw value reference with whether it actually
// resolved.
//
// This package resolves nothing: resolution is a collection-side fact
// (internal/observe/collect, NOT AUTHORIZED, P-3 unrealized) and is carried
// here, never performed here. An unresolvable raw reference is one of the
// named provenance defects of section 10.3.
//
// The zero value is the absence of a raw reference, which is fail-closed.
type RawReference struct {
	present  bool
	resolved bool
	ref      RawValueReference
}

// NoRawReference returns the raw reference of a field for which none exists.
func NoRawReference() RawReference {
	return RawReference{}
}

// ResolvedRawReference returns a raw reference the collection side reported as
// resolvable.
func ResolvedRawReference(r RawValueReference) RawReference {
	return RawReference{present: true, resolved: true, ref: r}
}

// UnresolvedRawReference returns a raw reference the collection side reported
// as not resolvable. It is a provenance defect (section 10.3).
func UnresolvedRawReference(r RawValueReference) RawReference {
	return RawReference{present: true, resolved: false, ref: r}
}

// Reference reports the raw value reference and whether one exists at all.
func (r RawReference) Reference() (RawValueReference, bool) {
	return r.ref, r.present
}

// IsResolved reports whether a raw reference exists and resolved.
func (r RawReference) IsResolved() bool {
	return r.present && r.resolved
}

// FieldProvenance is the conceptual minimum provenance set for one observed
// field (section 13), specializing Field Provenance standard section 6. It
// introduces no new schema.
//
// The set of section 13 is carried jointly by this type and the
// FieldObservation that holds it: subject reference, canonical field identity,
// source_identity, source_class, source_reference, collection_mechanism_class,
// observed_at, received_at (conditional), freshness, raw value reference,
// transformation reference (conditional), field_observation_outcome,
// validation status, conflict state and audit reference. Field identity, field
// outcome and freshness live on FieldObservation; nothing is duplicated.
//
// The zero value of each required element is not a value: it is the absence of
// that element, and absence makes the provenance invalid (fail-closed).
type FieldProvenance struct {
	// Subject is the subject reference of the field record (section 13).
	Subject SubjectRef

	// SourceIdentity is required for each observed field. Without it the field
	// is provenance-invalid and unusable (section 9).
	SourceIdentity SourceIdentity

	// SourceClass is required for each observed field, with the same absence
	// semantics as SourceIdentity (section 9).
	SourceClass SourceClass

	// SourceReference references the concrete source (section 13).
	SourceReference SourceReference

	// CollectionMechanismClass is required for each observed field, with the
	// same absence semantics as SourceIdentity (section 9).
	CollectionMechanismClass CollectionMechanismClass

	// ObservedAt is the observation time this field is attributed to
	// (section 11). It is required and unknown where undeterminable.
	ObservedAt ObservedAt

	// ReceivedAt is conditional and not-applicable where collection and intake
	// are not materially separate (section 11).
	ReceivedAt ReceivedAt

	// Raw references the retained raw value and whether it resolved
	// (sections 12 and 13).
	Raw RawReference

	// Transformation is conditional: it exists only where a normalization
	// actually ran (sections 12 and 13).
	Transformation TransformationReference

	// Validation carries the field validation status (section 13). Its
	// vocabulary is owned outside the Observation Contract.
	Validation ValidationStatus

	// Conflict carries the conflict state (section 13). Conflicts stay visible
	// and are never silently resolved.
	Conflict ConflictState

	// Audit is the audit reference of the field record (section 13).
	Audit AuditReference
}

// RequiredElementsPresent reports whether the provenance elements the contract
// requires for a usable field are present.
//
// The named defects of section 10.3 are a missing source_identity, a missing
// collection_mechanism_class and an unresolvable raw reference. Section 9
// additionally makes source_class required per observed field with the same
// absence semantics as source_identity, so it is checked here on the same
// basis.
//
// A false result means the field is provenance-invalid: its
// field_observation_outcome is provenance-invalid, its emission disposition is
// field-withheld, and it does not enter a normalized representation
// (fail-closed, section 13). Those consequences are applied by composition.go;
// this predicate only reports the defect and changes nothing.
//
// The predicate is read-only and infers nothing beyond usability. A true
// result is not trust, not validation, not authorization, not compliance, not
// support and not correctness of the value.
func (p FieldProvenance) RequiredElementsPresent() bool {
	if p.SourceIdentity == "" {
		return false
	}
	if p.SourceClass == "" {
		return false
	}
	if p.CollectionMechanismClass == "" {
		return false
	}
	return p.Raw.IsResolved()
}
