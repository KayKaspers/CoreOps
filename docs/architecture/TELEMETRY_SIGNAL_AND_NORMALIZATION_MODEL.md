# CoreOps – Telemetry Signal and Normalization Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation telemetry and normalization model
> Implementation Status: Not implemented
> Telemetry Protocol: Not selected
> Schema Format: Not selected
> Storage Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-019 (docs-only / telemetry, observability and normalization foundation)

## 1. Status

Technologieunabhängiges Modell für **Telemetry Signals, Signal-Identität, Raw/Normalized/Derived/Aggregated Telemetry, Metric-/Log-/Trace-/Health-Semantik, Canonical Fields und Normalization**. Companion zu [TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md](TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md) und [TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md](../security/TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md). Es implementiert **kein** OpenTelemetry/Prometheus/Grafana/Loki/Elasticsearch/SNMP/Syslog, kein Agent-Protokoll, kein Schemaformat.

## 2. Purpose

Telemetrie wird leicht mit Wahrheit verwechselt: ein Signal ist keine autoritative State-Aussage, `health signal ≠ verified health`, `missing ≠ zero` und `alert condition ≠ incident`. Dieses Modell trennt Telemetry von Event/Evidence/Command, trennt Raw/Normalized/Derived/Aggregated und bindet jede Transformation an Provenance — damit Telemetrie nie still zu autoritativem CoreOps-State oder Execution Authority wird.

## 3. Scope

Telemetry-Begriffe · Signal Classes · Signal Identity · Producers/Sources · Resource/Subject Binding · Raw/Normalized/Derived/Aggregated · Metric-/Log-/Trace-/Health-Semantik · Canonical Fields · Normalization Profiles · Mapping/Transformation · Units/Scale/Precision · Quality/Confidence · Freshness/Staleness · Sampling · Aggregation · Cardinality · State-Authority-Boundary · Event/Evidence-Boundary · Offline. Mapping/Quality-Detail (Companion 2), Trust/Privacy/Disclosure (Companion 3).

## 4. Non-Goals

- Keine OpenTelemetry-/Prometheus-/Grafana-/Loki-/Elasticsearch-/SIEM-Auswahl, kein SNMP/Syslog.
- Kein Agent-/Collector-Protokoll, kein Metric-/Log-/Trace-Schema, keine Serialisierung (JSON/Protobuf).
- Keine Label-Namen/Dashboards, keine Time-Series-DB/Storage, keine Alerting-/Rule-Engine, kein Runtime-Code.
- Keine Behauptung implementierter Telemetrie/Observability.

## 5. Concepts

Begriffe (mindestens): telemetry · telemetry signal · signal sample · signal series · metric · metric sample · log record · diagnostic record · trace · span · health signal · status signal · observation · event · audit event · evidence · raw/normalized/derived/aggregated telemetry · canonical field · source field · normalization profile · mapping · transformation · unit · scale · precision · sampling · aggregation · cardinality · quality · confidence · freshness · staleness.

**Grundregeln:**
```text
telemetry            ≠ audit event ≠ evidence ≠ command
signal received      ≠ source state currently true
raw telemetry        ≠ authoritative state
normalized telemetry ≠ validated telemetry
derived telemetry    ≠ independently observed telemetry
aggregated telemetry ≠ complete underlying population
health signal        ≠ verified system health
alert condition      ≠ incident
```

## 6. Signal Classes

Je Klasse: Purpose · Typical producer · Typical source · Subject/resource class · Workspace/scope · Signal shape · Expected frequency · Freshness expectation · Quality considerations · Cardinality considerations · Privacy/disclosure · Evidence relevance · Threat references. **Keine vollständige Signal-ID-Liste.**

