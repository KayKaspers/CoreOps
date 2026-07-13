# CoreOps – Break-Glass and Emergency Access Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation break-glass policy
> Implementation Status: Not implemented
> Emergency Credential Mechanism: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-009` (docs-only / security-baseline)

## 1. Status

Foundation policy for controlled emergency (break-glass) access. Technology-independent; no emergency credential, token, key, or hardware mechanism is selected and nothing is implemented. Companion to [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md) and [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md).

## 2. Purpose

Define break-glass access as an exceptional, temporary, named, reason-bound, scope-bound, audited, and reviewed path — so emergencies can be handled without creating permanent privilege, shared accounts, or unaudited shortcuts.

## 3. Scope

Emergency-access definition; eligibility; activation preconditions; approval; scope/duration; permission elevation; session boundary; audit/evidence; notification; expiry/revocation; post-event review; offline; abuse prevention; invariants; threat references; exceptions.

## 4. Non-Goals

- No emergency credential/token/key/hardware mechanism selection.
- No runtime implementation; no certification assessment; no ADR.
- No module/capability/threat-file change.

## 5. Emergency Access Definition

Break glass is: `exceptional` · `temporary` · `explicitly activated` · `reason-bound` · `scope-bound` · `audited` · `reviewed afterwards` · `revocable`.

Break glass is **not**: a permanent administrator role · a shared everyday account · a way to bypass missing normal permissions · a substitute for account recovery · an unaudited offline shortcut.

## 6. Eligibility

Only a named human identity (an Emergency Access Operator account attributable to a real person) may activate break glass. Eligibility is governed in advance; it is not open to any authenticated user by default.

## 7. Activation Preconditions

Required at activation: named human identity · reason · incident/emergency reference · requested scope · requested actions · maximum duration · activation authority · elevated permissions · audit start record · notification target · post-event review owner. Depending on profile, additionally: second approval · prior strong identity confirmation · local approval · physical/organisational control.

## 8. Approval

Activation authority is explicit. Where the profile allows, the break-glass requester ≠ the activation approver and ≠ the post-event reviewer (separation of duties). Small deployments may reduce separation with visible, auditable compensating controls; no automatic full-SoD claim.

## 9. Scope and Duration

Break glass is scope-bound (specific resources/actions) and time-bound (maximum duration). Scope expansion requires a new decision. No unlimited activation; a mandatory expiry limit is a design requirement; no silent extension.

## 10. Permission Elevation

Elevation grants only the requested emergency permissions for the requested scope/duration; normal role assignments remain unchanged. Elevation does not alter standing RBAC; it is layered and removed on expiry/revocation.

## 11. Session Boundary

A break-glass session is bound to the activation record and its expiry; revoked/expired break glass must not remain usable. Sensitive actions within break glass remain attributable and audited. Session technology is deferred.

## 12. Audit and Evidence

Required evidence: activation identity · activation time · reason · scope · permissions granted · approvals · actions performed · resources affected · expiry/revocation · notifications · post-event reviewer · review outcome · open exceptions. Evidence capability ≠ evidence availability ≠ requirement satisfaction. Records owned by MOD-EVD-001; not silently deleted.

## 13. Notification

Activation notifies a defined target (e.g. Platform Owner/security role). Notification is a record, not a control channel (module boundary: notification ≠ command).

## 14. Expiry and Revocation

Break glass expires automatically at maximum duration or is explicitly revoked. On expiry/revocation, elevated permissions are removed. Continued need requires a new activation decision, not extension.

## 15. Post-Event Review

Post-event review is **mandatory**: a named reviewer (≠ requester where possible) assesses reason, scope, actions, and outcome, and records review outcome and any open exceptions. Unreviewed break glass is an open governance item.

## 16. Offline and Isolated Environments

An isolated deployment may need a local emergency-access path; permanent cloud/external IdP availability must not be mandatory; local emergency authority is governed in advance; offline access is neither unrestricted nor anonymous; unavoidable audit interruptions must be later-synchronisable or otherwise evidenced. No claim of suitability for classified networks. Concrete credential/token/key/hardware mechanisms are deferred.

## 17. Abuse Prevention

Break glass must not become a routine workaround for missing permissions, a shared account, a permanent role, or an unaudited shortcut. Frequency/pattern of activations is itself review-relevant.

## 18. Security Invariants

Design requirements (not implemented controls):

```text
Break-glass access must be attributable to a named human identity.
Break-glass access must not silently alter normal role assignments.
Break-glass activation must create audit evidence.
Emergency access must expire or be explicitly revoked.
Privilege elevation must be time- and scope-bounded.
Offline emergency access must remain governed, attributable and non-anonymous.
Post-event review of break-glass use is mandatory.
```

## 19. Threat References

Relevant existing scenarios (from the [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md)): THR-001 (compromised identity), THR-002 (escalation), THR-003 (approval bypass), THR-004 (approval replay), THR-037 (insider), THR-016/017 (audit tampering), THR-024 (offline import). No parallel threat list; no invented THR IDs.

## 20. Exceptions

Any deviation (e.g. reduced separation in a single-operator deployment) requires a documented exception with rationale, a risk entry, Nova review, and Human-Maintainer approval, plus visible, auditable compensating controls. No silent exceptions.

## 21. Technology Boundary

No emergency credential, token, key, hardware, MFA, or session mechanism is selected; deferred to later ADR-governed work.

## 22. Compatibility

Additive; consistent with the identity governance and RBAC model, MOD-IAM-001/MOD-POL-001/MOD-EVD-001, threat model, and NDF rules. No technology; no ADR; claims no certification. Breaking-change potential: low.

## 23. Open Questions

- Which activation-approval model per profile (single-operator vs. multi-person)?
- Which maximum durations are recommended defaults?
- How is offline break-glass evidence synchronised later?

## 24. Next Decision

Nova review, then Human-Maintainer commit. Emergency credential mechanisms, activation enforcement, and session handling remain separate later work packages (ADR-governed).
