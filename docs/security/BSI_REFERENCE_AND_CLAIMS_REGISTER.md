# CoreOps – BSI Reference and Claims Register

> Document Status: Implemented, pending Nova review
> Certification Status: None claimed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004C` (docs-only / security-governance baseline)

## 1. Status

Foundation register of official BSI (and conditional cloud/sovereignty) references and the CoreOps claim classes. References are registered as **external references**, not as invented clauses. No requirement number, verbatim requirement, or full coverage is asserted. Version-accurate source ingestion is a later step.

## 2. Purpose

Give CoreOps a single, honest place that (a) lists the official reference set with provenance and version-verification status, (b) defines which claims are allowed, prohibited, or conditional, and (c) records the pending source-ingestion follow-up — so that later control-mapping work has a controlled starting point.

## 3. Reference Classes

Allowed normative-boundary values for any registered reference within CoreOps:

```text
external-reference · applicability-review-required · profile-guidance ·
deployment-dependent · future-control-mapping-source
```

External BSI documents do **not** automatically become globally normative CoreOps rules merely by being registered here.

## 4. Official Reference Set

Reference date: 2026-07-13 (registered by Nova from official BSI sources).

| Reference ID | Title | Publisher | Reference class | Current version/edition | Version verification status | Applicability | Profile relevance | Normative within CoreOps | Source-ingestion status | Notes |
|---|---|---|---|---|---|---|---|---|---|---|
| REF-BSI-200-1 | BSI-Standard 200-1 — Managementsysteme für Informationssicherheit | BSI | external-reference | 200-1 | reference-registered, version-verification-pending | ISMS governance | Standard/Hardened/Government | external-reference | pending | ISMS support (PSR-01) |
| REF-BSI-200-2 | BSI-Standard 200-2 — IT-Grundschutz-Methodik | BSI | external-reference | 200-2 | reference-registered, version-verification-pending | methodology | Standard/Hardened/Government | external-reference | pending | Methodik (PSR-01/02) |
| REF-BSI-200-3 | BSI-Standard 200-3 — Risikomanagement | BSI | external-reference | 200-3 | reference-registered, version-verification-pending | risk management | Hardened/Government | external-reference | pending | Risk (PSR-01, risk register) |
| REF-BSI-200-4 | BSI-Standard 200-4 — Business Continuity Management | BSI | external-reference | 200-4 | reference-registered, version-verification-pending | BCM | Hardened/Government | external-reference | pending | Continuity (PSR-09/10) |
| REF-BSI-GS-KOMP | BSI IT-Grundschutz-Kompendium | BSI | future-control-mapping-source | edition to be confirmed | version-confirmation-required | control families | profile-dependent | applicability-review-required | pending | Edition must be version-confirmed before any detail mapping |
| REF-BSI-MIN-BV | Aktuelle Mindeststandards des BSI für die Bundesverwaltung | BSI | deployment-dependent | current set | reference-registered, version-verification-pending | federal administration | Government | applicability-review-required | pending | Applicability depends on deployment context |
| REF-BSI-MIN-PROTO | Mindeststandard zur Protokollierung und Detektion von Cyberangriffen | BSI | profile-guidance | v2.1 (2024-11-11) | version-registered (2.1) | logging & detection | Hardened/Government | applicability-review-required | pending | Especially relevant for PSR-07/08 |

## 5. Source Provenance

The full official documents were **not** provided to this work package as a local source pack. This register therefore records titles, publisher, and (where explicitly provided by Nova) a version marker, but no clause-level content. No requirement text is reproduced or paraphrased as if authoritative.

## 6. Version Verification Status

- Only the logging/detection Mindeststandard carries an explicit version (v2.1, 2024-11-11).
- The IT-Grundschutz-Kompendium edition is explicitly `version-confirmation-required` before any detail mapping.
- All other references are `version-verification-pending` and must be confirmed version-accurately during a later source-ingestion work package.

## 7. Applicability Status

Applicability is deployment- and profile-dependent. Registration here does not assert that any reference applies to a given CoreOps deployment. Federal-administration minimum standards (REF-BSI-MIN-BV) are `applicability-review-required` and depend on the concrete operator context.

## 8. Claim Classes

```text
allowed-claim · prohibited-claim · conditional-claim
```

## 9. Allowed Claims

```text
CoreOps is designed with BSI-oriented security and governance principles in mind.
CoreOps provides a foundation for deployments that may need to align with
  IT-Grundschutz-oriented or public-sector security processes.
The Government profile is an internal CoreOps readiness profile.
CoreOps can provide evidence-supporting capabilities where those capabilities are
  implemented, configured and validated.
```

## 10. Prohibited Claims

```text
CoreOps is BSI-certified.
CoreOps is fully IT-Grundschutz compliant.
CoreOps is approved for the German federal administration.
CoreOps is automatically suitable for classified information.
CoreOps is VS-NfD approved.
Using CoreOps makes an organisation compliant.
The Government profile is a certification level.
```

## 11. Conditional Claims

```text
Alignment with a specific IT-Grundschutz requirement family
  — conditional on version-accurate source ingestion and later control mapping.
Support for a specific logging/detection expectation
  — conditional on deployment configuration and evidence.
Cloud security alignment (C5) / cloud sovereignty alignment (C3A)
  — conditional on a cloud-involving deployment and an applicability review.
```

## 12. Cloud and Sovereignty References

| Reference ID | Title | Publisher | Reference class | Version | Applicability | Notes |
|---|---|---|---|---|---|---|
| REF-C5-2026 | C5:2026 — Cloud Computing Compliance Criteria Catalogue | BSI | deployment-dependent | 2026 | conditional (cloud only) | Not relevant for pure on-prem/offline CoreOps |
| REF-C3A | C3A — Criteria enabling Cloud Computing Autonomy | BSI | deployment-dependent | source verification pending | conditional (cloud only) | Sovereignty reference for relevant cloud scenarios; published 2026-04-27 |

Using a C5-assessed service does not automatically make a CoreOps deployment compliant. On-premises/offline CoreOps is not automatically C5- or C3A-relevant.

## 13. Source-Ingestion Follow-up

A later, dedicated work package must: obtain version-accurate official sources (with documented provenance and identity/completeness/truncation checks), confirm the IT-Grundschutz-Kompendium edition, and only then produce any clause-to-capability mapping. Until then, no clause-level claim is permitted.

## 14. Review Cadence

References must be re-checked for version currency at each security-baseline or public-sector work package, and at minimum whenever a control-mapping work package is planned. Outdated versions are a registered risk (RISK-41).

## 15. Human-Maintainer Gate

Registration, version confirmation, applicability decisions, and any future normative promotion of a reference require Nova review and Human-Maintainer approval. No reference is promoted to a globally normative CoreOps rule in this work package.

## 16. Open Questions

- Which IT-Grundschutz-Kompendium edition is current and applicable per profile?
- Which minimum standards for the federal administration apply to which deployment classes?
- What triggers a C5/C3A applicability review in practice?
- Which references become `profile-guidance` vs. remain `external-reference` after ingestion?
