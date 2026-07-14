# CoreOps – NDF Feedback Candidates

> Prozessreferenz: [NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md)
> Erzeugt durch `CO-WP-004B` (docs-only / governance).

**Statuswerte, die der Implementation Agent setzen darf:** `candidate` · `validation-required` · `ready-for-bundling` · `project-local-only` · `deferred` · `duplicate`.
**Ausschließlich durch Nova Review + Human-Maintainer-Entscheidung:** `approved-for-transfer` · `transferred-to-ndf` · `adopted-in-ndf` · `rejected` · `superseded`.

> **Bestätigung: Kein Kandidat in diesem Dokument wurde an das NDF übertragen. Kein Kandidat ist `transferred-to-ndf` oder `adopted-in-ndf`.**
> **CO-WP-004B1-Update:** Nova hat die Übergabeschwelle als erreicht bestätigt und die Vorbereitung eines ersten Transferpakets für alle 7 Kandidaten autorisiert (siehe [NDF_FEEDBACK_TRANSFER_PACKAGE_001.md](../docs/governance/NDF_FEEDBACK_TRANSFER_PACKAGE_001.md)). Alle 7 Kandidaten stehen daher auf `approved-for-transfer` mit `Nova Gate: approved` und `Human-Maintainer Gate: pending until commit`. Der Human-Maintainer-Commit dieses Work Packages bildet das zweite Transfer-Gate — er bestätigt die Transferfreigabe, **nicht** den tatsächlichen NDF-Transfer. `NDF Work Package` und `NDF Adoption Version` bleiben leer; `Backlink Status` ist höchstens `transfer-package-prepared`.

## NDF-FC-COREOPS-001

- **Candidate ID:** NDF-FC-COREOPS-001
- **Source Lesson ID:** LL-001
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-001
- **Category:** Process / Prompt Governance
- **Title:** Git Read versus Git Write in Work-Package-Prompts
- **Generalized Observation:** Work-Package-Prompts, die Git-Interaktion pauschal verbieten, ohne zwischen lesenden und schreibenden Operationen zu unterscheiden, erzeugen vermeidbare Abweichungsmeldungen bei harmlosen read-only-Prüfungen (z. B. Status-/Log-Abfragen zur Scope-Verifikation).
- **Cross-Project Impact:** Betrifft jedes NDF-basierte Projekt, das Implementation Agents mit read-only-Verifikationspflichten, aber ohne Git-Schreibrechte einsetzt.
- **Recommended NDF Change:** Standard-Prompt-Bausteine für Work Packages um eine explizite Zweiteilung „Erlaubte read-only-Git-Befehle" / „Verbotene Git-Schreibbefehle" ergänzen (bereits als De-facto-Muster in allen CoreOps-Work-Packages ab CO-WP-001A umgesetzt).
- **Potential NDF Target Area:** NDF-Skill `ndf-work-package-runner` bzw. allgemeine Prompt-Vorlagen für `docs-only`/`review-only`-Work-Packages.
- **Evidence:** CO-WP-001 Rückmeldung an Nova, Abweichungsabschnitt; Nova-Bewertung `GO WITH NOTES` mit genau dieser Note.
- **Security Relevance:** no
- **Privacy Review:** PASS — keine privaten Daten enthalten.
- **Public-Neutrality Review:** PASS — ausschließlich generisches Prozessmuster, kein Projektinterna.
- **Existing NDF Rule Check:** nicht durchgeführt (keine Netzwerk-/NDF-Repository-Recherche in diesem Work Package; zu prüfen bei tatsächlichem Transfer).
- **Duplicate Check:** kein bekanntes Duplikat innerhalb der CoreOps-lokalen Register.
- **Breaking-Change Potential:** gering — additive Klarstellung bestehender Prompt-Konventionen, keine Verhaltensänderung bestehender Skills.
- **Suggested Bundle:** Bundle „Prompt Governance Patterns" (mit NDF-FC-COREOPS-004).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001A
- **NDF Adoption Commit:** 1ebffa6
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001A (Commit 1ebffa6). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-002

