# CoreOps – Telemetry Trust, Privacy and Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation telemetry trust, privacy and disclosure policy
> Implementation Status: Not implemented
> Telemetry Collector: Not selected
> Redaction Mechanism: Not selected
> Export Mechanism: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-019 (docs-only / telemetry, observability and normalization foundation)

## 1. Status

Technologieunabhängige Policy für **Telemetry-Vertrauen, Privacy, Cardinality/Labels, Disclosure/Export, Workspace-Isolation, Telemetry-to-Event/-Evidence-Grenzen, Offline und Fail-Closed**. Companion zu [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](../architecture/TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) und [TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md](../architecture/TELEMETRY_MAPPING_QUALITY_AND_COMPATIBILITY_MODEL.md). Kein Collector, kein Redaction-/Export-Mechanismus.

## 2. Purpose

Telemetrie ist eine Angriffs-, Fehl- und Datenschutzfläche: manipulierte Observations, Label-Leakage, Cardinality-Missbrauch, Threshold-getriebene Fehlaktionen. Diese Policy legt fest, wie wenig einem Signal vertraut wird und warum `telemetry ≠ authoritative state`, `metric threshold ≠ execution authorization`, `available label ≠ safe disclosure dimension` und `missing telemetry ≠ evidence of absence`.

## 3. Scope

Trust Boundary · Producer/Source Trust · Resource/Workspace Binding · Raw Signal Handling · Normalization Trust · Derived/Aggregated Signals · Freshness/Staleness · Quality/Confidence · Sampling · Cardinality · Labels/Attributes · Logs/Diagnostic Payloads · Privacy/Data Minimization · Disclosure · Export · Cross-Workspace Aggregation · Telemetry-to-Event/-Evidence · Offline · Failure/Unknown · Fail-Closed.

## 4. Non-Goals

- Kein Collector/Protokoll/Storage, kein Redaction-/Export-/Alerting-Mechanismus.
- Keine Behauptung implementierter/validierter Kontrollen; keine Health-/Observability-/Compliance-Garantie; kein Runtime-Code.

## 5. Trust Boundary

Telemetrie stammt jenseits einer Vertrauensgrenze (TB-03/TB-06/TB-07). `signal received ≠ source state currently true`; Trust ist scope-/quality-gebunden. Telemetrie erteilt **keine** Execution Authority und wird nicht still autoritativer State (Signal Model §27).

## 6. Producer and Source Trust

Producer-Identität und Source Trust explizit; `producer ≠ authoritative source`. Kompromittierte/gefälschte Producer (Bezug THR-008, THR-009, THR-010) liefern keine vertrauenswürdige Telemetrie-Wahrheit; manipulierte Observations (THR-012) bleiben quality-/provenance-kennzeichenbar.

## 7. Resource and Workspace Binding

Signale binden an Resource/Subject und Workspace/Scope; falsche/unklare Bindung ⇒ Quality `unknown`. Cross-Workspace-Zuordnung nur mit Autorität (§20, Bezug THR-035, THR-014).

## 8. Raw Signal Handling

Raw-Signale bleiben source-/provenance-erhaltend; `raw ≠ trustworthy ≠ authoritative`. Raw Logs/Payloads unterliegen Disclosure-Grenzen (§16).

## 9. Normalization Trust

Normalisierung erhält Source-Provenance/Transformation History; `normalized ≠ validated`; `mapping succeeded ≠ semantic correctness verified`. Fallback/Default-Mapping ist explizit (Companion 2).

## 10. Derived and Aggregated Signals

`derived ≠ independently observed`; `aggregated ≠ complete population`; Sampling-/Aggregations-Limitationen bleiben sichtbar (§13). Abgeleitete Health ist keine Policy-Entscheidung.

## 11. Freshness and Staleness

`recently ingested ≠ recently observed`; `recently observed ≠ current target state`; `stale ≠ invalid automatically`; `no recent sample ≠ zero ≠ resource unavailable` (Signal Model §23, Bezug THR-013, THR-027).

## 12. Quality and Confidence

Quality/Confidence getrennt; `high confidence ≠ authoritative`; `validated ≠ fresh`. Unbekannte Quality wird nicht als validiert behandelt.

## 13. Sampling

`sampled ≠ complete population`; `no sampled error ≠ no error occurred`; `sampling state unknown ≠ unsampled`. Sampling-Limitationen bleiben in Interpretation und Evidence-Claims sichtbar (Bezug THR-031).

## 14. Cardinality

Hohe Cardinality ist Betriebs-/Disclosure-Risiko; `available label ≠ safe telemetry dimension`; `unique request identity ≠ suitable persistent metric label`; `high cardinality ≠ higher observability quality`. Keine Cardinality-Grenze/Storage-Technologie ausgewählt (Bezug THR-036).

## 15. Labels and Attributes

Labels/Attribute enthalten keine unnötigen personenbezogenen Daten, keine Secrets/Credential-Inhalte, keine unautorisierten Cross-Workspace-Identifikatoren; `external attribute ≠ trusted canonical label` (Bezug THR-019, THR-020, THR-035).

## 16. Logs and Diagnostic Payloads

Raw Logs/Diagnostic Payloads benötigen besondere Disclosure-Grenzen; `diagnostic payload ≠ unrestricted disclosure`; getrennte consumer-safe/interne Diagnosedaten. Keine Secrets/Credentials.

## 17. Privacy and Data Minimization

