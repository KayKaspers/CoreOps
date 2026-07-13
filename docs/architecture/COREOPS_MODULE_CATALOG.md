# CoreOps – Module Catalog

> Document Status: Implemented, pending Nova review
> Catalog Status: Foundation logical module catalog
> Module IDs: Stable and not reusable
> Implementation Status: Not implemented
> Support Status: Not assessed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-008` (docs-only / logical-architecture foundation)

## 1. Status

Foundation register of the seventeen logical modules (companion to [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md)). No module is implemented or supported. This catalog does **not** replace the [Capability Matrix](FOUNDATION_CAPABILITY_MATRIX.md) and creates no second 94-row capability list.

## 2. Purpose

Provide per-module responsibilities, data ownership, dependencies, threat/invariant references, capability-domain references, offline role, and status, so later component work has a fixed logical reference.

## 3. Scope

The seventeen modules and their register fields. Threat scenarios live in the [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md); capabilities in the Capability Matrix.

## 4. Module-ID Rules

IDs are `MOD-<AREA>-<NNN>`, stable, unique, and **never reused**. A retired module changes status; its ID is not reassigned. New modules take the next free ID in their area.

## 5. Classification Model

`foundation-core` · `foundation-cross-cutting` · `optional-platform` · `optional-integration` · `profile-dependent` · `future-extension`. Separate status dimensions: Logical Module Status · Implementation Status · Deployment Requirement · Support Status · Evidence Status. Foundation defaults: Implementation `not-implemented`, Support `not-assessed`, Evidence `not-assessed`, Deployment Requirement `not-selected`.

## 6. Module Register

Each entry: Classification · Purpose · Primary/Secondary planes · Owned concepts (authoritative data) · Accepted inputs · Produced outputs · Commands accepted · Events/observations produced · Allowed dependencies · Forbidden dependencies · Security invariants · Threat refs · Capability-domain refs · Offline role · Deployment requirement · Impl/Support/Evidence status · Open decisions.

**MOD-EXP-001 — Experience and Interface** · foundation-core · Purpose: UI/CLI/API entry, presentation, input-validation boundary; no direct privileged execution. Planes: Experience (primary). Owned: none authoritative (presents others' data). Inputs: user intents/queries. Outputs: requests to control/policy modules. Commands accepted: user requests (validated). Events: user-action events. Allowed deps: MOD-WFL-001, MOD-POL-001, MOD-OBS-001 (read), MOD-EVD-001 (read). Forbidden: direct MOD-EXE-001, MOD-ADP-001 write. Invariants: read≠write, approval-before-execution. Threats: THR-001/002/036/038. Capability domains: Platform. Offline: local UI usable offline. Deploy: not-selected. Impl/Support/Evidence: not-implemented/not-assessed/not-assessed. Open: authN/session model.

**MOD-IAM-001 — Identity and Access Governance** · foundation-core · Purpose: identity/role abstraction, authorization-decision inputs; no concrete authN technology. Planes: Policy/Governance. Owned: identity/role abstraction (authoritative). Inputs: identity assertions, role definitions. Outputs: authorization-relevant attributes. Commands: role/permission changes (authorised). Events: identity/role change events. Allowed deps: MOD-POL-001, MOD-EVD-001. Forbidden: MOD-EXE-001, MOD-ADP-001. Invariants: privilege separation, attributable-audit. Threats: THR-001/002/035/037/038. Capability domains: Identity. Offline: local identity/roles. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: authN method, tenancy.

**MOD-POL-001 — Policy, Approval and Exception** · foundation-core · Purpose: policies, approval gates, exceptions, risk acceptance, policy-to-action boundary. Planes: Policy/Governance. Owned: effective policy, approval state (authoritative). Inputs: policy defs, approval requests, IAM attributes. Outputs: approval decisions, effective policy. Commands: policy/approval/exception changes (authorised). Events: approval/decision events. Allowed deps: MOD-IAM-001, MOD-EVD-001. Forbidden: performing target execution. Invariants: approval-before-execution, control/execution separable, fail-closed. Threats: THR-003/004/030/037. Capability domains: Identity/Policy. Offline: local policy evaluation. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: policy model, approval workflow.

**MOD-INV-001 — Inventory and Resource Registry** · foundation-core · Purpose: registered resources, management status, resource identity; discovered≠managed. Planes: Data/State (primary), Observation. Owned: registered-resource identity and management status (authoritative). Inputs: discovery results, imports, registrations. Outputs: authoritative resource registry. Commands: register/authorise-management (authorised). Events: registration/status-change events. Allowed deps: MOD-STA-001, MOD-EVD-001. Forbidden: MOD-OBS-001 auto-granting management authority. Invariants: discovered≠managed. Threats: THR-014/081-context. Capability domains: Inventory. Offline: local registry. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: provenance model.

**MOD-OBS-001 — Discovery and Observation** · foundation-core · Purpose: discovery, telemetry, observed state; stale≠current. Planes: Observation/Telemetry. Owned: observed state (authoritative for observation only). Inputs: discovery/telemetry from adapters/agents. Outputs: observed state, events. Commands: scan/observe (authorised, read). Events/observations: telemetry, health, reachability. Allowed deps: MOD-ADP-001, MOD-STA-001 (write observed), MOD-EVD-001. Forbidden: creating management authority; overwriting desired state. Invariants: unknown≠healthy, stale≠current, observation≠management. Threats: THR-012/013/034. Capability domains: Discovery, Monitoring. Offline: local observation. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: normalization, freshness thresholds.

**MOD-TOP-001 — Topology Model** · future-extension · Purpose: derived relationships/topologies, topology-data provenance; no implicit control authority. Planes: Observation, Data/State. Owned: topology relationships (with confidence/provenance). Inputs: network/observation reads. Outputs: topology graph (evidence-based). Commands: manual topology authority (authorised). Events: topology-change events. Allowed deps: MOD-OBS-001, MOD-STA-001, MOD-EVD-001. Forbidden: becoming authoritative inventory silently. Invariants: derived≠authoritative. Threats: THR-015. Capability domains: Network, Topology. Offline: last-known topology. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: evidence/confidence model.

**MOD-WFL-001 — Workflow and Job Orchestration** · foundation-core · Purpose: workflow/job lifecycle, execution intent, scheduling; queued≠executed. Planes: Control/Orchestration. Owned: job/workflow state and execution intent (authoritative). Inputs: intents (from Experience), approvals (from Policy). Outputs: approved execution requests. Commands: create/schedule/cancel jobs (authorised). Events: job-state events. Allowed deps: MOD-POL-001, MOD-EXE-001, MOD-STA-001, MOD-EVD-001. Forbidden: bypassing policy; direct target execution. Invariants: queued≠executed, approval-before-execution. Threats: THR-004/005/026/028. Capability domains: Automation, Deployments. Offline: local queue (controlled). Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: idempotency, replay protection.

**MOD-EXE-001 — Execution Coordination** · foundation-core · Purpose: handover of approved actions, execution authority, execution status; executed≠successful. Planes: Execution/Automation. Owned: execution authority and execution status (authoritative). Inputs: approved actions (from Workflow). Outputs: execution results, evidence. Commands: execute approved action (authorised). Events: execution-result events. Allowed deps: MOD-ADP-001, MOD-AGT-001 (optional), MOD-EVD-001. Forbidden: mutating policy; executing unapproved actions. Invariants: control/execution separable, executed≠successful, unverified-artifacts-not-trusted. Threats: THR-005/006/026/029/031. Capability domains: Automation, Deployments. Offline: local execution (controlled). Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: execution abstraction.

**MOD-ADP-001 — Integration and Adapter Framework** · optional-integration · Purpose: vendor/platform adapters, protocol abstraction, capability detection; no self-granted policy bypass. Planes: Integration/Adapter. Owned: adapter capability metadata (not governance). Inputs: external API/protocol data; execution requests (authorised). Outputs: normalized observations; adapter capability info. Commands: adapter read; authorised write (via Execution). Events: adapter events. Allowed deps: MOD-OBS-001, MOD-EXE-001 (as callee), MOD-SEC-001 (credential-use), MOD-EVD-001. Forbidden: owning governance; self-granting write; bypassing Policy. Invariants: read≠write, provider-failure≠authority. Threats: THR-007/010/011/039. Capability domains: all integration-facing. Offline: adapters degrade gracefully. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: adapter isolation, verification.

**MOD-AGT-001 — Edge and Agent Coordination** · profile-dependent · Purpose: optional agents/relays/local collectors, agent identity and trust status; remains optional. Planes: Edge/Agent (optional). Owned: agent identity/trust status. Inputs: agent enrollment, collected data. Outputs: telemetry, authorised local actions. Commands: authorised agent actions. Events: agent status/telemetry. Allowed deps: MOD-EXE-001, MOD-OBS-001, MOD-SEC-001, MOD-EVD-001. Forbidden: altering global policy; becoming mandatory. Invariants: agent-compromise-not-global, optional-agent. Threats: THR-008/009/039. Capability domains: Trust, Discovery. Offline: agent supports offline/edge. Deploy: not-selected (optional). Status: not-implemented/not-assessed/not-assessed. Open: enrollment/attestation.

**MOD-DEP-001 — Deployment and Update Management** · profile-dependent · Purpose: deployment/update intents, artifact references, rollback planning; no unverified artifact execution. Planes: Control, Execution. Owned: deployment/update intent and rollback plan. Inputs: artifact references, approvals, verified intake. Outputs: deployment requests (to Execution). Commands: plan/deploy/rollback (authorised). Events: deployment-state events. Allowed deps: MOD-WFL-001, MOD-EXE-001, MOD-OFF-001, MOD-EVD-001. Forbidden: executing unverified artifacts. Invariants: unverified-artifacts-not-trusted, recovery-needs-evidence. Threats: THR-021/022/031/032. Capability domains: Deployments. Offline: offline deploy via verified intake. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: rollback model, verification.

**MOD-STA-001 — Configuration and State** · foundation-core · Purpose: configuration metadata, desired and observed state, state provenance; no database selection. Planes: Data/State. Owned: configuration and desired state (authoritative); observed state (from Observation). Inputs: config changes, observations. Outputs: state views, derived views (non-authoritative). Commands: config/desired-state changes (authorised). Events: state-transition events. Allowed deps: MOD-EVD-001. Forbidden: observed overwriting desired silently; summary becoming authoritative. Invariants: observed≠desired silent overwrite, stale≠current. Threats: THR-013/028/029. Capability domains: Inventory, Monitoring. Offline: local state. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: persistence technology (deferred).

**MOD-SEC-001 — Credential and Secret Governance Boundary** · foundation-cross-cutting · Purpose: secret references, credential-use limits, access control and audit requirements; no concrete secret-store technology. Planes: Data/State (protected). Owned: secret references and credential-use policy (authoritative). Inputs: credential-use requests. Outputs: scoped credential references (never plaintext secrets). Commands: credential-use authorisation. Events: credential-use audit events. Allowed deps: MOD-POL-001, MOD-EVD-001. Forbidden: exposing secrets to logs/events/evidence. Invariants: no-secret-exposure. Threats: THR-019/020/040. Capability domains: Trust. Offline: local secret governance. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: secret-store technology (deferred).

**MOD-EVD-001 — Evidence and Audit** · foundation-cross-cutting · Purpose: audit events, approval/execution proof, evidence references; evidence capability ≠ requirement satisfaction. Planes: Evidence/Audit. Owned: audit records and evidence references (authoritative). Inputs: audit-relevant events from all modules. Outputs: audit trail, exportable evidence. Commands: controlled evidence export (authorised). Events: audit records. Allowed deps: none downstream (records only). Forbidden: granting execution authority; being an execution trigger. Invariants: attributable-audit, evidence≠satisfaction, no-secret-exposure. Threats: THR-016/017/018/030. Capability domains: Platform (Audit), Protection. Offline: local audit, controlled export. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: integrity/tamper-evidence mechanism.

**MOD-NOT-001 — Notification and External Communication** · optional-platform · Purpose: notifications, external channels, data-minimisation and disclosure limits. Planes: Integration. Owned: notification routing config. Inputs: events to notify. Outputs: notifications (minimised). Commands: send notification (authorised). Events: notification-sent events. Allowed deps: MOD-EVD-001. Forbidden: becoming a control channel. Invariants: notification≠command, no-secret-exposure. Threats: THR-018/025. Capability domains: Platform. Offline: queued/optional. Deploy: not-selected (optional). Status: not-implemented/not-assessed/not-assessed. Open: channel selection (deferred).

**MOD-OFF-001 — Offline Transfer and Artifact Intake** · optional-integration · Purpose: controlled import/export, provenance/integrity requirements, quarantine/approval; signing/crypto deferred. Planes: Integration, Trust. Owned: intake quarantine/verification state. Inputs: offline transfer packages. Outputs: verified/quarantined artifacts. Commands: import/export (authorised, verified). Events: intake/verification events. Allowed deps: MOD-DEP-001, MOD-SEC-001, MOD-EVD-001. Forbidden: directly triggering privileged execution; bypassing provenance/approval. Invariants: offline-provenance, unverified-artifacts-not-trusted, fail-closed. Threats: THR-024/025. Capability domains: Trust. Offline: core offline role. Deploy: not-selected (optional). Status: not-implemented/not-assessed/not-assessed. Open: verification mechanism (deferred).

**MOD-EXT-001 — Extension and Plugin Governance** · future-extension · Purpose: extension registration, provenance, permission limits, compatibility and lifecycle; no concrete plugin technology. Planes: Integration, Policy. Owned: extension registry and permission scope. Inputs: extension registrations/manifests. Outputs: governed extension activation state. Commands: register/enable/disable extension (authorised). Events: extension-lifecycle events. Allowed deps: MOD-POL-001, MOD-SEC-001, MOD-EVD-001. Forbidden: bypassing module contracts. Invariants: plugin-must-not-bypass-contracts. Threats: THR-023 (supply chain analog). Capability domains: Domain-pack/Extension. Offline: local extension governance. Deploy: not-selected. Status: not-implemented/not-assessed/not-assessed. Open: plugin technology (deferred).

## 7. Plane Mapping

Each module maps to ≥1 primary plane (see §6 and architecture §8). Multiple modules may serve one plane (e.g. Observation served by MOD-OBS-001 and MOD-TOP-001). Rejected: one-plane-one-module, one-module-one-deployment, one-module-one-microservice.

## 8. Data and State Ownership

Authoritative ownership (single owner per concept): identities/roles → MOD-IAM-001; effective policy/approval → MOD-POL-001; resource identity/management status → MOD-INV-001; observed state → MOD-OBS-001; topology → MOD-TOP-001; job/execution intent → MOD-WFL-001; execution status → MOD-EXE-001; configuration/desired state → MOD-STA-001; secret references → MOD-SEC-001; audit/evidence → MOD-EVD-001; deployment/update intent → MOD-DEP-001. Ground rules: observed state does not overwrite desired silently; derived topology is not authoritative inventory; summaries/dashboards are not authoritative; execution results do not rewrite approval history; adapters are not authoritative policy sources; audit/evidence does not grant execution authority.

## 9. Threat and Invariant References

Each module references relevant THR-IDs and invariants (§6). No parallel threat register is created; the [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) remains authoritative. Invariants are design requirements, not implemented controls.

## 10. Capability Relationship

Capability-domain references (§6) point to the [Capability Matrix](FOUNDATION_CAPABILITY_MATRIX.md) domains; this catalog does not duplicate the 94-capability list and does not change capability status.

## 11. Optionality

Optional: MOD-AGT-001 (agents), MOD-NOT-001 (notification), MOD-ADP-001/MOD-OFF-001 (integration), external cloud/identity/time/key/evidence systems. The underlying necessary security/governance functions (policy, secret governance, audit) are **not** optional. Offline core is preserved.

## 12. Validation Boundary

No module is implemented, supported, or validated. Status fields are design-state only; upgrades require evidence in a later WP.

## 13. Historical Retention

Modules are retained across changes (status change, not deletion). Module IDs are never reused.

## 14. Open Questions

- Which modules co-locate physically (later architecture WP + ADRs)?
- Which capability rows map to which module (later mapping, without duplicating the matrix)?
- Which modules are first to enter design/implementation?
