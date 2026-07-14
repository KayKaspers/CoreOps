# CoreOps – Observed, Desired and Effective State Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation state semantics
> Implementation Status: Not implemented
> State Engine: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-012` (docs-only / state-management and safe-remediation foundation)

## 1. Status

Foundation, technology-independent semantics for desired, observed, effective, reported, and last-known state, building on the source-of-truth and provenance model (`CO-WP-011`). Selects no state engine and implements nothing. Companions: [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md) and [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md).

## 2. Purpose

Make state distinctions and authority explicit so later state-management work can compute effective state and detect drift honestly — without treating observed as desired, a missing observation as no drift, or an unresolved conflict as a normal trusted state.

## 3. Scope

State concepts; desired-state lifecycle; observed-state model; effective-state boundary; reported/last-known state; unknown/conflicted state; state authority; freshness/validation; state transitions; offline; security invariants; threat references.

## 4. Non-Goals

- No state engine, scheduler, queue, event bus, policy engine, or automatic remediation.
- No API/database schema, retry strategy, agent protocol, or remote-execution technology.
- No actual configuration change; no ADR; no change to SoT/provenance/reconciliation/identity/module/capability/threat files.

## 5. Concepts

`desired state` · `proposed desired state` · `approved desired state` · `active desired state` · `observed state` · `reported state` · `last-known state` · `effective state` · `derived state` · `unknown state` · `conflicted state` · `drift` · `potential drift` · `confirmed drift` · `intentional divergence` · `accepted exception` · `convergence` · `remediation` · `verification` · `partial failure`.

Ground rules:

```text
reported state ≠ trusted observed state · observed state ≠ desired state ·
effective state ≠ compliance · no detected drift ≠ proof of correctness ·
no observation ≠ no drift · remediation requested ≠ approved ≠ executed ≠ successful ·
remediation successful ≠ convergence verified
```

## 6. Desired-State Lifecycle

States (conceptual guidance, not implemented enums): `proposed` · `review-pending` · `approved` · `active` · `superseded` · `suspended` · `withdrawn` · `expired` · `conflicted`.

Desired-state record (minimum): desired-state identity · owner · scope · target entities/resource class · authority source · policy/configuration reference · proposal time · approval state · activation state · effective period · supersession reference · exception reference · audit reference.

Ground rules: only authorised sources create active desired state; a proposal is not active; approval is not execution; expired/withdrawn desired state does not stay silently active; original history is not deleted on supersession.

## 7. Observed-State Model

Observed-state record (minimum): source identity · source class · entity identity · field identity · observed value classification · observation time · received time · freshness · trust status · validation status · provenance reference · sequence/correlation reference · conflict state.

```text
recently received ≠ recently observed · observation present ≠ observation trusted ·
agent report ≠ physical truth · external API response ≠ authoritative desired state ·
stale observation ≠ current state
```

## 8. Effective-State Boundary

Effective state is the state CoreOps currently uses after documented authority/provenance/conflict/freshness/override rules. It must **not** be produced automatically when authority is unclear, active conflicts exist, required provenance is missing, a relevant source is revoked, freshness is insufficient, or a partial import applies.

Effective-state values: `determined` · `determined-with-warning` · `last-known` · `indeterminate` · `conflicted` · `unavailable` · `not-applicable`.

```text
Unresolved authority or provenance conflict must not result in a normal trusted effective state.
For security-relevant or privileged changes: indeterminate or conflicted effective state
→ no automatic privileged remediation.
```

## 9. Reported and Last-Known State

**Reported state** is a source's value before authority/conflict rules apply. **Last-known state** is the most recent sufficiently-trusted value with its own freshness; `last-known ≠ current`. Neither is automatically the effective state.

## 10. Unknown and Conflicted State

`unknown` is not positive/healthy; `conflicted` is not silently resolved (source-of-truth model §7/§15). Unknown/conflicted state stays visible and blocks a normal trusted effective state for the affected field.

## 11. State Authority

State authority follows the [source-of-truth model](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md): desired state owned via MOD-POL-001/MOD-WFL-001/MOD-STA-001; observed via MOD-OBS-001; one authoritative owner per field concept; observed state does not silently overwrite desired.

## 12. Freshness and Validation

Freshness (`current`…`unknown`) and validation status follow the [provenance standard](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md). `stale ≠ current`; `unknown ≠ healthy`; `recent timestamp ≠ trusted`. Freshness bounds are per field/data class, defined later.

## 13. State Transitions

Transitions (desired proposed→approved→active→superseded; observed received→evaluated; effective determined/indeterminate/conflicted) are auditable and provenance-preserving. No transition silently upgrades trust, overwrites desired with observed, or discards conflict. No transition engine selected.

## 14. Offline Considerations

Delayed observations, delayed desired-state distribution, and conflicting online/offline authority are handled per the [reconciliation policy](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md): provenance-aware, scope-bound, conflict-aware, audit-preserving, fail-closed for unclear privileged authority. No automatic conflict-free synchronisation claimed.

## 15. Security Invariants

Design requirements (not implemented controls):

```text
Observed state must not silently overwrite desired state.
Unknown or stale state must not be interpreted as healthy or current.
No observation must not be interpreted as no drift.
Unresolved authority or provenance conflict must not produce a normal trusted effective state.
A summary or dashboard must not become an authoritative state source.
Effective state is not compliance.
```

## 16. Threat References

Relevant existing scenarios: THR-005 (job manipulation), THR-012/013 (telemetry/stale), THR-014/015 (inventory/topology), THR-024 (offline import), THR-026 (replay), THR-027 (time). No parallel threat list; no invented IDs.

## 17. Technology Boundary

No state engine, storage, scheduler, queue, event bus, or policy-engine technology selected; deferred to later ADR-governed work.

## 18. Compatibility

Additive; consistent with the source-of-truth model, provenance standard, reconciliation policy, module architecture (MOD-STA-001/MOD-OBS-001/MOD-POL-001), threat model, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 19. Open Questions

- Which effective-state computation rules per field class (later)?
- How are last-known and freshness thresholds set (later)?
- Which state-engine model per deployment (later)?

## 20. Next Decision

Nova review, then Human-Maintainer commit. State-engine technology and effective-state computation remain separate later work packages (ADR-governed).
