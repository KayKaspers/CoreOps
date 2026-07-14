# CoreOps – Integration Capability and Operation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation integration capability and operation model
> Implementation Status: Not implemented
> Capability Negotiation: Not implemented
> Protocol Binding: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-014 (docs-only / integration-architecture and security foundation)

## 1. Status

Technologieunabhängiges Modell für **Integration Capabilities und Operationen**: wie Capabilities deklariert, erkannt, erlaubt und validiert werden und wie Operationen nach Privileg klassifiziert sind. Companion zu [COREOPS_INTEGRATION_CONTRACT_V0_1.md](COREOPS_INTEGRATION_CONTRACT_V0_1.md) und [INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md). Keine Capability-Negotiation-Technologie, kein Protokoll, kein Schema.

## 2. Purpose

Capabilities sind der gefährlichste Punkt einer Integration: eine deklarierte Fähigkeit darf niemals still zu einer erlaubten oder implementierten werden. Dieses Modell trennt sechs Capability-Dimensionen strikt und ordnet jede Operation einer Privileg- und Autorisierungsklasse zu, damit `advertised ≠ detected ≠ permitted ≠ implemented ≠ supported ≠ validated`.

## 3. Scope

Capability-Konzepte/-Dimensionen · Capability Declaration · Operation Classes und Privilege Classification · Read-only/Write/Execution/Deployment/Evidence-Operationen · Capability Detection/Permission/Validation · Operation Preconditions · Result Classification · Compatibility · Offline · Audit/Evidence.

## 4. Non-Goals

- Keine Capability-Negotiation-/Handshake-Technologie.
- Kein Protokoll/Schema/SDK/Adapter.
- Keine konkrete Privilege-Enforcement-Implementierung.
- Keine Behauptung implementierter/validierter Capabilities.

## 5. Capability Concepts

Eine Capability beschreibt, welche Operationsklasse eine Integration auf welcher Zielressourcenklasse in welcher Privilegstufe **beanspruchen** kann. Beanspruchung ist keine Erlaubnis. Konsistent mit der [Capability Matrix Spec](../project-system/CAPABILITY_MATRIX_SPEC.md) (dortige mehrdimensionale Status).

## 6. Capability Dimensions

Sechs **getrennte** Dimensionen — keine wird still aus einer anderen abgeleitet:
```text
advertised  — von der Integration deklariert
detected    — von CoreOps beobachtet/erkannt
permitted   — durch Policy/Authorization erlaubt
implemented — tatsächlich vorhandene Funktion
supported   — im CoreOps-Sinne unterstützt (Supportgrenze)
validated   — durch Evidenz nachgewiesen
```
```text
advertised  ≠ implemented ≠ validated
detected    ≠ permitted
supported   ≠ validated
```
`evidence capability ≠ evidence available ≠ requirement satisfied` (aus CO-WP-004E).

## 7. Capability Declaration

Enthält mindestens: capability identity · capability class · contract version · declaring integration · target resource class · supported operation classes · read/write classification · execution privilege level · required authorization class · offline availability · known limitations · version/compatibility constraints · validation state · evidence reference. **Ein Adapter darf sich keine Permission selbst erteilen, indem er eine Capability deklariert** (Declaration ⊂ advertised, nicht permitted).

## 8. Operation Classes

`discover · query · read · observe · enumerate · export · register · configure · create · update · delete · execute · deploy · start · stop · restart · verify · collect-evidence`. Jede Klasse trägt eine eindeutige Privilege- und Authorization-Einordnung (§9).

## 9. Privilege Classification

| Privilegstufe | Operationsklassen | Authorization-Anforderung |
|---|---|---|
| discovery-only | discover | Enrollment-Scope; keine Write-Autorität |
| read/observe | query, read, observe, enumerate | Read-Scope; read-first |
| export/evidence | export, collect-evidence | eigener Disclosure-/Authorization-Scope |
| controlled write | register, configure, create, update, delete | Policy + ggf. Approval |
| privileged execution | execute, start, stop, restart | Execution Authorization (CO-WP-013) |
| deployment | deploy | Policy + Approval + Execution Authorization |
| verify | verify | separater Nachweis; request verify ≠ verified result |

```text
read      ≠ export
observe   ≠ configure
configure ≠ execute
execute   ≠ deploy
capability to execute ≠ authorization to execute
```

## 10. Read-only Operations

`discover/query/read/observe/enumerate` verändern kein Ziel. Read-only ist der Default nach Enrollment; Detection einer Write-Capability aktiviert **keine** Write-Autorität.

## 11. Write Operations

