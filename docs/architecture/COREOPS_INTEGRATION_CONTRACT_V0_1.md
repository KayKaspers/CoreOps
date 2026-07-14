# CoreOps Integration Contract v0.1

> Contract Name: CoreOps Integration Contract
> Contract Version: 0.1
> Document Status: Implemented, pending Nova review
> Contract Maturity: Foundation draft contract
> Implementation Status: Not implemented
> Protocol Binding: Not selected
> Schema Binding: Not selected
> SDK Status: Not implemented
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-014 (docs-only / integration-architecture and security foundation)

## 1. Status

Erster technologie- und herstellerunabhängiger CoreOps Integration Contract. Er beschreibt die **konzeptionellen** Grenzen zwischen CoreOps und Managed Resources, Agenten/Relays, Adaptern, externen Plattformen und Systemklassen. Companion zu [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md) und [INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md). Er implementiert **keine** API, kein SDK, keinen Adapter und kein Protokoll. **`Contract Version 0.1` ist nicht gleichzusetzen mit einer CoreOps-Produktversion oder einem Release.**

## 2. Purpose

CoreOps integriert eine Vielzahl heterogener Systeme, ohne an einen Hersteller oder ein Protokoll gebunden zu sein. Dieser Contract legt fest, wie eine Integration sich identifiziert, welche Capabilities sie deklariert, welche Operationen sie in welcher Privileg-/Autorisierungsklasse anfragen kann und wie Request, Acceptance, Execution, Result, Failure und Verifikation streng getrennt bleiben — damit `request accepted ≠ executed ≠ successful ≠ verified` und eine Integrationsantwort niemals still zu autoritativem CoreOps-State wird.

## 3. Scope

Integrationsidentität · Contract-Version/Compatibility · Integration Classes · Capability Declaration · Operation Classes · Read/Write/Execution-Boundaries · Request-/Response-/Acceptance-Semantik · Operation Lifecycle · Result-/Failure-Semantik · Async/Long-Running · Retry/Replay/Duplicate · Cancellation/Rollback/Recovery · Offline Integration · Provenance/Field Authority · Policy-/Execution-Authorization-Bindung · Audit/Evidence · Extension-/Vendor-Namespace.

## 4. Non-Goals

- Keine REST-/GraphQL-/gRPC-/WebSocket-Auswahl, keine OpenAPI-/AsyncAPI-Spezifikation.
- Kein JSON-/YAML-/Protobuf-/XML-Schema, kein Message-Broker, kein Event-Bus.
- Keine SDK-/Plugin-Implementierung, keine konkrete Adapterarchitektur.
- Keine Port-/Endpoint-Definition, kein Datenbankschema.
- Keine konkrete Authentisierung/Kryptografie, kein Runtime-Code.
- Keine Behauptung implementierter/validierter/zertifizierter Integrationen.

## 5. Contract Version

- **Contract Version 0.1** — Foundation-Draft; getrennt von Implementation-, Integration-, Adapter- und Managed-System-Version (§24).
- **Contract Version ≠ CoreOps-Produktversion/Release.** Eine gleiche Contract-Version garantiert keine identischen Capabilities; eine neuere Version ist nicht automatisch kompatibel (§24).
- Protocol Binding, Schema Binding, SDK: **not selected / not implemented.**

## 6. Concepts

Begriffe (mindestens): integration · integration instance · integration class · adapter · agent · relay · managed resource · external system · integration principal · contract version · capability · advertised/detected/permitted/supported/implemented/validated capability · operation · request · acceptance · authorization · execution · result · observation · event · error · failure · retry · replay · duplicate · correlation · cancellation.

**Grundregeln:**
```text
integration            ≠ adapter
adapter                ≠ managed resource
advertised capability  ≠ implemented capability
implemented capability ≠ validated capability
detected capability    ≠ permitted capability
request accepted       ≠ operation executed
operation executed     ≠ operation successful
operation successful   ≠ desired outcome verified
error response         ≠ proof that no target-side action occurred
```

## 7. Integration Classes

