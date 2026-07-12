# CoreOps – Project Brain

Kompakter, fortschreibbarer Wissensstand des Projekts. Wird pro Work Package aktualisiert.

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist **nicht** normativ.

---

## Projektstatus

Foundation – Initial Bootstrap. Es existiert ausschließlich das Core Governance Skeleton (Manifest, Profile, Brain, Context Pack, Queue, Next Phase, Roadmap).

## Aktuelle Phase

`Foundation 0.1 – Platform Foundation` (vorläufiger Arbeitsname).

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

Architekturform, Frontend, Backend/API, Datenhaltung, Topologie-Persistenz, Caching/Queue, Agent-Technologie, Observability-Stack, Policy-Engine, Workflow-Engine, Delivery Baseline (`docker-first`). **Keiner ist Accepted.**

## Bekannte Risiken

- NDF-Level-Ambiguität (Level 1 vs. Starter-Vorlage Level 2).
- Unbekannter Repository-Status (`pending-human-maintainer`).
- Offene Release-Taxonomie (`Foundation 0.1` vs. `Release 0.1 – Observe`).
- Breiter Scope mit hohem Sicherheitsanspruch (siehe [RISK_REGISTER.md](../project-system/RISK_REGISTER.md), 18 Foundation-Risiken).
- Offene Konflikte aus der Concept-Klassifikation (u. a. Plane-Taxonomie, Offline-Policy, Machine Identity vs. Air-Gap, Audit vs. Datenschutz, Docker-first-Einordnung).

## Offene Fragen

- Verbindliche NDF-Level-Vorlage?
- Bestätigung `docker-first` als Delivery Baseline?
- Finale Release-Taxonomie (→ `CO-WP-003`)?
- Technologie-Auswahl per ADR?

## Lessons Learned

- `CO-WP-001` erhielt von Nova `GO WITH NOTES`. Ursache der Note: eine nicht ausdrücklich dokumentierte read-only-Git-Abweichung (`git status --porcelain` unter einem pauschalen Git-Verbot). Konsequenz: Künftige Work Packages unterscheiden ausdrücklich zwischen **Git Read** (erlaubt) und **Git Write** (verboten ohne Human-Maintainer-Freigabe).
- `CO-WP-001A`: Der vollständige NDF-v1.0.0-Skills-Pack wurde lokal bereitgestellt (byte-identisch, verifiziert) und von Nova mit `GO` bewertet. Skills-first für spätere Work Packages ist damit lokal möglich; Auswahl erfolgt selektiv pro Work Package.
- `CO-WP-002`: Fail-closed-Blocker (fehlende Quelle) war korrekt (`GO – Blocker bestätigt`); nach Bereitstellung der lokalen read-only-Quelle wurde das Concept v3.0 vollständig registriert. Lesson: Quelle als verlässliche lokale Datei vor der Registrierung verifizieren (Identität, Abschnitte, Ende, Trunkierung).

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

## Letztes Work Package

`CO-WP-002 – Concept v3.0 Registration and Decision Classification` (docs-only). Umsetzung abgeschlossen, Nova Review ausstehend. Vorheriges WP: `CO-WP-001A – GO`.

## Nächstes Work Package

`CO-WP-003 – Project Brief, Scope Lock and Release Taxonomy` (planned-next; pending Nova review und Human-Maintainer-Freigabe).

## Human-Maintainer-Gates

Freigabe, Staging, Commit, Push, Merge, Tags, Releases, produktive Deployments sowie jede irreversible oder privilegierte Aktion sind ausschließlich dem Human Maintainer vorbehalten.

## Rückmeldung-an-Nova-Historie

- `CO-WP-001`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO WITH NOTES`.
- `CO-WP-001A`: strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova-Bewertung: `GO`.
- `CO-WP-002`: Fail-closed-Blocker gemeldet (`GO – Blocker bestätigt`); nach Source-Handoff regulär umgesetzt und strukturierte Rückmeldung mit Compact Context Summary geliefert; Nova Review `pending`.

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
