// Package observe realizes the canonical observation domain of the first
// CoreOps Observe value slice — Local Linux Host Identity & Basic System
// Observation.
//
// Authority boundary. This package is an implementation. It realizes
// authority; it does not create authority. The Observation Contract
// (docs/architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) is the
// sole normative source for every semantic modelled here:
//
//	internal/observe        != Observation Contract
//	internal/observe        != MOD-OBS-001
//	Go type                 != canonical governance vocabulary
//	implementation constant != authority to extend a vocabulary
//	repository package      != logical module authority
//	source code             != source of truth for the contract
//
// Where this package and the contract disagree, this package is wrong.
//
// Scope. The package holds the canonical observation envelope (section 7), the
// canonical field identities (section 9), the three bounded vocabularies
// (section 10), the per-field provenance structure (section 13), the bounded
// freshness vocabulary (section 14) and the deterministic composition rules R1
// to R6 (section 10.5).
//
// The package is I/O-free by authorized design. It reads no file, no /proc
// entry, no /etc entry, no environment variable and no runtime host fact; it
// runs no process and opens no network connection. Collection is the
// responsibility of internal/observe/collect, which is NOT AUTHORIZED and does
// not exist: P-3 is decided as a mechanism class and remains unrealized. The
// dependency direction internal/observe -> internal/observe/collect is
// forbidden (Productive-Code Transition Prerequisites section 9.2).
//
// The package never calls time.Now. No timestamp is fabricated (section 11).
package observe

import "time"

// ObservationID is the stable, non-reusable identifier of one collection run
// (section 7). It is not the identity of the observed subject.
type ObservationID string

// SubjectRef references the observation subject (section 8.1): exactly one
// local Linux host. It is not a managed-resource identity.
type SubjectRef string

// CollectionAttemptRef identifies the concrete collection attempt. One subject
// may have several attempts (section 7).
type CollectionAttemptRef string

// ContractRef references the Observation Contract and its revision (section 7).
type ContractRef string

// ProvenanceRef references the provenance root of the run (sections 7 and 13).
// A provenance reference is not trust:
//
//	provenance present  != trusted
//	provenance present  != validated
//	provenance complete != value correct
type ProvenanceRef string

// EvidenceRef is an Evidence Reference in the sense of the Evidence model
// section 9.
//
//	reference present != artifact present
//	reference present != access authorized
//	reference present != evidence sufficient
type EvidenceRef string

// ObservedAt carries the instant an observation is attributed to (section 11).
// It is required. Where it cannot be determined it is unknown — it is never
// replaced by received_at, never by "now" and never by a default. The zero
// value is unknown, which is the fail-closed state; it is not a substitute
// value.
type ObservedAt struct {
	determined bool
	instant    time.Time
}

// UnknownObservedAt returns the required-but-undeterminable observed_at.
// An undeterminable instant stays undeterminable (section 11).
func UnknownObservedAt() ObservedAt {
	return ObservedAt{}
}

// ObservedAtInstant returns a determined observed_at. The caller supplies the
// instant; this package neither reads a clock nor derives one.
func ObservedAtInstant(t time.Time) ObservedAt {
	return ObservedAt{determined: true, instant: t}
}

// Instant reports the observed_at instant and whether it was determined. The
// instant is meaningless when the second result is false.
func (o ObservedAt) Instant() (time.Time, bool) {
	return o.instant, o.determined
}

// IsUnknown reports whether observed_at is unknown. Unknown is not
// not-applicable: received_at, not observed_at, carries not-applicable
// (section 11).
func (o ObservedAt) IsUnknown() bool {
	return !o.determined
}

// ReceivedAt carries the intake instant (section 11). It is conditional: it is
// kept only where collection and intake are materially separate. Where they
// are not, the field is not-applicable — never unknown and never a copy of
// observed_at. The zero value is not-applicable.
type ReceivedAt struct {
	separate bool
	instant  time.Time
}

// NotApplicableReceivedAt returns the received_at of a run in which collection
// and intake are not materially separate (section 11).
func NotApplicableReceivedAt() ReceivedAt {
	return ReceivedAt{}
}

// ReceivedAtInstant returns a received_at for a run in which collection and
// intake are materially separate. The caller supplies the instant.
func ReceivedAtInstant(t time.Time) ReceivedAt {
	return ReceivedAt{separate: true, instant: t}
}

// Instant reports the received_at instant and whether intake was materially
// separate from collection.
func (r ReceivedAt) Instant() (time.Time, bool) {
	return r.instant, r.separate
}