Herstellerneutral; je Klasse: Purpose · Typical managed resources · Typical operation classes · Expected authority · Read-only possibility · Write/execution possibility · Offline considerations · Identity requirement · Provenance expectation · Audit requirement · Threat references. Keine Anbieterabhängigkeit ausgewählt.

| Klasse | Purpose (kurz) | Typ. Operationen | Read-only mögl. | Write/Exec mögl. | Threat refs |
|---|---|---|---|---|---|
| virtualisation platform | VM-Lebenszyklus beobachten/steuern | discover/query/observe/start/stop/execute | ja | ja (autorisiert) | THR-005, THR-034 |
| container platform | Container/Workloads | observe/deploy/start/stop | ja | ja (autorisiert) | THR-022, THR-034 |
| orchestrator | Cluster-/Scheduling-Zustand | observe/query/configure | ja | ja (autorisiert) | THR-005, THR-034 |
| operating system | Host-Zustand/-Konfiguration | read/observe/configure/execute | ja | ja (autorisiert) | THR-006, THR-034 |
| bare-metal management | Hardware/BMC | discover/observe/restart | ja | ja (autorisiert) | THR-034, THR-039 |
| network management | Netzobjekte/Konfiguration | read/observe/configure | ja | ja (autorisiert) | THR-039, THR-034 |
| printer / print-management | Druck-/Geräteverwaltung | discover/observe/configure | ja | ja (autorisiert) | THR-025, THR-034 |
| identity or directory system | Identitäts-/Verzeichnisdaten | read/observe/query | ja | selten (autorisiert) | THR-001, THR-035 |
| monitoring/observability | Telemetrie/Health | observe/query/collect-evidence | ja | nein/selten | THR-012, THR-013 |
| ITSM/ticketing | Tickets/Changes | read/create/update | ja | ja (autorisiert) | THR-030, THR-018 |
| source-control/CI/CD | Pipelines/Artefakte | read/observe/deploy | ja | ja (autorisiert) | THR-022, THR-023 |
| artifact/package repository | Artefakte/Pakete | read/export/deploy | ja | ja (autorisiert) | THR-021, THR-023 |
| backup/recovery | Backups/Restore | observe/execute/verify | ja | ja (autorisiert) | THR-033, THR-032 |
| notification system | Benachrichtigungen | read/create | ja | ja (autorisiert) | THR-018 |
| cloud/external service | externe Dienste | read/observe/execute | ja | ja (autorisiert) | THR-011, THR-039 |
| evidence/audit system | Evidenz/Audit | read/export/collect-evidence | ja | export (autorisiert) | THR-016, THR-018 |
| agent/relay | lokale Vermittlung | observe/execute (vermittelt) | ja | ja (autorisiert) | THR-008, THR-009 |
| offline transfer integration | Offline Import/Export | export/import (paketiert) | ja | ja (autorisiert, aktiviert) | THR-024, THR-025 |

Offline-Betrachtungen, Identity-Anforderung, Provenance- und Audit-Erwartung sind für jede Klasse in Companion 2/3 konkretisiert; keine Klasse erhält Write/Execution ohne explizite Policy und Authorization (§18).

## 8. Integration Identity

Eine Integration Instance benötigt konzeptionell mindestens: stable integration identity · integration class · owner · workspace/environment · managed-resource scope · integration principal · lifecycle state · contract version · declared capabilities · detected capabilities · permitted capabilities · trust status · validation status · offline status · audit reference.

**Grundregeln:**
```text
discovered integration ≠ enrolled integration
enrolled integration   ≠ trusted integration
trusted integration    ≠ globally authorised integration
integration identity   ≠ credential
```
Konsistent mit [Machine Identity](../security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md) und [Machine Enrollment](../security/MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md).

## 9. Integration Lifecycle

Konzeptionelle Statuswerte (keine implementierten Enums):
```text
discovered → registration-pending → enrollment-pending → active-read-only → active-restricted →
active-write-enabled → suspended → degraded → revoked → decommissioning → decommissioned → archived
```
`active-write-enabled` benötigt **explizite Policy- und Authorization-Grundlage** (CO-WP-013). Decommissioned IDs dürfen **nicht** still wiederverwendet werden. Der Default nach Enrollment ist `active-read-only` (read-first, DEC-P-04).

