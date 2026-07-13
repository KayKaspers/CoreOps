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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Bündelt zwei verwandte Lessons (Selektionsmodell + Provenance/Lock). Teil von Transfer Package 001, Bundle 2 (Skills Availability and Context Economy). Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Teil von Transfer Package 001, Bundle 1 (Work-Package Safety and Source Handling). Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Teil von Transfer Package 001, Bundle 3 (Governance and Status Modeling). Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Teil von Transfer Package 001, Bundle 3 (Governance and Status Modeling). Im NDF Intake Review als Guidance/optionales Muster eingeordnet. Übertragen zum NDF-Intake; Adoption ausstehend.

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
- **Status:** transferred-to-ndf
- **Nova Gate:** approved
- **Human-Maintainer Gate:** approved
- **Transfer Readiness:** Kriterien 1–12 erfüllt; Transfer durchgeführt über NDF-INTAKE-COREOPS-001 (Commit d08e35e).
- **NDF Work Package:** NDF-INTAKE-COREOPS-001
- **NDF Adoption Version:** —
- **Backlink Status:** intake-review-committed
- **Notes:** Tatsächliche Tailoring-Entscheidung für CoreOps selbst erfolgt in `CO-WP-004D`; dieser Kandidat betrifft nur das generalisierte NDF-Prozessmuster. Teil von Transfer Package 001, Bundle 3. Im NDF Intake Review als Guidance/optionales Muster eingeordnet. Übertragen zum NDF-Intake; Adoption ausstehend.

## Zusammenfassung

- **Anzahl Kandidaten:** 7 (alle reservierten IDs NDF-FC-COREOPS-001…007 bewertet)
- **Status-Verteilung (nach CO-WP-004B3):** `transferred-to-ndf` 7 · `adopted-in-ndf` 0
- **Bundles (Transfer Package 001):** Bundle 1 „Work-Package Safety and Source Handling" (001, 003, 004), Bundle 2 „Skills Availability and Context Economy" (002), Bundle 3 „Governance and Status Modeling" (005, 006, 007)
- **Nova Gate:** approved für alle 7. **Human-Maintainer Gate:** approved für alle 7. **NDF Work Package:** NDF-INTAKE-COREOPS-001 für alle 7.
- **Alle sieben Kandidaten wurden zum NDF-Intake übertragen (Commit d08e35e).** Kein Kandidat ist `adopted-in-ndf`. Die Adoption erfolgt ausschließlich durch künftige NDF-Work-Packages gemäß NDF-seitigem Entscheidungsprozess.
