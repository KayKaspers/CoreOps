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

## CO-WP-004 Konsolidierung

- **Foundation Capability Matrix:** [docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md) — 74 Capabilities über 12 Domains; alle `not-implemented`/`not-supported` (Proposed for acceptance).
- **Initiale Observe-Supportgrenze:** [docs/integrations/INITIAL_SUPPORT_BOUNDARY.md](../docs/integrations/INITIAL_SUPPORT_BOUNDARY.md) — Observe vollständig read-only, max. Integrationslevel 2; Map/Write ausgeschlossen (Proposed for acceptance).
- **Statusdimensionen und Evidence:** drei getrennte Statusdimensionen (Roadmap/Implementation/Support) und ein standardisierter Support-Evidence-Satz definiert; Herausgeberklasse ≠ Supportstatus.
- **Herstellersupport-Grenze:** `CCR-12` vorgeschlagen aufgelöst — Herstellernennung = Kandidat, kein Support/Partnerschaft/Zertifizierung. Aktuell keine Integration `supported`, keine Runtime-Capability implementiert.

## CO-WP-004A Konsolidierung (Sovereignty & BSI)

- **Concept-v3.1-Amendment:** [docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md](../docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md) (Accepted Product Direction; erweitert v3.0).
- **Sovereignty & Dependency Policy:** [docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](../docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) — CoreOps-Kern ohne verpflichtende externe Management-/CMDB-/ITSM-/Automation-/Container-/GitOps-Plattform; technische Basisabhängigkeiten getrennt und offen (keine akzeptiert).
- **BSI-Positionierung:** [docs/security/BSI_ALIGNMENT_POSITIONING.md](../docs/security/BSI_ALIGNMENT_POSITIONING.md) — BSI-orientiert entwickelt; **keine** Zertifizierung/Konformität/Zulassung/VS-Eignung behauptet.
- **Betriebsprofile (Zielmodell):** Standard / Hardened / Government (Government = späteres Nachweisprofil, keine Zertifizierung).
- **Lessons Learned & NDF-Feedback:** als Governance-Richtung bzw. kontrollierter Kandidatenprozess registriert; Detailprozess in `CO-WP-004B`.
- **ITIL/PRINCE2:** nur Foundation-Kandidaten; Tailoring in `CO-WP-004D`.

## CO-WP-004B Konsolidierung (Lessons Learned & NDF Feedback)

- **Lessons-Learned-Prozess:** [docs/governance/LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md) — Lifecycle, Klassen, Pflichtfelder, Statuswerte (Proposed for acceptance).
- **NDF-Feedback-Prozess:** [docs/governance/NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md) — kontrollierte, manuelle Rückführung; kein automatischer NDF-Rückfluss; Human-Maintainer-Gate für jeden Transfer.
- **Register:** [LESSONS_LEARNED_REGISTER.md](LESSONS_LEARNED_REGISTER.md) (15 Lessons) und [NDF_FEEDBACK_CANDIDATES.md](NDF_FEEDBACK_CANDIDATES.md) (7 Kandidaten).
- **Status:** Kein Kandidat an das NDF übertragen oder adoptiert; kein NDF-Repository verändert.

## CO-WP-004B1 Konsolidierung (First NDF Feedback Transfer Package)

