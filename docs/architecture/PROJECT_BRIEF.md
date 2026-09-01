# CoreOps – Project Brief

> Document Status: Accepted
> Decision Type: Product and Foundation Scope
> Technical Architecture Status: Unconfirmed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ

Erzeugt durch `CO-WP-003 – Project Brief, Foundation Scope Lock and Release Taxonomy` (docs-only). Keine Technologie ist als gewählt dargestellt; keine technische Architektur ist akzeptiert.

## 1. Projektname und Slogan

- **Projektname:** CoreOps
- **Slogan:** One Dashboard. Controlled Operations.

## 2. Problem Statement

Betriebsteams verwalten heterogene Infrastruktur (Anwendungen, Server, Virtualisierung, Container, Netzwerke, Drucker, verteilte Standorte) über viele getrennte Werkzeuge — ohne gemeinsame Source of Truth, ohne durchgängige Nachvollziehbarkeit von Änderungen und ohne konsistente Policy- und Freigabekontrolle. Das erzeugt blinde Flecken, riskante Änderungen und schwer auditierbare Abläufe, insbesondere in self-hosted, eingeschränkten oder air-gapped Umgebungen.

## 3. Produktvision

CoreOps ist eine universelle, self-hosted und offline-fähige Operations Control Plane, die **Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen** in einer modularen Oberfläche verbindet. Vollständige Produktvision: [COREOPS_CONCEPT_V3.md](COREOPS_CONCEPT_V3.md) (Accepted Product Vision).

## 4. Zielgruppen

Homelabs, Selfhoster, Vereine, kleine und mittlere Unternehmen, Bildungseinrichtungen, Entwickler, Content Creator, IT-Abteilungen, verteilte Standorte sowie eingeschränkte/Offline-/Air-Gap-Umgebungen; später MSPs und Rechenzentren. Core-Produkte (z. B. SpeakCore, CastCore, AirCore, OrgaCore) sind hochwertige First-Party-Integrationen, aber weder Voraussetzung noch alleiniger Mittelpunkt.

## 5. Zentrale Nutzerprobleme

- Kein einheitlicher Überblick über Systeme, Standorte, Zuordnungen und Abhängigkeiten.
- Unklare Ist-/Sollzustände und unbemerkter Drift.
- Riskante, schlecht nachvollziehbare Änderungen ohne Policy und Freigabe.
- Fragmentierte Hersteller-/Protokolllandschaft.
- Fehlende Offline-/Air-Gap-Betriebsfähigkeit vieler bestehender Werkzeuge.

## 6. Erwarteter Produktnutzen

- Eine kontrollierte Oberfläche für Beobachtung und Steuerung heterogener Infrastruktur.
- Durchgängige Nachvollziehbarkeit (Audit, Evidence, Provenance).
- Policy- und freigabegebundene, fail-closed Operationen.
- Self-hosted, offline-/air-gap-fähiger Betrieb ohne Cloudpflicht.

## 7. Zentrale Produktprinzipien

Documentation First · Architecture First · Security First · Privacy by Design · Fail Closed · AI creates, Humans approve · Read-only vor Write · Preview vor Execute · Plan vor Deployment · Backup vor gefährlicher Änderung · Verifikation nach jeder Änderung · Auditierbarkeit · Offline First · Public Neutrality.

## 8. Akzeptierte Produktanforderungen

Aus der akzeptierten Produktvision (siehe [DECISION_INDEX.md](../../project-system/DECISION_INDEX.md), `DEC-P-*`):

- Universelle, herstellerunabhängige Operations Control Plane (nicht Core-exklusiv).
- Self-hosted, offline-/Air-Gap-fähig, keine Cloudpflicht für Kernfunktionen.
- Zweisprachigkeit DE/EN ab Foundation (Deutsch als möglicher Default).
- Produktformel *SoT + Observed + Desired + Policy + Trusted Execution + Verification + Audit*; Read-only vor Write.
- KI nur beratend, kein autonomer Administrator.
- Fehlende Daten dürfen nicht automatisch als gesund gewertet werden.
- Strikte Trennung Observed / Desired / Effective State und Drift (kein `Enforce` im MVP).
- Privacy by Design (Datenklassen, Redaction, Retention).
- **Docker-first** als Delivery- und Betriebsanforderung für die zentrale Plattform (Einordnung siehe [FOUNDATION_SCOPE_LOCK.md](../governance/FOUNDATION_SCOPE_LOCK.md) und [RELEASE_TAXONOMY.md](../governance/RELEASE_TAXONOMY.md); noch nicht implementiert/getestet).

