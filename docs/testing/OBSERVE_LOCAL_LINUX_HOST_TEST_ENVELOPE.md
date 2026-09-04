# CoreOps – Observe Local Linux Host Test Envelope

> Document Status: Implemented; Nova Review abgeschlossen (Initial `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`**); `completed-go-with-notes`
> Human-Maintainer-Repository-Integration: COMPLETE — integrierter Commit `9999114200bf18baaadfb508e8464720b75e352e`; Push COMPLETE; Remote-Integration COMPLETE; `origin/main` = `9999114200bf18baaadfb508e8464720b75e352e`
> Artifact Class: Test-Design-Artefakt (docs-only) — **keine** Testimplementierung, **keine** Testausführung
> Phase: `Observe` — betreten und mit Grenze autorisiert (`HM-O1` `APPROVED WITH BOUNDARY`)
> Value Slice: Local Linux Host Identity & Basic System Observation — SELECTED (`HM-O2` `APPROVED`)
> Work Package: `CO-WP-032`, primärer Typ `docs-only` — autorisiert durch `HM-O3` `APPROVED`; Ausführungsgrenze `HM-O4` `APPROVED WITH EXACT BOUNDARY` (docs-only)
> Successor Work Package: none created, none reserved
> Tests planned: YES · Tests implemented: **NO** · Tests executed: **NO** · Test evidence produced: **NONE**
> Target Authorization (`P-1`): Not authorized — Gate B; **not** a prerequisite for writing productive code
> No-Mutation Evidence (`P-2`): method/plan **DEFINED** (§8.3) · Human-Maintainer-Disposition **APPROVED** (Gate-A-Punkt `A-12` erfüllt, [Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §16.1) · satisfaction **NOT SATISFIED** — Gate B, downstream of implementation and execution
> Collection Mechanism (`P-3`): **SELECTED** — Mechanismusklasse ausdrücklich entschieden ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5): Option A primär, Option B ausschließlich ergänzend, Option C nicht Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport. **Keine** Auswahl von Pfad, API, Bibliothek, Werkzeug, Sprache oder Runtime; **keine** Testimplementierungs- oder Testausführungsautorisierung.
> Language / Runtime: **SELECTED** — **Go**, ausdrückliche Human-Maintainer-Entscheidung ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4); ausgewählt ist ausschließlich die **Sprach-/Runtime-Klasse** für den ersten Observe-Slice. **Keine** Go-Version, **keine** Distribution, **keine** Toolchain, **kein** Testframework, **keine** Abhängigkeit; **keine** Testimplementierungs- oder Testausführungsautorisierung.
> Source Tree: **DECIDED** — ausdrückliche Human-Maintainer-Entscheidung ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §9.2): repository-verwurzelte Go-Struktur, `cmd/coreops/`, `internal/observe/`, `internal/observe/collect/`; Normalisierung zunächst ohne eigenes Paket innerhalb `internal/observe/`. Künftige Testplatzierungspolitik: **colocated `*_test.go`**, **kein** Top-Level-`tests/`, paketlokales `testdata/` ausschließlich als **künftige** Fixture-Ablage, falls und sobald Fixture-/Testautorität besteht. **NOT CREATED** — **kein** Verzeichnis, **keine** Datei, **kein** `testdata/`, **kein** `go.mod`, **kein** `go.sum`. **Keine** Testframework-, Testdateinamen-, Testkonventions-, Fixture-, Testimplementierungs- oder Testausführungsauswahl und **keine** entsprechende Autorisierung.
> Dependency Admission: **IN FORCE** — ausdrückliche Human-Maintainer-Entscheidung ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §10); Gate-A-Punkt `A-9` **erfüllt**. Zugelassene Drittabhängigkeiten: **0** · konkret zugelassen: **keine** · erster Observe-Slice: **Zero-Third-Party-Default**. Testabhängigkeitsgrenze: die Testeinrichtungen der Go-Standardbibliothek benötigen **keine** individuelle Drittzulassung; ein **separat bezogenes** Testframework, -werkzeug oder eine -bibliothek benötigt die **vollständige** `A-9`-Zulassung; transitive Testabhängigkeiten sind **im Scope**. **Kein** Testframework, **keine** Testbibliothek, **kein** Testwerkzeug und **keine** Testabhängigkeit ausgewählt oder zugelassen; **keine** Testimplementierungs- oder Testausführungsautorisierung.
> Test Execution Authorization: Not granted
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)
> Nachträgliche docs-only Änderung: konzeptioneller Mindestfall `OBS-LLH-TC-11` (heterogene All-Failure-Zusammensetzung, [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.5 **R6**) ergänzt; §13 auf den tatsächlichen Stand korrigiert. Grundlage ist die ausdrückliche Human-Maintainer-Entscheidung zur heterogenen All-Failure-Zusammensetzung; **kein** Decision-Identifier vergeben, **kein** ADR, **kein** neues und **kein** reserviertes Work Package. `OBS-LLH-TC-11` ist **nicht implementiert**, **nicht ausgeführt** und trägt `not run`.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-6`/`P-3`-Disposition ist angewandt. `P-3` ist **entschieden** (Mechanismusklasse, [Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5); Gate-A-Punkt `A-6` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **7 von 14 erfüllt, 0 teilweise, 7 offen — nicht passiert** (durch die spätere `A-7`-Anwendung fortgeschrieben, siehe folgende Zeile), Gate B unverändert auf **0 von 8 — nicht passiert**. Die Ausführungsvoraussetzungen (§10) bleiben **insgesamt unerfüllt**: Testimplementierung und Testausführung **NOT AUTHORIZED**, **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, jeder Fall unverändert `not run`. `OBS-LLH-TC-11`-Semantik, **R6**-Semantik, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen und die `P-2`-Evidenzanforderungen aus §8.3 bleiben **unverändert**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-7`-Disposition ist angewandt. Die Sprach-/Runtime-Entscheidung ist **getroffen** — **Go**, ausschließlich als **Sprach-/Runtime-Klasse** für den ersten Observe-Slice ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4); Gate-A-Punkt `A-7` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **8 von 14 erfüllt, 0 teilweise, 6 offen — nicht passiert** (durch die spätere `A-8`-Anwendung fortgeschrieben, siehe folgende Zeile; offen blieben `A-8`, `A-9`, `A-10`, `A-11`, `A-13`, `A-14`), Gate B unverändert auf **0 von 8 — nicht passiert**. Die Ausführungsvoraussetzungen (§10) bleiben **insgesamt unerfüllt**: Testimplementierung und Testausführung **NOT AUTHORIZED**, **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, Fixtures und Lab **nicht bereitgestellt**, `P-1` **`NOT AUTHORIZED`**, jeder Fall unverändert `not run`. `OBS-LLH-TC-11`-Semantik, **R6**-Semantik, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen, Ebenenzuordnungen, Testintentionen, erwartete Ausgänge und die `P-2`-Evidenzanforderungen aus §8.3 bleiben **unverändert**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung.
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-8`-Disposition ist angewandt. Die Source-Tree-Entscheidung ist **getroffen** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §9.2); Gate-A-Punkt `A-8` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **9 von 14 erfüllt, 0 teilweise, 5 offen — nicht passiert** (durch die spätere `A-9`-Anwendung fortgeschrieben, siehe folgende Zeile; offen blieben `A-9`, `A-10`, `A-11`, `A-13`, `A-14`), Gate B unverändert auf **0 von 8 — nicht passiert**. Für diesen Envelope folgt daraus ausschließlich eine **künftige Platzierungspolitik**: Tests künftig als **colocated `*_test.go`**-Dateien neben ihren Paketen, **kein** Top-Level-`tests/`, paketlokales `testdata/` ausschließlich als **künftige** Ablage für Fixtures, falls und sobald Fixture- und Testautorität besteht. Dieser Envelope bleibt ein **Governance-/Testdesign-Artefakt** und begründet **kein** Quellbaumverzeichnis; es wird **kein** Verzeichnis, **kein** `testdata/` und **keine** Testdatei angelegt. **Nicht** entschieden bleiben: Testframework, konkrete Testdateinamen, In-Package- gegen External-Package-Testkonvention, Fixture-Inhalt und Fixture-Identität. Die Ausführungsvoraussetzungen (§10) bleiben **insgesamt unerfüllt**: Testimplementierung und Testausführung **NOT AUTHORIZED**, **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, Fixtures und Lab **nicht bereitgestellt**, `P-1` **`NOT AUTHORIZED`**, jeder Fall unverändert `not run`. `OBS-LLH-TC-01`…`OBS-LLH-TC-11`-Identitäten und -Semantik, **R6**-Semantik, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen, Ebenenzuordnungen, Testintentionen, erwartete und verbotene Ausgänge, die `P-2`-Evidenzanforderungen aus §8.3 und die No-Mutation-Anforderungen bleiben **unverändert**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**; `A-11` blieb nach dieser Anwendung **offen** (historischer Stand; inzwischen **erfüllt**, siehe die letzte Zeile dieses Blocks).
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-9`-Disposition ist angewandt. Das **Dependency-Admission-Gate ist in Kraft** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §10, Disposition `APPROVE OPTION B WITH NOVA BOUNDARY CLARIFICATIONS`); Gate-A-Punkt `A-9` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **10 von 14 erfüllt, 0 teilweise, 4 offen — nicht passiert** (durch die spätere `A-10`-Anwendung fortgeschrieben, siehe folgende Zeile; offen blieben `A-10`, `A-11`, `A-13`, `A-14`), Gate B unverändert auf **0 von 8 — nicht passiert**. Zugelassene Drittabhängigkeiten: **0**; konkret zugelassen: **keine**; erster Observe-Slice: **Zero-Third-Party-Default**. Für diesen Envelope folgt daraus ausschließlich eine **Testabhängigkeitsgrenze**: die Testeinrichtungen der Go-Standardbibliothek benötigen **keine** individuelle Drittzulassung; ein **separat bezogenes** Testframework, -werkzeug oder eine -bibliothek benötigt die **vollständige** `A-9`-Zulassung; transitive Testabhängigkeiten sind **im Scope**. Es ist **kein** Testframework, **keine** Testbibliothek, **kein** Testwerkzeug, **keine** Assertions-, Mock- oder Fixture-Bibliothek und **keine** Testabhängigkeit ausgewählt, empfohlen oder zugelassen. Die Ausführungsvoraussetzungen (§10) bleiben **insgesamt unerfüllt**: Testimplementierung und Testausführung **NOT AUTHORIZED**, **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, Fixtures und Lab **nicht bereitgestellt**, `P-1` **`NOT AUTHORIZED`**, jeder konzeptionelle Fall unverändert `not run`. `OBS-LLH-TC-01`…`OBS-LLH-TC-11`-Identitäten und -Semantik, **R6**-Semantik, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen, Ebenenzuordnungen, Testintentionen, erwartete und verbotene Ausgänge, die `P-2`-Evidenzmethode aus §8.3 und die No-Mutation-Anforderungen bleiben **unverändert**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **kein** Dependency-Admission-Register und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**; `A-11` blieb nach dieser Anwendung **offen** (historischer Stand; inzwischen **erfüllt**, siehe die letzte Zeile dieses Blocks).
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-10`-Disposition ist angewandt. Die `NEW-8`-Entscheidung (`README` / `LICENSE`) ist **getroffen** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §13, Disposition `APPROVE README + APACHE-2.0 PUBLICATION MODEL WITH NOVA BOUNDARY CLARIFICATIONS`); Gate-A-Punkt `A-10` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **11 von 14 erfüllt, 0 teilweise, 3 offen — nicht passiert** (durch die spätere `A-13`-Anwendung fortgeschrieben, siehe folgende Zeile; offen blieben `A-11`, `A-13`, `A-14`), Gate B unverändert auf **0 von 8 — nicht passiert**. `README.md` **soll existieren** — zu erstellen in einer eigenen, ausdrücklich autorisierten Repository-Änderung **vor** dem ersten produktiven Source-Commit — und ist inzwischen durch eine davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung **erstellt und vorhanden**; `LICENSE` **soll existieren** und ist ebenso **erstellt und vorhanden** (standardmäßiger, **unveränderter** Apache-2.0-Text); Veröffentlichungsrechte-Modell **PUBLIC / OPEN SOURCE**, Outbound-Lizenz **Apache-2.0**; Beitragsprogramm **NICHT AKTIV**. `A-10` selbst hat die Dateien **nicht** erstellt; die Erstellungsautorität stammt ausschließlich aus jener gesonderten Autorisierung. `A-9` bleibt **erfüllt**, zugelassene Drittabhängigkeiten bleiben **0**. **Diese Angaben sind ausschließlich Current-State-Kontext und berühren die Testsemantik dieses Envelope nicht.** Für diesen Envelope folgt **keine** inhaltliche Änderung: die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die Beobachtungsausgangs- und Ergebnisvokabulare, die Emissions-Dispositionen, die Ebenenzuordnungen, die `P-2`-Evidenzmethode aus §8.3, die No-Mutation-Anforderungen und die aus `A-9` folgende Testabhängigkeitsgrenze bleiben **unverändert**. Es ist **kein** Testframework vorhanden oder ausgewählt; Testimplementierung und Testausführung bleiben **NOT AUTHORIZED**; **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, `P-1` **`NOT AUTHORIZED`**, Fixtures und Lab **nicht bereitgestellt**, jeder konzeptionelle Fall unverändert `not run`. Die Disposition erstellt **keine** Datei, autorisiert **keine** Dateierstellung, **keine** Testimplementierung, **keine** Testausführung, **keine** Fixture, **keinen** Zielzugriff, **keine** reale Beobachtung und **keine** Veröffentlichung produktiven Quellcodes; die spätere, gesondert autorisierte Publikations-Realisierung hat ausschließlich `README.md` und `LICENSE` erstellt und autorisiert von alledem ebenfalls **nichts**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**; `A-11` blieb nach dieser Anwendung **offen** (historischer Stand; inzwischen **erfüllt**, siehe die letzte Zeile dieses Blocks).
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-13`-Disposition ist angewandt. Die Build-/Packaging-Entscheidung ist **getroffen** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §12, Disposition `APPROVE MINIMUM-SUFFICIENT BUILD/PACKAGING DISPOSITION WITH NOVA BOUNDARY CLARIFICATIONS`, Geltung ausschließlich erster Observe-Slice); Gate-A-Punkt `A-13` ist **erfüllt**, Gate A stand nach dieser Anwendung auf **12 von 14 erfüllt, 0 teilweise, 2 offen — nicht passiert** (durch die spätere `A-11`-Anwendung fortgeschrieben, siehe die folgende Zeile; offen blieben: `A-11`, `A-14`), Gate B unverändert auf **0 von 8 — nicht passiert**. Auf **Governance-Ebene** entschieden sind: Reproduzierbarkeitsmodell **DECLARED-INPUT / DETERMINISTIC-IDENTITY** (Bit-für-Bit-Gleichheit **Design-Ziel**, **keine** Evidenzaussage), **Offline-Build erforderlich** (`GOTOOLCHAIN=local`, `GOPROXY=off`, fail-closed netzwerkfreie Ausführungsgrenze, konkrete Realisierung **zurückgestellt**), Build-Mechanismus **standardmäßige Go-Toolchain unmittelbar**, Go-Distributionsherkunftsklasse **offizielles Upstream-Go-Projekt-Release** mit exakter Patch-Pinning-**Regel** und Digest-Verifikationspflicht, `CGO_ENABLED=0`, `-trimpath`, `-buildvcs=false`, Quellrevision **über Provenance**, Go-Modulmechanismus **ausgewählt**, Artefaktklasse **Anwendungs-/Service-Artefakt** nach dem bestehenden Artifact-Identity-Modell, Ziel-OS-Klasse **Linux**, erste Artefaktform **natives Go-Executable**, SBOM-/Provenance-**Behandlung definiert**. **Zurückgestellt** bleiben: exakte Go-Version · Modulpfad · ausführbarer Basisname · `GOARCH` · Digest-Algorithmus · SBOM-Format/-Generator · Signatur-/Hash-Technologie · Container · CI. **Diese Angaben sind ausschließlich Current-State-Kontext und berühren die Testsemantik dieses Envelope nicht.** Für diesen Envelope folgt **keine** inhaltliche Änderung: die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die Beobachtungsausgangs- und Ergebnisvokabulare, die Record-Discarded- und Provenance-Semantik, die Emissions-Dispositionen, die Ebenenzuordnungen, die `P-2`-Evidenzmethode aus §8.3, die No-Mutation-Anforderungen und die aus `A-9` folgende Testabhängigkeitsgrenze bleiben **unverändert**. `A-13` wählt **kein** Testframework, **keine** Testbibliothek und **kein** Testwerkzeug aus und lässt **keine** Abhängigkeit zu — zugelassene Drittabhängigkeiten bleiben **0**; ein separat bezogenes Test- oder Build-Werkzeug bleibt vollständig `A-9`-pflichtig. Es ist **kein** Testkommando, **keine** Toolchain-Ausführung und **kein** Build autorisiert: Toolchain-Beschaffung und -Installation **NICHT AUTORISIERT**, Build-Ausführung **NICHT AUTORISIERT**, Testimplementierung und Testausführung bleiben **NOT AUTHORIZED**; **0** Tests implementiert, **0** Tests ausgeführt, `P-2` **`NOT SATISFIED`**, `P-1` **`NOT AUTHORIZED`**, Fixtures und Lab **nicht bereitgestellt**, jeder konzeptionelle Fall unverändert `not run`. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**; `A-11` blieb nach dieser Anwendung **offen** (historischer Stand; inzwischen **erfüllt**, siehe die letzte Zeile dieses Blocks).
> Nachträgliche docs-only Current-State-Anwendung: Die freigegebene Human-Maintainer-`A-11`-Disposition ist angewandt. Die ADR-/Decision-Disposition ist **getroffen** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §14): `APPROVE ADR-TREATMENT MATRIX WITH DEFERRED TRIGGERS WITH NOVA BOUNDARY CLARIFICATIONS`, Geltung ausschließlich **erster Observe-Slice**, Gegenstand ausschließlich die ADR-/Decision-Behandlung der sieben §14-Zeilen: `P-3` **`T2`** · Sprache/Runtime **`T2`** · Source Tree **`T4`** · Dependency-Admission **`T4`** · `NEW-8` **`T4`** · Build/Packaging **`T5`** · Envelope-Zusammensetzung (R6) **`T4`**. Gate-A-Punkt `A-11` steht damit auf **erfüllt**. Gate A nach dieser Anwendung: **13 von 14 erfüllt, 0 teilweise, 1 offen — nicht passiert**; offen bleibt genau `A-14`. Gate B unverändert **0 von 8 — nicht passiert**. **Für diesen Envelope folgt daraus keine inhaltliche Änderung.** Die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, das Beobachtungsergebnis- und Feldergebnis-Vokabular, die Verarbeitungs- und Emissions-Dispositionen, die Provenance- und Record-Discarded-Semantik, die Testebenenzuordnung sowie die `P-2`-Evidenzanforderungen und die No-Mutation-Evidenzanforderungen aus §8.3 bleiben **unverändert** — `A-11` ändert die Testsemantik **nicht**. Testimplementierung und Testausführung bleiben **NOT AUTHORIZED**; Tests bleiben **0 implementiert / 0 ausgeführt**; `P-2` bleibt **`NOT SATISFIED`**; zugelassene Drittabhängigkeiten bleiben **0**. Es entsteht **kein** Work Package, **keine** `CO-WP-033`, **kein** ADR, **keine** ADR-Datei, **keine** ADR-Nummer (**keine** vergeben, **keine** reserviert), **kein** ADR-Verzeichnis, **keine** ADR-Vorlage und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; ADR-Dateien bleiben **0**, akzeptierte ADRs bleiben **0**, Decision Index und Risk Register bleiben **unverändert**. Zurückgestellt ist **nicht** erloschen: `T2` und `T5` halten die dokumentierte ADR-Relevanz aufrecht.

