# CoreOps Concept v3.1 – Amendment: Sovereignty and BSI Orientation

> Document Status: Accepted Product Direction
> Technical Architecture Status: Unconfirmed
> Certification Status: No certification claimed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Supersedes: No existing concept section
> Extends: CoreOps Concept v3.0

Erzeugt durch `CO-WP-004A – Sovereignty, BSI Orientation and Concept Amendment Registration` (docs-only / governance-baseline). Dieses Amendment registriert eine **akzeptierte strategische Produktrichtung**. Es wählt **keine** Technologie, trifft **keine** Architekturentscheidung, erzeugt **keine** ADR und behauptet **keine** Zertifizierung.

## 1. Status

Akzeptierte Produktrichtung (durch Human Maintainer). Konkrete technische Umsetzung, BSI-Mapping, Zertifizierung, ITIL-Tailoring und PRINCE2-Tailoring bleiben **offen** und späteren Work Packages vorbehalten.

## 2. Relationship to Concept v3.0

Dieses Dokument **erweitert** [COREOPS_CONCEPT_V3.md](COREOPS_CONCEPT_V3.md); es ersetzt keinen bestehenden Concept-Abschnitt. Die akzeptierte Produktvision aus v3.0 bleibt gültig; v3.1 ergänzt Souveränitäts-, Dependency- und BSI-Ausrichtung. Bei Widersprüchen gilt die in [FOUNDATION_SCOPE_LOCK.md](../governance/FOUNDATION_SCOPE_LOCK.md) definierte Binding Source Hierarchy.

## 3. Product Sovereignty

```text
CoreOps ist ein eigenständiges Produkt.

Die Kernplattform muss ohne verpflichtende externe Monitoring-,
CMDB-, ITSM-, Automation-, Container-Management- oder
GitOps-Plattform betrieben werden können.
```

Dies bedeutet **nicht**, dass CoreOps jede technische Basisfunktion selbst neu entwickelt. Externe Softwareprodukte dürfen als Markt-, Architektur-, Bedienungs- und Sicherheitsreferenz analysiert werden; CoreOps darf daraus Ideen und allgemeine Muster ableiten, **ohne** deren Code, geschützte Inhalte oder Produktabhängigkeiten ungeprüft zu übernehmen.

## 4. Controlled Dependency Policy

Es wird strikt getrennt zwischen **Managementprodukt-Abhängigkeiten** (nicht verpflichtend für den Kern) und **technischen Basisabhängigkeiten** (nach späterer Bewertung grundsätzlich zulässig). Verbindliche Details: [SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](SOVEREIGNTY_AND_DEPENDENCY_POLICY.md). In diesem Work Package wird **keine** konkrete Technologie ausgewählt.

## 5. Open Standards