- **Candidate ID:** NDF-FC-COREOPS-002
- **Source Lesson ID:** LL-002, LL-003
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-001A
- **Category:** Tooling / Skills Governance
- **Title:** Vollständige lokale Skills-Verfügbarkeit bei selektiver Aktivierung
- **Generalized Observation:** Ein vollständiger, byte-identisch und provenance-gesicherter lokaler Skills-Bestand kann bereitgestellt werden, ohne dass jedes Work Package alle Skills lädt; Selektionskriterien (Work-Package-Typ, Scope, Sicherheitsgrenzen) steuern die tatsächliche Nutzung.
- **Cross-Project Impact:** Relevant für jedes NDF-Projekt mit größerem lokalem Skills-Pack, bei dem Context Economy und vollständige Verfügbarkeit gemeinsam gewährleistet werden sollen.
- **Recommended NDF Change:** Referenzmuster „vollständiger Import + Provenance/Lock + selektive Aktivierung" als dokumentiertes Skills-first-Onboarding-Verfahren in den NDF-Kern aufnehmen.
- **Potential NDF Target Area:** NDF Skills-Onboarding-Dokumentation / `ndf-skill-quality-reviewer`-Umfeld.
- **Evidence:** CO-WP-001A Rückmeldung „Skills-first-Nutzungsmodell"; `project-system/NDF_SKILLS_PROVENANCE.md`, `project-system/ndf-skills-lock.json`.
- **Security Relevance:** yes (Provenance/Hash-Lock als Supply-Chain-Kontrolle)
- **Privacy Review:** PASS — keine privaten Daten.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** nicht durchgeführt (kein Netzwerkzugriff in diesem Work Package).
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering — beschreibt ein bereits kompatibles Vorgehen, keine Änderung an bestehenden NDF-Skill-Definitionen.
- **Suggested Bundle:** Bundle „Skills Onboarding & Provenance".
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001B
- **NDF Adoption Commit:** e894c6f
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Bündelt zwei verwandte Lessons (Selektionsmodell + Provenance/Lock). Teil von Transfer Package 001, Bundle 2 (Skills Availability and Context Economy). Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001B (Commit e894c6f). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-003

- **Candidate ID:** NDF-FC-COREOPS-003
- **Source Lesson ID:** LL-004
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-002
- **Category:** Process / Source Governance
- **Title:** Source-Handoff für Chat- und externe Dokumente
- **Generalized Observation:** Umfangreiche, im Chat referenzierte externe Quelldokumente benötigen einen expliziten, lokal verifizierbaren Source-Handoff (Datei mit Identitätsprüfung, Vollständigkeitsprüfung, Trunkierungsprüfung), bevor eine „vollständige Registrierung" behauptet werden darf.
- **Cross-Project Impact:** Betrifft jedes Projekt, in dem umfangreiche Ausgangsdokumente (Konzepte, Spezifikationen) per Chat übergeben und anschließend in ein Repository überführt werden sollen.
- **Recommended NDF Change:** Standardverfahren „Source Preflight" (Identität, Abschnittsvollständigkeit, Trunkierungserkennung, unveränderter Quellstatus) als NDF-Referenzmuster für Registrierungs-Work-Packages dokumentieren.
- **Potential NDF Target Area:** NDF-Skill/Vorlage für „docs-only" Registrierungs-Work-Packages, ggf. `ndf-existing-project-analysis-runner`-Umfeld.
- **Evidence:** CO-WP-002 Blocked-Report; CO-WP-002-Fortsetzungsprompt „Source Preflight"-Kriterien; erfolgreiche Anwendung dokumentiert in der CO-WP-002-Rückmeldung.
- **Security Relevance:** no
- **Privacy Review:** PASS — Quelle selbst bleibt außerhalb des Kandidaten; nur das Verfahren wird generalisiert.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** nicht durchgeführt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering — zusätzliche Prüfschritte vor Registrierung, keine Änderung bestehender Abläufe.
- **Suggested Bundle:** Bundle „Prompt Governance Patterns" (mit NDF-FC-COREOPS-004).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001A
- **NDF Adoption Commit:** 1ebffa6
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001A (Commit 1ebffa6). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-004

