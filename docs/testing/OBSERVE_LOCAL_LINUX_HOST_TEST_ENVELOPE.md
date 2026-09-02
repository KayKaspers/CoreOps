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
> No-Mutation Evidence (`P-2`): method/plan **DEFINED** (§8.3) · satisfaction **NOT SATISFIED** — Gate B, downstream of implementation and execution
> Collection Mechanism (`P-3`): Not selected
> Language / Runtime: Not selected
> Test Execution Authorization: Not granted
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)
> Nachträgliche docs-only Änderung: konzeptioneller Mindestfall `OBS-LLH-TC-11` (heterogene All-Failure-Zusammensetzung, [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) §10.5 **R6**) ergänzt; §13 auf den tatsächlichen Stand korrigiert. Grundlage ist die ausdrückliche Human-Maintainer-Entscheidung zur heterogenen All-Failure-Zusammensetzung; **kein** Decision-Identifier vergeben, **kein** ADR, **kein** neues und **kein** reserviertes Work Package. `OBS-LLH-TC-11` ist **nicht implementiert**, **nicht ausgeführt** und trägt `not run`.

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

Ebenenzuordnung folgt der bestehenden Taxonomie (§8 der Teststrategie). Eine Ebenenzuordnung, die vom noch nicht entschiedenen Erhebungsmechanismus abhängt, ist als **`level assignment pending P-3`** gekennzeichnet und **nicht** vorentschieden.

## 7. Mindestfälle

### 7.1 Übersicht

| ID | Fall | Ebene | Erwarteter `observation_outcome` | Heutiges Testergebnis |
| -- | ---- | ----- | -------------------------------- | --------------------- |
| `OBS-LLH-TC-01` | Normale Beobachtung | `TL-2` + `TL-3` | `success` | `not run` |
| `OBS-LLH-TC-02` | Quelle abwesend | `TL-3` (+ `TL-2`) | `source-absent` | `not run` |
| `OBS-LLH-TC-03` | Quelle nicht verfügbar | `TL-3`, *level assignment pending `P-3`* | `source-unavailable` | `not run` |
| `OBS-LLH-TC-04` | Quelle malformed | `TL-2` + `TL-3` | `source-malformed` | `not run` |
| `OBS-LLH-TC-05` | Quelle stale | `TL-8` (+ `TL-2`) | `success` **mit** `freshness = stale` | `not run` |
| `OBS-LLH-TC-06` | Berechtigung verweigert | `TL-7` | `permission-denied` | `not run` |
| `OBS-LLH-TC-07` | Quelle/Plattform nicht unterstützt | `TL-3` (+ `TL-2`) | `unsupported` | `not run` |
| `OBS-LLH-TC-08` | Provenance ungültig | `TL-2` + `TL-7` | `field_observation_outcome = provenance-invalid` (Feld-Vokabular); Envelope nach §10.5 R3 `partial` **oder** nach R5 kein Envelope-Wert bei Emissions-Disposition `record-discarded` | `not run` |
| `OBS-LLH-TC-09` | Normalisierung fehlgeschlagen | `TL-2` | `normalization-failed` | `not run` |
| `OBS-LLH-TC-10` | No-Mutation-Sentinel | `TL-7` | ergebnisunabhängig — siehe §8 | `not run` |
| `OBS-LLH-TC-11` | Heterogene All-Failure-Zusammensetzung | `TL-2` + `TL-3`, *level assignment pending `P-3`* | **kein** Envelope-Wert (§10.5 **R6**); Emissions-Disposition `record-discarded` | `not run` |

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
- **Setup concept:** Quelle existiert, ist aber im Erhebungsversuch nicht lesbar (transiente Nichtverfügbarkeit, Zeitüberschreitung, Ressourcenkonflikt). Fixture-Klasse `unavailable`. *Level assignment pending `P-3`*, weil die konkrete Unverfügbarkeitsform vom Erhebungsmechanismus abhängt.
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
- **Setup concept:** Ein Beobachtungssubjekt, bei dem **kein** Feld `success` erreicht, **mindestens zwei** Felder **nicht** `provenance-invalid` sind und diese Felder **verschiedene** Erhebungsausgänge tragen — etwa ein Feld `permission-denied` und ein Feld `source-absent`. Kombination der bestehenden Fixture-Klassen (`permission-denied` + `missing`); die konkrete Kombinierbarkeit ist erst nach `P-3` bestimmbar. *Level assignment pending `P-3`*.
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
| **`P-2`-Evidenzmethode / -plan definiert** | Gate A — vor Produktivcode | **DEFINIERT** — die sieben kumulativen Anforderungen aus §8.3 |
| **`P-2` erfüllt** | Gate B — vor realer Beobachtung | **`NOT SATISFIED`** |

```text
P-2 evidence plan != P-2 satisfied
test design       != test execution
```

**Warum die Erfüllung erst nach der Implementierung möglich ist.** Die Anforderungen aus §8.3 setzen einen **ausführbaren** Beobachtungspfad, freigegebene Fixtures, eine deklarierte Umgebung sowie getrennte Ziel- und Ausführungsautorität voraus. Keine davon kann existieren, bevor überhaupt implementiert wurde. `P-2` als Vorbedingung für die Autorisierung produktiven Codes zu führen wäre daher **zirkulär**: die Evidenz verlangte genau das, was sie freigeben soll. Auf Gate-A-Ebene ist deshalb ausschließlich verlangt, dass die **Methode feststeht** — was mit §8.3 der Fall ist.

