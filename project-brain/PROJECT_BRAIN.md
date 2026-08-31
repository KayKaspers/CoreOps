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

- Foundation Capability Matrix: [docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md) — **94 Capabilities über 13 Domains** (autoritativ; in der CO-WP-004-Zusammenfassung seinerzeit als „74 Capabilities, 12 Domains" ausgewiesen — Zählkorrektur CO-WP-004E/CO-WP-029); alle `not-implemented`/`not-supported` (Proposed for acceptance).
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
- **Zählkorrektur:** Matrix enthält **94** Capabilities (grep-verifiziert); frühere Summe „74" korrigiert. Der Abgleich der übrigen „74"-Referenzen sowie die Korrektur der Domain-Summe (`12 → 13`) sind in **CO-WP-029** erfolgt.
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
- Status: `completed-go-with-notes` (Commit 7170c84).

## Event, Audit Correlation and Evidence Model (CO-WP-018)

- Drei neue Dokumente: [EVENT_AND_AUDIT_CORRELATION_MODEL.md](../docs/architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md), [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../docs/architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](../docs/security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md).
- Konsolidiert bestehende Audit-/Evidence-Begriffe ohne Parallelmodell. `event ≠ command ≠ notification ≠ audit event ≠ evidence`; `event identity ≠ correlation/request/operation/attempt`; `event producer ≠ authoritative source`. 20 Event-Klassen. Vier Zeitbegriffe getrennt (`occurrence/observation/recording/ingestion`); `timestamp/sequence ≠ authoritative global ordering`; `correlation ≠ causation`.
- Audit Events: `recorded ≠ validated`; `validated ≠ complete`; `correlated ≠ causally proven`; `purged ≠ historical event never existed`. Mehrere Attempts nicht zu einem Resultat verschmolzen. Duplicate/Replay/legitime Wiederholung getrennt (Historie erhalten). Audit Gaps sichtbar: `no event present ≠ action did not occur`; `missing evidence ≠ evidence of absence`; Completeness scope-bound/`unknown` möglich.
- Sechs Evidence-Dimensionen getrennt (capability/availability/freshness/integrity/validation/sufficiency); `available ≠ valid ≠ sufficient`; `stated integrity ≠ verified`; Sufficiency decision-/scope-/time-bound; `handling history ≠ legal admissibility` (keine rechtliche Beweiskraft). Audit Integrity: `append-only-governed ≠ technisch unveränderlich`; `audit administrator ≠ unrestricted disclosure/editor`; `read permission ≠ export permission`. Retention/Purge (keine Fristen); Offline mit Provenance/Integrität/Target-Binding/expliziter Governance; Failure sichtbar (`audit failure ≠ target action failed`); `closed ≠ complete/validated/sufficient/compliance`.
- Decision Index +17 (DEC-S-220…236), Risk Register +10 (RISK-250…259, gesamt 259).
- **Keine Event-Bus-/Logging-/SIEM-/Storage-/Schema-/Ordering-/Hash-/Signatur-/WORM-/Redaction-Technologie ausgewählt; keine ADR; API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit c7c3d90).

## Telemetry and Normalization Schema (CO-WP-019)

- Drei neue Dokumente: [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](../docs/architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md](../docs/architecture/TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md), [TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md](../docs/security/TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md).
- `telemetry ≠ event ≠ audit event ≠ evidence ≠ command`; `signal received ≠ source state currently true`. 22 Signal-Klassen; `signal/series/sample/resource identity` getrennt; `producer ≠ source`. Raw/Normalized/Derived/Aggregated getrennt (`raw ≠ validated`; `normalized ≠ lossless`; `derived ≠ independently observed`; `aggregated ≠ complete population`).
- Metric-Semantik (`counter ≠ rate`; `cumulative ≠ delta`; `zero ≠ missing ≠ unknown`); Logs (`log record ≠ governed audit record`; `absence of log ≠ no action`); Traces (`trace ≠ audit history`; `span completed ≠ operation successful`); Health (`self-reported healthy ≠ externally verified`; `telemetry absent ≠ service down`; `telemetry present ≠ healthy`).
- Canonical Fields (`same label ≠ same semantics`; `mapped field ≠ authoritative field`); Normalization Profiles (`profile exists ≠ mapping validated`); Units/Scale/Precision getrennt (`unknown/conflicting unit ≠ safe automatic conversion`); Quality/Confidence/Freshness/Validation getrennt; `recently ingested ≠ recently observed`; `stale ≠ invalid automatically`; `missing ≠ zero ≠ resource unavailable`. Sampling/Aggregation sichtbar (`sampled ≠ complete population`); Cardinality-/Label-Grenze (`available label ≠ safe dimension`).
- State-Authority-Boundary: `telemetry ≠ authoritative state`; `metric threshold ≠ execution authorization` (SoT/State autoritativ, kein Parallelmodell). Event/Evidence-Boundary: `anomaly ≠ incident`; `threshold crossed ≠ execution command`; Telemetry-to-Event/-Evidence mit expliziter Klassifikation/Provenance. Offline mit Provenance/Integrität/Target-Binding/Governance; Failure sichtbar (`telemetry failure ≠ target failure`; `missing telemetry ≠ evidence of absence`).
- Decision Index +18 (DEC-S-237…254), Risk Register +10 (RISK-260…269, gesamt 269).
- **Keine Telemetry-/Protokoll-/Schema-/Collector-/Storage-/Mapping-/Unit-/Aggregation-/Alerting-/Dashboard-Technologie ausgewählt; keine ADR; Event-/Audit-/Evidence-/API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 9bb12b2).

## Topology Graph, Evidence and Manual Authority (CO-WP-020)

- Drei neue Dokumente: [TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md](../docs/architecture/TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md), [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md](../docs/architecture/TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md), [TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md](../docs/security/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md).
- `topology graph ≠ authoritative physical reality`; `node ≠ managed resource automatically`; `edge ≠ verified connectivity`; `relationship assertion ≠ validated relationship`; `missing edge ≠ no relationship exists`. 24 Node-Klassen, 21 Relationship-Klassen; Node/Edge/Assertion Identity getrennt; 10 Assertion Origins (`manual ≠ fact`; `inferred ≠ independently observed`).
- Canonical Identity/Aliases (`alias match ≠ identity match`; `same hostname/address/external ID ≠ same node`); Identity Resolution (candidate…merged); Merge/Split erhalten historische Identität/Assertions/Evidence. Temporal Validity (`last observed ≠ currently valid indefinitely`; kein Last-Write-Wins); Views/Snapshots (`current view ≠ complete reality`; `snapshot ≠ alle Quellen zeitgleich/immutable`).
- Evidence-Dimensionen getrennt (Source Trust/Authority/Confidence/Validation/Evidence/Sufficiency/Completeness); `high confidence ≠ authoritative`; `many assertions ≠ independent sources`; `complete local graph ≠ complete physical topology`. Konflikte sichtbar (keine automatische Auflösung/Timestamp-Priorität); ungelöste sicherheitsrelevante Konflikte blockieren privilegierte Automation.
- Manual Authority human-attributable/scope-bound/reviewbar; Machine Principals imitieren keine Manual Authority; Override löscht keine konkurrierenden Observations/Evidence; `suppressed from view ≠ relationship does not exist`. State-/Execution-Boundary: `topology assertion ≠ authoritative state`; `graph path ≠ authorised network path`; `node selection ≠ authorised target scope`; Write/Execution bleibt an CO-WP-013 gebunden. Offline mit Provenance/Integrität/Target-Binding/Governance.
- Decision Index +19 (DEC-S-255…273), Risk Register +10 (RISK-270…279, gesamt 279).
- **Keine Graph-DB-/Discovery-/Query-/Identity-Resolution-/Conflict-Resolution-/Visualization-/Layout-Technologie ausgewählt; keine ADR; Telemetry-/Event-/Audit-/Evidence-/API-/Integration-/Data-/Identity-/Authorization-/State-/Provenance-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit 2c6d416).

## Milestone Lessons Review (CO-WP-013…020)

- Gebündelter docs-only Review der acht Foundation-WPs (24 Dokumente) abgeschlossen; Nova Review pending. Neues Dokument [MILESTONE_REVIEW_CO_WP_013_TO_020.md](MILESTONE_REVIEW_CO_WP_013_TO_020.md). Ergebnis: **GO WITH NOTES FOR CO-WP-021**.
- Foundation-Kette (Policy/Approval/Execution → Integration → Domain Pack → Data/Migration → API → Event/Evidence → Telemetry → Topology) kohärent; alle Cross-Foundation-Invarianten gehalten; keine Technologie ausgewählt; keine ADR; keine Implementierungs-/Compliance-Behauptung.
- Acht konsolidierte Lessons LL-023…030 (Lessons gesamt 30); drei NDF-Feedback-Kandidaten NDF-FC-COREOPS-011…013 (`candidate-pending-nova-review`, Kandidaten gesamt 13; 008…013 pending).
- Read-only-Befunde (nur Follow-up): Risk Register 279 (high 138/medium 117/low 24) → Konsolidierung ~CO-WP-029/030; Decision Index DEC-S-273 (Alt-Kombistatus DEC-S-01…37 vs. getrennte Dimensionen) → Harmonisierung ~CO-WP-029; Capability-Count „74→94" ~CO-WP-029; Dokumentationsökonomie (gemeinsames Invarianten-/Template-Referenzdokument) ~CO-WP-030.
- **Decision Index und Risk Register read-only unverändert; kein NDF-Transfer/keine Adoption; CO-WP-021 nicht begonnen; genau acht Allowed Files geändert.**
- Status: `completed-go-with-notes` (Commit d09d91b).

