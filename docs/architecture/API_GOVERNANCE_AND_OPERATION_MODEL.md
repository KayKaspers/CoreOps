# CoreOps – API Governance and Operation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation API-governance and operation model
> Implementation Status: Not implemented
> Transport Binding: Not selected
> API Style: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-017 (docs-only / API architecture and security foundation)

## 1. Status

Technologieunabhängiges Modell für **API-Surface, API-Identität, API-Klassen, Operationen und Side-Effect-Klassen**. Konkretisiert den [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md) für programmatische Schnittstellen. Companion zu [API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md](API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md) und [API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md](../security/API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md). Es wählt **kein** Transport, keinen API-Style, kein Schema, keine Statuscodes.

## 2. Purpose

Eine API macht Operationen programmatisch verfügbar — Verfügbarkeit ist aber keine Autorisierung, und eine Antwort ist kein Nachweis. Dieses Modell legt fest, wie APIs identifiziert, klassifiziert und operationsweise nach Side-Effect/Privileg eingeordnet werden, damit `API availability ≠ authorization`, `request accepted ≠ executed`, `successful response ≠ verified outcome` und `error response ≠ proof of no side effect`.

## 3. Scope

API-Begriffe · API-Klassen · API Identity · API Lifecycle · Producers/Consumers · Operation Classes · Side-Effect-Klassifikation · Request/Response Model · Acceptance/Result-Semantik · Bulk · Pagination/Continuation · Async · Policy-/Authorization-Bindung · Workspace/Scope · Audit/Evidence. Versionierung/Compatibility (Companion 2), Error/Idempotency/Replay (Companion 3).

## 4. Non-Goals

- Keine REST-/GraphQL-/gRPC-/WebSocket-Auswahl, kein HTTP-Statuscode-Mapping.
- Keine OpenAPI-/AsyncAPI-Spezifikation, kein JSON-/YAML-/XML-/Protobuf-Schema.
- Kein Endpoint-/URL-Design, keine API-Gateway-/Rate-Limit-Technologie.
- Keine Idempotency-Key-/Nonce-/Deduplication-Technologie, kein Runtime-Code.
- Keine Behauptung implementierter/validierter APIs.

## 5. Concepts

Begriffe (mindestens): API · API surface · API identity · API class · API producer · API consumer · API operation · operation version · request · response · acceptance · authorization · execution · result · error · problem detail · warning · partial result · bulk operation · asynchronous operation · continuation · pagination · idempotency · retry · replay · duplicate · correlation · deprecation · retirement.

**Grundregeln:**
```text
API                 ≠ transport protocol
API operation       ≠ Integration Contract capability
request received    ≠ request accepted
request accepted    ≠ authorised
authorised          ≠ executed
successful response ≠ desired outcome verified
error response      ≠ proof of no side effect
duplicate response  ≠ duplicate execution
```

## 6. API Classes

Je Klasse: Purpose · Owner · Consumers · Producer · Scope · Typical operation classes · Privilege level · Identity requirement · Authorization requirement · Versioning expectation · Compatibility expectation · Disclosure considerations · Offline considerations · Audit requirement · Threat references. **Keine Netzwerk-/Transportgrenze ausgewählt.**

| API-Klasse | Zweck (kurz) | Priv.-Level | Auth-Anforderung | Threat refs |
|---|---|---|---|---|
| public external API | öffentlich erreichbar (falls aktiviert) | niedrig/read | mind. authN wo nötig | THR-036, THR-039 |
| authenticated external API | authentifizierte externe Nutzung | mittel | authN + Policy | THR-038, THR-001 |
| administrative API | privilegierte Verwaltung | hoch | Policy + Approval | THR-002, THR-037 |
| workspace-scoped API | auf Workspace begrenzt | mittel | RBAC/scope | THR-035 |
| internal module API | Modul-zu-Modul | intern | Modulgrenzen | THR-034 |
| agent and relay API | Agent/Relay-Vermittlung | mittel | Machine Identity | THR-008, THR-009 |
| adapter and integration API | Adapter/Integration | mittel | Integration Trust | THR-010, THR-011 |
| automation-client API | programmatische Clients | mittel | Client-authN/scope | THR-038 |
| read-only observation API | reine Beobachtung | niedrig | read scope | THR-012, THR-013 |
| write and configuration API | Konfiguration/Write | hoch | Policy + Approval | THR-007 |
| execution-control API | Ausführungssteuerung | hoch | Execution Authorization | THR-005, THR-026 |
| evidence and export API | Evidenz/Export | hoch | Disclosure/Authorization | THR-018, THR-020 |
| offline import/export contract surface | Offline-Austausch | hoch | Provenance/Aktivierung | THR-024 |

