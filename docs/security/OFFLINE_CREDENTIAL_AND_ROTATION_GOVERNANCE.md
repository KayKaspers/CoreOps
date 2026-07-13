# CoreOps – Offline Credential and Rotation Governance

> Document Status: Implemented, pending Nova review
> Governance Status: Foundation credential lifecycle
> Implementation Status: Not implemented
> Raw Secret Storage: Not decided
> Cryptographic Mechanism: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-010` (docs-only / security-baseline)

## 1. Status

Foundation, technology-independent governance for machine credential lifecycle, rotation, and offline distribution/reconciliation. Selects no cryptographic mechanism and does not decide whether CoreOps stores raw secret material. Companion to [MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md) and [MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md](MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md).

## 2. Purpose

Define credential concepts, metadata, and lifecycle (issuance, rotation, renewal, revocation, compromise, offline distribution, reconciliation, decommissioning) so credential governance is explicit and auditable — without claiming ownership or storage of raw secret material.

## 3. Scope

Credential concepts/metadata; raw-secret boundary; issuance/activation; rotation; renewal; overlap; expiry; suspension; revocation; compromise; offline distribution; reconciliation; decommissioning; audit; invariants; threat references.

## 4. Non-Goals

- No cryptographic mechanism, key algorithm, certificate format, secret-store, or lifetime selection.
- No raw-secret storage decision; no runtime code; no ADR.
- No human-identity/module/capability/threat-file change.

## 5. Credential Concepts

`credential reference` · `credential metadata` · `credential status` · `credential owner` · `credential holder` · `credential issuer` · `raw secret or private material` · `public identity material`. CoreOps governance may manage credential references and lifecycle status.

## 6. Credential Metadata

Conceptual metadata: reference · status · owner · holder · issuer · scope · principal · lifecycle state · provenance · issuance/activation/expiry markers · rotation lineage (old/new reference) · revocation path · audit reference. No concrete format is selected.

## 7. Raw Secret Boundary

```text
Credential governance ≠ ownership of raw secret material
```

Whether CoreOps stores raw secret or private-key values is a later security-architecture decision (`Raw Secret Storage: Not decided`). Governance here manages **references and status**, not raw values; `manage-secret-reference ≠ read-secret-value` (RBAC model). Consistent with MOD-SEC-001.

## 8. Issuance and Activation

Issuance binds a credential reference to a principal, owner, and scope after approved trust establishment. Activation makes it usable within scope. Issuance/activation are audited; issuance does not expand scope beyond the approved request.

## 9. Rotation

```text
Rotation ≠ automatic trust expansion
New credential ≠ automatic invalidation evidence for the old credential
```

Rotation requirements: owner · scope (unchanged unless re-approved) · activation time · old credential reference · new credential reference · overlap boundary · revocation path · audit · offline reconciliation. No rotation technique or interval is selected.

## 10. Renewal

```text
Renewal ≠ silent scope expansion
```

Renewal extends validity for the same scope; scope change requires a new decision. Renewal is audited and bounded.

## 11. Overlap

Rotation/renewal may allow a bounded overlap window where old and new credentials are both temporarily valid; the overlap must be bounded, documented, and the old credential must be revoked/expired at the end. No indefinite overlap.

## 12. Expiry

Credentials carry an expiry (mechanism/lifetime deferred). Expired credentials are not authoritative. Expiry is audited; approaching expiry drives `rotation-due`/`renewal-pending` states.

## 13. Suspension

Temporary, reversible, reasoned, auditable limitation/stop of a credential's use without destroying it; resuming requires an explicit decision.

## 14. Revocation

Revocation ends a credential's authority; it must not be only a UI/documentation status and requires later technical enforcement and evidence. Revoked credentials are not authoritative (invariant). Offline revocation distribution is a known challenge (§16).

## 15. Compromise

On suspected/confirmed compromise: treat the credential as untrusted; limit dependent authority; review dependent principals/jobs/sessions; create audit/incident references; require an explicit re-enrollment/re-issuance decision (not automatic). No claim of implemented immediate revocation.

## 16. Offline Distribution

Offline/isolated credential distribution: pre-authorised package · provenance/integrity verification · bounded/single-use validity · target-environment binding · import quarantine · approval before activation · audit continuity. Revocation distribution to isolated systems is a recognised challenge requiring later design (reconciliation, revocation lists, or equivalent). No crypto mechanism selected; no classified-network claim.

## 17. Reconciliation

Offline credential/rotation actions must be later-reconcilable: local audit continuity plus synchronisation or equivalent evidence when connectivity resumes. Local availability ≠ trustworthiness; reconciliation confirms consistency and surfaces conflicts (e.g. a locally active but centrally revoked credential).

## 18. Decommissioning

Stop new use; revoke/expire the credential; preserve credential-metadata history (`archived-metadata`); preserve relevant audit/evidence; record final owner and reason. Decommissioned credential references are not reused silently.

## 19. Audit and Evidence

Issuance, activation, rotation lineage, renewal, overlap, suspension, revocation, compromise, offline distribution, reconciliation, and decommissioning are audit-relevant and attributable (MOD-EVD-001). Evidence capability ≠ availability ≠ requirement satisfaction. No secrets in audit/evidence (invariant).

## 20. Security Invariants

Design requirements (not implemented controls):

```text
Credential governance must not imply ownership of raw secret material.
manage-secret-reference must not imply read-secret-value.
Rotation must not expand scope silently.
Renewal must not expand scope silently.
A new credential must not be treated as automatic invalidation evidence for the old one.
Revoked or expired credentials must not remain authoritative.
Offline credential distribution requires provenance, integrity and approval.
Offline credential actions must be later-reconcilable.
Compromised credentials require containment and an explicit re-issuance decision.
Decommissioned credential references must not be reused silently.
```

## 21. Threat References

THR-019/020 (secret exposure), THR-023 (supply chain), THR-024 (offline import), THR-025 (offline export), THR-026 (replay), THR-008/010 (agent/adapter). No parallel threat list; no invented IDs.

## 22. Technology Boundary

No cryptographic mechanism, key algorithm, certificate format, secret store, credential lifetime, or raw-secret-storage decision is made; deferred to later ADR-governed work.

## 23. Compatibility

Additive; consistent with machine-identity governance, enrollment lifecycle, module architecture (MOD-SEC-001/MOD-OFF-001/MOD-EVD-001), threat model, sovereignty policy, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 24. Open Questions

- Whether CoreOps stores raw secret material (later security-architecture decision)?
- Which rotation intervals and overlap windows are recommended defaults?
- How is offline revocation distributed and reconciled reliably?

## 25. Next Decision

Nova review, then Human-Maintainer commit. Cryptographic mechanisms, secret storage, and rotation implementation remain separate later work packages (ADR-governed).
