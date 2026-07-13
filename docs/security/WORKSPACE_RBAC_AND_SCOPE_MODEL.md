# CoreOps – Workspace, RBAC and Scope Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation workspace and RBAC model
> Implementation Status: Not implemented
> Tenant Isolation Status: Not designed
> Policy Engine: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-009` (docs-only / security-baseline)

## 1. Status

Foundation workspace, role-based access control (RBAC), and scope model. Technology-independent; no policy engine, tenant-isolation design, or database is selected. Companion to [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md) and [BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md).

## 2. Purpose

Define workspaces, membership, roles, a permission taxonomy, and resource/action scopes so authority is always explicit, scope-bound, deny-by-default, and auditable — never implied by membership or a broad role name.

## 3. Scope

Workspace definition/lifecycle; membership; roles; permission taxonomy; resource/action scopes; role assignment; delegation; cross-workspace access; separation of duties; sensitive operations; audit; offline; invariants; threat references.

## 4. Non-Goals

- No policy engine, tenant-isolation implementation, database, or API spec.
- No full technical permission list; no runtime code; no ADR.
- No module/capability/threat-file change.

## 5. Workspace Definition

A workspace is an administrative and functional grouping boundary. It is **not** automatically: a legal organisation · a security tenant · a network segment · a deployment · a billing account.

Documented per workspace: Workspace ID · Name · Owner · Membership · Assigned roles · Managed-resource scope · Policy profile · Audit scope · Lifecycle status · Parent/relationship model · Cross-workspace sharing boundary. Workspace IDs are stable and not reused.

## 6. Workspace Lifecycle

Conceptual states: `provisioning` · `active` · `suspended` · `archived` · `deleted-logically`. Deletion/archival must not destroy audit/evidence references. Lifecycle changes are audited.

## 7. Membership

Membership is explicit; no implicit access by name similarity or group membership. A member has only the roles explicitly assigned in that workspace. Membership does not confer global platform authority. Membership changes are sensitive and audited.

## 8. Roles

Deny-by-default role families (Foundation guidance, not implemented enums): Platform Owner · Platform Administrator · Workspace Owner · Workspace Administrator · Operator · Deployment Operator · Read-Only Observer · Auditor · Integration Administrator · Emergency Access Operator. A role name does not imply scope; scope is assigned explicitly (§10).

## 9. Permission Taxonomy

Foundation permission taxonomy (not a full technical list). Conceptual permission fields: Permission ID · Action · Resource type · Scope type · Privilege level · Approval requirement · Audit requirement · Delegability · Offline availability · Break-glass eligibility.

Actions: `read` · `discover` · `observe` · `export` · `configure` · `register` · `approve` · `execute` · `deploy` · `update` · `manage-identity` · `manage-membership` · `manage-role` · `manage-integration` · `manage-secret-reference` · `view-audit` · `export-evidence` · `manage-policy` · `manage-emergency-access`.

Ground rules:

```text
read ≠ export · observe ≠ configure · configure ≠ execute ·
approve ≠ execute · execute ≠ deploy · manage-membership ≠ manage-role ·
view-audit ≠ alter-audit · manage-secret-reference ≠ read-secret-value
```

## 10. Resource and Action Scopes

Scope types: `global platform scope` · `workspace scope` · `environment scope` · `resource-group scope` · `individual resource scope` · `workflow scope` · `job scope` · `evidence scope`.

```text
A broader role name does not automatically imply a broader scope.
```

An authorization decision must be derivable from: principal · role/permission bundle · action · resource · scope · current lifecycle state · applicable approval state. No policy engine is selected.

## 11. Role Assignment

Deny-by-default; least privilege; explicit grant; scope-bound authority; auditable assignment; revocable; no implicit privilege inheritance. Role assignment is a sensitive operation (requester ≠ approver where the profile allows).

## 12. Delegation

Per the identity governance §12: explicit, scope-bound, time-bound where appropriate, revocable, auditable, non-transitive by default; never beyond one's own scope; break-glass delegation needs special governance.

## 13. Cross-Workspace Access

Cross-workspace access must be explicit, bounded, and auditable. Switching workspaces does not carry rights from the previous scope. No workspace creates global platform authority. Threat ref THR-035.

## 14. Separation of Duties

Recommended separations (role/deployment/policy requester ≠ approver; break-glass requester ≠ reviewer; offline importer ≠ final execution approver). Small deployments: full separation may be impossible; compensating controls and visible, auditable reduced separation are required; no automatic full-SoD claim.

## 15. Sensitive Operations

Treated as sensitive (possible reauthentication, extra approval, time-limit, SoD, audit/notification, post-review): role assignment · privilege elevation · workspace ownership transfer · break-glass activation · secret-reference changes · integration write-enablement · deployment approval · deployment execution · system-wide policy change · audit export · evidence export · offline artifact approval · identity recovery · account reactivation. No MFA/reauth technology selected.

## 16. Audit Requirements

Membership, role, scope, delegation, cross-workspace, and sensitive-operation changes are audit-relevant and attributable (MOD-EVD-001). `view-audit` never implies `alter-audit`. Evidence capability ≠ requirement satisfaction.

## 17. Offline Considerations

Local RBAC evaluation must be possible offline; no mandatory cloud/IdP for core authorization; offline emergency access is governed (break-glass policy). No classified-network claim.

## 18. Security Invariants

Design requirements (not implemented controls):

```text
Workspace membership must not imply global authority.
Role assignment must be explicit and auditable.
Cross-workspace access must be explicit and auditable.
Privilege elevation must be time- and scope-bounded.
Revoked authority must not remain indefinitely usable.
A broader role name must not imply a broader scope.
manage-secret-reference must not imply read-secret-value.
view-audit must not imply alter-audit.
```

## 19. Threat References

THR-001 (compromised identity), THR-002 (privilege escalation), THR-003/004 (approval bypass/replay), THR-035 (tenant boundary), THR-037 (insider), THR-016/017 (audit). No parallel threat list; no invented IDs.

## 20. Technology Boundary

No policy engine, tenant-isolation implementation, database, token, or session technology is selected; deferred to later ADR-governed work.

## 21. Compatibility

Additive; consistent with MOD-IAM-001/MOD-POL-001, system context, threat model, disclosure policy, and NDF rules. No technology; no ADR. Breaking-change potential: low.

## 22. Open Questions

- Which tenant-isolation model (if any) per profile (later)?
- Which permission bundles are recommended defaults?
- How is scope enforced technically (later policy engine)?

## 23. Next Decision

Nova review, then Human-Maintainer commit. Policy engine, tenant isolation, and permission implementation remain separate later work packages (ADR-governed).
