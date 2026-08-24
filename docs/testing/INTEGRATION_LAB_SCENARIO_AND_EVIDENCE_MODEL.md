# CoreOps – Integration Lab, Scenario and Evidence Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation integration-lab, scenario and evidence model
> Implementation Status: Not implemented
> Integration Lab Status: Not provisioned
> Lab Environment Status: Not created
> Test Execution Status: Not performed
> Validation Status: Not performed
> Technology Selection: None
> Container/Virtualization/Orchestrator/Network-Emulator/Simulator/CI Technology: Not selected
> Target Connection: None established
> External API Calls: None performed
> Production Validation: Not performed
> Accessibility Validation: Not performed
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-028 (docs-only / test strategy, fixtures and integration lab)

## 1. Status

Technologieunabhängiges Modell eines **logischen Integration Lab** für CoreOps: Autoritätsgrenze, Lab-Identität, konzeptionelle Rollen, Umgebungsprofile, Test-Double-Klassen, Fidelity-Deklaration, Sicherheitskontrollen, Zielbindung, Credential-Grenze, Reset und Disposability, Szenariomodell, Integration-Contract-Szenarien, Degraded-/Offline-/Recovery-Übung, Lab-Evidenz und Reproduzierbarkeitsgrenzen. Companion zu [FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) und [SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md).

**Das Lab ist nicht bereitgestellt.** Es wurde keine Umgebung erzeugt, kein Container gestartet, keine VM provisioniert, keine Verbindung zu einem Ziel aufgebaut und kein externer Aufruf ausgeführt.

## 2. Purpose

Ein Labor, das die Produktion erreichen kann, ist kein Labor. Das Modell legt fest, warum `lab role ≠ production authority`, `lab target ≠ production target`, `lab credential ≠ production credential`, `lab approval ≠ production deployment authorization`, `simulator success ≠ provider compatibility` und `integration-lab success ≠ production readiness` — und wie CoreOps-Betriebsszenarien später beobachtbar geübt werden können, **ohne** dass eine Übung als Betriebsnachweis gelesen wird.

## 3. Scope

Lab-Autoritätsgrenze · Lab Identity und Environment Declaration · konzeptionelle Lab-Rollen · Umgebungsprofile · Test-Double-Klassen · Fidelity-Deklaration · Lab-Sicherheitskontrollen · Zielbindung und Produktionsreichweite · Credential-Grenze · Reset/Cleanup/Disposability · Szenariomodell · Integration-Contract-Szenarien · Degraded-/Offline-/Recovery-Übung · Lab-Evidenz · Reproduzierbarkeit · Observation Points · Ergebnis-Interpretationsgrenzen · Provisioning Gate · Failure und Unknown · CDS-Grenze.

## 4. Non-Goals

- **Keine** Bereitstellung, Provisionierung, Konfiguration oder Inbetriebnahme eines Labors.
- **Keine** Auswahl von Container-, Virtualisierungs-, Orchestrierungs-, Netzwerkemulations-, Simulator-, CI- oder Test-Management-Technologie.
- **Keine** Verbindung zu einem Zielsystem, kein externer API-Aufruf, kein Scan, kein Deployment, keine Recovery-Ausführung.
- **Keine** neuen Laufzeitmodule, keine neuen Autoritäten und kein zweites Betriebsmodusmodell.
- **Keine** Aussage über Produktionsreife, Herstellersupport, Resilienz oder Sicherheit.

## 5. Concepts

Mindestbegriffe: `integration lab` · `lab environment` · `lab environment identity` · `lab environment profile` · `lab role` · `subject-under-test instance` · `controlled test target` · `test double` · `fidelity declaration` · `scenario` · `scenario step` · `stimulus` · `observation point` · `lab run` · `lab evidence` · `reset` · `provisioning gate`.

```text
lab environment    ≠ deployment environment
lab run            ≠ operation
lab evidence       ≠ operational evidence for production
scenario           ≠ test case
controlled target  ≠ managed production target
```

## 6. Lab Authority Boundary

Das Lab ist eine **Beobachtungsumgebung**, keine Autorität und kein Governance-Träger.

```text
lab role                ≠ production authority
lab environment         ≠ production environment
lab target              ≠ production target
lab credential          ≠ production credential
lab approval            ≠ production deployment authorization
lab recovery success    ≠ production recovery readiness
lab profile             ≠ CoreOps operational authority
simulated degraded state ≠ evidence of real degraded-mode resilience
integration-lab success ≠ production readiness
```

