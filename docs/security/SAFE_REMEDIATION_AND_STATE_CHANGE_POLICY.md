# CoreOps – Safe Remediation and State Change Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation safe-remediation policy
> Implementation Status: Not implemented
> Automatic Remediation: Not selected
> Execution Technology: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-012` (docs-only / state-management and safe-remediation foundation)

## 1. Status

Foundation, technology-independent policy for safe remediation and state change. It strictly separates detection, recommendation, planning, approval, execution, and verification, and selects no execution technology or automatic remediation. Companion to [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](../architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) and [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](../architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md).

## 2. Purpose

Ensure drift is never fixed by the component that detected it without explicit authority and approval, and that "executed" is never mistaken for "successful", "successful" for "verified converged", or "converged" for "compliant".

## 3. Scope

Authority model; read-only boundary; detection-vs-execution; recommendation; planning; approval; execution; verification; remediation lifecycle; exceptions; partial failure; rollback/recovery; offline; audit; fail-closed rules; invariants; threat references.

## 4. Non-Goals

- No drift/policy engine, scheduler, queue, retry, or remote-execution technology.
- No automatic remediation; no runtime execution; no ADR.
- No change to SoT/provenance/reconciliation/identity/module/capability/threat files.

## 5. Authority Model

Remediation authority derives from principal/role/scope/lifecycle/approval/policy state (`CO-WP-009`/`CO-WP-010`/`CO-WP-011`). Observation/drift modules (MOD-OBS-001/detection) hold no execution authority; execution is MOD-EXE-001 only on approved actions; policy/approval is MOD-POL-001. Repository Human-Maintainer authority ≠ runtime authority.

## 6. Read-only Boundary

```text
observe · detect · recommend · simulate/preview · prepare plan · approve · execute · verify
```

```text
read-only discovery must not silently gain write authority ·
drift detection must not activate write access · preview must not mutate target state ·
approval must not execute the action · verification must not silently retry or remediate
```

## 7. Detection versus Execution

Detection identifies drift; it does not approve or execute. A detector never self-grants privileged execution authority. Threat refs THR-007 (read-to-write), THR-005 (job manipulation).

## 8. Recommendation

A recommendation is advisory (`recommendation ≠ command`); it proposes an action for planning/approval and carries drift/impact references. Recommendation does not schedule or execute.

## 9. Planning

A plan describes intended actions, affected resources, risk/impact, execution boundary, and rollback/recovery expectation. `plan ≠ execution authority`. Plans are auditable and reviewable.

## 10. Approval

Privileged remediation requires explicit approval (MOD-POL-001), attributable to a named identity, with separation of duties where the profile allows (requester ≠ approver). `approval ≠ execution`. Approval is time-/scope-bound and audited (consistent with `CO-WP-009`).

## 11. Execution

Execution runs only approved, in-scope actions via MOD-EXE-001 and authorised adapter/agent paths. `executed ≠ successful`. Execution records result and evidence; unverified artifacts are not trusted for privileged execution.

## 12. Verification

Verification confirms convergence with a fresh, trusted observation, ideally independent of the execution source. `successful ≠ verified convergence`; `verified convergence ≠ compliance`. Verification does not silently retry or remediate.

## 13. Remediation Lifecycle

`identified` · `analysis-pending` · `plan-proposed` · `approval-pending` · `approved` · `scheduled` · `execution-pending` · `executing` · `executed` · `partial-failure` · `failed` · `verification-pending` · `verified-converged` · `verified-not-converged` · `rollback-pending` · `rolled-back` · `exception-open` · `closed`. Ground rules: approved ≠ executed; executed ≠ successful; successful ≠ verified; verified-converged needs current observation + evidence; failure/partial-failure stay visible; retry does not lose original history.

## 14. Exceptions

An accepted divergence follows the drift model §11–§12 (explicit, named-human, reason-/scope-/time-bound, auditable, reviewable). Exceptions do not suppress drift silently, propagate automatically, or rewrite audit history; a machine principal is not a sole risk acceptor.

## 15. Partial Failure

Partial failure is explicitly represented (drift model §15): `partial success ≠ complete success`; `unknown result ≠ failed ≠ successful`. Mixed-result remediations surface per-resource status and remain visible until resolved.

## 16. Rollback and Recovery

Rollback/recovery expectations are defined per plan; `rollback requested ≠ completed`; `rollback completed ≠ original state restored unless verified`. Rollback claims require verification evidence (invariant). No rollback technology selected. Threat ref THR-032.

## 17. Offline Remediation

Offline remediation plans/execution records reconcile provenance-aware, scope-bound, conflict-aware, audit-preserving, and fail-closed for unclear privileged authority (reconciliation policy). Delayed revocation and conflicting online/offline authority are handled explicitly; no automatic conflict-free merge; no classified-network claim.

## 18. Audit and Evidence

Detection, recommendation, plan, approval decision, execution request/result, verification, exception, rollback, and closure are audit-relevant and attributable (MOD-EVD-001). `evidence capability ≠ available ≠ verified success ≠ requirement satisfaction`. Audit history is not rewritten.

## 19. Fail-Closed Rules

On unclear authority, unresolved conflict on a security-relevant field, insufficient/stale observation for a privileged change, indeterminate/conflicted effective state, or unavailable approval: **no automatic privileged remediation**; keep drift/conflict visible; escalate to the responsible owner. `unknown ≠ healthy`.

## 20. Security Invariants

Design requirements (not implemented controls):

```text
Drift detection must not grant write authority.
Recommendation must not become execution automatically.
Approval must precede privileged remediation where required.
Executed must not be interpreted as successful.
Successful must not be interpreted as verified convergence.
Verified convergence must not be interpreted as compliance.
Unresolved conflicts must block automatic privileged remediation.
Rollback claims require verification evidence.
Exceptions must remain explicit, time- or review-bound and auditable.
Verification must not silently retry or remediate.
```

## 21. Threat References

THR-003 (approval bypass), THR-005 (job manipulation), THR-007 (read-to-write), THR-026 (replay), THR-028/029/030 (queued/executed/successful/compliant), THR-031 (partial failure), THR-032 (rollback), THR-016/017 (audit), THR-024 (offline). No parallel threat list; no invented IDs.

## 22. Technology Boundary

No execution, scheduler, queue, retry, policy-engine, or automatic-remediation technology selected; deferred to later ADR-governed work.

## 23. Compatibility

Additive; consistent with the state and drift models, identity/RBAC governance, module architecture (MOD-POL-001/MOD-WFL-001/MOD-EXE-001/MOD-EVD-001), threat model, and NDF rules. No technology; no ADR; claims no certification. Breaking-change potential: low.

## 24. Open Questions

- Which remediation approvals per impact class and profile (later)?
- How is verification-source independence enforced (later)?
- Whether any bounded automatic remediation is later allowed (separate ADR-governed decision)?

## 25. Next Decision

Nova review, then Human-Maintainer commit. Execution technology, remediation implementation, and any automatic-remediation decision remain separate later work packages (ADR-governed).