**`P-2` ist keine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.** Sie ist eine Voraussetzung dafür, ein reales Ziel zu beobachten.

### 8.6 Aktueller Stand

```text
P-2 evidence plan:    DEFINED  (§8.3, seven cumulative requirements)
P-2 satisfaction:     NOT SATISFIED
Tests implemented:    0
Tests executed:       0
No-mutation evidence: NONE
```

Dieses Dokument **beschafft** diese Evidenz nicht und darf nicht so gelesen werden, als hätte es sie beschafft. Ein definierter Evidenzweg ist keine Evidenz.

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

> **Gate-Zuordnung.** Die folgende Liste ist eine **Ausführungs**-Voraussetzungsliste. Sie beschreibt, was vor **Testausführung und realer Beobachtung** vorliegen muss — Gate B im Sinne von [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §16.2. **Keiner** ihrer Punkte ist eine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren. Auf Gate-A-Ebene ist aus diesem Dokument ausschließlich verlangt, dass Testdesign und `P-2`-Evidenzmethode **definiert** sind — beides ist mit §7 und §8.3 der Fall.

Eine spätere Ausführung erfordert **kumulativ** (konsistent mit Teststrategie §20, ergänzt um die Slice-Preconditions):

1. `P-3` — Entscheidung über Erhebungsmechanismus und Transport (**`NOT SELECTED`**).
2. Sprach-/Runtime-Entscheidung (**`NOT SELECTED`**).
3. Implementierung des Beobachtungspfads (**nicht autorisiert**).
4. Freigegebene Fixtures mit Identity, Revision und Provenance (**existieren nicht**).
5. Bereitgestelltes Lab-Environment mit deklarierter Environment Identity (**nicht bereitgestellt**).
6. `P-1` — separate Human-Maintainer-Zielautorisierung (**`NOT AUTHORIZED`**).
7. Explizite Human-Maintainer-Autorisierung der **Testausführung** (**nicht erteilt**).

Solange auch nur eine dieser Voraussetzungen offen ist, bleibt jeder Fall `not run` — und `not run != passed`.

## 11. Compatibility

Additiv. **Keine** Änderung an Teststrategie, Fixture-Governance, Lab-Modell, Ergebnisvokabular, Claim Boundary Set, Decision Index, Risk Register oder Capability Matrix. **Kein** neues Ergebnis-, Evidenz- oder Coverage-Modell. **Kein** ADR. Breaking-Change-Potenzial: gering.

## 12. Open Questions

- Konkrete Fixture-Klassen und -Revisionen je Fall (erst nach `P-3`).
- Environment-Profil für einen lokalen Linux-Host im Lab-Modell (erst nach `P-3`).
- Ob der Sentinel zusätzlich eine Langzeit-/Wiederholungsdimension benötigt (Regressionsvertrauen).
- Welche Nebeneffekte am Ziel mit der später gewählten Methode überhaupt beobachtbar sind.

## 13. Next Decision

**Aktueller Stand.** `CO-WP-032` ist **abgeschlossen und remote integriert**: Nova Final Review `GO`; Human-Maintainer-Integrationscommit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht; die anschließende Post-Integrations-Reconciliation steht auf `390da5cc8629dfa9cbea990c0a3c4ba4cb156e9b`. Die Repository-, Staging-, Commit- und Push-Gates für `CO-WP-032` sind damit **erledigt** und **nicht** mehr der nächste Schritt.

Der vorliegende docs-only Nachtrag (`OBS-LLH-TC-11` und diese Stands-Korrektur) liegt im Arbeitsverzeichnis. Staging, Commit, Push, Tag und jede weitere Repository-Aktion liegen unverändert **ausschließlich** beim Human Maintainer; ein künftiger Integrationsstand wird hier **nicht** vorweggenommen und **kein** Commit-Identifier dafür vorhergesagt.

**Unverändert und ausdrücklich nicht erteilt:**

```text
Testimplementierung:   NOT AUTHORIZED
Testausführung:        NOT AUTHORIZED
Fixtures:              nicht bereitgestellt
Lab-Environment:       nicht bereitgestellt
P-1 / Zielzugriff:     NOT AUTHORIZED
P-2 Evidenzplan:       DEFINED
P-2 Erfüllung:         NOT SATISFIED
P-3:                   NOT SELECTED
Sprache / Runtime:     NOT SELECTED
Tests implementiert:   0
Tests ausgeführt:      0
```

Testimplementierung, Testausführung, Fixture- und Lab-Bereitstellung bleiben getrennte, ausdrücklich zu autorisierende spätere Arbeit. **Dieser Envelope autorisiert keine davon.** Die Aufnahme von `OBS-LLH-TC-11` ist eine **Testfalldefinition** und **keine** Implementierung, **keine** Ausführung und **keine** Evidenz: `case enumerated != case implemented`.