## Deployment Control Plane and Blueprint Schema (CO-WP-021)

- Drei neue Dokumente: [DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md](../docs/architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md](../docs/architecture/DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md), [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](../docs/security/DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md).
- Deployment Control Plane = Governance-/Koordinationsgrenze (≠ Runtime/Engine/Orchestrator). `blueprint ≠ plan`; `deployment intent ≠ approved plan`; `approved plan ≠ execution authorization`; `deployment executed ≠ successful`; `successful ≠ desired state verified`; `closed ≠ successful`. Bindet jede Write-Aktion an CO-WP-013 (keine parallele Authorization Authority).
- Blueprint Identity stabil (≠ file name/repo path/product version); getrennte Versionsdimensionen; Inputs/Parameters/Defaults/Secret References (`default ≠ safe value`; `secret reference ≠ raw secret`); Environment Overlays (`overlay ≠ permission to weaken security`); Effective Blueprint mit Provenance; Artifact-Bindung (`available ≠ trusted ≠ compatible`; `newer ≠ authorised replacement`).
- Dynamische Topology-Selektion → **materialisierter, begrenzter Target-Set-Snapshot** (`topology query ≠ authorised target set`); conflicted/unresolved Identität wird nicht still privilegiertes Ziel; Pre-Execution-Revalidation (`approved target snapshot ≠ permanently valid`). Waves (`wave success ≠ remaining targets authorised`); Partial/Unknown Per-Target sichtbar (`unknown → kein Auto-Retry`); Cancellation (`cancel accepted ≠ no side effects`); Verification getrennt (`artifact present/health green/telemetry available ≠ verified/sufficient`); Rollback braucht Verifikation; Forward Recovery.
- State-Authority-Boundary (`deployment executed ≠ observed matches ≠ effective verified`; SoT/Drift autoritativ). Machine Principals imitieren keine Human Deployment Approval/Manual Authority. Offline mit Provenance/Integrität/Target-Binding/gebundener Authorization/expliziter Aktivierung.
- Decision Index +15 (DEC-S-274…288), Risk Register +5 (RISK-280…284, gesamt 284; Wachstumsgrenze §4 eingehalten).
- **Keine Deployment-Engine-/Orchestrator-/Pipeline-/Blueprint-Format-/Schema-/Agent-/Registry-/Rollback-Technologie ausgewählt; keine ADR; Policy-/Integration-/Domain-Pack-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**
- Status: `completed-go-with-notes` (Commit a152091).

## Artifact Trust, SBOM, Provenance and Revocation (CO-WP-022)

- Drei neue Dokumente: [ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md](../docs/architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md), [ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md](../docs/architecture/ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md), [ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md](../docs/security/ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md).
- `artifact available ≠ trusted`; `identity ≠ alias ≠ file name ≠ repository path`; `version ≠ revision`; `same version label ≠ same revision`; `newer ≠ safer ≠ compatible ≠ authorised replacement`. 19 Artifact-Klassen; 13 Rollen getrennt (`repository operator ≠ artifact owner`; `distributor ≠ trust authority`; `maintainer ≠ revocation authority`). Mutable Alias unterstützt Discovery, aber privilegiertes Deployment bindet aufgelöste Identität/Version/Revision.
- Provenance/Integrity/Validation/Trust/Support/Compatibility getrennt (`provenance available ≠ validated`; `integrity checked ≠ safe`; `validated ≠ universally trusted`; Trust-States untrusted…revoked; `trusted-for-bounded-scope` mit Scope/Versions/Evidence/Owner/Review). SBOM/Components (`SBOM available ≠ complete ≠ accurate`; `component absent from SBOM ≠ component absent`; `multiple SBOMs ≠ independent sources`); Component Identity (`same name/version ≠ same component/revision`).
- Vulnerability getrennt (`vulnerability reference ≠ artifact affected`; `component affected ≠ deployment exploitable`; `no scanner finding ≠ no vulnerability`; `fixed version ≠ compatible ≠ authorised`; `severity ≠ deployment-context risk`). Trust Decision use-/target-/scope-/version-/time-bound (`trusted-for-one-use ≠ trusted globally`; ≠ Execution Authorization). Quarantine (`quarantined ≠ malicious proven`; `quarantine released ≠ deployment authorised`).
- Withdrawal/Revocation/Repository-Removal/Support-Withdrawal getrennt; `revocation issued ≠ delivered to every offline environment`; **`revocation received → new deployment blocked unless explicit governed exception`**; `existing deployment ≠ permission for new deployment`; `artifact revoked ≠ automatic destructive removal`. Reinstatement (`reinstated ≠ historical revocation erased ≠ globally trusted`). Deployment Binding an konkrete Resolution + Pre-Execution-Recheck. Offline mit Provenance/Integrität/Target-Binding/gebundenem Trust/expliziter Import-Governance; Delayed Revocation sichtbar.
- Decision Index +15 (DEC-S-289…303), Risk Register +5 (RISK-285…289, gesamt 289; Wachstumsgrenze §4 eingehalten).
- **Keine Registry-/Package-/SBOM-/Hash-/Signing-/Trust-Anchor-/Transparency-/Scanner-/Build-/Dependency-Resolution-Technologie ausgewählt; keine ADR; Deployment-/Integration-/Domain-Pack-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Restricted, Isolated, Air-Gapped Operation and CorePack (CO-WP-023)

- Drei neue Dokumente: [RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](../docs/architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md), [COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md](../docs/architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md), [OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md](../docs/security/OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md).
- Connectivity Classes getrennt (`connected/restricted-connected/intermittently-connected/isolated/air-gapped/recovery-only`); `offline ≠ air-gapped`; `isolated ≠ trusted`; `network unavailable ≠ security controls optional`; keine Klassifiziertnetz-Eignung. Authority Boundary: `central unavailable ≠ local authority expands`; `local activation authority ≠ deployment execution authorization`; Delegation scope-/purpose-/profile-/version-/zeitgebunden.
- `CorePack ≠ Domain Pack/artifact/deployment blueprint/backup/evidence package`; Identity/Version/Revision/Assembly-/Transfer-/Import-/Activation-Instance getrennt; `transferred ≠ imported ≠ trusted ≠ approved-for-activation ≠ activated ≠ deployment authorised`. Content Resolution bindet konkrete Revision (`mutable alias ≠ final binding`); Contents erben Trust/Compatibility nicht vom Container.
- Target Binding durch Transfer/Import/Activation erhalten; `quarantine release ≠ activation authorization`; Partial/Unknown Activation sichtbar (`no automatic retry → reconciliation`); Delta erfordert bestätigte exakte Baseline; Delayed Revocation sichtbar (`no local revocation entry ≠ not revoked centrally`); Offline Authorization action-/target-/scope-/version-/purpose-/time-bound und nicht wiederverwendbar; Rollback erfordert aktuelle Trust-/Revocation-/Compatibility-Bewertung; Evidence Return (`returned ≠ complete ≠ sufficient`); Export-Autorität getrennt; keine Rohsecrets.
- Decision Index +14 (DEC-S-304…317), Risk Register +5 (RISK-290…294, gesamt 294; Wachstumsgrenze §4 eingehalten).
- **Keine Package-/Manifest-/Installer-/Update-/Transfer-/Removable-Media-/Hash-/Signing-/Trust-Anchor-/PKI-/Encryption-/Synchronisations-/Reconciliation-Technologie und keine Offline-Runtime ausgewählt; keine ADR; Artifact-/Deployment-/Domain-Pack-/Integration-/Data-/API-/Event-/Evidence-/Telemetry-/Topology-/Identity-/Authorization-/State-/Capability-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Secrets, Configuration Vault and Key Custody (CO-WP-024)

