# CoreOps – Recovery Mode Authority and Controlled Restoration Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation recovery-mode authority and controlled-restoration policy
> Implementation Status: Not implemented
> Recovery Engine: Not selected
> Backup/Restore Technology: Not selected
> Orchestration Technology: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Recovery-Readiness Claim: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-026 (docs-only / control-plane self-protection, degraded-operation and governed-recovery foundation)

## 1. Status

Technologieunabhängige Policy für **Recovery Authority, Recovery Mode, Recovery-Stufen, Inputs, Partial/Unknown Recovery, Offline Recovery, Break-Glass, Trust-Reassessment, Reconciliation und Recovery Exit**. Companion zu [SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md](../security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md) und [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md). Baut auf [Deployment Rollback/Forward Recovery (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Key/Recovery (CO-WP-024)](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md), [Break-Glass (CO-WP-009)](../security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) auf; kein Parallelmodell. Keine Recovery-/Backup-/Restore-/Orchestration-Technologie implementiert.

## 2. Purpose

Recovery Mode ist ein governter, temporärer Betriebszustand, kein Administrator-Freibrief. Die Policy legt fest, warum `recovery mode ≠ ordinary operating mode`, `recovery mode ≠ backup restore`, `recovery mode ≠ rollback`, `recovery authority ≠ normal policy authority`, `service restored ≠ authority restored`, `rollback completed ≠ governance restored` und `unknown recovery outcome ≠ recovery success`.

## 3. Scope

Recovery Authority · Recovery Mode · Recovery Stages · Recovery Inputs · Trust/Revocation/Compatibility Reassessment · Partial/Unknown Recovery · Offline/Agent Recovery · Break-Glass · Audit/Evidence/Reconciliation · Data Classification/Disclosure · Recovery Exit · Profile · Failure/Fail-Closed.

## 4. Non-Goals

- Keine Recovery-/Self-Healing-Engine, kein Backup-/Restore-/Snapshot-System, kein Orchestrator/Failover/HA, keine Replication/Sync/Reconciliation-Engine.
- Keine Behauptung implementierter/validierter/zertifizierter Recovery; keine Recovery-Readiness-/Resilience-Reife; kein Runtime-Code; keine Recovery-/Restore-Ausführung gegen reale Systeme.

## 5. Concepts

`recovery mode` · `recovery authority` · `accountable human recovery owner` · `recovery operator/approver` · `recovery plan/action/stage/checkpoint/instance` · `recovery verification` · `recovery reconciliation` · `recovery exit decision` · `recovery input` · `source/target state` · `partial recovery` · `outcome-unknown` · `rollback/forward-recovery boundary`.

## 6. Recovery Authority

Recovery Authority ist eigenständig und begrenzt, getrennt von normaler Policy-/Execution-Autorität ([CO-WP-013](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md)).

```text
recovery authority ≠ normal policy authority
recovery mode      ≠ unrestricted administrator access
recovery operator  ≠ recovery approver
local administrator ≠ recovery authority
break-glass authority ≠ unrestricted recovery authority
machine principal  ≠ human recovery approval
```

Recovery Authority ist scope-/target-/purpose-/zeitgebunden, menschlich zurechenbar und auditiert.

## 7. Recovery Mode

Ein Recovery-Mode-Eintrag: recovery-mode identity · initiating condition · accountable human recovery owner · recovery authority · scope · workspace · environment · affected targets · recovery objectives · permitted recovery capabilities · prohibited normal capabilities · source state · target state · recovery plan · checkpoints · verification plan · rollback/forward-recovery boundary · evidence plan · time/freshness boundary · exit criteria · audit reference · known limitations.

```text
recovery mode      ≠ automatic break-glass
backup available   ≠ restore authorized
restore authorized ≠ recovered state trusted
rollback available ≠ rollback currently safe
```

## 8. Recovery Stages

Getrennte Stufen: **(1)** detect · **(2)** declare · **(3)** contain · **(4)** preserve evidence · **(5)** establish accountable authority · **(6)** assess current state · **(7)** establish trusted recovery inputs · **(8)** approve recovery plan · **(9)** execute bounded recovery action · **(10)** verify immediate outcome · **(11)** reconcile state and evidence · **(12)** reassess trust/policy/identity/secrets · **(13)** determine partial/unknown outcomes · **(14)** approve recovery exit · **(15)** return to an explicitly selected operational mode.

```text
stage completed        ≠ next stage automatically authorized
technical restoration  ≠ governance restoration
data restored          ≠ state reconciled
service reachable      ≠ recovery verified
```

## 9. Recovery Inputs

Mögliche Inputs (je: identity · revision · source · provenance · integrity state · trust state · revocation state · compatibility · target binding · freshness · owner · assessment · approval · known limitations): known-good configuration reference · policy snapshot · identity/role snapshot · secret/key references · trusted artifact revision · CorePack revision · backup/snapshot reference · audit/evidence package · deployment blueprint · recovery runbook · manually approved corrective action.

```text
previously trusted  ≠ currently trusted
previously deployed ≠ safe recovery input
backup exists       ≠ backup is valid for this target
old configuration   ≠ known-good configuration
```

Mutable Aliases sind **keine** finale Recovery-Bindung. Keine Recovery-/Backup-Technologie ausgewählt.

## 10. Trust, Revocation and Compatibility Reassessment

Recovery Inputs werden **aktuell** neu bewertet auf Trust, Revocation, Provenance, Integrity, Compatibility, Target Binding und konkrete Revision ([CO-WP-021](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)/[022](../security/ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md)/[023](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md)):

```text
artifact previously trusted ≠ safe recovery artifact
deployment rollback available ≠ rollback authorized
CorePack imported  ≠ recovery authorized
CorePack activated ≠ governance restored
recovery deployment completed ≠ desired state verified
```

**Secrets/Keys ([CO-WP-024](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md)):** `secret restored ≠ secret trusted`; `key recovered ≠ key use authorized`; `old credential available ≠ rollback credential valid`. Recovery legt keine Raw Secrets offen; Retrieval/Use bleiben getrennt autorisiert; widerrufene/abgelaufene/kompromittierte Secrets werden **nicht** durch Rollback reaktiviert; Recovery eines Key Materials ist keine Reinstatement-Entscheidung (DEC-S-329); Recovery-Evidence enthält keine Secret Values.

## 11. Partial and Unknown Recovery

Zustände: recovered · partially-recovered · verification-pending · reconciliation-pending · rollback-pending · forward-recovery-pending · outcome-unknown · recovery-failed · recovery-aborted.

```text
partial recovery       ≠ normal operation
unknown recovery outcome ≠ safe retry
recovery failed        ≠ no side effects
recovery aborted       ≠ prior state restored
```

Unknown Outcomes bleiben sichtbar, blockieren unsichere automatische Wiederholung, erfordern neue Zustandsaufnahme + Reconciliation und berücksichtigen mögliche Side Effects (Bezug THR-031, THR-032).

## 12. Offline and Agent Recovery

Bindet an [CO-WP-023](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md): zentrale Nichterreichbarkeit erweitert lokale Autorität nicht; Agenten setzen nur innerhalb bereits delegierter Grenzen fort; abgelaufene Delegationen werden nicht still verlängert; lokale Recovery Authority muss ausdrücklich vorhanden sein; lokale Recovery-Ergebnisse benötigen Evidence Return; fehlende zentrale Bestätigung bedeutet weder Fehlschlag noch Erfolg; Reconciliation bleibt erforderlich; Clock-/Freshness-Uncertainty explizit.

```text
agent operational        ≠ agent centrally current
local recovery completed ≠ central reconciliation complete
offline evidence returned ≠ evidence complete or accepted
```

Keine Agent-/Sync-/Offline-Runtime-Technologie ausgewählt.

## 13. Break-Glass in Recovery

Break-Glass ([CO-WP-009](../security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md)) bleibt action-/scope-/target-/purpose-/profile-/version-/zeitgebunden, einmalig/verbrauchsbegrenzt, menschlich zurechenbar und nachträglich reviewpflichtig. `break-glass authority ≠ unrestricted recovery authority`; Break-Glass erzeugt **keine** dauerhafte Policy- oder Recovery-Autorität.

## 14. Audit, Evidence and Reconciliation

Bindet an [CO-WP-018](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md): Audit-Ausfall ≠ Abwesenheit von Operationen; Telemetry-Ausfall ≠ gesunder Zustand; Recovery-Evidence und Source Evidence bleiben getrennt; Schutz-/Recovery-Maßnahmen erzeugen nachvollziehbare Ereignisse, soweit Evidence-Fähigkeit verfügbar; Evidence-Ausfall kann privilegierte Operationen einschränken. `evidence package available ≠ recovery claim proven`. Reconciliation gleicht wiederhergestellten Zustand mit autoritativem Zustand ab; keine automatische Reconciliation-Engine.

## 15. Data Classification and Disclosure in Recovery

Bindet an [CO-WP-025](DATA_CLASSIFICATION_AND_HANDLING_MODEL.md): Recovery-Daten können sensitive/restricted/secret-bearing/evidence-protected sein; Recovery Mode erweitert **keine** Disclosure-/Exportautorität; Redaction erzeugt nur eine Derived View; redigierte Recovery Reports sind nicht automatisch disclosure-safe; Preservation Holds können Recovery-Evidence betreffen; `retention expiry ≠ deletion completion`; secret-bearing Recovery-Daten bleiben an CO-WP-024 gebunden.

## 16. Recovery Exit

Recovery Mode wird nur verlassen, wenn bewertet: accountable authority · identity state · authorization state · policy state · configuration state · source-of-truth state · secret/key state · artifact trust/revocation · deployment state · workspace/target binding · audit/evidence continuity · offline delegation · time/freshness · remaining degraded capabilities · unresolved unknown outcomes · follow-up actions.

```text
recovery action succeeded ≠ recovery exit authorized
recovery exit authorized  ≠ normal mode required
recovery may exit into guarded, restricted or degraded mode
historical incident state remains auditable
```

Die Exit-Entscheidung ist menschlich zurechenbar, scope-/zeitgebunden und auditierbar.

## 17. Profiles

`Standard`/`Hardened`/`Government` variieren zusätzliche Human Approvals, Recovery-Input-Bewertung, Rollentrennung, Evidence-Anforderungen, Exit-Freigabe und Unknown-State-Verhalten. `Government profile ≠ government certification`; `Hardened ≠ proven resilient`; `Standard ≠ self-protection optional`.

## 18. Failure and Fail-Closed

Keine privilegierte Recovery-Aktion bei: fehlender Recovery Authority · unbewerteten Recovery Inputs · stale/revoked/inkompatiblen Inputs · unbekanntem Source-/Target-State · fehlender Approval · unbekannter Offline-Reconciliation · fehlendem Audit-Start · Alias-only-Binding. `unknown ≠ safe`; `service reachable ≠ recovery verified`; Unknown Outcomes → Reconciliation, kein unsicherer Retry.

## 19. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Recovery mode remains separate from ordinary operation, backup restore and rollback.
2. Recovery authority is self-contained and bounded, separate from normal policy and execution authority; recovery operator ≠ recovery approver.
3. Recovery inputs require current identity, provenance, integrity, trust, revocation, compatibility, freshness and target-binding assessment; previously trusted ≠ currently trusted.
4. Recovered secrets, keys, artifacts, configurations and backups are not automatically trusted; rollback does not reactivate revoked/expired/compromised material.
5. Technical restoration is not governance restoration; service restored ≠ authority restored.
6. Partial and unknown recovery outcomes remain visible and block unsafe retry.
7. Offline operation does not expand local recovery authority; expired delegations are not silently extended.
8. Break-glass does not create permanent policy or recovery authority.
9. Recovery exit requires reassessment, verification and reconciliation; recovery may exit into guarded/restricted/degraded rather than normal mode.
10. Recovery evidence contains no raw secrets; missing audit does not prove no action occurred.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 20. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): rollback fails THR-032; partial failure without safe state THR-031; backup/restore manipulated THR-033; manipulated update/deployment artifact THR-021/THR-022; offline import THR-024; stolen admin identity THR-001; reuse expired approval THR-004; replay THR-026; false time THR-027; secret exposure THR-019/THR-020; audit deletion/manipulation THR-016/THR-017; automation client THR-038. Keine erfundenen IDs; kein Parallelregister.

