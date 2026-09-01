# CoreOps – Capability Security and Governance Alignment

> Document Status: Implemented, pending Nova review
> Alignment Status: Foundation alignment
> Implementation Status: Documentation only
> Detailed BSI Mapping: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004E` (docs-only / capability-governance alignment)

## 1. Status

Foundation alignment that connects the CoreOps Foundation Capability Matrix to the multi-dimensional status model, the BSI/public-sector readiness domains (PSR-01…PSR-18), product/operator/shared responsibility, the evidence three-state model, and the internal readiness and governance profiles. It implements no capability and performs no clause-level BSI control mapping.

## 2. Purpose

Make it separately visible, per capability, whether something is on the roadmap, designed, implemented, supported, evidenced, and how it relates to security/governance readiness — so that roadmap inclusion is never mistaken for implementation, implementation for support, support for validated evidence, evidence capability for available evidence, evidence for requirement satisfaction, profile relevance for certification, or PSR mapping for BSI compliance.

## 3. Scope

- A multi-dimensional capability status model (five status dimensions).
- Profile relevance (three readiness + three governance profiles).
- PSR-domain mapping (readiness relationship only).
- Product/operator/shared responsibility codes.
- The evidence three-state model applied to the matrix.
- Mapping rules and pending/unknown-state handling.
- A supplementary per-capability Security/Governance mapping table lives in the Foundation Capability Matrix (same IDs, additive).

## 4. Non-Goals

- No capability implementation.
- No clause-level BSI control mapping (see §16).
- No certification, compliance, or government-approval claim.
- No technology or architecture selection; no ADR.
- No capability deletion; no parallel capability list with different IDs.
- No upgrade of implementation/support/evidence status without evidence.

## 5. Source Documents

- [FOUNDATION_CAPABILITY_MATRIX.md](../architecture/FOUNDATION_CAPABILITY_MATRIX.md) (authoritative capabilities and IDs).
- [CAPABILITY_MATRIX_SPEC.md](../project-system/CAPABILITY_MATRIX_SPEC.md) (specification, extended in this WP).
- [BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md) (PSR domains, responsibility, evidence).
- [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](BSI_REFERENCE_AND_CLAIMS_REGISTER.md), [BSI_ALIGNMENT_POSITIONING.md](BSI_ALIGNMENT_POSITIONING.md).
- [PUBLIC_SECTOR_READINESS_PROFILE.md](../governance/PUBLIC_SECTOR_READINESS_PROFILE.md), [ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](../governance/ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md), [COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md](../governance/COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md).

## 6. Capability Identity

- Capabilities use stable IDs of the form `CAP-<DOMAIN>-<NNN>` (e.g. `CAP-IDENTITY-003`).
- IDs and names are preserved; this WP adds dimensions additively and deletes nothing.
- Identity verification (this WP): **94** unique capability rows, **94** unique IDs, no duplicates (grep-verified). See the count-reconciliation note in the Foundation Capability Matrix summary.

## 7. Status Dimensions

Five independent dimensions (never combined into one overloaded field):

- **Roadmap Status:** `not-planned` · `observe` · `planned` · `target-foundation` · `target-later` · `deferred` · `out-of-scope`. Existing project-valid roadmap values (e.g. `target-observe`, `target-map`, `target-plan`, `target-deploy`, `target-automate`, `target-extend`) remain valid and are treated as the authoritative roadmap dimension.
- **Implementation Status:** `not-designed` · `design-required` · `designed` · `implementation-required` · `in-progress` · `implemented` · `partially-implemented` · `not-applicable`. `implemented` only with technical evidence. Current matrix: all `not-implemented` (maps to `implementation-required`/`not-designed`).
- **Support Status:** `not-assessed` · `not-supported` · `experimental` · `planned-support` · `supported` · `deprecated` · `not-applicable`. `supported` requires the defined support-evidence set. Current matrix: all `not-supported`.
- **Evidence Status:** `not-assessed` · `evidence-model-required` · `evidence-capability-designed` · `evidence-capability-implemented` · `evidence-available` · `validated` · `not-applicable`. `validated` only with documented assessment. Current alignment: all `not-assessed` (no per-capability evidence assessed in this WP).
- **Security/Governance Status:** `not-assessed` · `baseline-mapped` · `design-required` · `controls-required` · `evidence-required` · `review-required` · `exception-required` · `not-applicable`. No value `compliant` exists.

## 8. Profile Dimensions

Readiness profiles (security posture): `Standard` (Std) · `Hardened` (Hard) · `Government` (Gov).
Governance profiles (optional, operations/project): `Service Operations` (SO) · `Major Deployment Project` (MDP) · `Public-Sector Delivery` (PSD).

Profiles may be combined. Profile relevance does **not** mean certification:

```text
Government profile        ≠ government approval
Public-Sector profile     ≠ authority approval
Hardened profile          ≠ certification
Service Operations profile≠ ITIL certification
Major Deployment profile  ≠ PRINCE2 certification
```

## 9. PSR Mapping

- Only the existing readiness domains PSR-01…PSR-18 are used; no new PSR domain is invented.
- A capability may map to several PSR domains, to none (if not readiness-relevant), or carry `mapping-review-required` when the mapping is still uncertain.
- Domain-level mapping guidance (applied per capability in the Foundation Capability Matrix supplementary table):

| CoreOps domain | Primary PSR domains |
|---|---|
| Platform / Experience | PSR-01, PSR-16 (Audit Explorer → PSR-07/16) |
| Identity and Governance | PSR-03, PSR-04, PSR-01 (Machine Identity → PSR-03/12) |
| Inventory / Source of Truth | PSR-02, PSR-16 (Field Provenance → PSR-02/14) |
| Discovery | PSR-02, PSR-11, PSR-05 (Controlled Ranges → PSR-11/13) |
| Monitoring / Observability | PSR-07, PSR-08, PSR-16 |
| Network | PSR-11, PSR-02 (Write Actions → PSR-13) |
| Topology | PSR-02, PSR-11 |
| Print | PSR-02, PSR-15 (Queue → PSR-15/07) |
| Virtualization / Containers | PSR-02, PSR-13, PSR-14 |
| Deployments and Change | PSR-13, PSR-06, PSR-09, PSR-14 |
| Automation and Response | PSR-08, PSR-13, PSR-01 |
| Trust and Offline | PSR-12, PSR-14, PSR-10, PSR-17, PSR-18 |
| Protection and Recovery | PSR-09, PSR-10, PSR-15, PSR-16 |

PSR mapping means **relevant to the readiness domain** — not BSI requirement satisfied, control implemented, or deployment compliant.

## 10. Responsibility Model

Compact codes: `P` product · `O` operator · `S` shared · combinations `P/O`, `P/S`, `O/S`, `P/O/S`.

- `P` — CoreOps provides the capability (defaults, model, functions).
- `O` — operator decides/operates (assignment, retention, processes, legal applicability).
- `S` — shared (configuration, testing, updates, evidence collection) requiring explicit assignment in a deployment.

A single `shared` value alone is insufficient unless product and operator duties are otherwise explained; codes reference the responsibility model in the BSI baseline §10.

## 11. Evidence Model

The three states remain separate at all times:

```text
Evidence capability:   CoreOps can generate or export evidence (may be planned/implemented as a product feature).
Evidence available:    a concrete deployment actually holds the evidence (deployment-specific).
Requirement satisfied: a responsible body assessed and accepted the evidence (never asserted by the product matrix).
```

The Foundation Capability Matrix therefore uses no blanket values such as `compliant`, `BSI-compliant`, `requirement-satisfied`, or `government-approved`.

## 12. Claim Boundaries

Allowed: capabilities are mapped to readiness domains and profiles for planning; CoreOps can provide evidence-supporting capabilities where implemented, configured and validated. Prohibited: capability certification/compliance status, BSI compliance, government approval, or any claim that PSR mapping equals control satisfaction. Follows the BSI reference/claims register §9–§11.

## 13. Mapping Rules

1. Preserve capability IDs and names; add dimensions additively.
2. Do not upgrade Implementation/Support/Evidence without evidence.
3. Map PSR domains by the domain-level guidance (§9), refining per capability where a specific capability clearly differs.
4. Assign responsibility codes reflecting product vs operator duties.
5. Use `not-assessed` / `mapping-review-required` for uncertain values rather than guessing.
6. Never introduce a `compliant`/`requirement-satisfied` value.
7. Keep roadmap/implementation/support/evidence/security-governance in separate columns.

## 14. Unknown and Deferred States

- Uncertain PSR mapping → `mapping-review-required`.
- Unassessed evidence → `not-assessed`.
- Security/governance not yet designed → `design-required` (or `not-assessed`).
- Unknown states are never silently interpreted as positive (no implicit `implemented`/`supported`/`validated`/`compliant`).

## 15. Exception Handling

Any deviation touching security invariants, accepted ADRs, or NDF rules requires a documented exception with rationale, a risk entry, Nova review, and Human-Maintainer approval. No silent exceptions. Security/Governance Status `exception-required` flags a capability needing such a decision.

## 16. Relationship to Detailed BSI Mapping

```text
PSR-domain mapping is a readiness and governance relationship.
It is not a clause-level BSI control mapping.
```

This WP maps no individual BSI requirement numbers, Kompendium controls, certification requirements, authority decrees, VS-NfD requirements, or concrete retention periods. Version-accurate BSI source ingestion and control mapping remain separate later work (see [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](BSI_REFERENCE_AND_CLAIMS_REGISTER.md) §13).

## 17. Compatibility

Additive and optional. Existing accepted decisions, the BSI/public-sector baseline, the ITIL/PRINCE2 tailoring, the sovereignty policy, and NDF rules remain valid. Single-status capability matrices remain allowed elsewhere; the Foundation Matrix needs the extra dimensions due to its security/public-sector scope. No ADR; no technology selected; breaking-change potential: low.

## 18. Validation Boundary

No capability is validated in this WP. Alignment status is documentation-only. Validation of any capability requires implementation plus evidence plus, where required, external assessment.

## 19. Open Questions

- Which capabilities become mandatory per profile vs. optional?
- Which PSR mappings currently marked `mapping-review-required` need refinement?
- Reconciliation of the project-wide capability count (see Foundation Matrix count note): **completed**. The authoritative count is **94** unique capability IDs; the reconciliation of the remaining "74" references was carried out in `CO-WP-029` (commit `6afa7ab`) and deterministically re-verified in `CO-WP-031`. No document states "74" as the authoritative count any more; the remaining occurrences are historical review and lessons records and are retained as such (`RISK-66` `closed`, HM-8).
- When is version-accurate BSI control mapping scheduled (separate WP)?

## 20. Next Decision

Nova review of this alignment, followed by a Human-Maintainer commit. Only afterwards may the locally planned `CO-WP-005` begin. Detailed BSI control mapping remains separate, later work.
