# CoreOps – Degraded Mode and Capability Restriction Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation operational-mode and capability-restriction model
> Implementation Status: Not implemented
> Mode Engine: Not selected
> Policy Engine: Not selected
> Circuit-Breaker/Backpressure Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-026 (docs-only / control-plane self-protection, degraded-operation and governed-recovery foundation)

## 1. Status

Technologieunabhängiges Modell für **Operational Modes, Capability Restriction Matrix, Read-only/Restricted/Degraded/Containment/Emergency-Stop und Mode Entry/Exit**. Companion zu [SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md](../security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md) und [RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md). Erweitert die Operational States aus [Restricted Operation (CO-WP-023) §14](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md) zu einem vollständigen Modell; DEC-S-316 (degraded = named/bounded/reviewable) bleibt gültig und wird nicht dupliziert. Keine Mode-/Policy-/Circuit-Breaker-Technologie implementiert.

## 2. Purpose

Operational Modes sind explizit und capability-begrenzt. Das Modell legt fest, warum `degraded mode ≠ unrestricted emergency mode`, `degraded mode ≠ permanent operating profile`, `read-only mode ≠ guaranteed absence of side effects`, `restricted mode ≠ read-only mode`, `containment ≠ recovery` und `unknown state ≠ normal or safe`.

## 3. Scope

Operational Modes · Mode Entry/Exit · Capability Restriction Matrix · Read-only Mode · Restricted/Guarded Mode · Degraded Mode · Containment · Emergency Stop · Unknown Operational State · Profile · Failure. Self-Protection-Trigger im [Companion Safety Model](../security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md); Recovery im [Companion Recovery Policy](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md).

## 4. Non-Goals

- Keine Mode-/State-Machine-/Policy-Engine, kein Circuit-Breaker/Rate-Limiter/Backpressure-Produkt, kein Scheduler/Supervisor/Service-Manager, keine HTTP-/API-/DB-Semantik.
- Keine Behauptung implementierter Modi/Kontrollen; kein Runtime-Code; keine Security-/Resilience-Reife.

## 5. Concepts

`operational mode` · `normal mode` · `guarded mode` · `restricted mode` · `read-only mode` · `degraded mode` · `containment mode` · `recovery-only mode` · `recovery mode` · `emergency-stop` · `unknown operational state` · `capability` · `critical/privileged capability` · `capability restriction/suspension` · `mode entry/exit condition` · `write behavior` · `side effect`.

## 6. Operational Modes

Zehn Modi, je: purpose · entry conditions · accountable authority · permitted capabilities · prohibited capabilities · write behavior · deployment behavior · secret-use behavior · export behavior · agent behavior · offline behavior · evidence requirements · validity boundary · review trigger · exit conditions · known limitations.

| Mode | Kurzcharakter |
| ---- | ------------- |
| `normal` | regulärer governter Betrieb (≠ proven secure) |
| `guarded` | erhöhte Wachsamkeit, zusätzliche Approvals |
| `restricted` | verengte Fähigkeiten bei eingeschränkter Autorität/Freshness |
| `read-only` | keine autorisierten Writes (≠ garantiert side-effect-free) |
| `degraded` | begrenzte, benannte, reviewpflichtige Einschränkung |
| `containment` | Isolierung eines Scopes (≠ recovery) |
| `recovery-only` | nur Recovery-Vorbereitung/-Fähigkeiten |
| `recovery` | governter temporärer Wiederherstellungsbetrieb ([Companion Recovery](../governance/RECOVERY_MODE_AUTHORITY_AND_CONTROLLED_RESTORATION_POLICY.md)) |
| `emergency-stop` | kontrollierter Stopp privilegierter Aktionen (≠ permanent shutdown) |
| `unknown operational state` | Mode nicht bestimmt → fail-closed |

```text
normal mode ≠ proven secure
guarded mode ≠ degraded mode
restricted mode ≠ read-only mode
read-only mode ≠ no side effects
degraded mode ≠ normal operation with warnings
containment ≠ recovery
recovery-only ≠ unrestricted recovery
recovery mode ≠ normal mode
emergency-stop ≠ permanent shutdown
unknown state ≠ normal or safe
```

Konsistent mit den Operational States aus [CO-WP-023 §14](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md).

## 7. Mode Entry and Exit

Mode-Übergänge sind explizit, menschlich zurechenbar, scope-/zeitgebunden und auditiert. Entry benötigt Trigger/Assessment ([Safety Model §8–§9](../security/SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md)); Exit benötigt Exit-Bedingungen und Review. `mode entered ≠ root cause resolved`; `system still running ≠ acceptable degraded operation`. Kein stiller Mode-Wechsel; Unknown → fail-closed in den restriktiveren Mode.

## 8. Capability Restriction Matrix

Capability-Gruppen: read inventory · read monitoring data · read audit/evidence · acknowledge alerts · modify configuration · modify policy · manage identities · manage authorization · resolve secrets · export data · generate reports · execute scripts · perform deployments · perform updates · perform destructive operations · perform backup · perform restore · perform rollback · import CorePacks · activate CorePacks · reconcile state · enter recovery · exit recovery.

Je Mode-/Capability-Kombination konzeptioneller Status: `permitted` · `permitted with additional approval` · `permitted only for bounded recovery` · `suspended` · `prohibited` · `unknown/fail-closed`. Grundmuster: privilegierte/destruktive Capabilities (modify policy/manage identities/deployments/updates/destructive/restore/rollback/resolve secrets/export) werden in restricted/read-only/degraded/containment `suspended`/`prohibited` oder nur `permitted with additional approval`; Read-/Diagnose-Capabilities bleiben nach Klassifikation zulässig; `enter/exit recovery` nur mit Recovery Authority. Keine technische Policy Engine ausgewählt.