- **Transferpaket:** [docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md](../docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md) — approved for NDF Intake; 3 Bundles; alle 7 Kandidaten enthalten.
- **Kandidatenstatus:** alle 7 auf `approved-for-transfer` (Nova Gate: approved; Human-Maintainer Gate: approved, Commit 4ad3111; `Backlink Status: transfer-package-prepared`).
- **Klassenstatistik-Korrektur:** Zusammenfassungsfehler im Lessons-Learned-Register behoben (LL-PROCESS-Liste enthielt fälschlich „001"; keine Lesson umklassifiziert).
- **Status:** Kein Kandidat an das NDF übertragen oder adoptiert; kein NDF-Repository verändert.

## CO-WP-004B2 Konsolidierung (Finalize NDF Intake Approval Gates)

- **Governance-Korrektur:** Transfer Package Status und alle 7 Human-Maintainer-Gates explizit in Repository-Dokumenten auf "approved" gesetzt.
- **Neue Lesson:** LL-016 erfasst — „Commit-gated status transitions must be represented in repository state" — Evidenz für Blockierungserlebnis und Ursachenanalyse.
- **Lehrpunkt:** Ein Commit liefert Evidenz für eine Freigabe, verändert aber keine Statuswerte in Repository-Dokumenten. Freigabefelder müssen VOR dem Commit auf den beabsichtigten Zielzustand gesetzt werden.
- **Status:** Kein Kandidat übertragen/adoptiert; NDF-Repository unverändert.

## CO-WP-004B3 Konsolidierung (Record Completed NDF Intake Transfer)

- **Intake-Dokumentation:** Transfer Package 001 dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e).
- **Kandidaten-Status:** Alle 7 auf `transferred-to-ndf` (NDF Work Package: `NDF-INTAKE-COREOPS-001`; Backlink Status: `intake-review-committed`).
- **Adoption:** Kein Kandidat adoptiert; Adoption Version leer; Adoption-Entscheidungen ausstehend.
- **Status:** Intake abgeschlossen; NDF-Repository unverändert; keine Adoption-WPs gestartet.

## CO-WP-004B4 Konsolidierung (Record Completed NDF Adoption and Close Transfer Package 001)

- **Adoption-Dokumentation:** Alle 7 Kandidaten von `transferred-to-ndf` auf `adopted-in-ndf` gesetzt, verteilt auf drei Adoption-Work-Packages.
- **Adoption A** (`NDF-ADOPT-COREOPS-001A`, Commit 1ebffa6): NDF-FC-COREOPS-001, -003, -004.
- **Adoption B** (`NDF-ADOPT-COREOPS-001B`, Commit e894c6f): NDF-FC-COREOPS-002.
- **Adoption C** (`NDF-ADOPT-COREOPS-001C`, Commit ebf716c): NDF-FC-COREOPS-005, -006, -007.
- **Transfer Package 001:** Status `Closed – all candidates processed`; Intake, Adoption und CoreOps-Backlink vollständig.
- **Release-Grenze:** NDF-Release-Zuordnung bleibt `not yet assigned`; keine NDF-Version behauptet. `adopted-in-ndf` ≠ in veröffentlichter NDF-Version enthalten.
- **Status:** Adoption für alle 7 abgeschlossen; NDF-Repository unverändert; kein weiteres Adoption-WP gestartet.

## CO-WP-004C Konsolidierung (BSI and Public-Sector Readiness Baseline)

- **Readiness-Baseline:** [BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](../docs/security/BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md) — Claim Boundaries, 18 Readiness-Domänen (PSR-01…PSR-18), Verantwortungs- und Evidenzmodell (capability ≠ available ≠ satisfied), Logging/Detection, Offline/Air-Gap, Souveränität/Cloud.
- **Referenz-/Claims-Register:** [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](../docs/security/BSI_REFERENCE_AND_CLAIMS_REGISTER.md) — offizielles Referenzset (BSI 200-1…200-4, IT-Grundschutz-Kompendium versionsoffen, Mindeststandards, Protokollierung/Detektion v2.1; C5/C3A bedingt), keine erfundenen Anforderungen.
- **Public-Sector-Profil:** [PUBLIC_SECTOR_READINESS_PROFILE.md](../docs/governance/PUBLIC_SECTOR_READINESS_PROFILE.md) — interne Profile Standard ⊂ Hardened ⊂ Government (keine Zertifizierungsstufen).
- **Governance-Fortschreibung:** Decision Index +7 (DEC-S-16…22), Risk Register +10 (RISK-40…49). BSI-Positionierung additiv aktualisiert.
- **Claim-Grenzen:** Keine BSI-Zertifizierung, keine IT-Grundschutz-Vollkonformität, keine Behörden-/VS-NfD-Freigabe; Produktreife ≠ Deployment-Compliance.
- **Status:** Baseline etabliert; Capability Matrix unverändert (→ CO-WP-004E); keine ADR, keine Technologieauswahl.

## CO-WP-004D Konsolidierung (ITIL and PRINCE2 Applicability and Tailoring Decision)

