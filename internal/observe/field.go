package observe

// Canonical field identities of the slice (Observation Contract section 9).
// Field names are canonically English. The list is exactly these six; this
// package neither adds a field nor renames one, and it does not produce a
// complete CoreOps field-ID list — canonical field IDs and their versioning
// are deliberately left open by the contract (section 22).
type FieldIdentity string

const (
	// FieldTargetID is the observation target identity (section 8.2). Its
	// derivation rule is OPEN (sections 8.3 and 22) and is not implemented
	// anywhere in this package.
	FieldTargetID FieldIdentity = "target_id"

	// FieldHostname is the observed hostname of the local host. A hostname is
	// not an identity, not an FQDN, not DNS resolvability and not unique.
	FieldHostname FieldIdentity = "hostname"

	// FieldOSIdentity is the operating-system / distribution identity. It is
	// not a support status, not compatibility, not a lifecycle status and not
	// a vendor relationship.
	FieldOSIdentity FieldIdentity = "os_identity"

	// FieldOSRelease is the operating-system release/version. It is not a
	// patch level, not currency, not absence of vulnerabilities and not a need
	// to update. It is never inferred from os_identity.
	FieldOSRelease FieldIdentity = "os_release"

	// FieldKernelRelease is the kernel release identifier. It is not the
	// operating-system version and not a security level; a running kernel is
	// not an installed kernel. It is never inferred from os_release.
	FieldKernelRelease FieldIdentity = "kernel_release"

	// FieldArchitecture is the hardware / instruction architecture. It is not
	// a CPU model, not a virtualization status and not a compatibility
	// commitment. It has no default.
	FieldArchitecture FieldIdentity = "architecture"
)

// canonicalFields is the closed set of canonical field identities of this
// slice, in contract order (section 9).
var canonicalFields = [...]FieldIdentity{
	FieldTargetID,
	FieldHostname,
	FieldOSIdentity,
	FieldOSRelease,
	FieldKernelRelease,
	FieldArchitecture,
}

// CanonicalFields returns the canonical field identities of the slice in
// contract order. The returned slice is a copy: the canonical set is not
// extensible from outside this package.
func CanonicalFields() []FieldIdentity {
	out := make([]FieldIdentity, len(canonicalFields))
	copy(out, canonicalFields[:])
	return out
}

// IsCanonical reports whether f is one of the six canonical field identities.
// Membership only: it asserts nothing about whether the field was observed,
// is valid, is trusted or is supported.
func (f FieldIdentity) IsCanonical() bool {
	for _, c := range canonicalFields {
		if c == f {
			return true
		}
	}
	return false
}

// TargetID is the observed target identity within the observation domain
// (section 8.2). It is an observation reference identity and nothing else:
//
//	target_id                != managed-resource identity
//	target_id present        != ownership
//	target_id present        != management authority
//	target_id present        != authorized target
//	target_id present        != enrollment
//	target_id stable         != identity verified
//	DISCOVERED               != MANAGED
//	DISCOVERED               != TRUSTED
//	hostname                 != identity
//
// The derivation rule for target_id is OPEN (sections 8.3 and 22) and bound to
// the concrete P-3 realization. This package derives nothing: there is no
// generator, no UUID strategy, no hostname fallback, no machine-id fallback,
// no hashing and no default. A target_id is supplied to this package or it is
// absent.
type TargetID string

// TargetAttribution carries target_id together with its provenance (sections
// 8.2, 9 and 13). It is the subject attribution that rule R1 evaluates.
//
// The zero value is the absence of target_id. An empty TargetID is absence,
// never an anonymous identity: without target_id the envelope is invalid and
// the observation is discarded rather than stored anonymously (section 9).
type TargetAttribution struct {
	id         TargetID
	provenance FieldProvenance
}

// NoTargetAttribution returns the attribution of a run that produced no
// target_id.
func NoTargetAttribution() TargetAttribution {
	return TargetAttribution{}
}

// AttributedTarget returns the attribution for a supplied target_id and its
// provenance. An empty id yields an absent attribution.
func AttributedTarget(id TargetID, p FieldProvenance) TargetAttribution {
	return TargetAttribution{id: id, provenance: p}
}

// ID reports target_id and whether it is present.
func (t TargetAttribution) ID() (TargetID, bool) {
	return t.id, t.id != ""
}

// Provenance reports the provenance of target_id.
func (t TargetAttribution) Provenance() FieldProvenance {
	return t.provenance
}

