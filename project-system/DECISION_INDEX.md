# CoreOps – Decision Index

> Initialer Decision Index. NDF-Basis: `v1.0.0` (Commit `9dcadc1`) — `main` nicht normativ.
> Erzeugt durch `CO-WP-002`. Quelle der Einträge: [COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md) und [CONCEPT_DECISION_CLASSIFICATION.md](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md).

**Grundregeln:** Keine technische Entscheidung erhält `accepted`. Produktentscheidungen dürfen `accepted-product` sein, wenn sie im Concept ausdrücklich als akzeptierte Vision vorgegeben sind. ADRs sind ausschließlich `adr-candidate`.

Statuswerte: `accepted-product` · `binding-governance` · `open` · `adr-candidate` · `deferred` · `non-goal` · `blocked`
Zusatzstatus (eingeführt in CO-WP-003, alle vor dem Human-Maintainer-Commit unverbindlich): `proposed` · `proposed-binding-governance` · `clarified` · `verified`. Diese kennzeichnen in CO-WP-003 aufgelöste Governance-/Scope-Punkte, die mit dem Human-Maintainer-Commit bindend werden.
Zusatzstatus für Produkt-/Strategierichtung (eingeführt in CO-WP-004A): `accepted-product-direction` · `prohibited` · `not-claimed` · `not-current-target` · `binding-governance-direction` · `controlled-candidate-process` · `foundation-candidate`. Keiner dieser Werte markiert eine technische Architektur- oder Technologieentscheidung als `accepted`.
Owner-Werte: `Human Maintainer` (HM) · `Nova`

## Accepted Product Decisions

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-P-01 | Universelle, herstellerunabhängige Operations Control Plane | accepted-product | APC | Concept §3,§5 | HM | CO-WP-003 | ja (ADR-0001) | Vision, nicht Core-exklusiv |
| DEC-P-02 | Self-hosted, offline-/Air-Gap-fähig, keine Cloudpflicht | accepted-product | APC | §3,§9.5,§38 | HM | CO-WP-023 | ja (ADR-0004) | Offline First |
| DEC-P-03 | Zweisprachigkeit DE/EN ab Foundation | accepted-product | APC | §3,§53 | HM | CO-WP-005 | nein | Deutsch als möglicher Default |
| DEC-P-04 | Produktformel + Read-only vor Write | accepted-product | APC | §8,§9.3 | HM | CO-WP-011 | ja (ADR-0003) | Sicherheitsziel |
| DEC-P-05 | KI nur beratend, kein autonomer Admin | accepted-product | APC | §7,§45 | HM | CO-WP-013 | ja (ADR-0029) | zugleich Non-Goal |
| DEC-P-06 | Fehlende Daten ≠ gesund | accepted-product | APC | §25 | HM | CO-WP-004 | nein | CoreScore-Ehrlichkeit |
| DEC-P-07 | Trennung Observed/Desired/Effective + Drift, kein Enforce im MVP | accepted-product | APC | §20 | HM | CO-WP-012 | ja (ADR-0010/0011) | — |
| DEC-P-08 | Privacy by Design (Klassen, Redaction, Retention) | accepted-product | APC | §41 | HM | CO-WP-025 | ja (ADR-0024) | — |

## Binding Governance Decisions

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-G-01 | Human-Maintainer-only für Freigabe/Commit/Push/Merge/Tags/Releases/Deployments | binding-governance | BG | §1 | HM | laufend | nein | NDF-Kern |
| DEC-G-02 | Ein Work Package pro Commit, keine verdeckte Scope-Erweiterung | binding-governance | BG | §1,§2 | Nova | laufend | nein | — |
| DEC-G-03 | Nova reviewt/bewertet; Agent committet/pusht/taggt nicht | binding-governance | BG | §1 | Nova | laufend | nein | — |
| DEC-G-04 | Fail Closed, Read-only vor Write, Preview vor Execute, Backup vor Änderung, Verifikation | binding-governance | BG | §2,§9.6 | Nova | CO-WP-007 | nein | — |
| DEC-G-05 | Skills-first ohne Netzwerk-/Secret-/Ausführungs-/Git-Autonomie | binding-governance | BG | §2 | Nova | laufend | nein | siehe NDF Skills Provenance |
| DEC-G-06 | Public Neutrality; keine Secrets in Logs/Prompts/Exports | binding-governance | BG | §2,§52 | Nova | CO-WP-005 | nein | — |
| DEC-G-07 | Control Plane nicht durch Plugins/Agents umgehbar | binding-governance | BG | §12,§16 | Nova | CO-WP-013 | ja (ADR-0022) | — |
| DEC-G-08 | 20 Sicherheitsinvarianten als verbindliches Zielbild | binding-governance | BG | §52 | Nova | CO-WP-007 | teils | — |

