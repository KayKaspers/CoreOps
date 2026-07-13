# CoreOps – Capability Matrix Specification

> Document Status: Implemented, pending Nova review
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004E` (docs-only / capability-governance alignment)

> **Hinweis:** Diese Spezifikation existierte in CoreOps zuvor nicht als eigenständige Datei; die Statusdefinitionen waren in der [FOUNDATION_CAPABILITY_MATRIX.md](../architecture/FOUNDATION_CAPABILITY_MATRIX.md) eingebettet. `CO-WP-004E` erstellt sie additiv als projektlokale Spezifikation und erweitert das Modell um Security-/Governance-Dimensionen. Sie ersetzt keine bestehende Matrixzeile und wählt keine Technologie.

## 1. Purpose

Define the field and status model for CoreOps capability matrices, including the multi-dimensional status model required by the Foundation Capability Matrix due to its security and public-sector scope. Companion: [CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md](../security/CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md).

## 2. Capability Identity

- Stable IDs `CAP-<DOMAIN>-<NNN>`, unique and preserved across work packages.
- Names are unique within the matrix.
- No capability is deleted; changes are additive.
- No parallel capability list with divergent IDs.

## 3. Multi-Dimensional Status Fields

A capability carries five independent status dimensions (a single overloaded status is not used where these vary independently):

- **Roadmap Status** — planning only: `not-planned` · `observe` · `planned` · `target-foundation` · `target-later` · `deferred` · `out-of-scope`. Project-valid target values (`target-observe`, `target-map`, `target-plan`, `target-deploy`, `target-automate`, `target-extend`) remain valid.
- **Implementation Status** — `not-designed` · `design-required` · `designed` · `implementation-required` · `in-progress` · `implemented` · `partially-implemented` · `not-applicable`. `implemented` requires technical evidence. (Legacy value `not-implemented` maps to `implementation-required`.)
- **Support Status** — `not-assessed` · `not-supported` · `experimental` · `planned-support` · `supported` · `deprecated` · `not-applicable`. `supported` requires the defined support-evidence set.
- **Evidence Status** — `not-assessed` · `evidence-model-required` · `evidence-capability-designed` · `evidence-capability-implemented` · `evidence-available` · `validated` · `not-applicable`. `validated` requires documented assessment.
- **Security/Governance Status** — `not-assessed` · `baseline-mapped` · `design-required` · `controls-required` · `evidence-required` · `review-required` · `exception-required` · `not-applicable`. No `compliant` value.

## 4. Profile Relevance

Readiness profiles: `Standard` (Std), `Hardened` (Hard), `Government` (Gov). Governance profiles: `Service Operations` (SO), `Major Deployment Project` (MDP), `Public-Sector Delivery` (PSD). Profiles may be combined. Profile relevance is not certification.

## 5. PSR Domain Mapping

- Uses only PSR-01…PSR-18 from the readiness baseline.
- A capability may map to several, none, or `mapping-review-required`.
- PSR mapping means "relevant to the readiness domain," not "BSI requirement satisfied / control implemented / deployment compliant."

## 6. Responsibility Codes

`P` product · `O` operator · `S` shared · combinations `P/O`, `P/S`, `O/S`, `P/O/S`. A single `shared` value alone is insufficient without explaining product vs operator duties. Semantics reference the responsibility model in the BSI/public-sector baseline §10.

## 7. Evidence Model

Three states kept separate: `Evidence capability` (product feature, may be planned/implemented) ≠ `Evidence available` (deployment-specific) ≠ `Requirement satisfied` (assessed and accepted; never asserted by the product matrix).

## 8. Claim Boundaries

No blanket values such as `compliant`, `BSI-compliant`, `requirement-satisfied`, or `government-approved`. Allowed and prohibited claims follow the BSI reference/claims register §9–§11. No automatic compliance inference from any status combination.

## 9. Unknown and Not-Assessed States

Uncertain or unassessed fields use `not-assessed` / `mapping-review-required` / `design-required`. Unknown states are never silently interpreted as positive. No status is upgraded without evidence.

## 10. Single-Status Matrices

Single-status capability matrices remain allowed where the additional dimensions are not needed (small/simple projects). The Foundation Capability Matrix requires the multi-dimensional model because of its security/public-sector scope.

## 11. Migration of Existing Single-Status Matrices

To migrate an existing single-status matrix additively: preserve IDs/names; keep existing roadmap/implementation/support columns; add Evidence Status, Security/Governance Status, Profile Relevance, PSR Domains, Responsibility; set uncertain values to pending; do not upgrade any status without evidence; do not create a second parallel list.

## 12. Compatibility

Additive to existing CoreOps documents; changes no accepted decision; selects no technology; creates no ADR. Breaking-change potential: low.
