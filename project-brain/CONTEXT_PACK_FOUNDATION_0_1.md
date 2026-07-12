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

`CO-WP-001 – NDF Project Bootstrap – Core Governance Skeleton` (docs-only). Nach Bearbeitung im **Nova Review**.

## Nächstes Work Package

`CO-WP-002 – Concept v3.0 Registration and Decision Classification` (vorgeschlagen; pending Nova review und Human-Maintainer-Freigabe).

## Aktuelle Blocker

- Keine harten Blocker. Offene Punkte: NDF-Level-Ambiguität, unbekannter Repository-Status, offene Release-Taxonomie, noch ausstehende vollständige Concept-v3.0-Übernahme.

## Zentrale Sicherheitsgrenzen

- Human-Maintainer-only für Freigabe, Staging, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen.
- Kein Code in Foundation.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Keine Übernahme privater Daten in Projektdateien.

## Akzeptierte ADRs

- **Keine.** Es existiert keine akzeptierte technische ADR.

## Bekannte Notes und Risiken

- NDF-Level-Ambiguität: Manifest `ndf_level: 1` vs. Starter-Vorlage `ndf_level: 2` (offen, nicht aufgelöst; nur Bootstrap-Status).
- Repository-Status `pending-human-maintainer` (keine erfundene URL).
- Release-Taxonomie offen: Kollision `Foundation 0.1` vs. `Release 0.1 – Observe` (→ `CO-WP-003`).
- Concept v3.0 noch nicht vollständig ins Repository übernommen.
- Fehlende Capability Matrix (folgt `CO-WP-004`).
- Keine akzeptierten technischen ADRs.

## Relevante Quelldateien

- [project-system/project-manifest.yaml](../project-system/project-manifest.yaml)
- [project-system/PROJECT_PROFILE.md](../project-system/PROJECT_PROFILE.md)
- [project-system/WORK_PACKAGE_QUEUE.md](../project-system/WORK_PACKAGE_QUEUE.md)
- [project-system/NEXT_PHASE.md](../project-system/NEXT_PHASE.md)
- [project-brain/PROJECT_BRAIN.md](PROJECT_BRAIN.md)
- [ROADMAP.md](../ROADMAP.md)

## Kompakte Historie

- `CO-WP-001` (docs-only): Core Governance Skeleton erstellt (7 Dateien). Kein Code, keine ADR, keine Technologieauswahl. Umsetzung abgeschlossen; Nova Review ausstehend.