## Open Foundation Decisions

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-O-01 | Plane-Taxonomie (Domain Packs als Plane?) | open | CCR | §10 | Nova | CO-WP-006/008 | ja (ADR-0005) | CCR-01 |
| DEC-O-02 | Release-Taxonomie / SemVer / 0.1-Kollision | proposed | CCR | §49 | HM | CO-WP-003 | nein | CCR-02 aufgelöst in RELEASE_TAXONOMY.md (Proposed for acceptance); Tag-Kandidaten `v0.0.1-foundation`, `v0.1.0-alpha.1` |
| DEC-O-03 | NDF-Level (1 vs. 2) | clarified | CCR | Manifest | HM | CO-WP-005 | nein | CCR-03: Semantik geklärt (Bootstrap-Status), Zahlenwert unverändert; Quellambiguität dokumentiert, kein Blocker |
| DEC-O-04 | `main`-Normativität (Concept vs. Governance) | proposed-binding-governance | BG | §Ausgangskontext | HM | CO-WP-003 | nein | CCR-04 geklärt: `v1.0.0`/`9dcadc1` normativ, `main` informativ; Übernahme nur via eigenes WP |
| DEC-O-05 | Offline-Policy bei getrennter Control Plane | open | CCR/FR | §21,§38 | Nova | CO-WP-013/023 | ja | CCR-05 |
| DEC-O-06 | Machine Identity vs. Air-Gap-Laufzeiten | open | CCR/FR | §31,§38 | Nova | CO-WP-010 | ja (ADR-0013) | CCR-06 |
| DEC-O-07 | Privilegierte Ausführung vs. „keine Remote-Root-Shell" | open | CCR/FR | §52 | Nova | CO-WP-013 | ja | CCR-07 |
| DEC-O-08 | Immutable Audit vs. Redaction/Retention | open | CCR/FR | §35,§41 | Nova | CO-WP-025 | ja (ADR-0024) | CCR-08 |
| DEC-O-09 | Offline-First-Facetten (Runtime/Install/Update/Recovery/Build) | open | CCR/FR | §38 | Nova | CO-WP-023 | ja (ADR-0030) | CCR-09 |
| DEC-O-10 | Docker-first: Anforderung/Baseline/Architektur/Kandidat | proposed | APC | §3,§53 | HM | CO-WP-003 | nein | CCR-10 eingeordnet: akzeptierte Delivery-/Betriebsanforderung (Compose-Standardinstallation), keine Anwendungsarchitektur, kein K8s-Zwang; noch nicht implementiert |
| DEC-O-11 | Foundation-Queue-Autorität (Concept §50 vs. aktive Queue) | proposed | CCR | §50 | Nova | CO-WP-003 | nein | CCR-11 geklärt: aktive Queue ausschließlich WORK_PACKAGE_QUEUE.md; Concept-Queue historischer Vorschlag |
| DEC-O-16 | Repository-Referenz | verified | — | origin | HM | CO-WP-003 | nein | `https://github.com/KayKaspers/CoreOps` verifiziert und im Manifest gesetzt |
| DEC-O-12 | Herstellersupport-Grenze | proposed | APC/CCR | §18,§26,§27 | Nova | CO-WP-004 | nein | CCR-12 vorgeschlagen aufgelöst in INITIAL_SUPPORT_BOUNDARY.md: Nennung = Kandidat, kein Support/Partnerschaft/Zertifizierung |
| DEC-O-13 | Source-of-Truth-Konfliktprioritäten je Datenklasse | open | FR/AC | §19 | Nova | CO-WP-011 | ja (ADR-0009) | — |
| DEC-O-17 | Foundation Capability Matrix | proposed | — | CO-WP-004 | Nova | CO-WP-004 | nein | Planungs-/Governance-Landkarte; 74 Capabilities, alle not-implemented/not-supported |
| DEC-O-18 | Initial Observe Support Boundary | proposed | — | CO-WP-004 | HM | CO-WP-004 | nein | Observe read-only bis Level 2; Write ausgeschlossen |
| DEC-O-19 | Drei getrennte Statusdimensionen | proposed-binding-governance | BG | CO-WP-004 | Nova | CO-WP-004 | nein | Roadmap/Implementation/Support unabhängig |
| DEC-O-20 | Support-Evidence-Anforderungen | proposed-binding-governance | BG | CO-WP-004 | Nova | CO-WP-004 | nein | 21-Punkte-Evidence-Satz; ohne Evidenz max. experimental |
| DEC-O-21 | Herstellerlisten sind keine Supportzusage | proposed-binding-governance | BG | CO-WP-004 | Nova | CO-WP-004 | nein | Vendor priority ≠ vendor support |
| DEC-O-14 | OIC v0.1 Vertragsumfang | open | FR | §17,§18 | Nova | CO-WP-014 | ja (ADR-0006) | — |
| DEC-O-15 | Topologie-Evidence und manuelle Autorität | open | FR/AC | §28 | Nova | CO-WP-020 | ja (ADR-0020/0021) | — |