- Drei neue Dokumente: [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../docs/security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md), [KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md](../docs/security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md), [CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md](../docs/architecture/CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md). Baut auf MOD-SEC-001/MOD-STA-001, CO-WP-010 (Credential Lifecycle) und CO-WP-009 (Break-Glass) auf; kein Parallelmodell.
- `secret ≠ ordinary configuration`; `secret reference ≠ secret value`; `credential ≠ identity`; `credential possession ≠ authorization`; `logical secret identity ≠ version ≠ value revision ≠ value instance ≠ lease`; `retrieval permitted ≠ use permitted ≠ export permitted`; `key identity ≠ key material ≠ custody`; `key custody ≠ unrestricted key use`; `vault availability ≠ secret trust`.
- Vault = logische Governance-Grenze (kein Produkt). 15 Secret-/Config-Klassen; getrennte Autoritäten (owner ≠ custodian ≠ user; retrieval ≠ use ≠ export; rotation ≠ reinstatement; backup operator ≠ recovery authority; machine principal ≠ human approval). Secret Lifecycle (proposed…destroyed/outcome-unknown); Retrieval/Distribution/Injection/Use getrennt (`injected ≠ consumed`).
- Rotation vollständig erst nach Consumer-State-Assessment (`new version ≠ rotation complete`; `unknown consumer ≠ safe`); Revocation mit Identity Binding + Freshness (`no local entry ≠ not revoked centrally`); Bootstrap/Root/Recovery/Break-Glass bounded/one-time, mandatory Rotation/Invalidation, keine dauerhafte Autorität; Recovery ≠ Reinstatement ≠ globaler Trust; Rollback reaktiviert kein revoked/expired/compromised Material. Configuration Source of Truth ≠ Runtime State; Secret References fail-closed; Drift ≠ auto-remediated. CorePack Trust ≠ Secret Trust; Deployment Authorization ≠ Secret-Use Authorization; Offline erweitert keine Autorität, Offline-Authorization nicht wiederverwendbar; Audit/Evidence ohne Secret Values.
- Decision Index +14 (DEC-S-318…331; nach Nova-GO-WITH-NOTES-Deduplizierung, ursprünglich 16), Risk Register +5 (RISK-295…299, gesamt 299; Wachstumsgrenze §36 eingehalten).
- **Keine Vault-/KMS-/HSM-/TPM-/PKI-/Key-Format-/Algorithmus-/Secret-Store-/Injection-/Configuration-/Rotation-/Backup-/Transfer-/Sync-Technologie ausgewählt; Rohsecret-/Rohschlüsselspeicherung nicht entschieden; keine ADR; bestehende Identity-/Authorization-/Trust-/Deployment-/Artifact-/Evidence-/Offline-/CorePack-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Data Classification, Retention and Redaction (CO-WP-025)

- Drei neue Dokumente: [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../docs/governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md), [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](../docs/governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md), [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../docs/security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md). Konkretisiert DEC-P-08; baut auf CO-WP-016/018/019/024 auf; kein Parallelmodell.
- `data classification ≠ deployment profile ≠ connectivity class ≠ national-security classification`; `classification label ≠ proven handling compliance`; `unknown-classification → fail-closed`; 7 Klassen (public/internal/sensitive/restricted/secret-bearing/evidence-protected/unknown); Classification Identity/Scope/Version/Freshness getrennt (materiale Änderung → Reassessment; reclassification erhält Historie). Getrennte Daten-Autoritäten (owner/steward/custodian/collection/use/disclosure/export/retention/hold/deletion/redaction).
- `collection permitted ≠ every later use`; `useful ≠ necessary`; Retention beginnt an definiertem Start Event; `retention expired ≠ deletion completed`; Preservation Hold und Retention getrennte Autoritäten (`hold ≠ unrestricted access`; neutraler Begriff, keine Rechtswirkung); Deletion request/authorization/execution/verification getrennt (`primary deletion ≠ all copies`; `logical ≠ destruction`; `unknown ≠ deleted`); Kopien/Backups/Caches/Derived im Scope.
- Redaction erzeugt gebundene Derived View (`redacted view ≠ source`; `redaction applied ≠ disclosure safe`); `masked/pseudonymized ≠ anonymous`; `hash ≠ anonymization`; Disclosure/Export/Publication getrennt von read/local (`data available ≠ disclosure authorized`); Secret-bearing Data an CO-WP-024 gebunden (Redaction ≠ Secret Removal); Evidence Retention ≠ Source Retention; Offline erweitert keine Retention-/Deletion-/Disclosure-Autorität; Unknown Outcomes fail-closed; Profile = Kontrollstärke (keine Zertifizierung).
- Decision Index +13 (DEC-S-332…344), Risk Register +5 (RISK-300…304, gesamt 304; Wachstumsgrenze §33 eingehalten).
- **Keine Retention-/Deletion-/Archiv-/DLP-/Discovery-/Redaction-/Masking-/Anonymisierungs-/Storage-/Backup-/Encryption-/Sync-Technologie ausgewählt; keine ADR; keine reine Technology-Deferral-Decision; keine Compliance-/Rechts-/Zertifizierungsbehauptung; bestehende Data-/Audit-/Evidence-/Telemetry-/Topology-/Secret-/Offline-/CorePack-/Threat-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Self-Protection, Degraded Modes and Recovery Mode (CO-WP-026)

- Drei neue Dokumente: [SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md](../docs/security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md), [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../docs/architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md), [RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md](../docs/governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md). Adressiert RISK-11 (Self-Dependency); erweitert die Operational States aus CO-WP-023 §14; kein Parallelmodell.
- `CoreOps self-protection ≠ protection of every managed asset`; `control-plane health ≠ managed-system health`; `process running ≠ platform governable`; **`governability ≠ bloße technische Erreichbarkeit`**; `monitoring unavailable ≠ system healthy`; `audit unavailable ≠ no operation occurred`; `trigger observed ≠ compromise proven`; `trigger absent ≠ platform safe`. 15 Schutzgüter, 23 Trigger-Kategorien, Protection Assessment (20 Felder), 12 Fault Domains (`shared infrastructure ≠ shared authority`; `fault in one workspace ≠ automatic global shutdown`), 20 Schutzmaßnahmen mit sechsstufiger Priorität (`protective action ≠ punishment/root-cause remediation`).
- Zehn Operational Modes (normal/guarded/restricted/read-only/degraded/containment/recovery-only/recovery/emergency-stop/unknown) mit Capability Restriction Matrix über 23 Capability-Gruppen (permitted/with-approval/bounded-recovery/suspended/prohibited/unknown-fail-closed). `read-only mode ≠ no side effects`; `read-only UI ≠ read-only platform`; `restricted ≠ read-only`; `guarded ≠ degraded`; `degraded ≠ authority expansion/indefinite exception` (DEC-S-316); `containment ≠ recovery`; `emergency-stop ≠ permanent shutdown`; `unknown ≠ normal/safe` → fail-closed.
- Recovery: `recovery mode ≠ ordinary mode ≠ backup restore ≠ rollback`; Recovery Authority eigenständig/begrenzt (`recovery operator ≠ recovery approver`; `local administrator ≠ recovery authority`; `break-glass ≠ unrestricted/permanent recovery authority`); 15 Stufen (`technical restoration ≠ governance restoration`; `service reachable ≠ recovery verified`); 11 Input-Klassen mit aktueller Trust-/Revocation-/Compatibility-Bewertung (`previously trusted ≠ currently trusted`; `backup exists ≠ valid for this target`); `secret restored ≠ secret trusted`; `key recovered ≠ key use authorized`; Partial/Unknown blockieren unsichere Wiederholung; Offline erweitert keine Recovery-Autorität; Recovery Exit braucht Reassessment/Verification/Reconciliation und kann in guarded/restricted/degraded austreten.
- Decision Index +13 (DEC-S-345…357), Risk Register +5 (RISK-305…309, gesamt 309; Wachstumsgrenze §39 eingehalten).
- **Keine Health-/Watchdog-/HA-/Failover-/Cluster-/Quorum-/Recovery-/Self-Healing-/Monitoring-/Backup-/Isolation-/Integrity-Scanning-/Orchestration-/Sync-/Reconciliation-Technologie ausgewählt; keine ADR; keine Technology-Deferral-Decision; keine Security-/Resilience-/Recovery-Readiness-/Compliance-Behauptung; bestehende Foundation-Dateien, Lessons-Learned-Register und NDF-Feedback-Kandidaten unverändert; keine NDF-Rückführung.**

## Foundation Milestone Review (CO-WP-021…026)

