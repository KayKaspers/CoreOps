# CoreOps – Machine Enrollment and Trust Lifecycle

> Document Status: Implemented, pending Nova review
> Lifecycle Status: Foundation enrollment lifecycle
> Implementation Status: Not implemented
> Enrollment Protocol: Not selected
> Trust Mechanism: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-010` (docs-only / security-baseline)

## 1. Status

Foundation, technology-independent enrollment and trust lifecycle for machine identities. Selects no enrollment protocol or trust mechanism and implements nothing. Companion to [MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md) and [OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md).

## 2. Purpose

Define how machine identities are enrolled, trusted, activated, suspended, revoked, handled on compromise, re-enrolled, and decommissioned — so trust is always explicit, owner-bound, scope-bound, and auditable, and never established by network position alone.

## 3. Scope

Enrollment concepts/preconditions/request/approval; bootstrap boundary; trust establishment; activation; suspension; revocation; compromise; re-enrollment; decommissioning; offline enrollment; audit; invariants; threat references.

## 4. Non-Goals

- No enrollment protocol, trust anchor, certificate, or cryptographic mechanism selection.
- No runtime code; no certification assessment; no ADR.
- No human-identity/module/capability/threat-file change.

## 5. Enrollment Concepts

Discovery, enrollment request, approval, registration, trust establishment, activation are distinct steps. `discovery ≠ enrollment`; `registration ≠ trust`; `enrollment ≠ write authority`.

## 6. Enrollment Preconditions

An owner exists; the target identity class and requested scope are defined; a bootstrap source and its provenance/integrity status are available; an approver is identified; the environment/workspace is bound. Enrollment is not initiated by network reachability alone.

## 7. Enrollment Request

Contains at least: requesting principal/source · target machine/workload identity · identity class · requested scope · requested capabilities · owner · approver · bootstrap source · provenance status · integrity status · environment/workspace · enrollment-request expiry · approval decision · issued identity reference · audit record.

Enrollment remains `explicit` · `scope-bound` · `owner-bound` · `auditable` · `revocable` · `non-transitive`. Not permitted: enrollment by network reachability; automatic global authority; silent self-enrollment of privileged agents; enrollment without owner; enrollment from unverified offline packages; reuse of an old identity after decommissioning without review.

## 8. Approval

Enrollment approval is explicit and attributable to an authorising identity (human or governed authority). Where the profile allows, requester ≠ approver (separation of duties). Approval is audited; approval of privileged scope is a sensitive operation.

## 9. Bootstrap Boundary

```text
Bootstrap material ≠ long-term identity
Bootstrap success ≠ permanent trust
Bootstrap possession ≠ unrestricted authorization
```

Bootstrap material must be time- or single-use-bounded where later design allows, scope-bound, controlled at issuance, auditable, and must not remain effective indefinitely after use/expiry. No one-time-token/certificate/key technology is selected.

## 10. Trust Establishment

Requires at least: identity binding · owner binding · scope binding · environment/workspace binding · provenance assessment · integrity assessment · approval · activation status · audit evidence · revocation path.

```text
Registration alone must not establish privileged trust.
```

Trust must not arise from IP address, hostname, MAC address, device label, network position, an agent's self-assertion, or an imported configuration file without provenance.

## 11. Activation

Activation requires a defined scope and an approved, provenance/integrity-checked trust binding. An activated identity is limited to its scope; activation does not expand authority beyond the approved request.

## 12. Suspension

Suspension is temporary, reversible, reasoned, and auditable. It limits or stops use without ending the identity; resuming requires an explicit decision.

## 13. Revocation

Revocation ends authority. It must not exist only as a UI/documentation status; it requires later technical enforcement and evidence. Revoked credentials must not remain authoritative (invariant). Revocation is audited.

## 14. Compromise

On suspected or confirmed compromise: limit authority; treat the credential as untrusted; review dependent principals and scopes; review relevant jobs and sessions; create audit/incident references; do **not** automatically re-enroll. No claim of implemented immediate revocation.

## 15. Re-Enrollment

Required on: compromise; lost trust basis; identity conflict; decommissioning + later resumption; untrustworthy credential lifecycle; material scope change; owner change where governance requires. Re-enrollment must not silently reuse an old identity, must not automatically restore prior authority, must not delete audit history, and must not overwrite a compromise marking.

## 16. Decommissioning

At least: stop new authority; revoke/expire active credentials; disable future enrollment reuse; preserve identity history; preserve relevant audit/evidence; remove/archive active registrations; review dependent integrations; review pending jobs; record final owner and reason. Decommissioned identity IDs must not be reused silently.

## 17. Offline Enrollment

Offline/isolated enrollment documents at least: pre-authorised enrollment package · controlled physical/administrative handoff · provenance verification requirement · integrity verification requirement · expiry/bounded validity · target-environment binding · single-use/bounded-use requirement · import quarantine · approval before activation · audit continuity · later reconciliation · revocation-distribution challenge.

```text
Offline enrollment packages require verifiable provenance, integrity and explicit approval.
The concrete signing, cryptographic and trust-anchor mechanisms remain deferred.
```

Not claimed: offline enrollment implemented; all isolated environments supported; suitability for classified networks; a concrete signing/crypto architecture.

## 18. Audit and Evidence

Enrollment requests, approvals, bootstrap source/provenance, activation, credential references, suspension, revocation, compromise, re-enrollment, and decommissioning are audit-relevant and attributable to an owner (MOD-EVD-001). Evidence capability ≠ availability ≠ requirement satisfaction.

## 19. Security Invariants

Design requirements (not implemented controls):

```text
Discovery must not imply enrollment.
Registration alone must not establish privileged trust.
Enrollment must not imply unrestricted authority.
Bootstrap material must not become a permanent identity.
Revoked credentials must not remain indefinitely authoritative.
Compromised identities require explicit containment; re-enrollment is not automatic.
Re-enrollment must not silently restore prior authority.
Offline enrollment requires provenance, integrity and approval.
Decommissioned identities must not be reused silently.
```

## 20. Threat References

THR-008/009 (compromised/forged agent), THR-010/011 (adapter/API), THR-023 (supply chain), THR-024 (offline import), THR-026 (replay), THR-034 (managed resource), THR-038 (automation client). No parallel threat list; no invented IDs.

## 21. Technology Boundary

No enrollment protocol, trust anchor, certificate, key, or cryptographic mechanism is selected; deferred to later ADR-governed work.

## 22. Compatibility

Additive; consistent with machine-identity governance, module architecture (MOD-AGT-001/MOD-ADP-001/MOD-OFF-001/MOD-EVD-001), threat model, sovereignty policy, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 23. Open Questions

- Which enrollment/attestation model per environment (later)?
- How is offline revocation distributed to isolated systems (later)?
- Which bootstrap validity bounds are recommended defaults?

## 24. Next Decision

Nova review, then Human-Maintainer commit. Enrollment protocol, trust mechanism, and attestation remain separate later work packages (ADR-governed).
