# CoreOps – Next Phase

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist **nicht** normativ.

## Current Phase

`Foundation 0.1 – Platform Foundation`

> Vorläufiger Arbeitsname. Die endgültige Phasen- und Release-Taxonomie wird durch `CO-WP-003` festgelegt.

## Local NDF Skills Baseline

Complete NDF v1.0.0 skills pack implemented (CO-WP-001A: GO).

## Concept Registration

Concept v3.0 vollständig registriert (CO-WP-002: GO WITH NOTES); Decision Classification, Decision Index und Risk Register vorhanden.

## Brief, Scope Lock and Release Taxonomy

Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (CO-WP-003: GO WITH NOTES, alle `Proposed for acceptance`); durch Human-Maintainer-Commit angenommen. Docker-first eingeordnet; aktive Queue autoritativ; NDF-`main` und NDF-Level geklärt; Repository-URL verifiziert.

## Capability Matrix and Support Boundary

Foundation Capability Matrix (74 Capabilities) und initiale Observe-Supportgrenze erstellt (CO-WP-004: GO WITH NOTES, `Proposed for acceptance`); durch Human-Maintainer-Commit angenommen. Drei Statusdimensionen und Support-Evidence definiert; `CCR-12` vorgeschlagen aufgelöst. Keine Runtime-Capability implementiert, keine Integration `supported`.

## Sovereignty and BSI Orientation

Concept-v3.1-Amendment, Sovereignty-and-Dependency-Policy und BSI-Alignment-Positioning erstellt (CO-WP-004A: GO WITH NOTES). Produktsouveränität und BSI-orientierte Entwicklung akzeptiert; verpflichtende externe Managementprodukte als Kernabhängigkeit ausgeschlossen; keine Zertifizierung/VS-Eignung behauptet; Standard-/Hardened-/Government-Profile registriert.

## Lessons Learned and NDF Feedback Governance

Lessons-Learned-Prozess und NDF-Feedback-Prozess etabliert (CO-WP-004B: GO WITH NOTES); 15 Lessons retrospektiv/neu erfasst; 7 reservierte NDF-Feedback-Kandidaten bewertet.

## First NDF Feedback Transfer Package