## 9. Frühe Produktgrenzen

CoreOps baut keine vollständigen Fremdsysteme nach, sondern integriert mit ihnen. Schreibende und privilegierte Aktionen sind stets policy- und freigabegebunden. In frühen Phasen dominiert Read-only.

## 10. Nicht-Ziele

Kein vollständiges CI-Build-System, SIEM, ITSM/Helpdesk, EDR/Antivirus, keine allgemeine Remote-Desktop-Plattform, keine vollständige PKI als Ersatz, kein vollständiger Low-Code-Builder, keine universelle Cloud-Management-Suite, kein autonomer KI-Administrator, kein vollständiger Ersatz jedes Herstellerinterfaces, keine vollständige ERP-/Lizenz-/Vertragsverwaltung, kein `Enforce`-Auto-Reconcile im MVP. Vollständige Liste: [DECISION_INDEX.md](../../project-system/DECISION_INDEX.md) (`DEC-N-*`).

## 11. Foundation-Ziel

Foundation 0.1 liefert ausschließlich Dokumentation, Architekturmodelle, Sicherheitsgrundlagen, Governance und Verträge — **keinen produktiven Anwendungscode**. Verbindliche Abgrenzung: [FOUNDATION_SCOPE_LOCK.md](../governance/FOUNDATION_SCOPE_LOCK.md).

## 12. Langfristige Produktmeilensteine

Foundation → Observe → Map → Plan → Deploy → Automate → Extend → Scale. Meilensteinnamen und Versions-Taxonomie: [RELEASE_TAXONOMY.md](../governance/RELEASE_TAXONOMY.md). Die Foundation-/Observe-Release-Taxonomie ist für Foundation 0.1 `Accepted` (`HM-1`); `v0.0.1-foundation` bleibt ein Foundation-**Tag-Kandidat** bis zu einer separaten Autorisierung. Die endgültige Versionsnummerierung der Meilensteine nach `Observe` bleibt offen und vorläufig. Diese Klarstellung autorisiert **keinen** Tag und **kein** Release.

## 13. Erfolgskriterien der Foundation

- Prüfbarer Project Brief, akzeptierter Scope Lock, eindeutige Release-Taxonomie.
- Konsistente, autoritative Work-Package-Queue.
- Threat Model, Trust Boundaries und Security-Invarianten dokumentiert.
- Kern-Foundation-Modelle (SoT, State/Drift, Policy, OIC, Machine Identity, Deployment-Blueprint, Artifact Trust, Offline/Air-Gap, Topologie-Evidence, Datenklassifizierung, Teststrategie) vorhanden.
- Decision Index und Risk Register konsistent, ohne unbehandelten Foundation-Blocker.
- Cross-Document Consistency Review und Foundation Readiness Review abgeschlossen.

Diese Kriterien sind **nicht** als erfüllt markiert; sie sind die Exit-Bedingungen der Foundation-Phase.

## 14. Offene technische Entscheidungen

Architektur (modularer Monolith, Plane-Modell), Frontend, Backend/API, Datenhaltung, Topologie-Persistenz, Caching/Queue, Agent-Technologie, Observability-Stack, Policy-Engine, Workflow-Engine sowie die konkrete Umsetzung der Docker-first-Delivery bleiben **offen** und ADR-pflichtig. Es ist keine Technologie ausgewählt. Siehe [CONCEPT_DECISION_CLASSIFICATION.md](CONCEPT_DECISION_CLASSIFICATION.md) (AC/TC/ADR).

## 15. Verweise

- Produktvision: [COREOPS_CONCEPT_V3.md](COREOPS_CONCEPT_V3.md)
- Klassifikation: [CONCEPT_DECISION_CLASSIFICATION.md](CONCEPT_DECISION_CLASSIFICATION.md)
- Scope Lock: [FOUNDATION_SCOPE_LOCK.md](../governance/FOUNDATION_SCOPE_LOCK.md)
- Release-Taxonomie: [RELEASE_TAXONOMY.md](../governance/RELEASE_TAXONOMY.md)
- Decision Index: [DECISION_INDEX.md](../../project-system/DECISION_INDEX.md)
- Risk Register: [RISK_REGISTER.md](../../project-system/RISK_REGISTER.md)
