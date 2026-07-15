# CoreOps – Key Material, Rotation, Revocation and Recovery Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation key custody, rotation, revocation and recovery policy
> Implementation Status: Not implemented
> KMS Technology: Not selected
> HSM/TPM: Not selected
> PKI/Trust Anchor: Not selected
> Key Algorithm/Format: Not selected
> Raw Key Storage: Not decided
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-024 (docs-only / security, configuration, secret-lifecycle and key-custody governance foundation)

## 1. Status

Technologieunabhängige Policy für **Key Custody, Root-/Bootstrap-Material, Rotation, Revocation, Suspension, Reinstatement, Recovery, Destruction, Break-Glass-Material und Offline-Grenzen**. Companion zu [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) und [CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md](../architecture/CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md). Erweitert [Machine Credential Lifecycle (CO-WP-010)](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md) und [Break-Glass (CO-WP-009)](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) ohne Parallelmodell. Kein KMS/HSM/TPM/PKI, kein Key-Algorithmus/-Format implementiert.

## 2. Purpose

Key Custody trennt Verwahrung von Nutzung. Diese Policy legt fest, warum `key identity ≠ key material`, `key custody ≠ unrestricted key use`, `private material possession ≠ operation authorization`, `key rotation ≠ certificate/credential rollout completion`, `key recovery ≠ reinstatement` und `break-glass available ≠ unrestricted emergency authority` — ohne Algorithmus, Format, Hardware oder PKI auszuwählen.

## 3. Scope

Key Custody · Key Identity/Instances · Bootstrap/Root/Recovery Material · Rotation · Revocation/Suspension/Reinstatement · Recovery · Destruction · Break-Glass Material · Offline/Isolated/Air-Gapped Key Use · Backup/Recovery Boundary · Deployment/Rollback Boundary · Audit (no key material) · Isolation · Failure/Fail-Closed.

## 4. Non-Goals

- Kein KMS/HSM/TPM/PKI/CA/Trust Anchor, kein Key-Algorithmus/-Format, kein Hash-/Encryption-/Signatur-/KDF-Verfahren, kein RNG, keine Envelope Encryption, kein Secret Sharing, keine Threshold-/Quorum-Technologie, keine Backup-Encryption, keine Hardware Root of Trust.
- Keine Behauptung implementierter/validierter/zertifizierter Key-Kontrollen; kein Runtime-Code; keine Compliance-/Einsatzreife-Behauptung; keine Rohschlüsselspeicher-Entscheidung; keine Key-/Secret-Generierung.

## 5. Concepts

`key identity` · `key purpose` · `key classification` · `key version` · `key material` · `public key material` · `private/restricted key material` · `key custody instance` · `key material state` · `certificate reference` · `bootstrap material` · `root material` · `recovery material` · `emergency/break-glass material` · `rotation event` · `revocation event` · `suspension` · `reinstatement` · `recovery instance` · `destruction event`.

```text
key identity  ≠ key material
key version   ≠ key custody instance
key custody   ≠ unrestricted key use
```

## 6. Key Custody

Mindestattribute je Key: key identity · purpose · classification · version · key material state · owner · custodian · permitted operations · permitted consumers · environment/target binding · activation boundary · rotation state · revocation state · recovery state · destruction state · export prohibition/authority · audit reference · unknown-state handling.

```text
key identity              ≠ key material
public material           ≠ unrestricted trust
private material possession ≠ operation authorization
key custody               ≠ policy authority
key recovery              ≠ reinstatement
key rotation              ≠ certificate or credential rollout completion
```

Keine Algorithmen, Formate, Hardware oder PKI definiert.

## 7. Key Identity and Instances

Getrennt: stable key identity · key version · key material state · key custody instance · rotation event · revocation event · recovery instance · destruction event. `same key material ≠ same custody state`; `same key version ≠ same distribution history`; `same reference ≠ same resolved key`. Mutable Aliases sind keine finale privilegierte Bindung; privilegierte Operationen zeigen auf eine konkrete, bewertete Key-Version.

## 8. Bootstrap, Root and Recovery Material

Besondere Governance für: bootstrap secret · initial enrollment material · root material · recovery material · emergency material · break-glass material · offline bootstrap snapshot. Mindestens: accountable human owner · creation/receipt evidence · scope · one-time/bounded-use expectation · target binding · validity · custody · access approval · use evidence · post-use review · mandatory rotation/invalidation · recovery limitations · destruction/retirement expectation.

```text
bootstrap completed  ≠ bootstrap material reusable
break-glass used     ≠ normal authority restored
recovery succeeded   ≠ prior compromised material trustworthy
root material available ≠ unrestricted authority
```

Bezug [Machine Enrollment (CO-WP-010)](MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE.md), THR-001.

## 9. Rotation

