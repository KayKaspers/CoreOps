# CoreOps – API Error, Idempotency and Replay Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation API error, idempotency and replay policy
> Implementation Status: Not implemented
> Error-Code Binding: Not selected
> Idempotency Mechanism: Not selected
> Replay Protection: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-017 (docs-only / API architecture and security foundation)

## 1. Status

Technologieunabhängige Policy für **Error-/Problem-Semantik, Disclosure-Grenzen, Retry-Klassifikation, Idempotency, Duplicate/Replay und Unknown Outcome** an der API-Grenze. Companion zu [API_GOVERNANCE_AND_OPERATION_MODEL.md](../architecture/API_GOVERNANCE_AND_OPERATION_MODEL.md) und [API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md](../architecture/API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md). Kein Error-Code-Binding, kein Idempotency-/Replay-Mechanismus.

## 2. Purpose

Fehler und Wiederholungen sind sicherheitskritisch: eine Fehlerantwort beweist keine Nichtausführung, ein Duplicate ist keine Doppelausführung, und ein Idempotency-Key ist keine Autorisierung. Diese Policy legt fest, wie Fehler klassifiziert/offengelegt werden und warum `error response ≠ no side effect`, `unknown outcome → no automatic retry` und `idempotency context ≠ execution authorization`.

## 3. Scope

Error Boundary/Classes · Consumer-Safe Error Information · Internal Diagnostics · Retry Classification · Idempotency Concepts/Context · Duplicate Handling · Replay Boundary · Unknown Outcome · Bulk/Partial · Async · Pagination/Continuation · Authorization Boundary · Workspace Isolation · Audit/Evidence · Fail-Closed.

## 4. Non-Goals

- Keine numerischen/protokollspezifischen Fehlercodes, kein HTTP-Statuscode-Mapping.
- Kein Idempotency-Key-/Nonce-/Deduplication-Mechanismus, keine Replay-Protection-Technologie.
- Keine Redaction-/Rate-Limit-Technologie, kein Runtime-Code.
- Keine Behauptung implementierter/validierter/zertifizierter Kontrollen.

## 5. Error Boundary

Eine Fehlerantwort ist eine **Aussage über die Anfrageverarbeitung**, kein Nachweis über den Zielzustand. `error response ≠ proof of no side effect`; `transport interruption ≠ operation did not execute`; `timeout ≠ failed operation`.

## 6. Error Classes

`request-contract error · validation error · authentication error · authorization error · scope error · policy conflict · state conflict · resource-not-found · operation-not-supported · compatibility error · rate-or-capacity constraint · dependency failure · integration failure · target rejection · timeout · transport interruption · partial failure · replay suspected · duplicate request · outcome unknown · internal processing failure · verification failure`. Ein Fehler benötigt: error identity/class · request/correlation references · operation reference · consumer-safe summary · retry classification · side-effect uncertainty · affected scope · problem owner · support reference (falls anwendbar) · audit reference. **Keine numerischen/protokollspezifischen Codes ausgewählt.**

## 7. Consumer-Safe Error Information

Consumer-facing Fehler offenbaren **nicht** unnötig: Secrets · Credential-Material · interne Pfade · private Hostnamen · Stack Traces · vollständige interne Queries · nicht autorisierte Ressourcenexistenz · Daten anderer Workspaces · interne Policy-/Security-Details.
```text
consumer-facing error       ≠ internal diagnostic record
internal diagnostic record  ≠ unrestricted disclosure
support correlation reference ≠ secret
```
Keine Redaction-Technologie ausgewählt (Bezug THR-019, THR-020, THR-018).

## 8. Internal Diagnostics

Interne Diagnostik kann mehr Detail enthalten, unterliegt aber Disclosure-/Access-Grenzen und Audit; sie ist nicht frei zugänglich und wird nicht an Consumer geleakt.

## 9. Retry Classification

`not-retryable · retryable-with-same-request · retryable-with-new-authorization · retryable-after-reconciliation · retryable-after-dependency-recovery · consumer-decision-required · unknown`.
```text
server error          ≠ automatically retryable
timeout               ≠ safe retry
transport interruption ≠ operation did not execute
unknown outcome       → no automatic retry
revoked or consumed authorization → no reuse
```
Keine Retry-/Backoff-Strategie ausgewählt.

## 10. Idempotency Concepts

`side-effect-free · naturally idempotent · conditionally idempotent · idempotent within defined scope · non-idempotent · idempotency context · duplicate detection · replay protection · retry · resumption`.
```text
idempotent        ≠ side-effect-free
same request body ≠ same authorised operation
duplicate detected ≠ prior result verified
idempotency context ≠ execution authorization
retry             ≠ replay
resumption        ≠ new operation
```

## 11. Idempotency Context

Konzeptionell gebunden an: API identity · operation identity/version · requesting principal · consumer identity · workspace/environment · target identity · action/scope · request input classification · policy/authorization reference · validity boundary · attempt history · result state · consumption state · audit reference. Eine Idempotency-Bindung darf **nicht**: Autorisierung verlängern · Scope erweitern · anderen Principal übernehmen · Target/Operation austauschen · Unknown Outcome automatisch als Erfolg behandeln. Keine Key-/Hash-/Nonce-/Speicherungstechnologie ausgewählt.

## 12. Duplicate Handling

