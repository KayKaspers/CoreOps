package observe

// This file realizes the three separate bounded vocabularies of Observation
// Contract section 10. No value moves between them; none is a superset of
// another.
//
//	envelope outcome     != per-field outcome
//	per-field outcome    != emission disposition
//	emission disposition != observation outcome
//	record discarded     != observation outcome
//	partial              != failed
//	partial              != complete
//
// No severity order, no precedence table, no winner semantics and no
// aggregation across heterogeneous failure kinds is defined here or anywhere
// else in this package. The predicates below are read-only membership and
// validity checks; they rank nothing.

// ObservationOutcome is the envelope vocabulary of section 10.2 — the result
// of the collection run as a whole. Exactly eight values; no others.
//
// It states something about the collection run, never about the target:
//
//	none of these != target unhealthy
//	none of these != failed
//	observation error != target unhealthy
//	permission denied != target unavailable
//	unsupported       != failed
//
// Explicitly rejected and not members: unknown, failed, error, timeout,
// healthy, unhealthy, degraded. provenance-invalid is explicitly NOT a ninth
// envelope value — it is a field value only (section 10.3).
type ObservationOutcome string

const (
	// OutcomeSuccess — source reachable, read, value obtained and, where
	// required, normalized. Not: value correct, target healthy, value trusted.
	OutcomeSuccess ObservationOutcome = "success"

	// OutcomePartial — at least one field success, at least one field not.
	// Not: failure, completeness, or "essentially successful".
	OutcomePartial ObservationOutcome = "partial"

	// OutcomeSourceAbsent — the expected source does not exist on this host.
	// Not: target absent, unsupported, or an error.
	OutcomeSourceAbsent ObservationOutcome = "source-absent"

	// OutcomeSourceUnavailable — source exists but was not readable in this
	// attempt. Not: permission denied, absent, or target offline. A timeout is
	// this value with provenance detail, not a value of its own.
	OutcomeSourceUnavailable ObservationOutcome = "source-unavailable"

	// OutcomePermissionDenied — read access was refused by the target
	// environment. Not: unavailable, absent, target unreachable, or an attack.
	OutcomePermissionDenied ObservationOutcome = "permission-denied"

	// OutcomeUnsupported — source or platform is outside the contractually
	// supported scope. Not: failed, absent, or a defect.
	OutcomeUnsupported ObservationOutcome = "unsupported"

	// OutcomeSourceMalformed — source read, but its content is not
	// interpretable per contract. Not: normalization-failed, not target
	// faulty.
	OutcomeSourceMalformed ObservationOutcome = "source-malformed"

	// OutcomeNormalizationFailed — a raw value exists, the mapping to the
	// canonical representation failed. Not: source-malformed, not value
	// invalid. The raw value is retained.
	OutcomeNormalizationFailed ObservationOutcome = "normalization-failed"
)

// observationOutcomes is the closed eight-value envelope vocabulary
// (section 10.2), in contract order.
var observationOutcomes = [...]ObservationOutcome{
	OutcomeSuccess,
	OutcomePartial,
	OutcomeSourceAbsent,
	OutcomeSourceUnavailable,
	OutcomePermissionDenied,
	OutcomeUnsupported,
	OutcomeSourceMalformed,
	OutcomeNormalizationFailed,
}

// ObservationOutcomes returns the eight envelope outcome values in contract
// order. The returned slice is a copy: the vocabulary is not extensible from
// outside this package, and this package has no authority to extend it.
func ObservationOutcomes() []ObservationOutcome {
	out := make([]ObservationOutcome, len(observationOutcomes))
	copy(out, observationOutcomes[:])
	return out
}

// IsCanonical reports whether o is one of the eight envelope outcome values.
// Membership only; it ranks nothing and asserts nothing about the target.
func (o ObservationOutcome) IsCanonical() bool {
	for _, c := range observationOutcomes {
		if c == o {
			return true
		}
	}
	return false
}

// FieldOutcome is the field vocabulary of section 10.3 — the result for
// exactly one observed field. Exactly nine values: the eight collection
// outcomes of section 10.2 plus provenance-invalid.
type FieldOutcome string

