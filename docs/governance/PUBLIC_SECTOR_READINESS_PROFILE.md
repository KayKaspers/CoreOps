# CoreOps – Public-Sector Readiness Profile

> Document Status: Implemented, pending Nova review
> Profile Status: Internal readiness model
> Certification Status: None
> Government Approval Status: None
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004C` (docs-only / security-governance baseline)

## 1. Status

Internal readiness model defining three CoreOps product profiles (Standard, Hardened, Government) and their inheritance, responsibilities, evidence and offline/sovereignty expectations, and claim boundaries. These are **internal product profiles**, not BSI classifications and not certification levels.

## 2. Purpose

Give operators a clear, honest profile model for public-sector-adjacent deployments, with explicit boundaries so that no profile is mistaken for certification, government approval, or deployment compliance.

## 3. Internal Profile Boundary

```text
Standard / Hardened / Government are internal CoreOps product profiles.

They are NOT:
- BSI classifications,
- certification levels,
- government approvals,
- deployment-compliance statements.
```

Profiles describe intended product configuration and readiness rigor. Actual compliance depends on deployment, operator processes, configuration, evidence, legal applicability, and external assessment (see [BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](../security/BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md) §5.3).

## 4. Standard Profile

- Secure default configuration for general self-hosted and enterprise environments.
- Basic audit and role capability.
- No special authority claim.
- Readiness domains at baseline rigor (see PSR-01…PSR-18).

## 5. Hardened Profile

Extends Standard with at least:
- more restrictive security defaults,
- stronger authentication and role limitation,
- stricter logging and audit requirements,
- extended hardening and integrity checks,
- higher evidence expectations,
- more controlled change processes,
- stronger offline and recovery capability.

## 6. Government Profile

Extends Hardened as an internal readiness profile with at least:
- public-sector-oriented governance,
- traceable responsibility assignment,
- more comprehensive audit and evidence expectations,
- clear separation of product and operator responsibility,
- controlled dependencies,
- documented offline/air-gap capability,
- sovereignty and data-control expectations,
- stronger change, approval, and separation-of-duties limits,
- exportable evidence data,
- explicit claim boundaries.

```text
Government Profile
≠ BSI certification
≠ government approval
≠ VS-NfD approval
≠ deployment compliance
```

## 7. Profile Inheritance

```text
Standard  ⊂  Hardened  ⊂  Government
```

Each higher profile inherits all expectations of the lower profile and adds stricter requirements. A higher profile never relaxes a lower profile's security expectation. Choosing a profile does not by itself implement it; each expectation still requires design, implementation, configuration, and evidence.

## 8. Mandatory Profile Decisions

An operator adopting a profile must decide at least:
- protection needs (Schutzbedarf) for the deployment,
- role and privilege assignment,
- logging retention and access policy,
- backup targets and restore-test cadence,
- network segmentation and trust boundaries,
- dependency acceptance and update policy,
- offline/air-gap mode (if applicable),
- evidence collection and review responsibilities.

## 9. Product Responsibility

Secure product defaults; role and permission model; audit functions; configuration options; evidence export; integrity and update functions; documentation. The product provides capabilities; it does not, by itself, satisfy a requirement.

## 10. Operator Responsibility

Schutzbedarfsfeststellung; concrete network segmentation; user/role assignment; retention periods; backup targets; operational processes; incident response; legal applicability; local approvals; assessor interaction.

## 11. Shared Responsibility

Logging configuration; patch/maintenance windows; integrations; backup/restore tests; dependency updates; risk acceptance; evidence collection. Shared items require explicit assignment in a concrete deployment.

## 12. Evidence Expectations

Per profile, increasing rigor across the evidence categories defined in the baseline (§11). At minimum, evidence must be traceable, time-referenced, scoped, attributable to a role, integrity-protected, and retention/export-capable. The three-state distinction (evidence capability ≠ evidence available ≠ requirement satisfied) applies to every profile.

## 13. Offline and Air-Gap Expectations

- Standard: no mandatory cloud for core function.
- Hardened: offline-capable core, local identity/audit, controlled update import.
- Government: documented offline/air-gap capability, provenance/integrity verification, evidence export without permanent external connectivity, backup/restore without cloud dependency, defined handling of unreachable external services, exit/restart strategy.

Not claimed: fully implemented; suitable for every air-gap tier; approved for classified networks.

## 14. Sovereignty Expectations

- No mandatory external control plane for the core.
- Cloud services remain optional integrations.
- Data location, jurisdiction, operator access, and exit assessed (especially for Government).
- C5:2026 / C3A are conditional references for relevant cloud scenarios only.
- Sovereignty is multi-dimensional (technical, operational, legal, exit).

## 15. Exceptions

Any deviation from a profile's expectations requires a documented exception with rationale, risk entry, Nova review, and Human-Maintainer approval. No silent exceptions.

## 16. Claims Boundary

Allowed and prohibited claims follow [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](../security/BSI_REFERENCE_AND_CLAIMS_REGISTER.md) §9–§11 and the baseline §5. A profile name must never be presented as a certification, approval, or compliance statement.

## 17. Validation Boundary

No profile is validated in this work package. Profile status is `Internal readiness model`. Validation of any profile expectation requires implementation plus evidence plus, where required, external assessment.

## 18. Future Capability Mapping

Mapping profiles and PSR domains onto concrete capabilities and their status dimensions is the subject of `CO-WP-004E`. This document does not modify the Capability Matrix.

## 19. Open Questions

- Which PSR domains are mandatory vs. optional per profile?
- What is the minimum evidence set to call a profile expectation "design-complete"?
- How are profile exceptions tracked over a deployment lifecycle?
- How does the Government profile interact with ITIL/PRINCE2 tailoring (`CO-WP-004D`)?
