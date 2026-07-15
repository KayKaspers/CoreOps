# CoreOps – Self-Protection and Control-Plane Safety Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation self-protection and control-plane safety model
> Implementation Status: Not implemented
> Health/Watchdog Technology: Not selected
> HA/Failover Technology: Not selected
> Isolation Technology: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Resilience Claim: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-026 (docs-only / control-plane self-protection, degraded-operation and governed-recovery foundation)

## 1. Status

Technologieunabhängiges Modell für **Self-Protection des CoreOps Control Planes**: Control-/Managed-Plane-Grenze, Schutzgüter, Protective Conditions/Trigger, Protection Assessment, Fault Domains/Blast Radius und Schutzmaßnahmen. Companion zu [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) und [RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md). Adressiert die Self-Dependency-Grenze ([RISK-11](../../project-system/RISK_REGISTER.md)); baut auf den bestehenden Operational States aus [Restricted Operation (CO-WP-023) §14](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md) auf; kein Parallelmodell. Keine Health-/Watchdog-/HA-/Failover-/Isolation-Technologie implementiert.

## 2. Purpose

Ein laufender Prozess ist kein governbarer Control Plane. Das Modell legt fest, warum `CoreOps self-protection ≠ protection of every managed asset`, `control-plane health ≠ managed-system health`, `process running ≠ platform governable`, `health check passed ≠ platform secure`, `monitoring unavailable ≠ system healthy`, `audit unavailable ≠ no operation occurred`, `self-protection trigger ≠ proof of compromise` und `absence of trigger ≠ proof of safety`.

## 3. Scope

Control-/Managed-Plane-Grenze · Schutzgüter/Self-Protection Scope · Protective Conditions/Trigger · Protection Assessment · Fault Domains/Blast Radius · Protective Actions/Priority · Integrationsgrenzen · Profile · Failure/Unknown State. Operational Modes/Capability Matrix im [Companion Degraded-Mode Model](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md); Recovery im [Companion Recovery Policy](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md).

## 4. Non-Goals

- Keine Health-Monitoring-/Watchdog-/Heartbeat-/Leader-Election-/Cluster-/HA-/Failover-/Load-Balancer-/Consensus-/Quorum-/Circuit-Breaker-/Rate-Limiter-/Isolation-/Sandbox-/Firewall-/Runtime-Security-/Integrity-Scanner-/EDR-Technologie.
- Keine Behauptung implementierter/validierter/zertifizierter Self-Protection; keine Security-/Resilience-/Recovery-Reife; keine Eignung für physische Steuerungssysteme/Safety; kein Runtime-Code; keine Health Checks gegen reale Systeme.

## 5. Concepts

`CoreOps control plane` · `managed plane` · `self-protection` · `protective condition` · `protective trigger` · `protection assessment` · `protection action` · `capability` · `critical/privileged capability` · `capability restriction/suspension` · `containment` · `quarantine` · `operational mode` · `fault domain` · `failure domain` · `blast radius` · `dependency` · `critical dependency` · `health signal` · `health assessment` · `liveness` · `readiness` · `governability` · `integrity/trust/authority/evidence state` · `unknown operational state`.

```text
governability = Fähigkeit, Operationen unter gültiger Autorität, Policy, State,
                Evidence und Schutzgrenzen kontrolliert auszuführen
governability ≠ bloße technische Erreichbarkeit
```

## 6. Control-Plane and Managed-Plane Boundary

Der **CoreOps Control Plane** ist die Governance-/Koordinationsgrenze (übergreifend MOD-POL-001/MOD-EXE-001/MOD-WFL-001/MOD-STA-001/MOD-SEC-001/MOD-EVD-001). Getrennt: Agents (MOD-AGT-001, optional) · Adapter (MOD-ADP-001) · verwaltete Systeme · lokale delegierte Ausführung · zentrale Autorität · lokale Authority Snapshots · Evidence Return · State Reconciliation.

```text
control-plane failure  ≠ managed asset failure
managed asset reachable ≠ control plane governable
agent running          ≠ agent authorized for new actions
central unavailable    ≠ local delegation expands
local execution completed ≠ central state reconciled
```

Self-Protection ist **kein** allgemeiner Schutz aller verwalteten Geräte.

## 7. Protected Assets and Self-Protection Scope

