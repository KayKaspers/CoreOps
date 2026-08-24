# CoreOps – Synthetic Fixture and Test Data Governance

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation synthetic fixture and test data governance
> Implementation Status: Not implemented
> Fixture Implementation Status: Not implemented
> Test Execution Status: Not performed
> Validation Status: Not performed
> Integration Lab Status: Not provisioned
> Technology Selection: None
> Fixture-Format/Generator/Data-Store/Mock/Anonymization Technology: Not selected
> Production Validation: Not performed
> Accessibility Validation: Not performed
> Production Data Use: None
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-028 (docs-only / test strategy, fixtures and integration lab)

## 1. Status

Technologieunabhängige Governance für **synthetische Fixtures** und **Testdaten** in CoreOps: Fixture-Autoritätsgrenze, Identity/Revision/Provenance, Fixture-Prinzipien, Fixture-Klassen und deren Domänen-Anwendbarkeit, Expected-Outcome-Deklaration, Repräsentativitätsgrenze, Bindung an bestehende Klassifikations-, Minimierungs-, Redaction-, Retention-, Secret- und Isolationsgovernance, Disclosure-Grenzen von Testartefakten sowie Fixture-Lebenszyklus. Companion zu [FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) und [INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md).

**Es existiert keine Fixture.** Dieses Dokument definiert, wie künftige Fixtures beschaffen sein müssen; es erzeugt keine, verwendet keine und validiert keine.

## 2. Purpose

Testdaten wirken harmlos, bis sie es nicht sind. Die Governance legt fest, warum `synthetic ≠ safe by definition`, `masked secret ≠ non-secret`, `test data ≠ unclassified data`, `fixture correctness ≠ system correctness` und `fixture representativeness ≠ production equivalence` — und wie Fixtures aussagekräftig sein können, **ohne** reale Geheimnisse, reale Ziele oder personenbezogene Daten in Testartefakte zu tragen.

## 3. Scope

Fixture-Autoritätsgrenze · Fixture Identity, Revision und Provenance · Fixture-Prinzipien · Fixture-Klassen · Domänen-Anwendbarkeit · Expected-Outcome-Deklaration · Repräsentativität und Fidelity · Klassifikationsbindung von Testdaten · Minimierung und Redaction · Secret- und Key-Grenze · Workspace-, Environment- und Target-Isolation · Retention und Löschung von Testdaten und Test-Evidenz · Disclosure-Flächen von Testartefakten · Grenze für produktionsabgeleitete Daten · Fixture-Lebenszyklus · Review und Änderungskontrolle · Failure und Unknown.

## 4. Non-Goals

- **Keine** Auswahl von Fixture-Format, Generator, Daten-Store, Mock-Framework, Anonymisierungs- oder Maskierungstechnologie.
- **Keine** erzeugten, generierten oder eingecheckten Fixture-Dateien.
- **Keine** Verwendung, Kopie oder Ableitung realer Produktionsdaten.
- **Kein** neues Klassifikations-, Retention-, Redaction- oder Secret-Modell — ausschließlich Bindung an die bestehenden.
- **Keine** Aussage über Datenschutz-Konformität, Anonymitätsgrad oder Re-Identifikationssicherheit.

## 5. Concepts

Mindestbegriffe: `fixture` · `fixture identity` · `fixture revision` · `fixture provenance` · `fixture class` · `fixture set` · `expected outcome declaration` · `fidelity declaration` · `test data object` · `test data class` · `derived test data` · `test artefact` · `test evidence artefact` · `debug output` · `fixture lifecycle state` · `fixture review`.

```text
fixture                ≠ test case
fixture set            ≠ scenario
test data object       ≠ production record
test artefact          ≠ evidence artefact
debug output           ≠ internal-only by nature
```

## 6. Fixture Authority Boundary

Eine Fixture ist ein **kontrollierter Testeingang**, kein Zustand, keine Quelle und keine Autorität.

```text
fixture                     ≠ authoritative state
fixture                     ≠ source of truth
fixture accepted by a test  ≠ data valid in CoreOps
fixture correctness         ≠ system correctness
fixture representativeness  ≠ production equivalence
fixture passed              ≠ production validated
```

Eine Fixture erteilt keine Berechtigung, erzeugt keinen Trust und setzt keinen Compatibility-, Approval- oder Validation-Zustand. Registriert als DEC-S-371.

## 7. Fixture Identity, Revision and Provenance

Jede künftige Fixture führt konzeptionell mindestens:

