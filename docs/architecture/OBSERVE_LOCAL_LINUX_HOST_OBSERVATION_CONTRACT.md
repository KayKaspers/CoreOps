# CoreOps – Observe Local Linux Host Observation Contract

> Document Status: Implemented; Nova Review abgeschlossen (Initial `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`**); `completed-go-with-notes`
> Human-Maintainer-Repository-Integration: PENDING — kein Commit, kein Push, keine Remote-Integration
> Contract Status: Observe slice contract (docs-only, technologieunabhängig)
> Phase: `Observe` — betreten und mit Grenze autorisiert (`HM-O1` `APPROVED WITH BOUNDARY`)
> Value Slice: Local Linux Host Identity & Basic System Observation — SELECTED (`HM-O2` `APPROVED`)
> Work Package: `CO-WP-032`, primärer Typ `docs-only` — autorisiert durch `HM-O3` `APPROVED`; Ausführungsgrenze `HM-O4` `APPROVED WITH EXACT BOUNDARY` (docs-only)
> Successor Work Package: none created, none reserved
> Implementation Status: Not implemented
> Collection Mechanism (`P-3`): Not selected
> Language / Runtime: Not selected
> Target Authorization (`P-1`): Not authorized
> No-Mutation Evidence (`P-2`): Not satisfied
> Tests implemented / executed: None / None
> Real Observation Performed: None
> Support Status: `not-supported` · Capability Implementation Status: `not-implemented`
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)

## 1. Status

Technologieunabhängiger **Mindest-Read-only-Vertrag** für den ersten Observe-Wertslice *Local Linux Host Identity & Basic System Observation*. Der Vertrag **spezialisiert** bestehende CoreOps-Modelle; er etabliert **kein** Parallelmodell. Autoritativ bleiben:

- [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) — Observed-/Effective-State-Semantik
- [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) — Feldidentität, Provenance, Freshness, Trust, Validation
- [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) — Autoritätsklassen
- [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) — Raw/Normalized/Derived, Canonical Fields, Freshness/Staleness
- [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) — Evidence References
- [COREOPS_MODULE_CATALOG.md](COREOPS_MODULE_CATALOG.md) — `MOD-OBS-001` / `MOD-INV-001` / `MOD-STA-001` / `MOD-EVD-001`
- [INITIAL_SUPPORT_BOUNDARY.md](../integrations/INITIAL_SUPPORT_BOUNDARY.md) — Observe-Zielumfang und Sicherheitsgrenzen

Dieser Vertrag **implementiert nichts**, wählt **keine** Technologie aus und beschreibt **keinen** durchgeführten Zugriff.

## 2. Purpose

Festlegen, **welche Fakten** ein künftiger lokaler Linux-Host-Beobachtungspfad überhaupt behaupten darf, **mit welcher Herkunft**, **mit welcher Zeitbindung** und **mit welchen ausdrücklichen Nicht-Implikationen** — bevor irgendein Erhebungsmechanismus, eine Sprache oder ein Zielzugriff gewählt oder autorisiert wird. Ziel ist, dass ein späteres Ergebnis nicht mehr behauptet, als tatsächlich beobachtet wurde, und dass Abwesenheit, Fehler, Berechtigungsverweigerung und Nichtunterstützung **unterscheidbar** bleiben.

## 3. Scope

Beobachtungsidentität · Beobachtungssubjekt · `target_id`-Semantik · Hostname · Betriebssystem-/Distributionsidentität · Betriebssystem-Release · Kernel · Architektur · `observed_at` · `received_at` · Provenance · Quellidentität · Freshness · Raw-Wert · normalisierte CoreOps-Repräsentation · `observation_outcome` · Absenz-/Fehlersemantik · Nicht-Implikationen.

## 4. Non-Goals

- **Kein** Erhebungsmechanismus, **kein** Transport, **kein** Agent/Agentless-Entscheid (`P-3` bleibt `NOT SELECTED`).
- **Keine** Sprache, **kein** Runtime, **kein** Schemaformat, **kein** Serialisierungsformat, **keine** Storage-Technologie.
- **Keine** Health-Semantik, **kein** Score, **kein** Aggregat, **kein** Statusrollup.
- **Keine** Discovery über den lokalen Host hinaus, **kein** Netzwerkziel, **keine** Credentials, **keine** Secrets, **keine** privilegierte Ausführung.
- **Keine** Managed-Resource-Registrierung, **keine** Enrollment-, Trust- oder Ownership-Aussage.
- **Kein** ADR, **keine** Decision-/Risk-ID, **keine** Änderung an Capability-, Support- oder Kompatibilitätsstatus.

## 5. Authority Boundary

```text
Observe entered               != target access authorized
value slice selected          != productive code authorized
contract defined              != implementation authorized
contract defined              != observation performed
productive code authorization != target authorization
implementation authorization  != target authorization
source code present           != target access granted
OBSERVED                      != AUTHORIZED
EVIDENCE                      != AUTHORITY
VERIFY                        != APPROVE
DETECTED                      != PERMITTED
PERMITTED                     != EXECUTED
```

