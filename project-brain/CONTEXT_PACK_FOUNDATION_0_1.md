# CoreOps – Context Pack: Foundation 0.1

Kompakter, wiederverwendbarer Kontext für die Foundation-Phase. Enthält keine Rohlogs, keine Chain-of-Thought, keine Secrets, keine privaten Daten und keine vollständige Kopie des Concept v3.0.

## Projekt und Phase

- **Projekt:** CoreOps — One Dashboard. Controlled Operations.
- **Phase:** `Foundation 0.1` (interner Phasenname; Release-Taxonomie: Foundation `v0.0.1-foundation`, Observe `v0.1.0-alpha.1`, beide `Proposed for acceptance`).

## NDF-Basis

- Nova-Development-Framework `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`).
- Branch `main` ist informativ, **nicht** normativ; Übernahme aus `main` nur via eigenes freigegebenes WP (CCR-04 geklärt).

## Akzeptierte Produktvision

Universelle, self-hosted und offline-fähige Operations Control Plane, die Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen modular verbindet. Herkunft: CoreOps Concept v3.0 (bereitgestellt 12. Juli 2026). Technische Detailentscheidungen sind unbestätigte Foundation-Kandidaten.

## Aktueller Scope

- **In Scope:** Foundation-Dokumentation, Sicherheitsgrundlagen, Architekturmodelle, Capability-/Supportgrenzen, ADR-Vorbereitung, Teststrategie, Readiness Review.
- **Out of Scope:** Anwendungscode, produktive Integrationen, Agents, Deployments, Scans, Netzwerkänderungen, Druckerverwaltung, Workflow-Ausführung, Secrets-Verarbeitung, produktive Installation.

## Letztes Work Package

`CO-WP-004 – Foundation Capability Matrix and Initial Support Boundary` (docs-only). Nach Bearbeitung im **Nova Review** (`pending`). Vorheriges WP: `CO-WP-003 – GO WITH NOTES`.

## Capability Matrix und Support Boundary (CO-WP-004)

- Foundation Capability Matrix (`docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md`): 74 Capabilities, 12 Domains; alle `not-implemented`/`not-supported` (Proposed for acceptance).
- Initiale Observe-Supportgrenze (`docs/integrations/INITIAL_SUPPORT_BOUNDARY.md`): Observe vollständig read-only, max. Integrationslevel 2; Map/Write ausgeschlossen.
- Drei getrennte Statusdimensionen (Roadmap/Implementation/Support), Herausgeberklassen und 21-Punkte-Support-Evidence-Satz definiert.
- `CCR-12` vorgeschlagen aufgelöst: Herstellernennung = Kandidat, kein Support.
- **Keine Runtime-Capability implementiert; keine Integration `supported`; keine technische Architektur Accepted; keine ADR.**

## Brief, Scope Lock und Release-Taxonomie (CO-WP-003)

- Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (`Proposed for acceptance`); Foundation `v0.0.1-foundation`, Observe `v0.1.0-alpha.1`.
- Docker-first eingeordnet; aktive Queue autoritativ; NDF-`main` und NDF-Level geklärt; Repository verifiziert.

## Concept-Registrierung

- Concept v3.0 vollständig registriert (57 Abschnitte); Decision Classification, Decision Index und Risk Register vorhanden.

## Lokaler NDF-Skills-Bestand