## 1. Status

Test-**Design**-Envelope für den ersten Observe-Wertslice. Das Dokument benennt, **welche Fälle** ein späterer Beobachtungspfad abdecken muss und **welche Evidenz** ein späterer Lauf erzeugen müsste. Es implementiert **keinen** Test, führt **keinen** Test aus und behauptet **kein** Testergebnis.

Autoritativ und **unverändert** bleiben:

- [FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) — Validierungs-Autoritätsgrenze, Claim Boundary Set, Ebenen-Taxonomie, Traceability Contract, Ergebnisvokabular, Negativ-/Fail-Closed-Semantik, Evidenzbindung, Execution Gates
- [SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) — Fixture-Identity, -Revision und -Provenance
- [INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md) — Environment Identity und Profile
- [OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) — der zu prüfende Vertrag

Dieser Envelope erzeugt **kein** zweites Test-, Ergebnis- oder Evidenzvokabular.

## 2. Purpose

Sicherstellen, dass die Vertragsunterscheidungen des Slices — Absenz, Nichtverfügbarkeit, Berechtigungsverweigerung, Nichtunterstützung, Malformedness, Normalisierungsfehler, Staleness, Provenance-Ungültigkeit, heterogene All-Failure-Zusammensetzung (Envelope-Zusammensetzbarkeit) und Nicht-Mutation — **prüfbar formuliert** sind, **bevor** ein Erhebungsmechanismus, eine Sprache oder ein Zielzugriff gewählt wird. Ein Fall, dessen Beobachtungspunkt heute nicht benennbar ist, wäre später auch nicht belegbar.

