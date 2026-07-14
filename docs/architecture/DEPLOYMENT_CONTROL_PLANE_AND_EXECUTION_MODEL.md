# CoreOps – Deployment Control Plane and Execution Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation deployment-control-plane and execution model
> Implementation Status: Not implemented
> Deployment Engine: Not selected
> Orchestrator: Not selected
> Agent Protocol: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-021 (docs-only / deployment architecture, targeting and execution-governance foundation)

## 1. Status

Technologieunabhängiges Modell für den **CoreOps Deployment Control Plane**: Intent, Plan, Target Set, Waves, Verification, Rollback/Recovery und deren Governance. Companion zu [DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md](DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md) und [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](../security/DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md). Es implementiert **keine** Deployment Engine, keinen Orchestrator, keine CI/CD-Plattform, kein Blueprint-Format, kein Agent-Protokoll.

## 2. Purpose

Ein Deployment-Vorgang ist der gefährlichste privilegierte Vorgang: er wählt Ziele, verändert Zustand und kann teilweise oder unbekannt scheitern. Dieses Modell trennt Blueprint → Intent → Plan → Approval → Execution Authorization → Execution → Verification → Closure und bindet jede Write-Aktion an die bestehende Execution-Authorization (CO-WP-013) — damit `blueprint availability ≠ deployment authorization`, `executed ≠ successful ≠ verified` und `topology query ≠ authorised target set`.

## 3. Scope

Control-Plane-Grenze · Deployment Intent/Plan · Target Sets/Snapshots · Topology Selection/Revalidation · Artifacts/Dependencies · Preflight · Deployment Lifecycle · Waves/Batches · Concurrency/Ordering · Pause/Resume/Cancellation · Partial/Unknown Outcome · Verification · Rollback/Forward Recovery · State-Authority-/Policy-/Authorization-Bindung · Offline · Audit/Evidence.

## 4. Non-Goals

- Keine CI/CD-/Deployment-Engine-/Orchestrator-Auswahl, kein Agent-/Runner-Protokoll.
- Kein Blueprint-/Manifest-Format, kein YAML/JSON/DSL-Schema, keine Pipeline-Syntax.
- Keine Container-/VM-Plattform, Artifact Registry, Secret-Technologie, Rollout-/Rollback-Engine.
- Kein Runtime-Code, kein tatsächliches Deployment; keine Behauptung implementierter/erfolgreicher Deployments.

## 5. Concepts

Begriffe (mindestens): deployment control plane · deployment blueprint · blueprint identity/version · blueprint input · parameter · environment overlay · deployment intent · deployment plan · deployment target set · target-set snapshot · deployment operation/attempt · wave · batch · artifact · dependency · preflight · approval · execution authorization · execution · result · verification · rollback · forward recovery · closure.

**Grundregeln:**
```text
blueprint            ≠ deployment plan
deployment intent    ≠ approved plan
approved plan        ≠ execution authorization
target query         ≠ authorised target set
artifact available   ≠ artifact trusted
deployment executed  ≠ deployment successful
deployment successful ≠ desired state verified
rollback started     ≠ prior state restored
closed               ≠ successful
```

## 6. Control-Plane Boundary

Der Deployment Control Plane ist die **Governance-/Koordinationsgrenze**, nicht automatisch Runtime, Deployment Engine, Agent, Pipeline, Artifact Store, Secret Store, Orchestrator oder autoritative Resource Source.
```text
control-plane decision   ≠ target execution
control-plane availability ≠ deployment authorization
```

## 7. Deployment Intent

Ein Intent beschreibt die **gewünschte Veränderung**: intent identity · owner · workspace/environment · blueprint reference · requested version · business/operational reason · requested target class/selection · desired outcome · urgency · risk classification · requested window · evidence references · audit reference.
```text
intent recorded ≠ plan approved ≠ deployment authorised
```

## 8. Deployment Plan

Ein Plan benötigt mindestens: plan identity · intent reference · accountable human owner · blueprint identity/version · resolved effective parameters · artifact/dependency references · **materialised target-set snapshot** · target inclusion/exclusion · waves/batches · ordering/dependencies · concurrency boundary · preconditions · change/maintenance window · policy decision reference · approval requirements · execution-authorization expectation · verification plan · rollback feasibility · forward-recovery path · offline considerations · audit reference. **Ein Plan ist keine Execution Authorization.**

## 9. Target Sets

Ein Target Set ist die Menge autorisierter Ziele eines Plans. Eine dynamische Topology Query/View dient **nicht** direkt als Scope (§11). Nodes werden aus der [Topology](TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md) selektiert, aber in einen begrenzten Snapshot materialisiert (§10).

## 10. Target-Set Snapshots

Ein Snapshot benötigt: target-set identity · topology view/query reference · resolved canonical node identities · resource identity references · workspace/environment · target classes · inclusion/exclusion rules · resolution time · topology freshness · identity-resolution state · relationship-conflict state · completeness limitation · evidence references · owner · audit reference.
```text
topology query result ≠ approved deployment target set
```
**Keine konfliktbehaftete/ungeklärte Identität wird still zum privilegierten Ziel.**

## 11. Topology Selection

