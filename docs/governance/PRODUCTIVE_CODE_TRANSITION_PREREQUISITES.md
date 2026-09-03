# CoreOps – Productive-Code Transition Prerequisites

> Document Status: Implemented; Nova Review abgeschlossen (Initial `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`**); `completed-go-with-notes`
> Human-Maintainer-Repository-Integration: COMPLETE — integrierter Commit `9999114200bf18baaadfb508e8464720b75e352e`; Push COMPLETE; Remote-Integration COMPLETE; `origin/main` = `9999114200bf18baaadfb508e8464720b75e352e`
> Artifact Class: Governance-Gate-Dokument (docs-only)
> Phase: `Observe` — betreten und mit Grenze autorisiert (`HM-O1` `APPROVED WITH BOUNDARY`)
> Productive Code: **NOT AUTHORIZED** · Implementation: **NOT AUTHORIZED** (Gate A, §16.1)
> `P-1` Target Authorization: **NOT AUTHORIZED** — Gate B (§16.2); **not** a Gate-A prerequisite
> `P-2` No-Mutation Evidence: method/plan **DEFINED** (Gate A) · Human-Maintainer-Disposition **APPROVED** (Gate-A-Punkt `A-12` erfüllt, §16.1) · satisfaction **NOT SATISFIED** — Gate B, downstream of implementation and execution
> `P-3` Collection Mechanism: **SELECTED** — ausdrückliche Human-Maintainer-Entscheidung innerhalb der unveränderten Decision Ceiling: Option A primär, Option B ausschließlich ergänzend, Option C **nicht** Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport (§7.5). **Kein** konkreter Pfad, **keine** API, **keine** Bibliothek, **kein** Werkzeug, **keine** Sprache und **keine** Runtime ausgewählt.
> Language / Runtime: **SELECTED** — **Go**, ausdrückliche Human-Maintainer-Entscheidung; ausgewählt ist ausschließlich die **Sprach-/Runtime-Klasse** für den ersten Observe-Slice (§8.4). **Keine** Go-Version, **keine** Distribution, **keine** Toolchain-Version, **keine** konkrete API, **keine** Abhängigkeit, **kein** Build-Werkzeug, **kein** Paketformat und **kein** breiterer CoreOps-Technologie-Stack ausgewählt. Modul-/Paketlayout und Quellbaumplatzierung sind durch **diese** Entscheidung ebenfalls **nicht** ausgewählt worden; sie sind inzwischen durch die davon getrennte Source-Tree-Entscheidung (§9.2) entschieden — siehe die folgende Zeile.
> Source Tree: **DECIDED** — ausdrückliche Human-Maintainer-Entscheidung (§9.2): repository-verwurzelte Go-Struktur; künftiger Modulwurzelort Repository-Wurzelverzeichnis; Einstiegspunkt `cmd/coreops/`; kanonische Beobachtungsdomäne `internal/observe/`; Erhebung `internal/observe/collect/`; Normalisierung zunächst **ohne** eigenes Paket innerhalb `internal/observe/`; Tests künftig **colocated `*_test.go`**; **kein** `pkg/`, **kein** `src/`, **kein** Top-Level-`tests/`, **keine** öffentliche Go-API. **NOT CREATED** — **kein** Verzeichnis, **keine** Datei, **kein** `cmd/`, **kein** `internal/`, **kein** `go.mod`, **kein** `go.sum` angelegt oder autorisiert; Modulpfad, Go-Version, Go-Distribution und Toolchain-Version **nicht** entschieden · Dependencies: **NONE ADMITTED** · Lockfile: **NONE**
> Dependency Admission: **IN FORCE** — ausdrückliche Human-Maintainer-Entscheidung (§10), Disposition `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`; Gate-A-Punkt `A-9` **erfüllt** (§16.1). Zugelassene Drittabhängigkeiten: **0** · durch `A-9` konkret zugelassen: **keine** · erster Observe-Slice: **Zero-Third-Party-Default**. Die Nutzung einer Drittabhängigkeit ist **untersagt**, solange die konkrete Abhängigkeit nicht das vollständige Admission-Verfahren durchlaufen und eine ausdrückliche Human-Maintainer-Zulassung erhalten hat; Bequemlichkeit ist **keine** hinreichende Begründung. **Kein** Paket, **kein** Modul, **keine** Bibliothek, **kein** Werkzeug, **kein** Testframework, **kein** Build-Werkzeug, **kein** Codegenerator und **kein** Vendor-Inhalt ausgewählt; **kein** `go.mod`, **kein** `go.sum` und **kein** Vendor-Baum angelegt oder autorisiert.
> `README.md` / `LICENSE`: **SHALL EXIST** — ausdrückliche Human-Maintainer-Entscheidung (§13), Disposition `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS`; Gate-A-Punkt `A-10` **erfüllt** (§16.1). Künftiges Veröffentlichungsrechte-Modell: **PUBLIC / OPEN SOURCE** · künftige Outbound-Lizenz: **Apache-2.0**, standardmäßiger und **unveränderter** Lizenztext, erst bei späterer, gesondert autorisierter Erstellung. Beide Dateien: **CREATED / PRESENT** — erstellt in einer davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung **nach** `A-10`; durch `A-10` selbst waren sie **nicht** zur Erstellung autorisiert (`Entscheidung != Artefaktautorisierung`), und die Erstellungsautorität stammt ausschließlich aus jener gesonderten Autorisierung. `README.md` liegt damit **vor** dem ersten produktiven Source-Commit vor; er bleibt öffentliche Zusammenfassung **ohne** normative Autorität. Beitragsprogramm: **NICHT AKTIV** · `CONTRIBUTING.md`: **nicht autorisiert, nicht erstellt** · CLA: **keine** · DCO: **keiner** · `SECURITY.md`: **nicht autorisiert, nicht durch `A-10` erstellt** · `NOTICE`: **nicht erstellt, nicht autorisiert** · gesonderte Markenrichtlinie: **zurückgestellt**
> Work Package: `CO-WP-032`, primärer Typ `docs-only` — autorisiert durch `HM-O3` `APPROVED`; Ausführungsgrenze `HM-O4` `APPROVED WITH EXACT BOUNDARY` (docs-only)
> Successor Work Package: none created, none reserved
> Accepted ADRs: **0** — dieses Dokument akzeptiert keinen
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)
> Nachträgliche docs-only Änderung: Die Human-Maintainer-Entscheidung zur heterogenen All-Failure-Zusammensetzung (Observation Contract §10.5 **R6**) ist in §14 als Decision-Disposition aufgenommen und wird vom bestehenden Gate-A-Punkt `A-11` getragen. Es entsteht **keine** zusätzliche Gate-A-Checklistenzeile, **kein** Decision-Identifier und **kein** ADR. `A-11` bleibt **offen**.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-12`-Disposition ist angewandt. Die Abfolge **R6-Semantikdisposition → Definition von `OBS-LLH-TC-11` → Human-Maintainer-Disposition von `A-12`** ist damit **vollständig**; `A-12` steht auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **6 von 14 erfüllt, 0 teilweise, 8 offen — nicht passiert** (durch die spätere `A-6`/`P-3`-Anwendung fortgeschrieben, siehe folgende Zeile); Gate B unverändert **0 von 8 — nicht passiert**. `A-12` erfüllt bedeutet ausschließlich, dass Testdesign und `P-2`-Evidenzmethode als **definiert** akzeptiert sind: **0** Tests implementiert, **0** Tests ausgeführt, No-Mutation-Evidenz **keine**, `P-2` **`NOT SATISFIED`**, produktiver Anwendungscode und Implementierung **NOT AUTHORIZED**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-6`/`P-3`-Disposition ist angewandt. `P-3` ist **entschieden** (§7.5): Option A primär, Option B ausschließlich ergänzend, Option C **nicht** Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport; die Decision Ceiling bleibt unverändert. Gate-A-Punkt `A-6` steht damit auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **7 von 14 erfüllt, 0 teilweise, 7 offen — nicht passiert** (durch die spätere `A-7`-Anwendung fortgeschrieben, siehe folgende Zeile); offen blieben genau `A-7`, `A-8`, `A-9`, `A-10`, `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. Die Auswahl einer Mechanismusklasse ist **keine** Sprach-/Runtime-Auswahl, **keine** Source-Tree-Autorisierung, **keine** Dependency-Zulassung, **keine** Produktivcode- oder Implementierungsautorisierung, **keine** Testimplementierung, **keine** Testausführung, **keine** `P-1`-Erteilung, **keine** Zielzugriffs- oder Beobachtungsautorisierung, **keine** `P-2`-Erfüllung und **keine** Erfüllung von `A-11`; die `target_id`-Ableitungsregel des Observation Contract ist dadurch **nicht** entschieden. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-7`-Disposition ist angewandt. Die Sprach-/Runtime-Entscheidung ist **getroffen** (§8.4): **Go** — ausschließlich als **Sprach-/Runtime-Klasse** für den ersten Observe-Slice; Rust bleibt als dokumentierte stärkste Alternative erhalten. Gate-A-Punkt `A-7` steht damit auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **8 von 14 erfüllt, 0 teilweise, 6 offen — nicht passiert** (durch die spätere `A-8`-Anwendung fortgeschrieben, siehe folgende Zeile); offen blieben genau `A-8`, `A-9`, `A-10`, `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. Die Auswahl einer Sprach-/Runtime-Klasse ist **keine** Go-Versions-, Distributions- oder Toolchain-Auswahl, **kein** Modul-/Paketlayout, **keine** API- oder Quellpfadauswahl, **keine** Source-Tree-Autorisierung, **keine** Dependency-Zulassung, **keine** Build-/Packaging-Disposition, **keine** Produktivcode- oder Implementierungsautorisierung, **keine** Testimplementierung, **keine** Testausführung, **keine** `P-1`-Erteilung, **keine** Zielzugriffs- oder Beobachtungsautorisierung, **keine** `P-2`-Erfüllung, **keine** Erfüllung von `A-11` und **keine** Auswahl eines breiteren CoreOps-Technologie-Stacks; die technische Architektur bleibt im Übrigen **unbestätigt**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-8`-Disposition ist angewandt. Die Source-Tree-Entscheidung ist **getroffen** (§9.2): `DECIDED — APPROVE OPTION A WITH NOVA BOUNDARY CLARIFICATIONS`. Entschieden ist eine repository-verwurzelte, Go-idiomatische Anwendungsstruktur: künftiger Modulwurzelort Repository-Wurzelverzeichnis; Einstiegspunkt `cmd/coreops/` ausschließlich für Komposition/Bootstrap — **keine** kanonische Beobachtungsdomänenlogik, **keine** Erhebungssemantik, **keine** Normalisierungssemantik; kanonische Beobachtungsdomäne `internal/observe/` (**internal/privat**); Erhebung `internal/observe/collect/` mit der `P-3`-Realisierung (Option A primär, Option B ausschließlich ergänzend), ohne Envelope-Zusammensetzung, ohne R1–R6, ohne kanonisches Vokabular und ohne Normalisierung; Normalisierung bleibt zunächst **ohne** eigenes Paket als quellunabhängige Rolle innerhalb `internal/observe/`; Tests künftig **colocated `*_test.go`**, **kein** Top-Level-`tests/`, paketlokales `testdata/` ausschließlich als **künftige** Fixture-Ablage, falls und sobald Fixture-/Testautorität besteht; **kein** `pkg/`, **kein** `src/`, **keine** öffentliche Go-API. Verbindliche Abhängigkeitsrichtung: `internal/observe/collect` → `internal/observe` zulässig; `internal/observe` → `internal/observe/collect` **untersagt**. Gate-A-Punkt `A-8` steht damit auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **9 von 14 erfüllt, 0 teilweise, 5 offen — nicht passiert** (durch die spätere `A-9`-Anwendung fortgeschrieben, siehe folgende Zeile); offen blieben genau `A-9`, `A-10`, `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. Die Entscheidung legt **kein** Verzeichnis und **keine** Datei an: `cmd/` und `internal/` sind **nicht** angelegt; `go.mod` ist **nicht** angelegt und **nicht** autorisiert; ein `go.sum` ist durch `A-8` **nicht gefordert**, **nicht** angelegt und **nicht** autorisiert und wäre nur dann am Repository-Wurzelort zu erwarten, wenn eine solche Datei später durch einen autorisierten Modul-, Dependency- oder Toolchain-Stand rechtmäßig erzeugt wird; Modulpfad, Go-Version, Go-Distribution und Toolchain-Version bleiben **nicht** entschieden. Die spätere Anlage der neuen Top-Level-Struktur `cmd/` und `internal/` benötigt eine ausdrücklich autorisierte Architektur-/Governance-Work-Package-Grenze nach [Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §12; ein späteres `A-14`-Work-Package genügt dafür **nur dann**, wenn sein freigegebener Scope diese Autorität ausdrücklich trägt — **automatisch** autorisiert `A-14` `cmd/` und `internal/` **nicht**. Die Auswahl einer Quellbaumstruktur ist **keine** Verzeichnis- oder Dateianlage, **keine** Dependency-Zulassung, **keine** Build-/Packaging-Disposition, **keine** Produktivcode- oder Implementierungsautorisierung, **keine** Testimplementierung, **keine** Testausführung, **keine** Fixture-Freigabe, **keine** `P-1`-Erteilung, **keine** Zielzugriffs- oder Beobachtungsautorisierung, **keine** `P-2`-Erfüllung und **keine** Erfüllung von `A-11`, `A-13` oder `A-14`. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**. Der [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) bleibt **unverändert** und **technologieunabhängig**; eine Paketplatzierung begründet **keine** neue semantische Autorität: `internal/observe` ist **nicht** `MOD-OBS-001`, ein Repository-Paket ist **keine** logische Modulautorität, und die kanonische Go-Domäne ist **nicht** der Observation Contract.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-9`-Disposition ist angewandt. Die Dependency-Admission-Entscheidung ist **getroffen** (§10): `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`. Das **Dependency-Admission-Gate ist in Kraft**; zugelassene Drittabhängigkeiten: **0**; durch `A-9` konkret zugelassene Abhängigkeiten: **keine**. Für den ersten Observe-Slice gilt der **Zero-Third-Party-Default**: die Nutzung einer Drittabhängigkeit ist **untersagt**, solange die konkrete Abhängigkeit nicht das vollständige `A-9`-Admission-Verfahren durchlaufen und eine ausdrückliche Human-Maintainer-Zulassung erhalten hat; Bequemlichkeit ist **keine** hinreichende Begründung, und ein Vorschlag ist **keine** Zulassung. Eine Drittabhängigkeit darf überhaupt nur vorgeschlagen werden, wenn **beide** Bedingungen zutreffen: ein konkreter, benannter Bedarf besteht **und** die Standardbibliotheks-/Plattformfähigkeit für genau diesen Bedarf nicht ausreicht. Verbindliche Klassifikation: Pakete der Standardbibliothek der gewählten Go-Distribution und ohne separaten Softwarebezug genutzte Plattformfähigkeit benötigen **keine** individuelle `A-9`-Drittzulassung — das bedeutet ausdrücklich **nicht**, dass deren Lieferkette geklärt wäre: Go-Distribution, Standardbibliothek und Toolchain bleiben `A-13` beziehungsweise nachgelagerter Governance zugeordnet und sind durch `A-9` **nicht** aufgelöst. Jedes **separat bezogene** Go-Modul, -Paket, jede Bibliothek und jedes Werkzeug ist **Drittabhängigkeit** und benötigt das vollständige Verfahren — einschließlich `golang.org/x/...`, sofern die betreffende Software nicht tatsächlich Bestandteil der Standardbibliothek der gewählten Go-Distribution ist; die Reputation eines Importpfads ändert die Herkunftsklassifikation **nicht**. Transitive, reine Build-, reine Test- sowie separat bezogene Tooling-/Codegenerierungsabhängigkeiten sind **im Scope**; vendorierte Drittsoftware bleibt **Drittsoftware**, da Vendoring ausschließlich ein Offline-Verfügbarkeits- und Auflösungsmechanismus ist. Die Zulassungsautorität liegt **ausschließlich** beim Human Maintainer und verlangt kumulativ eine ausdrücklich autorisierte Work-Package-Grenze, deren freigegebener Scope die konkrete Zulassungstätigkeit trägt, **und** eine ausdrückliche Zulassung je konkreter Abhängigkeit; die Zulassung ist an Version beziehungsweise unveränderliche Revision, Digest/Prüfsumme, Herkunfts-/Quellidentität, maßgeblichen Lizenzstand und aufgelöste transitive Schließung **gebunden**; eine Update-Policy autorisiert **keinen** künftigen Abhängigkeitsstand. Gate-A-Punkt `A-9` steht damit auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **10 von 14 erfüllt, 0 teilweise, 4 offen — nicht passiert** (durch die spätere `A-10`-Anwendung fortgeschrieben, siehe folgende Zeile); offen blieben genau `A-10`, `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. `A-10` blieb zu diesem Zeitpunkt **offen**: eine erste konkrete Drittzulassung konnte erst abgeschlossen werden, wenn die anwendbare Veröffentlichungs-/Lizenzdisposition zur Bewertung des Lizenzkriteriums (§10.2 Nr. 4) ausreicht — dieser Engpass ist durch die spätere `A-10`-Anwendung **behoben**, ohne dass dadurch eine Abhängigkeit bewertet oder zugelassen würde. Die Inkraftsetzung der Policy ist **keine** Zulassung, **keine** Installation, **keine** Auflösung, **keine** Paket-, Modul-, Bibliotheks-, Werkzeug- oder Testframework-Auswahl, **keine** Go-Versions-, Distributions- oder Toolchain-Auswahl, **keine** Auswahl einer konkreten Standardbibliotheks-API, **keine** `go.mod`-, `go.sum`- oder Vendor-Baum-Autorisierung, **keine** Modulpfadentscheidung, **keine** Build-/Packaging-Disposition, **keine** Source-Tree-Anlage, **keine** Produktivcode-, Implementierungs-, Testimplementierungs- oder Testausführungsautorisierung, **keine** Fixture-Freigabe, **keine** `P-1`-Erteilung, **keine** Zielzugriffs- oder Beobachtungsautorisierung, **keine** `P-2`-Erfüllung und **keine** Erfüllung von `A-10`, `A-11`, `A-13` oder `A-14`. `P-3` Option B bleibt **ausschließlich ergänzend**; `A-9` erweitert `P-3` **nicht** und hebt die Decision Ceiling **nicht** an. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer, **kein** ADR-Kandidat, **kein** Dependency-Admission-Register und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**, `A-11` bleibt **offen**, und bindend gilt: `A-9`-Disposition != `DEC-S-03`-Disposition.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-10`-Disposition ist angewandt. Die `NEW-8`-Entscheidung (`README` / `LICENSE`) ist **getroffen** (§13): `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS`. `README.md` **soll existieren**; Zeitpunkt: eine **eigene, ausdrücklich autorisierte** Repository-Änderung **vor** dem ersten produktiven Source-Commit, einschließlich Remote-Integration. Der `README` bleibt eine **öffentliche Zusammenfassung ohne normative Autorität**: er überschreibt **keine** autoritative Governance-Quelle, und die Source-of-Truth-Hierarchie ([Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §7) bleibt unverändert bindend — `README and public summaries` bleibt auf Rang 19. Quantitative Angaben im künftigen `README` sind **Spiegel, keine Autorität**: duplizierte Zähler sind zu **minimieren**, Verweise auf autoritative Current-State-Quellen sind vorzuziehen, und eine hart kodierte Zahl ist nur zulässig, wenn sie ausdrücklich als aktuell, gepflegt und **nicht-autoritativ** ausgewiesen ist. `LICENSE` **soll existieren**; das Veröffentlichungsrechte-Modell ist **PUBLIC / OPEN SOURCE**, die künftige Outbound-Lizenz ist **Apache-2.0** im standardmäßigen, **unveränderten** Lizenztext. Gate-A-Punkt `A-10` steht damit auf **erfüllt** (§16.1). Gate A nach dieser Anwendung: **11 von 14 erfüllt, 0 teilweise, 3 offen — nicht passiert**; offen bleiben genau `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. **Weder `README.md` noch `LICENSE` wird durch diese Anwendung erstellt**, und `A-10` autorisiert ihre Erstellung **nicht**; erforderlich sind dafür kumulativ (1) eine ausdrücklich autorisierte Repository-/Work-Package-Grenze, die `README.md` beziehungsweise `LICENSE` **und** deren öffentliche Inhaltsautorität benennt, (2) die Autorisierung des konkreten öffentlichen Inhalts beziehungsweise des standardmäßigen Lizenztextes und (3) die Human-Maintainer-Repository-, Staging-, Commit- und Push-Gates. Der Apache-2.0-Lizenztext wird hier **nicht** wiedergegeben, **nicht** entworfen und **nicht** verändert; **kein** Rechteinhaber-Wortlaut, **kein** Copyright-Jahr, **keine** jurisdiktionsspezifische Rechtsfolge, **keine** Durchsetzbarkeits- und **keine** Patentaussage wird hier getroffen. Ein öffentliches Beitragsprogramm ist **nicht aktiv**: `CONTRIBUTING.md` ist **nicht autorisiert** und **nicht erstellt**, es existiert **keine** CLA und **kein** DCO, und `A-10` disponiert **keinen** Beitragsworkflow, **keine** Beitragsannahme, **keine** Beitragsbedingungen und **keine** Beitragsrichtlinie; bindend bleibt: `Repository sichtbar != Beitragsprogramm aktiv`, `Pull Request möglich != Beitrag angenommen`, `unaufgeforderter Beitrag eingegangen != Beitrag angenommen, geprüft oder gemerged`. Urheberrechtliche Rechteeinräumung ist **keine** Namens- oder Markenrechtseinräumung: Apache-2.0 gewährt **keine** Rechte am Projektnamen CoreOps, am Logo, an der Wortmarke oder an der Core-Familien-Markenführung; eine sachliche Bezugnahme auf CoreOps impliziert für sich **keine** Billigung und **keine** Verbundenheit; durch `A-10` wird **keine** eingetragene Marke beansprucht, und eine gesonderte Markenrichtlinie ist **zurückgestellt** und wird hier **nicht** erstellt. Die Behauptungsgrenzen bleiben bestehen: ein künftiger `README` darf Produktionsreife, Sicherheits-, Zertifizierungs- oder Compliance-Zusicherungen, Supportzusagen, SLA, Verfügbarkeits-, Wartungs- oder Kompatibilitätsgarantien, Installationsreife und Zeitpläne **weder** aussagen **noch** nahelegen; unverändert gilt `Repository enthält Quellcode != Produkt veröffentlicht`, `Quellcode sichtbar != unterstützt`, `baubar != produktionsreif`, `lizenziert != gewährleistet`, `Open Source != unterstützter Dienst` und `Haftungsausschluss der Lizenz != zutreffende Produktstatusaussage`. `SECURITY.md` ist **nicht autorisiert** und wird durch `A-10` **nicht** erstellt; ein künftiger Meldeweg für Schwachstellen kann eine eigene Human-Maintainer-Disposition erfordern, ist aber **kein** `A-10`-Objekt und entsteht hier **nicht**. `NOTICE` wird **nicht** erstellt und **nicht** autorisiert; dass Apache-2.0 für CoreOps ohne tatsächlich vorhandenes Notice-Material zwingend eine `NOTICE`-Datei verlangte, wird hier **nicht** festgestellt. Für `A-9` folgt ausschließlich: das Lizenzkriterium (§10.2 Nr. 4) hat nunmehr ein **definiertes** CoreOps-Outbound-Veröffentlichungs-/Lizenzmodell, gegen das ein künftiger konkreter Abhängigkeitskandidat grundsätzlich bewertet werden **könnte**; `A-9` bleibt **erfüllt**, zugelassene Drittabhängigkeiten bleiben **0**, und in dieser Anwendung wird **keine** konkrete Abhängigkeit bewertet und **kein** Adoptionskandidat benannt. Die Disposition ist **keine** Artefakterstellung, **keine** Quellbaumanlage, **keine** Produktivcode-, Implementierungs-, Testimplementierungs- oder Testausführungsautorisierung, **keine** Autorisierung der Veröffentlichung produktiven Quellcodes, **kein** funktionaler Produktrelease, **kein** GitHub Release, **keine** Dependency-Zulassung, **keine** `go.mod`- oder `go.sum`-Anlage, **keine** Go-Toolchain-Auswahl, **keine** `P-1`-Erteilung, **keine** `P-2`-Erfüllung und **keine** Erfüllung von `A-11`, `A-13` oder `A-14`. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer, **kein** ADR-Kandidat, **keine** Markenrichtlinie, **keine** Beitragsrichtlinie, **keine** Sicherheitsrichtlinie, **kein** Release-Objekt und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**, und `A-11` bleibt **offen**.
> Nachträgliche docs-only Current-State-Anwendung (Publikations-Realisierung `README.md` / `LICENSE`): Eine davon getrennte, ausdrückliche Human-Maintainer-Publikations-Realisierungsautorisierung hat die **Artefakterstellung** autorisiert, die `A-10` selbst ausdrücklich **nicht** erteilt hat. `README.md` und `LICENSE` sind damit **erstellt und vorhanden**; `LICENSE` trägt den vollständigen, standardmäßigen und **unveränderten** Apache-License-2.0-Text (Version 2.0, January 2004) einschließlich `END OF TERMS AND CONDITIONS`, Appendix und dessen unveränderten Platzhaltern — **ohne** Rechteinhaber-Wortlaut, **ohne** Copyright-Jahr, **ohne** CoreOps-Einfügung, **ohne** SPDX-Zusatz und **ohne** Markdown-Formatierung. Verbindliche Autoritätskette: `A-10` entschied, dass beide Dateien existieren **sollen** → die spätere, gesonderte Human-Maintainer-Publikations-Realisierungsautorisierung erteilte die **Erstellungsautorität** → diese Realisierung hat die Dateien **erstellt** und die bestehenden Current-State-Träger abgeglichen. Historisch unverändert zutreffend: **`A-10` selbst hat die Dateien nicht erstellt und ihre Erstellung nicht autorisiert.** `A-10` wird dadurch **nicht** wiedereröffnet und bleibt **erfüllt**. Unverändert: Veröffentlichungsrechte-Modus **PUBLIC / OPEN SOURCE** · Outbound-Lizenz **Apache-2.0** · Beitragsprogramm **NICHT AKTIV** · `CONTRIBUTING.md` **nicht autorisiert / nicht erstellt** · CLA **keine** · DCO **keiner** · `SECURITY.md` **nicht autorisiert / nicht erstellt** · `NOTICE` **nicht erstellt / nicht autorisiert** · gesonderte Markenrichtlinie **zurückgestellt**, **keine** eingetragene Marke beansprucht. Gate A unverändert **11 von 14 erfüllt, 0 teilweise, 3 offen — nicht passiert**; offen bleiben genau `A-11`, `A-13` und `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. Diese Realisierung ist **keine** Produktivcode-, Implementierungs-, Source-Tree-, Build-/Packaging-, Testimplementierungs-, Testausführungs- oder Zielautorisierung, **keine** Autorisierung der Veröffentlichung produktiven Quellcodes, **keine** Dependency-Zulassung — zugelassene Drittabhängigkeiten bleiben **0** —, **keine** `go.mod`- oder `go.sum`-Anlage, **kein** funktionaler Produktrelease und **kein** GitHub Release. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Nummer, **kein** ADR-Kandidat, **keine** Marken-, Beitrags- oder Sicherheitsrichtlinie, **kein** Release-Objekt und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**. Bindend: `README erstellt != produktiver Code autorisiert` · `LICENSE erstellt != produktiver Code autorisiert` · `README vorhanden != Produkt benutzbar` · `LICENSE vorhanden != Produkt unterstützt` · `Apache-2.0 vorhanden != Abhängigkeit zugelassen` · `Open-Source-Lizenz vorhanden != Gate A passiert` · `Gate A passiert != Gate B passiert` · `Publikationsartefakte vorhanden != funktionaler Produktrelease` · `Publikationsartefakte vorhanden != GitHub Release` · `README/LICENSE-Realisierung != Source-Tree-Realisierung` · `README/LICENSE-Realisierung != A-13 erfüllt` · `README/LICENSE-Realisierung != A-14 erfüllt` · `README/LICENSE-Realisierung != CO-WP-033 erzeugt` · `Lizenz ausgewählt != Abhängigkeit kompatibel` · `Abhängigkeit kompatibel != Abhängigkeit zugelassen`.

## 1. Status und Zweck

Dieses Dokument definiert das **Governance-Gate vor späterem produktivem Anwendungscode**. Es bereitet Entscheidungen **vor**; es trifft keine. Jede hier enthaltene Empfehlung ist ausdrücklich `PROPOSED / UNACCEPTED` und wird erst durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung verbindlich.

```text
recommendation           != authorization
decision packet prepared != decision made
option evaluated         != option selected
gate documented          != gate passed
```

## 2. Foundation Relationship

**`Foundation 0.1` ist `CLOSED`** — förmlich geschlossen durch die Human-Maintainer-Entscheidung `HM-F1` (`APPROVED / COMPLETE`). Der Stand ist als annotierter, unsignierter Tag `v0.0.1-foundation` veröffentlicht; ein GitHub Release existiert nicht (`HM-R11` `NOT AUTHORIZED / CLOSED`).

Der [FOUNDATION_SCOPE_LOCK.md](FOUNDATION_SCOPE_LOCK.md) **governt und beschreibt die abgeschlossene Foundation-0.1-Phase** und bleibt für diese verbindlich. Er ist **keine dauerhafte Verbotsnorm für jede künftige produktive Arbeit**. Seine Liste *Forbidden Implementation Types* beschreibt, was **innerhalb der Foundation-Phase** unzulässig war — nicht, was für alle Zeit unzulässig bleibt.

Daraus folgt jedoch **nicht** das Gegenteil:

```text
Foundation phase closed  != productive code authorized
scope lock scoped to a closed phase != post-Foundation freedom
tag published            != functional product release
Observe entered          != implementation authorized
```

Post-Foundation-Arbeit — insbesondere produktiver Anwendungscode — benötigt eine **eigene, ausdrückliche Human-Maintainer-Autorisierung**. Eine solche ist **nicht erteilt**. Der Scope Lock selbst wird durch dieses Dokument **nicht** geändert; eine Änderung benötigte ein eigenes Work Package mit Nova Review und Human-Maintainer-Freigabe (Scope Lock, *Change Control*).

## 3. Observe Authority Boundary

`HM-O1` hat den Eintritt in die Phase `Observe` **mit Grenze** autorisiert. Der Eintritt in eine Phase ist eine Governance-Aussage, keine technische Freigabe.

| Durch Observe-Eintritt erteilt | Durch Observe-Eintritt **nicht** erteilt |
| ------------------------------ | ---------------------------------------- |
| Arbeit an Observe-Verträgen, Testdesign und Übergangs-Governance (docs-only) | Zielzugriff · reale Beobachtung · Testausführung |
| Auswahl **eines Wertslice** als Betrachtungsgegenstand (`HM-O2`) | Implementierung dieses Slice |
| Autorisierung **eines** Work Packages (`HM-O3`) mit docs-only Ausführungsgrenze (`HM-O4`) | produktiver Anwendungscode · Technologieauswahl · Nachfolge-Work-Package |

```text
Observe entered              != target access
value slice selected         != productive code authorized
CO-WP-032 authorized         != implementation authorized
work package execution authorized != productive code authorized
```

## 4. Implementation Authority

### 4.1 Zwei getrennte Gates

Produktiven Quellcode **zu schreiben** und ein **reales Ziel zu berühren** sind zwei verschiedene Autoritätsfragen mit zwei verschiedenen Gates. Sie werden hier ausdrücklich getrennt:

| Gate | Frage | Gegenstand |
| ---- | ----- | ---------- |
| **Gate A** | Darf produktiver Anwendungscode entstehen? | Implementierungs-/Produktivcode-Autorisierung (§16.1) |
| **Gate B** | Darf ein reales Ziel beobachtet werden? | Zielzugriff und Ausführung (§16.2) |

```text
productive code authorization != target authorization
implementation authorization  != target authorization
source code present           != target access granted
Gate A passed                 != Gate B passed
```

**`P-1` und die Erfüllung von `P-2` gehören zu Gate B, nicht zu Gate A.** Quellcode zu schreiben berührt kein Ziel. Ein implementierter, nie ausgeführter Beobachtungspfad greift auf nichts zu; Zielzugriff entsteht erst mit Ausführung gegen ein reales Ziel, und die dafür nötige Autorität wird eigenständig und später erteilt.

Für `P-2` gilt darüber hinaus zwingend: die No-Mutation-Evidenz **kann** vor der ersten Implementierung nicht vorliegen, weil sie einen ausführbaren Beobachtungspfad, Fixtures, eine Umgebung sowie getrennte Ziel- und Ausführungsautorität voraussetzt. `P-2` als Vorbedingung für Gate A zu führen wäre eine **zirkuläre Anforderung**. Auf Gate-A-Ebene ist deshalb ausschließlich verlangt:

```text
P-2 evidence method / plan DEFINED     <- maximal zulässige Gate-A-Anforderung
P-2 SATISFIED                          <- Gate B, nach Implementierung und Ausführung
P-2 evidence plan != P-2 satisfied
test design       != test execution
```

### 4.2 Gate A — Voraussetzungen der Implementierungsautorität

Implementierungsautorität ist **eigenständig** und entsteht nicht als Nebenwirkung. Sie setzt kumulativ voraus:

1. ein **ausdrücklich autorisiertes** implementierungsorientiertes Work Package mit benanntem Scope,
2. eine getroffene `P-3`-Entscheidung,
3. eine getroffene Sprach-/Runtime-Entscheidung,
4. eine getroffene Source-Tree-Entscheidung,
5. eine in Kraft gesetzte Dependency-Admission-Policy,
6. die Disposition von `NEW-8` (`README` / `LICENSE`),
7. die Disposition der erforderlichen ADR-/Decision-Punkte,
8. eine definierte Teststrategie beziehungsweise definierte künftige Validierungsanforderungen — einschließlich der **definierten `P-2`-Evidenzmethode**, ausdrücklich **nicht** deren Erfüllung,
9. die Build-/Packaging-Disposition, soweit dieses Dokument sie als tatsächlich erforderlich ausweist (§12).

Fehlt auch nur eine, ist **keine** Implementierungsautorität gegeben. Derzeit sind die neun Voraussetzungen **nicht kumulativ erfüllt**: mehrere stehen offen. Die `P-3`-Voraussetzung (Punkt 2) ist **erfüllt** — der lokale Erhebungsmechanismus ist als **Mechanismusklasse** ausdrücklich entschieden (§7.5; Gate-A-Punkt `A-6` **erfüllt**, §16.1); das ist ausdrücklich **keine** Sprach-/Runtime-, Source-Tree- oder Dependency-Entscheidung. Die Sprach-/Runtime-Voraussetzung (Punkt 3) ist durch eine davon getrennte, eigene Human-Maintainer-Entscheidung **erfüllt** — **Go** als **Sprach-/Runtime-Klasse** für den ersten Observe-Slice (§8.4; Gate-A-Punkt `A-7` **erfüllt**, §16.1); auch das ist ausdrücklich **keine** Source-Tree-, Dependency- oder Build-/Packaging-Entscheidung. Die Source-Tree-Voraussetzung (Punkt 4) ist durch eine davon wiederum getrennte, eigene Human-Maintainer-Entscheidung **erfüllt** — die Quellbaumstruktur ist ausdrücklich entschieden (§9.2; Gate-A-Punkt `A-8` **erfüllt**, §16.1); auch das ist ausdrücklich **keine** Verzeichnis- oder Dateianlage, **keine** Dependency-Zulassung und **keine** Build-/Packaging-Entscheidung. Die Dependency-Admission-Voraussetzung (Punkt 5) ist durch eine abermals davon getrennte, eigene Human-Maintainer-Entscheidung **erfüllt** — die Admission-Policy ist **in Kraft gesetzt** (§10; Gate-A-Punkt `A-9` **erfüllt**, §16.1). Erfüllt ist ausschließlich die **Inkraftsetzung des Verfahrens**: es ist **0** Drittabhängigkeit zugelassen, **keine** konkrete Abhängigkeit bewertet und **keine** ausgewählt — `Policy in Kraft != Abhängigkeit zugelassen`. Die `NEW-8`-Voraussetzung (Punkt 6) ist durch eine erneut davon getrennte, eigene Human-Maintainer-Entscheidung **erfüllt** — `README` und `LICENSE` sind ausdrücklich dispositioniert (§13; Gate-A-Punkt `A-10` **erfüllt**, §16.1): beide **sollen existieren**, das künftige Veröffentlichungsrechte-Modell ist **PUBLIC / OPEN SOURCE** und die künftige Outbound-Lizenz **Apache-2.0**. Erfüllt ist durch `A-10` ausschließlich die **Disposition**; durch `A-10` selbst war **keine** der beiden Dateien erstellt und ihre Erstellung **nicht** autorisiert — `Disposition getroffen != Artefakt erstellt`. Beide Dateien sind inzwischen durch die davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung **erstellt und vorhanden**; das erweitert `A-10` **nicht** und autorisiert **keine** Implementierung. Die Voraussetzung zu Teststrategie und `P-2`-Evidenzmethode (Punkt 8) ist **erfüllt** — Testdesign und Evidenzmethode sind definiert, und die erforderliche Human-Maintainer-Disposition dieser Definition ist erteilt (Gate-A-Punkt `A-12` **erfüllt**, §16.1). Das betrifft ausschließlich den **Definitionsstand**: `P-2` bleibt **`NOT SATISFIED`**, es ist **kein** Test implementiert und **kein** Test ausgeführt.

**Ausdrücklich nicht Bestandteil von Gate A:** `P-1` · die Erfüllung von `P-2` · Testausführungsautorität · Fixture-Freigabe · Lab-Bereitstellung. Diese gehören zu Gate B (§16.2).

Unverändert gilt der Human-Maintainer-Vorbehalt ([Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §6, `DEC-G-01`/`DEC-G-03`): Freigabe, Staging, Commit, Push, Merge, Branch, Tag, Release und jede irreversible Aktion liegen ausschließlich beim Human Maintainer.

## 5. `P-1` — Zielautorisierung (Gate B)

**Status: `NOT AUTHORIZED`.**

**Gate-Zuordnung: `P-1` ist eine Gate-B-Voraussetzung.** Sie ist **keine** Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren. `implementation authorization != target authorization`; Code, der nie gegen ein reales Ziel ausgeführt wird, berührt kein Ziel.

`P-1` ist die separate Human-Maintainer-Autorisierung für **das konkrete Ziel** und **die konkrete read-only Beobachtung**. Weder die Foundation, noch das Readiness Review, noch der Observe-Eintritt, noch der Observation Contract, noch dieses Dokument erteilt sie.

Eine spätere `P-1`-Erteilung muss **mindestens** benennen:

| Element | Anforderung |
| ------- | ----------- |
| **Konkretes Ziel** | welcher Host, in welcher Umgebung, in welcher Rolle; **keine** Zielklasse als Blankettautorisierung |
| **Konkrete Beobachtung** | welche Felder des [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md), aus welchen Quellklassen, mit welcher Erhebungsmechanismusklasse |
| **Ausführungsgrenze** | read-only; keine Schreib-, Management- oder Ausführungsautorität; keine Rechteeskalation; keine Credentials; keine Secrets; kein Netzwerkziel; zeitliche Begrenzung; Widerrufbarkeit; Auditpflicht |

```text
target class authorized != every target authorized
one observation authorized != a standing observation right
P-1 granted != P-2 satisfied
P-1 granted != productive code authorized
```

Dieses Dokument **autorisiert kein Ziel** und benennt keines.

## 6. `P-2` — No-Mutation- und Ausführungssicherheits-Evidenz

### 6.1 Zwei getrennte Zustände

| Zustand | Gate | Stand heute |
| ------- | ---- | ----------- |
| **`P-2`-Evidenzmethode / -plan definiert** | Gate A (§4.2 Punkt 8) | **DEFINIERT** — die sieben kumulativen Anforderungen stehen im [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §8.3; Human-Maintainer-Disposition **APPROVED** (`A-12` erfüllt, §16.1) |
| **`P-2` erfüllt** | Gate B | **`NOT SATISFIED`** |

```text
P-2 evidence plan != P-2 satisfied
```

### 6.2 Warum die Erfüllung erst nach Gate A möglich ist

Die `P-2`-Evidenz ist **beobachtetes geschütztes Verhalten**. Sie setzt voraus: einen ausführbaren Beobachtungspfad (also eine bereits erfolgte Implementierung), freigegebene Fixtures, eine deklarierte Umgebung, `P-1` und eine ausdrückliche Testausführungsautorität. Keine dieser Voraussetzungen kann vor der ersten Implementierung erfüllt sein.

`P-2` als Vorbedingung für die **Autorisierung** produktiven Codes zu führen wäre daher eine **zirkuläre Anforderung**: die Evidenz verlangt genau das, was sie angeblich freigeben soll. Die einzige auf Gate-A-Ebene zulässige Anforderung ist deshalb, dass die **Methode** feststeht — nicht, dass die Evidenz vorliegt.

```text
P-2 satisfaction is downstream of implementation and execution
defining how it will be proven != having proven it
test design                    != test execution
```

### 6.3 Inhalt der künftigen Evidenz

**Status der Erfüllung: `NOT SATISFIED`.**

Der No-Mutation-Vertrag muss durch **beobachtetes geschütztes Verhalten** belegt werden, nicht durch das Ausbleiben eines Fehlers. Die erforderliche künftige Evidenz ist vollständig im [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §8 beschrieben und wird hier **nicht** dupliziert. Kern:

```text
read-only                  != side-effect-free
no error observed          != protected behavior observed
"built as read-only"       != no mutation observed
one successful run         != regression confidence
not observable             != not present
```

Mindestens die Negativ-Familien `unknown target`, `wrong target` und `missing execution authorization` müssen als **implementierte und ausgeführte** Tests vorliegen ([Teststrategie](../testing/FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §12). Aktuell: **null Tests implementiert, null Tests ausgeführt, keine Evidenz vorhanden.**

**Zur Beschaffung dieser Evidenz wird in diesem Work Package nichts ausgeführt** — kein Zielzugriff, kein Test, kein Kommando, keine Beobachtung.

## 7. `P-3` — Decision Packet: lokaler Erhebungsmechanismus

**Status: `SELECTED` — durch eigene, ausdrückliche Human-Maintainer-Entscheidung (§7.5).** Dieser Abschnitt ist als **Entscheidungspaket** entstanden — Optionsvergleich und Empfehlung, ausdrücklich `PROPOSED / UNACCEPTED`. Diese Herkunft bleibt festgehalten: Optionsklassen (§7.2), Vergleichskriterien (§7.3) sowie die ursprüngliche Empfehlung und ihre Begründung (§7.4) bleiben unverändert erhalten und werden **nicht** rückwirkend zur Entscheidung umgedeutet. Die Empfehlung war eine Empfehlung; die Auswahl ist eine davon getrennte, spätere Human-Maintainer-Entscheidung (§7.5). Die Decision Ceiling (§7.1) bleibt unverändert.

### 7.1 Aktuelle Decision Ceiling

Das Paket bewegt sich ausschließlich innerhalb dieser Grenze:

```text
local Linux host only · no remote host · no network transport ·
no credentials · no secrets · no privileged execution
```

Optionen, die diese Grenze überschreiten, werden **nicht** bewertet, sondern ausgeschlossen.

### 7.2 Optionen (Mechanismusklassen, keine Produkte)

| Option | Beschreibung |
| ------ | ------------ |
| **A — Direktes Lesen standardisierter OS-Identitätsdateien** | Lesen der vom Betriebssystem bereitgestellten, standardisierten textuellen Identitäts- und Systeminformationsdateien im Dateisystem des lokalen Hosts |
| **B — Systeminformationsaufrufe der Standardbibliothek/Plattform** | In-Process-Abfrage der von der gewählten Sprachplattform bereitgestellten Standard-Systeminformationsschnittstellen |
| **C — Aufruf von OS-Kommandozeilenwerkzeugen und Parsen ihrer Ausgabe** | Starten mitgelieferter Systemwerkzeuge als Subprozess und Interpretieren der Textausgabe |
| **D — Abfrage eines lokalen System- oder Paketverwaltungsdienstes** | Nutzung einer lokalen Dienst-/Datenbankschnittstelle des Hosts als Informationsquelle |
| **E — Lokaler privilegierter Agent oder Daemon** | Ein dauerhaft laufender lokaler Sammler mit erhöhten Rechten |

### 7.3 Bewertung

Skala: **hoch** / **mittel** / **gering** — jeweils im Sinne der Spaltenüberschrift günstig zu lesen (bei Risiko- und Aufwandsspalten ist *gering* günstig).

| Kriterium | A — Dateien | B — Plattform-Aufrufe | C — Werkzeugaufruf | D — lokaler Dienst | E — Agent |
| --------- | ----------- | --------------------- | ------------------ | ------------------ | --------- |
| **Semantische Treue** | hoch — Quellsemantik bleibt sichtbar, Raw ist der Dateiinhalt | mittel — Plattform hat bereits interpretiert; Raw ist ein API-Rückgabewert | mittel — Ausgabe ist bereits formatiert und für Menschen gedacht | mittel bis gering — Dienstsicht ist ein weiteres Modell | gering — zusätzliche Sammler-Semantik |
| **Mutations-/Seiteneffektrisiko** | gering — lesender Dateizugriff; Rest-Seiteneffekte (Zugriffszeiten, Auditspuren) bleiben deklarationspflichtig | gering | mittel — Prozessstart, Ressourcenbindung, Logspuren | mittel — Dienstinteraktion, ggf. Zustandsänderung im Dienst | hoch — dauerhafte Präsenz, Konfigurations- und Zustandswirkung |
| **Privilegienbedarf** | gering — üblicherweise unprivilegiert lesbar | gering | mittel — Werkzeugverfügbarkeit und -rechte variieren | mittel bis hoch | **hoch — verletzt die Decision Ceiling** |
| **Portabilität (Linux)** | hoch für die standardisierten Identitätsquellen; einzelne Distributionen weichen ab | hoch | gering bis mittel — Werkzeugpräsenz, -version und -ausgabeformat variieren stark | gering — dienstabhängig, nicht überall vorhanden | mittel |
| **Provenance-Qualität** | hoch — Quelle ist eindeutig benennbar und referenzierbar | mittel — Quelle ist die Plattform, nicht die Ursprungsquelle | gering — Provenance endet beim Werkzeug, dessen eigene Quelle verborgen bleibt | mittel | gering |
| **Fehlerbeobachtbarkeit** | hoch — Abwesenheit, Verweigerung und Malformedness sind sauber trennbar | mittel — Fehlerklassen sind plattformabhängig zusammengefasst | gering — Exit-Code und Text vermischen Ursachen | mittel | gering |
| **Abhängigkeitslast** | gering — keine Drittabhängigkeit erforderlich | gering | mittel — Laufzeitabhängigkeit von externen Werkzeugen | hoch | hoch |
| **Offline-Eignung** | hoch | hoch | hoch | mittel | mittel |
| **Testbarkeit** | hoch — Quellen sind als Dateien fixture-fähig; alle zehn Envelope-Fälle sind darstellbar | mittel — Plattformaufrufe brauchen Abstraktion und Test Doubles | gering — Subprozessverhalten ist schwer deterministisch nachzustellen | gering | gering |
| **Implementierungskomplexität** | gering bis mittel — Parsing, aber klar begrenzt | gering | mittel | hoch | hoch |

### 7.4 Empfehlung

```text
RECOMMENDATION (P-3):  Option A als primärer Mechanismus,
                       Option B ergänzend, wo A keine Quelle bietet.
                       Option C nicht als Standardweg.
                       Option D und Option E für diesen Slice ausgeschlossen.
RECOMMENDATION STATUS: PROPOSED / UNACCEPTED (Stand der Empfehlung)
P-3 (heute):           SELECTED — eigene Human-Maintainer-Entscheidung (§7.5)
```

**Begründung.** Option A liefert als einzige Option gleichzeitig die höchste Provenance-Qualität (die Quelle ist benennbar und der Raw-Wert ist der Quellinhalt selbst), die sauberste Fehlerbeobachtbarkeit (Abwesenheit, Verweigerung und Malformedness bleiben trennbar, was Abschnitt §10 des Observation Contract überhaupt erst prüfbar macht) und die höchste Testbarkeit (alle zehn Envelope-Fälle sind mit Datei-Fixtures darstellbar). Option B schließt Lücken dort, wo keine standardisierte Datei existiert, verliert dabei aber Provenance-Tiefe: die Plattform hat bereits interpretiert. Option C wird **nicht** als Standardweg empfohlen, weil Werkzeugausgaben für Menschen und nicht für Verträge formatiert sind, weil Exit-Code und Textausgabe die vier zu trennenden Fehlerursachen vermischen und weil ein Subprozessstart das Seiteneffektprofil unnötig vergrößert. Option D und E überschreiten Abhängigkeits- beziehungsweise Privilegiengrenze und sind für einen ersten read-only Slice nicht vertretbar; E verletzt die Decision Ceiling unmittelbar.

**Diese Empfehlung ist keine Auswahl.** Sie benennt keine konkreten Pfade, keine Bibliothek und kein Werkzeug. Die spätere Auswahl ist **nicht** durch diese Empfehlung entstanden, sondern durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung (§7.5).

### 7.5 Human-Maintainer-Entscheidung

Der Human Maintainer hat `P-3` ausdrücklich **entschieden**. Die Entscheidung ist eine **Mechanismusklassen-Entscheidung** und bleibt vollständig innerhalb der unveränderten Decision Ceiling aus §7.1.

```text
DECISION (P-3):        SELECTED
PRIMARY:               Option A — direktes Lesen standardisierter OS-Identitäts-
                       und Systeminformationsdateien auf dem lokalen Host
SUPPLEMENTARY ONLY:    Option B — In-Process-Systeminformationsschnittstellen der
                       Standardbibliothek/Plattform, ausschließlich dort, wo die
                       verlangte Beobachtung aus einer geeigneten standardisierten
                       Dateiquelle nicht wahrheitsgemäß gewonnen werden kann
NOT THE STANDARD PATH: Option C
EXCLUDED (this slice): Option D · Option E
TRANSPORT:             kein Netzwerktransport
CEILING:               unverändert — local Linux host only · no remote host ·
                       no network transport · no credentials · no secrets ·
                       no privileged execution
DECIDED BY:            Human Maintainer — eigene, ausdrückliche Entscheidung
```

**Was diese Entscheidung ausdrücklich nicht auswählt.** Sie benennt **keinen** konkreten Quellpfad, **keine** API, **keine** Bibliothek, **kein** Kommando, **kein** Werkzeug, **keine** Sprache und **keine** Runtime. Ausgewählt ist die **Klasse** des Erhebungsmechanismus, nicht ihre Realisierung.

```text
mechanism class selected != source path selected
mechanism class selected != API selected
mechanism class selected != library selected
mechanism class selected != tool selected
mechanism class selected != language/runtime selected
mechanism class selected != target_id derivation decided
mechanism class selected != source tree authorized
mechanism class selected != dependency admitted
mechanism class selected != productive code authorized
mechanism class selected != implementation authorized
mechanism class selected != tests implemented
mechanism class selected != tests executed
mechanism class selected != P-1 authorized
mechanism class selected != target access authorized
mechanism class selected != real observation authorized
mechanism class selected != P-2 satisfied
mechanism class selected != A-11 satisfied
mechanism class selected != ADR accepted
A-6 satisfied            != Gate A passed
Gate A passed            != Gate B passed
```

Die Entscheidung erfüllt den Gate-A-Punkt `A-6` (§16.1). Sie erzeugt **keine** ADR-Datei, **keine** ADR-Nummer, **keine** Decision-ID, **keine** Zeile im Decision Index oder Risk Register, **keine** Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung und **kein** Work Package; akzeptierte ADRs bleiben **0** und `A-11` bleibt **offen** (§14). Im [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) bleibt `collection_mechanism_class` unverändert **Provenance-Metadatum**; die dortige `target_id`-Ableitungsregel ist durch diese Entscheidung **nicht** entschieden. Der Observation Contract wird durch sie **nicht** geändert.

## 8. Decision Packet: Sprache und Runtime

**Status: `SELECTED` — Go, durch eigene, ausdrückliche Human-Maintainer-Entscheidung (§8.4).** Dieser Abschnitt ist als **Entscheidungspaket** entstanden — Kandidatenvergleich und Empfehlung, ausdrücklich `PROPOSED / UNACCEPTED`. Diese Herkunft bleibt festgehalten: die Randbedingungen (§8.1), der Sechs-Kandidaten-Vergleich (§8.2) sowie die ursprüngliche Empfehlung und ihre Begründung (§8.3) bleiben unverändert erhalten und werden **nicht** rückwirkend zur Entscheidung umgedeutet. Die Empfehlung war eine Empfehlung; die Auswahl ist eine davon getrennte, spätere Human-Maintainer-Entscheidung (§8.4). Es wird **keine** frühere Technologieentscheidung unterstellt: die technischen Entscheidungen des Repositories umfassen ausschließlich die `P-3`-**Mechanismusklasse** (§7.5) und diese **Sprach-/Runtime-Klasse** (§8.4); ein breiterer CoreOps-Technologie-Stack ist **nicht** ausgewählt, die technische Architektur bleibt im Übrigen **unbestätigt**, und es bleibt bei **0** akzeptierten ADRs.

### 8.1 Randbedingungen aus bestehender Governance

- **Standalone-First / Souveränität** — der Kern muss ohne verpflichtende externe Plattform betreibbar sein ([Sovereignty and Dependency Policy](../architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) §2–§4).
- **Offline-/Air-Gap-Fähigkeit** — jede zugelassene Basisabhängigkeit muss offline verfügbar sein (ebenda §7).
- **Docker-first ist Delivery Baseline** (`HM-2` `APPROVED WITH BOUNDARY`) — und ausdrücklich `Docker-first != Docker-only != zwingende interne Anwendungsarchitektur != zwingende Runtime-Abhängigkeit != Observe-Voraussetzung`. Docker-first **präjudiziert die Sprachwahl nicht**.
- **DE/EN als Produktsprachen**, Englisch kanonisch für maschinenbezogene Artefakte ([Language Standard](COREOPS_LANGUAGE_STANDARD.md) §6).

### 8.2 Kandidaten und Bewertung

Bewertet werden **Sprach-/Runtime-Klassen**, nicht Versionen oder Distributionen.

| Kriterium | Go | Rust | Python | TypeScript / Node.js | JVM (Java/Kotlin) | .NET (C#) |
| --------- | -- | ---- | ------ | -------------------- | ----------------- | --------- |
| **Standalone/self-hosted-Eignung** | hoch — einzelnes statisch gelinktes Binary | hoch — einzelnes Binary | gering — Interpreter plus Umgebung | gering bis mittel — Runtime plus Modulbaum | mittel — Runtime nötig; self-contained Images möglich | mittel — self-contained Publish möglich |
| **Offline-Betrieb** | hoch | hoch | mittel — Umgebungsauflösung offline aufwendig | mittel | mittel | mittel |
| **Minimaler Deployment-Footprint** | hoch | hoch | gering | gering | gering bis mittel | mittel |
| **Linux-Portabilität** | hoch — breite Architekturabdeckung, einfache Cross-Kompilierung | hoch | hoch (Quellebene), gering (Verteilung) | mittel | hoch | mittel bis hoch |
| **Sicherheitsoberfläche** | mittel bis hoch — speichersicher im Normalfall, kleine Laufzeitoberfläche | hoch — stärkste Speicher- und Nebenläufigkeitsgarantien | mittel — große Laufzeit- und Paketoberfläche | gering bis mittel — große Paketoberfläche | mittel | mittel |
| **Deterministische Paketierung** | hoch — Modul-Prüfsummen, Vendoring, reproduzierbare Builds gut erreichbar | hoch — Lockfile plus Vendoring | gering bis mittel | gering bis mittel | mittel | mittel |
| **Abhängigkeitslast für diesen Slice** | sehr gering — Standardbibliothek genügt plausibel | sehr gering bis gering | mittel | mittel bis hoch | mittel | mittel |
| **Wartbarkeit** | hoch — kleine Sprache, flache Lernkurve, stabile Kompatibilitätszusage | mittel — höhere Einstiegs- und Umbaukosten | hoch (Lesbarkeit), mittel (Betriebsstabilität) | mittel | mittel bis hoch | mittel bis hoch |
| **Testbarkeit** | hoch — Testwerkzeuge in der Standarddistribution | hoch | hoch | hoch | hoch | hoch |
| **Künftige Cross-Plattform-Implikationen** | hoch — Cross-Kompilierung ohne Fremdtoolchain | hoch — mit höherem Toolchain-Aufwand | mittel | mittel | hoch | hoch |
| **Supply-Chain-Implikationen** | günstig — wenige Abhängigkeiten, prüfsummengebundene Auflösung | günstig — Lockfile, aber tendenziell tiefere Abhängigkeitsbäume | ungünstig — breite, transitiv tiefe Ökosystemoberfläche | ungünstig — sehr breite, tiefe Ökosystemoberfläche | mittel | mittel |

### 8.3 Empfehlung

```text
RECOMMENDATION (language/runtime):  Go als bevorzugter Kandidat.
                                    Rust als stärkste Alternative.
RECOMMENDATION STATUS:              PROPOSED / UNACCEPTED (Stand der Empfehlung)
LANGUAGE / RUNTIME (heute):         SELECTED — Go, eigene Human-Maintainer-
                                    Entscheidung (§8.4)
```

**Begründung.** Der erste Slice ist bewusst schmal: ein read-only Erhebungspfad, der Dateien liest, normalisiert und ein Ergebnis mit Provenance erzeugt. Für diese Aufgabe schneidet Go auf den CoreOps-tragenden Kriterien am besten ab — einzelnes selbstständiges Binary ohne Laufzeitinstallation (Standalone-First, Offline, minimaler Footprint), prüfsummengebundene Modulauflösung mit Vendoring (deterministische Paketierung, Supply-Chain), plausibel **null** Drittabhängigkeiten für diesen Slice (Dependency-Admission bleibt leer), Testwerkzeuge in der Standarddistribution und einfache Cross-Kompilierung für spätere Plattformerweiterung.

Rust ist bei Sicherheitsgarantien überlegen und bei Paketierung gleichwertig; die höheren Einstiegs-, Umbau- und Toolchain-Kosten stehen für einen ersten, absichtlich kleinen Slice in keinem günstigen Verhältnis. Python und TypeScript/Node.js scheitern nicht an der Aufgabe, sondern an den CoreOps-Randbedingungen: Laufzeitabhängigkeit, größerer Footprint, schwächere deterministische Paketierung und eine deutlich breitere Supply-Chain-Oberfläche. JVM und .NET sind tragfähig, bringen aber ohne Gegenwert für diesen Slice eine schwerere Laufzeit- und Paketierungsgeschichte mit.

**Diese Empfehlung ist keine Auswahl.** Sie war ausdrücklich ADR-relevant und blieb bis zu einer eigenen Human-Maintainer-Entscheidung `PROPOSED / UNACCEPTED`. Die spätere Auswahl ist **nicht** durch diese Empfehlung entstanden, sondern durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung (§8.4). Die dokumentierte ADR-Relevanz bleibt bestehen und ist **nicht** eingelöst: es existiert **kein** ADR.

```text
language recommendation != language selection
Docker-first            != language constraint
```

### 8.4 Human-Maintainer-Entscheidung

Der Human Maintainer hat die Sprach-/Runtime-Frage ausdrücklich **entschieden**. Die Entscheidung ist eine **Sprach-/Runtime-Klassen-Entscheidung** für den **ersten Observe-Slice** und geht über diesen Zuschnitt nicht hinaus.

```text
DECISION (language/runtime):  SELECTED — Go
SCOPE:                        Sprach-/Runtime-Klasse ausschließlich für den
                              ersten Observe-Slice
STRONGEST ALTERNATIVE:        Rust — als dokumentierte Alternative unverändert
                              erhalten (§8.2, §8.3)
NOT SELECTED:                 Go-Version · Go-Distribution · Toolchain-Version ·
                              Modul-/Paketlayout · Modul-/Paketname ·
                              konkrete API · Quellpfad · Abhängigkeit ·
                              Build-Werkzeug · Paketformat ·
                              breiterer CoreOps-Technologie-Stack
DECIDED BY:                   Human Maintainer — eigene, ausdrückliche Entscheidung
```

**Was diese Entscheidung ausdrücklich nicht auswählt.** Sie benennt **keine** Go-Version, **keine** Go-Distribution, **keine** Toolchain- oder Compilerversion, **kein** Modul- oder Paketlayout, **keinen** Modul- oder Paketnamen, **keine** konkrete API, **keinen** Quellpfad, **keine** Abhängigkeit, **kein** Build-Werkzeug und **kein** Paketformat. Ausgewählt ist die **Klasse** der Implementierungsplattform, nicht ihre Realisierung — und ausdrücklich **kein** breiterer CoreOps-Technologie-Stack.

```text
Go selected   != Go version selected
Go selected   != Go distribution selected
Go selected   != toolchain selected
Go selected   != module/package layout selected
Go selected   != concrete API selected
Go selected   != source path selected
Go selected   != source tree authorized
Go selected   != dependency admitted
Go selected   != build/packaging approved
Go selected   != broader technology stack selected
Go selected   != architecture confirmed
Go selected   != productive code authorized
Go selected   != implementation authorized
Go selected   != tests implemented
Go selected   != tests executed
Go selected   != P-1 authorized
Go selected   != target access authorized
Go selected   != real observation authorized
Go selected   != P-2 satisfied
Go selected   != A-11 satisfied
Go selected   != ADR accepted
A-7 satisfied != Gate A passed
Gate A passed != Gate B passed
```

Die Entscheidung erfüllt den Gate-A-Punkt `A-7` (§16.1). Sie erzeugt **keine** ADR-Datei, **keine** ADR-Nummer, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register, **keine** Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung und **kein** Work Package; akzeptierte ADRs bleiben **0** und `A-11` bleibt **offen** (§14). Sie legt **kein** Verzeichnis und **keine** Datei an — insbesondere **kein** `go.mod`, **kein** `go.sum`, **kein** Package-Manifest, **keinen** Vendor-Baum und **kein** Lockfile; `A-8`, `A-9` und `A-13` blieben zum Zeitpunkt dieser Entscheidung **offen** — `A-8` ist inzwischen durch eine davon getrennte, spätere Human-Maintainer-Entscheidung **erfüllt** (§9.2), `A-9` und `A-13` bleiben **offen**. Der [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) bleibt **technologieunabhängig**: die Entscheidung führt dort **keine** Go-spezifischen Typen, Schnittstellen, Schemata, Paket- oder Modulnamen, APIs, Quellpfade oder Implementierungskonventionen ein und ändert seine normative Beobachtungssemantik **nicht**.

## 9. Source-Tree-Entscheidung

**Status: `DECIDED` — durch eigene, ausdrückliche Human-Maintainer-Entscheidung (§9.2). Es wird kein Verzeichnis angelegt.** Dieser Abschnitt ist als **Vorschlag** entstanden — eine rollenbasierte, ausdrücklich `PROPOSED / UNACCEPTED` gehaltene Struktur ohne Sprachidiom. Diese Herkunft bleibt festgehalten: der ursprüngliche Vorschlag (§9.1) bleibt unverändert erhalten und wird **nicht** rückwirkend zur Entscheidung umgedeutet. Der Vorschlag war ein Vorschlag; die Auswahl ist eine davon getrennte, spätere Human-Maintainer-Entscheidung (§9.2). Das konkrete Verzeichnisidiom folgt der Sprach-/Runtime-Entscheidung aus §8 und wurde erst nach ihr bestimmt.

### 9.1 Ursprünglicher rollenbasierter Vorschlag (historisch, `PROPOSED / UNACCEPTED`)

Der Vorschlag war ausschließlich Markdown und beschrieb **Rollen**, keine Sprachidiome:

```text
<repository root>
├─ docs/                      bestehend, unverändert
├─ project-brain/             bestehend, unverändert
├─ project-system/            bestehend, unverändert
└─ <application root>/        VORSCHLAG — nicht angelegt
   ├─ <entry point>/          Programmeinstieg; enthält keine Domänenlogik
   ├─ <contract>/             Vertragstypen des Observation Contract:
   │                          Feldidentität, observation_outcome, Zeitmodell,
   │                          Provenance, Freshness — ohne Erhebungswissen
   ├─ <collection>/           Erhebung; kapselt den P-3-Mechanismus vollständig
   │                          hinter der Vertragsgrenze
   ├─ <normalization>/        Raw auf kanonische Repräsentation; kennt keine Quelle
   └─ <tests>/                Testfälle des Test Envelope
```

**Optimierungsziele** und ihre Umsetzung im Vorschlag:

| Ziel | Umsetzung |
| ---- | --------- |
| Schmaler erster Slice | genau vier Rollen plus Einstieg; keine Schichten „auf Vorrat" |
| Klare Domänengrenzen | `<contract>` kennt keinen Mechanismus; `<collection>` kennt keine Normalisierung; `<normalization>` kennt keine Quelle |
| Künftige Testbarkeit | die `P-3`-Abhängigkeit liegt vollständig in `<collection>`; die übrigen Rollen sind ohne Ziel und ohne Umgebung prüfbar |
| Minimum an verfrühter Abstraktion | keine Plugin-, Adapter-, Registry-, Transport- oder Konfigurationsschicht; keine Datenbank; kein Service-Layer |
| Standalone-First | ein Anwendungswurzelverzeichnis, ein Einstiegspunkt, keine externe Laufzeitvoraussetzung im Entwurf |

**Ausdrücklich nicht Bestandteil des Vorschlags:** ein Frontend, eine API-Schicht, eine Persistenzschicht, ein Agent, ein Scheduler, ein Konfigurationssystem, eine CI-Definition, ein Containerartefakt. Diese Ausschlüsse gelten für die Entscheidung aus §9.2 unverändert fort.

### 9.2 Human-Maintainer-Entscheidung

Der Human Maintainer hat die Source-Tree-Frage ausdrücklich **entschieden**. Die Entscheidung ist eine **Struktur- und Platzierungsentscheidung** für den **ersten Observe-Slice** und geht über diesen Zuschnitt nicht hinaus.

```text
DECISION (source tree):  DECIDED — APPROVE OPTION A WITH NOVA BOUNDARY CLARIFICATIONS
LAYOUT:                  repository-verwurzelte, Go-idiomatische Anwendungsstruktur
MODULE ROOT (künftig):   Repository-Wurzelverzeichnis
ENTRY POINT (künftig):   cmd/coreops/
OBSERVATION DOMAIN:      internal/observe/           — internal/privat
COLLECTION:              internal/observe/collect/   — internal/privat
NORMALIZATION:           zunächst ohne eigenes Paket; quellunabhängige Rolle
                         innerhalb internal/observe/
TESTS (künftig):         colocated *_test.go
NOT USED:                pkg/ · src/ · Top-Level-tests/ · öffentliche Go-API
CREATED:                 nichts — kein Verzeichnis, keine Datei
DECIDED BY:              Human Maintainer — eigene, ausdrückliche Entscheidung
```

**Künftige Struktur — entschieden, nicht angelegt.**

```text
<repository root>/
├─ .claude/                   bestehend, unverändert
├─ docs/                      bestehend, unverändert
├─ project-brain/             bestehend, unverändert
├─ project-system/            bestehend, unverändert
├─ ROADMAP.md                 bestehend, unverändert
├─ go.mod                     NUR künftiger Ort — nicht angelegt, nicht autorisiert
├─ go.sum                     NUR bedingter künftiger Ort — ausschließlich falls
│                             später durch einen autorisierten Modul-, Dependency-
│                             oder Toolchain-Stand rechtmäßig erzeugt; durch A-8
│                             nicht gefordert, nicht angelegt, nicht autorisiert
├─ cmd/                       ENTSCHIEDEN — nicht angelegt
│  └─ coreops/                Einstiegspunkt; ausschließlich Komposition/Bootstrap;
│                             keine kanonische Beobachtungsdomänenlogik, keine
│                             Erhebungssemantik, keine Normalisierungssemantik
└─ internal/                  ENTSCHIEDEN — nicht angelegt
   └─ observe/                kanonische Beobachtungsdomäne: kanonische
      │                       Beobachtungsidentität · kanonische Feldidentitäten ·
      │                       die gebundenen 8/9/3-Vokabulare · deterministische
      │                       R1–R6-Zusammensetzung · observed_at/received_at ·
      │                       Raw und Normalized getrennt gehalten ·
      │                       Provenance-Struktur · Freshness-Vokabular ·
      │                       Absenz-/Fehlersemantik · Normalisierung als
      │                       quellunabhängige Rolle
      └─ collect/             Erhebung: P-3-Realisierung (Option A primär,
                              Option B ausschließlich ergänzend) · Raw-Beobachtungs-
                              werte · feldweise Erhebungsausgänge · feldweise
                              Provenance-Fakten
```

**Verbindliche Abhängigkeitsrichtung.**

```text
cmd/coreops              -> internal/observe             zulässig
cmd/coreops              -> internal/observe/collect     zulässig, soweit die
                                                         Komposition es erfordert
internal/observe/collect -> internal/observe             zulässig
internal/observe         -> internal/observe/collect     UNTERSAGT
```

`internal/observe` trägt **keine** Erhebungssemantik. `internal/observe/collect` trägt **keine** Envelope-Zusammensetzung, **kein** R1–R6, **kein** kanonisches Vokabular und **keine** Normalisierung. `cmd/coreops` trägt **keine** kanonische Beobachtungsdomänenlogik, **keine** Erhebungssemantik und **keine** Normalisierungssemantik.

**Sichtbarkeit.** Alle künftigen Implementierungspakete liegen unterhalb von `internal/` und sind damit **internal/privat**. Eine öffentliche Go-API ist **nicht** ausgewählt und **nicht** autorisiert; `pkg/` wird **nicht** verwendet, `src/` wird **nicht** verwendet.

**Normalisierung.** Zunächst **kein** eigenes Paket. Die Rolle bleibt vollständig erhalten und **quellunabhängig** innerhalb `internal/observe/`. Eine spätere Aufteilung in ein eigenes Paket benötigt eine **eigene, begründete Entscheidung**.

**Tests.** Künftige Platzierungspolitik: Go-idiomatische, **colocated** `*_test.go`-Dateien neben ihren Paketen. **Kein** Top-Level-`tests/`. Ein **paketlokales** `testdata/` ist ausschließlich als **künftige** Ablage für Fixtures zulässig, **falls und sobald** Fixture- und Testautorität später besteht. Es wird **jetzt kein** `testdata/`-Verzeichnis angelegt und **keine** Fixture autorisiert. Der [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) bleibt ein Governance-/Testdesign-Artefakt und begründet **kein** Quellbaumverzeichnis.

```text
test placement decided       != test implemented
testdata placement permitted != fixture authorized
test design defined          != test execution authorized
```

**Nicht** entschieden bleiben: konkrete Testdateinamen · In-Package- gegen External-Package-Testkonvention · Testframework · Fixture-Inhalt · Fixture-Identität · Testimplementierung · Testausführung · Testautorität.

**`go.mod` und `go.sum`.** `A-8` entscheidet ausschließlich den **künftigen Ort** als Folge der Strukturwahl — mehr nicht.

```text
go.mod künftiger Ort:  Repository-Wurzelverzeichnis
go.mod:                NICHT ANGELEGT · NICHT AUTORISIERT
Modulpfad:             NICHT ENTSCHIEDEN
Go-Version:            NICHT ENTSCHIEDEN
Go-Distribution:       NICHT ENTSCHIEDEN
Toolchain-Version:     NICHT ENTSCHIEDEN

go.sum bedingter künftiger Ort:  Repository-Wurzelverzeichnis — ausschließlich
                       dann, wenn eine solche Datei später durch einen
                       autorisierten Modul-, Dependency- oder Toolchain-Stand
                       rechtmäßig erzeugt wird
go.sum:                DURCH A-8 NICHT GEFORDERT · NICHT ANGELEGT ·
                       NICHT AUTORISIERT
```

Ein `go.sum` ist damit **weder zugesichert noch notwendigerweise zu erwarten**. Ein benannter Ort ist keine erzeugte Datei.

### 9.3 Verzeichnisanlage — eigenständige `RGS`-§12-Grenze

**Die Entscheidung wählt eine künftige Struktur; sie autorisiert deren Anlage nicht.**

Der [Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §12 verlangt für **neue Top-Level-Struktur** ein eigenes, ausdrücklich autorisiertes Architektur- oder Governance-Work-Package. `cmd/` und `internal/` wären neue Top-Level-Struktur. Diese Anforderung besteht **eigenständig** neben `A-14` und wird durch `A-8` **nicht** erfüllt:

```text
A-8 entscheidet:  künftige Struktur und Platzierung
A-14 bleibt:      ausdrückliche implementierungsorientierte Work-Package-Autorität
RGS §12 bleibt:   eigenständige Scope-Anforderung für die tatsächliche Anlage
                  neuer Top-Level-Verzeichnisse
```

Ein späteres `A-14`-Work-Package genügt `RGS` §12 **nur dann**, wenn sein **freigegebener Scope ausdrücklich** die dafür erforderliche Architektur-/Governance-Autorität für die Anlage von `cmd/` und `internal/` trägt. Ein `A-14`-Work-Package autorisiert `cmd/` und `internal/` **nicht** automatisch.

### 9.4 Was diese Entscheidung ausdrücklich nicht auswählt

Sie benennt **keinen** Modulpfad, **keine** Go-Version, **keine** Go-Distribution, **keine** Toolchain-Version, **keine** konkrete API, **keine** konkrete Standardbibliotheksfunktion, **keinen** konkreten Erhebungsquellpfad, **keine** Abhängigkeit, **kein** Build-Werkzeug, **kein** Paketformat, **kein** Testframework und **keinen** breiteren CoreOps-Technologie-Stack. Sie legt **kein** Verzeichnis und **keine** Datei an.

```text
A-8 satisfied != source tree created
A-8 satisfied != cmd/ created
A-8 satisfied != internal/ created
A-8 satisfied != go.mod created
A-8 satisfied != go.sum created
A-8 satisfied != module path decided
A-8 satisfied != Go version decided
A-8 satisfied != dependency admitted
A-8 satisfied != build/packaging decided
A-8 satisfied != productive code authorized
A-8 satisfied != implementation authorized
A-8 satisfied != test implementation authorized
A-8 satisfied != test execution authorized
A-8 satisfied != fixture authorized
A-8 satisfied != P-1 granted
A-8 satisfied != target access authorized
A-8 satisfied != real observation authorized
A-8 satisfied != P-2 satisfied
A-8 satisfied != A-11 satisfied
A-8 satisfied != A-13 satisfied
A-8 satisfied != A-14 satisfied
A-8 satisfied != CO-WP-033 created or reserved
A-8 satisfied != Gate A passed
Gate A passed != Gate B passed
source-tree decision != source tree created
directory named      != directory exists
```

Die Entscheidung erfüllt den Gate-A-Punkt `A-8` (§16.1). Sie erzeugt **keine** ADR-Datei, **keine** ADR-Nummer, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register, **keine** Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung und **kein** Work Package; akzeptierte ADRs bleiben **0** und `A-11` bleibt **offen** (§14). Der [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) bleibt **unverändert** und **technologieunabhängig**: die Entscheidung führt dort **keine** Go-spezifischen Typen, Schnittstellen, Schemata, Paket- oder Modulnamen, APIs, Quellpfade oder Implementierungskonventionen ein und ändert seine normative Beobachtungssemantik **nicht**. Eine Paketplatzierung begründet **keine** neue semantische Autorität.

```text
internal/observe     != MOD-OBS-001
kanonische Go-Domäne != Observation Contract
Repository-Paket     != logische Modulautorität
shared semantics     != shared authority
integration          != runtime dependency
evidence             != authorization
roadmap              != work-package authorization
```

## 10. Dependency-Admission-Gate

**Status: `IN FORCE` — in Kraft gesetzt durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung, Disposition `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`; Gate-A-Punkt `A-9` **erfüllt** (§16.1). Zugelassene Drittabhängigkeiten: 0. Durch `A-9` konkret zugelassene Abhängigkeiten: keine. Keine wird durch dieses Dokument zugelassen.**

Der zuvor **vorbereitete** Abschnitt ist damit von *nicht in Kraft* auf **verbindlich** gestellt. Die Grundunterscheidung in §10.1, die zehn Kriterien in §10.2 und die Zulassungsautorität bleiben inhaltlich **unverändert** und gelten ab jetzt **bindend**; es entsteht **kein** elftes Kriterium, **kein** neues Kennungsschema und **kein** Dependency-Admission-Register.

```text
Policy in Kraft         != Abhängigkeit zugelassen
Vorschlag               != Zulassung
Bedarf erkannt          != Abhängigkeit zugelassen
WP autorisiert          != Abhängigkeit zugelassen
Abhängigkeit zugelassen != Abhängigkeit installiert
null Drittabhängigkeiten != null Lieferkette
```

Das Gate **erweitert** die [Sovereignty and Dependency Policy](../architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) §6 (deren zehn Zulassungskriterien unverändert gelten) um die Unterscheidung, die vor dem ersten produktiven Commit gebraucht wird.

### 10.1 Grundunterscheidung

| Klasse | Definition | Behandlung |
| ------ | ---------- | ---------- |
| **Standardbibliothek / Plattformfähigkeit** | Bestandteil der gewählten Sprach-/Runtime-Distribution selbst; kein separater Bezug, keine separate Versionierung, keine eigene Lieferkette | **kein** Admission-Verfahren; die Zulassung erfolgte implizit mit der Sprach-/Runtime-Entscheidung, deren Lieferkette gesondert zu bewerten ist |
| **Drittabhängigkeit** | jedes separat bezogene Paket, Modul, jede Bibliothek, jedes Werkzeug oder Artefakt mit eigener Herkunft, Versionierung und Lieferkette — **auch** wenn es klein, populär oder scheinbar trivial ist | **vollständiges** Admission-Verfahren je Abhängigkeit |

```text
popular           != vetted
small             != harmless
transitive        != out of scope
build-only        != outside the supply chain
test-only         != automatically trusted
vendored          != first-party
golang.org-Präfix != Standardbibliothek
```

Transitive und reine Build-/Test-Abhängigkeiten unterliegen demselben Verfahren; sie sind Teil der Lieferkette.

**Verbindliche Anwendung auf die gewählte Sprach-/Runtime-Klasse (Go).** Die folgende Zuordnung ist mit der Inkraftsetzung des Gates **verbindlich**. Sie ändert die Klassenlogik oben **nicht**, sondern macht sie für Go ausdrücklich:

| Gegenstand | Einordnung |
| ---------- | ---------- |
| Paket der Standardbibliothek der gewählten Go-Distribution | **keine** individuelle Drittzulassung — Bestandteil der gewählten Sprach-/Runtime-Distribution, **nicht** separat bezogene Software |
| Plattformfähigkeit, die **ohne** separaten Bezug zusätzlicher Software genutzt wird | **keine** individuelle Drittzulassung; eine **separat bezogene** native Bibliothek oder ein separat bezogenes Plattformwerkzeug erbt diese Ausnahme **nicht** |
| Fähigkeit des **beobachteten** Hosts | Beobachtungsfläche beziehungsweise Ziel — **keine** CoreOps-Softwareabhängigkeit; die Zielautorisierung bleibt `P-1`/Gate B (§5, §16.2) |
| jedes separat bezogene Go-Modul, -Paket, jede Bibliothek, jedes Werkzeug | **Drittabhängigkeit** — vollständiges Verfahren je Abhängigkeit |
| `golang.org/x/...` | **Drittabhängigkeit**, es sei denn, die betreffende Software ist tatsächlich Bestandteil der Standardbibliothek der gewählten Go-Distribution; die Reputation eines Importpfads ändert die Herkunftsklassifikation **nicht** |
| transitive Abhängigkeiten | **im Scope**; eine direkte Abhängigkeit darf **nicht** zugelassen werden, solange ihre aufgelöste transitive Schließung unbekannt oder unbegrenzt ist. Die Zulassung muss die aufgelöste Schließung so weit abdecken, dass Herkunft, Pinning, Integrität, Lizenzierung, Schwachstellenbehandlung, Offline-Verfügbarkeit und Exit-Anforderung erfüllt sind |
| reine Build-Abhängigkeiten | **im Scope** — `build-only != außerhalb der Lieferkette` |
| reine Test-Abhängigkeiten | **im Scope**; die Testeinrichtungen der Go-Standardbibliothek benötigen **keine** individuelle Drittzulassung, ein **separat bezogenes** Testframework, -werkzeug oder eine -bibliothek dagegen **schon** |
| separat bezogene Linter, Analyzer, Generatoren, Codegenerierungspakete, Build-Helfer und Entwicklerwerkzeuge, die governte Quell-, Build- oder Testeingänge beeinflussen | **im Scope** als Lieferketteneingang; ob sie zur Laufzeit ausgeliefert werden, kann die Risikobewertung beeinflussen, **nicht** aber die Prüfpflicht |
| vendorierte Drittsoftware | bleibt **Drittsoftware**; Vendoring ist ausschließlich ein Offline-Verfügbarkeits- und Auflösungsmechanismus und ändert die **Herkunft** nicht |

**Keine Lieferkettenklärung durch die Klassenzuordnung.** Dass Standardbibliothek und ohne separaten Softwarebezug genutzte Plattformfähigkeit **keine** individuelle Drittzulassung benötigen, bedeutet ausdrücklich **nicht**, dass ihre Lieferkette geklärt wäre. Go-Distribution, Standardbibliothek und Toolchain bleiben Gegenstand von `A-13` und nachgelagerter Governance; `A-9` löst sie **nicht** auf.

```text
Standardbibliothek       != Lieferkette geklärt
null Drittabhängigkeiten != null Lieferkette
Importpfad-Reputation    != Herkunft
vendored                 != first-party
beobachteter Host        != CoreOps-Softwareabhängigkeit
```

### 10.2 Admission-Kriterien je Drittabhängigkeit

| # | Kriterium | Anforderung |
| - | --------- | ----------- |
| 1 | **Notwendigkeit** | Warum ist die Standardbibliothek nicht ausreichend? Was ist der konkrete, benannte Bedarf? Bequemlichkeit ist kein Bedarf |
| 2 | **Herkunft** | Herausgeber, Repository, Bezugsweg, Verantwortlichkeit — identifizierbar und dokumentiert |
| 3 | **Pinning** | exakte Version, nicht Bereich; Auflösung an Prüfsumme oder Digest gebunden |
| 4 | **Lizenzierung** | Lizenz benannt, mit dem CoreOps-Veröffentlichungsmodell (§13.3–§13.4) vereinbar, Weitergabepflichten benannt |
| 5 | **Offline-Verfügbarkeit** | über kontrolliertes lokales Repository oder Vendoring beziehbar; eine zwingend online auflösende Abhängigkeit ist für Kernfunktionen unzulässig |
| 6 | **Schwachstellenbehandlung** | benannter Prozess: Beobachtung, Bewertung, Reaktionsweg, Verantwortlicher |
| 7 | **Integrität** | Prüfsumme oder Digest verifizierbar; Signatur/Provenance sofern verfügbar; SBOM-Erfassung nach [Artifact Identity, Provenance and SBOM Model](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md) |
| 8 | **Reproduzierbarkeit** | derselbe Eingang erzeugt denselben Auflösungsstand; Auflösung ist nicht zeit- oder netzwerkabhängig |
| 9 | **Update-Policy** | wer aktualisiert wann, nach welcher Prüfung, mit welchem Rückfallweg |
| 10 | **Ersetzungs-/Entfernungskosten** | Exit-Strategie nach Sovereignty Policy §9: wie teuer ist der Ausbau bei Lizenz-, Sicherheits- oder Verfügbarkeitsproblem |

**Zulassung erfolgt ausschließlich durch eine ausdrückliche Human-Maintainer-Entscheidung, je Abhängigkeit, dokumentiert.** Es gibt keine pauschale Zulassung einer Ökosystem-Familie und keine stille Zulassung durch Verwendung.

Diese zehn Kriterien sind mit der Inkraftsetzung des Gates **bindend**. Sie werden **nicht** ersetzt, **nicht** erweitert und **nicht** umnummeriert; ein **elftes** Kriterium entsteht **nicht**. Zusätzlich — und unverändert — gilt die Autoritätsanforderung: **ausdrückliche Human-Maintainer-Zulassung**.

**Zulassungsautorität (verbindlich).** Die Zulassungsautorität liegt **ausschließlich** beim Human Maintainer. **Kein** KI-Agent, **kein** Implementierungsagent, **kein** Paketmanager, **kein** Build-Werkzeug, **keine** Quelldatei, **kein** Importstatement, **kein** Modulmanifest, **kein** Lockfile und **keine** Vendoring-Aktion kann eine Abhängigkeit zulassen. Die Zulassung einer konkreten Abhängigkeit setzt **kumulativ** voraus:

1. eine **ausdrücklich autorisierte Work-Package-Grenze**, deren freigegebener Scope die konkrete Dependency-Admission-Tätigkeit ausdrücklich trägt — entweder als eigens dafür bestimmte Zulassungsarbeit oder als anderes autorisiertes implementierungs-/buildorientiertes Work Package, dessen freigegebener Scope diese konkrete Zulassung ausdrücklich enthält; **und**
2. eine **ausdrückliche Human-Maintainer-Zulassung je konkreter Abhängigkeit**.

Eine Work-Package-Autorisierung allein lässt **nichts** zu; die bloße Nennung einer Abhängigkeit in einem Work Package ebenfalls nicht. Dieses Dokument erzeugt **kein** solches Work Package.

```text
WP autorisiert               != Abhängigkeit zugelassen
Abhängigkeit im WP genannt  != Abhängigkeit zugelassen
KI-Empfehlung                != Zulassung
Paketmanager-Auflösung       != Zulassung
Importstatement              != Zulassung
go.mod-Eintrag               != Zulassung
go.sum-Eintrag               != Zulassung
Vendoring                    != Zulassung
```

**Zulassungsgranularität.** Die Zulassung erfolgt **je konkreter Abhängigkeit** und ist mindestens gebunden an: erklärte Notwendigkeit beziehungsweise Verwendungszweck · exakte Herkunfts-/Quellidentität · exakte Version oder unveränderliche Revision · Prüfsumme/Digest · aufgelöste transitive Schließung · anwendbaren Lizenzstand. Es gibt **keine** Blankettzulassung eines Ökosystems, einer Organisation, eines Herausgebers, einer Importpfadfamilie oder einer Abhängigkeitskategorie und **keine** stille Zulassung durch Verwendung.

**Versions- und Update-Bindung.** Die Zulassung ist an den zugelassenen Abhängigkeitsstand gebunden. Eine Änderung an Version, unveränderlicher Revision, Digest/Prüfsumme, Herkunfts-/Quellidentität, maßgeblich relevantem Lizenzstand oder aufgelöster transitiver Schließung verlangt eine **erneuerte beziehungsweise überarbeitete** Human-Maintainer-Zulassung. Eine Update-Policy (Kriterium 9) beschreibt, **wie** Updates bewertet werden, **wann** die Prüfung erfolgt, **welche** Prüfungen erwartet werden und **welcher** Rückfall-/Rollback-Weg gilt — sie autorisiert **keinen** künftigen Abhängigkeitsstand.

```text
Version X zugelassen       != Version Y zugelassen
Update-Policy dokumentiert != Update autorisiert
neuer Digest               != alte Integritätszulassung gilt fort
```

**Offline-/Souveränitätsanforderung.** Jede künftig zugelassene Drittabhängigkeit muss für den governten CoreOps-Kernpfad **kontrolliert und reproduzierbar offline auflösbar** sein; eine zwingend online auflösende Abhängigkeit ist für Kernfunktionen **unzulässig**. Erforderlich sind kontrollierte lokale Verfügbarkeit, reproduzierbare Auflösung, Provenance-/Integritätsprüfung, SBOM-Behandlung, Schwachstellenbehandlung sowie eine Exit-/Ersetzungsstrategie. Der **konkrete Mechanismus** ist durch `A-9` **nicht** ausgewählt: weder Go-Proxy noch Mirror, lokales Repository, CorePack-Realisierung, Vendor-Layout noch Transferformat; diese bleiben, soweit einschlägig, `A-13` beziehungsweise nachgelagerter Governance vorbehalten.

**Lizenzierung und `A-10`.** Die `A-9`-Policy konnte **in Kraft** sein, während `A-10` noch **offen** war; eine **erste konkrete** Drittzulassung war in diesem Zustand **nicht** abschließbar, weil die anwendbare Veröffentlichungs-/Lizenzdisposition zur Bewertung des Lizenzkriteriums (Nr. 4) nicht ausreichte. Dieser Engpass ist **behoben**: `A-10` ist durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung **erfüllt** (§13, §16.1), und das CoreOps-Outbound-Veröffentlichungs-/Lizenzmodell ist **definiert** — **PUBLIC / OPEN SOURCE** unter **Apache-2.0** (§13.3, §13.4). Ein künftiger **konkreter** Abhängigkeitskandidat kann damit hinsichtlich Kriterium Nr. 4 grundsätzlich **bewertbar** werden. Das ist ausschließlich eine Aussage über die **Bewertbarkeit**, nicht über eine Zulassung: durch diesen Abschnitt wird **keine** konkrete Abhängigkeit benannt, **keine** bewertet und **keine** zugelassen; zugelassene Drittabhängigkeiten bleiben **0**, und die Autoritäts-, Granularitäts- und Bindungsanforderungen dieses Abschnitts gelten unverändert vollständig fort.

```text
A-10 erfüllt              != Abhängigkeit zugelassen
Apache-2.0 ausgewählt     != Kandidat kompatibel
Kandidat kompatibel       != Kandidat zugelassen
Lizenzkompatibilität bewertet != Abhängigkeit zugelassen
Lizenzkriterium bewertbar != Lizenzkriterium bewertet
```

**Verhältnis zu `P-3` Option B.** `P-3` Option B bleibt **ausschließlich ergänzend** (§7.5); `A-9` erweitert `P-3` **nicht** und hebt die Decision Ceiling **nicht** an. Lässt sich Option B mit Standardbibliotheks- beziehungsweise Plattformfähigkeit **ohne** separat bezogene Software realisieren, ist **keine** individuelle `A-9`-Drittzulassung erforderlich. Wird dafür ein **separat bezogenes** Modul, eine Bibliothek oder ein Werkzeug vorgeschlagen, ist die **vollständige** `A-9`-Zulassung erforderlich — zusätzlich zu der unverändert eigenständig zu erfüllenden Anwendbarkeitsbedingung von Option B.

```text
Option B zulässig       != Abhängigkeit zugelassen
Abhängigkeit zugelassen != Option-B-Bedingung erfüllt
A-9 in Kraft            != P-3-Entscheidungsdecke erweitert
```

### 10.3 Zielzustand für den ersten Slice

Für den ersten Observe-Slice gilt der **Zero-Third-Party-Default**: Drittabhängigkeiten werden **nicht** verwendet, solange die konkrete Abhängigkeit nicht das vollständige Admission-Verfahren durchlaufen und eine ausdrückliche Human-Maintainer-Zulassung erhalten hat. Bequemlichkeit ist **keine** hinreichende Begründung. Eine Drittabhängigkeit darf überhaupt nur dann vorgeschlagen werden, wenn **beide** Bedingungen zutreffen: (1) es besteht ein konkreter, benannter Bedarf **und** (2) die Standardbibliotheks- beziehungsweise Plattformfähigkeit reicht für genau diesen Bedarf nicht aus. Auch dann gilt: `Vorschlag != Zulassung` und `Bedarf erkannt != Abhängigkeit zugelassen`.

Der entschiedene Zuschnitt (§7.5 Option A primär / Option B ausschließlich ergänzend, §8.4 Go) kommt plausibel **ohne** Drittabhängigkeit aus. Eine leere Abhängigkeitsliste ist für den ersten produktiven Commit ein **erklärtes und bevorzugtes Ziel**, keine bloße Möglichkeit — aber ausdrücklich **keine** Zusicherung: `plausibel != garantiert`. Durch die Sprach-/Runtime-Entscheidung wird **keine** Abhängigkeit zugelassen; **0** Abhängigkeiten sind zugelassen und es ist **keine** konkrete Standardbibliotheks-API ausgewählt. Auch durch die **Source-Tree-Entscheidung** (§9.2) wird **keine** Abhängigkeit zugelassen: sie benennt keine, setzt keine voraus, legt **kein** `go.mod` und **kein** `go.sum` an und erzeugt **keinen** Vendor-Baum. Sollte ein späterer, gesondert autorisierter Modul-/Dependency-/Toolchain-Stand einen Vendor-Baum vorsehen, wäre dessen Ort eine bloße Folge der Strukturwahl (Repository-Wurzelverzeichnis) und **keine** Entscheidung darüber, dass er entsteht. Auch die **Inkraftsetzung des Gates selbst** lässt **keine** Abhängigkeit zu: `A-9` ist **erfüllt**, die zugelassene Menge bleibt **leer**, und es wurde **keine** konkrete Abhängigkeit bewertet — es existiert **keine** Herkunfts-, Versions-, Digest-, Lizenz-, Schwachstellen-, SBOM- oder Transitivschließungsbewertung, weil **keine** konkrete Abhängigkeit zugelassen wird.

```text
A-9 erfüllt != Abhängigkeit zugelassen
A-9 erfüllt != Abhängigkeit installiert
A-9 erfüllt != Abhängigkeit aufgelöst
A-9 erfüllt != Drittpaket ausgewählt
A-9 erfüllt != Testframework ausgewählt
```

## 11. Package-Manager- und Lockfile-Implikationen

Der [Foundation Scope Lock](FOUNDATION_SCOPE_LOCK.md) führte *Lockfiles von Package Managern, Dependencies* als *Forbidden Implementation Type* der **abgeschlossenen** Foundation-Phase. Mit `HM-F1` ist diese Phase geschlossen; die Regel bleibt für sie gültig und wird hier **nicht** aufgehoben. Für Post-Foundation-Arbeit gilt:

- Ein Lockfile ist ein **Lieferketten-Governance-Artefakt**, kein Nebenprodukt. Seine erstmalige Erzeugung ist ein eigenes Autoritätsereignis, kein Implementierungsdetail.
- Ein Lockfile darf erst entstehen, **nachdem** Sprache/Runtime entschieden und die Dependency-Admission-Policy in Kraft ist — sonst dokumentiert es Zulassungen, die nie erteilt wurden.
- Ein Lockfile mit **null** Drittabhängigkeiten ist ein zulässiges und erwünschtes Ergebnis; er belegt die leere Lieferkette ausdrücklich.
- Jede Lockfile-Änderung ist reviewpflichtig wie eine Codeänderung: sie kann Herkunft, Version und Integrität stillschweigend verschieben.
- Der Auflösungsvorgang muss **offline** reproduzierbar sein (Vendoring oder kontrolliertes lokales Repository).
- Ein Lockfile belegt **Auflösungsstand**, nicht Sicherheit.

```text
lockfile present    != dependencies vetted
lockfile present    != build reproducible
resolution pinned   != artifact integrity verified
zero dependencies   != zero supply chain (toolchain remains in scope)
```

**Derzeit existiert kein Package Manifest, kein Lockfile und kein Package Manager im Repository. Keines wird durch dieses Dokument autorisiert.**

**Verhältnis zur Source-Tree-Entscheidung (§9.2).** `A-8` entscheidet ausschließlich den **künftigen Ort** eines Modulmanifests als Folge der Strukturwahl — **nicht** dessen Erzeugung und **nicht** dessen Inhalt. `go.mod` wäre am Repository-Wurzelverzeichnis zu erwarten, ist aber **nicht angelegt** und **nicht autorisiert**; Modulpfad, Go-Version, Go-Distribution und Toolchain-Version sind **nicht entschieden**. Ein `go.sum` ist durch `A-8` **nicht gefordert**, **nicht angelegt** und **nicht autorisiert**; es wäre — ebenfalls am Repository-Wurzelverzeichnis — ausschließlich dann zu erwarten, **falls** eine solche Datei später durch einen autorisierten Modul-, Dependency- oder Toolchain-Stand rechtmäßig erzeugt wird. Es ist damit **weder zugesichert noch notwendigerweise zu erwarten**.

**Verhältnis zur Dependency-Admission-Entscheidung (§10).** Die oben genannte Bedingung — ein Lockfile darf erst entstehen, **nachdem** Sprache/Runtime entschieden **und** die Dependency-Admission-Policy in Kraft ist — ist mit `A-9` nun **erfüllt**. Das ist ausschließlich eine **erfüllte Vorbedingung** für spätere, gesondert autorisierte dependency-governte Modul- und Build-Arbeit und **keine** Artefaktautorisierung: `go.mod`, `go.sum`, ein Lockfile, ein Package-Manifest und ein Vendor-Baum bleiben **nicht angelegt** und **nicht autorisiert**; Modulpfad, Go-Version, Go-Distribution und Toolchain-Version bleiben **nicht entschieden**; **0** Abhängigkeiten sind zugelassen. Ein Lockfile mit **null** Drittabhängigkeiten bliebe ein zulässiges und erwünschtes künftiges Ergebnis — seine Erzeugung ist damit weiterhin **nicht** autorisiert.

```text
Ort benannt         != Datei erzeugt
go.mod-Ort entschieden != go.mod autorisiert
go.sum bedingt möglich != go.sum erwartet
manifest absent      = the current state
Vorbedingung erfüllt != Artefakt autorisiert
A-9 erfüllt          != go.mod autorisiert
A-9 erfüllt          != go.sum autorisiert
A-9 erfüllt          != Vendor-Baum autorisiert
A-9 erfüllt          != Modulpfad entschieden
A-9 erfüllt          != Go-Version entschieden
A-9 erfüllt          != Go-Distribution entschieden
A-9 erfüllt          != Toolchain ausgewählt
A-9 erfüllt          != Build autorisiert
```

## 12. Build- und Packaging-Implikationen

Vorbereitend, **nicht** entschieden:

| Thema | Anforderung an eine spätere Entscheidung |
| ----- | ---------------------------------------- |
| **Reproduzierbarkeit** | derselbe Quellstand erzeugt ein identifizierbar gleiches Artefakt; Build-Eingänge sind deklariert |
| **Offline-Build** | der Build muss ohne Netzwerkzugriff durchführbar sein (Sovereignty Policy §7) |
| **Artefaktidentität** | Artefakt-Identität, Provenance und SBOM nach dem bestehenden [Artifact Identity, Provenance and SBOM Model](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md); `SBOM vorhanden != vollständig != sicher` |
| **Toolchain** | die Build-Toolchain ist selbst Teil der Lieferkette und unterliegt Herkunfts-, Pinning- und Integritätsanforderungen |
| **Container** | `HM-2` setzt Docker-first als **Delivery Baseline** — `Docker-first != Docker-only != zwingende Runtime-Abhängigkeit`. Ein Containerartefakt ist eine **Auslieferungsform**, keine Architekturvorgabe und keine Voraussetzung des Slice |
| **CI** | das Manifest führt `ci: recommended` mit `Status: pending`. Eine CI-Definition ist **nicht** autorisiert und **nicht** Voraussetzung des ersten produktiven Commits |

```text
build succeeded     != artifact verified
artifact built      != artifact released
container image     != product release
CI configured       != quality established
```

**Derzeit existiert keine Build-Definition, kein Containerartefakt und kein CI-Workflow. Keines wird durch dieses Dokument autorisiert.**

**Verhältnis zur Source-Tree-Entscheidung (§9.2).** Die Source-Tree-Entscheidung liefert für `A-13` ausschließlich einen **Eingang** — den künftigen Ort eines Modulmanifests — und **keine** Disposition. Reproduzierbarkeit, Offline-Build, Artefaktidentität/SBOM, Build-Toolchain, Go-Version, Distribution, Toolchain-Version, Containerform, Artefakt- beziehungsweise Binärname, Paketformat und CI bleiben sämtlich **unentschieden**; `A-13` bleibt **offen**. Ein Verzeichnisname ist **kein** Artefaktname.

```text
A-8 satisfied     != A-13 satisfied
go.mod-Ort        != Go-Version entschieden
go.mod-Ort        != Toolchain entschieden
Verzeichnisname   != Artefaktname
build definition absent = the current state
```

**Verhältnis zur Dependency-Admission-Entscheidung (§10).** `A-9` regelt die **Zulassung von Abhängigkeiten** — und sonst nichts. Die nachgelagerte konkrete Modul-, Build-, Toolchain-, Reproduzierbarkeits- und Packaging-Disposition bleibt Gegenstand von `A-13` und ist unverändert **offen**. `A-9` entscheidet insbesondere **nicht** über Go-Version, Go-Distribution, Toolchain-Version, Build-Werkzeug, Paketformat, Containerform, Artefakt- oder Binärnamen und CI und wählt auch **keinen** Auflösungs-, Proxy-, Mirror- oder Vendor-Mechanismus aus.

```text
A-9 satisfied     != A-13 satisfied
A-9 satisfied     != Go-Distribution ausgewählt
A-9 satisfied     != Toolchain ausgewählt
A-9 satisfied     != build/packaging entschieden
zero dependencies != zero supply chain
```

## 13. `NEW-8` — `README.md` und `LICENSE`

`NEW-8` war der in `CO-WP-031` **deferred** deklarierte Restpunkt: im öffentlichen Repository existierte damals weder `README.md` noch `LICENSE` (**historischer Stand**; beide sind inzwischen durch die gesondert autorisierte Publikations-Realisierung **erstellt und vorhanden**). `CO-WP-031` hat ihn ausdrücklich nicht behoben und kein Risiko registriert. Dieser Abschnitt lieferte ursprünglich die verlangte **Dispositionsempfehlung** — und **nur** diese; die Empfehlungen sind historisch erhalten.

**`NEW-8` ist inzwischen `DECIDED`.** Durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung — Disposition `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS` — ist der Punkt dispositioniert; Gate-A-Punkt `A-10` steht damit auf **erfüllt** (§16.1). Entschieden ist:

| Gegenstand | Entschiedener Stand |
| ---------- | ------------------- |
| `README.md` | **SHALL EXIST** — zu erstellen und remote zu integrieren in einer **eigenen, ausdrücklich autorisierten** Repository-Änderung **vor** dem ersten produktiven Source-Commit. **Current State: ERSTELLT / VORHANDEN** durch die gesondert autorisierte Publikations-Realisierung — **nicht** durch `A-10` selbst |
| `README`-Autorität | öffentliche Zusammenfassung **ohne** normative Autorität; überschreibt **keine** autoritative Governance-Quelle; Rang 19 der Source-of-Truth-Hierarchie unverändert |
| `README`-Current-State-Politik | duplizierte quantitative Spiegel **minimieren**; autoritative Verweise bevorzugen; jede quantitative Momentaufnahme ausdrücklich **aktuell, gepflegt und nicht-autoritativ** |
| `LICENSE` | **SHALL EXIST** — standardmäßiger, **unveränderter** Apache-2.0-Lizenztext, erst bei späterer, gesondert autorisierter Erstellung. **Current State: ERSTELLT / VORHANDEN** — die gesondert autorisierte Publikations-Realisierung hat die Datei mit dem vollständigen, unveränderten Apache-2.0-Text angelegt; **nicht** durch `A-10` selbst |
| Veröffentlichungsrechte-Modus | **PUBLIC / OPEN SOURCE** |
| Outbound-Lizenz | **Apache-2.0** |
| Beitragsprogramm | **NICHT AKTIV** · `CONTRIBUTING.md` **nicht autorisiert / nicht erstellt** · CLA **keine** · DCO **keiner** |
| Marke / Projektname | urheberrechtliche Rechteeinräumung **≠** Namens-/Markenrechtseinräumung; **keine** eingetragene Marke beansprucht; gesonderte Markenrichtlinie **zurückgestellt** |
| Gewährleistung / Support / Sicherheit | Behauptungsverbote bleiben in Kraft; `SECURITY.md` **nicht autorisiert / nicht durch `A-10` erstellt** |
| `NOTICE` | **nicht erstellt / nicht autorisiert** |

**Weder `README.md` noch `LICENSE` wird durch diese Entscheidung erstellt oder zur Erstellung autorisiert** (§13.5) — das bleibt unverändert zutreffend. Die Erstellung ist erst später durch eine davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung erfolgt; beide Dateien sind seitdem **vorhanden**. Der Apache-2.0-Lizenztext wird in diesem Dokument **nicht** wiedergegeben, **nicht** entworfen und **nicht** verändert.

```text
NEW-8 entschieden        != README erstellt
NEW-8 entschieden        != LICENSE erstellt
NEW-8 entschieden        != NOTICE erstellt
NEW-8 entschieden        != CONTRIBUTING erstellt
NEW-8 entschieden        != SECURITY.md erstellt
Apache-2.0 ausgewählt    != Apache-2.0-Datei erstellt
A-10 erfüllt             != Gate A passiert
```

### 13.1 Soll `README.md` vor dem ersten produktiven Source-Commit existieren?

```text
DECIDED:   YES — README SHALL EXIST
STATUS:    DECIDED — ausdrückliche Human-Maintainer-Entscheidung (`A-10`)
TIMING:    eigene, ausdrücklich autorisierte Repository-Änderung
           VOR dem ersten produktiven Source-Commit
CREATION:  NICHT durch diese Entscheidung autorisiert
CURRENT:   README.md ERSTELLT / VORHANDEN — durch die davon getrennte,
           gesondert autorisierte Publikations-Realisierung
```

*Historisch:* Diese Zeile stand ursprünglich als `RECOMMEND: YES` / `STATUS: PROPOSED / UNACCEPTED`; die folgende Begründung ist die ursprüngliche und bleibt als tragende Erwägung erhalten.

**Begründung.** Ab dem ersten produktiven Source-Commit ändert sich, was ein Betrachter des öffentlichen Repositories vernünftigerweise annimmt: aus einem erkennbaren Dokumentationskorpus wird ein Repository mit Code. Ohne Einstiegspunkt ist die naheliegende Fehlannahme „hier gibt es ein benutzbares Produkt" — genau die Behauptung, die CoreOps durchgängig vermeidet. Der `README` ist hier ein **Ehrlichkeits-Instrument**, nicht Marketing. Die [Source-of-Truth-Hierarchie](REPOSITORY_GOVERNANCE_STANDARD.md) §7 führt `README and public summaries` auf Rang 19 — bewusst **unter** allen normativen Quellen: ein `README` fasst zusammen, er darf nichts überschreiben.

### 13.2 Welchen Mindestinhalt sollte er tragen?

| Element | Anforderung |
| ------- | ----------- |
| Was CoreOps ist | Projektname, Slogan, eine ehrliche Kurzbeschreibung |
| **Ehrlicher Status** | Foundation 0.1 geschlossen · Observe betreten mit Grenze · **kein** funktionaler Produktrelease · **kein** benutzbares Produkt · **0** implementierte Runtime-Capabilities · **0** `supported` Integrationen · **0** akzeptierte ADRs |
| Was das Repository **enthält** | Dokumentation, Governance, Register — und ab dem ersten produktiven Commit: welcher eng begrenzte Slice |
| Was es **nicht** enthält | Installation, Betrieb, Support, Zusicherung, Zeitplan |
| Einstiegspunkte | Verweise auf Project Brief, Scope Lock, Release-Taxonomie, Work Package Queue |
| Lizenz-/Nutzungshinweis | Verweis auf den tatsächlichen Lizenzstand — **ohne** Rechte zu behaupten, die nicht eingeräumt sind |
| Governance-Hinweis | Human-Maintainer-Vorbehalt; keine Beitragszusage ohne eigenen Prozess |
| Sprache | DE oder EN mit klarem Sprachstatus; **keine** Paritätsbehauptung ohne geprüfte Übersetzung |

**Nicht** in den `README`: Reife-, Sicherheits-, Zertifizierungs-, Compliance- oder Supportbehauptungen; Zeitpläne; Fertigstellungsversprechen; `v1.0`-Andeutungen; private Daten. Mit der `A-10`-Entscheidung ausdrücklich ergänzt: ein künftiger `README` darf Produktionsreife, Sicherheits-, Zertifizierungs- oder Compliance-Zertifizierung, Supportzusage, SLA, Verfügbarkeits-, Wartungs- oder Kompatibilitätsgarantie, Installationsreife und Releasezeitplan **weder aussagen noch nahelegen**.

**Quantitative Angaben — Spiegel, nicht Autorität (`A-10`).** Der `README` ist bewusst eine **niedrigrangige öffentliche Zusammenfassung**; die Source-of-Truth-Hierarchie bleibt unverändert bindend, und `A-10` ändert sie **nicht**. Daraus folgt für den künftigen `README`:

- duplizierte, hart kodierte Zähler sind zu **minimieren**;
- **Verweise** auf die autoritativen Current-State-Quellen sind, wo praktikabel, vorzuziehen;
- eine hart kodierte Zahl ist nur zulässig, wenn sie ausdrücklich **aktuell**, **gepflegt** und **nicht-autoritativ** ausgewiesen ist.

```text
README-Zahl        != autoritative Zahl
README-Status      != normative Statusquelle
README vorhanden   != Produkt benutzbar
LICENSE vorhanden  != Produkt unterstützt
```

Dieses Dokument entwirft **keine** `README`-Prosa und legt **keinen** `README`-Inhalt fest; der konkrete öffentliche Inhalt bleibt gesondert zu autorisieren (§13.5).

### 13.3 Muss die Lizenz-Disposition vor der Veröffentlichung produktiven Quellcodes erfolgen?

```text
DECIDED:   YES — Lizenz-Disposition liegt vor
STATUS:    DECIDED — ausdrückliche Human-Maintainer-Entscheidung (`A-10`)
MODE:      PUBLIC / OPEN SOURCE
LICENSE:   Apache-2.0 (standardmäßiger, unveränderter Lizenztext)
CREATION:  Erstellung NICHT durch diese Entscheidung autorisiert
CURRENT:   LICENSE ERSTELLT / VORHANDEN — durch die davon getrennte,
           gesondert autorisierte Publikations-Realisierung;
           standardmäßiger, unveränderter Apache-2.0-Text
```

*Historisch:* Diese Zeile stand ursprünglich als `RECOMMEND: YES` / `STATUS: PROPOSED / UNACCEPTED`; die folgende Begründung ist die ursprüngliche und bleibt als tragende Erwägung erhalten. Die dort geforderte Disposition ist inzwischen **getroffen**: das künftige Veröffentlichungsrechte-Modell für produktiven CoreOps-Quellcode ist **PUBLIC / OPEN SOURCE**, die exakte künftige Outbound-Lizenz ist **Apache-2.0**. Die spätere `LICENSE`-Datei hat den **standardmäßigen, unveränderten** Apache-2.0-Text zu tragen; dieser wird hier **nicht** wiedergegeben und **nicht** verändert, und **kein** Rechteinhaber-Wortlaut, **kein** Copyright-Jahr, **keine** jurisdiktionsspezifische Rechtsfolge, **keine** Durchsetzbarkeits- und **keine** Patentaussage wird hier getroffen.

**Begründung.** Solange nur Dokumentation im öffentlichen Repository liegt, ist die fehlende Lizenz ein erklärter, dispositionierter Restpunkt. Mit veröffentlichtem produktivem Quellcode wird sie materiell: ein öffentliches Repository ohne Lizenz räumt Dritten **keine** Nutzungs-, Änderungs- oder Weitergaberechte ein — sichtbar ist nicht dasselbe wie nutzbar. Die Disposition muss also **vor** der Veröffentlichung getroffen werden, damit der veröffentlichte Zustand das ausdrückt, was tatsächlich gewollt ist. Eine bewusste Entscheidung, **keine** Lizenz zu vergeben, ist dabei eine vollwertige Disposition — sie muss nur getroffen und sichtbar sein statt zu unterbleiben.

> Dies ist eine **Governance-Empfehlung, keine Rechtsberatung.** Die Wahl eines Lizenzmodells hat rechtliche Wirkung; der Human Maintainer entscheidet sie eigenständig und zieht dafür bei Bedarf fachkundige Beratung hinzu.

### 13.4 Abgrenzung: privat/intern · Weitergabe · öffentliche Veröffentlichung

| Modus | Bedeutung | Lizenzbedarf |
| ----- | --------- | ------------ |
| **Privat / intern** | Nutzung ausschließlich durch den Rechteinhaber und Personen unter seiner Kontrolle; keine Weitergabe an Dritte | Lizenz für die eigene Nutzung nicht erforderlich; die Frage bleibt offen, nicht beantwortet |
| **Weitergabe (Redistribution)** | Dritte erhalten das Artefakt — auch bei Weitergabe an einen einzelnen Empfänger, auch ohne Entgelt | Eine ausdrückliche Rechteeinräumung ist erforderlich; ohne sie erhält der Empfänger keine Nutzungsrechte |
| **Öffentliche Veröffentlichung / Open Source** | jedermann kann beziehen, nutzen, ändern und weitergeben | Erfordert eine ausdrückliche Open-Source-Lizenz **und** eine Entscheidung über Beitragsannahme, Marken, Namensnennung und Haftungsausschluss |

```text
public repository != open source
public repository != open source by visibility alone
publicly visible  != publicly licensed
no license        != permission granted
code published    != redistribution right granted
```

**Entschiedener Modus (`A-10`).** Der für die **künftige** Veröffentlichung produktiven CoreOps-Quellcodes vorgesehene Modus ist die dritte Zeile dieser Tabelle: **öffentliche Veröffentlichung / Open Source** unter **Apache-2.0**. Die Unterscheidung dieser Tabelle bleibt dabei vollständig bindend — insbesondere gilt weiterhin, dass Sichtbarkeit für sich **keine** Rechteeinräumung ist: Open Source entsteht erst mit der tatsächlich vorhandenen, wirksamen Lizenz, **nicht** durch die Entscheidung und **nicht** durch die öffentliche Sichtbarkeit des Repositories. *Stand zum Zeitpunkt der `A-10`-Entscheidung (historisch):* damals existierte **keine** `LICENSE`-Datei; der tatsächliche Lizenzstand des Repositories war dadurch **unverändert**, und Dritten waren **keine** Nutzungs-, Änderungs- oder Weitergaberechte eingeräumt. **Current State:** die `LICENSE`-Datei ist durch die gesondert autorisierte Publikations-Realisierung **vorhanden** und trägt den standardmäßigen, unveränderten Apache-2.0-Text. Dieses Dokument trifft dazu **keine** Aussage über Durchsetzbarkeit, Patentwirkung oder jurisdiktionsspezifische Rechtsfolgen.

Die von dieser Tabelle für den Open-Source-Modus zusätzlich geforderten Dispositionen sind wie folgt entschieden beziehungsweise ausdrücklich **nicht** entschieden:

| Zusätzlich geforderte Disposition | Stand |
| --------------------------------- | ----- |
| Open-Source-Lizenz | **entschieden** — Apache-2.0, standardmäßiger unveränderter Text; durch `A-10` selbst **keine** Datei erstellt, inzwischen durch die gesondert autorisierte Publikations-Realisierung **erstellt und vorhanden** |
| Beitragsannahme | **entschieden: kein öffentliches Beitragsprogramm** — `NICHT AKTIV`; `CONTRIBUTING.md` **nicht autorisiert / nicht erstellt**; CLA **keine**; DCO **keiner**; `A-10` disponiert **keinen** Beitragsworkflow und **keine** Beitragsbedingungen |
| Marken und Namensnennung | urheberrechtliche Rechteeinräumung **≠** Namens-/Markenrechtseinräumung; **keine** eingetragene Marke beansprucht; gesonderte Markenrichtlinie **zurückgestellt**, hier **nicht** erstellt |
| Haftungs-/Gewährleistungsausschluss | Behauptungsverbote (§13.2) bleiben in Kraft; ein Lizenz-Haftungsausschluss ist **keine** zutreffende Produktstatusaussage |

```text
Apache-2.0 ausgewählt  != Apache-2.0-Datei erstellt
lizenzierter Quellcode != unterstütztes Produkt
Open Source            != unterstützter Dienst
Pull Request möglich   != Beitrag angenommen
Repository sichtbar    != Beitragsprogramm aktiv
```

### 13.5 Welche Human-Maintainer-Autorität ist erforderlich?

| Artefakt | Erforderliche Autorität |
| -------- | ----------------------- |
| `README.md` | Ausdrückliche Human-Maintainer-Autorisierung der **Erstellung**, plus Autorisierung des konkreten Inhalts (der `README` trifft öffentliche Statusaussagen), plus die üblichen Staging-/Commit-/Push-Gates |
| `LICENSE` | Ausdrückliche Human-Maintainer-Entscheidung über das **Lizenzmodell** (eine Rechtsentscheidung des Rechteinhabers), **danach** Autorisierung der Dateierstellung, plus die üblichen Staging-/Commit-/Push-Gates. Zusätzlich zu dispositionieren: Beitragsannahme, Marken- und Namensnennung, Haftungs- und Gewährleistungsausschluss |

```text
README/LICENSE recommendation != artifact authorization
recommendation to create      != permission to create
```

**Autoritäts-Firewall (`A-10`).** Die getroffene `A-10`-Entscheidung ist eine **Disposition**, keine **Artefakterstellungsautorität**. Sie beantwortet die Fragen *ob* (`SHALL EXIST`), *in welchem Modus* (`PUBLIC / OPEN SOURCE`), *unter welcher Lizenz* (`Apache-2.0`) und *wann frühestens* (`README` vor dem ersten produktiven Source-Commit) — sie erteilt jedoch **keine** Erlaubnis, eine Datei anzulegen. Die tatsächliche Erstellung von `README.md` beziehungsweise `LICENSE` setzt **kumulativ** voraus:

1. eine **ausdrücklich autorisierte Repository-/Work-Package-Grenze**, deren freigegebener Scope die betreffende Datei **und** deren öffentliche Inhaltsautorität ausdrücklich benennt; **und**
2. die **Autorisierung des konkreten Inhalts** — beim `README` die konkreten öffentlichen Statusaussagen, bei der `LICENSE` der standardmäßige, unveränderte Apache-2.0-Text; **und**
3. die **Human-Maintainer-Repository-, Staging-, Commit- und Push-Gates** ([Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §6, `DEC-G-01`/`DEC-G-03`).

Fehlt auch nur eine dieser drei Bedingungen, ist die Erstellung **nicht** autorisiert. Dieses Dokument erzeugt **kein** solches Work Package, **reserviert keine** Kennung und **benennt keinen** Nachfolger (§17); `CO-WP-033` bleibt **nicht erzeugt und nicht reserviert**. **Current State:** diese drei Bedingungen sind inzwischen durch eine davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung erfüllt worden — sie hat die exakte Dateigrenze benannt, den konkreten öffentlichen Inhalt beziehungsweise den standardmäßigen Lizenztext autorisiert und die Repository-, Staging-, Commit- und Push-Gates unverändert dem Human Maintainer vorbehalten. `README.md` und `LICENSE` sind seitdem **erstellt und vorhanden**. Diese Erfüllung erfolgte **ohne** Work Package: `CO-WP-033` bleibt **nicht erzeugt und nicht reserviert**.

```text
A-10 disposition       != artifact creation authority
A-10 satisfied         != README created
A-10 satisfied         != LICENSE created
A-10 satisfied         != NOTICE created
A-10 satisfied         != CONTRIBUTING created
A-10 satisfied         != SECURITY.md created
decision made          != permission to create
license selected       != artifact created
scope names the file   != content authorized
content authorized     != commit/push authorized
A-10 satisfied         != CO-WP-033 created or reserved
```

**Weder `README.md` noch `LICENSE` wird durch dieses Work Package oder durch die `A-10`-Entscheidung erstellt.** Das bleibt unverändert zutreffend; beide Dateien sind erst durch die gesondert autorisierte Publikations-Realisierung **erstellt** worden und seitdem **vorhanden**. `NOTICE`, `CONTRIBUTING.md` und `SECURITY.md` bleiben **nicht existent** und **nicht autorisiert**.

## 14. ADR- und Decision-Voraussetzungen

`CO-WP-032` erzeugt **keine** ADR-Datei, **keine** Decision-ID, **keine** Risk-ID und keine Zeile im Decision Index oder Risk Register. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs. Die Entscheidungspakete in §7, §8 und §9 sind **ADR-Vorbereitung**, nicht ADR.

Vor einer produktiven Implementierung sind mindestens folgende Punkte zu dispositionieren — jeweils durch den Human Maintainer, jeweils unter der bestehenden Relevanzregel `HM-3`:

| Punkt | Gegenstand | Aktueller Stand |
| ----- | ---------- | --------------- |
| `P-3` | Erhebungsmechanismus und Transport für `CAP-DISCOVERY-004` — im Readiness Review ausdrücklich als ADR-pflichtig bezeichnet | **entschieden** — `SELECTED` / dispositioniert durch ausdrückliche Human-Maintainer-Entscheidung (§7.5): Option A primär, Option B ausschließlich ergänzend, Option C nicht Standardweg, Option D/E ausgeschlossen, kein Netzwerktransport; **kein** Decision-Identifier, **kein** ADR |
| Sprache/Runtime | erste Implementierungsplattform | **entschieden** — `SELECTED` — **Go** / dispositioniert durch ausdrückliche Human-Maintainer-Entscheidung (§8.4): Sprach-/Runtime-**Klasse** ausschließlich für den ersten Observe-Slice, Rust als dokumentierte stärkste Alternative erhalten; **keine** Go-Version, **keine** Distribution, **keine** Toolchain, **keine** API, **keine** Abhängigkeit; Paketlayout und Quellbaumplatzierung **nicht** durch diese, sondern durch die Zeile *Source Tree* entschieden; **kein** Decision-Identifier, **kein** ADR |
| Source Tree | Anwendungsstruktur | **entschieden** — `DECIDED` / dispositioniert durch ausdrückliche Human-Maintainer-Entscheidung (§9.2): repository-verwurzelte Go-Struktur, `cmd/coreops/`, `internal/observe/`, `internal/observe/collect/`, Normalisierung zunächst ohne eigenes Paket innerhalb `internal/observe/`, colocated `*_test.go`, **kein** `pkg/`, **kein** `src/`, **kein** Top-Level-`tests/`, **keine** öffentliche Go-API; **kein** Verzeichnis und **keine** Datei angelegt, **kein** `go.mod`, **kein** `go.sum`, **kein** Modulpfad, **keine** Go-Version, **keine** Toolchain; **kein** Decision-Identifier, **kein** ADR |
| Dependency-Admission | Inkraftsetzung des Gates aus §10 | **entschieden** — `IN FORCE` / dispositioniert durch ausdrückliche Human-Maintainer-Entscheidung (§10), Disposition `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`: Admission-Gate **in Kraft**, zugelassene Drittabhängigkeiten **0**, konkret zugelassen **keine**, erster Observe-Slice **Zero-Third-Party-Default**; Zulassung ausschließlich je konkreter Abhängigkeit durch den Human Maintainer und nur innerhalb eines ausdrücklich autorisierten Work-Package-Scopes, gebunden an Version/Revision, Digest, Herkunft, Lizenzstand und aufgelöste transitive Schließung; **kein** Decision-Identifier, **kein** ADR, **kein** Admission-Register |
| `NEW-8` | `README` / `LICENSE` | **entschieden** — `DECIDED` / dispositioniert durch ausdrückliche Human-Maintainer-Entscheidung (§13), Disposition `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS`: `README.md` **SHALL EXIST** (eigene, ausdrücklich autorisierte Repository-Änderung **vor** dem ersten produktiven Source-Commit; öffentliche Zusammenfassung ohne normative Autorität; quantitative Angaben als gepflegter, nicht-autoritativer Spiegel), `LICENSE` **SHALL EXIST**, Veröffentlichungsrechte-Modus **PUBLIC / OPEN SOURCE**, Outbound-Lizenz **Apache-2.0** im standardmäßigen unveränderten Text; Beitragsprogramm **NICHT AKTIV**, `CONTRIBUTING.md` nicht autorisiert, CLA **keine**, DCO **keiner**; **keine** Marken-/Namensrechteeinräumung, **keine** eingetragene Marke beansprucht, Markenrichtlinie **zurückgestellt**; `SECURITY.md` und `NOTICE` **nicht autorisiert / nicht erstellt**. Durch `A-10` selbst **keine** Datei erstellt und **keine** Erstellung autorisiert; `README.md` und `LICENSE` sind inzwischen durch die gesondert autorisierte Publikations-Realisierung **erstellt und vorhanden**; **kein** Decision-Identifier, **kein** ADR |
| Build/Packaging | Reproduzierbarkeit, Offline-Build, Artefaktidentität | nicht entschieden |
| Envelope-Zusammensetzung (R6) | heterogene All-Failure-Zusammensetzung im [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.5 — kanonische Nicht-Emission bei evidenzerhaltender Retention des erzeugten Materials | **entschieden** — ausdrückliche Human-Maintainer-Entscheidung; dokumentarisch umgesetzt (Observation Contract §9, §10.4, §10.5, §22; Test Envelope `OBS-LLH-TC-11`); **kein** Decision-Identifier, **kein** ADR |

**Zur Zeile `P-3`.** Diese Entscheidung ist **getroffen** und wird hier als **Decision-Disposition** geführt. Sie wird — wie die R6-Zeile — vom bestehenden Gate-A-Punkt `A-11` (§16.1) **erfasst**, erfüllt ihn aber **nicht**. Wie die für diese Zeile dokumentierte ADR-Pflicht dauerhaft eingelöst wird, ist hier **nicht** entschieden: es entsteht **keine** ADR-Datei, **keine** ADR-Nummer, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register und **kein** Work Package. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

**Zur Zeile „Sprache/Runtime“.** Diese Entscheidung ist **getroffen** und wird hier als **Decision-Disposition** geführt. Sie wird — wie die `P-3`- und die R6-Zeile — vom bestehenden Gate-A-Punkt `A-11` (§16.1) **erfasst**, erfüllt ihn aber **nicht**. Wie die dokumentierte ADR-Relevanz der Sprach-/Runtime-Wahl dauerhaft eingelöst wird, ist hier **nicht** entschieden: es entsteht **keine** ADR-Datei, **keine** ADR-Nummer, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register und **kein** Work Package. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

**Zur Zeile „Source Tree“.** Diese Entscheidung ist **getroffen** und wird hier als **Decision-Disposition** geführt. Sie wird — wie die `P-3`-, die Sprach-/Runtime- und die R6-Zeile — vom bestehenden Gate-A-Punkt `A-11` (§16.1) **erfasst**, erfüllt ihn aber **nicht**. Ob und wie die Struktur- und Platzierungsentscheidung dauerhaft ADR-seitig eingelöst wird, ist hier **nicht** entschieden. Das autoritative ADR-Kandidateninventar führt **keinen** Kandidaten für Quellbaum oder Repository-Layout; `ADR-0002` (modularer Monolith) betrifft die **Architekturform** und ist eine davon getrennte, **unentschiedene** Frage — die Source-Tree-Entscheidung erfüllt, präjudiziert und disponiert sie **nicht**. Die Vergabe von ADR-Nummern bleibt ausdrücklich Human-Maintainer-Sache: es entsteht **keine** ADR-Datei, **keine** ADR-Nummer, **kein** ADR-Kandidat, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register und **kein** Work Package. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

```text
source tree decided != architecture form decided
source tree decided != ADR-0002 disposed
source tree decided != ADR required
source tree decided != ADR accepted
source tree decided != A-11 satisfied
```

**Zur Zeile „Envelope-Zusammensetzung (R6)“.** Diese Entscheidung ist **getroffen** und in den Slice-Dokumenten dokumentarisch umgesetzt. Sie wird hier als **Decision-Disposition** geführt und vom bestehenden Gate-A-Punkt `A-11` (§16.1) **getragen**. Es entsteht dafür **keine** zusätzliche Gate-A-Checklistenzeile, **keine** neue Checklistenkennung, **keine** Decision-, Risk-, ADR-, CCR- oder Capability-Kennung, **kein** Support-Status-Eintrag, **keine** Lesson, **kein** NDF-Feedback-Eintrag und **kein** Work Package — insbesondere wird **kein** Nachfolge-Work-Package erzeugt oder reserviert (§17). Auch der zugehörige docs-only Nachtrag erzeugt **keine** ADR-Datei und **keine** Zeile im Decision Index oder Risk Register; es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

**Zur Zeile „Dependency-Admission“.** Diese Entscheidung ist **getroffen** und wird hier als **Decision-Disposition** geführt. Sie wird — wie die `P-3`-, die Sprach-/Runtime-, die Source-Tree- und die R6-Zeile — vom bestehenden Gate-A-Punkt `A-11` (§16.1) **erfasst**, erfüllt ihn aber **nicht**. Entschieden ist ausschließlich die **Inkraftsetzung des Verfahrens** samt seiner Klassifikations-, Autoritäts-, Granularitäts- und Bindungsregeln; **keine** konkrete Abhängigkeit ist zugelassen, bewertet oder als Adoptionskandidat benannt. Ob und wie die Dependency-Admission-Governance dauerhaft ADR-seitig eingelöst wird, ist hier **nicht** entschieden: es entsteht **keine** ADR-Datei, **keine** ADR-Nummer, **kein** ADR-Kandidat, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** Zeile im Decision Index oder Risk Register, **kein** Dependency-Admission-Register und **kein** Work Package. `DEC-S-03` (*Technical foundation dependencies*) wird durch diese Disposition **nicht** berührt, **nicht** geschlossen und **nicht** umgestuft. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

```text
Admission-Gate in Kraft != ADR akzeptiert
Admission-Gate in Kraft != Abhängigkeit zugelassen
A-9-Disposition         != DEC-S-03-Disposition
A-9 erfüllt             != A-11 erfüllt
```

**Zur Zeile `NEW-8`.** Diese Entscheidung ist **getroffen** und wird hier als **Decision-Disposition** geführt. Sie wird — wie die `P-3`-, die Sprach-/Runtime-, die Source-Tree-, die Dependency-Admission- und die R6-Zeile — vom bestehenden Gate-A-Punkt `A-11` (§16.1) **erfasst**, erfüllt ihn aber **nicht**. Entschieden ist ausschließlich die **Disposition** über Existenz, Zeitpunkt, Veröffentlichungsmodus, Outbound-Lizenz sowie die Beitrags-, Marken- und Behauptungsgrenzen; **keine** Datei ist erstellt, und **keine** Erstellung ist autorisiert. Ob und wie die Veröffentlichungs-/Lizenzentscheidung dauerhaft ADR-seitig eingelöst wird, ist hier **nicht** entschieden: es entsteht **keine** ADR-Datei, **keine** ADR-Nummer, **kein** ADR-Kandidat, **keine** Decision-ID, **keine** `HM-*`-Kennung, **keine** `DEC-S-*`-Disposition, **keine** Zeile im Decision Index oder Risk Register, **keine** Markenrichtlinie, **keine** Beitragsrichtlinie, **keine** Sicherheitsrichtlinie, **kein** Release-Objekt und **kein** Work Package. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs.

```text
NEW-8 entschieden != ADR erforderlich
NEW-8 entschieden != ADR akzeptiert
NEW-8 entschieden != A-11 erfüllt
A-10 erfüllt      != A-11 erfüllt
A-10 erfüllt      != A-13 erfüllt
A-10 erfüllt      != A-14 erfüllt
```

**`A-11` bleibt `offen`.** `A-11` verlangt die Disposition **der Punkte dieser §14**. Die verbleibende Zeile — Build/Packaging — ist weiterhin unentschieden. Sechs entschiedene Zeilen (`P-3`, die R6-Zusammensetzung, Sprache/Runtime, Source Tree, Dependency-Admission und `NEW-8`) erfüllen `A-11` daher **nicht**; ebenso bleibt offen, wie die dokumentierte ADR-Pflicht dauerhaft eingelöst wird.

**`A-12` ist `erfüllt`.** Die freigegebene Abfolge lautete: **R6-Semantikdisposition → Definition von `OBS-LLH-TC-11` → Human-Maintainer-Disposition von `A-12`**. Die ersten beiden Schritte sind mit dem docs-only Nachtrag erfolgt; der dritte — die Human-Maintainer-Disposition der Testdesign- und `P-2`-Methodendefinition — ist mit der freigegebenen Human-Maintainer-`A-12`-Disposition **erfolgt**. Die Abfolge ist damit **vollständig**, und `A-12` steht auf **erfüllt** (§16.1). Die Gate-A-Bilanz lautete damit **6 von 14 erfüllt, 0 teilweise, 8 offen**; nach der späteren, davon getrennten Human-Maintainer-`A-6`/`P-3`-Disposition **7 von 14 erfüllt, 0 teilweise, 7 offen**, nach der wiederum davon getrennten Human-Maintainer-`A-7`-Disposition **8 von 14 erfüllt, 0 teilweise, 6 offen**, nach der abermals davon getrennten Human-Maintainer-`A-8`-Disposition **9 von 14 erfüllt, 0 teilweise, 5 offen**, nach der wiederum davon getrennten Human-Maintainer-`A-9`-Disposition **10 von 14 erfüllt, 0 teilweise, 4 offen**, und nach der erneut davon getrennten Human-Maintainer-`A-10`-Disposition lautet sie **11 von 14 erfüllt, 0 teilweise, 3 offen — Gate A ist weiterhin nicht passiert**; offen bleiben genau `A-11`, `A-13` und `A-14`. Die Gate-B-Bilanz bleibt **0 von 8 — realer Zielzugriff bleibt untersagt**. Die Disposition akzeptiert ausschließlich den **Definitionsstand** von Teststrategie und `P-2`-Evidenzmethode; sie implementiert **keinen** Test, führt **keinen** Test aus, erzeugt **keine** Evidenz, erfüllt `P-2` **nicht** und berührt `A-11` **nicht**.

```text
one decision disposed != A-11 satisfied
R6 decided            != A-12 disposed
A-12 disposed         != A-11 satisfied
A-12 satisfied        != Gate A passed
A-12 satisfied        != P-2 satisfied
A-12 satisfied        != tests implemented
A-12 satisfied        != tests executed
A-12 satisfied        != implementation authorized
R6 decided            != productive code authorized
R6 decided            != implementation authorized
R6 decided            != target access authorized
P-3 selected          != language/runtime selected
P-3 selected          != target_id derivation decided
P-3 selected          != A-11 satisfied
P-3 selected          != ADR accepted
P-3 selected          != Gate A passed
Go selected           != A-11 satisfied
Go selected           != ADR accepted
Go selected           != Gate A passed
source tree decided   != source tree created
source tree decided   != cmd/ or internal/ created
source tree decided   != go.mod created
source tree decided   != go.sum created
source tree decided   != module path decided
source tree decided   != dependency admitted
source tree decided   != build/packaging decided
source tree decided   != implementation authorized
source tree decided   != A-11 satisfied
source tree decided   != A-13 satisfied
source tree decided   != A-14 satisfied
source tree decided   != ADR accepted
source tree decided   != Gate A passed
A-8 satisfied         != Gate A passed
A-9 satisfied         != dependency admitted
A-9 satisfied         != dependency installed
A-9 satisfied         != dependency resolved
A-9 satisfied         != third-party package selected
A-9 satisfied         != test framework selected
A-9 satisfied         != go.mod created
A-9 satisfied         != go.sum created
A-9 satisfied         != vendor tree created
A-9 satisfied         != module path decided
A-9 satisfied         != Go version selected
A-9 satisfied         != Go distribution selected
A-9 satisfied         != Go toolchain selected
A-9 satisfied         != concrete stdlib API selected
A-9 satisfied         != build authorized
A-9 satisfied         != build/packaging decided
A-9 satisfied         != source tree created
A-9 satisfied         != productive code authorized
A-9 satisfied         != implementation authorized
A-9 satisfied         != test implementation authorized
A-9 satisfied         != test execution authorized
A-9 satisfied         != fixture authorized
A-9 satisfied         != P-1 granted
A-9 satisfied         != target access authorized
A-9 satisfied         != real observation authorized
A-9 satisfied         != P-2 satisfied
A-9 satisfied         != A-10 satisfied
A-9 satisfied         != A-11 satisfied
A-9 satisfied         != A-13 satisfied
A-9 satisfied         != A-14 satisfied
A-9 satisfied         != CO-WP-033 created or reserved
A-9 satisfied         != Gate A passed
A-10 satisfied        != README created
A-10 satisfied        != LICENSE created
A-10 satisfied        != NOTICE created
A-10 satisfied        != CONTRIBUTING created
A-10 satisfied        != SECURITY.md created
A-10 satisfied        != source tree created
A-10 satisfied        != productive code authorized
A-10 satisfied        != implementation authorized
A-10 satisfied        != productive-source publication authorized
A-10 satisfied        != product release
A-10 satisfied        != GitHub Release
A-10 satisfied        != dependency admitted
A-10 satisfied        != go.mod created
A-10 satisfied        != go.sum created
A-10 satisfied        != Go toolchain selected
A-10 satisfied        != A-11 satisfied
A-10 satisfied        != A-13 satisfied
A-10 satisfied        != A-14 satisfied
A-10 satisfied        != CO-WP-033 created or reserved
A-10 satisfied        != Gate A passed
README present        != product usable
LICENSE present       != product supported
public repository     != open source by visibility alone
Apache-2.0 selected   != Apache-2.0 file created
licensed source       != supported product
pull request possible != contribution accepted
Gate A passed         != Gate B passed
```

**Verhältnis zu den offenen `CCR`.** Die sechs offenen `CCR` bleiben offen und werden hier **nicht** geschlossen. `CCR-05` und `CCR-07` tragen `MUST CLOSE BEFORE DEPLOY`; das Readiness Review hat festgestellt, dass beide privilegierte beziehungsweise Offline-Ausführung betreffen und einen read-only Slice **nicht** berühren — sie sind ausdrücklich **keine** Observe-Precondition. Diese Feststellung wird hier nur referenziert, nicht neu getroffen.

```text
ADR candidate exists       != ADR required
ADR prepared               != ADR accepted
decision packet            != decision
foundation semantics established != ADR accepted
```

## 15. Künftige Testvoraussetzungen

Die Testvoraussetzungen verteilen sich auf **beide** Gates und dürfen nicht vermengt werden. Der [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §10 führt die Ausführungsvoraussetzungen vollständig.

### 15.1 Gate A — vor produktivem Code zu dispositionieren

Auf dieser Ebene ist ausschließlich verlangt, dass **definiert ist, was später geprüft und belegt werden muss** — nicht, dass es bereits geprüft oder belegt wäre.

1. **Teststrategie beziehungsweise künftige Validierungsanforderungen definiert** — für den Slice erfüllt durch den Test Envelope (elf Mindestfälle mit Intent, erwartetem Ausgang, Datensemantik, Provenance-Verhalten, verbotener Inferenz und künftiger Evidenzanforderung).
2. **`P-2`-Evidenzmethode definiert** — erfüllt durch Test Envelope §8.3 (sieben kumulative Anforderungen). **Nicht** verlangt: dass die Evidenz vorliegt.
3. `P-3` entschieden — **erfüllt**; der Erhebungsmechanismus ist als Mechanismusklasse ausdrücklich entschieden (§7.5). Das autorisiert **keine** Implementierung eines Falls.
4. Sprache/Runtime entschieden — **erfüllt**; die Sprach-/Runtime-**Klasse** ist ausdrücklich entschieden (§8.4: **Go**). Das autorisiert **keine** Testimplementierung und **keine** Testausführung: es ist **kein** Testwerkzeug, **keine** Toolchain und **keine** Go-Version ausgewählt.

**Disposition dieser Definitionen.** Die Punkte 1 und 2 sind durch die freigegebene Human-Maintainer-`A-12`-Disposition als für Gate A **hinreichend definiert akzeptiert** (Gate-A-Punkt `A-12` **erfüllt**, §16.1). Die Akzeptanz betrifft ausschließlich den **Definitionsstand**: sie ändert die `P-2`-Evidenzmethode **nicht**, autorisiert **keine** Testimplementierung und **keine** Testausführung, erzeugt **keine** Evidenz und erfüllt `P-2` **nicht**. Punkt 3 ist durch die ausdrückliche Human-Maintainer-`P-3`-Entscheidung **erfüllt** (§7.5); Punkt 4 ist durch die davon getrennte, ausdrückliche Human-Maintainer-Sprach-/Runtime-Entscheidung **erfüllt** (§8.4). Beides betrifft ausschließlich die jeweilige **Klasse** und autorisiert **keine** Testimplementierung und **keine** Testausführung; ein Testwerkzeug, eine Toolchain und eine Go-Version sind **nicht** ausgewählt. Es bleibt **kein** Test implementiert und **kein** Test ausgeführt, und `P-2` bleibt **`NOT SATISFIED`**.

### 15.2 Gate B — vor Testausführung und realer Beobachtung

5. Fixtures mit Identity, Revision und Provenance freigegeben — **existieren nicht**.
6. Lab-Environment mit deklarierter Environment Identity bereitgestellt — **nicht bereitgestellt**.
7. `P-1` erteilt, soweit reale Ziele berührt werden — **nicht erteilt**.
8. **Testausführung ausdrücklich autorisiert** — **nicht erteilt**.
9. `P-2` durch beobachtetes Schutzverhalten belegt — **nicht belegt**.

```text
test planned      != test implemented != test executed != test passed
test design       != test execution
test strategy defined != tests implemented
test result       != support status != production readiness
zero tests executed = the current state
```

**Keiner der Punkte 5 bis 9 ist eine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.**

## 16. Transition-Checklisten (fail-closed, zwei Gates)

Beide Checklisten sind **fail-closed**: `nicht dispositioniert` zählt als **nicht erfüllt** — nicht als „vermutlich in Ordnung". Sie sind **getrennt** zu bewerten; eine bestandene Gate-A-Checkliste sagt über Gate B nichts aus und umgekehrt.

### 16.1 Gate A — vor Produktivcode-/Implementierungsautorisierung

Eine spätere Produktivcode-Autorisierung gilt **nicht** als vorbereitet, solange auch nur eine dieser Zeilen offen ist.

| # | Punkt | Anforderung an die Disposition | Stand heute |
| - | ----- | ------------------------------ | ----------- |
| A-1 | **Observe-Phasenautorität** | Observe betreten und die Grenze von `HM-O1` ausdrücklich benannt | **erfüllt** — `HM-O1` `APPROVED WITH BOUNDARY` |
| A-2 | **Gewählter Value Slice** | ein benannter, eng begrenzter Slice | **erfüllt** — `HM-O2` `APPROVED` |
| A-3 | **`CO-WP-032` geschlossen** | Nova Review abgeschlossen **und** Human-Maintainer-Integration erfolgt | **erfüllt** — Nova Final Review `GO`, Human-Maintainer-Integration abgeschlossen (Commit `9999114`) |
| A-4 | **Nova Review** | Bewertung des Slice-Vertrags, des Test Envelope und dieses Gate-Dokuments liegt vor | **erfüllt** — Initial Review `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`** |
| A-5 | **Repository-Integration** | Human-Maintainer-Staging, -Commit und -Push erfolgreich durchgeführt | **erfüllt** — Commit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht, `origin/main` gleichauf |
| A-6 | **`P-3`-Entscheidung** | Erhebungsmechanismus und Transport ausdrücklich entschieden | **erfüllt** — `SELECTED` (§7.5): Option A primär, Option B ausschließlich ergänzend, Option C nicht Standardweg, Option D/E für diesen Slice ausgeschlossen, **kein** Netzwerktransport. Ausdrücklich **keine** Sprach-/Runtime-, Pfad-, API-, Bibliotheks- oder Werkzeugauswahl |
| A-7 | **Sprach-/Runtime-Entscheidung** | Plattform ausdrücklich entschieden | **erfüllt** — `SELECTED` — **Go** (§8.4), ausschließlich als **Sprach-/Runtime-Klasse** für den ersten Observe-Slice; Rust bleibt dokumentierte stärkste Alternative. Ausdrücklich **keine** Go-Versions-, Distributions-, Toolchain-, API-, Dependency- oder Build-/Packaging-Auswahl und **kein** breiterer Technologie-Stack; Paketlayout und Quellbaumplatzierung sind **nicht** durch `A-7`, sondern durch `A-8` entschieden (§9.2) |
| A-8 | **Source-Tree-Entscheidung** | Struktur ausdrücklich entschieden | **erfüllt** — `DECIDED` (§9.2), Disposition `APPROVE OPTION A WITH NOVA BOUNDARY CLARIFICATIONS`: repository-verwurzelte Go-Struktur; künftiger Modulwurzelort Repository-Wurzelverzeichnis; `cmd/coreops/` (nur Komposition/Bootstrap); `internal/observe/` (internal/privat); `internal/observe/collect/`; Normalisierung zunächst ohne eigenes Paket innerhalb `internal/observe/`; colocated `*_test.go`; **kein** `pkg/`, **kein** `src/`, **kein** Top-Level-`tests/`, **keine** öffentliche Go-API. Ausdrücklich **keine** Verzeichnis- oder Dateianlage, **kein** `go.mod`, **kein** `go.sum`, **kein** Modulpfad, **keine** Go-Versions-, Distributions- oder Toolchain-Auswahl, **keine** Dependency-Zulassung, **keine** Build-/Packaging-, Implementierungs-, Testimplementierungs-, Testausführungs- oder Fixture-Autorisierung; die Anlage von `cmd/` und `internal/` bleibt nach [Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §12 eigenständig gate-pflichtig |
| A-9 | **Dependency-Admission-Disposition** | Admission-Gate in Kraft; jede vorgesehene Abhängigkeit einzeln zugelassen oder ausdrücklich keine | **erfüllt** — Admission-Gate `IN FORCE` (§10), Disposition `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`; **ausdrücklich keine** Abhängigkeit vorgesehen und **0** zugelassen; erster Observe-Slice **Zero-Third-Party-Default**. Verbindlich: Standardbibliothek der gewählten Go-Distribution und ohne separaten Softwarebezug genutzte Plattformfähigkeit **ohne** individuelle Drittzulassung — ohne dass deren Lieferkette dadurch geklärt wäre; jedes separat bezogene Modul, Paket, jede Bibliothek und jedes Werkzeug — einschließlich `golang.org/x/...`, sofern nicht tatsächlich Standardbibliothek — als **Drittabhängigkeit**; transitive, Build-only-, Test-only-, separat bezogene Tooling-/Codegenerierungs- und vendorierte Abhängigkeiten **im Scope**. Zulassung ausschließlich durch den Human Maintainer, je konkreter Abhängigkeit und nur innerhalb eines ausdrücklich autorisierten Work-Package-Scopes; gebunden an Version/Revision, Digest, Herkunft, Lizenzstand und aufgelöste transitive Schließung. Ausdrücklich **keine** Zulassung, **keine** Auswahl, **keine** Installation, **keine** Auflösung, **kein** Testframework, **kein** `go.mod`, **kein** `go.sum`, **kein** Vendor-Baum und **keine** Go-Versions-, Distributions-, Toolchain- oder Build-/Packaging-Auswahl |
| A-10 | **`NEW-8`-Disposition** | `README` und `LICENSE` entschieden — auch eine bewusste Nicht-Erstellung ist eine Disposition | **erfüllt** — `DECIDED` (§13), Disposition `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS`: `README.md` **SHALL EXIST**, zu erstellen und remote zu integrieren in einer **eigenen, ausdrücklich autorisierten** Repository-Änderung **vor** dem ersten produktiven Source-Commit, als öffentliche Zusammenfassung **ohne** normative Autorität und mit minimierten, ausdrücklich gepflegten und **nicht-autoritativen** quantitativen Spiegeln; `LICENSE` **SHALL EXIST**; Veröffentlichungsrechte-Modus **PUBLIC / OPEN SOURCE**; Outbound-Lizenz **Apache-2.0** im standardmäßigen, **unveränderten** Text; öffentliches Beitragsprogramm **NICHT AKTIV**, `CONTRIBUTING.md` **nicht autorisiert / nicht erstellt**, CLA **keine**, DCO **keiner**; **keine** Marken-/Namensrechteeinräumung durch die Lizenz, **keine** eingetragene Marke beansprucht, gesonderte Markenrichtlinie **zurückgestellt**; Behauptungsverbote zu Reife, Sicherheit, Zertifizierung, Compliance, Support, SLA, Verfügbarkeit, Wartung, Kompatibilität, Installationsreife und Zeitplan bleiben in Kraft; `SECURITY.md` und `NOTICE` **nicht autorisiert / nicht erstellt**. Durch `A-10` selbst ausdrücklich **keine** Artefakterstellung und **keine** Erstellungsautorisierung; `README.md` und `LICENSE` sind inzwischen durch die gesondert autorisierte Publikations-Realisierung **erstellt und vorhanden**, wobei `LICENSE` den standardmäßigen, **unveränderten** Apache-2.0-Text trägt und dieser Text in diesem Dokument **nicht** wiedergegeben und **nicht** verändert ist; **keine** Quellbaumanlage, **keine** Dependency-Zulassung, **keine** Produktivcode-, Implementierungs-, Testimplementierungs-, Testausführungs- oder Quellcode-Veröffentlichungsautorisierung, **kein** Produktrelease, **kein** GitHub Release |
| A-11 | **ADR-/Decision-Disposition** | die Punkte aus §14 entschieden oder ausdrücklich zurückgestellt | **offen** |
| A-12 | **Teststrategie / künftige Validierungsanforderungen definiert** | §15.1 Punkte 1 und 2 — einschließlich der **definierten `P-2`-Evidenzmethode**; ausdrücklich **nicht** `P-2` erfüllt | **erfüllt** — Testdesign **definiert**, `P-2`-Evidenzmethode **definiert**, Human-Maintainer-Disposition dieser Definition **erteilt**. Ausdrücklich **nicht** `P-2` erfüllt: **0** Tests implementiert, **0** ausgeführt, No-Mutation-Evidenz **keine** |
| A-13 | **Build-/Packaging-Disposition** | Reproduzierbarkeit, Offline-Build und Artefaktidentität entschieden, soweit §12 sie als erforderlich ausweist | **offen** |
| A-14 | **Ausdrückliche Implementierungs-Work-Package-Autorität** | ein eigenes, ausdrücklich autorisiertes implementierungsorientiertes Work Package mit benanntem Scope | **offen** — nicht erteilt |

**Bilanz Gate A: 11 von 14 erfüllt, 0 teilweise, 3 offen. Gate A ist nicht passiert.** Offen bleiben genau drei Punkte: `A-11`, `A-13` und `A-14`.

**Ausdrücklich nicht in dieser Checkliste:** `P-1` · `P-2` erfüllt · Testausführungsautorität · Fixtures · Lab. Sie sind **keine** Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.

### 16.2 Gate B — vor realer Beobachtung / Zielausführung

Dieses Gate wird **erst nach** Gate A und **erst nach** einer vorhandenen Implementierung überhaupt bewertbar.

| # | Punkt | Anforderung an die Disposition | Stand heute |
| - | ----- | ------------------------------ | ----------- |
| B-1 | **`P-1` konkrete Zielautorisierung** | konkretes Ziel benannt (keine Zielklasse als Blankett) | **offen** — `NOT AUTHORIZED` |
| B-2 | **Exakt erlaubte Beobachtung** | welche Felder, aus welchen Quellklassen, mit welcher Erhebungsmechanismusklasse | **offen** |
| B-3 | **Exakte Ausführungsgrenze** | read-only; keine Rechteeskalation, keine Credentials, keine Secrets, kein Netzwerkziel; zeitliche Begrenzung, Widerrufbarkeit, Auditpflicht | **offen** |
| B-4 | **`P-2` No-Mutation-/Ausführungssicherheits-Evidenz** | beobachtetes geschütztes Verhalten nach Test Envelope §8.3 — **erst nach** Implementierung und Ausführung möglich | **offen** — `NOT SATISFIED` |
| B-5 | **Testausführungsautorität** | ausdrückliche Human-Maintainer-Autorisierung der Ausführung | **offen** — nicht erteilt |
| B-6 | **Fixtures** | Identity, Revision und Provenance freigegeben | **offen** — existieren nicht |
| B-7 | **Lab-Environment** | deklarierte Environment Identity bereitgestellt | **offen** — nicht bereitgestellt |
| B-8 | **Sonstige ausführungsspezifische Autorität** | jede weitere für den konkreten Zugriff erforderliche Freigabe | **offen** |

**Bilanz Gate B: 0 von 8 erfüllt. Realer Zielzugriff bleibt untersagt.**

```text
checklist documented != checklist satisfied
Gate A open          = productive code NOT AUTHORIZED
Gate B open          = target access NOT AUTHORIZED
Gate A passed        != Gate B passed
productive code authorization != target authorization
```

## 17. Kein automatischer Nachfolger

```text
Es erfolgt keine automatische Freigabe eines Folge-Work-Packages.
```

Der Abschluss von `CO-WP-032` — einschließlich Nova Review und Human-Maintainer-Integration — erzeugt **keinen** Nachfolger, reserviert **keine** Kennung und begründet **keine** Erwartung.

Zulässige Formulierung, und nur diese:

> **Ein späteres implementierungsorientiertes Work Package kann in Betracht gezogen werden, nachdem `CO-WP-032` geschlossen ist und die erforderlichen Human-Maintainer-Entscheidungen getroffen sind.**

```text
candidate successor != reserved work package
recommendation      != authorization
closure             != successor
```

Es wird **keine** Nachfolger-Kennung vergeben, **keine** reserviert und **keine** genannt. Es wird **kein** neuer primärer Work-Package-Typ eingeführt.

## 18. Compatibility

Additiv. **Keine** Änderung an Foundation Scope Lock, Release-Taxonomie, Decision Index, Risk Register, Lessons-Learned-Register, NDF-Feedback-Kandidaten, Capability Matrix, Initial Support Boundary, Roadmap oder Teststrategie. **Kein** ADR erzeugt oder akzeptiert. **Keine** Decision-, Risk- oder ADR-ID vergeben. **Keine** Technologie wird durch dieses Dokument selbst ausgewählt; die `P-3`-Mechanismusklassen-, die Sprach-/Runtime-, die Source-Tree- und die Dependency-Admission-Entscheidung sind eigenständige Human-Maintainer-Entscheidungen (§7.5, §8.4, §9.2, §10) und ausdrücklich **kein** breiterer Technologie-Stack. Die Inkraftsetzung des Dependency-Admission-Gates lässt **keine** Abhängigkeit zu, wählt **kein** Paket, **kein** Modul, **keine** Bibliothek, **kein** Werkzeug und **kein** Testframework aus, erzeugt **kein** Dependency-Admission-Register und berührt `DEC-S-03` **nicht**. Die Source-Tree-Entscheidung ist eine **Struktur- und Platzierungsentscheidung**; sie legt **kein** Verzeichnis und **keine** Datei an und ändert den [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) **nicht**. Die `NEW-8`-/`A-10`-Entscheidung (§13) ist eine **Veröffentlichungs- und Lizenzdisposition**; sie erstellt **keine** Datei — weder `README.md` noch `LICENSE`, `NOTICE`, `CONTRIBUTING.md` oder `SECURITY.md` — und autorisiert deren Erstellung **nicht** (`README.md` und `LICENSE` sind erst später durch eine davon getrennte, gesondert autorisierte Publikations-Realisierung entstanden und sind seitdem **vorhanden**; `NOTICE`, `CONTRIBUTING.md` und `SECURITY.md` bleiben **nicht erstellt**), gibt **keinen** Apache-2.0-Lizenztext wieder, ändert die Source-of-Truth-Hierarchie **nicht**, aktiviert **kein** Beitragsprogramm, erzeugt **keine** Marken-, Beitrags- oder Sicherheitsrichtlinie und beansprucht **keine** eingetragene Marke. **Kein** Artefakt außerhalb der neun autorisierten Pfade erzeugt. Breaking-Change-Potenzial: gering.

## 19. Next Decision

Der Nova Final Review dieses Dokuments ist erfolgt (`GO`) und die Human-Maintainer-Repository-Integration ist abgeschlossen (Commit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht); die Gate-A-Punkte `A-1` bis `A-5` sind damit erfüllt; `A-12` ist durch die freigegebene Human-Maintainer-`A-12`-Disposition ebenfalls **erfüllt**, `A-6` durch die ausdrückliche Human-Maintainer-`P-3`-Entscheidung (§7.5), `A-7` durch die davon getrennte, ausdrückliche Human-Maintainer-Sprach-/Runtime-Entscheidung (§8.4), `A-8` durch die abermals davon getrennte, ausdrückliche Human-Maintainer-Source-Tree-Entscheidung (§9.2) `A-9` durch die wiederum davon getrennte, ausdrückliche Human-Maintainer-Dependency-Admission-Entscheidung (§10) und `A-10` durch die erneut davon getrennte, ausdrückliche Human-Maintainer-`NEW-8`-Entscheidung (§13). **Als Nächstes** — und jeweils als eigene, ausdrückliche Human-Maintainer-Entscheidung — die verbleibenden **drei** offenen Gate-A-Punkte `A-11`, `A-13` und `A-14` (§16.1). Die Erstellung von `README.md` und `LICENSE` war **kein** offener Gate-A-Punkt, sondern eine davon getrennte, **eigens zu autorisierende Repository-Aktion** (§13.5): `A-10` ist **erfüllt**, und die Dateien sind durch die dafür gesondert erteilte Human-Maintainer-Publikations-Realisierungsautorisierung inzwischen **erstellt und vorhanden**; `A-10` selbst hat sie **nicht** erstellt und ihre Erstellung **nicht** autorisiert. Die `P-3`-Entscheidung hat `A-7` **nicht** autorisiert: sie wählte **keine** Sprache und **keine** Runtime aus; die Sprach-/Runtime-Auswahl war eine davon getrennte, spätere Entscheidung (§8.4). Die Sprach-/Runtime-Entscheidung ihrerseits hat `A-8` **nicht** autorisiert: die Source-Tree-Auswahl war eine davon getrennte, spätere Entscheidung (§9.2). Die Source-Tree-Entscheidung ihrerseits autorisiert **weder** `A-9` **noch** `A-13` **noch** `A-14`: sie lässt **keine** Abhängigkeit zu, disponiert Build und Packaging **nicht** und erteilt **keine** implementierungsorientierte Work-Package-Autorität. Die Dependency-Admission-Entscheidung ihrerseits autorisierte **weder** `A-10` **noch** `A-11` **noch** `A-13` **noch** `A-14`: sie ließ **keine** konkrete Abhängigkeit zu, disponierte `README`/`LICENSE` **nicht**, akzeptierte **keinen** ADR, disponierte Build und Packaging **nicht** und erteilte **keine** implementierungsorientierte Work-Package-Autorität; die `NEW-8`-Disposition war eine davon getrennte, spätere Entscheidung (§13). Sie setzt ausschließlich das Verfahren in Kraft; die zugelassene Menge bleibt **leer**. Die `NEW-8`-Entscheidung ihrerseits autorisiert **weder** `A-11` **noch** `A-13` **noch** `A-14`: sie akzeptiert **keinen** ADR, disponiert Build und Packaging **nicht** und erteilt **keine** implementierungsorientierte Work-Package-Autorität. Sie lässt **keine** Abhängigkeit zu — zugelassene Drittabhängigkeiten bleiben **0** —, und dass das Lizenzkriterium (§10.2 Nr. 4) nunmehr gegen ein definiertes Outbound-Modell **bewertbar** wäre, ist **keine** Bewertung und **keine** Zulassung. Sie autorisiert **keine** Veröffentlichung produktiven Quellcodes, **keinen** funktionalen Produktrelease und **kein** GitHub Release. Sie legt **kein** Verzeichnis und **keine** Datei an; die Anlage von `cmd/` und `internal/` bleibt nach [Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §12 eigenständig gate-pflichtig und wird durch ein späteres `A-14`-Work-Package **nur dann** getragen, wenn dessen freigegebener Scope diese Architektur-/Governance-Autorität ausdrücklich enthält. Es wird **kein** Nachfolge-Work-Package erzeugt und **keines** reserviert; `CO-WP-033` bleibt **nicht erzeugt und nicht reserviert** (§17). **Erst danach und erst mit einer vorhandenen Implementierung** werden die Gate-B-Punkte `B-1` bis `B-8` (§16.2) überhaupt bewertbar. Dieses Dokument trifft **keine** dieser Entscheidungen.
