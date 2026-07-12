# CoreOps – Foundation Risk Register

> Initiales Foundation Risk Register. NDF-Basis: `v1.0.0` (Commit `9dcadc1`) — `main` nicht normativ.
> Erzeugt durch `CO-WP-002`. Quelle: [COREOPS_CONCEPT_V3.md](../docs/architecture/COREOPS_CONCEPT_V3.md).

**Wertebereiche** — Likelihood: `low`/`medium`/`high` · Impact: `low`/`medium`/`high`/`critical` · Status: `open`/`monitored`/`treatment-planned`/`accepted-by-human`/`closed`.

Risk Level ist eine qualitative Ableitung aus Likelihood × Impact. Kein Risiko ist ohne Human-Maintainer-Entscheidung `accepted-by-human`. Es werden **keine** hypothetischen Schwachstellen als bestätigte Sicherheitsbefunde dargestellt; die Einträge sind Foundation-Risiken der Produkt- und Architekturkonzeption, keine verifizierten Vulnerabilities.

| Risk ID | Title | Description | Category | Likelihood | Impact | Risk Level | Current Controls | Required Treatment | Owner | Target WP | Status | Evidence |
| ------- | ----- | ----------- | -------- | ---------- | ------ | ---------- | ---------------- | ------------------ | ----- | --------- | ------ | -------- |
| RISK-01 | Überbreiter Produktscope | Sehr große Feature-Vision (Netzwerk, Druck, Deployment, Topologie, Policy, Offline) über viele Domänen | scope | high | high | high | Foundation nur docs-only; Non-Goals §7; phasenweise Roadmap | Scope Lock, klare Phasengrenzen | Nova | CO-WP-003 | open | Concept §3–§7, §49 |
| RISK-02 | Sicherheitskomplexität | Policy, Trust Plane, mTLS, PKI, Supply Chain, Break-Glass gleichzeitig | security | medium | critical | high | Fail-closed-Prinzip; Security-first; getrennte security-baseline-WPs | Threat Model, Security-Invarianten dokumentieren | Nova | CO-WP-007 | open | Concept §15,§52 |
| RISK-03 | Air-Gap-Trust | Signierte Offline-Transfers (CorePack) und Trust bei fehlender Konnektivität | security | medium | high | high | CorePack-Importprozess mit Signatur/Hash/Freigabe §38 | Offline-Trust-Modell definieren | Nova | CO-WP-023 | open | Concept §38 |
| RISK-04 | Privilegierte Remote-Ausführung | Ausführung auf Zielsystemen trotz „keine Remote-Root-Shell" | security | medium | critical | high | Execution nur mit signierter, befristeter Freigabe §14; Invariante §52.1 | Autorisierungsmodell + Wartungs-Ausnahmen klären (CCR-07) | Nova | CO-WP-013 | open | Concept §14,§52 |
| RISK-05 | Plugin-/Domain-Pack-Supply-Chain | Fremde Packs/Skripte als Angriffsvektor | security | medium | high | high | Publisher Trust, Signatur, Quarantäne, Lebenszyklus §18,§23,§32 | Domain-Pack-Governance + Supply-Chain-Baseline | Nova | CO-WP-015/022 | open | Concept §16,§23,§32 |
| RISK-06 | Hersteller-API-/Firmware-Fragmentierung | Viele Hersteller/Firmwares mit instabilen Schnittstellen | integration | high | medium | medium | Standards First §9.4; getestete Modelle/Firmware dokumentieren §26 | Integration Quality Levels + Support Boundary | Nova | CO-WP-004/015 | open | Concept §18,§26,§27 |
| RISK-07 | Topologie-Fehlinterpretation | Abgeleitete Verbindungen als Fakt dargestellt | data-quality | medium | high | medium | Confidence-Level, Evidence Model, „nicht als zweifelsfreie Tatsache" §28 | Evidence-/Confidence-Modell + manuelle Autorität | Nova | CO-WP-020 | open | Concept §28 |
| RISK-08 | Audit vs. Datenschutz | Immutable Audit kollidiert mit Redaction/Retention | privacy | medium | high | high | Datenklassen und Retention §41 | Konfliktauflösung Audit/Redaction (CCR-08) | Nova | CO-WP-025 | open | Concept §35,§41 |
| RISK-09 | Secret Custody | Zentraler Secret Store, Rotation, Key Custody | security | medium | critical | high | Grundregeln §40; verschlüsselte Speicherung; Referenzen statt Klartext | Secrets/Vault/Key-Custody-Baseline | Nova | CO-WP-024 | open | Concept §37,§40 |
| RISK-10 | Workflow-Wiederaufnahme/Idempotenz | Langlaufende Workflows nach Neustart, doppelte Aktionen | reliability | medium | high | medium | Idempotenz, persistenter Zustand, Retries §33 | Workflow-/Idempotenz-Modell; Temporal-Bewertung | Nova | CO-WP-018/021 | open | Concept §33 |
| RISK-11 | Self-Dependency | CoreOps schaltet eigene Betriebsgrundlage ab | reliability | medium | critical | high | Self-Dependency-Erkennung, Exclusion-Tags, Quorum §42 | Self-Protection-Modell dokumentieren | Nova | CO-WP-026 | open | Concept §42 |
| RISK-12 | Falsche Health-Anzeigen bei fehlenden Daten | Datenlücken als „gesund" | data-quality | medium | high | medium | „Fehlende Daten ≠ gesund" §25; Datenabdeckung anzeigen | CoreScore-Semantik + Datenabdeckung | Nova | CO-WP-004 | open | Concept §25 |
| RISK-13 | Unklare Supportversprechen | Herstellernennung als Support/Partnerschaft missverstanden | communication | medium | medium | medium | „Herstellerangabe ≠ Supportversprechen" §26; Herausgeberklassen §18 | Support Boundary + Public-Neutrality-Wording | Nova | CO-WP-004/005 | open | Concept §18,§26 |
| RISK-14 | Release-Taxonomie-Mehrdeutigkeit | `Foundation 0.1` vs. `Release 0.1 – Observe` | planning | high | medium | medium | RELEASE_TAXONOMY.md (Proposed): Foundation `v0.0.1-foundation` vs. Observe `v0.1.0-alpha.1` getrennt | Human-Maintainer-Commit der Taxonomie; danach schließbar | HM | CO-WP-003 | treatment-planned | Concept §49; docs/governance/RELEASE_TAXONOMY.md |
| RISK-15 | Technologie-Lock-in | Frühe implizite Bindung an Stack-Kandidaten | architecture | low | high | medium | Alle Technologien nur Kandidaten §46; ADR-Pflicht | ADR-Prozess vor Festlegung | Nova | diverse ADR-WPs | open | Concept §46,§51 |
| RISK-16 | Observability-Betriebskomplexität | Prometheus/Grafana/Loki/OTel/Alertmanager Betriebsaufwand | operations | medium | medium | medium | Native Widgets zusätzlich zu Grafana §24; Degraded Modes §42 | Observability-Baseline + Betriebsgrenzen | Nova | CO-WP-019 | open | Concept §24,§42 |
| RISK-17 | NDF-Level-Ambiguität | `ndf_level:1` vs. Starter-Vorlage `2` als Reifeaussage missverstanden | governance | medium | low | low | Semantik in CO-WP-003 geklärt (Bootstrap-Status) in Manifest/Profile/Brain; Zahlenwert unverändert | Human-Maintainer-Commit der Klarstellung; danach schließbar | HM | CO-WP-003 | treatment-planned | Manifest, docs/architecture/PROJECT_BRIEF.md |
| RISK-18 | Offline-First-Mehrdeutigkeit | Runtime/Install/Update/Recovery/Build nicht klar getrennt | architecture | medium | medium | medium | Betriebsmodi Connected/Restricted/Isolated/Air-Gapped §38 | Offline-Facetten präzisieren (CCR-09) | Nova | CO-WP-023 | open | Concept §38 |
| RISK-19 | Queue-Autoritätsverwirrung | Concept-§50-Queue (21 WPs) vs. aktive Queue (31 WPs) als verbindlich missverstanden | planning | medium | medium | medium | FOUNDATION_SCOPE_LOCK.md „Superseded Planning Material"; aktive Queue autoritativ | Human-Maintainer-Commit des Scope Lock; danach schließbar | Nova | CO-WP-003 | treatment-planned | docs/governance/FOUNDATION_SCOPE_LOCK.md (CCR-11) |
| RISK-20 | Docker-first-Missverständnis | „Docker-first" als bereits akzeptierte interne Anwendungsarchitektur oder K8s-Zwang fehlinterpretiert | architecture | medium | medium | medium | PROJECT_BRIEF/RELEASE_TAXONOMY: Delivery-/Betriebsanforderung, keine Architekturentscheidung | Human-Maintainer-Commit der Einordnung; ADR erst bei konkreter Umsetzung | HM | CO-WP-003 | treatment-planned | docs/architecture/PROJECT_BRIEF.md (CCR-10) |
| RISK-21 | NDF-Basis-Missverständnis | `main` fälschlich als normativ behandelt | governance | low | high | medium | CO-WP-003-Klarstellung: `v1.0.0`/`9dcadc1` normativ, `main` informativ; Übernahme nur via eigenes WP | Human-Maintainer-Commit der Klarstellung; danach schließbar | HM | CO-WP-003 | treatment-planned | docs/architecture/PROJECT_BRIEF.md (CCR-04) |