`identical request repeated · semantically equivalent request · duplicated transport delivery · duplicate asynchronous completion · duplicate partial result`. Duplicate Detection beruht **nicht** allein auf untrusted Timestamps; `duplicate detected ≠ prior result verified`; eine Duplicate Response überschreibt Attempt-/Result-Historie **nicht** (Bezug THR-026).

## 13. Replay Boundary

`replayed old request · replayed approval · replayed execution authorization · repeated offline package`. Ein altes Request-/Authorization-Artefakt autorisiert **nicht** erneut; verbrauchte/widerrufene Authorization nicht wiederverwendbar; Replay-Verdacht bleibt sichtbar/auditierbar (Bezug THR-026, THR-004). Keine Replay-Protection-Technologie ausgewählt.

## 14. Unknown Outcome

```text
outcome unknown ≠ failed ≠ not executed ≠ successful
```
Ursprüngliche Authorization nicht still wiederverwendbar · Versuchshistorie erhalten · Reconciliation erforderlich · Consumer über Unsicherheit informiert · **keine automatische Wiederholung** · neue Authorization nur nach expliziter Governance-Entscheidung · Side-Effect-Unsicherheit sichtbar (Bezug THR-029, THR-031).

## 15. Bulk and Partial Results

`aggregate success ≠ every item verified`; `single item failure ≠ complete bulk failure`; `bulk authorization ≠ unlimited target expansion`; `partial response ≠ complete result`. Per-Target-Autorität bleibt erhalten; Partial-Failure-State sichtbar (aus [Governance §15]).

## 16. Async Operations

Acceptance ≠ Execution; Statusabfrage ≠ Verification; Completion Notification ≠ Verification; Unknown Outcome → Reconciliation (aus [Governance §17]). Duplicate async completion überschreibt Historie nicht.

## 17. Pagination and Continuation

`next page ≠ same underlying state`; `continuation reference ≠ execution authorization`; `stable ordering ≠ snapshot consistency`; `page count ≠ total authoritative count`. Continuation ist scope-/authorization-/expiry-gebunden; darf keine Cross-Workspace-Daten offenlegen (§19, Bezug THR-035).

## 18. Authorization Boundary

API Acceptance ersetzt Policy/Approval/Execution Authorization **nicht** (CO-WP-013). Idempotency-/Continuation-/Retry-Kontexte erzeugen/verlängern **keine** Autorisierung. API/API-Gateway erzeugen keine parallele Autorisierungsautorität.

## 19. Workspace Isolation

Workspace Scope explizit; Cross-Workspace-Zugriff nur mit ausdrücklicher Autorität; Ressourcenexistenz nicht unautorisiert offengelegt (auch nicht über Fehler/Pagination/Enumerate); Read-only ohne stille Write-Eskalation (Bezug THR-035, THR-007).

## 20. Audit and Evidence

Erfasst wie [Governance §20]. Trennung `request ≠ acceptance ≠ execution ≠ success ≠ verification evidence ≠ compliance`. Replay-/Duplicate-Findings, Unknown-Outcome und Error-Disclosure-Ereignisse werden auditiert; Historie nicht umgeschrieben (Bezug THR-016, THR-017).

## 21. Fail-Closed Rules

Keine privilegierte/mutierende Verarbeitung bei: fehlender/unklarer Authorization · `indeterminate`/`conflicted` Acceptance · unbekanntem Outcome · Replay-Verdacht · abgelaufener/widerrufener/konsumierter Authorization · Plan-/Target-/Scope-Mismatch · unklarer Workspace-Grenze · unsicherer Side-Effect-Lage. Im Zweifel wird nicht wiederholt und nicht als Erfolg gewertet.

## 22. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. API availability must not imply authorisation.
2. Request acceptance must not imply execution.
3. Successful response must not imply verified outcome.
4. Error response must not prove absence of side effects.
5. Unknown outcome must block automatic retry and requires reconciliation.
6. Idempotency context must not extend, replace or recreate execution authorization.
7. Duplicate detection must not overwrite attempt or result history.
8. A replayed request or authorization must grant no new authority.
9. Read-only API must not silently gain write authority.
10. Cross-workspace access must remain explicit and scope-bound.
11. Error responses must not disclose secrets or unauthorised resource data.
12. Bulk operations must preserve per-target authority and partial-result visibility.

## 23. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-004, THR-005, THR-006, THR-007, THR-011, THR-016, THR-017, THR-018, THR-019, THR-020, THR-026, THR-029, THR-031, THR-035, THR-036, THR-038. Keine Duplikation, kein Parallelregister.

## 24. Technology Boundary

Nicht ausgewählt/implementiert: Error-Code/Statuscode, Idempotency-Key/Nonce/Deduplication, Replay-Protection, Redaction, Rate-Limit, Transport/API-Style, Runtime-Code.

## 25. Compatibility

Konsistent mit [API Governance](../architecture/API_GOVERNANCE_AND_OPERATION_MODEL.md), [API Versioning](../architecture/API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md), [Integration Trust/Failure](INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Public Neutrality/Disclosure](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md). Konkretisiert DEC-G-06 (keine Secrets in Logs/Exports), DEC-P-04.

## 26. Open Questions

- Konkreter Idempotency-/Replay-Schutz (spätere ADR).
- Consumer-safe Error-Taxonomie im Detail.
- Rate-Limit-/Quota-Modell (THR-036).

## 27. Next Decision

Event/Audit Correlation (CO-WP-018) und Telemetry (CO-WP-019) referenzieren Error-/Result-Evidenz. Secrets/Redaction-Detail (CO-WP-024/025). Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
