# CoreOps – Deployment Targeting, Execution and Recovery Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation deployment-targeting, execution and recovery policy
> Implementation Status: Not implemented
> Execution Technology: Not selected
> Rollback Technology: Not selected
> Offline Activation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-021 (docs-only / deployment architecture, targeting and execution-governance foundation)

## 1. Status

Technologieunabhängige Policy für **Deployment-Autorität, Target-Set-Governance, Preflight, Waves, Partial/Unknown Outcome, Verification, Rollback/Recovery, Manual Authority, Offline und Fail-Closed**. Companion zu [DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md) und [DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md](../architecture/DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md). Keine Execution-/Rollback-Technologie, keine Offline-Activation implementiert.

## 2. Purpose

Deployment ist privilegierte Massenveränderung: Target-Scope, Artifact-Trust und Erfolgsbewertung sind die kritischen Fehlerpunkte. Diese Policy legt fest, dass jede Ausführung an Policy/Approval/Execution Authorization (CO-WP-013) gebunden bleibt und warum `blueprint availability ≠ authorization`, `topology query ≠ authorised target set`, `executed ≠ verified` und `rollback completed ≠ restored unless verified`.

## 3. Scope

Authority Model · Blueprint/Artifact/Dependency Trust · Target-Set Authority/Revalidation · Parameter/Overlay Boundary · Policy/Approval/Execution Authorization · Preflight/Audit-Start · Waves · Partial Deployment · Pause/Resume/Cancellation · Unknown Outcome · Verification · Rollback/Forward Recovery · Manual Authority · Offline · Workspace Isolation · Failure/Closure · Fail-Closed.

## 4. Non-Goals

- Keine Execution-/Rollback-/Recovery-/Rollout-Engine, kein Orchestrator/Agent/Runner.
- Keine Registry-/Signing-/Secret-/Offline-Activation-Technologie.
- Keine Behauptung implementierter/erfolgreicher/zertifizierter Deployments; kein Runtime-Code.

## 5. Authority Model

Getrennt: Deployment Owner · Blueprint Owner · Artifact Owner · **Deployment Approval** (human) · **Execution Authorization** (CO-WP-013) · Recovery Owner · Manual Authority (human). `control-plane decision ≠ execution`; **Blueprint/Plan/Control Plane erzeugen keine parallele Authorization Authority**.

## 6. Blueprint Trust

Blueprint-Verfügbarkeit/-Aktivierung ist keine Autorisierung; `active ≠ validated`; Effective-Konfiguration-Provenance bleibt erhalten; Overlays schwächen Security nicht (Companion 2 §12).

## 7. Artifact and Dependency Trust

`artifact available ≠ trusted ≠ compatible`; `integrity checked ≠ safe`; Provenance/Integrität/Compatibility sind eigene Dimensionen (Companion 2 §15-16, Bezug THR-021, THR-022, THR-023). Widerrufene/zurückgezogene Artifacts sind nicht autoritativ.

## 8. Target-Set Authority

Ein Deployment wirkt nur auf ein **materialisiertes, autorisiertes Target Set** (Companion 1 §10); `topology query ≠ authorised target set`; conflicted/unresolved Identität wird nicht still privilegiertes Ziel (Bezug THR-014, THR-035). Excluded Targets bleiben ausgeschlossen.

## 9. Topology Revalidation

Vor privilegierter Ausführung (und bei Resume) wird das Target Set revalidiert (Companion 1 §12): `approved target snapshot ≠ permanently valid target scope`. Wesentliche Änderung → Re-Evaluation/neue Approval/neue Execution Authorization (Bezug THR-013, THR-015).

## 10. Parameter and Overlay Boundary

Parameter/Overlays sind provenance-erfasst und **schwächen keine Security-Requirements still**; `default ≠ safe`; `target-derived input ≠ target authorization`; keine Raw Secrets (Companion 2 §9-12).

## 11. Policy Decision

Deployment-Aktionen benötigen anwendbare Policy-Version/Decision (CO-WP-013); `permit ≠ authorization`; `indeterminate/conflicted → keine privilegierte Ausführung`.

## 12. Approval

Privilegierte/destruktive Deployments benötigen Approval (scope-/plan-/target-set-gebunden); `approved ≠ executed`; Machine Principals können nicht selbst genehmigen (§24).

## 13. Execution Authorization

Ausführung benötigt gültige Execution Authorization (action-/target-/scope-/plan-/time-bound, CO-WP-013); `approved plan ≠ execution authorization`; expired/revoked/consumed nicht wiederverwendbar (Bezug THR-026).

## 14. Preflight

Wie Companion 1 §15; **fail-closed** bei unbekanntem Target Scope, konflikthafter Identität, fehlender Authorization, unklarer Artifact-Provenance, fehlendem Audit-Start, unbekanntem Recovery-Pfad bei destruktivem Scope.

## 15. Audit Start Boundary

Ohne möglichen Audit-Start-Record erfolgt **keine** privilegierte Ausführung (konsistent mit [Audit Policy](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), Bezug THR-016, THR-017).

## 16. Waves and Batches

`wave success ≠ remaining targets authorised automatically`; jede Wave-Expansion ist eine explizite Entscheidung mit Entry/Stop-Conditions und Verifikation (Companion 1 §17).

## 17. Partial Deployment

`partial ≠ complete success ≠ complete failure`; Per-Target-Zustände bleiben sichtbar; Attempt-/Per-Target-Historie nicht durch Aggregate überschrieben (Bezug THR-031).

## 18. Pause and Resume

`pause requested ≠ paused`; `resume requested ≠ safe continuation`; Resume revalidiert Authorization/Target Set/Artifacts/Dependencies/Environment/Topology/Remaining Plan/Recovery (Companion 1 §19).