Schutzgrenzen (je: accountable owner · authoritative source · required freshness · required trust · failure impact · permitted degraded behavior · prohibited behavior · evidence requirement · recovery requirement · unknown-state behavior) für: Identity State · Authorization State · Policy State · Configuration State · Secret/Key State · Artifact Trust State · Deployment State · Audit/Evidence · Event-/Telemetry-Integrität · Workspace-/Mandantenisolation · Target Binding · Offline Delegation · CorePack State · Recovery State · Reporting-/Exportautorität.

## 8. Protective Conditions and Triggers

Technologieunabhängige Trigger-Kategorien: suspected control-plane compromise · integrity conflict · unauthorized configuration change · policy unavailable/stale · identity authority unavailable · authorization state unknown · secret/key state unavailable · trust/revocation state stale · audit/evidence pipeline unavailable · state-store conflict · partial deployment outcome · update/migration outcome unknown · clock/freshness uncertainty · workspace/target-binding conflict · dependency unavailable · critical resource exhaustion · storage pressure · uncontrolled queue growth · repeated failure loop · reconciliation backlog · offline delegation expiry · unexpected recovery activity · manual protective declaration.

```text
trigger observed      ≠ compromise proven
trigger absent        ≠ platform safe
single signal         ≠ unrestricted automatic shutdown
multiple weak signals ≠ automatically conclusive
unknown assessment    ≠ normal operation permitted
```

Keine konkreten Schwellenwerte definiert.

## 9. Protection Assessment

Ein Assessment erfasst: assessment identity · trigger identity · observation time · source · affected scope · workspace · environment · affected capabilities · confidence · freshness · integrity state · trust state · authority state · evidence availability · possible side effects · recommended restriction · accountable reviewer · review deadline · audit reference · known limitations.

```text
assessment complete ≠ root cause known
high confidence     ≠ human approval unnecessary
low confidence      ≠ unrestricted continuation
```

## 10. Fault Domains and Blast Radius

Fault Domains: platform-wide · workspace · environment · site · target · integration · agent · artifact/deployment · policy · identity · secret/key · evidence.

```text
shared infrastructure ≠ shared authority
shared dependency     ≠ identical impact
fault in one workspace ≠ automatic global shutdown
global signal         ≠ every workspace affected
containment scope must be explicit
```

Unbekannter Scope wird konservativ behandelt, aber **nicht** als nachgewiesener globaler Fehler dargestellt.

## 11. Protective Actions and Priority

Mögliche Maßnahmen: suspend privileged writes · suspend deployment execution · suspend destructive actions · freeze policy changes · freeze identity/authorization changes · restrict secret resolution · restrict export/reporting · isolate workspace/target scope · quarantine artifact/CorePack · require additional human approval · preserve volatile evidence · stop unsafe retries · apply bounded backpressure · pause external integrations · restrict local agents to existing delegation · enter read-only/containment/recovery-only mode · controlled service stop · emergency stop.

```text
protective action ≠ punishment
protective action ≠ root-cause remediation
restriction applied ≠ restriction effective everywhere
controlled stop ≠ evidence preserved automatically
```

Konzeptionelle Prioritätsreihenfolge: **(1)** Schutz von Personen/Umgebung, soweit CoreOps beeinflussbar; **(2)** Verhinderung weiterer privilegierter/destruktiver Aktionen; **(3)** Erhalt von Identity-/State-/Audit-/Evidence-Bindungen; **(4)** Begrenzung des Blast Radius; **(5)** Erhalt zulässiger Read-/Diagnosefähigkeiten; **(6)** kontrollierte Recovery-Vorbereitung. Keine Safety-/Eignungsbehauptung für physische Steuerungssysteme.

## 12. Integration Boundaries

- **Identity/Authorization ([CO-WP-009](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md)/[010](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md)/[013](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md)):** Protective Actions ändern keine Rollen/Autorität still; `machine principal ≠ human protective approval`.
- **Secrets/Keys ([CO-WP-024](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md)):** Restriction von Secret Resolution legt keine Raw Secrets offen; `secret available ≠ secret use authorized`.
- **Artifact/Deployment/CorePack ([CO-WP-021](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)/[022](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md)/[023](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md)):** Quarantine bindet an bestehende Trust-/Revocation-Governance.
- **Audit/Evidence/Telemetry ([CO-WP-018](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)/[019](TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md)):** `missing telemetry ≠ healthy`; `missing audit event ≠ no action occurred`; Evidence-Ausfall kann privilegierte Operationen einschränken.
- **Data Classification ([CO-WP-025](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md)):** Protection-/Recovery-Daten unterliegen Klassifikation/Disclosure.
- **Offline ([CO-WP-023](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md)):** `central unavailable ≠ local authority expands`.
- **Reporting-/Vulnerability-Roadmap:** Self-Protection-Ereignisse sind spätere Report-Inhalte; Vulnerability Findings können Protective Conditions auslösen; `CVE/product-name match ≠ automatic self-protection trigger`; Reporting erzeugt **keine** Recovery-/Exit-Autorität.

