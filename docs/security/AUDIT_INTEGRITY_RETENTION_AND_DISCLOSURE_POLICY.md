# CoreOps – Audit Integrity, Retention and Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation audit-integrity, retention and disclosure policy
> Implementation Status: Not implemented
> Audit Storage: Not selected
> Integrity Mechanism: Not selected
> Retention Automation: Not implemented
> Disclosure Automation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-018 (docs-only / event, audit, evidence and security-governance foundation)

## 1. Status

Technologieunabhängige Policy für **Audit-Integrität, Completeness, Retention, Disclosure/Export, Workspace-Isolation, Offline-Continuity und Fail-Closed**. Companion zu [EVENT_AND_AUDIT_CORRELATION_MODEL.md](../architecture/EVENT_AND_AUDIT_CORRELATION_MODEL.md) und [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md). Kein Audit-Storage, kein Integrity-Mechanismus, keine Retention-/Disclosure-Automation.

## 2. Purpose

Audit- und Evidenzdaten sind ein Angriffs- und Fehlerziel: Manipulation, fehlende Aufzeichnung, unautorisierte Offenlegung. Diese Policy legt fest, wer auditiert/offenlegt und warum `append-only-governed ≠ technically immutable`, `audit administrator ≠ unrestricted disclosure authority`, `read permission ≠ export permission` und `missing evidence ≠ evidence of absence`.

## 3. Scope

Authority Model · Audit Event Boundary · Producer/Source Trust · Recording · Validation · Correlation · Completeness/Gaps · Integrity · Transformations/Redaction · Retention · Archival · Purge · Disclosure · Export · Workspace Isolation · Privacy/Data Minimization · Offline Audit Continuity · Failure/Unknown State · Closure · Audit/Evidence · Fail-Closed.

## 4. Non-Goals

- Kein Audit-Storage/SIEM/Logging-Framework, kein Hash-/Signatur-/Ledger-/WORM-Mechanismus.
- Keine Retention-/Redaction-/Export-Automation, keine konkreten Fristen.
- Keine rechtliche Beweiskraft/Compliance behauptet; kein Runtime-Code.

## 5. Authority Model

Getrennt: Audit Owner · Retention Owner · Disclosure Owner · Evidence Owner · Storage Responsibility · Audit Administrator. `storage administrator ≠ audit owner`; `audit administrator ≠ unrestricted event editor`; `audit administrator ≠ unrestricted disclosure authority`. Aufbauend auf [Data Ownership](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md) (kein Parallelmodell).

## 6. Audit Event Boundary

Ein Audit Event ist eine governance-relevante Beobachtung, **keine** automatische Wahrheits-/Vollständigkeits-/Compliance-Behauptung (aus [Event Model §18]). `audit record ≠ validated evidence`; `event ≠ command`; **Events erteilen keine Execution Authority**.

## 7. Producer and Source Trust

Producer-Identität und Source Trust sind explizit; `event producer ≠ authoritative source`. Untrusted/kompromittierte Producer (Bezug THR-008, THR-010) produzieren keine vertrauenswürdige Audit-Wahrheit.

## 8. Recording

Recording ist ein eigener Schritt (`recording-pending`/`recorded`); `recorded ≠ validated`. Recording-Failure ist sichtbar zu behandeln (§22). Recording erzeugt keine technische Unveränderlichkeit.

## 9. Validation

Audit-/Evidence-Validierung ist eigenständig (`validation-pending`/`validated`/`validation-failed`); `validated ≠ complete ≠ sufficient` (Companion 2 §16-17).

## 10. Correlation

Correlation verknüpft Events (Companion 1 §13); `correlated ≠ causally proven`; Correlation-Zugriff gewährt keinen Zugriff auf alle korrelierten Events (§18).

## 11. Completeness and Gaps

Audit Completeness ist scopegebunden und kann `unknown` bleiben; `no event present ≠ action did not occur`; `gap suspected ≠ gap confirmed`; `collector healthy ≠ every producer reported` (Companion 1 §20). Gaps werden erfasst, nicht verborgen.

## 12. Integrity

Berücksichtigt: producer identity · source trust · event identity · sequence context · schema reference · provenance · recording status · transformation history · integrity status · validation status · gap status · retention status.
```text
append-only-governed     ≠ technically immutable
integrity metadata present ≠ integrity verified
```
Keine Hash-/Signatur-/Ledger-/WORM-Technologie ausgewählt (Bezug THR-016, THR-017).

## 13. Transformations and Redaction

Normalisierung/Redaction erhalten Provenance und ursprüngliche Referenzen; Transformation entfernt **keine** erforderliche Provenance und schreibt Historie nicht still um. Redaction dient Datenminimierung/Disclosure, nicht Verschleierung. Keine Redaction-Technologie ausgewählt.

## 14. Retention

Retention Owner/Class/Trigger/Review-Trigger dokumentiert; **keine konkreten Fristen**. `retention expiry ≠ automatic purge authority`; Holds/Exceptions möglich (`legal/organisational hold ≠ permanent retention`).

## 15. Archival

`archived ≠ unavailable`; `archived ≠ safe to purge`. Archivierte Audit-/Evidenzdaten bleiben referenzierbar und provenance-erhaltend.

## 16. Purge

Purge benötigt Purge Authority + Approval; erfasst affected references · evidence impact · disclosure impact · audit reference. `purged event ≠ historical event never existed`; **Purge orphant keine erforderlichen Referenzen/Governance-Nachweise still**.

