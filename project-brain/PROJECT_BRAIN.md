# CoreOps – Project Brain

Kompakter, fortschreibbarer Wissensstand des Projekts. Wird pro Work Package aktualisiert.

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist informativ, **nicht** normativ (CCR-04 geklärt: Übernahme aus `main` nur via eigenes freigegebenes WP).

---

## Projektstatus

Foundation 0.1. Concept v3.0 registriert und klassifiziert; Project Brief, Foundation Scope Lock und Release-Taxonomie erstellt (`Proposed for acceptance`). Weiterhin kein Anwendungscode.

## Aktuelle Phase

`Foundation 0.1` (interner Phasenname der dokumentations-/governanceorientierten Foundation; Release-Taxonomie: Foundation-Tag-Kandidat `v0.0.1-foundation`, erster Observe-Prerelease-Kandidat `v0.1.0-alpha.1`).

## Akzeptierte Produktvision

CoreOps als universelle, self-hosted und offline-fähige Operations Control Plane, die Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen modular verbindet. Herkunft: CoreOps Concept v3.0, bereitgestellt am 12. Juli 2026.

## Bindende Prozessgrenzen

- NDF v1.0.0 normativ (Tag/Commit), `main` nicht normativ.
- Human-Maintainer-only für Freigabe, Staging, Commit, Push, Merge, Tags, Releases, Deployments, irreversible/privilegierte Aktionen.
- Kein Code in Foundation.
- Keine autonomen Git-, Netzwerk- oder Infrastrukturaktionen.
- Fail-closed bei Unklarheit.
- Bewertung des Ergebnisses (`GO`/`GO WITH NOTES`/`REWORK`/`SPLIT`/`STOP`) trifft Nova, nicht der Implementation Assistant.

## Architekturstand

Kein verbindlicher Architekturstand. Nur unbestätigte Kandidaten (z. B. modularer Monolith). Nichts finalisiert.

## Technische Entscheidungen

Keine technische Entscheidung getroffen. Technologie-Stack nicht ausgewählt.

## Noch nicht akzeptierte ADR-Kandidaten

Architekturform, Frontend, Backend/API, Datenhaltung, Topologie-Persistenz, Caching/Queue, Agent-Technologie, Observability-Stack, Policy-Engine, Workflow-Engine. **Keiner ist Accepted.** (Docker-first ist in CO-WP-003 als Delivery-/Betriebsanforderung eingeordnet — `proposed` —, nicht als Anwendungsarchitektur; die konkrete Umsetzung bleibt ADR-relevant.)

## Bekannte Risiken

- Breiter Scope mit hohem Sicherheitsanspruch (siehe [RISK_REGISTER.md](../project-system/RISK_REGISTER.md), 39 Foundation-Risiken; RISK-02/04/09/11 mit Impact `critical`).
- Governance-/Scope-Risiken auf `treatment-planned`: RISK-13/14/17/19/20/21 (CO-WP-003), RISK-22/23/24/25 (CO-WP-004), RISK-26…34 (CO-WP-004A) und RISK-35…39 (CO-WP-004B: Transferprozesslast, Security-Learning-Verspätung, Evidenz-Verallgemeinerung, Duplikate, inkonsistente Statuswerte).
- Weiterhin offene Konflikte aus der Concept-Klassifikation: Plane-Taxonomie (CCR-01), Offline-Policy (CCR-05), Machine Identity vs. Air-Gap (CCR-06), privilegierte Ausführung (CCR-07), Audit vs. Datenschutz (CCR-08), Offline-First-Facetten (CCR-09). CCR-12 (Herstellersupport-Grenze) in CO-WP-004 vorgeschlagen aufgelöst.

## Offene Fragen

- Technologie-/Architekturauswahl per ADR (weiterhin vollständig offen).
- Auflösung der verbleibenden Konflikte CCR-01, 05, 06, 07, 08, 09, 12 in den jeweiligen Foundation-WPs.
- Capability Matrix und Support Boundary (→ `CO-WP-004`).

## Lessons Learned

