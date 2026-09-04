package observe

// This file realizes the deterministic envelope composition of Observation
// Contract section 10.5, rules R1 to R6.
//
// The contract's evaluation ORDER is binding and is followed literally below.
// Nothing else is introduced:
//
//   - no "most severe failure wins" logic
//   - no priority table
//   - no severity hierarchy
//   - no winner semantics
//   - no synthetic aggregation of heterogeneous failure kinds
//   - no ninth envelope value
//
// partial is not weakened: it requires the success share the contract demands.
//
// R1, R5 and R6 assert NO envelope outcome. That state is modelled by an
// unasserted EnvelopeOutcomeSlot, not by a new vocabulary value.
//
//	R1 = invalid subject attribution       (target_id)
//	R5 = nothing provenance-bearing/usable (all fields provenance-invalid)
//	R6 = no truthful envelope composition  (heterogeneous all-failure)
//
// R4 and R6 are disjoint and gapless: if every non-provenance-invalid field
// carries the SAME collection outcome, R4 applies; if they carry DIFFERENT
// ones, R6 applies. provenance-invalid fields never yield an envelope value;
// they can trigger R3, or R5 when they affect every field, and they do NOT
// count towards R6.

// Rule identifies the composition rule of section 10.5 that applied.
//
// These identifiers name existing normative contract rules; they are not an
// outcome vocabulary, carry no ordering and carry no severity. They exist so
// the applied rule stays auditable.
type Rule uint8

const (
	// RuleNone means no contract rule applied. See Compose.
	RuleNone Rule = iota
	RuleR1
	RuleR2
	RuleR3
	RuleR4
	RuleR5
	RuleR6
)

// Composition is the deterministic result of composing an envelope from the
// per-field outcomes (section 10.5).
type Composition struct {
	// Rule is the contract rule that applied.
	Rule Rule

	// Outcome is the envelope observation_outcome slot. It asserts nothing
	// under R1, R5 and R6.
	Outcome EnvelopeOutcomeSlot

	// Disposition is the record-level emission disposition (section 10.4).
	Disposition EmissionDisposition

	// WithheldFields lists the fields whose emission disposition is
	// field-withheld — every provenance-invalid field of an emitted record
	// (sections 10.4, 10.5 and 13). It is nil for a discarded record, which is
	// not emitted at all; the per-field causes, provenance and raw values of a
	// discarded record are retained by the caller under section 10.4 and are
	// not this result's concern.
	WithheldFields []FieldIdentity
}