- **Applicability/Tailoring:** [ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](../docs/governance/ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md) — ITIL `adopted-with-tailoring` (20-Element-Matrix), PRINCE2 Version 7 `optional-profile` (12-Muster-Matrix), ITIL-4-/Version-5-Grenze, Konflikthierarchie, Overload-Guards.
- **Governance-Profile:** [COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md](../docs/governance/COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md) — Service Operations Guidance, Major Deployment Project Profile, Public-Sector Delivery Profile (intern, optional, keine Zertifizierungsstufen).
- **NDF-Primat:** NDF bleibt primäres Softwareentwicklungs-/Repository-Governance-Framework; keine parallele WP-/Decision-/Risk-Struktur.
- **Governance-Fortschreibung:** Decision Index +8 (DEC-S-23…30, löst DEC-S-10/11), Risk Register +11 (RISK-50…60).
- **Claim-Grenzen:** keine ITIL-/PRINCE2-Zertifizierung, keine vollständige Implementierung, kein Endorsement (u. a. PeopleCert), keine Tool-Abhängigkeit.
- **Status:** Tailoring-Entscheidung umgesetzt; Capability Matrix und Lessons-Learned-Register unverändert; keine ADR; keine NDF-Rückführung.

## CO-WP-004E Konsolidierung (Capability Matrix Security and Governance Alignment)

- **Alignment-Dokument:** [CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md](../docs/security/CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md) — fünf Statusdimensionen (Roadmap/Implementation/Support/Evidence/Security-Governance), Profile, PSR-Mapping, Verantwortung, Evidenz-Drei-Zustands-Modell.
- **Capability-Matrix-Spezifikation:** [CAPABILITY_MATRIX_SPEC.md](../docs/project-system/CAPABILITY_MATRIX_SPEC.md) — neu erstellt (existierte lokal nicht); mehrdimensionales Modell + Migrationsregeln.
- **Foundation Capability Matrix:** additiv um Evidence-/Security-Governance-Status, Profile Relevance, PSR-Zuordnung und Responsibility-Codes für alle Capabilities erweitert; keine gelöscht/umbenannt/hochgestuft.
- **Zählkorrektur:** Matrix enthält **94** Capabilities (frühere Summe „74" grep-verifiziert korrigiert; „74"-Referenzen anderer Dokumente später abzugleichen).
- **Governance-Fortschreibung:** Decision Index +7 (DEC-S-31…37), Risk Register +6 (RISK-61…66).
- **Claim-Grenzen:** PSR-Mapping = Readiness-Relevanz, keine BSI-Compliance; kein detailliertes Control-Mapping; kein `compliant`/`requirement-satisfied`/`government-approved`.
- **Status:** Alignment dokumentiert; keine Capability implementiert; keine ADR; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-005 Konsolidierung (Language Standard, Public Neutrality and Repository Governance)

- **Language Standard:** [COREOPS_LANGUAGE_STANDARD.md](../docs/governance/COREOPS_LANGUAGE_STANDARD.md) — Englisch kanonisch für maschinenbezogene Bezeichner; DE/EN primäre Produktsprachen; Translation-Status-Modell; semantische Parität Pflicht; keine automatische Paritätsbehauptung.
- **Public Neutrality & Disclosure:** [PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](../docs/governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md) — Hersteller-/Organisations-/institutionelle Neutralität; Secrets-/Datenschutz-/Infrastrukturgrenzen; sichere synthetische Beispieldaten; Redaction; Endorsement-Grenzen.
- **Repository Governance:** [REPOSITORY_GOVERNANCE_STANDARD.md](../docs/governance/REPOSITORY_GOVERNANCE_STANDARD.md) — Human-Maintainer-Gates erhalten; Source-of-Truth-Hierarchie (erweitert FOUNDATION_SCOPE_LOCK); authoritative/derived; Dokumentstatus/Supersession; Dateinamen/stabile IDs; UTF-8/Zeilenenden; PowerShell-Korrekturstandard; Public-Hygiene.
- **Governance-Fortschreibung:** Decision Index +6 (DEC-S-38…43, getrennte Dimensionen Decision Class/Lifecycle/Binding Level), Risk Register +13 (RISK-67…79).
- **Grenzen:** keine automatisierte Durchsetzung, keine README/CONTRIBUTING/.gitignore/.gitattributes-Änderung, keine Übersetzung, keine Massenkonvertierung, keine ADR.
- **Status:** Governance-Foundation etabliert; nächstes WP `CO-WP-006`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-006 Konsolidierung (System Context, Plane Taxonomy and External Boundaries)