Telemetrie erfasst nur Notwendiges: Identitätsreferenzen statt Vollprofile · keine Raw Secrets/Credential-Inhalte · keine vollständigen Payloads ohne Notwendigkeit · getrennte consumer-safe/interne Views · Workspace-Isolation · Subject Access/Disclosure als spätere Governance · Retention Review · Redaction History. **Keine vollständige Datenschutz-/Rechtskonformität behauptet** (Bezug DEC-P-08, CO-WP-025).

## 18. Disclosure

Disclosure ist scope-/authority-gebunden; `read permission ≠ export permission`; Ressourcenexistenz nicht unautorisiert offengelegt (auch nicht via Labels/Aggregation).

## 19. Export

Telemetry-Export benötigt eigene Autorität, Scope, Redaction-Anforderung, Provenance-Erhalt, Integrity-Status, Validity-Boundary, Audit-Referenz (konsistent mit [Audit Disclosure](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)). Keine Export-/Redaction-Technologie ausgewählt.

## 20. Cross-Workspace Aggregation

Cross-Workspace-Aggregation benötigt explizite Governance; Aggregation legt keine Cross-Workspace-Rohdaten/-Identifikatoren offen (Bezug THR-035).

## 21. Telemetry-to-Event Boundary

`telemetry ≠ audit event automatically`; `anomaly ≠ confirmed incident`; `alert condition ≠ approved remediation`; `threshold crossed ≠ execution command`. Eine Telemetry→Event-Umwandlung braucht explizite Klassifikation und Provenance (Companion 1 §28).

## 22. Telemetry-to-Evidence Boundary

`telemetry ≠ sufficient evidence automatically`; Telemetrie als Evidence braucht Provenance, Quality/Validation und Sufficiency-Bewertung ([Evidence Model](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md)); `available ≠ valid ≠ sufficient`.

## 23. Offline Telemetry

Offline-Telemetry-Pakete benötigen: producer/source identity · target environment · workspace/scope · signal classes · source schema · normalization profile · observation boundary · clock uncertainty · sequence context · sampling state · provenance/integrity/quality status · partial package state · validity boundary · duplicate/replay context · import quarantine · explicit import · reconciliation. Nicht behauptet: implementiert, konfliktfreie automatische Reconciliation, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance.

## 24. Failure and Unknown State

`producer/collector unavailable · source schema unknown · mapping missing/failed · unit unknown/conflict · normalisation failed · partial transformation · sampling/quality/freshness/provenance/integrity unknown`.
```text
telemetry failure     ≠ target failure
normalisation failure ≠ source value invalid automatically
missing telemetry     ≠ evidence of absence
unknown quality       ≠ validated quality
```

## 25. Audit and Evidence

Telemetry-Trust-/Disclosure-/Export-Ereignisse werden auditiert; Provenance/Scope erhalten. Trennung `telemetry ≠ event ≠ evidence ≠ compliance`.

## 26. Fail-Closed Rules

Keine privilegierte Nutzung/Offenlegung bei: unklarer Producer-/Source-Trust · unbekannter/`conflicted` Unit für automatische Konversion · unbekannter Quality/Provenance für sicherheitsrelevante Entscheidung · fehlender Disclosure-Autorität · unklarer Workspace-Grenze · unbekanntem Offline-Provenance · Threshold-getriebener Aktion ohne Policy/Approval/Execution Authorization (CO-WP-013).

## 27. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Telemetry must not grant execution authority.
2. Telemetry must not inherit authoritative state automatically.
3. A producer is not automatically the authoritative source.
4. Unknown or conflicting units must block unsafe automatic conversion.
5. Missing telemetry must not be interpreted as zero, inactivity or target failure.
6. Sampling and aggregation limitations must remain visible.
7. Telemetry labels must not expose secrets or unauthorised workspace data.
8. High confidence must not imply authoritative truth.
9. Telemetry-to-event and telemetry-to-evidence transitions require explicit classification and provenance.
10. A metric threshold or anomaly must not by itself authorise remediation.
11. Offline telemetry import requires provenance, integrity, target binding and explicit governance.
12. Cross-workspace aggregation and disclosure require explicit authority.

## 28. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-008, THR-009, THR-010, THR-011, THR-012, THR-013, THR-014, THR-019, THR-020, THR-024, THR-026, THR-027, THR-029, THR-031, THR-034, THR-035, THR-036. Keine Duplikation, kein Parallelregister.

## 29. Technology Boundary

Nicht ausgewählt/implementiert: Telemetry-Collector/Protokoll, Redaction-/Export-/Alerting-Mechanismus, Storage/Time-Series-DB, Dashboard, Runtime-Code.

## 30. Compatibility

Konsistent mit Signal-/Mapping-Companion, [Event/Audit](../architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Evidence](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [SoT/State](../architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)/[Drift](../architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md), [Policy/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md), [Public Neutrality/Disclosure](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md). Konkretisiert DEC-P-06, DEC-G-06, DEC-P-08.

## 31. Open Questions

- Cardinality-Grenzen/Quota (THR-036).
- Trusted-Time-/Freshness-Schwellen (THR-027).
- Redaction-/Export-Mechanismus (CO-WP-025).

## 32. Next Decision

Topology Graph (CO-WP-020), Data Classification/Retention (CO-WP-025) und Self-Protection (CO-WP-026) konkretisieren Cardinality-/Disclosure-/Alerting-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