## 7. API Identity

Eine API-Surface benötigt mindestens: stable API identity · canonical name · owning module · API class · producer identity · consumer classes · lifecycle state · API version · operation versions · schema/data-contract references · Integration Contract version · Domain-Pack references (falls anwendbar) · scope · support status · compatibility status · deprecation state · security-review state · validation state · audit reference.
```text
display name    ≠ API identity
route or URL    ≠ API identity
transport binding ≠ API identity
API version     ≠ CoreOps product version
```
**Keine vollständige API-ID-/Operationsliste.**

## 8. API Lifecycle

Konzeptionelle Statuswerte:
```text
proposed → draft → preview → active → restricted → deprecated → security-maintenance-only → retired → archived
```
- `preview` erzeugt keinen stabilen Compatibility-/Support-Claim.
- `active` ≠ automatisch supported/validated.
- `restricted` bleibt scope-/consumergebunden.
- `deprecated` benötigt Migrations-/Retirement-Grenzen (Companion 2 §14).
- `retired` erhält keine neuen Consumers; historische API-/Operation-IDs bleiben erhalten.

## 9. Producers and Consumers

Producer (stellt API bereit, = owning module) und Consumer (nutzt) sind getrennt. Consumer-Klassen (extern/intern/agent/adapter/automation) tragen unterschiedliche Identity-/Authorization-Anforderungen. Ein Consumer erhält keine zusätzliche Autorität durch API-Version oder Consumer-Typ (§19).

## 10. Operation Classes

`query · read · observe · enumerate · export · create · configure · update · delete · execute · deploy · start · stop · restart · verify · collect-evidence · administrative-management`. Je Operation klassifiziert: read-only/mutating · side-effect-free expectation · idempotency expectation · privilege level · authorization requirement · approval requirement · target scope · result semantics · retry suitability · audit requirement.

## 11. Side-Effect Classification

```text
read      ≠ export
read-only ≠ disclosure-free
configure ≠ execute
execute   ≠ deploy
request verify ≠ verified result
operation available ≠ operation authorised
```
Read-only-Operationen ändern kein Ziel, können aber Offenlegung bedeuten (Export/Enumerate → Disclosure-Scope). Mutating/Execution/Deploy-Operationen benötigen Policy/Approval/Execution Authorization (§18).

## 12. Request Model

Konzeptionell (keine Serialisierung): API identity · API/operation version · request identity · correlation identity · requesting principal · accountable human owner (wo nötig) · consumer identity · workspace/environment · target identity/class · operation identity · requested scope · input classification · policy-decision reference · approval reference (falls anwendbar) · execution-authorization reference (falls anwendbar) · idempotency context (falls anwendbar) · validity/timeout boundary · retry/replay context · provenance reference · audit reference.

## 13. Response Model

Mindestens: request identity · correlation identity · operation identity · attempt identity · acceptance state · execution state (falls anwendbar) · reported outcome · result classification · partial-result state · warning state · error/problem reference · continuation state · verification state · provenance reference · freshness state · audit reference.
```text
response delivered  ≠ operation completed
operation completed ≠ successful
successful          ≠ verified
empty response      ≠ no result
warning             ≠ success ≠ failure
```

## 14. Acceptance and Result Semantics

Acceptance States: `received · accepted-for-evaluation · rejected · authorization-required · not-supported · indeterminate · conflicted`. Result States: `succeeded · failed · partially-succeeded · not-executed · outcome-unknown · cancelled-before-execution · cancelled-after-partial-execution · verification-failed · indeterminate`. Konsistent mit [Integration Contract §13/§15](COREOPS_INTEGRATION_CONTRACT_V0_1.md); **kein Parallelmodell mit widersprüchlichen Bedeutungen**.

## 15. Bulk Operations

Bulk benötigt: bulk request identity · individual target references · individual operation identities/suboperations · per-item authorization expectation · per-item result · aggregate result · partial-failure state · continuation/resume state · rollback/recovery expectation · audit reference.
```text
aggregate success  ≠ every item verified
single item failure ≠ complete bulk failure
bulk authorization ≠ unlimited target expansion
partial response   ≠ complete result
```
Keine Batch-/Transaktionstechnologie ausgewählt.