- Gebündelter, unnummerierter **review-only** Milestone über sechs Foundation-WPs und 18 Dokumente. Neues Dokument: [MILESTONE_REVIEW_CO_WP_021_TO_026.md](MILESTONE_REVIEW_CO_WP_021_TO_026.md). Status `completed-go-with-notes`. Ergebnis **GO WITH NOTES**; Empfehlung **GO WITH NOTES FOR CO-WP-027** (keine Freigabe erteilt).
- **Foundation-Kohärenz:** Deployment (021) → Artifact Trust (022) → Restricted/Offline/CorePack (023) → Secrets/Key Custody (024) → Data Classification/Retention/Disclosure (025) → Self-Protection/Degraded/Recovery (026) bilden eine geschlossene Kette; jede Stufe konsumiert die vorige als autoritative Grenze. **0 Konflikte, 0 Parallelmodelle** für Autorität, State, Health, Trust oder Evidence.
- **Invarianten gehalten:** `plan ≠ approval ≠ execution authorization`; `executed ≠ successful ≠ verified`; `closed ≠ successful`; `partial ≠ complete`; `unknown ≠ failed ≠ safe`; `artifact available ≠ trusted`; `integrity verified ≠ safe`; `CorePack available ≠ trusted ≠ activation authorized`; `secret reference ≠ secret value`; `retrieval ≠ use ≠ export`; `retention expiry ≠ deletion`; `backup available ≠ restore authorized`; `service reachable ≠ recovery verified`; `degraded mode ≠ authority expansion`; `recovery authority ≠ ordinary policy authority`; `technical restoration ≠ governance restoration`.
- **Stärkster Dedup-Befund:** CO-WP-023 definiert die Offline-Autoritätsgrenze **einmal autoritativ**; 024…026 referenzieren sie statt eigener Offline-Modelle (`central authority unavailable ≠ local authority expands automatically`).
- **Register read-only:** Decision Index `DEC-S-357` (+84, lückenlos, je WP 15/15/14/14/13/13); Risk Register 309 (+30, exakt 5/WP; high 167/medium 118/low 24; treatment-planned 292/open 17). **Neue Registerbefunde:** Severity verliert Trennschärfe (29 von 30 neuen Risiken `high`), Lifecycle bleibt uniform (`open` unverändert 17) → Rekalibrierung `CO-WP-029/030`.
- **Reale Betriebsevidenz** (extern, Stärke `limited`, generisch): bestätigt die Foundation. Coverage-Matrix über 15 Konzepte — 5 `already covered`, 4 `partially covered`, 4 `genuine extension` (alle eng, alle später), 2 `duplicate/reject`. „Safe Change Transaction" = **Kompositionsmuster**, kein neuer Lebenszyklus. Geschichtete, scope-gebundene Health bereits abgedeckt (Telemetry §17: sechs Health-Klassen + Source/Freshness/Observation Scope); offen: consumer-gebundene Vertragsprüfung. Echte Erweiterungen: Protected-Drift-Änderungskontext, Recovery-Set-Komposit, Bindung an den **erfassten** Vorzustand, Expected-vs-Actual-Vergleich.
- **Lessons LL-031…038** (gesamt 38); **NDF-Kandidaten NDF-FC-COREOPS-014…015** (`candidate-pending-nova-review`, gesamt 15). Vier weitere Muster bewertet und ausdrücklich **nicht** promoted.
- **Keine** NDF-Rückführung (NDF v1.0.0 bleibt normativ; post-v1.0-`main` nur informativ/methodisch genutzt); **keine** CDS-Adoption/-Import/-Pilot (eine nicht-aktivierende Gate-Notiz für `CO-WP-027`); Decision Index und Risk Register unverändert; kein Foundation-Dokument geändert; `CO-WP-027` nicht begonnen; keine Git-Schreibaktion.

## UX Information Architecture and Dashboard System (CO-WP-027)

- Drei neue Foundation-Dokumente: [UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md](../docs/architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md), [DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md](../docs/architecture/DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md), [UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md](../docs/security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md).
- **Experience-Grenze unverändert:** MOD-EXP-001 stellt dar, validiert, nimmt Intents entgegen und routet; `presentation ≠ source of truth`; `UI intent ≠ approval`; keine neue Experience-Autorität, kein paralleles Datenmodell (DEC-S-68 referenziert, nicht dupliziert).
- **IA:** Presentation Context über die **acht bestehenden RBAC-Scopes**; **13 Top-Level-Informationsbereiche**, jeder auf eine bestehende Capability-Domäne und ein Modul rückführbar; Netzwerk/Druck/Virtualisierung als **Domain Lenses** ohne eigene Autorität. `navigation ≠ domain ownership ≠ target binding`; `selection ≠ authorization`; `filter ≠ execution scope`.
- **Dashboard:** `dashboard representation ≠ authoritative domain state`; kein zweites System of Record; eine Zusammenfassung darf priorisieren, aber **keine Dimension löschen**. **Elf getrennte Statusdimensionen** (condition · severity · confidence · freshness · evidence availability · data coverage · capability availability · permission state · operational mode · conflict state · verification state); **kein** globaler Health-Boolean, **kein** normativer Aggregat-Score; Ampel/Farbe/Icon/Position sind Darstellungsform, keine Statussemantik. CoreScore bleibt objektbezogen mit Pflichtbegleitangaben.
- **Ehrlichkeitsgrenzen:** `unknown ≠ healthy`; `stale ≠ current`; `partial ≠ complete`; `degraded ≠ failed`; `hidden ≠ nonexistent`; `permission denied ≠ capability unavailable ≠ unknown ≠ restricted by mode`; `provider/process healthy ≠ consumer capability usable`.
- **Simple/Expert:** im Produktmodell (Concept V3, UX) **bereits benannt**, dort **ohne** Semantik; UX-Semantik definiert, **ohne** Autoritätswirkung — `Simple mode ≠ different authorization model`, `Expert mode ≠ authority expansion`, fehlende Präferenz → sicherer expliziter Fallback, Expert ≠ Break-Glass.
- **Dangerous Action:** zehnstufiger Pfad nicht kollabierbar; `UI confirmation ≠ approval record`; `preview ≠ execution`; `policy permit ≠ approval ≠ execution authorization`; `executed ≠ successful ≠ verified`; acht Pflichtangaben inkl. `unknown`.
- **Accessibility:** zwölf Designanforderungen (Tastatur, Fokus, Nicht-nur-Farbe, Textäquivalente, klare Sprache, dichte Daten, Fehleridentifikation, DE/EN-Textexpansion, reduzierte Bewegung, nicht-visuelle Unterscheidbarkeit von unknown/stale/partial/degraded/denied). **Keine** WCAG-/Validierungs-/Screenreader-/Keyboard-Test-/Zertifizierungsaussage; **keine** Accessibility-Evidenz erzeugt.
- **Disclosure:** `data readable ≠ data discloseable`; `redacted view ≠ source deleted`; `evidence reference ≠ raw evidence discloseable`; `secret reference ≠ secret value`; sieben Export-/Copy-/Rendering-Flächen als Disclosure-Flächen benannt; bestehende Familien (CO-WP-024/025, RISK-70/302/304) referenziert statt dupliziert.
- **CDS:** `Candidate` / **read-only** Vergleichseingabe. Achsenunabhängigkeit und „Farbe allein trägt keine Bedeutung" sind `compatible overlap`; Scope-Bindung, vier Ursachenklassen, Operational Modes und die elf Dimensionen sind `CoreOps-local`; die Fünf-Achsen-Abbildung ist `future mapping candidate`; eine 1:1-Reduktion wäre Bedeutungsverlust → `not adopted`. **Keine** Adoption, **kein** Import, **keine** Tokens, **kein** aktivierter Pilot, **keine** Adoption-Evidenz.
- Decision Index +8 (DEC-S-358…365), Risk Register +4 (RISK-310…313, gesamt 313; Severity bewusst differenziert; RISK-313 erster Accessibility-Eintrag). Keine ADR; keine Technologieauswahl; keine Mockups; kein Implementierungsanspruch. Lessons-Learned-Register und NDF-Feedback-Kandidaten **unverändert**.
- **Nova-Notes-Closure (1–3):** (1) die elf Dimensionen sind eine **Darstellungsordnung**, **kein** neues autoritatives Statusschema (`UX presentation dimensions ≠ new authoritative status schema`; `presentation ≠ source-of-truth ownership`); kein Subjekt muss alle elf materialisieren; `not applicable ≠ unknown` und `dimension not applicable to a subject ≠ applicable dimension whose value is unknown`; jede Dimension bleibt durch ihr bestehendes autoritatives Modell governt (§11.1). (2) `presentation context ≠ authorization scope`; `referencing a scope identity ≠ performing authorization`; `context narrowing ≠ permission mutation`; `navigation ≠ authorization`; `UI visibility ≠ permission result` — die Experience-Ebene bindet eine Sicht an eine bestehende Scope-Identität, MOD-IAM-001/MOD-POL-001 bleiben allein autoritativ; kein zweites Scope-Modell, RBAC-Foundation unverändert (§7.1). (3) `safe presentation fallback: Simple / reduced-complexity presentation` — Darstellungsentscheidung, **kein** anderer Autorisierungszustand; `Simple mode ≠ fewer permissions`, `Expert mode ≠ more permissions`, `Expert mode ≠ Break Glass`, `mode preference ≠ authorization`; Presentation Modes und Operational Modes ausdrücklich abgegrenzt. Keine neuen Decisions/Risks; keine Renumerierung; DEC-S-358…365 und RISK-310…313 unverändert.

## Test Strategy, Fixtures and Integration Lab (CO-WP-028)

