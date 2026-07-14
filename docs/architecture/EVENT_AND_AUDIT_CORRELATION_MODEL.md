# CoreOps – Event and Audit Correlation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation event and audit-correlation model
> Implementation Status: Not implemented
> Event Transport: Not selected
> Event Schema: Not selected
> Ordering Mechanism: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-018 (docs-only / event, audit, evidence and security-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Event-Identität, Event-Klassen, Correlation/Causation, Zeit-/Sequenzunsicherheit, Audit Events und Audit Completeness**. Companion zu [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) und [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Es wählt **keinen** Event Bus, kein Logging-Framework, kein SIEM, kein Schema, keinen Ordering-Mechanismus.

## 2. Purpose

Ein Ereignis ist kein Befehl, eine Korrelation ist keine Kausalität, und ein aufgezeichneter Audit-Datensatz ist nicht automatisch wahr oder vollständig. Dieses Modell trennt Event-Identität von Correlation/Request/Operation, trennt vier Zeitbegriffe und macht Audit Gaps sichtbar — damit `event ≠ command`, `correlation ≠ causation`, `timestamp ≠ authoritative global ordering` und `no event present ≠ action did not occur`.

## 3. Scope

Event-Begriffe/-Klassen · Event Identity · Producers/Sources · Subjects/Objects · Principals/Workspace/Scope · Event Time · Clock Uncertainty · Correlation/Causation · Operation/Attempt-Linkage · Ordering/Sequence · Duplicate/Replay/Repeated · Audit Events · Audit Lifecycle · Audit Gaps/Completeness · Offline Events · Audit/Evidence References. Evidence-Detail (Companion 2), Integrity/Retention/Disclosure (Companion 3).

## 4. Non-Goals

- Kein Event-Bus/Broker/Queue, kein Logging-Framework/Logformat, kein SIEM.
- Kein Datenbank-/Storage-Schema, kein UUID-/ID-Format, keine Timestamp-/Clock-Synchronisation.
- Kein Hash-/Signatur-/Ledger-/WORM-Mechanismus, keine Ordering-Technologie, kein Runtime-Code.
- Keine Behauptung implementierter Audits oder technischer Unveränderlichkeit.

## 5. Concepts

Begriffe (mindestens): event · domain event · audit event · security event · state observation · notification · command · request · operation · attempt · result · evidence · evidence reference · evidence source · evidence producer · correlation · causation · sequence · event/observation/recording/ingestion time · principal · accountable human owner · subject · object · scope · workspace · retention · disclosure · export · audit gap.

**Grundregeln:**
```text
event              ≠ command
event              ≠ notification
audit event        ≠ proof that the described action occurred exactly as reported
log entry          ≠ governed audit record
audit record       ≠ validated evidence
evidence available ≠ evidence valid ≠ evidence sufficient
correlation        ≠ causation
timestamp          ≠ authoritative global ordering
received time      ≠ occurrence time
export capability  ≠ disclosure authorization
```

## 6. Event Classes

Je Klasse: Purpose · Authoritative producer/owner · Typical subject · Typical object · Workspace/scope expectation · Principal attribution · Correlation expectation · Evidence relevance · Retention consideration · Disclosure consideration · Offline behavior · Threat references. **Keine vollständige Event-ID-Liste.**

| Event-Klasse | Autoritativer Producer (Konzept) | Threat refs |
|---|---|---|
| identity and access event | MOD-IAM | THR-001, THR-002 |
| machine-identity event | MOD-IAM/SEC | THR-008, THR-009 |
| policy-evaluation event | MOD-POL | THR-003 |
| approval event | MOD-POL | THR-003, THR-004 |
| authorization event | MOD-POL/EXE | THR-004, THR-026 |
| request and API event | MOD-EXP | THR-038, THR-035 |
| integration event | MOD-ADP | THR-010, THR-011 |
| operation event | MOD-EXE/WFL | THR-005, THR-028 |
| execution event | MOD-EXE | THR-005, THR-029 |
| verification event | MOD-STA/EVD | THR-029, THR-030 |
| resource and inventory event | MOD-INV | THR-014 |
| state and drift event | MOD-STA | THR-012, THR-013 |
| deployment event | MOD-DEP | THR-021, THR-022 |
| migration event | MOD-STA/data | THR-014, THR-033 |
| Domain-Pack lifecycle event | MOD-EXT | THR-023 |
| security event | MOD-SEC | THR-037 |
| audit-governance event | MOD-EVD | THR-016, THR-017 |
| offline import and export event | MOD-OFF | THR-024, THR-025 |
| system-health event | MOD-OBS | THR-036 |
| administrative event | MOD-POL/SEC | THR-002, THR-037 |

## 7. Event Identity

Ein Event benötigt konzeptionell mindestens: stable event identity · event class · event type · producer · source · source trust state · subject · object · principal · accountable human owner (wo nötig) · workspace/environment · scope · event time · observation time · recording time · sequence context · correlation references · causation references · request/operation/attempt/result reference · policy/authorization references (falls anwendbar) · provenance reference · integrity state · audit classification · retention classification · disclosure classification · evidence references.
```text
event identity ≠ correlation identity ≠ request identity ≠ operation identity
event producer ≠ automatically authoritative source
event recorded ≠ event validated
```
Keine ID-Technologie ausgewählt.

## 8. Producers and Sources

Producer (erzeugt Event) und Source (Ursprung der beschriebenen Beobachtung) sind getrennt; `event producer ≠ automatically authoritative source`. Source Trust State ist explizit (konsistent mit [Integration Trust](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md) und [SoT](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)).

## 9. Subjects and Objects

Subject (handelnde/betroffene Identität) und Object (betroffene Ressource) sind getrennt; beide scope-/workspace-gebunden. Ein Event erteilt keine Autorität über Subject/Object (§23 Invarianten).

## 10. Principals, Workspace and Scope

Principal-Zuordnung (human/machine) mit `accountable human owner` wo nötig; Workspace/Scope explizit. Cross-Workspace-Korrelation gewährt keinen Cross-Workspace-Zugriff (Companion 3 §19).

## 11. Event Time

Getrennt: `event occurrence time · source observation time · CoreOps recording time · ingestion time · reconciliation time`.
```text
newer timestamp ≠ newer authoritative state
same timestamp  ≠ same event
earlier received ≠ earlier occurred
clock synchronized ≠ globally ordered
```

## 12. Clock Uncertainty

Berücksichtigt: unsichere/fehlende Source-Zeit · Clock Drift · Offline-/Air-Gap-Verzögerung · Replayed Events · verspätete Events · Zeitkorrekturen · Sequenz- statt Zeitbezug · Conflict-/Provenance-Auswirkung. Keine Clock-Synchronisation ausgewählt (Bezug THR-027).

## 13. Correlation

Typen: request · operation · attempt · workflow · resource · principal · workspace correlation · causation · parent-child · reconciliation relationship.
```text
same correlation identity ≠ same event
correlated events         ≠ proven causal relationship
causation reference       ≠ execution authorization
shared resource           ≠ shared authority
```
**Correlation-IDs sind keine Secrets/Berechtigungsartefakte.**

## 14. Causation

Causation (parent-child/reconciliation) ist eine explizite, belegte Beziehung, nicht aus Correlation abgeleitet. `correlation ≠ causation`; eine Causation-Referenz gewährt keine Execution Authority.

## 15. Operation and Attempt Linkage

Audit Events verknüpfen (wo anwendbar): request/operation/attempt identity · idempotency context reference · executing principal · accountable owner · target · scope · policy/approval/execution-authorization decision · acceptance/execution/result/verification/closure state. **Mehrere Attempts werden nicht zu einem scheinbar eindeutigen Resultat verschmolzen.**
```text
one operation may have multiple attempts
attempt completed ≠ operation verified
```

## 16. Event Ordering and Sequence

Dokumentiert: producer-local · operation attempt · workspace-local · offline package · reconciliation sequence · fehlende/unterbrochene Sequenzen · duplicate sequence values · sequence reset · source-Neuregistrierung · mehrere unabhängige Producer.
```text
sequence         ≠ global total order
higher sequence  ≠ greater authority
sequence gap     ≠ definite event loss without further evidence
duplicate sequence ≠ definite duplicate event
```
Keine Sequence-/Ordering-Technologie ausgewählt.

## 17. Duplicate, Replay and Repeated Events

Zu unterscheiden: duplicate recording · duplicate delivery · replayed event · repeated real-world event · repeated operation attempt · reconciliation copy · historical re-import. Regeln: Duplicate-Detection nicht nur per Timestamp; Replay-Verdacht sichtbar; legitimes wiederholtes Event nicht still dedupliziert; Attempt-/Event-Historie nicht überschrieben; Offline Re-Import erzeugt keine neue Autorität. Kein Deduplication-/Replay-Mechanismus (Bezug THR-026).

## 18. Audit Events

Ein Audit Event beschreibt eine governance-relevante Beobachtung/Statusänderung. Es ist **keine** automatische Behauptung vollständiger Wahrheit, technischer Unveränderlichkeit, rechtlicher Beweiskraft, Compliance oder erfolgreicher Kontrolle. Mindestens auditiert: identity lifecycle transition · role/membership change · machine enrollment/revocation · policy evaluation · approval decision · execution authorization · break-glass use · request acceptance · execution start/result · verification · state-authority conflict · drift/remediation · migration · API deprecation use · Domain-Pack support/trust change · offline import/activation · evidence export · audit configuration change · retention/purge decision.

## 19. Audit Lifecycle

Konzeptionelle Statuswerte:
```text
generated → received → recording-pending → recorded → validation-pending → validated →
validation-failed → correlation-pending → correlated → gap-suspected → gap-confirmed →
superseded → archival-pending → archived → retention-review → purge-pending → purged → invalidated
```
```text
recorded   ≠ validated
validated  ≠ complete
correlated ≠ causally proven
archived   ≠ unavailable
purged     ≠ historical event never existed
```
Purge orphant keine erforderlichen Referenzen/Governance-Nachweise (Companion 3 §16).

## 20. Audit Gaps and Completeness

Gap-Typen: expected event missing · producer/collector unavailable · recording failure · partial export/import · sequence/time/correlation gap · retention-related absence · unauthorised disclosure restriction · unknown completeness.
```text
no event present    ≠ action did not occur
audit gap suspected ≠ audit gap confirmed
complete local audit ≠ globally complete evidence
collector healthy   ≠ every producer reported
```
**Audit Completeness ist scopegebunden und kann `unknown` bleiben.**

## 21. Offline Events

Offline-Event-Pakete folgen Companion 3 §21: producer identity · target environment · workspace/scope · event sequence · clock uncertainty · provenance/integrity status · validity boundary · duplicate/replay context · local recording · local activation/import · partial package state · revocation-distribution challenge · reconciliation · conflict state · audit continuity. Nicht behauptet: implementiert, beliebige Air-Gap-Stufen, konfliktfreie automatische Reconciliation, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024).

## 22. Audit and Evidence References

Events referenzieren Evidence über stabile Evidence References (Companion 2); eine Referenz ist kein Ersatz für das Artefakt und kein Nachweis seiner Existenz/Gültigkeit.

## 23. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Events must not grant execution authority.
2. Audit records must not be treated as inherently true or complete.
3. Correlation must not be treated as proven causation.
4. Timestamps and sequences must not be treated as authoritative global ordering.
5. Missing events must not prove that an action did not occur.
6. Duplicate handling must preserve event and attempt history.
7. An event producer is not automatically the authoritative source.
8. Correlation identities must not be secrets or authorisation artifacts.
9. Multiple attempts must not be merged into one apparently unique result.

## 24. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-001, THR-002, THR-003, THR-004, THR-005, THR-008, THR-009, THR-010, THR-011, THR-012, THR-013, THR-014, THR-015, THR-016, THR-017, THR-018, THR-021, THR-022, THR-023, THR-024, THR-025, THR-026, THR-027, THR-028, THR-029, THR-030, THR-035, THR-036, THR-037, THR-040.

## 25. Technology Boundary

Nicht ausgewählt/implementiert: Event-Bus/Broker/Queue, Logging-Framework/Logformat, SIEM, DB/Storage/Schema, ID-Format, Timestamp-/Clock-Synchronisation, Ordering-Mechanismus, Hash/Signatur/Ledger/WORM, Runtime-Code.

## 26. Compatibility

Konsistent mit [API Governance](API_GOVERNANCE_AND_OPERATION_MODEL.md), [Integration Contract/Trust](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [SoT/Provenance/State](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Data Ownership/Migration](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Policy/Approval/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Threat Model](../security/COREOPS_FOUNDATION_THREAT_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert die Evidence-/Audit-Aspekte von DEC-P-06, DEC-G-08 (Invarianten) und das Evidence-Capability-Modell (CO-WP-004E).

## 27. Open Questions

- Event-Schema-/Ordering-Ansatz (spätere ADR).
- Standardisierte Correlation-Konventionen im Detail.
- Trusted-Time-/Sequence-Ansatz (mit THR-027).

## 28. Next Decision

Companion 2 trägt das Evidence-Modell, Companion 3 die Integrity-/Retention-/Disclosure-Policy. Telemetry/Normalization (CO-WP-019) referenziert Event-Zeit/Provenance. Transport-/Schema-Wahl bleibt einer späteren ADR-Runde vorbehalten.