- `CO-WP-001` erhielt von Nova `GO WITH NOTES`. Ursache der Note: eine nicht ausdrücklich dokumentierte read-only-Git-Abweichung (`git status --porcelain` unter einem pauschalen Git-Verbot). Konsequenz: Künftige Work Packages unterscheiden ausdrücklich zwischen **Git Read** (erlaubt) und **Git Write** (verboten ohne Human-Maintainer-Freigabe).
- `CO-WP-001A`: Der vollständige NDF-v1.0.0-Skills-Pack wurde lokal bereitgestellt (byte-identisch, verifiziert) und von Nova mit `GO` bewertet. Skills-first für spätere Work Packages ist damit lokal möglich; Auswahl erfolgt selektiv pro Work Package.
- `CO-WP-002`: Fail-closed-Blocker (fehlende Quelle) war korrekt (`GO – Blocker bestätigt`); nach Bereitstellung der lokalen read-only-Quelle wurde das Concept v3.0 vollständig registriert und mit `GO WITH NOTES` bewertet. Lesson: Quelle als verlässliche lokale Datei vor der Registrierung verifizieren (Identität, Abschnitte, Ende, Trunkierung).
- `CO-WP-003`: Governance-Konsolidierung — Konflikte CCR-02/04/10/11 vorgeschlagen aufgelöst; Zusatzstatus (`proposed`/`clarified`/`verified`) im Decision Index eingeführt, ohne technische Architektur zu akzeptieren. Nova-Bewertung: `GO WITH NOTES`; durch Human-Maintainer-Commit angenommen.
- `CO-WP-004`: Capability Matrix und Observe-Supportgrenze — drei getrennte Statusdimensionen trennen Planung, Implementierung und Support sauber. Lesson: „Capability target ≠ implemented ≠ verified ≠ supported"; Herstellernennung ist kein Support (CCR-12 vorgeschlagen aufgelöst). Nova-Bewertung: `GO WITH NOTES`; durch Human-Maintainer-Commit angenommen.
- `CO-WP-004A`: Strategische Erweiterung registriert (Souveränität, BSI-Orientierung). Lesson: Ausrichtung akzeptieren, ohne Technologie, Konformität oder Zertifizierung zu behaupten; Unabhängigkeit ≠ vollständige Eigenentwicklung; ITIL/PRINCE2/BSI bewusst schlank halten (Framework-Overload-Risiko). Nova-Bewertung: `GO WITH NOTES`; durch Human-Maintainer-Commit angenommen.
- `CO-WP-004B`: Lessons-Learned- und NDF-Feedback-Governance etabliert; 15 Lessons erfasst (13 retrospektiv, 2 aus dem WP selbst), 7 reservierte NDF-Kandidaten bewertet. Lesson: getrennte Statusmodelle für Lessons vs. NDF-Kandidaten und eine explizite Positivliste erlaubter Agent-Status verhindern eine versehentliche Simulation von Freigabeentscheidungen. Nova-Bewertung: `GO WITH NOTES`; durch Human-Maintainer-Commit angenommen.
- `CO-WP-004B1`: Erstes NDF-Transferpaket vorbereitet. Lesson: Eine gemeldete Statistik-Inkonsistenz (LL-PROCESS „8 gemeldet, 9 IDs gelistet") ließ sich auf einen Zusammenfassungsfehler zurückführen (Vermischung von Primär- und Sekundärklasse bei LL-001) — Korrektur betraf nur den Text, keine Lesson wurde umklassifiziert. Konsequenz: Zusammenfassungssätze über klassifizierte Register stets gegen die Einzelfelder gegenprüfen, nicht nur gegen frühere Zusammenfassungen.

## Lokaler NDF-Skills-Bestand

- Vollständiger NDF-v1.0.0-Skills-Pack lokal unter `.claude/skills/` bereitgestellt.
- Normative Quelle: Tag `v1.0.0`, Commit `9dcadc1` (`main` nicht verwendet).
- 38 Skills, 39 Dateien, byte-identisch übernommen.
- Provenance: [NDF_SKILLS_PROVENANCE.md](../project-system/NDF_SKILLS_PROVENANCE.md); Lock: [ndf-skills-lock.json](../project-system/ndf-skills-lock.json).
- Keine Skill-Datei wurde ausgeführt oder verändert.
- Keine automatische Vollaktivierung; Auswahl selektiv pro Work Package.

## Concept-Registrierung und Klassifikation

- Concept v3.0 vollständig registriert unter [docs/architecture/COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md) (alle 57 Abschnitte; ein privater Standortname öffentlich neutralisiert).
- Decision Classification: [CONCEPT_DECISION_CLASSIFICATION.md](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md) — 10 Klassen; technische Aussagen klassifiziert, **nicht** akzeptiert.
- Decision Index ([DECISION_INDEX.md](../project-system/DECISION_INDEX.md)) und Risk Register ([RISK_REGISTER.md](../project-system/RISK_REGISTER.md)) initial vorhanden.
- 32 ADR-Kandidaten erfasst, **kein** ADR Accepted, **keine** ADR-Datei erzeugt.
- 12 offene Konflikte (CCR-01…12) dokumentiert und bleiben offen.

## Brief, Scope Lock und Release-Taxonomie (CO-WP-003)

- Project Brief: [docs/architecture/PROJECT_BRIEF.md](../docs/architecture/PROJECT_BRIEF.md) (Proposed for acceptance).
- Foundation Scope Lock: [docs/governance/FOUNDATION_SCOPE_LOCK.md](../docs/governance/FOUNDATION_SCOPE_LOCK.md) (In/Out of Scope, Source Hierarchy, Change Control, Exit Gates offen).
- Release-Taxonomie: [docs/governance/RELEASE_TAXONOMY.md](../docs/governance/RELEASE_TAXONOMY.md) — Foundation `v0.0.1-foundation`, Observe `v0.1.0-alpha.1` (Proposed; kein Tag/Release erzeugt).
- Docker-first als akzeptierte Delivery-/Betriebsanforderung eingeordnet (keine Anwendungsarchitektur, kein K8s-Zwang, noch nicht implementiert).
- Aktive Queue autoritativ (WORK_PACKAGE_QUEUE.md) gegenüber Concept-§50-Queue geklärt.
- NDF-`main`-Auslegung und NDF-Level-Semantik geklärt; Repository-URL verifiziert (`https://github.com/KayKaspers/CoreOps`).
- **Technische Architektur weiterhin offen; keine Technologie ausgewählt; keine ADR erzeugt/akzeptiert.**

## Capability Matrix und Support Boundary (CO-WP-004)

- Foundation Capability Matrix: [docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md) — 74 Capabilities, 12 Domains; alle `not-implemented`/`not-supported` (Proposed for acceptance).
- Initiale Observe-Supportgrenze: [docs/integrations/INITIAL_SUPPORT_BOUNDARY.md](../docs/integrations/INITIAL_SUPPORT_BOUNDARY.md) — Observe vollständig read-only, max. Integrationslevel 2; Map/Write ausgeschlossen.
- Drei getrennte Statusdimensionen (Roadmap/Implementation/Support), Integrationslevel präzisiert, Herausgeberklassen und 21-Punkte-Support-Evidence-Satz definiert.
- `CCR-12` vorgeschlagen aufgelöst: Herstellernennung = Kandidat, kein Support/Partnerschaft/Zertifizierung.
- **Aktuell keine Runtime-Capability implementiert, keine Integration `supported`. Technische Architektur/Technologie weiterhin offen; keine ADR erzeugt.**

## Souveränität und BSI-Orientierung (CO-WP-004A)

- Concept-v3.1-Amendment ([COREOPS_CONCEPT_V3_1_AMENDMENT.md](../docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md)), Sovereignty-and-Dependency-Policy ([SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](../docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md)) und BSI-Positionierung ([BSI_ALIGNMENT_POSITIONING.md](../docs/security/BSI_ALIGNMENT_POSITIONING.md)) erstellt.
- Produktsouveränität akzeptiert; **keine** externe Managementplattform wird Kernabhängigkeit; technische Basisabhängigkeiten getrennt und offen (keine akzeptiert).
- BSI-orientierte Entwicklung akzeptiert; **keine** Zertifizierung/Konformität/Zulassung/VS-Eignung behauptet.
- Standard-/Hardened-/Government-Profile als Zielmodell registriert (Government = späteres Nachweisprofil, keine Zertifizierung).
- Lessons Learned als Governance-Richtung, kontrollierter NDF-Rückfluss als Kandidatenprozess registriert (Detail in `CO-WP-004B`).
- ITIL und PRINCE2 bleiben Foundation-Kandidaten (Tailoring in `CO-WP-004D`).
- **Capability Matrix unverändert; technische Architektur weiterhin offen; keine ADR.**

## Lessons Learned und NDF Feedback (CO-WP-004B)

- Lessons-Learned-Prozess ([LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md)) und NDF-Feedback-Prozess ([NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md)) etabliert (Proposed for acceptance).
- Retrospektive Ersterfassung: 13 Lessons aus CO-WP-001…004A plus 2 Lessons aus CO-WP-004B selbst (insgesamt 15) im [LESSONS_LEARNED_REGISTER.md](../project-system/LESSONS_LEARNED_REGISTER.md).
- 7 reservierte NDF-Feedback-Kandidaten (NDF-FC-COREOPS-001…007) bewertet, alle `ready-for-bundling`, im [NDF_FEEDBACK_CANDIDATES.md](../project-system/NDF_FEEDBACK_CANDIDATES.md).
- **Kein Kandidat wurde an das NDF übertragen oder adoptiert; kein NDF-Repository verändert; kein NDF-Work-Package erstellt.**
- Transferentscheidungen (`approved-for-transfer` und höher) bleiben ausschließlich Nova Review + Human-Maintainer-Gate vorbehalten.

## Erstes NDF-Transferpaket (CO-WP-004B1)

- Nova hat die Übergabeschwelle (5–10 Kandidaten) als erreicht festgestellt und die Vorbereitung für alle 7 reservierten Kandidaten autorisiert.
- Alle 7 Kandidaten auf `approved-for-transfer` (Nova Gate: approved; Human-Maintainer Gate: approved, Commit 4ad3111).

## NDF Intake Approval Gate Finalization (CO-WP-004B2)

- Transfer Package Status auf `Approved for NDF Intake` gesetzt; alle 7 Human-Maintainer-Gates auf `approved` gesetzt.
- Neue Lesson LL-016 erfasst: „Commit-gated status transitions must be represented in repository state".

## NDF Intake Transfer Recording (CO-WP-004B3)

- Transfer Package 001 wurde dem NDF-Intake übertragen (NDF-INTAKE-COREOPS-001, Commit d08e35e).
- Alle 7 Kandidaten auf `transferred-to-ndf` gesetzt (NDF Work Package: NDF-INTAKE-COREOPS-001).

## NDF Adoption Completion and Transfer Package Closure (CO-WP-004B4)

- Alle 7 Kandidaten von `transferred-to-ndf` auf `adopted-in-ndf` gesetzt, verteilt auf drei Adoption-Work-Packages:
  - Adoption A (`NDF-ADOPT-COREOPS-001A`, Commit 1ebffa6): 001, 003, 004.
  - Adoption B (`NDF-ADOPT-COREOPS-001B`, Commit e894c6f): 002.
  - Adoption C (`NDF-ADOPT-COREOPS-001C`, Commit ebf716c): 005, 006, 007.
- Transfer Package 001 geschlossen (`Closed – all candidates processed`).
- **Release-Zuordnung bleibt `not yet assigned`; keine NDF-Version behauptet. `adopted-in-ndf` ≠ in veröffentlichter NDF-Version enthalten.**
- NDF-Repository unverändert; keine weiteren Adoption-WPs gestartet.

## BSI and Public-Sector Readiness Baseline (CO-WP-004C)

- Drei neue Dokumente: Readiness-Baseline ([BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](../docs/security/BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md), 18 PSR-Domänen), Reference-/Claims-Register ([BSI_REFERENCE_AND_CLAIMS_REGISTER.md](../docs/security/BSI_REFERENCE_AND_CLAIMS_REGISTER.md)), Public-Sector-Profil ([PUBLIC_SECTOR_READINESS_PROFILE.md](../docs/governance/PUBLIC_SECTOR_READINESS_PROFILE.md)).
- Interne Profile Standard ⊂ Hardened ⊂ Government (keine Zertifizierung). Verantwortungs- und Evidenzmodell (capability ≠ available ≠ satisfied). Offline/Air-Gap und Souveränität/Cloud (C5/C3A bedingt) eingeordnet.
- Governance: Decision Index +7 (DEC-S-16…22), Risk Register +10 (RISK-40…49). BSI-Positionierung additiv aktualisiert.
- **Keine BSI-Zertifizierung/Vollkonformität/Behörden-/VS-NfD-Freigabe behauptet; keine einzelnen BSI-Anforderungen erfunden; Capability Matrix unverändert (→ CO-WP-004E); keine ADR; keine Technologie ausgewählt.**

## ITIL and PRINCE2 Applicability and Tailoring (CO-WP-004D)

- Zwei neue Dokumente: [ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](../docs/governance/ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md) und [COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md](../docs/governance/COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md).
- ITIL `adopted-with-tailoring` (Service-Management-Guidance); PRINCE2 Version 7 `optional-profile`; vollständige Implementierung beider `rejected`. ITIL 4 und Version 5 mit expliziten Versionsgrenzen (Bridge-Pfad anerkannt).
- NDF bleibt primäres Entwicklungs-/Repository-Governance-Framework; Konflikthierarchie und Overload-Guards dokumentiert; keine parallele Governance.
- Drei interne, optionale Governance-Profile (Service Operations / Major Deployment Project / Public-Sector Delivery), keine Zertifizierungsstufen.
- Decision Index +8 (DEC-S-23…30, löst DEC-S-10/11), Risk Register +11 (RISK-50…60).
- **Keine Zertifizierung/Endorsement, keine Tool-Abhängigkeit, keine ADR; Capability Matrix und Lessons-Learned-Register unverändert; keine NDF-Rückführung gestartet.**

## Capability Matrix Security and Governance Alignment (CO-WP-004E)

- Zwei neue Dokumente: [CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md](../docs/security/CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md) und [CAPABILITY_MATRIX_SPEC.md](../docs/project-system/CAPABILITY_MATRIX_SPEC.md) (letztere existierte lokal nicht, neu erstellt).
- Foundation Capability Matrix additiv erweitert: fünf getrennte Statusdimensionen (Roadmap/Implementation/Support/Evidence/Security-Governance), Profile Relevance, PSR-01…18-Zuordnung, Responsibility-Codes (P/O/S) — für alle Capabilities; keine gelöscht/umbenannt/hochgestuft.
- **Zählkorrektur:** Matrix enthält **94** Capabilities (grep-verifiziert); frühere Summe „74" korrigiert; andere „74"-Referenzen später abzugleichen (an Nova gemeldet).
- PSR-Mapping = Readiness-Relevanz, **nicht** BSI-Compliance; kein detailliertes Control-Mapping; Evidence durchgehend `not-assessed`; kein `compliant`/`requirement-satisfied`.
- Decision Index +7 (DEC-S-31…37, erweitert DEC-O-19), Risk Register +6 (RISK-61…66).
- **Keine Capability implementiert; keine ADR; keine Technologieauswahl; Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Language Standard, Public Neutrality and Repository Governance (CO-WP-005)

- Drei neue Governance-Dokumente: [COREOPS_LANGUAGE_STANDARD.md](../docs/governance/COREOPS_LANGUAGE_STANDARD.md), [PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](../docs/governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md), [REPOSITORY_GOVERNANCE_STANDARD.md](../docs/governance/REPOSITORY_GOVERNANCE_STANDARD.md).
- Sprachstandard: Englisch kanonisch für maschinenbezogene Bezeichner/IDs/Commit-Messages; DE/EN primäre Produktsprachen; Translation-Status-Modell; semantische Parität Pflicht; keine automatische Paritätsbehauptung.
- Public Neutrality: Hersteller-/Organisations-/institutionelle Neutralität; Nennung ≠ Endorsement; Public-Sector ≠ Behördenfreigabe. Disclosure: keine Secrets/personenbezogenen/privaten Infrastrukturdaten; synthetische Beispieldaten; Redaction.
- Repository Governance: Human-Maintainer-Gates erhalten; Source-of-Truth-Hierarchie (erweitert FOUNDATION_SCOPE_LOCK, „summary must not override source"); Dokumentstatus/Supersession; Dateinamen/stabile IDs; UTF-8/Zeilenenden; PowerShell-Korrekturstandard; Public-Hygiene.
- Decision Index +6 (DEC-S-38…43, getrennte Dimensionen), Risk Register +13 (RISK-67…79).
- **Keine automatisierte Durchsetzung; keine README/.gitignore/.gitattributes-Änderung; keine Übersetzung; keine ADR; Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## System Context, Plane Taxonomy and External Boundaries (CO-WP-006)

- Drei neue Dokumente: [SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md](../docs/architecture/SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md), [COREOPS_PLANE_TAXONOMY.md](../docs/architecture/COREOPS_PLANE_TAXONOMY.md), [TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md](../docs/security/TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md).
- Systemkontext: 7 Akteure/Rollen; Produkt- ≠ Deployment- ≠ Managed-Grenze; 20 externe Systemklassen (alle optional); 15 Interaktionsklassen (read ≠ write); Kontrollautorität-Zustände; Connected/Restricted/Offline-Modi; Failure-Grundregeln (unknown ≠ healthy usw.); 13 Datenklassen; Mermaid-Kontextdiagramm.
- Plane Taxonomy: 10 logische Planes; plane ≠ Deployment-Einheit; Edge/Agent Plane optional (agentless möglich); Managed Resource Plane außerhalb Produktgrenze; Mermaid-Plane-Diagramm; Deployment-Unabhängigkeit.
- Trust Boundaries: 11 Vertrauensgrenzen; Policy-to-Action und Control-to-Execution getrennt; fail-closed; Online-zu-Offline-Import mit Provenance; Threat Model deferred zu CO-WP-007.
- Decision Index +10 (DEC-S-44…53), Risk Register +14 (RISK-80…93).
- **Keine Technologie/Architektur/Threat Model ausgewählt; keine ADR; Capability Matrix unverändert; Lessons-Learned-Register unverändert; keine NDF-Rückführung.**
- Pfadhinweis: Scope Lock liegt unter `docs/governance/FOUNDATION_SCOPE_LOCK.md` (Prompt nannte `project-system/…`).

## Threat Model and Trust Boundaries (CO-WP-007)

- Zwei neue Dokumente: [COREOPS_FOUNDATION_THREAT_MODEL.md](../docs/security/COREOPS_FOUNDATION_THREAT_MODEL.md) und [THREAT_SCENARIO_REGISTER.md](../docs/security/THREAT_SCENARIO_REGISTER.md); Trust-Boundary-Dokument additiv erweitert (TB-IDs + Threat-Verweise).
- 24 Assets (AST-01…24), 16 Threat Actors, 18 Threat-Kategorien; Angriffsflächen für alle 10 Planes; Trust-Boundary-Analyse TB-01…11; 40 Threat Scenarios (THR-001…040) mit stabilen IDs und qualitativen evidence-bounded Ratings; 17 Sicherheitsinvarianten; 5 Abuse Cases (AB-1…5); 2 Mermaid-Diagramme (conceptual).
- **Threat Model = Bedrohungen + Sicherheitsanforderungen, keine implementierten/validierten Kontrollen. Invarianten = Designanforderungen. Kein `mitigated`/`closed` ohne Evidenz (aktuell keiner).**
- Decision Index +10 (DEC-S-54…63), Risk Register +10 (RISK-94…103, nur Threat-Model-Governance; einzelne Threats im Register, nicht dupliziert).
- **Kein Pentest/Scan; keine Technologie-/Krypto-/Authentisierungsauswahl; keine ADR; Capability Matrix + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Architecture and Module Boundaries (CO-WP-008)

- Drei neue Dokumente: [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](../docs/architecture/COREOPS_LOGICAL_MODULE_ARCHITECTURE.md), [COREOPS_MODULE_CATALOG.md](../docs/architecture/COREOPS_MODULE_CATALOG.md), [MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md](../docs/architecture/MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md).
- 17 logische Module (MOD-*) mit stabilen IDs, Klassifikation, Verantwortungen, Daten-/Zustandsownership (ein Owner pro Konzept), erlaubten/verbotenen Abhängigkeiten, Threat-/Invarianten-Referenzen; 2 Mermaid-Diagramme (conceptual).
- Policy/Control/Execution getrennt; Experience löst Execution nicht direkt aus; Adapter umgehen Governance nicht; Agenten optional (agentless möglich); Offline-Intake keine direkte Ausführung; Evidence ohne Execution-Autorität; Notification ≠ Kommandokanal; Plugins umgehen Verträge nicht.
- **module ≠ microservice/process/container/deployment unit; keine Technologie-/Protokoll-/Deployment-Auswahl; Invarianten = Designanforderungen, keine implementierten Kontrollen.**
- Decision Index +12 (DEC-S-64…75), Risk Register +15 (RISK-104…118).
- **Keine ADR; Capability Matrix + Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Human Identity, Workspaces, RBAC and Break Glass (CO-WP-009)

- Drei neue Dokumente: [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](../docs/security/HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../docs/security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md), [BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](../docs/security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md).
- Getrennte Konzepte person/identity/account/principal; Repository-Autorität ≠ Runtime-Autorität; Auth ≠ Authz; Account-Lifecycle; Sessions gebunden an aktuellen Status; Delegation explizit/non-transitive; SoD mit Kleininstallations-Ausnahmen.
- Workspace ≠ Security-Tenant; deny-by-default/least-privilege/scope-bound RBAC; Permission-Taxonomie (19 Aktionen); 8 Scope-Typen; Membership ohne globale Autorität; Cross-Workspace explizit/auditierbar.
- Break Glass benannt/temporär/reason-/scope-bound/auditiert mit Ablaufpflicht und verpflichtendem Post-Event Review; Offline-Emergency governed (nicht anonym); keine Dauer-Admin-Rolle.
- **14 Security-Invarianten (Designanforderungen, keine Kontrollen); keine Auth-/IdP-/Session-/MFA-/Policy-Engine-Auswahl.**
- Decision Index +14 (DEC-S-76…89), Risk Register +18 (RISK-119…136).
- **Keine ADR; Modul-/Capability-/Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Machine Identity, Enrollment and Offline Credential Lifecycle (CO-WP-010)

- Drei neue Dokumente: [MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md](../docs/security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md](../docs/security/MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md), [OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md](../docs/security/OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md).
- Human ≠ Machine Identity; Identity ≠ Credential; Discovery ≠ Enrollment; Registration ≠ Trust; Enrollment ≠ Write Authority. 10 Principal-Klassen (managed-resource/agent/relay/adapter/integration/automation-client/deployment-runner/observation-collector/offline-transfer/evidence-export).
- Lifecycle (discovered…archived); Bootstrap ≠ permanente Identität; Trust nicht aus IP/Hostname/MAC/Netzposition/Self-Assertion. Rotation/Renewal ohne stille Scope-Erweiterung; Revocation nicht nur Doku-Status; Compromise → explizite Re-Enrollment-Entscheidung; Decommissioned-IDs nicht still wiederverwendet.
- **Credential-Governance ≠ Rohsecret-Ownership (Speicherung deferred); Offline-Enrollment/-Credentials mit Provenance/Integrität/Approval; Offline-Revocation als anerkannte Herausforderung. 12 Machine-Identity- + weitere Invarianten (Designanforderungen).**
- Decision Index +16 (DEC-S-90…105), Risk Register +19 (RISK-137…155).
- **Keine PKI-/Krypto-/Protokoll-/Secret-Store-Auswahl; keine ADR; Human-Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Source of Truth and Field Provenance (CO-WP-011)

- Drei neue Dokumente: [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](../docs/architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](../docs/architecture/FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), [OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md](../docs/security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md).
- SoT ≠ System of Record; 10 Authority-Klassen; autoritative Modulownership (ein Owner pro Feldkonzept). Desired/Observed/Effective/Derived/Cached getrennt; Effective ≠ compliant. Feldidentität stabil/canonical (≠ UI-Label/Adapter-Feld); 22-Feld-Provenance-Metadaten; Freshness/Trust/Validation getrennt; Transformation-Lineage erhält Input-Autorität.
- Konfliktmodell (7 States); kein universelles Last-Write-Wins; neuester Timestamp gewinnt nicht automatisch; Konflikte bleiben sichtbar; widerrufene Quellen nicht autoritativ; importierte Daten erben keine Autorität; Overrides mit Provenance-Erhalt; Offline-Reconciliation fail-closed; Zeit-/Sequenzunsicherheit sichtbar; Partial Import ≠ complete.
- **12 Security-Invarianten (Designanforderungen); Audit-History nicht durch Reconciliation umgeschrieben.**
- Decision Index +16 (DEC-S-106…121), Risk Register +17 (RISK-156…172).
- **Keine Storage-/DB-/Merge-/CRDT-/Messaging-/Sync-/Krypto-Provenance-Auswahl; keine ADR; Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**

## Observed, Desired, Effective State and Drift (CO-WP-012)

- Drei neue Dokumente: [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](../docs/architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](../docs/architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md), [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../docs/security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md).
- Desired/Observed/Reported/Effective/Last-Known getrennt; Effective State bleibt indeterminate/conflicted bei unklarer Autorität. Keine Beobachtung ≠ keine Drift; kein erkannter Drift ≠ Compliance; stale ≠ current.
- Drift-Definition + 12 Arten + 14 Detection States; Exceptions explizit/reviewbar; Konvergenz (`verified-converged` braucht Evidenz) ≠ Compliance. Detection/Recommendation/Plan/Approval/Execution/Verification strikt getrennt; Drift-Erkennung gewährt keine Write Authority; Executed ≠ successful ≠ verified ≠ compliant; Partial Failure sichtbar; Rollback braucht Verifikation; fail-closed.
- **Auto-Remediation nicht ausgewählt/implementiert; 12+ Security-Invarianten (Designanforderungen).**
- Decision Index +15 (DEC-S-122…136), Risk Register +17 (RISK-173…189).
- **Keine Engine-/Scheduler-/Queue-/Policy-Engine-/Auto-Remediation-Auswahl; keine ADR; SoT-/Identity-/Modul-/Capability-/Threat-Dateien + Lessons-Learned-Register unverändert; keine NDF-Rückführung.**
- **Milestone:** CO-WP-005…012 = acht WPs → Milestone-Lessons-Review-Eignung `yes` (Bündelentscheidung Nova/Human Maintainer; nicht automatisch gestartet).

## Milestone Lessons Review (CO-WP-005…012)

- Gebündelter docs-only Review der acht Foundation-WPs abgeschlossen; Nova Review pending. Neues Dokument [MILESTONE_REVIEW_CO_WP_005_TO_012.md](MILESTONE_REVIEW_CO_WP_005_TO_012.md). Ergebnis: **GO WITH NOTES FOR CO-WP-013**.
- Sechs konsolidierte Lessons LL-017…022 (Lessons gesamt 22); drei NDF-Feedback-Kandidaten NDF-FC-COREOPS-008…010 (`candidate-pending-nova-review`).
- Read-only-Befunde (nur als Follow-up dokumentiert, nichts umgesetzt): Risk Register 189 Einträge → Konsolidierungslauf ~CO-WP-029/030; Decision Index mischt Alt-Kombistatus (DEC-S-01…37) mit getrennten Dimensionen (DEC-S-38…136); Capability-Count „74→94" in Alt-Abschnitten.
- **Risk Register und Decision Index unverändert (read-only); kein NDF-Transfer/keine Adoption; CO-WP-013 nicht begonnen.**
- Status: `completed-go-with-notes` (Commit 74f8e32).

## Policy, Approval and Execution Authorization (CO-WP-013)

- Drei neue Dokumente: [POLICY_DECISION_AND_EVALUATION_MODEL.md](../docs/security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](../docs/security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md), [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](../docs/security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md).
- Policy Evaluation, Approval und Execution Authorization sind getrennte Verantwortlichkeiten. Policy permit ≠ Approval ≠ Execution Authorization; kein Default-Permit (deny/indeterminate außer explizit autorisiert); `indeterminate`/`conflicted` fail-closed; Policy-Konflikte bleiben sichtbar.
- Approval explizit/zurechenbar/scope-bound/widerrufbar; Machine Principals ohne Self-Approval; Approver-Rollenname ≠ unbegrenzter Scope; Separation of Duties (kleine Deployments mit sichtbaren kompensierenden Kontrollen).
- Execution Authorization action-/target-/scope-/plan-/time-bound; materielle Plan-/Target-Änderung → Re-Evaluation; expired/revoked/consumed nicht wiederverwendbar; Pre-Execution Guards + fail-closed; Adapter/Agent erweitern Scope nicht; Break Glass kein Parallelmodell; Offline Authorization mit Target-Binding/Provenance/Integrität/expliziter Aktivierung.
- Executed ≠ successful ≠ verified; closed ≠ successful; audit evidence ≠ execution authority.
- Decision Index +16 (DEC-S-137…152), Risk Register +17 (RISK-190…206, gesamt 206).
- **Keine Policy-/Approval-/Execution-Engine, kein Autorisierungsartefakt/Token, kein Replay-Mechanismus, keine Queue/Workflow-Runtime ausgewählt; keine ADR; Identity-/Modul-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 438a5a0).

## CoreOps Integration Contract v0.1 (CO-WP-014)

- Drei neue Dokumente: [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../docs/architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md), [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](../docs/architecture/INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md), [INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md](../docs/security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md).
- Contract Version 0.1 ≠ CoreOps-Produkt-/Release-Version. 18 Integration Classes herstellerneutral. Integration Identity/Lifecycle (`discovered`…`archived`; `active-write-enabled` braucht Policy+Authorization; keine stille ID-Wiederverwendung).
- Sechs Capability-Dimensionen getrennt (advertised/detected/permitted/implemented/supported/validated); Adapter erteilt sich keine Permission per Declaration. 18 Operation Classes mit Privilege/Authorization-Einordnung; read ≠ export, observe ≠ configure, execute ≠ deploy, capability-to-execute ≠ authorization-to-execute.
- request acceptance ≠ authorization ≠ execution; execution-completed ≠ successful; successful ≠ verified; transport failure ≠ keine Nebenwirkung; unknown outcome sichtbar und blockiert automatischen Retry; partial failure sichtbar; rollback/recovery brauchen Verifikation.
- Read-only ohne stille Write Authority; Write/Execution brauchen explizite Policy+Authorization (CO-WP-013 Guards autoritativ, kein Parallelmodell); Adapter/Agent erweitern Target/Action/Scope nicht; Integrationsergebnisse erben keine Autorität (SoT/Provenance CO-WP-011 autoritativ); Offline mit Target-Binding/Provenance/Integrität/expliziter Aktivierung; Extensions überschreiben keine Core-Invariante.
- Decision Index +16 (DEC-S-153…168), Risk Register +13 (RISK-207…219, gesamt 219).
- **Keine Protokoll-/Schema-/Transport-/SDK-/Adapter-/Replay-/Queue-/Messaging-Technologie ausgewählt; keine ADR; Identity-/Authorization-/State-/Provenance-/Modul-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 611773b).