## ADR Candidates

Alle 30 Concept-ADR-Kandidaten (§51) plus zwei Foundation-Klärungen (Delivery Baseline, Release-Taxonomie) sind mit Status `adr-candidate` erfasst. Vollständige Liste: [CONCEPT_DECISION_CLASSIFICATION.md → ADR-Required Topics](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md). **Kein ADR ist Accepted; es wurden keine ADR-Dateien erzeugt.**

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ |
| DEC-A-0001…0030 | Concept-ADR-Kandidaten ADR-0001…ADR-0030 | adr-candidate | ADR | §51 | HM | diverse Foundation-WPs | ja |
| DEC-A-0031 | Delivery Baseline (Docker-first-Klassifikation) | adr-candidate | ADR | §3,§53 | HM | CO-WP-003 | ja |
| DEC-A-0032 | Release-Taxonomie / SemVer | adr-candidate | ADR | §49 | HM | CO-WP-003 | ja |

## Product Direction and Governance Decisions (CO-WP-004A)

> `accepted-product-direction` = durch den Human Maintainer akzeptierte Produkt-/Strategierichtung (kein technischer Architektur-/Technologie-`accepted`). Registriert über [COREOPS_CONCEPT_V3_1_AMENDMENT.md](../docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md).

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-S-01 | Product Sovereignty | accepted-product-direction | APC | CO-WP-004A | HM | CO-WP-004A | nein | Eigenständiges Kernprodukt |
| DEC-S-02 | Mandatory external management products | prohibited | BG/NG | CO-WP-004A | HM | CO-WP-004A | nein | Nicht verpflichtend für CoreOps-Kern |
| DEC-S-03 | Technical foundation dependencies | open | FR | CO-WP-004A | Nova | later review | teils | Keine konkrete Dependency akzeptiert |
| DEC-S-04 | BSI-oriented development | accepted-product-direction | APC | CO-WP-004A | HM | CO-WP-007/Baseline | nein | Orientierung, keine Konformität |
| DEC-S-05 | Government Profile | accepted-product-direction | RC | CO-WP-004A | HM | später | nein | Roadmap-Richtung, keine Zertifizierung |
| DEC-S-06 | BSI certification | not-claimed | NG | CO-WP-004A | HM | — | nein | Keine Zertifizierung behauptet |
| DEC-S-07 | VS-NfD | not-current-target | NG | CO-WP-004A | HM | — | nein | Kein aktuelles Releaseziel |
| DEC-S-08 | Lessons Learned | binding-governance-direction | BG | CO-WP-004A | Nova | CO-WP-004B | nein | Prüfpflicht; Detailprozess in 004B |
| DEC-S-09 | NDF Feedback | controlled-candidate-process | BG | CO-WP-004A | HM | CO-WP-004B | nein | Übernahme nur via eigenes NDF-WP |
| DEC-S-10 | ITIL alignment | foundation-candidate | RC | CO-WP-004A | Nova | CO-WP-004D | nein | Nur Kandidat |
| DEC-S-11 | PRINCE2-derived governance | foundation-candidate | RC | CO-WP-004A | Nova | CO-WP-004D | nein | Nur Kandidat |
| DEC-S-12 | Lessons-Learned-Prozess | proposed-binding-governance | BG | CO-WP-004B | Nova | CO-WP-004B | nein | [LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md) |
| DEC-S-13 | NDF-Feedback-Prozess | proposed-binding-governance | BG | CO-WP-004B | Nova | CO-WP-004B | nein | [NDF_FEEDBACK_PROCESS.md](../docs/governance/NDF_FEEDBACK_PROCESS.md) |
| DEC-S-14 | Kein automatischer NDF-Rückfluss | proposed-binding-governance | BG | CO-WP-004B | HM | CO-WP-004B | nein | „No Automatic Synchronization" §18 |
| DEC-S-15 | Human-Maintainer-Gate für NDF-Transfer | proposed-binding-governance | BG | CO-WP-004B | HM | CO-WP-004B | nein | Nur Nova+HM dürfen `approved-for-transfer` setzen |

## Public-Sector Readiness Decisions (CO-WP-004C)