// Compose applies rules R1 to R6 of section 10.5 in the contract's binding
// order and reports the resulting envelope composition.
//
// The second result reports whether a contract rule applied at all. It is
// false only for inputs the contract does not cover:
//
//   - an observation carrying no field observations, which section 9 excludes
//     by making all six canonical fields required; and
//   - a field carrying an outcome outside the nine-value field vocabulary.
//
// In those cases NO envelope outcome and NO emission disposition is asserted:
// the closed record-discarded trigger list of section 10.4 (R1, R5, R6) is not
// extended, and no seventh rule is invented. The returned Composition then
// carries RuleNone, an unasserted outcome and an empty Disposition, which is
// not a member of the emission vocabulary.
//
// Compose is pure and read-only. It performs no I/O, mutates no input,
// normalizes nothing, derives no target_id, computes no freshness and
// fabricates no timestamp. It also does not enforce that a provenance-invalid
// field carries no normalized value: that invariant belongs to the
// normalization role, which is not realized in this work package.
func Compose(target TargetAttribution, fields []FieldObservation) (Composition, bool) {
	// R1 — target_id missing or its provenance invalid. Evaluated first; the
	// order of section 10.5 is binding. No envelope outcome is asserted.
	if !target.Valid() {
		return Composition{
			Rule:        RuleR1,
			Outcome:     NoOutcomeAsserted(),
			Disposition: DispositionRecordDiscarded,
		}, true
	}

	// Input the contract does not cover. Checked before R2 so that an empty
	// field set cannot satisfy "all fields success" vacuously and fabricate a
	// success the run never produced.
	if len(fields) == 0 {
		return Composition{Rule: RuleNone, Outcome: NoOutcomeAsserted()}, false
	}

	successes := 0
	usable := 0
	var usableOutcomes []ObservationOutcome
	var withheld []FieldIdentity

	for _, f := range fields {
		if !f.Outcome.IsCanonical() {
			return Composition{Rule: RuleNone, Outcome: NoOutcomeAsserted()}, false
		}

		if f.IsProvenanceInvalid() {
			// Section 13: a provenance-invalid field is withheld and is not
			// normalized. Section 10.5: it never yields an envelope value and
			// does not count towards R6.
			withheld = append(withheld, f.Identity)
			continue
		}

		usable++
		if f.Outcome == FieldOutcomeSuccess {
			successes++
		}

		o, ok := f.Outcome.CollectionOutcome()
		if !ok {
			return Composition{Rule: RuleNone, Outcome: NoOutcomeAsserted()}, false
		}
		if !containsOutcome(usableOutcomes, o) {
			usableOutcomes = append(usableOutcomes, o)
		}
	}

	// R2 — all fields success.
	if successes == len(fields) {
		return Composition{
			Rule:        RuleR2,
			Outcome:     AssertOutcome(OutcomeSuccess),
			Disposition: DispositionEmitted,
		}, true
	}

	// R3 — at least one field success and at least one field a different
	// value, provenance-invalid included. The envelope is partial and the
	// record is emitted; each provenance-invalid field is additionally
	// field-withheld. Every per-field cause stays individually visible:
	// partial swallows none of them.
	if successes > 0 {
		return Composition{
			Rule:           RuleR3,
			Outcome:        AssertOutcome(OutcomePartial),
			Disposition:    DispositionEmitted,
			WithheldFields: withheld,
		}, true
	}

	// R4 — no field success, but every non-provenance-invalid field carries
	// the SAME collection outcome X. The envelope value is X itself. No
	// precedence, no severity and no selection among outcomes is applied: with
	// exactly one distinct outcome present there is nothing to select. The
	// guard on usable keeps R4 from firing vacuously on an all-
	// provenance-invalid field set, which belongs to R5.
	//
	// X must further be legally assertable under the no-success premise R4
	// operates in. Section 10.2 defines partial as "at least one field
	// success, at least one field not", so an envelope partial requires at
	// least one success. R4 is reached only when successes == 0, so X ==
	// partial can never be asserted here — asserting it would silently widen
	// partial, which the contract rejects. This condition enforces the
	// existing section 10.2 definition; it is not an additional rule, not a
	// severity or precedence choice, and not a seventh rule.
	if usable > 0 && len(usableOutcomes) == 1 && usableOutcomes[0] != OutcomePartial {
		return Composition{
			Rule:           RuleR4,
			Outcome:        AssertOutcome(usableOutcomes[0]),
			Disposition:    DispositionEmitted,
			WithheldFields: withheld,
		}, true
	}

	// R5 — every otherwise observed field is provenance-invalid. The run
	// produced nothing provenance-bearing: no envelope value is asserted and
	// the record is not emitted as an Observation. The provenance defect
	// record is retained by the caller (section 10.4).
	if usable == 0 {
		return Composition{
			Rule:        RuleR5,
			Outcome:     NoOutcomeAsserted(),
			Disposition: DispositionRecordDiscarded,
		}, true
	}

	// R6 — no field success, at least two non-provenance-invalid fields, and
	// more than one distinct collection outcome among them. Within the bounded
	// eight-value vocabulary no envelope statement would be true, so none is
	// asserted: no synthetic aggregate is formed, partial is not weakened (its
	// required success share is absent), no outcome wins, and no ninth
	// envelope value arises. The record is discarded; the per-field causes,
	// provenance and any raw values are retained by the caller (section 10.4).
	if usable >= 2 && len(usableOutcomes) > 1 {
		return Composition{
			Rule:        RuleR6,
			Outcome:     NoOutcomeAsserted(),
			Disposition: DispositionRecordDiscarded,
		}, true
	}

	// Reached only by a usable field set that satisfies no normative rule:
	// every usable field carries the field-level partial outcome while no
	// field is success, so R4 is barred by the section 10.2 partial
	// definition and R6 does not apply because only one distinct outcome is
	// present. Section 10.3 records that field-level partial does not occur in
	// this slice, so this input lies outside the contract's valid composition
	// domain. Neither an envelope outcome nor an emission disposition is
	// asserted — the same fail-closed non-assertion this function already uses
	// for uncovered input. No replacement outcome is invented, no failure
	// value is substituted and no disposition is fabricated.
	return Composition{Rule: RuleNone, Outcome: NoOutcomeAsserted()}, false
}

// containsOutcome reports whether outcomes already holds o.
func containsOutcome(outcomes []ObservationOutcome, o ObservationOutcome) bool {
	for _, c := range outcomes {
		if c == o {
			return true
		}
	}
	return false
}
