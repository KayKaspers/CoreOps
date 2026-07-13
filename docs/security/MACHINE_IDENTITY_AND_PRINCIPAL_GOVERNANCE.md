# CoreOps – Machine Identity and Principal Governance

> Document Status: Implemented, pending Nova review
> Governance Status: Foundation machine-identity governance
> Implementation Status: Not implemented
> Credential Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-010` (docs-only / security-baseline)

## 1. Status

Foundation, technology-independent governance for non-human (machine) identities and principals: workloads, devices, agents, adapters, relays, and automation clients. It selects no PKI, certificate, key algorithm, TPM/HSM, mTLS, SSH-key model, enrollment protocol, or secret store, and implements nothing. Companions: [MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md](MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md) and [OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md).

## 2. Purpose

Give CoreOps clear machine-identity concepts and authority boundaries so later identity-architecture work can design enrollment and credentials against a shared model — without conflating identity with credential, discovery with enrollment, or registration with trust.

## 3. Scope

Concepts; human-vs-machine identity; principal classes; identity ownership; lifecycle; scope/authorization; agent/adapter boundary; automation clients; audit/evidence; offline; security invariants; threat references.

## 4. Non-Goals

- No PKI, certificate format, key algorithm, TPM/HSM, mTLS, SSH-key, enrollment-protocol, or secret-store selection.
- No database schema, API spec, runtime code, or hardware-attestation design.
- No certification assessment; no ADR; no human-identity/module/capability/threat-file change.

## 5. Concepts

`machine identity` · `machine account` · `machine principal` · `workload identity` · `device identity` · `agent identity` · `adapter identity` · `automation-client identity` · `credential` · `secret` · `private key material` · `public identity material` · `enrollment request` · `bootstrap material` · `enrollment approval` · `registration` · `trust establishment` · `rotation` · `renewal` · `suspension` · `revocation` · `re-enrollment` · `decommissioning`.

Ground rules:

```text
machine identity ≠ credential · credential ≠ secret-governance metadata ·
registration ≠ trust · discovery ≠ enrollment · enrollment ≠ write authority ·
successful authentication ≠ unrestricted authorization ·
agent identity ≠ global platform authority · adapter identity ≠ permission to bypass policy
```

## 6. Human versus Machine Identity

```text
Human identity ≠ machine identity
A machine identity must not be used as an anonymous replacement for a human operator.
Human approval must remain attributable to a human identity even when execution
is performed by a machine principal.
```

A machine identity may execute actions, read data, supply telemetry, and invoke approved workflows — only within its explicit scope. It must not inherit Human-Maintainer or Platform-Owner authority (consistent with [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md) §7).

## 7. Principal Classes

For each class: purpose · owning module · expected scope · read/write/execution authority · enrollment requirement · offline considerations · audit requirement · revocation owner · known threats. No credential technology is fixed.

| Principal class | Owning module | Default authority | Known threats |
|---|---|---|---|
| managed-resource principal | MOD-INV-001 | read (observed) | THR-034 |
| agent principal | MOD-AGT-001 | scoped collect/execute (authorised) | THR-008/009 |
| relay principal | MOD-AGT-001 | forward (scoped) | THR-009/039 |
| adapter principal | MOD-ADP-001 | read-first; authorised write via Execution | THR-007/010/011 |
| integration principal | MOD-ADP-001 | scoped read/write (authorised) | THR-010/011 |
| automation-client principal | MOD-EXP-001/MOD-WFL-001 | scoped requests (no self-approve) | THR-038 |
| deployment-runner principal | MOD-EXE-001/MOD-DEP-001 | approved deploy/execute | THR-021/022/026 |
| observation collector principal | MOD-OBS-001 | read/telemetry | THR-012/013 |
| offline-transfer principal | MOD-OFF-001 | verified import/export | THR-024/025 |
| evidence-export principal | MOD-EVD-001 | controlled export | THR-018/020 |

## 8. Identity Ownership

Every machine identity has an explicit owner (a human identity or governed role) and a revocation owner. Ownerless machine identities are not permitted. Ownership and revocation authority are auditable (MOD-EVD-001).

## 9. Lifecycle

States (conceptual guidance, not implemented enums): `discovered` · `enrollment-pending` · `approval-pending` · `enrolled` · `active` · `rotation-due` · `renewal-pending` · `suspended` · `revoked` · `compromised` · `re-enrollment-required` · `decommissioning` · `decommissioned` · `archived`.

Ground rules: discovery creates no identity; enrollment creates no automatic write authority; activation requires a defined scope; rotation/renewal must not leave old credentials indefinitely valid; suspension limits/stops use; revocation ends authority; compromise overrides normal lifecycle; re-enrollment requires a new trust decision; decommissioning preserves audit/evidence references. Detailed lifecycle in the enrollment and credential documents.

## 10. Scope and Authorization

A machine identity gains permissions only via: principal identity · principal class · workspace/environment · resource scope · action scope · lifecycle state · credential state · approval state · policy state.

```text
Enrollment ≠ management authority · Authentication ≠ write permission ·
Agent identity ≠ permission to execute arbitrary commands ·
Adapter identity ≠ permission to alter policy · Automation client ≠ permission to self-approve
```

## 11. Agent and Adapter Boundary

Agent principals (MOD-AGT-001) are optional; agentless operation remains possible; agent compromise is contained, not global (THR-008/009; invariant). Adapter principals (MOD-ADP-001) are read-first and never bypass policy or self-grant write (THR-007/010/011). Both remain scope-bound.

## 12. Automation Clients

Automation-client principals make scoped requests through Experience/Orchestration; they cannot self-approve their own actions (approval remains a separate authority, THR-038). Client authentication/scoping mechanisms are deferred.

## 13. Audit and Evidence

Machine-identity creation/enrollment, ownership, approval, scope, activation, credential references, rotation/renewal, suspension/revocation, compromise, re-enrollment, decommissioning, and actions executed under the principal are audit-relevant and attributable to an owner (MOD-EVD-001). Evidence capability ≠ evidence availability ≠ requirement satisfaction.

## 14. Offline Considerations

Machine identities must be operable in isolated environments without mandatory cloud/external services; offline enrollment/credential handling is governed (see companion documents); no claim of suitability for classified networks; concrete mechanisms deferred.

## 15. Security Invariants

Design requirements (not implemented controls):

```text
Discovery must not imply enrollment.
Enrollment must not imply unrestricted authority.
Machine authentication must not imply human approval.
Machine principals must remain scope-bound.
Machine identities must not silently impersonate human identities.
Revoked credentials must not remain indefinitely authoritative.
Compromised identities require explicit containment and review.
Re-enrollment must not silently restore prior authority.
Offline enrollment requires provenance, integrity and approval.
Raw secret material ownership remains an explicit later decision.
Agent or adapter identity must not bypass policy boundaries.
Decommissioned identities must not be reused silently.
```

## 16. Threat References

Relevant existing scenarios: THR-007 (read-to-write), THR-008/009 (compromised/forged agent), THR-010/011 (adapter/API), THR-023 (supply chain), THR-024 (offline import), THR-026 (replay), THR-019/020 (secret exposure), THR-034 (managed resource vs CoreOps), THR-038 (automation client), THR-016/017 (audit). No parallel threat list; no invented IDs.

## 17. Technology Boundary

No PKI, certificate, key algorithm, TPM/HSM, mTLS, SSH-key, enrollment protocol, secret store, or credential technology is selected; deferred to later ADR-governed work.

## 18. Compatibility

Additive; consistent with the module architecture (MOD-AGT-001/MOD-ADP-001/MOD-SEC-001/MOD-EVD-001), threat model, human-identity governance, sovereignty policy, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 19. Open Questions

- Which credential/identity mechanism per principal class (later)?
- How is machine-to-CoreOps ingestion isolated (later)?
- Which attestation (if any) is used for high-trust principals?

## 20. Next Decision

Nova review, then Human-Maintainer commit. Credential technology, enrollment protocol, and attestation remain separate later work packages (ADR-governed).