- **Candidate ID:** NDF-FC-COREOPS-004
- **Source Lesson ID:** LL-005
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-002
- **Category:** Process / Blocker Handling
- **Title:** Blocked Report ohne künstlichen Zwischen-Commit
- **Generalized Observation:** Ein fail-closed-Stopp ohne jegliche Dateiänderung ist ein vollständiges, review-fähiges Ergebnis, wenn er strukturiert (Grund, betroffene Prüfung, Entsperrungsbedingung) dokumentiert wird; ein Zwischen-Commit „für nichts" ist nicht nötig.
- **Cross-Project Impact:** Betrifft jeden NDF-Workflow mit Fail-Closed-Prinzip und Human-Maintainer-Commit-Gate.
- **Recommended NDF Change:** „Blocked Report"-Struktur als Standardformat in die NDF-Prozessdokumentation aufnehmen (Grund, Preflight-Ergebnis, benötigte Entsperrung, keine Dateiänderung).
- **Potential NDF Target Area:** NDF-Kernprinzipien „Fail Closed" / Work-Package-Runner-Dokumentation.
- **Evidence:** CO-WP-002-Fortsetzungsprompt Abschnitt 1; Nova-Bewertung „GO – Blocker bestätigt".
- **Security Relevance:** no
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** nicht durchgeführt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering.
- **Suggested Bundle:** Bundle „Prompt Governance Patterns" (mit NDF-FC-COREOPS-001, 003).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001A
- **NDF Adoption Commit:** 1ebffa6
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001A (Commit 1ebffa6). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-005

- **Candidate ID:** NDF-FC-COREOPS-005
- **Source Lesson ID:** LL-011
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004A
- **Category:** Governance / Decision Modeling
- **Title:** Accepted Product Direction ohne technische Festlegung
- **Generalized Observation:** Strategische Produktrichtungen (Souveränität, Sicherheitsausrichtung) lassen sich mit einem eigenen Statuswert „Accepted Product Direction" registrieren, der ausdrücklich von technischer Architektur-/Technologie-Akzeptanz getrennt ist, wodurch frühe strategische Klarheit ohne verfrühte technische Festlegung möglich wird.
- **Cross-Project Impact:** Relevant für jedes Projekt, das strategische Leitplanken vor abschließender technischer Architektur formalisieren will.
- **Recommended NDF Change:** Statuswert-Familie „Accepted Product Direction" mit definierten Unterwerten (`prohibited`, `not-claimed`, `foundation-candidate`, `binding-governance-direction` u. a.) als optionales NDF-Decision-Index-Muster dokumentieren.
- **Potential NDF Target Area:** NDF Decision-Index-Vorlage / Governance-Dokumentation.
- **Evidence:** `project-system/DECISION_INDEX.md` Abschnitt „Product Direction and Governance Decisions"; `docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md`.
- **Security Relevance:** no
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS — Muster ist abstrakt, keine CoreOps-spezifischen Inhalte nötig für den NDF-Transfer.
- **Existing NDF Rule Check:** nicht durchgeführt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering bis mittel — neue optionale Statuswert-Familie, keine Änderung bestehender Pflichtfelder.
- **Suggested Bundle:** Bundle „Decision & Status Modeling" (mit NDF-FC-COREOPS-006).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001C
- **NDF Adoption Commit:** ebf716c
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Teil von Transfer Package 001, Bundle 3 (Governance and Status Modeling). Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001C (Commit ebf716c). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-006

- **Candidate ID:** NDF-FC-COREOPS-006
- **Source Lesson ID:** LL-008
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004
- **Category:** Architecture / Status Modeling
- **Title:** Mehrdimensionale Statusmodelle statt überladener Einzelstatus
- **Generalized Observation:** Wo Planung, tatsächlicher Implementierungsfortschritt und Vertrauens-/Supportstatus unabhängig variieren können, verhindert ein einziges Statusfeld pro Element Missverständnisse; drei getrennte, unabhängige Dimensionen (z. B. Roadmap/Implementation/Support) sind robuster.
- **Cross-Project Impact:** Allgemein anwendbar auf jede Capability-, Feature- oder Integrationsmatrix in NDF-basierten Projekten.
- **Recommended NDF Change:** Mehrdimensionales Statusmodell als empfohlenes Muster für Capability-/Feature-Matrizen in der NDF-Dokumentationsvorlage aufnehmen.
- **Potential NDF Target Area:** NDF-Vorlagen für Foundation-/Capability-Dokumentation.
- **Evidence:** `docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md`.
- **Security Relevance:** no
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** nicht durchgeführt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering.
- **Suggested Bundle:** Bundle „Decision & Status Modeling" (mit NDF-FC-COREOPS-005).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001C
- **NDF Adoption Commit:** ebf716c
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Teil von Transfer Package 001, Bundle 3 (Governance and Status Modeling). Im NDF Intake Review als Guidance/optionales Muster eingeordnet. Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001C (Commit ebf716c). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-007