## 3. Scope

Elf Mindestfälle (§7) · Traceability-Bindung (§5) · Ergebnisvokabular-Bindung (§6) · No-Mutation-Sentinel (§8) · Test-Claim-Grenze (§9) · Ausführungsvoraussetzungen (§10).

## 4. Non-Goals

- **Keine** Testimplementierung, **kein** Testcode, **keine** Fixtures, **keine** Harness, **kein** Runner, **kein** Framework.
- **Keine** Testausführung, **keine** Testevidenz, **keine** Coverage-Aussage.
- **Keine** Auswahl von Sprache, Runtime, Testframework, Lab-Technologie oder Erhebungsmechanismus.
- **Kein** realer Zielzugriff, **keine** reale Beobachtung, **keine** Netzwerk-, Credential- oder Secret-Nutzung.
- **Keine** Sicherheits-, Support-, Produktions- oder Accessibility-Aussage.

## 5. Traceability-Bindung

Jeder Fall dieses Envelopes ist ein **Testfall-Entwurf** im Sinne des [Test Case Traceability Contract](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §10. Er trägt hier bereits: Test Case Identity, Subject under Test, Source Contract, Declared Scope, Stimulus, Expected Outcome, Expected Prohibited Outcome, Observation Points, Evidence Requirements und Result Limitations. **Nicht** gefüllt und ausdrücklich offen bleiben: Fixture References (Fixtures existieren nicht), Environment/Profile (kein Lab bereitgestellt), Test Case Revision jenseits der Erstfassung.

Die Test-Case-Kennungen `OBS-LLH-TC-01` … `OBS-LLH-TC-11` sind **Testfall-Identitäten** — ausdrücklich **keine** Decision-, Risk-, ADR-, Capability- oder Work-Package-IDs und **kein** neues Register.

```text
test case references authority != test case becomes authority
traceability                   != requirement coverage
a referenced invariant         != a validated invariant
```

## 6. Ergebnis- und Ebenenbindung

Ergebniswerte kommen **ausschließlich** aus [Test Outcome Semantics](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §11: `not run` · `passed` · `failed` · `blocked` · `inconclusive` · `not applicable`. **Heute trägt jeder Fall dieses Envelopes `not run`.**

Erwartete `observation_outcome`-Werte kommen **ausschließlich** aus dem [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.2. Die beiden Vokabulare bleiben getrennt:

```text
test result        != observation_outcome
test passed        != observation succeeded
observation_outcome = the expectation a test case asserts against
```

Ebenenzuordnung folgt der bestehenden Taxonomie (§8 der Teststrategie). Eine Ebenenzuordnung, die von der konkreten Realisierung des Erhebungsmechanismus abhängt, ist als **`level assignment pending concrete implementation/runtime realization; P-3 selected`** gekennzeichnet und **nicht** vorentschieden. Die `P-3`-**Mechanismusklasse** ist entschieden ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5); die konkrete Implementierungs- und Runtime-Realisierung ist es **nicht**: die Sprach-/Runtime-**Klasse** ist zwar entschieden (**Go**, [Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4), aber es sind **keine** exakte Go-Version, **keine** beschaffte oder installierte Toolchain, **kein** Testframework, **keine** konkrete API und **keine** Implementierung ausgewählt oder vorhanden — die durch `A-13` entschiedene Build-/Toolchain-**Governance** (§12) ändert daran **nichts**; es existiert **keine** Testimplementierung und **keine** Testausführungsautorisierung. `Sprach-/Runtime-Klasse ausgewählt != konkrete Implementierungs-/Runtime-Realisierung`. Die Markierung bleibt daher — mit dieser präzisierten Formulierung — bestehen.

## 7. Mindestfälle

### 7.1 Übersicht

| ID | Fall | Ebene | Erwarteter `observation_outcome` | Heutiges Testergebnis |
| -- | ---- | ----- | -------------------------------- | --------------------- |
| `OBS-LLH-TC-01` | Normale Beobachtung | `TL-2` + `TL-3` | `success` | `not run` |
| `OBS-LLH-TC-02` | Quelle abwesend | `TL-3` (+ `TL-2`) | `source-absent` | `not run` |
| `OBS-LLH-TC-03` | Quelle nicht verfügbar | `TL-3`, *level assignment pending concrete implementation/runtime realization; `P-3` selected* | `source-unavailable` | `not run` |
| `OBS-LLH-TC-04` | Quelle malformed | `TL-2` + `TL-3` | `source-malformed` | `not run` |
| `OBS-LLH-TC-05` | Quelle stale | `TL-8` (+ `TL-2`) | `success` **mit** `freshness = stale` | `not run` |
| `OBS-LLH-TC-06` | Berechtigung verweigert | `TL-7` | `permission-denied` | `not run` |
| `OBS-LLH-TC-07` | Quelle/Plattform nicht unterstützt | `TL-3` (+ `TL-2`) | `unsupported` | `not run` |
| `OBS-LLH-TC-08` | Provenance ungültig | `TL-2` + `TL-7` | `field_observation_outcome = provenance-invalid` (Feld-Vokabular); Envelope nach §10.5 R3 `partial` **oder** nach R5 kein Envelope-Wert bei Emissions-Disposition `record-discarded` | `not run` |
| `OBS-LLH-TC-09` | Normalisierung fehlgeschlagen | `TL-2` | `normalization-failed` | `not run` |
| `OBS-LLH-TC-10` | No-Mutation-Sentinel | `TL-7` | ergebnisunabhängig — siehe §8 | `not run` |
| `OBS-LLH-TC-11` | Heterogene All-Failure-Zusammensetzung | `TL-2` + `TL-3`, *level assignment pending concrete implementation/runtime realization; `P-3` selected* | **kein** Envelope-Wert (§10.5 **R6**); Emissions-Disposition `record-discarded` | `not run` |

### 7.2 `OBS-LLH-TC-01` — Normale Beobachtung

- **Test intent:** Belegen, dass ein vollständiger, vertragskonformer Erhebungsvorgang alle Pflichtfelder mit Raw-Wert, normalisierter Repräsentation, vollständiger Provenance und `observed_at` erzeugt.
- **Setup concept:** Ein Beobachtungssubjekt, dessen sämtliche Vertragsquellen vorhanden, lesbar und vertragskonform formatiert sind. Fixture-Klasse `nominal`.
- **Expected `observation_outcome`:** `success` auf Envelope-Ebene; `success` für jedes Pflichtfeld.
- **Expected data semantics:** `target_id` gesetzt; `hostname`, `os_identity`, `os_release`, `kernel_release`, `architecture` je mit Raw-Wert **und** normalisierter Repräsentation; `observed_at` gesetzt (nicht `unknown`); `freshness` bestimmt; **kein** Feld mit Default oder Platzhalter.
- **Expected provenance behavior:** Je Feld `source_identity`, `source_class`, `collection_mechanism_class`, Raw-Referenz und Transformationsreferenz vorhanden und auflösbar; Raw bleibt neben Normalized erhalten.
- **Prohibited inference:** Ein `success` belegt **nicht**, dass die Werte korrekt sind, dass das Ziel gesund, vertrauenswürdig, verwaltet oder autorisiert ist, dass die Capability implementiert ist oder dass eine Integration `supported` wäre.
- **Future evidence requirement:** Vollständiger Observation-Record mit allen Pflichtfeldern, Raw- **und** Normalized-Werten, Provenance-Satz je Feld, `observed_at` mit Zeitquelle, plus die Revisionsbindung nach Teststrategie §18.

### 7.3 `OBS-LLH-TC-02` — Quelle abwesend

- **Test intent:** Belegen, dass eine nicht existierende Vertragsquelle als **Abwesenheit** geführt wird und nicht als Fehler, Default, Leerwert oder Nichtunterstützung.
- **Setup concept:** Beobachtungssubjekt, bei dem genau eine erwartete Quelle nicht existiert. Fixture-Klasse `missing`.
- **Expected `observation_outcome`:** Feld `source-absent`; Envelope `partial`, sofern andere Felder `success` sind.
- **Expected data semantics:** Das betroffene Feld trägt **keinen** Wert — kein Leerstring, keine Null, kein Platzhalter, keine Ableitung aus einem Nachbarfeld. Die Abwesenheit ist **positiv** dargestellt.
- **Expected provenance behavior:** Provenance existiert trotzdem: erwartete Quelle, Quellklasse, Erhebungsmechanismusklasse, Versuchszeit und Ausgang sind aufgezeichnet. Ein abwesendes Feld ist **kein** provenance-loses Feld.
- **Prohibited inference:** `source-absent` bedeutet **nicht**, dass das Ziel abwesend, defekt, nicht unterstützt oder gesund ist; `missing data != healthy`.
- **Future evidence requirement:** Record, der die erwartete Quelle **benennt**, die Abwesenheit als Ausgang ausweist und zeigt, dass kein Ersatzwert entstand.

### 7.4 `OBS-LLH-TC-03` — Quelle nicht verfügbar

- **Test intent:** Belegen, dass eine existierende, aber im Versuch nicht lesbare Quelle von Abwesenheit **und** von Berechtigungsverweigerung unterscheidbar bleibt.
- **Setup concept:** Quelle existiert, ist aber im Erhebungsversuch nicht lesbar (transiente Nichtverfügbarkeit, Zeitüberschreitung, Ressourcenkonflikt). Fixture-Klasse `unavailable`. *Level assignment pending concrete implementation/runtime realization; `P-3` selected*, weil die konkrete Unverfügbarkeitsform vom Erhebungsmechanismus abhängt.
- **Expected `observation_outcome`:** Feld `source-unavailable`; Envelope `partial`, sofern andere Felder `success` sind.
- **Expected data semantics:** Kein Wert, kein Default, kein Rückgriff auf einen früheren Wert **innerhalb desselben Vorgangs**. Ein Rückgriff auf einen früheren Wert wäre `last-known state` und gehört zum State-Modell, nicht zum Erhebungsergebnis.
- **Expected provenance behavior:** Ausgang, Versuchszeit, Erhebungsmechanismusklasse und — sofern vorhanden — Unverfügbarkeitsdetail als Provenance-Detail, **nicht** als eigener Outcome-Wert.
- **Prohibited inference:** `source-unavailable` ist **nicht** `permission-denied`, **nicht** `source-absent`, **nicht** „Ziel offline", **nicht** „Ziel nicht erreichbar" und **kein** Zielzustand.
- **Future evidence requirement:** Record, der Nichtverfügbarkeit von Abwesenheit und von Verweigerung **explizit** trennt, samt Beleg, dass kein Ersatz- oder Altwert eingesetzt wurde.

### 7.5 `OBS-LLH-TC-04` — Quelle malformed

- **Test intent:** Belegen, dass eine gelesene, aber nicht vertragskonform interpretierbare Quelle abgelehnt wird, ohne dass der Parser rät.
- **Setup concept:** Quelle vorhanden und lesbar, Inhalt strukturell oder semantisch nicht vertragskonform. Fixture-Klasse `malformed`.
- **Expected `observation_outcome`:** Feld `source-malformed`; Envelope `partial`, sofern andere Felder `success` sind.
- **Expected data semantics:** **Kein** normalisierter Wert. Der Raw-Inhalt bleibt — soweit disclosure-zulässig — als Raw referenzierbar; er wird **nicht** teilweise geraten, nicht repariert und nicht mit einem plausiblen Wert ersetzt.
- **Expected provenance behavior:** Ablehnungs-Record mit Grund und Feldbezug; Raw-Referenz bleibt erhalten; die fehlgeschlagene Interpretation ist als solche gekennzeichnet.
- **Prohibited inference:** `source-malformed` ist **nicht** `normalization-failed` (dort liegt ein gültiger Raw-Wert vor), **nicht** „Ziel fehlerhaft" und **kein** Sicherheitsbefund.
- **Future evidence requirement:** Record, der zeigt, dass **kein** Wert entstand, welcher Vertragsteil verletzt wurde und dass keine Parser-Annahme als beobachtete Evidenz geführt wird.

### 7.6 `OBS-LLH-TC-05` — Quelle stale

- **Test intent:** Belegen, dass Alter als eigenständige Dimension sichtbar bleibt und **nicht** in den Ausgang oder in eine Zustandsaussage kollabiert.
- **Setup concept:** Erfolgreiche Beobachtung, deren `observed_at` außerhalb der für die Entscheidung gebundenen Altersgrenze liegt. Fixture-Klasse `stale`. Die Grenze selbst ist heute `PROPOSED / UNACCEPTED` (Observation Contract §14) — der Fall prüft die **Sichtbarkeit**, nicht einen Zahlenwert.
- **Expected `observation_outcome`:** `success` — Staleness ist **kein** Erhebungsfehler. `freshness = stale` erscheint als eigenständiges Feld.
- **Expected data semantics:** Wert vorhanden, Raw und Normalized vorhanden, `observed_at` gesetzt, `freshness = stale` sichtbar und nicht überschrieben.
- **Expected provenance behavior:** `observed_at` und — sofern materiell getrennt — `received_at` bleiben unterscheidbar; die Ableitung der Freshness ist nachvollziehbar; Uhrenunsicherheit wird nicht wegnormalisiert.
- **Prohibited inference:** `stale != unavailable`, `stale != invalid`, `stale != failed`, `fresh != healthy`, `fresh != trusted`, `fresh != authorized`. Ein `success` mit `stale` ist **keine** Aussage über den aktuellen Zielzustand.
- **Future evidence requirement:** Record mit getrennten `observed_at`/`received_at`-Angaben, ausgewiesener Freshness und Beleg, dass Staleness weder den Ausgang verändert noch den Wert entfernt hat.

### 7.7 `OBS-LLH-TC-06` — Berechtigung verweigert

- **Test intent:** Als **Negativfall erster Klasse** das spezifizierte geschützte Verhalten bei verweigertem Lesezugriff **beobachten** — nicht lediglich eine Fehlermeldung feststellen.
- **Setup concept:** Quelle existiert, der Lesezugriff wird durch die Zielumgebung verweigert. Fixture-Klasse `permission-denied`.
- **Expected `observation_outcome`:** Feld `permission-denied`; Envelope `partial`, sofern andere Felder `success` sind.
- **Expected data semantics:** Kein Wert; **keine** Eskalation; **kein** zweiter Versuch mit erweiterten Rechten; **kein** alternativer privilegierter Pfad; kein Default.
- **Expected provenance behavior:** Denial-Record mit Quelle, Erhebungsmechanismusklasse, Zeit und **spezifiziertem Grund**; der Grund ist der spezifizierte Grund, nicht ein generischer Fehler.
- **Prohibited inference:** `permission-denied != source-unavailable`, `!= source-absent`, `!= target unavailable`, `!= target unhealthy`, `!= Angriff`. Eine erschienene Fehlermeldung ist **kein** beobachtetes Schutzverhalten.
- **Future evidence requirement:** Record des **beobachteten geschützten Verhaltens** gemäß Teststrategie §12: Zugriff blieb blockiert · **kein** Seiteneffekt am Ziel · Denial ist als Record sichtbar · der Grund ist der spezifizierte Grund · **keine** Rechteeskalation fand statt.

### 7.8 `OBS-LLH-TC-07` — Quelle oder Plattform nicht unterstützt

- **Test intent:** Belegen, dass Nichtunterstützung als eigene, nicht-defekthafte Kategorie geführt wird.
- **Setup concept:** Beobachtungssubjekt oder Quelltyp außerhalb des vertraglich unterstützten Umfangs. Fixture-Klasse `unsupported`.
- **Expected `observation_outcome`:** Feld oder Envelope `unsupported`.
- **Expected data semantics:** Kein Wert, kein Best-Effort-Parsing, keine Näherung, keine Teilinterpretation einer fremden Quellform.
- **Expected provenance behavior:** Record benennt, **welche** Vertragsbedingung nicht erfüllt ist — nicht bloß „nicht unterstützt".
- **Prohibited inference:** `unsupported != failed`, `!= absent`, `!= Defekt`, `!= Ziel ungeeignet`, `!= Supportstatus`. Nichtunterstützung im Vertrag ist **keine** Aussage über die Capability-Matrix.
- **Future evidence requirement:** Record mit benannter Vertragsbedingung und Beleg, dass kein Wert und keine Näherung entstand.

### 7.9 `OBS-LLH-TC-08` — Provenance ungültig

- **Test intent:** Belegen, dass ein Feld ohne verwertbare Provenance **fail-closed** ausscheidet und nicht als Beobachtung weiterläuft.
- **Setup concept:** Ein Feldergebnis, dessen Provenance unvollständig, widersprüchlich oder nicht auflösbar ist — etwa fehlende `source_identity`, fehlende `collection_mechanism_class` oder eine nicht auflösbare Raw-Referenz. Fixture-Klasse `provenance-invalid`.
- **Erwartete Vokabularzuordnung (drei getrennte Ebenen, [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.1):**
  - **Feldebene:** `field_observation_outcome = provenance-invalid` — ein Wert des **Feld**-Vokabulars (§10.3). Er ist ausdrücklich **kein** `observation_outcome`-Wert; das Envelope-Vokabular bleibt bei acht Werten.
  - **Emissionsebene:** `field-withheld` für das betroffene Feld — bzw. `record-discarded`, wo der Defekt den gesamten Datensatz trägt (§10.4). `record-discarded` ist eine **Verarbeitungs-/Ausgabe-Disposition**, **kein** Beobachtungsergebnis.
  - **Envelopeebene:** nach §10.5 — **R3** (einige Felder gültig, einige `provenance-invalid`) → Envelope `partial`; **R5** (alle sonst beobachteten Felder `provenance-invalid`) → **kein** Envelope-Wert wird behauptet, Datensatz `record-discarded`; **R1** (`target_id` fehlt oder ist provenance-ungültig) → ebenfalls kein Envelope-Wert, `record-discarded`.
- **Expected data semantics:** **Kein** normalisierter Wert; das Feld fließt **nicht** in eine CoreOps-Repräsentation ein; es entsteht kein „Wert ohne Herkunft". Bei R3 bleiben die übrigen Felder mit ihren eigenen Ausgängen unverändert erhalten.
- **Expected provenance behavior:** Der Provenance-Mangel selbst wird aufgezeichnet — welches Provenance-Element fehlte oder widersprüchlich war. Bei `record-discarded` bleibt ein Provenance-Defekt-Record für Audit erhalten; der Verwurf ist nicht spurlos.
- **Prohibited inference:** `provenance present != trusted`; umgekehrt bedeutet `provenance-invalid` **nicht**, dass der Wert falsch, das Ziel fehlerhaft oder die Quelle bösartig ist. Zusätzlich verbindlich:

```text
provenance invalid != target unhealthy
provenance invalid != source unavailable
provenance invalid != source absent
record discarded   != observation outcome
record discarded   != failed observation
partial            != failed
```

- **Future evidence requirement:** Record, der das fehlende Provenance-Element benennt, die drei Ebenen (Feldausgang, Emissions-Disposition, Envelope) getrennt ausweist und belegt, dass der Wert **nicht** verwendet wurde (fail-closed, nicht fail-open).

### 7.10 `OBS-LLH-TC-09` — Normalisierung fehlgeschlagen

- **Test intent:** Belegen, dass ein gültiger Raw-Wert erhalten bleibt, wenn die Abbildung auf die kanonische Repräsentation scheitert.
- **Setup concept:** Quelle gelesen, Raw-Wert vertragskonform gewonnen, die kanonische Abbildung ist jedoch nicht durchführbar (unbekannte Ausprägung, uneindeutige Zuordnung, fehlende Abbildungsregel).
- **Expected `observation_outcome`:** Feld `normalization-failed`; Envelope `partial`, sofern andere Felder `success` sind.
- **Expected data semantics:** Raw-Wert **bleibt erhalten, referenzierbar und sichtbar**; **kein** normalisierter Wert; **kein** Näherungswert; **kein** Default; **kein** stiller Fallback auf eine andere Abbildungsregel.
- **Expected provenance behavior:** Transformationsversuch, verwendete Regelreferenz und Fehlgrund sind aufgezeichnet; die Raw-Provenance bleibt unverändert erhalten.
- **Prohibited inference:** `normalization-failed != source-malformed` (der Raw-Wert war gültig), `!= Wert ungültig`, `!= Ziel fehlerhaft`. `normalization success != semantic validation` bleibt auch im Erfolgsfall gültig.
- **Future evidence requirement:** Record mit erhaltenem Raw-Wert, benannter Abbildungsregel, Fehlgrund und Beleg, dass **kein** Ersatzwert kanonisiert wurde.

### 7.11 `OBS-LLH-TC-10` — No-Mutation-Sentinel

Siehe §8. Der Fall ist heute `not run`; `P-2` bleibt **`NOT SATISFIED`**.

### 7.12 `OBS-LLH-TC-11` — Heterogene All-Failure-Zusammensetzung

- **Test intent:** Belegen, dass ein Erhebungsvorgang **ohne** jedes `success`-Feld und mit **verschiedenen** Fehlerursachen unter den verwertbaren Feldern **keinen** Envelope-Wert erfindet, sondern nach [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.5 **R6** ohne kanonische Observation endet — bei vollständig erhaltener feldweiser Unterscheidbarkeit.
- **Setup concept:** Ein Beobachtungssubjekt, bei dem **kein** Feld `success` erreicht, **mindestens zwei** Felder **nicht** `provenance-invalid` sind und diese Felder **verschiedene** Erhebungsausgänge tragen — etwa ein Feld `permission-denied` und ein Feld `source-absent`. Kombination der bestehenden Fixture-Klassen (`permission-denied` + `missing`); die konkrete Kombinierbarkeit ist erst nach `P-3` bestimmbar. *Level assignment pending concrete implementation/runtime realization; `P-3` selected*.
- **Expected `observation_outcome`:** **keiner.** Auf Envelope-Ebene wird **kein** Wert des Acht-Werte-Vokabulars (§10.2) behauptet.
- **Erwartete Vokabularzuordnung (drei getrennte Ebenen, [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.1):**
  - **Feldebene:** jedes Feld trägt seinen **eigenen** `field_observation_outcome` aus dem Neun-Werte-Feld-Vokabular (§10.3); die Ursachen bleiben **einzeln unterscheidbar** und werden **nicht** zusammengefasst.
  - **Envelopeebene:** **kein** `observation_outcome`. Ausdrücklich **kein** `partial`, **kein** synthetisches Aggregat, **kein** neunter Wert.
  - **Emissionsebene:** `record-discarded` (§10.4) — ausschließlich eine **Verarbeitungs-/Ausgabe-Disposition**, **kein** Beobachtungsergebnis und **kein** Löschvorgang.
- **Expected data semantics:** Der Datensatz wird **nicht** als kanonische Observation ausgegeben. Es entsteht **kein** Ersatz-, Default- oder Näherungswert und **keine** synthetische Envelope-Aussage. Es wird **keine** Präzedenz- und **keine** Schwereordnung zwischen den Fehlerursachen angewandt; kein Ausgang „gewinnt“. `partial` wird **nicht** aufgeweicht — der von `partial` geforderte `success`-Anteil fehlt.
- **Expected provenance behavior:** Die Provenance bleibt je Feld verfügbar; ein Raw-Wert bleibt dort verfügbar, wo einer entstanden ist; der Kontext des Erhebungsversuchs bleibt erhalten. Das erzeugte Diagnose-/Evidenzmaterial wird **nicht** stillschweigend gelöscht (§10.4). Die Retention erfolgt unter den **bestehenden** Semantiken — es wird **keine** Speicher-, Persistenz-, Aufbewahrungs- oder Löschtechnik geprüft oder vorausgesetzt.
- **Prohibited inference:** ausdrücklich verbindlich:

```text
kein Envelope-Wert  != partial
kein Envelope-Wert  != failed observation
record discarded    != observation outcome
record discarded    != failed observation
record discarded    != deletion
record discarded    != target state
diagnostic retained != canonical Observation emitted
diagnostic material != governed audit record
R6 decided          != productive code authorized
R6 decided          != implementation authorized
R6 decided          != test execution authorized
```

Aus R6 folgt **nichts** über das Ziel: weder Gesundheit noch Erreichbarkeit noch Defekt. Aus R6 folgt ebenso **keine** Autorisierung: die Regel ist eine Semantikfestlegung, **keine** Freigabe von Implementierung, Zielzugriff oder Testausführung.

- **Future evidence requirement:** Ein Record, der zeigt: (a) es wurde **kein** Envelope-Wert behauptet; (b) `partial` wurde **nicht** verwendet; (c) es wurde **keine** Präzedenz- oder Schwereordnung angewandt; (d) es wurde **keine** kanonische Observation ausgegeben; (e) `record-discarded` ist als reine Verarbeitungs-/Emissions-Disposition ausgewiesen; (f) jede feldweise Fehlerursache ist einzeln nachvollziehbar geblieben; (g) die Provenance ist weiterhin verfügbar; (h) ein vorhandener Raw-Wert ist weiterhin verfügbar; (i) das Diagnose-/Evidenzmaterial wurde **nicht** stillschweigend gelöscht; (j) es entstand **kein** neunter Envelope-Wert; (k) es wurde **keine** Aussage über die Zielgesundheit und **keine** Autorisierungsaussage abgeleitet. Die Revisionsbindung nach [Teststrategie](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §18 gilt unverändert.
- **Heutiges Ergebnis:** `not run`. Der Fall ist **nicht implementiert**, **nicht ausgeführt** und **nicht bestanden**; er erzeugt **keine** Evidenz. `not run != passed`.

## 8. No-Mutation-Sentinel

### 8.1 Gegenstand

Der Sentinel prüft nicht, ob eine Beobachtung „funktioniert", sondern ob der Beobachtungspfad **innerhalb seiner autorisierten Grenze** das beobachtete Ziel **nicht verändert**. Das ist die `P-2`-Frage.

### 8.2 Die entscheidende Grenze

```text
read-only != side-effect-free
no error observed != protected behavior observed
an error message appeared != protected behavior observed
absence of a detected change != absence of change
```

Ein Pfad kann read-only gegenüber den *Nutzdaten* sein und dennoch Seiteneffekte erzeugen — Zugriffszeiten, Logeinträge, Zählerstände, Ressourcenbindung, Sperren, Cache-Zustände, Auditspuren am Ziel. Ein „read-only"-Etikett ist deshalb **keine** No-Mutation-Evidenz.

### 8.3 Welche künftige Evidenz `P-2` erfüllen müsste

Erforderlich ist **beobachtetes geschütztes Verhalten**, nicht das Ausbleiben eines Fehlers. Kumulativ:

1. **Deklarierte Mutationsgrenze.** Vorab und schriftlich: welche Zielzustände als „unverändert" gelten müssen und welche Nebeneffekte innerhalb der autorisierten Grenze ausdrücklich zulässig sind. Ohne diese Vorab-Deklaration ist jedes spätere Ergebnis unfalsifizierbar.
2. **Vorher-/Nachher-Beobachtung an benannten Beobachtungspunkten** — mit derselben Erhebungsmethode, unter deklarierter Environment Identity, mit Zeitbindung und Clock-Unsicherheit.
3. **Ausgeführte Negativ-Familien der Teststrategie §12** — mindestens `unknown target`, `wrong target` und `missing execution authorization` als **implementierte und ausgeführte** Tests, deren Record das Schutzverhalten selbst zeigt: Aktion blieb blockiert · **kein** Seiteneffekt am Ziel · Denial als Record sichtbar · Grund ist der spezifizierte Grund.
4. **Fehlerpfad-Abdeckung.** Auch die Fälle `OBS-LLH-TC-02` bis `OBS-LLH-TC-09` **sowie** `OBS-LLH-TC-11` dürfen keine Mutation erzeugen — insbesondere kein Retry-Sturm, keine Rechteeskalation, keine Reparatur, kein Anlegen fehlender Quellen. Für `OBS-LLH-TC-11` gilt zusätzlich, dass auch die Nicht-Emission (`record-discarded`) und die Retention des erzeugten Diagnosematerials keine Mutation am beobachteten Ziel erzeugen dürfen. Alle genannten Fälle sind **künftige** Anforderungen: sie tragen heute `not run`, sind **nicht implementiert** und **nicht ausgeführt** — aus ihrer Nennung folgt **keine** bereits erzeugte Evidenz (`case enumerated != case implemented`).
5. **Deklarierte Restunsicherheit.** Welche Nebeneffekte mit der gewählten Methode **nicht** beobachtbar waren. Nicht beobachtbar ist nicht dasselbe wie nicht vorhanden.
6. **Unabhängige Reproduzierbarkeit** durch eine Person, die den Pfad nicht gebaut hat, unter deklarierter Revisionskombination (Teststrategie §18).
7. **Explizite Human-Maintainer-Autorisierung** sowohl des Zielzugriffs (`P-1`) als auch der Testausführung.

### 8.4 Was den Sentinel **nicht** erfüllt

```text
"es lief ohne Fehler durch"          != no mutation observed
"der Pfad ist als read-only gebaut"  != no mutation observed
"es wurde keine Änderung bemerkt"    != no mutation observed
eine Codereview-Aussage              != beobachtetes Verhalten
ein einzelner erfolgreicher Lauf     != Regressionsvertrauen
```

### 8.5 `P-2`-Lebenszyklus: Plan gegen Erfüllung

`P-2` hat **zwei getrennte Zustände**, die nie zusammengezogen werden dürfen:

| Zustand | Gate | Stand |
| ------- | ---- | ----- |
| **`P-2`-Evidenzmethode / -plan definiert** | Gate A — vor Produktivcode | **DEFINIERT** — die sieben kumulativen Anforderungen aus §8.3; die Human-Maintainer-Disposition dieser Definition ist **APPROVED** (Gate-A-Punkt `A-12` erfüllt) |
| **`P-2` erfüllt** | Gate B — vor realer Beobachtung | **`NOT SATISFIED`** |

```text
P-2 evidence plan      != P-2 satisfied
evidence plan approved != evidence produced
test design            != test execution
test design accepted   != test implemented
```

**Warum die Erfüllung erst nach der Implementierung möglich ist.** Die Anforderungen aus §8.3 setzen einen **ausführbaren** Beobachtungspfad, freigegebene Fixtures, eine deklarierte Umgebung sowie getrennte Ziel- und Ausführungsautorität voraus. Keine davon kann existieren, bevor überhaupt implementiert wurde. `P-2` als Vorbedingung für die Autorisierung produktiven Codes zu führen wäre daher **zirkulär**: die Evidenz verlangte genau das, was sie freigeben soll. Auf Gate-A-Ebene ist deshalb ausschließlich verlangt, dass die **Methode feststeht** — was mit §8.3 der Fall ist.

**`P-2` ist keine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.** Sie ist eine Voraussetzung dafür, ein reales Ziel zu beobachten.

### 8.6 Aktueller Stand

```text
P-2 evidence plan:    DEFINED  (§8.3, seven cumulative requirements)
P-2 plan disposition: HUMAN-MAINTAINER DISPOSITION APPROVED (Gate A, A-12)
P-2 satisfaction:     NOT SATISFIED
Tests implemented:    0
Tests executed:       0
No-mutation evidence: NONE
```

Dieses Dokument **beschafft** diese Evidenz nicht und darf nicht so gelesen werden, als hätte es sie beschafft. Ein definierter Evidenzweg ist keine Evidenz — und ein **dispositionierter** Evidenzweg ebenso wenig: die freigegebene Human-Maintainer-`A-12`-Disposition akzeptiert die Methode als **definiert** und erzeugt dadurch **keine** Evidenz.

## 9. Test-Claim-Grenze

Die bestehende Trennung der [Teststrategie](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §7 gilt unverändert und wird hier **nicht** erweitert:

```text
test planned     != test implemented
test implemented != test executed
test executed    != test passed
test result      != support status
test result      != production readiness
test result      != authoritative state
test result      != approval
test evidence    != validation evidence for a requirement
```

Zusätzlich für diesen Envelope:

```text
test envelope defined  != test suite exists
case enumerated        != case implemented
expected outcome named != outcome ever observed
eleven cases           != coverage complete
test strategy defined  != tests implemented
P-2 evidence plan      != P-2 satisfied
test design            != test execution
read-only              != side-effect-free
source tree decided          != test suite exists
test placement decided       != test implemented
testdata placement permitted != fixture authorized
test design defined          != test execution authorized
```

Und für die Autoritätsseite, damit dieses Dokument nicht als Vorbedingung für Gate A missverstanden wird:

```text
productive code authorization != target authorization
implementation authorization  != target authorization
source code present           != target access granted
test execution authority      != productive code authority
```

Das Claim Boundary Set ist **keine** Reifeleiter; `support`, `production readiness`, `security validation` und `operational validation` bleiben eigenständige Dimensionen mit eigener Autorität.

## 10. Ausführungsvoraussetzungen

Kein Fall dieses Envelopes ist heute ausführbar.

> **Gate-Zuordnung.** Die folgende Liste ist eine **Ausführungs**-Voraussetzungsliste. Sie beschreibt, was vor **Testausführung und realer Beobachtung** vorliegen muss — Gate B im Sinne von [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §16.2. **Keiner** ihrer Punkte ist eine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren. Auf Gate-A-Ebene ist aus diesem Dokument ausschließlich verlangt, dass Testdesign und `P-2`-Evidenzmethode **definiert** sind — beides ist mit §7 und §8.3 der Fall, und die Human-Maintainer-Disposition dieser Definition ist erteilt (Gate-A-Punkt `A-12` **erfüllt**). An der folgenden Ausführungs-Voraussetzungsliste ändert das **nichts**: jeder Punkt bleibt offen, jeder Fall bleibt `not run`.

> **Verhältnis zur Source-Tree-Entscheidung (`A-8`).** Die entschiedene Quellbaumstruktur ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §9.2) legt ausschließlich eine **künftige Platzierungspolitik** fest — colocated `*_test.go`, **kein** Top-Level-`tests/`, paketlokales `testdata/` nur als künftige Fixture-Ablage. An dieser Voraussetzungsliste ändert sie **nichts**: sie fügt **keinen** Punkt hinzu, erfüllt **keinen** Punkt und hebt **keinen** Punkt auf. Sie ist **keine** Testimplementierungs-, Testausführungs- oder Fixture-Autorisierung, wählt **kein** Testframework und legt **keine** Datei und **kein** Verzeichnis an. Jeder Punkt 3 bis 7 bleibt offen, jeder Fall bleibt `not run`.

Eine spätere Ausführung erfordert **kumulativ** (konsistent mit Teststrategie §20, ergänzt um die Slice-Preconditions):

1. `P-3` — Entscheidung über Erhebungsmechanismus und Transport (**`SELECTED`**) — **erfüllt**. Der Human Maintainer hat die **Mechanismusklasse** entschieden ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5): Option A primär, Option B ausschließlich ergänzend, Option C nicht Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport. Das ist **keine** Ausführungsautorisierung und **keine** Auswahl von Pfad, API, Bibliothek, Werkzeug, Sprache oder Runtime.
2. Sprach-/Runtime-Entscheidung (**`SELECTED` — Go**) — **erfüllt**. Der Human Maintainer hat die **Sprach-/Runtime-Klasse** für den ersten Observe-Slice entschieden ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4). Das ist **keine** Ausführungsautorisierung und **keine** Auswahl von Go-Version, Distribution, Toolchain, Testframework, konkreter API oder Abhängigkeit.
3. Implementierung des Beobachtungspfads (**nicht autorisiert**).
4. Freigegebene Fixtures mit Identity, Revision und Provenance (**existieren nicht**).
5. Bereitgestelltes Lab-Environment mit deklarierter Environment Identity (**nicht bereitgestellt**).
6. `P-1` — separate Human-Maintainer-Zielautorisierung (**`NOT AUTHORIZED`**).
7. Explizite Human-Maintainer-Autorisierung der **Testausführung** (**nicht erteilt**).

Die Punkte 1 und 2 sind **erfüllt**; die Punkte 3 bis 7 bleiben **offen**. Die Ausführungs-Voraussetzungsmenge bleibt damit **insgesamt unerfüllt**. Solange auch nur eine dieser Voraussetzungen offen ist, bleibt jeder Fall `not run` — und `not run != passed`. Es sind **0** Tests implementiert und **0** Tests ausgeführt; `P-2` bleibt **`NOT SATISFIED`**.

## 11. Compatibility

Additiv. **Keine** Änderung an Teststrategie, Fixture-Governance, Lab-Modell, Ergebnisvokabular, Claim Boundary Set, Decision Index, Risk Register oder Capability Matrix. **Kein** neues Ergebnis-, Evidenz- oder Coverage-Modell. **Kein** ADR. Breaking-Change-Potenzial: gering.

## 12. Open Questions

- Konkrete Fixture-Klassen und -Revisionen je Fall — **offen**. Die `P-3`-Mechanismusklasse ist entschieden; die konkreten Fixture-Klassen und -Revisionen sind damit **nicht** bestimmt, **nicht** bereitgestellt und **nicht** freigegeben. Auch die Source-Tree-Entscheidung bestimmt sie **nicht**: sie erlaubt lediglich ein **paketlokales** `testdata/` als **künftige** Ablage, falls und sobald Fixture- und Testautorität besteht — `testdata placement permitted != fixture authorized`. Ihre Festlegung bleibt spätere, ausdrücklich zu autorisierende Arbeit.
- Environment-Profil für einen lokalen Linux-Host im Lab-Modell — **offen**. Die `P-3`-Entscheidung liefert **kein** Lab-Profil; es ist weder definiert noch bereitgestellt noch autorisiert.
- Ob der Sentinel zusätzlich eine Langzeit-/Wiederholungsdimension benötigt (Regressionsvertrauen).
- Welche Nebeneffekte am Ziel mit der später gewählten Methode überhaupt beobachtbar sind.

## 13. Next Decision

**Aktueller Stand.** `CO-WP-032` ist **abgeschlossen und remote integriert**: Nova Final Review `GO`; Human-Maintainer-Integrationscommit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht; die anschließende Post-Integrations-Reconciliation steht auf `390da5cc8629dfa9cbea990c0a3c4ba4cb156e9b`. Die Repository-, Staging-, Commit- und Push-Gates für `CO-WP-032` sind damit **erledigt** und **nicht** mehr der nächste Schritt.

Der vorliegende docs-only Nachtrag (`OBS-LLH-TC-11` und diese Stands-Korrektur) liegt im Arbeitsverzeichnis. Staging, Commit, Push, Tag und jede weitere Repository-Aktion liegen unverändert **ausschließlich** beim Human Maintainer; ein künftiger Integrationsstand wird hier **nicht** vorweggenommen und **kein** Commit-Identifier dafür vorhergesagt.

**`P-3` ist entschieden.** Der Human Maintainer hat den lokalen Erhebungsmechanismus als **Mechanismusklasse** ausgewählt ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5). Gate-A-Punkt `A-6` ist damit **erfüllt**.

```text
P-3:                   SELECTED — Mechanismusklasse
                       Option A primär · Option B ausschließlich ergänzend ·
                       Option C nicht Standardweg · Option D/E ausgeschlossen ·
                       kein Netzwerktransport
A-6:                   erfüllt
Gate A:                7 von 14 erfüllt, 0 teilweise, 7 offen — nicht passiert
                       (Stand dieser Anwendung; siehe die folgende
                       Sprach-/Runtime-Anwendung)
Gate B:                0 von 8 — nicht passiert
```

**Die Sprach-/Runtime-Entscheidung ist getroffen.** Der Human Maintainer hat die **Sprach-/Runtime-Klasse** für den ersten Observe-Slice ausgewählt ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4). Gate-A-Punkt `A-7` ist damit **erfüllt**.

```text
Sprache / Runtime:     SELECTED — Go (Sprach-/Runtime-Klasse, erster Observe-Slice)
                       Rust bleibt dokumentierte stärkste Alternative
NICHT ausgewählt       Go-Version · Distribution · Toolchain · Testframework ·
(durch A-7):           Modul-/Paketlayout · konkrete API · Quellpfad ·
                       Abhängigkeit · Build/Packaging ·
                       breiterer Technologie-Stack
                       (das Modul-/Paketlayout ist inzwischen durch die
                       folgende Source-Tree-Entscheidung entschieden; die
                       übrigen bleiben unausgewählt)
A-7:                   erfüllt
Gate A:                8 von 14 erfüllt, 0 teilweise, 6 offen — nicht passiert
                       (Stand dieser Anwendung; siehe die folgende
                       Source-Tree-Anwendung)
                       offen blieben: A-8 · A-9 · A-10 · A-11 · A-13 · A-14
Gate B:                0 von 8 — nicht passiert
```

Auch diese Entscheidung autorisiert **keine** Testimplementierung, **keine** Testausführung, **keinen** Zielzugriff und **keine** reale Beobachtung, erfüllt `P-2` **nicht** und ändert an diesem Envelope inhaltlich **nichts**: Testintentionen, erwartete Ausgänge, Ebenenzuordnungen, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen sowie die sieben `P-2`-Evidenzanforderungen bleiben **unverändert**; **0** Tests implementiert, **0** Tests ausgeführt, jeder Fall unverändert `not run`.

Die Entscheidung benennt **keinen** konkreten Quellpfad, **keine** API, **keine** Bibliothek, **kein** Kommando, **kein** Werkzeug, **keine** Sprache und **keine** Runtime. Sie autorisiert **keine** Testimplementierung, **keine** Testausführung, **keinen** Zielzugriff und **keine** reale Beobachtung, erfüllt `P-2` **nicht** und ändert an diesem Envelope inhaltlich **nichts**: **0** Tests implementiert, **0** Tests ausgeführt, jeder Fall unverändert `not run`.

**Die Source-Tree-Entscheidung ist getroffen.** Der Human Maintainer hat die Quellbaumstruktur für den ersten Observe-Slice ausgewählt ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §9.2). Gate-A-Punkt `A-8` ist damit **erfüllt**.

```text
Source Tree:           DECIDED — repository-verwurzelte Go-Struktur
                       cmd/coreops/ · internal/observe/ ·
                       internal/observe/collect/
                       Normalisierung zunächst ohne eigenes Paket innerhalb
                       internal/observe/
Testplatzierung:       künftig colocated *_test.go
                       kein Top-Level-tests/
                       paketlokales testdata/ nur als künftige Fixture-Ablage,
                       falls und sobald Fixture-/Testautorität besteht
NICHT ausgewählt:      Testframework · konkrete Testdateinamen ·
                       In-Package- gegen External-Package-Testkonvention ·
                       Fixture-Inhalt · Fixture-Identität · Modulpfad ·
                       Go-Version · Distribution · Toolchain · Abhängigkeit ·
                       Build/Packaging
NICHT angelegt:        Verzeichnis · Datei · cmd/ · internal/ · testdata/ ·
                       go.mod · go.sum · Testdatei · Fixture
A-8:                   erfüllt
Gate A:                9 von 14 erfüllt, 0 teilweise, 5 offen — nicht passiert
                       (Stand dieser Anwendung; siehe die folgende
                       Dependency-Admission-Anwendung)
                       offen blieben: A-9 · A-10 · A-11 · A-13 · A-14
Gate B:                0 von 8 — nicht passiert
```

Auch diese Entscheidung autorisiert **keine** Testimplementierung, **keine** Testausführung, **keine** Fixture, **keinen** Zielzugriff und **keine** reale Beobachtung, erfüllt `P-2` **nicht** und ändert an diesem Envelope inhaltlich **nichts**: Testfall-Identitäten, Testintentionen, erwartete und verbotene Ausgänge, Ebenenzuordnungen, Ergebnis- und Ausgangsvokabulare, Emissions-Dispositionen sowie die sieben `P-2`-Evidenzanforderungen bleiben **unverändert**; **0** Tests implementiert, **0** Tests ausgeführt, jeder Fall unverändert `not run`. Sie entscheidet ausschließlich eine **künftige Platzierungspolitik** — `test placement decided != test implemented`.

**Die Dependency-Admission-Entscheidung ist getroffen.** Der Human Maintainer hat das Dependency-Admission-Gate **in Kraft gesetzt** ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §10). Gate-A-Punkt `A-9` ist damit **erfüllt**.

```text
Dependency Admission:  IN FORCE — Disposition APPROVE OPTION B WITH NOVA
                       BOUNDARY CLARIFICATIONS
Zugelassene Dritt-     0
abhängigkeiten:
Konkret zugelassen:    keine
Erster Observe-Slice:  ZERO-THIRD-PARTY-DEFAULT
Testabhängigkeits-     Testeinrichtungen der Go-Standardbibliothek benötigen
grenze:                keine individuelle Drittzulassung; ein separat
                       bezogenes Testframework, -werkzeug oder eine
                       -bibliothek benötigt die vollständige A-9-Zulassung;
                       transitive Testabhängigkeiten sind im Scope
NICHT ausgewählt:      Testframework · Testbibliothek · Testwerkzeug ·
                       Assertions-/Mock-/Fixture-Bibliothek · Build-Werkzeug ·
                       Codegenerator · Linter · Analyzer · Vendor-Inhalt
NICHT zugelassen:      jede konkrete Abhängigkeit
A-9:                   erfüllt
Gate A:                10 von 14 erfüllt, 0 teilweise, 4 offen — nicht passiert
                       (Stand dieser Anwendung; siehe die folgende
                       NEW-8-/A-10-Anwendung)
                       offen blieben: A-10 · A-11 · A-13 · A-14
Gate B:                0 von 8 — nicht passiert
```

Für diesen Envelope folgt daraus **keine** inhaltliche Änderung. Die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die Ergebnis- und Ausgangsvokabulare, die Emissions-Dispositionen, die Ebenenzuordnungen, die `P-2`-Evidenzanforderungen aus §8.3 und die No-Mutation-Anforderungen bleiben **unverändert**. Die Inkraftsetzung autorisiert **keine** Testimplementierung, **keine** Testausführung, **keine** Fixture, **keinen** Zielzugriff und **keine** reale Beobachtung, erfüllt `P-2` **nicht** und lässt **keine** Testabhängigkeit zu: **0** Tests implementiert, **0** Tests ausgeführt, jeder konzeptionelle Fall unverändert `not run`. Es wird **kein** konkretes Testframework und **keine** konkrete Testbibliothek benannt, empfohlen oder ausgewählt.

```text
Testabhängigkeitsgrenze geklärt != Testframework ausgewählt
Policy in Kraft                 != Abhängigkeit zugelassen
stdlib-Testeinrichtung          != separat bezogenes Testframework
test-only                       != automatisch vertrauenswürdig
transitiv                       != außerhalb des Scopes
A-9 erfüllt                     != Testimplementierung autorisiert
A-9 erfüllt                     != Testausführung autorisiert
A-9 erfüllt                     != P-2 erfüllt
```

**Die `NEW-8`-Entscheidung (`README` / `LICENSE`) ist getroffen.** Der Human Maintainer hat das künftige Veröffentlichungs- und Lizenzmodell dispositioniert ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §13). Gate-A-Punkt `A-10` ist damit **erfüllt**. Diese Zeilen sind **Current-State-Kontext** und **keine** Testsemantik.

```text
README.md:             SHALL EXIST — ERSTELLT / VORHANDEN (durch die davon
                       getrennte, gesondert autorisierte Publikations-
                       Realisierung; NICHT durch A-10 selbst)
README-Zeitpunkt:      eigene, ausdrücklich autorisierte Repository-Änderung
                       VOR dem ersten produktiven Source-Commit
README-Autorität:      öffentliche Zusammenfassung ohne normative Autorität
LICENSE:               SHALL EXIST — ERSTELLT / VORHANDEN (standardmaessiger,
                       unveraenderter Apache-2.0-Text; NICHT durch A-10 selbst)
Veröffentlichungsmodus: PUBLIC / OPEN SOURCE
Outbound-Lizenz:       Apache-2.0 (standardmäßiger, unveränderter Text;
                       hier NICHT wiedergegeben; das LICENSE-Artefakt liegt
                       inzwischen im Repository-Wurzelverzeichnis vor)
Beitragsprogramm:      NICHT AKTIV — CONTRIBUTING.md nicht autorisiert,
                       CLA keine, DCO keiner
NOTICE / SECURITY.md:  nicht autorisiert · nicht erstellt
Markenrichtlinie:      zurückgestellt · keine eingetragene Marke beansprucht
NICHT autorisiert:     Testimplementierung · Testausführung · Fixture ·
                       Zielzugriff · reale Beobachtung · Veröffentlichung
                       produktiven Quellcodes · jede Dateierstellung
                       außerhalb von README.md und LICENSE
A-9:                   unverändert erfüllt — 0 Drittabhängigkeiten zugelassen
A-10:                  erfüllt
Gate A:                11 von 14 erfüllt, 0 teilweise, 3 offen — nicht passiert
                       (durch die spätere A-13-Anwendung fortgeschrieben)
                       offen blieben: A-11 · A-13 · A-14
Gate B:                0 von 8 — nicht passiert
```

Für diesen Envelope folgt daraus **keine** inhaltliche Änderung. Die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die Beobachtungsausgangs- und Ergebnisvokabulare, die Emissions-Dispositionen, die Ebenenzuordnungen, die `P-2`-Evidenzmethode aus §8.3, die No-Mutation-Anforderungen und die aus `A-9` folgende Testabhängigkeitsgrenze bleiben **unverändert**. Es ist weiterhin **kein** Testframework vorhanden, ausgewählt oder zugelassen; **0** Tests implementiert, **0** Tests ausgeführt, jeder konzeptionelle Fall unverändert `not run`.

**Publikations-Realisierung (Current-State-Kontext).** `README.md` und `LICENSE` sind inzwischen durch eine davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung **erstellt und vorhanden**; `LICENSE` trägt den standardmäßigen, **unveränderten** Apache-2.0-Text. Auch das ist ausschließlich **Current-State-Kontext** und **keine** Testsemantik: Testfallidentitäten, Testintentionen, erwartete und verbotene Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die `P-2`-Evidenzmethode aus §8.3, die No-Mutation-Anforderungen, die aus `A-9` folgende Testabhängigkeitsgrenze und die Ebenenzuordnungen bleiben **unverändert**. Unverändert bleiben ebenso: **0** Tests implementiert · **0** Tests ausgeführt · `P-2` **`NOT SATISFIED`** · `P-1` **`NOT AUTHORIZED`** · Testimplementierung und Testausführung **NOT AUTHORIZED** · Fixtures und Lab **nicht bereitgestellt** · **0** zugelassene Drittabhängigkeiten · **kein** Testframework ausgewählt. Bindend: `README erstellt != Test implementiert` · `LICENSE erstellt != Testausführung autorisiert` · `Publikationsartefakt vorhanden != P-2 erfüllt` · `Publikationsartefakt vorhanden != Zielzugriff autorisiert` · `Publikationsartefakt vorhanden != Fixture freigegeben`.

```text
A-10 erfüllt   != README erstellt
A-10 erfüllt   != LICENSE erstellt
A-10 erfüllt   != NOTICE erstellt
A-10 erfüllt   != Testframework ausgewählt
A-10 erfüllt   != Testimplementierung autorisiert
A-10 erfüllt   != Testausführung autorisiert
A-10 erfüllt   != Fixture freigegeben
A-10 erfüllt   != P-1 erteilt
A-10 erfüllt   != P-2 erfüllt
A-10 erfüllt   != Abhängigkeit zugelassen
A-10 erfüllt   != Gate A passiert
```

**Die Build-/Packaging-Entscheidung ist getroffen.** Der Human Maintainer hat die Build-, Reproduzierbarkeits-, Offline-Build-, Toolchain-Governance- und Artefaktidentitätsfragen für den ersten Observe-Slice auf **Governance-Ebene** dispositioniert ([Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §12). Gate-A-Punkt `A-13` ist damit **erfüllt**. Diese Zeilen sind **Current-State-Kontext** und **keine** Testsemantik.

```text
Reproduzierbarkeit:    DECLARED-INPUT / DETERMINISTIC-IDENTITY-Modell;
                       Bit-für-Bit-Gleichheit ist DESIGN-ZIEL und
                       ausdrücklich KEINE Gate-A-Evidenzaussage
Offline-Build:         ERFORDERLICH — GOTOOLCHAIN=local, GOPROXY=off,
                       zusätzlich eine fail-closed netzwerkfreie
                       Ausführungsgrenze (konkrete Realisierung
                       ZURÜCKGESTELLT); kein Proxy, kein Mirror, kein
                       Vendor-Mechanismus, kein Paket-Repository ausgewählt
Build-Mechanismus:     standardmäßige Go-Toolchain unmittelbar; kein Make,
                       kein Task Runner, kein Wrapper, kein Codegenerator,
                       kein Build-only-Drittwerkzeug zugelassen
Go-Distribution:       Herkunftsklasse offizielles Upstream-Go-Projekt-
                       Release; exakte Go-Version ZURÜCKGESTELLT;
                       Toolchain-Beschaffung und -Installation
                       NICHT AUTORISIERT
CGO / Build-Flags:     CGO_ENABLED=0 · -trimpath · -buildvcs=false;
                       Quellrevision über Provenance
Modulmodell:           Go-Modulmechanismus AUSGEWÄHLT; go.mod und go.sum
                       NICHT ANGELEGT / NICHT AUTORISIERT; Modulpfad
                       ZURÜCKGESTELLT
Artefakt:              Klasse Anwendungs-/Service-Artefakt nach dem
                       bestehenden Artifact-Identity-Modell; Digest später
                       erforderlich, Algorithmus ZURÜCKGESTELLT;
                       ausführbarer Basisname ZURÜCKGESTELLT
Build-Ziel:            Ziel-OS-Klasse Linux; GOARCH ZURÜCKGESTELLT
Erste Artefaktform:    natives Go-Executable; Container und CI
                       ZURÜCKGESTELLT
SBOM / Provenance:     Behandlung DEFINIERT — NICHT ERZEUGT
Testwerkzeuge:         durch A-13 KEIN Testframework, KEINE Testbibliothek
                       und KEIN Testwerkzeug ausgewählt oder zugelassen;
                       separat bezogene Werkzeuge bleiben A-9-pflichtig
NICHT autorisiert:     Toolchain-Beschaffung · Toolchain-Installation ·
                       Build-Ausführung · Artefakterstellung ·
                       Digest-/Provenance-/SBOM-Erzeugung ·
                       Testimplementierung · Testausführung · Fixture ·
                       Zielzugriff · reale Beobachtung · jede
                       Dateierstellung
A-9:                   unverändert erfüllt — 0 Drittabhängigkeiten zugelassen
A-13:                  erfüllt
Gate A:                12 von 14 erfüllt, 0 teilweise, 2 offen — nicht passiert
                       (durch die spätere A-11-Anwendung fortgeschrieben)
                       offen blieben: A-11 · A-14
Gate B:                0 von 8 — nicht passiert
```

Für diesen Envelope folgt daraus **keine** inhaltliche Änderung. Die Identitäten `OBS-LLH-TC-01`…`OBS-LLH-TC-11`, ihre Testintentionen, erwarteten und verbotenen Ausgänge, die **R6**- und `OBS-LLH-TC-11`-Semantik, die Beobachtungsausgangs- und Ergebnisvokabulare, die Record-Discarded- und Provenance-Semantik, das Testergebnisvokabular, die Emissions-Dispositionen, die Ebenenzuordnungen, die `P-2`-Evidenzmethode aus §8.3, die No-Mutation-Anforderungen und die aus `A-9` folgende Testabhängigkeitsgrenze bleiben **unverändert**. **Kein** konkretes Testkommando und **keine** Toolchain-Ausführung ist autorisiert; **0** Tests implementiert, **0** Tests ausgeführt, jeder konzeptionelle Fall unverändert `not run`.

```text
A-13 erfüllt   != Toolchain beschafft
A-13 erfüllt   != Toolchain installiert
A-13 erfüllt   != Build ausgeführt
A-13 erfüllt   != Artefakt erzeugt
A-13 erfüllt   != Testframework ausgewählt
A-13 erfüllt   != Testimplementierung autorisiert
A-13 erfüllt   != Testausführung autorisiert
A-13 erfüllt   != Fixture freigegeben
A-13 erfüllt   != P-1 erteilt
A-13 erfüllt   != P-2 erfüllt
A-13 erfüllt   != Abhängigkeit zugelassen
A-13 erfüllt   != Gate A passiert
```

**Die ADR-/Decision-Disposition ist getroffen.** Der Human Maintainer hat die ADR-Behandlung der sieben Punkte aus [Transition Prerequisites](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §14 ausdrücklich dispositioniert; Gate-A-Punkt `A-11` ist damit **erfüllt**. Für diesen Envelope folgt daraus **keine** inhaltliche Änderung.

```text
A-11:                  erfüllt — ADR-Behandlung DECIDED, Geltung
                       ausschließlich erster Observe-Slice
ADR-Behandlung:        P-3 T2 · Sprache/Runtime T2 · Source Tree T4 ·
                       Dependency-Admission T4 · NEW-8 T4 ·
                       Build/Packaging T5 · R6 T4
                       T2 = ADR-Record zurückgestellt bis Relevanzauslöser
                       T4 = kein gesonderter ADR erforderlich
                       T5 = bestehende Kandidaten-/Decision-Gegenstände
                            bleiben getrennt
Akzeptierte ADRs:      0
ADR-Dateien:           0
ADR-Nummer:            keine vergeben · keine reserviert
ADR-Ort / -Vorlage:    nicht entschieden
Decision Index:        unverändert
Risk Register:         unverändert
Testsemantik:          UNVERÄNDERT — TC-Identitäten, R1–R6, TC-11,
                       Ergebnis- und Dispositionsvokabulare,
                       P-2-Evidenzmethode und No-Mutation-Evidenz
                       bleiben unberührt
NICHT autorisiert:     Implementierung · Produktivcode · Quellbaumanlage ·
                       go.mod / go.sum · Toolchain-Beschaffung ·
                       Build-Ausführung · Testimplementierung ·
                       Testausführung · Fixture · Zielzugriff
Gate A:                13 von 14 erfüllt, 0 teilweise, 1 offen — nicht passiert
                       offen: A-14
Gate B:                0 von 8 — nicht passiert
```

```text
A-11 erfüllt   != ADR erzeugt
A-11 erfüllt   != ADR akzeptiert
A-11 erfüllt   != ADR-Nummer vergeben
A-11 erfüllt   != Decision Index geändert
A-11 erfüllt   != Risk Register geändert
A-11 erfüllt   != Testsemantik geändert
A-11 erfüllt   != Testframework ausgewählt
A-11 erfüllt   != Testimplementierung autorisiert
A-11 erfüllt   != Testausführung autorisiert
A-11 erfüllt   != Fixture freigegeben
A-11 erfüllt   != P-1 erteilt
A-11 erfüllt   != P-2 erfüllt
A-11 erfüllt   != Abhängigkeit zugelassen
A-11 erfüllt   != A-14 erfüllt
A-11 erfüllt   != Gate A passiert
ADR zurückgestellt != ADR-Pflicht erloschen
```

**Im Übrigen unverändert und ausdrücklich nicht erteilt:**

```text
Testimplementierung:   NOT AUTHORIZED
Testausführung:        NOT AUTHORIZED
Fixtures:              nicht bereitgestellt
Lab-Environment:       nicht bereitgestellt
P-1 / Zielzugriff:     NOT AUTHORIZED
P-2 Evidenzplan:       DEFINED / HUMAN-MAINTAINER-DISPOSITION APPROVED
P-2 Erfüllung:         NOT SATISFIED
Sprache / Runtime:     SELECTED — Go (nur Klasse; keine Version, keine
                       Toolchain, kein Testframework, keine Implementierung)
Source Tree:           DECIDED — nicht angelegt; nur künftige Platzierungs-
                       politik (colocated *_test.go; kein Top-Level-tests/;
                       paketlokales testdata/ nur künftig und nur bei später
                       bestehender Fixture-/Testautorität)
Dependency Admission:  IN FORCE — 0 Drittabhängigkeiten zugelassen,
                       keine konkret zugelassen, kein Testframework und
                       keine Testbibliothek ausgewählt
README / LICENSE:      SHALL EXIST — beide ERSTELLT / VORHANDEN; Erstellung
                       nicht durch A-10, sondern durch die davon getrennte,
                       gesondert autorisierte Publikations-Realisierung;
                       Veröffentlichungsmodus PUBLIC / OPEN SOURCE;
                       Outbound-Lizenz Apache-2.0; Beitragsprogramm
                       NICHT AKTIV
Build / Packaging:     DECIDED (A-13 erfüllt) — nur Governance-Ebene:
                       Reproduzierbarkeits-, Offline-Build-, Toolchain-
                       Governance-, Modul-, Artefaktidentitäts-, Ziel- und
                       Packaging-Disposition entschieden; exakte Go-Version,
                       Modulpfad, Executable-Name, GOARCH, Digest-
                       Algorithmus, SBOM-Format, Container und CI
                       ZURÜCKGESTELLT; kein Build ausgeführt, kein Artefakt,
                       kein Digest, keine Provenance, keine SBOM
ADR-Behandlung:        DECIDED für den ersten Observe-Slice (A-11 erfüllt) —
                       P-3 T2 · Sprache/Runtime T2 · Source Tree T4 ·
                       Dependency-Admission T4 · NEW-8 T4 ·
                       Build/Packaging T5 · R6 T4;
                       zurückgestellt ist NICHT erloschen
Akzeptierte ADRs:      0
ADR-Dateien:           0
ADR-Nummer:            keine vergeben · keine reserviert
Decision Index:        unverändert
Risk Register:         unverändert
Gate A:                13 von 14 erfüllt, 0 teilweise, 1 offen —
                       nicht passiert (offen: A-14)
Gate B:                0 von 8 — nicht passiert
Tests implementiert:   0
Tests ausgeführt:      0
Implementierung:       NOT AUTHORIZED
CO-WP-033:             nicht erzeugt / nicht reserviert
```

Testimplementierung, Testausführung, Fixture- und Lab-Bereitstellung bleiben getrennte, ausdrücklich zu autorisierende spätere Arbeit. **Dieser Envelope autorisiert keine davon.** Die Aufnahme von `OBS-LLH-TC-11` ist eine **Testfalldefinition** und **keine** Implementierung, **keine** Ausführung und **keine** Evidenz: `case enumerated != case implemented`.