- Drei neue Foundation-Dokumente: [FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md](../docs/testing/FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md), [SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md](../docs/testing/SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md), [INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md](../docs/testing/INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md).
- **Keine zweite Evidence-Autorität:** Test-Evidenz ist eine Evidence-Klasse **innerhalb** des bestehenden Evidence-/Audit-/Provenance-Modells (MOD-EVD-001, CO-WP-018); `test evidence ≠ operational authority`; `test evidence ≠ approval`; `validation ≠ support`; `support ≠ implementation`.
- **Claim Boundary Set (unabhängige Nicht-Implikationen, `claim boundary set ≠ universal maturity ladder`):** nur `test planned → test implemented → test executed → test result` ist eine echte lokale Progression; alle übrigen Aussagen stehen nebeneinander und sind keine Sprossen — `test planned ≠ test implemented ≠ test executed ≠ test passed ≠ requirement universally satisfied`; `fixture passed ≠ production validated`; `synthetic fixture ≠ real environment`; `simulator success ≠ provider compatibility`; `integration-lab success ≠ production readiness`; `one successful run ≠ regression confidence`; `no failing test ≠ absence of defect`; `coverage reported ≠ coverage complete`; `not tested ≠ passed`; `blocked ≠ failed`; `inconclusive ≠ passed`; `not applicable ≠ unknown`.
- **Testebenen:** zehn aus CoreOps abgeleitete Ebenen `TL-1`…`TL-10` — **keine** Testpyramide, keine Mengenverteilung, `test level ≠ authority level`. Je Ebene: Purpose, Subject under Test, autoritativer Input, Fixture-/Umgebungstyp, erwartete Evidenz, Failure-Bedeutung, ausdrücklich Nicht-Belegbares, Automatisierungs- und Offline-Eignung.
- **Traceability:** 18 Mindestfelder je Testfall mit Bezug auf Contract/Requirement/Invariant und ggf. Decision/Risk/Threat; Identity und Revision getrennt; `test case references authority ≠ test case becomes authority`.
- **Ergebnissemantik:** sechs Zustände (`not run`, `passed`, `failed`, `blocked`, `inconclusive`, `not applicable`) — **ausschließlich** für Testausführung/-evidenz, **kein** globales CoreOps-Statusmodell. `partially passed`, `stale`/`expired` und `error` wurden geprüft und **bewusst verworfen** (Teilergebnisse auf Observation-Ebene; Alterung über die bestehende Evidence-Freshness; Harness-Fehler als `blocked`/`inconclusive`).
- **Negativ-/Fail-Closed-Primat:** Negativszenarien sind erstklassig; `passed` nur bei **beobachtetem geschütztem Verhalten** — `an error message appeared ≠ protected behavior observed`; fehlende Beobachtung des Ausbleibens einer Wirkung → `inconclusive`. 17 verpflichtende Negativfamilien.
- **Szenarienfamilien:** State/Authority/Evidence, Migration, Telemetry, Deployment, Artifact/SBOM/Vulnerability, Offline/CorePack, Secrets, Klassifikation/Retention/Disclosure, Self-Protection/Degraded/Recovery, UX/Dashboard — jeweils mit den zu erhaltenden Unterscheidungen der jeweiligen Foundation-Modelle, referenziert statt dupliziert.
- **Integration Contract:** die sechs Capability-Dimensionen bleiben getrennt; `supported` ist durch **keinen** Test belegbar (Governance-Entscheidung mit eigener Evidenzanforderung). Consumer/Provider-Mismatch als Pflichtszenario: `provider process healthy ≠ protocol reachable ≠ consumer can use dependency` — **ohne** die Dependency-Contract-Health-Erweiterung zu implementieren oder zu registrieren.
- **CO-WP-027 als Testsubjekt:** neun UX-Designannahmen (13 Informationsbereiche, Presentation Context, elf Dimensionen, `not applicable`/`unknown`, vier Unavailability-Ursachen, Simple/Expert, Overview→List→Detail→Evidence, Dangerous-Action-Pfad, Dashboard-Priorisierung) sind **zu prüfende Annahmen**, keine belegten Wahrheiten; Methoden benannt, **keine** Usability-Validierung behauptet.
- **Accessibility:** neun Evidenz**anforderungen** (Tastatur, Fokus/Reihenfolge, Nicht-nur-Farbe, Textäquivalente, Zustandsunterscheidung, Gefahrenkommunikation, DE/EN-Expansion, reduzierte Bewegung, dichte Daten) — **keine** WCAG-Konformität, **keine** Validierung, **kein** Screenreader-Nachweis, **kein** Werkzeug ausgewählt oder ausgeführt.
- **Fixtures:** synthetisch als Default; Identity, Revision, Provenance und Expected Outcome explizit; 19 Fixture-Klassen; **kein** reales Secret, **kein** wiederverwendbares Credential, **kein** realer privater Hostname/Netzbereich, **keine** unredigierten personenbezogenen Daten. `synthetic ≠ safe by definition`; `test data ≠ unclassified data`; `masked secret ≠ non-secret automatically`; `fixture correctness ≠ system correctness`; `fixture representativeness ≠ production equivalence`. Produktionsdaten als Fixture-Quelle standardmäßig ausgeschlossen; sieben kumulative Ausnahmebedingungen beschrieben, **keine** autorisiert.
- **Integration Lab:** logisch, `non-production`, isoliert, verwerfbar, selbst-hostbar, offline-fähig, public-neutral, credential-safe, scope-gebunden, evidenzerzeugend. Sechs konzeptionelle Rollen (**keine** Laufzeitmodule, **keine** neuen `MOD-*`); sieben Umgebungsprofile auf den bestehenden Connectivity-Klassen und Operational Modes; sechs Test-Double-Klassen mit **verpflichtender** Fidelity-Deklaration. `lab role ≠ production authority`; `lab target ≠ production target`; `lab credential ≠ production credential`; `lab approval ≠ production deployment authorization`; `lab profile ≠ CoreOps operational authority`; `simulated degraded state ≠ evidence of real degraded-mode resilience`. Achtstufiges Provisioning Gate — **keine** Bedingung erfüllt, Lab `not provisioned`.
- **Coverage:** 14 Dimensionen; **keine** einzelne Prozentzahl als alleiniges Maß; unbekannte Coverage bleibt explizit; `coverage metric ≠ quality`; `100% of declared subset ≠ 100% of Foundation`. Aktueller Stand in allen Dimensionen `not measured`.
- **Risikoabdeckung:** alle 21 auf `CO-WP-028` gerichteten Risiken in einer Risk-to-Test-Matrix mit Test Layer, Fixture-/Szenarioklasse, erwartetem geschütztem Verhalten, geforderter Evidenz, Grenzen und Future Execution Gate — keiner zurückgestellt, keiner ausgelassen, **keiner geschlossen**.
- Decision Index +8 (DEC-S-366…373), Risk Register +3 (RISK-314…316, gesamt 316; Severity differenziert 2× high / 1× medium; erste Test-/Fixture-/Coverage-/Laboreinträge des Registers; Level-Summen aus den ID-Listen neu berechnet). Keine ADR; keine Technologieauswahl; keine ausführbaren Tests, Skripte, CI-Jobs oder Container; kein bestehendes Foundation-Dokument geändert. Lessons-Learned-Register und NDF-Feedback-Kandidaten **unverändert**.
- **CDS:** Pilot **nicht aktiviert**, keine CDS-Testevidenz importiert oder referenziert — `CoreOps test strategy ≠ CDS consumer evidence automatically`; `CDS Candidate evidence ≠ CoreOps validation`.
- **Nova-Notes-Closure (1–3):** (1) **Claim Boundary Set statt Reifeleiter** — die Claim-Aussagen sind eine Menge voneinander unabhängiger Nicht-Implikationen (`claim boundary set ≠ universal maturity ladder`); nur `test planned → implemented → executed → result` ist eine echte lokale Progression; `support`, `production readiness`, `security validation`, `accessibility validation`, `fixture representativeness` und `operational validation` bleiben eigenständige Dimensionen mit eigener Autorität und eigener Evidenzanforderung und werden **nicht** zu Pflichtsprossen (`eigenständige Dimension ≠ Endstufe einer Testprogression`; `Testfall ohne Bezug zu diesen Dimensionen ≠ unvollständiger Testfall`). (2) **Revision ≠ Staleness** — Evidenz zu Revision A bleibt **historische Evidenz für Revision A**; eine neue Test-Case-, Fixture- oder Source-Revision macht sie **nicht** automatisch `stale` (`neue Revision B ≠ Evidenz zu Revision A wird stale`; `Evidenz zu Revision A ≠ Validierungsevidenz für Revision B`). Eine materielle Revisionsänderung etabliert lediglich **keine Anwendbarkeit** auf die neue Revision und erfordert **Reassessment bzw. Revalidierung**; `stale` gilt ausschließlich dort, wo das bestehende Evidence-Freshness-Modell Staleness feststellt. Historische Evidenz wird nicht umgeschrieben; **kein** eigener Test-Evidence-Lebenszyklus. (3) **Queue-Kopfzeile neutralisiert** — die veraltete Router-Aussage zu `CO-WP-013` ist entfernt; die Kopfzeile benennt keinen nächsten Schritt und erteilt keine Freigabe, der aktuelle Stand ergibt sich aus Statusspalten und Current-State-Spiegeln. **Keine** automatische Autorisierung von `CO-WP-029`. Keine neuen Decisions, keine neuen Risks, keine Renumerierung; DEC-S-366…373 und RISK-314…316 unverändert (nur Note-1-Wortlautkorrektur in DEC-S-366).