Ein Lab-Ergebnis ändert keinen autoritativen CoreOps-Zustand, erteilt keine Freigabe und erzeugt keinen Support-Status. Registriert als DEC-S-372.

## 7. Lab Identity and Environment Declaration

Jede künftige Lab-Umgebung führt konzeptionell mindestens: `lab environment identity` (stabil, nicht wiederverwendbar) · `environment class` (ausdrücklich `non-production`) · `environment profile` (§9) · `owner` · `workspace binding` · `permitted target set` · `credential boundary` · `network boundary` · `data boundary` · `reset expectation` · `evidence expectation` · `known limitations` · `lifecycle state`.

```text
same lab name        ≠ same lab environment identity
environment declared ≠ environment verified
non-production label ≠ non-production reach
```

Eine Umgebung ohne deklarierte Identität und Profil ist **nicht** verwendbar und erzeugt keine referenzierbare Evidenz.

## 8. Lab Roles

Konzeptionelle Rollen — **keine** Laufzeitmodule, **keine** neuen `MOD-*`-Einträge, **keine** neuen Autoritäten:

| Rolle | Aufgabe | Ausdrücklich nicht |
| ----- | ------- | ------------------ |
| CoreOps Subject-under-Test | die geprüfte CoreOps-Instanz | keine produktive Control-Plane-Instanz |
| Managed-Target Simulator oder kontrolliertes Ziel | stellt ein Zielsystemverhalten bereit | kein reales verwaltetes Produktionsziel |
| Adapter- oder Provider-Test-Double | ersetzt eine Integration an ihrer Contract-Grenze | kein Herstellerprodukt und kein Support-Nachweis |
| Evidence Collector | sammelt Beobachtungen zu einem Lauf | keine Evidence-Autorität; erzeugt keine Gültigkeit |
| Test Controller | löst Szenarioschritte aus und protokolliert sie | keine Approval-Instanz und keine Execution Authority |
| Human Observer / Reviewer | beobachtet, bewertet, dokumentiert Grenzen | keine automatische Freigabe |

```text
lab role                 ≠ CoreOps module
evidence collector       ≠ evidence authority
test controller          ≠ execution authorization
human observer in a lab  ≠ approver of a production action
```

## 9. Lab Environment Profiles

Profile bilden **bestehende** CoreOps-Konzepte ab; es entsteht **kein** paralleles globales Modusmodell. Ein Profil ist eine **Laborbedingung**, kein CoreOps-Betriebszustand.

| Profil | Bildet ab | Übungszweck | Ausdrückliche Grenze |
| ------ | --------- | ----------- | -------------------- |
| `lab-connected` | Connectivity-Klasse `connected` | Regelbetrieb mit erreichbarer zentraler Autorität | keine Produktionslast, keine Produktionstopologie |
| `lab-restricted` | `restricted-connected` und `intermittently-connected` | gefilterte, unidirektionale oder lückenhafte Pfade, Freshness-Grenzen | simulierte Filterung ≠ reale Netzarchitektur |
| `lab-isolated` | `isolated` | kontrollierter Transfer, Import- und Activation-Grenzen | kein physischer Transfer |
| `lab-air-gapped` | `air-gapped` | vollständige Trennung, CorePack-Import, Evidence Return | simuliert; **kein** Nachweis der Air-Gap-Eignung |
| `lab-degraded` | Operational Mode `degraded` sowie `guarded`/`restricted`/`read-only` | Capability-Suspension, Sichtbarkeit von `unknown` | simulierter Ausfall ≠ reale Resilienz |
| `lab-recovery` | `recovery-only` und Recovery Mode | Recovery-Stufen, Input-Reassessment, Exit-Bedingungen | Lab-Recovery-Erfolg ≠ Produktions-Recovery-Bereitschaft |
| `lab-unknown-connectivity` | `unknown-connectivity` | konservative Behandlung bei unbestimmter Konnektivität | `unknown ≠ safe` |

```text
lab profile        ≠ CoreOps operational mode authority
profile named      ≠ condition faithfully reproduced
air-gapped profile ≠ statement of suitability for classified networks
```

Es wird **keine** Eignungsaussage für klassifizierte oder VS-Netze getroffen.

## 10. Test Double Classes