| Signal-Klasse | Typ. Producer (Konzept) | Threat refs |
|---|---|---|
| resource inventory observation | MOD-INV | THR-014 |
| resource state observation | MOD-OBS/STA | THR-012, THR-013 |
| performance metric | MOD-OBS | THR-012 |
| capacity metric | MOD-OBS | THR-036 |
| availability signal | MOD-OBS | THR-013 |
| health signal | MOD-OBS | THR-012, THR-013 |
| usage signal | MOD-OBS | THR-012 |
| workflow or job telemetry | MOD-WFL | THR-028 |
| execution telemetry | MOD-EXE | THR-005, THR-029 |
| integration telemetry | MOD-ADP | THR-010, THR-011 |
| network telemetry | MOD-OBS/ADP | THR-039 |
| printer and print telemetry | MOD-ADP | THR-025 |
| container and workload telemetry | MOD-ADP | THR-034 |
| virtualisation telemetry | MOD-ADP | THR-034 |
| operating-system telemetry | MOD-ADP | THR-034 |
| security telemetry | MOD-SEC | THR-037 |
| application telemetry | MOD-ADP | THR-011 |
| topology observation | MOD-TOP | THR-015 |
| synthetic observation | MOD-OBS | THR-013 |
| diagnostic log | MOD-OBS/EVD | THR-019, THR-020 |
| trace or operation path signal | MOD-EXE/OBS | THR-005 |
| offline telemetry package | MOD-OFF | THR-024 |

## 7. Signal Identity

Ein Telemetry Signal benötigt konzeptionell mindestens: stable signal identity · signal class/type · producer · source · source trust state · resource/subject identity · workspace/environment · scope · series identity (falls anwendbar) · sample identity · source schema reference · normalization profile reference · source observation time · recording/ingestion time · sequence context · unit · scale · precision · value classification · quality state · confidence state · freshness state · provenance reference · transformation references · retention classification · disclosure classification · evidence relevance.
```text
signal identity ≠ series identity ≠ sample identity ≠ resource identity
producer        ≠ source
signal recorded ≠ signal validated
```
Keine ID-Technologie ausgewählt.

## 8. Producers and Sources

Producer (erzeugt/liefert Signal) ≠ Source (Ursprung der Messung); Source Trust State explizit (konsistent mit [Integration Trust](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Event Model](EVENT_AND_AUDIT_CORRELATION_MODEL.md)). Ein Producer ist nicht automatisch autoritative Quelle.

## 9. Resource and Subject Binding

Signale binden an Resource/Subject-Identität (kanonisch, [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md)); `adapter field ≠ canonical field`. Falsche/unklare Bindung ⇒ Quality `unknown`, nicht stille Zuordnung (Bezug THR-014).

## 10. Raw Telemetry

Quellnah erfasste Werte mit ursprünglicher Semantik. `raw ≠ automatically trustworthy`; `raw ≠ authoritative state`. Raw behält Source/Provenance.

## 11. Normalized Telemetry

Auf kanonische CoreOps-Felder/Semantik abgebildete Werte (Companion 2). `normalized ≠ lossless`; `normalized ≠ validated`; Normalisierung erhält Source-Provenance/Transformation History (§20).

## 12. Derived Telemetry

Aus anderen Signalen/Zuständen berechnete Werte. `derived ≠ independently observed`; abgeleitete Werte sind keine unabhängige zweite Quelle.

## 13. Aggregated Telemetry

Über Zeit/Ressourcen/Signalgruppen zusammengefasste Werte. `aggregated ≠ complete population`; `transformed ≠ authoritative`. Jede Transformation bleibt nachvollziehbar.

## 14. Metric Semantics

`gauge-like · counter-like · monotonic counter · non-monotonic counter · delta · cumulative · rate · ratio · percentage · duration · timestamp-like · distribution/histogram-like · unknown metric semantics`.
```text
counter    ≠ rate
cumulative ≠ delta
percentage ≠ ratio without scale definition
zero       ≠ missing ≠ unknown
missing sample ≠ zero value
unknown semantics ≠ safe for automatic aggregation
```
Keine Metric-Technologie ausgewählt.

## 15. Log and Diagnostic Semantics