Selektion nutzt Topology-Assertions (observed/declared/manual), respektiert aber deren Grenzen: `graph presence ≠ resource reachable`; `node present ≠ available`; conflicted/unresolved Identity blockiert Aufnahme ins Target Set. Auswahl erzeugt keine Autorität (`node selection ≠ authorised target scope`, CO-WP-020).

## 12. Target Revalidation

Unmittelbar vor privilegierter Ausführung wird geprüft: target identities still resolve · workspace/environment unchanged · target eligibility still valid · topology conflicts not materially worsened · policy/approval scope still match · artifact/blueprint versions still match · execution authorization still valid · excluded targets remain excluded.
```text
approved target snapshot ≠ permanently valid target scope
```
Eine **wesentliche Änderung** erfordert neue Bewertung, aktualisierten Plan, neue/aktualisierte Approval und neue Execution Authorization (abhängig von der Governance, CO-WP-013).

## 13. Artifacts

Detail in Companion 2 §15. Ein Artifact: identity · version · class · source · owner · producer · provenance/integrity/validation status · compatibility state · target constraints · dependency references · known limitations · withdrawal/revocation state · audit reference.
```text
artifact downloaded ≠ artifact trusted
integrity checked   ≠ artifact safe
artifact available  ≠ artifact compatible
newer artifact      ≠ authorised replacement
```
Keine Registry-/Signing-/Hash-Technologie (CO-WP-022 offen).

## 14. Dependencies

Dependency-Bindung (Companion 2 §16): identity · type · version/compatibility scope · required/optional · target relevance · availability · trust/validation state · failure behavior · offline availability · audit. Dependency-Änderungen machen Plan/Compatibility/Verification erneut bewertbar; keine automatische Dependency Resolution.

## 15. Preflight

Mindestens geprüft: blueprint identity/version known · effective parameters resolved · required inputs present · unknown inputs absent/explizit behandelt · target set materialised · target identities/scopes valid · topology conflicts assessed · artifacts/dependencies available · provenance/integrity assessed · policy decision current · approval current · execution authorization current · audit start record possible · change window valid · concurrency conflicts assessed · rollback/forward-recovery path assessed · verification plan present · offline nodes/delayed state considered. **Fail-closed** bei unbekanntem Target Scope, konflikthafter Target Identity, fehlender Authorization, unklarer Artifact-Provenance, fehlendem Audit-Start, unbekanntem Recovery-Pfad bei destruktivem Scope.

## 16. Deployment Lifecycle

Konzeptionelle Statuswerte:
```text
proposed → analysis-pending → plan-ready → approval-pending → approved → scheduled →
preflight-pending → preflight-failed → execution-pending → executing → partially-deployed →
executed → verification-pending → verified → verification-failed → pause-pending → paused →
resume-pending → cancellation-pending → cancelled-before-execution → cancelled-after-partial-execution →
rollback-pending → rolling-back → rolled-back → forward-recovery-pending → recovered → outcome-unknown →
suspended → closed
```
```text
approved  ≠ authorised execution
executed  ≠ successful
verified  ≠ compliant
rolled-back ≠ original state restored unless verified
recovered ≠ complete service recovery unless verified
closed    ≠ successful
```

## 17. Waves and Batches

Modi: `single-target · single-batch · multiple-batch · wave-based · canary-like limited wave · progressive expansion · manual checkpoint · unknown rollout mode` (konzeptionell, keine Rollout-Technologie). Jede Wave: wave identity · target subset · entry conditions · execution boundary · verification conditions · stop conditions · expansion decision · remaining targets · result state · audit reference.
```text
wave success ≠ remaining targets authorised automatically
```

## 18. Concurrency and Ordering

Dokumentiert: abhängige Targets · shared Resources · Exclusive Operations · konkurrierende Deployment Plans · laufende Workflows · Migration/Maintenance · Offline Targets · Target Drift · Dependency Ordering. Keine Locking-/Queue-/Transaction-/Scheduler-Technologie.

## 19. Pause and Resume

`pause requested ≠ deployment paused`; `resume requested ≠ safe continuation`. **Resume revalidiert** Authorization, Target Set, Artifacts, Dependencies, Environment, Topology, Remaining Plan, Recovery Options (§12).

## 20. Cancellation

`cancel accepted ≠ execution stopped ≠ no side effects`. Cancellation nach Partial Execution bleibt sichtbar (`cancelled-after-partial-execution`) und benötigt Reconciliation.

## 21. Partial Deployment

`partial deployment ≠ complete success ≠ complete failure`. Per-Target-Zustände (succeeded/failed/not executed/outcome-unknown/rolled-back/forward-recovery-required/offline-unresolved) bleiben sichtbar; **Attempt-/Per-Target-Historie wird nicht durch Aggregate Results überschrieben** (Bezug THR-031).

## 22. Unknown Outcome