- **Candidate ID:** NDF-FC-COREOPS-007
- **Source Lesson ID:** LL-013
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004A
- **Category:** Governance / Framework Adoption
- **Title:** Framework-Tailoring vor Framework-Übernahme
- **Generalized Observation:** Bei mehreren potenziell relevanten externen Rahmenwerken (z. B. Sicherheitsstandards, Projekt-/Servicemanagement-Frameworks) verhindert eine explizite „Kandidat, kein Accepted"-Phase mit separater Tailoring-Entscheidung Framework-Overload und verfrühte Vollübernahme.
- **Cross-Project Impact:** Relevant für jedes Projekt, das mehrere externe Governance-/Compliance-Rahmenwerke evaluiert.
- **Recommended NDF Change:** „Framework Candidate → Tailoring Decision"-Zweiphasenmodell als NDF-Referenzmuster für Rahmenwerk-Übernahmen dokumentieren.
- **Potential NDF Target Area:** NDF Governance-Dokumentation / ADR-Vorbereitung.
- **Evidence:** `docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md` §11–12; `project-system/RISK_REGISTER.md` RISK-32.
- **Security Relevance:** no
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** nicht durchgeführt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering.
- **Suggested Bundle:** Bundle „Decision & Status Modeling" (thematisch verwandt, aber eigenständig wegen Fokus auf externe Rahmenwerke statt interne Stati).
- **Status:** adopted-in-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **Intake Work Package:** NDF-INTAKE-COREOPS-001
- **Intake Commit:** d08e35e
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **Adoption Work Package:** NDF-ADOPT-COREOPS-001C
- **NDF Adoption Commit:** ebf716c
- **Adoption Status:** adopted
- **NDF Adoption Version:** — (not yet assigned; no NDF release version claimed)
- **Backlink Status:** adoption-recorded
- **Notes:** Tatsächliche Tailoring-Entscheidung für CoreOps selbst erfolgt in `CO-WP-004D`; dieser Kandidat betrifft nur das generalisierte NDF-Prozessmuster. Teil von Transfer Package 001, Bundle 3. Im NDF Intake Review als Guidance/optionales Muster eingeordnet. Adoptiert in den aktuellen NDF-Entwicklungsstand über NDF-ADOPT-COREOPS-001C (Commit ebf716c). Release-Versionszuordnung noch offen.

## NDF-FC-COREOPS-008