`application log · system log · agent diagnostic · adapter diagnostic · security diagnostic · user-visible message · internal diagnostic · audit-relevant diagnostic`.
```text
log record        ≠ governed audit record
error log         ≠ target action failed
absence of log    ≠ action did not occur
diagnostic payload ≠ unrestricted disclosure
```
Keine Log Levels/Logformate erzwungen (Bezug THR-019, THR-020).

## 16. Trace and Span Semantics

Konzeptionell: trace/span identity · parent relationship · operation/attempt reference · producer · source · start/end boundary · reported duration · status · error reference · workspace/scope · sampling state · provenance.
```text
trace              ≠ audit history
parent relationship ≠ proven causation
span completed     ≠ operation successful
missing span       ≠ operation did not occur
```
Keine Trace-Technologie ausgewählt.

## 17. Health and Status Signals

`self-reported · externally observed · synthetic · dependency · derived · unknown health`.
```text
self-reported healthy ≠ externally verified healthy
single successful probe ≠ sustained availability
dependency healthy   ≠ service healthy
telemetry absent     ≠ service down
telemetry present    ≠ service healthy
```
Health erhält Source, Freshness, Observation Scope.

## 18. Canonical Fields

Ein Canonical Field benötigt: stable field identity · canonical name · owning module · data domain · semantic definition · value classification · expected unit/unit class · scale · precision expectation · nullable/unknown behavior · source mappings · transformation constraints · quality expectations · compatibility state · deprecation state · evidence reference.
```text
display name ≠ field identity
same label   ≠ same semantics
same unit    ≠ same scale or precision
mapped field ≠ authoritative field
```
**Keine vollständige Feld-ID-Liste.** Owning module folgt [Data Ownership](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md).

## 19. Normalization Profiles

Detail in Companion 2 §7-9. Ein Profile bindet Source-Schema/-Version an einen Target-Canonical-Field-Set mit Mapping-/Unit-/Precision-Regeln, Lossiness, Compatibility/Validation-State. `profile exists ≠ mapping validated`; `fallback mapping ≠ safe mapping`.

## 20. Mapping and Transformation

Transformation History (Companion 2 §11) erfasst source/target field, Regelreferenz, conversion/normalisation/filtering/redaction/aggregation/derivation, precision/unit change, default substitution, quality impact, lossiness, producer/processor, time boundary, provenance. `transformation history ≠ legal chain of custody`; **Transformationen entfernen Raw Provenance nicht still**.

## 21. Units, Scale and Precision

Unit Class · Source Unit · Canonical Unit · Scale · Precision · Rounding · Conversion Loss · Unknown Unit · Conflicting Unit · Unit Version/Definition Context.
```text
same numeric value ≠ same measurement
unit converted     ≠ precision preserved
rounded value      ≠ source value
unknown unit       ≠ canonical unit
conflicting unit   ≠ safe automatic conversion
```
Keine Unit Library ausgewählt.

## 22. Quality and Confidence

Detail Companion 2 §17-18. Quality (not-assessed…unknown) und Confidence (not-assessed…unknown) getrennt. `high confidence ≠ authoritative`; `validated ≠ fresh`; `normalization success ≠ semantic validation`.

## 23. Freshness and Staleness

Freshness gebunden an: signal class · source · resource · expected observation frequency · last source observation · last successful ingestion · clock uncertainty · offline state · quality state · decision/consumer context.
```text
recently ingested ≠ recently observed
recently observed ≠ current target state
stale             ≠ invalid automatically
no recent sample  ≠ zero ≠ resource unavailable
```

## 24. Sampling

`unsampled · periodic · event-triggered · adaptive · partial · unknown · source-side · collector-side · downstream sampling`.
```text
sampled telemetry   ≠ complete population
no sampled error    ≠ no error occurred
sampling state unknown ≠ unsampled
sampling rate       ≠ reliability guarantee
```
Keine Sampling-Technologie ausgewählt.

## 25. Aggregation

`time-window · resource · workspace · class · percentile-like · count · rate · rollup · unknown aggregation`. Je Aggregation: source signals · scope · time boundary · aggregation function class · sampling state · missing-value behavior · quality impact · lossiness · provenance.
```text
aggregate success ≠ every source valid
rollup            ≠ raw data preserved
aggregate value   ≠ authoritative individual value
```
Keine Aggregation Engine ausgewählt.