## 10. Capability Declaration

Zusammenfassung (Detail in Companion 2 §7). Eine Declaration enthält mindestens: capability identity · capability class · contract version · declaring integration · target resource class · supported operation classes · read/write classification · execution privilege level · required authorization class · offline availability · known limitations · version/compatibility constraints · validation state · evidence reference. Die sechs Dimensionen `advertised/detected/permitted/implemented/supported/validated` sind **getrennt**; keine wird still aus einer anderen abgeleitet. **Ein Adapter darf sich keine Permission selbst erteilen, indem er eine Capability deklariert.**

## 11. Operation Classes

Klassen (Detail/Privileg in Companion 2 §8/§9): discover · query · read · observe · enumerate · export · register · configure · create · update · delete · execute · deploy · start · stop · restart · verify · collect-evidence.

**Grundregeln:**
```text
read     ≠ export
observe  ≠ configure
configure ≠ execute
execute  ≠ deploy
request verify ≠ verified result
capability to execute ≠ authorization to execute
```
Jede Operation Class hat eine eindeutige Privilege- und Authorization-Einordnung.

## 12. Request Model

Konzeptionelles Mindestmodell (keine Serialisierung): contract version · request identity · correlation identity · operation identity · requesting principal · accountable human owner (wo erforderlich) · integration identity · workspace/environment · target resource identity/class · operation class · requested action · requested scope · input classification · desired-state/workflow reference · policy-decision reference · approval reference (wo erforderlich) · execution-authorization reference (wo erforderlich) · credential-state expectation · freshness/trust expectations · timeout/validity boundary · retry/replay context · audit reference.

## 13. Acceptance Model

Statuswerte: `received` · `accepted-for-evaluation` · `rejected` · `authorization-required` · `not-supported` · `indeterminate` · `conflicted`.
```text
received                ≠ accepted
accepted-for-evaluation ≠ authorised
authorised              ≠ execution-started
rejected response       ≠ proof that no earlier target-side action occurred
```

## 14. Operation Lifecycle

Konzeptionelle Statuswerte:
```text
requested → received → evaluation-pending → authorization-pending → accepted → queued-or-deferred →
execution-started → execution-in-progress → execution-completed → result-pending → verification-pending →
completed → partially-completed → failed → cancel-pending → cancelled → outcome-unknown →
reconciliation-required → closed
```
`queued-or-deferred` ist konzeptionell und wählt **keine** Queue-Technologie.
```text
accepted           ≠ queued
queued             ≠ execution-started
execution-completed ≠ successful
completed          ≠ verified
cancel requested   ≠ cancelled
closed             ≠ successful
```

## 15. Result Semantics

Result-States: `succeeded` · `failed` · `partially-succeeded` · `not-executed` · `outcome-unknown` · `cancelled-before-execution` · `cancelled-after-partial-execution` · `verification-failed` · `indeterminate`. Ein Result benötigt mindestens: operation identity · attempt identity · integration identity · target identity · reported outcome · result source · start/completion references · affected-resource references · partial-result information · error/failure classification · rollback/recovery status · verification status · provenance reference · audit reference. **`outcome-unknown` bleibt sichtbar** und gilt weder als Erfolg noch als sichere Nichtausführung.

## 16. Error and Failure Model

Zu unterscheiden: contract error · validation error · authorization error · capability mismatch · target rejection · transport interruption · integration failure · target-side failure · partial failure · timeout · stale response · duplicate response · replay suspected · outcome unknown · verification failure.
```text
transport failure ≠ target action did not occur
timeout           ≠ failed operation
error response    ≠ no side effect
partial failure   ≠ complete failure ≠ complete success
```
Keine konkrete Error-Code-Struktur ausgewählt.

## 17. Read, Write and Execution Boundary