## Cross-Document Consistency and ADR Candidate Review (CO-WP-029)

- Ein neues Review-Dokument: [CROSS_DOCUMENT_CONSISTENCY_AND_ADR_CANDIDATE_REVIEW.md](CROSS_DOCUMENT_CONSISTENCY_AND_ADR_CANDIDATE_REVIEW.md). **review-only** — kein Foundation-Modell erweitert, keine neue Autorität, kein neues Statusmodell, keine Technologie, kein ADR.
- **Konsistenzergebnis:** Foundation-weit über **105** Dokumente (26.245 Zeilen) **konsistent**. **0** materielle Contradictions, **0** Parallelmodelle, **0** konkurrierende Autoritätsmodelle. Autoritätsketten (Repository ≠ Runtime; Policy ≠ Approval ≠ Execution Authorization; Recovery Authority eigenständig; Experience ohne Execution; Adapter ohne Scope-Erweiterung) sind durchgängig referenziert statt dupliziert. `CoreScore` bleibt objektbezogen — **kein** globaler Health-Boolean, **kein** normativer Aggregat-Score.
- **Mechanische Korrekturen (6 Klassen, 9 Pfade inkl. Review-Dokument):** (1) **DEC-S-88** wurde in drei Prosa-Referenzen fälschlich als Break-Glass-Entscheidung geführt — autoritativ ist *Session technology* (`deferred`/`non-binding`, CO-WP-009); korrigiert auf **DEC-S-84/DEC-S-85/DEC-S-147**, die authoritative Zeile blieb unverändert. (2) Capability-Count `74 → 94` in vier Current-State-Spiegeln und in der `DEC-O-17`-Notiz. (3) **Domain-Count `12 → 13`** — die Capability Matrix selbst nannte „12", zählte in derselben Aufzählung aber 13 und trägt 13 ID-Präfixe. (4) Zwei fehlgeleitete interne Verweise in einem Deployment-Policy-Dokument. (5) Current-State-Spiegel auf Baseline **`b7827b8`** reconciled (CO-WP-028 ist committet, nicht „ausstehend"). (6) Widersprüchliche CO-WP-024-Commit-Aussage innerhalb desselben Abschnitts bereinigt.
- **Register unverändert in Substanz:** Decision Index `DEC-S-373` — 373 IDs, lückenlos, keine Duplikate. Risk Register 316 — lückenlos, keine Duplikate; Level 170/122/24 = 316 ✓; Status `treatment-planned` 299 / `open` 17 = 316 ✓; die publizierten Summen stimmen exakt mit den ID-Listen überein. **Keine** neue Decision, **kein** neues Risiko, **keine** Renumerierung, **keine** Statusmigration, **keine** Severity-Änderung.
- **ADR-Kandidaten:** **32** (ADR-0001…0030 aus Concept §51 plus zwei Foundation-Klärungen `DEC-A-0031`/`DEC-A-0032`), vollständig inventarisiert und dispositioniert. Ergebnis: **18 `foundation-semantics-established-technical-choice-pending`** · **6 `still-open`** · **3 `deferred-post-foundation`** · **3 `candidate-may-no-longer-be-required`** · **2 `requires-human-decision-before-readiness`** (Summe 32, je Kandidat genau eine primäre Disposition) · **7 Konsolidierungs-Cluster** über 15 Kandidaten. **Kein ADR akzeptiert, keine ADR-Datei erzeugt, keine Technologie ausgewählt.** `Foundation semantics established ≠ ADR accepted`.
- **Strukturbefunde für `CO-WP-030` (MAJOR):** (a) Die Tabellen `DEC-S-38…373` (336 Zeilen) besitzen **keine** `ADR Required`-Spalte — der ADR-Bedarf ist dort aus dem Schema **nicht ableitbar**. Das ist eine Aussage über die *Ableitbarkeit*, **nicht** über den Bedarf: `fehlende Spalte ≠ ADR erforderlich` und `fehlende Spalte ≠ ADR nicht erforderlich`; ebenso `DEC-S Lifecycle Status ≠ ADR-Disposition`. Die readiness-relevante ADR-Basis ist das autoritative `DEC-A`-Inventar (32) plus die Decision-Zuordnungen der Tabellen **mit** `ADR Required`-Spalte plus die CO-WP-029-Dispositionsmatrix plus explizite HM-Entscheidungen — sie ist **vollständig** und durch (a) **nicht** beeinträchtigt. (b) Vier ADR-pflichtige Decisions (`DEC-O-05`, `DEC-O-07`, `DEC-D-03`, `DEC-D-06`) tragen `ADR Required = ja` **ohne** ADR-Kandidaten-ID. (c) Zwei Decision-Konventionen koexistieren weiter (`DEC-S-01…37` kombiniert vs. `DEC-S-38…373` getrennte Dimensionen). (d) `DEC-S-52…317` enthalten ~30 „Technologie deferred"-Decisions; ab `CO-WP-024` gilt die gegenteilige, von Nova bestätigte Regel, dass eine reine Deferral **keine** Decision rechtfertigt.
- **Severity-Kalibrierung (nur Empfehlung):** Das Register veröffentlicht **keine** verbindliche Likelihood × Impact → Level-Matrix. Die von diesem Review aus dem Bestand abgeleitete Modal-Abbildung ist eine **rein analytische Heuristik**, **keine** Registernorm. Gemessen daran fallen **24** Zeilen auf (`RISK-07, 10, 12, 15, 21, 23, 26, 34, 36, 39, 46, 47, 55, 62, 67, 73, 77, 78, 85, 109, 113, 115, 152, 310`). Sie sind **Severity-Kalibrierungs-Reviewkandidaten**, ausdrücklich **keine** nachgewiesenen Fehlbewertungen: `Abweichung von der Heuristik ≠ registrierter Risiko-Defekt`. Da keine Norm existiert, gibt es auch keine verletzte Norm — **alle registrierten Severities bleiben unverändert gültig**; die Liste geht als Kalibrierungskandidat an `CO-WP-030`/Human Maintainer.
- **HM-Inputs (14, `HM-1`…`HM-14`):** Entscheidungen bzw. Inputs, die vor einem **positiven** Foundation-Readiness-Verdikt aufzulösen sind — **keine** Vorbedingung für den *Beginn* von `CO-WP-030`. `CO-WP-030` darf mit sämtlichen unaufgelöst starten und klassifiziert sie dann selbst als `HUMAN DECISION REQUIRED` / `READINESS BLOCKER` / `READINESS NOTE` / `POST-FOUNDATION / NON-BLOCKING`. **Vor Beginn zwingend erforderlich: keiner.** `CO-WP-029` entscheidet das Readiness-Ergebnis nicht.
- **Befund-Lifecycle:** alle 25 Befunde tragen eine explizite Disposition — `corrected` 7 · `open-human-decision` 8 · `open-readiness-review` 5 · `deferred-post-foundation` 1 · `note-only` 4 (Ledger §5.4 des Review-Dokuments). Severity und Disposition sind getrennte Dimensionen: `finding severity ≠ finding lifecycle status`; `MAJOR ≠ automatischer READINESS BLOCKER`; `weitergereicht ≠ stillschweigend akzeptiert`. **Kein** Befund dieses Reviews ist ein Readiness-Blocker.
- **Grenzen:** Kein Readiness-Urteil (`CO-WP-030`), keine Release-Bewertung (`CO-WP-031`), keine Testausführung, kein Integration Lab, kein NDF-Transfer, keine CDS-Adoption. Lessons-Learned-Register und NDF-Feedback-Kandidaten **unverändert**.
- **Nova-Notes-Closure (1–3, geschlossen):** (1) ADR-Readiness-Basis präzisiert (Ableitbarkeit ≠ Bedarf; vier-teilige Basis benannt; HM-Punkte sind keine Startvorbedingung für `CO-WP-030`); (2) Befund-Lifecycle vollständig (25 Befunde mit Disposition, Ledger §5.4, Severity ≠ Lifecycle); (3) Severity-Modal-Abbildung ausdrücklich als analytische Heuristik gekennzeichnet (24 Kalibrierungs-Reviewkandidaten, **keine** nachgewiesenen Fehlbewertungen). Keine neuen Decisions/Risks/ADRs, keine Renumerierung, keine Severity-Änderung; die akzeptierten Korrekturen bleiben unverändert und wurden nicht ausgeweitet.