> Registriert über [BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](../docs/security/BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md), [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](../docs/security/BSI_REFERENCE_AND_CLAIMS_REGISTER.md) und [PUBLIC_SECTOR_READINESS_PROFILE.md](../docs/governance/PUBLIC_SECTOR_READINESS_PROFILE.md). Keine technische Architektur ausgewählt, keine ADR erzeugt, keine Zertifizierung behauptet.

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-S-16 | Internal profiles Standard, Hardened and Government | accepted-product-direction | RC | CO-WP-004C | HM | CO-WP-004E | nein | Accepted product and governance direction; interne Profile, keine Zertifizierungsstufen |
| DEC-S-17 | BSI-oriented and IT-Grundschutz-aligned positioning | accepted-product-direction | APC | CO-WP-004C | HM | later baseline | nein | Accepted **with explicit claim boundaries**; keine Konformität/Zertifizierung |
| DEC-S-18 | Public-sector readiness baseline | accepted-product-direction | APC | CO-WP-004C | HM | CO-WP-004E | nein | Accepted **foundation baseline**; 18 PSR-Domänen, Verantwortungs-/Evidenzmodell |
| DEC-S-19 | Detailed BSI requirement mapping | deferred | DF | CO-WP-004C | Nova | later ingestion WP | teils | Deferred to later work; versionsgenaue Source-Ingestion erforderlich |
| DEC-S-20 | Government profile certification or approval | not-claimed | NG | CO-WP-004C | HM | — | nein | Not claimed; Government Profile ist internes Readiness-Profil |
| DEC-S-21 | Cloud C5/C3A applicability | open | CCR | CO-WP-004C | Nova | CO-WP-023 | teils | **Conditional and deployment-dependent**; nur bei Cloud-Szenarien relevant |
| DEC-S-22 | VS-NfD suitability | not-claimed | NG | CO-WP-004C | HM | — | nein | **Not assessed and not claimed** (vgl. DEC-S-07 not-current-target) |

## Framework Tailoring Decisions (CO-WP-004D)

> Registriert über [ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md](../docs/governance/ITIL_AND_PRINCE2_APPLICABILITY_AND_TAILORING.md) und [COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md](../docs/governance/COREOPS_SERVICE_AND_PROJECT_GOVERNANCE_PROFILES.md). Löst die Kandidaten DEC-S-10 (ITIL) und DEC-S-11 (PRINCE2) auf. Keine Zertifizierung/Endorsement, keine Tool-Abhängigkeit, keine ADR.

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-S-23 | Selected ITIL concepts | adopted-with-tailoring | RC | CO-WP-004D | HM | CO-WP-004E | nein | Service-management guidance; löst DEC-S-10 |
| DEC-S-24 | ITIL 4 and ITIL Version 5 | binding-governance-direction | RC | CO-WP-004D | Nova | later ops WP | nein | Both tracked **with explicit version boundaries**; Bridge-Pfad anerkannt |
| DEC-S-25 | Full ITIL implementation | rejected | NG | CO-WP-004D | HM | — | nein | Rejected; Framework-Overload-Schutz |
| DEC-S-26 | PRINCE2 Project Management Version 7 | optional-profile | RC | CO-WP-004D | HM | CO-WP-004E | nein | Optional project/deployment governance guidance; löst DEC-S-11 |
| DEC-S-27 | Full PRINCE2 implementation | rejected | NG | CO-WP-004D | HM | — | nein | Rejected; NDF bleibt maßgeblich |
| DEC-S-28 | NDF primacy | binding-governance-direction | BG | CO-WP-004D | Nova | laufend | nein | NDF remains primary software-development and repository-governance framework |
| DEC-S-29 | External certification or endorsement | not-claimed | NG | CO-WP-004D | HM | — | nein | Not claimed; keine PeopleCert-/Akkreditierungs-Claims |
| DEC-S-30 | Framework-specific tooling dependency | not-selected | NG | CO-WP-004D | HM | — | nein | Kein ITSM-/PM-Produkt ausgewählt |

## Capability Governance Alignment Decisions (CO-WP-004E)

