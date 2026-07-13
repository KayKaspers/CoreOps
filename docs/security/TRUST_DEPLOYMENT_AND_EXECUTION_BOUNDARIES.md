# CoreOps – Trust, Deployment and Execution Boundaries

> Document Status: Implemented, pending Nova review
> Boundary Status: Foundation security boundary model
> Threat Model Status: Not yet performed
> Technology Selection: None
> Certification Status: None claimed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-006` (docs-only / architecture-context foundation)

## 1. Status

Foundation security boundary model: the conceptual trust boundaries between users, CoreOps, managed environments, external providers, control/execution, policy/action, agents, telemetry, evidence, offline transfer, privilege, and organisation. It performs no threat model and selects no cryptography, network architecture, or authentication method.

## 2. Purpose

Make the trust boundaries explicit so later threat modeling (`CO-WP-007`) and security baselines can build on a shared model, and so read-only observation is never conflated with write/execution authority.

## 3. Scope

Trust boundary model and the individual boundaries with, for each: assets crossing · required authorization · integrity expectations · confidentiality expectations · audit expectations · failure behavior · open design decisions.

## 4. Non-Goals

- No threat-model execution (belongs to `CO-WP-007`).
- No cryptography, network architecture, or authentication method selected.
- No certification/compliance claim; no technology; no ADR.

## 5. Trust Boundary Model

A trust boundary is a point where assets cross between parties/components with different trust levels and where authorization, integrity, confidentiality, and audit expectations must be defined. Boundaries below are conceptual; enforcement mechanisms are later decisions.

## 6. User Boundary

- **Assets crossing:** user intents, credentials/tokens (exchange only), requests, evidence reads.
- **Required authorization:** authenticated identity + role (mechanism deferred).
- **Integrity/Confidentiality:** requests integrity-protected; secrets never exposed in the clear.
- **Audit:** privileged and change-relevant user actions are auditable.
- **Failure behavior:** fail-closed on unclear authorization.
- **Open decisions:** authentication method, session model.

## 7. Managed Environment Boundary

- **Assets crossing:** observations (read), authorised changes (write), collected telemetry.
- **Required authorization:** explicit per-target authority (`managed-read-only` … `execution-authorised`).
- **Integrity/Confidentiality:** target data integrity respected; sensitive data minimised/redacted.
- **Audit:** all writes/executions on targets are evidenced.
- **Failure behavior:** unreachable target ≠ healthy; discovered ≠ managed.
- **Open decisions:** transport, agent vs agentless, credential handling for read integrations.

## 8. External Provider Boundary

- **Assets crossing:** imported artifacts/data, exported evidence/notifications, optional identity/observability data.
- **Required authorization:** provider trust review before reliance; provenance for imports.
- **Integrity/Confidentiality:** provenance/integrity for imports; no secrets to unreviewed providers.
- **Audit:** provider interactions logged.
- **Failure behavior:** provider failure must not break local core function; provider is optional.
- **Open decisions:** which providers, provenance mechanism.

## 9. Policy-to-Action Boundary

- **Assets crossing:** an intent becomes an authorised action only after policy evaluation.
- **Required authorization:** policy/approval decision (fail-closed).
- **Integrity/Audit:** decision and its inputs are evidenced.
- **Failure behavior:** on policy-engine unavailability, deny (fail-closed).
- **Open decisions:** policy model, approval workflow.

## 10. Control-to-Execution Boundary

- **Assets crossing:** approved actions handed from orchestration to execution.
- **Required authorization:** only `execution-authorised`/`deployment-authorised` actions cross.
- **Integrity/Audit:** action provenance and result evidenced.
- **Failure behavior:** queued ≠ executed; executed ≠ successful.
- **Open decisions:** execution abstraction, idempotency mechanism.

## 11. Agent Boundary

- **Assets crossing:** agent-collected telemetry, agent-executed approved actions.
- **Required authorization:** agent identity and enrollment (mechanism deferred); optional plane.
- **Integrity/Confidentiality:** agent channel integrity; least privilege.
- **Failure behavior:** agent failure is degraded, not silent success; agentless remains possible.
- **Open decisions:** agent identity/enrollment, offline agent behavior.

## 12. Telemetry Boundary

- **Assets crossing:** observations from targets into CoreOps state.
- **Integrity:** telemetry is evidence of observation, not ground truth; stale ≠ current.
- **Audit:** source and time reference retained.
- **Failure behavior:** missing telemetry ≠ healthy.
- **Open decisions:** normalization schema, freshness thresholds.

## 13. Evidence Boundary

- **Assets crossing:** audit/evidence data exported to an evidence/audit store.
- **Required authorization:** controlled export; redaction of secrets/personal data.
- **Integrity/Confidentiality:** integrity and time reference; no secrets in unsuitable reports.
- **Audit:** export itself is auditable.
- **Failure behavior:** unavailable export is surfaced, not silently skipped.
- **Open decisions:** evidence format, external store trust.

## 14. Online-to-Offline Transfer Boundary

- **Assets crossing:** artifacts/updates imported into an offline/isolated deployment; evidence exported out.
- **Required authorization:** provenance and integrity verification on import; Human-Maintainer approval.
- **Integrity:** signed/verified transfer package (mechanism deferred; see CorePack capability).
- **Failure behavior:** unverified import is rejected (fail-closed); local availability ≠ trustworthiness.
- **Open decisions:** transfer package format, verification mechanism.

## 15. Privilege Boundary

- **Assets crossing:** escalation from non-privileged to privileged actions (e.g. break-glass).
- **Required authorization:** explicit approval + separation of duties; audited.
- **Failure behavior:** fail-closed; privileged use is exceptional and evidenced.
- **Open decisions:** privilege model, break-glass mechanism.

## 16. Organisational Boundary

- **Assets crossing:** data between tenants/organisational units where applicable.
- **Required authorization:** organisational/tenant scoping.
- **Confidentiality:** no cross-organisation leakage; neutrality per disclosure policy.
- **Open decisions:** tenancy model (if any), scoping mechanism.

## 17. Failure Behavior

Ground rules across boundaries (consistent with the system context):

```text
unknown ≠ healthy · stale ≠ current · discovered ≠ managed ·
queued ≠ executed · executed ≠ successful · successful ≠ compliant
```

On any unclear authorization or unavailable safety-relevant component: fail closed.

## 18. Audit Expectations

Boundary crossings that change state, authorise actions, exchange credentials, or export evidence must be auditable with traceable origin, time reference, scope, and responsible role — consistent with the evidence model in the BSI/public-sector baseline (§11) and the capability alignment.

## 19. Open Security Decisions

Authentication/session model; policy and approval model; execution/idempotency abstraction; agent identity/enrollment; cryptography and key custody; transfer-package verification; tenancy model. These are inputs to the threat model (`CO-WP-007`) and later security baselines; none is decided here.

## 20. Compatibility

Additive; consistent with the system context, plane taxonomy, sovereignty policy, BSI/public-sector baseline, and NDF rules. Selects no technology; performs no threat model; creates no ADR; claims no certification. Breaking-change potential: low.

## 21. Next Decision

Nova review, then Human-Maintainer commit. The detailed threat model and trust-boundary enforcement are `CO-WP-007` and later security work; technology selection remains deferred.