| Feld | Bedeutung |
| ---- | --------- |
| Fixture Identity | stabile, nicht wiederverwendbare Kennung |
| Fixture Revision | getrennt von der Identity; jede Inhaltsänderung erzeugt eine neue Revision |
| Fixture Class | genau eine primäre Klasse aus §9 |
| Domain / Subject Binding | auf welches CoreOps-Modell oder Modul sie sich bezieht |
| Provenance | wie sie entstanden ist (synthetisch konstruiert, abgeleitet, kuratiert) |
| Origin Declaration | ausdrücklich: `synthetic` oder — nur im Ausnahmefall §19 — `production-derived` |
| Classification | Datenklasse nach [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) §6 |
| Expected Outcome | erwartetes Verhalten des Subjekts (§11) |
| Known Limitations | wofür sie ausdrücklich nicht repräsentativ ist |
| Lifecycle State | Zustand nach §20 |
| Review Reference | Freigabebezug, sofern erforderlich |

```text
same fixture identity ≠ same fixture revision
fixture revision unchanged ≠ meaning unchanged when its source contract changed
provenance stated ≠ provenance verified
```

Retired Fixture-Identities werden **nicht** wiederverwendet; historische Test-Evidenz bleibt referenzierbar.

## 8. Fixture Principles

Verbindliche Grundregeln für jede künftige Fixture:

- **synthetisch als Default** — konstruiert, nicht kopiert.
- **keine Produktionskopie als Default** — Ausnahmen nur unter §19.
- **kein wiederverwendbares Credential** — keine Passwörter, Tokens, Schlüssel, Zertifikate mit realer Gültigkeit.
- **kein reales Secret** — auch nicht abgelaufen, auch nicht „nur zum Testen".
- **kein realer privater Hostname, Domainname, Netzbereich oder Adressbezug.**
- **keine unredigierten personenbezogenen Daten.**
- **keine implizite Autorität** — eine Fixture beweist nichts über Berechtigung.
- **Identity explizit.**
- **Revision explizit.**
- **Provenance explizit.**
- **Expected Outcome explizit.**

```text
synthetic         ≠ safe by definition
synthetic         ≠ unclassified
test-only label   ≠ handling exemption
expired credential ≠ non-credential
placeholder value ≠ absence of a secret reference
```

Registriert als DEC-S-371.

## 9. Fixture Classes

Klassen als Testeingangs-Charakter, nicht als Datentyp:

| Klasse | Kurzcharakter |
| ------ | ------------- |
| `nominal` | erwarteter, wohlgeformter Regelfall |
| `boundary` | Grenzwert an einer deklarierten Schwelle |
| `negative` | gezielt unzulässiger Eingang zur Beobachtung des Schutzverhaltens |
| `malformed` | strukturell defekt oder nicht parsbar |
| `missing` | erwarteter Bestandteil fehlt |
| `stale` | inhaltlich gültig, aber jenseits der Freshness-Grenze |
| `unknown` | Wert oder Zustand ist nicht bestimmt |
| `partial` | unvollständige, aber sichtbar unvollständige Menge |
| `conflicted` | zwei Quellen behaupten Unvereinbares |
| `revoked` | zuvor gültig, inzwischen widerrufen |
| `expired` | Gültigkeitsfenster überschritten |
| `unsupported` | außerhalb des deklarierten Support-Umfangs |
| `incompatible` | Version oder Contract passt nachweislich nicht |
| `wrong-target` | korrekt geformt, aber an das falsche Ziel gebunden |
| `wrong-environment` | korrekt geformt, aber an die falsche Umgebung gebunden |
| `permission-denied` | Aufruf ist zulässig geformt, aber nicht berechtigt |
| `capability-unavailable` | Fähigkeit ist nicht vorhanden (≠ nicht berechtigt) |
| `degraded` | Eingang unter eingeschränktem Betriebsmodus |
| `recovery` | Eingang im Kontext einer Wiederherstellung |

```text
capability-unavailable ≠ permission-denied
unknown                ≠ missing
stale                  ≠ invalid
partial                ≠ malformed
revoked                ≠ expired
unsupported            ≠ incompatible
```

**Nicht jede Klasse gilt für jede Domäne.** Eine Klasse, die für ein Subjekt keinen Sinn ergibt, wird als `not applicable` geführt — ausdrücklich **nicht** als `unknown` und **nicht** als Lücke.

## 10. Fixture Class Applicability by Domain

Orientierung, keine Pflichtmatrix. Die verbindliche Zuordnung entsteht mit dem jeweiligen Testfall (§10 der Teststrategie).

