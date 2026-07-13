# CoreOps – Threat Scenario Register

> Document Status: Implemented, pending Nova review
> Register Status: Foundation threat register
> Threats Closed: None claimed without evidence
> Automated Validation: Not implemented
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-007` (docs-only / security-baseline)

## 1. Status

Authoritative foundation threat inventory for CoreOps (40 scenarios, THR-001…THR-040). Companion to the [Foundation Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md). No threat is `mitigated`/`closed` without evidence; in the current foundation state none is.

## 2. Purpose

Maintain stable-ID threat scenarios with qualitative, evidence-bounded ratings, mitigation states, ownership, and follow-up, so later security-architecture work can design and validate controls against a fixed reference.

## 3. Scope

Threat scenarios only. Project-wide governance risks live in the [Risk Register](../../project-system/RISK_REGISTER.md); trust boundaries in [TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md](TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md).

## 4. Threat-ID Rules

IDs are `THR-NNN`, stable, unique, and **never reused**. A mitigated scenario is not deleted; its status changes with evidence. New scenarios take the next free ID.

## 5. Rating Model

Likelihood: `low` · `medium` · `high` · `unassessed`. Impact: `low` · `medium` · `high` · `critical` · `unassessed`. Risk Level: `low` · `medium` · `high` · `critical` · `unassessed`. Ratings are qualitative and evidence-bounded — no numeric precision is implied. Absent a deployment to measure, likelihood is often `unassessed`.

## 6. Mitigation States

`not-defined` · `design-required` · `implementation-required` · `evidence-required` · `validation-required` · `partially-addressed`. `mitigated`/`closed` require real evidence (not present in the foundation state).

## 7. Scenario Register

Each entry: Category · Actor · Target asset · Entry point · Planes · Boundaries · Preconditions → Threat event → Potential impact · Existing conceptual safeguards · Missing mitigations · L/I/Risk · Mitigation state · Evidence required · Owner · Follow-up · Status.

**THR-001 — Stolen or compromised administrator identity** · identity-and-authentication · compromised platform administrator · AST-01 · Entry: Experience (UI/CLI/API) · Planes: Experience, Policy · TB-01/TB-11 · Pre: valid admin credential obtained → Event: attacker acts as admin → Impact: broad unauthorised control. Existing: role model concept, HM gates. Missing: authN/session model, anomaly detection. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: authN + session + audit tests · Owner: Nova · Follow-up: CO-WP-009 · Status: open

**THR-002 — Privilege escalation** · authorization-and-privilege · compromised operator account · AST-03 · Entry: Policy/Control · Planes: Policy, Control · TB-11 · Pre: limited access → Event: escalates to privileged role → Impact: unauthorised privileged actions. Existing: RBAC concept, least-privilege intent. Missing: enforced RBAC, SoD. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: RBAC + escalation tests · Owner: Nova · Follow-up: CO-WP-009 · Status: open

**THR-003 — Approval gate bypass** · policy-and-approval-bypass · compromised platform administrator · AST-05 · Entry: Policy · Planes: Policy, Control · TB-04 · Pre: approval required → Event: privileged action proceeds without valid approval → Impact: unauthorised change. Existing: policy-to-action separation (invariant 3), fail-closed intent. Missing: approval engine, enforcement. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: approval-gate tests · Owner: Nova · Follow-up: CO-WP-013 · Status: open

**THR-004 — Reuse of expired or one-time approval** · policy-and-approval-bypass · malicious authenticated user · AST-05 · Entry: Control · Planes: Policy, Control · TB-04/TB-05 · Pre: prior approval exists → Event: expired/one-time approval replayed → Impact: unauthorised repeat action. Existing: time-bound approval concept. Missing: replay/sequence protection. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: replay tests · Owner: Nova · Follow-up: CO-WP-013 · Status: open

**THR-005 — Job manipulation before execution** · execution-and-command-abuse · compromised operator account · AST-08 · Entry: Control→Execution · Planes: Control, Execution · TB-05 · Pre: job queued → Event: job altered before execution → Impact: unintended action on target. Existing: preview/approval concept, control-to-execution separation. Missing: integrity of queued jobs. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: job-integrity tests · Owner: Nova · Follow-up: CO-WP-013/021 · Status: open

**THR-006 — Command or script injection** · execution-and-command-abuse · malicious authenticated user · AST-08 · Entry: Execution · Planes: Execution · TB-02 · Pre: parameterised action → Event: injected command/script executes → Impact: unauthorised target command. Existing: signed-script concept (CAP-AUTOMATION-001). Missing: input handling, signed lifecycle. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: injection tests · Owner: Security Review · Follow-up: CO-WP-013 · Status: open

**THR-007 — Unauthorised write activation of a read-only integration** · adapter-and-integration-compromise · compromised adapter/integration · AST-19 · Entry: Integration · Planes: Integration, Managed Resource · TB-03/TB-02 · Pre: read-only connector → Event: write path invoked → Impact: unauthorised target change. Existing: read≠write invariant (1). Missing: authority enforcement. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: authority tests · Owner: Nova · Follow-up: CO-WP-013 · Status: open

**THR-008 — Compromised agent** · agent-and-relay-compromise · compromised agent/relay · AST-20 · Entry: Edge/Agent · Planes: Edge/Agent, Control · TB-06 · Pre: agent deployed → Event: agent controlled by attacker → Impact: false data / unauthorised local actions. Existing: optional-agent design, least-privilege intent. Missing: agent identity/enrollment, isolation. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: agent-compromise containment tests · Owner: Nova · Follow-up: CO-WP-010 · Status: open

**THR-009 — Forged or spoofed agent** · agent-and-relay-compromise · unauthenticated external attacker · AST-20 · Entry: Edge/Agent · Planes: Edge/Agent, Control · TB-06 · Pre: agent channel exists → Event: rogue agent impersonates a real one → Impact: injected telemetry/actions. Existing: agent-identity requirement (invariant 16). Missing: enrollment/attestation. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: enrollment tests · Owner: Nova · Follow-up: CO-WP-010 · Status: open

**THR-010 — Compromised adapter** · adapter-and-integration-compromise · compromised adapter/integration · AST-12 · Entry: Integration · Planes: Integration · TB-03 · Pre: adapter trusted → Event: adapter manipulated → Impact: false inventory / unauthorised calls. Existing: adapter≠support, read-first. Missing: adapter isolation, verification. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: adapter-integrity tests · Owner: Nova · Follow-up: CO-WP-014 · Status: open

**THR-011 — Malicious external API response** · adapter-and-integration-compromise · malicious external provider · AST-12 · Entry: Integration · Planes: Integration, Observation · TB-03 · Pre: external API used → Event: crafted response → Impact: corrupted state / injection. Existing: provider-review requirement. Missing: response validation, provider trust review. L: unassessed · I: medium · Risk: medium · MitState: design-required · Evidence: response-validation tests · Owner: Nova · Follow-up: CO-WP-014 · Status: open

**THR-012 — Manipulated telemetry** · telemetry-and-observed-state-manipulation · compromised managed resource · AST-14 · Entry: Observation · Planes: Observation, Data/State · TB-07 · Pre: telemetry ingested → Event: falsified telemetry → Impact: wrong state/decisions. Existing: telemetry-to-state boundary, evidence-of-observation semantics. Missing: telemetry integrity, source trust. L: unassessed · I: medium · Risk: medium · MitState: design-required · Evidence: telemetry-integrity tests · Owner: Nova · Follow-up: CO-WP-012/019 · Status: open

**THR-013 — Stale telemetry treated as current** · stale-or-inconsistent-state · environmental/infrastructure failure · AST-07 · Entry: Observation · Planes: Observation, Data/State · TB-07 · Pre: telemetry gap → Event: stale data read as current → Impact: wrong health/decisions. Existing: stale≠current invariant (6), "missing data ≠ healthy". Missing: freshness thresholds. L: medium · I: medium · Risk: medium · MitState: design-required · Evidence: freshness tests · Owner: Nova · Follow-up: CO-WP-012 · Status: open

**THR-014 — Manipulated inventory** · inventory-and-topology-manipulation · compromised operator account · AST-12 · Entry: Data/State · Planes: Data/State, Observation · TB-07 · Pre: inventory writable via import → Event: false inventory entries → Impact: wrong scope/authority. Existing: field provenance concept. Missing: provenance integrity. L: unassessed · I: medium · Risk: medium · MitState: design-required · Evidence: provenance tests · Owner: Nova · Follow-up: CO-WP-011 · Status: open

**THR-015 — Manipulated topology** · inventory-and-topology-manipulation · compromised managed resource · AST-13 · Entry: Observation · Planes: Observation · TB-07 · Pre: topology derived → Event: falsified adjacency → Impact: wrong impact analysis. Existing: confidence/evidence model, manual authority. Missing: evidence integrity. L: unassessed · I: medium · Risk: medium · MitState: design-required · Evidence: topology-evidence tests · Owner: Nova · Follow-up: CO-WP-020 · Status: open

**THR-016 — Audit-log deletion** · audit-and-evidence-tampering · malicious privileged insider · AST-16 · Entry: Evidence/Audit · Planes: Evidence/Audit · TB-08/TB-11 · Pre: audit store accessible → Event: records deleted → Impact: lost accountability. Existing: attributable-records invariant (14). Missing: tamper-evident/immutable audit. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: tamper-evidence tests · Owner: Security Review · Follow-up: CO-WP-018 · Status: open

**THR-017 — Audit-log manipulation** · audit-and-evidence-tampering · malicious privileged insider · AST-16 · Entry: Evidence/Audit · Planes: Evidence/Audit · TB-08 · Pre: audit writable → Event: records altered → Impact: false history. Existing: integrity requirement. Missing: integrity protection. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: integrity tests · Owner: Security Review · Follow-up: CO-WP-018 · Status: open

**THR-018 — Evidence export to wrong recipient** · audit-and-evidence-tampering · accidental operator · AST-17 · Entry: Evidence-export · Planes: Evidence/Audit · TB-08 · Pre: export enabled → Event: evidence sent to wrong destination → Impact: disclosure. Existing: controlled-export requirement, redaction. Missing: export authorisation/controls. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: export-control tests · Owner: Nova · Follow-up: CO-WP-018 · Status: open

**THR-019 — Secret leak in logs** · secret-and-credential-exposure · accidental operator · AST-02 · Entry: Observation/Evidence · Planes: Observation, Evidence · TB-08 · Pre: logging active → Event: secret written to log → Impact: credential exposure. Existing: "no secrets in logs" invariant (13), disclosure policy. Missing: redaction enforcement. L: medium · I: high · Risk: high · MitState: design-required · Evidence: redaction tests · Owner: Security Review · Follow-up: CO-WP-024/025 · Status: open

**THR-020 — Secret leak in exports** · secret-and-credential-exposure · accidental operator · AST-02 · Entry: Evidence-export · Planes: Evidence/Audit · TB-08 · Pre: export/support bundle → Event: secret included → Impact: credential exposure. Existing: support-bundle redaction concept (CAP-PROTECT-007). Missing: enforced redaction. L: medium · I: high · Risk: high · MitState: design-required · Evidence: export-redaction tests · Owner: Security Review · Follow-up: CO-WP-024/025 · Status: open

**THR-021 — Manipulated update artifact** · deployment-and-update-integrity · artifact/update source attacker · AST-10 · Entry: Integration/Execution · Planes: Integration, Execution · TB-03/TB-09 · Pre: update imported → Event: tampered update deployed → Impact: compromised system. Existing: artifact-verification concept (CAP-TRUST-004), invariant 11. Missing: provenance/signature verification. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: artifact-verification tests · Owner: Security Review · Follow-up: CO-WP-022 · Status: open

**THR-022 — Manipulated deployment artifact** · deployment-and-update-integrity · software supply-chain attacker · AST-09 · Entry: Execution · Planes: Execution · TB-05 · Pre: deploy uses artifact → Event: tampered artifact deployed → Impact: compromised deployment. Existing: artifact trust concept, SBOM/provenance. Missing: verification enforcement. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: verification tests · Owner: Security Review · Follow-up: CO-WP-022 · Status: open

**THR-023 — Dependency or supply-chain compromise** · software-supply-chain · software supply-chain attacker · AST-21 · Entry: Integration · Planes: Integration, Data/State · TB-03 · Pre: dependency admitted → Event: compromised dependency introduced → Impact: broad compromise. Existing: dependency admission criteria, provenance/lock direction. Missing: SBOM/verification enforcement. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: SBOM/provenance tests · Owner: Security Review · Follow-up: CO-WP-022 · Status: open

**THR-024 — Offline import without provenance** · offline-import-and-export · artifact/update source attacker · AST-21 · Entry: Online-to-offline · Planes: Integration, Execution · TB-09 · Pre: offline import → Event: unverified artifact imported → Impact: compromised isolated deployment. Existing: CorePack signed-import concept, invariant 12, fail-closed. Missing: verification mechanism. L: unassessed · I: critical · Risk: high · MitState: design-required · Evidence: import-verification tests · Owner: Security Review · Follow-up: CO-WP-023 · Status: open

**THR-025 — Offline export with sensitive data** · offline-import-and-export · accidental operator · AST-24 · Entry: Online-to-offline · Planes: Evidence/Audit, Data/State · TB-09/TB-08 · Pre: offline export → Event: sensitive/personal data exported unredacted → Impact: disclosure. Existing: redaction requirement, disclosure policy. Missing: enforced redaction/minimisation. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: export-redaction tests · Owner: Nova · Follow-up: CO-WP-023/025 · Status: open

**THR-026 — Replay of a previously authorised command** · execution-and-command-abuse · network-position attacker · AST-08 · Entry: Control→Execution · Planes: Control, Execution · TB-05 · Pre: authorised command observed → Event: command replayed → Impact: unauthorised repeat. Existing: control-to-execution separation. Missing: replay/sequence protection. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: replay tests · Owner: Nova · Follow-up: CO-WP-013 · Status: open

**THR-027 — False or manipulated time source** · stale-or-inconsistent-state · network-position attacker · AST-22 · Entry: Observation/Execution · Planes: all · TB-07 · Pre: time relied upon → Event: skewed/false time → Impact: wrong sequencing, expired approvals honored. Existing: time/sequence integrity need (AST-22). Missing: trusted time, local fallback. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: time-integrity tests · Owner: Nova · Follow-up: CO-WP-007-followup · Status: open

**THR-028 — Queue entry falsely treated as executed** · stale-or-inconsistent-state · environmental/infrastructure failure · AST-08 · Entry: Control · Planes: Control, Execution · TB-05 · Pre: job queued → Event: queued read as executed → Impact: false success/skip. Existing: queued≠executed invariant (7). Missing: state model enforcement. L: medium · I: medium · Risk: medium · MitState: design-required · Evidence: state tests · Owner: Nova · Follow-up: CO-WP-012/013 · Status: open

**THR-029 — Executed action falsely treated as successful** · stale-or-inconsistent-state · environmental/infrastructure failure · AST-08 · Entry: Execution · Planes: Execution · TB-05 · Pre: action executed → Event: result unverified but marked success → Impact: undetected failure. Existing: executed≠successful invariant (8), health-check concept. Missing: verification enforcement. L: medium · I: medium · Risk: medium · MitState: design-required · Evidence: verification tests · Owner: Nova · Follow-up: CO-WP-013 · Status: open

**THR-030 — Successful job falsely treated as compliant** · stale-or-inconsistent-state · malicious privileged insider · AST-23 · Entry: Evidence/Audit · Planes: Policy, Evidence · TB-08 · Pre: job succeeded → Event: success equated with compliance → Impact: false compliance impression. Existing: successful≠compliant invariant (9), evidence≠satisfaction. Missing: compliance evaluation separation. L: unassessed · I: high · Risk: medium · MitState: design-required · Evidence: compliance-separation review · Owner: Nova · Follow-up: CO-WP-018 · Status: open

**THR-031 — Partial job failure without a safe state** · recovery-and-rollback-failure · environmental/infrastructure failure · AST-08 · Entry: Execution · Planes: Execution · TB-05 · Pre: multi-step job → Event: partial failure leaves inconsistent state → Impact: unsafe/unknown state. Existing: idempotency requirement, fail-closed. Missing: transactional/rollback model. L: medium · I: high · Risk: high · MitState: design-required · Evidence: partial-failure tests · Owner: Nova · Follow-up: CO-WP-021/026 · Status: open

**THR-032 — Rollback fails** · recovery-and-rollback-failure · environmental/infrastructure failure · AST-06 · Entry: Execution · Planes: Execution, Data/State · TB-05 · Pre: rollback invoked → Event: rollback fails/incomplete → Impact: stuck bad state. Existing: last-known-good concept, invariant 17. Missing: tested rollback. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: rollback tests · Owner: Nova · Follow-up: CO-WP-026 · Status: open

**THR-033 — Backup or restore manipulated** · recovery-and-rollback-failure · malicious privileged insider · AST-18 · Entry: Data/State · Planes: Data/State, Evidence · TB-11 · Pre: backup exists → Event: backup/restore tampered → Impact: unrecoverable/compromised restore. Existing: restore-readiness concept (backup≠restorable). Missing: backup integrity, restore tests. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: restore-integrity tests · Owner: Security Review · Follow-up: CO-WP-026 · Status: open

**THR-034 — Managed resource compromises CoreOps** · execution-and-command-abuse · compromised managed resource · AST-07 · Entry: Managed Resource→CoreOps · Planes: Integration, Observation · TB-02 · Pre: CoreOps reads/acts on target → Event: malicious target exploits CoreOps ingestion → Impact: CoreOps compromise. Existing: read-first, least-privilege, self-protection concept. Missing: input isolation. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: ingestion-isolation tests · Owner: Security Review · Follow-up: CO-WP-026 · Status: open

**THR-035 — Organisation or tenant data boundary violated** · organisational-and-tenant-boundary-failure · malicious authenticated user · AST-24 · Entry: Data/State · Planes: Data/State, Policy · TB-10 · Pre: multi-org/tenant scope → Event: cross-boundary access → Impact: data leakage. Existing: organisational-boundary concept, neutrality policy. Missing: tenancy/scoping model. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: isolation tests · Owner: Nova · Follow-up: CO-WP-009 · Status: open

**THR-036 — Resource exhaustion or denial of service** · availability-and-resource-exhaustion · unauthenticated external attacker · AST-08 · Entry: Experience/Integration · Planes: Experience, Control · TB-01/TB-03 · Pre: request handling → Event: resource exhaustion → Impact: degraded/unavailable operations. Existing: degraded-modes concept, self-dependency protection. Missing: rate-limiting/quota model. L: medium · I: medium · Risk: medium · MitState: design-required · Evidence: load/limit tests · Owner: Nova · Follow-up: CO-WP-026 · Status: open

**THR-037 — Malicious privileged insider** · authorization-and-privilege · malicious privileged insider · AST-03 · Entry: Policy/Execution · Planes: Policy, Execution · TB-11 · Pre: privileged access → Event: abuse of legitimate privilege → Impact: unauthorised change/disclosure. Existing: SoD/break-glass concept, audit. Missing: SoD enforcement, monitoring. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: SoD tests · Owner: Security Review · Follow-up: CO-WP-009/013 · Status: open

**THR-038 — Malicious automation client** · identity-and-authentication · malicious automation client · AST-01 · Entry: API · Planes: Experience, Control · TB-01 · Pre: automation client trusted → Event: client abused/compromised → Impact: unauthorised automated actions. Existing: automation-client role concept. Missing: client authN/scoping. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: client-auth tests · Owner: Nova · Follow-up: CO-WP-009/014 · Status: open

**THR-039 — Network-position attacker (interception/MITM)** · data-disclosure-and-privacy · network-position attacker · AST-19 · Entry: any network path · Planes: Integration, Edge/Agent · TB-03/TB-06 · Pre: network traffic → Event: interception/tampering → Impact: disclosure/integrity loss. Existing: mTLS direction (CAP-TRUST-003), integrity expectations. Missing: transport security selection. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: transport-security tests · Owner: Security Review · Follow-up: CO-WP-010 · Status: open

**THR-040 — Stolen backup or evidence recipient** · data-disclosure-and-privacy · stolen backup/evidence recipient · AST-18 · Entry: Evidence/Backup export · Planes: Evidence/Audit, Data/State · TB-08 · Pre: backup/evidence stored/exported → Event: copy obtained by attacker → Impact: disclosure of sensitive data. Existing: controlled-export, redaction. Missing: encryption-at-rest/in-transit selection, access control. L: unassessed · I: high · Risk: high · MitState: design-required · Evidence: at-rest protection tests · Owner: Security Review · Follow-up: CO-WP-024/026 · Status: open

## 8. Cross-References

Threat model: [COREOPS_FOUNDATION_THREAT_MODEL.md](COREOPS_FOUNDATION_THREAT_MODEL.md) (§7 assets AST-xx, §11 boundaries TB-xx, §13 abuse cases AB-x, §14 invariants). Boundaries: [TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md](TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md).

## 9. Validation Boundary

No scenario is validated; no mitigation is implemented. Ratings are qualitative. `mitigated`/`closed` require documented evidence in a later WP. Absence of evidence keeps a scenario `open`.

## 10. Review Triggers

A major architecture change; a new plane/boundary/external class; a new capability entering implementation; a security-baseline or threat-model-followup WP; discovery of a real vulnerability (handled via the security process, not this foundation register).

## 11. Historical Retention

Scenarios are retained even after mitigation (status change, not deletion). THR IDs are never reused.

## 12. Open Questions

- Which scenarios are highest priority for the first security-architecture WP?
- Which mitigations become the invariants' enforcement mechanisms (later design)?
- When is likelihood re-rated against a real deployment?
