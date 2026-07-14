# CoreOps – Drift Detection and Convergence Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation drift and convergence model
> Implementation Status: Not implemented
> Detection Engine: Not selected
> Automated Remediation: Not implemented
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-012` (docs-only / state-management and safe-remediation foundation)

## 1. Status

Foundation, technology-independent model for drift detection, classification, exceptions, and convergence. Selects no detection engine and implements no automatic remediation. Companion to [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) and [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md).

## 2. Purpose

Define what counts as drift, how it is classified and tracked, and what convergence means — so later work detects drift honestly (no observation ≠ no drift; no drift ≠ compliance) and never treats detection as authority to change target systems.

## 3. Scope

Drift definition/types; detection preconditions/states; drift record; impact/urgency; intentional divergence; exceptions; convergence; verification; partial failure; offline; audit; invariants; threat references.

## 4. Non-Goals

- No detection engine, scheduler, queue, or automatic remediation.
- No API/database schema; no retry strategy; no execution technology.
- No ADR; no change to SoT/provenance/reconciliation/identity/module/capability/threat files.

## 5. Drift Definition

Drift is a traceable divergence between a valid desired state and a sufficiently trustworthy observed/effective state. Not every difference is confirmed drift. Before confirming, check: desired-state validity · scope match · field identity · source identity · provenance · freshness · trust status · validation status · conflict state · exception state · comparison semantics.

## 6. Drift Types

`configuration drift` · `version drift` · `deployment drift` · `policy drift` · `permission drift` · `resource-registration drift` · `inventory drift` · `topology drift` · `integration-capability drift` · `agent or adapter drift` · `credential-reference drift` · `evidence or audit-availability drift`. A drift type is not a statement about cause, fault, or compliance.

## 7. Detection Preconditions

Detection requires a valid active desired state, a sufficiently fresh/trusted observed or effective state, matching scope and canonical field identity, and no blocking conflict/unknown authority. Otherwise the detection status reflects the gap (`insufficient-data`, `stale-observation`, `active-conflict`) rather than asserting drift or no-drift.

## 8. Detection States

`not-assessed` · `insufficient-data` · `comparison-pending` · `no-drift-observed` · `potential-drift` · `confirmed-drift` · `stale-observation` · `active-conflict` · `exception-active` · `remediation-pending` · `remediation-in-progress` · `verification-pending` · `resolved` · `invalidated`.

```text
no-drift-observed ≠ compliant · insufficient-data ≠ no drift ·
stale-observation ≠ confirmed current state ·
resolved ≠ technically verified unless verification evidence exists
```

## 9. Drift Record

Conceptual minimum (no JSON/DB structure): drift record identity · entity/resource identity · field/state identity · workspace/environment scope · desired-state reference · observed-state reference · effective-state reference · drift type · detection status · detected time · freshness status · authority status · conflict status · exception status · impact · urgency · owner · recommended action · approval requirement · remediation reference · verification reference · audit reference.

## 10. Impact and Urgency

Impact: `low` · `medium` · `high` · `critical` · `unassessed`. Urgency: `routine` · `planned` · `prompt` · `immediate-review` · `unassessed`. These are qualitative, not a mathematical risk computation; `critical` must not automatically lead to unreviewed execution.

## 11. Intentional Divergence

A divergence counts as intentional/accepted only if explicit, named-human attributable, reason-bound, scope-bound, time-bound where appropriate, risk-visible, auditable, reviewable, and revocable.

## 12. Exceptions

Exception record: exception identity · human owner · reason · scope · affected desired-state reference · affected fields/resources · start · expiry/review date · approval · risk reference · compensating controls · audit reference. Not permitted: silent drift suppression; unbounded exception without review; automatic propagation to other resources; a machine principal as sole risk acceptor; an exception that rewrites audit history.

## 13. Convergence

Convergence means a valid desired state and a sufficiently fresh, trusted observed state match per documented comparison rules. Convergence is not automatically compliance, security validation, business correctness, successful deployment certification, or permanent stability.

Convergence states: `not-assessed` · `not-converged` · `convergence-pending` · `apparently-converged` · `verified-converged` · `temporarily-divergent` · `exception-authorised` · `indeterminate`. `verified-converged` requires actual verification evidence.

## 14. Verification

Verification uses a fresh, trusted observation (ideally not the same potentially-compromised source used for execution) to confirm the target reached the desired state. `executed ≠ successful`; `successful ≠ verified convergence`; `verified convergence ≠ compliance`. No verification technology selected.

## 15. Partial Failure

Consider: mixed results across resources; multi-step partial failure; execution succeeded but verification failed; partial rollback; unreachable managed resource; stale post-observation; conflicting agent/adapter data; unavailable audit/evidence export.

```text
partial success ≠ complete success · unknown result ≠ failed ≠ successful ·
rollback requested ≠ rollback completed · rollback completed ≠ original state restored unless verified
```

## 16. Offline and Delayed State

Delayed observations, delayed desired-state distribution, offline exceptions/plans/execution records, delayed revocation, replayed reports, clock uncertainty, partial synchronisation, and conflicting online/offline authority are handled provenance-aware, scope-bound, conflict-aware, audit-preserving, and fail-closed for unclear privileged authority (reconciliation policy). No automatic conflict-free synchronisation claimed.

## 17. Audit and Evidence

Drift detection, classification, exceptions, and convergence/verification outcomes are audit-relevant and attributable (MOD-EVD-001). `evidence capability ≠ evidence available`; `evidence available ≠ verified success`; `verified convergence ≠ requirement satisfaction`. Audit history is not rewritten.

## 18. Security Invariants

Design requirements (not implemented controls):

```text
No observation must not be interpreted as no drift.
No detected drift must not be interpreted as compliance.
Unknown or stale state must not be interpreted as healthy or current.
Drift detection must not grant write authority.
Unresolved conflicts must block automatic privileged remediation.
Exceptions must remain explicit, time- or review-bound and auditable.
Verified convergence must not be interpreted as compliance.
```

## 19. Threat References

THR-005 (job manipulation), THR-012/013 (telemetry/stale), THR-014/015 (inventory/topology), THR-028/029/030 (queued/executed/successful/compliant), THR-031 (partial failure), THR-032 (rollback), THR-016/017 (audit), THR-024 (offline), THR-026 (replay). No parallel threat list; no invented IDs.

## 20. Technology Boundary

No detection engine, scheduler, queue, retry, or remediation technology selected; deferred to later ADR-governed work.

## 21. Compatibility

Additive; consistent with the state model, source-of-truth/provenance model, reconciliation policy, module architecture, threat model, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 22. Open Questions

- Which comparison semantics per drift type (later)?
- Which impact/urgency defaults per field class (later)?
- How is verification independence enforced (later)?

## 23. Next Decision

Nova review, then Human-Maintainer commit. Detection-engine technology and remediation implementation remain separate later work packages (ADR-governed).