## Domain Pack Governance, Support Levels and Compatibility (CO-WP-015)

- Drei neue Dokumente: [DOMAIN_PACK_GOVERNANCE_MODEL.md](../docs/architecture/DOMAIN_PACK_GOVERNANCE_MODEL.md), [DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md](../docs/architecture/DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md), [DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md](../docs/security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md).
- Domain Pack = versionierte Governance-/Produktgrenze (≠ adapter/plugin/deployment unit/certification). 15 Pack-Klassen herstellerneutral; stabile, nicht wiederverwendbare Pack-IDs. Neun Statusdimensionen getrennt (Lifecycle/Maintenance/Support/Implementation/Validation/Evidence/Security-Review/Compatibility/Release).
- `active ≠ implemented ≠ maintained ≠ supported ≠ validated ≠ universally compatible`. Support Levels SUP-0/1/2/3/D; Support Level ≠ SLA; project-maintained ≠ validated; SUP-3 version-/target-/profil-/limitation-/evidence-bound. Compatibility Claims version-/target-/profil-/evidence-bound; expected ≠ validated; unknown ≠ compatible.
- Community/External ≠ trust/support/endorsement; public availability ≠ verified provenance; Vendor-Bezug ≠ Endorsement/Zertifizierung. Pack-Aktivierung gewährt keine Runtime-Autorität; Dependencies nicht still Core-Pflicht; Overlap ≠ doppelte Autorität. Offline mit Target-Binding/Provenance/Integrität/expliziter Aktivierung; kompromittierter Pack/Maintainer suspendierbar/widerrufbar; deprecated ≠ removed; retired ≠ historische Evidenz gelöscht; retired IDs nicht wiederverwendet.
- Decision Index +17 (DEC-S-169…185), Risk Register +10 (RISK-220…229, gesamt 229).
- **Keine Packaging-/Marketplace-/Plugin-/Update-/Dependency-Resolution-/Signaturtechnologie ausgewählt; keine ADR; Integration-/Identity-/Authorization-/State-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 8191e06).