`CO-WP-004B1 – First NDF Feedback Transfer Package` completed; Nova Review pending.
Übergabeschwelle mit sieben Kandidaten erreicht (Nova-Feststellung); Transfer Package 001 mit drei Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer` (Nova Gate: approved, Human-Maintainer Gate: approved, Commit 4ad3111). Kein Kandidat übertragen/adoptiert, kein NDF-Repository verändert, kein NDF-Work-Package erstellt.

## NDF Intake Approval Gate Finalization

`CO-WP-004B2 – Finalize NDF Intake Approval Gates` completed; Nova Review: `GO WITH NOTES`.
Transfer Package Status auf `Approved for NDF Intake` gesetzt; alle 7 Human-Maintainer-Gates auf `approved` gesetzt. Neue Lesson LL-016 erfasst. NDF-Repository unverändert.

## NDF Intake Transfer Recording

`CO-WP-004B3 – Record Completed NDF Intake Transfer` completed; Nova Review pending.
Transfer Package 001 wurde dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e). Alle 7 Kandidaten auf `transferred-to-ndf` gesetzt.

## NDF Adoption Completion and Transfer Package Closure

`CO-WP-004B4 – Record Completed NDF Adoption and Close Transfer Package 001` completed; Nova Review pending.
Alle 7 Kandidaten auf `adopted-in-ndf` (drei Adoption-Commits 1ebffa6 / e894c6f / ebf716c); Transfer Package 001 geschlossen.

## BSI and Public-Sector Readiness Baseline

`CO-WP-004C – BSI and Public-Sector Readiness Baseline` completed; Nova Review pending. Drei Dokumente (Readiness-Baseline mit 18 PSR-Domänen, Reference-/Claims-Register, Public-Sector-Profil); keine Zertifizierung/Compliance behauptet.

## ITIL and PRINCE2 Applicability and Tailoring

`CO-WP-004D – ITIL and PRINCE2 Applicability and Tailoring Decision` completed; Nova Review pending. ITIL `adopted-with-tailoring`, PRINCE2 V7 `optional-profile`, NDF bleibt primär; drei interne Governance-Profile; keine Zertifizierung.

## Capability Matrix Security and Governance Alignment

`CO-WP-004E – Capability Matrix Security and Governance Alignment` completed; Nova Review pending. Foundation Capability Matrix mehrdimensional ausgerichtet (94 Capabilities); PSR-Mapping ≠ Compliance.

## Language Standard, Public Neutrality and Repository Governance

`CO-WP-005 – Language Standard, Public Neutrality and Repository Governance` completed; Nova Review pending. Englisch kanonisch (maschinenbezogen), DE/EN Produkt; Neutralitäts-/Disclosure-Grenzen; Source-of-Truth-Hierarchie; Human-Maintainer-Gates.

## System Context, Plane Taxonomy and External Boundaries

`CO-WP-006 – System Context, Plane Taxonomy and External Boundaries` completed; Nova Review pending. Systemkontext, 10 Planes und 11 Vertrauensgrenzen definiert; keine Technologieauswahl.

## Threat Model and Trust Boundaries

`CO-WP-007 – Threat Model and Trust Boundaries` completed; Nova Review pending. Foundation Threat Model (24 Assets, 40 Threat Scenarios THR-001…040, 17 Invarianten, 5 Abuse Cases); keine implementierten/validierten Kontrollen.

## Architecture and Module Boundaries

`CO-WP-008 – Architecture and Module Boundaries` completed; Nova Review pending. 17 logische Module (MOD-*), Autoritätsgrenzen, Datenownership, verbotene Bypässe; keine Technologie-/Deployment-Auswahl.

## Human Identity, Workspaces, RBAC and Break Glass

`CO-WP-009 – Human Identity, Workspaces, RBAC and Break Glass` completed; Nova Review pending. Human-Identity-, Workspace/RBAC- und Break-Glass-Governance; keine Auth-/IdP-/Policy-Engine-Auswahl.

## Machine Identity, Enrollment and Offline Credential Lifecycle

`CO-WP-010 – Machine Identity, Enrollment and Offline Credential Lifecycle` completed; Nova Review pending. Machine-Identity-, Enrollment-/Trust- und Offline-Credential-Governance; keine PKI-/Krypto-/Protokoll-Auswahl.

## Source of Truth and Field Provenance

`CO-WP-011 – Source of Truth and Field Provenance` completed; Nova Review pending. SoT ≠ System of Record, Field Provenance, Konfliktmodell; keine Storage-/Merge-/Krypto-Auswahl.

## Observed, Desired, Effective State and Drift

Current implemented WP: `CO-WP-012 – Observed, Desired, Effective State and Drift` – pending Nova review.

Current state-management foundation:
Desired, observed, effective and last-known state semantics, drift detection, convergence and safe-remediation governance defined.

Drift engine:
not implemented

Automatic remediation:
not selected

Runtime verification:
not implemented

Drei neue Dokumente: Observed/Desired/Effective State Model (State-Semantik, Desired-State-Lifecycle, Effective-State-Grenze, Unknown/Conflicted), Drift Detection and Convergence Model (Drift-Definition/-Arten, Detection States, Drift Record, Impact/Urgency, Exceptions, Konvergenz, Partial Failure) und Safe Remediation and State Change Policy (Detection/Recommendation/Plan/Approval/Execution/Verification getrennt, Read-only-Grenze, Remediation-Lifecycle, Rollback, fail-closed). Decision Index +15 (DEC-S-122…136), Risk Register +17 (RISK-173…189). Keine Beobachtung ≠ keine Drift; kein Drift ≠ Compliance; Drift-Erkennung ohne Write Authority; Executed ≠ successful ≠ verified ≠ compliant; keine Engine-/Scheduler-/Queue-/Auto-Remediation-Auswahl; keine ADR; SoT-/Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.

**Milestone Lessons Review Eligibility:** yes (CO-WP-005…012 = acht WPs erreicht) — gebündelte Bündelentscheidung durch Nova/Human Maintainer nach CO-WP-012-Commit; kein Review/Transfer automatisch gestartet.

Next planned WP: `CO-WP-013 – Policy, Approval and Execution Authorization`.

## Milestone Lessons Review (CO-WP-005…012)

`Milestone Lessons Review CO-WP-005 through CO-WP-012` completed; Nova Review pending. Gebündelter Review der acht Foundation-WPs (docs-only). Ergebnis: **GO WITH NOTES FOR CO-WP-013**. Neues Dokument `project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md`. Sechs konsolidierte Lessons (LL-017…022) und drei NDF-Feedback-Kandidaten (NDF-FC-COREOPS-008…010, `candidate-pending-nova-review`) erfasst. Read-only-Befunde: Risk Register (189 Einträge) braucht einen späteren Konsolidierungslauf; Decision Index mischt Alt-Kombistatus (DEC-S-01…37) mit getrennten Dimensionen (DEC-S-38…136); Capability-Count „74→94" in Alt-Abschnitten. Follow-ups bis ~CO-WP-029/030. **Risk Register und Decision Index unverändert (read-only); keine NDF-Rückführung; CO-WP-013 nicht begonnen.**

Milestone Review Status: `completed-go-with-notes` (Commit 74f8e32).

## Policy, Approval and Execution Authorization

WP Status:
`CO-WP-013 – Policy, Approval and Execution Authorization` – `completed-go-with-notes` (Commit 438a5a0).

Current authorization foundation:
Policy evaluation, approval and execution authorization are separate, bounded and auditable. A policy permit does not imply approval or execution authorization; execution authorization is action-, target-, scope-, plan- and time-bound; expired/revoked/consumed authorization is not reusable; machine principals cannot self-approve.

Policy engine:
not selected

Approval engine:
not selected

Execution authorization mechanism:
not implemented

Drei neue Dokumente: Policy Decision and Evaluation Model (Policy-Klassen, Lifecycle, Evaluation Inputs, Decision Outcomes, Default-Deny, Konflikte/Präzedenz, Exceptions, Offline Evaluation), Approval and Authorization Lifecycle (Approval Requirements/Requests/Decisions, Approver Authority, Separation of Duties, Expiry/Revocation/Consumption, Replay-Grenze, Offline Approval) und Execution Authorization and Guard Policy (Execution Intent/Plan/Authorization, Authorization Lifecycle, Pre-Execution Guards, Execution Boundary, Replay/Duplicate, Verification/Closure, Fail-Closed). Decision Index +16 (DEC-S-137…152), Risk Register +17 (RISK-190…206). Keine Policy-/Approval-/Execution-Engine, kein Autorisierungsartefakt/Token, kein Replay-Mechanismus, keine Queue/Workflow-Runtime ausgewählt; keine ADR; Identity-/Modul-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

## CoreOps Integration Contract v0.1

WP Status:
`CO-WP-014 – CoreOps Integration Contract v0.1` – `completed-go-with-notes` (Commit 611773b).

Current integration foundation:
CoreOps Integration Contract v0.1, capability and operation model, and integration trust/failure policy defined. Advertised/detected/permitted/implemented/supported/validated capabilities are separate; request acceptance ≠ authorization ≠ execution; completion ≠ success; success ≠ verification; unknown outcome stays explicit and blocks automatic retry; read-only integration gains no silent write authority; integration results inherit no authoritative state automatically; adapters/agents cannot expand scope; contract extensions cannot override security invariants.

Protocol and schema:
not selected

SDK and adapters:
not implemented

Runtime integration:
not implemented

Drei neue Dokumente: CoreOps Integration Contract v0.1 (Contract-Version, Integration Classes/Identity/Lifecycle, Request-/Acceptance-/Operation-Lifecycle, Result-/Failure-Semantik, Async, Retry/Replay, Cancellation/Rollback/Recovery, Offline, Provenance, Versioning, Extensions), Integration Capability and Operation Model (sechs Capability-Dimensionen, Operation Classes, Privilege Classification, Detection/Permission/Validation) und Integration Trust, Failure and Recovery Policy (Trust Boundary, Failure Classification, Unknown Outcome, Retry/Replay, Partial Failure, Rollback/Recovery, Fail-Closed). Decision Index +16 (DEC-S-153…168), Risk Register +13 (RISK-207…219, gesamt 219). Keine Protokoll-/Schema-/Transport-/SDK-/Adapter-/Replay-/Queue-/Messaging-Technologie ausgewählt; keine ADR; Identity-/Authorization-/State-/Provenance-/Modul-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

## Domain Pack Governance, Support Levels and Compatibility

WP Status:
`CO-WP-015 – Domain Pack Governance, Support Levels and Compatibility` – `completed-go-with-notes` (Commit 8191e06).

Current modular-product foundation:
Domain-Pack governance, lifecycle, support levels, compatibility claims and trust boundaries defined. A Domain Pack is a versioned governance/product boundary (≠ adapter/plugin/deployment unit/certification); pack lifecycle/maintenance/support/implementation/validation/evidence/security-review/compatibility are separate dimensions; support level ≠ SLA; SUP-3 is version-/target-/profile-/limitation-/evidence-bound; community/external origin ≠ trust/support/endorsement; pack activation grants no runtime authority; offline pack use needs provenance/integrity/target binding/explicit activation.

Domain-Pack implementation:
not started

Packaging and plugin runtime:
not selected

Compatibility validation:
not performed

Drei neue Dokumente: Domain Pack Governance Model (Pack-Begriff/-Klassen/-Identität, Statusdimensionen, Lifecycle, Ownership/Maintenance, Dependencies, Composition/Overlap, Community/External, Vendor Neutrality, Offline), Domain Pack Support and Compatibility Model (SUP-0…SUP-3/SUP-D, Maintenance, Implementation/Validation/Evidence, Compatibility Status/Dimensions/Claims, Deprecation/Retirement) und Domain Pack Trust, Provenance and Lifecycle Policy (Trust/Provenance/Integrity, Security Review/Response, Suspension/Withdrawal, Offline Distribution, Fail-Closed). Decision Index +17 (DEC-S-169…185), Risk Register +10 (RISK-220…229, gesamt 229). Keine Packaging-/Marketplace-/Plugin-/Update-/Dependency-Resolution-/Signaturtechnologie ausgewählt; keine ADR; Integration-/Identity-/Authorization-/State-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

## Data Ownership, Persistence, Schema Versioning and Migration

WP Status:
`CO-WP-016 – Data Ownership, Persistence, Schema Versioning and Migration` – `completed-go-with-notes` (Commit 69b3334).

Current data foundation:
Data ownership, persistence classes, schema versioning, migration safety and recovery governance defined. Data owner/steward/storage/write/migration/retention/recovery are separate responsibilities; storage responsibility ≠ authoritative ownership; migration authority ≠ unrestricted write; schema version ≠ data version; unknown/conflicted compatibility blocks automatic destructive migration; executed migration ≠ validated integrity; backup exists ≠ restorable; restore completed ≠ service recovery; migration does not reactivate revoked authority; audit/evidence provenance preserved; offline migration needs provenance/integrity/target binding/explicit activation.

Storage and database technology:
not selected

Migration implementation:
not started

Backup and recovery validation:
not performed

Drei neue Dokumente: Data Ownership and Persistence Model (Ownership-Dimensionen, Datenklassen/autoritative Module, Persistence-Klassen, Data Lifecycle, Retention/Recovery/Privacy), Schema Versioning and Migration Model (Schema Identity, Versionsdimensionen, Compatibility-/Change-Klassen, Migration Plan/Preflight/Lifecycle, Mixed-Version/Partial, Rollback/Forward Recovery) und Data Migration Integrity and Recovery Policy (Authority, Klassifikation, Backup/Recovery, destruktive Migration, Identity-/Audit-Daten, Offline, Concurrency, Fail-Closed). Decision Index +18 (DEC-S-186…203), Risk Register +10 (RISK-230…239, gesamt 239). Keine Storage-/DB-/ORM-/Schema-/Serialisierungs-/Migration-/Backup-/Replikations-/Cluster-/Transaction-Technologie ausgewählt; keine ADR; Integration-/Domain-Pack-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

## API Governance, Versioning, Errors and Idempotency

WP Status:
`CO-WP-017 – API Governance, Versioning, Errors and Idempotency` – `completed-go-with-notes` (Commit 7170c84).

Current API foundation:
API governance, independent version dimensions, compatibility, error semantics, idempotency and replay boundaries defined. API identity ≠ route/URL/transport/product version; API availability ≠ authorization; request acceptance ≠ execution; successful response ≠ verified outcome; error response ≠ proof of no side effect; idempotency context ≠ execution authorization; unknown outcome blocks automatic retry and requires reconciliation; bulk preserves per-target authority and partial results; pagination/continuation imply no snapshot/current authorization; read-only API gains no silent write authority.

API transport and style:
not selected

API implementation:
not started

Compatibility and idempotency validation:
not performed

Drei neue Dokumente: API Governance and Operation Model (API Classes/Identity/Lifecycle, Producers/Consumers, Operation Classes, Side-Effect-Klassifikation, Request/Response, Acceptance/Result, Bulk, Pagination, Async, Policy/Authorization-Bindung, Workspace-Isolation), API Versioning, Compatibility and Deprecation Model (zwölf Versionsdimensionen, Compatibility-/Change-Klassen, Request/Response/Error/Behavioural Compatibility, Deprecation/Retirement) und API Error, Idempotency and Replay Policy (Error Classes/Disclosure, Retry Classification, Idempotency Context, Duplicate/Replay, Unknown Outcome, Fail-Closed). Decision Index +16 (DEC-S-204…219), Risk Register +10 (RISK-240…249, gesamt 249). Keine Transport-/API-Style-/Statuscode-/Schema-/Gateway-/Idempotency-/Replay-/Rate-Limit-Technologie ausgewählt; keine ADR; Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

## Event, Audit Correlation and Evidence Model

Current implemented WP:
`CO-WP-018 – Event, Audit Correlation and Evidence Model` – pending Nova review.

Current audit and evidence foundation:
Event identity, correlation, causation, audit completeness, evidence references, validation, sufficiency, retention, disclosure and offline continuity defined. Event ≠ command ≠ notification ≠ audit event ≠ evidence; event identity ≠ correlation/request/operation/attempt; occurrence/observation/recording/ingestion time separate; timestamp/sequence ≠ authoritative global ordering; correlation ≠ causation; recorded ≠ validated/complete; missing event ≠ no action; evidence capability/availability/freshness/integrity/validation/sufficiency separate; sufficiency decision-/scope-/time-bound; audit administrator ≠ unrestricted disclosure; closed ≠ complete/validated/sufficient/compliance.

Event transport and storage:
not selected

Audit implementation:
not started

Evidence validation:
not performed

Drei neue Dokumente: Event and Audit Correlation Model (Event Classes/Identity, Producers/Sources, vier Zeitbegriffe, Clock Uncertainty, Correlation/Causation, Operation/Attempt-Linkage, Ordering/Sequence, Duplicate/Replay, Audit Events/Lifecycle, Gaps/Completeness), Evidence Reference, Validation and Lineage Model (Evidence-Dimensionen, References/Sets, Availability/Freshness/Integrity/Provenance/Validation/Sufficiency, Conflicts/Supersession) und Audit Integrity, Retention and Disclosure Policy (Authority, Integrity, Completeness/Gaps, Retention/Archival/Purge, Disclosure/Export, Workspace-Isolation, Privacy, Offline, Failure/Unknown, Fail-Closed). Decision Index +17 (DEC-S-220…236), Risk Register +10 (RISK-250…259, gesamt 259). Keine Event-Bus-/Logging-/SIEM-/Storage-/Schema-/Ordering-/Hash-/Signatur-/WORM-/Redaction-Technologie ausgewählt; keine ADR; API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.

Next planned WP nach Nova Review und Human-Maintainer-Commit: `CO-WP-019 – Telemetry and Normalization Schema`.

## Current Goal

Dokumentation, Scope, Architektur, Security, Governance und Verträge definieren — ohne Anwendungscode, ohne verbindliche Technologieauswahl und ohne akzeptierte ADRs.

## In Scope

- Foundation-Dokumentation
- Sicherheitsgrundlagen
- Architekturmodelle
- Capability- und Supportgrenzen
- ADR-Vorbereitung
- Teststrategie
- Readiness Review

## Out of Scope

- Anwendungscode
- produktive Integrationen
- Agents
- Deployments
- Scans
- Netzwerkänderungen
- Druckerverwaltung
- Workflow-Ausführung
- Secrets-Verarbeitung
- produktive Installation

## Next Work Package

`CO-WP-004C – BSI and Public-Sector Readiness Baseline` (after CO-WP-004B2 Nova review and Kay-Commit)

## Foundation Exit Conditions

> **Vorläufige Gates.** Es wird **nicht** behauptet, dass diese Gates bereits erfüllt sind.

- Foundation-Queue abgeschlossen
- Scope Lock vorhanden
- Threat Model vorhanden
- kritische Trust Boundaries beschrieben
- Architektur konsistent
- relevante ADRs durch Human Maintainer entschieden
- Security-Invarianten dokumentiert
- Teststrategie vorhanden
- keine offenen Foundation-Blocker
- Readiness Review abgeschlossen
- Release Preparation separat erfolgt