## Foundation Readiness Review (CO-WP-030)

- Ein neues Review-Dokument: [FOUNDATION_0_1_READINESS_REVIEW.md](FOUNDATION_0_1_READINESS_REVIEW.md). `review-only` mit ausdrücklich begrenzten Dokumentationsschreibungen; Baseline `6afa7ab`. **Kein** Foundation-Modell erweitert, **keine** neue Autorität, **kein** Statusmodell, **keine** Technologie, **kein** ADR, **keine** Release-Bewertung.
- **Verdikt: `FOUNDATION_READINESS: READY WITH NOTES`.** Ausdrücklich **kein** Foundation-inhaltlicher Blocker: kein fehlendes Artefakt, kein materieller Widerspruch, kein Parallelmodell, keine Autoritätsumgehung, keine Standalone-First-Verletzung, kein unbehandelter Foundation-Blocker im Risk Register. Es verbleibt auch **kein** Blocker der Entscheidungsverfügbarkeit mehr.
- **Human-Maintainer-Entscheidungen (Final Readiness Reconciliation):** `HM-1` **`APPROVED`** — Foundation-Release-Taxonomie für Foundation 0.1 verbindlich; `DEC-A-0032` aufgelöst; **keine** Tag- oder Release-Autorisierung. `HM-2` **`APPROVED WITH BOUNDARY`** — Docker-first als Foundation-0.1-Delivery-Baseline; `DEC-A-0031` aufgelöst; bindend `Docker-first ≠ Docker-only`, `≠ zwingende Runtime-Abhängigkeit`, `≠ Observe-Voraussetzung`; autorisiert **nicht** `CO-WP-031`, Foundation Release, Tag-Erzeugung, Observe- oder Technologie-Implementierung. `HM-3` **`APPROVED`** — kriterienbasierte Relevanzregel: Foundation-0.1-relevant ist ein ADR-Kandidat nur, wo seine offene Entscheidung einen verbindlichen Foundation-Vertrag, eine Authority-Grenze oder ein Exit Gate eindeutig bestimmt; reine post-Foundation-Technologie darf `deferred` bleiben.
- **Foundation-relevante ADR-Menge:** Angewandt auf das **bestehende** `CO-WP-029`-Inventar (32 Kandidaten, kein neues Inventar) umfasst sie genau **zwei** Kandidaten — `DEC-A-0032` (durch `HM-1`) und `DEC-A-0031` (durch `HM-2`), beide aufgelöst. **Unaufgelöste Foundation-relevante ADR-Kandidaten: keine.** Die sechs `still-open`-Kandidaten (`ADR-0002`, `ADR-0005`/`CCR-01`, `ADR-0009`/`DEC-O-13`, `ADR-0013`/`CCR-06`, `ADR-0024`/`CCR-08`, `ADR-0030`/`CCR-09`) wurden **einzeln** gegen alle drei `HM-3`-Kriterien geprüft und sind **nicht** Foundation-relevant, weil der tragende Foundation-Vertrag bzw. die Authority-Grenze jeweils bereits anderweitig etabliert ist; sie bleiben **offen** und über `HM-5` sichtbar — **nicht** geschlossen. `ADR-Kandidat existiert ≠ Foundation-relevanter ADR`; `Technologie deferred ≠ Foundation unvollständig`.
- **Warum `READY WITH NOTES` und nicht `READY`:** reale, nicht-blockierende Restpunkte — `HM-4`, `HM-5`, `HM-6`, `HM-8`, `HM-13`, `HM-14` vor dem Foundation-Release, die sechs offenen `CCR`, die Statushygiene der 17 `open`-Risiken (`NF-2`), die Dokumentstatus-Header und `NF-1`…`NF-4`. `REWORK REQUIRED` und `BLOCKED` sind unzutreffend: es fehlt **keine** Foundation-Anforderung und **keine** logisch vorgängige Entscheidung.
- **Exit Gates (24):** **1** `SATISFIED` (Gate 22, Cross-Document Consistency Review) · **22** `SATISFIED WITH NOTE` (einschließlich der durch `HM-1`/`HM-2`/`HM-3` aufgelösten Gates 3 und 8) · **0** `HUMAN DECISION REQUIRED` · **1** `NOT SATISFIED – NON-BLOCKING FOR FOUNDATION 0.1` (Gate 24, `CO-WP-031`) · **0** `NOT SATISFIED – BLOCKING` · **0** `POST-FOUNDATION / NOT APPLICABLE` · **0** `UNKNOWN`. **23 von 24 Gates bewertet und erfüllt**; die übrigen 22 Gates wurden bei der Reconciliation **nicht** erneut geöffnet. **Gate 24 ist Foundation-Arbeit**, die sequenziell auf `CO-WP-030` folgt — ausdrücklich **nicht** post-Foundation und **nicht** „not applicable"; es hindert die Empfehlung nach `CO-WP-031` nicht, bleibt aber vor jeder Release-Reife-Aussage erforderlich (`Foundation-Korpus hinreichend vollständig ≠ Foundation-Phase abgeschlossen`). Gate 23 ist erfüllt, weil das Review **durchgeführt** wurde — `Review-Artefakt existiert ≠ positives Readiness-Gate bestanden`. Der Scope Lock definiert die Gates und wurde **nicht** verändert: `scope definition ≠ readiness evidence`. `SATISFIED WITH NOTE` bedeutet durchgängig *Artefakt existiert, ist tragfähig und committet* — **nicht** *Kontrolle implementiert oder verifiziert*.
- **Vier neue Befunde (`CONFIRMED WITH NEW FINDINGS`), keiner ein Blocker:** `NF-1` — `CO-WP-029` §8.6 nennt „zehn" Einträge, listet aber zwölf IDs; `DEC-O-03` ist `clarified` und `DEC-O-16` ist `verified`, also betrifft `HM-4` genau zehn IDs (`DEC-O-02`, `-04`, `-10`, `-11`, `-12`, `-17`, `-18`, `-19`, `-20`, `-21`). `NF-2` — 17 Risiken tragen `open`, obwohl ihre Target-WPs sämtlich abgeschlossen sind und das geforderte Artefakt jeweils existiert; behandelt ist die Substanz, nicht der Statuswert. `NF-3` — zehn der elf auf `CO-WP-030` gerichteten Risiken beschreiben Design-Treatments und sind durch ein Review nicht behandelbar; nur `RISK-314` benennt das Readiness Review als Durchsetzungsort und ist erfüllt. `NF-4` — der Gate-Bezug „Security-Invarianten" ist zwischen Concept §52 (20 Verbote, `DEC-G-08`) und Threat Model §14 (17 Design Requirements) nicht benannt; beide Sätze sind widerspruchsfrei und werden downstream konsistent verwendet.
- **Standalone-First `PASS`:** **kein** Foundation-Dokument unter `docs/` nennt Core Vision, Core Brain, Core-Dev oder CDF. CDS ist ausdrücklich `Adoption: Not started` / `Pilot: Inactive`; MCP, Datenhaltung und Event Bus sind nicht ausgewählt. [Sovereignty Policy §3/§4](../docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) verankert: `optionale Integration ≠ Laufzeitvoraussetzung`.
- **Security `PASS WITH PRECONDITIONS`, kein Foundation Blocker:** ein künftiger read-only Observe-Slice müsste **keine** der neun Grenzen umgehen (read-only-before-write · target scope · provenance · source trust · permission denial · secrets boundary · fail-closed · keine implizite Management-Autorität · keine verdeckte Execution-Autorität). `CCR-05`/`CCR-07` sind vor **Deploy** zu schließen, nicht vor Observe.
- **Testing/Verification `PASS WITH NOTES`:** Exit Gate 19 fordert Strategie, nicht Ausführung. Nicht etabliert bleiben Testimplementierung, Testausführung, Coverage-Messung (überall `not measured`), Lab-Bereitstellung und Foundation-Validierung. Die 21 auf `CO-WP-028` gerichteten Risiken bleiben unverändert `treatment-planned`.
- **14 HM-Inputs vollständig klassifiziert:** **3 durch den Human Maintainer entschieden** (`HM-1` resolved · `HM-3` resolved · `HM-2` resolved with boundary) · **6** vor Foundation-Release (`HM-4`, `HM-5`, `HM-6`, `HM-8`, `HM-13`, `HM-14`) · **5** aufschiebbar (`HM-7`, `HM-9`, `HM-10`, `HM-11`, `HM-12`) · **0** informational · **0** obsolet. `HM-4`…`HM-14` bleiben unentschieden und behalten ihre Timing-Klassifikationen unverändert. **Human decisions made by Claude: 0** — die drei Entscheidungen stammen vom Human Maintainer und wurden hier nur protokolliert und angewandt.
- **Register unabhängig nachgerechnet:** Decision Index 373 `DEC-S` lückenlos und duplikatfrei. Risk Register 316 lückenlos und duplikatfrei; Level `high 170` / `medium 122` / `low 24` und Status `treatment-planned 299` / `open 17` exakt — vollständig deckungsgleich mit `CO-WP-029`, **keine** Abweichung. **Keine** Registeränderung vorgenommen.
- **Observe `READY WITH PRECONDITIONS`:** alle zehn Elemente des vordeklarierten Wertkriteriums haben einen benannten Foundation-Eigentümer; das Kriterium wurde **nicht** abgeschwächt. Kandidat *Local Linux Host Identity & Basic System Observation* weiterhin gültig (`CAP-DISCOVERY-004`, `target-observe`, `read`, `not-implemented`/`not-supported`). Harte Preconditions: `P-1` separate Human-Maintainer-Zielautorisierung · `P-2` No-Mutation-Vertrag durch **beobachtetes** geschütztes Verhalten belegt (`an error message appeared ≠ protected behavior observed`; derzeit null Tests implementiert und ausgeführt) · `P-3` Transport-/Erhebungsentscheidung für `CAP-DISCOVERY-004`. **Keine** Observe-WP-ID, **keine** Implementierung, **kein** Zielzugriff.
- **Semantische Grenze:** CoreOps besitzt `operational condition`, `operational state`/`operational mode` und `operational severity`/`impact` bereits mit benannten autoritativen Quellen. `BOUNDARY_RECONCILIATION_REQUIRED` gegenüber CDS besteht, ist im Repository bereits als `future mapping candidate` registriert und wird als **POST-FOUNDATION** klassifiziert — weder Foundation Blocker noch `MUST CLOSE BEFORE OBSERVE`. **Kein** fremdes Vokabular importiert.
- **Grenzen:** keine Testausführung, keine Laufzeitprüfung, kein realer Zielzugriff, kein Netzwerk-, Credential- oder Secret-Zugriff, keine Technologieauswahl, **kein ADR akzeptiert, abgelehnt oder erzeugt**, keine neuen Decisions/Risks/ADRs, keine Severity-, Target- oder Statusänderung, keine Decision-Statusmigration, keine Prosa-Bereinigung, kein Git-Write. Evidenzstärke **`moderate`**: Register-, ID-, Tally- und Referenzprüfungen deterministisch (`strong`), Gate-Bewertungen dokumentarisch durch eine einzelne Reviewinstanz (`limited`).