| Domäne | Typisch erforderliche Klassen | Typisch `not applicable` |
| ------ | ----------------------------- | ------------------------ |
| Identity und Autorisierung | `negative`, `expired`, `revoked`, `permission-denied`, `boundary` | `degraded`, sofern die Identitätsprüfung modusunabhängig ist |
| Policy und Approval | `nominal`, `negative`, `conflicted`, `unknown`, `expired` | `malformed` bei rein logischen Regeln |
| Integration und Adapter | `nominal`, `unsupported`, `incompatible`, `malformed`, `capability-unavailable` | `recovery`, sofern die Integration keine Recovery-Rolle hat |
| Migration | `nominal`, `partial`, `unknown`, `incompatible`, `recovery` | `permission-denied` auf Datenebene |
| Telemetry | `nominal`, `missing`, `malformed`, `conflicted`, `partial`, `unknown` | `revoked` |
| Deployment | `nominal`, `partial`, `unknown`, `wrong-target`, `wrong-environment` | `stale` auf Blueprint-Ebene, sofern versioniert |
| Artifact und SBOM | `nominal`, `partial`, `missing`, `revoked`, `incompatible` | `degraded` |
| Offline und CorePack | `nominal`, `revoked`, `stale`, `wrong-target`, `partial`, `expired` | `capability-unavailable` |
| Secrets | Platzhalter-`nominal`, `permission-denied`, `expired`, `revoked`, `unknown` | **jede Klasse, die einen realen Wert erfordern würde** |
| Klassifikation und Retention | `nominal`, `missing`, `stale`, `unknown`, `partial` | `incompatible` |
| Degraded und Recovery | `degraded`, `recovery`, `partial`, `unknown`, `expired` | `unsupported` |
| Experience und Darstellung | alle Zustandsklassen als Darstellungseingang | `malformed`, sofern die Ebene nur Ergebnisse darstellt |

## 11. Expected Outcome Declaration

Jede Fixture deklariert, **was mit ihr geschehen soll** — nicht nur, was sie enthält:

- erwartetes zulässiges Ergebnis;
- erwartetes **verbotenes** Ergebnis;
- ob sie ein Positiv- oder ein Negativszenario bedient;
- bei Negativszenarien: das **konkrete geschützte Verhalten**, das beobachtet werden muss.

```text
fixture accepted   ≠ expected outcome met
fixture rejected   ≠ protected behavior observed
rejection reason unspecified ≠ negative scenario passed
```

Eine Fixture ohne deklariertes Expected Outcome ist unvollständig und darf keinen Testfall stützen.

## 12. Representativeness and Fidelity

Jede Fixture deklariert ihre **Fidelity** gegenüber dem, was sie nachbildet, und ihre bekannten Grenzen.

```text
fixture representativeness ≠ production equivalence
synthetic fixture          ≠ real environment
plausible shape            ≠ observed real behavior
one vendor shape           ≠ vendor family behavior
fixture volume             ≠ production volume
```

Eine hohe erklärte Fidelity ist eine **Behauptung**, kein Nachweis. Fidelity wird nie aus einem bestandenen Test abgeleitet.

## 13. Test Data Classification Binding

Testdaten sind **klassifizierte Daten**. Es gilt unverändert [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md):

- Jedes Testdatenobjekt trägt eine Datenklasse; fehlt sie, gilt `unknown-classification` mit **fail-closed** zur strengeren anwendbaren Grenze.
- Eine aus dem Container geerbte Klassifikation (Fixture-Set, Lab-Umgebung, Evidenzpaket) ist **nicht automatisch gültig**.
- Materielle Änderung an Inhalt, Herkunft, Zweck, Ziel, Workspace oder Disclosure Scope löst **Reassessment** aus.

```text
test data        ≠ unclassified data
synthetic        ≠ public
internal test data ≠ safe for external disclosure
container-inherited classification ≠ automatically valid
```

Registriert als DEC-S-371.

## 14. Minimization and Redaction

Es gilt unverändert [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md):

- Eine Fixture enthält nur die Felder, die das Szenario tatsächlich benötigt.
- Redaction in Testartefakten ist eine **Derived View**, keine Löschung an der Quelle.
- Maskierung und Pseudonymisierung erzeugen **keine** Anonymität.

```text
redacted           ≠ disclosure-safe
masked/pseudonym   ≠ anonymous
minimized          ≠ non-sensitive
redacted view      ≠ source deleted
```

Es wird **keine** Aussage über Re-Identifikationssicherheit getroffen und **kein** Anonymisierungsverfahren ausgewählt.

## 15. Secret and Key Boundary in Fixtures

Absolut und ausnahmslos:

- **Kein realer Secret Value** in einer Fixture, einem Testfall, einer Lab-Konfiguration, einem Testprotokoll, einer Fehlerausgabe oder einem Evidenzartefakt.
- Zulässig sind ausschließlich **Platzhalter-Referenzen** ohne Wert, entsprechend dem Configuration-Reference-Modell.
- Ein maskiertes Secret bleibt `secret-bearing` und bleibt an [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) gebunden.

```text
masked secret       ≠ non-secret automatically
secret reference    ≠ secret value
test credential     ≠ harmless credential
expired key         ≠ safe to publish
no raw credentials in test artefacts
```

Key-Material-Handhabung, Rotation und Custody werden durch Fixtures **nicht** geübt und **nicht** belegt.

## 16. Workspace, Environment and Target Isolation

Testdaten sind an Workspace, Environment und Ziel gebunden — genau wie Betriebsdaten:

- Keine Wiederverwendung von Testdaten über Workspace-Grenzen hinweg ohne explizite, autorisierte Entscheidung.
- Eine Fixture, die an ein Ziel oder eine Umgebung gebunden ist, ist in einer anderen `wrong-target` bzw. `wrong-environment` — und damit ein **Negativfall**, kein Bequemlichkeitspfad.
- Lab-Testdaten dürfen kein Produktionsziel adressieren ([Lab Companion](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md) §13).

```text
same shape        ≠ same scope
reusable fixture  ≠ cross-workspace authorization
lab test data     ≠ production data
```

## 17. Retention and Deletion of Test Data and Test Evidence

Es gilt unverändert [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](../governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md):

- Fixtures, Lab-Daten und Test-Evidenz haben **getrennte** Retention-Erwartungen.
- Die Aufbewahrung eines Testergebnisses rechtfertigt **nicht** die unbegrenzte Aufbewahrung der zugrunde liegenden Daten.
- Abgelaufene Retention ist weder eine Löschautorisierung noch ein Löschabschluss.
- Kopien in Backups, Caches, Derived Data und Evidenzpaketen sind Teil des Kopien-Inventars.

```text
test evidence retention ≠ permission to retain source data indefinitely
retention expired       ≠ deletion complete
primary copy deleted    ≠ all copies removed
deletion outcome unknown ≠ deleted
```

## 18. Disclosure Surfaces of Test Artefacts

Testartefakte sind **Disclosure-Flächen**. Mindestens zu governen: Testprotokolle · Fehlerausgaben und Stack-Traces · Diff-/Vergleichsausgaben · Screenshots und Renderings · Evidenzpakete · Export- und Reportausgaben · Debug- und Trace-Ausgaben.

```text
debug output   ≠ free disclosure surface
failure dump   ≠ exempt from classification
screenshot     ≠ redacted by default
test log       ≠ internal-only by nature
read           ≠ export
```

Ein Testartefakt, das eine Datenklasse enthält, unterliegt deren Export- und Logging-Grenze — unabhängig davon, dass es „nur ein Testlauf" war.

## 19. Production-Derived Data Boundary

Produktionsdaten sind als Fixture-Quelle **standardmäßig ausgeschlossen**. Eine spätere Ausnahme wäre nur zulässig mit kumulativ:

1. expliziter, zurechenbarer Human-Maintainer-Autorisierung je Vorgang;
2. dokumentierter Notwendigkeit, die synthetisch nachweislich nicht erreichbar ist;
3. Klassifikation und Reassessment der abgeleiteten Daten;
4. Minimierung und Redaction vor jeder Verwendung;
5. `production-derived` als Origin Declaration in der Fixture-Identität;
6. eigener, begrenzter Retention- und Löschgrenze;
7. Ausschluss aus jeder öffentlichen oder repositorygebundenen Ablage.

```text
production copy           ≠ default fixture source
anonymization claimed     ≠ anonymity achieved
authorized once           ≠ standing authorization
production-derived fixture ≠ synthetic fixture
```

**Dieses Dokument autorisiert keinen solchen Vorgang.** Es beschreibt ausschließlich die Bedingungen, unter denen eine spätere Autorisierung überhaupt geprüft werden dürfte.

## 20. Fixture Lifecycle

Zustände: `draft` · `review-pending` · `active` · `deprecated` · `retired` · `quarantined`.

```text
draft        ≠ usable in an evidence-bearing run
active       ≠ correct
deprecated   ≠ removed
retired identity ≠ reusable identity
quarantined  ≠ deleted
```