Stufen: `discovery-only` · `read-only` · `observation` · `controlled write` · `privileged execution` · `deployment` · `evidence export`. Regeln:
- Read-only erhält **nicht** still Write.
- Capability Detection verursacht **keine** Write-Aktivierung.
- Write Enablement benötigt explizite Policy und — wo gefordert — Approval.
- Privileged Execution benötigt Execution Authorization (CO-WP-013).
- Adapter/Agent erweitern Scope/Target **nicht**.
- Evidence Export hat eigenen Disclosure-/Authorization-Scope.

## 18. Policy and Authorization Binding

Für Write-/Execution-Operationen mindestens: applicable policy version · policy decision · approval requirement · approval decision · execution authorization · action · target · scope · plan/workflow reference · validity boundary · consumption state · revocation state.

**Grundregel:**
```text
Integration Contract acceptance does not replace Policy, Approval or Execution Authorization.
```
Die Pre-Execution Guards aus [CO-WP-013](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) bleiben autoritativ; **kein Parallelmodell**.

## 19. Async and Long-Running Operations

- Operationen können länger als eine Verbindung dauern; Request-Verbindung und Operation-Lifecycle sind getrennt.
- Statusabfrage ist **keine** Result-Verifikation; Zwischenstände sind nicht automatisch final.
- Verlorene Verbindung beendet **nicht** automatisch die Operation.
- Cancellation benötigt eine eigene bestätigte Statusänderung.
- Bei unbekanntem Outcome ist Reconciliation erforderlich.
- Keine Queue-/Worker-/Polling-/Callback-/Streaming-Technologie ausgewählt.

## 20. Retry, Replay and Duplicate Handling

Fälle: retry after explicit failure · retry after unknown outcome · duplicate request · duplicate result · replayed request · replayed authorization · repeated offline package · resumed operation. Regeln:
- Retry benötigt eine begründete Governance-Entscheidung; **Unknown Outcome verhindert automatische Wiederholung.**
- Duplicate Detection beruht **nicht** allein auf untrusted Timestamps.
- Wiederverwendung verbrauchter Authorization ist unzulässig.
- Ursprüngliche Versuche/Ergebnisse bleiben historisch erhalten.
- **Keine** Idempotency-/Nonce-/Deduplication-Technologie ausgewählt (Bezug THR-026).

## 21. Cancellation, Rollback and Recovery

Zustände: cancel requested/accepted/rejected/in progress · cancelled before/after partial execution · rollback requested/in progress/completed · recovery required/completed · verification pending.
```text
cancel accepted   ≠ operation cancelled
rollback completed ≠ original state restored
recovery completed ≠ desired outcome verified
```
Rollback-/Recovery-Claims benötigen Verifikationsevidenz (Bezug THR-031, THR-032).

## 22. Offline Integration

Mindestens: offline request/result package · target-environment binding · contract-version binding · integration identity · principal identity · accountable owner · operation/scope binding · authorization reference · provenance status · integrity status · validity boundary · usage boundary · import quarantine · local activation · result reconciliation · revocation-distribution challenge · clock/sequence uncertainty. Nicht behauptet: implementierte Offline Integration, automatische konfliktfreie Synchronisation, beliebige Air-Gap-Unterstützung, Eignung für klassifizierte Netze, konkrete Signing-/Token-/Trust-Anchor-Technologie. Konsistent mit [Offline Reconciliation Policy](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md).

## 23. Provenance and State Authority

```text
integration response  ≠ authoritative CoreOps state
adapter field         ≠ canonical CoreOps field
successful parsing    ≠ validated provenance
recently received     ≠ recently observed
vendor or external value ≠ CoreOps policy
```
Die [Source-of-Truth-](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) und [Field-Provenance-Dokumente](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) (CO-WP-011) bleiben autoritativ; **kein paralleles Provenance-Modell**.

## 24. Versioning and Compatibility

