# CoreOps Concept v3.0 – Decision Classification

> Companion zu [COREOPS_CONCEPT_V3.md](COREOPS_CONCEPT_V3.md)
> NDF-Basis: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` nicht normativ
> Erzeugt durch: `CO-WP-002 – Concept v3.0 Registration and Decision Classification`

## Purpose

Dieses Dokument ordnet die wesentlichen Aussagen des CoreOps Concept v3.0 systematisch definierten Klassen zu. Es trennt akzeptierte Produktanforderungen von unbestätigten technischen Kandidaten, kennzeichnet bindende Governance-Grenzen, identifiziert ADR-pflichtige Themen und macht Konflikte und Klärungsbedarf sichtbar.

Abgrenzung:

- Dieses Dokument **entscheidet keine** technische Architektur.
- Es **akzeptiert keine** Technologie.
- Es **erzeugt keine** ADR und markiert keine ADR als Accepted.
- Die Klassifikation erfolgt überwiegend auf Abschnittsebene; eine Zeile pro Einzelsatz wäre künstlich zu kleinteilig.

## Classification Model

| Klasse | Bedeutung |
| ------ | --------- |
| `APC` – Accepted Product Constraint | Durch den Human Maintainer akzeptierte produktbezogene Anforderung (Produktvision). |
| `BG` – Binding Governance | Bindende NDF-, Rollen-, Freigabe- oder Prozessregel. |
| `FR` – Foundation Requirement | Thema, das während Foundation 0.1 geklärt oder dokumentiert werden muss. |
| `AC` – Architecture Candidate | Vorgeschlagene, noch nicht akzeptierte Architekturentscheidung. |
| `TC` – Technology Candidate | Konkrete, noch zu evaluierende Technologie oder Bibliothek. |
| `ADR` – ADR Required | Thema, das vor verbindlicher Festlegung einen dokumentierten Optionsvergleich und eine Human-Maintainer-Entscheidung benötigt. |
| `RC` – Roadmap Candidate | Geplanter späterer Funktionsumfang, noch nicht scope-locked. |
| `NG` – Non-Goal | Explizit ausgeschlossener oder bewusst nicht angestrebter Funktionsumfang. |
| `DF` – Deferred | Bewusst später zu behandelndes Thema. |
| `CCR` – Conflict or Clarification Required | Widerspruch, Mehrdeutigkeit, unvollständige Definition oder notwendige Human-Maintainer-Klärung. |

Jede Aussage besitzt genau **eine Primärklasse** und optional zusätzliche Sekundärklassen.

## Accepted Product Constraints

| ID | Concept Section | Statement | Primary | Secondary | Rationale | Foundation Follow-up |
| -- | --------------- | --------- | ------- | --------- | --------- | -------------------- |
| APC-01 | 3, 5 | Universelle, herstellerunabhängige Operations Control Plane, nicht exklusiv für Core-Produkte | APC | — | Ausdrücklich akzeptierte Produktvision | CO-WP-003 Scope Lock |
| APC-02 | 3, 9.5, 38 | Vollständig self-hosted, offline- und Air-Gap-fähig; keine Cloudpflicht für Kernfunktionen | APC | FR | Kernanforderung der Vision | CO-WP-023 |
| APC-03 | 3, 53 | Zweisprachigkeit DE/EN ab Foundation; Deutsch als mögliche Standardsprache | APC | — | Produktanforderung | CO-WP-005 |
| APC-04 | 8, 9.3 | Produktformel SoT + Observed + Desired + Policy + Trusted Execution + Verification + Audit; Read-only vor Write | APC | BG | Zentrales Produkt- und Sicherheitsziel | CO-WP-011/012/013 |
| APC-05 | 7, 45 | Kein autonomer KI-Administrator; KI nur beratend | APC | BG, NG | Sicherheits- und Produktgrenze | CO-WP-013 |
| APC-06 | 18, 26, 54 | Herstellerlisten sind kein bestätigter Support, keine Partnerschaft, keine Zertifizierung | APC | CCR | Ehrlichkeit gegenüber Nutzern | CO-WP-004 Support Boundary |
| APC-07 | 25 | Fehlende Daten dürfen nicht automatisch als gesund gewertet werden | APC | FR | Vermeidung falscher Health-Anzeigen | CO-WP-004 |
| APC-08 | 20 | Strikte Trennung Observed / Desired / Effective State und Drift; `Enforce` kein MVP-Ziel | APC | AC, FR | Kernmodell der Plattform | CO-WP-012 |
| APC-09 | 41 | Datenschutz by Design: Datenklassen, Redaction, Retention | APC | FR, BG | Privacy-Anforderung | CO-WP-025 |
| APC-10 | 42, 43 | Self-Protection, Degraded Modes, Recovery Mode, Self-Dependency-Schutz | APC | FR | Betriebssicherheit | CO-WP-026 |

## Binding Governance

| ID | Concept Section | Statement | Primary | Secondary |
| -- | --------------- | --------- | ------- | --------- |
| BG-01 | 1 | Human Maintainer entscheidet allein über Freigabe, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen | BG | — |
| BG-02 | 1, 2 | Ein kohärentes Work Package pro Commit; keine verdeckten Scope-Erweiterungen | BG | — |
| BG-03 | 1 | Nova plant/reviewt und bewertet (GO/GO WITH NOTES/REWORK/SPLIT/STOP); Implementation Agent setzt nur um, committet/pusht/taggt nicht | BG | — |
| BG-04 | 2, 9.6 | Fail Closed; Read-only vor Write; Preview vor Execute; Plan vor Deployment; Backup vor gefährlicher Änderung; Verifikation nach jeder Änderung | BG | APC |
| BG-05 | 2 | Skills-first Operating Mode; Skills ohne Netzwerk-, Secret-, Ausführungs-, Git- oder Release-Autonomie | BG | — |
| BG-06 | 2, 52 | Public Neutrality für öffentliche Repository-Inhalte; keine Secrets in Logs/Prompts/Exports | BG | APC |
| BG-07 | 12, 16 | Control Plane darf nicht durch Plugins oder Agents umgangen werden | BG | AC |
| BG-08 | 52 | 20 projektweite Sicherheitsinvarianten gelten verbindlich als Zielbild | BG | FR |

## Foundation Requirements

Themen, die in Foundation 0.1 dokumentiert oder entschieden werden müssen (Auswahl, nach Concept-Abschnitt):

- FR: Plane-Architektur und Modulgrenzen (§10–16) → CO-WP-008
- FR: Source of Truth und Field Provenance (§19) → CO-WP-011
- FR: Observed/Desired/Effective State und Drift (§20) → CO-WP-012
- FR: Policy-, Approval- und Ausführungsautorisierung (§21) → CO-WP-013
- FR: OIC Integration Contract v0.1 (§17, 18) → CO-WP-014
- FR: Threat Model und Trust Boundaries (§15, 52) → CO-WP-007
- FR: Machine Identity, Enrollment, Offline Credential Lifecycle (§30, 31) → CO-WP-010
- FR: Artifact Trust, SBOM, Provenance, Revocation (§23) → CO-WP-022
- FR: Offline/Isolated/Air-Gapped Operation und CorePack (§38) → CO-WP-023
- FR: Secrets, Configuration Vault, Key Custody (§37, 40) → CO-WP-024
- FR: Data Classification, Retention, Redaction (§41) → CO-WP-025
- FR: Topology Graph, Evidence, Manual Authority (§28) → CO-WP-020
- FR: Event, Audit Correlation, Evidence Model (§34, 35) → CO-WP-018
- FR: Telemetry und Normalization Schema (§24) → CO-WP-019
- FR: Deployment Control Plane und Blueprint Schema (§22) → CO-WP-021
- FR: Domain Pack Governance, Support Levels (§16, 18) → CO-WP-015
- FR: Self-Protection, Degraded Modes, Recovery (§42, 43) → CO-WP-026
- FR: UX Information Architecture und Dashboard System (§11) → CO-WP-027
- FR: Test Strategy, Fixtures, Integration Lab (§44) → CO-WP-028

## Architecture Candidates

> Noch nicht akzeptiert. Jeweils ADR-pflichtig.

| ID | Concept Section | Candidate |
| -- | --------------- | --------- |
| AC-01 | 9.1, 46 | Modularer Monolith mit separat betreibbaren Komponenten |
| AC-02 | 10 | Sechs-Ebenen-Plane-Architektur (Experience/Control/Data/Execution/Trust + Domain Packs) |
| AC-03 | 9.2 | API-first mit versionierten APIs |
| AC-04 | 30, 31 | Ausgehende mTLS-Agent-/Relay-Verbindungen; kurzlebige Machine Identities |
| AC-05 | 22 | Trennung von Build und Deployment; unveränderliche Artefakte |
| AC-06 | 46 | PostgreSQL-basierte Topologie mit vorberechneten Graphansichten statt Graphdatenbank im MVP |
| AC-07 | 20 | Drift Detection vor Reconciliation; Betriebsmodi Observe→Recommend→Approve |
| AC-08 | 19 | Feldbasiertes Source-of-Truth-Modell mit Prioritäts-/Konfliktregeln |

## Technology Candidates

> Reine Kandidaten (§46 und weitere). Keiner ist ausgewählt oder akzeptiert.

| ID | Concept Section | Candidate | Bereich |
| -- | --------------- | --------- | ------- |
| TC-01 | 46 | Next.js, React, TypeScript, Tailwind CSS, PWA | Frontend |
| TC-02 | 46 | Node.js, TypeScript, Fastify, OpenAPI | Backend |
| TC-03 | 46 | PostgreSQL, Redis, optional S3-kompatibler Objektspeicher | Datenhaltung |
| TC-04 | 30, 46 | Go | Agent/Relay |
| TC-05 | 24, 46 | Prometheus, Grafana, OpenTelemetry Collector, Loki, Alertmanager | Observability |
| TC-06 | 21, 46 | Open Policy Agent | Policy |
| TC-07 | 33, 46 | Temporal (oder eigene persistente Job Engine) | Workflows |
| TC-08 | 23 | Sigstore/Cosign, CycloneDX, SLSA, TUF-ähnlich | Supply Chain |
| TC-09 | 19 | NetBox (optional) | Inventar/IPAM |
| TC-10 | 40 | HashiCorp Vault, Bitwarden Secrets Manager (optional, später) | Secrets Backends |

## ADR-Required Topics

> Nummerierte Kandidatenliste. **Kein Accepted-Status.** Es wurden keine ADR-Dateien erzeugt. Quelle: Concept §51 (30 Kandidaten) plus zwei Foundation-Klärungen.

1. ADR-0001 CoreOps als universelle Operations Control Plane
2. ADR-0002 Modularer Monolith
3. ADR-0003 Read-only First
4. ADR-0004 Offline First
5. ADR-0005 Plane-Architektur
6. ADR-0006 CoreOps Integration Contract
7. ADR-0007 Standards First
8. ADR-0008 Native Widgets vor allgemeinen iFrames
9. ADR-0009 Source of Truth und feldbasierte Provenance
10. ADR-0010 Desired State getrennt von Observed State
11. ADR-0011 Drift Detection vor Reconciliation
12. ADR-0012 Zentrale fail-closed Policy Engine
13. ADR-0013 Kurzlebige Machine Identities
14. ADR-0014 Ausgehende Agent-Verbindungen
15. ADR-0015 Build und Deployment getrennt
16. ADR-0016 Unveränderliche Deployment-Artefakte
17. ADR-0017 Artifact Trust, SBOM und Provenance
18. ADR-0018 Versioniertes Eventmodell
19. ADR-0019 Persistente und wiederaufnehmbare Workflows
20. ADR-0020 Evidenzbasierte Topologie
21. ADR-0021 Manuelle Topologieautorität
22. ADR-0022 Domain Packs umgehen die Control Plane nicht
23. ADR-0023 Configuration Vault
24. ADR-0024 Data Classification und Redaction
25. ADR-0025 CoreOps Self-Protection
26. ADR-0026 Integration Quality Levels
27. ADR-0027 PostgreSQL vor Graphdatenbank im MVP
28. ADR-0028 Grafana als Analysewerkzeug
29. ADR-0029 KI nur beratend
30. ADR-0030 CorePack und Offline Trust
31. ADR-Candidate (Foundation) Delivery Baseline: Bedeutung von „Docker-first" (Produktanforderung vs. Delivery Baseline vs. Architekturentscheidung)
32. ADR-Candidate (Foundation) Release-Taxonomie und SemVer-Zuordnung (Auflösung der `Foundation 0.1` / `Release 0.1` Kollision)

## Roadmap Candidates

> Meilensteine ohne finale Releasezusage (§49). Nummern vorläufig; verbindliche Taxonomie via CO-WP-003.

- RC-01 Foundation 0.1 – Platform Foundation (nur Doku/Architektur/Governance/Security)
- RC-02 Release 0.1 – Observe (read-only Inventory/Monitoring/Discovery)
- RC-03 Release 0.2 – Map (Topologie v1, Path Explorer, Basis-IPAM)
- RC-04 Release 0.3 – Plan (Desired State, Drift, Blueprints, Policy Evaluation)
- RC-05 Release 0.4 – Deploy (kontrollierte Deployments, Rollback, Evidence)
- RC-06 Release 0.5 – Automate (Workflows, Runbooks, Incident Response)
- RC-07 Release 0.6 – Extend (Windows, Hyper-V, Kubernetes, Netzwerk-/Druckerbereitstellung)
- RC-08 Release 0.7 – Scale (Multi-Site, HA, Multi-Tenant, Compliance Packs)

## Non-Goals

> §7 (frühe Nicht-Ziele) und §45.

- NG-01 Kein vollständiges CI-Build-System
- NG-02 Kein vollständiges SIEM
- NG-03 Kein vollständiges ITSM-/Helpdesk-System
- NG-04 Keine EDR / kein Antivirus
- NG-05 Keine allgemeine Remote-Desktop-Plattform
- NG-06 Keine vollständige Public-Key-Infrastruktur als Ersatz
- NG-07 Kein vollständiger Low-Code-App-Builder
- NG-08 Keine universelle Cloud-Management-Suite
- NG-09 Kein autonomer KI-Administrator
- NG-10 Kein vollständiger Ersatz jedes Herstellerinterfaces
- NG-11 Keine vollständige ERP-/Lizenz-/Vertragsverwaltung
- NG-12 `Enforce`-Auto-Reconciliation kein MVP-Ziel (§20)

## Deferred Topics

- DF-01 Visueller Drag-and-Drop-Workflow-Editor (erst nach stabilem deklarativem Schema, §33)
- DF-02 Externe PKI, TPM, Hardware Attestation, SPIFFE (§31, „Später")
- DF-03 Externe Secret-Backends (Vault, Bitwarden, Enterprise Vaults; §40, „Später")
- DF-04 Tracing in Observability (§24, „optional später")
- DF-05 MSP-/Rechenzentrums-Fokus (§5, „später")
- DF-06 Bidirektionaler NetBox-Sync über read-only hinaus (§19)

## Conflicts and Clarifications

| Conflict ID | Topic | Current Statements | Why Ambiguous | Risk | Recommended Resolving WP | Human Decision Required |
| ----------- | ----- | ------------------ | ------------- | ---- | ------------------------ | ----------------------- |
| CCR-01 | Plane-Taxonomie | §10 nennt sechs Ebenen, davon „Domain Packs" als eine Plane | Domain Packs sind wahrscheinlich plane-übergreifende Integrationspakete, keine eigene Plane | Inkonsistentes Architekturmodell | CO-WP-006 / CO-WP-008 | Ja |
| CCR-02 | Release-Taxonomie | `Foundation 0.1` (§49) und `Release 0.1 – Observe` (§49) tragen identische Nummer | Doppelte 0.1-Nummerierung ist mehrdeutig | Verwirrung bei Versionierung/Kommunikation | CO-WP-003 | Ja |
| CCR-03 | NDF-Level | Manifest `ndf_level: 1` als Bootstrap-Wert; Starter-Vorlage nutzt teils `ndf_level: 2` | Bootstrap-Wert vs. Reifeaussage | Falsch interpretierte NDF-Reife | CO-WP-005 | Ja |
| CCR-04 | `main`-Normativität | Concept §Ausgangskontext: `main` „darf ergänzend geprüft werden"; CO-WP-Governance: `main` nicht normativ | Zwei Formulierungen zur `main`-Rolle | Uneinheitliche NDF-Basisregel | CO-WP-005 | Ja |
| CCR-05 | Offline Policy | §21/§38: Fail-closed Policy Engine vs. Offline-/Air-Gap-Betrieb | Unklar, welche zuvor autorisierten Aktionen offline bei getrennter Control Plane erlaubt sind | Blockierte Betriebsfähigkeit oder unsichere Offline-Freigabe | CO-WP-013 / CO-WP-023 | Ja |
| CCR-06 | Machine Identity vs. Air-Gap | §31: kurzlebige Zertifikate; §38: lang laufende isolierte/Air-Gap-Standorte | Kurzlebige Credentials vs. lange Offline-Phasen | Ausfall der Maschinenidentität in Air-Gap | CO-WP-010 | Ja |
| CCR-07 | Privilegierte Ausführung | §52.1 „keine allgemeine Remote-Root-Shell" vs. notwendige privilegierte Wartung | Wartung benötigt teils privilegierte Aktionen | Entweder Wartung unmöglich oder Sicherheitsinvariante verletzt | CO-WP-013 | Ja |
| CCR-08 | Audit vs. Datenschutz | §35/§52 unveränderliches Audit vs. §41 Löschung/Redaction/Retention | Immutable Audit kollidiert mit Redaction/Retention | Rechts-/Datenschutzkonflikt | CO-WP-025 | Ja |
| CCR-09 | Offline-First-Facetten | §38 unterscheidet nicht klar Runtime/Installation/Updates/Recovery/Builds | „Offline First" ist mehrdeutig | Unterschiedliche Erwartungen an Offline-Fähigkeit | CO-WP-023 | Ja |
| CCR-10 | Docker-first | §3/§53 „Docker-first" ohne Klassifikation | Produktanforderung vs. Delivery Baseline vs. Architekturentscheidung vs. Kandidat | Verfrühte oder unklare Delivery-Festlegung | CO-WP-003 / ADR | Ja |
| CCR-11 | Foundation-Queue-Abweichung | Concept §50 (21 WPs) vs. verbindliche CoreOps-Queue (CO-WP-001…031) | Zwei unterschiedliche Queue-Schnitte | Verwirrung über den verbindlichen Plan | CO-WP-003 | Ja |
| CCR-12 | Herstellersupport | §26/§27/§18 Herstellerlisten vs. „kein Supportversprechen" | Namensnennung könnte als Support missverstanden werden | Falsche Support-/Kompatibilitätserwartung | CO-WP-004 | Ja |

## Classification Summary

Zählung nach Primärklasse (Einträge in diesem Dokument; Klassifikation überwiegend auf Abschnittsebene):

| Klasse | Anzahl |
| ------ | ------ |
| APC | 10 |
| BG | 8 |
| FR | 19 |
| AC | 8 |
| TC | 10 |
| ADR | 32 |
| RC | 8 |
| NG | 12 |
| DF | 6 |
| CCR | 12 |

> Hinweis: Die Zahlen bilden die dokumentierten Einträge ab, nicht jeden Einzelsatz des Concepts. Sekundärklassen sind in den Zeilen vermerkt und hier nicht separat gezählt. Es besteht keine künstliche Präzision.
