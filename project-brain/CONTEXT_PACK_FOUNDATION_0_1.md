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

`CO-WP-004B3 – Record Completed NDF Intake Transfer` (docs-only / cross-project-traceability). Nach Bearbeitung im **Nova Review** (`pending`). Vorheriges WP: `CO-WP-004B2 – GO WITH NOTES`.

## Erstes NDF-Transferpaket (CO-WP-004B1)

- Übergabeschwelle (5–10 Kandidaten) erreicht; Nova autorisiert Vorbereitung für alle 7 reservierten Kandidaten.
- Transfer Package 001 mit 3 Bundles erstellt (Work-Package Safety/Source Handling; Skills Availability/Context Economy; Governance/Status Modeling).
- Alle 7 Kandidaten auf `approved-for-transfer` (Nova Gate: approved; Human-Maintainer Gate: approved, Commit 4ad3111).
- **Kein Kandidat übertragen/adoptiert; kein NDF-Repository verändert; kein NDF-Work-Package erstellt.**

## NDF Intake Approval Gate Finalization (CO-WP-004B2)

- Transfer Package Status auf `Approved for NDF Intake` gesetzt.
- Alle 7 Human-Maintainer-Gates auf `approved` gesetzt (evidenziert durch Commit 4ad3111).
- Neue Lesson LL-016 erfasst: „Commit-gated status transitions must be represented in repository state".
- Lehrpunkt: Commit liefert Evidenz, verändert aber nicht automatisch Statuswerte. Freigabefelder müssen VOR dem Commit gesetzt werden.
- **Kein Kandidat übertragen/adoptiert; kein NDF-Repository verändert; kein NDF-Work-Package erstellt.**

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

`CO-WP-004C – BSI and Public-Sector Readiness Baseline` (planned-next; pending Nova review von CO-WP-004B3 und Human-Maintainer-Freigabe). `CO-WP-005` bleibt `planned-after-CO-WP-004E`. External: NDF Adoption-Entscheidungen ausstehend.

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
- `CO-WP-004B1` (docs-only / transfer-preparation): Übergabeschwelle erreicht; Transfer Package 001 mit 3 Bundles erstellt; alle 7 Kandidaten auf `approved-for-transfer` (Nova Gate: approved; Human-Maintainer Gate: pending until commit); Zusammenfassungsfehler im Lessons-Register korrigiert (keine Umklassifizierung). Kein Kandidat übertragen/adoptiert, kein NDF-Repository verändert, kein Git Write. Nova Review ausstehend; erster NDF Intake-Versuch fail-closed blockiert.
- `CO-WP-004B2` (docs-only / governance-correction): Transfer Package Status und alle 7 Human-Maintainer-Gates auf `approved` gesetzt (Commit 4ad3111 als Evidenz); neue Lesson LL-016 erfasst („Commit-gated status transitions must be represented in repository state"); Lehrpunkt: Statuswerte müssen VOR dem Commit gesetzt werden, nicht durch den Commit. Kein Kandidat übertragen/adoptiert, kein NDF-Repository verändert. Nova Review ausstehend.