const (
	// Seven of the eight collection outcomes, meaning as in section 10.2 but
	// scoped to exactly one field. The eighth, partial, follows below with its
	// own note.
	FieldOutcomeSuccess             FieldOutcome = "success"
	FieldOutcomeSourceAbsent        FieldOutcome = "source-absent"
	FieldOutcomeSourceUnavailable   FieldOutcome = "source-unavailable"
	FieldOutcomePermissionDenied    FieldOutcome = "permission-denied"
	FieldOutcomeUnsupported         FieldOutcome = "unsupported"
	FieldOutcomeSourceMalformed     FieldOutcome = "source-malformed"
	FieldOutcomeNormalizationFailed FieldOutcome = "normalization-failed"

	// FieldOutcomePartial is admissible at field level only where a field is
	// itself composed from several source parts. In this slice that case does
	// not occur and the value stays unused; no composite-field mechanism is
	// implemented. The value is defined because the vocabulary has exactly
	// nine members (section 10.3).
	FieldOutcomePartial FieldOutcome = "partial"

	// FieldOutcomeProvenanceInvalid — a value was collected, but its
	// provenance is incomplete, contradictory or unresolvable (missing
	// source_identity, missing collection_mechanism_class, unresolvable raw
	// reference). The field is therefore unusable and is not normalized
	// (fail-closed, section 13).
	//
	// It is a statement about the usability of the field record, not about the
	// target and not about the source:
	//
	//	provenance invalid != target unhealthy
	//	provenance invalid != source unavailable
	//	provenance invalid != source absent
	//	provenance invalid != permission denied
	//	provenance invalid != value wrong
	//	provenance invalid != source malicious
	//	provenance invalid != an observation_outcome value
	//
	// This value exists at field level ONLY. It is never an envelope outcome.
	FieldOutcomeProvenanceInvalid FieldOutcome = "provenance-invalid"
)

// fieldOutcomes is the closed nine-value field vocabulary (section 10.3), in
// contract order.
var fieldOutcomes = [...]FieldOutcome{
	FieldOutcomeSuccess,
	FieldOutcomePartial,
	FieldOutcomeSourceAbsent,
	FieldOutcomeSourceUnavailable,
	FieldOutcomePermissionDenied,
	FieldOutcomeUnsupported,
	FieldOutcomeSourceMalformed,
	FieldOutcomeNormalizationFailed,
	FieldOutcomeProvenanceInvalid,
}

// FieldOutcomes returns the nine field outcome values in contract order. The
// returned slice is a copy.
func FieldOutcomes() []FieldOutcome {
	out := make([]FieldOutcome, len(fieldOutcomes))
	copy(out, fieldOutcomes[:])
	return out
}

// IsCanonical reports whether f is one of the nine field outcome values.
// Membership only.
func (f FieldOutcome) IsCanonical() bool {
	for _, c := range fieldOutcomes {
		if c == f {
			return true
		}
	}
	return false
}

// CollectionOutcome reports the envelope-vocabulary value that corresponds to
// this field outcome, and whether such a correspondence exists.
//
// The eight collection outcomes carry the same canonical name at both levels
// (section 10.3). provenance-invalid has NO envelope counterpart: it is a
// field value only, it never yields an envelope value, and the second result
// is false for it (sections 10.2, 10.3 and 10.5).
//
// This is a validity mapping over identical canonical names, not a
// translation, not a promotion and not an ordering. A non-member value has no
// correspondence either.
func (f FieldOutcome) CollectionOutcome() (ObservationOutcome, bool) {
	if f == FieldOutcomeProvenanceInvalid {
		return "", false
	}
	o := ObservationOutcome(f)
	if !o.IsCanonical() {
		return "", false
	}
	return o, true
}

// EmissionDisposition is the processing vocabulary of section 10.4 — whether
// and how a record is emitted at all. Exactly three values.
//
// It describes processing and output, never the observation result:
//
//	record discarded     != observation outcome
//	record discarded     != failed observation
//	record discarded     != target unhealthy
//	record discarded     != target state
//	record discarded     != deletion
//	field withheld       != field absent
//	emission disposition != observation outcome
type EmissionDisposition string

const (
	// DispositionEmitted — the record is emitted, with all fields and their
	// field_observation_outcome values.
	DispositionEmitted EmissionDisposition = "emitted"

	// DispositionFieldWithheld — a single field does not enter the normalized
	// representation; its raw value, its provenance and its
	// field_observation_outcome stay visible in the record.
	DispositionFieldWithheld EmissionDisposition = "field-withheld"

	// DispositionRecordDiscarded — the record is not emitted as a canonical
	// Observation because a mandatory validity condition cannot be met
	// truthfully. No observation_outcome is asserted. This is not deletion:
	// the material already produced is retained under the existing semantics
	// (section 10.4).
	//
	// The trigger list is closed: R1, R5 and R6 (section 10.5) and nothing
	// else. Extending it requires its own explicit Human-Maintainer
	// authorization; this package does not extend it.
	DispositionRecordDiscarded EmissionDisposition = "record-discarded"
)

// emissionDispositions is the closed three-value processing vocabulary
// (section 10.4), in contract order.
var emissionDispositions = [...]EmissionDisposition{
	DispositionEmitted,
	DispositionFieldWithheld,
	DispositionRecordDiscarded,
}

// EmissionDispositions returns the three emission dispositions in contract
// order. The returned slice is a copy.
func EmissionDispositions() []EmissionDisposition {
	out := make([]EmissionDisposition, len(emissionDispositions))
	copy(out, emissionDispositions[:])
	return out
}

// IsCanonical reports whether d is one of the three emission dispositions.
// Membership only.
func (d EmissionDisposition) IsCanonical() bool {
	for _, c := range emissionDispositions {
		if c == d {
			return true
		}
	}
	return false
}
