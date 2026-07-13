# CoreOps – Service and Project Governance Profiles

> Document Status: Implemented, pending Nova review
> Profile Type: Internal optional governance profiles
> Certification Status: None
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004D` (docs-only / governance-framework applicability review)

## 1. Status

Three internal, optional, tailorable CoreOps governance profiles derived from the ITIL/PRINCE2 tailoring decision ([ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md)). These are **internal product/operations profiles**, not certification levels and not government approvals.

## 2. Purpose

Give operators and larger CoreOps initiatives a lightweight, optional governance vocabulary for service operations and project delivery, without mandating full ITIL or PRINCE2 and without replacing NDF as the primary development and repository-governance framework.

## 3. Profile Model

```text
CoreOps Service Operations Guidance      (ITIL-inspired, operations)
CoreOps Major Deployment Project Profile (PRINCE2-inspired, projects)
CoreOps Public-Sector Delivery Profile   (combines both + BSI/public-sector baseline)
```

All three are `internal`, `optional`, `tailorable`, and `not certification levels`. They may be combined.

## 4. Service Operations Guidance

- ITIL-inspired, for ongoing operations.
- Scalable and optional; small deployments may ignore it.
- Covers monitoring/event handling, incident/problem handling, change enablement, continual improvement, service-level and supplier governance (as guidance, not implemented capabilities).
- No ITIL certification; no mandatory external ITSM product.
- Maps operations concepts onto CoreOps capabilities (subject to `CO-WP-004E`), not onto a separate tool.

## 5. Major Deployment Project Profile

- PRINCE2-inspired and tailored, for larger deployments and migrations.
- Provides optional project governance: business justification, defined roles, stage/phase control, management by exception and tolerances, risk/quality/change governance, progress reporting, project-level tailoring.
- NDF remains authoritative for software development; the Human Maintainer retains commit, push, tag, and release control.
- A project risk log under this profile is separate from and does not replace the CoreOps Risk Register.

## 6. Public-Sector Delivery Profile

- Combines the [Public-Sector Readiness Baseline / Profiles](../security/BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md) with selected ITIL service governance and optional PRINCE2 project tailoring.
- For public-sector-adjacent introduction projects that need documented governance and evidence.
- No government approval; no certification level; claim boundaries per the BSI reference/claims register apply.

## 7. Profile Selection

- Profiles are opt-in per deployment or initiative.
- Small/self-hosted deployments may use none.
- Larger deployments may use Service Operations Guidance and/or the Major Deployment Project Profile.
- Public-sector initiatives may use the Public-Sector Delivery Profile (which references the other two).
- Selection and tailoring are Human-Maintainer decisions.

## 8. Roles and Responsibilities

- No new global authorities are created.
- Framework roles map onto existing NDF roles: Nova (planning/architecture/review), implementation agent (execution within a work package), Human Maintainer (approval, commit, push, tag, release, irreversible/privileged actions).
- Project-profile roles (e.g. a deployment lead) are project-local and never override NDF Human-Maintainer gates.

## 9. NDF Integration

- NDF remains the primary software-development and repository-governance framework.
- These profiles add operations/project governance around NDF, not a parallel development process.
- No duplicate work-package system; the CoreOps Work Package Queue remains authoritative.
- No parallel decision or risk register.

## 10. Change and Approval Gates

- All commit, push, tag, and release actions remain Human-Maintainer-only (NDF core rule DEC-G-01).
- Change/approval guidance from ITIL/PRINCE2 is advisory and routes through existing NDF gates.
- No profile can authorise an autonomous git or release action.

## 11. Evidence Expectations

- Where a profile references evidence, it reuses the evidence model and three-state distinction (capability ≠ available ≠ satisfied) from the BSI/public-sector baseline (§11 there).
- Project profiles may add lightweight project evidence (justification, stage decisions, tolerances) as project-local artifacts.

## 12. Tailoring

- Every profile is tailored to the concrete deployment/initiative; nothing is adopted wholesale.
- Prefer the lightest form that solves the problem; avoid ceremonies/artifacts with no benefit.
- Tailoring decisions are recorded per initiative, not globally mandated.

## 13. Exceptions

Any deviation that touches security invariants, accepted ADRs, or NDF rules requires a documented exception with rationale, a risk entry, Nova review, and Human-Maintainer approval. No silent exceptions.

## 14. Claim Boundaries

Follows [ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md) §15. No certification, endorsement, accreditation, or full-framework-implementation claim. Profile names are never presented as certification levels.

## 15. Compatibility

Additive and optional. Existing accepted decisions, the BSI/public-sector baseline, the sovereignty policy, and NDF rules remain valid. No ADR is created; no technology or tool is selected; no Capability Matrix change (belongs to `CO-WP-004E`). Breaking-change potential: low.

## 16. Open Questions

- Which profile elements become recommended defaults vs. remain fully optional?
- How do these profiles reference concrete capabilities after `CO-WP-004E`?
- Whether a lightweight tailoring checklist template is added later (not created here).