## 13. Profiles

`Standard`, `Hardened`, `Government` als Governance-Stärke (zusätzliche Human Approvals · zulässige Degraded Capabilities · Delegationslänge/-Freshness · Evidence-Anforderungen · Recovery-Input-Bewertung · Rollentrennung · Review-Intervalle · Exit-Freigabe · Unknown-State-Verhalten · zulässige Export-/Diagnosefähigkeit).

```text
Government profile ≠ government certification
Hardened profile   ≠ proven resilient
Standard profile   ≠ self-protection optional
```

## 14. Failure and Unknown State

`protective assessment incomplete · trigger source unknown · fault scope undetermined · dependency state unknown · evidence pipeline unavailable · protection action outcome unknown`.

```text
unknown ≠ safe
failure ≠ no side effects
retry   ≠ automatically permitted
monitoring unavailable ≠ system healthy
audit unavailable ≠ no operation occurred
```

## 15. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Control-plane self-protection remains separate from protection of managed assets.
2. Technical reachability does not imply governability; a running process does not prove a healthy, trusted or governable platform.
3. Missing monitoring or audit evidence does not prove a healthy state or absence of operations.
4. Protective conditions and actions do not expand authority.
5. A protective trigger is not proof of compromise; absence of a trigger is not proof of safety.
6. Fault scope and blast radius must be explicit; unknown scope is handled conservatively but not asserted as a proven global fault.
7. Protective actions do not punish, do not remediate root cause, and are not effective everywhere merely by being applied.
8. A controlled stop does not automatically preserve evidence.
9. Unknown operational state is handled fail-closed.
10. Evidence or telemetry unavailability may restrict privileged operations.
11. Reporting and vulnerability findings do not create automatic self-protection, recovery or exit authority.
12. Self-protection must not disable CoreOps' own operating basis without governed authority (self-dependency).

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 16. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): managed resource vs CoreOps THR-034; resource exhaustion/DoS THR-036; partial job failure without safe state THR-031; rollback fails THR-032; backup/restore manipulated THR-033; privilege escalation THR-002; stolen admin identity THR-001; audit deletion/manipulation THR-016/THR-017; manipulated/stale telemetry THR-012/THR-013; false time THR-027; replay THR-026; insider THR-037; automation client THR-038. Keine erfundenen IDs; kein Parallelregister.

## 17. Technology Boundary

Nicht ausgewählt: Health-Monitoring · Watchdog · Heartbeat · Leader Election · Cluster · HA · Failover · Load Balancer · Consensus · Quorum · Distributed Lock · Circuit Breaker · Rate Limiter · Backpressure Engine · Isolation/Sandbox · Network Segmentation · Firewall · Runtime Security · Integrity/Malware Scanner · EDR. Schutzmuster (Capability Restriction, Backpressure, Containment, Circuit-Breaking) sind konzeptionell; keine technische Umsetzung ausgewählt. Alle `deferred`.

## 18. Compatibility

Adressiert [RISK-11](../../project-system/RISK_REGISTER.md) (Self-Dependency); konsistent mit [Restricted/Offline (CO-WP-023)](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md), [Execution Authorization (CO-WP-013)](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [Deployment (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Audit/Evidence (CO-WP-018)](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), MOD-EVD-001 (Protection-Domäne). Additiv; keine bestehende Invariante geschwächt; kein Parallelmodell.

## 19. Open Questions

- Konkrete Trigger-Schwellen und Confidence-Bewertung (deferred).
- Quorum-/Exclusion-Mechanismen gegen Self-Dependency (Konzept §42, deferred; keine Consensus-Technologie).
- Verhältnis Protective Actions zu physischer Safety (bewusst nicht behauptet).

## 20. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): gebündelter `CO-WP-021…026` Foundation Milestone Review (noch nicht terminiert, keine WP-Nummer). Zuerst Nova Review von CO-WP-026, danach Human-Maintainer-Commit.