- **System Context:** [SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md](../docs/architecture/SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md) — Akteure/Rollen (7), Produkt-/Deployment-/Managed-/Operator-/Provider-/Evidence-Grenzen, 20 externe Systemklassen, 15 Interaktionsklassen, Kontrollautorität, Connected/Restricted/Offline-Modi, Failure-/Degradationsgrenzen, 13 Datenklassen, Mermaid-Kontextdiagramm.
- **Plane Taxonomy:** [COREOPS_PLANE_TAXONOMY.md](../docs/architecture/COREOPS_PLANE_TAXONOMY.md) — 10 logische Planes (plane ≠ Prozess/Container/Microservice/Segment/Deployment-Einheit); Edge/Agent Plane optional (agentless möglich); Managed Resource Plane außerhalb Produktgrenze; Mermaid-Plane-Diagramm.
- **Trust Boundaries:** [TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md](../docs/security/TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md) — 11 Vertrauensgrenzen; Policy-to-Action und Control-to-Execution getrennt; fail-closed; Threat Model deferred zu CO-WP-007.
- **Governance-Fortschreibung:** Decision Index +10 (DEC-S-44…53, getrennte Dimensionen), Risk Register +14 (RISK-80…93).
- **Grenzen:** keine Technologieauswahl, keine Komponentenarchitektur, kein Threat Model, keine ADR; Diagramme `conceptual`, keine Deployment-Topologie.
- **Status:** Architektur-Kontext-Foundation etabliert; nächstes WP `CO-WP-007`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-007 Konsolidierung (Threat Model and Trust Boundaries)

- **Foundation Threat Model:** [COREOPS_FOUNDATION_THREAT_MODEL.md](../docs/security/COREOPS_FOUNDATION_THREAT_MODEL.md) — 24 Assets (AST-01…24), 16 Threat Actors, 18 Threat-Kategorien, Angriffsflächen für alle 10 Planes, Trust-Boundary-Analyse (TB-01…11), 17 Sicherheitsinvarianten, 5 Abuse Cases (AB-1…5), 2 Mermaid-Diagramme (Threat Surface Overlay, Abuse-Path).
- **Threat Scenario Register:** [THREAT_SCENARIO_REGISTER.md](../docs/security/THREAT_SCENARIO_REGISTER.md) — 40 Szenarien (THR-001…040), stabile IDs, qualitative evidence-bounded Ratings, Mitigation States; kein `mitigated`/`closed` ohne Evidenz.
- **Trust-Boundary-Dokument:** additiv um stabile Boundary-IDs (TB-01…11) + Threat-Verweise + Threat-Model-Status erweitert (kein Parallelmodell).
- **Governance-Fortschreibung:** Decision Index +10 (DEC-S-54…63, getrennte Dimensionen), Risk Register +10 (RISK-94…103, nur Threat-Model-Governance; einzelne Threats nicht dupliziert).
- **Grenzen:** keine Sicherheitskontrolle implementiert/validiert; kein Pentest/Scan; keine Technologie-/Krypto-/Authentisierungsauswahl; keine ADR; Capability Matrix unverändert. Invarianten = Designanforderungen, keine implementierten Kontrollen.
- **Status:** Foundation-Security-Baseline etabliert; nächstes WP `CO-WP-008`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-008 Konsolidierung (Architecture and Module Boundaries)

