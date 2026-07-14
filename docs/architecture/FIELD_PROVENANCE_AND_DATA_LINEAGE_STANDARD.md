# CoreOps – Field Provenance and Data Lineage Standard

> Document Status: Implemented, pending Nova review
> Standard Status: Foundation field-provenance and lineage standard
> Implementation Status: Not implemented
> Field Schema: Not selected
> Cryptographic Provenance Mechanism: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-011` (docs-only / data-governance and architecture foundation)

## 1. Status

Foundation, technology-independent standard for field-level provenance and data lineage: field identity, provenance metadata, freshness, trust/validation status, transformation lineage, imported/derived data, and privacy. Selects no field schema or cryptographic provenance mechanism and implements nothing. Companion to [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) and [OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md).

## 2. Purpose

Make each field's origin, authority, freshness, trust, and lineage explicit and attached through transformations and overrides — so later data-architecture work can implement provenance without losing the authority or uncertainty of inputs.

## 3. Scope

Field identity; provenance metadata; source identity; authority/freshness/trust/validation status; transformation lineage; manual corrections; imported/derived data; audit references; privacy; offline; invariants; threat references.

## 4. Non-Goals

- No field schema, JSON/database schema, cryptographic signature field, or timestamp format selection.
- No runtime code; no full field-ID list; no ADR.
- No identity/module/capability/threat-file change.

## 5. Field Identity

Field references are conceptually stable and unique, derived from a canonical CoreOps field identity — **not** from UI labels, translated display names, table positions, or temporary adapter fields.

```text
display name ≠ stable field identity
adapter field ≠ canonical CoreOps field
translated label ≠ field identifier
```

No full field-ID list is created here.

## 6. Provenance Metadata

Conceptual minimum model (no concrete schema): Entity identity · Field identity/path · Value classification · Source identity · Source class · Source reference · Authority class · Owning module · Workspace/environment scope · Resource scope · Observed/produced time · Received time · Effective time · Freshness state · Trust state · Validation state · Transformation/derivation reference · Correlation/operation reference · Override reference · Conflict state · Audit reference · Retention owner. Not required: real sensitive field values in governance docs; concrete JSON/DB schemas; cryptographic signature fields; fixed timestamp formats.

## 7. Source Identity

Each field value carries an attributable source identity and source class (human/module/managed-resource/agent/adapter/external/import/migration/derived/manual). Source identity binds to an owner (human or governed machine principal per `CO-WP-009`/`CO-WP-010`). `successful import ≠ validated provenance`.

## 8. Authority Status

Each field value carries an authority class (source-of-truth model §7). Authority status is not validation, compliance, or freshness. Multiple sources may report a field; only the documented authority may set the effective value.

## 9. Freshness

Freshness states (conceptual guidance, not implemented enums): `current` · `aging` · `stale` · `expired` · `unknown` · `not-applicable`. `stale ≠ current`; `unknown ≠ healthy`; `recently received ≠ recently observed`; `recent timestamp ≠ trusted timestamp`; `cached value ≠ live value`. No universal time bounds; freshness bounds are later definable per field/data class.

## 10. Trust Status

`unassessed` · `untrusted` · `conditionally-trusted` · `trusted-for-specified-scope` · `revoked` · `conflicted`. Governance guidance only; no technical/cryptographic check claimed. `trusted-for-specified-scope` is field/source/operation-scoped, not a global trust grant.

## 11. Validation Status

`not-validated` · `format-validated` · `provenance-checked` · `integrity-checked` · `semantically-reviewed` · `operationally-verified` · `invalid`. Governance guidance; no implemented validation claimed. Evidence capability ≠ evidence available ≠ requirement satisfaction.

## 12. Transformation and Lineage

Derived values retain their input references. Minimum lineage: derivation identity · input references · transformation class · producing module · produced time · algorithm/rule-version reference · validation status · freshness dependency · conflict propagation · audit reference.

```text
A derived value must not hide the authority or uncertainty of its inputs.
```

No transformation engine or algorithm language selected.

## 13. Manual Corrections

Manual corrections/overrides record: override identity · human actor · reason · affected entity · affected field · previous effective value reference · override value classification · scope · start · expiry · approval · risk/exception reference · audit reference · review owner. Not permitted: override without provenance; silent replacement of the original source; unbounded override without review; changing audit history; automatic global effect; override by a machine principal without an explicitly allowed workflow. No override engine selected.

## 14. Imported Data

```text
import ≠ authority transfer · successful parsing ≠ provenance validation ·
external API response ≠ trusted effective state · vendor data ≠ CoreOps policy ·
offline package metadata ≠ verified package
```

Imported data retains at least: source identity · source class · scope · import time · original time · provenance status · integrity status · validation status · conflict status · import operation reference.

## 15. Derived Data

Derived data is not authoritative by default; it carries lineage (§12) and propagates conflict/uncertainty from inputs. A dashboard/summary derived view is never an authoritative source (repository-governance source-of-truth rule).

## 16. Audit References

Provenance changes, overrides, imports, derivations, and conflict resolutions carry audit references (MOD-EVD-001), attributable to a source/owner. Audit history is not rewritten by normal reconciliation. No secrets in provenance/audit.

## 17. Privacy and Data Minimization

Provenance metadata may be personal or operationally sensitive. Record only necessary identity references; no unnecessary plaintext credentials; no unfiltered copy of complete source datasets; disclosure limits for exports; pseudonymous/technical references where appropriate; retention decided by responsible owners; audit attributability without unnecessary data collection. No concrete retention periods. Consistent with the [Public Neutrality and Disclosure Policy](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md) and BSI baseline PSR-15.

## 18. Offline Considerations

Offline-produced provenance retains source/original-time/provenance/integrity status and is reconciled per the [reconciliation policy](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md); local availability ≠ trustworthiness. No cryptographic mechanism selected; no classified-network claim.

## 19. Security Invariants

Design requirements (not implemented controls):

```text
Field provenance must remain attached through transformations, overrides and reconciliation.
Imported data must not inherit authority automatically.
Manual override must preserve original provenance.
Derived state must not hide the authority or uncertainty of its inputs.
Stale must not be interpreted as current; unknown must not be interpreted as healthy.
Latest timestamp must not automatically win.
Provenance metadata must not expose secrets or excessive personal data.
Adapter/display field identity must not replace canonical field identity.
```

## 20. Threat References

THR-011 (API), THR-012/013 (telemetry/stale), THR-014/015 (inventory/topology), THR-016/017 (audit), THR-024 (offline import), THR-026 (replay), THR-027 (time). No parallel threat list; no invented IDs.

## 21. Technology Boundary

No field schema, cryptographic provenance, signature, timestamp format, or storage technology selected; deferred to later ADR-governed work.

## 22. Compatibility

Additive; consistent with the source-of-truth model, module architecture, threat model, disclosure policy, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 23. Open Questions

- Which provenance fields are mandatory vs. optional per data class (later)?
- Whether cryptographic provenance is used for high-trust fields (later)?
- How is canonical field identity defined and versioned (later)?

## 24. Next Decision

Nova review, then Human-Maintainer commit. Field schema, provenance mechanism, and lineage implementation remain separate later work packages (ADR-governed).
