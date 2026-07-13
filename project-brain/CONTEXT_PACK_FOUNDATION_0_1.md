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

`CO-WP-006 – System Context, Plane Taxonomy and External Boundaries` (docs-only / architecture-context foundation). Nach Bearbeitung im **Nova Review** (`pending`). Vorheriges WP: `CO-WP-005 – GO WITH NOTES`.

## System Context, Plane Taxonomy and External Boundaries (CO-WP-006)

- Drei neue Dokumente: System Context and External Boundaries, Plane Taxonomy, Trust/Deployment/Execution Boundaries.
- 7 Akteure; Produkt- ≠ Deployment- ≠ Managed-Grenze; 20 externe Systemklassen (optional); 15 Interaktionsklassen (read ≠ write); Kontrollautorität; Connected/Restricted/Offline-Modi; Failure-Grundregeln; 13 Datenklassen; zwei Mermaid-Diagramme (conceptual).
- 10 logische Planes (plane ≠ Deployment-Einheit); Edge/Agent Plane optional; Managed Resource Plane außerhalb Produktgrenze. 11 Vertrauensgrenzen; Control-to-Execution/Policy-to-Action getrennt; fail-closed; Threat Model deferred (CO-WP-007).
- Decision Index +10 (DEC-S-44…53), Risk Register +14 (RISK-80…93). Keine Technologie/Architektur/Threat Model; keine ADR; Lessons-Learned unverändert; keine NDF-Rückführung.

## Language Standard, Public Neutrality and Repository Governance (CO-WP-005)

- Drei neue Governance-Dokumente: Language Standard, Public-Neutrality-and-Disclosure-Policy, Repository-Governance-Standard.
- Englisch kanonisch für maschinenbezogene Bezeichner; DE/EN primäre Produktsprachen; Translation-Status-Modell; semantische Parität Pflicht; keine automatische Parität.
- Hersteller-/Organisations-/institutionelle Neutralität; Nennung ≠ Endorsement; Public-Sector ≠ Behördenfreigabe; Secrets-/Datenschutzgrenzen; synthetische Beispieldaten; Redaction.
- Human-Maintainer-Gates erhalten; Source-of-Truth-Hierarchie (erweitert FOUNDATION_SCOPE_LOCK); Dokumentstatus/Supersession; Dateinamen/stabile IDs; UTF-8/Zeilenenden; PowerShell-Korrekturstandard.
- Decision Index +6 (DEC-S-38…43), Risk Register +13 (RISK-67…79). Keine ADR; keine automatisierte Durchsetzung; Lessons-Learned unverändert; keine NDF-Rückführung.

## Capability Matrix Security and Governance Alignment (CO-WP-004E)

- Zwei neue Dokumente: Capability-Security-and-Governance-Alignment + Capability-Matrix-Spezifikation (letztere neu erstellt).
- Foundation Capability Matrix additiv um fünf Statusdimensionen (Roadmap/Implementation/Support/Evidence/Security-Governance), Profile Relevance, PSR-01…18-Zuordnung, Responsibility-Codes erweitert; keine Capability gelöscht/umbenannt/hochgestuft.
- **Zählkorrektur: 94 Capabilities (grep-verifiziert); frühere „74"-Summe korrigiert; andere Referenzen später abzugleichen.**
- PSR-Mapping = Readiness-Relevanz, nicht BSI-Compliance; kein Control-Mapping; Evidence `not-assessed`; kein `compliant`. Decision Index +7 (DEC-S-31…37), Risk Register +6 (RISK-61…66). Keine ADR; Lessons-Learned unverändert; keine NDF-Rückführung.

## ITIL and PRINCE2 Applicability and Tailoring (CO-WP-004D)

