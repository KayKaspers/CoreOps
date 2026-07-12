# CoreOps – Decision Index

> Initialer Decision Index. NDF-Basis: `v1.0.0` (Commit `9dcadc1`) — `main` nicht normativ.
> Erzeugt durch `CO-WP-002`. Quelle der Einträge: [COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md) und [CONCEPT_DECISION_CLASSIFICATION.md](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md).

**Grundregeln:** Keine technische Entscheidung erhält `accepted`. Produktentscheidungen dürfen `accepted-product` sein, wenn sie im Concept ausdrücklich als akzeptierte Vision vorgegeben sind. ADRs sind ausschließlich `adr-candidate`.

Statuswerte: `accepted-product` · `binding-governance` · `open` · `adr-candidate` · `deferred` · `non-goal` · `blocked`
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
| DEC-O-02 | Release-Taxonomie / SemVer / 0.1-Kollision | open | CCR | §49 | HM | CO-WP-003 | ja | CCR-02 |
| DEC-O-03 | NDF-Level (1 vs. 2) | open | CCR | Manifest | HM | CO-WP-005 | nein | CCR-03 |
| DEC-O-04 | `main`-Normativität (Concept vs. Governance) | open | CCR | §Ausgangskontext | HM | CO-WP-005 | nein | CCR-04 |
| DEC-O-05 | Offline-Policy bei getrennter Control Plane | open | CCR/FR | §21,§38 | Nova | CO-WP-013/023 | ja | CCR-05 |
| DEC-O-06 | Machine Identity vs. Air-Gap-Laufzeiten | open | CCR/FR | §31,§38 | Nova | CO-WP-010 | ja (ADR-0013) | CCR-06 |
| DEC-O-07 | Privilegierte Ausführung vs. „keine Remote-Root-Shell" | open | CCR/FR | §52 | Nova | CO-WP-013 | ja | CCR-07 |
| DEC-O-08 | Immutable Audit vs. Redaction/Retention | open | CCR/FR | §35,§41 | Nova | CO-WP-025 | ja (ADR-0024) | CCR-08 |
| DEC-O-09 | Offline-First-Facetten (Runtime/Install/Update/Recovery/Build) | open | CCR/FR | §38 | Nova | CO-WP-023 | ja (ADR-0030) | CCR-09 |
| DEC-O-10 | Docker-first: Anforderung/Baseline/Architektur/Kandidat | open | CCR | §3,§53 | HM | CO-WP-003 | ja | CCR-10 |
| DEC-O-11 | Foundation-Queue-Abweichung (Concept §50 vs. verbindliche Queue) | open | CCR | §50 | Nova | CO-WP-003 | nein | CCR-11 |
| DEC-O-12 | Herstellersupport-Grenze | open | CCR/APC | §18,§26,§27 | Nova | CO-WP-004 | nein | CCR-12 |
| DEC-O-13 | Source-of-Truth-Konfliktprioritäten je Datenklasse | open | FR/AC | §19 | Nova | CO-WP-011 | ja (ADR-0009) | — |
| DEC-O-14 | OIC v0.1 Vertragsumfang | open | FR | §17,§18 | Nova | CO-WP-014 | ja (ADR-0006) | — |
| DEC-O-15 | Topologie-Evidence und manuelle Autorität | open | FR/AC | §28 | Nova | CO-WP-020 | ja (ADR-0020/0021) | — |

## ADR Candidates

Alle 30 Concept-ADR-Kandidaten (§51) plus zwei Foundation-Klärungen (Delivery Baseline, Release-Taxonomie) sind mit Status `adr-candidate` erfasst. Vollständige Liste: [CONCEPT_DECISION_CLASSIFICATION.md → ADR-Required Topics](../docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md). **Kein ADR ist Accepted; es wurden keine ADR-Dateien erzeugt.**

| Decision ID | Topic | Status | Class | Source | Owner | Target WP | ADR Required |
| ----------- | ----- | ------ | ----- | ------ | ----- | --------- | ------------ |
| DEC-A-0001…0030 | Concept-ADR-Kandidaten ADR-0001…ADR-0030 | adr-candidate | ADR | §51 | HM | diverse Foundation-WPs | ja |
| DEC-A-0031 | Delivery Baseline (Docker-first-Klassifikation) | adr-candidate | ADR | §3,§53 | HM | CO-WP-003 | ja |
| DEC-A-0032 | Release-Taxonomie / SemVer | adr-candidate | ADR | §49 | HM | CO-WP-003 | ja |

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
- Open Foundation Decisions: 15
- ADR Candidates: 32 (30 Concept + 2 Foundation)
- Deferred Decisions: 6
- Non-Goals: 12

**Bestätigung:** Keine technische Entscheidung trägt den Status `accepted`. Keine ADR ist Accepted. Es wurde keine ADR-Datei erzeugt.
