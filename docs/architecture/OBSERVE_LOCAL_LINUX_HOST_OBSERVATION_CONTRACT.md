# CoreOps – Observe Local Linux Host Observation Contract

> Document Status: Implemented; Nova Review abgeschlossen (Initial `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`**); `completed-go-with-notes`
> Human-Maintainer-Repository-Integration: COMPLETE — integrierter Commit `9999114200bf18baaadfb508e8464720b75e352e`; Push COMPLETE; Remote-Integration COMPLETE; `origin/main` = `9999114200bf18baaadfb508e8464720b75e352e`
> Contract Status: Observe slice contract (docs-only, technologieunabhängig)
> Phase: `Observe` — betreten und mit Grenze autorisiert (`HM-O1` `APPROVED WITH BOUNDARY`)
> Value Slice: Local Linux Host Identity & Basic System Observation — SELECTED (`HM-O2` `APPROVED`)
> Work Package: `CO-WP-032`, primärer Typ `docs-only` — autorisiert durch `HM-O3` `APPROVED`; Ausführungsgrenze `HM-O4` `APPROVED WITH EXACT BOUNDARY` (docs-only)
> Successor Work Package: none created, none reserved
> Implementation Status: Not implemented
> Collection Mechanism (`P-3`): **SELECTED** — ausschließlich als **Mechanismusklasse**, entschieden **außerhalb** dieses Vertrags ([PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5): Option A primär, Option B ausschließlich ergänzend, Option C **nicht** Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport. **Kein** konkreter Quellpfad, **keine** API, **keine** Bibliothek, **kein** Werkzeug und **kein** Erhebungsintervall ausgewählt; die `target_id`-Ableitungsregel ist dadurch **nicht** entschieden.
> Language / Runtime: **SELECTED** — ausschließlich als **Sprach-/Runtime-Klasse** für den ersten Observe-Slice: **Go**, entschieden **außerhalb** dieses Vertrags ([PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §8.4). **Keine** Go-Version, **keine** Distribution, **keine** Toolchain, **keine** API, **keine** Bibliothek, **kein** Paket-/Modulname, **kein** Quellpfad, **keine** Implementierung und **kein** breiterer Technologie-Stack ausgewählt. Dieser Vertrag bleibt **technologieunabhängig**.
> Target Authorization (`P-1`): Not authorized
> No-Mutation Evidence (`P-2`): Not satisfied
> Tests implemented / executed: None / None
> Real Observation Performed: None
> Support Status: `not-supported` · Capability Implementation Status: `not-implemented`
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)
> Nachträgliche docs-only Änderung: Regel **R6** — heterogene All-Failure-Zusammensetzung — ergänzt in §9, §10.4, §10.5, §22 und §23 (kanonische Nicht-Emission bei evidenzerhaltender Retention des erzeugten Materials). Grundlage ist die ausdrückliche Human-Maintainer-Entscheidung zur heterogenen All-Failure-Zusammensetzung; **kein** Decision-Identifier vergeben, **kein** ADR, **keine** Risk-, CCR- oder Capability-Kennung, **kein** neues und **kein** reserviertes Work Package.
> Nachträgliche docs-only Current-State-Reconciliation: Die veralteten Statusaussagen zu `P-3` und zur Sprach-/Runtime-Frage sind auf den tatsächlichen Stand gebracht (Kopf, §4, §8.3, §13, §14, §20, §22, §23). Beide Entscheidungen sind **außerhalb** dieses Vertrags durch eigene, ausdrückliche Human-Maintainer-Entscheidungen getroffen und in [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5 und §8.4 dokumentiert. Diese Reconciliation ist **ausschließlich** eine Statusangleichung: die normative Beobachtungssemantik dieses Vertrags — Feldidentitäten, Feld-/Semantikmatrix, Envelope-Semantik, das achtwertige `observation_outcome`-Vokabular, das neunwertige Feld-Vokabular, die drei Emissions-Dispositionen, **R1** bis **R6**, Raw-/Normalized-Trennung, Provenance-Anforderungen, `observed_at`/`received_at`, Freshness-Vokabular, Absenz-/Fehlersemantik, Modulautorität, Sicherheitsinvarianten und `target_id`-Semantik — bleibt **unverändert**. Es entsteht **kein** Work Package, **kein** ADR, **keine** ADR-Nummer und **keine** Decision-, Risk-, CCR-, Capability-, Support-Status-, Lessons- oder NDF-Feedback-Kennung; akzeptierte ADRs bleiben **0**.

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

- **Kein** Erhebungsmechanismus, **kein** Transport, **kein** Agent/Agentless-Entscheid — dieser Vertrag trifft **keine** solche Auswahl. Die `P-3`-**Mechanismusklasse** ist außerhalb dieses Vertrags entschieden ([PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5); ihre konkrete Realisierung — Quellpfad, API, Bibliothek, Werkzeug, Erhebungsintervall — ist es **nicht**.
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

Die konkrete Ableitungsregel für `target_id` ist **offen**. Sie ist an die konkrete `P-3`-**Realisierung** gebunden: die `P-3`-Mechanismusklasse ist zwar entschieden, die Ableitungsregel dadurch aber ausdrücklich **nicht** (`mechanism class selected != target_id derivation decided`). Verbindlich ist nur:

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
| `observation_outcome` | Ergebnis des Erhebungsvorgangs — Envelope-Vokabular §10.2 (acht Werte) | M | — | **required** für jede tatsächlich ausgegebene kanonische Observation (§7); ausgenommen sind ausschließlich §10.5 **R1**, **R5** und **R6** — dort wird **keiner** behauptet **und** der Datensatz wird **nicht** als kanonische Observation ausgegeben (`record-discarded`, §10.4) | — | nicht Feld-Ergebnis · nicht Health · nicht Support · nicht Autorisierung · `provenance-invalid` ist **kein** Wert dieses Vokabulars |
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

Die Requiredness-Ausnahme für `observation_outcome` weicht §7 **nicht** auf. Greift §10.5 R1, R5 oder R6, entsteht **keine** kanonische Observation: es wird weder ein `observation_outcome` behauptet noch ein Datensatz als Observation ausgegeben. Umgekehrt trägt jede tatsächlich ausgegebene kanonische Observation weiterhin **zwingend** ein `observation_outcome` (§7) — diese Mindestbestandteilspflicht bleibt unverändert und wird **nicht** bedingt gestellt.

```text
emitted canonical Observation   => observation_outcome present
no observation_outcome asserted => record not emitted as a canonical Observation
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
| `record-discarded` | Der Datensatz wird **nicht** als kanonische Observation ausgegeben, weil eine **zwingende Gültigkeitsbedingung** einer kanonischen Observation nicht wahrheitsgemäß erfüllbar ist. Es wird **kein** `observation_outcome` behauptet. Das bereits erzeugte Material bleibt nach den unten stehenden Erhaltungsregeln erhalten; in den Fällen R1 und R5 bleibt insbesondere der Provenance-Defekt-Record für Audit erhalten |

```text
record discarded     != observation outcome
record discarded     != failed observation
record discarded     != target unhealthy
record discarded     != target state
record discarded     != deletion
field withheld       != field absent
emission disposition != observation outcome
diagnostic retained  != canonical Observation emitted
diagnostic material  != governed audit record
```

**Tragendes Prinzip.** Eine kanonische Observation wird **nicht** ausgegeben, wenn eine **zwingende Gültigkeitsbedingung** einer kanonischen Observation nicht **wahrheitsgemäß** erfüllt werden kann. Das ist ausdrücklich **nicht** gleichbedeutend damit, dass in jedem Auslöserfall ein einzelnes §7-Feld syntaktisch unbefüllbar wäre: bei R6 sind die §7-Bestandteile technisch vorhanden — unbefüllbar ist allein eine **wahre** Envelope-Aussage innerhalb des begrenzten Acht-Werte-Vokabulars (§10.2).

**Geschlossene Auslöserliste.** `record-discarded` tritt ausschließlich in den folgenden drei, in §10.5 abschließend definierten Fällen ein:

1. **R1** — ungültige Subjektzuordnung: `target_id` fehlt oder seine Provenance ist ungültig (§8.2, §9).
2. **R5** — keine provenance-tragfähige, verwertbare Beobachtung: alle sonst beobachteten Felder sind `provenance-invalid` (§13).
3. **R6** — keine wahrheitsgemäße Envelope-Zusammensetzung innerhalb des begrenzten Acht-Werte-Vokabulars: kein Feld `success`, mindestens zwei nicht-`provenance-invalid`-Felder, und unter diesen mehr als ein unterschiedlicher Erhebungsausgang.

Die Liste ist **abschließend**. Ein weiterer Auslöser entsteht **nicht** durch Auslegung; die Erweiterung dieser Auslöserliste benötigt eine **eigene, ausdrückliche Human-Maintainer-Autorisierung**.

**Erhaltung des erzeugten Materials.** `record-discarded` ist **kein** Löschvorgang. Nach den **bestehenden** Semantiken bleiben erhalten:

- die feldweisen `field_observation_outcome`-Werte mit ihrer je eigenen Ursache, einzeln unterscheidbar (§10.3),
- die Provenance-Angaben je Feld (§13),
- der Raw-Wert dort, wo ein Raw-Wert entstanden ist (§12),
- der Kontext des Erhebungsversuchs (`collection_attempt_ref`, §7) — in den Fällen R1 und R5 einschließlich des Provenance-Defekt-Records.

Dieses Material ist `internal diagnostic` beziehungsweise `audit-relevant diagnostic` im Sinne von [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) §15. Es entsteht dadurch **keine** neue Artefaktklasse und **kein** neues Vokabular. Die Einstufung als **governed audit record** ist eine eigene Aussage des Audit-/Evidence-Modells (`MOD-EVD-001`, §16) und folgt **nicht** automatisch aus der Retention des Diagnosematerials.

Ob und wie dieses Material geführt wird, ist hier **nicht** entschieden: Es wird **keine** Storage-Technologie, **keine** Datenbank, **keine** Persistenzform, **keine** Aufbewahrungsdauer, **kein** Löschmechanismus, **keine** Logging-Implementierung und **kein** Telemetrieprotokoll ausgewählt (§20).

### 10.5 Deterministische Zusammensetzung

Die folgenden Regeln sind verbindlich und in dieser Reihenfolge auszuwerten:

| # | Bedingung | `observation_outcome` | Emissions-Disposition |
| - | --------- | --------------------- | --------------------- |
| **R1** | `target_id` fehlt **oder** seine Provenance ist ungültig | **keiner** — es wird keiner behauptet | `record-discarded` (§8.2, §9) |
| **R2** | alle Felder `success` | `success` | `emitted` |
| **R3** | mindestens ein Feld `success` **und** mindestens ein Feld ein anderer Wert (einschließlich `provenance-invalid`) | `partial` | `emitted`; jedes `provenance-invalid`-Feld zusätzlich `field-withheld` |
| **R4** | kein Feld `success`, aber alle nicht-`provenance-invalid`-Felder tragen **denselben** Erhebungsausgang X | X | `emitted` |
| **R5** | **alle** sonst beobachteten Felder sind `provenance-invalid` | **keiner** — es wird keiner behauptet | `record-discarded`; Provenance-Defekt-Record bleibt erhalten |
| **R6** | kein Feld `success`, **mindestens zwei** nicht-`provenance-invalid`-Felder, und unter diesen **mehr als ein** unterschiedlicher Erhebungsausgang | **keiner** — es wird keiner behauptet | `record-discarded`; feldweise Ursachen, Provenance, vorhandenes Raw und Erhebungskontext bleiben erhalten (§10.4) |

Zu den von R3, R5 und R6 abgedeckten Fällen ausdrücklich:

- **Einige Felder gültig, einige `provenance-invalid`** → R3: Envelope `partial`; jede feldweise Ursache bleibt einzeln erhalten und sichtbar; die betroffenen Felder werden `field-withheld` und **nicht** normalisiert; der Datensatz wird ausgegeben.
- **Alle sonst beobachteten Felder `provenance-invalid`** → R5: Der Erhebungsvorgang hat nichts Provenance-Tragfähiges erzeugt; es wird **kein** Envelope-Wert behauptet und der Datensatz wird **nicht** als Beobachtung ausgegeben. Das ist eine Emissions-Disposition, **kein** `observation_outcome`.
- **Kein Feld `success` und verschiedene Erhebungsausgänge unter mindestens zwei verwertbaren Feldern** → R6: Innerhalb des begrenzten Acht-Werte-Vokabulars (§10.2) existiert **keine** Envelope-Aussage, die wahr wäre. Es wird deshalb **kein** Envelope-Wert behauptet, **kein** synthetisches Aggregat gebildet, `partial` **nicht** aufgeweicht (der von `partial` geforderte `success`-Anteil fehlt) und **keine** Präzedenz- oder Schwereordnung zwischen den Ausgängen eingeführt — kein Ausgang „gewinnt“. Der Datensatz wird `record-discarded` und **nicht** als kanonische Observation ausgegeben; die feldweisen Ursachen bleiben einzeln unterscheidbar, Provenance und vorhandenes Raw bleiben erhalten, das erzeugte Diagnosematerial bleibt unter den bestehenden Semantiken erhalten (§10.4). Es entsteht **kein** neunter Envelope-Wert.

`provenance-invalid`-Felder liefern **niemals** einen Envelope-Wert; sie können lediglich R3 auslösen oder — wenn sie alle Felder betreffen — R5. Für R6 zählen sie **nicht** mit: R6 bewertet ausschließlich die nicht-`provenance-invalid`-Felder.

R6 ist **deterministisch**: Bedingung und Ergebnis sind vollständig aus den feldweisen Ausgängen bestimmt — ohne Ermessen, ohne Reihenfolgeabhängigkeit unter den Feldern und ohne Auswahl eines „führenden“ Ausgangs. R4 und R6 sind disjunkt und lückenlos: tragen alle nicht-`provenance-invalid`-Felder **denselben** Ausgang, greift R4; tragen sie **verschiedene**, greift R6. R4 bleibt dadurch **unverändert**.

R1, R5 und R6 bleiben unterscheidbar:

```text
R1 = invalid subject attribution       (target_id)
R5 = nothing provenance-bearing/usable (all fields provenance-invalid)
R6 = no truthful envelope composition  (heterogeneous all-failure)
```

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

`collection_mechanism_class` benennt bewusst eine **Klasse**, nicht ein Werkzeug. Die **Werkzeugwahl** ist **nicht getroffen**; außerhalb dieses Vertrags entschieden sind ausschließlich die `P-3`-**Mechanismusklasse** und der Transport (**kein** Netzwerktransport) — **kein** konkretes Werkzeug, **keine** konkrete Quelle und **keine** konkrete API ([PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §7.5). An der Provenance-Anforderung ändert das **nichts**.

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

Zahlenwerte werden hier **bewusst nicht** gesetzt: die sinnvolle Grenze hängt vom **Erhebungsintervall** ab. Die `P-3`-Mechanismusklasse ist entschieden, das Erhebungsintervall selbst ist es **nicht** (`mechanism class selected != collection cadence selected`). Die Schwellenfestlegung ist eine spätere, ausdrücklich Human-Maintainer-gebundene Entscheidung.

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

**Durch diesen Vertrag nicht ausgewählt:** Erhebungsmechanismus · Transport · Agent/Agentless · Sprache · Runtime · Schemaformat · Serialisierung · Storage · Scheduler · Zeitquelle · Parserbibliothek. Dieser Vertrag bleibt **technologieunabhängig** und trifft **keine** dieser Auswahlen.

**Stand außerhalb dieses Vertrags (Current State).** Zwei eng begrenzte Entscheidungen sind durch eigene, ausdrückliche Human-Maintainer-Entscheidungen getroffen und in [PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) dokumentiert: `P-3` ist **`SELECTED`** — ausschließlich als **Mechanismusklasse** (dort §7.5: Option A primär, Option B ausschließlich ergänzend, Option C nicht Standardweg, Option D und E für diesen Slice ausgeschlossen, **kein** Netzwerktransport) — und die **Sprach-/Runtime-Klasse** ist **`SELECTED`**: **Go** (dort §8.4), ausschließlich als Klasse für den ersten Observe-Slice. **Nicht** ausgewählt bleiben: konkreter Quellpfad · konkrete API · Bibliothek · Werkzeug · Erhebungsintervall · Go-Version · Distribution · Toolchain · Modul-/Paketlayout · Schemaformat · Serialisierung · Storage · Scheduler · Zeitquelle · Parserbibliothek · breiterer Technologie-Stack. Die übrigen dortigen Decision-Pakete sind inzwischen ebenfalls entschieden — ebenso **außerhalb** dieses Vertrags: der **Source Tree** ist **`DECIDED`** (dort §9.2; Gate-A-Punkt `A-8` erfüllt), aber **`NOT CREATED`** — **kein** Verzeichnis und **keine** Datei angelegt; und die `README`/`LICENSE`-Disposition ist **`DECIDED`** (dort §13; Gate-A-Punkt `A-10` erfüllt) mit Veröffentlichungsrechte-Modus **PUBLIC / OPEN SOURCE** und Outbound-Lizenz **Apache-2.0**. `README.md` und `LICENSE` sind inzwischen **vorhanden**; ihre Erstellung erfolgte **nicht** durch `A-10`, sondern durch eine davon getrennte, ausdrücklich autorisierte Human-Maintainer-Publikations-Realisierung. Alle diese Stände sind **außerhalb** dieses Vertrags definiert: er trifft **keine** dieser Entscheidungen, ist für sie **nicht** die Autorität und leitet aus ihnen **keine** Autorität ab.

Für diesen Vertrag folgt daraus **nichts**: er führt **keine** sprachspezifischen Typen, Schnittstellen, Schemata, Paket- oder Modulnamen, APIs, Quellpfade oder Implementierungskonventionen ein, und seine normative Beobachtungssemantik bleibt **unverändert**.

```text
mechanism class selected        != source path selected
mechanism class selected        != API selected
mechanism class selected        != collection cadence selected
mechanism class selected        != target_id derivation decided
language/runtime class selected != language version selected
language/runtime class selected != toolchain selected
language/runtime class selected != implementation selected
language/runtime class selected != broader technology stack selected
technology decided outside      != this contract changed
```

## 21. Compatibility

Additiv und spezialisierend. **Kein** Parallelmodell zu State-, Provenance-, Telemetry-, Evidence- oder Testmodell. **Keine** Änderung an Capability Matrix, Initial Support Boundary, Decision Index, Risk Register, Foundation Scope Lock oder Release-Taxonomie. **Kein** ADR. Breaking-Change-Potenzial: gering.

## 22. Open Questions

**Geschlossen — nicht mehr offen.** Die zuvor hier geführte Zusammensetzungslage (§10.5: kein Feld `success` **und** verschiedene Erhebungsausgänge unter den nicht-`provenance-invalid`-Feldern) ist **entschieden**. Sie ist durch die ausdrückliche Human-Maintainer-Entscheidung zur heterogenen All-Failure-Zusammensetzung als Regel **R6** (§10.5) aufgelöst: **kein** Envelope-Wert wird behauptet, der Datensatz wird **nicht** als kanonische Observation ausgegeben (`record-discarded`), und das erzeugte Material bleibt nach den bestehenden Semantiken erhalten (§10.4). Die Entscheidung weicht `partial` **nicht** auf, führt **keinen** neunten Envelope-Wert ein und begründet **keine** Präzedenz- oder Schwereordnung. Für sie wurde **kein** Decision-Identifier, **kein** ADR, **keine** Risk-, CCR- oder Capability-Kennung und **kein** Work Package erzeugt oder reserviert.

Weiterhin offen:

- Ableitungsregel und Stabilitätsgarantie für `target_id` — an die konkrete `P-3`-**Realisierung** gebunden; die entschiedene Mechanismusklasse entscheidet sie **nicht**.
- Konkrete Freshness-Schwellen je Feldgruppe — erst nach der konkreten `P-3`-Realisierung und dem daraus folgenden Erhebungsintervall, dann Human-Maintainer-gebunden.
- Ob und wo `received_at` im Slice überhaupt materiell getrennt auftritt.
- Ob eine Evidence Reference je Observation verpflichtend wird oder bedingt bleibt.
- Kanonische Feld-IDs und ihre Versionierung (bewusst nicht hier festgelegt).

## 23. Next Decision

**Aktueller Stand.** `CO-WP-032` ist **abgeschlossen und remote integriert**: Nova Final Review `GO`; Human-Maintainer-Integrationscommit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht; die anschließende Post-Integrations-Reconciliation steht auf `390da5cc8629dfa9cbea990c0a3c4ba4cb156e9b`. Die Repository-, Staging-, Commit- und Push-Gates für `CO-WP-032` sind damit **erledigt** und **nicht** mehr der nächste Schritt.

Der vorliegende R6-Nachtrag ist eine docs-only Änderung im Arbeitsverzeichnis. Staging, Commit, Push, Tag und jede weitere Repository-Aktion liegen unverändert **ausschließlich** beim Human Maintainer ([Repository Governance Standard](../governance/REPOSITORY_GOVERNANCE_STANDARD.md) §6). Ein künftiger Integrationsstand dieses Nachtrags wird hier **nicht** vorweggenommen und **kein** Commit-Identifier dafür vorhergesagt.

**Offen — und jeweils nur durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung:** die verbleibenden **sechs** Gate-A-Punkte ([PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md](../governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) §16.1) — Source-Tree-Entscheidung (`A-8`), Dependency-Admission (`A-9`), `NEW-8` / `README`-`LICENSE` (`A-10`), ADR-/Decision-Disposition (`A-11`), Build-/Packaging-Disposition (`A-13`) und ein ausdrücklich autorisiertes implementierungsorientiertes Work Package (`A-14`). **Nicht** mehr offen sind `P-3` (`A-6`) und die Sprach-/Runtime-Entscheidung (`A-7`): beide sind ausschließlich als **Klasse** entschieden, autorisieren **keine** Implementierung und ändern diesen Vertrag **nicht**. **Erst danach und erst mit vorhandener Implementierung** werden `P-1` (Zielautorisierung) und die Gate-B-Punkte (§16.2) überhaupt bewertbar.

```text
Gate A:            NOT PASSED — 8 von 14 erfüllt, 0 teilweise, 6 offen
                   (offen: A-8 · A-9 · A-10 · A-11 · A-13 · A-14)
Gate B:            NOT PASSED — 0 von 8
productive code:   NOT AUTHORIZED
implementation:    NOT AUTHORIZED
P-1:               NOT AUTHORIZED
P-2:               NOT SATISFIED
P-3:               SELECTED — nur Mechanismusklasse; kein Quellpfad,
                   keine API, keine Bibliothek, kein Werkzeug,
                   kein Erhebungsintervall, target_id-Ableitung
                   nicht entschieden
language/runtime:  SELECTED — nur Klasse (Go); keine Version, keine
                   Distribution, keine Toolchain, keine API, keine
                   Implementierung, kein breiterer Technologie-Stack
target access:     NOT AUTHORIZED
test execution:    NOT AUTHORIZED
accepted ADRs:     0
successor WP:      none created, none reserved
```

Dieser Vertrag autorisiert **keine** davon. Die R6-Entscheidung ist eine Semantikfestlegung und **keine** Freigabe: `R6 decided != productive code authorized`, `R6 decided != implementation authorized`.
