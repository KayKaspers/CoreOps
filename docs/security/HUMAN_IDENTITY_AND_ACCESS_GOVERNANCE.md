# CoreOps – Human Identity and Access Governance

> Document Status: Implemented, pending Nova review
> Governance Status: Foundation human-identity governance
> Implementation Status: Not implemented
> Authentication Technology: Not selected
> Identity Provider: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-009` (docs-only / security-baseline)

## 1. Status

Foundation, technology-independent governance for human identities, accounts, principals, project-vs-runtime roles, account lifecycle, sessions, recovery, delegation, and separation of duties. It selects no authentication protocol, identity provider, session, or MFA technology and implements nothing. Companion documents: [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md) and [BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md).

## 2. Purpose

Provide clear identity/account concepts and authority boundaries so later identity-architecture work can design authentication and authorization against a shared, honest model — without conflating authentication with authorization or repository authority with runtime authority.

## 3. Scope

Identity concepts; human-identity/account boundary; project-vs-runtime roles; account lifecycle; authentication-vs-authorization; sessions/reauthentication; recovery; delegation; separation of duties; privacy/minimisation; audit/evidence; offline; security invariants; threat references.

## 4. Non-Goals

- No identity-provider, OIDC/OAuth/SAML/LDAP, password-hashing, MFA, token-format, or session-store selection.
- No database schema, API spec, UI, or runtime code.
- No concrete tenant-isolation design; no certification assessment; no ADR.

## 5. Identity Concepts

`person` · `human identity` · `account` · `principal` · `credential` · `authentication event` · `session` · `workspace` · `organisation` · `membership` · `role` · `permission` · `permission bundle` · `resource scope` · `action scope` · `delegation` · `approval` · `emergency access` · `break-glass event`.

Ground rules:

```text
person ≠ account · account ≠ role · role ≠ permission ·
membership ≠ administrative authority · authentication ≠ authorization ·
workspace membership ≠ access to every managed resource ·
successful login ≠ permission for privileged execution
```

## 6. Human Identity and Account Boundary

A human identity refers to a real person; an account is a managed principal that a person authenticates as. One person may hold several accounts (e.g. normal + emergency), each attributable. Accounts carry no authority by existence — authority derives from explicit role/scope assignment (see RBAC model).

## 7. Project versus Runtime Roles

**Repository authority ≠ runtime authority.** The **Human Maintainer** is a project/repository-governance role (commit/push/tag/release/final decisions) and does **not** automatically become a privileged runtime or deployment role. Distinct roles: Human Maintainer · Platform Owner · Platform Administrator · Operator · Auditor/Reviewer · Managed-System Administrator · Workspace Member · Emergency Access Operator.

```text
Repository authority ≠ runtime authority
Platform ownership ≠ unrestricted operational access
Managed-system administration ≠ CoreOps administration
```

## 8. Account Lifecycle

States (conceptual guidance, not implemented enums): `invited` · `pending-verification` · `active` · `suspended` · `locked` · `recovery-pending` · `deactivated` · `departed` · `archived`.

Requirements: unique account identity; no silent account reuse; joiner/mover/leaver process; role review on task change; timely deactivation; audit evidence; no deletion of relevant audit history; recovery without bypassing identity verification.

## 9. Authentication versus Authorization

Authentication establishes *who*; authorization establishes *what is allowed*. A successful authentication grants no privileged authority by itself. Authentication technology (protocol, MFA, credential handling) is deferred; the model only requires that authentication is attributable and that authorization is evaluated separately (deny-by-default).

## 10. Sessions and Reauthentication

- A session is not an account; session authority is bound to the current account and membership status.
- Revoked authority must not remain indefinitely usable in existing sessions.
- Sensitive actions may require reauthentication.
- Session expiry, revocation, and maximum duration are later technical decisions.
- Offline operation must not force indefinitely valid privileged sessions.
- No cookie/token/session technology is selected.

## 11. Identity Recovery

Recovery re-establishes access for a legitimate identity without bypassing identity verification; it is not an account-reuse shortcut and not a substitute for break-glass. Recovery is a sensitive operation (reauthentication/approval may apply) and is audited. Mechanism deferred.

## 12. Delegation

Delegation is `explicit` · `scope-bound` · `time-bound where appropriate` · `revocable` · `auditable` · `non-transitive by default`. Not permitted: unlimited silent delegation; delegation beyond one's own scope; automatic re-delegation; delegation without owner/expiry check; delegation of break-glass authority without special governance.

## 13. Separation of Duties

Recommended separations (requester ≠ approver for roles, deployments, policy; break-glass requester ≠ reviewer; audit subject ≠ sole reviewer; offline importer ≠ final execution approver). Single-person operation and small homelabs remain possible: full personnel separation is not possible in every profile; compensating controls may be required; reduced separation must be visible and auditable; no automatic claim of full separation of duties.

## 14. Privacy and Data Minimization

Identity, profile, and audit data are minimised (only what is needed), consistent with the [Public Neutrality and Disclosure Policy](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md) and BSI baseline PSR-15. No unnecessary personal data; audit records must be attributable but not carry excessive personal information; retention/disclosure are operator decisions with minimisation as the default.

## 15. Audit and Evidence

Authentication events, role/membership changes, sensitive operations, delegations, recoveries, and break-glass activations are audit-relevant and must be attributable (identity, time, scope, action). Evidence capability ≠ evidence availability ≠ requirement satisfaction. Audit records are owned by MOD-EVD-001 and are not silently deleted.

## 16. Offline Considerations

Local identity and role operation must be possible without mandatory cloud or external IdP; offline emergency access is governed (see break-glass policy); unavoidable audit interruptions must be later-synchronisable or otherwise evidenced. No claim of suitability for classified networks.

## 17. Security Invariants

Design requirements (not implemented controls):

```text
Authentication must not imply unrestricted authorization.
Workspace membership must not imply global authority.
Role assignment must be explicit and auditable.
Privilege elevation must be time- and scope-bounded.
Revoked authority must not remain indefinitely usable.
Approval and execution authority remain separate where required.
Repository Human-Maintainer authority is not runtime platform authority.
A managed-system administrator is not automatically a CoreOps administrator.
Secret-governance permission must not imply access to raw secret values.
```

## 18. Threat References

Relevant existing scenarios (from the [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md)): THR-001 (stolen/compromised admin identity), THR-002 (privilege escalation), THR-003 (approval bypass), THR-004 (approval replay), THR-035 (tenant/organisation boundary), THR-037 (malicious privileged insider), THR-038 (malicious automation client), THR-016/017 (audit tampering). No parallel threat list; no invented THR IDs.

## 19. Technology Boundary

No authentication protocol, identity provider, session technology, MFA mechanism, credential storage, or policy engine is selected. These are later ADR-governed decisions.

## 20. Compatibility

Additive; consistent with the module architecture (MOD-IAM-001/MOD-POL-001), system context, threat model, disclosure policy, and NDF rules. No technology; no ADR; no module/capability/threat-file change. Breaking-change potential: low.

## 21. Open Questions

- Which authentication/MFA model per profile (later)?
- How is reauthentication for sensitive operations enforced (later)?
- Which minimal identity attributes are stored?

## 22. Next Decision

Nova review, then Human-Maintainer commit. Identity architecture, authentication technology, and session mechanisms remain separate later work packages (ADR-governed).