## Data Ownership, Persistence, Schema Versioning and Migration (CO-WP-016)

- Drei neue Dokumente: [DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md](../docs/architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md](../docs/architecture/SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md](../docs/security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md).
- Neun Ownership-Dimensionen getrennt (Data Owner/Steward/Storage/Write/Migration/Retention/Disclosure/Recovery/Evidence). `data owner ≠ storage operator`; `storage responsibility ≠ write authority ≠ migration authority`; ein Adapter/Pack wird nicht durch Speicherung zum Owner. 16 Datenklassen mit autoritativem Modul; 10 Persistence-Klassen (`cached ≠ authoritative`, `append-only-governed ≠ technisch unveränderlich`, `unknown persistence ≠ safe for destructive migration`).
- Schema Identity stabil (≠ storage location/file name); 8 Versionsdimensionen getrennt (`schema version ≠ data version`); 11 Compatibility-Klassen (`read ≠ write ≠ round-trip`; `unknown ≠ compatible`; `conflicted ≠ safe for automatic migration`); 12 Change Classes. Migration Plan/Preflight/Lifecycle; `executed ≠ validated`; `rolled-back ≠ restored unless verified`; `closed ≠ successful`.
- Backup exists ≠ restorable; restore ≠ service recovery; rollback braucht Verifikation. Partial Migration sichtbar; Mixed-Version bounded. Destruktive Migration: gebundene Autorität + Approval, fail-closed. Migration erzeugt keine neue Autorität, reaktiviert keine widerrufene/konsumierte Authorization; Audit-/Evidence-Provenance erhalten; keine Roh-Secret-Speicherung entschieden. Offline mit Target-Binding/Provenance/Integrität/expliziter Aktivierung.
- Decision Index +18 (DEC-S-186…203), Risk Register +10 (RISK-230…239, gesamt 239).
- **Keine Storage-/DB-/ORM-/Schema-/Serialisierungs-/Migration-/Backup-/Replikations-/Cluster-/Transaction-Technologie ausgewählt; keine ADR; Integration-/Domain-Pack-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 69b3334).