Konzeptionelle Klassen von Testquellen — **keine** Technologieauswahl:

| Klasse | Charakter | Zulässige Aussage | Unzulässige Aussage |
| ------ | --------- | ----------------- | ------------------- |
| `stub` | liefert feste, minimale Antworten | Aufrufpfad existiert | Verhalten des echten Gegenübers |
| `fake` | vereinfachte, aber funktionsfähige Ersatzimplementierung | interne Logik reagiert konsistent | Produktverhalten |
| `simulator` | modelliert Zielverhalten einschließlich Fehlerfällen | Szenario ist beobachtbar durchlaufbar | Providerkompatibilität |
| `emulator` | bildet ein Protokoll oder eine Schnittstelle näher nach | Protokollpfad ist begehbar | Herstellerkonformität |
| `recorded synthetic response` | vorab konstruierte Antwortfolge | deterministische Wiederholbarkeit | Repräsentativität realer Antworten |
| `controlled real test target` | reales, aber ausschließlich für das Lab bereitgestelltes Ziel | Verhalten eines konkreten Exemplars unter Laborbedingungen | Familien-, Firmware- oder Produktaussage |

```text
simulator pass        ≠ real provider validation
mocked success        ≠ integration success
test double behavior  ≠ external product behavior
recorded response     ≠ current provider behavior
controlled real target ≠ supported target
```

Registriert als DEC-S-372.

## 11. Fidelity Declaration

Jede Test-Double-Instanz **muss** ihre Fidelity deklarieren: nachgebildeter Umfang · nicht nachgebildeter Umfang · Fehlerfälle, die sie erzeugen kann · Fehlerfälle, die sie **nicht** erzeugen kann · Herkunft des nachgebildeten Verhaltens · bekannte Abweichungen.

```text
fidelity declared ≠ fidelity verified
high fidelity     ≠ equivalence
passing against a double ≠ evidence about the real system
undeclared fidelity → Ergebnis trägt eine unbekannte Limitation
```

Ein Ergebnis gegen ein Test Double ohne Fidelity-Deklaration wird als `inconclusive` geführt.

## 12. Lab Safety Controls

Neun künftige Kontrollbereiche. Sie sind **Anforderungen an ein späteres Labor**, keine implementierten Kontrollen:

| Kontrolle | Anforderung |
| --------- | ----------- |
| Environment Identity | jede Umgebung eindeutig, `non-production` deklariert, in jeder Evidenz referenziert |
| Target Binding | ausschließlich Ziele aus dem deklarierten `permitted target set` (§13) |
| Network Boundary | keine Route, kein Namensdienst und kein Proxy-Pfad zu Produktionsnetzen |
| Credential Boundary | ausschließlich laborgebundene, nicht produktionsgültige Credentials (§14) |
| Test-Data Boundary | ausschließlich freigegebene Fixtures und Testdaten nach dem [Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) |
| Reset und Cleanup | definierter, wiederholbarer Rücksetzpfad; unbekannter Cleanup-Ausgang failt closed (§15) |
| Artifact Provenance | jedes eingebrachte Artefakt trägt Identity, Provenance und Integritätsangabe |
| Test Execution Authorization | Ausführung erfordert explizite Autorisierung; Laborzugang ist keine Ausführungsberechtigung |
| Evidence Capture | Beobachtungen werden unter dem bestehenden Evidence-Modell erfasst (§19) |

Zusätzlich gilt die **Public-Neutrality-Grenze**: Lab-Konfigurationen, Szenariodefinitionen und Evidenzartefakte im Repository enthalten keine realen privaten Hostnamen, Domainnamen, Netzbereiche, Organisationsnamen oder personenbezogenen Daten ([PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md)).

## 13. Target Binding and Production Reach

Das Lab darf **niemals still** Produktionsreichweite erlangen.

- Ein Ziel ist entweder im deklarierten `permitted target set` oder es ist unzulässig — es gibt keinen impliziten Pfad.
- Ein Ziel außerhalb der Menge ist ein `wrong-target`-**Negativfall**: erwartetes Verhalten ist die Verweigerung, nicht die Ausführung.
- Eine Erweiterung des Target Set ist eine explizite, zurechenbare Entscheidung und löst eine Re-Evaluation aus.
- Ein Produktionsziel ist in keinem Profil zulässig.