## 19. Cancellation

`cancel accepted ≠ execution stopped ≠ no side effects`; Cancellation nach Partial Execution sichtbar, Reconciliation erforderlich.

## 20. Unknown Outcome

`unknown target outcome → no automatic retry → reconciliation required`; ursprüngliche Authorization nicht still wiederverwendbar (Bezug THR-026, THR-028).

## 21. Verification

`execution completion ≠ desired-state verification`; `artifact present ≠ desired state achieved`; `health signal green ≠ verified`; `telemetry available ≠ sufficient evidence`. Verification target-/scope-/version-/zeitgebunden (Evidence-Sufficiency CO-WP-018, Telemetry-Grenzen CO-WP-019).

## 22. Rollback

`rollback available ≠ safe ≠ validated`; `rollback completed ≠ prior state restored unless verified`. Rollback benötigt gebundene Autorität + Verifikationsevidenz (Bezug THR-032). Keine Rollback-Technologie.

## 23. Forward Recovery

Forward Recovery bei unmöglichem/unvollständigem/riskanterem Rollback; `recovered ≠ complete service recovery unless verified`. Recovery Owner verantwortet Nachweis.

## 24. Manual Authority

Manuelle Aktionen (Target Exclusion, Pause, Resume, Wave Expansion, Rollback, Forward Recovery, Known Exception, Verification Note) bleiben human-attributable, scope-bound, zeit-/reviewgebunden, auditierbar und **nicht-destruktiv** gegenüber bestehender Evidence. **Machine Principals imitieren keine Human Approval/Manual Authority** (konsistent mit [Topology Manual Authority](../architecture/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md)).

## 25. Offline Deployment

Offline-Deployment-Pakete benötigen: blueprint identity/version · effective parameters · target environment · materialised target set · artifact/dependency set · policy/approval references · execution authorization · provenance/integrity status · validity/usage boundary · import quarantine · local activation · execution attempts · partial results · unknown outcomes · result package · reconciliation · delayed revocation · clock/sequence uncertainty. Nicht behauptet: implementiert, automatische konfliktfreie Ausführung, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance/Autorität.

## 26. Workspace Isolation

Target Set, Blueprint und Deployment sind workspace-/environment-gebunden; Cross-Workspace-Deployment nur mit expliziter Autorität; Ressourcenexistenz nicht unautorisiert offengelegt (Bezug THR-035).

## 27. Failure and Closure

`closed ≠ successful ≠ complete ≠ compliant`; Closure benötigt Closure Reason · Final Outcome · Verification Status · Per-Target-Zusammenfassung · Remaining Risk/Exception · Owner · Audit Reference.

## 28. Audit and Evidence

Erfasst wie Companion 1 §29. Trennung `plan ≠ authorization ≠ execution ≠ verification evidence ≠ compliance`; Historie nicht umgeschrieben.

## 29. Fail-Closed Rules

Keine privilegierte Deployment-Ausführung bei: unbekanntem/unbounded Target Scope · conflicted Identity · fehlender/abgelaufener/widerrufener Authorization · unklarer Artifact-Provenance/-Integrity · Blueprint-/Artifact-Version-Mismatch · fehlendem Audit-Start · unbekanntem Recovery-Pfad bei destruktivem Scope · `indeterminate`/`conflicted` Policy · unklarer Workspace-Grenze · unbekanntem Offline-Provenance.

## 30. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Blueprint availability must not imply deployment authorization.
2. Dynamic topology selection must be materialised into a bounded target set before privileged deployment.
3. Target-set approval must not remain valid after material target/topology/blueprint/artifact/authorization change without re-evaluation.
4. Blueprint parameters and overlays must not weaken security requirements or expand target scope silently.
5. Artifact availability must not imply provenance, integrity, compatibility or authorization.
6. Execution completion must not imply desired-state verification.
7. Partial and unknown per-target outcomes must remain visible and block unsafe automatic retry.
8. Cancellation must not prove absence of side effects.
9. Rollback completion must not prove restoration without verification.
10. Machine principals must not imitate human deployment approval or manual authority.
11. Offline deployment requires provenance, integrity, target binding, bounded authorization and explicit activation.
12. No privileged deployment without a possible audit-start record.

## 31. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-002, THR-007, THR-008, THR-010, THR-013, THR-014, THR-015, THR-016, THR-017, THR-021, THR-022, THR-023, THR-024, THR-026, THR-028, THR-031, THR-032, THR-034, THR-035. Keine Duplikation, kein Parallelregister.

## 32. Technology Boundary

Nicht ausgewählt/implementiert: Execution-/Rollback-/Recovery-/Rollout-Engine, Orchestrator/Agent/Runner, Registry/Signing/Secret, Offline-Activation, Runtime-Code.

## 33. Compatibility

Konsistent mit Control-Plane-/Blueprint-Companion, [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Integration Trust](INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Migration Integrity](DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md), [Topology Manual Authority](../architecture/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md), [Audit Policy](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md). Konkretisiert DEC-P-02, DEC-P-04, DEC-G-04, DEC-G-07.

## 34. Open Questions

- Konkrete Rollback-/Recovery-Mechanismen (spätere ADR, mit CO-WP-026).
- Wave-Approval-Schwellen je Risikoklasse.
- Offline-Activation-Mechanismus (CO-WP-023).

## 35. Next Decision

Artifact Trust/SBOM (CO-WP-022), Restricted/Air-Gapped Operation (CO-WP-023), Secrets/Key Custody (CO-WP-024) und Self-Protection/Recovery (CO-WP-026) konkretisieren Trust-/Recovery-/Offline-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