Vollständiger NDF-v1.0.0-Skills-Pack lokal unter `.claude/skills/` verfügbar (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`). Skills werden **pro Work Package selektiv** verwendet, nicht automatisch vollständig aktiviert. Provenance/Lock: `project-system/NDF_SKILLS_PROVENANCE.md`, `project-system/ndf-skills-lock.json`.

## Nächstes Work Package

`CO-WP-005 – Language Standard, Public Neutrality and Repository Governance` (vorgeschlagen; pending Nova review und Human-Maintainer-Freigabe).

## Aktuelle Blocker

- Keine harten Blocker. In CO-WP-003/004 adressiert (proposed/treatment-planned): Release-Taxonomie, Docker-first, Queue-Autorität, NDF-`main`, NDF-Level, Herstellersupport-Grenze (CCR-12), Observe-Scope, Legacy-Protokolle, Capability-Maturity, Drucktelemetrie.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09 sowie alle technischen ADR-Kandidaten.

## Zentrale Sicherheitsgrenzen

- Human-Maintainer-only für Freigabe, Staging, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen.
- Kein Code in Foundation.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Keine Übernahme privater Daten in Projektdateien.

## Akzeptierte ADRs

- **Keine.** Es existiert keine akzeptierte technische ADR.

## Bekannte Notes und Risiken

- NDF-Level-Semantik geklärt (CCR-03: Bootstrap-Status, Zahlenwert unverändert).
- Repository-URL im Manifest verifiziert und gesetzt: `https://github.com/KayKaspers/CoreOps`.
- Release-Taxonomie vorgeschlagen aufgelöst (CCR-02): `v0.0.1-foundation` / `v0.1.0-alpha.1`.
- Docker-first eingeordnet (CCR-10); aktive Queue-Autorität geklärt (CCR-11); NDF-`main` geklärt (CCR-04); Herstellersupport-Grenze vorgeschlagen aufgelöst (CCR-12).
- Capability Matrix (74 Capabilities) und Observe-Supportgrenze vorhanden; alle Capabilities `not-implemented`/`not-supported`.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09; 25 Foundation-Risiken (RISK-01…25, davon 10 `treatment-planned`).
- Keine akzeptierten technischen ADRs.

## Relevante Quelldateien

- [project-system/project-manifest.yaml](../project-system/project-manifest.yaml)
- [project-system/PROJECT_PROFILE.md](../project-system/PROJECT_PROFILE.md)
- [project-system/WORK_PACKAGE_QUEUE.md](../project-system/WORK_PACKAGE_QUEUE.md)
- [project-system/NEXT_PHASE.md](../project-system/NEXT_PHASE.md)
- [project-system/DECISION_INDEX.md](../project-system/DECISION_INDEX.md)
- [project-system/RISK_REGISTER.md](../project-system/RISK_REGISTER.md)
- [project-system/NDF_SKILLS_PROVENANCE.md](../project-system/NDF_SKILLS_PROVENANCE.md)
- [project-system/ndf-skills-lock.json](../project-system/ndf-skills-lock.json)
- [docs/architecture/COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md)
- [docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md)
- [docs/architecture/PROJECT_BRIEF.md](../docs/architecture/PROJECT_BRIEF.md)
- [docs/governance/FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md)
- [docs/governance/RELEASE_TAXONOMY.md](../docs/governance/RELEASE_TAXONOMY.md)
- [docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md)
- [docs/integrations/INITIAL_SUPPORT_BOUNDARY.md](../docs/integrations/INITIAL_SUPPORT_BOUNDARY.md)
- [project-brain/PROJECT_BRAIN.md](PROJECT_BRAIN.md)
- [ROADMAP.md](../ROADMAP.md)

## Kompakte Historie

- `CO-WP-001` (docs-only): Core Governance Skeleton erstellt (7 Dateien). Nova-Bewertung: `GO WITH NOTES` (read-only-Git-Note).
- `CO-WP-001A` (docs-only): Vollständiger NDF-v1.0.0-Skills-Pack lokal bereitgestellt (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`), Provenance + Lock erstellt. Nova-Bewertung: `GO`.
- `CO-WP-002` (docs-only): Nach fail-closed-Blocker und Source-Handoff Concept v3.0 vollständig registriert (57 Abschnitte), Decision Classification + Decision Index + Risk Register erstellt. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-003` (docs-only): Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (Proposed); CCR-02/04/10/11 aufgelöst, NDF-Level geklärt, Repository verifiziert. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004` (docs-only): Foundation Capability Matrix (74 Capabilities) und initiale Observe-Supportgrenze erstellt (Proposed); drei Statusdimensionen, Support-Evidence, CCR-12 vorgeschlagen aufgelöst. Keine Capability implementiert, keine Integration supported, keine ADR, kein Git Write. Nova Review ausstehend.