CoreOps bevorzugt offene Standards und dokumentierte Schnittstellen (konsistent mit Concept v3.0 §9.4 „Standards First"). Herstellerspezifische Verfahren bleiben nur unter den bereits dokumentierten Bedingungen zulässig. Dieses Amendment legt keine konkreten Standards verbindlich fest.

## 6. BSI-oriented Development

```text
CoreOps wird BSI-orientiert und an IT-Grundschutz-Prinzipien
ausgerichtet entwickelt.

Das Projekt behauptet ohne formale Prüfung keine vollständige
BSI-Konformität, Zertifizierung, Zulassung oder VS-Eignung.
```

Zulässige und unzulässige Formulierungen sind verbindlich in [BSI_ALIGNMENT_POSITIONING.md](../security/BSI_ALIGNMENT_POSITIONING.md) geregelt.

## 7. Security and Deployment Profiles

Als Zielmodell werden drei Profile registriert (noch nicht implementiert):

- **Standard Profile** — sicherer Standardbetrieb, Self-Hosting, Homelabs, Vereine, Bildung, KMU.
- **Hardened Profile** — restriktivere Defaults, erhöhte Schutzbedarfe, getrennte Netze, Offline-/Air-Gap-Umgebungen, stärkere Funktionstrennung, strengere Audit- und Updatevorgaben.
- **Government Profile** — späteres Konfigurations-, Betriebs- und Nachweisprofil: IT-Grundschutz-Mapping als Ziel, Public-Sector-Dokumentationspaket, Härtungsbaseline, Rollen- und Funktionstrennung, Protokollierungs- und Nachweisanforderungen.

Das Government Profile ist **keine** Zertifizierung oder Zulassung.

## 8. Public-Sector Readiness

Public-Sector-Readiness ist ein **Produktziel**: auditierbare, sicherheitsorientierte, nachweisfähige Betriebsumgebungen. Readiness ist kein Nachweis von Konformität und keine behördliche Anerkennung.

## 9. Certification and Approval Boundaries

CoreOps behauptet **nicht**: BSI-zertifiziert, vollständig BSI-konform, BSI-zugelassen, VS-NfD-zugelassen, behördlich zertifiziert.

```text
VS-NfD und höhere Einstufungen sind für frühe CoreOps-Releases
keine zugesicherte Produkteigenschaft und kein aktuelles Releaseziel.
```

Die Architektur darf eine spätere Evaluierung nicht absichtlich verhindern; es wird jedoch keine Eignung versprochen.

## 10. Lessons Learned and NDF Feedback Direction

```text
Relevante Work Packages, Reviews, Fehler, Blocker und
Architekturentscheidungen müssen auf Lessons Learned geprüft werden.

Wiederverwendbare Erkenntnisse können als NDF Feedback Candidates
klassifiziert werden.

Eine Übernahme in das NDF erfolgt ausschließlich kontrolliert
durch ein separates NDF Work Package, Nova Review und
Human-Maintainer-Freigabe.
```

Der vollständige Prozess wird erst in `CO-WP-004B` dokumentiert. Dieses Amendment nimmt **keine** NDF-Änderung vor.

## 11. ITIL Alignment Candidate

```text
ITIL:
Kandidat für Operations-, Change-, Incident-, Deployment-,
Release- und Continual-Improvement-Terminologie.
```

Nur Kandidat. Keine vollständige Übernahme. Die Tailoring-Entscheidung erfolgt erst in `CO-WP-004D`.

## 12. PRINCE2-derived Governance Candidate

```text
PRINCE2:
Kandidat für schlanke Projekt-Governance, Phasensteuerung,
Toleranzen, Eskalation und Lessons Learned.
```

Nur Kandidat. Keine vollständige Übernahme. Die Tailoring-Entscheidung erfolgt erst in `CO-WP-004D`.

## 13. Extended Non-Goals

- Keine verpflichtende externe Monitoring-, CMDB-, ITSM-, Automation-, Container-Management- oder GitOps-Plattform als Kernabhängigkeit.
- Keine verpflichtende SaaS-/Cloud-Abhängigkeit für Kernfunktionen.
- Keine Zertifizierungs-, Zulassungs- oder VS-Eignungsbehauptung ohne formale Prüfung.
- Keine vollständige Übernahme von ITIL oder PRINCE2 ohne dokumentierte Tailoring-Entscheidung.
- Keine Erweiterung der Capability Matrix in diesem Work Package.

## 14. Open Decisions

- Konkrete technische Basisabhängigkeiten (Datenbanken, Frameworks, Krypto-/Protokoll-/Telemetrie-/Build-Komponenten) — offen, spätere Bewertung.
- Konkretes IT-Grundschutz-Mapping und BSI-Nachweisartefakte — spätere Security-Baseline-Work-Packages.
- ITIL-Tailoring (`CO-WP-004D`) und PRINCE2-Tailoring (`CO-WP-004D`).
- Ausgestaltung der Government/Hardened-Profile — spätere Work Packages.
- Vollständiger Lessons-Learned-/NDF-Feedback-Prozess (`CO-WP-004B`).

## 15. Follow-up Work Packages

- `CO-WP-004B` — Lessons Learned and NDF Feedback Governance
- `CO-WP-004C` — (geplant; Souveränitäts-/Security-Vertiefung)
- `CO-WP-004D` — ITIL and PRINCE2 Alignment Tailoring Decision
- `CO-WP-004E` — (geplant; Abschluss der 004er-Erweiterung)

> Die genaue Ausgestaltung der 004er-Folge-Work-Packages wird durch Nova geplant und durch den Human Maintainer freigegeben.