## 26. Cardinality

Berücksichtigt: Resource/User/Principal/Workspace/Request IDs · dynamic values · paths · error messages · external labels · vendor attributes · unbounded strings.
```text
available label      ≠ safe telemetry dimension
unique request identity ≠ suitable persistent metric label
external attribute   ≠ trusted canonical label
high cardinality     ≠ higher observability quality
```
Keine Cardinality-Grenze/Storage-Technologie ausgewählt (Bezug THR-036).

## 27. State-Authority Boundary

```text
telemetry signal      ≠ authoritative CoreOps state
normalized observation ≠ effective state
derived health        ≠ policy decision
telemetry freshness   ≠ state authority
metric threshold      ≠ execution authorization
```
[SoT](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) und [State-Modell](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) bleiben autoritativ; **kein Parallelmodell**.

## 28. Event and Evidence Boundary

Telemetrie **kann** zu einem Event beitragen, als Evidence referenziert werden, Verifikation stützen oder spätere Policy-Evaluation auslösen. **Aber**: `telemetry ≠ audit event automatically`; `telemetry ≠ sufficient evidence automatically`; `alert condition ≠ approved remediation`; `anomaly ≠ confirmed incident`; `threshold crossed ≠ execution command`. Umwandlung/Referenz braucht Provenance und klare Klassifikation (konsistent mit [Event/Audit](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Evidence](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md)).

## 29. Offline Telemetry

Offline-Telemetry-Pakete folgen Companion 3 §23 (target-environment binding, provenance, integrity, import governance). Nicht behauptet: implementiert, konfliktfreie automatische Reconciliation, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-Technologie (Bezug THR-024).

## 30. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Telemetry must not grant execution authority.
2. Telemetry signals must not inherit authoritative state automatically.
3. Raw, normalized, derived and aggregated telemetry remain separate.
4. Normalization must preserve source provenance and transformation history.
5. Unknown or conflicting units must block unsafe automatic conversion.
6. Missing telemetry must not be interpreted as zero or absence of activity.
7. Freshness must distinguish observation time from ingestion time.
8. Sampling and aggregation limitations must remain visible.
9. High confidence must not imply authoritative truth.
10. Telemetry-to-event and telemetry-to-evidence transitions require explicit classification and provenance.

## 31. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-005, THR-008, THR-009, THR-010, THR-011, THR-012, THR-013, THR-014, THR-015, THR-019, THR-020, THR-024, THR-025, THR-028, THR-029, THR-034, THR-035, THR-036, THR-037, THR-039.

## 32. Technology Boundary

Nicht ausgewählt/implementiert: Telemetry-Protokoll (OTel/Prometheus/SNMP/Syslog), Metric-/Log-/Trace-Schema, Serialisierung, Collector/Agent, Time-Series-DB/Storage, Dashboard, Alerting/Rule-Engine, Unit-Library, Runtime-Code.

## 33. Compatibility

Konsistent mit [Event/Audit](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Evidence](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [SoT/Provenance/State](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)/[Drift](DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md), [Integration Contract/Trust](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [API Governance](API_GOVERNANCE_AND_OPERATION_MODEL.md), [Data Ownership](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Policy/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert DEC-P-06 (fehlende Daten ≠ gesund), CAP-MONITORING-*.

## 34. Open Questions

- Telemetry-Protokoll-/Schema-Ansatz (spätere ADR).
- Trusted-Time-/Freshness-Schwellen je Signal-Klasse.
- Cardinality-Grenzen (mit THR-036).

## 35. Next Decision

Companion 2 trägt Mapping/Quality/Compatibility, Companion 3 die Trust/Privacy/Disclosure-Policy. Topology Graph (CO-WP-020) referenziert topology observation. Protokoll-/Storage-Wahl bleibt einer späteren ADR-Runde vorbehalten.
