# CoreOps – ITIL and PRINCE2 Applicability and Tailoring

> Document Status: Implemented, pending Nova review
> Decision Status: Tailoring decision implemented
> ITIL Status: Adopted with tailoring
> PRINCE2 Status: Optional guidance/profile
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004D` (docs-only / governance-framework applicability review)

## 1. Status

Tailoring decision for two external governance frameworks: ITIL (service-management) is **adopted-with-tailoring** as guidance; PRINCE2 Project Management Version 7 is an **optional project/deployment governance profile**. NDF remains the primary software-development and repository-governance framework. No certification, endorsement, or full-framework implementation is claimed. This work package selects no tooling and creates no ADR.

## 2. Purpose

Decide, traceably, which high-level ITIL and PRINCE2 concepts CoreOps uses as guidance, which it explicitly does not adopt, how the two frameworks relate to NDF, and how framework overload and duplicate governance are prevented — consistent with the [NDF External Framework Tailoring Guidance](../../../Nova-Development-Framework/docs/governance/NDF_EXTERNAL_FRAMEWORK_TAILORING_GUIDANCE.md) (read-only reference; adopted from CoreOps cross-project candidate 007).

## 3. Scope

- Applicability review and tailoring for ITIL (service-management guidance).
- Applicability review and tailoring for PRINCE2 Project Management Version 7 (project governance).
- The NDF relationship and a framework conflict hierarchy.
- Framework-overload guards and claim boundaries.
- Registration of internal CoreOps governance profiles (defined in the companion profiles document).

## 4. Non-Goals

- No ITSM system, ticketing system, or project-management engine.
- No technical integration or tool selection.
- No full ITIL or PRINCE2 methodology adoption.
- No certification, accreditation, endorsement, or training/exam content.
- No reproduction of extensive licensed framework content.
- No Capability Matrix change (belongs to `CO-WP-004E`).
- No ADR, no technology selection.

## 5. Reference and Source Boundary

Nova registered the following official reference status (July 2026). The full licensed framework publications were **not** handed over as a local source pack. Therefore this document uses only high-level, publicly acknowledged concepts; it reconstructs no full framework structure, reproduces no long official text, and invents no page/clause references or certification requirements.

- **ITIL:** ITIL 4 Foundation; ITIL Foundation (Version 5); ITIL Foundation Bridge (Version 5).
- **PRINCE2:** PRINCE2 Project Management (Version 7).

Version-accurate migration or clause-level mapping requires a later, reviewed official source pack.

## 6. NDF Relationship

```text
NDF remains the primary software-development and repository-governance framework for CoreOps.
```

- **ITIL** complements NDF with service-management guidance, operations processes, support/service governance, and continual improvement in operations.
- **PRINCE2** complements NDF optionally with governance for larger deployments, organisational rollout, complex projects, and project-level steering outside the software-development process.

Not permitted: a parallel second software-development process; a duplicate work-package structure; replacement of NDF Human-Maintainer gates; a parallel decision or risk register; automatic adoption of proprietary roles; ITIL or PRINCE2 as an overarching CoreOps development framework.

## 7. ITIL Applicability

ITIL is **adopted-with-tailoring** for: service-management guidance; ongoing operations; monitoring and event handling; incident and problem handling; change and deployment governance; configuration and asset reference; continual improvement; supplier and service-level governance; availability, continuity, and capacity as later operations domains.

Not automatically adopted: the full ITIL operating model; all practices; all roles; the full service-desk model; the certification structure; training/exam content; a mandatory external ITSM platform.

## 8. ITIL Applicability Matrix

Adoption levels: `core-guidance` · `profile-guidance` · `module-guidance` · `future-capability-reference` · `operator-responsibility` · `not-adopted` · `deferred`.

| Element | Version boundary | Applicability | Target CoreOps area | Profile relevance | Adoption level | Reason | Required evidence | NDF/CoreOps overlap | Follow-up |
|---|---|---|---|---|---|---|---|---|---|
| Guiding principles | version-robust | high | Governance mindset | Service Ops | core-guidance | Version-robust, aligns with NDF principles | design docs | Overlaps NDF principles | — |
| Service-management dimensions | version-robust | medium | Operations model | Service Ops | profile-guidance | Useful lens, not mandatory | design docs | Partial NDF overlap | CO-WP-004E |
| Value-system/value-chain concepts | v4/v5-tagged | medium | Operations model | Service Ops | profile-guidance | Version-tagged; high-level only | design docs | Low overlap | deferred detail |
| Continual improvement | version-robust | high | Ops improvement | Service Ops | core-guidance | Aligns with Lessons-Learned process | LL register (unchanged here) | Overlaps LL process | — |
| Monitoring and event management | version-robust | high | Observe domain | Service Ops | module-guidance | Maps to Observe capabilities | capability evidence | Overlaps Capability Matrix | CO-WP-004E/019 |
| Incident management | version-robust | high | Ops support | Service Ops | module-guidance | Core operations need | ops evidence | Low NDF overlap | later ops WP |
| Problem management | version-robust | medium | Ops support | Service Ops | module-guidance | Follows incident handling | ops evidence | Low overlap | later ops WP |
| Change enablement | version-robust | high | Change governance | Service Ops/Deployment | module-guidance | Aligns with PSR-13 | change/approval evidence | Overlaps PSR-13 | CO-WP-013 |
| Release management | version-robust | high | Release governance | Deployment | module-guidance | Aligns with deployment control | deployment evidence | Overlaps PSR-13/DEC | CO-WP-021 |
| Deployment management | version-robust | high | Deployment | Deployment | module-guidance | Deployment control plane | deployment evidence | Overlaps PSR-13 | CO-WP-021 |
| Service configuration management | version-robust | medium | Source of truth | Service Ops | future-capability-reference | Maps to config inventory | config evidence | Overlaps PSR-02 | CO-WP-011 |
| IT asset management | version-robust | medium | Asset inventory | Service Ops | future-capability-reference | Maps to asset inventory | config evidence | Overlaps PSR-02 | CO-WP-011 |
| Information security management | version-robust | high | Security governance | Hardened/Government | profile-guidance | Aligns with BSI baseline | security docs | Overlaps PSR-01/BSI baseline | CO-WP-007 |
| Service level management | version-robust | medium | Service governance | Service Ops | profile-guidance | Operator-facing | service docs | Low overlap | later ops WP |
| Supplier management | version-robust | medium | Supply chain | Service Ops/Government | profile-guidance | Aligns with dependency policy | dependency evidence | Overlaps PSR-14 | CO-WP-022 |
| Availability management | version-robust | medium | Operations | Hardened/Government | future-capability-reference | Later operations domain | ops evidence | Overlaps PSR-10 | deferred |
| Capacity and performance management | version-robust | medium | Operations | Hardened/Government | future-capability-reference | Later operations domain | ops evidence | Low overlap | deferred |
| Service continuity management | version-robust | high | Continuity | Hardened/Government | profile-guidance | Aligns with BCM/PSR-10 | continuity-test evidence | Overlaps PSR-09/10 | CO-WP-026 |
| Knowledge management | version-robust | low | Documentation | Service Ops | operator-responsibility | Mostly operator-owned | documentation evidence | Low overlap | — |
| Responsible automation and AI governance | version-robust | high | AI/automation limits | All | core-guidance | Aligns with "AI advisory only" (DEC-P-05) | design docs | Overlaps DEC-P-05/DEC-N-09 | CO-WP-013 |
| Full ITIL operating model | v4/v5-tagged | n/a | — | — | not-adopted | Would cause framework overload | — | — | — |

No full ITIL practice descriptions are reproduced.

## 9. ITIL Version Boundary

- ITIL 4 remains available as an official line.
- ITIL Foundation Version 5 is officially available.
- A Bridge path from ITIL 4 to Version 5 exists.
- CoreOps prefers version-robust governance principles.
- Version-specific concepts are tagged (e.g. `v4/v5-tagged`).
- No full ITIL Version 5 compatibility is claimed.
- Existing ITIL-4-oriented decisions are not automatically migrated to Version 5.
- A review trigger is set for material changes to the official ITIL line (see §17).
- Detailed version migration requires a reviewed official source pack.

## 10. PRINCE2 Applicability

PRINCE2 Project Management Version 7 is an **optional-profile** (with `guidance-only` elements where justified) for: larger CoreOps deployments; complex integration projects; migrations; multi-stakeholder releases; public-sector rollout projects; business justification; role/responsibility definition; stage/phase control; tolerances and escalation; risk/quality/change control; project-level tailoring.

Not globally mandatory: the full PRINCE2 role structure; mandatory ceremonies/meetings; all management products; full process adoption; certification of participants; PRINCE2 as a replacement for NDF.

## 11. PRINCE2 Applicability Matrix

Adoption levels: `optional-project-profile` · `deployment-project-guidance` · `release-governance-guidance` · `public-sector-project-guidance` · `already-covered-by-ndf` · `not-adopted` · `deferred`.

| Pattern | Applicability | Target scenario | Adoption level | Reason | NDF overlap | Required tailoring | Follow-up |
|---|---|---|---|---|---|---|---|
| Continued business justification | high | Larger deployments | optional-project-profile | Keeps projects purpose-driven | Low | Lightweight justification only | — |
| Defined roles and responsibilities | medium | Multi-stakeholder projects | deployment-project-guidance | Clear accountability | Maps to NDF roles | Map onto Human Maintainer/Nova, no new authorities | — |
| People and stakeholder considerations | medium | Public-sector rollout | public-sector-project-guidance | Stakeholder alignment | Low | Optional | — |
| Stage or phase control | high | Migrations/rollouts | deployment-project-guidance | Controlled progression | Partial (NDF WP flow) | Do not duplicate WP structure | — |
| Management by exception and tolerances | high | Larger deployments | optional-project-profile | Efficient escalation | Low | Define tolerances per project | — |
| Focus on products and outcomes | medium | All projects | already-covered-by-ndf | NDF is outcome/WP-focused | High | None (NDF covers) | — |
| Risk governance | high | Complex projects | optional-project-profile | Project-level risk | Overlaps Risk Register | Project risk log ≠ CoreOps Risk Register | — |
| Quality governance | medium | Deployments | deployment-project-guidance | Quality gates | Overlaps NDF review gates | Reuse NDF gates | — |
| Issue and change governance | high | Deployments/releases | release-governance-guidance | Controlled change | Overlaps PSR-13/NDF | Reuse Human-Maintainer gates | — |
| Progress reporting | low | Larger projects | deployment-project-guidance | Visibility | Low | Lightweight only | — |
| Tailoring to project context | high | All projects | optional-project-profile | Core tailoring principle | Aligns w/ NDF tailoring | Mandatory before use | — |
| Lessons and continual learning | medium | All projects | already-covered-by-ndf | NDF Lessons-Learned process | High | None (LL process covers) | — |

No full PRINCE2 processes or management products are reconstructed.

## 12. Tailoring Decisions

- **ITIL:** `adopted-with-tailoring` — selected concepts used as core/profile/module guidance; full operating model, all practices/roles, service-desk model, certification structure `not-adopted`.
- **PRINCE2 Version 7:** `optional-profile` — used as optional project/deployment governance; full role structure, mandatory ceremonies, all management products, full process adoption, participant certification `not-adopted`; outcome-focus and lessons `already-covered-by-ndf`.
- Both are tailored to CoreOps scope; neither replaces NDF; neither creates a compliance/certification claim.

## 13. Framework Conflict Hierarchy

```text
Applicable law and regulation
Human-Maintainer authority
Security and safety invariants
Accepted CoreOps ADRs and scope locks
Normative NDF rules
Accepted CoreOps governance
Tailored ITIL and PRINCE2 guidance
Advisory framework material
```

Tailored framework guidance never overrides a higher tier.

## 14. Framework-Overload Guards

- Adopt only elements that solve a real CoreOps problem; exclude/defer the rest explicitly.
- No parallel second software-development process; no duplicate work-package structure.
- No parallel decision or risk register; the existing Decision Index and Risk Register remain authoritative.
- Map any framework role onto existing NDF roles (Nova, implementation agent, Human Maintainer).
- Prefer the lightest form that solves the problem; avoid meetings/artifacts/terminology with no benefit.
- Prefer optional profiles over universal mandates.

## 15. Claim Boundaries

Allowed:

```text
CoreOps uses selected ITIL-aligned service-management concepts.
CoreOps provides optional project-governance guidance informed by PRINCE2 Project Management Version 7.
CoreOps tailors external frameworks to its own scope and governance.
NDF remains the primary development and repository-governance framework.
```

Prohibited:

```text
CoreOps is ITIL certified.
CoreOps is PRINCE2 certified.
CoreOps fully implements ITIL.
CoreOps fully implements PRINCE2.
Using CoreOps makes an organisation ITIL compliant.
CoreOps requires certified ITIL or PRINCE2 personnel.
CoreOps is officially endorsed by PeopleCert.
CoreOps implements every ITIL practice.
The CoreOps profiles are certification levels.
```

No partner, accreditation, or tool-certification claims.

## 16. Licensing and Attribution Boundary

ITIL and PRINCE2 are externally authored, licensed frameworks. This document uses only high-level, publicly acknowledged concept names for applicability assessment. It reproduces no extensive licensed content, no verbatim practice/process descriptions, and no exam/training material. No trademark, endorsement, or accreditation is claimed. Any version-accurate detail work requires a properly licensed source pack.

## 17. Review Triggers

- A material change to the official ITIL line (e.g. new version or bridge changes).
- A material change to PRINCE2's official version.
- A CoreOps operations or deployment work package that needs a specific framework element.
- A conflict discovered between tailored guidance and a higher hierarchy tier.
- Availability of a properly licensed source pack for version-accurate mapping.

## 18. Compatibility

Additive and optional. Existing accepted decisions, the BSI/public-sector baseline, the sovereignty policy, and NDF rules remain valid. No ADR hierarchy is overridden; no compliance/certification claim is created; no retroactive migration is forced; no new NDF or CoreOps version is claimed. Breaking-change potential: low.

## 19. Open Questions

- Which ITIL operations elements become concrete capabilities in `CO-WP-004E` and later operations WPs?
- Which PRINCE2 tolerances are recommended defaults for public-sector delivery projects?
- When is a licensed ITIL/PRINCE2 source pack obtained for version-accurate detail?
- How do the internal governance profiles interact with future ITSM/PM tool evaluations (if any)?

## 20. Next Decision

Nova review of this tailoring decision, followed by a Human-Maintainer commit. Only afterwards may `CO-WP-004E` begin. The internal governance profiles are defined in [COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md](COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md).