- Zwei neue Dokumente: ITIL/PRINCE2-Applicability-and-Tailoring und interne Service-/Project-Governance-Profile.
- ITIL `adopted-with-tailoring`; PRINCE2 Version 7 `optional-profile`; vollständige Implementierung beider `rejected`. ITIL 4 & Version 5 mit Versionsgrenzen (Bridge anerkannt).
- NDF bleibt primäres Framework; Konflikthierarchie + Overload-Guards; keine parallele Governance. Drei interne, optionale Profile (keine Zertifizierungsstufen).
- Decision Index +8 (DEC-S-23…30), Risk Register +11 (RISK-50…60). Capability Matrix + Lessons-Learned-Register unverändert; keine NDF-Rückführung; keine Zertifizierung/Endorsement/Tool-Abhängigkeit; keine ADR.

## BSI and Public-Sector Readiness Baseline (CO-WP-004C)

- Drei neue Dokumente: Readiness-Baseline (18 PSR-Domänen, Verantwortungs-/Evidenzmodell), Reference-/Claims-Register (offizielles BSI-Referenzset; C5/C3A bedingt), Public-Sector-Profil (Standard ⊂ Hardened ⊂ Government, intern).
- Claim Boundaries: keine BSI-Zertifizierung/Vollkonformität/Behörden-/VS-NfD-Freigabe; Produktreife ≠ Deployment-Compliance; keine einzelnen BSI-Anforderungen erfunden.
- Evidenz-Drei-Zustands-Modell (capability ≠ available ≠ satisfied). Offline/Air-Gap und Souveränität/Cloud eingeordnet.
- Decision Index +7 (DEC-S-16…22), Risk Register +10 (RISK-40…49). BSI-Positionierung additiv aktualisiert.
- **Capability Matrix unverändert (→ CO-WP-004E); keine ADR; keine Technologie ausgewählt.**

## NDF-Transferpaket – Vollständiger Lebenszyklus (CO-WP-004B1 bis 004B4)

