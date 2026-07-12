# CoreOps – Project Profile

> **Slogan:** One Dashboard. Controlled Operations.

**Status:** Foundation – Initial Bootstrap
**NDF-Basis:** Nova-Development-Framework `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — der Branch `main` ist **nicht** normativ.
**Erzeugt durch:** Work Package `CO-WP-001` (docs-only, Core Governance Skeleton)

---

## 1. Projektname und Slogan

- **Projektname:** CoreOps
- **Slogan:** One Dashboard. Controlled Operations.

## 2. Kurzbeschreibung

CoreOps ist eine universelle, self-hosted und offline-fähige Operations Control Plane für Anwendungen, Deployments, Server, Virtualisierung, Container, Netzwerke, Drucker und verteilte Infrastruktur.

## 3. Zu lösendes Problem

Betriebsteams verwalten heterogene Infrastruktur über viele getrennte Werkzeuge, ohne eine gemeinsame Source of Truth, ohne durchgängige Nachvollziehbarkeit von Änderungen und ohne konsistente Policy- und Freigabekontrolle. CoreOps adressiert diese Fragmentierung durch eine einheitliche, kontrollierte Oberfläche.

## 4. Produktvision

CoreOps verbindet **Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen** in einer modularen Oberfläche. Details siehe Herkunft in Abschnitt 17.

## 5. Zielgruppen

- Betreiber self-hosted Infrastruktur
- Plattform- und Operations-Teams
- Administratoren verteilter Systeme (Server, Container, Netzwerke, Drucker)
- Organisationen mit Anforderungen an Offline-/Air-Gapped-Betrieb

## 6. Allgemeiner Einsatzbereich

Zentralisierte, kontrollierte Steuerung und Beobachtung heterogener Infrastruktur — von Einzelservern bis zu verteilten, teils netzgetrennten Umgebungen.

## 7. Langfristige Kernfähigkeiten

- Source of Truth mit Field Provenance
- Observed / Desired / Effective State und Drift-Erkennung
- Topologie-Graph mit Evidenz und manueller Autorität
- Deployment Control Plane
- Policy, Freigabe und Ausführungsautorisierung
- Vertrauenswürdige Automatisierung
- Audit-, Evidence- und Telemetriemodell
- Modulare Domain Packs
- Restricted / Isolated / Air-Gapped Betrieb

## 8. Klare frühe Nicht-Ziele

- Kein Anwendungscode in der Foundation-Phase
- Keine produktiven Integrationen
- Keine autonomen Agents
- Keine Deployments, Scans oder Netzwerkänderungen durch das Framework
- Keine Druckerverwaltung im Betrieb
- Keine Workflow-Ausführung
- Keine Secrets-Verarbeitung
- Keine produktive Installation

## 9. Foundation-Status

Frühester Bootstrap. Es existiert ausschließlich das Core Governance Skeleton. Weder Architektur noch Technologie-Stack sind verbindlich. Keine ADR ist Accepted. Die Foundation-Queue ist noch nicht vollständig ausgeführt.

## 10. NDF-v1.0.0-Basis

Normativ ist ausschließlich `Nova-Development-Framework` am Tag `v1.0.0` / Commit `9dcadc1`. Der veränderliche Branch `main` ist nicht normativ und darf in diesem Work Package weder übernommen noch als verbindlich behandelt werden.

## 11. Rollenmodell

- **Owner / Human Maintainer:** entscheidet allein über Freigabe, Staging, Commit, Push, Merge, Tags, Releases, produktive Deployments und irreversible/privilegierte Aktionen.
- **Architecture Lead / Nova:** Planung, Architektur, Review, Bewertung von Rückmeldungen.
- **Implementation Assistant / Claude:** Umsetzung genau eines freigegebenen Work Packages, ohne eigenständige Freigabe- oder Bewertungsentscheidung.

## 12. Akzeptierte Prozess- und Sicherheitsgrenzen

- NDF-Prozess mit Human-Maintainer-Gates ist bindend.
- Kein Code in der Foundation-Phase.
- Keine autonomen Git- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Keine Übernahme privater Daten in öffentliche Projektdateien.

## 13. Technische Kandidaten (ohne Akzeptanzstatus)

Siehe Abschnitt **Foundation Candidates** unten. Kein Eintrag ist beschlossen.

## 14. Bekannte Initialrisiken

- NDF-Level-Ambiguität (Level 1 vs. Starter-Vorlage Level 2), siehe Notes.
- Unbekannter Repository-Status (`pending-human-maintainer`).
- Offene Release-Taxonomie (Kollision `Foundation 0.1` vs. `Release 0.1 – Observe`).
- Concept v3.0 noch nicht vollständig ins Repository übernommen.
- Breiter Produktscope mit hohem Sicherheitsanspruch (Policy, Air-Gap, Secrets).

## 15. Offene Foundation-Fragen

- Welche NDF-Level-Vorlage ist verbindlich?
- Wird `docker-first` als Delivery Baseline bestätigt?
- Wie lautet die finale Release-Taxonomie (→ `CO-WP-003`)?
- Welcher Technologie-Stack wird per ADR entschieden?

## 16. Nächster Meilenstein

`CO-WP-002 – Concept v3.0 Registration and Decision Classification` (pending Nova review and Human Maintainer approval).

## 17. Herkunft des Concept v3.0

CoreOps Concept v3.0 – Platform Foundation, durch den Human Maintainer am 12. Juli 2026 als Produktvision und Ausgangsbasis bereitgestellt. Technische Detailentscheidungen gelten weiterhin als unbestätigte Foundation-Kandidaten.

---

## Accepted Product Vision

Produktziel: eine einheitliche, kontrollierte Operations Control Plane, die Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen modular zusammenführt. Gewünschter Nutzen: weniger Werkzeugfragmentierung, durchgängige Nachvollziehbarkeit, kontrollierte und policy-gebundene Operationen — auch offline/air-gapped.

## Binding Project Governance

- NDF-v1.0.0-Prozess ist verbindlich (Tag `v1.0.0`, Commit `9dcadc1`; `main` nicht normativ).
- Human-Maintainer-Gates für alle irreversiblen und privilegierten Aktionen.
- Kein Code in der Foundation-Phase.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen durch den Implementation Assistant.
- Rückmeldung an Nova; Bewertung (`GO`, `GO WITH NOTES`, `REWORK`, `SPLIT`, `STOP`) trifft ausschließlich Nova.

## Foundation Candidates

> **Hinweis:** Diese Kandidaten sind **nicht beschlossen** und dienen nur als Diskussionsbasis für spätere ADRs.

- modularer Monolith
- Next.js
- Fastify
- PostgreSQL
- Redis
- Go-Agent
- Prometheus
- Grafana
- Loki
- OpenTelemetry
- Open Policy Agent
- Temporal
- PostgreSQL-basierte Topologie

## ADR Required

Folgende Themen benötigen vor jeder verbindlichen technischen Festlegung einen Vergleich und eine Human-Maintainer-Entscheidung:

- Architekturform (modularer Monolith vs. Alternativen)
- Frontend-Framework
- Backend-/API-Framework
- Primäre Datenhaltung und Topologie-Persistenz
- Caching-/Queue-Schicht
- Agent-Technologie und Enrollment
- Observability-Stack (Metriken, Logs, Traces)
- Policy-Engine
- Workflow-/Automatisierungs-Engine
- Delivery Baseline (`docker-first`)

---

## CO-WP-003 Konsolidierung

- **Project Brief:** [docs/architecture/PROJECT_BRIEF.md](../docs/architecture/PROJECT_BRIEF.md) (Proposed for acceptance).
- **Foundation Scope Lock:** [docs/governance/FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md) (Proposed for acceptance).
- **Release-Taxonomie:** [docs/governance/RELEASE_TAXONOMY.md](../docs/governance/RELEASE_TAXONOMY.md) — Foundation-Tag-Kandidat `v0.0.1-foundation`, erster Observe-Prerelease-Kandidat `v0.1.0-alpha.1` (Proposed for acceptance; kein Tag/Release erzeugt).
- **Docker-first-Einordnung:** akzeptierte Delivery-/Betriebsanforderung für die zentrale Plattform (dokumentierte Docker-/Compose-Standardinstallation, reproduzierbare self-hosted Bereitstellung ohne Cloudpflicht). **Nicht** implementiert/getestet; bestimmt **keine** interne Anwendungsarchitektur; Kubernetes ist keine Voraussetzung; Agents/Relays dürfen native Binärdateien sein.
- **Aktive Work-Package-Queue:** ausschließlich [WORK_PACKAGE_QUEUE.md](WORK_PACKAGE_QUEUE.md) (`CO-WP-001…031` + `CO-WP-001A`). Die Concept-§50-Queue (21 WPs) ist historischer Vorschlag.
- **NDF-`main`-Klarstellung:** `v1.0.0` / `9dcadc1` normativ; `main` ausschließlich informativ; jede Übernahme aus `main` benötigt ein eigenes freigegebenes Work Package (CCR-04 geklärt, keine ADR).
- **Repository-Status:** verifiziert und gesetzt — `https://github.com/KayKaspers/CoreOps` (origin-Clone-Remote endet auf `.git`).

## Notes

- **NDF-Level-Semantik:** `ndf_level: 1` bedeutet ausschließlich „initialer NDF-Projektbootstrap" — keine vollständige NDF-Konformität, technische Reife, Releasebereitschaft, Sicherheitsfreigabe oder erreichte Produktstufe. Der Zahlenwert bleibt unverändert. Die Starter-Vorlage mit `ndf_level: 2` bleibt dokumentierte NDF-Quellambiguität, ist aber kein Blocker (CCR-03).
- **Skills-first:** Lokaler NDF-v1.0.0-Skills-Pack seit CO-WP-001A vorhanden; pro Work Package selektiv genutzt.