```text
reachable        ≠ permitted
permitted        ≠ authorized for execution
target set change ≠ silent continuation of prior approval
lab target       ≠ production target
```

## 14. Credential Boundary

- Ausschließlich laborgebundene Credentials mit begrenzter Gültigkeit und begrenztem Scope.
- Kein Produktions-Credential, kein produktionsgültiges Zertifikat, kein produktives Schlüsselmaterial — auch nicht abgelaufen.
- Kein Credential-Wert in Szenariodefinitionen, Protokollen, Fehlerausgaben oder Evidenzartefakten; ausschließlich Referenzen ohne Wert.
- Custody, Rotation und Revocation von Lab-Credentials folgen [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](../security/SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md); das Lab erzeugt **kein** eigenes Secret-Modell.

```text
lab credential      ≠ production credential
credential retrieval ≠ use authorization
lab access          ≠ execution authorization
expired lab credential ≠ safe to disclose
```

## 15. Reset, Cleanup and Disposability

Das Lab ist **verwerfbar und rücksetzbar**.

- Ein Szenariolauf beginnt aus einem deklarierten Ausgangszustand.
- Reset ist Teil des Szenarios, nicht ein optionaler Nachlauf.
- Ein unbekannter Reset- oder Cleanup-Ausgang wird als `unknown` geführt und blockiert den nächsten evidenztragenden Lauf.
- Restdaten aus einem Lauf unterliegen der Retention-Grenze des [Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md) §17.

```text
reset requested   ≠ reset complete
cleanup outcome unknown ≠ environment clean
environment reused ≠ environment equivalent to its prior state
disposable         ≠ automatically clean
```

## 16. Scenario Model

Ein Szenario ist eine geordnete Folge von Stimuli mit deklarierten Beobachtungspunkten. Konzeptionelle Mindestangaben: `scenario identity` · `scenario revision` · `scenario family` (Teststrategie §13) · `test levels` · `required environment profile` · `required lab roles` · `fixture set` · `preconditions` · `steps` · `observation points` · `expected outcome` · `expected prohibited outcome` · `reset expectation` · `evidence requirements` · `result limitations`.

```text
scenario                ≠ test case
scenario executed       ≠ scenario passed
scenario step succeeded ≠ scenario outcome established
scenario revision unchanged ≠ meaning unchanged when its source contract changed
```

## 17. Integration Contract Scenario Families

Pflichtszenarien an der Integrationsgrenze, aufbauend auf [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md):

- **Capability-Dimensionen getrennt:** je Dimension ein eigener Beobachtungspunkt (Teststrategie §14).
- **Acceptance-Kette:** `request acceptance ≠ authorization ≠ execution`; `completion ≠ success`; `success ≠ verification`.
- **Unknown Outcome:** blockiert automatischen Retry und erfordert Reconciliation.
- **Read-only-Grenze:** eine read-only-Integration erhält nicht still Write Authority.
- **Consumer/Provider-Mismatch:** `provider process healthy ≠ protocol reachable ≠ consumer can use dependency` — als Szenarienfamilie gefordert. Die früher identifizierte Dependency-Contract-Health-Erweiterung wird **nicht** implementiert und **nicht** registriert.
- **Contract-Version-Mismatch:** `unsupported`, `incompatible` und `unknown compatibility` sind getrennte Szenarien; `unknown compatibility ≠ compatibility`.
- **Adapter-Scope:** ein Adapter erweitert Ziel, Aktion oder Scope nicht; die Verweigerung ist das erwartete Verhalten.

## 18. Degraded, Offline and Recovery Exercise Model

Diese Profile sind **Übungen**, keine Nachweise.

- **Degraded:** Capability-Suspension, Sichtbarkeit von `unknown`, benannte und befristete Gültigkeit mit Exit-Bedingung; `degraded ≠ authority expansion`.
- **Offline und CorePack:** Import, Quarantäne, Assessment, Approval, Activation und Evidence Return als **getrennte** Beobachtungspunkte; `container integrity ≠ content trust`; `evidence returned ≠ evidence complete`.
- **Recovery:** Input-Reassessment gegen den aktuellen Trust-, Revocation- und Compatibility-Stand; `previously trusted ≠ currently trusted`; `service restored ≠ authority restored`; Exit erfordert Reassessment, Verification und Reconciliation.
- **Clock- und Freshness-Unsicherheit:** simulierte Zeit ist als simuliert zu kennzeichnen; ein Ergebnis unter simulierter Zeit trägt diese Limitation.