Auslöser: scheduled · risk-triggered · compromise-triggered · dependency-triggered · owner-triggered · policy-triggered · emergency · offline-delayed. Rotation berücksichtigt: old version · new version · affected consumers · compatibility · distribution · activation sequence · overlap boundary · rollback feasibility · revocation timing · unknown consumers · offline consumers · evidence · completion criteria.

```text
new version created ≠ rotation complete
new version available ≠ all consumers updated
old version revoked ≠ all old uses stopped
partial rollout     ≠ successful rotation
unknown consumer state ≠ safe completion
rollback to old version ≠ automatically permitted
```

Overlap ist begrenzt/dokumentiert; kein unbegrenztes Overlap ([CO-WP-010 §11](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md)). Keine Rotation-Engine/-Intervall ausgewählt.

## 10. Revocation, Suspension and Reinstatement

Getrennt: secret revocation · credential revocation · key revocation · use prohibition · distribution prohibition · export prohibition · local emergency suspension · central revocation · delayed offline revocation. Revocation besitzt Identity Binding und Freshness.

```text
no local revocation entry ≠ not revoked centrally
revocation issued         ≠ observed by every consumer
suspension                ≠ reinstatement
reinstatement             ≠ historical revocation erased
reinstatement             ≠ use authorization
```

Reinstatement benötigt: verantwortliche Authority · Grund · neue Evidenz · Scope · Target · Consumer · aktualisierte Version/Bewertung · Gültigkeit · Review Trigger · Audit Reference (Bezug THR-026, THR-027).

## 11. Recovery

Recovery benötigt: Recovery Authority · Recovery Scope · Target Binding · Version Binding · Custody · Review · Post-Recovery Rotation · Revocation Reassessment · Evidence · Unknown Outcome.

```text
restore completed ≠ recovered secret valid
recovered value   ≠ reinstated value
historical secret ≠ safe rollback secret
```

`backup operator ≠ recovery authority`. Recovery impliziert keine Reinstatement- oder globale Trust-Entscheidung; wiederhergestelltes Material wird nach Recovery neu bewertet und ggf. rotiert (Bezug THR-032, THR-033).

## 12. Destruction

Destruction stoppt jede Nutzung, entfernt aktive Verwendbarkeit und erhält historische Audit-/Evidence-Referenzen (`destroyed ≠ historical evidence erased`). Destruction ist scope-/version-gebunden, approbiert und auditiert; Destruction-Outcome kann `unknown` sein → Reconciliation. Destroyed Key-/Secret-Identitäten werden nicht still wiederverwendet.

## 13. Break-Glass Material

Break-Glass-**Material** (emergency keys/credentials) folgt der [Break-Glass Policy (CO-WP-009)](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md): exceptional · temporary · named · scope-/time-bound · auditiert · mandatory expiry. Zusätzlich hier: bounded/one-time-use, verpflichtende Rotation/Invalidation nach Nutzung, post-use review. `break-glass available ≠ unrestricted emergency authority`; `break-glass used ≠ normal authority restored`; **keine dauerhafte Autorität**. Kein Parallel-Break-Glass-Prozess.

## 14. Offline, Isolated and Air-Gapped Key Use

Bindet an [CO-WP-023](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md): offline key/secret subset · local custody boundary · local policy/identity/role snapshot · revocation snapshot · freshness boundary · clock uncertainty · delayed central changes · evidence return · reconciliation.

```text
central unavailable    ≠ local authority expands
offline cache available ≠ use currently authorized
local key copy         ≠ unrestricted export
no local revocation entry ≠ centrally valid
local clock recent     ≠ central policy current
air-gapped classification ≠ security validation
```

Offline Authorizations sind action-, secret-, key-, target-, consumer-, workspace-, purpose-, version- und zeitgebunden und nicht wiederverwendbar. Keine Offline-Vault-/Synchronisationstechnologie.

## 15. Backup and Recovery Boundary

Getrennt: ordinary backup · configuration backup · secret backup · key recovery material · vault recovery state · evidence package.

```text
backup contains secret material ≠ ordinary backup
backup available               ≠ recovery authorized
restore completed              ≠ recovered secret valid
recovered value                ≠ reinstated value
historical secret              ≠ safe rollback secret
```

Secret-/Key-haltige Backups unterliegen gesonderter Custody/Export-Grenze ([Companion Governance §12](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md); Bezug THR-040). Keine Backup-/Encryption-/Recovery-Technologie ausgewählt.

## 16. Deployment and Rollback Boundary

Bindet an [CO-WP-021](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)/[CO-WP-022](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md):

```text
rollback artifact available ≠ rollback credential valid
SBOM available              ≠ secret inventory complete
artifact provenance         ≠ credential provenance
deployment execution authorization ≠ secret export authorization
```

Rollback bewertet Secret-/Credential-/Key-Zustände neu; ein Rollback reaktiviert **keine** widerrufene, abgelaufene oder kompromittierte Secret-/Key-Version.