// Valid reports whether the subject attribution holds: target_id present and
// its required provenance elements present. It is the negation of the R1
// trigger of section 10.5 — target_id missing or its provenance invalid.
func (t TargetAttribution) Valid() bool {
	if _, ok := t.ID(); !ok {
		return false
	}
	return t.provenance.RequiredElementsPresent()
}

// RawValue is the source-near observed value of a field (section 12): original
// semantics, unchanged, carried with its source and provenance.
//
//	raw != automatically trustworthy
//	raw != authoritative state
//	raw != validated
//
// Absence is carried positively. The zero value is the absence of a raw value,
// not an empty observed value: an absent field is never an empty string, never
// zero, never a placeholder and never a default (sections 9 and 15). Without a
// raw value no observation of the field exists.
type RawValue struct {
	present bool
	value   string
}

// NoRawValue returns the raw value of a field that was not observed.
func NoRawValue() RawValue {
	return RawValue{}
}

// ObservedRawValue returns a raw value that was actually observed. The caller
// supplies it; this package observes nothing.
func ObservedRawValue(v string) RawValue {
	return RawValue{present: true, value: v}
}

// Value reports the raw value and whether one exists. The value is meaningless
// when the second result is false.
func (r RawValue) Value() (string, bool) {
	return r.value, r.present
}

// NormalizedValue is the canonical CoreOps representation of a field
// (section 12). It is conditional: it exists only where normalization actually
// succeeded.
//
//	normalized            != lossless
//	normalized            != validated
//	normalization success != semantic validation
//	absent                != normalized to a default
//
// Raw and Normalized stay separate: a normalized value never replaces,
// overwrites or discards its raw value, which remains present, referenceable
// and visible — including where normalization failed.
//
// This package implements NO normalization transformation. There is no
// lowercasing rule, no trimming policy, no canonical mapping table, no OS
// alias map, no architecture mapping, no release parsing and no host
// normalization. Concrete normalization is not decided by the contract; only
// the raw/normalized structure is realized here.
type NormalizedValue struct {
	present bool
	value   string
}

// NoNormalizedValue returns the normalized value of a field that was not
// normalized — including a field withheld as provenance-invalid, which is
// never normalized (section 13).
func NoNormalizedValue() NormalizedValue {
	return NormalizedValue{}
}

// NormalizedTo returns a normalized value produced by a successful
// normalization. The caller supplies it; this package transforms nothing.
func NormalizedTo(v string) NormalizedValue {
	return NormalizedValue{present: true, value: v}
}

// Value reports the normalized value and whether one exists. The value is
// meaningless when the second result is false.
func (n NormalizedValue) Value() (string, bool) {
	return n.value, n.present
}

// FieldObservation is the result for exactly one canonical field (section 9).
//
// It keeps raw and normalized separate, carries the field-level outcome from
// the nine-value vocabulary (section 10.3), the per-field provenance
// (section 13) and the freshness state (section 14). Together with its
// provenance it carries the conceptual minimum set of section 13 for the
// field; nothing is duplicated between the two.
//
// Not this type's responsibility: collecting the value, normalizing it,
// deriving freshness from any age computation, deciding the envelope outcome,
// or deciding emission — those are composition (composition.go) and the
// unauthorized collection role.
type FieldObservation struct {
	// Identity is the canonical field identity (section 9).
	Identity FieldIdentity

	// Raw is raw_observed_value (sections 9 and 12), required where the field
	// was observed at all.
	Raw RawValue

	// Normalized is normalized_value (sections 9 and 12), conditional on
	// successful normalization.
	Normalized NormalizedValue

	// Outcome is field_observation_outcome from the nine-value field
	// vocabulary (section 10.3). It is not an envelope outcome and not an
	// emission disposition.
	Outcome FieldOutcome

	// Provenance is the per-field provenance set (section 13). Without usable
	// provenance the field outcome is provenance-invalid and the field is
	// withheld and not normalized (fail-closed).
	Provenance FieldProvenance

	// Freshness is the per-field freshness state (section 14), required for
	// each observed field.
	Freshness Freshness
}

// IsProvenanceInvalid reports whether the field carries the provenance-invalid
// field outcome. Such a field is withheld, is not normalized, never yields an
// envelope value, and does not count towards rule R6 (sections 10.5 and 13).
func (f FieldObservation) IsProvenanceInvalid() bool {
	return f.Outcome == FieldOutcomeProvenanceInvalid
}
