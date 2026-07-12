# CoreOps – Context Pack: Foundation 0.1

Kompakter, wiederverwendbarer Kontext für die Foundation-Phase. Enthält keine Rohlogs, keine Chain-of-Thought, keine Secrets, keine privaten Daten und keine vollständige Kopie des Concept v3.0.

## Projekt und Phase

- **Projekt:** CoreOps — One Dashboard. Controlled Operations.
- **Phase:** `Foundation 0.1 – Platform Foundation` (vorläufiger Arbeitsname).

## NDF-Basis

- Nova-Development-Framework `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`).
- Branch `main` ist **nicht** normativ.

## Akzeptierte Produktvision

Universelle, self-hosted und offline-fähige Operations Control Plane, die Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen modular verbindet. Herkunft: CoreOps Concept v3.0 (bereitgestellt 12. Juli 2026). Technische Detailentscheidungen sind unbestätigte Foundation-Kandidaten.

## Aktueller Scope

- **In Scope:** Foundation-Dokumentation, Sicherheitsgrundlagen, Architekturmodelle, Capability-/Supportgrenzen, ADR-Vorbereitung, Teststrategie, Readiness Review.
- **Out of Scope:** Anwendungscode, produktive Integrationen, Agents, Deployments, Scans, Netzwerkänderungen, Druckerverwaltung, Workflow-Ausführung, Secrets-Verarbeitung, produktive Installation.

## Letztes Work Package

`CO-WP-002 – Concept v3.0 Registration and Decision Classification` (docs-only). Nach Bearbeitung im **Nova Review** (`pending`). Vorheriges WP: `CO-WP-001A – GO`.

## Concept-Registrierung

- Concept v3.0 vollständig registriert (alle 57 Abschnitte) unter `docs/architecture/COREOPS_CONCEPT_V3.md`; ein privater Standortname öffentlich neutralisiert.
- Decision Classification (`docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md`), initialer Decision Index (`project-system/DECISION_INDEX.md`) und initiales Risk Register (`project-system/RISK_REGISTER.md`) vorhanden.
- Technische Aussagen klassifiziert, **nicht** akzeptiert; keine ADR Accepted; keine ADR-Datei erzeugt.

## Lokaler NDF-Skills-Bestand

Vollständiger NDF-v1.0.0-Skills-Pack lokal unter `.claude/skills/` verfügbar (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`). Skills werden **pro Work Package selektiv** verwendet, nicht automatisch vollständig aktiviert. Provenance/Lock: `project-system/NDF_SKILLS_PROVENANCE.md`, `project-system/ndf-skills-lock.json`.

## Nächstes Work Package

`CO-WP-003 – Project Brief, Scope Lock and Release Taxonomy` (vorgeschlagen; pending Nova review und Human-Maintainer-Freigabe).

## Aktuelle Blocker

- Keine harten Blocker. Offene Punkte: NDF-Level-Ambiguität, unbekannter Repository-Status, offene Release-Taxonomie sowie 12 klassifizierte Konflikte (CCR-01…12) aus der Concept-Klassifikation.

## Zentrale Sicherheitsgrenzen

- Human-Maintainer-only für Freigabe, Staging, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen.
- Kein Code in Foundation.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Keine Übernahme privater Daten in Projektdateien.

## Akzeptierte ADRs

- **Keine.** Es existiert keine akzeptierte technische ADR.

## Bekannte Notes und Risiken

- NDF-Level-Ambiguität: Manifest `ndf_level: 1` vs. Starter-Vorlage `ndf_level: 2` (offen; CCR-03).
- Repository-Status `pending-human-maintainer` im Manifest (Remote origin: `github.com/KayKaspers/CoreOps.git`).
- Release-Taxonomie offen: Kollision `Foundation 0.1` vs. `Release 0.1 – Observe` (CCR-02 → `CO-WP-003`).
- Fehlende Capability Matrix (folgt `CO-WP-004`).
- 12 klassifizierte Konflikte (CCR-01…12) und 18 Foundation-Risiken (RISK-01…18) offen.
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
- [project-brain/PROJECT_BRAIN.md](PROJECT_BRAIN.md)
- [ROADMAP.md](../ROADMAP.md)

## Kompakte Historie

- `CO-WP-001` (docs-only): Core Governance Skeleton erstellt (7 Dateien). Nova-Bewertung: `GO WITH NOTES` (read-only-Git-Note).
- `CO-WP-001A` (docs-only): Vollständiger NDF-v1.0.0-Skills-Pack lokal bereitgestellt (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`), Provenance + Lock erstellt. Nova-Bewertung: `GO`.
- `CO-WP-002` (docs-only): Zunächst fail-closed geblockt (fehlende Quelle, `GO – Blocker bestätigt`); nach read-only-Source-Handoff Concept v3.0 vollständig registriert (57 Abschnitte), Decision Classification + Decision Index + Risk Register erstellt, Statusdokumente fortgeschrieben. Kein Code, keine ADR, kein Git Write. Nova Review ausstehend.