- **Candidate ID:** NDF-FC-COREOPS-008
- **Source Lesson ID:** LL-017
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-005…CO-WP-012 (Milestone Review)
- **Category:** Governance / Decision Modeling
- **Title:** Getrennte Entscheidungsdimensionen (Decision Class · Lifecycle Status · Binding Level)
- **Generalized Observation:** Entscheidungsregister bleiben eindeutig auswertbar, wenn Art (Class), Reifegrad (Lifecycle Status) und Verbindlichkeit (Binding Level) als getrennte Dimensionen geführt werden statt als kombinierte Pseudo-Statuswerte.
- **Cross-Project Impact:** Relevant für jedes Projekt mit einem Decision Index / ADR-Vorstufe.
- **Recommended NDF Change:** Getrennte Entscheidungsdimensionen als optionales NDF-Decision-Index-Muster dokumentieren; kombinierte Statuswerte ausdrücklich als Anti-Muster benennen.
- **Potential NDF Target Area:** NDF Governance-/Decision-Index-Vorlagen.
- **Evidence:** `project-system/DECISION_INDEX.md` DEC-S-38…136 (getrennt) vs. DEC-S-01…37 (kombiniert); `project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md` §12.
- **Security Relevance:** partly (Verbindlichkeit von Sicherheitsentscheidungen muss eindeutig lesbar sein).
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** verwandt mit NDF-FC-COREOPS-006 (mehrdimensionaler Capability-Status), aber eigenständig (dort Capabilities, hier Entscheidungen); nicht als bereits verbindlich abgedeckt bekannt.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering (additive Vorlage/Guidance).
- **Suggested Bundle:** potenzielles Bundle „Decision & Status Modeling" (mit NDF-FC-COREOPS-009/010).
- **Status:** candidate-pending-nova-review
- **Nova Gate:** pending
- **Human-Maintainer Gate:** pending
- **Transfer Readiness:** nicht bewertet; kein Transfer gestartet.
- **Intake Work Package:** —
- **Intake Commit:** —
- **NDF Work Package:** —
- **Adoption Work Package:** —
- **NDF Adoption Commit:** —
- **Adoption Status:** not started
- **NDF Adoption Version:** —
- **Backlink Status:** none
- **Notes:** Evidenz aus acht Work Packages (CO-WP-005…012). Keine CoreOps-Technologieentscheidung enthalten. Erwartet Nova-/Human-Maintainer-Bündelentscheidung; kein Transfer/keine Adoption automatisch gestartet.

## NDF-FC-COREOPS-009

