package observe

// This file realizes the bounded freshness vocabulary of Observation Contract
// section 14.
//
// No new CoreOps freshness taxonomy is created. The slice USES the existing
// freshness state set of the Field Provenance standard section 9, bounded to
// the values materially applicable to identity and basic system facts.
//
// Freshness asserts nothing about the target and nothing about trust:
//
//	fresh             != healthy
//	fresh             != trusted
//	fresh             != authorized
//	stale             != unavailable
//	stale             != invalid
//	unknown freshness != current

// Freshness is the per-field freshness state (section 14). Exactly the five
// values below; no others.
//
// The value expired is deliberately NOT part of this vocabulary: it
// presupposes a declared validity period that the contract does not define.
type Freshness string

const (
	// FreshnessCurrent — the observation lies within the age limit accepted
	// for the decision.
	FreshnessCurrent Freshness = "current"

	// FreshnessAging — the observation is visibly ageing but still within the
	// limit.
	FreshnessAging Freshness = "aging"

	// FreshnessStale — the observation lies outside the limit and is to be
	// carried visibly as such. Stale is not unavailable and not invalid.
	FreshnessStale Freshness = "stale"

	// FreshnessUnknown — observed_at is unknown, or no limit is bound.
	// Unknown freshness is not current.
	FreshnessUnknown Freshness = "unknown"

	// FreshnessNotApplicable — no meaningful age binding exists for this
	// field. This is distinct from unknown and never interchangeable with it.
	FreshnessNotApplicable Freshness = "not-applicable"
)

// freshnessStates is the closed five-value freshness vocabulary of the slice
// (section 14), in contract order.
var freshnessStates = [...]Freshness{
	FreshnessCurrent,
	FreshnessAging,
	FreshnessStale,
	FreshnessUnknown,
	FreshnessNotApplicable,
}

// FreshnessStates returns the five freshness values in contract order. The
// returned slice is a copy: the vocabulary is not extensible from outside this
// package, and this package has no authority to extend it.
func FreshnessStates() []Freshness {
	out := make([]Freshness, len(freshnessStates))
	copy(out, freshnessStates[:])
	return out
}

// IsCanonical reports whether f is one of the five freshness values.
// Membership only.
//
// No freshness value is DERIVED here. Concrete age thresholds are
// PROPOSED / UNACCEPTED: the repository supplies no binding values, and the
// threshold choice is a later, explicitly Human-Maintainer-bound decision that
// depends on the collection cadence — which is not selected
// (mechanism class selected != collection cadence selected).
//
// This package therefore implements: no numeric threshold, no default limit,
// no age computation, no threshold function, no configuration, and no implicit
// interpretation of any time difference. A freshness state is supplied to this
// package or it is unknown.
func (f Freshness) IsCanonical() bool {
	for _, c := range freshnessStates {
		if c == f {
			return true
		}
	}
	return false
}
