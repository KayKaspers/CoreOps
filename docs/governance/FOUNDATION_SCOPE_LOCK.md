# CoreOps – Foundation 0.1 Scope Lock

> Status: **Proposed for acceptance**
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-003` (docs-only)

Bis zum Nova Review und Human-Maintainer-Commit ist dieser Scope Lock ein Vorschlag. Nach dem Human-Maintainer-Commit gilt er als bindende Projektgovernance.

## Purpose

Der Foundation Scope Lock grenzt verbindlich ab, was in Foundation 0.1 erarbeitet werden darf und was nicht. Er verhindert verdeckte Scope-Erweiterungen und stellt sicher, dass in der Foundation-Phase **kein produktiver Anwendungscode** entsteht.

## Binding Source Hierarchy

Bei Widersprüchen gilt diese Reihenfolge (höchste zuerst):

1. Human-Maintainer-Entscheidungen
2. akzeptierte ADRs
3. Foundation Scope Lock
4. Project Brief
5. aktive Work Package Queue (`project-system/WORK_PACKAGE_QUEUE.md`)
6. CoreOps Concept v3.0 (als Produktvision)
7. Decision Classification (als Analyseartefakt)
8. Roadmap (als Planungsartefakt)

Die NDF-v1.0.0-Regeln (Tag `v1.0.0`, Commit `9dcadc1`) bleiben darüber hinaus verbindliche Framework-Basis.

## Foundation In Scope

- Produkt- und Projektbrief
- Scope und Nicht-Ziele
- Rollen- und Governance-Modell
- Architecture Decision Preparation
- Architekturmodelle und Modulgrenzen
- Threat Model
- Trust Boundaries
- Identitätsmodelle
- Policy- und Approval-Modell
- Source of Truth und Provenance
- State- und Driftmodelle
- OIC und Domain-Pack-Governance
- API- und Eventmodelle
- Deployment-Blueprint-Modell
- Artifact-Trust-Modell
- Offline- und Air-Gap-Baseline
- Topologiemodell
- Telemetrieschema
- Datenklassifizierung
- Retention und Redaction
- Self-Protection und Recovery-Konzept
- UX-Informationsarchitektur
- Teststrategie und Integration Lab
- Risk Register
- Decision Index
- ADR-Kandidaten und akzeptierte ADRs
- Foundation Readiness Review
- Foundation Release Preparation

## Foundation Out of Scope

- produktiver Anwendungscode
- Frontend-Implementierung
- Backend-Implementierung
- Datenbankschema als ausführbare Migration
- Agent- oder Relay-Implementierung
- Discovery Scans
- Netzwerkzugriffe auf reale Zielsysteme
- SNMP-Abfragen realer Geräte
- Druckersteuerung
- Deployment-Ausführung
- Skriptausführung
- Workflow Runtime
- Secrets-Verarbeitung
- produktive Authentifizierung
- produktive PKI
- produktive Docker-Installation
- CI-Build einer Anwendung
- produktive Integrationen
- Firmware- oder Netzwerkänderungen
- produktive Releases
- produktive Deployments

## Allowed Artifact Types

- Markdown-Dokumentation (Briefs, Modelle, Governance, Security-Konzepte)
- nicht ausführbare Schemakandidaten und Beispiel-Snippets zur Illustration
- Pseudocode und Diagrammbeschreibungen
- Tabellen, Register und Indizes (Decision Index, Risk Register)

Dokumentationsbeispiele, Pseudocode, Schemakandidaten und nicht ausführbare Illustrationen bleiben zulässig, sofern sie **keine** technische Entscheidung vortäuschen.

## Forbidden Implementation Types

- ausführbarer Anwendungs-, Agent- oder Relay-Code
- ausführbare Migrationen, Build-/CI-Pipelines, Container-Builds
- Lockfiles von Package Managern, Dependencies
- Skripte mit realer Wirkung, Netzwerk-/Geräteinteraktion
- ADR-Dateien innerhalb dieses und der vorherigen Bootstrap-Work-Packages, sofern nicht ausdrücklich freigegeben

## Change Control

Änderungen am Scope Lock benötigen:

- ein eigenes Work Package,
- Nova Review,
- Human-Maintainer-Freigabe,
- ein Project-Brain- und Context-Pack-Update.

## Exit Gates

Foundation 0.1 darf erst als releasebereit bewertet werden, wenn mindestens die 24 Foundation Exit Gates erfüllt sind (Project Brief, Scope Lock, Release-Taxonomie, konsistente Queue, Threat Model, Trust Boundaries, Security-Invarianten, relevante ADRs durch Human Maintainer entschieden, SoT-Modell, State-/Driftmodell, OIC v0.1, Policy-/Approval-Modell, Machine-Identity-Modell, Deployment-Blueprint-Modell, Artifact-Trust-Modell, Offline-/Air-Gap-Baseline, Topologie-Evidence-Modell, Datenklassifizierung/Retention, Teststrategie, konsistenter Decision Index, Risk Register ohne unbehandelten Foundation-Blocker, Cross-Document Consistency Review, Foundation Readiness Review, separate Release Preparation).

> **Keines** dieser Gates ist in diesem Work Package als erfüllt markiert. Sie bleiben offen.

## Known Exceptions

Zu Beginn **keine** stillen Ausnahmen.

## Superseded Planning Material

Die im CoreOps Concept v3.0 (§50) enthaltene Queue mit `CO-WP-001` bis `CO-WP-021` ist ein **historischer initialer Planungsvorschlag**. Die aktive, verbindliche Foundation-Queue ist ausschließlich `project-system/WORK_PACKAGE_QUEUE.md` (aktuell `CO-WP-001` bis `CO-WP-031` sowie `CO-WP-001A`). Das Concept wird dadurch nicht rückwirkend umgeschrieben; die Abweichung ist hier erklärt (Auflösung von CCR-11).