- **CO-WP-004B1:** Transfer Package 001 mit 3 Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer`.
- **CO-WP-004B2:** Package Status und 7 Human-Maintainer-Gates auf `approved`; Lesson LL-016 erfasst.
- **CO-WP-004B3:** Alle 7 Kandidaten auf `transferred-to-ndf` (NDF-INTAKE-COREOPS-001, Commit d08e35e).
- **CO-WP-004B4:** Alle 7 Kandidaten auf `adopted-in-ndf`, verteilt auf drei Adoption-WPs:
  - Adoption A (`NDF-ADOPT-COREOPS-001A`, Commit 1ebffa6): 001, 003, 004.
  - Adoption B (`NDF-ADOPT-COREOPS-001B`, Commit e894c6f): 002.
  - Adoption C (`NDF-ADOPT-COREOPS-001C`, Commit ebf716c): 005, 006, 007.
- **Transfer Package 001 Status:** `Closed – all candidates processed`.
- **Release-Grenze:** NDF-Release-Zuordnung bleibt `not yet assigned`; keine NDF-Version behauptet. `adopted-in-ndf` ≠ in veröffentlichter NDF-Version enthalten.
- **NDF-Repository unverändert durch CoreOps; keine weiteren Adoption-WPs gestartet.**

## Lessons Learned und NDF Feedback (CO-WP-004B)

- Lessons-Learned-Prozess und NDF-Feedback-Prozess etabliert (Proposed for acceptance).
- 15 Lessons erfasst (13 retrospektiv aus CO-WP-001…004A, 2 aus CO-WP-004B selbst) im Lessons-Learned-Register.
- 7 reservierte NDF-Feedback-Kandidaten (NDF-FC-COREOPS-001…007) bewertet, alle `ready-for-bundling`.
- **Kein Kandidat übertragen/adoptiert; kein NDF-Repository verändert; kein NDF-Work-Package erstellt.**
- Transferentscheidungen bleiben Nova Review + Human-Maintainer-Gate vorbehalten.

## Souveränität und BSI-Orientierung (CO-WP-004A)

- Concept-v3.1-Amendment, Sovereignty-and-Dependency-Policy und BSI-Alignment-Positioning erstellt (Accepted Product Direction).
- Produktsouveränität akzeptiert; verpflichtende externe Managementprodukte als Kernabhängigkeit ausgeschlossen; technische Basisabhängigkeiten offen (keine akzeptiert).
- BSI-orientiert; **keine** Zertifizierung/Konformität/Zulassung/VS-Eignung behauptet. Standard-/Hardened-/Government-Profile registriert.
- Lessons Learned (Governance-Richtung) und kontrollierter NDF-Rückfluss (Kandidatenprozess) registriert; ITIL/PRINCE2 nur Kandidaten.
- Capability Matrix unverändert; technische Architektur offen; keine ADR.

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

`CO-WP-007 – Threat Model and Trust Boundaries` (security-baseline; planned-next; pending Nova review von CO-WP-006 und Human-Maintainer-Freigabe). Die 004er-Erweiterungsserie (004A…004E), CO-WP-005 und CO-WP-006 sind abgeschlossen. External: NDF-Release-Zuordnung für die drei Adoption-Commits ausstehend.

## Aktuelle Blocker

- Keine harten Blocker. In CO-WP-003/004/004A/004B/004B1 adressiert (proposed/treatment-planned/accepted-direction/approved-for-transfer): Release-Taxonomie, Docker-first, Queue-Autorität, NDF-`main`, NDF-Level, CCR-12, Observe-Scope, Legacy-Protokolle, Capability-Maturity, Drucktelemetrie, Produktsouveränität, BSI-Orientierung, Basisabhängigkeiten, Lessons-Learned-/NDF-Feedback-Prozess, erstes NDF-Transferpaket.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09; technische ADR-Kandidaten; ITIL/PRINCE2-Tailoring (CO-WP-004D); konkretes IT-Grundschutz-Mapping (CO-WP-004C); NDF-seitiger Intake Review für Transfer Package 001.

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
- Souveränität und BSI-Orientierung als Accepted Product Direction registriert (CO-WP-004A); keine Zertifizierung/VS-Eignung behauptet; keine konkrete Dependency akzeptiert.
- Lessons-Learned- und NDF-Feedback-Prozess etabliert (CO-WP-004B); 15 Lessons, 7 NDF-Kandidaten (alle `ready-for-bundling`); kein Kandidat übertragen/adoptiert.
- Weiter offen: Konflikte CCR-01, 05, 06, 07, 08, 09; 39 Foundation-Risiken (RISK-01…39, davon 22 `treatment-planned`).
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
- [docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md](../docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md)
- [docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](../docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md)
- [docs/security/BSI_ALIGNMENT_POSITIONING.md](../docs/security/BSI_ALIGNMENT_POSITIONING.md)
- [docs/governance/LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md)
- [docs/governance/NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md)
- [project-system/LESSONS_LEARNED_REGISTER.md](../project-system/LESSONS_LEARNED_REGISTER.md)
- [project-system/NDF_FEEDBACK_CANDIDATES.md](../project-system/NDF_FEEDBACK_CANDIDATES.md)
- [docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md](../docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md)
- [project-brain/PROJECT_BRAIN.md](PROJECT_BRAIN.md)
- [ROADMAP.md](../ROADMAP.md)

## Kompakte Historie

- `CO-WP-001` (docs-only): Core Governance Skeleton erstellt (7 Dateien). Nova-Bewertung: `GO WITH NOTES` (read-only-Git-Note).
- `CO-WP-001A` (docs-only): Vollständiger NDF-v1.0.0-Skills-Pack lokal bereitgestellt (38 Skills, 39 Dateien, byte-identisch, Commit `9dcadc1`), Provenance + Lock erstellt. Nova-Bewertung: `GO`.
- `CO-WP-002` (docs-only): Nach fail-closed-Blocker und Source-Handoff Concept v3.0 vollständig registriert (57 Abschnitte), Decision Classification + Decision Index + Risk Register erstellt. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-003` (docs-only): Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (Proposed); CCR-02/04/10/11 aufgelöst, NDF-Level geklärt, Repository verifiziert. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004` (docs-only): Foundation Capability Matrix (74 Capabilities) und initiale Observe-Supportgrenze erstellt (Proposed); drei Statusdimensionen, Support-Evidence, CCR-12 vorgeschlagen aufgelöst. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004A` (docs-only / gov-baseline): Concept-v3.1-Amendment, Sovereignty-and-Dependency-Policy und BSI-Alignment-Positioning erstellt; Souveränität und BSI-Orientierung als Accepted Product Direction; keine Zertifizierung/VS-Eignung; Standard-/Hardened-/Government-Profile; Lessons-Learned-/NDF-Feedback-Richtung; ITIL/PRINCE2 Kandidaten. Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B` (docs-only / governance): Lessons-Learned-Prozess und NDF-Feedback-Prozess etabliert; 15 Lessons erfasst; 7 reservierte NDF-Kandidaten bewertet (alle `ready-for-bundling`). Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B1` (docs-only / transfer-preparation): Transfer Package 001 mit 3 Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer`. Erster NDF Intake-Versuch fail-closed blockiert.
- `CO-WP-004B2` (docs-only / governance-correction): Package Status und alle 7 Human-Maintainer-Gates auf `approved`; neue Lesson LL-016 erfasst.
- `CO-WP-004B3` (docs-only / cross-project-traceability): Transfer Package 001 dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e); alle 7 Kandidaten auf `transferred-to-ndf`.
- `CO-WP-004B4` (docs-only / cross-project adoption traceability): Alle 7 Kandidaten auf `adopted-in-ndf` über drei Adoption-WPs (Commits 1ebffa6, e894c6f, ebf716c); Transfer Package 001 geschlossen; NDF-Release-Zuordnung bleibt offen. Nova Review ausstehend.
- `CO-WP-004C` (docs-only / security-governance baseline): BSI- und Public-Sector-Readiness-Baseline (18 PSR-Domänen), Reference-/Claims-Register und Public-Sector-Profil erstellt; Decision Index +7, Risk Register +10; BSI-Positionierung additiv aktualisiert. Keine Zertifizierung/Compliance/Behördenfreigabe behauptet; Capability Matrix unverändert; keine ADR. Nova Review ausstehend.
- `CO-WP-004D` (docs-only / governance-framework applicability review): ITIL `adopted-with-tailoring`, PRINCE2 Version 7 `optional-profile`; ITIL 4 & Version 5 mit Versionsgrenzen; drei interne Governance-Profile; NDF bleibt primär; Decision Index +8 (DEC-S-23…30), Risk Register +11 (RISK-50…60). Keine Zertifizierung/Endorsement/Tool-Abhängigkeit; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-004E` (docs-only / capability-governance alignment): Foundation Capability Matrix um fünf Statusdimensionen, Profile, PSR-Mapping und Responsibility-Codes erweitert (94 Capabilities; Zählkorrektur von „74"); Alignment-Dokument + Capability-Matrix-Spezifikation erstellt; Decision Index +7 (DEC-S-31…37), Risk Register +6 (RISK-61…66). PSR-Mapping ≠ BSI-Compliance; keine Capability implementiert; keine ADR; Lessons-Learned unverändert. Nova Review ausstehend.
- `CO-WP-005` (docs-only / repository-governance foundation): Language Standard, Public-Neutrality-and-Disclosure-Policy und Repository-Governance-Standard erstellt; Englisch kanonisch (maschinenbezogen), DE/EN Produkt; Neutralitäts-/Disclosure-Grenzen; Source-of-Truth-Hierarchie; Human-Maintainer-Gates; UTF-8/Zeilenenden; PowerShell-Korrekturstandard. Decision Index +6 (DEC-S-38…43), Risk Register +13 (RISK-67…79). Keine ADR; keine automatisierte Durchsetzung; Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
- `CO-WP-006` (docs-only / architecture-context foundation): System Context and External Boundaries, Plane Taxonomy (10 Planes) und Trust/Deployment/Execution Boundaries erstellt; Produkt- ≠ Deployment- ≠ Managed-Grenze; 20 externe Systemklassen; Kontrollautorität; Connected/Restricted/Offline-Modi; zwei Mermaid-Diagramme (conceptual). Decision Index +10 (DEC-S-44…53), Risk Register +14 (RISK-80…93). Keine Technologie/Architektur/Threat Model; keine ADR; Capability Matrix + Lessons-Learned unverändert; keine NDF-Rückführung. Nova Review ausstehend.