Die Trennung ist zweistufig: **Gate A** entscheidet, ob produktiver Anwendungscode entstehen darf; **Gate B** entscheidet, ob ein reales Ziel berührt werden darf. `P-1` und die Erfüllung von `P-2` gehören zu Gate B — sie sind **keine** Voraussetzung dafür, diesen Vertrag zu implementieren. Die vollständige Gate-Definition steht in [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §4 und §16.

Dieser Vertrag erteilt **keinen** Zielzugriff. Ein realer Zugriff auf ein Linux-Ziel benötigt die separate Human-Maintainer-Zielautorisierung `P-1` (**derzeit `NOT AUTHORIZED`**) und einen Execution-Authorization-Pfad gemäß [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md). Der Vertrag ist außerdem **keine** Support-Aussage: die vier referenzierten Capabilities bleiben `not-implemented` / `not-supported`.

## 6. Capability Basis

Der Slice referenziert **ausschließlich bestehende** Capabilities der [FOUNDATION_CAPABILITY_MATRIX.md](FOUNDATION_CAPABILITY_MATRIX.md); **keine** Zeile wird erzeugt, umbenannt, hochgestuft oder in Implementierungs-, Support-, Evidence- oder Kompatibilitätsstatus verändert.

| Capability | Titel | Rolle in diesem Slice | Status (unverändert) |
| ---------- | ----- | --------------------- | -------------------- |
| `CAP-DISCOVERY-004` | Linux Discovery (read-only) | Erhebung der Basisfakten des lokalen Linux-Hosts | `target-observe` · `not-implemented` · `not-supported` · Transport offen · Agent/Agentless offen |
| `CAP-INVENTORY-001` | Asset Inventory (read-only) | Aufnahme der beobachteten Fakten als **beobachtetes** Subjekt — **nicht** als managed resource | `target-observe` · `not-implemented` · `not-supported` |
| `CAP-INVENTORY-004` | Field Provenance | Feldbasierte Herkunft/Vertrauen/Status je Beobachtungsfeld | `target-observe` · `not-implemented` · `not-supported` |
| `CAP-MONITORING-007` | Telemetry Normalization | Abbildung Raw auf kanonische CoreOps-Repräsentation; Rohdaten bleiben verfügbar | `target-observe` · `not-implemented` · `not-supported` |

```text
capability selected for a slice != capability implemented
capability referenced           != capability verified
capability verified             != integration supported
```

## 7. Beobachtungsidentität (Observation Identity)

Eine **Observation** ist ein einzelner, zeitgebundener, read-only Erhebungsvorgang gegen genau **ein** Beobachtungssubjekt, mit genau **einer** Quelle je Feld und genau **einem** Ergebnis-Envelope.

Mindestbestandteile der Observation Identity (konzeptionell, **kein** Schema):

- `observation_id` — stabile, nicht wiederverwendbare Kennung des Erhebungsvorgangs
- `observation_subject_ref` — Referenz auf das Beobachtungssubjekt (§8)
- `collection_attempt_ref` — Kennung des konkreten Erhebungsversuchs (ein Subjekt kann mehrere Versuche haben)
- `contract_ref` — Referenz auf diesen Vertrag und seine Revision
- `observation_outcome` — Envelope-Ergebnis (§10)
- `observed_at` / `received_at` — Zeitmodell (§11)
- `provenance_ref` — Provenance-Wurzel des Vorgangs (§13)
- `evidence_ref` — optionale Evidence Reference gemäß Evidence-Modell §9

```text
observation_id      != subject identity
observation attempt != observation result
observation result  != authoritative CoreOps state
one observation     != a series
```

## 8. Beobachtungssubjekt und `target_id`

### 8.1 Subjektklasse

Das Beobachtungssubjekt dieses Slices ist genau: **ein lokaler Linux-Host, auf dem der Erhebungspfad selbst läuft**. Kein Remote-Host, kein Netzwerkbereich, kein Cluster, kein Container-Verbund, kein Gerät.

### 8.2 `target_id`

`target_id` ist die **Identität des beobachteten Ziels innerhalb der Beobachtungsdomäne** (`MOD-OBS-001`). Sie ist eine Beobachtungs-Referenzidentität, sonst nichts.

```text
target_id                  =  observed target identity
observed target identity   != managed-resource identity
DISCOVERED                 != MANAGED
DISCOVERED                 != TRUSTED
target_id present          != ownership
target_id present          != management authority
target_id present          != authorized target
target_id present          != enrollment
target_id stable           != identity verified
```

`target_id` darf **nicht** stillschweigend erzeugen: eine Managed-Inventory-Identität (`MOD-INV-001` bleibt allein autoritativ für registrierte Ressourcenidentität und Management-Status), eine Trust-Aussage (Trust Status folgt Field Provenance §10), eine Autorisierungsidentität (`MOD-POL-001` und die Execution-Authorization-Policy bleiben allein autoritativ) oder eine Ownership-/Verwaltungsbehauptung.

### 8.3 Identitätsableitung

Die konkrete Ableitungsregel für `target_id` ist **offen** und an `P-3` gebunden. Verbindlich ist nur:

- `target_id` ist innerhalb der Beobachtungsdomäne stabil und nicht wiederverwendbar zu behandeln ([Repository Governance Standard](../governance/REPOSITORY_GOVERNANCE_STANDARD.md) §13).
- `hostname` ist **kein** `target_id`. Ein Hostname ist ein **beobachtetes Feld**: veränderlich, nicht eindeutig und als Identität nicht tragfähig.
- Eine Identitätskollision oder -unklarheit bleibt als `conflicted` sichtbar und wird **nicht** still aufgelöst (SoT-Konfliktmodell, `CO-WP-011`).

```text
hostname       != identity
identity guess != observed identity
```

## 9. Feld- und Semantikmatrix (normativ)

Kanonische Feldidentität folgt [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) §5 und [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) §18. Feldnamen sind **kanonisch englisch** ([Language Standard](../governance/COREOPS_LANGUAGE_STANDARD.md) §6); es entsteht **keine** vollständige Feld-ID-Liste.

`R` = Raw · `N` = Normalized · `M` = Metadata (weder Raw noch Normalized, sondern Beobachtungsmetadatum).

| Feld / Konzept | Zweck | Raw/Norm. | Provenance-Anforderung | Required / Conditional | Absenzsemantik | Ausdrückliche Nicht-Implikation |
| -------------- | ----- | --------- | ---------------------- | ---------------------- | -------------- | ------------------------------- |
| `target_id` | Beobachtungs-Zielidentität | M | Ableitungsregel und Quelle referenziert | **required** | Ohne `target_id` ist der Envelope ungültig; die Observation wird verworfen, nicht anonym gespeichert | nicht managed identity · nicht trusted · nicht authorized · nicht ownership |
| `hostname` | beobachteter Hostname des lokalen Hosts | R + N | Quelle, Erhebungsmechanismusklasse, Zeit | **required** | `absent` mit `field_observation_outcome`; **nie** Leerstring, **nie** Platzhalter | nicht Identität · nicht FQDN · nicht DNS-Auflösbarkeit · nicht Eindeutigkeit |
| `os_identity` | Betriebssystem-/Distributionsidentität | R + N | Quelle, Quellfeldbezug, Zeit | **required** | `absent`; keine Ableitung aus anderen Feldern | nicht Supportstatus · nicht Kompatibilität · nicht Lifecycle-Status · nicht Herstellerbeziehung |
| `os_release` | Betriebssystem-Release/Version | R + N | Quelle, Quellfeldbezug, Zeit | **required** | `absent`; **kein** Rückschluss aus `os_identity` | nicht Patchstand · nicht Aktualität · nicht Schwachstellenfreiheit · nicht Update-Notwendigkeit |
| `kernel_release` | Kernel-Release-Kennung | R + N | Quelle, Erhebungsmechanismusklasse, Zeit | **required** | `absent`; **kein** Rückschluss aus `os_release` | nicht Betriebssystemversion · laufender nicht gleich installierter Kernel · nicht Sicherheitsstand |
| `architecture` | Hardware-/Instruktionsarchitektur | R + N | Quelle, Zeit | **required** | `absent`; **kein** Default | nicht CPU-Modell · nicht Virtualisierungsstatus · nicht Kompatibilitätszusage |
| `raw_observed_value` | quellnaher Wert je Feld | R | Quelle unverändert erhalten | **required**, wo ein Feld beobachtet wurde | Ohne Raw-Wert existiert keine Beobachtung des Feldes | nicht vertrauenswürdig · nicht autoritativ · nicht validiert |
| `normalized_value` | kanonische CoreOps-Repräsentation | N | Transformationsreferenz und Raw-Referenz | **conditional** — nur bei erfolgreicher Normalisierung | `normalization-failed`; Raw bleibt erhalten und sichtbar | nicht verlustfrei · nicht validiert · nicht semantisch geprüft |
| `source_identity` | Identität der konkreten Quelle je Feld | M | selbst Bestandteil der Provenance | **required** je beobachtetem Feld | Ohne Quellidentität ist das Feld `provenance-invalid` und **nicht** verwendbar | nicht Trust · nicht Autorität · nicht Validierung |
| `source_class` | Quellklasse gemäß Field Provenance §7 | M | — | **required** je beobachtetem Feld | wie `source_identity` | nicht Trust-Zuweisung |
| `collection_mechanism_class` | Klasse des Erhebungswegs (**nicht** das Werkzeug) | M | — | **required** je beobachtetem Feld | wie `source_identity` | nicht Technologieauswahl · nicht `P-3`-Entscheidung |
| `observed_at` | Zeitpunkt, dem die Beobachtung zugeordnet ist | M | Zeitquelle und Unsicherheit benannt | **required** | `unknown`; **nie** ersetzt durch `received_at` | nicht Ingestionszeit · nicht Zielzustand jetzt · nicht vertrauenswürdiger Zeitstempel |
| `received_at` | Zeitpunkt der Entgegennahme/Ingestion | M | — | **conditional** — nur wo Erhebung und Entgegennahme materiell getrennt sind | Feld entfällt als `not-applicable`, **nicht** `unknown` | nicht Beobachtungszeit |
| `freshness` | Aktualitätszustand je Feld | M | leitet sich aus `observed_at` und Kontext ab | **required** je beobachtetem Feld | `unknown` | nicht healthy · nicht trusted · nicht authorized · `stale` nicht unavailable |
| `field_observation_outcome` | Ergebnis je Feld — Feld-Vokabular §10.3 (neun Werte, einschließlich `provenance-invalid`) | M | — | **required** je Feld | — | nicht Envelope-Ergebnis · nicht Emissions-Disposition · nicht Zielzustand |
| `observation_outcome` | Ergebnis des Erhebungsvorgangs — Envelope-Vokabular §10.2 (acht Werte) | M | — | **required**, außer wo §10.5 R1 oder R5 greift; dann wird **keiner** behauptet | — | nicht Feld-Ergebnis · nicht Health · nicht Support · nicht Autorisierung · `provenance-invalid` ist **kein** Wert dieses Vokabulars |
| Emissions-Disposition | ob und wie der Datensatz ausgegeben wird — §10.4 (`emitted` · `field-withheld` · `record-discarded`) | M | — | **required** | — | nicht `observation_outcome` · `record-discarded` nicht Fehlschlag · nicht Zielzustand |
| `provenance_ref` | Verweis auf den Provenance-Satz | M | — | **required** | Ohne Provenance ist die Observation nicht verwertbar (fail-closed) | Provenance nicht Trust |
| `evidence_ref` | Evidence Reference gemäß Evidence-Modell §9 | M | — | **conditional** — nur wo eine Evidence Reference tatsächlich erzeugt wurde | Feld entfällt; **keine** fiktive Referenz | Reference nicht Artefakt · nicht Zugriffsberechtigung · nicht Sufficiency |

**Verbindliche Zusatzregeln zur Matrix**

```text
missing data      != healthy
absent field      != default value
absent field      != zero
absent field      != empty string
unsupported       != failed
observation error != target unhealthy
permission denied != target unavailable
```

Es wird **keine** Health-Semantik definiert. Der Slice erzeugt **kein** Feld, das eine Aussage über Gesundheit, Zustandsgüte, Compliance, Risiko oder Handlungsbedarf trifft.

## 10. Observation-Outcome-Modell

Der Slice verwendet ausdrücklich `observation_outcome` — **nicht** einen generischen Ziel-`status`. Ein `status` am Ziel würde eine Aussage über das Ziel behaupten; `observation_outcome` behauptet ausschließlich eine Aussage über den **Erhebungsvorgang**.

### 10.1 Drei getrennte Vokabulare

Der Slice führt **drei** ausdrücklich getrennte Vokabulare. Kein Wert wandert zwischen ihnen; keines ist eine Obermenge eines anderen.

| Vokabular | Ebene | Gegenstand | Werte |
| --------- | ----- | ---------- | ----- |
| `observation_outcome` | **Envelope** | Ergebnis des Erhebungsvorgangs als Ganzes | §10.2 — genau **acht** Werte |
| `field_observation_outcome` | **Feld** | Ergebnis je beobachtetem Feld | §10.3 — genau **neun** Werte |
| Emissions-Disposition | **Verarbeitung** | ob und wie das Ergebnis überhaupt ausgegeben wird | §10.4 — genau **drei** Werte |

```text
envelope outcome        != per-field outcome
per-field outcome       != emission disposition
emission disposition    != observation outcome
record discarded        != observation outcome
aggregated outcome must not erase a per-field cause
partial                 != failed
partial                 != complete
```

Die Zweistufigkeit von Envelope und Feld erhält die Unterscheidbarkeit: `partial` darf keine der feldweisen Ursachen verschlucken. Die Emissions-Disposition ist davon nochmals getrennt, damit ein **nicht ausgegebener** Datensatz nicht als Beobachtungsergebnis missverstanden wird.

### 10.2 `observation_outcome` — Envelope-Vokabular (bounded)

Genau **acht** Werte; **keine** weiteren. Dieses Set ist gegenüber der Erstfassung **unverändert**.

| Outcome | Bedeutung | Ausdrücklich **nicht** |
| ------- | --------- | ---------------------- |
| `success` | Quelle erreichbar, gelesen, Wert erhalten und — soweit gefordert — normalisiert | nicht Wert korrekt · nicht Ziel gesund · nicht Wert vertrauenswürdig |
| `partial` | mindestens ein Feld `success`, mindestens ein Feld nicht | nicht Fehlschlag · nicht Vollständigkeit · nicht im Wesentlichen erfolgreich |
| `source-absent` | die erwartete Quelle existiert auf diesem Host **nicht** | nicht Ziel nicht vorhanden · nicht unsupported · nicht Fehler |
| `source-unavailable` | Quelle existiert, war aber im Erhebungsversuch nicht lesbar (z. B. transient) | nicht permission denied · nicht absent · nicht Ziel offline |
| `permission-denied` | Lesezugriff wurde durch die Zielumgebung verweigert | nicht unavailable · nicht absent · nicht Ziel unerreichbar · nicht Angriff |
| `unsupported` | Quelle oder Plattform liegt außerhalb des vertraglich unterstützten Umfangs | nicht failed · nicht absent · nicht Defekt |
| `source-malformed` | Quelle gelesen, Inhalt aber nicht vertragskonform interpretierbar | nicht normalization-failed · nicht Ziel fehlerhaft |
| `normalization-failed` | Raw-Wert liegt vor, die Abbildung auf die kanonische Repräsentation schlug fehl | nicht source-malformed · nicht Wert ungültig · Raw **bleibt erhalten** |

**Nicht** Bestandteil des Sets und ausdrücklich verworfen: `unknown` (Unwissen ist Absenz **oder** einer der benannten Fälle, kein eigener Ausgang), `failed` (würde sechs unterscheidbare Ursachen kollabieren), `healthy` / `unhealthy` / `degraded` (Zielzustands-Semantik, die dieser Vertrag nicht erzeugt), `timeout` (ist `source-unavailable` mit Provenance-Detail), `error` (nicht unterscheidungsfähig).

```text
permission-denied  != unsupported
unsupported        != source-unavailable
source-unavailable != source-absent
none of these      != target unhealthy
none of these      != failed
provenance-invalid is NOT an observation_outcome value
```

**`provenance-invalid` ist ausdrücklich kein neunter `observation_outcome`-Wert.** Es ist ein reiner Feldwert (§10.3) und wird ohne eine eigene, ausdrückliche Human-Maintainer-Autorisierung **nicht** in das Envelope-Vokabular übernommen.

### 10.3 `field_observation_outcome` — Feld-Vokabular (bounded)

Genau **neun** Werte: die acht Erhebungsausgänge aus §10.2 **plus** `provenance-invalid`. `partial` ist auf Feldebene zulässig nur dort, wo ein Feld selbst aus mehreren Quellteilen zusammengesetzt wird; im vorliegenden Slice tritt dieser Fall nicht auf und der Wert bleibt für Feldebene ungenutzt.

| Wert | Ebene | Bedeutung |
| ---- | ----- | --------- |
| `success` · `partial` · `source-absent` · `source-unavailable` · `permission-denied` · `unsupported` · `source-malformed` · `normalization-failed` | Feld **und** Envelope | wie §10.2, bezogen auf genau ein Feld |
| **`provenance-invalid`** | **nur Feld** | Ein Wert wurde erhoben, aber die zugehörige Provenance ist unvollständig, widersprüchlich oder nicht auflösbar (fehlende `source_identity`, fehlende `collection_mechanism_class`, nicht auflösbare Raw-Referenz). Das Feld ist damit **nicht verwendbar** und wird **nicht** normalisiert (fail-closed, §13). |

Nicht-Implikationen von `provenance-invalid`:

```text
provenance invalid != target unhealthy
provenance invalid != source unavailable
provenance invalid != source absent
provenance invalid != permission denied
provenance invalid != value wrong
provenance invalid != source malicious
provenance invalid != an observation_outcome value
```

`provenance-invalid` ist eine Aussage über die **Verwertbarkeit des Feldsatzes**, nicht über das Ziel und nicht über die Quelle.

### 10.4 Emissions-Disposition (bounded)

Getrenntes Vokabular; beschreibt **Verarbeitung und Ausgabe**, nicht das Beobachtungsergebnis. Genau **drei** Werte:

| Wert | Bedeutung |
| ---- | --------- |
| `emitted` | Der Beobachtungsdatensatz wird ausgegeben — mit allen Feldern und ihren `field_observation_outcome`-Werten |
| `field-withheld` | Ein einzelnes Feld fließt **nicht** in die normalisierte Repräsentation ein; sein Raw-Wert, seine Provenance-Angaben und sein `field_observation_outcome` bleiben im Datensatz sichtbar |
| `record-discarded` | Der Datensatz wird als Beobachtung **nicht** ausgegeben, weil er nicht provenance-tragfähig ist. Es wird **kein** `observation_outcome` behauptet. Ein Provenance-Defekt-Record bleibt für Audit erhalten |

```text
record discarded  != observation outcome
record discarded  != failed observation
record discarded  != target unhealthy
field withheld    != field absent
emission disposition != observation outcome
```

### 10.5 Deterministische Zusammensetzung

Die folgenden Regeln sind verbindlich und in dieser Reihenfolge auszuwerten:

| # | Bedingung | `observation_outcome` | Emissions-Disposition |
| - | --------- | --------------------- | --------------------- |
| **R1** | `target_id` fehlt **oder** seine Provenance ist ungültig | **keiner** — es wird keiner behauptet | `record-discarded` (§8.2, §9) |
| **R2** | alle Felder `success` | `success` | `emitted` |
| **R3** | mindestens ein Feld `success` **und** mindestens ein Feld ein anderer Wert (einschließlich `provenance-invalid`) | `partial` | `emitted`; jedes `provenance-invalid`-Feld zusätzlich `field-withheld` |
| **R4** | kein Feld `success`, aber alle nicht-`provenance-invalid`-Felder tragen **denselben** Erhebungsausgang X | X | `emitted` |
| **R5** | **alle** sonst beobachteten Felder sind `provenance-invalid` | **keiner** — es wird keiner behauptet | `record-discarded`; Provenance-Defekt-Record bleibt erhalten |

Zu den beiden von R3 und R5 abgedeckten Fällen ausdrücklich:

- **Einige Felder gültig, einige `provenance-invalid`** → R3: Envelope `partial`; jede feldweise Ursache bleibt einzeln erhalten und sichtbar; die betroffenen Felder werden `field-withheld` und **nicht** normalisiert; der Datensatz wird ausgegeben.
- **Alle sonst beobachteten Felder `provenance-invalid`** → R5: Der Erhebungsvorgang hat nichts Provenance-Tragfähiges erzeugt; es wird **kein** Envelope-Wert behauptet und der Datensatz wird **nicht** als Beobachtung ausgegeben. Das ist eine Emissions-Disposition, **kein** `observation_outcome`.

`provenance-invalid`-Felder liefern **niemals** einen Envelope-Wert; sie können lediglich R3 auslösen oder — wenn sie alle Felder betreffen — R5.

### 10.6 Verhältnis zu bestehenden Vokabularen

`observation_outcome` ist **kein** globales CoreOps-Statusmodell, **kein** Effective-State-Wert und **kein** Testergebnis. Effective-State-Werte (`determined` bis `not-applicable`) bleiben beim [State-Modell](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) §8; Testergebnisse (`not run` bis `not applicable`) bleiben bei der [Teststrategie](../testing/FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §11. **Kein** Wert wird zwischen diesen Vokabularen übertragen.

## 11. Zeitmodell

```text
observed_at       != received_at
observed_at       != current target state
recently received != recently observed
recent timestamp  != trusted timestamp
```

- **`observed_at`** ist der Zeitpunkt, dem die Beobachtung zugeordnet ist. Er ist **required**. Ist er nicht ermittelbar, lautet der Wert `unknown` — er wird **nicht** durch `received_at`, nicht durch „jetzt" und nicht durch einen Default ersetzt.
- **`received_at`** wird **nur** dort geführt, wo Erhebung und Entgegennahme materiell getrennt sind (etwa bei späterer Ingestion eines zuvor erzeugten Ergebnisses). Wo sie es nicht sind, ist das Feld `not-applicable` — **nicht** `unknown` und **nicht** eine Kopie von `observed_at`.
- Die verwendete Zeitquelle und ihre Unsicherheit sind Bestandteil der Provenance. Uhrenunsicherheit wird **nicht** wegnormalisiert (Bezug THR-027).
- **Kein Zeitstempel wird fabriziert.** Ein nicht ermittelbarer Zeitpunkt bleibt nicht ermittelbar.

## 12. Raw- und Normalized-Modell

Der Slice unterscheidet strikt gemäß [Telemetry-Modell](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) §10–§12:

- **Raw observed value** — der quellnahe Wert mit ursprünglicher Semantik, unverändert, mit Quelle und Provenance.
- **Normalized CoreOps representation** — die Abbildung auf ein kanonisches CoreOps-Feld mit kanonischer Semantik.
- **Derived** — im Slice **nicht vorgesehen**. Es wird kein Wert aus anderen Werten errechnet.

Normalisierung **muss**: Provenance erhalten · Absenz erhalten · Unsicherheit erhalten · die Raw-Referenz erhalten · die Transformationsreferenz führen.

Normalisierung **darf nicht**: fehlende Fakten erzeugen · eine Parser-Annahme in beobachtete Evidenz verwandeln · Raw ersetzen oder verwerfen · aus einem Feld ein anderes ableiten · bei Unklarheit einen Default setzen.

```text
raw                   != automatically trustworthy
raw                   != authoritative state
normalized            != lossless
normalized            != validated
parser assumption     != observed evidence
normalization success != semantic validation
absent                != normalized to a default
```

Schlägt die Normalisierung fehl, ist der feldbezogene Ausgang `normalization-failed`; der Raw-Wert **bleibt erhalten, referenzierbar und sichtbar**.

## 13. Provenance

Für jede materielle Beobachtung muss die Provenance späteres Nachvollziehen erlauben zu: **Quelle** · **Erhebungsmechanismusklasse** · **Beobachtungszeit** · **Normalisierung/Transformation** · **Freshness** · **Evidenzherkunft**.

Konzeptioneller Mindestsatz je beobachtetem Feld (Spezialisierung von [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) §6, **kein** neues Schema):

`subject reference` · `canonical field identity` · `source_identity` · `source_class` · `source_reference` · `collection_mechanism_class` · `observed_at` · `received_at` (conditional) · `freshness` · `raw value reference` · `transformation reference` (conditional) · `field_observation_outcome` · `validation status` · `conflict state` · `audit reference`.

```text
provenance present         != trusted
provenance present         != validated
provenance complete        != value correct
parsed                     != validated provenance
collection mechanism class != selected technology
```

Ein Feld ohne verwertbare Provenance ist **nicht** verwendbar: sein `field_observation_outcome` lautet `provenance-invalid` (Feld-Vokabular §10.3 — **kein** `observation_outcome`-Wert), seine Emissions-Disposition lautet `field-withheld` (§10.4), und es fließt **nicht** in eine normalisierte Repräsentation ein (fail-closed). Betrifft der Provenance-Defekt `target_id` oder alle sonst beobachteten Felder, greifen §10.5 R1 beziehungsweise R5: der Datensatz wird `record-discarded` und es wird **kein** `observation_outcome` behauptet.

`collection_mechanism_class` benennt bewusst eine **Klasse**, nicht ein Werkzeug — die Werkzeug- und Transportwahl ist `P-3` und **nicht getroffen**.

## 14. Freshness-Modell (bounded)

Es wird **keine** neue CoreOps-Freshness-Taxonomie erzeugt. Der Slice **verwendet** die bestehende Freshness-Zustandsmenge aus [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) §9 und begrenzt sie auf die für Identitäts- und Basissystemfakten materiell anwendbaren Werte:

| Wert | Anwendung im Slice |
| ---- | ------------------ |
| `current` | Beobachtung liegt innerhalb der für die Entscheidung akzeptierten Altersgrenze |
| `aging` | Beobachtung altert erkennbar, ist aber noch innerhalb der Grenze |
| `stale` | Beobachtung liegt außerhalb der Grenze und ist als solche sichtbar zu führen |
| `unknown` | `observed_at` ist `unknown` oder es ist keine Grenze gebunden |
| `not-applicable` | für dieses Feld existiert keine sinnvolle Altersbindung |

`expired` wird im Slice **nicht** verwendet: es setzt eine deklarierte Gültigkeitsdauer voraus, die dieser Vertrag nicht definiert.

```text
fresh             != healthy
fresh             != trusted
fresh             != authorized
stale             != unavailable
stale             != invalid
unknown freshness != current
```

**Konkrete Altersschwellen: `PROPOSED / UNACCEPTED`.** Das Repository liefert **keine** bindenden Werte; Field Provenance §9 stellt ausdrücklich fest, dass Freshness-Grenzen später **je Feld und Datenklasse** zu definieren sind. Als Diskussionsgrundlage — **nicht** als Norm, **nicht** als Entscheidung — gilt die Änderungsrate als tragende Größe:

| Feldgruppe | Erwartete Änderungsrate | Konsequenz für die spätere Schwellenwahl |
| ---------- | ----------------------- | ---------------------------------------- |
| `os_identity`, `architecture` | sehr selten | lange Grenze vertretbar; ändert sich praktisch nur bei Neuinstallation |
| `os_release`, `kernel_release` | mittelfristig | Grenze an Update-/Reboot-Rhythmus zu binden |
| `hostname` | administrativ jederzeit änderbar | kurze Grenze; niedrige Beständigkeitsannahme |

Zahlenwerte werden hier **bewusst nicht** gesetzt: die sinnvolle Grenze hängt vom Erhebungsintervall ab, das an `P-3` gebunden und **nicht entschieden** ist. Die Schwellenfestlegung ist eine spätere, ausdrücklich Human-Maintainer-gebundene Entscheidung.

## 15. Absenz- und Fehlersemantik

```text
missing data      != healthy
missing data      != zero
missing data      != empty
missing data      != default
observation error != target unhealthy
permission denied != target unavailable
unsupported       != failed
source absent     != target absent
read-only         != side-effect-free
```

- Abwesenheit wird **positiv** als Abwesenheit geführt (Outcome plus Provenance), nicht durch Weglassen des Feldes und nicht durch einen Ersatzwert.
- Ein Fehler des Beobachtungspfads ist eine Aussage über den **Pfad**, nicht über das **Ziel**.
- Ein verweigerter Lesezugriff sagt über Erreichbarkeit oder Zustand des Ziels **nichts** aus.
- Ein nicht unterstützter Quelltyp ist **kein** Defekt und **kein** Fehlschlag.
- Konflikte — widersprüchliche Werte aus mehreren Quellen — bleiben sichtbar und werden **nicht** still aufgelöst (SoT-Konfliktmodell).

## 16. Modul- und Autoritätszuordnung

| Belang | Autoritatives Modul | Grenze |
| ------ | ------------------- | ------ |
| Beobachteter Zustand, `target_id`, `observation_outcome` | `MOD-OBS-001` | erzeugt **keine** Management-Autorität, überschreibt **keinen** Desired State |
| Registrierte Ressourcenidentität und Management-Status | `MOD-INV-001` | wird durch eine Beobachtung **nicht** automatisch erzeugt (`discovered != managed`) |
| Persistenz des beobachteten Zustands | `MOD-STA-001` | Observed überschreibt Desired **nicht** still |
| Audit- und Evidence-Referenzen | `MOD-EVD-001` | Evidence erteilt **keine** Ausführungsautorität |
| Zielautorisierung und Execution Boundary | `MOD-POL-001` und Execution-Authorization-Policy | `P-1` **nicht erteilt** |

## 17. Contract-Level Non-Implications

```text
contract defined           != implementation authorized
contract defined           != observation performed
contract defined           != technology selected
observation succeeded      != target healthy
observation succeeded      != target trusted
observation succeeded      != target managed
observation succeeded      != target authorized
field present              != field validated
field normalized           != field correct
freshness current          != state authoritative
evidence reference present != evidence sufficient
read-only path             != side-effect-free path
slice contract             != support commitment
slice contract             != capability implementation
```

## 18. Security Invariants

Designanforderungen, **keine** implementierten Kontrollen:

1. Ein Beobachtungspfad darf **keine** Schreib-, Management- oder Ausführungsautorität erlangen — auch nicht implizit.
2. Eine Beobachtung darf **keine** Managed-Resource-Identität erzeugen.
3. Beobachtungsfehler dürfen **nicht** als Zielzustand dargestellt werden.
4. `permission-denied`, `unsupported`, `source-unavailable` und `source-absent` müssen unterscheidbar bleiben.
5. Raw-Provenance darf durch Normalisierung **nicht** entfernt werden.
6. `observed_at` und `received_at` dürfen **nicht** vermischt oder fabriziert werden.
7. Fehlende Daten dürfen **nicht** als gesund, null, leer oder Default interpretiert werden.
8. Der Pfad darf **keine** Credentials, Secrets oder privilegierten Rechte voraussetzen oder anfordern.
9. Der Pfad darf **kein** Netzwerkziel und **kein** Remote-System berühren.
10. Ohne gültige Provenance fällt der Pfad **fail-closed** aus; ein bestmöglich geratenes Ergebnis ist unzulässig.

## 19. Threat References

Bestehende Szenarien, **keine** neuen IDs: THR-012 und THR-013 (Telemetrie, stale), THR-014 (Inventar), THR-026 (Replay), THR-027 (Zeit), THR-034 (Beobachtungspfad). Es entsteht **keine** Parallel-Threat-Liste.

## 20. Technology Boundary

**Nicht ausgewählt:** Erhebungsmechanismus · Transport · Agent/Agentless · Sprache · Runtime · Schemaformat · Serialisierung · Storage · Scheduler · Zeitquelle · Parserbibliothek. `P-3` bleibt `NOT SELECTED`; die Decision-Pakete stehen in [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) und sind `PROPOSED / UNACCEPTED`.

## 21. Compatibility

Additiv und spezialisierend. **Kein** Parallelmodell zu State-, Provenance-, Telemetry-, Evidence- oder Testmodell. **Keine** Änderung an Capability Matrix, Initial Support Boundary, Decision Index, Risk Register, Foundation Scope Lock oder Release-Taxonomie. **Kein** ADR. Breaking-Change-Potenzial: gering.

## 22. Open Questions

- **Nicht abgedeckte Zusammensetzungslage (§10.5):** kein Feld `success` **und** die nicht-`provenance-invalid`-Felder tragen **verschiedene** Erhebungsausgänge. R4 greift nicht (kein gemeinsamer Ausgang), R3 greift nicht (kein `success`), R5 greift nicht (nicht alle `provenance-invalid`). Es wird hier **bewusst kein** Wert erfunden und `partial` **nicht** aufgeweicht; die Regel ist vor einer Implementierung ausdrücklich zu entscheiden. Bis dahin gilt fail-closed: es wird **kein** Envelope-Wert behauptet.
- Ableitungsregel und Stabilitätsgarantie für `target_id` (an `P-3` gebunden).
- Konkrete Freshness-Schwellen je Feldgruppe (erst nach `P-3`, dann Human-Maintainer-gebunden).
- Ob und wo `received_at` im Slice überhaupt materiell getrennt auftritt.
- Ob eine Evidence Reference je Observation verpflichtend wird oder bedingt bleibt.
- Kanonische Feld-IDs und ihre Versionierung (bewusst nicht hier festgelegt).

## 23. Next Decision

Der Nova Final Review dieses Vertrags ist erfolgt (`GO`); als Nächstes folgen die Human-Maintainer-Repository-, Staging-, Commit- und Push-Gates. **Danach** — und nur durch eigene, ausdrückliche Human-Maintainer-Entscheidungen — `P-1` (Zielautorisierung), `P-3` (Erhebungsmechanismus) und die Sprach-/Runtime-Entscheidung. Dieser Vertrag autorisiert **keine** davon.