```text
unknown target outcome → no automatic retry → reconciliation required
```
Unbekannte Ziel-Ergebnisse bleiben explizit; ursprüngliche Authorization nicht still wiederverwendbar (konsistent mit [Integration Trust §13](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [API Error §14](../security/API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md), Bezug THR-026, THR-028).

## 23. Verification

Zu unterscheiden: execution completion · artifact presence · configuration · desired-state · service-health observation · dependency · functional · operational · manual verification · evidence sufficiency assessment.
```text
artifact present    ≠ desired state achieved
health signal green ≠ deployment verified
telemetry available ≠ sufficient deployment evidence
```
Verification bleibt target-/scope-/version-/zeitgebunden (Evidence-Sufficiency aus CO-WP-018).

## 24. Rollback

Rollback (Companion 3 §22): rollback identity · affected targets · source deployment attempt · prior blueprint/artifact context · rollback plan · authorization · preconditions · execution result · verification · remaining differences · audit reference.
```text
rollback available ≠ rollback safe ≠ rollback validated
rollback completed ≠ prior state restored unless verified
```
Keine Rollback-Technologie (Bezug THR-032).

## 25. Forward Recovery

Forward Recovery ist erforderlich, wenn Rollback unmöglich/unvollständig/riskanter ist. `recovered ≠ complete service recovery unless verified`. Recovery Owner verantwortet den Nachweis.

## 26. State-Authority Boundary

```text
deployment requested  ≠ desired state accepted
deployment executed   ≠ observed state matches
observed state matches ≠ effective state verified
```
Die [State-Authority-](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) und [Drift-Modelle](DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md) bleiben autoritativ; **kein Parallel-State-Modell**.

## 27. Policy and Authorization Binding

Jede privilegierte Deployment-Ausführung bindet: policy version · policy decision · approval requirement/decision · execution authorization · action · target set · scope · plan reference · blueprint/artifact versions · validity boundary · revocation/consumption state. **Blueprint/Plan/Control Plane erzeugen keine parallele Authorization Authority** ([CO-WP-013](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) autoritativ).

## 28. Offline Deployment

Offline-Deployment-Pakete folgen Companion 3 §25 (target-environment/blueprint/artifact binding, provenance, integrity, bounded authorization, explicit activation, import quarantine, reconciliation). Nicht behauptet: implementiert, automatische konfliktfreie Ausführung, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-Technologie (Bezug THR-024).

## 29. Audit and Evidence

Erfasst: intent · blueprint/version · effective parameters · target-set snapshot · topology reference · artifact/dependency references · policy/approval · execution authorization · preflight · wave/batch · attempt · per-target result · pause/resume/cancellation · unknown outcome · verification · rollback · forward recovery · manual decision · closure. Trennung:
```text
plan evidence ≠ authorization evidence ≠ execution evidence ≠ verification evidence ≠ compliance
```

## 30. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Blueprint availability must not imply deployment authorization.
2. Deployment intent, plan, approval, authorization, execution and verification remain separate.
3. Dynamic topology selection must be materialised into a bounded target set.
4. Target-set approval must not remain valid after material scope change without re-evaluation.
5. Blueprint parameters and overlays must not weaken security requirements silently.
6. Artifact availability must not imply provenance, integrity or compatibility.
7. Execution completion must not imply desired-state verification.
8. Partial and unknown per-target results must remain visible.
9. Cancellation must not imply absence of side effects.
10. Rollback claims require verification evidence.
11. Machine principals must not imitate human deployment approval or manual authority.
12. Offline deployment requires provenance, integrity, target binding, bounded authorization and explicit activation.

## 31. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-002, THR-005, THR-007, THR-008, THR-010, THR-013, THR-014, THR-015, THR-016, THR-017, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-026, THR-028, THR-031, THR-032, THR-034, THR-035.

## 32. Technology Boundary

Nicht ausgewählt/implementiert: Deployment Engine/Orchestrator/CI-CD, Agent-/Runner-/Transportprotokoll, Blueprint-/Manifest-Format, Schema/Template Engine, Pipeline-Syntax, Container-/VM-Plattform, Artifact Registry, Secret-Technologie, Rollout-/Rollback-/Recovery Engine, Runtime-Code.

## 33. Compatibility

Konsistent mit [Policy/Approval/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [Domain Pack](DOMAIN_PACK_GOVERNANCE_MODEL.md), [Data/Migration](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [API](API_GOVERNANCE_AND_OPERATION_MODEL.md), [Event/Evidence](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Telemetry](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [Topology](TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md), [SoT/State/Drift](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert DEC-P-02 (Offline-First), DEC-G-04/07, CAP-DEPLOY-*.

## 34. Open Questions

- Deployment-Engine-/Blueprint-Format-Wahl (spätere ADR, mit CO-WP-022 Artifact Trust verbunden).
- Konkrete Wave-/Rollout-Strategien je Risikoklasse.
- Snapshot-/Consistency-Boundary-Mechanismus (mit CO-WP-026).

## 35. Next Decision

Companion 2 (Blueprint/Versioning) und Companion 3 (Targeting/Execution/Recovery-Policy). Artifact Trust/SBOM/Provenance (CO-WP-022), Restricted/Air-Gapped Operation (CO-WP-023) und Self-Protection/Recovery (CO-WP-026) konkretisieren einzelne Grenzen. Engine-/Format-Wahl bleibt einer späteren ADR-Runde vorbehalten.
