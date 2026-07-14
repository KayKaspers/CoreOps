# CoreOps – Integration Trust, Failure and Recovery Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation integration trust and failure policy
> Implementation Status: Not implemented
> Trust Mechanism: Not selected
> Replay Protection: Not selected
> Recovery Automation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-014 (docs-only / integration-architecture and security foundation)

## 1. Status

Technologieunabhängige Policy für **Vertrauen, Fehlerklassifikation und Wiederherstellung** an der Integrationsgrenze. Companion zu [COREOPS_INTEGRATION_CONTRACT_V0_1.md](../architecture/COREOPS_INTEGRATION_CONTRACT_V0_1.md) und [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](../architecture/INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md). Kein Trust-Mechanismus, kein Replay-Schutz, keine Recovery-Automation ausgewählt/implementiert.

## 2. Purpose

Integrationen sind eine primäre Angriffs- und Fehlerfläche. Diese Policy legt fest, wie wenig einer Integrationsantwort vertraut wird, wie Fehler klassifiziert werden, warum `transport failure ≠ no side effect` und `unknown outcome ≠ safe non-execution`, und wie Retry/Rollback/Recovery fail-closed und verifikationsgebunden bleiben.

## 3. Scope

Integration Trust Boundary · Identity/Enrollment · Capability Trust · Policy/Authorization · Request Acceptance · Execution Boundary · Result Trust · Failure Classification · Unknown Outcome · Retry/Replay · Duplicate Handling · Partial Failure · Cancellation · Rollback/Recovery · Offline Integration · Provenance · Audit/Evidence · Fail-Closed Rules.

## 4. Non-Goals

- Kein Trust-/Attestierungs-Mechanismus, keine Kryptografie-/Signaturauswahl.
- Kein Replay-Schutz (Nonce/Idempotency), keine Recovery-Automation.
- Kein Protokoll/Schema/SDK/Adapter, kein Runtime-Code.
- Keine Behauptung implementierter/validierter/zertifizierter Kontrollen.

## 5. Integration Trust Boundary

Jede Integration liegt jenseits einer Vertrauensgrenze (TB-02/TB-03/TB-06 aus dem [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md)). `enrolled ≠ trusted`; `trusted ≠ globally authorised`. Vertrauen ist scope-/capability-gebunden und jederzeit widerrufbar.

## 6. Identity and Enrollment

Integrationsidentität, Enrollment und Trust folgen [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md) und [Machine Enrollment](MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md): `discovery ≠ enrollment`, `registration ≠ trust`, `identity ≠ credential`. Kein Parallelmodell.

## 7. Capability Trust

`advertised/detected` Capabilities sind **nicht** vertrauenswürdig-erlaubt; nur `permitted` (durch Policy) und — für Erfolgsaussagen — `validated` (durch Evidenz) tragen. Ein Adapter erhält keine Autorität durch Selbstdeklaration.

## 8. Policy and Authorization

Write/Execution über Integrationen benötigen anwendbare Policy-Version, Approval (wo gefordert) und Execution Authorization (CO-WP-013). **Contract acceptance ersetzt keine Authorization.** Pre-Execution Guards bleiben autoritativ.

## 9. Request Acceptance

`received ≠ accepted`; `accepted-for-evaluation ≠ authorised`; `authorised ≠ execution-started`. Ein `rejected` beweist **nicht**, dass keine frühere target-seitige Aktion stattfand.

## 10. Execution Boundary

Ausführung nur über vorgesehene Execution-/Adapter-/Agentengrenzen (CO-WP-013 §12). Adapter self-grant, Agent-Scope-Expansion, Evidence-als-Trigger, Notification-als-Command und sofortige Offline-Import-Ausführung sind unzulässig.

## 11. Result Trust

Eine Integrationsantwort ist eine **Beobachtung**, kein Nachweis. `integration response ≠ authoritative state`; `execution-completed ≠ successful`; `successful ≠ verified`. Erfolgsaussagen benötigen unabhängige, frische Verifikation (§18).

## 12. Failure Classification

Wie Contract §16: contract/validation/authorization error · capability mismatch · target rejection · transport interruption · integration/target-side/partial failure · timeout · stale/duplicate response · replay suspected · outcome unknown · verification failure.
```text
transport failure ≠ target action did not occur
timeout           ≠ failed operation
error response    ≠ no side effect
partial failure   ≠ complete failure ≠ complete success
```

## 13. Unknown Outcome

`outcome-unknown` bleibt **explizit sichtbar** und wird weder als Erfolg noch als sichere Nichtausführung behandelt. Er erfordert Reconciliation (§18) und **blockiert automatische Wiederholung** (§14).

## 14. Retry and Replay

Retry benötigt eine begründete Governance-Entscheidung; **Unknown Outcome verhindert automatischen Retry** (Bezug THR-029, THR-031). Replay einer Request/Authorization begründet keine neue Autorität; verbrauchte/widerrufene Authorization ist nicht wiederverwendbar (Bezug THR-026). Kein Nonce-/Idempotency-Mechanismus ausgewählt.