Eine Fixture wird `quarantined`, sobald der Verdacht besteht, dass sie reale Secrets, reale Ziele, personenbezogene Daten oder eine falsche Klassifikation trägt. Quarantäne ist **fail-closed**: bestehende Ergebnisse, die auf ihr beruhen, werden als `inconclusive` behandelt, bis die Prüfung abgeschlossen ist.

## 21. Fixture Review and Change Control

- Jede `active`-Fixture hat einen zurechenbaren Owner.
- Eine Inhaltsänderung erzeugt eine neue Revision; darauf beruhende Ergebnisse bleiben **historische Evidenz für die vorherige Fixture-Revision** und werden dadurch **nicht** automatisch `stale`. Ihre Anwendbarkeit auf die neue Revision ist jedoch nicht etabliert und erfordert Reassessment bzw. Revalidierung (Teststrategie §18).
- Änderungen an Klassifikation, Origin Declaration oder Expected Outcome sind reviewpflichtig.
- Ein bestandener Testlauf ist **kein** Freigabegrund für eine Fixture.

```text
fixture reviewed ≠ fixture validated
fixture in use   ≠ fixture approved
passing run      ≠ fixture quality evidence
```

## 22. Failure and Unknown State

- Unbekannte Klassifikation, Herkunft oder Zielbindung → **fail-closed**, Fixture nicht verwendbar.
- Unbekannter Löschausgang bei Testdaten → als `unknown` geführt, nie als gelöscht.
- Unbekannte Fidelity → Ergebnisse tragen eine ausdrückliche Limitation.

```text
unknown ≠ safe · unknown ≠ absent · unknown ≠ acceptable
```

## 23. Security Invariants

1. Kein reales Secret, kein wiederverwendbares Credential und kein reales Schlüsselmaterial in einer Fixture oder einem Testartefakt.
2. Kein realer privater Hostname, Domainname oder Netzbereich in einer Fixture oder einem Testartefakt.
3. Keine unredigierten personenbezogenen Daten in einer Fixture oder einem Testartefakt.
4. Produktionsdaten sind als Fixture-Quelle ausgeschlossen, solange keine explizite Einzelautorisierung nach §19 vorliegt.
5. Testdaten sind klassifizierte Daten; unbekannte Klassifikation failt closed zur strengeren Grenze.
6. Ein maskiertes Secret bleibt `secret-bearing`.
7. Eine Fixture erteilt keine Berechtigung, keinen Trust und keinen Compatibility-Zustand.
8. Fixture-Korrektheit belegt keine Systemkorrektheit; Repräsentativität belegt keine Produktionsäquivalenz.
9. Testartefakte unterliegen den Export-, Logging- und Disclosure-Grenzen ihrer Datenklasse.
10. Retention von Test-Evidenz verlängert die Retention der Quelldaten nicht.
11. Eine Fixture ohne Identity, Revision, Provenance und Expected Outcome stützt keinen Testfall.
12. Quarantäne einer Fixture wirkt fail-closed auf alle darauf beruhenden Ergebnisse.

## 24. Threat References

Referenziert, nicht dupliziert: THR-002, THR-013, THR-016, THR-018, THR-019, THR-020, THR-024, THR-025, THR-033 und THR-035 aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md). Diese Governance benennt Datenschutz- und Geheimnisgrenzen für Testdaten und erhebt **keinen** Verifikations- oder Konformitätsanspruch.

## 25. Technology Boundary

Fixture-Format, Fixture-Generator, Testdaten-Store, Mock-Framework, Anonymisierungs- und Maskierungstechnologie sowie Datensynthese-Werkzeuge: **nicht ausgewählt**. Die Auswahl gehört in spätere Implementierungs- und ADR-Arbeit.

## 26. Compatibility

Bindet an, ohne zu ändern: [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md), [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](../governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md), [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md), [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md), [KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md) und [PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md).

## 27. Open Questions

- Ob Fixtures im Repository oder ausschließlich in der Laborumgebung gehalten werden — abhängig von Klassifikation und Public-Neutrality-Grenze; Entscheidung gehört zu späterer Tooling- und ADR-Arbeit.
- Wie Fixture-Revisionen an Contract-Revisionen gebunden werden, ohne ein zweites Versionierungsmodell zu erzeugen.
- Wie weit Fidelity-Deklarationen formalisiert werden, ohne eine implizite Kompatibilitätsaussage zu erzeugen.

## 28. Next Decision

Nova Review dieses Dokuments gemeinsam mit der [Teststrategie](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) und dem [Lab-Modell](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md). **Keine** Fixture wird durch dieses Dokument erzeugt, verwendet, freigegeben oder validiert.
