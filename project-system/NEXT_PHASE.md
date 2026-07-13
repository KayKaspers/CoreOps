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

Current implemented WP: `CO-WP-008 – Architecture and Module Boundaries` – pending Nova review.

Current architecture state:
Foundation logical modules, stable module identities, authority boundaries and dependency rules defined.

Technology selection:
none

Deployment topology:
not selected

Implementation:
not started

Drei neue Dokumente: Logical Module Architecture (17 Module MOD-*, Modulklassifikation, Policy/Control/Execution-Trennung, Adapter-/Agentengrenzen, 2 Mermaid-Diagramme), Module Catalog (17-Modul-Register, Daten-/Zustandsownership, Threat-/Invarianten-Referenzen) und Module Boundary and Dependency Standard (Autoritätsgrenzen, erlaubte/verbotene Abhängigkeiten, verbotene Bypässe, Zyklusverbot). Decision Index +12 (DEC-S-64…75), Risk Register +15 (RISK-104…118). Module = logische Verantwortungsgrenzen (keine Microservices/Deployments); keine Technologie-/Deployment-Auswahl; keine ADR; Capability Matrix + Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.

Next planned WP: `CO-WP-009 – Human Identity, Workspaces, RBAC and Break Glass`.

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