## 9. Read-only Mode

Definiert ausdrücklich: welche Reads zulässig sind · welche Reads Disclosure-/Secret-Risiken tragen · welche Hintergrundaktionen weiter Side Effects erzeugen könnten · Behandlung von Caches/Leases/Sessions/lokalen Agent-Aktionen · Evidence-Erzeugung.

```text
read-only UI          ≠ read-only platform
GET-like operation    ≠ side-effect free
diagnostic read       ≠ export authorized
read-only mode        ≠ secret disclosure permitted
```

Keine konkrete HTTP-/API-/DB-Semantik vorgeschrieben.

## 10. Restricted and Guarded Mode

`guarded`: normaler Betrieb mit erhöhter Wachsamkeit und zusätzlichen Approvals für privilegierte Aktionen. `restricted`: bewusst verengte Fähigkeiten bei eingeschränkter Autorität/Freshness (read-only bevorzugt vor write; keine privilegierte Aktivierung ohne aktuelle Autorisierung; keine stille Scope-Erweiterung — konsistent mit [CO-WP-023 §13](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md)). `restricted ≠ read-only`; `guarded ≠ degraded`.

## 11. Degraded Mode

Ein Degraded-Mode-Eintrag: mode identity · reason · trigger · affected scope · accountable human owner · allowed/prohibited capabilities · local/central authority state · policy version · identity snapshot · trust/revocation snapshot · evidence state · validity · review interval · exit condition · audit reference · known limitations.

```text
degraded operation ≠ authority expansion
degraded operation ≠ indefinite exception
degraded operation ≠ suppression of unknown state
system still running ≠ acceptable degraded operation
```

Jeder Degraded Mode ist benannt, begrenzt, zeitlich überprüfbar, personell zurechenbar, widerrufbar, auditierbar und mit Exit-Bedingung versehen (DEC-S-316; hier für Self-Protection konkretisiert).

## 12. Containment

Containment: containment scope · authority · affected capabilities · isolation boundary · evidence preservation · communication boundary · local-agent behavior · secret/credential behavior · export prohibition · review · escalation · release conditions.

```text
containment applied  ≠ compromise removed
quarantine release   ≠ normal operation authorized
containment release  ≠ recovery exit
containment evidence ≠ root cause proven
```

Containment ist getrennt von Recovery.

## 13. Emergency Stop

`emergency-stop` stoppt kontrolliert privilegierte/destruktive Aktionen; `emergency-stop ≠ permanent shutdown`; `controlled stop ≠ evidence preserved automatically`. Wiederaufnahme benötigt Assessment und explizite Mode-Wahl (nicht automatisch `normal`).

## 14. Unknown Operational State

Ist der Mode nicht eindeutig bestimmbar, gilt `unknown operational state` → fail-closed in den restriktiveren Mode; privilegierte/destruktive Capabilities `suspended`; Read-/Diagnose nach Klassifikation; explizite Assessment-/Reconciliation-Pflicht. `unknown ≠ normal or safe`.

## 15. Profiles

`Standard`/`Hardened`/`Government` variieren zulässige Degraded Capabilities, zusätzliche Approvals, Review-Intervalle, Exit-Freigabe und Unknown-State-Verhalten. `Government profile ≠ government certification`; `Hardened ≠ proven resilient`; `Standard ≠ self-protection optional`.

## 16. Failure and Unknown State

`mode assessment incomplete · conflicting mode signals · exit condition unverifiable · capability status unknown`.

```text
unknown ≠ safe/normal
failure ≠ no side effects
retry   ≠ automatically permitted
```

## 17. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Operational modes remain explicit, accountable and capability-bounded.
2. Read-only mode does not automatically guarantee absence of side effects and does not permit secret disclosure.
3. Restricted, guarded, read-only, degraded and containment modes remain distinct.
4. Degraded modes are temporary, bounded, reviewable, attributable, revocable and carry an exit condition (DEC-S-316).
5. Containment remains separate from recovery; containment release is not normal-operation authorization.
6. Emergency stop is not a permanent shutdown and does not automatically preserve evidence.
7. Unknown operational state is handled fail-closed toward the more restrictive mode.
8. Mode entry and exit are explicit, human-attributable, scope- and time-bound, and audited.
9. Protective conditions and mode changes do not expand authority.
10. Privileged and destructive capabilities are suspended or prohibited unless explicitly permitted for the current mode and authority.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 18. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): managed resource vs CoreOps THR-034; resource exhaustion/DoS THR-036; partial failure without safe state THR-031; privilege escalation THR-002; stolen admin identity THR-001; reuse expired approval THR-004; audit deletion/manipulation THR-016/THR-017; stale telemetry THR-013; automation client THR-038. Keine erfundenen IDs; kein Parallelregister.

## 19. Technology Boundary

Nicht ausgewählt: Mode-/State-Machine-Framework · Policy Engine · Circuit Breaker · Rate Limiter · Backpressure Engine · Scheduler · Supervisor · Service Manager · Isolation/Sandbox. Alle `deferred`.

## 20. Compatibility

Erweitert die Operational States aus [CO-WP-023 §14](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md); konsistent mit [Execution Authorization (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [Break-Glass (CO-WP-009)](../security/BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md). DEC-S-316 nicht dupliziert. Additiv; keine bestehende Invariante geschwächt.

## 21. Open Questions

- Konkrete Mode-/Capability-Matrix-Ausprägung je Profil (deferred).
- Behandlung laufender langlebiger Aktionen bei Emergency Stop (deferred; keine Scheduler-Auswahl).

## 22. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): gebündelter `CO-WP-021…026` Foundation Milestone Review (noch nicht terminiert). Zuerst Nova Review von CO-WP-026, danach Human-Maintainer-Commit.