## 21. Technology Boundary

Nicht ausgewählt: Recovery-/Self-Healing-Engine · Backup-/Restore-/Snapshot-System · Orchestrator · Failover/HA · Replication · Synchronisations-/Reconciliation-Engine · State-Machine-Framework. Alle `deferred`.

## 22. Compatibility

Baut auf [Deployment Rollback/Forward Recovery (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Key/Recovery (CO-WP-024)](../security/KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md), [Break-Glass (CO-WP-009)](../security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md), [Data Classification (CO-WP-025)](DATA_CLASSIFICATION_AND_HANDLING_MODEL.md), [Offline/CorePack (CO-WP-023)](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md) auf. DEC-S-329 (recovery ≠ reinstatement) referenziert, nicht dupliziert. Additiv; keine bestehende Invariante geschwächt.

## 23. Open Questions

- Konkrete Recovery-Runbook-Struktur (deferred).
- Verfahren zur Bestätigung „known-good" ohne festgelegte Backup-/Hash-Technologie (deferred).
- Reconciliation-Verfahren für Offline-Recovery-Ergebnisse (deferred).

## 24. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): gebündelter `CO-WP-021…026` Foundation Milestone Review (noch nicht terminiert, keine WP-Nummer). Zuerst Nova Review von CO-WP-026, danach Human-Maintainer-Commit.