- **Logical Module Architecture:** [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](../docs/architecture/COREOPS_LOGICAL_MODULE_ARCHITECTURE.md) — 17 logische Module (MOD-EXP/IAM/POL/INV/OBS/TOP/WFL/EXE/ADP/AGT/DEP/STA/SEC/EVD/NOT/OFF/EXT-001), Modulklassifikation, Policy/Control/Execution-Trennung, 2 Mermaid-Diagramme (Logical Module, Prohibited Bypass).
- **Module Catalog:** [COREOPS_MODULE_CATALOG.md](../docs/architecture/COREOPS_MODULE_CATALOG.md) — vollständiges 17-Modul-Register; autoritative Daten-/Zustandsownership (ein Owner pro Konzept); Threat-/Invarianten-/Capability-Domänen-Referenzen; ersetzt **nicht** die Capability Matrix.
- **Boundary/Dependency Standard:** [MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md](../docs/architecture/MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md) — Autoritätsgrenzen, erlaubte/verbotene Abhängigkeiten, verbotene Bypässe, Zyklusverbot, Adapter-/Agent-/Evidence-/Offline-/Plugin-Grenzen.
- **Governance-Fortschreibung:** Decision Index +12 (DEC-S-64…75), Risk Register +15 (RISK-104…118).
- **Grenzen:** module ≠ microservice/process/container/deployment unit; keine Technologie-/Protokoll-/Deployment-Auswahl; keine ADR; Invarianten = Designanforderungen; Capability Matrix + Threat-Dateien unverändert.
- **Status:** Foundation-Logische-Architektur etabliert; nächstes WP `CO-WP-009`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-009 Konsolidierung (Human Identity, Workspaces, RBAC and Break Glass)

- **Identity Governance:** [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](../docs/security/HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md) — Begriffsmodell (person/identity/account/principal), Projekt- ≠ Runtime-Rollen (Human Maintainer ≠ Runtime-Autorität), Account-Lifecycle, Auth ≠ Authz, Sessions/Reauth, Recovery, Delegation, SoD.
- **Workspace/RBAC/Scope:** [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../docs/security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md) — Workspace ≠ Security-Tenant; deny-by-default/least-privilege RBAC; Permission-Taxonomie; 8 Scope-Typen; Cross-Workspace explizit/auditierbar.
- **Break-Glass:** [BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](../docs/security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) — benannt/temporär/reason-/scope-bound/auditiert; Ablaufpflicht; verpflichtender Post-Event Review; Offline-Emergency governed (nicht anonym).
- **Governance-Fortschreibung:** Decision Index +14 (DEC-S-76…89), Risk Register +18 (RISK-119…136).
- **Grenzen:** keine Auth-/IdP-/Session-/MFA-/Policy-Engine-/Tenant-Isolation-Auswahl; keine ADR; Invarianten = Designanforderungen; Modul-/Capability-/Threat-Dateien unverändert.
- **Status:** Foundation-Identity-/RBAC-/Break-Glass-Baseline etabliert; nächstes WP `CO-WP-010`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-010 Konsolidierung (Machine Identity, Enrollment and Offline Credential Lifecycle)

- **Machine Identity Governance:** [MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md](../docs/security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md) — Begriffsmodell, Human ≠ Machine, 10 Principal-Klassen (mit Owning-Modul/Threats), Lifecycle, Scope/Authorization, Agent-/Adapter-Grenze, Automation Clients.
- **Enrollment & Trust Lifecycle:** [MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md](../docs/security/MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md) — Enrollment (explizit/owner-/scope-bound), Bootstrap-Grenze, Trust Establishment (Registration ≠ Trust), Suspension/Revocation/Compromise, Re-Enrollment, Decommissioning, Offline-Enrollment.
- **Credential & Rotation Governance:** [OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md](../docs/security/OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md) — Credential-Metadaten, Raw-Secret-Grenze (Governance ≠ Rohsecret-Ownership; Speicherung deferred), Rotation/Renewal/Overlap, Offline-Distribution/Reconciliation.
- **Governance-Fortschreibung:** Decision Index +16 (DEC-S-90…105), Risk Register +19 (RISK-137…155).
- **Grenzen:** Discovery ≠ Enrollment; Enrollment ≠ Write Authority; keine PKI-/Zertifikats-/Krypto-/TPM-/mTLS-/SSH-/Enrollment-Protokoll-/Secret-Store-Auswahl; keine ADR; Invarianten = Designanforderungen.
- **Status:** Foundation-Machine-Identity-Baseline etabliert; nächstes WP `CO-WP-011`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-011 Konsolidierung (Source of Truth and Field Provenance)