## 16. Pagination and Continuation

Berücksichtigt: query/operation identity · consumer/scope · ordering expectation · continuation state · freshness/snapshot expectation · filter state · authorization state · expiry · duplicate/skipped-item risk · audit reference (wo nötig).
```text
next page             ≠ same underlying state automatically
continuation reference ≠ execution authorization
stable ordering       ≠ snapshot consistency
page count            ≠ total authoritative count
```
Keine Cursor-/Offset-/Token-/Snapshot-Technologie ausgewählt.

## 17. Async and Long-Running Operations

Request-Verbindung ≠ Operation-Lifecycle; Acceptance ≠ Execution; Statusabfrage ≠ Verification; Zwischenstatus nicht final; Verbindungsverlust beendet Operation nicht; Completion Notification ≠ Verification; Cancellation braucht bestätigten Outcome; Unknown Outcome braucht Reconciliation. Keine Polling-/Callback-/Streaming-/Queue-/Worker-Technologie ausgewählt.

## 18. Policy and Authorization Binding

Write-/Export-/Execution-Operationen referenzieren: policy version · policy decision · approval requirement · approval decision · execution authorization · action · target · scope · plan/workflow reference · expiry · revocation state · consumption state. **API Acceptance ersetzt diese Governance nicht; API/API-Gateway erzeugen keine parallele Autorisierungsautorität** (CO-WP-013 autoritativ).

## 19. Workspace and Scope Boundary

- Workspace Scope bleibt explizit; Cross-Workspace Requests benötigen ausdrückliche Autorität.
- Ressourcenexistenz wird nicht unautorisiert offengelegt (§Companion 3).
- Administrative API ist nicht automatisch global; Internal API nicht automatisch vertrauenswürdig.
- Agent-/Adapter-API erweitert Scope nicht selbst; API-Version/Consumer-Typ erzeugt keine zusätzliche Autorität; Read-only bleibt ohne stille Write-Eskalation (Bezug THR-007, THR-035).

## 20. Audit and Evidence

Erfasst: API/operation identity · versions · consumer · principal · workspace/scope · request · acceptance · policy/authorization references · idempotency context reference · attempt · response · error · partial result · retry · replay/duplicate finding · pagination/continuation · async status · result · verification · deprecation use · closure. Trennung:
```text
request evidence ≠ acceptance evidence ≠ execution evidence ≠ success evidence ≠ verification evidence ≠ compliance
```

## 21. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. API availability must not imply authorisation.
2. Request acceptance must not imply execution.
3. Successful response must not imply verified outcome.
4. Error response must not prove absence of side effects.
5. Read-only API must not silently gain write authority.
6. Cross-workspace access must remain explicit and scope-bound.
7. Bulk operations must preserve per-target authority and partial results.
8. API version or consumer type must not grant additional authority.
9. An API must not create a parallel authorization authority.

## 22. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-001, THR-002, THR-005, THR-007, THR-008, THR-009, THR-010, THR-011, THR-012, THR-013, THR-018, THR-020, THR-024, THR-026, THR-034, THR-035, THR-036, THR-037, THR-038, THR-039.

## 23. Technology Boundary

Nicht ausgewählt/implementiert: Transport (REST/GraphQL/gRPC/WebSocket), API-Style, Statuscodes, OpenAPI/AsyncAPI, Schema, Endpoint/URL, API-Gateway, Rate-Limit, Idempotency-/Replay-Mechanismus, Runtime-Code.

## 24. Compatibility

Konsistent mit [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md)/[Capability/Operation](INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md), [Module Catalog](COREOPS_MODULE_CATALOG.md), [SoT/Provenance](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Schema/Migration](SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [Policy/Approval/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [RBAC](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert DEC-G-07, DEC-P-04.

## 25. Open Questions

- Transport-/API-Style-Auswahl (spätere ADR).
- Verbindliche Versionsnotation (DEC-O-02, Companion 2).
- Rate-Limit-/Quota-Modell (mit THR-036, später).

## 26. Next Decision

Companion 2 (Versioning/Compatibility) und Companion 3 (Error/Idempotency/Replay) tragen Detailregeln. Event/Audit Correlation (CO-WP-018) und Telemetry (CO-WP-019) referenzieren API-Operationen. Transport-/Style-Wahl bleibt einer späteren ADR-Runde vorbehalten.