> Registriert über [CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md](../docs/security/CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md), [CAPABILITY_MATRIX_SPEC.md](../docs/project-system/CAPABILITY_MATRIX_SPEC.md) und der ausgerichteten [FOUNDATION_CAPABILITY_MATRIX.md](../docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md). Erweitert DEC-O-19 (drei Statusdimensionen). Keine Capability implementiert, keine BSI-Compliance, keine ADR.

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ | ----- |
| DEC-S-31 | Multi-dimensional capability status | accepted-product-direction | BG | CO-WP-004E | Nova | CO-WP-004E | nein | Accepted für die Foundation Capability Matrix; erweitert DEC-O-19 |
| DEC-S-32 | Roadmap/Implementation/Support/Evidence | binding-governance-direction | BG | CO-WP-004E | Nova | laufend | nein | Separate, unabhängige Dimensionen (+ Security/Governance) |
| DEC-S-33 | PSR-domain mapping | accepted-product-direction | RC | CO-WP-004E | Nova | later mapping WP | nein | Readiness-Relationship, **nicht** Compliance-Mapping |
| DEC-S-34 | Responsibility mapping | binding-governance-direction | BG | CO-WP-004E | Nova | CO-WP-004E | nein | Product/Operator/Shared-Dimensionen erforderlich |
| DEC-S-35 | Evidence capability | binding-governance-direction | BG | CO-WP-004E | Nova | laufend | nein | Getrennt von deployment evidence und requirement satisfaction |
| DEC-S-36 | Detailed BSI control mapping | deferred | DF | CO-WP-004E | Nova | later ingestion WP | teils | Deferred; versionsgenaue Source-Ingestion nötig |
| DEC-S-37 | Capability certification or compliance status | not-claimed | NG | CO-WP-004E | HM | — | nein | Not claimed; kein `compliant`-Wert in der Matrix |

## Language and Repository Governance Decisions (CO-WP-005)

> Registriert über [COREOPS_LANGUAGE_STANDARD.md](../docs/governance/COREOPS_LANGUAGE_STANDARD.md), [PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](../docs/governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md) und [REPOSITORY_GOVERNANCE_STANDARD.md](../docs/governance/REPOSITORY_GOVERNANCE_STANDARD.md). **Dimensionen getrennt** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte. Keine ADR, keine Technologieauswahl, keine automatisierte Durchsetzung.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-38 | Canonical language | governance-direction | accepted | binding-governance | CO-WP-005 | HM | English canonical for machine-facing/technical identifiers; DE/EN primary product languages |
| DEC-S-39 | Translation parity | governance-direction | accepted | binding-governance | CO-WP-005 | Nova | Parity must be explicit and evidence-based; no automatic full-bilingual claim |
| DEC-S-40 | Public neutrality and disclosure | governance-direction | accepted | binding-governance | CO-WP-005 | HM | Public artifacts organisation-/vendor-neutral, free of unnecessary private infrastructure data |
| DEC-S-41 | Human-Maintainer repository gates | governance-direction | accepted | binding-governance | CO-WP-005 | HM | HM gates mandatory for git writes, commits, pushes, tags, releases and final governance decisions |
| DEC-S-42 | Source-of-truth precedence | governance-direction | accepted | binding-governance | CO-WP-005 | Nova | Summaries and derived artifacts must not override authoritative sources |
| DEC-S-43 | Encoding and line endings | governance-direction | accepted | binding-governance | CO-WP-005 | Nova | New text artifacts use UTF-8; existing line-ending conventions preserved unless a dedicated migration is approved |

## System Context and Architecture Taxonomy Decisions (CO-WP-006)