## 17. Disclosure

Disclosure ist scope-/authority-gebunden; `read permission ≠ disclosure/export permission`; `audit administrator ≠ unrestricted disclosure authority`; `correlation reference ≠ permission to access all correlated events`.

## 18. Export

Audit-/Evidence-Export benötigt: requesting principal · consumer · workspace · subject/object scope · event/evidence classes · time/sequence scope · purpose · disclosure classification · redaction requirement · approval requirement (falls anwendbar) · export format reference · provenance preservation · integrity status · validity boundary · audit reference. `export completed ≠ recipient authorised for every contained item` (Bezug THR-018, THR-040). Keine Export-/Redaction-Technologie ausgewählt.

## 19. Workspace Isolation

Workspace Scope explizit; Cross-Workspace-Audit/-Evidence-Zugriff nur mit ausdrücklicher Autorität; Correlation/Export legt keine Cross-Workspace-Daten offen (Bezug THR-035).

## 20. Privacy and Data Minimization

Audit-/Evidence-Daten erfassen nur Notwendiges: Identitätsreferenzen statt Vollprofile · keine Raw Secrets/Credential-Inhalte · keine vollständigen Payloads ohne Notwendigkeit · getrennte consumer-safe/interne Diagnosedaten · Workspace-Isolation · Subject Access/Disclosure als spätere Governance · Retention Review · Redaction History. **Keine vollständige Datenschutz-/Rechtskonformität behauptet** (Bezug DEC-P-08, CO-WP-025).

## 21. Offline Audit Continuity

Offline-Event-/Evidence-Pakete benötigen: producer identity · target environment · workspace/scope · event sequence · clock uncertainty · provenance/integrity status · validity boundary · duplicate/replay context · local recording · local activation/import · partial package state · revocation-distribution challenge · reconciliation · conflict state · audit continuity. Nicht behauptet: implementiert, beliebige Air-Gap-Stufen, konfliktfreie automatische Reconciliation, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance.

## 22. Failure and Unknown State

`producer/collector/recording failed · storage unavailable · partial write · validation failed · correlation incomplete · export/offline import incomplete · integrity/provenance/completeness unknown`.
```text
audit failure          ≠ target action failed
recording failure      ≠ action did not occur
unknown audit completeness ≠ complete audit
missing evidence       ≠ evidence of absence
```
Security-relevante Audit-Failures werden sichtbar/auditierbar behandelt; keine Alerting-Technologie ausgewählt.

## 23. Closure

Audit-/Evidence-Vorgänge dürfen `closed` werden, ohne erfolgreich/vollständig zu sein. Closure benötigt: closure reason · final audit state · final evidence state · completeness status · validation status · remaining gap/exception · owner · audit reference.
```text
closed ≠ complete ≠ validated ≠ sufficient
```

## 24. Audit and Evidence

Audit-Governance-Ereignisse (config change, retention/purge decision, disclosure/export, integrity finding) sind selbst auditiert; Historie nicht umgeschrieben.

## 25. Fail-Closed Rules

Keine privilegierte Offenlegung/Purge/Import bei: unklarer Disclosure-Autorität · fehlender Provenance (offline) · unbekannter Integrität · unklarer Workspace-Grenze · `conflicted` Evidenz für privilegierte Entscheidung · fehlender Approval für destruktive Retention/Purge · unbekannter Audit-Completeness für sicherheitsrelevante Entscheidungen.

## 26. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Events must not grant execution authority.
2. Audit records must not be treated as inherently true or complete.
3. Correlation must not be treated as proven causation.
4. Missing events must not prove that an action did not occur.
5. Duplicate handling must preserve event and attempt history.
6. Evidence availability must not imply validity or sufficiency.
7. Evidence sufficiency remains decision-, scope- and time-bound.
8. Append-only-governed is not technically immutable; audit administrators are not unrestricted editors.
9. Audit administrators must not receive unrestricted disclosure authority.
10. Audit and evidence exports must preserve provenance and scope.
11. Offline audit import requires provenance, integrity, target binding and explicit import governance.
12. Audit closure must not imply completeness, validation or compliance.

## 27. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-002, THR-008, THR-010, THR-013, THR-016, THR-017, THR-018, THR-019, THR-020, THR-024, THR-025, THR-026, THR-027, THR-035, THR-037, THR-040. Keine Duplikation, kein Parallelregister.

## 28. Technology Boundary

Nicht ausgewählt/implementiert: Audit-Storage/SIEM/Logging, Hash/Signatur/Ledger/WORM, Retention-/Redaction-/Export-Automation, Alerting, Runtime-Code.

## 29. Compatibility

Konsistent mit Event-/Evidence-Companion, [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Data Ownership/Migration](../architecture/DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Offline Reconciliation](OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md), [Public Neutrality/Disclosure](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md). Konkretisiert DEC-G-06 (keine Secrets in Logs/Exports), DEC-G-08, DEC-P-08.

## 30. Open Questions

- Integrity-Mechanismus (Hash/Signatur/WORM) — spätere ADR.
- Konkrete Retention-Fristen/Holds (CO-WP-025).
- Alerting für sicherheitsrelevante Audit-Failures (CO-WP-026).

## 31. Next Decision

Telemetry/Normalization (CO-WP-019), Data Classification/Retention (CO-WP-025) und Self-Protection/Recovery (CO-WP-026) konkretisieren Retention-/Integrity-/Alerting-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