```text
simulated outage      ≠ observed real outage
exercise completed    ≠ resilience demonstrated
recovery rehearsed    ≠ recovery readiness established
offline profile passed ≠ offline operation proven
```

## 19. Lab Evidence Model

Lab-Evidenz wird unter dem **bestehenden** Evidence-, Audit- und Provenance-Modell geführt ([EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [EVENT_AND_AUDIT_CORRELATION_MODEL.md](../architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md), MOD-EVD-001). Es entsteht **keine** zweite Evidence-Autorität.

Konzeptionelle Mindestbindung eines Lauf-Ergebnisses: `scenario revision` · `test case revisions` · `fixture revisions` · `source revision` · `lab environment identity` · `environment profile` · `lab role instances` · `test double fidelity declarations` · `target identity/class` · `configuration` · `time` mit Clock-Kennzeichnung · `observations` je Observation Point · `result` · `result limitations` · `artifact references` · `reviewer note`, sofern zutreffend.

```text
lab evidence            ≠ operational authority
lab evidence            ≠ production evidence
evidence exists         ≠ evidence sufficient
evidence collected      ≠ evidence validated
evidence returned       ≠ evidence complete
```

## 20. Reproducibility Semantics

```text
reproducible once           ≠ universally reproducible
same scenario revision      ≠ same environment
same environment identity   ≠ same environment state
deterministic double        ≠ deterministic system
executor-produced result    ≠ independently verified result
```

Ein Ergebnis gilt für die Kombination aus Szenario-, Testfall-, Fixture-, Source- und Environment-Revision, unter der es entstanden ist, und bleibt für diese Kombination **historische Evidenz**. Ändert sich eine davon materiell, wird das Ergebnis **nicht** automatisch `stale`; seine **Anwendbarkeit auf die neue Kombination ist nicht etabliert** und erfordert Reassessment bzw. Revalidierung (Teststrategie §18).

```text
Evidenz zu Revision A ≠ Validierungsevidenz für Revision B
neue Revision         ≠ automatische Staleness der bestehenden Evidenz
```

## 21. Observation Points

Ein Szenario beobachtet mindestens dort, wo eine Grenze gehalten werden muss: Autorisierungsentscheidung · Ausführungsauslösung · Zielwirkung oder deren Ausbleiben · Zustandsänderung oder deren Ausbleiben · Audit- und Evidenzerzeugung · Fehler- und Ablehnungsgrund · Darstellung, sofern die Experience-Ebene beteiligt ist.

```text
observing an outcome ≠ observing the absence of a side effect
absence not observed ≠ absence established
missing observation  → inconclusive, nicht passed
```

Ein Negativszenario ohne Beobachtung des **Ausbleibens** der Wirkung ist `inconclusive`, nicht `passed`.

## 22. Result Interpretation Limits

Jedes künftige Lab-Ergebnis trägt seine Grenzen mit:

```text
lab success             ≠ production readiness
lab success             ≠ security verification
lab success             ≠ vendor support
lab success             ≠ regression confidence
one environment profile ≠ all operational conditions
one controlled target   ≠ a target family
no observed failure     ≠ absence of defect
```

Ein Lab-Ergebnis darf in Release-, Readiness- oder Support-Aussagen **nicht** als Validierung, Verifikation oder Zusicherung verwendet werden.

## 23. Lab Provisioning Gate

Eine spätere Bereitstellung erfordert kumulativ:

1. explizite Human-Maintainer-Autorisierung der Laborbereitstellung;
2. Technologieauswahl über ADR-Arbeit (§27);
3. deklarierte Lab Environment Identity und Profil (§7);
4. deklariertes und begrenztes `permitted target set` (§13);
5. laborgebundene Credential-Grenze (§14);
6. freigegebene Fixtures und Testdaten ([Fixture Companion](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md));
7. definierten Reset- und Cleanup-Pfad (§15);
8. Evidenzerfassung unter dem bestehenden Evidence-Modell (§19).

**Keine dieser Bedingungen ist erfüllt.** Das Lab bleibt `not provisioned`.

## 24. Failure and Unknown State

- Unbekannte Environment-Identität, unbekanntes Profil oder unbekannte Zielbindung → **fail-closed**, kein evidenztragender Lauf.
- Unbekannter Reset- oder Cleanup-Ausgang → `unknown`, blockiert den nächsten Lauf.
- Unbekannte Test-Double-Fidelity → Ergebnis `inconclusive`.
- Unbekannte Zeitbasis → Ergebnis trägt eine ausdrückliche Clock-Limitation.

```text
unknown ≠ safe · unknown ≠ clean · unknown ≠ passed
```

## 25. Security Invariants

1. Das Lab ist `non-production` und erlangt keine Produktionsreichweite.
2. Ein Lab-Ziel ist nie ein Produktionsziel; ein Ziel außerhalb des `permitted target set` ist ein Negativfall.
3. Ein Lab-Credential ist nie ein Produktions-Credential; kein Credential-Wert erscheint in Szenarien, Protokollen oder Evidenz.
4. Eine Lab-Rolle ist keine CoreOps-Autorität und kein Laufzeitmodul.
5. Ein Lab-Profil ist keine CoreOps-Betriebsautorität.
6. Laborzugang ist keine Ausführungsberechtigung.
7. Lab-Evidenz wird unter dem bestehenden Evidence-Modell geführt; es entsteht keine zweite Evidence-Autorität.
8. Ein Test Double ohne Fidelity-Deklaration erzeugt kein `passed`.
9. Ein Negativszenario ohne beobachtetes Ausbleiben der Wirkung erzeugt kein `passed`.
10. Ein Lab-Ergebnis belegt weder Produktionsreife noch Sicherheit, Recovery-Fähigkeit, Offline-Betrieb noch Herstellersupport.
11. Ein unbekannter Reset- oder Cleanup-Ausgang blockiert fail-closed den nächsten evidenztragenden Lauf.
12. Lab-Artefakte im Repository bleiben public-neutral: keine realen privaten Hostnamen, Domainnamen, Netzbereiche, Organisationsnamen oder personenbezogenen Daten.

## 26. Threat References

Referenziert, nicht dupliziert: THR-001, THR-002, THR-019, THR-021, THR-022, THR-023, THR-028, THR-029, THR-031, THR-032, THR-034, THR-036 und THR-038 aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md). Dieses Modell benennt Laborsicherheits- und Szenarioanforderungen und erhebt **keinen** Verifikationsanspruch.