Zu unterscheiden: contract version · implementation version · integration version · adapter version · managed-system version · capability version · extension version. Regeln:
- Gleiche Contract-Version ≠ identische Capabilities; neuere Version ≠ automatisch kompatibel.
- Unbekannte Pflichtfelder oder Semantikänderungen dürfen **nicht** still ignoriert werden (fail-safe).
- Additive optionale Erweiterungen benötigen Namespace-/Compatibility-Regeln (§25).
- Breaking Changes benötigen eine neue Contract-Version oder eine klar definierte Compatibility-Entscheidung.
- Keine konkrete Semantic-Versioning-Policy erzwungen, soweit nicht bereits entschieden (DEC-O-02 offen).

## 25. Extension Boundary

Hersteller-/integrationsspezifische Erweiterungen dürfen den Core Contract ergänzen, aber **keine** Core-Sicherheitsinvariante überschreiben, **keine** Pflichtfelder entfernen, **keine** Autorisierung umgehen, **keine** globale Authority erzeugen. Mindestens: extension identity · owner · namespace · contract compatibility · capability impact · security impact · fallback behavior · unknown-extension behavior · audit requirement. **Unknown Extensions mit sicherheitsrelevanter Bedeutung werden nicht still positiv interpretiert.**

## 26. Audit and Evidence

Mindestens erfasst: integration identity · contract version · capability declaration · request · acceptance decision · policy/authorization references · operation attempt · target · scope · execution start · intermediate status · result · partial result · error/failure · retry · cancellation · rollback · recovery · offline import/export · verification · closure. Trennung:
```text
request evidence   ≠ acceptance evidence ≠ execution evidence ≠ success evidence ≠ verification evidence ≠ compliance
```

## 27. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Advertised capability must not imply implemented, permitted or validated capability.
2. Read-only integration must not silently gain write authority.
3. Request acceptance must not imply execution.
4. Execution completion must not imply success.
5. Success must not imply verification.
6. Unknown outcome must not trigger automatic retry.
7. Consumed or revoked authorization must not be reused.
8. Adapter or agent must not expand target, action or scope.
9. Integration data must not inherit authoritative status automatically.
10. Offline integration requires provenance, integrity, target binding and explicit activation.
11. Partial failure must remain visible.
12. Contract extensions must not bypass core security invariants.

## 28. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-001, THR-005, THR-006, THR-007, THR-008, THR-009, THR-010, THR-011, THR-012, THR-013, THR-014, THR-016, THR-017, THR-018, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-025, THR-026, THR-028, THR-029, THR-030, THR-031, THR-032, THR-033, THR-034, THR-035, THR-039.

## 29. Technology Boundary

Nicht ausgewählt/implementiert: Protokoll (REST/GraphQL/gRPC/WebSocket), Spezifikation (OpenAPI/AsyncAPI), Schema (JSON/YAML/Protobuf/XML), Message-Broker/Event-Bus, SDK/Plugin/Adapter, Port/Endpoint, DB-Schema, Authentisierung/Kryptografie, Runtime-Code.

## 30. Compatibility

Konsistent mit [System Context](SYSTEM_CONTEXT_AND_EXTERNAL_BOUNDARIES.md), [Plane Taxonomy](COREOPS_PLANE_TAXONOMY.md), [Module Architecture](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md)/[Catalog](COREOPS_MODULE_CATALOG.md), [Machine Identity](../security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [SoT/Provenance](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [State/Drift](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [Policy/Approval/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Ergänzt DEC-P-01/04, DEC-G-07 und die Integrationsaspekte der Capability Matrix.

## 31. Open Questions

- Verbindliche Semantic-Versioning-Policy des Contracts (DEC-O-02, später).
- Minimale Pflichtfeld-Menge je Operation Class im Detail (Companion 2, spätere Iteration).
- Verhältnis Contract ↔ späteres Job-/Deployment-Modell (CO-WP-021).

## 32. Next Decision

Companion 2 (Capability/Operation) und Companion 3 (Trust/Failure/Recovery) konkretisieren dieses Contract-Gerüst. Domain Pack Governance (CO-WP-015), Deployment Control Plane (CO-WP-021) und Artifact Trust (CO-WP-022) bauen darauf auf. Protokoll-/Schema-/SDK-Wahl bleibt einer späteren ADR-Runde vorbehalten.
