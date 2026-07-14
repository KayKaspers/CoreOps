# CoreOps – Offline Data Reconciliation and Conflict Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation reconciliation and conflict policy
> Implementation Status: Not implemented
> Merge Technology: Not selected
> Synchronisation Protocol: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-011` (docs-only / data-governance and architecture foundation)

## 1. Status

Foundation, technology-independent policy for detecting and resolving data conflicts and reconciling offline/delayed changes. Selects no merge technology or synchronisation protocol and implements nothing. Companion to [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](../architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) and [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](../architecture/FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md).

## 2. Purpose

Ensure conflicts stay visible and are resolved by documented authority (not automatic last-write-wins), and that offline/delayed changes reconcile with provenance, integrity, authority review, and fail-closed handling.

## 3. Scope

Reconciliation concepts/preconditions; conflict detection/states; resolution inputs; overrides; offline changes; delayed observations; revoked sources; replay/duplication; time/sequence uncertainty; partial import; audit; privacy; fail-closed rules; invariants; threat references; exceptions.

## 4. Non-Goals

- No merge technology, CRDT, synchronisation protocol, storage, or clock technology selection.
- No runtime code; no certification assessment; no ADR.
- No identity/module/capability/threat-file change.

## 5. Reconciliation Concepts

Reconciliation aligns local and central state after offline/delayed changes, applying authority, provenance, and conflict rules. `local availability ≠ trustworthiness`; reconciliation confirms consistency and surfaces conflicts (e.g. a locally active but centrally revoked source/credential).

## 6. Authority Preconditions

Reconciliation requires known authority rules for the affected fields (source-of-truth model), source identity and provenance/integrity status, and an owner/approver where the profile requires. On unclear authority, reconciliation is fail-closed (§19).

## 7. Conflict Detection

A conflict exists at least when: multiple authoritative sources disagree; desired and observed state are incompatible; source identity is ambiguous; the same source reports different sequences; offline and online changes collide; an override contradicts a newer policy; an imported value comes from a revoked source; topology/inventory relationships are inconsistent.

## 8. Conflict States

`no-conflict` · `potential-conflict` · `active-conflict` · `resolution-pending` · `resolved` · `accepted-exception` · `invalidated`. Conflicts are not silently discarded; they remain visible until explicitly resolved; `resolved` does not make the source retroactively trustworthy.

## 9. Resolution Inputs

A conceptual decision model considers: field authority · source authority · scope · lifecycle state · revocation state · provenance status · integrity status · freshness · explicit approval · manual override · security invariant. No universal blanket priority list applies to all field types. Instead: field classes need defined authority policies; security-sensitive fields need stricter rules; human override is not automatically highest authority; revoked/invalidated sources must not win; audit/evidence data is not rewritten by normal operator overrides.

## 10. Manual Overrides

Overrides follow the provenance standard §13 (explicit, named-human, reason-/scope-/field-bound, time-bound, auditable, provenance-preserving). An override resolving a conflict is itself audited and visible; it does not delete original provenance or conflicting-source records.

## 11. Offline Changes

Offline change queues, offline overrides, and offline enrollment/credential changes reconcile explicitly with provenance/integrity checks and conflict review; no automatic conflict-free merge; no CRDT/merge technology claimed.

## 12. Delayed Observations

Offline sources may deliver delayed but valid observations. Delayed observations are placed by event/observation time (not merely receipt time) and may create conflicts (§7) that stay visible. `received later ≠ happened later`.

## 13. Revoked Sources

A source revoked (identity/credential/trust) must not remain authoritative even if it delivers later data. Offline revocation distribution is a recognised challenge; a locally-active-but-centrally-revoked source is a conflict requiring review, not silent acceptance. Consistent with `CO-WP-010` revocation governance.

## 14. Replay and Duplication

Replay and duplicate delivery must be later-detectable; a duplicate is not automatically harmless. Ordering does not follow from an untrusted timestamp/sequence alone. Threat refs THR-026/027.

## 15. Time and Sequence Uncertainty

Event/observation/receipt/effective times differ; clocks may be wrong/manipulated; clock skew and sequence gaps surface as uncertainty, not hidden. `timestamp present ≠ timestamp trusted`; `sequence present ≠ sequence complete`. No time/sequence technology selected.

## 16. Partial Import

A partial import is not treated as complete; incomplete imports are flagged, quarantined where needed, and reconciled with explicit status. Provenance/integrity of the package apply (offline enrollment/credential governance).

## 17. Audit and Evidence

Conflict detection, resolution decisions, overrides, offline reconciliation, and invalidations are audit-relevant and attributable (MOD-EVD-001). Audit/evidence history is not rewritten by normal reconciliation. Evidence capability ≠ availability ≠ requirement satisfaction.

## 18. Privacy and Disclosure

Reconciliation and conflict records minimise personal/operational data (provenance standard §17); no unnecessary full source-dataset copies; disclosure limits for exports; consistent with the disclosure policy.

## 19. Fail-Closed Rules

On unclear authority, unresolved conflict on a security-relevant field, unverified provenance/integrity, or unavailable approval, reconciliation is fail-closed: do not silently pick a value, keep the conflict visible, and escalate to the responsible owner. `unknown ≠ healthy`.

## 20. Security Invariants

Design requirements (not implemented controls):

```text
Field conflicts must remain visible until explicitly resolved.
Security-relevant fields must not use automatic last-write-wins.
Latest timestamp must not automatically win.
Revoked sources must not remain authoritative.
Offline reconciliation must not bypass provenance, integrity or approval.
Manual override must preserve original provenance.
Audit and evidence history must not be rewritten by normal reconciliation.
Partial imports must not be treated as complete.
On unclear authority, reconciliation fails closed.
```

## 21. Threat References

THR-011 (API), THR-012/013 (telemetry/stale), THR-014/015 (inventory/topology), THR-016/017/018 (audit/evidence), THR-024 (offline import), THR-026 (replay), THR-027 (time), THR-035 (tenant boundary). No parallel threat list; no invented IDs.

## 22. Exceptions

Any deviation (e.g. an accepted-exception conflict state) requires a documented rationale, a risk entry, Nova review, and Human-Maintainer approval, with visible, auditable compensating handling. No silent exceptions.

## 23. Technology Boundary

No merge, CRDT, synchronisation, storage, or clock technology selected; deferred to later ADR-governed work.

## 24. Compatibility

Additive; consistent with the source-of-truth model, provenance standard, module architecture (MOD-OFF-001/MOD-STA-001/MOD-EVD-001), threat model, offline credential governance, and NDF rules. No technology; no ADR; claims no certification. Breaking-change potential: low.

## 25. Open Questions

- Which conflict-resolution defaults per field class (later)?
- How is offline revocation reliably distributed and reconciled (later)?
- Which reconciliation model per deployment/connectivity tier (later)?

## 26. Next Decision

Nova review, then Human-Maintainer commit. Merge/synchronisation technology and reconciliation implementation remain separate later work packages (ADR-governed).