## Letztes Work Package

`CO-WP-030 – Foundation Readiness Review` (`review-only` mit ausdrücklich begrenzten Dokumentationsschreibungen) — durch ausdrückliche Human-Maintainer-Freigabe und eine Scope Clarification bearbeitet; Nova Review `GO WITH NOTES`, Notes-Runde (1–3) geschlossen, **Nova Final Review `GO`**; **`completed-go-with-notes`**; Human-Maintainer-Commit ausstehend. Repository-Baseline `6afa7ab`. Ein neues Dokument: [FOUNDATION_0_1_READINESS_REVIEW.md](FOUNDATION_0_1_READINESS_REVIEW.md). Verdikt **`FOUNDATION_READINESS: READY WITH NOTES`** nach Auflösung von `HM-1`, `HM-2` und `HM-3` durch den Human Maintainer; **kein** Foundation-inhaltlicher Blocker und **kein** verbleibender Blocker der Entscheidungsverfügbarkeit; erforderliche Foundation-Nacharbeit **keine**. Empfehlung `CO-WP-031: PROCEED WITH NOTES`. Vorher: `CO-WP-029 – Cross-Document Consistency and ADR Candidate Review` (review-only) — Nova Review `GO WITH NOTES`; Notes-Runde (Notes 1–3) geschlossen; **`completed-go-with-notes`**; **committet und gepusht als `6afa7ab`**. Davor: `CO-WP-028 – Test Strategy, Fixtures and Integration Lab` (docs-only) — Nova Review **`GO WITH NOTES`**, Notes-Runde (1–3) geschlossen, **`completed-go-with-notes`**, Human-Maintainer-Commit **`b7827b8`** (`b7827b89f76aba61fb255cfbf5a6682d4191cefe`, gepusht; Branch gleichauf mit origin/main). Davor: `CO-WP-027 – UX Information Architecture and Dashboard System` (docs-only) — Nova Review **`GO WITH NOTES`**, Notes-Runde (1–3) geschlossen, **`completed-go-with-notes`**, Human-Maintainer-Commit **`1dee29d`** (`1dee29d19fe500588161c5878cbbb76e5dbb0812`, gepusht; Branch gleichauf mit origin/main). Zuvor: `Foundation Milestone Review CO-WP-021…026` (unnummeriert, review-only / docs-only) — Nova Review `GO WITH NOTES`, nachgelagerte Incident-Coverage-Tally-Korrektur `GO`, `completed-go-with-notes`, Human-Maintainer-Commit **`2e1ab66`** (gepusht; Branch gleichauf mit origin/main). Letztes reguläres WP davor: `CO-WP-026 – completed-go-with-notes` (Commit 399de21, gepusht). Davor: `CO-WP-025 – completed-go-with-notes` (Commit 3419664, gepusht).

## Zukünftige Roadmap und Milestone-Vormerkung

- **Reporting-/Vulnerability-Roadmap** (in [WORK_PACKAGE_QUEUE.md](../project-system/WORK_PACKAGE_QUEUE.md) registriert, `roadmap-candidate` · `not scheduled` · `WP identifier pending queue review` · `not implemented`, **keine WP-Nummern, keine Decision/kein Risk, keine Technologie**): Reporting Foundation → Asset and Component Inventory → Vulnerability Intelligence Ingestion → Vulnerability Correlation → Exposure and Remediation → Reporting Implementation. Reporting: professionelle PDF-Berichte, CoreOps-Design/Logo/Mandantenprofile, DE/EN, Inventar-/Log-/Update-/Deployment-/Topologie-/Audit-/Evidence-/Vulnerability-Berichte, Redaction vor Rendering, getrennte Disclosure-/Export-Autorisierung, keine Raw Secrets. Vulnerability: Inventar (HW/OS/Firmware/SW/Pakete/Container/Images/SBOM), Identitäten (CPE/purl/Firmware-ID/Image Digest/SBOM Component), Quellen (CVE/NVD/Advisories/CISA KEV/EPSS/VEX), Match Confidence, getrennte Zustände, Produktnamensmatch ≠ Betroffenheit, Offline-Snapshots/CorePacks, Remediation über Deployment Governance.
- **Milestone-Hinweis:** Der gebündelte `CO-WP-021…026` Foundation Milestone Review ist **abgeschlossen** (`completed-go-with-notes`, Nova Review `GO WITH NOTES`); er hat die Reporting-/Vulnerability-Roadmap bewertet und die Governance-vor-Capability-Sequenzierung bestätigt, **ohne** WP-Nummern zu vergeben oder eine Capability freizugeben.

## Nächstes Work Package

**Verbindlicher nächster Schritt: Human-Maintainer-Commit von `CO-WP-030`.** `CO-WP-030 – Foundation Readiness Review` ist durch ausdrückliche Human-Maintainer-Freigabe und eine Scope Clarification bearbeitet; Nova Review `GO WITH NOTES`, Notes-Runde (1–3) geschlossen, **Nova Final Review `GO`**, Status `completed-go-with-notes`. Das Readiness-Verdikt ist **`FOUNDATION_READINESS: READY WITH NOTES`**: `HM-1`, `HM-2` und `HM-3` sind durch den Human Maintainer entschieden und in der eng begrenzten Neubewertung der Exit Gates 3 und 8 angewandt; die übrigen 22 Gates wurden **nicht** erneut geöffnet. Erforderliche Foundation-Nacharbeit: **keine**; es verbleibt **kein** Blocker. Das vorangegangene `CO-WP-029` ist `completed-go-with-notes` und committet (`6afa7ab`, gepusht); davor `CO-WP-028` (`b7827b8`, gepusht), `CO-WP-027` (`1dee29d`, gepusht) und der `CO-WP-021…026` Foundation Milestone Review (`2e1ab66`, gepusht). `CO-WP-031` bleibt **retained**, ist nicht freigegeben, nicht begonnen und nicht automatisch autorisiert; die Empfehlung **`CO-WP-031: PROCEED WITH NOTES`** ist eine Empfehlung, **keine** Autorisierung (`Empfehlung ≠ Autorisierung`). Foundation Readiness ist mit `CO-WP-030` bewertet; das Ergebnis ist **kein** Release-Urteil: `READY WITH NOTES = bereit zum Eintritt in die Foundation Release Preparation`, **nicht** `Foundation 0.1 released` und **nicht** `Foundation-Phase abgeschlossen`. Der nicht-aktivierende CDS-Gate-Hinweis wurde in `CO-WP-027` und `CO-WP-028` eingehalten (CDS read-only bzw. gar nicht genutzt, keine Adoption, kein Pilot, keine CDS-Testevidenz) und bleibt vor jeder substanziellen Designübernahme erneut zu prüfen.

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