`register/configure/create/update/delete` benötigen explizite Policy und — wo gefordert — Approval. Scope-/Target-Bindung nach [Execution Authorization §10](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md); Adapter erweitern Targets nicht.

## 12. Execution Operations

`execute/start/stop/restart` benötigen Execution Authorization (action-/target-/scope-/plan-/time-bound). Agent/Adapter erweitern Action-Scope nicht.

## 13. Deployment Operations

`deploy` benötigt Policy + Approval + Execution Authorization und — wo relevant — Artefakt-Provenance (CO-WP-022). `deploy ≠ execute` (breiterer Impact).

## 14. Evidence and Export Operations

`export/collect-evidence` haben einen eigenen Disclosure-/Authorization-Scope; Redaction/Datenminimierung gilt (Bezug THR-018, THR-020, THR-025). Evidence-Export ist keine Ausführungsbefugnis.

## 15. Capability Detection

CoreOps kann Capabilities beobachten (`detected`). Detection ist **read-only** und erzeugt **keine** Permission. `detected ≠ permitted`; erkannte Write-/Execution-Fähigkeit bleibt bis zu expliziter Policy/Authorization ungenutzt.

## 16. Capability Permission

`permitted` entsteht ausschließlich durch Policy-Entscheidung/Authorization (CO-WP-013), nicht durch Declaration oder Detection. Fehlende/unklare Permission ⇒ keine privilegierte Operation (fail-closed).

## 17. Capability Validation

`validated` benötigt Evidenz; ohne Nachweis bleibt eine Capability höchstens `advertised`/`detected`/`permitted`/`implemented`, aber nicht `validated`. Keine stille Ableitung.

## 18. Operation Preconditions

Vor privilegierter Operation mindestens: gültige/erlaubte Capability · anwendbare Policy-Version · gültige Approval (wo gefordert) · aktive, nicht abgelaufene/widerrufene/konsumierte Execution Authorization · Plan-/Target-/Scope-Match · ausreichend bestimmter Effective State · Provenance verfügbar · Adapter/Agent-Trust akzeptabel · Offline-Paket gültig · Audit-Start möglich. Bei Unklarheit fail-closed (Companion 3 §22, konsistent mit CO-WP-013 Pre-Execution Guards).

## 19. Result Classification

Wie Contract §15: `succeeded/failed/partially-succeeded/not-executed/outcome-unknown/cancelled-*/verification-failed/indeterminate`. `outcome-unknown` bleibt sichtbar; `execution-completed ≠ successful`; `successful ≠ verified`.

## 20. Compatibility

Capability-Version, Contract-Version, Extension-Version getrennt (Contract §24). Unbekannte Pflicht-Capability-Semantik wird nicht still ignoriert (fail-safe).

## 21. Offline Considerations

Capabilities können offline eingeschränkt sein (`offline availability`). Offline erkannte/erlaubte Capabilities benötigen Provenance/Integrität/Target-Binding/explizite Aktivierung (Contract §22). Keine automatische konfliktfreie Synchronisation.

## 22. Audit and Evidence

Erfasst: capability declaration · detection · permission decision · validation state/evidence · operation attempt/result. Trennung `advertised/detected/permitted/implemented/supported/validated` bleibt im Audit sichtbar; `evidence capability ≠ available ≠ satisfied`.

## 23. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Advertised capability must not imply detected, permitted, implemented or validated capability.
2. Detected capability must not become permitted automatically.
3. An adapter must not self-grant permission by declaring a capability.
4. Read/observe operations must not gain write or execution authority.
5. Privileged operations require explicit policy and applicable authorization.
6. `capability to execute` is not `authorization to execute`.
7. Validation requires evidence; it is not inferred.
8. Unknown required capability semantics must fail safely, not silently.

## 24. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-005, THR-006, THR-007, THR-010, THR-011, THR-012, THR-014, THR-018, THR-020, THR-024, THR-029, THR-034. Keine Duplikation, kein Parallelregister.

## 25. Technology Boundary

Nicht ausgewählt/implementiert: Capability-Negotiation/Handshake, Protokoll, Schema, SDK, Adapter, Privilege-Enforcement-Runtime.

## 26. Open Questions

- Minimale Pflichtfelder je Operation Class (spätere Iteration).
- Formales Verhältnis `supported` (CoreOps-Supportgrenze) ↔ `validated` je Integration Class.

## 27. Next Decision

Domain Pack Governance (CO-WP-015) und Deployment Control Plane (CO-WP-021) konkretisieren Capability-/Operation-Aspekte. Negotiation-/Enforcement-Technologie bleibt einer späteren ADR-Runde vorbehalten.
