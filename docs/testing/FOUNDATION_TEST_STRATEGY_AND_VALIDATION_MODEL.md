# CoreOps – Foundation Test Strategy and Validation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation test strategy and validation model
> Implementation Status: Not implemented
> Test Implementation Status: Not implemented
> Test Execution Status: Not performed
> Validation Status: Not performed
> Integration Lab Status: Not provisioned
> Technology Selection: None
> Test-Framework/Language/CI/Container/Virtualization/Mock/Scanner/Coverage/Reporting Technology: Not selected
> Production Validation: Not performed
> Accessibility Validation: Not performed
> Security Verification: Not performed
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-028 (docs-only / test strategy, fixtures and integration lab)

## 1. Status

Technologieunabhängige **Foundation Test Strategy** für CoreOps: Testebenen, Subject-under-Test-Klassen, Traceability zu bestehenden Contracts/Invarianten, Test-Ergebnissemantik, Negativ-/Fail-Closed-Testing, Foundation-Szenarienfamilien, Contract-Testing über die sechs Capability-Dimensionen, Coverage-Modell und Risk-to-Test-Mapping. Companions: [SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) und [INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md).

Dieses Dokument **plant** Tests. Es implementiert keine, führt keine aus und erzeugt keine Validierungsevidenz. Es erzeugt **keine** zweite Evidence-Autorität und **kein** zweites Statusmodell.

## 2. Purpose

Eine grüne Testsuite ist kein Betriebsnachweis. Das Modell legt fest, warum `test planned ≠ test implemented ≠ test executed ≠ test passed ≠ requirement universally satisfied`, warum `no failing test ≠ absence of defect` und `coverage reported ≠ coverage complete` — und wie eine spätere Implementierung ausführbare Tests ergänzen kann, **ohne** die Foundation-Testsemantik zu verändern.

## 3. Scope

Validierungs-Autoritätsgrenze · Claim Boundary Set · Testebenen-Taxonomie mit Ebenenkontrakt · Subject-under-Test-Klassen · Test-Case-Traceability-Contract · Test-Ergebnissemantik · Negativ-/Fail-Closed-Testing · Foundation-Szenarienfamilien · Integration-Contract-Testing · CO-WP-027-Designannahmen als Testsubjekte · Accessibility-/Lokalisierungs-Evidenzanforderungen · Coverage-Modell · Test-Evidenz und Reproduzierbarkeit · Risk-to-Test-Matrix · Execution Gates · Exit-/Acceptance-Semantik. Fixtures und Testdaten im [Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md); Laborumgebung, Test Doubles und Lab-Evidenz im [Lab Companion](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md).

## 4. Non-Goals

- **Keine** Auswahl von Test-Framework, Programmiersprache, CI-Provider, Container-/Virtualisierungstechnologie, Mock-Framework, Browser-Automation, Load-Test-, Accessibility-Scanner-, Security-Scanner-, Datenbank-, Message-Broker-, Lab-Orchestrator-, Netzwerkemulator-, Test-Management-, Coverage- oder Reporting-Technologie.
- **Keine** ausführbaren Tests, Testskripte, CI-Jobs, Container, VMs, Scans oder Tool-Läufe.
- **Kein** neues Autoritäts-, Evidence-, Approval- oder Statusmodell für CoreOps.
- **Keine** Aussage über Sicherheit, Recovery-Fähigkeit, Offline-Betrieb, Usability, Accessibility-Konformität, Herstellersupport, Produktions- oder Releasereife.
- **Keine** Retrofit-Änderung an bestehenden Foundation-Dokumenten.

## 5. Concepts

Mindestbegriffe: `test subject` · `test level` · `test case` · `test case revision` · `test scenario` · `scenario family` · `stimulus` · `observation point` · `expected outcome` · `expected prohibited outcome` · `protected behavior` · `test result` · `test run` · `test evidence` · `coverage dimension` · `execution gate` · `result limitation` · `applicability`.

```text
test subject        ≠ test case
test case           ≠ test run
test run            ≠ test evidence
test evidence       ≠ validation evidence for a requirement
protected behavior observed ≠ system universally secure
```

## 6. Validation Authority Boundary

Die Teststrategie ist **abgeleitet**, nicht autoritativ. Sie referenziert bestehende CoreOps-Modelle und ersetzt keines.