## 27. Technology Boundary

Container-, Virtualisierungs-, Orchestrierungs-, Netzwerkemulations-, Simulator-, Mock-, CI-, Test-Management-, Browser-Automation-, Load-Test- und Scanner-Technologie: **nicht ausgewählt**. Die Auswahl gehört in spätere Implementierungs- und ADR-Arbeit.

## 28. Compatibility

Bindet an, ohne zu ändern: [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md), [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](../architecture/INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md), [RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md), [COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md), [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md), [RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md), [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](../security/DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md), [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) und [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md).

## 29. CDS Boundary

Der CDS-Pilot ist **nicht aktiviert**. Es wird keine CDS-Testevidenz importiert, referenziert oder verwendet. Für die UX-Semantik aus CO-WP-027 gilt ausdrücklich:

```text
CoreOps test strategy ≠ CDS consumer evidence automatically
CDS Candidate evidence ≠ CoreOps validation
```

Eine spätere CDS-Pilot-Evidenz erfordert eine **eigene, explizite Autorisierung** sowie einen Re-Check der dann aktuellen CDS-Reife und Consumer-Evidenz. Es werden keine CDS-Dateien, -Tokens oder -Abhängigkeiten eingeführt.

## 30. Open Questions

- Ob `controlled real test target` überhaupt zugelassen wird oder ob das Lab dauerhaft auf Test Doubles beschränkt bleibt — Entscheidung gehört zu späterer Design- und ADR-Arbeit.
- Wie Zeit- und Freshness-Simulation formalisiert wird, ohne die Clock-Unsicherheitssemantik des Offline-Modells zu verdoppeln.
- Wie Lab-Evidenz von Betriebsevidenz dauerhaft unterscheidbar bleibt, wenn beide unter demselben Evidence-Modell geführt werden.

## 31. Next Decision

Nova Review dieses Dokuments gemeinsam mit der [Teststrategie](FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) und der [Fixture-Governance](SYNTHETIC_FIXTURE_AND_TEST_DATA_GOVERNANCE.md). **Kein** Labor wird durch dieses Dokument bereitgestellt, konfiguriert, gestartet oder verbunden.