- **Source of Truth / State Authority:** [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](../docs/architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) — SoT ≠ System of Record; 10 Authority-Klassen; autoritative Modulownership (ein Owner pro Feldkonzept); Desired/Observed/Effective/Derived/Cached; Konfliktmodell.
- **Field Provenance / Lineage:** [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](../docs/architecture/FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) — stabile Feldidentität (≠ UI-Label/Adapter-Feld), 22-Feld-Provenance-Metadatenmodell, Freshness/Trust/Validation getrennt, Transformation-Lineage, Import/Derived, Privacy/Minimierung.
- **Offline Reconciliation / Conflict:** [OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md](../docs/security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md) — Conflict Detection/States, Resolution Inputs (kein universelles Last-Write-Wins), Revoked Sources, Replay/Duplication, Time-Uncertainty, Partial Import, fail-closed.
- **Governance-Fortschreibung:** Decision Index +16 (DEC-S-106…121), Risk Register +17 (RISK-156…172).
- **Grenzen:** Observed ≠ Desired; Derived/Cached nicht autoritativ; Import ≠ Authority-Transfer; keine Storage-/DB-/Merge-/CRDT-/Messaging-/Sync-/Krypto-Provenance-Auswahl; keine ADR; Invarianten = Designanforderungen.
- **Status:** Foundation-Data-Governance-Baseline etabliert; nächstes WP `CO-WP-012`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## CO-WP-012 Konsolidierung (Observed, Desired, Effective State and Drift)

- **State-Semantik:** [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](../docs/architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) — Desired/Observed/Reported/Effective/Last-Known getrennt; Desired-State-Lifecycle; Effective-State-Grenze (indeterminate/conflicted bei unklarer Autorität); Unknown/Conflicted sichtbar.
- **Drift & Konvergenz:** [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](../docs/architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md) — Drift-Definition/12 Arten, 14 Detection States, Drift Record, Impact/Urgency, Exceptions, Konvergenz (`verified-converged` braucht Evidenz), Partial Failure.
- **Safe Remediation:** [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../docs/security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md) — Detection/Recommendation/Plan/Approval/Execution/Verification getrennt; Read-only-Grenze; Remediation-Lifecycle; Rollback braucht Verifikation; fail-closed; **Auto-Remediation deferred**.
- **Governance-Fortschreibung:** Decision Index +15 (DEC-S-122…136), Risk Register +17 (RISK-173…189).
- **Grenzen:** keine Beobachtung ≠ keine Drift; kein Drift ≠ Compliance; Executed ≠ successful ≠ verified ≠ compliant; keine Engine-/Scheduler-/Queue-/Auto-Remediation-Auswahl; keine ADR; Invarianten = Designanforderungen.
- **Milestone:** CO-WP-005…012 (acht WPs) → Milestone-Lessons-Review-Eignung `yes`; Bündelentscheidung bei Nova/Human Maintainer; nicht automatisch gestartet.
- **Status:** Foundation-State-Management-Baseline etabliert; nächstes WP `CO-WP-013`; Lessons-Learned-Register unverändert; keine NDF-Rückführung.

## Milestone Lessons Review (CO-WP-005…012)

- **Ergebnis:** GO WITH NOTES FOR CO-WP-013. Neues Dokument [MILESTONE_REVIEW_CO_WP_005_TO_012.md](../project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md); Nova Review pending.
- **Lessons:** LL-017…022 konsolidiert (Lessons gesamt 22).
- **NDF-Kandidaten:** NDF-FC-COREOPS-008…010 (`candidate-pending-nova-review`); keine NDF-Rückführung.
- **Read-only-Befunde:** Risk Register (189) → späterer Konsolidierungslauf; Decision Index Alt-Kombistatus (DEC-S-01…37) vs. getrennte Dimensionen (DEC-S-38…136); Capability-Count „74→94" in Alt-Abschnitten. Follow-ups ~CO-WP-029/030.
- **Grenzen:** Risk Register und Decision Index unverändert (read-only); kein NDF-Transfer; CO-WP-013 nicht begonnen.

## Notes

- **NDF-Level-Semantik:** `ndf_level: 1` bedeutet ausschließlich „initialer NDF-Projektbootstrap" — keine vollständige NDF-Konformität, technische Reife, Releasebereitschaft, Sicherheitsfreigabe oder erreichte Produktstufe. Der Zahlenwert bleibt unverändert. Die Starter-Vorlage mit `ndf_level: 2` bleibt dokumentierte NDF-Quellambiguität, ist aber kein Blocker (CCR-03).
- **Skills-first:** Lokaler NDF-v1.0.0-Skills-Pack seit CO-WP-001A vorhanden; pro Work Package selektiv genutzt.
