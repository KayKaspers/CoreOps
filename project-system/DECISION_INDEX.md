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

## Human Identity, RBAC and Break-Glass Decisions (CO-WP-009)

> Registriert über [HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](../docs/security/HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../docs/security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md) und [BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](../docs/security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine Identity-/Auth-/Session-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-76 | Human identity and account | security-context | accepted | binding-governance | CO-WP-009 | Nova | Getrennte Konzepte (person/identity/account/principal) |
| DEC-S-77 | Repository vs runtime authority | security-context | accepted | binding-governance | CO-WP-009 | HM | Getrennt; Human Maintainer ≠ Runtime-Autorität |
| DEC-S-78 | Workspace boundary | security-context | accepted | binding-governance | CO-WP-009 | Nova | Administrative/Ressourcen-Scope-Grenze, **nicht** automatisch Security-Tenant |
| DEC-S-79 | RBAC direction | security-context | accepted | binding-governance | CO-WP-009 | Nova | deny-by-default, least-privilege, scope-bound |
| DEC-S-80 | Workspace membership | security-context | accepted | binding-governance | CO-WP-009 | Nova | Erzeugt **keine** globale Autorität |
| DEC-S-81 | Role assignment | governance-direction | accepted | binding-governance | CO-WP-009 | Nova | Explizit, auditierbar, widerrufbar |
| DEC-S-82 | Sensitive operations | security-context | accepted | binding-governance | CO-WP-009 | Nova | Können Reauthentication und Approval verlangen |
| DEC-S-83 | Delegation | governance-direction | accepted | binding-governance | CO-WP-009 | Nova | Explizit, scope-bound, non-transitive by default |
| DEC-S-84 | Break glass | security-context | accepted | binding-governance | CO-WP-009 | HM | Temporär, benannt, reason-/scope-bound, auditiert |
| DEC-S-85 | Break-glass permissions | security-context | accepted | binding-governance | CO-WP-009 | Nova | Müssen ablaufen oder widerrufen werden |
| DEC-S-86 | Offline emergency access | security-context | accepted | binding-governance | CO-WP-009 | Nova | Governed Design-Anforderung; Mechanismus deferred |
| DEC-S-87 | Identity provider / authentication technology | architecture-context | deferred | non-binding | CO-WP-009 | Nova | Deferred |
| DEC-S-88 | Session technology | architecture-context | deferred | non-binding | CO-WP-009 | Nova | Deferred |
| DEC-S-89 | Tenant-isolation implementation | architecture-context | deferred | non-binding | CO-WP-009 | Nova | Deferred |

## Machine Identity and Credential Lifecycle Decisions (CO-WP-010)

> Registriert über [MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md](../docs/security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md](../docs/security/MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md) und [OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md](../docs/security/OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine PKI-/Krypto-/Protokoll-/Secret-Store-Auswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-90 | Human vs machine identity | security-context | accepted | binding-governance | CO-WP-010 | Nova | Getrennt; keine anonyme Human-Ersetzung |
| DEC-S-91 | Machine identity vs credential | security-context | accepted | binding-governance | CO-WP-010 | Nova | Getrennte Konzepte |
| DEC-S-92 | Discovery vs enrollment | security-context | accepted | binding-governance | CO-WP-010 | Nova | Discovery erzeugt keine Identität |
| DEC-S-93 | Registration vs trust | security-context | accepted | binding-governance | CO-WP-010 | Nova | Registration allein ≠ privilegiertes Vertrauen |
| DEC-S-94 | Enrollment governance | security-context | accepted | binding-governance | CO-WP-010 | Nova | Explizit, owner-/scope-bound, auditierbar; keine unbeschränkte Autorität |
| DEC-S-95 | Machine principal scope | security-context | accepted | binding-governance | CO-WP-010 | Nova | Scope-bound; keine geerbte Human-/Owner-Autorität |
| DEC-S-96 | Agent/adapter identity boundary | security-context | accepted | binding-governance | CO-WP-010 | Nova | Umgehen Policy-Grenzen **nicht** |
| DEC-S-97 | Bootstrap material | security-context | accepted | binding-governance | CO-WP-010 | Nova | Keine permanente Identität; scope-/zeitgebunden |
| DEC-S-98 | Credential governance vs raw secret | security-context | accepted | binding-governance | CO-WP-010 | Nova | Governance ≠ Rohsecret-Ownership |
| DEC-S-99 | Rotation/renewal scope | security-context | accepted | binding-governance | CO-WP-010 | Nova | Keine stille Scope-Erweiterung |
| DEC-S-100 | Compromise handling | security-context | accepted | binding-governance | CO-WP-010 | Nova | Containment + explizite Re-Enrollment-Entscheidung; nicht automatisch |
| DEC-S-101 | Decommissioned ID reuse | governance-direction | accepted | binding-governance | CO-WP-010 | Nova | Keine stille Wiederverwendung |
| DEC-S-102 | Offline enrollment | security-context | accepted | binding-governance | CO-WP-010 | Nova | Provenance, Integrität, Approval erforderlich; Mechanismus deferred |
| DEC-S-103 | PKI/credential format/crypto/trust anchors | architecture-context | deferred | non-binding | CO-WP-010 | Nova | Deferred |
| DEC-S-104 | Enrollment protocol | architecture-context | deferred | non-binding | CO-WP-010 | Nova | Deferred |
| DEC-S-105 | Raw secret storage | architecture-context | deferred | non-binding | CO-WP-010 | Nova | Deferred (ob CoreOps Rohsecrets speichert) |

## Source of Truth and Field Provenance Decisions (CO-WP-011)

> Registriert über [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](../docs/architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](../docs/architecture/FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) und [OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md](../docs/security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine Storage-/Merge-/Sync-/Krypto-Auswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-106 | Source of truth vs system of record | architecture-context | accepted | binding-governance | CO-WP-011 | Nova | Getrennte Konzepte |
| DEC-S-107 | Authoritative field ownership | governance-direction | accepted | binding-governance | CO-WP-011 | Nova | Explizit und modulgebunden (ein Owner pro Feldkonzept) |
| DEC-S-108 | Desired/observed/effective/derived state | architecture-context | accepted | binding-governance | CO-WP-011 | Nova | Getrennte Konzepte |
| DEC-S-109 | Derived and cached data | security-context | accepted | binding-governance | CO-WP-011 | Nova | Nicht autoritativ per Default |
| DEC-S-110 | Imported data authority | security-context | accepted | binding-governance | CO-WP-011 | Nova | Erbt Autorität **nicht** automatisch |
| DEC-S-111 | Timestamp conflict rule | security-context | accepted | binding-governance | CO-WP-011 | Nova | Neuester Timestamp gewinnt nicht automatisch |
| DEC-S-112 | Manual overrides | security-context | accepted | binding-governance | CO-WP-011 | Nova | Explizit, feld-/reason-bound, auditierbar |
| DEC-S-113 | Override provenance | security-context | accepted | binding-governance | CO-WP-011 | Nova | Ursprüngliche Provenance bleibt erhalten |
| DEC-S-114 | Conflict visibility | security-context | accepted | binding-governance | CO-WP-011 | Nova | Konflikte bleiben sichtbar bis explizit aufgelöst |
| DEC-S-115 | Revoked/invalidated sources | security-context | accepted | binding-governance | CO-WP-011 | Nova | Bleiben **nicht** autoritativ |
| DEC-S-116 | Offline reconciliation | security-context | accepted | binding-governance | CO-WP-011 | Nova | Erfordert Provenance/Integrität/Authority/Conflict-Review; fail-closed |
| DEC-S-117 | Field provenance persistence | governance-direction | accepted | binding-governance | CO-WP-011 | Nova | Bleibt durch Transformationen erhalten |
| DEC-S-118 | Audit/evidence history | security-context | accepted | binding-governance | CO-WP-011 | Nova | Nicht durch normale Reconciliation umgeschrieben |
| DEC-S-119 | Storage/merge/synchronisation technology | architecture-context | deferred | non-binding | CO-WP-011 | Nova | Deferred |
| DEC-S-120 | Cryptographic provenance | architecture-context | deferred | non-binding | CO-WP-011 | Nova | Deferred |
| DEC-S-121 | Universal field schema | architecture-context | deferred | non-binding | CO-WP-011 | Nova | Deferred |

## State, Drift and Safe-Remediation Decisions (CO-WP-012)

> Registriert über [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](../docs/architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](../docs/architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md) und [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../docs/security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md). **Getrennte Dimensionen**; keine kombinierten Pseudostatuswerte; keine Engine-/Scheduler-/Queue-/Remediation-Auswahl; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-122 | State concepts | architecture-context | accepted | binding-governance | CO-WP-012 | Nova | Desired/observed/reported/effective/last-known getrennt |
| DEC-S-123 | Effective state under conflict | security-context | accepted | binding-governance | CO-WP-012 | Nova | Bleibt indeterminate/conflicted bei unklarer Autorität |
| DEC-S-124 | No observation semantics | governance-direction | accepted | binding-governance | CO-WP-012 | Nova | Keine Beobachtung ≠ keine Drift |
| DEC-S-125 | No-drift semantics | governance-direction | accepted | binding-governance | CO-WP-012 | Nova | Kein erkannter Drift ≠ Compliance |
| DEC-S-126 | Drift detection authority | security-context | accepted | binding-governance | CO-WP-012 | Nova | Drift-Erkennung gewährt **keine** Write Authority |
| DEC-S-127 | Remediation responsibility separation | security-context | accepted | binding-governance | CO-WP-012 | Nova | Detection/Recommendation/Plan/Approval/Execution/Verification getrennt |
| DEC-S-128 | Privileged remediation authority | security-context | accepted | binding-governance | CO-WP-012 | HM | Explizite Autorität + Approval erforderlich |
| DEC-S-129 | Divergence exceptions | security-context | accepted | binding-governance | CO-WP-012 | Nova | Explizit, scope-bound, zurechenbar, reviewbar |
| DEC-S-130 | Execution semantics | security-context | accepted | binding-governance | CO-WP-012 | Nova | Executed ≠ successful |
| DEC-S-131 | Success semantics | security-context | accepted | binding-governance | CO-WP-012 | Nova | Successful ≠ verified convergence |
| DEC-S-132 | Convergence semantics | governance-direction | accepted | binding-governance | CO-WP-012 | Nova | Verified convergence ≠ compliance |
| DEC-S-133 | Partial failure | governance-direction | accepted | binding-governance | CO-WP-012 | Nova | Bleibt explizit repräsentiert |
| DEC-S-134 | Automatic remediation | security-context | deferred | non-binding | CO-WP-012 | HM | Deferred; nicht ausgewählt/implementiert |
| DEC-S-135 | State/drift/remediation engine technology | architecture-context | deferred | non-binding | CO-WP-012 | Nova | Deferred |
| DEC-S-136 | Retry/queue/scheduling technology | architecture-context | deferred | non-binding | CO-WP-012 | Nova | Deferred |

## Policy, Approval and Execution Authorization Decisions (CO-WP-013)

> Registriert über [POLICY_DECISION_AND_EVALUATION_MODEL.md](../docs/security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](../docs/security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) und [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](../docs/security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Policy-/Approval-/Execution-Engine, kein Autorisierungsartefakt/Token, kein Replay-Mechanismus ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-137 | Policy evaluation / approval / execution authorization | security-context | accepted | binding-governance | CO-WP-013 | Nova | Getrennte Verantwortlichkeiten |
| DEC-S-138 | Policy permit semantics | security-context | accepted | binding-governance | CO-WP-013 | Nova | permit ≠ approval ≠ execution authorization |
| DEC-S-139 | Default privileged outcome | security-context | accepted | binding-governance | CO-WP-013 | Nova | deny/indeterminate außer explizit autorisiert; kein Default-Permit |
| DEC-S-140 | Policy conflicts | security-context | accepted | binding-governance | CO-WP-013 | Nova | Bleiben sichtbar bis explizit aufgelöst |
| DEC-S-141 | Approval properties | security-context | accepted | binding-governance | CO-WP-013 | Nova | Explizit, zurechenbar, scope-bound, widerrufbar |
| DEC-S-142 | Machine self-approval | security-context | accepted | binding-governance | CO-WP-013 | HM | Machine Principals können sich **nicht** selbst genehmigen |
| DEC-S-143 | Approver scope | security-context | accepted | binding-governance | CO-WP-013 | Nova | Rollenname ≠ unbegrenzter Approval-Scope |
| DEC-S-144 | Execution authorization binding | security-context | accepted | binding-governance | CO-WP-013 | Nova | action-/target-/scope-/plan-/time-bound |
| DEC-S-145 | Plan/target change | security-context | accepted | binding-governance | CO-WP-013 | Nova | Materielle Änderung → Re-Evaluation/ggf. neue Approval |
| DEC-S-146 | Expired/revoked/consumed authorization | security-context | accepted | binding-governance | CO-WP-013 | Nova | Nicht wiederverwendbar |
| DEC-S-147 | Break-glass relationship | security-context | accepted | binding-governance | CO-WP-013 | HM | Außergewöhnlich, nicht permanent; referenziert bestehende Break-Glass-Policy (kein Parallelmodell) |
| DEC-S-148 | Offline authorization | security-context | accepted | binding-governance | CO-WP-013 | Nova | Target-/Scope-/Plan-Binding, Provenance, Integrität, explizite Aktivierung |
| DEC-S-149 | Execution result semantics | governance-direction | accepted | binding-governance | CO-WP-013 | Nova | Execution result ≠ success |
| DEC-S-150 | Success/closure semantics | governance-direction | accepted | binding-governance | CO-WP-013 | Nova | Success ≠ verification; closed ≠ success |
| DEC-S-151 | Policy/approval/authorization engine technology | architecture-context | deferred | non-binding | CO-WP-013 | Nova | Deferred |
| DEC-S-152 | Authorization artifact and replay-protection mechanism | architecture-context | deferred | non-binding | CO-WP-013 | Nova | Deferred |

## Integration Contract Decisions (CO-WP-014)

> Registriert über [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../docs/architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md), [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](../docs/architecture/INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md) und [INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md](../docs/security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Protokoll-/Schema-/Transport-/SDK-/Adapter-/Replay-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-153 | CoreOps Integration Contract | architecture-context | accepted | binding-governance | CO-WP-014 | Nova | Foundation Contract Version 0.1 |
| DEC-S-154 | Contract version scope | governance-direction | accepted | binding-governance | CO-WP-014 | Nova | Contract-Version ≠ CoreOps-Produkt-/Release-Version |
| DEC-S-155 | Capability dimensions | security-context | accepted | binding-governance | CO-WP-014 | Nova | advertised/detected/permitted/implemented/supported/validated getrennt |
| DEC-S-156 | Request acceptance semantics | security-context | accepted | binding-governance | CO-WP-014 | Nova | Acceptance ≠ Authorization ≠ Execution |
| DEC-S-157 | Operation completion semantics | governance-direction | accepted | binding-governance | CO-WP-014 | Nova | Completion ≠ success |
| DEC-S-158 | Operation success semantics | governance-direction | accepted | binding-governance | CO-WP-014 | Nova | Success ≠ verification |
| DEC-S-159 | Unknown outcome | security-context | accepted | binding-governance | CO-WP-014 | Nova | Erfordert Reconciliation; blockiert automatischen Retry |
| DEC-S-160 | Read-only integration | security-context | accepted | binding-governance | CO-WP-014 | Nova | Erhält **nicht** still Write Authority |
| DEC-S-161 | Write/execution authority | security-context | accepted | binding-governance | CO-WP-014 | Nova | Erfordern explizite Policy + anwendbare Authorization |
| DEC-S-162 | Integration result authority | security-context | accepted | binding-governance | CO-WP-014 | Nova | Erben **nicht** automatisch autoritativen State |
| DEC-S-163 | Adapter/agent scope | security-context | accepted | binding-governance | CO-WP-014 | Nova | Erweitern Target/Action/Scope **nicht** |
| DEC-S-164 | Offline integration | security-context | accepted | binding-governance | CO-WP-014 | Nova | Provenance, Integrität, Target-Binding, explizite Aktivierung |
| DEC-S-165 | Contract extensions | security-context | accepted | binding-governance | CO-WP-014 | Nova | Überschreiben **keine** Core-Sicherheitsinvariante |
| DEC-S-166 | Protocol/schema/transport/SDK technology | architecture-context | deferred | non-binding | CO-WP-014 | Nova | Deferred |
| DEC-S-167 | Replay/deduplication/idempotency mechanism | architecture-context | deferred | non-binding | CO-WP-014 | Nova | Deferred |
| DEC-S-168 | Recovery automation | architecture-context | deferred | non-binding | CO-WP-014 | Nova | Deferred |

## Domain Pack Governance Decisions (CO-WP-015)

> Registriert über [DOMAIN_PACK_GOVERNANCE_MODEL.md](../docs/architecture/DOMAIN_PACK_GOVERNANCE_MODEL.md), [DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md](../docs/architecture/DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md) und [DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md](../docs/security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Packaging-/Marketplace-/Plugin-/Update-/Signaturtechnologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-169 | Domain Pack definition | architecture-context | accepted | binding-governance | CO-WP-015 | Nova | Versionierte Governance-/Produktgrenze; ≠ adapter/plugin/deployment unit/certification |
| DEC-S-170 | Pack identity | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Stabile IDs, nicht wiederverwendbar |
| DEC-S-171 | Pack status dimensions | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Lifecycle/Maintenance/Support/Implementation/Validation/Evidence/Security-Review/Compatibility getrennt |
| DEC-S-172 | Support level semantics | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Support Level ≠ SLA |
| DEC-S-173 | Project-maintained semantics | security-context | accepted | binding-governance | CO-WP-015 | Nova | project-maintained ≠ validated |
| DEC-S-174 | SUP-3 scope | security-context | accepted | binding-governance | CO-WP-015 | Nova | Version-/Profil-/Limitation-/Evidence-bound |
| DEC-S-175 | Compatibility claims | security-context | accepted | binding-governance | CO-WP-015 | Nova | Version-/Target-/Profil-/Evidence-bound |
| DEC-S-176 | Expected compatibility | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | expected ≠ validated compatibility |
| DEC-S-177 | Community/external packs | security-context | accepted | binding-governance | CO-WP-015 | Nova | Nicht automatisch trusted/supported |
| DEC-S-178 | Vendor reference | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Kein Endorsement/Zertifizierung/Herstellersupport |
| DEC-S-179 | Pack activation authority | security-context | accepted | binding-governance | CO-WP-015 | Nova | Gewährt **keine** Runtime-Autorität automatisch |
| DEC-S-180 | Pack dependencies | security-context | accepted | binding-governance | CO-WP-015 | Nova | Werden **nicht** still zur Core-Pflicht |
| DEC-S-181 | Offline pack use | security-context | accepted | binding-governance | CO-WP-015 | Nova | Provenance, Integrität, Target-Binding, explizite Aktivierung |
| DEC-S-182 | Deprecation | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Erfordert Migration und Maintenance-Boundary |
| DEC-S-183 | Retired pack IDs | governance-direction | accepted | binding-governance | CO-WP-015 | Nova | Nicht wiederverwendbar; historische Evidenz erhalten |
| DEC-S-184 | Packaging/marketplace/plugin/update technology | architecture-context | deferred | non-binding | CO-WP-015 | Nova | Deferred |
| DEC-S-185 | Pack verification/signing/trust-anchor technology | architecture-context | deferred | non-binding | CO-WP-015 | Nova | Deferred |

## Data Ownership, Persistence and Migration Decisions (CO-WP-016)

> Registriert über [DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md](../docs/architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md](../docs/architecture/SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md) und [DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md](../docs/security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Storage-/DB-/ORM-/Schema-/Migration-/Backup-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-186 | Ownership responsibilities | security-context | accepted | binding-governance | CO-WP-016 | Nova | Owner/Steward/Storage/Write/Migration/Retention/Recovery getrennt |
| DEC-S-187 | Storage responsibility | security-context | accepted | binding-governance | CO-WP-016 | Nova | ≠ autoritative Ownership |
| DEC-S-188 | Migration authority | security-context | accepted | binding-governance | CO-WP-016 | Nova | ≠ uneingeschränkte Write-Autorität |
| DEC-S-189 | Schema identity | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Stabil; ≠ storage location/file name |
| DEC-S-190 | Version dimensions | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Schema/Data/Producer/Consumer/Migration/Contract/Product getrennt |
| DEC-S-191 | Compatibility dimensions | security-context | accepted | binding-governance | CO-WP-016 | Nova | Read/Write/Round-Trip getrennt |
| DEC-S-192 | Unknown/conflicted compatibility | security-context | accepted | binding-governance | CO-WP-016 | Nova | Blockiert automatische destruktive Migration |
| DEC-S-193 | Migration responsibilities | security-context | accepted | binding-governance | CO-WP-016 | Nova | Plan/Approval/Execution/Validation/Recovery getrennt |
| DEC-S-194 | Executed migration semantics | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Executed ≠ validierte Integrität |
| DEC-S-195 | Backup semantics | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Backup exists ≠ restorable |
| DEC-S-196 | Restore semantics | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Restore completed ≠ service recovery |
| DEC-S-197 | Partial migration | governance-direction | accepted | binding-governance | CO-WP-016 | Nova | Bleibt explizit |
| DEC-S-198 | Mixed-version operation | security-context | accepted | binding-governance | CO-WP-016 | Nova | Explizite Compatibility-/Dauergrenzen |
| DEC-S-199 | Destructive migration | security-context | accepted | binding-governance | CO-WP-016 | HM | Gebundene Autorität + anwendbare Approval |
| DEC-S-200 | Migration and authority | security-context | accepted | binding-governance | CO-WP-016 | Nova | Reaktiviert keine widerrufene Autorität/konsumierte Authorization |
| DEC-S-201 | Audit/evidence provenance | security-context | accepted | binding-governance | CO-WP-016 | Nova | Bleibt durch Migration erhalten |
| DEC-S-202 | Offline migration | security-context | accepted | binding-governance | CO-WP-016 | Nova | Provenance, Integrität, Target-Binding, explizite Aktivierung |
| DEC-S-203 | Storage/DB/schema/migration/backup/recovery technology | architecture-context | deferred | non-binding | CO-WP-016 | Nova | Deferred |

## API Governance Decisions (CO-WP-017)

> Registriert über [API_GOVERNANCE_AND_OPERATION_MODEL.md](../docs/architecture/API_GOVERNANCE_AND_OPERATION_MODEL.md), [API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md](../docs/architecture/API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md) und [API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md](../docs/security/API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Transport-/API-Style-/Schema-/Statuscode-/Idempotency-/Replay-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-204 | API identity | governance-direction | accepted | binding-governance | CO-WP-017 | Nova | Stabil; ≠ route/URL/transport binding |
| DEC-S-205 | API version dimensions | governance-direction | accepted | binding-governance | CO-WP-017 | Nova | Surface/Operation/Request/Response/Error/Schema/Contract/Pack/Product getrennt |
| DEC-S-206 | API availability | security-context | accepted | binding-governance | CO-WP-017 | Nova | Impliziert **keine** Authorization |
| DEC-S-207 | Request acceptance | security-context | accepted | binding-governance | CO-WP-017 | Nova | Impliziert **keine** Execution |
| DEC-S-208 | Successful response | governance-direction | accepted | binding-governance | CO-WP-017 | Nova | Impliziert **kein** verifiziertes Ergebnis |
| DEC-S-209 | Error response | security-context | accepted | binding-governance | CO-WP-017 | Nova | Beweist **keine** Abwesenheit von Nebenwirkungen |
| DEC-S-210 | Read-only API | security-context | accepted | binding-governance | CO-WP-017 | Nova | Erhält **nicht** still Write Authority |
| DEC-S-211 | Compatibility dimensions | security-context | accepted | binding-governance | CO-WP-017 | Nova | Request/Response/Error/Behavioural getrennt |
| DEC-S-212 | Unknown compatibility | security-context | accepted | binding-governance | CO-WP-017 | Nova | Impliziert **keine** Compatibility |
| DEC-S-213 | Idempotency context | security-context | accepted | binding-governance | CO-WP-017 | Nova | Verlängert/ersetzt **keine** Authorization |
| DEC-S-214 | Unknown outcome | security-context | accepted | binding-governance | CO-WP-017 | Nova | Blockiert automatischen Retry; erfordert Reconciliation |
| DEC-S-215 | Duplicate/replay handling | security-context | accepted | binding-governance | CO-WP-017 | Nova | Erhält Request-/Attempt-/Result-Historie |
| DEC-S-216 | Bulk operations | security-context | accepted | binding-governance | CO-WP-017 | Nova | Erhalten Per-Target-Autorität und Partial Results |
| DEC-S-217 | Pagination/continuation | security-context | accepted | binding-governance | CO-WP-017 | Nova | Impliziert **kein** Snapshot/keine aktuelle Authorization automatisch |
| DEC-S-218 | Deprecated/retired API IDs | governance-direction | accepted | binding-governance | CO-WP-017 | Nova | Nicht still wiederverwendbar |
| DEC-S-219 | Transport/API-style/schema/status-code/pagination/idempotency/replay technology | architecture-context | deferred | non-binding | CO-WP-017 | Nova | Deferred |

## Event, Audit and Evidence Decisions (CO-WP-018)

> Registriert über [EVENT_AND_AUDIT_CORRELATION_MODEL.md](../docs/architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md), [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../docs/architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) und [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](../docs/security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Event-Bus-/Logging-/SIEM-/Storage-/Schema-/Ordering-/Hash-/Signatur-/WORM-/Redaction-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-220 | Event identity | governance-direction | accepted | binding-governance | CO-WP-018 | Nova | Getrennt von Correlation/Request/Operation/Attempt |
| DEC-S-221 | Event concept separation | architecture-context | accepted | binding-governance | CO-WP-018 | Nova | Event/Audit Event/Notification/Command/Evidence getrennt |
| DEC-S-222 | Event time dimensions | architecture-context | accepted | binding-governance | CO-WP-018 | Nova | Occurrence/Observation/Recording/Ingestion getrennt |
| DEC-S-223 | Timestamp ordering | security-context | accepted | binding-governance | CO-WP-018 | Nova | Liefert **keine** autoritative globale Ordnung |
| DEC-S-224 | Correlation vs causation | security-context | accepted | binding-governance | CO-WP-018 | Nova | Correlation beweist **keine** Causation |
| DEC-S-225 | Event recorded semantics | governance-direction | accepted | binding-governance | CO-WP-018 | Nova | Recorded ≠ validated/complete |
| DEC-S-226 | Audit record semantics | security-context | accepted | binding-governance | CO-WP-018 | Nova | ≠ validated evidence; nicht inhärent wahr |
| DEC-S-227 | Missing event semantics | security-context | accepted | binding-governance | CO-WP-018 | Nova | Beweist **keine** Nichtausführung |
| DEC-S-228 | Evidence dimensions | security-context | accepted | binding-governance | CO-WP-018 | Nova | Capability/Availability/Freshness/Integrity/Validation/Sufficiency getrennt |
| DEC-S-229 | Evidence sufficiency | security-context | accepted | binding-governance | CO-WP-018 | Nova | Decision-/Scope-/Time-bound |
| DEC-S-230 | Audit completeness | governance-direction | accepted | binding-governance | CO-WP-018 | Nova | Scope-bound; kann `unknown` bleiben |
| DEC-S-231 | Duplicate/replay handling | security-context | accepted | binding-governance | CO-WP-018 | Nova | Erhält Event-/Attempt-Historie |
| DEC-S-232 | Audit administrator authority | security-context | accepted | binding-governance | CO-WP-018 | HM | Keine uneingeschränkte Disclosure Authority |
| DEC-S-233 | Audit/evidence export | security-context | accepted | binding-governance | CO-WP-018 | Nova | Explizite Scope; Provenance erhalten |
| DEC-S-234 | Offline audit/evidence import | security-context | accepted | binding-governance | CO-WP-018 | Nova | Provenance, Integrität, Target-Binding, explizite Governance |
| DEC-S-235 | Audit closure | governance-direction | accepted | binding-governance | CO-WP-018 | Nova | ≠ Completeness/Validation/Sufficiency/Compliance |
| DEC-S-236 | Event-bus/logging/SIEM/storage/schema/ordering/hash/signature/WORM/redaction technology | architecture-context | deferred | non-binding | CO-WP-018 | Nova | Deferred |

## Telemetry and Normalization Decisions (CO-WP-019)

> Registriert über [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](../docs/architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md](../docs/architecture/TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md) und [TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md](../docs/security/TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Telemetry-/Protokoll-/Schema-/Collector-/Storage-/Mapping-/Unit-/Aggregation-/Alerting-/Dashboard-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-237 | Telemetry concept separation | architecture-context | accepted | binding-governance | CO-WP-019 | Nova | Telemetry/Event/Audit Event/Evidence/Notification/Command getrennt |
| DEC-S-238 | Signal identity dimensions | governance-direction | accepted | binding-governance | CO-WP-019 | Nova | Signal/Series/Sample/Resource Identity getrennt |
| DEC-S-239 | Producer vs source | security-context | accepted | binding-governance | CO-WP-019 | Nova | Producer getrennt von Source |
| DEC-S-240 | Telemetry state dimensions | architecture-context | accepted | binding-governance | CO-WP-019 | Nova | Raw/Normalized/Derived/Aggregated getrennt |
| DEC-S-241 | Telemetry state authority | security-context | accepted | binding-governance | CO-WP-019 | Nova | Erbt **keinen** autoritativen State |
| DEC-S-242 | Normalization provenance | security-context | accepted | binding-governance | CO-WP-019 | Nova | Erhält Source-Provenance und Transformation History |
| DEC-S-243 | Canonical field identity | governance-direction | accepted | binding-governance | CO-WP-019 | Nova | Stabil; ≠ display label/source field name |
| DEC-S-244 | Unit/scale/precision | security-context | accepted | binding-governance | CO-WP-019 | Nova | Getrennte Dimensionen |
| DEC-S-245 | Unknown/conflicting unit | security-context | accepted | binding-governance | CO-WP-019 | Nova | Blockiert unsafe automatische Konversion |
| DEC-S-246 | Quality/confidence/freshness/integrity/validation | security-context | accepted | binding-governance | CO-WP-019 | Nova | Getrennte Dimensionen |
| DEC-S-247 | Observation vs recording/ingestion time | security-context | accepted | binding-governance | CO-WP-019 | Nova | Observation Time getrennt |
| DEC-S-248 | Missing telemetry semantics | security-context | accepted | binding-governance | CO-WP-019 | Nova | ≠ zero/inactivity/target failure |
| DEC-S-249 | Sampling visibility | governance-direction | accepted | binding-governance | CO-WP-019 | Nova | Bleibt in Interpretation/Evidence sichtbar |
| DEC-S-250 | Aggregation semantics | governance-direction | accepted | binding-governance | CO-WP-019 | Nova | ≠ vollständige/unabhängig validierte Population |
| DEC-S-251 | Telemetry label disclosure | security-context | accepted | binding-governance | CO-WP-019 | Nova | ≠ sichere Disclosure-Dimension |
| DEC-S-252 | Telemetry-to-event/-evidence | security-context | accepted | binding-governance | CO-WP-019 | Nova | Erfordert explizite Klassifikation und Provenance |
| DEC-S-253 | Offline telemetry import | security-context | accepted | binding-governance | CO-WP-019 | Nova | Provenance, Integrität, Target-Binding, explizite Governance |
| DEC-S-254 | Telemetry protocol/schema/collector/storage/mapping/unit/aggregation/alerting/dashboard technology | architecture-context | deferred | non-binding | CO-WP-019 | Nova | Deferred |

## Topology Graph Decisions (CO-WP-020)

> Registriert über [TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md](../docs/architecture/TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md), [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md](../docs/architecture/TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) und [TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md](../docs/security/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md). **Getrennte Dimensionen** (Decision Class · Lifecycle Status · Binding Level); keine kombinierten Pseudostatuswerte; keine Graph-DB-/Discovery-/Query-/Identity-Resolution-/Conflict-Resolution-/Visualization-/Layout-Technologie ausgewählt; keine ADR.

| Decision ID | Topic | Decision Class | Lifecycle Status | Binding Level | Source | Owner | Notes |
| ----------- | ----- | -------------- | ---------------- | ------------- | ------ | ----- | ----- |
| DEC-S-255 | Topology graph boundary | architecture-context | accepted | binding-governance | CO-WP-020 | Nova | Verwaltete Assertions; ≠ autoritative physische Realität |
| DEC-S-256 | Node identity dimensions | governance-direction | accepted | binding-governance | CO-WP-020 | Nova | Node/Resource Identity/Alias/Display Name getrennt |
| DEC-S-257 | Edge identity | governance-direction | accepted | binding-governance | CO-WP-020 | Nova | Getrennt von Source/Target Node und Relationship Class |
| DEC-S-258 | Assertion origin classes | security-context | accepted | binding-governance | CO-WP-020 | Nova | Discovered/Observed/Declared/Imported/Manual/Derived/Inferred getrennt |
| DEC-S-259 | Relationship assertion semantics | security-context | accepted | binding-governance | CO-WP-020 | Nova | ≠ validierte/aktuell aktive Beziehung |
| DEC-S-260 | Identity match semantics | security-context | accepted | binding-governance | CO-WP-020 | Nova | Same name/address/alias ≠ same canonical node |
| DEC-S-261 | Identity resolution | security-context | accepted | binding-governance | CO-WP-020 | Nova | Erhält Evidence/Scope/Temporal Context |
| DEC-S-262 | Merge and split | security-context | accepted | binding-governance | CO-WP-020 | Nova | Erhalten historische Identitäten/Assertions/Evidence |
| DEC-S-263 | Current view semantics | governance-direction | accepted | binding-governance | CO-WP-020 | Nova | ≠ vollständige Topologie |
| DEC-S-264 | Snapshot semantics | governance-direction | accepted | binding-governance | CO-WP-020 | Nova | ≠ alle Quellen zeitgleich beobachtet/immutable Evidence |
| DEC-S-265 | Evidence dimensions | security-context | accepted | binding-governance | CO-WP-020 | Nova | Source Trust/Authority/Confidence/Validation/Evidence/Completeness getrennt |
| DEC-S-266 | Timestamp precedence | security-context | accepted | binding-governance | CO-WP-020 | Nova | Kein stilles Last-Write-Wins |
| DEC-S-267 | Manual authority | security-context | accepted | binding-governance | CO-WP-020 | HM | Human-attributable, scope-bound, reviewbar |
| DEC-S-268 | Manual override | security-context | accepted | binding-governance | CO-WP-020 | Nova | Löscht **keine** konkurrierenden Observations/Evidence |
| DEC-S-269 | Suppression | security-context | accepted | binding-governance | CO-WP-020 | Nova | ≠ Relationship-Abwesenheit |
| DEC-S-270 | Topology authority boundary | security-context | accepted | binding-governance | CO-WP-020 | Nova | Gewährt keine Netzwerk-/Konfigurations-/Execution-Autorität |
| DEC-S-271 | Unresolved conflict | security-context | accepted | binding-governance | CO-WP-020 | Nova | Blockiert unsafe privilegierte Automatisierung |
| DEC-S-272 | Offline topology import | security-context | accepted | binding-governance | CO-WP-020 | Nova | Provenance, Integrität, Target-Binding, explizite Governance |
| DEC-S-273 | Graph-DB/discovery/query/identity-resolution/conflict-resolution/visualization/layout technology | architecture-context | deferred | non-binding | CO-WP-020 | Nova | Deferred |

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

**CO-WP-009-Registrierungen:** 14 Human-Identity-/RBAC-/Break-Glass-Entscheidungen (DEC-S-76…89), **getrennte Dimensionen**: getrennte Identity-/Account-Konzepte; Repository- ≠ Runtime-Autorität; Workspace ≠ automatischer Security-Tenant; RBAC deny-by-default/least-privilege/scope-bound; Membership ohne globale Autorität; Rollenzuweisung explizit/auditierbar/widerrufbar; sensible Aktionen mit Reauth/Approval; Delegation explizit/scope-bound/non-transitive; Break Glass temporär/benannt/auditiert mit Ablaufpflicht; Offline-Emergency-Access governed (Mechanismus deferred); Identity-Provider-/Auth-/Session-Technologie und Tenant-Isolation `deferred`. Keine Technologieauswahl; keine ADR.

**CO-WP-010-Registrierungen:** 16 Machine-Identity-/Credential-Lifecycle-Entscheidungen (DEC-S-90…105), **getrennte Dimensionen**: Human ≠ Machine Identity; Identity ≠ Credential; Discovery ≠ Enrollment; Registration ≠ Trust; Enrollment explizit/owner-/scope-bound (keine unbeschränkte Autorität); Machine Principals scope-bound; Agent/Adapter-Identität umgeht Policy nicht; Bootstrap ≠ permanente Identität; Credential-Governance ≠ Rohsecret-Ownership; Rotation/Renewal ohne stille Scope-Erweiterung; Compromise mit Containment + expliziter Re-Enrollment-Entscheidung; keine stille Decommission-ID-Wiederverwendung; Offline-Enrollment mit Provenance/Integrität/Approval; PKI/Krypto/Trust-Anchors, Enrollment-Protokoll und Rohsecret-Speicherung `deferred`. Keine Technologieauswahl; keine ADR.

**CO-WP-011-Registrierungen:** 16 Source-of-Truth-/Field-Provenance-Entscheidungen (DEC-S-106…121), **getrennte Dimensionen**: Source of Truth ≠ System of Record; autoritative Feldownership explizit/modulgebunden; Desired/Observed/Effective/Derived getrennt; Derived/Cached nicht autoritativ per Default; importierte Daten erben keine Autorität; neuester Timestamp gewinnt nicht automatisch; Overrides explizit/feld-/reason-bound mit Provenance-Erhalt; Konflikte bleiben sichtbar; widerrufene Quellen nicht autoritativ; Offline-Reconciliation mit Provenance/Integrität/Authority/Conflict-Review (fail-closed); Field-Provenance durch Transformationen erhalten; Audit-/Evidence-History nicht durch Reconciliation umgeschrieben; Storage/Merge/Sync-Technologie, kryptografische Provenance und universelles Feldschema `deferred`. Keine Technologieauswahl; keine ADR.

**CO-WP-012-Registrierungen:** 15 State-/Drift-/Safe-Remediation-Entscheidungen (DEC-S-122…136), **getrennte Dimensionen**: Desired/Observed/Reported/Effective/Last-Known getrennt; Effective State bleibt indeterminate/conflicted bei unklarer Autorität; keine Beobachtung ≠ keine Drift; kein erkannter Drift ≠ Compliance; Drift-Erkennung ohne Write Authority; Detection/Recommendation/Plan/Approval/Execution/Verification getrennt; privilegierte Remediation mit expliziter Autorität + Approval; Exceptions explizit/reviewbar; Executed ≠ successful; Successful ≠ verified convergence; Verified convergence ≠ Compliance; Partial Failure explizit; automatische Remediation, Engine-/Retry-/Queue-/Scheduling-Technologie `deferred`. Keine Technologieauswahl; keine ADR.

**CO-WP-020-Registrierungen:** 19 Topology-Graph-Entscheidungen (DEC-S-255…273), **getrennte Dimensionen**: Topology Graph als verwaltete Assertion-Darstellung (≠ autoritative physische Realität); Node/Resource Identity/Alias/Display Name getrennt; Edge Identity getrennt von Source/Target/Relationship Class; sieben Origin-Klassen getrennt; Relationship Assertion ≠ validierte/aktive Beziehung; same name/address/alias ≠ same canonical node; Identity Resolution erhält Evidence/Scope/Temporal Context; Merge/Split erhalten historische Identitäten/Assertions/Evidence; Current View ≠ vollständige Topologie; Snapshot ≠ zeitgleich/immutable; Source Trust/Authority/Confidence/Validation/Evidence/Completeness getrennt; kein stilles Timestamp-Last-Write-Wins; Manual Authority human-attributable/scope-bound/reviewbar; Manual Override löscht keine konkurrierenden Observations; Suppression ≠ Relationship-Abwesenheit; Topology gewährt keine Netzwerk-/Konfigurations-/Execution-Autorität; unresolved Conflict blockiert unsafe privilegierte Automatisierung; Offline-Import mit Provenance/Integrität/Target-Binding/expliziter Governance; Graph-DB-/Discovery-/Query-/Identity-Resolution-/Conflict-Resolution-/Visualization-/Layout-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-019-Registrierungen:** 18 Telemetry-/Normalization-Entscheidungen (DEC-S-237…254), **getrennte Dimensionen**: Telemetry/Event/Audit Event/Evidence/Notification/Command getrennt; Signal/Series/Sample/Resource Identity getrennt; Producer ≠ Source; Raw/Normalized/Derived/Aggregated getrennt; Telemetry erbt keinen autoritativen State; Normalisierung erhält Source-Provenance/Transformation History; Canonical Field Identity stabil (≠ display label/source field name); Unit/Scale/Precision getrennt; Unknown/Conflicting Unit blockiert unsafe automatische Konversion; Quality/Confidence/Freshness/Integrity/Validation getrennt; Observation Time getrennt von Recording/Ingestion; Missing Telemetry ≠ zero/inactivity/target failure; Sampling bleibt sichtbar; Aggregation ≠ vollständige/unabhängig validierte Population; Telemetry Label ≠ sichere Disclosure-Dimension; Telemetry-to-Event/-Evidence mit expliziter Klassifikation/Provenance; Offline-Import mit Provenance/Integrität/Target-Binding/expliziter Governance; Telemetry-Protokoll-/Schema-/Collector-/Storage-/Mapping-/Unit-/Aggregation-/Alerting-/Dashboard-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-018-Registrierungen:** 17 Event-/Audit-/Evidence-Entscheidungen (DEC-S-220…236), **getrennte Dimensionen**: Event Identity getrennt von Correlation/Request/Operation/Attempt; Event/Audit Event/Notification/Command/Evidence getrennte Konzepte; vier Zeitdimensionen getrennt; Timestamp/Sequence ohne autoritative globale Ordnung; Correlation ≠ Causation; Recorded ≠ validated/complete; Audit Record ≠ validated evidence; Missing Event ≠ Nichtausführung; sechs Evidence-Dimensionen getrennt; Sufficiency decision-/scope-/time-bound; Audit Completeness scope-bound (kann unknown bleiben); Duplicate/Replay erhält Event-/Attempt-Historie; Audit Administrator ohne uneingeschränkte Disclosure Authority; Export mit expliziter Scope + Provenance-Erhalt; Offline-Import mit Provenance/Integrität/Target-Binding/expliziter Governance; Audit Closure ≠ Completeness/Validation/Sufficiency/Compliance; Event-Bus-/Logging-/SIEM-/Storage-/Schema-/Ordering-/Hash-/Signatur-/WORM-/Redaction-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-017-Registrierungen:** 16 API-Governance-Entscheidungen (DEC-S-204…219), **getrennte Dimensionen**: API Identity stabil (≠ route/URL/transport); zwölf Versionsdimensionen getrennt; API Availability ≠ Authorization; Request Acceptance ≠ Execution; Successful Response ≠ verifiziertes Ergebnis; Error Response ≠ Abwesenheit von Nebenwirkungen; Read-only-API erhält nicht still Write Authority; Request/Response/Error/Behavioural-Compatibility getrennt; Unknown Compatibility ≠ Compatibility; Idempotency Context verlängert/ersetzt keine Authorization; Unknown Outcome blockiert automatischen Retry und erfordert Reconciliation; Duplicate/Replay erhält Request-/Attempt-/Result-Historie; Bulk erhält Per-Target-Autorität und Partial Results; Pagination/Continuation impliziert kein Snapshot/keine aktuelle Authorization; deprecated/retired API-IDs nicht still wiederverwendbar; Transport-/API-Style-/Schema-/Statuscode-/Pagination-/Idempotency-/Replay-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-016-Registrierungen:** 18 Data-Ownership-/Persistence-/Migration-Entscheidungen (DEC-S-186…203), **getrennte Dimensionen**: Owner/Steward/Storage/Write/Migration/Retention/Recovery getrennt; Storage ≠ autoritative Ownership; Migration Authority ≠ uneingeschränkte Write-Autorität; Schema Identity stabil (≠ storage location/file name); Schema-/Data-/Producer-/Consumer-/Migration-/Contract-/Product-Version getrennt; Read/Write/Round-Trip-Compatibility getrennt; Unknown/Conflicted Compatibility blockiert automatische destruktive Migration; Plan/Approval/Execution/Validation/Recovery getrennt; Executed ≠ validierte Integrität; Backup exists ≠ restorable; Restore ≠ Service Recovery; Partial Migration explizit; Mixed-Version mit Compatibility-/Dauergrenzen; destruktive Migration mit gebundener Autorität + Approval; Migration reaktiviert keine widerrufene Autorität/konsumierte Authorization; Audit-/Evidence-Provenance bleibt erhalten; Offline-Migration mit Provenance/Integrität/Target-Binding/expliziter Aktivierung; Storage-/DB-/Schema-/Migration-/Backup-/Recovery-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-015-Registrierungen:** 17 Domain-Pack-Governance-Entscheidungen (DEC-S-169…185), **getrennte Dimensionen**: Domain Pack als versionierte Governance-/Produktgrenze (≠ adapter/plugin/deployment unit/certification); stabile, nicht wiederverwendbare Pack-IDs; neun Statusdimensionen getrennt; Support Level ≠ SLA; project-maintained ≠ validated; SUP-3 version-/profil-/limitation-/evidence-bound; Compatibility Claims version-/target-/profil-/evidence-bound; expected ≠ validated compatibility; Community/External Packs nicht automatisch trusted/supported; Vendor-Bezug ≠ Endorsement/Zertifizierung/Herstellersupport; Pack-Aktivierung gewährt keine Runtime-Autorität; Dependencies werden nicht still Core-Pflicht; Offline-Pack-Nutzung mit Provenance/Integrität/Target-Binding/expliziter Aktivierung; Deprecation mit Migration/Maintenance-Boundary; retired Pack-IDs nicht wiederverwendbar (historische Evidenz erhalten); Packaging-/Marketplace-/Plugin-/Update- und Pack-Verification-/Signing-/Trust-Anchor-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-014-Registrierungen:** 16 Integration-Contract-Entscheidungen (DEC-S-153…168), **getrennte Dimensionen**: CoreOps Integration Contract als Foundation Version 0.1; Contract-Version ≠ Produkt-/Release-Version; sechs Capability-Dimensionen (advertised/detected/permitted/implemented/supported/validated) getrennt; Request Acceptance ≠ Authorization ≠ Execution; Completion ≠ success; Success ≠ verification; Unknown Outcome erfordert Reconciliation und blockiert automatischen Retry; Read-only-Integration erhält nicht still Write Authority; Write/Execution erfordern explizite Policy + Authorization; Integrationsergebnisse erben nicht automatisch autoritativen State; Adapter/Agent erweitern Target/Action/Scope nicht; Offline Integration mit Provenance/Integrität/Target-Binding/expliziter Aktivierung; Contract-Extensions überschreiben keine Core-Sicherheitsinvariante; Protokoll-/Schema-/Transport-/SDK-, Replay-/Deduplication-/Idempotency- und Recovery-Automation-Technologie `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert.

**CO-WP-013-Registrierungen:** 16 Policy-/Approval-/Execution-Authorization-Entscheidungen (DEC-S-137…152), **getrennte Dimensionen**: Policy Evaluation, Approval und Execution Authorization getrennt; Policy permit ≠ Approval ≠ Execution Authorization; kein Default-Permit (deny/indeterminate außer explizit autorisiert); Policy-Konflikte bleiben sichtbar; Approval explizit/zurechenbar/scope-bound/widerrufbar; Machine Principals ohne Self-Approval; Approver-Rollenname ≠ unbegrenzter Scope; Execution Authorization action-/target-/scope-/plan-/time-bound; materielle Plan-/Target-Änderung → Re-Evaluation; expired/revoked/consumed nicht wiederverwendbar; Break Glass außergewöhnlich (kein Parallelmodell); Offline Authorization mit Target-Binding/Provenance/Integrität/expliziter Aktivierung; Execution result ≠ success; Success ≠ verification, closed ≠ success; Policy-/Approval-/Authorization-Engine, Autorisierungsartefakt/Token und Replay-Mechanismus `deferred`. Keine Technologieauswahl; keine ADR. `DEC-S-01…37` bleiben unverändert (Legacy-Format-Follow-up für CO-WP-029).

**Bestätigung:** Keine technische Architektur-, Technologie- oder Implementierungsentscheidung trägt den Status `accepted`. Keine Integration ist `supported`. Keine Zertifizierung/VS-Eignung behauptet. Keine ADR ist Accepted. Es wurde keine ADR-Datei erzeugt.