- **Candidate ID:** NDF-FC-COREOPS-009
- **Source Lesson ID:** LL-018
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-007…CO-WP-012 (Milestone Review)
- **Category:** Security / Documentation Coherence
- **Title:** Wiederverwendbares Vokabular an Sicherheitsinvarianten (die „≠"-Grenzkette)
- **Generalized Observation:** Ein zentral definiertes, benanntes Set an Sicherheitsinvarianten (read≠write, unknown≠healthy, stale≠current, queued≠executed, executed≠successful, successful≠compliant, evidence-capability≠available≠satisfied), das in Folgedokumenten referenziert statt neu formuliert wird, hält einen großen Dokumentensatz kohärent.
- **Cross-Project Impact:** Relevant für jedes Projekt mit umfangreicher Sicherheits-/Architektur-Dokumentation.
- **Recommended NDF Change:** Ein referenzierbares Sicherheitsinvarianten-Vokabular als NDF-Security-Baseline-Muster dokumentieren; Folgedokumente referenzieren statt Invarianten neu zu formulieren.
- **Potential NDF Target Area:** NDF Security-Baseline / Threat-Model-Vorlagen.
- **Evidence:** `docs/security/COREOPS_FOUNDATION_THREAT_MODEL.md` §14 (17 Invarianten); Wiederverwendung in CO-WP-008/009/010/011/012; `project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md` §6, §10.
- **Security Relevance:** yes.
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** kein bekanntes verbindliches NDF-Äquivalent.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering (additive Guidance).
- **Suggested Bundle:** potenzielles Bundle „Security Coherence" (mit NDF-FC-COREOPS-010).
- **Status:** candidate-pending-nova-review
- **Nova Gate:** pending
- **Human-Maintainer Gate:** pending
- **Transfer Readiness:** nicht bewertet; kein Transfer gestartet.
- **Intake Work Package:** —
- **Intake Commit:** —
- **NDF Work Package:** —
- **Adoption Work Package:** —
- **NDF Adoption Commit:** —
- **Adoption Status:** not started
- **NDF Adoption Version:** —
- **Backlink Status:** none
- **Notes:** Evidenz aus sechs Work Packages (CO-WP-007…012). Invarianten sind Designanforderungen, keine implementierten Kontrollen. Kein Transfer/keine Adoption automatisch gestartet.

## NDF-FC-COREOPS-010

- **Candidate ID:** NDF-FC-COREOPS-010
- **Source Lesson ID:** LL-019
- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-007, CO-WP-006…CO-WP-012 (Milestone Review)
- **Category:** Governance / Risk Modeling
- **Title:** Trennung von Threat-Scenario-Register und Projekt-Risk-Register
- **Generalized Observation:** Individuelle Bedrohungen in einem Threat-Scenario-Register zu führen und sie nicht als eigene Governance-Risiken zu duplizieren, hält das Risk Register fokussiert; ergänzend ist ein periodischer Konsolidierungstakt nötig, damit das Risk Register nicht unwartbar wächst.
- **Cross-Project Impact:** Relevant für jedes Projekt, das sowohl Threat Modeling als auch ein Governance-Risk-Register führt.
- **Recommended NDF Change:** Guidance zur Registertrennung (Threat Scenario ≠ Governance Risk) plus empfohlener Konsolidierungstakt als NDF-Referenzmuster dokumentieren.
- **Potential NDF Target Area:** NDF Risk-/Threat-Management-Vorlagen.
- **Evidence:** `docs/security/THREAT_SCENARIO_REGISTER.md` (THR-001…040); `project-system/RISK_REGISTER.md` RISK-94…103 (Hinweis) und Wachstum 66→189; `project-brain/MILESTONE_REVIEW_CO_WP_005_TO_012.md` §11.
- **Security Relevance:** partly.
- **Privacy Review:** PASS.
- **Public-Neutrality Review:** PASS.
- **Existing NDF Rule Check:** kein bekanntes verbindliches NDF-Äquivalent.
- **Duplicate Check:** kein bekanntes Duplikat.
- **Breaking-Change Potential:** gering (additive Guidance).
- **Suggested Bundle:** potenzielles Bundle „Security Coherence" (mit NDF-FC-COREOPS-009).
- **Status:** candidate-pending-nova-review
- **Nova Gate:** pending
- **Human-Maintainer Gate:** pending
- **Transfer Readiness:** nicht bewertet; kein Transfer gestartet.
- **Intake Work Package:** —
- **Intake Commit:** —
- **NDF Work Package:** —
- **Adoption Work Package:** —
- **NDF Adoption Commit:** —
- **Adoption Status:** not started
- **NDF Adoption Version:** —
- **Backlink Status:** none
- **Notes:** Evidenz aus dem Threat Model (CO-WP-007) und der Registerpflege über CO-WP-006…012. Kein Transfer/keine Adoption automatisch gestartet.

## Zusammenfassung

- **Anzahl Kandidaten:** 10 (NDF-FC-COREOPS-001…007 adoptiert; NDF-FC-COREOPS-008…010 neu aus dem Milestone Review CO-WP-005…012)
- **Status-Verteilung (nach Milestone Review):** `adopted-in-ndf` 7 · `candidate-pending-nova-review` 3 · `transferred-to-ndf` 0
- **Bundles (Transfer Package 001):** Bundle 1 „Work-Package Safety and Source Handling" (001, 003, 004), Bundle 2 „Skills Availability and Context Economy" (002), Bundle 3 „Governance and Status Modeling" (005, 006, 007)
- **Nova Gate:** approved für alle 7. **Human-Maintainer Gate:** approved für alle 7. **Intake Work Package:** NDF-INTAKE-COREOPS-001 für alle 7 (Commit d08e35e).
- **Adoption-Zuordnung:** Adoption A `NDF-ADOPT-COREOPS-001A` (Commit 1ebffa6) — Kandidaten 001, 003, 004. Adoption B `NDF-ADOPT-COREOPS-001B` (Commit e894c6f) — Kandidat 002. Adoption C `NDF-ADOPT-COREOPS-001C` (Commit ebf716c) — Kandidaten 005, 006, 007.
- **Alle sieben Kandidaten wurden über drei geprüfte Human-Maintainer-Commits in den aktuellen NDF-Entwicklungszweig adoptiert.** Release-Versionszuordnung bleibt offen (`not yet assigned`); keine NDF-Version wird behauptet. `adopted-in-ndf` bedeutet **nicht** "in einer veröffentlichten NDF-Version enthalten".
- **Neue Kandidaten aus dem Milestone Review (008, 009, 010):** Status `candidate-pending-nova-review`; Nova Gate und Human-Maintainer Gate `pending`; kein Transfer, keine Adoption und kein NDF-Commit gestartet. Erwartete Nova-/Human-Maintainer-Bündelentscheidung. Die Zeilen 356–359 beziehen sich ausschließlich auf Transfer Package 001 (Kandidaten 001–007).