> Registriert über [SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md](../docs/architecture/SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md), [COREOPS_PLANE_TAXONOMY.md](../docs/architecture/COREOPS_PLANE_TAXONOMY.md) und [TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md](../docs/security/TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Technologieauswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-44 | System context | architecture-context | accepted | binding-governance | CO-WP-006 | HM | Accepted foundation architecture context; technologieunabhängig |
| DEC-S-45 | Plane taxonomy | architecture-context | accepted | binding-governance | CO-WP-006 | Nova | Accepted conceptual architecture taxonomy (10 Planes) |
| DEC-S-46 | Plane semantics | architecture-context | accepted | guidance | CO-WP-006 | Nova | Logische Verantwortungs-/Datenflussbereiche, **keine** verpflichtenden Deployment-Einheiten |
| DEC-S-47 | Managed resources | architecture-context | accepted | binding-governance | CO-WP-006 | Nova | Außerhalb der CoreOps-Produktgrenze |
| DEC-S-48 | External services | architecture-context | accepted | binding-governance | CO-WP-006 | HM | Optional, sofern nicht später ausdrücklich akzeptiert |
| DEC-S-49 | Agent plane | architecture-context | accepted | binding-governance | CO-WP-006 | Nova | Optional; agentless-Betrieb bleibt möglich |
| DEC-S-50 | Offline core | product-direction | accepted | binding-governance | CO-WP-006 | HM | Accepted product and governance direction; lokale Kernfunktion ohne Cloudpflicht |
| DEC-S-51 | Control authority | governance-direction | accepted | binding-governance | CO-WP-006 | Nova | Explizite Autorisierung für Write- und Execution-Aktionen erforderlich |
| DEC-S-52 | Detailed technology architecture | architecture-context | deferred | non-binding | CO-WP-006 | Nova | Deferred; spätere Architektur-WPs/ADRs |
| DEC-S-53 | Detailed threat model | security-context | deferred | non-binding | CO-WP-006 | Nova | Deferred zu CO-WP-007 |

## Threat Model and Security Baseline Decisions (CO-WP-007)

> Registriert über [COREOPS_FOUNDATION_THREAT_MODEL.md](../docs/security/COREOPS_FOUNDATION_THREAT_MODEL.md) und [THREAT_SCENARIO_REGISTER.md](../docs/security/THREAT_SCENARIO_REGISTER.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine Sicherheitskontrolle implementiert; keine Technologie-/Kryptoauswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-54 | Foundation threat model | security-context | accepted | binding-governance | CO-WP-007 | HM | Accepted security-governance baseline; nicht implementiert/validiert |
| DEC-S-55 | Threat scenario register | security-context | accepted | binding-governance | CO-WP-007 | Nova | Accepted authoritative threat inventory (THR-001…040) |
| DEC-S-56 | Threat IDs | governance-direction | accepted | binding-governance | CO-WP-007 | Nova | Stabil, nicht wiederverwendet, nicht still gelöscht |
| DEC-S-57 | Threat and risk ratings | governance-direction | accepted | guidance | CO-WP-007 | Nova | Qualitativ, evidence-bounded; keine numerische Präzision |
| DEC-S-58 | Security invariants | security-context | accepted | binding-governance | CO-WP-007 | Nova | Bindende Designanforderungen, **keine** implementierten Kontrollen |
| DEC-S-59 | Mitigation implementation | security-context | not-claimed | non-binding | CO-WP-007 | HM | Not claimed; keine Kontrolle implementiert |
| DEC-S-60 | Mitigation validation | security-context | deferred | non-binding | CO-WP-007 | Nova | Deferred; keine Validierung durchgeführt |
| DEC-S-61 | Penetration testing | security-context | deferred | non-binding | CO-WP-007 | HM | Deferred; kein Pentest durchgeführt |
| DEC-S-62 | Detailed security architecture | architecture-context | deferred | non-binding | CO-WP-007 | Nova | Deferred zu späteren Security-Architektur-WPs |
| DEC-S-63 | Technology and crypto selection | architecture-context | deferred | non-binding | CO-WP-007 | Nova | Deferred; keine Krypto-/Authentisierungs-/Netzauswahl |

## Logical Module Architecture Decisions (CO-WP-008)

> Registriert über [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](../docs/architecture/COREOPS_LOGICAL_MODULE_ARCHITECTURE.md), [COREOPS_MODULE_CATALOG.md](../docs/architecture/COREOPS_MODULE_CATALOG.md) und [MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md](../docs/architecture/MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine Technologie-/Deployment-Auswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-64 | Logical module architecture | architecture-context | accepted | binding-governance | CO-WP-008 | HM | Accepted foundation architecture direction (17 Module) |
| DEC-S-65 | Module IDs | governance-direction | accepted | binding-governance | CO-WP-008 | Nova | Stabil, nicht wiederverwendbar |
| DEC-S-66 | Module semantics | architecture-context | accepted | guidance | CO-WP-008 | Nova | Logische Verantwortungsgrenzen, **keine** Pflicht-Services/Deployment-Einheiten |
| DEC-S-67 | Policy/Control/Execution separation | security-context | accepted | binding-governance | CO-WP-008 | Nova | Getrennte Autoritätsgrenzen |
| DEC-S-68 | Experience module boundary | security-context | accepted | binding-governance | CO-WP-008 | Nova | Darf privilegierte Ausführung **nicht** direkt auslösen |
| DEC-S-69 | Adapter boundary | security-context | accepted | binding-governance | CO-WP-008 | Nova | Adapter besitzen/umgehen globale Governance **nicht** |
| DEC-S-70 | Agent module | architecture-context | accepted | binding-governance | CO-WP-008 | Nova | Optional; agentless möglich |
| DEC-S-71 | Offline transfer boundary | security-context | accepted | binding-governance | CO-WP-008 | Nova | Getrennte Intake-/Approval-Grenze; keine direkte Ausführung |
| DEC-S-72 | Authoritative data ownership | governance-direction | accepted | binding-governance | CO-WP-008 | Nova | Muss explizit sein (ein Owner pro Konzept) |
| DEC-S-73 | Communication technology | architecture-context | deferred | non-binding | CO-WP-008 | Nova | Deferred; nur konzeptionelle Muster |
| DEC-S-74 | Deployment topology | architecture-context | deferred | non-binding | CO-WP-008 | Nova | Deferred |
| DEC-S-75 | Implementation architecture | architecture-context | deferred | non-binding | CO-WP-008 | Nova | Deferred (Monolith/Microservices offen) |

## Deferred Decisions

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ |
| DEC-D-01 | Visueller Workflow-Editor | deferred | DF | §33 | Nova | post-Foundation | ja (ADR-0019) |
| DEC-D-02 | Externe PKI/TPM/HW-Attestation/SPIFFE | deferred | DF | §31 | Nova | post-Foundation | ja (ADR-0013) |
| DEC-D-03 | Externe Secret-Backends | deferred | DF | §40 | Nova | post-Foundation | ja |
| DEC-D-04 | Tracing in Observability | deferred | DF | §24 | Nova | post-Foundation | ja (ADR-0028) |
| DEC-D-05 | MSP-/Rechenzentrums-Fokus | deferred | DF | §5 | HM | post-Foundation | nein |
| DEC-D-06 | Bidirektionaler NetBox-Sync (über read-only hinaus) | deferred | DF | §19 | Nova | post-Foundation | ja |

## Rejected or Non-Goal Decisions

| Decision ID | Topic | Status | Class | Source | Owner | Notes |
| ----------- | ----- | ------ | ----- | ------ | ----- | ----- |
| DEC-N-01 | Vollständiges CI-Build-System | non-goal | NG | §7 | HM | Integration statt Nachbau |
| DEC-N-02 | Vollständiges SIEM | non-goal | NG | §7 | HM | — |
| DEC-N-03 | Vollständiges ITSM/Helpdesk | non-goal | NG | §7 | HM | — |
| DEC-N-04 | EDR / Antivirus | non-goal | NG | §7 | HM | — |
| DEC-N-05 | Allgemeine Remote-Desktop-Plattform | non-goal | NG | §7 | HM | — |
| DEC-N-06 | Vollständige PKI als Ersatz | non-goal | NG | §7 | HM | — |
| DEC-N-07 | Vollständiger Low-Code-App-Builder | non-goal | NG | §7 | HM | — |
| DEC-N-08 | Universelle Cloud-Management-Suite | non-goal | NG | §7 | HM | — |
| DEC-N-09 | Autonomer KI-Administrator | non-goal | NG | §7,§45 | HM | — |
| DEC-N-10 | Vollständiger Ersatz jedes Herstellerinterfaces | non-goal | NG | §7 | HM | — |
| DEC-N-11 | Vollständige ERP-/Lizenz-/Vertragsverwaltung | non-goal | NG | §7 | HM | — |
| DEC-N-12 | Enforce-Auto-Reconciliation im MVP | non-goal | NG | §20 | HM | später möglich |

## Zusammenfassung

- Accepted Product Decisions: 8
- Binding Governance Decisions: 8
- Foundation Decisions (Sektion „Open Foundation Decisions"): 21 — davon offen 9, `proposed`/`proposed-binding-governance` 10, `clarified` 1, `verified` 1
- ADR Candidates: 32 (30 Concept + 2 Foundation)
- Deferred Decisions: 6
- Non-Goals: 12

**CO-WP-003-Auflösungen (vor Human-Maintainer-Commit unverbindlich):** DEC-O-02 `proposed`, DEC-O-03 `clarified`, DEC-O-04 `proposed-binding-governance`, DEC-O-10 `proposed`, DEC-O-11 `proposed`, DEC-O-16 `verified`.

**CO-WP-004-Auflösungen (vor Human-Maintainer-Commit unverbindlich):** DEC-O-12 Herstellersupport-Grenze `proposed` (CCR-12), DEC-O-17 Capability Matrix `proposed`, DEC-O-18 Observe Support Boundary `proposed`, DEC-O-19 drei Statusdimensionen `proposed-binding-governance`, DEC-O-20 Support-Evidence `proposed-binding-governance`, DEC-O-21 Herstellerlisten ≠ Support `proposed-binding-governance`.

**CO-WP-004A-Registrierungen:** 11 Produkt-/Strategie-/Governance-Entscheidungen (DEC-S-01…11): Souveränität und BSI-orientierte Entwicklung als `accepted-product-direction`; verpflichtende externe Managementprodukte `prohibited`; Zertifizierung `not-claimed`, VS-NfD `not-current-target`; Lessons Learned `binding-governance-direction`, NDF-Feedback `controlled-candidate-process`; ITIL und PRINCE2 `foundation-candidate`; technische Basisabhängigkeiten `open`.

**CO-WP-004B-Registrierungen:** 4 Governance-Entscheidungen (DEC-S-12…15): Lessons-Learned-Prozess, NDF-Feedback-Prozess, kein automatischer NDF-Rückfluss und Human-Maintainer-Gate für NDF-Transfer — alle `proposed-binding-governance`.

**CO-WP-004C-Registrierungen:** 7 Public-Sector-Readiness-Entscheidungen (DEC-S-16…22): interne Profile Standard/Hardened/Government und BSI-orientierte Positionierung sowie Public-Sector-Readiness-Baseline als `accepted-product-direction` (mit expliziten Claim Boundaries); detailliertes BSI-Requirement-Mapping `deferred`; Government-Profil-Zertifizierung `not-claimed`; Cloud C5/C3A `open` (conditional/deployment-dependent); VS-NfD `not-claimed` (not assessed). Keine technische Architektur ausgewählt, keine ADR erzeugt.

**CO-WP-004D-Registrierungen:** 8 Framework-Tailoring-Entscheidungen (DEC-S-23…30): ausgewählte ITIL-Konzepte `adopted-with-tailoring`; ITIL 4 und Version 5 mit expliziten Versionsgrenzen; vollständige ITIL-Implementierung `rejected`; PRINCE2 Version 7 `optional-profile`; vollständige PRINCE2-Implementierung `rejected`; NDF-Primat `binding-governance-direction`; externe Zertifizierung/Endorsement `not-claimed`; Framework-Tool-Abhängigkeit `not-selected`. Löst DEC-S-10/DEC-S-11 (ITIL/PRINCE2 foundation-candidate) auf. Keine ADR, keine Tool-Auswahl.

**CO-WP-004E-Registrierungen:** 7 Capability-Governance-Alignment-Entscheidungen (DEC-S-31…37): mehrdimensionaler Capability-Status `accepted-product-direction`; Roadmap/Implementation/Support/Evidence als getrennte Dimensionen und Evidence-Capability-Trennung `binding-governance-direction`; PSR-Mapping = Readiness-Relationship (nicht Compliance); Responsibility-Mapping erforderlich; detailliertes BSI-Control-Mapping `deferred`; Capability-Zertifizierung/Compliance `not-claimed`. Erweitert DEC-O-19; keine ADR, keine Technologieauswahl.

**CO-WP-005-Registrierungen:** 6 Language-/Repository-Governance-Entscheidungen (DEC-S-38…43), mit **getrennten Dimensionen** (Decision Class · Lifecycle Status · Binding Level), alle `governance-direction` / `accepted` / `binding-governance`: kanonische Sprache (Englisch maschinenbezogen, DE/EN Produkt); Übersetzungsparität evidenzbasiert; Public Neutrality/Disclosure; Human-Maintainer-Repository-Gates; Source-of-Truth-Präzedenz; UTF-8/Zeilenenden. Keine kombinierten Pseudostatuswerte; keine ADR; keine automatisierte Durchsetzung.

**CO-WP-006-Registrierungen:** 10 System-Context-/Architektur-Taxonomie-Entscheidungen (DEC-S-44…53), **getrennte Dimensionen**: System Context und Plane Taxonomy `accepted`/`binding-governance`; Plane-Semantik = logische Bereiche (kein Deployment-Zwang); Managed Resources außerhalb Produktgrenze; External Services optional; Agent Plane optional (agentless möglich); Offline-Core als Produktrichtung; Control Authority erfordert explizite Autorisierung; Detailtechnologie und Threat Model `deferred`. Keine Technologieauswahl; keine ADR.

**CO-WP-007-Registrierungen:** 10 Threat-Model-/Security-Baseline-Entscheidungen (DEC-S-54…63), **getrennte Dimensionen**: Foundation Threat Model und Threat Scenario Register `accepted`/`binding-governance`; Threat-IDs stabil; Ratings qualitativ/evidence-bounded; Security-Invarianten bindende Designanforderungen (keine implementierten Kontrollen); Mitigation-Implementierung `not-claimed`; Mitigation-Validierung, Penetration Testing, Detail-Security-Architektur und Technologie-/Kryptoauswahl `deferred`. Keine Sicherheitskontrolle implementiert; keine ADR.

**CO-WP-008-Registrierungen:** 12 Logical-Module-Architecture-Entscheidungen (DEC-S-64…75), **getrennte Dimensionen**: logische Modularchitektur und Module-Semantik `accepted`; Module-IDs stabil; Policy/Control/Execution-Trennung, Experience-, Adapter-, Offline-Grenzen als Security-Kontext `binding-governance`; Agent-Modul optional; autoritative Datenownership explizit; Communication-Technologie, Deployment-Topologie und Implementation-Architektur `deferred`. Keine Technologie-/Deployment-Auswahl; keine ADR.

**Bestätigung:** Keine technische Architektur-, Technologie- oder Implementierungsentscheidung trägt den Status `accepted`. Keine Integration ist `supported`. Keine Zertifizierung/VS-Eignung behauptet. Keine ADR ist Accepted. Es wurde keine ADR-Datei erzeugt.