## 15. Duplicate Handling

Duplicate Detection beruht **nicht** allein auf untrusted Timestamps. Ein Duplicate darf keine doppelte Target-Aktion erzeugen; ursprüngliche Versuchs-/Ergebnishistorie bleibt erhalten. Keine Deduplication-Technologie ausgewählt.

## 16. Partial Failure

Teilfehler bleiben **sichtbar** (`partial failure ≠ complete success`) und hinterlassen einen erfassten, nicht als vollständig markierten Zustand (Bezug THR-031). Betroffene Ressourcen werden im Result referenziert.

## 17. Cancellation

`cancel requested ≠ cancelled`; Cancellation benötigt eine eigene bestätigte Statusänderung. `cancelled-after-partial-execution` bleibt sichtbar und ist kein sauberer Abbruch.

## 18. Rollback and Recovery

```text
rollback completed ≠ original state restored
recovery completed ≠ desired outcome verified
```
Rollback-/Recovery-Claims benötigen **Verifikationsevidenz** über eine möglichst unabhängige, frische Quelle (konsistent mit [Safe Remediation §12](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md), Bezug THR-032). Keine Recovery-Automation implementiert.

## 19. Offline Integration

Offline-Pakete benötigen target-environment/contract-version/operation/scope-Binding, Provenance, Integrität, Import-Quarantäne, explizite lokale Aktivierung, Result-Reconciliation und Revocation-Distribution-Behandlung (Contract §22, Bezug THR-024). Fail-closed bei unklarer Autorität; keine automatische konfliktfreie Synchronisation; keine Klassifiziertnetz-Eignung behauptet.

## 20. Provenance

Integrationsergebnisse erhalten Source-/Provenance-Informationen; `adapter field ≠ canonical field`; `recently received ≠ recently observed`; `successful parsing ≠ validated provenance`. Die SoT-/Provenance-Dokumente (CO-WP-011) bleiben autoritativ; kein Parallelmodell.

## 21. Audit and Evidence

Erfasst wie Contract §26. Trennung `request ≠ acceptance ≠ execution ≠ success ≠ verification evidence ≠ compliance`. Audit-Records sind keine Ausführungsauslöser (Bezug THR-016, THR-017); Secret-Redaction gilt (Bezug THR-019, THR-020).

## 22. Fail-Closed Rules

Bei sicherheitsrelevanter Unklarheit **keine** privilegierte Integration: unklare/fehlende Permission · `indeterminate`/`conflicted` Policy · fehlende/abgelaufene/widerrufene/konsumierte Authorization · Plan-/Target-/Scope-Mismatch · unzureichend bestimmter Effective State · fehlende Provenance · unklarer Adapter-/Agent-Trust · ungültiges Offline-Paket · unbekannter Outcome · Contract-Version-Mismatch mit unbekannter Pflichtsemantik · unbekannte sicherheitsrelevante Extension.

## 23. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Enrolled integration is not trusted; trusted is not globally authorised.
2. Advertised/detected capability must not imply permitted or validated capability.
3. Request acceptance must not imply execution.
4. Transport failure must not be treated as proof of no target-side action.
5. Execution completion must not imply success; success must not imply verification.
6. Unknown outcome must remain explicit and must not trigger automatic retry.
7. Consumed or revoked authorization must not be reused.
8. Adapter or agent must not expand target, action or scope.
9. Integration results must not inherit authoritative state automatically.
10. Rollback/recovery claims require verification evidence.
11. Offline integration requires provenance, integrity, target binding and explicit activation.
12. Contract extensions must not bypass core security invariants.

## 24. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-005, THR-007, THR-008, THR-009, THR-010, THR-011, THR-012, THR-016, THR-017, THR-019, THR-020, THR-023, THR-024, THR-026, THR-028, THR-029, THR-031, THR-032, THR-034, THR-039. Keine Duplikation, kein Parallelregister.

## 25. Technology Boundary

Nicht ausgewählt/implementiert: Trust-/Attestierungs-Mechanismus, Replay-Schutz, Recovery-Automation, Protokoll/Schema/SDK/Adapter, Kryptografie/Signatur.

## 26. Compatibility

Konsistent mit Contract/Capability-Companion, [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md), [Trust/Deployment/Execution Boundaries](TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md), [Machine Identity/Enrollment](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Offline Reconciliation](OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md). Konkretisiert DEC-P-04, DEC-G-04, DEC-G-07.

## 27. Open Questions

- Konkreter Replay-/Duplicate-Schutz je Operation Class (spätere ADR).
- Unabhängige Verifikationsquellen je Integration Class.
- Reconciliation-Ablauf nach Unknown Outcome im Detail.

## 28. Next Decision

Artifact Trust/Provenance (CO-WP-022), Restricted/Air-Gapped Operation (CO-WP-023) und Self-Protection/Recovery Mode (CO-WP-026) konkretisieren Trust-/Failure-/Recovery-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
