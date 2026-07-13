# CoreOps – Module Boundary and Dependency Standard

> Document Status: Implemented, pending Nova review
> Standard Status: Foundation module-boundary standard
> Automated Enforcement: Not implemented
> Technology Selection: None
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-008` (docs-only / logical-architecture foundation)

## 1. Status

Foundation standard for module boundaries and dependencies: authority boundaries, data ownership, allowed/forbidden dependencies, communication concepts, and per-boundary rules. It introduces no automated enforcement and selects no technology. Companion to [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md) and [COREOPS_MODULE_CATALOG.md](COREOPS_MODULE_CATALOG.md).

## 2. Purpose

Keep module boundaries and dependency direction explicit so authority separation and security invariants survive into later component design, and so no module silently gains authority it should not have.

## 3. Scope

Module-vs-component distinction; authority and data-ownership boundaries; dependency direction; allowed/forbidden dependencies; communication concepts; per-boundary rules (policy/control/adapter/agent/evidence/offline/plugin); dependency cycles; exceptions.

## 4. Non-Goals

- No technology, protocol, or deployment topology.
- No automated dependency enforcement/linting (later tooling).
- No application code; no ADR; no Capability Matrix or threat-file change.

## 5. Module versus Component

```text
module ≠ microservice · module ≠ process · module ≠ container ·
module ≠ repository package · module ≠ mandatory deployment unit
```

A module is a logical responsibility boundary. One component may implement several modules; one module may later be realized by several components. Physical separation is a later ADR-governed decision.

## 6. Authority Boundaries

Three separable authority kinds: **policy authority** (MOD-POL-001/MOD-IAM-001), **control/intent authority** (MOD-WFL-001), **execution authority** (MOD-EXE-001). Policy authorises; control expresses intent; execution acts only on approved actions. No module holds all three.

## 7. Data Ownership

Each authoritative concept has exactly one owner module (catalog §8). A `summary`/`derived`/`generated` view never overrides its authoritative source (repository governance source-of-truth rule). Observed state does not silently overwrite desired state; execution results do not rewrite approval history.

## 8. Dependency Direction

```text
Experience → control/policy-facing modules → not directly to privileged execution
Policy/Approval → authorises orchestration/execution intent → does not perform target execution
Workflow Orchestration → requests approved execution → does not bypass Policy
Execution Coordination → uses approved adapter/agent paths → does not mutate policy
Adapters → isolate vendor/protocol specifics → do not own global governance
Observation → supplies observed state → does not create management authority
Evidence/Audit → records relevant actions → does not control those actions
Offline Intake → supplies quarantined/verified inputs → does not directly trigger privileged execution
```

## 9. Allowed Dependencies

Per-module allowed dependencies are listed in the [catalog](COREOPS_MODULE_CATALOG.md) §6. General rule: depend "downward" toward data/evidence and "inward" toward policy for authorisation; never "sideways" to bypass an authority boundary.

## 10. Forbidden Dependencies

```text
Experience directly invokes privileged target execution.
Adapter grants itself write authority.
Agent changes global policy.
Observation data becomes approved desired state automatically.
Notification channel becomes a control channel.
Evidence store becomes an execution trigger.
Offline import bypasses provenance or approval.
Optional cloud integration becomes required for local core operation.
Plugin or extension bypasses module contracts.
```

## 11. Commands, Events and Observations

Conceptual communication only (no protocol selected):

```text
Command ≠ event · Observation ≠ approval · Notification ≠ command ·
Query result ≠ authoritative state · Artifact reference ≠ trusted artifact
```

Patterns: request · command · approval decision · event · observation · state transition · query · evidence record · artifact reference · notification. No REST/GraphQL/gRPC/WebSocket/broker/event-bus/DB-trigger selection.

## 12. Policy-to-Action

An intent becomes an authorised action only after policy/approval (MOD-POL-001). Fail-closed on policy unavailability. Threats THR-003/004; invariant 3.

## 13. Control-to-Execution

Only approved actions cross from MOD-WFL-001 to MOD-EXE-001; `execution-authorised`/`deployment-authorised` required; replay/sequence protection needed. Threats THR-026/028/029; invariant 4.

## 14. Adapter Boundary

MOD-ADP-001 isolates vendor/protocol specifics, read-first; it does not own governance and does not self-grant write authority. Read capability ≠ write capability (invariant 1). Threats THR-007/010/011.

## 15. Agent Boundary

MOD-AGT-001 is optional; agentless operation remains possible. Agent identity/enrollment required; agent compromise is contained, not global (invariant 16). Threats THR-008/009.

## 16. Evidence Boundary

MOD-EVD-001 records audit/evidence and supports controlled export; it holds no execution authority and is not an execution trigger. Attributable records; no secrets in evidence (invariants 13/14). Threats THR-016/017/018.

## 17. Offline Intake Boundary

MOD-OFF-001 supplies quarantined/verified inputs; provenance and integrity verification required; import does not directly trigger privileged execution and does not bypass approval (invariant 12; fail-closed). Threats THR-024/025.

## 18. Plugin and Extension Boundary

MOD-EXT-001 governs extensions with provenance, permission limits, compatibility, and lifecycle; extensions do not bypass module contracts, policy, or authority boundaries. Threat THR-023 (supply-chain analog).

## 19. Dependency Cycles

Authority-crossing dependency cycles are forbidden (e.g. Execution → Policy → Execution). Acyclic authority direction is a design invariant; cycles that hide coupling must be refactored, not accepted silently.

## 20. Exception Process

Any deviation from a boundary/dependency rule that touches security invariants, accepted ADRs, or NDF rules requires a documented exception with rationale, a risk entry, Nova review, and Human-Maintainer approval. No silent exceptions.

## 21. Validation Boundary

No boundary is enforced or validated in this WP; rules are design requirements. Automated enforcement (linting, dependency checks) is later tooling.

## 22. Compatibility

Additive; consistent with the logical architecture, module catalog, system context, plane taxonomy, threat model, repository-governance source-of-truth rule, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 23. Open Questions

- Which dependencies are checked by later automated tooling?
- How are module contracts specified for validation?
- Which exceptions (if any) are anticipated for early prototypes?

## 24. Next Decision

Nova review, then Human-Maintainer commit. Component design, dependency-enforcement tooling, and technology selection remain separate later work packages (ADR-governed).