// IsNotApplicable reports whether received_at is not-applicable. This is
// distinct from unknown, which received_at never carries (section 11).
func (r ReceivedAt) IsNotApplicable() bool {
	return !r.separate
}

// EnvelopeOutcomeSlot carries the envelope result of a collection run.
//
// R1, R5 and R6 assert no envelope outcome at all (section 10.5). The slot
// models that state through the absence of an assertion rather than through a
// ninth vocabulary value: the eight-value vocabulary of section 10.2 stays
// bounded and provenance-invalid remains outside it.
//
//	emitted canonical Observation   => observation_outcome present
//	no observation_outcome asserted => record not emitted as a canonical Observation
//
// The zero value asserts nothing, which is the fail-closed state.
type EnvelopeOutcomeSlot struct {
	asserted bool
	outcome  ObservationOutcome
}

// NoOutcomeAsserted returns the slot of a run for which the contract asserts
// no envelope outcome (section 10.5, rules R1, R5 and R6).
func NoOutcomeAsserted() EnvelopeOutcomeSlot {
	return EnvelopeOutcomeSlot{}
}

// AssertOutcome returns a slot asserting o.
func AssertOutcome(o ObservationOutcome) EnvelopeOutcomeSlot {
	return EnvelopeOutcomeSlot{asserted: true, outcome: o}
}

// Outcome reports the asserted envelope outcome and whether one was asserted.
// The outcome is meaningless when the second result is false.
func (s EnvelopeOutcomeSlot) Outcome() (ObservationOutcome, bool) {
	return s.outcome, s.asserted
}

// IsAsserted reports whether an envelope outcome was asserted at all.
func (s EnvelopeOutcomeSlot) IsAsserted() bool {
	return s.asserted
}

// EvidenceReference carries the conditional Evidence Reference of a run
// (section 9). It is kept only where a reference was actually produced; no
// fictitious reference is created. The zero value carries none.
type EvidenceReference struct {
	present bool
	ref     EvidenceRef
}

// NoEvidenceReference returns the reference of a run that produced none.
func NoEvidenceReference() EvidenceReference {
	return EvidenceReference{}
}

// RecordedEvidenceReference returns a reference that was actually produced.
func RecordedEvidenceReference(r EvidenceRef) EvidenceReference {
	return EvidenceReference{present: true, ref: r}
}

// Reference reports the Evidence Reference and whether one exists.
func (e EvidenceReference) Reference() (EvidenceRef, bool) {
	return e.ref, e.present
}

// Observation is the canonical observation envelope of section 7: a single,
// time-bound, read-only collection run against exactly one subject, with
// exactly one source per field and exactly one result envelope.
//
//	observation_id      != subject identity
//	observation attempt != observation result
//	observation result  != authoritative CoreOps state
//	one observation     != a series
//
// The envelope asserts nothing about the target: an observation_outcome is a
// statement about the collection run, never a target status, never health,
// never support and never authorization.
//
// Not this type's responsibility: collecting any value, deriving target_id,
// normalizing anything, computing freshness, resolving evidence, persisting or
// serializing the record, or deciding emission — composition of the envelope
// from the per-field outcomes lives in composition.go.
type Observation struct {
	// ID is observation_id (section 7).
	ID ObservationID

	// Subject is observation_subject_ref (sections 7 and 8.1).
	Subject SubjectRef

	// CollectionAttempt is collection_attempt_ref (section 7).
	CollectionAttempt CollectionAttemptRef

	// Contract is contract_ref (section 7).
	Contract ContractRef

	// Target carries target_id and its provenance (sections 8.2 and 9). It is
	// the subject attribution rule R1 evaluates.
	Target TargetAttribution

	// Outcome is the envelope observation_outcome slot (section 10.2). It
	// asserts nothing under R1, R5 and R6.
	Outcome EnvelopeOutcomeSlot

	// ObservedAt is observed_at (section 11), required, unknown where
	// undeterminable.
	ObservedAt ObservedAt

	// ReceivedAt is received_at (section 11), conditional, not-applicable
	// where collection and intake are not materially separate.
	ReceivedAt ReceivedAt

	// Provenance is provenance_ref, the provenance root of the run
	// (sections 7 and 13).
	Provenance ProvenanceRef

	// Evidence is the conditional evidence_ref (section 9).
	Evidence EvidenceReference

	// Fields carries the per-field observations (section 9). All six canonical
	// fields are required; this type enforces no cardinality of its own.
	Fields []FieldObservation
}