## Verteilung nach Risk Level

- high: 8 (RISK-01, 02, 03, 04, 05, 08, 09, 11)
- medium: 12 (RISK-06, 07, 10, 12, 13, 14, 15, 16, 18, 19, 20, 21)
- low: 1 (RISK-17)

Gesamt: 21 Risiken. Status: `open` 16 · `treatment-planned` 5 (RISK-14, 17, 19, 20, 21 — durch CO-WP-003 adressiert, Schließung erst nach Human-Maintainer-Commit).

## Höchste Risiken (Fokus)

RISK-02 (Sicherheitskomplexität), RISK-04 (privilegierte Remote-Ausführung), RISK-09 (Secret Custody), RISK-11 (Self-Dependency) — jeweils Impact `critical`. Behandlung in den security-baseline-Work-Packages CO-WP-007, CO-WP-013, CO-WP-024, CO-WP-026.

## CO-WP-003 – geänderte Risiken

RISK-14 (Release-Taxonomie), RISK-17 (NDF-Level), RISK-19 (Queue-Autorität), RISK-20 (Docker-first), RISK-21 (NDF-Basis/`main`) auf `treatment-planned` gesetzt bzw. neu erfasst. **Kein Risiko ohne Evidenz geschlossen; kein Risiko `accepted-by-human`.**

**Bestätigung:** Kein erfundener Sicherheitsbefund. Kein Risiko `accepted-by-human` ohne Human-Maintainer-Entscheidung. Alle Risiken sind mit Ziel-Work-Packages verknüpft.
