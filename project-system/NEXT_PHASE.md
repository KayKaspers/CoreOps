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

Current implemented WP: `CO-WP-004B3 – Record Completed NDF Intake Transfer` – pending Nova review.
Transfer Package 001 wurde dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e).
Alle 7 Kandidaten von `approved-for-transfer` auf `transferred-to-ndf` gesetzt; NDF Work Package: `NDF-INTAKE-COREOPS-001` für alle 7; Adoption Version bleibt leer.
Kein Kandidat adoptiert; NDF-Repository unverändert.

Next planned CoreOps WP: `CO-WP-004C – BSI and Public-Sector Readiness Baseline` (planned-next; pending Nova review und Human-Maintainer-Freigabe).

External NDF status: Transfer Package 001 Intake abgeschlossen; Adoption ausstehend.

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