## API Governance, Versioning, Errors and Idempotency (CO-WP-017)

- Drei neue Dokumente: [API_GOVERNANCE_AND_OPERATION_MODEL.md](../docs/architecture/API_GOVERNANCE_AND_OPERATION_MODEL.md), [API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md](../docs/architecture/API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md), [API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md](../docs/security/API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md).
- Konkretisiert den Integration Contract für programmatische Schnittstellen (kein Parallelmodell). 13 API-Klassen; API Identity stabil (`route/URL/transport ≠ identity`; `API version ≠ product version`); Lifecycle (proposed…archived). 17 Operation Classes mit Side-Effect-/Privilege-Klassifikation (`read ≠ export`, `configure ≠ execute`, `operation available ≠ authorised`).
- `API availability ≠ authorization`; `request acceptance ≠ execution`; `successful response ≠ verified outcome`; `error response ≠ proof of no side effect`; `duplicate response ≠ duplicate execution`. Zwölf Versionsdimensionen getrennt; Request/Response/Error/Behavioural-Compatibility getrennt (formal additiv kann breaking sein; `unknown ≠ compatible`); 14 Change Classes.
- Idempotency Context gebunden, verlängert/ersetzt keine Authorization; retry ≠ replay; `unknown outcome → kein Auto-Retry`, Reconciliation erforderlich; Duplicate/Replay erhält Historie und autorisiert nicht neu. Bulk erhält Per-Target-Autorität, Partial sichtbar; Pagination ≠ Snapshot/aktuelle Authorization/authoritative count. Consumer-safe Error-Disclosure (keine Secrets/Ressourcenexistenz); Workspace-Isolation explizit; API erzeugt keine parallele Autorisierungsautorität (CO-WP-013 autoritativ).
- Decision Index +16 (DEC-S-204…219), Risk Register +10 (RISK-240…249, gesamt 249).
- **Keine Transport-/API-Style-/Statuscode-/Schema-/Gateway-/Idempotency-/Replay-/Rate-Limit-Technologie ausgewählt; keine ADR; Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Letztes Work Package