## 17. Audit and Evidence

Audit/Evidence enthalten **kein** Key-/Secret-Material — nur sichere Metadaten (key/secret identity · version · request/requester · consumer · target · purpose · policy decision · approval · authorization · event type · outcome · freshness · audit reference). `access logged ≠ access authorized`; `missing audit event ≠ no key use occurred`. Records gehören MOD-EVD-001; Historie nicht umgeschrieben (Bezug THR-016, THR-017, THR-019).

## 18. Workspace, Environment and Target Isolation

Key-/Custody-Bindungen bleiben durch Custody, Rotation, Revocation, Recovery und Offline Use erhalten. Ein Key aus Workspace/Environment A wird nicht still in B verwendet; Environment Binding nicht durch Rename ersetzbar; geteilte Infrastruktur erzeugt keine gemeinsame Key Authority; `same key material ≠ same authorization` (Bezug THR-035).

## 19. Failure and Unknown State

`key material unavailable · custody state unknown · rotation partially completed · old key still in use · revocation delivery unknown · recovery outcome unknown · destruction outcome unknown · offline freshness unknown · policy conflict · target/workspace mismatch · unauthorized export suspected`.

```text
unknown          ≠ safe
failure          ≠ no side effects
retry            ≠ automatically permitted
missing evidence ≠ operation did not occur
```

## 20. Fail-Closed Rules

Keine privilegierte Key-Operation bei: unbekanntem/unaufgelöstem Key-Identity/Version · fehlender Operation-Authorization für Ziel/Consumer/Scope · widerrufener/abgelaufener/suspendierter Key-Version · unklarer Custody · zu alter/unklarer Revocation-Freshness ohne Exception · Target-/Workspace-Mismatch · Alias-only-Binding · vermutetem unautorisiertem Export. Recovery/Bootstrap/Break-Glass nur mit expliziter, benannter, zeitgebundener Autorität und mandatory Post-Use-Rotation.

## 21. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Key identity remains separate from key material and custody; custody does not imply use authority.
2. Public key material does not imply unrestricted trust; private material possession does not imply operation authorization.
3. Rotation is complete only after affected/unknown/offline consumer states are assessed.
4. Revocation possesses identity binding and freshness; a missing local revocation entry does not prove central non-revocation.
5. Suspension, reinstatement and recovery remain distinct; reinstatement does not erase history or imply use authorization.
6. Recovery does not imply reinstatement or a global trust decision; recovered material is reassessed and rotated.
7. Rollback does not reactivate revoked, expired or compromised secret or key material.
8. Bootstrap, root, recovery and break-glass material are bounded/one-time and require mandatory rotation/invalidation; they create no permanent authority.
9. Offline operation does not expand authority; offline key authorizations are bounded and non-reusable.
10. Destruction stops use and preserves historical evidence; destroyed identities are not silently reused.
11. Unknown rotation/revocation/recovery/destruction outcomes remain visible and block unsafe retry.
12. Audit and evidence contain no key or secret material.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 22. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): Secret/Key Exposure THR-019/THR-020/THR-040; offline export sensitive THR-025; offline import THR-024; replay THR-026; false time THR-027; reuse expired approval THR-004; stolen admin identity THR-001; privilege escalation THR-002; rollback fails THR-032; backup/restore manipulated THR-033; audit deletion/manipulation THR-016/THR-017; tenant boundary THR-035; insider THR-037; automation client THR-038. Keine erfundenen IDs; kein Parallelregister.

## 23. Technology Boundary

Nicht ausgewählt: KMS · HSM · TPM · PKI · CA · Trust Anchor · Key-Algorithmus/-Format · Hash-/Encryption-/Signatur-/KDF-Verfahren · RNG · Envelope Encryption · Secret Sharing · Threshold/Quorum · Rotation-Engine · Certificate Manager · Backup-Encryption · Hardware Root of Trust. Alle bleiben `deferred`; Rohschlüsselspeicherung nicht entschieden.

## 24. Compatibility

Konsistent mit [Companion Governance](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md), [Machine Credential Lifecycle (CO-WP-010)](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md), [Break-Glass (CO-WP-009)](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md), [Deployment (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Artifact Trust (CO-WP-022)](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md), [Offline/CorePack (CO-WP-023)](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md). Additiv; keine bestehende Invariante geschwächt.

## 25. Open Questions

- Ob CoreOps Rohschlüsselmaterial speichert (spätere Sicherheitsarchitektur-Entscheidung).
- Empfohlene Rotationsintervalle/Overlap-Fenster je Key-Klasse (deferred).
- Zuverlässige Offline-Revocation-Distribution für Keys (deferred).

## 26. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-025 – Data Classification, Retention and Redaction`. Zuerst Nova Review von CO-WP-024, danach Human-Maintainer-Commit.