- Autoritative Evidenz-, Audit- und Provenance-Semantik bleibt bei [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [EVENT_AND_AUDIT_CORRELATION_MODEL.md](../architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md) und MOD-EVD-001. Test-Evidenz ist eine **Evidence-Klasse innerhalb** dieses Modells, keine parallele Autorität.
- Autoritative Zustands-, Policy-, Approval- und Execution-Semantik bleibt bei [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](../architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [POLICY_DECISION_AND_EVALUATION_MODEL.md](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](../security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) und [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md).
- Ein Testergebnis erteilt keine Berechtigung, ändert keinen autoritativen Zustand und ersetzt keine Freigabe.

```text
test evidence           ≠ operational authority
test evidence           ≠ approval
validation evidence     ≠ approval
validation              ≠ support
support                 ≠ implementation
test result             ≠ authoritative state
test strategy           ≠ second evidence authority
```

Registriert als DEC-S-366.

## 7. Claim Boundary Set

Die folgenden Aussagen bilden ein **Claim Boundary Set** — eine Menge voneinander **unabhängiger Nicht-Implikationen**, die für jede Aussage über CoreOps-Tests in Dokumentation, Berichten, Release-Texten und Reviews gilt. Das Set ist ausdrücklich **keine** universelle Reifeleiter: es gibt keine Pflichtfolge, die jede Testaussage durchlaufen müsste, und keine Dimension wird durch das Erreichen einer anderen „erklommen".

```text
claim boundary set ≠ universal maturity ladder
```

### 7.1 Zulässige lokale Progression (Ausführungsfolge)

Nur diese Folge ist eine echte Abfolge, weil jede Stufe die vorherige tatsächlich voraussetzt:

```text
test planned → test implemented → test executed → test result
```

mit den zugehörigen Grenzen:

```text
test planned     ≠ test implemented
test implemented ≠ test executed
test executed    ≠ test passed
```

### 7.2 Unabhängige Claim-Grenzen (keine Sprossen)

Die folgenden Aussagen stehen **nebeneinander**. Keine ist eine höhere Stufe der Ausführungsfolge aus §7.1, keine wird durch fortgesetzte Testausführung erreicht, und keine ist als Pflichtziel eines Testfalls zu lesen:

```text
test passed             ≠ requirement universally satisfied
fixture passed          ≠ production validated
synthetic fixture       ≠ real environment
simulator success       ≠ provider compatibility
integration-lab success ≠ production readiness
one successful run      ≠ regression confidence
no failing test         ≠ absence of defect
coverage reported       ≠ coverage complete
```

### 7.3 Ergebnis-Nicht-Implikationen

Ebenfalls unabhängig, auf der Ergebnisebene (§11):

```text
not tested     ≠ passed
blocked        ≠ failed
inconclusive   ≠ passed
not applicable ≠ unknown
```

### 7.4 Ausdrücklich keine Sprossen

`support` · `production readiness` · `security validation` · `accessibility validation` · `fixture representativeness` · `operational validation` sind **eigenständige Dimensionen mit eigener Autorität und eigener Evidenzanforderung** — **nicht** die oberen Enden einer Testprogression.

```text
eigenständige Dimension ≠ Endstufe einer Testprogression
kein Test „steigt auf" zu support, production readiness oder validation
Testfall ohne Bezug zu diesen Dimensionen ≠ unvollständiger Testfall
```

Eine Aussage, die eine dieser Grenzen **überschreitet** — also die eine Seite mit der anderen gleichsetzt —, ist **Claim Inflation** und in CoreOps unzulässig. Registriert als DEC-S-366.

## 8. Test Level Taxonomy

Die Ebenen sind aus CoreOps-Gegenständen abgeleitet, **nicht** aus einer konventionellen Testpyramide. Es gibt keine implizite Mengenverteilung zwischen den Ebenen und keine Ebene ist „wichtiger" als eine andere. Eine Ebene ist ein **Betrachtungsgegenstand**, keine Autoritätsstufe.

```text
test level   ≠ authority level
higher level ≠ stronger evidence
lower level  ≠ weaker requirement
```

| Ebene | Purpose | Subject under Test | Authoritative Input | Fixture / Environment Type |
| ----- | ------- | ------------------ | ------------------- | -------------------------- |
| `TL-1` Document and Contract Consistency | Widerspruchsfreiheit der Foundation-Aussagen, Invarianten, IDs und Register | Foundation-Dokumente, Decision Index, Risk Register, Capability Matrix | die Dokumente selbst | keine Umgebung; Repository-Inhalt |
| `TL-2` Schema and Model Conformance | Konformität deklarierter Modelle/Schemata zu ihrem Contract | Blueprint-, Telemetry-, Evidence-, CorePack-, Klassifikations- und API-Schemata | [SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md](../architecture/SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [API_GOVERNANCE_AND_OPERATION_MODEL.md](../architecture/API_GOVERNANCE_AND_OPERATION_MODEL.md) | synthetische Instanzdaten, keine Laufzeit |
| `TL-3` Module Contract Behaviour | Verhalten eines einzelnen Moduls an seiner deklarierten Grenze | ein `MOD-*` aus dem [Module Catalog](../architecture/COREOPS_MODULE_CATALOG.md) | [MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md](../architecture/MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md) | isolierte Modulumgebung, Test Doubles für Nachbarn |
| `TL-4` Authority and Policy Decision | Policy-Auswertung, Approval und Execution Authorization als getrennte Schritte | MOD-POL-001, MOD-IAM-001, MOD-EXE-001 | [POLICY_DECISION_AND_EVALUATION_MODEL.md](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](../security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) | synthetische Principals, Scopes, Authorizations |
| `TL-5` Adapter and Provider Integration | Adapter/Agent gegen den Integration Contract | MOD-ADP-001, MOD-AGT-001, Domain Packs | [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md) | Test Double **oder** kontrolliertes Lab-Zielsystem |
| `TL-6` Control-Plane Scenario | mehrstufige Betriebsszenarien über mehrere Module | Deployment-, Migration-, Recovery-, Reconciliation-Abläufe | CO-WP-016/021/026-Dokumente | Integration Lab, mehrere Zielrollen |
| `TL-7` Security Negative and Fail-Closed | Beobachtung des geschützten Verhaltens bei fehlender/ungültiger Autorisierung | Autorisierungs-, Trust-, Revocation- und Isolationsgrenzen | [COREOPS_FOUNDATION_THREAT_MODEL.md](../security/COREOPS_FOUNDATION_THREAT_MODEL.md), [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) | Negativ-/Malformed-/Revoked-/Expired-Fixtures |
| `TL-8` Operational Mode and Connectivity | Verhalten je Connectivity-Klasse und Operational Mode | Mode-Übergänge, Offline-Import/Activation, Degraded/Recovery | [RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md), [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) | Lab-Environment-Profile (simuliert) |
| `TL-9` Data Governance and Disclosure | Klassifikation, Redaction, Retention, Secret- und Export-Grenzen | Testdatenobjekte, Derived Views, Evidenz- und Exportartefakte, Secret-Referenzen | [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md), [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) | synthetische klassifizierte Testdaten, Platzhalter-Secret-Referenzen |
| `TL-10` Experience and Accessibility | Darstellung, Aktionssicherheit, Zustandsunterscheidbarkeit, DE/EN | Experience-Ebene (MOD-EXP-001), Darstellungsflächen und Zustandsanzeigen | die drei CO-WP-027-Dokumente | manuelle/assistierte Prüfung an einer späteren Oberfläche |

| Ebene | Expected Evidence | Failure Meaning | Cannot Prove | Automation Suitability | Offline Suitability |
| ----- | ----------------- | --------------- | ------------ | ---------------------- | ------------------- |
| `TL-1` | Konsistenzbefund mit Dokument-/Abschnittsreferenz | dokumentarischer Widerspruch, **kein** Systemdefekt | dass die beschriebene Semantik implementierbar oder korrekt ist | hoch | vollständig offline |
| `TL-2` | Konformitäts-/Ablehnungsbefund je Instanz mit Feldreferenz | Modell oder Instanz verletzt den deklarierten Contract | dass reale Produzenten dieselben Instanzen erzeugen | hoch | vollständig offline |
| `TL-3` | beobachtetes Modulverhalten je Grenzfall | Modul verletzt seine deklarierte Grenze | Verhalten im Zusammenspiel mit realen Nachbarn | hoch | vollständig offline |
| `TL-4` | Entscheidungs-, Approval- und Authorization-Records getrennt | Autorisierungskette kollabiert oder erlaubt zu viel | universelle Sicherheit der Autorisierungslogik | hoch | vollständig offline |
| `TL-5` | Request-/Acceptance-/Result-Records plus Fidelity-Deklaration | Adapter verletzt Contract oder überschreitet Scope | Kompatibilität mit einem realen Herstellerprodukt | mittel | teilweise (nur mit Test Double) |
| `TL-6` | Szenario-Verlauf mit Per-Target-Ergebnissen und Verifikations-Records | Ablauf verletzt Reihenfolge, Autorität oder Sichtbarkeit | Produktionsreife, Skalierung, Topologiegröße | mittel | teilweise (Lab-abhängig) |
| `TL-7` | Record des **beobachteten geschützten Verhaltens**, nicht nur einer Fehlermeldung | Schutzverhalten wurde **nicht** beobachtet | Abwesenheit weiterer Angriffspfade | mittel | vollständig offline |
| `TL-8` | Mode-/Profil-Records mit Freshness-, Revocation- und Clock-Angaben | Modus- oder Konnektivitätsverhalten verletzt Autoritätsgrenzen | reale Resilienz unter echten Ausfällen | mittel | vollständig offline (simuliert) |
| `TL-9` | Klassifikations-, Redaction-, Retention- und Denial-Records | Datenhandhabung verletzt eine Governance-Grenze | Re-Identifikationssicherheit; reale Kopienvollständigkeit | mittel | vollständig offline |
| `TL-10` | strukturierte Zustands-, Textäquivalent- und Tastaturpfad-Inventare | Designanforderung ist nicht erfüllt | Usability, Accessibility-Konformität, Screenreader-Eignung | gering (überwiegend manuell) | offline möglich |

Registriert als DEC-S-367.

## 9. Subject-under-Test Classes

Jeder Testfall benennt genau **eine** primäre Subject-Klasse: `document/contract` · `schema/model instance` · `module` · `authority chain` · `adapter/integration` · `operation/scenario` · `operational mode` · `data object/class` · `evidence artefact` · `presentation surface`.

```text
subject class      ≠ module ownership
testing a subject  ≠ owning its authoritative state
```

## 10. Test Case Traceability Contract

Ein Testfall ist an bestehende Foundation-Semantik gebunden. Konzeptionelle Mindestfelder:

| Feld | Bedeutung |
| ---- | --------- |
| Test Case Identity | stabile, nicht wiederverwendbare Kennung |
| Test Case Revision | getrennt von der Identity; Inhaltsänderung erzeugt eine neue Revision |
| Source Contract / Requirement / Invariant | Dokument und Abschnitt der autoritativen Aussage |
| Decision Reference | `DEC-*`, sofern zutreffend |
| Risk Reference | `RISK-*`, sofern zutreffend |
| Threat Reference | `THR-*`, sofern zutreffend |
| Subject under Test | Subject-Klasse und konkretes Subjekt |
| Declared Scope | was der Fall abdeckt — und ausdrücklich nicht |
| Preconditions | erforderlicher Ausgangszustand |
| Fixture References | Fixture-Identity **und** -Revision |
| Environment / Profile | Lab-Environment-Identity und -Profil |
| Action / Stimulus | ausgelöste Handlung |
| Expected Outcome | erwartetes zulässiges Ergebnis |
| Expected Prohibited Outcome | Ergebnis, das **nicht** eintreten darf |
| Observation Points | wo beobachtet wird (Zustand, Record, Darstellung) |
| Evidence Requirements | welche Evidenz der Lauf erzeugen muss |
| Result Limitations | was das Ergebnis ausdrücklich nicht belegt |
| Applicability | Bedingungen, unter denen der Fall `not applicable` ist |

```text
test case references authority ≠ test case becomes authority
traceability                   ≠ requirement coverage
a referenced invariant         ≠ a validated invariant
```

Registriert als DEC-S-368.

## 11. Test Outcome Semantics

Das Vokabular ist **ausschließlich** auf Testausführung und Test-Evidenz bezogen. Es ist **kein** globales CoreOps-Statusmodell und wird nicht auf Betriebszustände, Capabilities, Deployments oder Domänenzustände angewandt.

| Result | Bedeutung |
| ------ | --------- |
| `not run` | nicht ausgeführt — weder Aussage noch Erwartung |
| `passed` | erwartetes Verhalten wurde an den deklarierten Observation Points beobachtet |
| `failed` | erwartetes Verhalten wurde nachweislich nicht beobachtet |
| `blocked` | Ausführung war nicht möglich (Precondition, Umgebung, Fixture, Autorisierung) |
| `inconclusive` | ausgeführt, aber Beobachtung unzureichend, widersprüchlich oder unzuverlässig |
| `not applicable` | Fall gilt für dieses Subjekt oder Profil ausdrücklich nicht |

```text
not run        ≠ passed
blocked        ≠ failed
inconclusive   ≠ passed
not applicable ≠ unknown
failed         ≠ vulnerability confirmed
passed         ≠ requirement universally satisfied
```

**Bewusst keine zusätzlichen Ergebniszustände.** Geprüft und verworfen:

- `partially passed` — würde `partial ≠ complete` aufweichen; Teilergebnisse gehören auf die Observation-Ebene, das Fallergebnis bleibt `failed` oder `inconclusive`, gemessen an der deklarierten Erwartung.
- `stale` / `expired` — ein Ergebnis altert nicht in einen neuen Testzustand, sondern wird zu **stale Evidenz** ausschließlich dort, wo die bestehende Evidence-Freshness-Semantik ([Evidence Model](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) §12) Staleness feststellt — **nicht** durch einen Revisionswechsel (§18).
- `error` — ein Harness- oder Laborfehler ist kein Subjektdefekt und wird als `blocked` (nicht ausführbar) oder `inconclusive` (unzuverlässig beobachtet) geführt.

Registriert als DEC-S-369.

## 12. Negative and Fail-Closed Testing

Negativszenarien sind **erstklassig**, nicht ergänzend. Ein Negativszenario ist genau dann `passed`, wenn das **spezifizierte geschützte Verhalten beobachtet** wurde.

```text
an error message appeared           ≠ protected behavior observed
operation rejected                  ≠ rejected for the specified reason
expected negative behavior observed ≠ system universally secure
```

Beispiel der geforderten Beobachtungstiefe:

```text
ungültige Autorisierung
→ privilegierte Aktion bleibt blockiert
→ kein Seiteneffekt am Ziel
→ Denial ist als Record sichtbar
→ Grund ist der spezifizierte Grund
```

nicht lediglich:

```text
eine Fehlermeldung erschien
```

Verpflichtende Negativ-Familien: `missing identity` · `missing approval` · `missing execution authorization` · `unknown target` · `wrong target` · `wrong environment` · `stale evidence` · `revoked artifact` · `invalid compatibility` · `unknown compatibility` · `partial state` · `conflict` · `missing telemetry` · `missing audit` · `offline authority expiry` · `secret cleanup unknown` · `recovery input invalid`.

Registriert als DEC-S-370.

## 13. Foundation Scenario Families

Jede Familie benennt die zu erhaltende Unterscheidung. Fixture-Klassen im [Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) §9; Umgebungsprofile im [Lab Companion](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md) §9.

### 13.1 State, Authority and Evidence

```text
unknown ≠ healthy · stale ≠ current · partial ≠ complete
conflicted ≠ resolved · revoked ≠ valid · not applicable ≠ unknown
missing evidence      ≠ negative evidence
evidence available    ≠ evidence sufficient
insufficient evidence ≠ evidence absent
authority mismatch    → fail-closed, keine stille Auflösung
```

### 13.2 Migration

```text
executed ≠ validated
partial migration bleibt partial und sichtbar
mixed-version state ist ein eigener Testzustand mit Compatibility-Grenze
rollback/recovery erfordert eigene Verifikationsevidenz
backup exists ≠ restorable · restore ≠ service recovery
```

### 13.3 Telemetry

```text
raw ≠ validated
missing ≠ zero ≠ inactivity ≠ target failure
ingested time ≠ observed time ≠ recording time
unknown/conflicting unit → keine automatische Konversion
sampled ≠ complete · derived ≠ independently observed
aggregate ≠ validated population
```

### 13.4 Deployment

```text
executed ≠ verified
partial wave bleibt per-target sichtbar
unknown outcome blockiert automatischen Retry
wave success ≠ remaining targets authorized
materielle Target-Set-Änderung → Re-Evaluation der Approval
cancellation ≠ Abwesenheit von Nebenwirkungen
```

### 13.5 Artifact, SBOM and Vulnerability

```text
artifact available ≠ trusted
SBOM available ≠ complete/accurate
component missing from SBOM ≠ component absent
integrity verified ≠ safe
vulnerability reference ≠ artifact affected
affected ≠ exploitable in this deployment
no finding ≠ no vulnerability
```

### 13.6 Offline and CorePack

```text
container integrity ≠ content trust
import ≠ activation authorization
quarantine release ≠ activation approval
partial activation ≠ complete success
evidence returned ≠ evidence complete
wrong target/environment → fail-closed
stale revocation state ≠ current revocation state
central unavailable ≠ local authority expands
```

### 13.7 Secrets

```text
retrieval ≠ use authorization
custody ≠ use authority
use ≠ export
cleanup requested ≠ secret absent
cleanup outcome unknown → fail-closed, niemals als „absent" gemeldet
masked secret bleibt secret-bearing
```

### 13.8 Classification, Retention and Disclosure

```text
unknown classification → fail-closed zur strengeren Grenze
stale classification ≠ current classification
container-inherited label ≠ automatically valid
retention expired ≠ deletion authorized ≠ deletion complete
primary copy deleted ≠ all copies removed
redacted ≠ disclosure-safe · masked/pseudonym ≠ anonymous
read ≠ export
evidence retention ≠ source retention
```

### 13.9 Self-Protection, Degraded and Recovery

```text
monitoring unavailable ≠ healthy
audit unavailable ≠ no action occurred
degraded ≠ permanent exception · degraded ≠ authority expansion
recovery input previously trusted ≠ currently trusted
service restored ≠ authority restored
partial recovery ≠ safe retry
recovery exit erfordert Reassessment, Verification, Reconciliation
break-glass ≠ permanent recovery authority
expired delegation ≠ silently extended delegation
```

### 13.10 UX and Dashboard

```text
hidden ≠ nonexistent
permission denied ≠ capability unavailable ≠ unknown ≠ restricted by operational mode
not applicable ≠ unknown
Simple ≠ fewer permissions · Expert ≠ more permissions · Expert ≠ Break Glass
dashboard summary ≠ authoritative state · summary ≠ completeness
UI confirmation ≠ approval · preview ≠ execution · selection ≠ target authorization
```

## 14. Integration Contract Testing

Die sechs Capability-Dimensionen des [Integration Contract v0.1](../architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md) §10 bleiben in Tests **getrennt**; keine wird aus einer anderen abgeleitet.

| Dimension | Was ein Test belegen darf | Was er nicht belegt |
| --------- | ------------------------- | ------------------- |
| `advertised` | eine Deklaration existiert und ist wohlgeformt | dass die Capability implementiert ist |
| `detected` | eine Erkennung hat stattgefunden | dass sie erlaubt ist |
| `permitted` | Policy erlaubt die Klasse im Scope | dass eine Ausführung autorisiert ist |
| `implemented` | Verhalten existiert gegen ein Test Double oder ein Lab-Ziel | dass ein realer Provider es unterstützt |
| `supported` | **nichts** — Support ist eine Governance-Entscheidung mit eigener Evidenzanforderung | dass Tests Support erzeugen |
| `validated` | eine Validierung wurde gegen einen deklarierten Umfang durchgeführt | universelle Gültigkeit |

```text
advertised ≠ implemented · implemented ≠ supported · supported ≠ validated universally
capability detected ≠ capability permitted · capability permitted ≠ execution authorized
```

**Consumer/Provider-Mismatch als Pflichtszenario:**

```text
provider process healthy ≠ protocol reachable ≠ consumer can use dependency
```

Dies ist eine **Testbarkeits- und Szenarioanforderung**. Die früher identifizierte Dependency-Contract-Health-Erweiterung wird hier ausdrücklich **nicht** implementiert und **nicht** registriert.

## 15. CO-WP-027 Design Assumptions as Test Subjects

Die folgenden Aussagen aus CO-WP-027 sind **Designannahmen**, die später zu prüfen sind — **keine** belegten Wahrheiten. Diese Strategie darf Methoden festlegen; sie darf **keine** Usability-Validierung behaupten.

| Annahme | Zu prüfende Frage | Zulässige Methode | Nicht belegbar |
| ------- | ----------------- | ----------------- | -------------- |
| 13 Top-Level-Informationsbereiche stützen Orientierung | finden Nutzer die zuständige Fläche ohne Fehlrouting? | strukturierte manuelle Aufgabenbeobachtung | Usability-Qualität |
| Presentation Context wird nicht als Autorisierung gelesen | wird `presentation context ≠ authorization scope` verstanden? | manuelle Interpretationsprüfung | Freiheit von Fehldeutung |
| 11 Darstellungsdimensionen bleiben verständlich | entsteht Status-Overload? | manuelle Dichte- und Verständnisprüfung | kognitive Belastung objektiv |
| `not applicable` und `unknown` sind unterscheidbar | sind beide nicht-visuell trennbar? | strukturierte Zustandsinventur plus manuelle Prüfung | Screenreader-Eignung |
| vier Unavailability-Ursachen bleiben getrennt | werden sie nie zu „nicht verfügbar" verschmolzen? | Zustandsinventur je Ursache | Vollständigkeit realer Ursachen |
| Simple/Expert kommuniziert Dichte, nicht Autorität | wird `Expert ≠ more permissions` verstanden? | manuelle Interpretationsprüfung | Abwesenheit von Fehlnutzung |
| Overview→List→Detail→Evidence erhält Kontext | bleiben Rückkehr- und Scope-Kontext erhalten? | Pfadverfolgung mit Kontextinventur | Navigationseffizienz |
| Dangerous-Action-Pfad bleibt verständlich | bleibt `confirmation ≠ approval` erkennbar? | Pfad- und Pflichtangabeninventur | Sicherheit vor Fehlbedienung |
| Dashboard-Priorisierung verbirgt nicht `unknown`/`stale`/`partial` | überlebt jede Dimension die Priorisierung? | Dimensionsinventur vor und nach Priorisierung | Aussagekraft der Priorisierung |

```text
assumption stated ≠ assumption validated
method specified  ≠ usability validated
```

## 16. Accessibility and Localization Evidence Requirements

Zukünftige Evidenzanforderungen — **ohne** jede Konformitäts-, Validierungs- oder Zertifizierungsaussage und **ohne** ausgewähltes oder ausgeführtes Prüfwerkzeug:

| Bereich | Geforderte künftige Evidenz |
| ------- | --------------------------- |
| Tastaturbedienbarkeit | vollständiger Tastaturpfad je Aktionsfläche, inklusive gefährlicher Aktionen |
| Fokus und Reihenfolge | dokumentierte Fokusreihenfolge und sichtbarer Fokuszustand je Ansicht |
| Nicht-nur-Farbe | Textäquivalent- und Strukturinventar je Zustand (`unknown`, `stale`, `partial`, `degraded`, `denied`, `not applicable`) |
| Textäquivalente | Äquivalent je nicht-textuellem Bedeutungsträger |
| Zustandsunterscheidung | nicht-visuelle Trennbarkeit der sechs Nullzustände und der vier Unavailability-Ursachen |
| Gefahrenkommunikation | nicht-visuelle Erkennbarkeit der acht Pflichtangaben des Dangerous-Action-Pfads |
| DE/EN | Textexpansionsnachweis ohne Bedeutungs- oder Zustandsverlust; Übersetzungsparität für Zustandsbegriffe |
| Reduzierte Bewegung | Bedeutungserhalt ohne Animation |
| Dichte Daten | Navigierbarkeit großer Listen und Tabellen ohne Kontextverlust |

```text
evidence requirement defined ≠ evidence produced
design requirement           ≠ WCAG conformance
keyboard path documented     ≠ screen-reader support proven
no accessibility tooling selected · no accessibility run performed
```

## 17. Coverage Model

Coverage ist **mehrdimensional**. Eine einzelne Prozentzahl ist als alleiniges Maß unzulässig.

Dimensionen: `requirements` · `invariants` · `decisions` · `risks` · `modules` · `capabilities` · `operations` · `failure modes` · `operational modes` · `offline profiles` · `data classes` · `authority boundaries` · `UX states` · `threat scenarios`.

Je Dimension wird geführt: Bezugsmenge · deklarierter Ausschnitt · abgedeckt · teilweise abgedeckt · nicht abgedeckt · **unbekannt** · nicht anwendbar.

```text
coverage metric              ≠ quality
100% of declared subset      ≠ 100% of Foundation
no uncovered item in matrix  ≠ no unknown requirement
dimension covered            ≠ dimension validated
```

Unbekannte Coverage bleibt explizit und wird nie als `0` oder als „abgedeckt" geführt.

**Aktueller Coverage-Stand:** in allen Dimensionen `not measured` — es existiert kein Testfall, kein Lauf und kein Ergebnis. Registriert als DEC-S-373.

## 18. Test Evidence and Reproducibility

Test-Evidenz wird unter dem **bestehenden** Evidence-, Audit- und Provenance-Modell geführt (§6). Konzeptionelle Mindestbindung eines künftigen Ergebnisses:

`test case revision` · `fixture revision` · `source revision` (Repository- und Contract-Stand) · `environment identity` · `target identity/class` · `configuration` · `time` (mit Clock-Unsicherheit, sofern zutreffend) · `observations` · `result` · `result limitations` · `artifact references` · `review status`, sofern zutreffend.

```text
evidence exists              ≠ evidence sufficient
reproducible once            ≠ universally reproducible
executor-produced result     ≠ independently verified result
evidence bound to a revision ≠ evidence valid for later revisions
missing evidence             ≠ negative result
```

Ein Ergebnis ist an die Revisionskombination gebunden, unter der es entstanden ist, und bleibt **historische Evidenz für genau diese Kombination**. Eine neue Test-Case-, Fixture- oder Source-Revision entwertet dieses Ergebnis **nicht** und macht es **nicht** automatisch `stale`:

```text
Evidenz zu Revision A         = historische Evidenz für Revision A
neue Revision B               ≠ Evidenz zu Revision A wird stale
Evidenz zu Revision A         ≠ Validierungsevidenz für Revision B
materielle Revisionsänderung  → Anwendbarkeit auf die neue Revision ist nicht etabliert
                              → Reassessment bzw. Revalidierung erforderlich
```

`stale` wird ausschließlich dort verwendet, wo das **bestehende** Evidence-Freshness-Modell ([Evidence Model](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) §12) Staleness tatsächlich feststellt. Historische Evidenz wird nicht umgeschrieben, nicht zurückgezogen und nicht neu bewertet; es entsteht **kein** eigener Test-Evidence-Lebenszyklus.

## 19. Existing CO-WP-028 Risk Coverage

Alle **21** derzeit auf `CO-WP-028` gerichteten Risk-Register-Einträge (aus dem sauberen Repository neu berechnet). Alle sind **abgedeckt**; keiner ist zurückgestellt, keiner ausgelassen. **Keiner wird durch dieses WP geschlossen** — eine dokumentierte Teststrategie ist keine Behandlungsevidenz.

| Risk ID | Failure Mode | Source | Test Layer | Fixture / Scenario Class | Expected Protected Behavior | Required Evidence | Limitations | Future Execution Gate |
| ------- | ------------ | ------ | ---------- | ------------------------ | --------------------------- | ----------------- | ----------- | --------------------- |
| RISK-233 | Ausgeführte Migration gilt als validiert | [Migration](../security/DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md) §13 | `TL-6`, `TL-2` | `nominal`, `partial`, `unknown`, `wrong-target` | Ausführung setzt keinen `validated`-Zustand; Validierung ist ein getrennter, evidenztragender Schritt; unvalidierte Migration blockiert abhängige privilegierte Operationen | getrennte Execution- und Validation-Records; `validated: no/unknown` explizit | Fixture-Datenmenge und -Heterogenität ≠ Produktion | Migration-Implementierung plus Lab |
| RISK-261 | Raw gilt als validiert; Normalisierung verliert Provenance | [Telemetry](../architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) §20 | `TL-2`, `TL-9` | `nominal`, `malformed`, `missing`, `conflicted` | Raw trägt nie `validated`; Normalisierung erhält Source-Provenance und Transformation History | Feld-Provenance im normalisierten Record; Validierungszustand getrennt | synthetische Signale ≠ reale Producer-Varianz | Telemetry-Normalisierung |
| RISK-263 | Unbekannte oder konfligierende Einheit still konvertiert | [Telemetry Mapping](../architecture/TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md) §12 | `TL-2`, `TL-7` | `unknown`, `conflicted`, `unsupported`, `incompatible` | Unknown/Conflicting Unit blockiert automatische Konversion und bleibt sichtbar; keine stille Kanonisierung | Ablehnungs-Record mit Grund und Feldbezug | reale Hersteller-Einheitenkodierungen nicht aufzählbar | Mapping-Implementierung |
| RISK-265 | Sampling-Charakter verborgen | [Telemetry](../architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) §24 | `TL-2`, `TL-10` | `partial`, `nominal` | Sampling-Metadaten überleben Ableitung, Aggregation und Darstellung; `sampled ≠ complete` bleibt sichtbar | Sampling-Attribut in Derived/Aggregated Output und in der Darstellung | Darstellungsprüfung manuell; keine Usability-Aussage | Telemetry- plus Experience-Implementierung |
| RISK-283 | Execution, Wave oder Partial als Erfolg | [Deployment](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md) §21 | `TL-6`, `TL-7` | `partial`, `unknown`, `wrong-target`, `nominal` | `executed ≠ verified`; Wave-Erfolg autorisiert den Rest nicht ohne explizite Expansionsentscheidung; Partial bleibt per-target sichtbar und blockiert unsafe Retry | Per-Target-Result-Set, Expansion-Decision-Record, Verification-Record | Lab-Target-Set ≠ Produktionstopologie und -skalierung | Deployment-Control-Plane plus Lab |
| RISK-287 | SBOM oder Integrity als Vollständigkeit/Sicherheit | [Artifact](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md) §18 | `TL-2`, `TL-7` | `partial` (unvollständige SBOM), `missing`, `nominal` | SBOM-Verfügbarkeit setzt keine Vollständigkeit; fehlende Komponente ist `unknown`, nicht `absent`; Integritätsprüfung setzt keine Sicherheitsaussage | Completeness-State (`unknown` zulässig) getrennt vom Integrity-State | synthetische SBOM ≠ reale Toolchain-Ausgabe | Artifact- und SBOM-Implementierung |
| RISK-288 | Vulnerability Reference als universelle Exploitability | [Artifact Dependency](../architecture/ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md) §16 | `TL-2`, `TL-9` | `nominal`, `unknown`, `not-applicable` | `vulnerability reference ≠ affected ≠ exploitable in this deployment`; „kein Fund" wird als `unknown` dargestellt, nicht als „nicht betroffen" | getrennte Candidate-, Applicability- und Confirmed-Zustände mit Match Confidence | keine reale Vulnerability-Quelle wird eingebunden | Vulnerability-Intelligence-Roadmap-Capability (**keine WP-Nummer vergeben**) |
| RISK-291 | CorePack-Contents erben Trust vom Container | [CorePack](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md) §6 | `TL-8`, `TL-7` | `nominal` Container mit `revoked`-, `incompatible`- oder `unsupported`-Inhalten | Container-Integrität setzt keinen Content-Trust; jeder Inhalt wird eigenständig bewertet; ein untrusted Inhalt vertraut Geschwister nicht implizit | Per-Content-Trust- und Compatibility-Assessment-Records | synthetisches Pack ≠ reale Publisher-Distribution | CorePack-Implementierung plus isoliertes Profil |
| RISK-293 | Import, Partial oder Evidence als Activation bzw. Vollständigkeit | [Offline Trust](../security/OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md) §21 | `TL-8`, `TL-7` | `partial`, `unknown`, `permission-denied` | `quarantine release ≠ activation approval`; `import assessment passed ≠ approved`; Partial Activation bleibt partial; `evidence returned ≠ evidence complete` | getrennte Quarantine-, Assessment-, Approval- und Activation-Records; Evidence-Completeness darf `unknown` sein | Air-Gapped-Profil ist simuliert; kein physischer Transfer | Offline-Trust-Implementierung plus Air-Gapped-Profil |
| RISK-299 | Retrieval als Use; Custody und Use vermischt; Cleanup-Persistenz | [Secrets](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) §11 | `TL-4`, `TL-7`, `TL-9` | Platzhalter-Referenzen; `permission-denied`, `expired`, `revoked`, `unknown` Cleanup-Ausgang | `retrieval ≠ use ≠ export`; `custody ≠ use authority`; unbekannter Cleanup-Ausgang failt closed und wird nie als „absent" gemeldet | getrennte Retrieval-, Use- und Cleanup-Records; Cleanup-State `unknown` zulässig und sichtbar | **kein reales Secret wird verwendet**; Key-Material-Handhabung wird nicht geübt | Secrets-Implementierung plus credential-safe Lab |
| RISK-300 | Falsche oder stale Klassifikation führt zu unzulässigem Handling | [Classification](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) §7 | `TL-9`, `TL-7` | `missing`, `stale`, `unknown`, container-inherited | Unknown-Klassifikation failt closed zur strengeren Grenze; geerbtes Label ist nicht automatisch gültig; materielle Änderung löst Reassessment aus | Classification-Decision-Record mit Assessment Time und Freshness Boundary | Fixture-Klassen ≠ reale Datenheterogenität | Classification-Implementierung |
| RISK-301 | Retention Expiry als Löschabschluss; Kopien außerhalb der Löschung | [Retention](../governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md) §12 | `TL-9`, `TL-6` | `nominal` plus Derived-, Backup- und Cache-Kopien, `partial`, `unknown` | `expiry ≠ deletion authorized ≠ deletion complete`; Primärlöschung entfernt nicht alle Kopien; unbekannter Ausgang wird nicht als gelöscht gemeldet | Copy-Inventar plus Per-Copy-Deletion-Verification; verbleibendes `unknown` sichtbar | Lab-Copy-Inventar ist synthetisch und endlich | Retention- und Deletion-Implementierung |
| RISK-302 | Redaction als sichere oder anonyme Disclosure; Lineage-Export | [Redaction](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md) §9 | `TL-9`, `TL-7` | `nominal`, `partial` Lineage, `boundary` | `redacted ≠ disclosure-safe`; `masked/pseudonym ≠ anonymous`; `read ≠ export`; unvollständige Derived-Data-Lineage blockiert Export | Residual-Disclosure-Assessment plus Lineage-Completeness-State | **kein Re-Identifikationstest**; Restrisiko wird nicht quantifiziert | Redaction- und Disclosure-Implementierung |
| RISK-304 | Secret-Downgrade; Cross-Workspace-Reuse; Evidence als Source Retention | [Classification](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) §12 | `TL-9`, `TL-7` | maskiert `secret-bearing`, cross-workspace, evidence-retention | Maskiertes Secret bleibt `secret-bearing`; Cross-Workspace-Wiederverwendung wird verweigert; Evidence Retention verlängert Source Retention nicht | Klassifikationspersistenz-Record; Isolation-Denial-Record; getrennte Retention-Uhren | nur Platzhalter-Secrets | Classification- plus Secrets-Implementierung |
| RISK-305 | Kompromittierter Control Plane behält privilegierte Fähigkeiten; Monitoring-/Audit-Ausfall als gesund | [Self-Protection](../security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md) §11 | `TL-7`, `TL-8` | `degraded`, `missing` (Audit/Telemetry), `unknown` | `monitoring/audit unavailable ≠ healthy`; fehlendes Audit impliziert keine Nichtausführung; privilegierte und destruktive Capabilities werden suspendiert | Capability-Suspension-Record plus expliziter `unknown`-Health-State | **simulierter Ausfall ≠ Nachweis realer Resilienz** | Self-Protection-Implementierung plus Degraded-Profil |
| RISK-306 | Degraded Mode wird dauerhafte Umgehung | [Degraded](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) §11 | `TL-8`, `TL-4` | `degraded` mit `expired` Gültigkeit | Degraded ist benannt, befristet, reviewpflichtig und widerrufbar mit Exit-Bedingung; Ablauf verlängert nicht still; `degraded ≠ authority expansion` | Mode-Entry- und Exit-Records mit Validity und Review-Referenz | Langzeitdrift in kurzen Läufen nicht beobachtbar; Zeit ist simuliert | Degraded-Mode-Implementierung |
| RISK-307 | Recovery nutzt stale, revoked oder inkompatible Inputs bzw. stellt nur Dienste her | [Recovery](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md) §10 | `TL-6`, `TL-7`, `TL-8` | `revoked`, `stale`, `incompatible`, `wrong-environment` | `previously trusted ≠ currently trusted`; Recovery-Inputs werden gegen den aktuellen Trust-, Revocation- und Compatibility-Stand neu bewertet; `service restored ≠ authority restored` | Input-Reassessment-Record; getrennte Service- und Authority-Restoration-States | **Lab-Recovery-Erfolg ≠ Produktions-Recovery-Bereitschaft** | Recovery-Implementierung plus Recovery-Profil |
| RISK-308 | Partial oder Unknown Recovery, falscher Fault Scope, verfrühter Exit | [Recovery](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md) §16 | `TL-6`, `TL-7` | `partial`, `unknown`, `conflicted` | Unknown Outcome blockiert unsafe Retry; Containment Scope ist explizit; Exit erfordert Reassessment, Verification und Reconciliation | Per-Stage-Checkpoint-Records; Exit-Decision mit Reassessment-Referenz | **ein erfolgreicher Lauf ≠ Regressionsvertrauen** | Recovery-Implementierung |
| RISK-309 | Break-Glass wird permanente Recovery-Autorität; Offline-Agent überschreitet Delegation | [Recovery](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md) §13 | `TL-4`, `TL-7`, `TL-8` | `expired` Delegation, `revoked`, `boundary` | Break-Glass bleibt außergewöhnlich und befristet; abgelaufene Delegation stoppt privilegierte Aktionen fail-closed; `agent operational ≠ centrally current` | Delegation-Expiry-Enforcement-Record; Denied-Action-Record | Clock-Manipulationsszenarien sind synthetisch | Offline-/Agent- plus Recovery-Implementierung |
| RISK-311 | Ausgeblendetes als nicht vorhanden; Berechtigung, Fähigkeit, Unknown und Mode vermischt | [IA](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) §13 | `TL-10`, `TL-3` | vier Unavailability-Ursachen plus sechs Nullzustände | Die vier Ursachen bleiben unterscheidbar und werden nie zu „nicht verfügbar" verschmolzen; Ausgeblendetes wird nie als nicht existent dargestellt; Simple/Expert ändert nur Dichte | Per-State-Presentation-Record; programmatisch bestimmbare Ursachenklasse; nicht-visuelle Unterscheidbarkeit | **nur manuelle oder assistierte Prüfung; keine Usability-Validierung** | Experience-Implementierung |
| RISK-313 | Statusbedeutung ohne Farb-, Positions- oder Bewegungswahrnehmung verloren | [UX Policy](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md) §16 | `TL-10` | `nominal`, `boundary` (DE/EN-Expansion), `degraded`, reduzierte Bewegung | `unknown`, `stale`, `partial`, `degraded` und `denied` bleiben ohne Farbe, Position und Bewegung durch Textäquivalent und Struktur unterscheidbar; Tastaturbedienbarkeit und Fokusreihenfolge bleiben erhalten | Textäquivalent-Inventar je Zustand; Tastaturpfad-Record; DE/EN-Expansionsnachweis — als **Designanforderungs-Evidenz** | **keine WCAG-Konformität, kein Screenreader-Nachweis, keine Accessibility-Validierung, keine Zertifizierung, kein Werkzeug ausgewählt oder ausgeführt** | Experience-Implementierung plus gesonderte Accessibility-Validierungs-Autorisierung |

```text
risk mapped into a test strategy ≠ risk treated
risk covered in this matrix      ≠ risk closed
```

## 20. Execution Gates

Kein Testfall dieser Strategie ist heute ausführbar. Eine spätere Ausführung erfordert kumulativ:

1. Implementierung des betroffenen Moduls oder der Capability.
2. Technologieauswahl über ADR-Arbeit (§24).
3. Bereitgestelltes Integration Lab mit deklarierter Environment Identity ([Lab Companion](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md) §23).
4. Freigegebene Fixtures mit Identity, Revision und Provenance ([Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) §7).
5. Explizite Human-Maintainer-Autorisierung der Testausführung.

Eine spätere Implementierung darf ausführbare Tests **ergänzen**, ohne §7, §10, §11, §12 oder §18 zu ändern: Ebenen, Traceability-Felder, Ergebnisvokabular, Negativsemantik und Evidenzbindung sind implementierungsunabhängig formuliert.

## 21. Exit and Acceptance Semantics

CO-WP-028 erzeugt **Test-Governance**, keine erfolgreiche Systemvalidierung.

Etabliert: Teststrategie dokumentiert · Fixture-Governance dokumentiert · Integration-Lab-Contract dokumentiert · Risk-to-Test-Mapping dokumentiert.

**Nicht** etabliert: Tests implementiert · Tests ausgeführt · Foundation validiert · Sicherheit verifiziert · Recovery nachgewiesen · Offline-Betrieb nachgewiesen · UX validiert · Accessibility validiert · Herstellersupport · Produktionsreife · Releasereife.

Foundation Readiness bleibt `CO-WP-030`.

## 22. Security Invariants

1. Test-Evidenz erzeugt keine Autorität, keine Freigabe und keinen autoritativen Zustand.
2. Es entsteht keine zweite Evidence-, Approval- oder Statusautorität.
3. Keine Aussage überschreitet eine Grenze des Claim Boundary Set (§7); das Set ist eine Menge unabhängiger Nicht-Implikationen und **keine** universelle Reifeleiter (`claim boundary set ≠ universal maturity ladder`).
4. Ein Negativszenario ist nur bei beobachtetem geschütztem Verhalten `passed`.
5. `not run`, `blocked`, `inconclusive` und `not applicable` werden nie als `passed` geführt.
6. `not applicable ≠ unknown` bleibt in Ergebnissen und Coverage erhalten.
7. Kein Testartefakt enthält reale Secrets, wiederverwendbare Credentials oder unredigierte personenbezogene Daten.
8. Ein Lab- oder Simulatorergebnis belegt weder Produktionsreife noch Providerkompatibilität noch Support.
9. Testdaten unterliegen den bestehenden Klassifikations-, Redaction- und Retention-Grenzen.
10. Unbekannte Coverage bleibt explizit und wird nie als abgedeckt geführt.
11. Kein Risiko wird durch die Existenz einer Teststrategie geschlossen.
12. Accessibility-Anforderungen erzeugen keine Konformitäts-, Validierungs- oder Zertifizierungsaussage.

## 23. Threat References

Referenziert, nicht dupliziert: THR-001, THR-002, THR-013, THR-016, THR-018, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-025, THR-028, THR-029, THR-031, THR-032, THR-033, THR-034, THR-035, THR-036 und THR-038 aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md). Diese Strategie benennt Testbarkeitsanforderungen zu diesen Szenarien und erhebt **keinen** Verifikationsanspruch.

## 24. Technology Boundary

Test-Framework, Programmiersprache, CI-Provider, Container- und Virtualisierungstechnologie, Mock-Framework, Browser-Automation, Load-Test-Werkzeug, Accessibility-Scanner, Security-Scanner, Datenbank, Message Broker, Lab-Orchestrator, Netzwerkemulator, Test-Management-Plattform, Coverage-Werkzeug und Reporting-Werkzeug: **nicht ausgewählt**. Die Auswahl gehört in spätere Implementierungs- und ADR-Arbeit; dieses Dokument registriert dazu **keine** Deferral-Decision, da die Technologiegrenze bereits Governance-Grenze der Non-Goals ist.

## 25. Compatibility

Ergänzt CO-WP-004…027, ohne ein Foundation-Dokument zu ändern. Bindet an: Integration Contract v0.1 (§14), Evidence- und Audit-Modell (§6, §18), Policy/Approval/Execution (§13.1), Migration (§13.2), Telemetry (§13.3), Deployment (§13.4), Artifact/SBOM (§13.5), Offline/CorePack (§13.6), Secrets (§13.7), Classification/Retention/Redaction (§13.8), Self-Protection/Degraded/Recovery (§13.9) sowie UX/Dashboard (§13.10, §15, §16).

## 26. Open Questions

- Ob `TL-1` vor der Implementierung eigenständig ausführbar gemacht wird (Repository-Konsistenzprüfung) — die Entscheidung gehört zu späterer Tooling- und ADR-Arbeit.
- Wie die Target-WP-Zuordnung der 21 auf `CO-WP-028` gerichteten Risiken nach diesem WP fortgeschrieben wird (Strategieanteil dokumentiert, Ausführungsanteil offen) — Follow-up `CO-WP-029`/`CO-WP-030`.
- Ob ein späterer CDS-Pilot eigene Consumer-Evidenz beisteuern darf — erfordert gesonderte Autorisierung, siehe [Lab Companion](INTEGRATION_LAB_SCENARIO_AND_EVIDENCE_MODEL.md) §29.

## 27. Next Decision

Nova Review dieses Dokuments und der beiden Companions. Danach Human-Maintainer-Entscheidung über Staging und Commit. **Keine** Testimplementierung, **keine** Testausführung, **keine** Laborbereitstellung und **kein** Start von `CO-WP-029` durch dieses Dokument.
