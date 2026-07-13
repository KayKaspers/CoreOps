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

## Letztes Work Package

`CO-WP-004C – BSI and Public-Sector Readiness Baseline` (docs-only / security-governance baseline). Umsetzung abgeschlossen, Nova Review ausstehend. Vorheriges WP: `CO-WP-004B4 – GO WITH NOTES`.

## Nächstes Work Package

`CO-WP-004D – ITIL and PRINCE2 Applicability and Tailoring Decision` (planned-next; pending Nova review von CO-WP-004C und Human-Maintainer-Freigabe). Danach `CO-WP-004E`. `CO-WP-005` bleibt `planned-after-CO-WP-004E`.

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