`CO-WP-017 – API Governance, Versioning, Errors and Idempotency` (docs-only / API architecture and security foundation). Umsetzung abgeschlossen, Nova Review ausstehend. Vorheriges WP: `CO-WP-016 – completed-go-with-notes` (Commit 69b3334).

## Nächstes Work Package

`CO-WP-018 – Event, Audit Correlation and Evidence Model` (docs-only; planned-next; erst nach Nova Review von CO-WP-017 und Human-Maintainer-Commit).

## Human-Maintainer-Gates

Freigabe, Staging, Commit, Push, Merge, Tags, Releases, produktive Deployments sowie jede irreversible oder privilegierte Aktion sind ausschließlich dem Human Maintainer vorbehalten.

## Rückmeldung-an-Nova-Historie

- `CO-WP-001`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-001A`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO`.
- `CO-WP-002`: Fail-closed-Blocker gemeldet (`GO – Blocker bestätigt`); nach Source-Handoff regulär umgesetzt; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-003`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004A`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-004B1`: Rückmeldung geliefert; Nova Review `pending`.
- `CO-WP-004B2`: Rückmeldung geliefert; Nova Review: `GO WITH NOTES`.
- `CO-WP-004B3`: Rückmeldung geliefert; Nova Review `pending`.
- `CO-WP-004B4`: Rückmeldung geliefert; Nova Review `pending`.
- `CO-WP-004C`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-004D`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-004E`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-005`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-006`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-007`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-008`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-009`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-010`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-011`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.
- `CO-WP-012`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`. Milestone-Lessons-Review-Eignung gemeldet (CO-WP-005…012).

---

## Initialer Entscheidungsstand

- Produktvision **akzeptiert**.
- NDF v1.0.0 **normativ**.
- **Keine** technische ADR Accepted.
- **Keine** Technologie final ausgewählt.
- **Keine** Foundation-Freigabe erteilt.
- **Kein** Release vorbereitet.

## Notes

- **NDF-Level-Ambiguität:** Manifest nutzt `ndf_level: 1` (allgemeine NDF-v1.0.0-Vorlage); eine Starter-Vorlage nutzt `ndf_level: 2`. Offen und nicht aufgelöst; `ndf_level: 1` = initialer Bootstrap-Status, keine Reife/Zertifizierung/Compliance behauptet.
